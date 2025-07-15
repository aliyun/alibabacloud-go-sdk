// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveApiProductsAuthoritiesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveApiProductsAuthoritiesResponseBody
	GetRequestId() *string
}

type RemoveApiProductsAuthoritiesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CEB6EC62-B6C7-5082-A45A-45A204724AC2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveApiProductsAuthoritiesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveApiProductsAuthoritiesResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveApiProductsAuthoritiesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveApiProductsAuthoritiesResponseBody) SetRequestId(v string) *RemoveApiProductsAuthoritiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveApiProductsAuthoritiesResponseBody) Validate() error {
	return dara.Validate(s)
}
