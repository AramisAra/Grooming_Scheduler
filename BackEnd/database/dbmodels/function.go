package dbmodels

// Remove the import cycle between utilsmodels and dbmodels

func CreateJoinResultClient(client Client, animal Animal) joinResultClient {
	return joinResultClient{Client: client, Animal: animal}
}

func CreateClientResponse(client Client) ClientSerializer {
	return ClientSerializer{ID: client.ID, FullName: client.Full_Name, Email: client.Email, Phone: client.Phone,
		Animals: client.Animals, Appointments: client.Appointments}
}

func CreateServiceResponse(service Service) ServiceSerializer {
	return ServiceSerializer{ID: service.ID, ServiceName: service.NameService, ServiceDesc: service.ServiceDesc, Price: service.Price}
}

func CreateOwnerResponse(owner Owner) OwnerSerializer {
	return OwnerSerializer{ID: owner.ID, Full_Name: owner.Full_Name, Phone: owner.Phone, Email: owner.Email,
		Career: owner.Career, Services: owner.Services, Appointments: owner.Appointments}
}

func CreateAnimalResponse(animal Animal) AnimalSerializer {
	return AnimalSerializer{ID: animal.ID, Animal_Name: animal.Animal_Name, Animal_Specie: animal.Animal_Specie, Animal_Age: animal.Animal_Age, Client: animal.Client_id}
}

func CreateAppointmentResponse(appointment Appointment) AppointmentSerializer {
	return AppointmentSerializer{ClientID: appointment.ClientID, OwnerID: appointment.OwnerID, ServiceID: appointment.ServiceID,
		Date: appointment.Date, Time: appointment.Time}
}
