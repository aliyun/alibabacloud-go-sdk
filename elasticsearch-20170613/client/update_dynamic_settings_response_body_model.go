// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDynamicSettingsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateDynamicSettingsResponseBody
	GetRequestId() *string
	SetResult(v bool) *UpdateDynamicSettingsResponseBody
	GetResult() *bool
}

type UpdateDynamicSettingsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s UpdateDynamicSettingsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDynamicSettingsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDynamicSettingsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDynamicSettingsResponseBody) GetResult() *bool {
	return s.Result
}

func (s *UpdateDynamicSettingsResponseBody) SetRequestId(v string) *UpdateDynamicSettingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDynamicSettingsResponseBody) SetResult(v bool) *UpdateDynamicSettingsResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateDynamicSettingsResponseBody) Validate() error {
	return dara.Validate(s)
}
