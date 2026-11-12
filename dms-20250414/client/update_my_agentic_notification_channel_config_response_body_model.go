// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMyAgenticNotificationChannelConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *UpdateMyAgenticNotificationChannelConfigResponseBodyData) *UpdateMyAgenticNotificationChannelConfigResponseBody
	GetData() *UpdateMyAgenticNotificationChannelConfigResponseBodyData
	SetErrorCode(v string) *UpdateMyAgenticNotificationChannelConfigResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateMyAgenticNotificationChannelConfigResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *UpdateMyAgenticNotificationChannelConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateMyAgenticNotificationChannelConfigResponseBody
	GetSuccess() *bool
}

type UpdateMyAgenticNotificationChannelConfigResponseBody struct {
	// The updated notification channel configuration.
	Data *UpdateMyAgenticNotificationChannelConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned when the request fails.
	//
	// example:
	//
	// INVALID_PARAMETER
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned when the request fails.
	//
	// example:
	//
	// channel must be FEISHU
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0C1CB646-1DE4-4AD0-B4A4-7D47DD52E931
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values: true: The request is successful. false: The request fails.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateMyAgenticNotificationChannelConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMyAgenticNotificationChannelConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMyAgenticNotificationChannelConfigResponseBody) GetData() *UpdateMyAgenticNotificationChannelConfigResponseBodyData {
	return s.Data
}

func (s *UpdateMyAgenticNotificationChannelConfigResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateMyAgenticNotificationChannelConfigResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateMyAgenticNotificationChannelConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMyAgenticNotificationChannelConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateMyAgenticNotificationChannelConfigResponseBody) SetData(v *UpdateMyAgenticNotificationChannelConfigResponseBodyData) *UpdateMyAgenticNotificationChannelConfigResponseBody {
	s.Data = v
	return s
}

func (s *UpdateMyAgenticNotificationChannelConfigResponseBody) SetErrorCode(v string) *UpdateMyAgenticNotificationChannelConfigResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateMyAgenticNotificationChannelConfigResponseBody) SetErrorMessage(v string) *UpdateMyAgenticNotificationChannelConfigResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateMyAgenticNotificationChannelConfigResponseBody) SetRequestId(v string) *UpdateMyAgenticNotificationChannelConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMyAgenticNotificationChannelConfigResponseBody) SetSuccess(v bool) *UpdateMyAgenticNotificationChannelConfigResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateMyAgenticNotificationChannelConfigResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateMyAgenticNotificationChannelConfigResponseBodyData struct {
	// The notification channel type. Valid values: FEISHU (Lark).
	//
	// example:
	//
	// FEISHU
	Channel *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	// The channel configuration information. Both keys and values are strings, including the bound identity information such as openId, mobile, or email. Sensitive values are always displayed as ******, and plaintext is not returned.
	//
	// example:
	//
	// {"openId":"******","mobile":"******","email":"******"}
	Config map[string]*string `json:"Config,omitempty" xml:"Config,omitempty"`
	// Indicates whether the notification channel is enabled. Returns false if the channel has never been configured.
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
}

func (s UpdateMyAgenticNotificationChannelConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateMyAgenticNotificationChannelConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateMyAgenticNotificationChannelConfigResponseBodyData) GetChannel() *string {
	return s.Channel
}

func (s *UpdateMyAgenticNotificationChannelConfigResponseBodyData) GetConfig() map[string]*string {
	return s.Config
}

func (s *UpdateMyAgenticNotificationChannelConfigResponseBodyData) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateMyAgenticNotificationChannelConfigResponseBodyData) SetChannel(v string) *UpdateMyAgenticNotificationChannelConfigResponseBodyData {
	s.Channel = &v
	return s
}

func (s *UpdateMyAgenticNotificationChannelConfigResponseBodyData) SetConfig(v map[string]*string) *UpdateMyAgenticNotificationChannelConfigResponseBodyData {
	s.Config = v
	return s
}

func (s *UpdateMyAgenticNotificationChannelConfigResponseBodyData) SetEnabled(v bool) *UpdateMyAgenticNotificationChannelConfigResponseBodyData {
	s.Enabled = &v
	return s
}

func (s *UpdateMyAgenticNotificationChannelConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}
