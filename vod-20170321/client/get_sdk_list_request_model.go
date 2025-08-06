// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSdkListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroup(v int32) *GetSdkListRequest
	GetGroup() *int32
	SetProduct(v string) *GetSdkListRequest
	GetProduct() *string
}

type GetSdkListRequest struct {
	Group   *int32  `json:"Group,omitempty" xml:"Group,omitempty"`
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
}

func (s GetSdkListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSdkListRequest) GoString() string {
	return s.String()
}

func (s *GetSdkListRequest) GetGroup() *int32 {
	return s.Group
}

func (s *GetSdkListRequest) GetProduct() *string {
	return s.Product
}

func (s *GetSdkListRequest) SetGroup(v int32) *GetSdkListRequest {
	s.Group = &v
	return s
}

func (s *GetSdkListRequest) SetProduct(v string) *GetSdkListRequest {
	s.Product = &v
	return s
}

func (s *GetSdkListRequest) Validate() error {
	return dara.Validate(s)
}
