// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTairSkvDdbTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeTairSkvDdbTableRequest
	GetInstanceId() *string
	SetSecurityToken(v string) *DescribeTairSkvDdbTableRequest
	GetSecurityToken() *string
}

type DescribeTairSkvDdbTableRequest struct {
	// The ID of the instance to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// tt-bp1zxszhcgatnx\\*\\*\\*\\*
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// ****
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeTairSkvDdbTableRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTairSkvDdbTableRequest) GoString() string {
	return s.String()
}

func (s *DescribeTairSkvDdbTableRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeTairSkvDdbTableRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeTairSkvDdbTableRequest) SetInstanceId(v string) *DescribeTairSkvDdbTableRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeTairSkvDdbTableRequest) SetSecurityToken(v string) *DescribeTairSkvDdbTableRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeTairSkvDdbTableRequest) Validate() error {
	return dara.Validate(s)
}
