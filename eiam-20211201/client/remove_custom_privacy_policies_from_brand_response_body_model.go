// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveCustomPrivacyPoliciesFromBrandResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveCustomPrivacyPoliciesFromBrandResponseBody
	GetRequestId() *string
}

type RemoveCustomPrivacyPoliciesFromBrandResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveCustomPrivacyPoliciesFromBrandResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveCustomPrivacyPoliciesFromBrandResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveCustomPrivacyPoliciesFromBrandResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveCustomPrivacyPoliciesFromBrandResponseBody) SetRequestId(v string) *RemoveCustomPrivacyPoliciesFromBrandResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveCustomPrivacyPoliciesFromBrandResponseBody) Validate() error {
	return dara.Validate(s)
}
