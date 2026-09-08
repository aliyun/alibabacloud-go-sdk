// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMarketingPreferenceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateMarketingPreferenceResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateMarketingPreferenceResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateMarketingPreferenceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateMarketingPreferenceResponseBody
	GetSuccess() *bool
}

type UpdateMarketingPreferenceResponseBody struct {
	// The business status code.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message returned when the call failed.
	//
	// example:
	//
	// /
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// /
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// - true: The call was successful.
	//
	// - false: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateMarketingPreferenceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMarketingPreferenceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMarketingPreferenceResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateMarketingPreferenceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateMarketingPreferenceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMarketingPreferenceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateMarketingPreferenceResponseBody) SetCode(v string) *UpdateMarketingPreferenceResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateMarketingPreferenceResponseBody) SetMessage(v string) *UpdateMarketingPreferenceResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateMarketingPreferenceResponseBody) SetRequestId(v string) *UpdateMarketingPreferenceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMarketingPreferenceResponseBody) SetSuccess(v bool) *UpdateMarketingPreferenceResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateMarketingPreferenceResponseBody) Validate() error {
	return dara.Validate(s)
}
