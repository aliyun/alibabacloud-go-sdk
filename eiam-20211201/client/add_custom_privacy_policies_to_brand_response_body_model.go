// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCustomPrivacyPoliciesToBrandResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddCustomPrivacyPoliciesToBrandResponseBody
	GetRequestId() *string
}

type AddCustomPrivacyPoliciesToBrandResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddCustomPrivacyPoliciesToBrandResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddCustomPrivacyPoliciesToBrandResponseBody) GoString() string {
	return s.String()
}

func (s *AddCustomPrivacyPoliciesToBrandResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddCustomPrivacyPoliciesToBrandResponseBody) SetRequestId(v string) *AddCustomPrivacyPoliciesToBrandResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddCustomPrivacyPoliciesToBrandResponseBody) Validate() error {
	return dara.Validate(s)
}
