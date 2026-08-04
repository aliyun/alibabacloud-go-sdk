// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCustomerSensitiveInfoLogicalDeleteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUid(v string) *CustomerSensitiveInfoLogicalDeleteRequest
	GetUid() *string
}

type CustomerSensitiveInfoLogicalDeleteRequest struct {
	// This parameter is required.
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s CustomerSensitiveInfoLogicalDeleteRequest) String() string {
	return dara.Prettify(s)
}

func (s CustomerSensitiveInfoLogicalDeleteRequest) GoString() string {
	return s.String()
}

func (s *CustomerSensitiveInfoLogicalDeleteRequest) GetUid() *string {
	return s.Uid
}

func (s *CustomerSensitiveInfoLogicalDeleteRequest) SetUid(v string) *CustomerSensitiveInfoLogicalDeleteRequest {
	s.Uid = &v
	return s
}

func (s *CustomerSensitiveInfoLogicalDeleteRequest) Validate() error {
	return dara.Validate(s)
}
