// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBrandResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBrandId(v string) *CreateBrandResponseBody
	GetBrandId() *string
	SetRequestId(v string) *CreateBrandResponseBody
	GetRequestId() *string
}

type CreateBrandResponseBody struct {
	// example:
	//
	// brand_xxxx
	BrandId *string `json:"BrandId,omitempty" xml:"BrandId,omitempty"`
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateBrandResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateBrandResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBrandResponseBody) GetBrandId() *string {
	return s.BrandId
}

func (s *CreateBrandResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateBrandResponseBody) SetBrandId(v string) *CreateBrandResponseBody {
	s.BrandId = &v
	return s
}

func (s *CreateBrandResponseBody) SetRequestId(v string) *CreateBrandResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBrandResponseBody) Validate() error {
	return dara.Validate(s)
}
