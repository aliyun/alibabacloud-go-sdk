// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCustomerLabelByConfigGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupType(v string) *QueryCustomerLabelByConfigGroupRequest
	GetGroupType() *string
	SetPK(v int64) *QueryCustomerLabelByConfigGroupRequest
	GetPK() *int64
	SetToken(v string) *QueryCustomerLabelByConfigGroupRequest
	GetToken() *string
}

type QueryCustomerLabelByConfigGroupRequest struct {
	// This parameter is required.
	GroupType *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	// This parameter is required.
	PK *int64 `json:"PK,omitempty" xml:"PK,omitempty"`
	// This parameter is required.
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s QueryCustomerLabelByConfigGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryCustomerLabelByConfigGroupRequest) GoString() string {
	return s.String()
}

func (s *QueryCustomerLabelByConfigGroupRequest) GetGroupType() *string {
	return s.GroupType
}

func (s *QueryCustomerLabelByConfigGroupRequest) GetPK() *int64 {
	return s.PK
}

func (s *QueryCustomerLabelByConfigGroupRequest) GetToken() *string {
	return s.Token
}

func (s *QueryCustomerLabelByConfigGroupRequest) SetGroupType(v string) *QueryCustomerLabelByConfigGroupRequest {
	s.GroupType = &v
	return s
}

func (s *QueryCustomerLabelByConfigGroupRequest) SetPK(v int64) *QueryCustomerLabelByConfigGroupRequest {
	s.PK = &v
	return s
}

func (s *QueryCustomerLabelByConfigGroupRequest) SetToken(v string) *QueryCustomerLabelByConfigGroupRequest {
	s.Token = &v
	return s
}

func (s *QueryCustomerLabelByConfigGroupRequest) Validate() error {
	return dara.Validate(s)
}
