// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEnvironmentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetEnvironmentResponseBody
	GetCode() *string
	SetData(v *GetEnvironmentResponseBodyData) *GetEnvironmentResponseBody
	GetData() *GetEnvironmentResponseBodyData
	SetMessage(v string) *GetEnvironmentResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetEnvironmentResponseBody
	GetRequestId() *string
}

type GetEnvironmentResponseBody struct {
	// Response code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Response data.
	Data *GetEnvironmentResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Request ID, used for tracing the API call chain.
	//
	// example:
	//
	// 3F8EE674-BB08-5E92-BE6F-E4756A748B0F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetEnvironmentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetEnvironmentResponseBody) GoString() string {
	return s.String()
}

func (s *GetEnvironmentResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetEnvironmentResponseBody) GetData() *GetEnvironmentResponseBodyData {
	return s.Data
}

func (s *GetEnvironmentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetEnvironmentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetEnvironmentResponseBody) SetCode(v string) *GetEnvironmentResponseBody {
	s.Code = &v
	return s
}

func (s *GetEnvironmentResponseBody) SetData(v *GetEnvironmentResponseBodyData) *GetEnvironmentResponseBody {
	s.Data = v
	return s
}

func (s *GetEnvironmentResponseBody) SetMessage(v string) *GetEnvironmentResponseBody {
	s.Message = &v
	return s
}

func (s *GetEnvironmentResponseBody) SetRequestId(v string) *GetEnvironmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEnvironmentResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetEnvironmentResponseBodyData struct {
	// Environment alias.
	//
	// example:
	//
	// 测试环境
	Alias *string `json:"alias,omitempty" xml:"alias,omitempty"`
	// Creation timestamp.
	//
	// example:
	//
	// 1719386834548
	CreateTimestamp *int64 `json:"createTimestamp,omitempty" xml:"createTimestamp,omitempty"`
	// Whether it is the default environment.
	//
	// example:
	//
	// true
	Default *bool `json:"default,omitempty" xml:"default,omitempty"`
	// Environment description.
	//
	// example:
	//
	// 这是xxx的xx项目测试环境
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Environment ID.
	//
	// example:
	//
	// env-cq7l5s5lhtgi6qasrdc0
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	// Gateway information
	GatewayInfo *GatewayInfo `json:"gatewayInfo,omitempty" xml:"gatewayInfo,omitempty"`
	// Environment name.
	//
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aekzzzntl5njbpi
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// Related resource information.
	StatisticsInfo *GetEnvironmentResponseBodyDataStatisticsInfo `json:"statisticsInfo,omitempty" xml:"statisticsInfo,omitempty" type:"Struct"`
	// List of subdomains.
	SubDomainInfos []*SubDomainInfo `json:"subDomainInfos,omitempty" xml:"subDomainInfos,omitempty" type:"Repeated"`
	// Update timestamp.
	//
	// example:
	//
	// 1719386834548
	UpdateTimestamp *int64 `json:"updateTimestamp,omitempty" xml:"updateTimestamp,omitempty"`
}

func (s GetEnvironmentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetEnvironmentResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetEnvironmentResponseBodyData) GetAlias() *string {
	return s.Alias
}

func (s *GetEnvironmentResponseBodyData) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *GetEnvironmentResponseBodyData) GetDefault() *bool {
	return s.Default
}

func (s *GetEnvironmentResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetEnvironmentResponseBodyData) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *GetEnvironmentResponseBodyData) GetGatewayInfo() *GatewayInfo {
	return s.GatewayInfo
}

func (s *GetEnvironmentResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetEnvironmentResponseBodyData) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetEnvironmentResponseBodyData) GetStatisticsInfo() *GetEnvironmentResponseBodyDataStatisticsInfo {
	return s.StatisticsInfo
}

func (s *GetEnvironmentResponseBodyData) GetSubDomainInfos() []*SubDomainInfo {
	return s.SubDomainInfos
}

func (s *GetEnvironmentResponseBodyData) GetUpdateTimestamp() *int64 {
	return s.UpdateTimestamp
}

func (s *GetEnvironmentResponseBodyData) SetAlias(v string) *GetEnvironmentResponseBodyData {
	s.Alias = &v
	return s
}

func (s *GetEnvironmentResponseBodyData) SetCreateTimestamp(v int64) *GetEnvironmentResponseBodyData {
	s.CreateTimestamp = &v
	return s
}

func (s *GetEnvironmentResponseBodyData) SetDefault(v bool) *GetEnvironmentResponseBodyData {
	s.Default = &v
	return s
}

func (s *GetEnvironmentResponseBodyData) SetDescription(v string) *GetEnvironmentResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetEnvironmentResponseBodyData) SetEnvironmentId(v string) *GetEnvironmentResponseBodyData {
	s.EnvironmentId = &v
	return s
}

func (s *GetEnvironmentResponseBodyData) SetGatewayInfo(v *GatewayInfo) *GetEnvironmentResponseBodyData {
	s.GatewayInfo = v
	return s
}

func (s *GetEnvironmentResponseBodyData) SetName(v string) *GetEnvironmentResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetEnvironmentResponseBodyData) SetResourceGroupId(v string) *GetEnvironmentResponseBodyData {
	s.ResourceGroupId = &v
	return s
}

func (s *GetEnvironmentResponseBodyData) SetStatisticsInfo(v *GetEnvironmentResponseBodyDataStatisticsInfo) *GetEnvironmentResponseBodyData {
	s.StatisticsInfo = v
	return s
}

func (s *GetEnvironmentResponseBodyData) SetSubDomainInfos(v []*SubDomainInfo) *GetEnvironmentResponseBodyData {
	s.SubDomainInfos = v
	return s
}

func (s *GetEnvironmentResponseBodyData) SetUpdateTimestamp(v int64) *GetEnvironmentResponseBodyData {
	s.UpdateTimestamp = &v
	return s
}

func (s *GetEnvironmentResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetEnvironmentResponseBodyDataStatisticsInfo struct {
	// The array of related resource information.
	ResourceStatistics []*ResourceStatistic `json:"resourceStatistics,omitempty" xml:"resourceStatistics,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 4
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s GetEnvironmentResponseBodyDataStatisticsInfo) String() string {
	return dara.Prettify(s)
}

func (s GetEnvironmentResponseBodyDataStatisticsInfo) GoString() string {
	return s.String()
}

func (s *GetEnvironmentResponseBodyDataStatisticsInfo) GetResourceStatistics() []*ResourceStatistic {
	return s.ResourceStatistics
}

func (s *GetEnvironmentResponseBodyDataStatisticsInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetEnvironmentResponseBodyDataStatisticsInfo) SetResourceStatistics(v []*ResourceStatistic) *GetEnvironmentResponseBodyDataStatisticsInfo {
	s.ResourceStatistics = v
	return s
}

func (s *GetEnvironmentResponseBodyDataStatisticsInfo) SetTotalCount(v int32) *GetEnvironmentResponseBodyDataStatisticsInfo {
	s.TotalCount = &v
	return s
}

func (s *GetEnvironmentResponseBodyDataStatisticsInfo) Validate() error {
	return dara.Validate(s)
}
