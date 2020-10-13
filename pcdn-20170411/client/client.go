// This file is auto-generated, don't edit it. Thanks.
package client

import (
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	rpc "github.com/alibabacloud-go/tea-rpc/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type StopDomainRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	Domain        *string `json:"Domain,omitempty" xml:"Domain,omitempty" require:"true"`
}

func (s StopDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s StopDomainRequest) GoString() string {
	return s.String()
}

func (s *StopDomainRequest) SetSecurityToken(v string) *StopDomainRequest {
	s.SecurityToken = &v
	return s
}

func (s *StopDomainRequest) SetVersion(v string) *StopDomainRequest {
	s.Version = &v
	return s
}

func (s *StopDomainRequest) SetDomain(v string) *StopDomainRequest {
	s.Domain = &v
	return s
}

type StopDomainResponse struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code       *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" require:"true"`
}

func (s StopDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s StopDomainResponse) GoString() string {
	return s.String()
}

func (s *StopDomainResponse) SetRequestId(v string) *StopDomainResponse {
	s.RequestId = &v
	return s
}

func (s *StopDomainResponse) SetCode(v int) *StopDomainResponse {
	s.Code = &v
	return s
}

func (s *StopDomainResponse) SetResourceId(v string) *StopDomainResponse {
	s.ResourceId = &v
	return s
}

type StartDomainRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	Domain        *string `json:"Domain,omitempty" xml:"Domain,omitempty" require:"true"`
}

func (s StartDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s StartDomainRequest) GoString() string {
	return s.String()
}

func (s *StartDomainRequest) SetSecurityToken(v string) *StartDomainRequest {
	s.SecurityToken = &v
	return s
}

func (s *StartDomainRequest) SetVersion(v string) *StartDomainRequest {
	s.Version = &v
	return s
}

func (s *StartDomainRequest) SetDomain(v string) *StartDomainRequest {
	s.Domain = &v
	return s
}

type StartDomainResponse struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code       *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" require:"true"`
}

func (s StartDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s StartDomainResponse) GoString() string {
	return s.String()
}

func (s *StartDomainResponse) SetRequestId(v string) *StartDomainResponse {
	s.RequestId = &v
	return s
}

func (s *StartDomainResponse) SetCode(v int) *StartDomainResponse {
	s.Code = &v
	return s
}

func (s *StartDomainResponse) SetResourceId(v string) *StartDomainResponse {
	s.ResourceId = &v
	return s
}

type DeleteDomainRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	Domain        *string `json:"Domain,omitempty" xml:"Domain,omitempty" require:"true"`
}

func (s DeleteDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDomainRequest) GoString() string {
	return s.String()
}

func (s *DeleteDomainRequest) SetSecurityToken(v string) *DeleteDomainRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteDomainRequest) SetVersion(v string) *DeleteDomainRequest {
	s.Version = &v
	return s
}

func (s *DeleteDomainRequest) SetDomain(v string) *DeleteDomainRequest {
	s.Domain = &v
	return s
}

type DeleteDomainResponse struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code       *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" require:"true"`
}

func (s DeleteDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDomainResponse) GoString() string {
	return s.String()
}

func (s *DeleteDomainResponse) SetRequestId(v string) *DeleteDomainResponse {
	s.RequestId = &v
	return s
}

func (s *DeleteDomainResponse) SetCode(v int) *DeleteDomainResponse {
	s.Code = &v
	return s
}

func (s *DeleteDomainResponse) SetResourceId(v string) *DeleteDomainResponse {
	s.ResourceId = &v
	return s
}

type AddDomainRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	BusinessType  *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty" require:"true"`
	Domain        *string `json:"Domain,omitempty" xml:"Domain,omitempty" require:"true"`
	LiveFormat    *string `json:"LiveFormat,omitempty" xml:"LiveFormat,omitempty"`
	SliceDomain   *string `json:"SliceDomain,omitempty" xml:"SliceDomain,omitempty"`
	Region        *string `json:"Region,omitempty" xml:"Region,omitempty"`
	DemoUrls      *string `json:"DemoUrls,omitempty" xml:"DemoUrls,omitempty"`
}

func (s AddDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s AddDomainRequest) GoString() string {
	return s.String()
}

func (s *AddDomainRequest) SetSecurityToken(v string) *AddDomainRequest {
	s.SecurityToken = &v
	return s
}

func (s *AddDomainRequest) SetVersion(v string) *AddDomainRequest {
	s.Version = &v
	return s
}

func (s *AddDomainRequest) SetBusinessType(v string) *AddDomainRequest {
	s.BusinessType = &v
	return s
}

func (s *AddDomainRequest) SetDomain(v string) *AddDomainRequest {
	s.Domain = &v
	return s
}

func (s *AddDomainRequest) SetLiveFormat(v string) *AddDomainRequest {
	s.LiveFormat = &v
	return s
}

func (s *AddDomainRequest) SetSliceDomain(v string) *AddDomainRequest {
	s.SliceDomain = &v
	return s
}

func (s *AddDomainRequest) SetRegion(v string) *AddDomainRequest {
	s.Region = &v
	return s
}

func (s *AddDomainRequest) SetDemoUrls(v string) *AddDomainRequest {
	s.DemoUrls = &v
	return s
}

type AddDomainResponse struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code       *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" require:"true"`
}

func (s AddDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s AddDomainResponse) GoString() string {
	return s.String()
}

func (s *AddDomainResponse) SetRequestId(v string) *AddDomainResponse {
	s.RequestId = &v
	return s
}

func (s *AddDomainResponse) SetCode(v int) *AddDomainResponse {
	s.Code = &v
	return s
}

func (s *AddDomainResponse) SetResourceId(v string) *AddDomainResponse {
	s.ResourceId = &v
	return s
}

type GetBalanceTrafficDataRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	DataInterval  *int    `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	ResourceId    *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" require:"true"`
}

func (s GetBalanceTrafficDataRequest) String() string {
	return tea.Prettify(s)
}

func (s GetBalanceTrafficDataRequest) GoString() string {
	return s.String()
}

func (s *GetBalanceTrafficDataRequest) SetSecurityToken(v string) *GetBalanceTrafficDataRequest {
	s.SecurityToken = &v
	return s
}

func (s *GetBalanceTrafficDataRequest) SetVersion(v string) *GetBalanceTrafficDataRequest {
	s.Version = &v
	return s
}

func (s *GetBalanceTrafficDataRequest) SetDataInterval(v int) *GetBalanceTrafficDataRequest {
	s.DataInterval = &v
	return s
}

func (s *GetBalanceTrafficDataRequest) SetResourceId(v string) *GetBalanceTrafficDataRequest {
	s.ResourceId = &v
	return s
}

type GetBalanceTrafficDataResponse struct {
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code      *int                                   `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	DataList  *GetBalanceTrafficDataResponseDataList `json:"DataList,omitempty" xml:"DataList,omitempty" require:"true" type:"Struct"`
	Labels    *GetBalanceTrafficDataResponseLabels   `json:"Labels,omitempty" xml:"Labels,omitempty" require:"true" type:"Struct"`
}

func (s GetBalanceTrafficDataResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBalanceTrafficDataResponse) GoString() string {
	return s.String()
}

func (s *GetBalanceTrafficDataResponse) SetRequestId(v string) *GetBalanceTrafficDataResponse {
	s.RequestId = &v
	return s
}

func (s *GetBalanceTrafficDataResponse) SetCode(v int) *GetBalanceTrafficDataResponse {
	s.Code = &v
	return s
}

func (s *GetBalanceTrafficDataResponse) SetDataList(v *GetBalanceTrafficDataResponseDataList) *GetBalanceTrafficDataResponse {
	s.DataList = v
	return s
}

func (s *GetBalanceTrafficDataResponse) SetLabels(v *GetBalanceTrafficDataResponseLabels) *GetBalanceTrafficDataResponse {
	s.Labels = v
	return s
}

type GetBalanceTrafficDataResponseDataList struct {
	UsageData []*GetBalanceTrafficDataResponseDataListUsageData `json:"UsageData,omitempty" xml:"UsageData,omitempty" require:"true" type:"Repeated"`
}

func (s GetBalanceTrafficDataResponseDataList) String() string {
	return tea.Prettify(s)
}

func (s GetBalanceTrafficDataResponseDataList) GoString() string {
	return s.String()
}

func (s *GetBalanceTrafficDataResponseDataList) SetUsageData(v []*GetBalanceTrafficDataResponseDataListUsageData) *GetBalanceTrafficDataResponseDataList {
	s.UsageData = v
	return s
}

type GetBalanceTrafficDataResponseDataListUsageData struct {
	Date   *string                                               `json:"Date,omitempty" xml:"Date,omitempty" require:"true"`
	Values *GetBalanceTrafficDataResponseDataListUsageDataValues `json:"Values,omitempty" xml:"Values,omitempty" require:"true" type:"Struct"`
}

func (s GetBalanceTrafficDataResponseDataListUsageData) String() string {
	return tea.Prettify(s)
}

func (s GetBalanceTrafficDataResponseDataListUsageData) GoString() string {
	return s.String()
}

func (s *GetBalanceTrafficDataResponseDataListUsageData) SetDate(v string) *GetBalanceTrafficDataResponseDataListUsageData {
	s.Date = &v
	return s
}

func (s *GetBalanceTrafficDataResponseDataListUsageData) SetValues(v *GetBalanceTrafficDataResponseDataListUsageDataValues) *GetBalanceTrafficDataResponseDataListUsageData {
	s.Values = v
	return s
}

type GetBalanceTrafficDataResponseDataListUsageDataValues struct {
	// Values
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" require:"true" type:"Repeated"`
}

func (s GetBalanceTrafficDataResponseDataListUsageDataValues) String() string {
	return tea.Prettify(s)
}

func (s GetBalanceTrafficDataResponseDataListUsageDataValues) GoString() string {
	return s.String()
}

func (s *GetBalanceTrafficDataResponseDataListUsageDataValues) SetValues(v []*string) *GetBalanceTrafficDataResponseDataListUsageDataValues {
	s.Values = v
	return s
}

type GetBalanceTrafficDataResponseLabels struct {
	Label []*string `json:"Label,omitempty" xml:"Label,omitempty" require:"true" type:"Repeated"`
}

func (s GetBalanceTrafficDataResponseLabels) String() string {
	return tea.Prettify(s)
}

func (s GetBalanceTrafficDataResponseLabels) GoString() string {
	return s.String()
}

func (s *GetBalanceTrafficDataResponseLabels) SetLabel(v []*string) *GetBalanceTrafficDataResponseLabels {
	s.Label = v
	return s
}

type AddPcdnControlRuleRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	Region        *string `json:"Region,omitempty" xml:"Region,omitempty" require:"true"`
	IspName       *string `json:"IspName,omitempty" xml:"IspName,omitempty" require:"true"`
	PlatformType  *string `json:"PlatformType,omitempty" xml:"PlatformType,omitempty" require:"true"`
	BusinessType  *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty" require:"true"`
	Market        *string `json:"Market,omitempty" xml:"Market,omitempty" require:"true"`
	AppVersion    *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty" require:"true"`
}

func (s AddPcdnControlRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s AddPcdnControlRuleRequest) GoString() string {
	return s.String()
}

func (s *AddPcdnControlRuleRequest) SetSecurityToken(v string) *AddPcdnControlRuleRequest {
	s.SecurityToken = &v
	return s
}

func (s *AddPcdnControlRuleRequest) SetVersion(v string) *AddPcdnControlRuleRequest {
	s.Version = &v
	return s
}

func (s *AddPcdnControlRuleRequest) SetName(v string) *AddPcdnControlRuleRequest {
	s.Name = &v
	return s
}

func (s *AddPcdnControlRuleRequest) SetRegion(v string) *AddPcdnControlRuleRequest {
	s.Region = &v
	return s
}

func (s *AddPcdnControlRuleRequest) SetIspName(v string) *AddPcdnControlRuleRequest {
	s.IspName = &v
	return s
}

func (s *AddPcdnControlRuleRequest) SetPlatformType(v string) *AddPcdnControlRuleRequest {
	s.PlatformType = &v
	return s
}

func (s *AddPcdnControlRuleRequest) SetBusinessType(v string) *AddPcdnControlRuleRequest {
	s.BusinessType = &v
	return s
}

func (s *AddPcdnControlRuleRequest) SetMarket(v string) *AddPcdnControlRuleRequest {
	s.Market = &v
	return s
}

func (s *AddPcdnControlRuleRequest) SetAppVersion(v string) *AddPcdnControlRuleRequest {
	s.AppVersion = &v
	return s
}

type AddPcdnControlRuleResponse struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code       *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" require:"true"`
}

func (s AddPcdnControlRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s AddPcdnControlRuleResponse) GoString() string {
	return s.String()
}

func (s *AddPcdnControlRuleResponse) SetRequestId(v string) *AddPcdnControlRuleResponse {
	s.RequestId = &v
	return s
}

func (s *AddPcdnControlRuleResponse) SetCode(v int) *AddPcdnControlRuleResponse {
	s.Code = &v
	return s
}

func (s *AddPcdnControlRuleResponse) SetResourceId(v string) *AddPcdnControlRuleResponse {
	s.ResourceId = &v
	return s
}

type AddConsumerRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Version              *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	BusinessType         *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty" require:"true"`
	Company              *string `json:"Company,omitempty" xml:"Company,omitempty" require:"true"`
	Site                 *string `json:"Site,omitempty" xml:"Site,omitempty" require:"true"`
	Requirement          *string `json:"Requirement,omitempty" xml:"Requirement,omitempty" require:"true"`
	Mobile               *string `json:"Mobile,omitempty" xml:"Mobile,omitempty" require:"true"`
	Ca                   *string `json:"Ca,omitempty" xml:"Ca,omitempty"`
	Operator             *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	Email                *string `json:"Email,omitempty" xml:"Email,omitempty"`
	BandwidthRequirement *string `json:"BandwidthRequirement,omitempty" xml:"BandwidthRequirement,omitempty"`
}

func (s AddConsumerRequest) String() string {
	return tea.Prettify(s)
}

func (s AddConsumerRequest) GoString() string {
	return s.String()
}

func (s *AddConsumerRequest) SetSecurityToken(v string) *AddConsumerRequest {
	s.SecurityToken = &v
	return s
}

func (s *AddConsumerRequest) SetVersion(v string) *AddConsumerRequest {
	s.Version = &v
	return s
}

func (s *AddConsumerRequest) SetBusinessType(v string) *AddConsumerRequest {
	s.BusinessType = &v
	return s
}

func (s *AddConsumerRequest) SetCompany(v string) *AddConsumerRequest {
	s.Company = &v
	return s
}

func (s *AddConsumerRequest) SetSite(v string) *AddConsumerRequest {
	s.Site = &v
	return s
}

func (s *AddConsumerRequest) SetRequirement(v string) *AddConsumerRequest {
	s.Requirement = &v
	return s
}

func (s *AddConsumerRequest) SetMobile(v string) *AddConsumerRequest {
	s.Mobile = &v
	return s
}

func (s *AddConsumerRequest) SetCa(v string) *AddConsumerRequest {
	s.Ca = &v
	return s
}

func (s *AddConsumerRequest) SetOperator(v string) *AddConsumerRequest {
	s.Operator = &v
	return s
}

func (s *AddConsumerRequest) SetEmail(v string) *AddConsumerRequest {
	s.Email = &v
	return s
}

func (s *AddConsumerRequest) SetBandwidthRequirement(v string) *AddConsumerRequest {
	s.BandwidthRequirement = &v
	return s
}

type AddConsumerResponse struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code       *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" require:"true"`
}

func (s AddConsumerResponse) String() string {
	return tea.Prettify(s)
}

func (s AddConsumerResponse) GoString() string {
	return s.String()
}

func (s *AddConsumerResponse) SetRequestId(v string) *AddConsumerResponse {
	s.RequestId = &v
	return s
}

func (s *AddConsumerResponse) SetCode(v int) *AddConsumerResponse {
	s.Code = &v
	return s
}

func (s *AddConsumerResponse) SetResourceId(v string) *AddConsumerResponse {
	s.ResourceId = &v
	return s
}

type GetAccessDataRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	Domain        *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Region        *string `json:"Region,omitempty" xml:"Region,omitempty" require:"true"`
	IspName       *string `json:"IspName,omitempty" xml:"IspName,omitempty" require:"true"`
	PlatformType  *string `json:"PlatformType,omitempty" xml:"PlatformType,omitempty" require:"true"`
	BusinessType  *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty" require:"true"`
	StartDate     *string `json:"StartDate,omitempty" xml:"StartDate,omitempty" require:"true"`
	EndDate       *string `json:"EndDate,omitempty" xml:"EndDate,omitempty" require:"true"`
}

func (s GetAccessDataRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAccessDataRequest) GoString() string {
	return s.String()
}

func (s *GetAccessDataRequest) SetSecurityToken(v string) *GetAccessDataRequest {
	s.SecurityToken = &v
	return s
}

func (s *GetAccessDataRequest) SetVersion(v string) *GetAccessDataRequest {
	s.Version = &v
	return s
}

func (s *GetAccessDataRequest) SetDomain(v string) *GetAccessDataRequest {
	s.Domain = &v
	return s
}

func (s *GetAccessDataRequest) SetRegion(v string) *GetAccessDataRequest {
	s.Region = &v
	return s
}

func (s *GetAccessDataRequest) SetIspName(v string) *GetAccessDataRequest {
	s.IspName = &v
	return s
}

func (s *GetAccessDataRequest) SetPlatformType(v string) *GetAccessDataRequest {
	s.PlatformType = &v
	return s
}

func (s *GetAccessDataRequest) SetBusinessType(v string) *GetAccessDataRequest {
	s.BusinessType = &v
	return s
}

func (s *GetAccessDataRequest) SetStartDate(v string) *GetAccessDataRequest {
	s.StartDate = &v
	return s
}

func (s *GetAccessDataRequest) SetEndDate(v string) *GetAccessDataRequest {
	s.EndDate = &v
	return s
}

type GetAccessDataResponse struct {
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code      *int                           `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	DataList  *GetAccessDataResponseDataList `json:"DataList,omitempty" xml:"DataList,omitempty" require:"true" type:"Struct"`
	Labels    *GetAccessDataResponseLabels   `json:"Labels,omitempty" xml:"Labels,omitempty" require:"true" type:"Struct"`
}

func (s GetAccessDataResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAccessDataResponse) GoString() string {
	return s.String()
}

func (s *GetAccessDataResponse) SetRequestId(v string) *GetAccessDataResponse {
	s.RequestId = &v
	return s
}

func (s *GetAccessDataResponse) SetCode(v int) *GetAccessDataResponse {
	s.Code = &v
	return s
}

func (s *GetAccessDataResponse) SetDataList(v *GetAccessDataResponseDataList) *GetAccessDataResponse {
	s.DataList = v
	return s
}

func (s *GetAccessDataResponse) SetLabels(v *GetAccessDataResponseLabels) *GetAccessDataResponse {
	s.Labels = v
	return s
}

type GetAccessDataResponseDataList struct {
	UsageData []*GetAccessDataResponseDataListUsageData `json:"UsageData,omitempty" xml:"UsageData,omitempty" require:"true" type:"Repeated"`
}

func (s GetAccessDataResponseDataList) String() string {
	return tea.Prettify(s)
}

func (s GetAccessDataResponseDataList) GoString() string {
	return s.String()
}

func (s *GetAccessDataResponseDataList) SetUsageData(v []*GetAccessDataResponseDataListUsageData) *GetAccessDataResponseDataList {
	s.UsageData = v
	return s
}

type GetAccessDataResponseDataListUsageData struct {
	Date   *string                                       `json:"Date,omitempty" xml:"Date,omitempty" require:"true"`
	Values *GetAccessDataResponseDataListUsageDataValues `json:"Values,omitempty" xml:"Values,omitempty" require:"true" type:"Struct"`
}

func (s GetAccessDataResponseDataListUsageData) String() string {
	return tea.Prettify(s)
}

func (s GetAccessDataResponseDataListUsageData) GoString() string {
	return s.String()
}

func (s *GetAccessDataResponseDataListUsageData) SetDate(v string) *GetAccessDataResponseDataListUsageData {
	s.Date = &v
	return s
}

func (s *GetAccessDataResponseDataListUsageData) SetValues(v *GetAccessDataResponseDataListUsageDataValues) *GetAccessDataResponseDataListUsageData {
	s.Values = v
	return s
}

type GetAccessDataResponseDataListUsageDataValues struct {
	// Values
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" require:"true" type:"Repeated"`
}

func (s GetAccessDataResponseDataListUsageDataValues) String() string {
	return tea.Prettify(s)
}

func (s GetAccessDataResponseDataListUsageDataValues) GoString() string {
	return s.String()
}

func (s *GetAccessDataResponseDataListUsageDataValues) SetValues(v []*string) *GetAccessDataResponseDataListUsageDataValues {
	s.Values = v
	return s
}

type GetAccessDataResponseLabels struct {
	Label []*string `json:"Label,omitempty" xml:"Label,omitempty" require:"true" type:"Repeated"`
}

func (s GetAccessDataResponseLabels) String() string {
	return tea.Prettify(s)
}

func (s GetAccessDataResponseLabels) GoString() string {
	return s.String()
}

func (s *GetAccessDataResponseLabels) SetLabel(v []*string) *GetAccessDataResponseLabels {
	s.Label = v
	return s
}

type EnablePcdnControlRuleRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	ResourceId    *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" require:"true"`
}

func (s EnablePcdnControlRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s EnablePcdnControlRuleRequest) GoString() string {
	return s.String()
}

func (s *EnablePcdnControlRuleRequest) SetSecurityToken(v string) *EnablePcdnControlRuleRequest {
	s.SecurityToken = &v
	return s
}

func (s *EnablePcdnControlRuleRequest) SetVersion(v string) *EnablePcdnControlRuleRequest {
	s.Version = &v
	return s
}

func (s *EnablePcdnControlRuleRequest) SetResourceId(v string) *EnablePcdnControlRuleRequest {
	s.ResourceId = &v
	return s
}

type EnablePcdnControlRuleResponse struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code       *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" require:"true"`
}

func (s EnablePcdnControlRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s EnablePcdnControlRuleResponse) GoString() string {
	return s.String()
}

func (s *EnablePcdnControlRuleResponse) SetRequestId(v string) *EnablePcdnControlRuleResponse {
	s.RequestId = &v
	return s
}

func (s *EnablePcdnControlRuleResponse) SetCode(v int) *EnablePcdnControlRuleResponse {
	s.Code = &v
	return s
}

func (s *EnablePcdnControlRuleResponse) SetResourceId(v string) *EnablePcdnControlRuleResponse {
	s.ResourceId = &v
	return s
}

type EditPcdnControlRuleRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	ResourceId    *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" require:"true"`
	Region        *string `json:"Region,omitempty" xml:"Region,omitempty" require:"true"`
	IspName       *string `json:"IspName,omitempty" xml:"IspName,omitempty" require:"true"`
	PlatformType  *string `json:"PlatformType,omitempty" xml:"PlatformType,omitempty" require:"true"`
	BusinessType  *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty" require:"true"`
	Market        *string `json:"Market,omitempty" xml:"Market,omitempty" require:"true"`
	AppVersion    *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty" require:"true"`
}

func (s EditPcdnControlRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s EditPcdnControlRuleRequest) GoString() string {
	return s.String()
}

func (s *EditPcdnControlRuleRequest) SetSecurityToken(v string) *EditPcdnControlRuleRequest {
	s.SecurityToken = &v
	return s
}

func (s *EditPcdnControlRuleRequest) SetVersion(v string) *EditPcdnControlRuleRequest {
	s.Version = &v
	return s
}

func (s *EditPcdnControlRuleRequest) SetName(v string) *EditPcdnControlRuleRequest {
	s.Name = &v
	return s
}

func (s *EditPcdnControlRuleRequest) SetResourceId(v string) *EditPcdnControlRuleRequest {
	s.ResourceId = &v
	return s
}

func (s *EditPcdnControlRuleRequest) SetRegion(v string) *EditPcdnControlRuleRequest {
	s.Region = &v
	return s
}

func (s *EditPcdnControlRuleRequest) SetIspName(v string) *EditPcdnControlRuleRequest {
	s.IspName = &v
	return s
}

func (s *EditPcdnControlRuleRequest) SetPlatformType(v string) *EditPcdnControlRuleRequest {
	s.PlatformType = &v
	return s
}

func (s *EditPcdnControlRuleRequest) SetBusinessType(v string) *EditPcdnControlRuleRequest {
	s.BusinessType = &v
	return s
}

func (s *EditPcdnControlRuleRequest) SetMarket(v string) *EditPcdnControlRuleRequest {
	s.Market = &v
	return s
}

func (s *EditPcdnControlRuleRequest) SetAppVersion(v string) *EditPcdnControlRuleRequest {
	s.AppVersion = &v
	return s
}

type EditPcdnControlRuleResponse struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code       *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" require:"true"`
}

func (s EditPcdnControlRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s EditPcdnControlRuleResponse) GoString() string {
	return s.String()
}

func (s *EditPcdnControlRuleResponse) SetRequestId(v string) *EditPcdnControlRuleResponse {
	s.RequestId = &v
	return s
}

func (s *EditPcdnControlRuleResponse) SetCode(v int) *EditPcdnControlRuleResponse {
	s.Code = &v
	return s
}

func (s *EditPcdnControlRuleResponse) SetResourceId(v string) *EditPcdnControlRuleResponse {
	s.ResourceId = &v
	return s
}

type DisablePcdnControlRuleRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	ResourceId    *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" require:"true"`
}

func (s DisablePcdnControlRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DisablePcdnControlRuleRequest) GoString() string {
	return s.String()
}

func (s *DisablePcdnControlRuleRequest) SetSecurityToken(v string) *DisablePcdnControlRuleRequest {
	s.SecurityToken = &v
	return s
}

func (s *DisablePcdnControlRuleRequest) SetVersion(v string) *DisablePcdnControlRuleRequest {
	s.Version = &v
	return s
}

func (s *DisablePcdnControlRuleRequest) SetResourceId(v string) *DisablePcdnControlRuleRequest {
	s.ResourceId = &v
	return s
}

type DisablePcdnControlRuleResponse struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code       *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" require:"true"`
}

func (s DisablePcdnControlRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DisablePcdnControlRuleResponse) GoString() string {
	return s.String()
}

func (s *DisablePcdnControlRuleResponse) SetRequestId(v string) *DisablePcdnControlRuleResponse {
	s.RequestId = &v
	return s
}

func (s *DisablePcdnControlRuleResponse) SetCode(v int) *DisablePcdnControlRuleResponse {
	s.Code = &v
	return s
}

func (s *DisablePcdnControlRuleResponse) SetResourceId(v string) *DisablePcdnControlRuleResponse {
	s.ResourceId = &v
	return s
}

type DeletePcdnControlRuleRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	ResourceId    *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" require:"true"`
}

func (s DeletePcdnControlRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeletePcdnControlRuleRequest) GoString() string {
	return s.String()
}

func (s *DeletePcdnControlRuleRequest) SetSecurityToken(v string) *DeletePcdnControlRuleRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeletePcdnControlRuleRequest) SetVersion(v string) *DeletePcdnControlRuleRequest {
	s.Version = &v
	return s
}

func (s *DeletePcdnControlRuleRequest) SetResourceId(v string) *DeletePcdnControlRuleRequest {
	s.ResourceId = &v
	return s
}

type DeletePcdnControlRuleResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code      *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
}

func (s DeletePcdnControlRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePcdnControlRuleResponse) GoString() string {
	return s.String()
}

func (s *DeletePcdnControlRuleResponse) SetRequestId(v string) *DeletePcdnControlRuleResponse {
	s.RequestId = &v
	return s
}

func (s *DeletePcdnControlRuleResponse) SetCode(v int) *DeletePcdnControlRuleResponse {
	s.Code = &v
	return s
}

type GetAllPlatformTypesRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
}

func (s GetAllPlatformTypesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAllPlatformTypesRequest) GoString() string {
	return s.String()
}

func (s *GetAllPlatformTypesRequest) SetSecurityToken(v string) *GetAllPlatformTypesRequest {
	s.SecurityToken = &v
	return s
}

func (s *GetAllPlatformTypesRequest) SetVersion(v string) *GetAllPlatformTypesRequest {
	s.Version = &v
	return s
}

type GetAllPlatformTypesResponse struct {
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code      *int                                 `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	DataList  *GetAllPlatformTypesResponseDataList `json:"DataList,omitempty" xml:"DataList,omitempty" require:"true" type:"Struct"`
}

func (s GetAllPlatformTypesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAllPlatformTypesResponse) GoString() string {
	return s.String()
}

func (s *GetAllPlatformTypesResponse) SetRequestId(v string) *GetAllPlatformTypesResponse {
	s.RequestId = &v
	return s
}

func (s *GetAllPlatformTypesResponse) SetCode(v int) *GetAllPlatformTypesResponse {
	s.Code = &v
	return s
}

func (s *GetAllPlatformTypesResponse) SetDataList(v *GetAllPlatformTypesResponseDataList) *GetAllPlatformTypesResponse {
	s.DataList = v
	return s
}

type GetAllPlatformTypesResponseDataList struct {
	UsageData []*GetAllPlatformTypesResponseDataListUsageData `json:"UsageData,omitempty" xml:"UsageData,omitempty" require:"true" type:"Repeated"`
}

func (s GetAllPlatformTypesResponseDataList) String() string {
	return tea.Prettify(s)
}

func (s GetAllPlatformTypesResponseDataList) GoString() string {
	return s.String()
}

func (s *GetAllPlatformTypesResponseDataList) SetUsageData(v []*GetAllPlatformTypesResponseDataListUsageData) *GetAllPlatformTypesResponseDataList {
	s.UsageData = v
	return s
}

type GetAllPlatformTypesResponseDataListUsageData struct {
	Code *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
}

func (s GetAllPlatformTypesResponseDataListUsageData) String() string {
	return tea.Prettify(s)
}

func (s GetAllPlatformTypesResponseDataListUsageData) GoString() string {
	return s.String()
}

func (s *GetAllPlatformTypesResponseDataListUsageData) SetCode(v int) *GetAllPlatformTypesResponseDataListUsageData {
	s.Code = &v
	return s
}

func (s *GetAllPlatformTypesResponseDataListUsageData) SetName(v string) *GetAllPlatformTypesResponseDataListUsageData {
	s.Name = &v
	return s
}

type GetAllMarketsRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
}

func (s GetAllMarketsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAllMarketsRequest) GoString() string {
	return s.String()
}

func (s *GetAllMarketsRequest) SetSecurityToken(v string) *GetAllMarketsRequest {
	s.SecurityToken = &v
	return s
}

func (s *GetAllMarketsRequest) SetVersion(v string) *GetAllMarketsRequest {
	s.Version = &v
	return s
}

type GetAllMarketsResponse struct {
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code      *int                           `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	DataList  *GetAllMarketsResponseDataList `json:"DataList,omitempty" xml:"DataList,omitempty" require:"true" type:"Struct"`
}

func (s GetAllMarketsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAllMarketsResponse) GoString() string {
	return s.String()
}

func (s *GetAllMarketsResponse) SetRequestId(v string) *GetAllMarketsResponse {
	s.RequestId = &v
	return s
}

func (s *GetAllMarketsResponse) SetCode(v int) *GetAllMarketsResponse {
	s.Code = &v
	return s
}

func (s *GetAllMarketsResponse) SetDataList(v *GetAllMarketsResponseDataList) *GetAllMarketsResponse {
	s.DataList = v
	return s
}

type GetAllMarketsResponseDataList struct {
	UsageData []*GetAllMarketsResponseDataListUsageData `json:"UsageData,omitempty" xml:"UsageData,omitempty" require:"true" type:"Repeated"`
}

func (s GetAllMarketsResponseDataList) String() string {
	return tea.Prettify(s)
}

func (s GetAllMarketsResponseDataList) GoString() string {
	return s.String()
}

func (s *GetAllMarketsResponseDataList) SetUsageData(v []*GetAllMarketsResponseDataListUsageData) *GetAllMarketsResponseDataList {
	s.UsageData = v
	return s
}

type GetAllMarketsResponseDataListUsageData struct {
	Code       *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	MarketCode *string `json:"MarketCode,omitempty" xml:"MarketCode,omitempty" require:"true"`
	MarketName *string `json:"MarketName,omitempty" xml:"MarketName,omitempty" require:"true"`
}

func (s GetAllMarketsResponseDataListUsageData) String() string {
	return tea.Prettify(s)
}

func (s GetAllMarketsResponseDataListUsageData) GoString() string {
	return s.String()
}

func (s *GetAllMarketsResponseDataListUsageData) SetCode(v int) *GetAllMarketsResponseDataListUsageData {
	s.Code = &v
	return s
}

func (s *GetAllMarketsResponseDataListUsageData) SetMarketCode(v string) *GetAllMarketsResponseDataListUsageData {
	s.MarketCode = &v
	return s
}

func (s *GetAllMarketsResponseDataListUsageData) SetMarketName(v string) *GetAllMarketsResponseDataListUsageData {
	s.MarketName = &v
	return s
}

type GetAllIspRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
}

func (s GetAllIspRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAllIspRequest) GoString() string {
	return s.String()
}

func (s *GetAllIspRequest) SetSecurityToken(v string) *GetAllIspRequest {
	s.SecurityToken = &v
	return s
}

func (s *GetAllIspRequest) SetVersion(v string) *GetAllIspRequest {
	s.Version = &v
	return s
}

type GetAllIspResponse struct {
	RequestId *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code      *int                       `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	DataList  *GetAllIspResponseDataList `json:"DataList,omitempty" xml:"DataList,omitempty" require:"true" type:"Struct"`
}

func (s GetAllIspResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAllIspResponse) GoString() string {
	return s.String()
}

func (s *GetAllIspResponse) SetRequestId(v string) *GetAllIspResponse {
	s.RequestId = &v
	return s
}

func (s *GetAllIspResponse) SetCode(v int) *GetAllIspResponse {
	s.Code = &v
	return s
}

func (s *GetAllIspResponse) SetDataList(v *GetAllIspResponseDataList) *GetAllIspResponse {
	s.DataList = v
	return s
}

type GetAllIspResponseDataList struct {
	UsageData []*GetAllIspResponseDataListUsageData `json:"UsageData,omitempty" xml:"UsageData,omitempty" require:"true" type:"Repeated"`
}

func (s GetAllIspResponseDataList) String() string {
	return tea.Prettify(s)
}

func (s GetAllIspResponseDataList) GoString() string {
	return s.String()
}

func (s *GetAllIspResponseDataList) SetUsageData(v []*GetAllIspResponseDataListUsageData) *GetAllIspResponseDataList {
	s.UsageData = v
	return s
}

type GetAllIspResponseDataListUsageData struct {
	NameCn     *string `json:"NameCn,omitempty" xml:"NameCn,omitempty" require:"true"`
	NameEn     *string `json:"NameEn,omitempty" xml:"NameEn,omitempty" require:"true"`
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" require:"true"`
}

func (s GetAllIspResponseDataListUsageData) String() string {
	return tea.Prettify(s)
}

func (s GetAllIspResponseDataListUsageData) GoString() string {
	return s.String()
}

func (s *GetAllIspResponseDataListUsageData) SetNameCn(v string) *GetAllIspResponseDataListUsageData {
	s.NameCn = &v
	return s
}

func (s *GetAllIspResponseDataListUsageData) SetNameEn(v string) *GetAllIspResponseDataListUsageData {
	s.NameEn = &v
	return s
}

func (s *GetAllIspResponseDataListUsageData) SetResourceId(v string) *GetAllIspResponseDataListUsageData {
	s.ResourceId = &v
	return s
}

type GetAllAppVersionsRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
}

func (s GetAllAppVersionsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAllAppVersionsRequest) GoString() string {
	return s.String()
}

func (s *GetAllAppVersionsRequest) SetSecurityToken(v string) *GetAllAppVersionsRequest {
	s.SecurityToken = &v
	return s
}

func (s *GetAllAppVersionsRequest) SetVersion(v string) *GetAllAppVersionsRequest {
	s.Version = &v
	return s
}

type GetAllAppVersionsResponse struct {
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code      *int                               `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	DataList  *GetAllAppVersionsResponseDataList `json:"DataList,omitempty" xml:"DataList,omitempty" require:"true" type:"Struct"`
}

func (s GetAllAppVersionsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAllAppVersionsResponse) GoString() string {
	return s.String()
}

func (s *GetAllAppVersionsResponse) SetRequestId(v string) *GetAllAppVersionsResponse {
	s.RequestId = &v
	return s
}

func (s *GetAllAppVersionsResponse) SetCode(v int) *GetAllAppVersionsResponse {
	s.Code = &v
	return s
}

func (s *GetAllAppVersionsResponse) SetDataList(v *GetAllAppVersionsResponseDataList) *GetAllAppVersionsResponse {
	s.DataList = v
	return s
}

type GetAllAppVersionsResponseDataList struct {
	UsageData []*GetAllAppVersionsResponseDataListUsageData `json:"UsageData,omitempty" xml:"UsageData,omitempty" require:"true" type:"Repeated"`
}

func (s GetAllAppVersionsResponseDataList) String() string {
	return tea.Prettify(s)
}

func (s GetAllAppVersionsResponseDataList) GoString() string {
	return s.String()
}

func (s *GetAllAppVersionsResponseDataList) SetUsageData(v []*GetAllAppVersionsResponseDataListUsageData) *GetAllAppVersionsResponseDataList {
	s.UsageData = v
	return s
}

type GetAllAppVersionsResponseDataListUsageData struct {
	Code  *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty" require:"true"`
}

func (s GetAllAppVersionsResponseDataListUsageData) String() string {
	return tea.Prettify(s)
}

func (s GetAllAppVersionsResponseDataListUsageData) GoString() string {
	return s.String()
}

func (s *GetAllAppVersionsResponseDataListUsageData) SetCode(v int) *GetAllAppVersionsResponseDataListUsageData {
	s.Code = &v
	return s
}

func (s *GetAllAppVersionsResponseDataListUsageData) SetValue(v string) *GetAllAppVersionsResponseDataListUsageData {
	s.Value = &v
	return s
}

type GetConsumerStatusRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
}

func (s GetConsumerStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetConsumerStatusRequest) GoString() string {
	return s.String()
}

func (s *GetConsumerStatusRequest) SetSecurityToken(v string) *GetConsumerStatusRequest {
	s.SecurityToken = &v
	return s
}

func (s *GetConsumerStatusRequest) SetVersion(v string) *GetConsumerStatusRequest {
	s.Version = &v
	return s
}

type GetConsumerStatusResponse struct {
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code               *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	IntegreatedMode    *int    `json:"IntegreatedMode,omitempty" xml:"IntegreatedMode,omitempty" require:"true"`
	Inservice          *bool   `json:"Inservice,omitempty" xml:"Inservice,omitempty" require:"true"`
	RealtimeMonitor    *bool   `json:"RealtimeMonitor,omitempty" xml:"RealtimeMonitor,omitempty" require:"true"`
	LiveMonitor        *bool   `json:"LiveMonitor,omitempty" xml:"LiveMonitor,omitempty" require:"true"`
	CdnUrlRedirectFlag *bool   `json:"CdnUrlRedirectFlag,omitempty" xml:"CdnUrlRedirectFlag,omitempty" require:"true"`
	BusinessType       *bool   `json:"BusinessType,omitempty" xml:"BusinessType,omitempty" require:"true"`
	Audit              *int    `json:"Audit,omitempty" xml:"Audit,omitempty" require:"true"`
	Comment            *string `json:"Comment,omitempty" xml:"Comment,omitempty" require:"true"`
	CreatedAt          *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty" require:"true"`
	UpdatedAt          *string `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty" require:"true"`
}

func (s GetConsumerStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetConsumerStatusResponse) GoString() string {
	return s.String()
}

func (s *GetConsumerStatusResponse) SetRequestId(v string) *GetConsumerStatusResponse {
	s.RequestId = &v
	return s
}

func (s *GetConsumerStatusResponse) SetCode(v int) *GetConsumerStatusResponse {
	s.Code = &v
	return s
}

func (s *GetConsumerStatusResponse) SetIntegreatedMode(v int) *GetConsumerStatusResponse {
	s.IntegreatedMode = &v
	return s
}

func (s *GetConsumerStatusResponse) SetInservice(v bool) *GetConsumerStatusResponse {
	s.Inservice = &v
	return s
}

func (s *GetConsumerStatusResponse) SetRealtimeMonitor(v bool) *GetConsumerStatusResponse {
	s.RealtimeMonitor = &v
	return s
}

func (s *GetConsumerStatusResponse) SetLiveMonitor(v bool) *GetConsumerStatusResponse {
	s.LiveMonitor = &v
	return s
}

func (s *GetConsumerStatusResponse) SetCdnUrlRedirectFlag(v bool) *GetConsumerStatusResponse {
	s.CdnUrlRedirectFlag = &v
	return s
}

func (s *GetConsumerStatusResponse) SetBusinessType(v bool) *GetConsumerStatusResponse {
	s.BusinessType = &v
	return s
}

func (s *GetConsumerStatusResponse) SetAudit(v int) *GetConsumerStatusResponse {
	s.Audit = &v
	return s
}

func (s *GetConsumerStatusResponse) SetComment(v string) *GetConsumerStatusResponse {
	s.Comment = &v
	return s
}

func (s *GetConsumerStatusResponse) SetCreatedAt(v string) *GetConsumerStatusResponse {
	s.CreatedAt = &v
	return s
}

func (s *GetConsumerStatusResponse) SetUpdatedAt(v string) *GetConsumerStatusResponse {
	s.UpdatedAt = &v
	return s
}

type GetClientsRatioRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
}

func (s GetClientsRatioRequest) String() string {
	return tea.Prettify(s)
}

func (s GetClientsRatioRequest) GoString() string {
	return s.String()
}

func (s *GetClientsRatioRequest) SetSecurityToken(v string) *GetClientsRatioRequest {
	s.SecurityToken = &v
	return s
}

func (s *GetClientsRatioRequest) SetVersion(v string) *GetClientsRatioRequest {
	s.Version = &v
	return s
}

type GetClientsRatioResponse struct {
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code      *int                             `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	DataList  *GetClientsRatioResponseDataList `json:"DataList,omitempty" xml:"DataList,omitempty" require:"true" type:"Struct"`
}

func (s GetClientsRatioResponse) String() string {
	return tea.Prettify(s)
}

func (s GetClientsRatioResponse) GoString() string {
	return s.String()
}

func (s *GetClientsRatioResponse) SetRequestId(v string) *GetClientsRatioResponse {
	s.RequestId = &v
	return s
}

func (s *GetClientsRatioResponse) SetCode(v int) *GetClientsRatioResponse {
	s.Code = &v
	return s
}

func (s *GetClientsRatioResponse) SetDataList(v *GetClientsRatioResponseDataList) *GetClientsRatioResponse {
	s.DataList = v
	return s
}

type GetClientsRatioResponseDataList struct {
	UsageData []*GetClientsRatioResponseDataListUsageData `json:"UsageData,omitempty" xml:"UsageData,omitempty" require:"true" type:"Repeated"`
}

func (s GetClientsRatioResponseDataList) String() string {
	return tea.Prettify(s)
}

func (s GetClientsRatioResponseDataList) GoString() string {
	return s.String()
}

func (s *GetClientsRatioResponseDataList) SetUsageData(v []*GetClientsRatioResponseDataListUsageData) *GetClientsRatioResponseDataList {
	s.UsageData = v
	return s
}

type GetClientsRatioResponseDataListUsageData struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	Rate  *string `json:"Rate,omitempty" xml:"Rate,omitempty" require:"true"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty" require:"true"`
}

func (s GetClientsRatioResponseDataListUsageData) String() string {
	return tea.Prettify(s)
}

func (s GetClientsRatioResponseDataListUsageData) GoString() string {
	return s.String()
}

func (s *GetClientsRatioResponseDataListUsageData) SetName(v string) *GetClientsRatioResponseDataListUsageData {
	s.Name = &v
	return s
}

func (s *GetClientsRatioResponseDataListUsageData) SetRate(v string) *GetClientsRatioResponseDataListUsageData {
	s.Rate = &v
	return s
}

func (s *GetClientsRatioResponseDataListUsageData) SetValue(v string) *GetClientsRatioResponseDataListUsageData {
	s.Value = &v
	return s
}

type GetBandwidthDataRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	Domain        *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Region        *string `json:"Region,omitempty" xml:"Region,omitempty" require:"true"`
	IspName       *string `json:"IspName,omitempty" xml:"IspName,omitempty" require:"true"`
	PlatformType  *string `json:"PlatformType,omitempty" xml:"PlatformType,omitempty" require:"true"`
	BusinessType  *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty" require:"true"`
	StartDate     *string `json:"StartDate,omitempty" xml:"StartDate,omitempty" require:"true"`
	EndDate       *string `json:"EndDate,omitempty" xml:"EndDate,omitempty" require:"true"`
}

func (s GetBandwidthDataRequest) String() string {
	return tea.Prettify(s)
}

func (s GetBandwidthDataRequest) GoString() string {
	return s.String()
}

func (s *GetBandwidthDataRequest) SetSecurityToken(v string) *GetBandwidthDataRequest {
	s.SecurityToken = &v
	return s
}

func (s *GetBandwidthDataRequest) SetVersion(v string) *GetBandwidthDataRequest {
	s.Version = &v
	return s
}

func (s *GetBandwidthDataRequest) SetDomain(v string) *GetBandwidthDataRequest {
	s.Domain = &v
	return s
}

func (s *GetBandwidthDataRequest) SetRegion(v string) *GetBandwidthDataRequest {
	s.Region = &v
	return s
}

func (s *GetBandwidthDataRequest) SetIspName(v string) *GetBandwidthDataRequest {
	s.IspName = &v
	return s
}

func (s *GetBandwidthDataRequest) SetPlatformType(v string) *GetBandwidthDataRequest {
	s.PlatformType = &v
	return s
}

func (s *GetBandwidthDataRequest) SetBusinessType(v string) *GetBandwidthDataRequest {
	s.BusinessType = &v
	return s
}

func (s *GetBandwidthDataRequest) SetStartDate(v string) *GetBandwidthDataRequest {
	s.StartDate = &v
	return s
}

func (s *GetBandwidthDataRequest) SetEndDate(v string) *GetBandwidthDataRequest {
	s.EndDate = &v
	return s
}

type GetBandwidthDataResponse struct {
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code      *int                              `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	DataList  *GetBandwidthDataResponseDataList `json:"DataList,omitempty" xml:"DataList,omitempty" require:"true" type:"Struct"`
	Labels    *GetBandwidthDataResponseLabels   `json:"Labels,omitempty" xml:"Labels,omitempty" require:"true" type:"Struct"`
}

func (s GetBandwidthDataResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBandwidthDataResponse) GoString() string {
	return s.String()
}

func (s *GetBandwidthDataResponse) SetRequestId(v string) *GetBandwidthDataResponse {
	s.RequestId = &v
	return s
}

func (s *GetBandwidthDataResponse) SetCode(v int) *GetBandwidthDataResponse {
	s.Code = &v
	return s
}

func (s *GetBandwidthDataResponse) SetDataList(v *GetBandwidthDataResponseDataList) *GetBandwidthDataResponse {
	s.DataList = v
	return s
}

func (s *GetBandwidthDataResponse) SetLabels(v *GetBandwidthDataResponseLabels) *GetBandwidthDataResponse {
	s.Labels = v
	return s
}

type GetBandwidthDataResponseDataList struct {
	UsageData []*GetBandwidthDataResponseDataListUsageData `json:"UsageData,omitempty" xml:"UsageData,omitempty" require:"true" type:"Repeated"`
}

func (s GetBandwidthDataResponseDataList) String() string {
	return tea.Prettify(s)
}

func (s GetBandwidthDataResponseDataList) GoString() string {
	return s.String()
}

func (s *GetBandwidthDataResponseDataList) SetUsageData(v []*GetBandwidthDataResponseDataListUsageData) *GetBandwidthDataResponseDataList {
	s.UsageData = v
	return s
}

type GetBandwidthDataResponseDataListUsageData struct {
	Date   *string                                          `json:"Date,omitempty" xml:"Date,omitempty" require:"true"`
	Values *GetBandwidthDataResponseDataListUsageDataValues `json:"Values,omitempty" xml:"Values,omitempty" require:"true" type:"Struct"`
}

func (s GetBandwidthDataResponseDataListUsageData) String() string {
	return tea.Prettify(s)
}

func (s GetBandwidthDataResponseDataListUsageData) GoString() string {
	return s.String()
}

func (s *GetBandwidthDataResponseDataListUsageData) SetDate(v string) *GetBandwidthDataResponseDataListUsageData {
	s.Date = &v
	return s
}

func (s *GetBandwidthDataResponseDataListUsageData) SetValues(v *GetBandwidthDataResponseDataListUsageDataValues) *GetBandwidthDataResponseDataListUsageData {
	s.Values = v
	return s
}

type GetBandwidthDataResponseDataListUsageDataValues struct {
	// Values
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" require:"true" type:"Repeated"`
}

func (s GetBandwidthDataResponseDataListUsageDataValues) String() string {
	return tea.Prettify(s)
}

func (s GetBandwidthDataResponseDataListUsageDataValues) GoString() string {
	return s.String()
}

func (s *GetBandwidthDataResponseDataListUsageDataValues) SetValues(v []*string) *GetBandwidthDataResponseDataListUsageDataValues {
	s.Values = v
	return s
}

type GetBandwidthDataResponseLabels struct {
	Label []*string `json:"Label,omitempty" xml:"Label,omitempty" require:"true" type:"Repeated"`
}

func (s GetBandwidthDataResponseLabels) String() string {
	return tea.Prettify(s)
}

func (s GetBandwidthDataResponseLabels) GoString() string {
	return s.String()
}

func (s *GetBandwidthDataResponseLabels) SetLabel(v []*string) *GetBandwidthDataResponseLabels {
	s.Label = v
	return s
}

type GetBalanceBandwidthDataRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	DataInterval  *int    `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	ResourceId    *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" require:"true"`
}

func (s GetBalanceBandwidthDataRequest) String() string {
	return tea.Prettify(s)
}

func (s GetBalanceBandwidthDataRequest) GoString() string {
	return s.String()
}

func (s *GetBalanceBandwidthDataRequest) SetSecurityToken(v string) *GetBalanceBandwidthDataRequest {
	s.SecurityToken = &v
	return s
}

func (s *GetBalanceBandwidthDataRequest) SetVersion(v string) *GetBalanceBandwidthDataRequest {
	s.Version = &v
	return s
}

func (s *GetBalanceBandwidthDataRequest) SetDataInterval(v int) *GetBalanceBandwidthDataRequest {
	s.DataInterval = &v
	return s
}

func (s *GetBalanceBandwidthDataRequest) SetResourceId(v string) *GetBalanceBandwidthDataRequest {
	s.ResourceId = &v
	return s
}

type GetBalanceBandwidthDataResponse struct {
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code      *int                                     `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	DataList  *GetBalanceBandwidthDataResponseDataList `json:"DataList,omitempty" xml:"DataList,omitempty" require:"true" type:"Struct"`
	Labels    *GetBalanceBandwidthDataResponseLabels   `json:"Labels,omitempty" xml:"Labels,omitempty" require:"true" type:"Struct"`
}

func (s GetBalanceBandwidthDataResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBalanceBandwidthDataResponse) GoString() string {
	return s.String()
}

func (s *GetBalanceBandwidthDataResponse) SetRequestId(v string) *GetBalanceBandwidthDataResponse {
	s.RequestId = &v
	return s
}

func (s *GetBalanceBandwidthDataResponse) SetCode(v int) *GetBalanceBandwidthDataResponse {
	s.Code = &v
	return s
}

func (s *GetBalanceBandwidthDataResponse) SetDataList(v *GetBalanceBandwidthDataResponseDataList) *GetBalanceBandwidthDataResponse {
	s.DataList = v
	return s
}

func (s *GetBalanceBandwidthDataResponse) SetLabels(v *GetBalanceBandwidthDataResponseLabels) *GetBalanceBandwidthDataResponse {
	s.Labels = v
	return s
}

type GetBalanceBandwidthDataResponseDataList struct {
	UsageData []*GetBalanceBandwidthDataResponseDataListUsageData `json:"UsageData,omitempty" xml:"UsageData,omitempty" require:"true" type:"Repeated"`
}

func (s GetBalanceBandwidthDataResponseDataList) String() string {
	return tea.Prettify(s)
}

func (s GetBalanceBandwidthDataResponseDataList) GoString() string {
	return s.String()
}

func (s *GetBalanceBandwidthDataResponseDataList) SetUsageData(v []*GetBalanceBandwidthDataResponseDataListUsageData) *GetBalanceBandwidthDataResponseDataList {
	s.UsageData = v
	return s
}

type GetBalanceBandwidthDataResponseDataListUsageData struct {
	Date   *string                                                 `json:"Date,omitempty" xml:"Date,omitempty" require:"true"`
	Values *GetBalanceBandwidthDataResponseDataListUsageDataValues `json:"Values,omitempty" xml:"Values,omitempty" require:"true" type:"Struct"`
}

func (s GetBalanceBandwidthDataResponseDataListUsageData) String() string {
	return tea.Prettify(s)
}

func (s GetBalanceBandwidthDataResponseDataListUsageData) GoString() string {
	return s.String()
}

func (s *GetBalanceBandwidthDataResponseDataListUsageData) SetDate(v string) *GetBalanceBandwidthDataResponseDataListUsageData {
	s.Date = &v
	return s
}

func (s *GetBalanceBandwidthDataResponseDataListUsageData) SetValues(v *GetBalanceBandwidthDataResponseDataListUsageDataValues) *GetBalanceBandwidthDataResponseDataListUsageData {
	s.Values = v
	return s
}

type GetBalanceBandwidthDataResponseDataListUsageDataValues struct {
	// Values
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" require:"true" type:"Repeated"`
}

func (s GetBalanceBandwidthDataResponseDataListUsageDataValues) String() string {
	return tea.Prettify(s)
}

func (s GetBalanceBandwidthDataResponseDataListUsageDataValues) GoString() string {
	return s.String()
}

func (s *GetBalanceBandwidthDataResponseDataListUsageDataValues) SetValues(v []*string) *GetBalanceBandwidthDataResponseDataListUsageDataValues {
	s.Values = v
	return s
}

type GetBalanceBandwidthDataResponseLabels struct {
	Label []*string `json:"Label,omitempty" xml:"Label,omitempty" require:"true" type:"Repeated"`
}

func (s GetBalanceBandwidthDataResponseLabels) String() string {
	return tea.Prettify(s)
}

func (s GetBalanceBandwidthDataResponseLabels) GoString() string {
	return s.String()
}

func (s *GetBalanceBandwidthDataResponseLabels) SetLabel(v []*string) *GetBalanceBandwidthDataResponseLabels {
	s.Label = v
	return s
}

type GetControlRulesRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	Page          *string `json:"Page,omitempty" xml:"Page,omitempty" require:"true"`
	PageSize      *string `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
}

func (s GetControlRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetControlRulesRequest) GoString() string {
	return s.String()
}

func (s *GetControlRulesRequest) SetSecurityToken(v string) *GetControlRulesRequest {
	s.SecurityToken = &v
	return s
}

func (s *GetControlRulesRequest) SetVersion(v string) *GetControlRulesRequest {
	s.Version = &v
	return s
}

func (s *GetControlRulesRequest) SetPage(v string) *GetControlRulesRequest {
	s.Page = &v
	return s
}

func (s *GetControlRulesRequest) SetPageSize(v string) *GetControlRulesRequest {
	s.PageSize = &v
	return s
}

type GetControlRulesResponse struct {
	RequestId   *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code        *int                                `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	SettingList *GetControlRulesResponseSettingList `json:"SettingList,omitempty" xml:"SettingList,omitempty" require:"true" type:"Struct"`
	Pager       *GetControlRulesResponsePager       `json:"Pager,omitempty" xml:"Pager,omitempty" require:"true" type:"Struct"`
}

func (s GetControlRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetControlRulesResponse) GoString() string {
	return s.String()
}

func (s *GetControlRulesResponse) SetRequestId(v string) *GetControlRulesResponse {
	s.RequestId = &v
	return s
}

func (s *GetControlRulesResponse) SetCode(v int) *GetControlRulesResponse {
	s.Code = &v
	return s
}

func (s *GetControlRulesResponse) SetSettingList(v *GetControlRulesResponseSettingList) *GetControlRulesResponse {
	s.SettingList = v
	return s
}

func (s *GetControlRulesResponse) SetPager(v *GetControlRulesResponsePager) *GetControlRulesResponse {
	s.Pager = v
	return s
}

type GetControlRulesResponseSettingList struct {
	Setting []*GetControlRulesResponseSettingListSetting `json:"Setting,omitempty" xml:"Setting,omitempty" require:"true" type:"Repeated"`
}

func (s GetControlRulesResponseSettingList) String() string {
	return tea.Prettify(s)
}

func (s GetControlRulesResponseSettingList) GoString() string {
	return s.String()
}

func (s *GetControlRulesResponseSettingList) SetSetting(v []*GetControlRulesResponseSettingListSetting) *GetControlRulesResponseSettingList {
	s.Setting = v
	return s
}

type GetControlRulesResponseSettingListSetting struct {
	PlatformType *string `json:"PlatformType,omitempty" xml:"PlatformType,omitempty" require:"true"`
	AppVersion   *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty" require:"true"`
	IspName      *string `json:"IspName,omitempty" xml:"IspName,omitempty" require:"true"`
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty" require:"true"`
	ClientId     *string `json:"ClientId,omitempty" xml:"ClientId,omitempty" require:"true"`
	CreatedAt    *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty" require:"true"`
	MarketType   *string `json:"MarketType,omitempty" xml:"MarketType,omitempty" require:"true"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	Onoff        *bool   `json:"Onoff,omitempty" xml:"Onoff,omitempty" require:"true"`
	Usable       *bool   `json:"Usable,omitempty" xml:"Usable,omitempty" require:"true"`
	Region       *string `json:"Region,omitempty" xml:"Region,omitempty" require:"true"`
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" require:"true"`
	UpdatedAt    *string `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty" require:"true"`
}

func (s GetControlRulesResponseSettingListSetting) String() string {
	return tea.Prettify(s)
}

func (s GetControlRulesResponseSettingListSetting) GoString() string {
	return s.String()
}

func (s *GetControlRulesResponseSettingListSetting) SetPlatformType(v string) *GetControlRulesResponseSettingListSetting {
	s.PlatformType = &v
	return s
}

func (s *GetControlRulesResponseSettingListSetting) SetAppVersion(v string) *GetControlRulesResponseSettingListSetting {
	s.AppVersion = &v
	return s
}

func (s *GetControlRulesResponseSettingListSetting) SetIspName(v string) *GetControlRulesResponseSettingListSetting {
	s.IspName = &v
	return s
}

func (s *GetControlRulesResponseSettingListSetting) SetBusinessType(v string) *GetControlRulesResponseSettingListSetting {
	s.BusinessType = &v
	return s
}

func (s *GetControlRulesResponseSettingListSetting) SetClientId(v string) *GetControlRulesResponseSettingListSetting {
	s.ClientId = &v
	return s
}

func (s *GetControlRulesResponseSettingListSetting) SetCreatedAt(v string) *GetControlRulesResponseSettingListSetting {
	s.CreatedAt = &v
	return s
}

func (s *GetControlRulesResponseSettingListSetting) SetMarketType(v string) *GetControlRulesResponseSettingListSetting {
	s.MarketType = &v
	return s
}

func (s *GetControlRulesResponseSettingListSetting) SetName(v string) *GetControlRulesResponseSettingListSetting {
	s.Name = &v
	return s
}

func (s *GetControlRulesResponseSettingListSetting) SetOnoff(v bool) *GetControlRulesResponseSettingListSetting {
	s.Onoff = &v
	return s
}

func (s *GetControlRulesResponseSettingListSetting) SetUsable(v bool) *GetControlRulesResponseSettingListSetting {
	s.Usable = &v
	return s
}

func (s *GetControlRulesResponseSettingListSetting) SetRegion(v string) *GetControlRulesResponseSettingListSetting {
	s.Region = &v
	return s
}

func (s *GetControlRulesResponseSettingListSetting) SetResourceId(v string) *GetControlRulesResponseSettingListSetting {
	s.ResourceId = &v
	return s
}

func (s *GetControlRulesResponseSettingListSetting) SetUpdatedAt(v string) *GetControlRulesResponseSettingListSetting {
	s.UpdatedAt = &v
	return s
}

type GetControlRulesResponsePager struct {
	Page     *int `json:"Page,omitempty" xml:"Page,omitempty" require:"true"`
	Total    *int `json:"Total,omitempty" xml:"Total,omitempty" require:"true"`
	PageSize *int `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
}

func (s GetControlRulesResponsePager) String() string {
	return tea.Prettify(s)
}

func (s GetControlRulesResponsePager) GoString() string {
	return s.String()
}

func (s *GetControlRulesResponsePager) SetPage(v int) *GetControlRulesResponsePager {
	s.Page = &v
	return s
}

func (s *GetControlRulesResponsePager) SetTotal(v int) *GetControlRulesResponsePager {
	s.Total = &v
	return s
}

func (s *GetControlRulesResponsePager) SetPageSize(v int) *GetControlRulesResponsePager {
	s.PageSize = &v
	return s
}

type GetDomainCountRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
}

func (s GetDomainCountRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDomainCountRequest) GoString() string {
	return s.String()
}

func (s *GetDomainCountRequest) SetSecurityToken(v string) *GetDomainCountRequest {
	s.SecurityToken = &v
	return s
}

func (s *GetDomainCountRequest) SetVersion(v string) *GetDomainCountRequest {
	s.Version = &v
	return s
}

type GetDomainCountResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code      *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Data      *int    `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
}

func (s GetDomainCountResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDomainCountResponse) GoString() string {
	return s.String()
}

func (s *GetDomainCountResponse) SetRequestId(v string) *GetDomainCountResponse {
	s.RequestId = &v
	return s
}

func (s *GetDomainCountResponse) SetCode(v int) *GetDomainCountResponse {
	s.Code = &v
	return s
}

func (s *GetDomainCountResponse) SetData(v int) *GetDomainCountResponse {
	s.Data = &v
	return s
}

type GetCurrentModeRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
}

func (s GetCurrentModeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCurrentModeRequest) GoString() string {
	return s.String()
}

func (s *GetCurrentModeRequest) SetSecurityToken(v string) *GetCurrentModeRequest {
	s.SecurityToken = &v
	return s
}

func (s *GetCurrentModeRequest) SetVersion(v string) *GetCurrentModeRequest {
	s.Version = &v
	return s
}

type GetCurrentModeResponse struct {
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code              *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	ModeCode          *int    `json:"ModeCode,omitempty" xml:"ModeCode,omitempty" require:"true"`
	PaddingModeCode   *int    `json:"PaddingModeCode,omitempty" xml:"PaddingModeCode,omitempty" require:"true"`
	EffectiveAt       *int    `json:"EffectiveAt,omitempty" xml:"EffectiveAt,omitempty" require:"true"`
	EstimateBandwidth *int    `json:"EstimateBandwidth,omitempty" xml:"EstimateBandwidth,omitempty" require:"true"`
}

func (s GetCurrentModeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCurrentModeResponse) GoString() string {
	return s.String()
}

func (s *GetCurrentModeResponse) SetRequestId(v string) *GetCurrentModeResponse {
	s.RequestId = &v
	return s
}

func (s *GetCurrentModeResponse) SetCode(v int) *GetCurrentModeResponse {
	s.Code = &v
	return s
}

func (s *GetCurrentModeResponse) SetModeCode(v int) *GetCurrentModeResponse {
	s.ModeCode = &v
	return s
}

func (s *GetCurrentModeResponse) SetPaddingModeCode(v int) *GetCurrentModeResponse {
	s.PaddingModeCode = &v
	return s
}

func (s *GetCurrentModeResponse) SetEffectiveAt(v int) *GetCurrentModeResponse {
	s.EffectiveAt = &v
	return s
}

func (s *GetCurrentModeResponse) SetEstimateBandwidth(v int) *GetCurrentModeResponse {
	s.EstimateBandwidth = &v
	return s
}

type GetCoverRateDataRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	Domain        *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Region        *string `json:"Region,omitempty" xml:"Region,omitempty" require:"true"`
	IspName       *string `json:"IspName,omitempty" xml:"IspName,omitempty" require:"true"`
	PlatformType  *string `json:"PlatformType,omitempty" xml:"PlatformType,omitempty" require:"true"`
	BusinessType  *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty" require:"true"`
	StartDate     *string `json:"StartDate,omitempty" xml:"StartDate,omitempty" require:"true"`
	EndDate       *string `json:"EndDate,omitempty" xml:"EndDate,omitempty" require:"true"`
}

func (s GetCoverRateDataRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCoverRateDataRequest) GoString() string {
	return s.String()
}

func (s *GetCoverRateDataRequest) SetSecurityToken(v string) *GetCoverRateDataRequest {
	s.SecurityToken = &v
	return s
}

func (s *GetCoverRateDataRequest) SetVersion(v string) *GetCoverRateDataRequest {
	s.Version = &v
	return s
}

func (s *GetCoverRateDataRequest) SetDomain(v string) *GetCoverRateDataRequest {
	s.Domain = &v
	return s
}

func (s *GetCoverRateDataRequest) SetRegion(v string) *GetCoverRateDataRequest {
	s.Region = &v
	return s
}

func (s *GetCoverRateDataRequest) SetIspName(v string) *GetCoverRateDataRequest {
	s.IspName = &v
	return s
}

func (s *GetCoverRateDataRequest) SetPlatformType(v string) *GetCoverRateDataRequest {
	s.PlatformType = &v
	return s
}

func (s *GetCoverRateDataRequest) SetBusinessType(v string) *GetCoverRateDataRequest {
	s.BusinessType = &v
	return s
}

func (s *GetCoverRateDataRequest) SetStartDate(v string) *GetCoverRateDataRequest {
	s.StartDate = &v
	return s
}

func (s *GetCoverRateDataRequest) SetEndDate(v string) *GetCoverRateDataRequest {
	s.EndDate = &v
	return s
}

type GetCoverRateDataResponse struct {
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code      *int                              `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	DataList  *GetCoverRateDataResponseDataList `json:"DataList,omitempty" xml:"DataList,omitempty" require:"true" type:"Struct"`
	Labels    *GetCoverRateDataResponseLabels   `json:"Labels,omitempty" xml:"Labels,omitempty" require:"true" type:"Struct"`
}

func (s GetCoverRateDataResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCoverRateDataResponse) GoString() string {
	return s.String()
}

func (s *GetCoverRateDataResponse) SetRequestId(v string) *GetCoverRateDataResponse {
	s.RequestId = &v
	return s
}

func (s *GetCoverRateDataResponse) SetCode(v int) *GetCoverRateDataResponse {
	s.Code = &v
	return s
}

func (s *GetCoverRateDataResponse) SetDataList(v *GetCoverRateDataResponseDataList) *GetCoverRateDataResponse {
	s.DataList = v
	return s
}

func (s *GetCoverRateDataResponse) SetLabels(v *GetCoverRateDataResponseLabels) *GetCoverRateDataResponse {
	s.Labels = v
	return s
}

type GetCoverRateDataResponseDataList struct {
	UsageData []*GetCoverRateDataResponseDataListUsageData `json:"UsageData,omitempty" xml:"UsageData,omitempty" require:"true" type:"Repeated"`
}

func (s GetCoverRateDataResponseDataList) String() string {
	return tea.Prettify(s)
}

func (s GetCoverRateDataResponseDataList) GoString() string {
	return s.String()
}

func (s *GetCoverRateDataResponseDataList) SetUsageData(v []*GetCoverRateDataResponseDataListUsageData) *GetCoverRateDataResponseDataList {
	s.UsageData = v
	return s
}

type GetCoverRateDataResponseDataListUsageData struct {
	Date   *string                                          `json:"Date,omitempty" xml:"Date,omitempty" require:"true"`
	Values *GetCoverRateDataResponseDataListUsageDataValues `json:"Values,omitempty" xml:"Values,omitempty" require:"true" type:"Struct"`
}

func (s GetCoverRateDataResponseDataListUsageData) String() string {
	return tea.Prettify(s)
}

func (s GetCoverRateDataResponseDataListUsageData) GoString() string {
	return s.String()
}

func (s *GetCoverRateDataResponseDataListUsageData) SetDate(v string) *GetCoverRateDataResponseDataListUsageData {
	s.Date = &v
	return s
}

func (s *GetCoverRateDataResponseDataListUsageData) SetValues(v *GetCoverRateDataResponseDataListUsageDataValues) *GetCoverRateDataResponseDataListUsageData {
	s.Values = v
	return s
}

type GetCoverRateDataResponseDataListUsageDataValues struct {
	// Values
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" require:"true" type:"Repeated"`
}

func (s GetCoverRateDataResponseDataListUsageDataValues) String() string {
	return tea.Prettify(s)
}

func (s GetCoverRateDataResponseDataListUsageDataValues) GoString() string {
	return s.String()
}

func (s *GetCoverRateDataResponseDataListUsageDataValues) SetValues(v []*string) *GetCoverRateDataResponseDataListUsageDataValues {
	s.Values = v
	return s
}

type GetCoverRateDataResponseLabels struct {
	Label []*string `json:"Label,omitempty" xml:"Label,omitempty" require:"true" type:"Repeated"`
}

func (s GetCoverRateDataResponseLabels) String() string {
	return tea.Prettify(s)
}

func (s GetCoverRateDataResponseLabels) GoString() string {
	return s.String()
}

func (s *GetCoverRateDataResponseLabels) SetLabel(v []*string) *GetCoverRateDataResponseLabels {
	s.Label = v
	return s
}

type GetFeeHistoryRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	Page          *string `json:"Page,omitempty" xml:"Page,omitempty" require:"true"`
	PageSize      *string `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
}

func (s GetFeeHistoryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFeeHistoryRequest) GoString() string {
	return s.String()
}

func (s *GetFeeHistoryRequest) SetSecurityToken(v string) *GetFeeHistoryRequest {
	s.SecurityToken = &v
	return s
}

func (s *GetFeeHistoryRequest) SetVersion(v string) *GetFeeHistoryRequest {
	s.Version = &v
	return s
}

func (s *GetFeeHistoryRequest) SetPage(v string) *GetFeeHistoryRequest {
	s.Page = &v
	return s
}

func (s *GetFeeHistoryRequest) SetPageSize(v string) *GetFeeHistoryRequest {
	s.PageSize = &v
	return s
}

type GetFeeHistoryResponse struct {
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code      *int                          `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	FeeList   *GetFeeHistoryResponseFeeList `json:"FeeList,omitempty" xml:"FeeList,omitempty" require:"true" type:"Struct"`
	Pager     *GetFeeHistoryResponsePager   `json:"Pager,omitempty" xml:"Pager,omitempty" require:"true" type:"Struct"`
}

func (s GetFeeHistoryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFeeHistoryResponse) GoString() string {
	return s.String()
}

func (s *GetFeeHistoryResponse) SetRequestId(v string) *GetFeeHistoryResponse {
	s.RequestId = &v
	return s
}

func (s *GetFeeHistoryResponse) SetCode(v int) *GetFeeHistoryResponse {
	s.Code = &v
	return s
}

func (s *GetFeeHistoryResponse) SetFeeList(v *GetFeeHistoryResponseFeeList) *GetFeeHistoryResponse {
	s.FeeList = v
	return s
}

func (s *GetFeeHistoryResponse) SetPager(v *GetFeeHistoryResponsePager) *GetFeeHistoryResponse {
	s.Pager = v
	return s
}

type GetFeeHistoryResponseFeeList struct {
	Fee []*GetFeeHistoryResponseFeeListFee `json:"Fee,omitempty" xml:"Fee,omitempty" require:"true" type:"Repeated"`
}

func (s GetFeeHistoryResponseFeeList) String() string {
	return tea.Prettify(s)
}

func (s GetFeeHistoryResponseFeeList) GoString() string {
	return s.String()
}

func (s *GetFeeHistoryResponseFeeList) SetFee(v []*GetFeeHistoryResponseFeeListFee) *GetFeeHistoryResponseFeeList {
	s.Fee = v
	return s
}

type GetFeeHistoryResponseFeeListFee struct {
	Date                *string `json:"Date,omitempty" xml:"Date,omitempty" require:"true"`
	Mode                *string `json:"Mode,omitempty" xml:"Mode,omitempty" require:"true"`
	TotalBandwidth      *int    `json:"TotalBandwidth,omitempty" xml:"TotalBandwidth,omitempty" require:"true"`
	LevelTwoBandwidth   *int    `json:"LevelTwoBandwidth,omitempty" xml:"LevelTwoBandwidth,omitempty" require:"true"`
	LevelThreeBandwidth *int    `json:"LevelThreeBandwidth,omitempty" xml:"LevelThreeBandwidth,omitempty" require:"true"`
	TotalTraffic        *int    `json:"TotalTraffic,omitempty" xml:"TotalTraffic,omitempty" require:"true"`
	LevelTwoTraffic     *int    `json:"LevelTwoTraffic,omitempty" xml:"LevelTwoTraffic,omitempty" require:"true"`
	LevelThreeTraffic   *int    `json:"LevelThreeTraffic,omitempty" xml:"LevelThreeTraffic,omitempty" require:"true"`
	TimeSpan            *string `json:"TimeSpan,omitempty" xml:"TimeSpan,omitempty" require:"true"`
	BusinessType        *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty" require:"true"`
	StartDate           *string `json:"StartDate,omitempty" xml:"StartDate,omitempty" require:"true"`
	EndDate             *string `json:"EndDate,omitempty" xml:"EndDate,omitempty" require:"true"`
	ResourceId          *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" require:"true"`
	FlowOut             *int    `json:"FlowOut,omitempty" xml:"FlowOut,omitempty" require:"true"`
}

func (s GetFeeHistoryResponseFeeListFee) String() string {
	return tea.Prettify(s)
}

func (s GetFeeHistoryResponseFeeListFee) GoString() string {
	return s.String()
}

func (s *GetFeeHistoryResponseFeeListFee) SetDate(v string) *GetFeeHistoryResponseFeeListFee {
	s.Date = &v
	return s
}

func (s *GetFeeHistoryResponseFeeListFee) SetMode(v string) *GetFeeHistoryResponseFeeListFee {
	s.Mode = &v
	return s
}

func (s *GetFeeHistoryResponseFeeListFee) SetTotalBandwidth(v int) *GetFeeHistoryResponseFeeListFee {
	s.TotalBandwidth = &v
	return s
}

func (s *GetFeeHistoryResponseFeeListFee) SetLevelTwoBandwidth(v int) *GetFeeHistoryResponseFeeListFee {
	s.LevelTwoBandwidth = &v
	return s
}

func (s *GetFeeHistoryResponseFeeListFee) SetLevelThreeBandwidth(v int) *GetFeeHistoryResponseFeeListFee {
	s.LevelThreeBandwidth = &v
	return s
}

func (s *GetFeeHistoryResponseFeeListFee) SetTotalTraffic(v int) *GetFeeHistoryResponseFeeListFee {
	s.TotalTraffic = &v
	return s
}

func (s *GetFeeHistoryResponseFeeListFee) SetLevelTwoTraffic(v int) *GetFeeHistoryResponseFeeListFee {
	s.LevelTwoTraffic = &v
	return s
}

func (s *GetFeeHistoryResponseFeeListFee) SetLevelThreeTraffic(v int) *GetFeeHistoryResponseFeeListFee {
	s.LevelThreeTraffic = &v
	return s
}

func (s *GetFeeHistoryResponseFeeListFee) SetTimeSpan(v string) *GetFeeHistoryResponseFeeListFee {
	s.TimeSpan = &v
	return s
}

func (s *GetFeeHistoryResponseFeeListFee) SetBusinessType(v string) *GetFeeHistoryResponseFeeListFee {
	s.BusinessType = &v
	return s
}

func (s *GetFeeHistoryResponseFeeListFee) SetStartDate(v string) *GetFeeHistoryResponseFeeListFee {
	s.StartDate = &v
	return s
}

func (s *GetFeeHistoryResponseFeeListFee) SetEndDate(v string) *GetFeeHistoryResponseFeeListFee {
	s.EndDate = &v
	return s
}

func (s *GetFeeHistoryResponseFeeListFee) SetResourceId(v string) *GetFeeHistoryResponseFeeListFee {
	s.ResourceId = &v
	return s
}

func (s *GetFeeHistoryResponseFeeListFee) SetFlowOut(v int) *GetFeeHistoryResponseFeeListFee {
	s.FlowOut = &v
	return s
}

type GetFeeHistoryResponsePager struct {
	Page     *int `json:"Page,omitempty" xml:"Page,omitempty" require:"true"`
	Total    *int `json:"Total,omitempty" xml:"Total,omitempty" require:"true"`
	PageSize *int `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
}

func (s GetFeeHistoryResponsePager) String() string {
	return tea.Prettify(s)
}

func (s GetFeeHistoryResponsePager) GoString() string {
	return s.String()
}

func (s *GetFeeHistoryResponsePager) SetPage(v int) *GetFeeHistoryResponsePager {
	s.Page = &v
	return s
}

func (s *GetFeeHistoryResponsePager) SetTotal(v int) *GetFeeHistoryResponsePager {
	s.Total = &v
	return s
}

func (s *GetFeeHistoryResponsePager) SetPageSize(v int) *GetFeeHistoryResponsePager {
	s.PageSize = &v
	return s
}

type GetExpenseSummaryRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	StartDate     *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	EndDate       *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	Domain        *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Region        *string `json:"Region,omitempty" xml:"Region,omitempty"`
	IspName       *string `json:"IspName,omitempty" xml:"IspName,omitempty"`
	PlatformType  *string `json:"PlatformType,omitempty" xml:"PlatformType,omitempty"`
	BusinessType  *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
}

func (s GetExpenseSummaryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetExpenseSummaryRequest) GoString() string {
	return s.String()
}

func (s *GetExpenseSummaryRequest) SetSecurityToken(v string) *GetExpenseSummaryRequest {
	s.SecurityToken = &v
	return s
}

func (s *GetExpenseSummaryRequest) SetVersion(v string) *GetExpenseSummaryRequest {
	s.Version = &v
	return s
}

func (s *GetExpenseSummaryRequest) SetStartDate(v string) *GetExpenseSummaryRequest {
	s.StartDate = &v
	return s
}

func (s *GetExpenseSummaryRequest) SetEndDate(v string) *GetExpenseSummaryRequest {
	s.EndDate = &v
	return s
}

func (s *GetExpenseSummaryRequest) SetDomain(v string) *GetExpenseSummaryRequest {
	s.Domain = &v
	return s
}

func (s *GetExpenseSummaryRequest) SetRegion(v string) *GetExpenseSummaryRequest {
	s.Region = &v
	return s
}

func (s *GetExpenseSummaryRequest) SetIspName(v string) *GetExpenseSummaryRequest {
	s.IspName = &v
	return s
}

func (s *GetExpenseSummaryRequest) SetPlatformType(v string) *GetExpenseSummaryRequest {
	s.PlatformType = &v
	return s
}

func (s *GetExpenseSummaryRequest) SetBusinessType(v string) *GetExpenseSummaryRequest {
	s.BusinessType = &v
	return s
}

type GetExpenseSummaryResponse struct {
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code      *int                           `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Data      *GetExpenseSummaryResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s GetExpenseSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetExpenseSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetExpenseSummaryResponse) SetRequestId(v string) *GetExpenseSummaryResponse {
	s.RequestId = &v
	return s
}

func (s *GetExpenseSummaryResponse) SetCode(v int) *GetExpenseSummaryResponse {
	s.Code = &v
	return s
}

func (s *GetExpenseSummaryResponse) SetData(v *GetExpenseSummaryResponseData) *GetExpenseSummaryResponse {
	s.Data = v
	return s
}

type GetExpenseSummaryResponseData struct {
	TotalTraffic    *int64   `json:"TotalTraffic,omitempty" xml:"TotalTraffic,omitempty" require:"true"`
	TotalUV         *int     `json:"TotalUV,omitempty" xml:"TotalUV,omitempty" require:"true"`
	ShareRate       *float32 `json:"ShareRate,omitempty" xml:"ShareRate,omitempty" require:"true"`
	CoverRate       *float32 `json:"CoverRate,omitempty" xml:"CoverRate,omitempty" require:"true"`
	ForecastFluency *float32 `json:"ForecastFluency,omitempty" xml:"ForecastFluency,omitempty" require:"true"`
	TopBandwidth    *int64   `json:"TopBandwidth,omitempty" xml:"TopBandwidth,omitempty" require:"true"`
}

func (s GetExpenseSummaryResponseData) String() string {
	return tea.Prettify(s)
}

func (s GetExpenseSummaryResponseData) GoString() string {
	return s.String()
}

func (s *GetExpenseSummaryResponseData) SetTotalTraffic(v int64) *GetExpenseSummaryResponseData {
	s.TotalTraffic = &v
	return s
}

func (s *GetExpenseSummaryResponseData) SetTotalUV(v int) *GetExpenseSummaryResponseData {
	s.TotalUV = &v
	return s
}

func (s *GetExpenseSummaryResponseData) SetShareRate(v float32) *GetExpenseSummaryResponseData {
	s.ShareRate = &v
	return s
}

func (s *GetExpenseSummaryResponseData) SetCoverRate(v float32) *GetExpenseSummaryResponseData {
	s.CoverRate = &v
	return s
}

func (s *GetExpenseSummaryResponseData) SetForecastFluency(v float32) *GetExpenseSummaryResponseData {
	s.ForecastFluency = &v
	return s
}

func (s *GetExpenseSummaryResponseData) SetTopBandwidth(v int64) *GetExpenseSummaryResponseData {
	s.TopBandwidth = &v
	return s
}

type GetDomainsRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	Page          *string `json:"Page,omitempty" xml:"Page,omitempty" require:"true"`
	PageSize      *string `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	Domain        *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s GetDomainsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDomainsRequest) GoString() string {
	return s.String()
}

func (s *GetDomainsRequest) SetSecurityToken(v string) *GetDomainsRequest {
	s.SecurityToken = &v
	return s
}

func (s *GetDomainsRequest) SetVersion(v string) *GetDomainsRequest {
	s.Version = &v
	return s
}

func (s *GetDomainsRequest) SetPage(v string) *GetDomainsRequest {
	s.Page = &v
	return s
}

func (s *GetDomainsRequest) SetPageSize(v string) *GetDomainsRequest {
	s.PageSize = &v
	return s
}

func (s *GetDomainsRequest) SetDomain(v string) *GetDomainsRequest {
	s.Domain = &v
	return s
}

type GetDomainsResponse struct {
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code      *int                        `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	DataList  *GetDomainsResponseDataList `json:"DataList,omitempty" xml:"DataList,omitempty" require:"true" type:"Struct"`
	Pager     *GetDomainsResponsePager    `json:"Pager,omitempty" xml:"Pager,omitempty" require:"true" type:"Struct"`
}

func (s GetDomainsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDomainsResponse) GoString() string {
	return s.String()
}

func (s *GetDomainsResponse) SetRequestId(v string) *GetDomainsResponse {
	s.RequestId = &v
	return s
}

func (s *GetDomainsResponse) SetCode(v int) *GetDomainsResponse {
	s.Code = &v
	return s
}

func (s *GetDomainsResponse) SetDataList(v *GetDomainsResponseDataList) *GetDomainsResponse {
	s.DataList = v
	return s
}

func (s *GetDomainsResponse) SetPager(v *GetDomainsResponsePager) *GetDomainsResponse {
	s.Pager = v
	return s
}

type GetDomainsResponseDataList struct {
	UsageData []*GetDomainsResponseDataListUsageData `json:"UsageData,omitempty" xml:"UsageData,omitempty" require:"true" type:"Repeated"`
}

func (s GetDomainsResponseDataList) String() string {
	return tea.Prettify(s)
}

func (s GetDomainsResponseDataList) GoString() string {
	return s.String()
}

func (s *GetDomainsResponseDataList) SetUsageData(v []*GetDomainsResponseDataListUsageData) *GetDomainsResponseDataList {
	s.UsageData = v
	return s
}

type GetDomainsResponseDataListUsageData struct {
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" require:"true"`
	Domain       *string `json:"Domain,omitempty" xml:"Domain,omitempty" require:"true"`
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty" require:"true"`
	Status       *bool   `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	CreatedAt    *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty" require:"true"`
	UpdatedAt    *string `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty" require:"true"`
	SliceFormat  *string `json:"SliceFormat,omitempty" xml:"SliceFormat,omitempty" require:"true"`
}

func (s GetDomainsResponseDataListUsageData) String() string {
	return tea.Prettify(s)
}

func (s GetDomainsResponseDataListUsageData) GoString() string {
	return s.String()
}

func (s *GetDomainsResponseDataListUsageData) SetResourceId(v string) *GetDomainsResponseDataListUsageData {
	s.ResourceId = &v
	return s
}

func (s *GetDomainsResponseDataListUsageData) SetDomain(v string) *GetDomainsResponseDataListUsageData {
	s.Domain = &v
	return s
}

func (s *GetDomainsResponseDataListUsageData) SetBusinessType(v string) *GetDomainsResponseDataListUsageData {
	s.BusinessType = &v
	return s
}

func (s *GetDomainsResponseDataListUsageData) SetStatus(v bool) *GetDomainsResponseDataListUsageData {
	s.Status = &v
	return s
}

func (s *GetDomainsResponseDataListUsageData) SetCreatedAt(v string) *GetDomainsResponseDataListUsageData {
	s.CreatedAt = &v
	return s
}

func (s *GetDomainsResponseDataListUsageData) SetUpdatedAt(v string) *GetDomainsResponseDataListUsageData {
	s.UpdatedAt = &v
	return s
}

func (s *GetDomainsResponseDataListUsageData) SetSliceFormat(v string) *GetDomainsResponseDataListUsageData {
	s.SliceFormat = &v
	return s
}

type GetDomainsResponsePager struct {
	Page     *int `json:"Page,omitempty" xml:"Page,omitempty" require:"true"`
	Total    *int `json:"Total,omitempty" xml:"Total,omitempty" require:"true"`
	PageSize *int `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
}

func (s GetDomainsResponsePager) String() string {
	return tea.Prettify(s)
}

func (s GetDomainsResponsePager) GoString() string {
	return s.String()
}

func (s *GetDomainsResponsePager) SetPage(v int) *GetDomainsResponsePager {
	s.Page = &v
	return s
}

func (s *GetDomainsResponsePager) SetTotal(v int) *GetDomainsResponsePager {
	s.Total = &v
	return s
}

func (s *GetDomainsResponsePager) SetPageSize(v int) *GetDomainsResponsePager {
	s.PageSize = &v
	return s
}

type GetLogsListRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	Domain        *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Date          *string `json:"Date,omitempty" xml:"Date,omitempty"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s GetLogsListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLogsListRequest) GoString() string {
	return s.String()
}

func (s *GetLogsListRequest) SetSecurityToken(v string) *GetLogsListRequest {
	s.SecurityToken = &v
	return s
}

func (s *GetLogsListRequest) SetVersion(v string) *GetLogsListRequest {
	s.Version = &v
	return s
}

func (s *GetLogsListRequest) SetDomain(v string) *GetLogsListRequest {
	s.Domain = &v
	return s
}

func (s *GetLogsListRequest) SetDate(v string) *GetLogsListRequest {
	s.Date = &v
	return s
}

func (s *GetLogsListRequest) SetStartTime(v string) *GetLogsListRequest {
	s.StartTime = &v
	return s
}

func (s *GetLogsListRequest) SetEndTime(v string) *GetLogsListRequest {
	s.EndTime = &v
	return s
}

type GetLogsListResponse struct {
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code      *int                        `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	LogList   *GetLogsListResponseLogList `json:"LogList,omitempty" xml:"LogList,omitempty" require:"true" type:"Struct"`
}

func (s GetLogsListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLogsListResponse) GoString() string {
	return s.String()
}

func (s *GetLogsListResponse) SetRequestId(v string) *GetLogsListResponse {
	s.RequestId = &v
	return s
}

func (s *GetLogsListResponse) SetCode(v int) *GetLogsListResponse {
	s.Code = &v
	return s
}

func (s *GetLogsListResponse) SetLogList(v *GetLogsListResponseLogList) *GetLogsListResponse {
	s.LogList = v
	return s
}

type GetLogsListResponseLogList struct {
	Log []*GetLogsListResponseLogListLog `json:"Log,omitempty" xml:"Log,omitempty" require:"true" type:"Repeated"`
}

func (s GetLogsListResponseLogList) String() string {
	return tea.Prettify(s)
}

func (s GetLogsListResponseLogList) GoString() string {
	return s.String()
}

func (s *GetLogsListResponseLogList) SetLog(v []*GetLogsListResponseLogListLog) *GetLogsListResponseLogList {
	s.Log = v
	return s
}

type GetLogsListResponseLogListLog struct {
	Url       *string `json:"Url,omitempty" xml:"Url,omitempty" require:"true"`
	FileName  *string `json:"FileName,omitempty" xml:"FileName,omitempty" require:"true"`
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty" require:"true"`
	EndDate   *string `json:"EndDate,omitempty" xml:"EndDate,omitempty" require:"true"`
}

func (s GetLogsListResponseLogListLog) String() string {
	return tea.Prettify(s)
}

func (s GetLogsListResponseLogListLog) GoString() string {
	return s.String()
}

func (s *GetLogsListResponseLogListLog) SetUrl(v string) *GetLogsListResponseLogListLog {
	s.Url = &v
	return s
}

func (s *GetLogsListResponseLogListLog) SetFileName(v string) *GetLogsListResponseLogListLog {
	s.FileName = &v
	return s
}

func (s *GetLogsListResponseLogListLog) SetStartDate(v string) *GetLogsListResponseLogListLog {
	s.StartDate = &v
	return s
}

func (s *GetLogsListResponseLogListLog) SetEndDate(v string) *GetLogsListResponseLogListLog {
	s.EndDate = &v
	return s
}

type GetFluencyDataRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	Domain        *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Region        *string `json:"Region,omitempty" xml:"Region,omitempty" require:"true"`
	IspName       *string `json:"IspName,omitempty" xml:"IspName,omitempty" require:"true"`
	PlatformType  *string `json:"PlatformType,omitempty" xml:"PlatformType,omitempty" require:"true"`
	BusinessType  *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty" require:"true"`
	StartDate     *string `json:"StartDate,omitempty" xml:"StartDate,omitempty" require:"true"`
	EndDate       *string `json:"EndDate,omitempty" xml:"EndDate,omitempty" require:"true"`
}

func (s GetFluencyDataRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFluencyDataRequest) GoString() string {
	return s.String()
}

func (s *GetFluencyDataRequest) SetSecurityToken(v string) *GetFluencyDataRequest {
	s.SecurityToken = &v
	return s
}

func (s *GetFluencyDataRequest) SetVersion(v string) *GetFluencyDataRequest {
	s.Version = &v
	return s
}

func (s *GetFluencyDataRequest) SetDomain(v string) *GetFluencyDataRequest {
	s.Domain = &v
	return s
}

func (s *GetFluencyDataRequest) SetRegion(v string) *GetFluencyDataRequest {
	s.Region = &v
	return s
}

func (s *GetFluencyDataRequest) SetIspName(v string) *GetFluencyDataRequest {
	s.IspName = &v
	return s
}

func (s *GetFluencyDataRequest) SetPlatformType(v string) *GetFluencyDataRequest {
	s.PlatformType = &v
	return s
}

func (s *GetFluencyDataRequest) SetBusinessType(v string) *GetFluencyDataRequest {
	s.BusinessType = &v
	return s
}

func (s *GetFluencyDataRequest) SetStartDate(v string) *GetFluencyDataRequest {
	s.StartDate = &v
	return s
}

func (s *GetFluencyDataRequest) SetEndDate(v string) *GetFluencyDataRequest {
	s.EndDate = &v
	return s
}

type GetFluencyDataResponse struct {
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code      *int                            `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	DataList  *GetFluencyDataResponseDataList `json:"DataList,omitempty" xml:"DataList,omitempty" require:"true" type:"Struct"`
	Labels    *GetFluencyDataResponseLabels   `json:"Labels,omitempty" xml:"Labels,omitempty" require:"true" type:"Struct"`
}

func (s GetFluencyDataResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFluencyDataResponse) GoString() string {
	return s.String()
}

func (s *GetFluencyDataResponse) SetRequestId(v string) *GetFluencyDataResponse {
	s.RequestId = &v
	return s
}

func (s *GetFluencyDataResponse) SetCode(v int) *GetFluencyDataResponse {
	s.Code = &v
	return s
}

func (s *GetFluencyDataResponse) SetDataList(v *GetFluencyDataResponseDataList) *GetFluencyDataResponse {
	s.DataList = v
	return s
}

func (s *GetFluencyDataResponse) SetLabels(v *GetFluencyDataResponseLabels) *GetFluencyDataResponse {
	s.Labels = v
	return s
}

type GetFluencyDataResponseDataList struct {
	UsageData []*GetFluencyDataResponseDataListUsageData `json:"UsageData,omitempty" xml:"UsageData,omitempty" require:"true" type:"Repeated"`
}

func (s GetFluencyDataResponseDataList) String() string {
	return tea.Prettify(s)
}

func (s GetFluencyDataResponseDataList) GoString() string {
	return s.String()
}

func (s *GetFluencyDataResponseDataList) SetUsageData(v []*GetFluencyDataResponseDataListUsageData) *GetFluencyDataResponseDataList {
	s.UsageData = v
	return s
}

type GetFluencyDataResponseDataListUsageData struct {
	Date   *string                                        `json:"Date,omitempty" xml:"Date,omitempty" require:"true"`
	Values *GetFluencyDataResponseDataListUsageDataValues `json:"Values,omitempty" xml:"Values,omitempty" require:"true" type:"Struct"`
}

func (s GetFluencyDataResponseDataListUsageData) String() string {
	return tea.Prettify(s)
}

func (s GetFluencyDataResponseDataListUsageData) GoString() string {
	return s.String()
}

func (s *GetFluencyDataResponseDataListUsageData) SetDate(v string) *GetFluencyDataResponseDataListUsageData {
	s.Date = &v
	return s
}

func (s *GetFluencyDataResponseDataListUsageData) SetValues(v *GetFluencyDataResponseDataListUsageDataValues) *GetFluencyDataResponseDataListUsageData {
	s.Values = v
	return s
}

type GetFluencyDataResponseDataListUsageDataValues struct {
	// Values
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" require:"true" type:"Repeated"`
}

func (s GetFluencyDataResponseDataListUsageDataValues) String() string {
	return tea.Prettify(s)
}

func (s GetFluencyDataResponseDataListUsageDataValues) GoString() string {
	return s.String()
}

func (s *GetFluencyDataResponseDataListUsageDataValues) SetValues(v []*string) *GetFluencyDataResponseDataListUsageDataValues {
	s.Values = v
	return s
}

type GetFluencyDataResponseLabels struct {
	Label []*string `json:"Label,omitempty" xml:"Label,omitempty" require:"true" type:"Repeated"`
}

func (s GetFluencyDataResponseLabels) String() string {
	return tea.Prettify(s)
}

func (s GetFluencyDataResponseLabels) GoString() string {
	return s.String()
}

func (s *GetFluencyDataResponseLabels) SetLabel(v []*string) *GetFluencyDataResponseLabels {
	s.Label = v
	return s
}

type GetFirstFrameDelayDataRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	Domain        *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Region        *string `json:"Region,omitempty" xml:"Region,omitempty" require:"true"`
	IspName       *string `json:"IspName,omitempty" xml:"IspName,omitempty" require:"true"`
	PlatformType  *string `json:"PlatformType,omitempty" xml:"PlatformType,omitempty" require:"true"`
	BusinessType  *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty" require:"true"`
	StartDate     *string `json:"StartDate,omitempty" xml:"StartDate,omitempty" require:"true"`
	EndDate       *string `json:"EndDate,omitempty" xml:"EndDate,omitempty" require:"true"`
}

func (s GetFirstFrameDelayDataRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFirstFrameDelayDataRequest) GoString() string {
	return s.String()
}

func (s *GetFirstFrameDelayDataRequest) SetSecurityToken(v string) *GetFirstFrameDelayDataRequest {
	s.SecurityToken = &v
	return s
}

func (s *GetFirstFrameDelayDataRequest) SetVersion(v string) *GetFirstFrameDelayDataRequest {
	s.Version = &v
	return s
}

func (s *GetFirstFrameDelayDataRequest) SetDomain(v string) *GetFirstFrameDelayDataRequest {
	s.Domain = &v
	return s
}

func (s *GetFirstFrameDelayDataRequest) SetRegion(v string) *GetFirstFrameDelayDataRequest {
	s.Region = &v
	return s
}

func (s *GetFirstFrameDelayDataRequest) SetIspName(v string) *GetFirstFrameDelayDataRequest {
	s.IspName = &v
	return s
}

func (s *GetFirstFrameDelayDataRequest) SetPlatformType(v string) *GetFirstFrameDelayDataRequest {
	s.PlatformType = &v
	return s
}

func (s *GetFirstFrameDelayDataRequest) SetBusinessType(v string) *GetFirstFrameDelayDataRequest {
	s.BusinessType = &v
	return s
}

func (s *GetFirstFrameDelayDataRequest) SetStartDate(v string) *GetFirstFrameDelayDataRequest {
	s.StartDate = &v
	return s
}

func (s *GetFirstFrameDelayDataRequest) SetEndDate(v string) *GetFirstFrameDelayDataRequest {
	s.EndDate = &v
	return s
}

type GetFirstFrameDelayDataResponse struct {
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code      *int                                    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	DataList  *GetFirstFrameDelayDataResponseDataList `json:"DataList,omitempty" xml:"DataList,omitempty" require:"true" type:"Struct"`
	Labels    *GetFirstFrameDelayDataResponseLabels   `json:"Labels,omitempty" xml:"Labels,omitempty" require:"true" type:"Struct"`
}

func (s GetFirstFrameDelayDataResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFirstFrameDelayDataResponse) GoString() string {
	return s.String()
}

func (s *GetFirstFrameDelayDataResponse) SetRequestId(v string) *GetFirstFrameDelayDataResponse {
	s.RequestId = &v
	return s
}

func (s *GetFirstFrameDelayDataResponse) SetCode(v int) *GetFirstFrameDelayDataResponse {
	s.Code = &v
	return s
}

func (s *GetFirstFrameDelayDataResponse) SetDataList(v *GetFirstFrameDelayDataResponseDataList) *GetFirstFrameDelayDataResponse {
	s.DataList = v
	return s
}

func (s *GetFirstFrameDelayDataResponse) SetLabels(v *GetFirstFrameDelayDataResponseLabels) *GetFirstFrameDelayDataResponse {
	s.Labels = v
	return s
}

type GetFirstFrameDelayDataResponseDataList struct {
	UsageData []*GetFirstFrameDelayDataResponseDataListUsageData `json:"UsageData,omitempty" xml:"UsageData,omitempty" require:"true" type:"Repeated"`
}

func (s GetFirstFrameDelayDataResponseDataList) String() string {
	return tea.Prettify(s)
}

func (s GetFirstFrameDelayDataResponseDataList) GoString() string {
	return s.String()
}

func (s *GetFirstFrameDelayDataResponseDataList) SetUsageData(v []*GetFirstFrameDelayDataResponseDataListUsageData) *GetFirstFrameDelayDataResponseDataList {
	s.UsageData = v
	return s
}

type GetFirstFrameDelayDataResponseDataListUsageData struct {
	Date   *string                                                `json:"Date,omitempty" xml:"Date,omitempty" require:"true"`
	Values *GetFirstFrameDelayDataResponseDataListUsageDataValues `json:"Values,omitempty" xml:"Values,omitempty" require:"true" type:"Struct"`
}

func (s GetFirstFrameDelayDataResponseDataListUsageData) String() string {
	return tea.Prettify(s)
}

func (s GetFirstFrameDelayDataResponseDataListUsageData) GoString() string {
	return s.String()
}

func (s *GetFirstFrameDelayDataResponseDataListUsageData) SetDate(v string) *GetFirstFrameDelayDataResponseDataListUsageData {
	s.Date = &v
	return s
}

func (s *GetFirstFrameDelayDataResponseDataListUsageData) SetValues(v *GetFirstFrameDelayDataResponseDataListUsageDataValues) *GetFirstFrameDelayDataResponseDataListUsageData {
	s.Values = v
	return s
}

type GetFirstFrameDelayDataResponseDataListUsageDataValues struct {
	// Values
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" require:"true" type:"Repeated"`
}

func (s GetFirstFrameDelayDataResponseDataListUsageDataValues) String() string {
	return tea.Prettify(s)
}

func (s GetFirstFrameDelayDataResponseDataListUsageDataValues) GoString() string {
	return s.String()
}

func (s *GetFirstFrameDelayDataResponseDataListUsageDataValues) SetValues(v []*string) *GetFirstFrameDelayDataResponseDataListUsageDataValues {
	s.Values = v
	return s
}

type GetFirstFrameDelayDataResponseLabels struct {
	Label []*string `json:"Label,omitempty" xml:"Label,omitempty" require:"true" type:"Repeated"`
}

func (s GetFirstFrameDelayDataResponseLabels) String() string {
	return tea.Prettify(s)
}

func (s GetFirstFrameDelayDataResponseLabels) GoString() string {
	return s.String()
}

func (s *GetFirstFrameDelayDataResponseLabels) SetLabel(v []*string) *GetFirstFrameDelayDataResponseLabels {
	s.Label = v
	return s
}

type GetTokenListRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
}

func (s GetTokenListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTokenListRequest) GoString() string {
	return s.String()
}

func (s *GetTokenListRequest) SetSecurityToken(v string) *GetTokenListRequest {
	s.SecurityToken = &v
	return s
}

func (s *GetTokenListRequest) SetVersion(v string) *GetTokenListRequest {
	s.Version = &v
	return s
}

type GetTokenListResponse struct {
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code      *int                           `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	TokenList *GetTokenListResponseTokenList `json:"TokenList,omitempty" xml:"TokenList,omitempty" require:"true" type:"Struct"`
}

func (s GetTokenListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTokenListResponse) GoString() string {
	return s.String()
}

func (s *GetTokenListResponse) SetRequestId(v string) *GetTokenListResponse {
	s.RequestId = &v
	return s
}

func (s *GetTokenListResponse) SetCode(v int) *GetTokenListResponse {
	s.Code = &v
	return s
}

func (s *GetTokenListResponse) SetTokenList(v *GetTokenListResponseTokenList) *GetTokenListResponse {
	s.TokenList = v
	return s
}

type GetTokenListResponseTokenList struct {
	Token []*GetTokenListResponseTokenListToken `json:"Token,omitempty" xml:"Token,omitempty" require:"true" type:"Repeated"`
}

func (s GetTokenListResponseTokenList) String() string {
	return tea.Prettify(s)
}

func (s GetTokenListResponseTokenList) GoString() string {
	return s.String()
}

func (s *GetTokenListResponseTokenList) SetToken(v []*GetTokenListResponseTokenListToken) *GetTokenListResponseTokenList {
	s.Token = v
	return s
}

type GetTokenListResponseTokenListToken struct {
	ClientId     *string `json:"ClientId,omitempty" xml:"ClientId,omitempty" require:"true"`
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" require:"true"`
	PlatformName *string `json:"PlatformName,omitempty" xml:"PlatformName,omitempty" require:"true"`
	PlatformType *string `json:"PlatformType,omitempty" xml:"PlatformType,omitempty" require:"true"`
	Token        *string `json:"Token,omitempty" xml:"Token,omitempty" require:"true"`
	CreatedAt    *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty" require:"true"`
	UpdatedAt    *string `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty" require:"true"`
}

func (s GetTokenListResponseTokenListToken) String() string {
	return tea.Prettify(s)
}

func (s GetTokenListResponseTokenListToken) GoString() string {
	return s.String()
}

func (s *GetTokenListResponseTokenListToken) SetClientId(v string) *GetTokenListResponseTokenListToken {
	s.ClientId = &v
	return s
}

func (s *GetTokenListResponseTokenListToken) SetResourceId(v string) *GetTokenListResponseTokenListToken {
	s.ResourceId = &v
	return s
}

func (s *GetTokenListResponseTokenListToken) SetPlatformName(v string) *GetTokenListResponseTokenListToken {
	s.PlatformName = &v
	return s
}

func (s *GetTokenListResponseTokenListToken) SetPlatformType(v string) *GetTokenListResponseTokenListToken {
	s.PlatformType = &v
	return s
}

func (s *GetTokenListResponseTokenListToken) SetToken(v string) *GetTokenListResponseTokenListToken {
	s.Token = &v
	return s
}

func (s *GetTokenListResponseTokenListToken) SetCreatedAt(v string) *GetTokenListResponseTokenListToken {
	s.CreatedAt = &v
	return s
}

func (s *GetTokenListResponseTokenListToken) SetUpdatedAt(v string) *GetTokenListResponseTokenListToken {
	s.UpdatedAt = &v
	return s
}

type GetShareRateDataRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	Domain        *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Region        *string `json:"Region,omitempty" xml:"Region,omitempty" require:"true"`
	IspName       *string `json:"IspName,omitempty" xml:"IspName,omitempty" require:"true"`
	PlatformType  *string `json:"PlatformType,omitempty" xml:"PlatformType,omitempty" require:"true"`
	BusinessType  *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty" require:"true"`
	StartDate     *string `json:"StartDate,omitempty" xml:"StartDate,omitempty" require:"true"`
	EndDate       *string `json:"EndDate,omitempty" xml:"EndDate,omitempty" require:"true"`
}

func (s GetShareRateDataRequest) String() string {
	return tea.Prettify(s)
}

func (s GetShareRateDataRequest) GoString() string {
	return s.String()
}

func (s *GetShareRateDataRequest) SetSecurityToken(v string) *GetShareRateDataRequest {
	s.SecurityToken = &v
	return s
}

func (s *GetShareRateDataRequest) SetVersion(v string) *GetShareRateDataRequest {
	s.Version = &v
	return s
}

func (s *GetShareRateDataRequest) SetDomain(v string) *GetShareRateDataRequest {
	s.Domain = &v
	return s
}

func (s *GetShareRateDataRequest) SetRegion(v string) *GetShareRateDataRequest {
	s.Region = &v
	return s
}

func (s *GetShareRateDataRequest) SetIspName(v string) *GetShareRateDataRequest {
	s.IspName = &v
	return s
}

func (s *GetShareRateDataRequest) SetPlatformType(v string) *GetShareRateDataRequest {
	s.PlatformType = &v
	return s
}

func (s *GetShareRateDataRequest) SetBusinessType(v string) *GetShareRateDataRequest {
	s.BusinessType = &v
	return s
}

func (s *GetShareRateDataRequest) SetStartDate(v string) *GetShareRateDataRequest {
	s.StartDate = &v
	return s
}

func (s *GetShareRateDataRequest) SetEndDate(v string) *GetShareRateDataRequest {
	s.EndDate = &v
	return s
}

type GetShareRateDataResponse struct {
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code      *int                              `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	DataList  *GetShareRateDataResponseDataList `json:"DataList,omitempty" xml:"DataList,omitempty" require:"true" type:"Struct"`
	Labels    *GetShareRateDataResponseLabels   `json:"Labels,omitempty" xml:"Labels,omitempty" require:"true" type:"Struct"`
}

func (s GetShareRateDataResponse) String() string {
	return tea.Prettify(s)
}

func (s GetShareRateDataResponse) GoString() string {
	return s.String()
}

func (s *GetShareRateDataResponse) SetRequestId(v string) *GetShareRateDataResponse {
	s.RequestId = &v
	return s
}

func (s *GetShareRateDataResponse) SetCode(v int) *GetShareRateDataResponse {
	s.Code = &v
	return s
}

func (s *GetShareRateDataResponse) SetDataList(v *GetShareRateDataResponseDataList) *GetShareRateDataResponse {
	s.DataList = v
	return s
}

func (s *GetShareRateDataResponse) SetLabels(v *GetShareRateDataResponseLabels) *GetShareRateDataResponse {
	s.Labels = v
	return s
}

type GetShareRateDataResponseDataList struct {
	UsageData []*GetShareRateDataResponseDataListUsageData `json:"UsageData,omitempty" xml:"UsageData,omitempty" require:"true" type:"Repeated"`
}

func (s GetShareRateDataResponseDataList) String() string {
	return tea.Prettify(s)
}

func (s GetShareRateDataResponseDataList) GoString() string {
	return s.String()
}

func (s *GetShareRateDataResponseDataList) SetUsageData(v []*GetShareRateDataResponseDataListUsageData) *GetShareRateDataResponseDataList {
	s.UsageData = v
	return s
}

type GetShareRateDataResponseDataListUsageData struct {
	Date   *string                                          `json:"Date,omitempty" xml:"Date,omitempty" require:"true"`
	Values *GetShareRateDataResponseDataListUsageDataValues `json:"Values,omitempty" xml:"Values,omitempty" require:"true" type:"Struct"`
}

func (s GetShareRateDataResponseDataListUsageData) String() string {
	return tea.Prettify(s)
}

func (s GetShareRateDataResponseDataListUsageData) GoString() string {
	return s.String()
}

func (s *GetShareRateDataResponseDataListUsageData) SetDate(v string) *GetShareRateDataResponseDataListUsageData {
	s.Date = &v
	return s
}

func (s *GetShareRateDataResponseDataListUsageData) SetValues(v *GetShareRateDataResponseDataListUsageDataValues) *GetShareRateDataResponseDataListUsageData {
	s.Values = v
	return s
}

type GetShareRateDataResponseDataListUsageDataValues struct {
	// Values
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" require:"true" type:"Repeated"`
}

func (s GetShareRateDataResponseDataListUsageDataValues) String() string {
	return tea.Prettify(s)
}

func (s GetShareRateDataResponseDataListUsageDataValues) GoString() string {
	return s.String()
}

func (s *GetShareRateDataResponseDataListUsageDataValues) SetValues(v []*string) *GetShareRateDataResponseDataListUsageDataValues {
	s.Values = v
	return s
}

type GetShareRateDataResponseLabels struct {
	Label []*string `json:"Label,omitempty" xml:"Label,omitempty" require:"true" type:"Repeated"`
}

func (s GetShareRateDataResponseLabels) String() string {
	return tea.Prettify(s)
}

func (s GetShareRateDataResponseLabels) GoString() string {
	return s.String()
}

func (s *GetShareRateDataResponseLabels) SetLabel(v []*string) *GetShareRateDataResponseLabels {
	s.Label = v
	return s
}

type GetTrafficDataRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	Domain        *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Region        *string `json:"Region,omitempty" xml:"Region,omitempty" require:"true"`
	IspName       *string `json:"IspName,omitempty" xml:"IspName,omitempty" require:"true"`
	PlatformType  *string `json:"PlatformType,omitempty" xml:"PlatformType,omitempty" require:"true"`
	BusinessType  *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty" require:"true"`
	StartDate     *string `json:"StartDate,omitempty" xml:"StartDate,omitempty" require:"true"`
	EndDate       *string `json:"EndDate,omitempty" xml:"EndDate,omitempty" require:"true"`
}

func (s GetTrafficDataRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTrafficDataRequest) GoString() string {
	return s.String()
}

func (s *GetTrafficDataRequest) SetSecurityToken(v string) *GetTrafficDataRequest {
	s.SecurityToken = &v
	return s
}

func (s *GetTrafficDataRequest) SetVersion(v string) *GetTrafficDataRequest {
	s.Version = &v
	return s
}

func (s *GetTrafficDataRequest) SetDomain(v string) *GetTrafficDataRequest {
	s.Domain = &v
	return s
}

func (s *GetTrafficDataRequest) SetRegion(v string) *GetTrafficDataRequest {
	s.Region = &v
	return s
}

func (s *GetTrafficDataRequest) SetIspName(v string) *GetTrafficDataRequest {
	s.IspName = &v
	return s
}

func (s *GetTrafficDataRequest) SetPlatformType(v string) *GetTrafficDataRequest {
	s.PlatformType = &v
	return s
}

func (s *GetTrafficDataRequest) SetBusinessType(v string) *GetTrafficDataRequest {
	s.BusinessType = &v
	return s
}

func (s *GetTrafficDataRequest) SetStartDate(v string) *GetTrafficDataRequest {
	s.StartDate = &v
	return s
}

func (s *GetTrafficDataRequest) SetEndDate(v string) *GetTrafficDataRequest {
	s.EndDate = &v
	return s
}

type GetTrafficDataResponse struct {
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code      *int                            `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	DataList  *GetTrafficDataResponseDataList `json:"DataList,omitempty" xml:"DataList,omitempty" require:"true" type:"Struct"`
	Labels    *GetTrafficDataResponseLabels   `json:"Labels,omitempty" xml:"Labels,omitempty" require:"true" type:"Struct"`
}

func (s GetTrafficDataResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTrafficDataResponse) GoString() string {
	return s.String()
}

func (s *GetTrafficDataResponse) SetRequestId(v string) *GetTrafficDataResponse {
	s.RequestId = &v
	return s
}

func (s *GetTrafficDataResponse) SetCode(v int) *GetTrafficDataResponse {
	s.Code = &v
	return s
}

func (s *GetTrafficDataResponse) SetDataList(v *GetTrafficDataResponseDataList) *GetTrafficDataResponse {
	s.DataList = v
	return s
}

func (s *GetTrafficDataResponse) SetLabels(v *GetTrafficDataResponseLabels) *GetTrafficDataResponse {
	s.Labels = v
	return s
}

type GetTrafficDataResponseDataList struct {
	UsageData []*GetTrafficDataResponseDataListUsageData `json:"UsageData,omitempty" xml:"UsageData,omitempty" require:"true" type:"Repeated"`
}

func (s GetTrafficDataResponseDataList) String() string {
	return tea.Prettify(s)
}

func (s GetTrafficDataResponseDataList) GoString() string {
	return s.String()
}

func (s *GetTrafficDataResponseDataList) SetUsageData(v []*GetTrafficDataResponseDataListUsageData) *GetTrafficDataResponseDataList {
	s.UsageData = v
	return s
}

type GetTrafficDataResponseDataListUsageData struct {
	Date   *string                                        `json:"Date,omitempty" xml:"Date,omitempty" require:"true"`
	Values *GetTrafficDataResponseDataListUsageDataValues `json:"Values,omitempty" xml:"Values,omitempty" require:"true" type:"Struct"`
}

func (s GetTrafficDataResponseDataListUsageData) String() string {
	return tea.Prettify(s)
}

func (s GetTrafficDataResponseDataListUsageData) GoString() string {
	return s.String()
}

func (s *GetTrafficDataResponseDataListUsageData) SetDate(v string) *GetTrafficDataResponseDataListUsageData {
	s.Date = &v
	return s
}

func (s *GetTrafficDataResponseDataListUsageData) SetValues(v *GetTrafficDataResponseDataListUsageDataValues) *GetTrafficDataResponseDataListUsageData {
	s.Values = v
	return s
}

type GetTrafficDataResponseDataListUsageDataValues struct {
	// Values
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" require:"true" type:"Repeated"`
}

func (s GetTrafficDataResponseDataListUsageDataValues) String() string {
	return tea.Prettify(s)
}

func (s GetTrafficDataResponseDataListUsageDataValues) GoString() string {
	return s.String()
}

func (s *GetTrafficDataResponseDataListUsageDataValues) SetValues(v []*string) *GetTrafficDataResponseDataListUsageDataValues {
	s.Values = v
	return s
}

type GetTrafficDataResponseLabels struct {
	Label []*string `json:"Label,omitempty" xml:"Label,omitempty" require:"true" type:"Repeated"`
}

func (s GetTrafficDataResponseLabels) String() string {
	return tea.Prettify(s)
}

func (s GetTrafficDataResponseLabels) GoString() string {
	return s.String()
}

func (s *GetTrafficDataResponseLabels) SetLabel(v []*string) *GetTrafficDataResponseLabels {
	s.Label = v
	return s
}

type GetTrafficByRegionRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
}

func (s GetTrafficByRegionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTrafficByRegionRequest) GoString() string {
	return s.String()
}

func (s *GetTrafficByRegionRequest) SetSecurityToken(v string) *GetTrafficByRegionRequest {
	s.SecurityToken = &v
	return s
}

func (s *GetTrafficByRegionRequest) SetVersion(v string) *GetTrafficByRegionRequest {
	s.Version = &v
	return s
}

type GetTrafficByRegionResponse struct {
	RequestId       *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code            *int                                       `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	TrafficDataList *GetTrafficByRegionResponseTrafficDataList `json:"TrafficDataList,omitempty" xml:"TrafficDataList,omitempty" require:"true" type:"Struct"`
}

func (s GetTrafficByRegionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTrafficByRegionResponse) GoString() string {
	return s.String()
}

func (s *GetTrafficByRegionResponse) SetRequestId(v string) *GetTrafficByRegionResponse {
	s.RequestId = &v
	return s
}

func (s *GetTrafficByRegionResponse) SetCode(v int) *GetTrafficByRegionResponse {
	s.Code = &v
	return s
}

func (s *GetTrafficByRegionResponse) SetTrafficDataList(v *GetTrafficByRegionResponseTrafficDataList) *GetTrafficByRegionResponse {
	s.TrafficDataList = v
	return s
}

type GetTrafficByRegionResponseTrafficDataList struct {
	TrafficData []*GetTrafficByRegionResponseTrafficDataListTrafficData `json:"TrafficData,omitempty" xml:"TrafficData,omitempty" require:"true" type:"Repeated"`
}

func (s GetTrafficByRegionResponseTrafficDataList) String() string {
	return tea.Prettify(s)
}

func (s GetTrafficByRegionResponseTrafficDataList) GoString() string {
	return s.String()
}

func (s *GetTrafficByRegionResponseTrafficDataList) SetTrafficData(v []*GetTrafficByRegionResponseTrafficDataListTrafficData) *GetTrafficByRegionResponseTrafficDataList {
	s.TrafficData = v
	return s
}

type GetTrafficByRegionResponseTrafficDataListTrafficData struct {
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	Traffic *int64  `json:"Traffic,omitempty" xml:"Traffic,omitempty" require:"true"`
}

func (s GetTrafficByRegionResponseTrafficDataListTrafficData) String() string {
	return tea.Prettify(s)
}

func (s GetTrafficByRegionResponseTrafficDataListTrafficData) GoString() string {
	return s.String()
}

func (s *GetTrafficByRegionResponseTrafficDataListTrafficData) SetName(v string) *GetTrafficByRegionResponseTrafficDataListTrafficData {
	s.Name = &v
	return s
}

func (s *GetTrafficByRegionResponseTrafficDataListTrafficData) SetTraffic(v int64) *GetTrafficByRegionResponseTrafficDataListTrafficData {
	s.Traffic = &v
	return s
}

type GetAllRegionsRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
}

func (s GetAllRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAllRegionsRequest) GoString() string {
	return s.String()
}

func (s *GetAllRegionsRequest) SetSecurityToken(v string) *GetAllRegionsRequest {
	s.SecurityToken = &v
	return s
}

func (s *GetAllRegionsRequest) SetVersion(v string) *GetAllRegionsRequest {
	s.Version = &v
	return s
}

type GetAllRegionsResponse struct {
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code      *int                           `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	DataList  *GetAllRegionsResponseDataList `json:"DataList,omitempty" xml:"DataList,omitempty" require:"true" type:"Struct"`
}

func (s GetAllRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAllRegionsResponse) GoString() string {
	return s.String()
}

func (s *GetAllRegionsResponse) SetRequestId(v string) *GetAllRegionsResponse {
	s.RequestId = &v
	return s
}

func (s *GetAllRegionsResponse) SetCode(v int) *GetAllRegionsResponse {
	s.Code = &v
	return s
}

func (s *GetAllRegionsResponse) SetDataList(v *GetAllRegionsResponseDataList) *GetAllRegionsResponse {
	s.DataList = v
	return s
}

type GetAllRegionsResponseDataList struct {
	UsageData []*GetAllRegionsResponseDataListUsageData `json:"UsageData,omitempty" xml:"UsageData,omitempty" require:"true" type:"Repeated"`
}

func (s GetAllRegionsResponseDataList) String() string {
	return tea.Prettify(s)
}

func (s GetAllRegionsResponseDataList) GoString() string {
	return s.String()
}

func (s *GetAllRegionsResponseDataList) SetUsageData(v []*GetAllRegionsResponseDataListUsageData) *GetAllRegionsResponseDataList {
	s.UsageData = v
	return s
}

type GetAllRegionsResponseDataListUsageData struct {
	Code *string `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
}

func (s GetAllRegionsResponseDataListUsageData) String() string {
	return tea.Prettify(s)
}

func (s GetAllRegionsResponseDataListUsageData) GoString() string {
	return s.String()
}

func (s *GetAllRegionsResponseDataListUsageData) SetCode(v string) *GetAllRegionsResponseDataListUsageData {
	s.Code = &v
	return s
}

func (s *GetAllRegionsResponseDataListUsageData) SetName(v string) *GetAllRegionsResponseDataListUsageData {
	s.Name = &v
	return s
}

type Client struct {
	rpc.Client
}

func NewClient(config *rpc.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *rpc.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("central")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("pcdn"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) StopDomainWithOptions(request *StopDomainRequest, runtime *util.RuntimeOptions) (_result *StopDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &StopDomainResponse{}
	_body, _err := client.DoRequest(tea.String("StopDomain"), tea.String("HTTPS"), tea.String("GET"), tea.String("2017-04-11"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopDomain(request *StopDomainRequest) (_result *StopDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopDomainResponse{}
	_body, _err := client.StopDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartDomainWithOptions(request *StartDomainRequest, runtime *util.RuntimeOptions) (_result *StartDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &StartDomainResponse{}
	_body, _err := client.DoRequest(tea.String("StartDomain"), tea.String("HTTPS"), tea.String("GET"), tea.String("2017-04-11"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartDomain(request *StartDomainRequest) (_result *StartDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartDomainResponse{}
	_body, _err := client.StartDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDomainWithOptions(request *DeleteDomainRequest, runtime *util.RuntimeOptions) (_result *DeleteDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteDomainResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteDomain"), tea.String("HTTPS"), tea.String("GET"), tea.String("2017-04-11"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDomain(request *DeleteDomainRequest) (_result *DeleteDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDomainResponse{}
	_body, _err := client.DeleteDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddDomainWithOptions(request *AddDomainRequest, runtime *util.RuntimeOptions) (_result *AddDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AddDomainResponse{}
	_body, _err := client.DoRequest(tea.String("AddDomain"), tea.String("HTTPS"), tea.String("GET"), tea.String("2017-04-11"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddDomain(request *AddDomainRequest) (_result *AddDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddDomainResponse{}
	_body, _err := client.AddDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetBalanceTrafficDataWithOptions(request *GetBalanceTrafficDataRequest, runtime *util.RuntimeOptions) (_result *GetBalanceTrafficDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetBalanceTrafficDataResponse{}
	_body, _err := client.DoRequest(tea.String("GetBalanceTrafficData"), tea.String("HTTPS"), tea.String("GET"), tea.String("2017-04-11"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetBalanceTrafficData(request *GetBalanceTrafficDataRequest) (_result *GetBalanceTrafficDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetBalanceTrafficDataResponse{}
	_body, _err := client.GetBalanceTrafficDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddPcdnControlRuleWithOptions(request *AddPcdnControlRuleRequest, runtime *util.RuntimeOptions) (_result *AddPcdnControlRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AddPcdnControlRuleResponse{}
	_body, _err := client.DoRequest(tea.String("AddPcdnControlRule"), tea.String("HTTPS"), tea.String("GET"), tea.String("2017-04-11"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddPcdnControlRule(request *AddPcdnControlRuleRequest) (_result *AddPcdnControlRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddPcdnControlRuleResponse{}
	_body, _err := client.AddPcdnControlRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddConsumerWithOptions(request *AddConsumerRequest, runtime *util.RuntimeOptions) (_result *AddConsumerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AddConsumerResponse{}
	_body, _err := client.DoRequest(tea.String("AddConsumer"), tea.String("HTTPS"), tea.String("GET"), tea.String("2017-04-11"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddConsumer(request *AddConsumerRequest) (_result *AddConsumerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddConsumerResponse{}
	_body, _err := client.AddConsumerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAccessDataWithOptions(request *GetAccessDataRequest, runtime *util.RuntimeOptions) (_result *GetAccessDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetAccessDataResponse{}
	_body, _err := client.DoRequest(tea.String("GetAccessData"), tea.String("HTTPS"), tea.String("GET"), tea.String("2017-04-11"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAccessData(request *GetAccessDataRequest) (_result *GetAccessDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAccessDataResponse{}
	_body, _err := client.GetAccessDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnablePcdnControlRuleWithOptions(request *EnablePcdnControlRuleRequest, runtime *util.RuntimeOptions) (_result *EnablePcdnControlRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &EnablePcdnControlRuleResponse{}
	_body, _err := client.DoRequest(tea.String("EnablePcdnControlRule"), tea.String("HTTPS"), tea.String("GET"), tea.String("2017-04-11"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnablePcdnControlRule(request *EnablePcdnControlRuleRequest) (_result *EnablePcdnControlRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnablePcdnControlRuleResponse{}
	_body, _err := client.EnablePcdnControlRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EditPcdnControlRuleWithOptions(request *EditPcdnControlRuleRequest, runtime *util.RuntimeOptions) (_result *EditPcdnControlRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &EditPcdnControlRuleResponse{}
	_body, _err := client.DoRequest(tea.String("EditPcdnControlRule"), tea.String("HTTPS"), tea.String("GET"), tea.String("2017-04-11"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EditPcdnControlRule(request *EditPcdnControlRuleRequest) (_result *EditPcdnControlRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EditPcdnControlRuleResponse{}
	_body, _err := client.EditPcdnControlRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisablePcdnControlRuleWithOptions(request *DisablePcdnControlRuleRequest, runtime *util.RuntimeOptions) (_result *DisablePcdnControlRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DisablePcdnControlRuleResponse{}
	_body, _err := client.DoRequest(tea.String("DisablePcdnControlRule"), tea.String("HTTPS"), tea.String("GET"), tea.String("2017-04-11"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisablePcdnControlRule(request *DisablePcdnControlRuleRequest) (_result *DisablePcdnControlRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisablePcdnControlRuleResponse{}
	_body, _err := client.DisablePcdnControlRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeletePcdnControlRuleWithOptions(request *DeletePcdnControlRuleRequest, runtime *util.RuntimeOptions) (_result *DeletePcdnControlRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeletePcdnControlRuleResponse{}
	_body, _err := client.DoRequest(tea.String("DeletePcdnControlRule"), tea.String("HTTPS"), tea.String("GET"), tea.String("2017-04-11"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeletePcdnControlRule(request *DeletePcdnControlRuleRequest) (_result *DeletePcdnControlRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeletePcdnControlRuleResponse{}
	_body, _err := client.DeletePcdnControlRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAllPlatformTypesWithOptions(request *GetAllPlatformTypesRequest, runtime *util.RuntimeOptions) (_result *GetAllPlatformTypesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetAllPlatformTypesResponse{}
	_body, _err := client.DoRequest(tea.String("GetAllPlatformTypes"), tea.String("HTTPS"), tea.String("GET"), tea.String("2017-04-11"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAllPlatformTypes(request *GetAllPlatformTypesRequest) (_result *GetAllPlatformTypesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAllPlatformTypesResponse{}
	_body, _err := client.GetAllPlatformTypesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAllMarketsWithOptions(request *GetAllMarketsRequest, runtime *util.RuntimeOptions) (_result *GetAllMarketsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetAllMarketsResponse{}
	_body, _err := client.DoRequest(tea.String("GetAllMarkets"), tea.String("HTTPS"), tea.String("GET"), tea.String("2017-04-11"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAllMarkets(request *GetAllMarketsRequest) (_result *GetAllMarketsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAllMarketsResponse{}
	_body, _err := client.GetAllMarketsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAllIspWithOptions(request *GetAllIspRequest, runtime *util.RuntimeOptions) (_result *GetAllIspResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetAllIspResponse{}
	_body, _err := client.DoRequest(tea.String("GetAllIsp"), tea.String("HTTPS"), tea.String("GET"), tea.String("2017-04-11"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAllIsp(request *GetAllIspRequest) (_result *GetAllIspResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAllIspResponse{}
	_body, _err := client.GetAllIspWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAllAppVersionsWithOptions(request *GetAllAppVersionsRequest, runtime *util.RuntimeOptions) (_result *GetAllAppVersionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetAllAppVersionsResponse{}
	_body, _err := client.DoRequest(tea.String("GetAllAppVersions"), tea.String("HTTPS"), tea.String("GET"), tea.String("2017-04-11"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAllAppVersions(request *GetAllAppVersionsRequest) (_result *GetAllAppVersionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAllAppVersionsResponse{}
	_body, _err := client.GetAllAppVersionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetConsumerStatusWithOptions(request *GetConsumerStatusRequest, runtime *util.RuntimeOptions) (_result *GetConsumerStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetConsumerStatusResponse{}
	_body, _err := client.DoRequest(tea.String("GetConsumerStatus"), tea.String("HTTPS"), tea.String("GET"), tea.String("2017-04-11"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetConsumerStatus(request *GetConsumerStatusRequest) (_result *GetConsumerStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetConsumerStatusResponse{}
	_body, _err := client.GetConsumerStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetClientsRatioWithOptions(request *GetClientsRatioRequest, runtime *util.RuntimeOptions) (_result *GetClientsRatioResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetClientsRatioResponse{}
	_body, _err := client.DoRequest(tea.String("GetClientsRatio"), tea.String("HTTPS"), tea.String("GET"), tea.String("2017-04-11"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetClientsRatio(request *GetClientsRatioRequest) (_result *GetClientsRatioResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetClientsRatioResponse{}
	_body, _err := client.GetClientsRatioWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetBandwidthDataWithOptions(request *GetBandwidthDataRequest, runtime *util.RuntimeOptions) (_result *GetBandwidthDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetBandwidthDataResponse{}
	_body, _err := client.DoRequest(tea.String("GetBandwidthData"), tea.String("HTTPS"), tea.String("GET"), tea.String("2017-04-11"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetBandwidthData(request *GetBandwidthDataRequest) (_result *GetBandwidthDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetBandwidthDataResponse{}
	_body, _err := client.GetBandwidthDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetBalanceBandwidthDataWithOptions(request *GetBalanceBandwidthDataRequest, runtime *util.RuntimeOptions) (_result *GetBalanceBandwidthDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetBalanceBandwidthDataResponse{}
	_body, _err := client.DoRequest(tea.String("GetBalanceBandwidthData"), tea.String("HTTPS"), tea.String("GET"), tea.String("2017-04-11"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetBalanceBandwidthData(request *GetBalanceBandwidthDataRequest) (_result *GetBalanceBandwidthDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetBalanceBandwidthDataResponse{}
	_body, _err := client.GetBalanceBandwidthDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetControlRulesWithOptions(request *GetControlRulesRequest, runtime *util.RuntimeOptions) (_result *GetControlRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetControlRulesResponse{}
	_body, _err := client.DoRequest(tea.String("GetControlRules"), tea.String("HTTPS"), tea.String("GET"), tea.String("2017-04-11"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetControlRules(request *GetControlRulesRequest) (_result *GetControlRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetControlRulesResponse{}
	_body, _err := client.GetControlRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDomainCountWithOptions(request *GetDomainCountRequest, runtime *util.RuntimeOptions) (_result *GetDomainCountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetDomainCountResponse{}
	_body, _err := client.DoRequest(tea.String("GetDomainCount"), tea.String("HTTPS"), tea.String("GET"), tea.String("2017-04-11"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDomainCount(request *GetDomainCountRequest) (_result *GetDomainCountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDomainCountResponse{}
	_body, _err := client.GetDomainCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCurrentModeWithOptions(request *GetCurrentModeRequest, runtime *util.RuntimeOptions) (_result *GetCurrentModeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetCurrentModeResponse{}
	_body, _err := client.DoRequest(tea.String("GetCurrentMode"), tea.String("HTTPS"), tea.String("GET"), tea.String("2017-04-11"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCurrentMode(request *GetCurrentModeRequest) (_result *GetCurrentModeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCurrentModeResponse{}
	_body, _err := client.GetCurrentModeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCoverRateDataWithOptions(request *GetCoverRateDataRequest, runtime *util.RuntimeOptions) (_result *GetCoverRateDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetCoverRateDataResponse{}
	_body, _err := client.DoRequest(tea.String("GetCoverRateData"), tea.String("HTTPS"), tea.String("GET"), tea.String("2017-04-11"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCoverRateData(request *GetCoverRateDataRequest) (_result *GetCoverRateDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCoverRateDataResponse{}
	_body, _err := client.GetCoverRateDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetFeeHistoryWithOptions(request *GetFeeHistoryRequest, runtime *util.RuntimeOptions) (_result *GetFeeHistoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetFeeHistoryResponse{}
	_body, _err := client.DoRequest(tea.String("GetFeeHistory"), tea.String("HTTPS"), tea.String("GET"), tea.String("2017-04-11"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetFeeHistory(request *GetFeeHistoryRequest) (_result *GetFeeHistoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetFeeHistoryResponse{}
	_body, _err := client.GetFeeHistoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetExpenseSummaryWithOptions(request *GetExpenseSummaryRequest, runtime *util.RuntimeOptions) (_result *GetExpenseSummaryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetExpenseSummaryResponse{}
	_body, _err := client.DoRequest(tea.String("GetExpenseSummary"), tea.String("HTTPS"), tea.String("GET"), tea.String("2017-04-11"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetExpenseSummary(request *GetExpenseSummaryRequest) (_result *GetExpenseSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetExpenseSummaryResponse{}
	_body, _err := client.GetExpenseSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDomainsWithOptions(request *GetDomainsRequest, runtime *util.RuntimeOptions) (_result *GetDomainsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetDomainsResponse{}
	_body, _err := client.DoRequest(tea.String("GetDomains"), tea.String("HTTPS"), tea.String("GET"), tea.String("2017-04-11"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDomains(request *GetDomainsRequest) (_result *GetDomainsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDomainsResponse{}
	_body, _err := client.GetDomainsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetLogsListWithOptions(request *GetLogsListRequest, runtime *util.RuntimeOptions) (_result *GetLogsListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetLogsListResponse{}
	_body, _err := client.DoRequest(tea.String("GetLogsList"), tea.String("HTTPS"), tea.String("GET"), tea.String("2017-04-11"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetLogsList(request *GetLogsListRequest) (_result *GetLogsListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetLogsListResponse{}
	_body, _err := client.GetLogsListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetFluencyDataWithOptions(request *GetFluencyDataRequest, runtime *util.RuntimeOptions) (_result *GetFluencyDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetFluencyDataResponse{}
	_body, _err := client.DoRequest(tea.String("GetFluencyData"), tea.String("HTTPS"), tea.String("GET"), tea.String("2017-04-11"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetFluencyData(request *GetFluencyDataRequest) (_result *GetFluencyDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetFluencyDataResponse{}
	_body, _err := client.GetFluencyDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetFirstFrameDelayDataWithOptions(request *GetFirstFrameDelayDataRequest, runtime *util.RuntimeOptions) (_result *GetFirstFrameDelayDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetFirstFrameDelayDataResponse{}
	_body, _err := client.DoRequest(tea.String("GetFirstFrameDelayData"), tea.String("HTTPS"), tea.String("GET"), tea.String("2017-04-11"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetFirstFrameDelayData(request *GetFirstFrameDelayDataRequest) (_result *GetFirstFrameDelayDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetFirstFrameDelayDataResponse{}
	_body, _err := client.GetFirstFrameDelayDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTokenListWithOptions(request *GetTokenListRequest, runtime *util.RuntimeOptions) (_result *GetTokenListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetTokenListResponse{}
	_body, _err := client.DoRequest(tea.String("GetTokenList"), tea.String("HTTPS"), tea.String("GET"), tea.String("2017-04-11"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTokenList(request *GetTokenListRequest) (_result *GetTokenListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTokenListResponse{}
	_body, _err := client.GetTokenListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetShareRateDataWithOptions(request *GetShareRateDataRequest, runtime *util.RuntimeOptions) (_result *GetShareRateDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetShareRateDataResponse{}
	_body, _err := client.DoRequest(tea.String("GetShareRateData"), tea.String("HTTPS"), tea.String("GET"), tea.String("2017-04-11"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetShareRateData(request *GetShareRateDataRequest) (_result *GetShareRateDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetShareRateDataResponse{}
	_body, _err := client.GetShareRateDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTrafficDataWithOptions(request *GetTrafficDataRequest, runtime *util.RuntimeOptions) (_result *GetTrafficDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetTrafficDataResponse{}
	_body, _err := client.DoRequest(tea.String("GetTrafficData"), tea.String("HTTPS"), tea.String("GET"), tea.String("2017-04-11"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTrafficData(request *GetTrafficDataRequest) (_result *GetTrafficDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTrafficDataResponse{}
	_body, _err := client.GetTrafficDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTrafficByRegionWithOptions(request *GetTrafficByRegionRequest, runtime *util.RuntimeOptions) (_result *GetTrafficByRegionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetTrafficByRegionResponse{}
	_body, _err := client.DoRequest(tea.String("GetTrafficByRegion"), tea.String("HTTPS"), tea.String("GET"), tea.String("2017-04-11"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTrafficByRegion(request *GetTrafficByRegionRequest) (_result *GetTrafficByRegionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTrafficByRegionResponse{}
	_body, _err := client.GetTrafficByRegionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAllRegionsWithOptions(request *GetAllRegionsRequest, runtime *util.RuntimeOptions) (_result *GetAllRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetAllRegionsResponse{}
	_body, _err := client.DoRequest(tea.String("GetAllRegions"), tea.String("HTTPS"), tea.String("GET"), tea.String("2017-04-11"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAllRegions(request *GetAllRegionsRequest) (_result *GetAllRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAllRegionsResponse{}
	_body, _err := client.GetAllRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
