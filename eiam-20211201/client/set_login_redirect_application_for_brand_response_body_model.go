// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLoginRedirectApplicationForBrandResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetLoginRedirectApplicationForBrandResponseBody
	GetRequestId() *string
}

type SetLoginRedirectApplicationForBrandResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetLoginRedirectApplicationForBrandResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetLoginRedirectApplicationForBrandResponseBody) GoString() string {
	return s.String()
}

func (s *SetLoginRedirectApplicationForBrandResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetLoginRedirectApplicationForBrandResponseBody) SetRequestId(v string) *SetLoginRedirectApplicationForBrandResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetLoginRedirectApplicationForBrandResponseBody) Validate() error {
	return dara.Validate(s)
}
