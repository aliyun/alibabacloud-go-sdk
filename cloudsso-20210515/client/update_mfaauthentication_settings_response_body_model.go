// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMFAAuthenticationSettingsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateMFAAuthenticationSettingsResponseBody
	GetRequestId() *string
}

type UpdateMFAAuthenticationSettingsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// A1C748E3-8944-5593-81BC-7D96AE24F77B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateMFAAuthenticationSettingsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMFAAuthenticationSettingsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMFAAuthenticationSettingsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMFAAuthenticationSettingsResponseBody) SetRequestId(v string) *UpdateMFAAuthenticationSettingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMFAAuthenticationSettingsResponseBody) Validate() error {
	return dara.Validate(s)
}
