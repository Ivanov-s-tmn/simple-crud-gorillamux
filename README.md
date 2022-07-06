Simple Go CRUD API with gorilla/mux

API endpoints:

    1. GET /cars --> getAllCars func 
    2. GET /cars/{id} --> getCarById func
    3. GET /cars/make/{make} --> getCarsByBrand func
    4. PUT /cars/{id} --> updateCar func
    5. POST /cars --> createCar func
    6. DELETE /cars/{id} --> deleteCar func