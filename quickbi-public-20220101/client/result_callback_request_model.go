// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResultCallbackRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *ResultCallbackRequest
	GetApplicationId() *string
	SetHandleReason(v string) *ResultCallbackRequest
	GetHandleReason() *string
	SetStatus(v int32) *ResultCallbackRequest
	GetStatus() *int32
}

type ResultCallbackRequest struct {
	// The ID of the approval process.
	//
	// This parameter is required.
	//
	// example:
	//
	// c5ea0db8-****-****-9081-04bc0df4c6a3
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The reason for the approval.
	//
	// This parameter is required.
	//
	// example:
	//
	// You are not a Division A analyst.
	HandleReason *string `json:"HandleReason,omitempty" xml:"HandleReason,omitempty"`
	// Approval result:
	//
	// 	- 1: passed
	//
	// 	- 2: rejected
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ResultCallbackRequest) String() string {
	return dara.Prettify(s)
}

func (s ResultCallbackRequest) GoString() string {
	return s.String()
}

func (s *ResultCallbackRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ResultCallbackRequest) GetHandleReason() *string {
	return s.HandleReason
}

func (s *ResultCallbackRequest) GetStatus() *int32 {
	return s.Status
}

func (s *ResultCallbackRequest) SetApplicationId(v string) *ResultCallbackRequest {
	s.ApplicationId = &v
	return s
}

func (s *ResultCallbackRequest) SetHandleReason(v string) *ResultCallbackRequest {
	s.HandleReason = &v
	return s
}

func (s *ResultCallbackRequest) SetStatus(v int32) *ResultCallbackRequest {
	s.Status = &v
	return s
}

func (s *ResultCallbackRequest) Validate() error {
	return dara.Validate(s)
}
