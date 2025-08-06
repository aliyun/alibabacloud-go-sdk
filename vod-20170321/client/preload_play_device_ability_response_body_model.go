// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreloadPlayDeviceAbilityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *PreloadPlayDeviceAbilityResponseBody
	GetRequestId() *string
	SetResult(v bool) *PreloadPlayDeviceAbilityResponseBody
	GetResult() *bool
}

type PreloadPlayDeviceAbilityResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s PreloadPlayDeviceAbilityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PreloadPlayDeviceAbilityResponseBody) GoString() string {
	return s.String()
}

func (s *PreloadPlayDeviceAbilityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PreloadPlayDeviceAbilityResponseBody) GetResult() *bool {
	return s.Result
}

func (s *PreloadPlayDeviceAbilityResponseBody) SetRequestId(v string) *PreloadPlayDeviceAbilityResponseBody {
	s.RequestId = &v
	return s
}

func (s *PreloadPlayDeviceAbilityResponseBody) SetResult(v bool) *PreloadPlayDeviceAbilityResponseBody {
	s.Result = &v
	return s
}

func (s *PreloadPlayDeviceAbilityResponseBody) Validate() error {
	return dara.Validate(s)
}
