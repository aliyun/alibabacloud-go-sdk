// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddConsumerRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Version              *string `json:"Version,omitempty" xml:"Version,omitempty"`
	BusinessType         *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	Company              *string `json:"Company,omitempty" xml:"Company,omitempty"`
	Site                 *string `json:"Site,omitempty" xml:"Site,omitempty"`
	Requirement          *string `json:"Requirement,omitempty" xml:"Requirement,omitempty"`
	Mobile               *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
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

type AddConsumerResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	Code       *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s AddConsumerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddConsumerResponseBody) GoString() string {
	return s.String()
}

func (s *AddConsumerResponseBody) SetRequestId(v string) *AddConsumerResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddConsumerResponseBody) SetResourceId(v string) *AddConsumerResponseBody {
	s.ResourceId = &v
	return s
}

func (s *AddConsumerResponseBody) SetCode(v int32) *AddConsumerResponseBody {
	s.Code = &v
	return s
}

type AddConsumerResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddConsumerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddConsumerResponse) String() string {
	return tea.Prettify(s)
}

func (s AddConsumerResponse) GoString() string {
	return s.String()
}

func (s *AddConsumerResponse) SetHeaders(v map[string]*string) *AddConsumerResponse {
	s.Headers = v
	return s
}

func (s *AddConsumerResponse) SetBody(v *AddConsumerResponseBody) *AddConsumerResponse {
	s.Body = v
	return s
}

type AddDomainRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty"`
	BusinessType  *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	Domain        *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
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

type AddDomainResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	Code       *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s AddDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddDomainResponseBody) GoString() string {
	return s.String()
}

func (s *AddDomainResponseBody) SetRequestId(v string) *AddDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddDomainResponseBody) SetResourceId(v string) *AddDomainResponseBody {
	s.ResourceId = &v
	return s
}

func (s *AddDomainResponseBody) SetCode(v int32) *AddDomainResponseBody {
	s.Code = &v
	return s
}

type AddDomainResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s AddDomainResponse) GoString() string {
	return s.String()
}

func (s *AddDomainResponse) SetHeaders(v map[string]*string) *AddDomainResponse {
	s.Headers = v
	return s
}

func (s *AddDomainResponse) SetBody(v *AddDomainResponseBody) *AddDomainResponse {
	s.Body = v
	return s
}

type AddPcdnControlRuleRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Region        *string `json:"Region,omitempty" xml:"Region,omitempty"`
	IspName       *string `json:"IspName,omitempty" xml:"IspName,omitempty"`
	PlatformType  *string `json:"PlatformType,omitempty" xml:"PlatformType,omitempty"`
	BusinessType  *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	Market        *string `json:"Market,omitempty" xml:"Market,omitempty"`
	AppVersion    *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
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

type AddPcdnControlRuleResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	Code       *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s AddPcdnControlRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddPcdnControlRuleResponseBody) GoString() string {
	return s.String()
}

func (s *AddPcdnControlRuleResponseBody) SetRequestId(v string) *AddPcdnControlRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddPcdnControlRuleResponseBody) SetResourceId(v string) *AddPcdnControlRuleResponseBody {
	s.ResourceId = &v
	return s
}

func (s *AddPcdnControlRuleResponseBody) SetCode(v int32) *AddPcdnControlRuleResponseBody {
	s.Code = &v
	return s
}

type AddPcdnControlRuleResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddPcdnControlRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddPcdnControlRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s AddPcdnControlRuleResponse) GoString() string {
	return s.String()
}

func (s *AddPcdnControlRuleResponse) SetHeaders(v map[string]*string) *AddPcdnControlRuleResponse {
	s.Headers = v
	return s
}

func (s *AddPcdnControlRuleResponse) SetBody(v *AddPcdnControlRuleResponseBody) *AddPcdnControlRuleResponse {
	s.Body = v
	return s
}

type DeleteDomainRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty"`
	Domain        *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
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

type DeleteDomainResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	Code       *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s DeleteDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDomainResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDomainResponseBody) SetRequestId(v string) *DeleteDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDomainResponseBody) SetResourceId(v string) *DeleteDomainResponseBody {
	s.ResourceId = &v
	return s
}

func (s *DeleteDomainResponseBody) SetCode(v int32) *DeleteDomainResponseBody {
	s.Code = &v
	return s
}

type DeleteDomainResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDomainResponse) GoString() string {
	return s.String()
}

func (s *DeleteDomainResponse) SetHeaders(v map[string]*string) *DeleteDomainResponse {
	s.Headers = v
	return s
}

func (s *DeleteDomainResponse) SetBody(v *DeleteDomainResponseBody) *DeleteDomainResponse {
	s.Body = v
	return s
}

type DeletePcdnControlRuleRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty"`
	ResourceId    *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
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

type DeletePcdnControlRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s DeletePcdnControlRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeletePcdnControlRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePcdnControlRuleResponseBody) SetRequestId(v string) *DeletePcdnControlRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePcdnControlRuleResponseBody) SetCode(v int32) *DeletePcdnControlRuleResponseBody {
	s.Code = &v
	return s
}

type DeletePcdnControlRuleResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeletePcdnControlRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeletePcdnControlRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePcdnControlRuleResponse) GoString() string {
	return s.String()
}

func (s *DeletePcdnControlRuleResponse) SetHeaders(v map[string]*string) *DeletePcdnControlRuleResponse {
	s.Headers = v
	return s
}

func (s *DeletePcdnControlRuleResponse) SetBody(v *DeletePcdnControlRuleResponseBody) *DeletePcdnControlRuleResponse {
	s.Body = v
	return s
}

type DisablePcdnControlRuleRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty"`
	ResourceId    *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
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

type DisablePcdnControlRuleResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	Code       *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s DisablePcdnControlRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisablePcdnControlRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DisablePcdnControlRuleResponseBody) SetRequestId(v string) *DisablePcdnControlRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisablePcdnControlRuleResponseBody) SetResourceId(v string) *DisablePcdnControlRuleResponseBody {
	s.ResourceId = &v
	return s
}

func (s *DisablePcdnControlRuleResponseBody) SetCode(v int32) *DisablePcdnControlRuleResponseBody {
	s.Code = &v
	return s
}

type DisablePcdnControlRuleResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DisablePcdnControlRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisablePcdnControlRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DisablePcdnControlRuleResponse) GoString() string {
	return s.String()
}

func (s *DisablePcdnControlRuleResponse) SetHeaders(v map[string]*string) *DisablePcdnControlRuleResponse {
	s.Headers = v
	return s
}

func (s *DisablePcdnControlRuleResponse) SetBody(v *DisablePcdnControlRuleResponseBody) *DisablePcdnControlRuleResponse {
	s.Body = v
	return s
}

type EditPcdnControlRuleRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ResourceId    *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	Region        *string `json:"Region,omitempty" xml:"Region,omitempty"`
	IspName       *string `json:"IspName,omitempty" xml:"IspName,omitempty"`
	PlatformType  *string `json:"PlatformType,omitempty" xml:"PlatformType,omitempty"`
	BusinessType  *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	Market        *string `json:"Market,omitempty" xml:"Market,omitempty"`
	AppVersion    *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
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

type EditPcdnControlRuleResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	Code       *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s EditPcdnControlRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EditPcdnControlRuleResponseBody) GoString() string {
	return s.String()
}

func (s *EditPcdnControlRuleResponseBody) SetRequestId(v string) *EditPcdnControlRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *EditPcdnControlRuleResponseBody) SetResourceId(v string) *EditPcdnControlRuleResponseBody {
	s.ResourceId = &v
	return s
}

func (s *EditPcdnControlRuleResponseBody) SetCode(v int32) *EditPcdnControlRuleResponseBody {
	s.Code = &v
	return s
}

type EditPcdnControlRuleResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EditPcdnControlRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EditPcdnControlRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s EditPcdnControlRuleResponse) GoString() string {
	return s.String()
}

func (s *EditPcdnControlRuleResponse) SetHeaders(v map[string]*string) *EditPcdnControlRuleResponse {
	s.Headers = v
	return s
}

func (s *EditPcdnControlRuleResponse) SetBody(v *EditPcdnControlRuleResponseBody) *EditPcdnControlRuleResponse {
	s.Body = v
	return s
}

type EnablePcdnControlRuleRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty"`
	ResourceId    *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
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

type EnablePcdnControlRuleResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	Code       *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s EnablePcdnControlRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnablePcdnControlRuleResponseBody) GoString() string {
	return s.String()
}

func (s *EnablePcdnControlRuleResponseBody) SetRequestId(v string) *EnablePcdnControlRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnablePcdnControlRuleResponseBody) SetResourceId(v string) *EnablePcdnControlRuleResponseBody {
	s.ResourceId = &v
	return s
}

func (s *EnablePcdnControlRuleResponseBody) SetCode(v int32) *EnablePcdnControlRuleResponseBody {
	s.Code = &v
	return s
}

type EnablePcdnControlRuleResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EnablePcdnControlRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnablePcdnControlRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s EnablePcdnControlRuleResponse) GoString() string {
	return s.String()
}

func (s *EnablePcdnControlRuleResponse) SetHeaders(v map[string]*string) *EnablePcdnControlRuleResponse {
	s.Headers = v
	return s
}

func (s *EnablePcdnControlRuleResponse) SetBody(v *EnablePcdnControlRuleResponseBody) *EnablePcdnControlRuleResponse {
	s.Body = v
	return s
}

type GetAccessDataRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty"`
	Domain        *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Region        *string `json:"Region,omitempty" xml:"Region,omitempty"`
	IspName       *string `json:"IspName,omitempty" xml:"IspName,omitempty"`
	PlatformType  *string `json:"PlatformType,omitempty" xml:"PlatformType,omitempty"`
	BusinessType  *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	StartDate     *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	EndDate       *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
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

type GetAccessDataResponseBody struct {
	DataList  *GetAccessDataResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Struct"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Labels    *GetAccessDataResponseBodyLabels   `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Struct"`
	Code      *int32                             `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetAccessDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAccessDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccessDataResponseBody) SetDataList(v *GetAccessDataResponseBodyDataList) *GetAccessDataResponseBody {
	s.DataList = v
	return s
}

func (s *GetAccessDataResponseBody) SetRequestId(v string) *GetAccessDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAccessDataResponseBody) SetLabels(v *GetAccessDataResponseBodyLabels) *GetAccessDataResponseBody {
	s.Labels = v
	return s
}

func (s *GetAccessDataResponseBody) SetCode(v int32) *GetAccessDataResponseBody {
	s.Code = &v
	return s
}

type GetAccessDataResponseBodyDataList struct {
	UsageData []*GetAccessDataResponseBodyDataListUsageData `json:"UsageData,omitempty" xml:"UsageData,omitempty" type:"Repeated"`
}

func (s GetAccessDataResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s GetAccessDataResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *GetAccessDataResponseBodyDataList) SetUsageData(v []*GetAccessDataResponseBodyDataListUsageData) *GetAccessDataResponseBodyDataList {
	s.UsageData = v
	return s
}

type GetAccessDataResponseBodyDataListUsageData struct {
	Values *GetAccessDataResponseBodyDataListUsageDataValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Struct"`
	Date   *string                                           `json:"Date,omitempty" xml:"Date,omitempty"`
}

func (s GetAccessDataResponseBodyDataListUsageData) String() string {
	return tea.Prettify(s)
}

func (s GetAccessDataResponseBodyDataListUsageData) GoString() string {
	return s.String()
}

func (s *GetAccessDataResponseBodyDataListUsageData) SetValues(v *GetAccessDataResponseBodyDataListUsageDataValues) *GetAccessDataResponseBodyDataListUsageData {
	s.Values = v
	return s
}

func (s *GetAccessDataResponseBodyDataListUsageData) SetDate(v string) *GetAccessDataResponseBodyDataListUsageData {
	s.Date = &v
	return s
}

type GetAccessDataResponseBodyDataListUsageDataValues struct {
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s GetAccessDataResponseBodyDataListUsageDataValues) String() string {
	return tea.Prettify(s)
}

func (s GetAccessDataResponseBodyDataListUsageDataValues) GoString() string {
	return s.String()
}

func (s *GetAccessDataResponseBodyDataListUsageDataValues) SetValues(v []*string) *GetAccessDataResponseBodyDataListUsageDataValues {
	s.Values = v
	return s
}

type GetAccessDataResponseBodyLabels struct {
	Label []*string `json:"Label,omitempty" xml:"Label,omitempty" type:"Repeated"`
}

func (s GetAccessDataResponseBodyLabels) String() string {
	return tea.Prettify(s)
}

func (s GetAccessDataResponseBodyLabels) GoString() string {
	return s.String()
}

func (s *GetAccessDataResponseBodyLabels) SetLabel(v []*string) *GetAccessDataResponseBodyLabels {
	s.Label = v
	return s
}

type GetAccessDataResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAccessDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAccessDataResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAccessDataResponse) GoString() string {
	return s.String()
}

func (s *GetAccessDataResponse) SetHeaders(v map[string]*string) *GetAccessDataResponse {
	s.Headers = v
	return s
}

func (s *GetAccessDataResponse) SetBody(v *GetAccessDataResponseBody) *GetAccessDataResponse {
	s.Body = v
	return s
}

type GetAllAppVersionsRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty"`
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

type GetAllAppVersionsResponseBody struct {
	DataList  *GetAllAppVersionsResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Struct"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *int32                                 `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetAllAppVersionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAllAppVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *GetAllAppVersionsResponseBody) SetDataList(v *GetAllAppVersionsResponseBodyDataList) *GetAllAppVersionsResponseBody {
	s.DataList = v
	return s
}

func (s *GetAllAppVersionsResponseBody) SetRequestId(v string) *GetAllAppVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAllAppVersionsResponseBody) SetCode(v int32) *GetAllAppVersionsResponseBody {
	s.Code = &v
	return s
}

type GetAllAppVersionsResponseBodyDataList struct {
	UsageData []*GetAllAppVersionsResponseBodyDataListUsageData `json:"UsageData,omitempty" xml:"UsageData,omitempty" type:"Repeated"`
}

func (s GetAllAppVersionsResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s GetAllAppVersionsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *GetAllAppVersionsResponseBodyDataList) SetUsageData(v []*GetAllAppVersionsResponseBodyDataListUsageData) *GetAllAppVersionsResponseBodyDataList {
	s.UsageData = v
	return s
}

type GetAllAppVersionsResponseBodyDataListUsageData struct {
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Code  *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetAllAppVersionsResponseBodyDataListUsageData) String() string {
	return tea.Prettify(s)
}

func (s GetAllAppVersionsResponseBodyDataListUsageData) GoString() string {
	return s.String()
}

func (s *GetAllAppVersionsResponseBodyDataListUsageData) SetValue(v string) *GetAllAppVersionsResponseBodyDataListUsageData {
	s.Value = &v
	return s
}

func (s *GetAllAppVersionsResponseBodyDataListUsageData) SetCode(v int32) *GetAllAppVersionsResponseBodyDataListUsageData {
	s.Code = &v
	return s
}

type GetAllAppVersionsResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAllAppVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAllAppVersionsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAllAppVersionsResponse) GoString() string {
	return s.String()
}

func (s *GetAllAppVersionsResponse) SetHeaders(v map[string]*string) *GetAllAppVersionsResponse {
	s.Headers = v
	return s
}

func (s *GetAllAppVersionsResponse) SetBody(v *GetAllAppVersionsResponseBody) *GetAllAppVersionsResponse {
	s.Body = v
	return s
}

type GetAllIspRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty"`
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

type GetAllIspResponseBody struct {
	DataList  *GetAllIspResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Struct"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *int32                         `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetAllIspResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAllIspResponseBody) GoString() string {
	return s.String()
}

func (s *GetAllIspResponseBody) SetDataList(v *GetAllIspResponseBodyDataList) *GetAllIspResponseBody {
	s.DataList = v
	return s
}

func (s *GetAllIspResponseBody) SetRequestId(v string) *GetAllIspResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAllIspResponseBody) SetCode(v int32) *GetAllIspResponseBody {
	s.Code = &v
	return s
}

type GetAllIspResponseBodyDataList struct {
	UsageData []*GetAllIspResponseBodyDataListUsageData `json:"UsageData,omitempty" xml:"UsageData,omitempty" type:"Repeated"`
}

func (s GetAllIspResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s GetAllIspResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *GetAllIspResponseBodyDataList) SetUsageData(v []*GetAllIspResponseBodyDataListUsageData) *GetAllIspResponseBodyDataList {
	s.UsageData = v
	return s
}

type GetAllIspResponseBodyDataListUsageData struct {
	NameEn     *string `json:"NameEn,omitempty" xml:"NameEn,omitempty"`
	NameCn     *string `json:"NameCn,omitempty" xml:"NameCn,omitempty"`
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s GetAllIspResponseBodyDataListUsageData) String() string {
	return tea.Prettify(s)
}

func (s GetAllIspResponseBodyDataListUsageData) GoString() string {
	return s.String()
}

func (s *GetAllIspResponseBodyDataListUsageData) SetNameEn(v string) *GetAllIspResponseBodyDataListUsageData {
	s.NameEn = &v
	return s
}

func (s *GetAllIspResponseBodyDataListUsageData) SetNameCn(v string) *GetAllIspResponseBodyDataListUsageData {
	s.NameCn = &v
	return s
}

func (s *GetAllIspResponseBodyDataListUsageData) SetResourceId(v string) *GetAllIspResponseBodyDataListUsageData {
	s.ResourceId = &v
	return s
}

type GetAllIspResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAllIspResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAllIspResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAllIspResponse) GoString() string {
	return s.String()
}

func (s *GetAllIspResponse) SetHeaders(v map[string]*string) *GetAllIspResponse {
	s.Headers = v
	return s
}

func (s *GetAllIspResponse) SetBody(v *GetAllIspResponseBody) *GetAllIspResponse {
	s.Body = v
	return s
}

type GetAllMarketsRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty"`
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

type GetAllMarketsResponseBody struct {
	DataList  *GetAllMarketsResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Struct"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *int32                             `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetAllMarketsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAllMarketsResponseBody) GoString() string {
	return s.String()
}

func (s *GetAllMarketsResponseBody) SetDataList(v *GetAllMarketsResponseBodyDataList) *GetAllMarketsResponseBody {
	s.DataList = v
	return s
}

func (s *GetAllMarketsResponseBody) SetRequestId(v string) *GetAllMarketsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAllMarketsResponseBody) SetCode(v int32) *GetAllMarketsResponseBody {
	s.Code = &v
	return s
}

type GetAllMarketsResponseBodyDataList struct {
	UsageData []*GetAllMarketsResponseBodyDataListUsageData `json:"UsageData,omitempty" xml:"UsageData,omitempty" type:"Repeated"`
}

func (s GetAllMarketsResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s GetAllMarketsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *GetAllMarketsResponseBodyDataList) SetUsageData(v []*GetAllMarketsResponseBodyDataListUsageData) *GetAllMarketsResponseBodyDataList {
	s.UsageData = v
	return s
}

type GetAllMarketsResponseBodyDataListUsageData struct {
	Code       *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	MarketCode *string `json:"MarketCode,omitempty" xml:"MarketCode,omitempty"`
	MarketName *string `json:"MarketName,omitempty" xml:"MarketName,omitempty"`
}

func (s GetAllMarketsResponseBodyDataListUsageData) String() string {
	return tea.Prettify(s)
}

func (s GetAllMarketsResponseBodyDataListUsageData) GoString() string {
	return s.String()
}

func (s *GetAllMarketsResponseBodyDataListUsageData) SetCode(v int32) *GetAllMarketsResponseBodyDataListUsageData {
	s.Code = &v
	return s
}

func (s *GetAllMarketsResponseBodyDataListUsageData) SetMarketCode(v string) *GetAllMarketsResponseBodyDataListUsageData {
	s.MarketCode = &v
	return s
}

func (s *GetAllMarketsResponseBodyDataListUsageData) SetMarketName(v string) *GetAllMarketsResponseBodyDataListUsageData {
	s.MarketName = &v
	return s
}

type GetAllMarketsResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAllMarketsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAllMarketsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAllMarketsResponse) GoString() string {
	return s.String()
}

func (s *GetAllMarketsResponse) SetHeaders(v map[string]*string) *GetAllMarketsResponse {
	s.Headers = v
	return s
}

func (s *GetAllMarketsResponse) SetBody(v *GetAllMarketsResponseBody) *GetAllMarketsResponse {
	s.Body = v
	return s
}

type GetAllPlatformTypesRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty"`
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

type GetAllPlatformTypesResponseBody struct {
	DataList  *GetAllPlatformTypesResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Struct"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *int32                                   `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetAllPlatformTypesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAllPlatformTypesResponseBody) GoString() string {
	return s.String()
}

func (s *GetAllPlatformTypesResponseBody) SetDataList(v *GetAllPlatformTypesResponseBodyDataList) *GetAllPlatformTypesResponseBody {
	s.DataList = v
	return s
}

func (s *GetAllPlatformTypesResponseBody) SetRequestId(v string) *GetAllPlatformTypesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAllPlatformTypesResponseBody) SetCode(v int32) *GetAllPlatformTypesResponseBody {
	s.Code = &v
	return s
}

type GetAllPlatformTypesResponseBodyDataList struct {
	UsageData []*GetAllPlatformTypesResponseBodyDataListUsageData `json:"UsageData,omitempty" xml:"UsageData,omitempty" type:"Repeated"`
}

func (s GetAllPlatformTypesResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s GetAllPlatformTypesResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *GetAllPlatformTypesResponseBodyDataList) SetUsageData(v []*GetAllPlatformTypesResponseBodyDataListUsageData) *GetAllPlatformTypesResponseBodyDataList {
	s.UsageData = v
	return s
}

type GetAllPlatformTypesResponseBodyDataListUsageData struct {
	Code *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetAllPlatformTypesResponseBodyDataListUsageData) String() string {
	return tea.Prettify(s)
}

func (s GetAllPlatformTypesResponseBodyDataListUsageData) GoString() string {
	return s.String()
}

func (s *GetAllPlatformTypesResponseBodyDataListUsageData) SetCode(v int32) *GetAllPlatformTypesResponseBodyDataListUsageData {
	s.Code = &v
	return s
}

func (s *GetAllPlatformTypesResponseBodyDataListUsageData) SetName(v string) *GetAllPlatformTypesResponseBodyDataListUsageData {
	s.Name = &v
	return s
}

type GetAllPlatformTypesResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAllPlatformTypesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAllPlatformTypesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAllPlatformTypesResponse) GoString() string {
	return s.String()
}

func (s *GetAllPlatformTypesResponse) SetHeaders(v map[string]*string) *GetAllPlatformTypesResponse {
	s.Headers = v
	return s
}

func (s *GetAllPlatformTypesResponse) SetBody(v *GetAllPlatformTypesResponseBody) *GetAllPlatformTypesResponse {
	s.Body = v
	return s
}

type GetAllRegionsRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty"`
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

type GetAllRegionsResponseBody struct {
	DataList  *GetAllRegionsResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Struct"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *int32                             `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetAllRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAllRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *GetAllRegionsResponseBody) SetDataList(v *GetAllRegionsResponseBodyDataList) *GetAllRegionsResponseBody {
	s.DataList = v
	return s
}

func (s *GetAllRegionsResponseBody) SetRequestId(v string) *GetAllRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAllRegionsResponseBody) SetCode(v int32) *GetAllRegionsResponseBody {
	s.Code = &v
	return s
}

type GetAllRegionsResponseBodyDataList struct {
	UsageData []*GetAllRegionsResponseBodyDataListUsageData `json:"UsageData,omitempty" xml:"UsageData,omitempty" type:"Repeated"`
}

func (s GetAllRegionsResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s GetAllRegionsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *GetAllRegionsResponseBodyDataList) SetUsageData(v []*GetAllRegionsResponseBodyDataListUsageData) *GetAllRegionsResponseBodyDataList {
	s.UsageData = v
	return s
}

type GetAllRegionsResponseBodyDataListUsageData struct {
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetAllRegionsResponseBodyDataListUsageData) String() string {
	return tea.Prettify(s)
}

func (s GetAllRegionsResponseBodyDataListUsageData) GoString() string {
	return s.String()
}

func (s *GetAllRegionsResponseBodyDataListUsageData) SetCode(v string) *GetAllRegionsResponseBodyDataListUsageData {
	s.Code = &v
	return s
}

func (s *GetAllRegionsResponseBodyDataListUsageData) SetName(v string) *GetAllRegionsResponseBodyDataListUsageData {
	s.Name = &v
	return s
}

type GetAllRegionsResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAllRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAllRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAllRegionsResponse) GoString() string {
	return s.String()
}

func (s *GetAllRegionsResponse) SetHeaders(v map[string]*string) *GetAllRegionsResponse {
	s.Headers = v
	return s
}

func (s *GetAllRegionsResponse) SetBody(v *GetAllRegionsResponseBody) *GetAllRegionsResponse {
	s.Body = v
	return s
}

type GetBalanceBandwidthDataRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty"`
	DataInterval  *int32  `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	ResourceId    *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
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

func (s *GetBalanceBandwidthDataRequest) SetDataInterval(v int32) *GetBalanceBandwidthDataRequest {
	s.DataInterval = &v
	return s
}

func (s *GetBalanceBandwidthDataRequest) SetResourceId(v string) *GetBalanceBandwidthDataRequest {
	s.ResourceId = &v
	return s
}

type GetBalanceBandwidthDataResponseBody struct {
	DataList  *GetBalanceBandwidthDataResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Struct"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Labels    *GetBalanceBandwidthDataResponseBodyLabels   `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Struct"`
	Code      *int32                                       `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetBalanceBandwidthDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetBalanceBandwidthDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetBalanceBandwidthDataResponseBody) SetDataList(v *GetBalanceBandwidthDataResponseBodyDataList) *GetBalanceBandwidthDataResponseBody {
	s.DataList = v
	return s
}

func (s *GetBalanceBandwidthDataResponseBody) SetRequestId(v string) *GetBalanceBandwidthDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBalanceBandwidthDataResponseBody) SetLabels(v *GetBalanceBandwidthDataResponseBodyLabels) *GetBalanceBandwidthDataResponseBody {
	s.Labels = v
	return s
}

func (s *GetBalanceBandwidthDataResponseBody) SetCode(v int32) *GetBalanceBandwidthDataResponseBody {
	s.Code = &v
	return s
}

type GetBalanceBandwidthDataResponseBodyDataList struct {
	UsageData []*GetBalanceBandwidthDataResponseBodyDataListUsageData `json:"UsageData,omitempty" xml:"UsageData,omitempty" type:"Repeated"`
}

func (s GetBalanceBandwidthDataResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s GetBalanceBandwidthDataResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *GetBalanceBandwidthDataResponseBodyDataList) SetUsageData(v []*GetBalanceBandwidthDataResponseBodyDataListUsageData) *GetBalanceBandwidthDataResponseBodyDataList {
	s.UsageData = v
	return s
}

type GetBalanceBandwidthDataResponseBodyDataListUsageData struct {
	Values *GetBalanceBandwidthDataResponseBodyDataListUsageDataValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Struct"`
	Date   *string                                                     `json:"Date,omitempty" xml:"Date,omitempty"`
}

func (s GetBalanceBandwidthDataResponseBodyDataListUsageData) String() string {
	return tea.Prettify(s)
}

func (s GetBalanceBandwidthDataResponseBodyDataListUsageData) GoString() string {
	return s.String()
}

func (s *GetBalanceBandwidthDataResponseBodyDataListUsageData) SetValues(v *GetBalanceBandwidthDataResponseBodyDataListUsageDataValues) *GetBalanceBandwidthDataResponseBodyDataListUsageData {
	s.Values = v
	return s
}

func (s *GetBalanceBandwidthDataResponseBodyDataListUsageData) SetDate(v string) *GetBalanceBandwidthDataResponseBodyDataListUsageData {
	s.Date = &v
	return s
}

type GetBalanceBandwidthDataResponseBodyDataListUsageDataValues struct {
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s GetBalanceBandwidthDataResponseBodyDataListUsageDataValues) String() string {
	return tea.Prettify(s)
}

func (s GetBalanceBandwidthDataResponseBodyDataListUsageDataValues) GoString() string {
	return s.String()
}

func (s *GetBalanceBandwidthDataResponseBodyDataListUsageDataValues) SetValues(v []*string) *GetBalanceBandwidthDataResponseBodyDataListUsageDataValues {
	s.Values = v
	return s
}

type GetBalanceBandwidthDataResponseBodyLabels struct {
	Label []*string `json:"Label,omitempty" xml:"Label,omitempty" type:"Repeated"`
}

func (s GetBalanceBandwidthDataResponseBodyLabels) String() string {
	return tea.Prettify(s)
}

func (s GetBalanceBandwidthDataResponseBodyLabels) GoString() string {
	return s.String()
}

func (s *GetBalanceBandwidthDataResponseBodyLabels) SetLabel(v []*string) *GetBalanceBandwidthDataResponseBodyLabels {
	s.Label = v
	return s
}

type GetBalanceBandwidthDataResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetBalanceBandwidthDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetBalanceBandwidthDataResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBalanceBandwidthDataResponse) GoString() string {
	return s.String()
}

func (s *GetBalanceBandwidthDataResponse) SetHeaders(v map[string]*string) *GetBalanceBandwidthDataResponse {
	s.Headers = v
	return s
}

func (s *GetBalanceBandwidthDataResponse) SetBody(v *GetBalanceBandwidthDataResponseBody) *GetBalanceBandwidthDataResponse {
	s.Body = v
	return s
}

type GetBalanceTrafficDataRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty"`
	DataInterval  *int32  `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	ResourceId    *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
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

func (s *GetBalanceTrafficDataRequest) SetDataInterval(v int32) *GetBalanceTrafficDataRequest {
	s.DataInterval = &v
	return s
}

func (s *GetBalanceTrafficDataRequest) SetResourceId(v string) *GetBalanceTrafficDataRequest {
	s.ResourceId = &v
	return s
}

type GetBalanceTrafficDataResponseBody struct {
	DataList  *GetBalanceTrafficDataResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Struct"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Labels    *GetBalanceTrafficDataResponseBodyLabels   `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Struct"`
	Code      *int32                                     `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetBalanceTrafficDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetBalanceTrafficDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetBalanceTrafficDataResponseBody) SetDataList(v *GetBalanceTrafficDataResponseBodyDataList) *GetBalanceTrafficDataResponseBody {
	s.DataList = v
	return s
}

func (s *GetBalanceTrafficDataResponseBody) SetRequestId(v string) *GetBalanceTrafficDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBalanceTrafficDataResponseBody) SetLabels(v *GetBalanceTrafficDataResponseBodyLabels) *GetBalanceTrafficDataResponseBody {
	s.Labels = v
	return s
}

func (s *GetBalanceTrafficDataResponseBody) SetCode(v int32) *GetBalanceTrafficDataResponseBody {
	s.Code = &v
	return s
}

type GetBalanceTrafficDataResponseBodyDataList struct {
	UsageData []*GetBalanceTrafficDataResponseBodyDataListUsageData `json:"UsageData,omitempty" xml:"UsageData,omitempty" type:"Repeated"`
}

func (s GetBalanceTrafficDataResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s GetBalanceTrafficDataResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *GetBalanceTrafficDataResponseBodyDataList) SetUsageData(v []*GetBalanceTrafficDataResponseBodyDataListUsageData) *GetBalanceTrafficDataResponseBodyDataList {
	s.UsageData = v
	return s
}

type GetBalanceTrafficDataResponseBodyDataListUsageData struct {
	Values *GetBalanceTrafficDataResponseBodyDataListUsageDataValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Struct"`
	Date   *string                                                   `json:"Date,omitempty" xml:"Date,omitempty"`
}

func (s GetBalanceTrafficDataResponseBodyDataListUsageData) String() string {
	return tea.Prettify(s)
}

func (s GetBalanceTrafficDataResponseBodyDataListUsageData) GoString() string {
	return s.String()
}

func (s *GetBalanceTrafficDataResponseBodyDataListUsageData) SetValues(v *GetBalanceTrafficDataResponseBodyDataListUsageDataValues) *GetBalanceTrafficDataResponseBodyDataListUsageData {
	s.Values = v
	return s
}

func (s *GetBalanceTrafficDataResponseBodyDataListUsageData) SetDate(v string) *GetBalanceTrafficDataResponseBodyDataListUsageData {
	s.Date = &v
	return s
}

type GetBalanceTrafficDataResponseBodyDataListUsageDataValues struct {
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s GetBalanceTrafficDataResponseBodyDataListUsageDataValues) String() string {
	return tea.Prettify(s)
}

func (s GetBalanceTrafficDataResponseBodyDataListUsageDataValues) GoString() string {
	return s.String()
}

func (s *GetBalanceTrafficDataResponseBodyDataListUsageDataValues) SetValues(v []*string) *GetBalanceTrafficDataResponseBodyDataListUsageDataValues {
	s.Values = v
	return s
}

type GetBalanceTrafficDataResponseBodyLabels struct {
	Label []*string `json:"Label,omitempty" xml:"Label,omitempty" type:"Repeated"`
}

func (s GetBalanceTrafficDataResponseBodyLabels) String() string {
	return tea.Prettify(s)
}

func (s GetBalanceTrafficDataResponseBodyLabels) GoString() string {
	return s.String()
}

func (s *GetBalanceTrafficDataResponseBodyLabels) SetLabel(v []*string) *GetBalanceTrafficDataResponseBodyLabels {
	s.Label = v
	return s
}

type GetBalanceTrafficDataResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetBalanceTrafficDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetBalanceTrafficDataResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBalanceTrafficDataResponse) GoString() string {
	return s.String()
}

func (s *GetBalanceTrafficDataResponse) SetHeaders(v map[string]*string) *GetBalanceTrafficDataResponse {
	s.Headers = v
	return s
}

func (s *GetBalanceTrafficDataResponse) SetBody(v *GetBalanceTrafficDataResponseBody) *GetBalanceTrafficDataResponse {
	s.Body = v
	return s
}

type GetBandwidthDataRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty"`
	Domain        *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Region        *string `json:"Region,omitempty" xml:"Region,omitempty"`
	IspName       *string `json:"IspName,omitempty" xml:"IspName,omitempty"`
	PlatformType  *string `json:"PlatformType,omitempty" xml:"PlatformType,omitempty"`
	BusinessType  *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	StartDate     *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	EndDate       *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
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

type GetBandwidthDataResponseBody struct {
	DataList  *GetBandwidthDataResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Struct"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Labels    *GetBandwidthDataResponseBodyLabels   `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Struct"`
	Code      *int32                                `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetBandwidthDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetBandwidthDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetBandwidthDataResponseBody) SetDataList(v *GetBandwidthDataResponseBodyDataList) *GetBandwidthDataResponseBody {
	s.DataList = v
	return s
}

func (s *GetBandwidthDataResponseBody) SetRequestId(v string) *GetBandwidthDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBandwidthDataResponseBody) SetLabels(v *GetBandwidthDataResponseBodyLabels) *GetBandwidthDataResponseBody {
	s.Labels = v
	return s
}

func (s *GetBandwidthDataResponseBody) SetCode(v int32) *GetBandwidthDataResponseBody {
	s.Code = &v
	return s
}

type GetBandwidthDataResponseBodyDataList struct {
	UsageData []*GetBandwidthDataResponseBodyDataListUsageData `json:"UsageData,omitempty" xml:"UsageData,omitempty" type:"Repeated"`
}

func (s GetBandwidthDataResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s GetBandwidthDataResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *GetBandwidthDataResponseBodyDataList) SetUsageData(v []*GetBandwidthDataResponseBodyDataListUsageData) *GetBandwidthDataResponseBodyDataList {
	s.UsageData = v
	return s
}

type GetBandwidthDataResponseBodyDataListUsageData struct {
	Values *GetBandwidthDataResponseBodyDataListUsageDataValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Struct"`
	Date   *string                                              `json:"Date,omitempty" xml:"Date,omitempty"`
}

func (s GetBandwidthDataResponseBodyDataListUsageData) String() string {
	return tea.Prettify(s)
}

func (s GetBandwidthDataResponseBodyDataListUsageData) GoString() string {
	return s.String()
}

func (s *GetBandwidthDataResponseBodyDataListUsageData) SetValues(v *GetBandwidthDataResponseBodyDataListUsageDataValues) *GetBandwidthDataResponseBodyDataListUsageData {
	s.Values = v
	return s
}

func (s *GetBandwidthDataResponseBodyDataListUsageData) SetDate(v string) *GetBandwidthDataResponseBodyDataListUsageData {
	s.Date = &v
	return s
}

type GetBandwidthDataResponseBodyDataListUsageDataValues struct {
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s GetBandwidthDataResponseBodyDataListUsageDataValues) String() string {
	return tea.Prettify(s)
}

func (s GetBandwidthDataResponseBodyDataListUsageDataValues) GoString() string {
	return s.String()
}

func (s *GetBandwidthDataResponseBodyDataListUsageDataValues) SetValues(v []*string) *GetBandwidthDataResponseBodyDataListUsageDataValues {
	s.Values = v
	return s
}

type GetBandwidthDataResponseBodyLabels struct {
	Label []*string `json:"Label,omitempty" xml:"Label,omitempty" type:"Repeated"`
}

func (s GetBandwidthDataResponseBodyLabels) String() string {
	return tea.Prettify(s)
}

func (s GetBandwidthDataResponseBodyLabels) GoString() string {
	return s.String()
}

func (s *GetBandwidthDataResponseBodyLabels) SetLabel(v []*string) *GetBandwidthDataResponseBodyLabels {
	s.Label = v
	return s
}

type GetBandwidthDataResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetBandwidthDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetBandwidthDataResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBandwidthDataResponse) GoString() string {
	return s.String()
}

func (s *GetBandwidthDataResponse) SetHeaders(v map[string]*string) *GetBandwidthDataResponse {
	s.Headers = v
	return s
}

func (s *GetBandwidthDataResponse) SetBody(v *GetBandwidthDataResponseBody) *GetBandwidthDataResponse {
	s.Body = v
	return s
}

type GetClientsRatioRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty"`
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

type GetClientsRatioResponseBody struct {
	DataList  *GetClientsRatioResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Struct"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *int32                               `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetClientsRatioResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetClientsRatioResponseBody) GoString() string {
	return s.String()
}

func (s *GetClientsRatioResponseBody) SetDataList(v *GetClientsRatioResponseBodyDataList) *GetClientsRatioResponseBody {
	s.DataList = v
	return s
}

func (s *GetClientsRatioResponseBody) SetRequestId(v string) *GetClientsRatioResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetClientsRatioResponseBody) SetCode(v int32) *GetClientsRatioResponseBody {
	s.Code = &v
	return s
}

type GetClientsRatioResponseBodyDataList struct {
	UsageData []*GetClientsRatioResponseBodyDataListUsageData `json:"UsageData,omitempty" xml:"UsageData,omitempty" type:"Repeated"`
}

func (s GetClientsRatioResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s GetClientsRatioResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *GetClientsRatioResponseBodyDataList) SetUsageData(v []*GetClientsRatioResponseBodyDataListUsageData) *GetClientsRatioResponseBodyDataList {
	s.UsageData = v
	return s
}

type GetClientsRatioResponseBodyDataListUsageData struct {
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Rate  *string `json:"Rate,omitempty" xml:"Rate,omitempty"`
}

func (s GetClientsRatioResponseBodyDataListUsageData) String() string {
	return tea.Prettify(s)
}

func (s GetClientsRatioResponseBodyDataListUsageData) GoString() string {
	return s.String()
}

func (s *GetClientsRatioResponseBodyDataListUsageData) SetValue(v string) *GetClientsRatioResponseBodyDataListUsageData {
	s.Value = &v
	return s
}

func (s *GetClientsRatioResponseBodyDataListUsageData) SetName(v string) *GetClientsRatioResponseBodyDataListUsageData {
	s.Name = &v
	return s
}

func (s *GetClientsRatioResponseBodyDataListUsageData) SetRate(v string) *GetClientsRatioResponseBodyDataListUsageData {
	s.Rate = &v
	return s
}

type GetClientsRatioResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetClientsRatioResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetClientsRatioResponse) String() string {
	return tea.Prettify(s)
}

func (s GetClientsRatioResponse) GoString() string {
	return s.String()
}

func (s *GetClientsRatioResponse) SetHeaders(v map[string]*string) *GetClientsRatioResponse {
	s.Headers = v
	return s
}

func (s *GetClientsRatioResponse) SetBody(v *GetClientsRatioResponseBody) *GetClientsRatioResponse {
	s.Body = v
	return s
}

type GetConsumerStatusRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty"`
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

type GetConsumerStatusResponseBody struct {
	Comment            *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	LiveMonitor        *bool   `json:"LiveMonitor,omitempty" xml:"LiveMonitor,omitempty"`
	Audit              *int32  `json:"Audit,omitempty" xml:"Audit,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	IntegreatedMode    *int32  `json:"IntegreatedMode,omitempty" xml:"IntegreatedMode,omitempty"`
	CreatedAt          *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	CdnUrlRedirectFlag *bool   `json:"CdnUrlRedirectFlag,omitempty" xml:"CdnUrlRedirectFlag,omitempty"`
	BusinessType       *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	Inservice          *bool   `json:"Inservice,omitempty" xml:"Inservice,omitempty"`
	RealtimeMonitor    *bool   `json:"RealtimeMonitor,omitempty" xml:"RealtimeMonitor,omitempty"`
	Code               *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	UpdatedAt          *string `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
}

func (s GetConsumerStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetConsumerStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetConsumerStatusResponseBody) SetComment(v string) *GetConsumerStatusResponseBody {
	s.Comment = &v
	return s
}

func (s *GetConsumerStatusResponseBody) SetLiveMonitor(v bool) *GetConsumerStatusResponseBody {
	s.LiveMonitor = &v
	return s
}

func (s *GetConsumerStatusResponseBody) SetAudit(v int32) *GetConsumerStatusResponseBody {
	s.Audit = &v
	return s
}

func (s *GetConsumerStatusResponseBody) SetRequestId(v string) *GetConsumerStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetConsumerStatusResponseBody) SetIntegreatedMode(v int32) *GetConsumerStatusResponseBody {
	s.IntegreatedMode = &v
	return s
}

func (s *GetConsumerStatusResponseBody) SetCreatedAt(v string) *GetConsumerStatusResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *GetConsumerStatusResponseBody) SetCdnUrlRedirectFlag(v bool) *GetConsumerStatusResponseBody {
	s.CdnUrlRedirectFlag = &v
	return s
}

func (s *GetConsumerStatusResponseBody) SetBusinessType(v string) *GetConsumerStatusResponseBody {
	s.BusinessType = &v
	return s
}

func (s *GetConsumerStatusResponseBody) SetInservice(v bool) *GetConsumerStatusResponseBody {
	s.Inservice = &v
	return s
}

func (s *GetConsumerStatusResponseBody) SetRealtimeMonitor(v bool) *GetConsumerStatusResponseBody {
	s.RealtimeMonitor = &v
	return s
}

func (s *GetConsumerStatusResponseBody) SetCode(v int32) *GetConsumerStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetConsumerStatusResponseBody) SetUpdatedAt(v string) *GetConsumerStatusResponseBody {
	s.UpdatedAt = &v
	return s
}

type GetConsumerStatusResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetConsumerStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetConsumerStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetConsumerStatusResponse) GoString() string {
	return s.String()
}

func (s *GetConsumerStatusResponse) SetHeaders(v map[string]*string) *GetConsumerStatusResponse {
	s.Headers = v
	return s
}

func (s *GetConsumerStatusResponse) SetBody(v *GetConsumerStatusResponseBody) *GetConsumerStatusResponse {
	s.Body = v
	return s
}

type GetControlRulesRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty"`
	Page          *string `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize      *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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

type GetControlRulesResponseBody struct {
	SettingList *GetControlRulesResponseBodySettingList `json:"SettingList,omitempty" xml:"SettingList,omitempty" type:"Struct"`
	RequestId   *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Pager       *GetControlRulesResponseBodyPager       `json:"Pager,omitempty" xml:"Pager,omitempty" type:"Struct"`
	Code        *int32                                  `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetControlRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetControlRulesResponseBody) GoString() string {
	return s.String()
}

func (s *GetControlRulesResponseBody) SetSettingList(v *GetControlRulesResponseBodySettingList) *GetControlRulesResponseBody {
	s.SettingList = v
	return s
}

func (s *GetControlRulesResponseBody) SetRequestId(v string) *GetControlRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetControlRulesResponseBody) SetPager(v *GetControlRulesResponseBodyPager) *GetControlRulesResponseBody {
	s.Pager = v
	return s
}

func (s *GetControlRulesResponseBody) SetCode(v int32) *GetControlRulesResponseBody {
	s.Code = &v
	return s
}

type GetControlRulesResponseBodySettingList struct {
	Setting []*GetControlRulesResponseBodySettingListSetting `json:"Setting,omitempty" xml:"Setting,omitempty" type:"Repeated"`
}

func (s GetControlRulesResponseBodySettingList) String() string {
	return tea.Prettify(s)
}

func (s GetControlRulesResponseBodySettingList) GoString() string {
	return s.String()
}

func (s *GetControlRulesResponseBodySettingList) SetSetting(v []*GetControlRulesResponseBodySettingListSetting) *GetControlRulesResponseBodySettingList {
	s.Setting = v
	return s
}

type GetControlRulesResponseBodySettingListSetting struct {
	CreatedAt    *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	ClientId     *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	Usable       *bool   `json:"Usable,omitempty" xml:"Usable,omitempty"`
	Region       *string `json:"Region,omitempty" xml:"Region,omitempty"`
	PlatformType *string `json:"PlatformType,omitempty" xml:"PlatformType,omitempty"`
	MarketType   *string `json:"MarketType,omitempty" xml:"MarketType,omitempty"`
	Onoff        *bool   `json:"Onoff,omitempty" xml:"Onoff,omitempty"`
	IspName      *string `json:"IspName,omitempty" xml:"IspName,omitempty"`
	AppVersion   *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	UpdatedAt    *string `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s GetControlRulesResponseBodySettingListSetting) String() string {
	return tea.Prettify(s)
}

func (s GetControlRulesResponseBodySettingListSetting) GoString() string {
	return s.String()
}

func (s *GetControlRulesResponseBodySettingListSetting) SetCreatedAt(v string) *GetControlRulesResponseBodySettingListSetting {
	s.CreatedAt = &v
	return s
}

func (s *GetControlRulesResponseBodySettingListSetting) SetClientId(v string) *GetControlRulesResponseBodySettingListSetting {
	s.ClientId = &v
	return s
}

func (s *GetControlRulesResponseBodySettingListSetting) SetBusinessType(v string) *GetControlRulesResponseBodySettingListSetting {
	s.BusinessType = &v
	return s
}

func (s *GetControlRulesResponseBodySettingListSetting) SetUsable(v bool) *GetControlRulesResponseBodySettingListSetting {
	s.Usable = &v
	return s
}

func (s *GetControlRulesResponseBodySettingListSetting) SetRegion(v string) *GetControlRulesResponseBodySettingListSetting {
	s.Region = &v
	return s
}

func (s *GetControlRulesResponseBodySettingListSetting) SetPlatformType(v string) *GetControlRulesResponseBodySettingListSetting {
	s.PlatformType = &v
	return s
}

func (s *GetControlRulesResponseBodySettingListSetting) SetMarketType(v string) *GetControlRulesResponseBodySettingListSetting {
	s.MarketType = &v
	return s
}

func (s *GetControlRulesResponseBodySettingListSetting) SetOnoff(v bool) *GetControlRulesResponseBodySettingListSetting {
	s.Onoff = &v
	return s
}

func (s *GetControlRulesResponseBodySettingListSetting) SetIspName(v string) *GetControlRulesResponseBodySettingListSetting {
	s.IspName = &v
	return s
}

func (s *GetControlRulesResponseBodySettingListSetting) SetAppVersion(v string) *GetControlRulesResponseBodySettingListSetting {
	s.AppVersion = &v
	return s
}

func (s *GetControlRulesResponseBodySettingListSetting) SetUpdatedAt(v string) *GetControlRulesResponseBodySettingListSetting {
	s.UpdatedAt = &v
	return s
}

func (s *GetControlRulesResponseBodySettingListSetting) SetName(v string) *GetControlRulesResponseBodySettingListSetting {
	s.Name = &v
	return s
}

func (s *GetControlRulesResponseBodySettingListSetting) SetResourceId(v string) *GetControlRulesResponseBodySettingListSetting {
	s.ResourceId = &v
	return s
}

type GetControlRulesResponseBodyPager struct {
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total    *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
	Page     *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
}

func (s GetControlRulesResponseBodyPager) String() string {
	return tea.Prettify(s)
}

func (s GetControlRulesResponseBodyPager) GoString() string {
	return s.String()
}

func (s *GetControlRulesResponseBodyPager) SetPageSize(v int32) *GetControlRulesResponseBodyPager {
	s.PageSize = &v
	return s
}

func (s *GetControlRulesResponseBodyPager) SetTotal(v int32) *GetControlRulesResponseBodyPager {
	s.Total = &v
	return s
}

func (s *GetControlRulesResponseBodyPager) SetPage(v int32) *GetControlRulesResponseBodyPager {
	s.Page = &v
	return s
}

type GetControlRulesResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetControlRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetControlRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetControlRulesResponse) GoString() string {
	return s.String()
}

func (s *GetControlRulesResponse) SetHeaders(v map[string]*string) *GetControlRulesResponse {
	s.Headers = v
	return s
}

func (s *GetControlRulesResponse) SetBody(v *GetControlRulesResponseBody) *GetControlRulesResponse {
	s.Body = v
	return s
}

type GetCoverRateDataRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty"`
	Domain        *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Region        *string `json:"Region,omitempty" xml:"Region,omitempty"`
	IspName       *string `json:"IspName,omitempty" xml:"IspName,omitempty"`
	PlatformType  *string `json:"PlatformType,omitempty" xml:"PlatformType,omitempty"`
	BusinessType  *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	StartDate     *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	EndDate       *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
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

type GetCoverRateDataResponseBody struct {
	DataList  *GetCoverRateDataResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Struct"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Labels    *GetCoverRateDataResponseBodyLabels   `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Struct"`
	Code      *int32                                `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetCoverRateDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCoverRateDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetCoverRateDataResponseBody) SetDataList(v *GetCoverRateDataResponseBodyDataList) *GetCoverRateDataResponseBody {
	s.DataList = v
	return s
}

func (s *GetCoverRateDataResponseBody) SetRequestId(v string) *GetCoverRateDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCoverRateDataResponseBody) SetLabels(v *GetCoverRateDataResponseBodyLabels) *GetCoverRateDataResponseBody {
	s.Labels = v
	return s
}

func (s *GetCoverRateDataResponseBody) SetCode(v int32) *GetCoverRateDataResponseBody {
	s.Code = &v
	return s
}

type GetCoverRateDataResponseBodyDataList struct {
	UsageData []*GetCoverRateDataResponseBodyDataListUsageData `json:"UsageData,omitempty" xml:"UsageData,omitempty" type:"Repeated"`
}

func (s GetCoverRateDataResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s GetCoverRateDataResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *GetCoverRateDataResponseBodyDataList) SetUsageData(v []*GetCoverRateDataResponseBodyDataListUsageData) *GetCoverRateDataResponseBodyDataList {
	s.UsageData = v
	return s
}

type GetCoverRateDataResponseBodyDataListUsageData struct {
	Values *GetCoverRateDataResponseBodyDataListUsageDataValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Struct"`
	Date   *string                                              `json:"Date,omitempty" xml:"Date,omitempty"`
}

func (s GetCoverRateDataResponseBodyDataListUsageData) String() string {
	return tea.Prettify(s)
}

func (s GetCoverRateDataResponseBodyDataListUsageData) GoString() string {
	return s.String()
}

func (s *GetCoverRateDataResponseBodyDataListUsageData) SetValues(v *GetCoverRateDataResponseBodyDataListUsageDataValues) *GetCoverRateDataResponseBodyDataListUsageData {
	s.Values = v
	return s
}

func (s *GetCoverRateDataResponseBodyDataListUsageData) SetDate(v string) *GetCoverRateDataResponseBodyDataListUsageData {
	s.Date = &v
	return s
}

type GetCoverRateDataResponseBodyDataListUsageDataValues struct {
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s GetCoverRateDataResponseBodyDataListUsageDataValues) String() string {
	return tea.Prettify(s)
}

func (s GetCoverRateDataResponseBodyDataListUsageDataValues) GoString() string {
	return s.String()
}

func (s *GetCoverRateDataResponseBodyDataListUsageDataValues) SetValues(v []*string) *GetCoverRateDataResponseBodyDataListUsageDataValues {
	s.Values = v
	return s
}

type GetCoverRateDataResponseBodyLabels struct {
	Label []*string `json:"Label,omitempty" xml:"Label,omitempty" type:"Repeated"`
}

func (s GetCoverRateDataResponseBodyLabels) String() string {
	return tea.Prettify(s)
}

func (s GetCoverRateDataResponseBodyLabels) GoString() string {
	return s.String()
}

func (s *GetCoverRateDataResponseBodyLabels) SetLabel(v []*string) *GetCoverRateDataResponseBodyLabels {
	s.Label = v
	return s
}

type GetCoverRateDataResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetCoverRateDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCoverRateDataResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCoverRateDataResponse) GoString() string {
	return s.String()
}

func (s *GetCoverRateDataResponse) SetHeaders(v map[string]*string) *GetCoverRateDataResponse {
	s.Headers = v
	return s
}

func (s *GetCoverRateDataResponse) SetBody(v *GetCoverRateDataResponseBody) *GetCoverRateDataResponse {
	s.Body = v
	return s
}

type GetCurrentModeRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty"`
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

type GetCurrentModeResponseBody struct {
	ModeCode          *int32  `json:"ModeCode,omitempty" xml:"ModeCode,omitempty"`
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PaddingModeCode   *int32  `json:"PaddingModeCode,omitempty" xml:"PaddingModeCode,omitempty"`
	EffectiveAt       *int32  `json:"EffectiveAt,omitempty" xml:"EffectiveAt,omitempty"`
	EstimateBandwidth *int32  `json:"EstimateBandwidth,omitempty" xml:"EstimateBandwidth,omitempty"`
	Code              *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetCurrentModeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCurrentModeResponseBody) GoString() string {
	return s.String()
}

func (s *GetCurrentModeResponseBody) SetModeCode(v int32) *GetCurrentModeResponseBody {
	s.ModeCode = &v
	return s
}

func (s *GetCurrentModeResponseBody) SetRequestId(v string) *GetCurrentModeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCurrentModeResponseBody) SetPaddingModeCode(v int32) *GetCurrentModeResponseBody {
	s.PaddingModeCode = &v
	return s
}

func (s *GetCurrentModeResponseBody) SetEffectiveAt(v int32) *GetCurrentModeResponseBody {
	s.EffectiveAt = &v
	return s
}

func (s *GetCurrentModeResponseBody) SetEstimateBandwidth(v int32) *GetCurrentModeResponseBody {
	s.EstimateBandwidth = &v
	return s
}

func (s *GetCurrentModeResponseBody) SetCode(v int32) *GetCurrentModeResponseBody {
	s.Code = &v
	return s
}

type GetCurrentModeResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetCurrentModeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCurrentModeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCurrentModeResponse) GoString() string {
	return s.String()
}

func (s *GetCurrentModeResponse) SetHeaders(v map[string]*string) *GetCurrentModeResponse {
	s.Headers = v
	return s
}

func (s *GetCurrentModeResponse) SetBody(v *GetCurrentModeResponseBody) *GetCurrentModeResponse {
	s.Body = v
	return s
}

type GetDomainCountRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty"`
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

type GetDomainCountResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *int32  `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetDomainCountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDomainCountResponseBody) GoString() string {
	return s.String()
}

func (s *GetDomainCountResponseBody) SetRequestId(v string) *GetDomainCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDomainCountResponseBody) SetData(v int32) *GetDomainCountResponseBody {
	s.Data = &v
	return s
}

func (s *GetDomainCountResponseBody) SetCode(v int32) *GetDomainCountResponseBody {
	s.Code = &v
	return s
}

type GetDomainCountResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDomainCountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDomainCountResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDomainCountResponse) GoString() string {
	return s.String()
}

func (s *GetDomainCountResponse) SetHeaders(v map[string]*string) *GetDomainCountResponse {
	s.Headers = v
	return s
}

func (s *GetDomainCountResponse) SetBody(v *GetDomainCountResponseBody) *GetDomainCountResponse {
	s.Body = v
	return s
}

type GetDomainsRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty"`
	Page          *string `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize      *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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

type GetDomainsResponseBody struct {
	DataList  *GetDomainsResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Struct"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Pager     *GetDomainsResponseBodyPager    `json:"Pager,omitempty" xml:"Pager,omitempty" type:"Struct"`
	Code      *int32                          `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetDomainsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *GetDomainsResponseBody) SetDataList(v *GetDomainsResponseBodyDataList) *GetDomainsResponseBody {
	s.DataList = v
	return s
}

func (s *GetDomainsResponseBody) SetRequestId(v string) *GetDomainsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDomainsResponseBody) SetPager(v *GetDomainsResponseBodyPager) *GetDomainsResponseBody {
	s.Pager = v
	return s
}

func (s *GetDomainsResponseBody) SetCode(v int32) *GetDomainsResponseBody {
	s.Code = &v
	return s
}

type GetDomainsResponseBodyDataList struct {
	UsageData []*GetDomainsResponseBodyDataListUsageData `json:"UsageData,omitempty" xml:"UsageData,omitempty" type:"Repeated"`
}

func (s GetDomainsResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s GetDomainsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *GetDomainsResponseBodyDataList) SetUsageData(v []*GetDomainsResponseBodyDataListUsageData) *GetDomainsResponseBodyDataList {
	s.UsageData = v
	return s
}

type GetDomainsResponseBodyDataListUsageData struct {
	Status       *bool   `json:"Status,omitempty" xml:"Status,omitempty"`
	Domain       *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	SliceFormat  *string `json:"SliceFormat,omitempty" xml:"SliceFormat,omitempty"`
	CreatedAt    *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	UpdatedAt    *string `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
}

func (s GetDomainsResponseBodyDataListUsageData) String() string {
	return tea.Prettify(s)
}

func (s GetDomainsResponseBodyDataListUsageData) GoString() string {
	return s.String()
}

func (s *GetDomainsResponseBodyDataListUsageData) SetStatus(v bool) *GetDomainsResponseBodyDataListUsageData {
	s.Status = &v
	return s
}

func (s *GetDomainsResponseBodyDataListUsageData) SetDomain(v string) *GetDomainsResponseBodyDataListUsageData {
	s.Domain = &v
	return s
}

func (s *GetDomainsResponseBodyDataListUsageData) SetSliceFormat(v string) *GetDomainsResponseBodyDataListUsageData {
	s.SliceFormat = &v
	return s
}

func (s *GetDomainsResponseBodyDataListUsageData) SetCreatedAt(v string) *GetDomainsResponseBodyDataListUsageData {
	s.CreatedAt = &v
	return s
}

func (s *GetDomainsResponseBodyDataListUsageData) SetUpdatedAt(v string) *GetDomainsResponseBodyDataListUsageData {
	s.UpdatedAt = &v
	return s
}

func (s *GetDomainsResponseBodyDataListUsageData) SetResourceId(v string) *GetDomainsResponseBodyDataListUsageData {
	s.ResourceId = &v
	return s
}

func (s *GetDomainsResponseBodyDataListUsageData) SetBusinessType(v string) *GetDomainsResponseBodyDataListUsageData {
	s.BusinessType = &v
	return s
}

type GetDomainsResponseBodyPager struct {
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total    *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
	Page     *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
}

func (s GetDomainsResponseBodyPager) String() string {
	return tea.Prettify(s)
}

func (s GetDomainsResponseBodyPager) GoString() string {
	return s.String()
}

func (s *GetDomainsResponseBodyPager) SetPageSize(v int32) *GetDomainsResponseBodyPager {
	s.PageSize = &v
	return s
}

func (s *GetDomainsResponseBodyPager) SetTotal(v int32) *GetDomainsResponseBodyPager {
	s.Total = &v
	return s
}

func (s *GetDomainsResponseBodyPager) SetPage(v int32) *GetDomainsResponseBodyPager {
	s.Page = &v
	return s
}

type GetDomainsResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDomainsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDomainsResponse) GoString() string {
	return s.String()
}

func (s *GetDomainsResponse) SetHeaders(v map[string]*string) *GetDomainsResponse {
	s.Headers = v
	return s
}

func (s *GetDomainsResponse) SetBody(v *GetDomainsResponseBody) *GetDomainsResponse {
	s.Body = v
	return s
}

type GetExpenseSummaryRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty"`
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

type GetExpenseSummaryResponseBody struct {
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *GetExpenseSummaryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *int32                             `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetExpenseSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetExpenseSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetExpenseSummaryResponseBody) SetRequestId(v string) *GetExpenseSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetExpenseSummaryResponseBody) SetData(v *GetExpenseSummaryResponseBodyData) *GetExpenseSummaryResponseBody {
	s.Data = v
	return s
}

func (s *GetExpenseSummaryResponseBody) SetCode(v int32) *GetExpenseSummaryResponseBody {
	s.Code = &v
	return s
}

type GetExpenseSummaryResponseBodyData struct {
	ForecastFluency *float32 `json:"ForecastFluency,omitempty" xml:"ForecastFluency,omitempty"`
	TopBandwidth    *int64   `json:"TopBandwidth,omitempty" xml:"TopBandwidth,omitempty"`
	TotalTraffic    *int64   `json:"TotalTraffic,omitempty" xml:"TotalTraffic,omitempty"`
	CoverRate       *float32 `json:"CoverRate,omitempty" xml:"CoverRate,omitempty"`
	ShareRate       *float32 `json:"ShareRate,omitempty" xml:"ShareRate,omitempty"`
	TotalUV         *int32   `json:"TotalUV,omitempty" xml:"TotalUV,omitempty"`
}

func (s GetExpenseSummaryResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetExpenseSummaryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetExpenseSummaryResponseBodyData) SetForecastFluency(v float32) *GetExpenseSummaryResponseBodyData {
	s.ForecastFluency = &v
	return s
}

func (s *GetExpenseSummaryResponseBodyData) SetTopBandwidth(v int64) *GetExpenseSummaryResponseBodyData {
	s.TopBandwidth = &v
	return s
}

func (s *GetExpenseSummaryResponseBodyData) SetTotalTraffic(v int64) *GetExpenseSummaryResponseBodyData {
	s.TotalTraffic = &v
	return s
}

func (s *GetExpenseSummaryResponseBodyData) SetCoverRate(v float32) *GetExpenseSummaryResponseBodyData {
	s.CoverRate = &v
	return s
}

func (s *GetExpenseSummaryResponseBodyData) SetShareRate(v float32) *GetExpenseSummaryResponseBodyData {
	s.ShareRate = &v
	return s
}

func (s *GetExpenseSummaryResponseBodyData) SetTotalUV(v int32) *GetExpenseSummaryResponseBodyData {
	s.TotalUV = &v
	return s
}

type GetExpenseSummaryResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetExpenseSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetExpenseSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetExpenseSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetExpenseSummaryResponse) SetHeaders(v map[string]*string) *GetExpenseSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetExpenseSummaryResponse) SetBody(v *GetExpenseSummaryResponseBody) *GetExpenseSummaryResponse {
	s.Body = v
	return s
}

type GetFeeHistoryRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty"`
	Page          *string `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize      *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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

type GetFeeHistoryResponseBody struct {
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	FeeList   *GetFeeHistoryResponseBodyFeeList `json:"FeeList,omitempty" xml:"FeeList,omitempty" type:"Struct"`
	Pager     *GetFeeHistoryResponseBodyPager   `json:"Pager,omitempty" xml:"Pager,omitempty" type:"Struct"`
	Code      *int32                            `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetFeeHistoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFeeHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *GetFeeHistoryResponseBody) SetRequestId(v string) *GetFeeHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFeeHistoryResponseBody) SetFeeList(v *GetFeeHistoryResponseBodyFeeList) *GetFeeHistoryResponseBody {
	s.FeeList = v
	return s
}

func (s *GetFeeHistoryResponseBody) SetPager(v *GetFeeHistoryResponseBodyPager) *GetFeeHistoryResponseBody {
	s.Pager = v
	return s
}

func (s *GetFeeHistoryResponseBody) SetCode(v int32) *GetFeeHistoryResponseBody {
	s.Code = &v
	return s
}

type GetFeeHistoryResponseBodyFeeList struct {
	Fee []*GetFeeHistoryResponseBodyFeeListFee `json:"Fee,omitempty" xml:"Fee,omitempty" type:"Repeated"`
}

func (s GetFeeHistoryResponseBodyFeeList) String() string {
	return tea.Prettify(s)
}

func (s GetFeeHistoryResponseBodyFeeList) GoString() string {
	return s.String()
}

func (s *GetFeeHistoryResponseBodyFeeList) SetFee(v []*GetFeeHistoryResponseBodyFeeListFee) *GetFeeHistoryResponseBodyFeeList {
	s.Fee = v
	return s
}

type GetFeeHistoryResponseBodyFeeListFee struct {
	EndDate             *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	TimeSpan            *string `json:"TimeSpan,omitempty" xml:"TimeSpan,omitempty"`
	Date                *string `json:"Date,omitempty" xml:"Date,omitempty"`
	StartDate           *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	LevelThreeTraffic   *int32  `json:"LevelThreeTraffic,omitempty" xml:"LevelThreeTraffic,omitempty"`
	Mode                *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	TotalTraffic        *int32  `json:"TotalTraffic,omitempty" xml:"TotalTraffic,omitempty"`
	BusinessType        *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	LevelTwoTraffic     *int32  `json:"LevelTwoTraffic,omitempty" xml:"LevelTwoTraffic,omitempty"`
	LevelThreeBandwidth *int32  `json:"LevelThreeBandwidth,omitempty" xml:"LevelThreeBandwidth,omitempty"`
	LevelTwoBandwidth   *int32  `json:"LevelTwoBandwidth,omitempty" xml:"LevelTwoBandwidth,omitempty"`
	FlowOut             *int32  `json:"FlowOut,omitempty" xml:"FlowOut,omitempty"`
	ResourceId          *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	TotalBandwidth      *int32  `json:"TotalBandwidth,omitempty" xml:"TotalBandwidth,omitempty"`
}

func (s GetFeeHistoryResponseBodyFeeListFee) String() string {
	return tea.Prettify(s)
}

func (s GetFeeHistoryResponseBodyFeeListFee) GoString() string {
	return s.String()
}

func (s *GetFeeHistoryResponseBodyFeeListFee) SetEndDate(v string) *GetFeeHistoryResponseBodyFeeListFee {
	s.EndDate = &v
	return s
}

func (s *GetFeeHistoryResponseBodyFeeListFee) SetTimeSpan(v string) *GetFeeHistoryResponseBodyFeeListFee {
	s.TimeSpan = &v
	return s
}

func (s *GetFeeHistoryResponseBodyFeeListFee) SetDate(v string) *GetFeeHistoryResponseBodyFeeListFee {
	s.Date = &v
	return s
}

func (s *GetFeeHistoryResponseBodyFeeListFee) SetStartDate(v string) *GetFeeHistoryResponseBodyFeeListFee {
	s.StartDate = &v
	return s
}

func (s *GetFeeHistoryResponseBodyFeeListFee) SetLevelThreeTraffic(v int32) *GetFeeHistoryResponseBodyFeeListFee {
	s.LevelThreeTraffic = &v
	return s
}

func (s *GetFeeHistoryResponseBodyFeeListFee) SetMode(v string) *GetFeeHistoryResponseBodyFeeListFee {
	s.Mode = &v
	return s
}

func (s *GetFeeHistoryResponseBodyFeeListFee) SetTotalTraffic(v int32) *GetFeeHistoryResponseBodyFeeListFee {
	s.TotalTraffic = &v
	return s
}

func (s *GetFeeHistoryResponseBodyFeeListFee) SetBusinessType(v string) *GetFeeHistoryResponseBodyFeeListFee {
	s.BusinessType = &v
	return s
}

func (s *GetFeeHistoryResponseBodyFeeListFee) SetLevelTwoTraffic(v int32) *GetFeeHistoryResponseBodyFeeListFee {
	s.LevelTwoTraffic = &v
	return s
}

func (s *GetFeeHistoryResponseBodyFeeListFee) SetLevelThreeBandwidth(v int32) *GetFeeHistoryResponseBodyFeeListFee {
	s.LevelThreeBandwidth = &v
	return s
}

func (s *GetFeeHistoryResponseBodyFeeListFee) SetLevelTwoBandwidth(v int32) *GetFeeHistoryResponseBodyFeeListFee {
	s.LevelTwoBandwidth = &v
	return s
}

func (s *GetFeeHistoryResponseBodyFeeListFee) SetFlowOut(v int32) *GetFeeHistoryResponseBodyFeeListFee {
	s.FlowOut = &v
	return s
}

func (s *GetFeeHistoryResponseBodyFeeListFee) SetResourceId(v string) *GetFeeHistoryResponseBodyFeeListFee {
	s.ResourceId = &v
	return s
}

func (s *GetFeeHistoryResponseBodyFeeListFee) SetTotalBandwidth(v int32) *GetFeeHistoryResponseBodyFeeListFee {
	s.TotalBandwidth = &v
	return s
}

type GetFeeHistoryResponseBodyPager struct {
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total    *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
	Page     *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
}

func (s GetFeeHistoryResponseBodyPager) String() string {
	return tea.Prettify(s)
}

func (s GetFeeHistoryResponseBodyPager) GoString() string {
	return s.String()
}

func (s *GetFeeHistoryResponseBodyPager) SetPageSize(v int32) *GetFeeHistoryResponseBodyPager {
	s.PageSize = &v
	return s
}

func (s *GetFeeHistoryResponseBodyPager) SetTotal(v int32) *GetFeeHistoryResponseBodyPager {
	s.Total = &v
	return s
}

func (s *GetFeeHistoryResponseBodyPager) SetPage(v int32) *GetFeeHistoryResponseBodyPager {
	s.Page = &v
	return s
}

type GetFeeHistoryResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetFeeHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetFeeHistoryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFeeHistoryResponse) GoString() string {
	return s.String()
}

func (s *GetFeeHistoryResponse) SetHeaders(v map[string]*string) *GetFeeHistoryResponse {
	s.Headers = v
	return s
}

func (s *GetFeeHistoryResponse) SetBody(v *GetFeeHistoryResponseBody) *GetFeeHistoryResponse {
	s.Body = v
	return s
}

type GetFirstFrameDelayDataRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty"`
	Domain        *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Region        *string `json:"Region,omitempty" xml:"Region,omitempty"`
	IspName       *string `json:"IspName,omitempty" xml:"IspName,omitempty"`
	PlatformType  *string `json:"PlatformType,omitempty" xml:"PlatformType,omitempty"`
	BusinessType  *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	StartDate     *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	EndDate       *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
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

type GetFirstFrameDelayDataResponseBody struct {
	DataList  *GetFirstFrameDelayDataResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Struct"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Labels    *GetFirstFrameDelayDataResponseBodyLabels   `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Struct"`
	Code      *int32                                      `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetFirstFrameDelayDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFirstFrameDelayDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetFirstFrameDelayDataResponseBody) SetDataList(v *GetFirstFrameDelayDataResponseBodyDataList) *GetFirstFrameDelayDataResponseBody {
	s.DataList = v
	return s
}

func (s *GetFirstFrameDelayDataResponseBody) SetRequestId(v string) *GetFirstFrameDelayDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFirstFrameDelayDataResponseBody) SetLabels(v *GetFirstFrameDelayDataResponseBodyLabels) *GetFirstFrameDelayDataResponseBody {
	s.Labels = v
	return s
}

func (s *GetFirstFrameDelayDataResponseBody) SetCode(v int32) *GetFirstFrameDelayDataResponseBody {
	s.Code = &v
	return s
}

type GetFirstFrameDelayDataResponseBodyDataList struct {
	UsageData []*GetFirstFrameDelayDataResponseBodyDataListUsageData `json:"UsageData,omitempty" xml:"UsageData,omitempty" type:"Repeated"`
}

func (s GetFirstFrameDelayDataResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s GetFirstFrameDelayDataResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *GetFirstFrameDelayDataResponseBodyDataList) SetUsageData(v []*GetFirstFrameDelayDataResponseBodyDataListUsageData) *GetFirstFrameDelayDataResponseBodyDataList {
	s.UsageData = v
	return s
}

type GetFirstFrameDelayDataResponseBodyDataListUsageData struct {
	Values *GetFirstFrameDelayDataResponseBodyDataListUsageDataValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Struct"`
	Date   *string                                                    `json:"Date,omitempty" xml:"Date,omitempty"`
}

func (s GetFirstFrameDelayDataResponseBodyDataListUsageData) String() string {
	return tea.Prettify(s)
}

func (s GetFirstFrameDelayDataResponseBodyDataListUsageData) GoString() string {
	return s.String()
}

func (s *GetFirstFrameDelayDataResponseBodyDataListUsageData) SetValues(v *GetFirstFrameDelayDataResponseBodyDataListUsageDataValues) *GetFirstFrameDelayDataResponseBodyDataListUsageData {
	s.Values = v
	return s
}

func (s *GetFirstFrameDelayDataResponseBodyDataListUsageData) SetDate(v string) *GetFirstFrameDelayDataResponseBodyDataListUsageData {
	s.Date = &v
	return s
}

type GetFirstFrameDelayDataResponseBodyDataListUsageDataValues struct {
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s GetFirstFrameDelayDataResponseBodyDataListUsageDataValues) String() string {
	return tea.Prettify(s)
}

func (s GetFirstFrameDelayDataResponseBodyDataListUsageDataValues) GoString() string {
	return s.String()
}

func (s *GetFirstFrameDelayDataResponseBodyDataListUsageDataValues) SetValues(v []*string) *GetFirstFrameDelayDataResponseBodyDataListUsageDataValues {
	s.Values = v
	return s
}

type GetFirstFrameDelayDataResponseBodyLabels struct {
	Label []*string `json:"Label,omitempty" xml:"Label,omitempty" type:"Repeated"`
}

func (s GetFirstFrameDelayDataResponseBodyLabels) String() string {
	return tea.Prettify(s)
}

func (s GetFirstFrameDelayDataResponseBodyLabels) GoString() string {
	return s.String()
}

func (s *GetFirstFrameDelayDataResponseBodyLabels) SetLabel(v []*string) *GetFirstFrameDelayDataResponseBodyLabels {
	s.Label = v
	return s
}

type GetFirstFrameDelayDataResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetFirstFrameDelayDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetFirstFrameDelayDataResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFirstFrameDelayDataResponse) GoString() string {
	return s.String()
}

func (s *GetFirstFrameDelayDataResponse) SetHeaders(v map[string]*string) *GetFirstFrameDelayDataResponse {
	s.Headers = v
	return s
}

func (s *GetFirstFrameDelayDataResponse) SetBody(v *GetFirstFrameDelayDataResponseBody) *GetFirstFrameDelayDataResponse {
	s.Body = v
	return s
}

type GetFluencyDataRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty"`
	Domain        *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Region        *string `json:"Region,omitempty" xml:"Region,omitempty"`
	IspName       *string `json:"IspName,omitempty" xml:"IspName,omitempty"`
	PlatformType  *string `json:"PlatformType,omitempty" xml:"PlatformType,omitempty"`
	BusinessType  *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	StartDate     *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	EndDate       *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
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

type GetFluencyDataResponseBody struct {
	DataList  *GetFluencyDataResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Struct"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Labels    *GetFluencyDataResponseBodyLabels   `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Struct"`
	Code      *int32                              `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetFluencyDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFluencyDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetFluencyDataResponseBody) SetDataList(v *GetFluencyDataResponseBodyDataList) *GetFluencyDataResponseBody {
	s.DataList = v
	return s
}

func (s *GetFluencyDataResponseBody) SetRequestId(v string) *GetFluencyDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFluencyDataResponseBody) SetLabels(v *GetFluencyDataResponseBodyLabels) *GetFluencyDataResponseBody {
	s.Labels = v
	return s
}

func (s *GetFluencyDataResponseBody) SetCode(v int32) *GetFluencyDataResponseBody {
	s.Code = &v
	return s
}

type GetFluencyDataResponseBodyDataList struct {
	UsageData []*GetFluencyDataResponseBodyDataListUsageData `json:"UsageData,omitempty" xml:"UsageData,omitempty" type:"Repeated"`
}

func (s GetFluencyDataResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s GetFluencyDataResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *GetFluencyDataResponseBodyDataList) SetUsageData(v []*GetFluencyDataResponseBodyDataListUsageData) *GetFluencyDataResponseBodyDataList {
	s.UsageData = v
	return s
}

type GetFluencyDataResponseBodyDataListUsageData struct {
	Values *GetFluencyDataResponseBodyDataListUsageDataValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Struct"`
	Date   *string                                            `json:"Date,omitempty" xml:"Date,omitempty"`
}

func (s GetFluencyDataResponseBodyDataListUsageData) String() string {
	return tea.Prettify(s)
}

func (s GetFluencyDataResponseBodyDataListUsageData) GoString() string {
	return s.String()
}

func (s *GetFluencyDataResponseBodyDataListUsageData) SetValues(v *GetFluencyDataResponseBodyDataListUsageDataValues) *GetFluencyDataResponseBodyDataListUsageData {
	s.Values = v
	return s
}

func (s *GetFluencyDataResponseBodyDataListUsageData) SetDate(v string) *GetFluencyDataResponseBodyDataListUsageData {
	s.Date = &v
	return s
}

type GetFluencyDataResponseBodyDataListUsageDataValues struct {
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s GetFluencyDataResponseBodyDataListUsageDataValues) String() string {
	return tea.Prettify(s)
}

func (s GetFluencyDataResponseBodyDataListUsageDataValues) GoString() string {
	return s.String()
}

func (s *GetFluencyDataResponseBodyDataListUsageDataValues) SetValues(v []*string) *GetFluencyDataResponseBodyDataListUsageDataValues {
	s.Values = v
	return s
}

type GetFluencyDataResponseBodyLabels struct {
	Label []*string `json:"Label,omitempty" xml:"Label,omitempty" type:"Repeated"`
}

func (s GetFluencyDataResponseBodyLabels) String() string {
	return tea.Prettify(s)
}

func (s GetFluencyDataResponseBodyLabels) GoString() string {
	return s.String()
}

func (s *GetFluencyDataResponseBodyLabels) SetLabel(v []*string) *GetFluencyDataResponseBodyLabels {
	s.Label = v
	return s
}

type GetFluencyDataResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetFluencyDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetFluencyDataResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFluencyDataResponse) GoString() string {
	return s.String()
}

func (s *GetFluencyDataResponse) SetHeaders(v map[string]*string) *GetFluencyDataResponse {
	s.Headers = v
	return s
}

func (s *GetFluencyDataResponse) SetBody(v *GetFluencyDataResponseBody) *GetFluencyDataResponse {
	s.Body = v
	return s
}

type GetLogsListRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty"`
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

type GetLogsListResponseBody struct {
	LogList   *GetLogsListResponseBodyLogList `json:"LogList,omitempty" xml:"LogList,omitempty" type:"Struct"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *int32                          `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetLogsListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLogsListResponseBody) GoString() string {
	return s.String()
}

func (s *GetLogsListResponseBody) SetLogList(v *GetLogsListResponseBodyLogList) *GetLogsListResponseBody {
	s.LogList = v
	return s
}

func (s *GetLogsListResponseBody) SetRequestId(v string) *GetLogsListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLogsListResponseBody) SetCode(v int32) *GetLogsListResponseBody {
	s.Code = &v
	return s
}

type GetLogsListResponseBodyLogList struct {
	Log []*GetLogsListResponseBodyLogListLog `json:"Log,omitempty" xml:"Log,omitempty" type:"Repeated"`
}

func (s GetLogsListResponseBodyLogList) String() string {
	return tea.Prettify(s)
}

func (s GetLogsListResponseBodyLogList) GoString() string {
	return s.String()
}

func (s *GetLogsListResponseBodyLogList) SetLog(v []*GetLogsListResponseBodyLogListLog) *GetLogsListResponseBodyLogList {
	s.Log = v
	return s
}

type GetLogsListResponseBodyLogListLog struct {
	EndDate   *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	Url       *string `json:"Url,omitempty" xml:"Url,omitempty"`
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	FileName  *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
}

func (s GetLogsListResponseBodyLogListLog) String() string {
	return tea.Prettify(s)
}

func (s GetLogsListResponseBodyLogListLog) GoString() string {
	return s.String()
}

func (s *GetLogsListResponseBodyLogListLog) SetEndDate(v string) *GetLogsListResponseBodyLogListLog {
	s.EndDate = &v
	return s
}

func (s *GetLogsListResponseBodyLogListLog) SetUrl(v string) *GetLogsListResponseBodyLogListLog {
	s.Url = &v
	return s
}

func (s *GetLogsListResponseBodyLogListLog) SetStartDate(v string) *GetLogsListResponseBodyLogListLog {
	s.StartDate = &v
	return s
}

func (s *GetLogsListResponseBodyLogListLog) SetFileName(v string) *GetLogsListResponseBodyLogListLog {
	s.FileName = &v
	return s
}

type GetLogsListResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetLogsListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetLogsListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLogsListResponse) GoString() string {
	return s.String()
}

func (s *GetLogsListResponse) SetHeaders(v map[string]*string) *GetLogsListResponse {
	s.Headers = v
	return s
}

func (s *GetLogsListResponse) SetBody(v *GetLogsListResponseBody) *GetLogsListResponse {
	s.Body = v
	return s
}

type GetShareRateDataRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty"`
	Domain        *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Region        *string `json:"Region,omitempty" xml:"Region,omitempty"`
	IspName       *string `json:"IspName,omitempty" xml:"IspName,omitempty"`
	PlatformType  *string `json:"PlatformType,omitempty" xml:"PlatformType,omitempty"`
	BusinessType  *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	StartDate     *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	EndDate       *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
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

type GetShareRateDataResponseBody struct {
	DataList  *GetShareRateDataResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Struct"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Labels    *GetShareRateDataResponseBodyLabels   `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Struct"`
	Code      *int32                                `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetShareRateDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetShareRateDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetShareRateDataResponseBody) SetDataList(v *GetShareRateDataResponseBodyDataList) *GetShareRateDataResponseBody {
	s.DataList = v
	return s
}

func (s *GetShareRateDataResponseBody) SetRequestId(v string) *GetShareRateDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetShareRateDataResponseBody) SetLabels(v *GetShareRateDataResponseBodyLabels) *GetShareRateDataResponseBody {
	s.Labels = v
	return s
}

func (s *GetShareRateDataResponseBody) SetCode(v int32) *GetShareRateDataResponseBody {
	s.Code = &v
	return s
}

type GetShareRateDataResponseBodyDataList struct {
	UsageData []*GetShareRateDataResponseBodyDataListUsageData `json:"UsageData,omitempty" xml:"UsageData,omitempty" type:"Repeated"`
}

func (s GetShareRateDataResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s GetShareRateDataResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *GetShareRateDataResponseBodyDataList) SetUsageData(v []*GetShareRateDataResponseBodyDataListUsageData) *GetShareRateDataResponseBodyDataList {
	s.UsageData = v
	return s
}

type GetShareRateDataResponseBodyDataListUsageData struct {
	Values *GetShareRateDataResponseBodyDataListUsageDataValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Struct"`
	Date   *string                                              `json:"Date,omitempty" xml:"Date,omitempty"`
}

func (s GetShareRateDataResponseBodyDataListUsageData) String() string {
	return tea.Prettify(s)
}

func (s GetShareRateDataResponseBodyDataListUsageData) GoString() string {
	return s.String()
}

func (s *GetShareRateDataResponseBodyDataListUsageData) SetValues(v *GetShareRateDataResponseBodyDataListUsageDataValues) *GetShareRateDataResponseBodyDataListUsageData {
	s.Values = v
	return s
}

func (s *GetShareRateDataResponseBodyDataListUsageData) SetDate(v string) *GetShareRateDataResponseBodyDataListUsageData {
	s.Date = &v
	return s
}

type GetShareRateDataResponseBodyDataListUsageDataValues struct {
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s GetShareRateDataResponseBodyDataListUsageDataValues) String() string {
	return tea.Prettify(s)
}

func (s GetShareRateDataResponseBodyDataListUsageDataValues) GoString() string {
	return s.String()
}

func (s *GetShareRateDataResponseBodyDataListUsageDataValues) SetValues(v []*string) *GetShareRateDataResponseBodyDataListUsageDataValues {
	s.Values = v
	return s
}

type GetShareRateDataResponseBodyLabels struct {
	Label []*string `json:"Label,omitempty" xml:"Label,omitempty" type:"Repeated"`
}

func (s GetShareRateDataResponseBodyLabels) String() string {
	return tea.Prettify(s)
}

func (s GetShareRateDataResponseBodyLabels) GoString() string {
	return s.String()
}

func (s *GetShareRateDataResponseBodyLabels) SetLabel(v []*string) *GetShareRateDataResponseBodyLabels {
	s.Label = v
	return s
}

type GetShareRateDataResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetShareRateDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetShareRateDataResponse) String() string {
	return tea.Prettify(s)
}

func (s GetShareRateDataResponse) GoString() string {
	return s.String()
}

func (s *GetShareRateDataResponse) SetHeaders(v map[string]*string) *GetShareRateDataResponse {
	s.Headers = v
	return s
}

func (s *GetShareRateDataResponse) SetBody(v *GetShareRateDataResponseBody) *GetShareRateDataResponse {
	s.Body = v
	return s
}

type GetTokenListRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty"`
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

type GetTokenListResponseBody struct {
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TokenList *GetTokenListResponseBodyTokenList `json:"TokenList,omitempty" xml:"TokenList,omitempty" type:"Struct"`
	Code      *int32                             `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetTokenListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTokenListResponseBody) GoString() string {
	return s.String()
}

func (s *GetTokenListResponseBody) SetRequestId(v string) *GetTokenListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTokenListResponseBody) SetTokenList(v *GetTokenListResponseBodyTokenList) *GetTokenListResponseBody {
	s.TokenList = v
	return s
}

func (s *GetTokenListResponseBody) SetCode(v int32) *GetTokenListResponseBody {
	s.Code = &v
	return s
}

type GetTokenListResponseBodyTokenList struct {
	Token []*GetTokenListResponseBodyTokenListToken `json:"Token,omitempty" xml:"Token,omitempty" type:"Repeated"`
}

func (s GetTokenListResponseBodyTokenList) String() string {
	return tea.Prettify(s)
}

func (s GetTokenListResponseBodyTokenList) GoString() string {
	return s.String()
}

func (s *GetTokenListResponseBodyTokenList) SetToken(v []*GetTokenListResponseBodyTokenListToken) *GetTokenListResponseBodyTokenList {
	s.Token = v
	return s
}

type GetTokenListResponseBodyTokenListToken struct {
	PlatformName *string `json:"PlatformName,omitempty" xml:"PlatformName,omitempty"`
	Token        *string `json:"Token,omitempty" xml:"Token,omitempty"`
	PlatformType *string `json:"PlatformType,omitempty" xml:"PlatformType,omitempty"`
	CreatedAt    *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	UpdatedAt    *string `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ClientId     *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
}

func (s GetTokenListResponseBodyTokenListToken) String() string {
	return tea.Prettify(s)
}

func (s GetTokenListResponseBodyTokenListToken) GoString() string {
	return s.String()
}

func (s *GetTokenListResponseBodyTokenListToken) SetPlatformName(v string) *GetTokenListResponseBodyTokenListToken {
	s.PlatformName = &v
	return s
}

func (s *GetTokenListResponseBodyTokenListToken) SetToken(v string) *GetTokenListResponseBodyTokenListToken {
	s.Token = &v
	return s
}

func (s *GetTokenListResponseBodyTokenListToken) SetPlatformType(v string) *GetTokenListResponseBodyTokenListToken {
	s.PlatformType = &v
	return s
}

func (s *GetTokenListResponseBodyTokenListToken) SetCreatedAt(v string) *GetTokenListResponseBodyTokenListToken {
	s.CreatedAt = &v
	return s
}

func (s *GetTokenListResponseBodyTokenListToken) SetUpdatedAt(v string) *GetTokenListResponseBodyTokenListToken {
	s.UpdatedAt = &v
	return s
}

func (s *GetTokenListResponseBodyTokenListToken) SetResourceId(v string) *GetTokenListResponseBodyTokenListToken {
	s.ResourceId = &v
	return s
}

func (s *GetTokenListResponseBodyTokenListToken) SetClientId(v string) *GetTokenListResponseBodyTokenListToken {
	s.ClientId = &v
	return s
}

type GetTokenListResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetTokenListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTokenListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTokenListResponse) GoString() string {
	return s.String()
}

func (s *GetTokenListResponse) SetHeaders(v map[string]*string) *GetTokenListResponse {
	s.Headers = v
	return s
}

func (s *GetTokenListResponse) SetBody(v *GetTokenListResponseBody) *GetTokenListResponse {
	s.Body = v
	return s
}

type GetTrafficByRegionRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty"`
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

type GetTrafficByRegionResponseBody struct {
	RequestId       *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TrafficDataList *GetTrafficByRegionResponseBodyTrafficDataList `json:"TrafficDataList,omitempty" xml:"TrafficDataList,omitempty" type:"Struct"`
	Code            *int32                                         `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetTrafficByRegionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTrafficByRegionResponseBody) GoString() string {
	return s.String()
}

func (s *GetTrafficByRegionResponseBody) SetRequestId(v string) *GetTrafficByRegionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTrafficByRegionResponseBody) SetTrafficDataList(v *GetTrafficByRegionResponseBodyTrafficDataList) *GetTrafficByRegionResponseBody {
	s.TrafficDataList = v
	return s
}

func (s *GetTrafficByRegionResponseBody) SetCode(v int32) *GetTrafficByRegionResponseBody {
	s.Code = &v
	return s
}

type GetTrafficByRegionResponseBodyTrafficDataList struct {
	TrafficData []*GetTrafficByRegionResponseBodyTrafficDataListTrafficData `json:"TrafficData,omitempty" xml:"TrafficData,omitempty" type:"Repeated"`
}

func (s GetTrafficByRegionResponseBodyTrafficDataList) String() string {
	return tea.Prettify(s)
}

func (s GetTrafficByRegionResponseBodyTrafficDataList) GoString() string {
	return s.String()
}

func (s *GetTrafficByRegionResponseBodyTrafficDataList) SetTrafficData(v []*GetTrafficByRegionResponseBodyTrafficDataListTrafficData) *GetTrafficByRegionResponseBodyTrafficDataList {
	s.TrafficData = v
	return s
}

type GetTrafficByRegionResponseBodyTrafficDataListTrafficData struct {
	Traffic *int64  `json:"Traffic,omitempty" xml:"Traffic,omitempty"`
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetTrafficByRegionResponseBodyTrafficDataListTrafficData) String() string {
	return tea.Prettify(s)
}

func (s GetTrafficByRegionResponseBodyTrafficDataListTrafficData) GoString() string {
	return s.String()
}

func (s *GetTrafficByRegionResponseBodyTrafficDataListTrafficData) SetTraffic(v int64) *GetTrafficByRegionResponseBodyTrafficDataListTrafficData {
	s.Traffic = &v
	return s
}

func (s *GetTrafficByRegionResponseBodyTrafficDataListTrafficData) SetName(v string) *GetTrafficByRegionResponseBodyTrafficDataListTrafficData {
	s.Name = &v
	return s
}

type GetTrafficByRegionResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetTrafficByRegionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTrafficByRegionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTrafficByRegionResponse) GoString() string {
	return s.String()
}

func (s *GetTrafficByRegionResponse) SetHeaders(v map[string]*string) *GetTrafficByRegionResponse {
	s.Headers = v
	return s
}

func (s *GetTrafficByRegionResponse) SetBody(v *GetTrafficByRegionResponseBody) *GetTrafficByRegionResponse {
	s.Body = v
	return s
}

type GetTrafficDataRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty"`
	Domain        *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Region        *string `json:"Region,omitempty" xml:"Region,omitempty"`
	IspName       *string `json:"IspName,omitempty" xml:"IspName,omitempty"`
	PlatformType  *string `json:"PlatformType,omitempty" xml:"PlatformType,omitempty"`
	BusinessType  *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	StartDate     *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	EndDate       *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
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

type GetTrafficDataResponseBody struct {
	DataList  *GetTrafficDataResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Struct"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Labels    *GetTrafficDataResponseBodyLabels   `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Struct"`
	Code      *int32                              `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetTrafficDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTrafficDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetTrafficDataResponseBody) SetDataList(v *GetTrafficDataResponseBodyDataList) *GetTrafficDataResponseBody {
	s.DataList = v
	return s
}

func (s *GetTrafficDataResponseBody) SetRequestId(v string) *GetTrafficDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTrafficDataResponseBody) SetLabels(v *GetTrafficDataResponseBodyLabels) *GetTrafficDataResponseBody {
	s.Labels = v
	return s
}

func (s *GetTrafficDataResponseBody) SetCode(v int32) *GetTrafficDataResponseBody {
	s.Code = &v
	return s
}

type GetTrafficDataResponseBodyDataList struct {
	UsageData []*GetTrafficDataResponseBodyDataListUsageData `json:"UsageData,omitempty" xml:"UsageData,omitempty" type:"Repeated"`
}

func (s GetTrafficDataResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s GetTrafficDataResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *GetTrafficDataResponseBodyDataList) SetUsageData(v []*GetTrafficDataResponseBodyDataListUsageData) *GetTrafficDataResponseBodyDataList {
	s.UsageData = v
	return s
}

type GetTrafficDataResponseBodyDataListUsageData struct {
	Values *GetTrafficDataResponseBodyDataListUsageDataValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Struct"`
	Date   *string                                            `json:"Date,omitempty" xml:"Date,omitempty"`
}

func (s GetTrafficDataResponseBodyDataListUsageData) String() string {
	return tea.Prettify(s)
}

func (s GetTrafficDataResponseBodyDataListUsageData) GoString() string {
	return s.String()
}

func (s *GetTrafficDataResponseBodyDataListUsageData) SetValues(v *GetTrafficDataResponseBodyDataListUsageDataValues) *GetTrafficDataResponseBodyDataListUsageData {
	s.Values = v
	return s
}

func (s *GetTrafficDataResponseBodyDataListUsageData) SetDate(v string) *GetTrafficDataResponseBodyDataListUsageData {
	s.Date = &v
	return s
}

type GetTrafficDataResponseBodyDataListUsageDataValues struct {
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s GetTrafficDataResponseBodyDataListUsageDataValues) String() string {
	return tea.Prettify(s)
}

func (s GetTrafficDataResponseBodyDataListUsageDataValues) GoString() string {
	return s.String()
}

func (s *GetTrafficDataResponseBodyDataListUsageDataValues) SetValues(v []*string) *GetTrafficDataResponseBodyDataListUsageDataValues {
	s.Values = v
	return s
}

type GetTrafficDataResponseBodyLabels struct {
	Label []*string `json:"Label,omitempty" xml:"Label,omitempty" type:"Repeated"`
}

func (s GetTrafficDataResponseBodyLabels) String() string {
	return tea.Prettify(s)
}

func (s GetTrafficDataResponseBodyLabels) GoString() string {
	return s.String()
}

func (s *GetTrafficDataResponseBodyLabels) SetLabel(v []*string) *GetTrafficDataResponseBodyLabels {
	s.Label = v
	return s
}

type GetTrafficDataResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetTrafficDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTrafficDataResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTrafficDataResponse) GoString() string {
	return s.String()
}

func (s *GetTrafficDataResponse) SetHeaders(v map[string]*string) *GetTrafficDataResponse {
	s.Headers = v
	return s
}

func (s *GetTrafficDataResponse) SetBody(v *GetTrafficDataResponseBody) *GetTrafficDataResponse {
	s.Body = v
	return s
}

type StartDomainRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty"`
	Domain        *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
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

type StartDomainResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	Code       *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s StartDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartDomainResponseBody) GoString() string {
	return s.String()
}

func (s *StartDomainResponseBody) SetRequestId(v string) *StartDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartDomainResponseBody) SetResourceId(v string) *StartDomainResponseBody {
	s.ResourceId = &v
	return s
}

func (s *StartDomainResponseBody) SetCode(v int32) *StartDomainResponseBody {
	s.Code = &v
	return s
}

type StartDomainResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s StartDomainResponse) GoString() string {
	return s.String()
}

func (s *StartDomainResponse) SetHeaders(v map[string]*string) *StartDomainResponse {
	s.Headers = v
	return s
}

func (s *StartDomainResponse) SetBody(v *StartDomainResponseBody) *StartDomainResponse {
	s.Body = v
	return s
}

type StopDomainRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty"`
	Domain        *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
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

type StopDomainResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	Code       *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s StopDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopDomainResponseBody) GoString() string {
	return s.String()
}

func (s *StopDomainResponseBody) SetRequestId(v string) *StopDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopDomainResponseBody) SetResourceId(v string) *StopDomainResponseBody {
	s.ResourceId = &v
	return s
}

func (s *StopDomainResponseBody) SetCode(v int32) *StopDomainResponseBody {
	s.Code = &v
	return s
}

type StopDomainResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s StopDomainResponse) GoString() string {
	return s.String()
}

func (s *StopDomainResponse) SetHeaders(v map[string]*string) *StopDomainResponse {
	s.Headers = v
	return s
}

func (s *StopDomainResponse) SetBody(v *StopDomainResponseBody) *StopDomainResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
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

func (client *Client) AddConsumerWithOptions(request *AddConsumerRequest, runtime *util.RuntimeOptions) (_result *AddConsumerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &AddConsumerResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddConsumer"), tea.String("2017-04-11"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) AddDomainWithOptions(request *AddDomainRequest, runtime *util.RuntimeOptions) (_result *AddDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &AddDomainResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddDomain"), tea.String("2017-04-11"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) AddPcdnControlRuleWithOptions(request *AddPcdnControlRuleRequest, runtime *util.RuntimeOptions) (_result *AddPcdnControlRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &AddPcdnControlRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddPcdnControlRule"), tea.String("2017-04-11"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DeleteDomainWithOptions(request *DeleteDomainRequest, runtime *util.RuntimeOptions) (_result *DeleteDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DeleteDomainResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteDomain"), tea.String("2017-04-11"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DeletePcdnControlRuleWithOptions(request *DeletePcdnControlRuleRequest, runtime *util.RuntimeOptions) (_result *DeletePcdnControlRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DeletePcdnControlRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeletePcdnControlRule"), tea.String("2017-04-11"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DisablePcdnControlRuleWithOptions(request *DisablePcdnControlRuleRequest, runtime *util.RuntimeOptions) (_result *DisablePcdnControlRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DisablePcdnControlRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DisablePcdnControlRule"), tea.String("2017-04-11"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) EditPcdnControlRuleWithOptions(request *EditPcdnControlRuleRequest, runtime *util.RuntimeOptions) (_result *EditPcdnControlRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &EditPcdnControlRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("EditPcdnControlRule"), tea.String("2017-04-11"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) EnablePcdnControlRuleWithOptions(request *EnablePcdnControlRuleRequest, runtime *util.RuntimeOptions) (_result *EnablePcdnControlRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &EnablePcdnControlRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("EnablePcdnControlRule"), tea.String("2017-04-11"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) GetAccessDataWithOptions(request *GetAccessDataRequest, runtime *util.RuntimeOptions) (_result *GetAccessDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetAccessDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetAccessData"), tea.String("2017-04-11"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) GetAllAppVersionsWithOptions(request *GetAllAppVersionsRequest, runtime *util.RuntimeOptions) (_result *GetAllAppVersionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetAllAppVersionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetAllAppVersions"), tea.String("2017-04-11"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) GetAllIspWithOptions(request *GetAllIspRequest, runtime *util.RuntimeOptions) (_result *GetAllIspResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetAllIspResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetAllIsp"), tea.String("2017-04-11"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) GetAllMarketsWithOptions(request *GetAllMarketsRequest, runtime *util.RuntimeOptions) (_result *GetAllMarketsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetAllMarketsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetAllMarkets"), tea.String("2017-04-11"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) GetAllPlatformTypesWithOptions(request *GetAllPlatformTypesRequest, runtime *util.RuntimeOptions) (_result *GetAllPlatformTypesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetAllPlatformTypesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetAllPlatformTypes"), tea.String("2017-04-11"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) GetAllRegionsWithOptions(request *GetAllRegionsRequest, runtime *util.RuntimeOptions) (_result *GetAllRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetAllRegionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetAllRegions"), tea.String("2017-04-11"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) GetBalanceBandwidthDataWithOptions(request *GetBalanceBandwidthDataRequest, runtime *util.RuntimeOptions) (_result *GetBalanceBandwidthDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetBalanceBandwidthDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetBalanceBandwidthData"), tea.String("2017-04-11"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) GetBalanceTrafficDataWithOptions(request *GetBalanceTrafficDataRequest, runtime *util.RuntimeOptions) (_result *GetBalanceTrafficDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetBalanceTrafficDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetBalanceTrafficData"), tea.String("2017-04-11"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) GetBandwidthDataWithOptions(request *GetBandwidthDataRequest, runtime *util.RuntimeOptions) (_result *GetBandwidthDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetBandwidthDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetBandwidthData"), tea.String("2017-04-11"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) GetClientsRatioWithOptions(request *GetClientsRatioRequest, runtime *util.RuntimeOptions) (_result *GetClientsRatioResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetClientsRatioResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetClientsRatio"), tea.String("2017-04-11"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) GetConsumerStatusWithOptions(request *GetConsumerStatusRequest, runtime *util.RuntimeOptions) (_result *GetConsumerStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetConsumerStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetConsumerStatus"), tea.String("2017-04-11"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) GetControlRulesWithOptions(request *GetControlRulesRequest, runtime *util.RuntimeOptions) (_result *GetControlRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetControlRulesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetControlRules"), tea.String("2017-04-11"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) GetCoverRateDataWithOptions(request *GetCoverRateDataRequest, runtime *util.RuntimeOptions) (_result *GetCoverRateDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetCoverRateDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetCoverRateData"), tea.String("2017-04-11"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) GetCurrentModeWithOptions(request *GetCurrentModeRequest, runtime *util.RuntimeOptions) (_result *GetCurrentModeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetCurrentModeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetCurrentMode"), tea.String("2017-04-11"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) GetDomainCountWithOptions(request *GetDomainCountRequest, runtime *util.RuntimeOptions) (_result *GetDomainCountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetDomainCountResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetDomainCount"), tea.String("2017-04-11"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) GetDomainsWithOptions(request *GetDomainsRequest, runtime *util.RuntimeOptions) (_result *GetDomainsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetDomainsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetDomains"), tea.String("2017-04-11"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) GetExpenseSummaryWithOptions(request *GetExpenseSummaryRequest, runtime *util.RuntimeOptions) (_result *GetExpenseSummaryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetExpenseSummaryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetExpenseSummary"), tea.String("2017-04-11"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) GetFeeHistoryWithOptions(request *GetFeeHistoryRequest, runtime *util.RuntimeOptions) (_result *GetFeeHistoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetFeeHistoryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetFeeHistory"), tea.String("2017-04-11"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) GetFirstFrameDelayDataWithOptions(request *GetFirstFrameDelayDataRequest, runtime *util.RuntimeOptions) (_result *GetFirstFrameDelayDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetFirstFrameDelayDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetFirstFrameDelayData"), tea.String("2017-04-11"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) GetFluencyDataWithOptions(request *GetFluencyDataRequest, runtime *util.RuntimeOptions) (_result *GetFluencyDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetFluencyDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetFluencyData"), tea.String("2017-04-11"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) GetLogsListWithOptions(request *GetLogsListRequest, runtime *util.RuntimeOptions) (_result *GetLogsListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetLogsListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetLogsList"), tea.String("2017-04-11"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) GetShareRateDataWithOptions(request *GetShareRateDataRequest, runtime *util.RuntimeOptions) (_result *GetShareRateDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetShareRateDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetShareRateData"), tea.String("2017-04-11"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) GetTokenListWithOptions(request *GetTokenListRequest, runtime *util.RuntimeOptions) (_result *GetTokenListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetTokenListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetTokenList"), tea.String("2017-04-11"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) GetTrafficByRegionWithOptions(request *GetTrafficByRegionRequest, runtime *util.RuntimeOptions) (_result *GetTrafficByRegionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetTrafficByRegionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetTrafficByRegion"), tea.String("2017-04-11"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) GetTrafficDataWithOptions(request *GetTrafficDataRequest, runtime *util.RuntimeOptions) (_result *GetTrafficDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetTrafficDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetTrafficData"), tea.String("2017-04-11"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) StartDomainWithOptions(request *StartDomainRequest, runtime *util.RuntimeOptions) (_result *StartDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &StartDomainResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StartDomain"), tea.String("2017-04-11"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) StopDomainWithOptions(request *StopDomainRequest, runtime *util.RuntimeOptions) (_result *StopDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &StopDomainResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StopDomain"), tea.String("2017-04-11"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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
