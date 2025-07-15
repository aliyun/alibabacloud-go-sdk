// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetApiProductsAuthoritiesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetApiProductsAuthoritiesResponseBody
	GetRequestId() *string
}

type SetApiProductsAuthoritiesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 2603F41E-77FC-59A3-840E-296578A9BDE0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetApiProductsAuthoritiesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetApiProductsAuthoritiesResponseBody) GoString() string {
	return s.String()
}

func (s *SetApiProductsAuthoritiesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetApiProductsAuthoritiesResponseBody) SetRequestId(v string) *SetApiProductsAuthoritiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetApiProductsAuthoritiesResponseBody) Validate() error {
	return dara.Validate(s)
}
