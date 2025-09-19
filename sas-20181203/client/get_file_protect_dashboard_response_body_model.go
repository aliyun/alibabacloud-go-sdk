// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileProtectDashboardResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetFileProtectDashboardResponseBodyData) *GetFileProtectDashboardResponseBody
	GetData() *GetFileProtectDashboardResponseBodyData
	SetRequestId(v string) *GetFileProtectDashboardResponseBody
	GetRequestId() *string
}

type GetFileProtectDashboardResponseBody struct {
	// The response parameters.
	Data *GetFileProtectDashboardResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 6673D49C-A9AB-40DD-B4A2-B92306701AE7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetFileProtectDashboardResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFileProtectDashboardResponseBody) GoString() string {
	return s.String()
}

func (s *GetFileProtectDashboardResponseBody) GetData() *GetFileProtectDashboardResponseBodyData {
	return s.Data
}

func (s *GetFileProtectDashboardResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFileProtectDashboardResponseBody) SetData(v *GetFileProtectDashboardResponseBodyData) *GetFileProtectDashboardResponseBody {
	s.Data = v
	return s
}

func (s *GetFileProtectDashboardResponseBody) SetRequestId(v string) *GetFileProtectDashboardResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFileProtectDashboardResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetFileProtectDashboardResponseBodyData struct {
	// The total number of enabled rules.
	//
	// example:
	//
	// 12
	EnableRuleCount *int32 `json:"EnableRuleCount,omitempty" xml:"EnableRuleCount,omitempty"`
	// The total number of servers on which the Security Center agent is installed.
	//
	// example:
	//
	// 12
	PluginCount *int32 `json:"PluginCount,omitempty" xml:"PluginCount,omitempty"`
	// The total number of servers on which the Security Center agent is offline.
	//
	// example:
	//
	// 1
	PluginOfflineCount *int32 `json:"PluginOfflineCount,omitempty" xml:"PluginOfflineCount,omitempty"`
	// The total number of servers on which the Security Center agent is online.
	//
	// example:
	//
	// 11
	PluginOnlineCount *int32 `json:"PluginOnlineCount,omitempty" xml:"PluginOnlineCount,omitempty"`
}

func (s GetFileProtectDashboardResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetFileProtectDashboardResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetFileProtectDashboardResponseBodyData) GetEnableRuleCount() *int32 {
	return s.EnableRuleCount
}

func (s *GetFileProtectDashboardResponseBodyData) GetPluginCount() *int32 {
	return s.PluginCount
}

func (s *GetFileProtectDashboardResponseBodyData) GetPluginOfflineCount() *int32 {
	return s.PluginOfflineCount
}

func (s *GetFileProtectDashboardResponseBodyData) GetPluginOnlineCount() *int32 {
	return s.PluginOnlineCount
}

func (s *GetFileProtectDashboardResponseBodyData) SetEnableRuleCount(v int32) *GetFileProtectDashboardResponseBodyData {
	s.EnableRuleCount = &v
	return s
}

func (s *GetFileProtectDashboardResponseBodyData) SetPluginCount(v int32) *GetFileProtectDashboardResponseBodyData {
	s.PluginCount = &v
	return s
}

func (s *GetFileProtectDashboardResponseBodyData) SetPluginOfflineCount(v int32) *GetFileProtectDashboardResponseBodyData {
	s.PluginOfflineCount = &v
	return s
}

func (s *GetFileProtectDashboardResponseBodyData) SetPluginOnlineCount(v int32) *GetFileProtectDashboardResponseBodyData {
	s.PluginOnlineCount = &v
	return s
}

func (s *GetFileProtectDashboardResponseBodyData) Validate() error {
	return dara.Validate(s)
}
