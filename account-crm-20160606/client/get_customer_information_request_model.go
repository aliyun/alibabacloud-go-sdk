// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomerInformationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserId(v int64) *GetCustomerInformationRequest
	GetUserId() *int64
}

type GetCustomerInformationRequest struct {
	// This parameter is required.
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetCustomerInformationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCustomerInformationRequest) GoString() string {
	return s.String()
}

func (s *GetCustomerInformationRequest) GetUserId() *int64 {
	return s.UserId
}

func (s *GetCustomerInformationRequest) SetUserId(v int64) *GetCustomerInformationRequest {
	s.UserId = &v
	return s
}

func (s *GetCustomerInformationRequest) Validate() error {
	return dara.Validate(s)
}
