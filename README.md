# Uyga Vazifa: Unary gRPC Xizmatini Implementatsiya Qilish

#### Maqsad:
Ushbu vazifaning maqsadi `Go` tilida `unary gRPC` xizmatini implementatsiya qilish va uni unit testidan o'tkazishdir. 

#### Talablar:

1. **gRPC Serviceni Ta'riflang**:
   - Kitob ma'lumotlarini boshqarish uchun `book.proto` faylini yarating.
   - Service quyidagi unary metodlarni o'z ichiga olishi kerak: Kitob haqida ma'lumot olish (`GetBookInfo`) va yangi kitob qo'shish (`AddBook`).

2. **gRPC Service ni Implementatsiya Qiling**:
   - Serviceni implementatsiya qiling, kitob haqida ma'lumotlarni olish va yangi kitoblarni qo'shish funksiyalarini taqdim eting.

3. **Ma'lumotlarni Saqlash**:
   - Kitob ma'lumotlarini ma'lumotlar bazasi yoki boshqa saqlash vositasida saqlang. Ma'lumotlarni saqlash uchun mos API yoki kutubxonani tanlang.

3. **Unit Test Yozing**:
   - Har bir unary method uchun unit test yozing.


### Message Types
`BookRequest`
- Fields:
    - `string title`: Ma'lumot so'raladigan kitobning nomi.

`BookResponse`
- Fields:
    - `string title`: Kitobning nomi.
    - `string author`: Kitob muallifi.
    - `string summary`: Kitobning qisqacha mazmuni.
    - `string timestamp`: Javobning vaqt belgisi.

`Book`
- Fields:
    - `string title`: Kitobning nomi.
    - `string author`:  Kitob muallifi.
    - `string summary`: Kitobning qisqacha mazmuni.





