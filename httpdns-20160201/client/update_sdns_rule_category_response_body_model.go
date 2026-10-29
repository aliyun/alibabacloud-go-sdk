// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSdnsRuleCategoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateSdnsRuleCategoryResponseBody
	GetRequestId() *string
}

type UpdateSdnsRuleCategoryResponseBody struct {
	// The ID of the request. You can use the ID to troubleshoot call issues. This ID is not the policy ID.
	//
	// example:
	//
	// 50F9C40E-188D-4208-BE2C-74271337****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSdnsRuleCategoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSdnsRuleCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSdnsRuleCategoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSdnsRuleCategoryResponseBody) SetRequestId(v string) *UpdateSdnsRuleCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSdnsRuleCategoryResponseBody) Validate() error {
	return dara.Validate(s)
}
