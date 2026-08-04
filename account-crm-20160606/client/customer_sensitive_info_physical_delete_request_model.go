// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCustomerSensitiveInfoPhysicalDeleteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUid(v string) *CustomerSensitiveInfoPhysicalDeleteRequest
	GetUid() *string
}

type CustomerSensitiveInfoPhysicalDeleteRequest struct {
	// This parameter is required.
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s CustomerSensitiveInfoPhysicalDeleteRequest) String() string {
	return dara.Prettify(s)
}

func (s CustomerSensitiveInfoPhysicalDeleteRequest) GoString() string {
	return s.String()
}

func (s *CustomerSensitiveInfoPhysicalDeleteRequest) GetUid() *string {
	return s.Uid
}

func (s *CustomerSensitiveInfoPhysicalDeleteRequest) SetUid(v string) *CustomerSensitiveInfoPhysicalDeleteRequest {
	s.Uid = &v
	return s
}

func (s *CustomerSensitiveInfoPhysicalDeleteRequest) Validate() error {
	return dara.Validate(s)
}
