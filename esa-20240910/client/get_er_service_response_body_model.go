// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetErServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPlanName(v string) *GetErServiceResponseBody
	GetPlanName() *string
	SetRequestId(v string) *GetErServiceResponseBody
	GetRequestId() *string
	SetStatus(v string) *GetErServiceResponseBody
	GetStatus() *string
}

type GetErServiceResponseBody struct {
	// The billing mode. Valid values:
	//
	// 	- er_paymode: billed for customers on the China site.
	//
	// 	- er_freemode: free for customers on the China site.
	//
	// 	- er_paymodeintl: billed for customers on the International site.
	//
	// 	- err_freemodeintl: free for customers on the International site
	//
	// example:
	//
	// er_paymode
	PlanName *string `json:"PlanName,omitempty" xml:"PlanName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CB1A380B-09F0-41BB-A198-72F8FD6DA2FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The service status. Valid values:
	//
	// 	- Creating
	//
	// 	- Running
	//
	// 	- NotOpened
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetErServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetErServiceResponseBody) GoString() string {
	return s.String()
}

func (s *GetErServiceResponseBody) GetPlanName() *string {
	return s.PlanName
}

func (s *GetErServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetErServiceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetErServiceResponseBody) SetPlanName(v string) *GetErServiceResponseBody {
	s.PlanName = &v
	return s
}

func (s *GetErServiceResponseBody) SetRequestId(v string) *GetErServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetErServiceResponseBody) SetStatus(v string) *GetErServiceResponseBody {
	s.Status = &v
	return s
}

func (s *GetErServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
