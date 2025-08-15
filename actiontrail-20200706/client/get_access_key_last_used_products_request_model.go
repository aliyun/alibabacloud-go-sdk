// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccessKeyLastUsedProductsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessKey(v string) *GetAccessKeyLastUsedProductsRequest
	GetAccessKey() *string
}

type GetAccessKeyLastUsedProductsRequest struct {
	// The AccessKey ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// LTAI****************
	AccessKey *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
}

func (s GetAccessKeyLastUsedProductsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAccessKeyLastUsedProductsRequest) GoString() string {
	return s.String()
}

func (s *GetAccessKeyLastUsedProductsRequest) GetAccessKey() *string {
	return s.AccessKey
}

func (s *GetAccessKeyLastUsedProductsRequest) SetAccessKey(v string) *GetAccessKeyLastUsedProductsRequest {
	s.AccessKey = &v
	return s
}

func (s *GetAccessKeyLastUsedProductsRequest) Validate() error {
	return dara.Validate(s)
}
