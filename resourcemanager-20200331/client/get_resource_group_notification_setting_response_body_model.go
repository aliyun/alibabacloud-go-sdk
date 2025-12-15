// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceGroupNotificationSettingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetResourceGroupNotificationSettingResponseBody
	GetRequestId() *string
	SetResourceGroupNotificationEnableStatus(v bool) *GetResourceGroupNotificationSettingResponseBody
	GetResourceGroupNotificationEnableStatus() *bool
}

type GetResourceGroupNotificationSettingResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 7F70D09B-0EE5-54A6-A09A-1EBDB9297172
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the group event notification is enabled.
	//
	// example:
	//
	// true
	ResourceGroupNotificationEnableStatus *bool `json:"ResourceGroupNotificationEnableStatus,omitempty" xml:"ResourceGroupNotificationEnableStatus,omitempty"`
}

func (s GetResourceGroupNotificationSettingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetResourceGroupNotificationSettingResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceGroupNotificationSettingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetResourceGroupNotificationSettingResponseBody) GetResourceGroupNotificationEnableStatus() *bool {
	return s.ResourceGroupNotificationEnableStatus
}

func (s *GetResourceGroupNotificationSettingResponseBody) SetRequestId(v string) *GetResourceGroupNotificationSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourceGroupNotificationSettingResponseBody) SetResourceGroupNotificationEnableStatus(v bool) *GetResourceGroupNotificationSettingResponseBody {
	s.ResourceGroupNotificationEnableStatus = &v
	return s
}

func (s *GetResourceGroupNotificationSettingResponseBody) Validate() error {
	return dara.Validate(s)
}
