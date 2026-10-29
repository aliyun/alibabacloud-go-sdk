// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSdnsRuleCategoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *CreateSdnsRuleCategoryResponseBody
	GetId() *string
	SetRequestId(v string) *CreateSdnsRuleCategoryResponseBody
	GetRequestId() *string
}

type CreateSdnsRuleCategoryResponseBody struct {
	// The ID of the policy created and saved in this request. This ID is consistent with the ID in the list and details, and is used for querying, updating, enabling, disabling, and deleting the policy. The ID is saved as a string to avoid precision loss of large integers.
	//
	// example:
	//
	// 123456789012345678
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the request. This ID is used to troubleshoot call issues and is not the policy ID.
	//
	// example:
	//
	// 50F9C40E-188D-4208-BE2C-74271337****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSdnsRuleCategoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSdnsRuleCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSdnsRuleCategoryResponseBody) GetId() *string {
	return s.Id
}

func (s *CreateSdnsRuleCategoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSdnsRuleCategoryResponseBody) SetId(v string) *CreateSdnsRuleCategoryResponseBody {
	s.Id = &v
	return s
}

func (s *CreateSdnsRuleCategoryResponseBody) SetRequestId(v string) *CreateSdnsRuleCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSdnsRuleCategoryResponseBody) Validate() error {
	return dara.Validate(s)
}
