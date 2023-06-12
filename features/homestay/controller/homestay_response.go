package controller

import (
	models "be-api/features"
)

type HomestayResponse struct {
	HostID      uint             `json:"host_id,omitempty"`
	Title       string           `json:"title,omitempty"`
	Description string           `json:"description,omitempty"`
	Location    string           `json:"location,omitempty"`
	Price       float64          `json:"price,omitempty"`
	Facilities  string           `json:"facilities,omitempty"`
	Images      string           `json:"image_links,omitempty"`
	Reviews     []ReviewResponse `json:"reviews,omitempty"`
}

type ReviewResponse struct {
	Ratings uint `json:"ratings,omitempty"`
}

func HomestayEntityToResponse(homestay models.HomestayEntity) HomestayResponse {
	var reviews []ReviewResponse
	for _, review := range homestay.Reviews {
		reviews = append(reviews, ReviewEntityToResponse(review))
	}

	return HomestayResponse{
		HostID:      homestay.HostID,
		Title:       homestay.Title,
		Description: homestay.Description,
		Location:    homestay.Location,
		Price:       homestay.Price,
		Facilities:  homestay.Facilities,
		Reviews:     reviews,
	}
}

func ReviewEntityToResponse(review models.ReviewEntity) ReviewResponse {
	return ReviewResponse{
		Ratings: review.Ratings,
	}
}
