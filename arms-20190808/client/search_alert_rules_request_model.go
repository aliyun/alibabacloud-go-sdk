// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchAlertRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertRuleId(v string) *SearchAlertRulesRequest
	GetAlertRuleId() *string
	SetAppType(v string) *SearchAlertRulesRequest
	GetAppType() *string
	SetCurrentPage(v int32) *SearchAlertRulesRequest
	GetCurrentPage() *int32
	SetPageSize(v int32) *SearchAlertRulesRequest
	GetPageSize() *int32
	SetPid(v string) *SearchAlertRulesRequest
	GetPid() *string
	SetRegionId(v string) *SearchAlertRulesRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *SearchAlertRulesRequest
	GetResourceGroupId() *string
	SetSystemRegionId(v string) *SearchAlertRulesRequest
	GetSystemRegionId() *string
	SetTags(v []*SearchAlertRulesRequestTags) *SearchAlertRulesRequest
	GetTags() []*SearchAlertRulesRequestTags
	SetTitle(v string) *SearchAlertRulesRequest
	GetTitle() *string
	SetType(v string) *SearchAlertRulesRequest
	GetType() *string
}

type SearchAlertRulesRequest struct {
	// The id of AlertRule.
	//
	// example:
	//
	// 12345
	AlertRuleId *string `json:"AlertRuleId,omitempty" xml:"AlertRuleId,omitempty"`
	// The type of the application that is associated with the alert rule. Valid values:
	//
	// 	- `TRACE`: application
	//
	// 	- `RETCODE`: browser
	//
	// example:
	//
	// TRACE
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// The page number of the page to return. Default value: `1`.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries to return per page. Default value: `10`.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The process identifier (PID) of the application that is associated with the alert rule. For more information about how to obtain the PID, see [Obtain the PID of an application](https://help.aliyun.com/document_detail/186100.html?spm=a2c4g.11186623.6.792.1b50654cqcDPyk#title-imy-7gj-qhr).
	//
	// example:
	//
	// atc889zkcf@d8deedfa9bf****
	Pid *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	// The region ID of the alert data. For more information about the mappings between **RegionId*	- and **SystemRegionId**, see the detailed description below the table.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group. You can obtain the resource group ID in the **Resource Management*	- console.
	//
	// example:
	//
	// rg-acfmxyexli2****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The region ID of the alert rule. For more information about the mappings between **RegionId*	- and **SystemRegionId**, see the detailed description below the table.
	//
	// example:
	//
	// cn-hangzhou
	SystemRegionId *string `json:"SystemRegionId,omitempty" xml:"SystemRegionId,omitempty"`
	// The list of tags.
	Tags []*SearchAlertRulesRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The alert rule name.
	//
	// example:
	//
	// AlertRuleTitle
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The alert rule type. Valid values:
	//
	// 	- `1`: custom alert rules that are used to monitor drill-down data sets
	//
	// 	- `3`: custom alert rules that are used to monitor tiled data sets
	//
	// 	- `4`: alert rules that are used to monitor the browser, including the default frontend alert rules
	//
	// 	- `5`: alert rules that are used to monitor applications, including the default application alert rules
	//
	// 	- `6`: the default browser alert rules
	//
	// 	- `7`: the default application alert rules
	//
	// 	- `8`: Tracing Analysis alert rules
	//
	// 	- `101`: Prometheus alert rules
	//
	// example:
	//
	// 4
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SearchAlertRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchAlertRulesRequest) GoString() string {
	return s.String()
}

func (s *SearchAlertRulesRequest) GetAlertRuleId() *string {
	return s.AlertRuleId
}

func (s *SearchAlertRulesRequest) GetAppType() *string {
	return s.AppType
}

func (s *SearchAlertRulesRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *SearchAlertRulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchAlertRulesRequest) GetPid() *string {
	return s.Pid
}

func (s *SearchAlertRulesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SearchAlertRulesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *SearchAlertRulesRequest) GetSystemRegionId() *string {
	return s.SystemRegionId
}

func (s *SearchAlertRulesRequest) GetTags() []*SearchAlertRulesRequestTags {
	return s.Tags
}

func (s *SearchAlertRulesRequest) GetTitle() *string {
	return s.Title
}

func (s *SearchAlertRulesRequest) GetType() *string {
	return s.Type
}

func (s *SearchAlertRulesRequest) SetAlertRuleId(v string) *SearchAlertRulesRequest {
	s.AlertRuleId = &v
	return s
}

func (s *SearchAlertRulesRequest) SetAppType(v string) *SearchAlertRulesRequest {
	s.AppType = &v
	return s
}

func (s *SearchAlertRulesRequest) SetCurrentPage(v int32) *SearchAlertRulesRequest {
	s.CurrentPage = &v
	return s
}

func (s *SearchAlertRulesRequest) SetPageSize(v int32) *SearchAlertRulesRequest {
	s.PageSize = &v
	return s
}

func (s *SearchAlertRulesRequest) SetPid(v string) *SearchAlertRulesRequest {
	s.Pid = &v
	return s
}

func (s *SearchAlertRulesRequest) SetRegionId(v string) *SearchAlertRulesRequest {
	s.RegionId = &v
	return s
}

func (s *SearchAlertRulesRequest) SetResourceGroupId(v string) *SearchAlertRulesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *SearchAlertRulesRequest) SetSystemRegionId(v string) *SearchAlertRulesRequest {
	s.SystemRegionId = &v
	return s
}

func (s *SearchAlertRulesRequest) SetTags(v []*SearchAlertRulesRequestTags) *SearchAlertRulesRequest {
	s.Tags = v
	return s
}

func (s *SearchAlertRulesRequest) SetTitle(v string) *SearchAlertRulesRequest {
	s.Title = &v
	return s
}

func (s *SearchAlertRulesRequest) SetType(v string) *SearchAlertRulesRequest {
	s.Type = &v
	return s
}

func (s *SearchAlertRulesRequest) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SearchAlertRulesRequestTags struct {
	// The key of the tag. The following system preset fields are provided:
	//
	// 	- traceId: the ID of the trace.
	//
	// 	- serverApp: the name of the server application.
	//
	// 	- clientApp: the name of the client application.
	//
	// 	- service: the name of the operation.
	//
	// 	- rpc: the type of the call.
	//
	// 	- msOfSpan: the duration exceeds a specific value.
	//
	// 	- clientIp: the IP address of the client.
	//
	// 	- serverIp: the IP address of the server.
	//
	// 	- isError: specifies whether the call is abnormal.
	//
	// 	- hasTprof: contains only thread profiling.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SearchAlertRulesRequestTags) String() string {
	return dara.Prettify(s)
}

func (s SearchAlertRulesRequestTags) GoString() string {
	return s.String()
}

func (s *SearchAlertRulesRequestTags) GetKey() *string {
	return s.Key
}

func (s *SearchAlertRulesRequestTags) GetValue() *string {
	return s.Value
}

func (s *SearchAlertRulesRequestTags) SetKey(v string) *SearchAlertRulesRequestTags {
	s.Key = &v
	return s
}

func (s *SearchAlertRulesRequestTags) SetValue(v string) *SearchAlertRulesRequestTags {
	s.Value = &v
	return s
}

func (s *SearchAlertRulesRequestTags) Validate() error {
	return dara.Validate(s)
}
