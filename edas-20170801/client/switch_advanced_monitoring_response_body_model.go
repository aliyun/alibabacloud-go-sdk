// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchAdvancedMonitoringResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAdvancedMonitoringEnabled(v bool) *SwitchAdvancedMonitoringResponseBody
	GetAdvancedMonitoringEnabled() *bool
	SetCode(v int32) *SwitchAdvancedMonitoringResponseBody
	GetCode() *int32
	SetMessage(v string) *SwitchAdvancedMonitoringResponseBody
	GetMessage() *string
	SetRequestId(v string) *SwitchAdvancedMonitoringResponseBody
	GetRequestId() *string
}

type SwitchAdvancedMonitoringResponseBody struct {
	// Indicates whether the advanced application monitoring feature is enabled. Valid values:
	//
	// 	- true: The advanced application monitoring feature is enabled.
	//
	// 	- false: The advanced application monitoring feature is disabled.
	//
	// example:
	//
	// false
	AdvancedMonitoringEnabled *bool `json:"AdvancedMonitoringEnabled,omitempty" xml:"AdvancedMonitoringEnabled,omitempty"`
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The additional information that is returned.
	//
	// example:
	//
	// The advanced monitoring status is disabled already for application which app_id is 9e224bc6-a646-4484-ae31-e617b7e7****
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 577AED12-32D8-40B6-991F-17D7A601****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SwitchAdvancedMonitoringResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SwitchAdvancedMonitoringResponseBody) GoString() string {
	return s.String()
}

func (s *SwitchAdvancedMonitoringResponseBody) GetAdvancedMonitoringEnabled() *bool {
	return s.AdvancedMonitoringEnabled
}

func (s *SwitchAdvancedMonitoringResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *SwitchAdvancedMonitoringResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SwitchAdvancedMonitoringResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SwitchAdvancedMonitoringResponseBody) SetAdvancedMonitoringEnabled(v bool) *SwitchAdvancedMonitoringResponseBody {
	s.AdvancedMonitoringEnabled = &v
	return s
}

func (s *SwitchAdvancedMonitoringResponseBody) SetCode(v int32) *SwitchAdvancedMonitoringResponseBody {
	s.Code = &v
	return s
}

func (s *SwitchAdvancedMonitoringResponseBody) SetMessage(v string) *SwitchAdvancedMonitoringResponseBody {
	s.Message = &v
	return s
}

func (s *SwitchAdvancedMonitoringResponseBody) SetRequestId(v string) *SwitchAdvancedMonitoringResponseBody {
	s.RequestId = &v
	return s
}

func (s *SwitchAdvancedMonitoringResponseBody) Validate() error {
	return dara.Validate(s)
}
