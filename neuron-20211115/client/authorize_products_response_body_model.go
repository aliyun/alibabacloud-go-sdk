// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeProductsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AuthorizeProductsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AuthorizeProductsResponseBody
	GetSuccess() *bool
}

type AuthorizeProductsResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// sdadawqewe
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AuthorizeProductsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeProductsResponseBody) GoString() string {
	return s.String()
}

func (s *AuthorizeProductsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AuthorizeProductsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AuthorizeProductsResponseBody) SetRequestId(v string) *AuthorizeProductsResponseBody {
	s.RequestId = &v
	return s
}

func (s *AuthorizeProductsResponseBody) SetSuccess(v bool) *AuthorizeProductsResponseBody {
	s.Success = &v
	return s
}

func (s *AuthorizeProductsResponseBody) Validate() error {
	return dara.Validate(s)
}
