# Worker-pool

Задание для VK стажировки 

В коде реализован примитивный worker-pool с возможностью динамически добавлять и удалять воркеры. Воркеры обрабатывают задания, поступающие через канал.
# Стек
Go 

# Описание работы кода:
1. Создается канал jobChan для заданий и список воркеров workers.
2. Процесс создания воркеров:
  Функция addWorker() создает и запускает нового воркера, добавляя его в пул.
3. Обработка заданий:
   ~ Воркеры постоянно ждут задание из jobChan;
   ~ когда задание получено -> воркер выводит свой ID и содержание задания. 
4. Добавление воркера: после запуска основной программы, через некоторое время добавляется новый воркер.
5. Удаление воркера: Через некоторое время после добавления, один из воркеров останавливается с помощью Stop().
6. Задания генерируются в отдельной горутине и отправляются в jobChan.
7. После отправки всех заданий -> jobChan закрывается (это сигнал воркерам о завершении, новые задания не поступают).
8. Воркеры завершают работу (после 7 пункта, когда обнаруживают, что канал закрыт).
9. Используется sync.WaitGroup для ожидания завершения всех воркеров перед выходом из программы.

# Пример работы кода:
![image](https://github.com/user-attachments/assets/950604ff-dd10-4793-a0b0-a00dffa62e24)
