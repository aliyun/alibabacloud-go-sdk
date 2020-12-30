// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddAccountRequest struct {
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	AccountName       *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	Password          *string `json:"Password,omitempty" xml:"Password,omitempty"`
	IsShort           *bool   `json:"IsShort,omitempty" xml:"IsShort,omitempty"`
	EnableKMS         *bool   `json:"EnableKMS,omitempty" xml:"EnableKMS,omitempty"`
	Remark            *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	RamUid            *string `json:"RamUid,omitempty" xml:"RamUid,omitempty"`
	UseRandomPassword *bool   `json:"UseRandomPassword,omitempty" xml:"UseRandomPassword,omitempty"`
}

func (s AddAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s AddAccountRequest) GoString() string {
	return s.String()
}

func (s *AddAccountRequest) SetRegionId(v string) *AddAccountRequest {
	s.RegionId = &v
	return s
}

func (s *AddAccountRequest) SetAccountName(v string) *AddAccountRequest {
	s.AccountName = &v
	return s
}

func (s *AddAccountRequest) SetPassword(v string) *AddAccountRequest {
	s.Password = &v
	return s
}

func (s *AddAccountRequest) SetIsShort(v bool) *AddAccountRequest {
	s.IsShort = &v
	return s
}

func (s *AddAccountRequest) SetEnableKMS(v bool) *AddAccountRequest {
	s.EnableKMS = &v
	return s
}

func (s *AddAccountRequest) SetRemark(v string) *AddAccountRequest {
	s.Remark = &v
	return s
}

func (s *AddAccountRequest) SetRamUid(v string) *AddAccountRequest {
	s.RamUid = &v
	return s
}

func (s *AddAccountRequest) SetUseRandomPassword(v bool) *AddAccountRequest {
	s.UseRandomPassword = &v
	return s
}

type AddAccountResponseBody struct {
	Account   *AddAccountResponseBodyAccount `json:"Account,omitempty" xml:"Account,omitempty" type:"Struct"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RegionId  *string                        `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AddAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddAccountResponseBody) GoString() string {
	return s.String()
}

func (s *AddAccountResponseBody) SetAccount(v *AddAccountResponseBodyAccount) *AddAccountResponseBody {
	s.Account = v
	return s
}

func (s *AddAccountResponseBody) SetRequestId(v string) *AddAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddAccountResponseBody) SetRegionId(v string) *AddAccountResponseBody {
	s.RegionId = &v
	return s
}

type AddAccountResponseBodyAccount struct {
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s AddAccountResponseBodyAccount) String() string {
	return tea.Prettify(s)
}

func (s AddAccountResponseBodyAccount) GoString() string {
	return s.String()
}

func (s *AddAccountResponseBodyAccount) SetPassword(v string) *AddAccountResponseBodyAccount {
	s.Password = &v
	return s
}

func (s *AddAccountResponseBodyAccount) SetUserName(v string) *AddAccountResponseBodyAccount {
	s.UserName = &v
	return s
}

type AddAccountResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s AddAccountResponse) GoString() string {
	return s.String()
}

func (s *AddAccountResponse) SetHeaders(v map[string]*string) *AddAccountResponse {
	s.Headers = v
	return s
}

func (s *AddAccountResponse) SetBody(v *AddAccountResponseBody) *AddAccountResponse {
	s.Body = v
	return s
}

type AddEndPointRequest struct {
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	VpcID       *string `json:"VpcID,omitempty" xml:"VpcID,omitempty"`
	Vswitch     *string `json:"Vswitch,omitempty" xml:"Vswitch,omitempty"`
	Zone        *string `json:"Zone,omitempty" xml:"Zone,omitempty"`
	Product     *string `json:"Product,omitempty" xml:"Product,omitempty"`
}

func (s AddEndPointRequest) String() string {
	return tea.Prettify(s)
}

func (s AddEndPointRequest) GoString() string {
	return s.String()
}

func (s *AddEndPointRequest) SetRegionId(v string) *AddEndPointRequest {
	s.RegionId = &v
	return s
}

func (s *AddEndPointRequest) SetNetworkType(v string) *AddEndPointRequest {
	s.NetworkType = &v
	return s
}

func (s *AddEndPointRequest) SetVpcID(v string) *AddEndPointRequest {
	s.VpcID = &v
	return s
}

func (s *AddEndPointRequest) SetVswitch(v string) *AddEndPointRequest {
	s.Vswitch = &v
	return s
}

func (s *AddEndPointRequest) SetZone(v string) *AddEndPointRequest {
	s.Zone = &v
	return s
}

func (s *AddEndPointRequest) SetProduct(v string) *AddEndPointRequest {
	s.Product = &v
	return s
}

type AddEndPointResponseBody struct {
	RequestId    *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	EndPointInfo *AddEndPointResponseBodyEndPointInfo `json:"EndPointInfo,omitempty" xml:"EndPointInfo,omitempty" type:"Struct"`
	RegionId     *string                              `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AddEndPointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddEndPointResponseBody) GoString() string {
	return s.String()
}

func (s *AddEndPointResponseBody) SetRequestId(v string) *AddEndPointResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddEndPointResponseBody) SetEndPointInfo(v *AddEndPointResponseBodyEndPointInfo) *AddEndPointResponseBody {
	s.EndPointInfo = v
	return s
}

func (s *AddEndPointResponseBody) SetRegionId(v string) *AddEndPointResponseBody {
	s.RegionId = &v
	return s
}

type AddEndPointResponseBodyEndPointInfo struct {
	Product     *string `json:"product,omitempty" xml:"product,omitempty"`
	Zone        *string `json:"zone,omitempty" xml:"zone,omitempty"`
	DomainURL   *string `json:"domainURL,omitempty" xml:"domainURL,omitempty"`
	VSwitch     *string `json:"vSwitch,omitempty" xml:"vSwitch,omitempty"`
	Host        *string `json:"host,omitempty" xml:"host,omitempty"`
	Status      *string `json:"status,omitempty" xml:"status,omitempty"`
	EndPointID  *string `json:"endPointID,omitempty" xml:"endPointID,omitempty"`
	AllowIP     *string `json:"allowIP,omitempty" xml:"allowIP,omitempty"`
	VpcID       *string `json:"vpcID,omitempty" xml:"vpcID,omitempty"`
	NetworkType *string `json:"networkType,omitempty" xml:"networkType,omitempty"`
	Port        *string `json:"port,omitempty" xml:"port,omitempty"`
}

func (s AddEndPointResponseBodyEndPointInfo) String() string {
	return tea.Prettify(s)
}

func (s AddEndPointResponseBodyEndPointInfo) GoString() string {
	return s.String()
}

func (s *AddEndPointResponseBodyEndPointInfo) SetProduct(v string) *AddEndPointResponseBodyEndPointInfo {
	s.Product = &v
	return s
}

func (s *AddEndPointResponseBodyEndPointInfo) SetZone(v string) *AddEndPointResponseBodyEndPointInfo {
	s.Zone = &v
	return s
}

func (s *AddEndPointResponseBodyEndPointInfo) SetDomainURL(v string) *AddEndPointResponseBodyEndPointInfo {
	s.DomainURL = &v
	return s
}

func (s *AddEndPointResponseBodyEndPointInfo) SetVSwitch(v string) *AddEndPointResponseBodyEndPointInfo {
	s.VSwitch = &v
	return s
}

func (s *AddEndPointResponseBodyEndPointInfo) SetHost(v string) *AddEndPointResponseBodyEndPointInfo {
	s.Host = &v
	return s
}

func (s *AddEndPointResponseBodyEndPointInfo) SetStatus(v string) *AddEndPointResponseBodyEndPointInfo {
	s.Status = &v
	return s
}

func (s *AddEndPointResponseBodyEndPointInfo) SetEndPointID(v string) *AddEndPointResponseBodyEndPointInfo {
	s.EndPointID = &v
	return s
}

func (s *AddEndPointResponseBodyEndPointInfo) SetAllowIP(v string) *AddEndPointResponseBodyEndPointInfo {
	s.AllowIP = &v
	return s
}

func (s *AddEndPointResponseBodyEndPointInfo) SetVpcID(v string) *AddEndPointResponseBodyEndPointInfo {
	s.VpcID = &v
	return s
}

func (s *AddEndPointResponseBodyEndPointInfo) SetNetworkType(v string) *AddEndPointResponseBodyEndPointInfo {
	s.NetworkType = &v
	return s
}

func (s *AddEndPointResponseBodyEndPointInfo) SetPort(v string) *AddEndPointResponseBodyEndPointInfo {
	s.Port = &v
	return s
}

type AddEndPointResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddEndPointResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddEndPointResponse) String() string {
	return tea.Prettify(s)
}

func (s AddEndPointResponse) GoString() string {
	return s.String()
}

func (s *AddEndPointResponse) SetHeaders(v map[string]*string) *AddEndPointResponse {
	s.Headers = v
	return s
}

func (s *AddEndPointResponse) SetBody(v *AddEndPointResponseBody) *AddEndPointResponse {
	s.Body = v
	return s
}

type CreateInstanceRequest struct {
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ChargeType   *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	Component    *string `json:"Component,omitempty" xml:"Component,omitempty"`
}

func (s CreateInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequest) SetRegionId(v string) *CreateInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateInstanceRequest) SetChargeType(v string) *CreateInstanceRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateInstanceRequest) SetInstanceType(v string) *CreateInstanceRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateInstanceRequest) SetComponent(v string) *CreateInstanceRequest {
	s.Component = &v
	return s
}

type CreateInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorInfo *string `json:"ErrorInfo,omitempty" xml:"ErrorInfo,omitempty"`
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Result    *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s CreateInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponseBody) SetRequestId(v string) *CreateInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInstanceResponseBody) SetErrorInfo(v string) *CreateInstanceResponseBody {
	s.ErrorInfo = &v
	return s
}

func (s *CreateInstanceResponseBody) SetErrorCode(v string) *CreateInstanceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateInstanceResponseBody) SetSuccess(v bool) *CreateInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *CreateInstanceResponseBody) SetResult(v string) *CreateInstanceResponseBody {
	s.Result = &v
	return s
}

type CreateInstanceResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponse) SetHeaders(v map[string]*string) *CreateInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateInstanceResponse) SetBody(v *CreateInstanceResponseBody) *CreateInstanceResponse {
	s.Body = v
	return s
}

type DeleteAccountRequest struct {
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	AccountName   *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	IsShort       *bool   `json:"IsShort,omitempty" xml:"IsShort,omitempty"`
	IsServiceUser *bool   `json:"IsServiceUser,omitempty" xml:"IsServiceUser,omitempty"`
}

func (s DeleteAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccountRequest) GoString() string {
	return s.String()
}

func (s *DeleteAccountRequest) SetRegionId(v string) *DeleteAccountRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteAccountRequest) SetAccountName(v string) *DeleteAccountRequest {
	s.AccountName = &v
	return s
}

func (s *DeleteAccountRequest) SetIsShort(v bool) *DeleteAccountRequest {
	s.IsShort = &v
	return s
}

func (s *DeleteAccountRequest) SetIsServiceUser(v bool) *DeleteAccountRequest {
	s.IsServiceUser = &v
	return s
}

type DeleteAccountResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccountResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAccountResponseBody) SetRequestId(v string) *DeleteAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAccountResponseBody) SetRegionId(v string) *DeleteAccountResponseBody {
	s.RegionId = &v
	return s
}

type DeleteAccountResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccountResponse) GoString() string {
	return s.String()
}

func (s *DeleteAccountResponse) SetHeaders(v map[string]*string) *DeleteAccountResponse {
	s.Headers = v
	return s
}

func (s *DeleteAccountResponse) SetBody(v *DeleteAccountResponseBody) *DeleteAccountResponse {
	s.Body = v
	return s
}

type GetAllowIPRequest struct {
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	Product     *string `json:"Product,omitempty" xml:"Product,omitempty"`
}

func (s GetAllowIPRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAllowIPRequest) GoString() string {
	return s.String()
}

func (s *GetAllowIPRequest) SetRegionId(v string) *GetAllowIPRequest {
	s.RegionId = &v
	return s
}

func (s *GetAllowIPRequest) SetNetworkType(v string) *GetAllowIPRequest {
	s.NetworkType = &v
	return s
}

func (s *GetAllowIPRequest) SetProduct(v string) *GetAllowIPRequest {
	s.Product = &v
	return s
}

type GetAllowIPResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	AllowIP   *string `json:"AllowIP,omitempty" xml:"AllowIP,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetAllowIPResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAllowIPResponseBody) GoString() string {
	return s.String()
}

func (s *GetAllowIPResponseBody) SetRequestId(v string) *GetAllowIPResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAllowIPResponseBody) SetAllowIP(v string) *GetAllowIPResponseBody {
	s.AllowIP = &v
	return s
}

func (s *GetAllowIPResponseBody) SetRegionId(v string) *GetAllowIPResponseBody {
	s.RegionId = &v
	return s
}

type GetAllowIPResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAllowIPResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAllowIPResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAllowIPResponse) GoString() string {
	return s.String()
}

func (s *GetAllowIPResponse) SetHeaders(v map[string]*string) *GetAllowIPResponse {
	s.Headers = v
	return s
}

func (s *GetAllowIPResponse) SetBody(v *GetAllowIPResponseBody) *GetAllowIPResponse {
	s.Body = v
	return s
}

type GetEndPointRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	EndPointID *string `json:"EndPointID,omitempty" xml:"EndPointID,omitempty"`
}

func (s GetEndPointRequest) String() string {
	return tea.Prettify(s)
}

func (s GetEndPointRequest) GoString() string {
	return s.String()
}

func (s *GetEndPointRequest) SetRegionId(v string) *GetEndPointRequest {
	s.RegionId = &v
	return s
}

func (s *GetEndPointRequest) SetEndPointID(v string) *GetEndPointRequest {
	s.EndPointID = &v
	return s
}

type GetEndPointResponseBody struct {
	RequestId    *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	EndPointInfo *GetEndPointResponseBodyEndPointInfo `json:"EndPointInfo,omitempty" xml:"EndPointInfo,omitempty" type:"Struct"`
	RegionId     *string                              `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetEndPointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetEndPointResponseBody) GoString() string {
	return s.String()
}

func (s *GetEndPointResponseBody) SetRequestId(v string) *GetEndPointResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEndPointResponseBody) SetEndPointInfo(v *GetEndPointResponseBodyEndPointInfo) *GetEndPointResponseBody {
	s.EndPointInfo = v
	return s
}

func (s *GetEndPointResponseBody) SetRegionId(v string) *GetEndPointResponseBody {
	s.RegionId = &v
	return s
}

type GetEndPointResponseBodyEndPointInfo struct {
	Product     *string `json:"product,omitempty" xml:"product,omitempty"`
	Zone        *string `json:"zone,omitempty" xml:"zone,omitempty"`
	DomainURL   *string `json:"domainURL,omitempty" xml:"domainURL,omitempty"`
	VSwitch     *string `json:"vSwitch,omitempty" xml:"vSwitch,omitempty"`
	Host        *string `json:"host,omitempty" xml:"host,omitempty"`
	Status      *string `json:"status,omitempty" xml:"status,omitempty"`
	AllowIP     *string `json:"allowIP,omitempty" xml:"allowIP,omitempty"`
	VpcID       *string `json:"vpcID,omitempty" xml:"vpcID,omitempty"`
	EndPointID  *string `json:"endPointID,omitempty" xml:"endPointID,omitempty"`
	NetworkType *string `json:"networkType,omitempty" xml:"networkType,omitempty"`
	Port        *string `json:"port,omitempty" xml:"port,omitempty"`
}

func (s GetEndPointResponseBodyEndPointInfo) String() string {
	return tea.Prettify(s)
}

func (s GetEndPointResponseBodyEndPointInfo) GoString() string {
	return s.String()
}

func (s *GetEndPointResponseBodyEndPointInfo) SetProduct(v string) *GetEndPointResponseBodyEndPointInfo {
	s.Product = &v
	return s
}

func (s *GetEndPointResponseBodyEndPointInfo) SetZone(v string) *GetEndPointResponseBodyEndPointInfo {
	s.Zone = &v
	return s
}

func (s *GetEndPointResponseBodyEndPointInfo) SetDomainURL(v string) *GetEndPointResponseBodyEndPointInfo {
	s.DomainURL = &v
	return s
}

func (s *GetEndPointResponseBodyEndPointInfo) SetVSwitch(v string) *GetEndPointResponseBodyEndPointInfo {
	s.VSwitch = &v
	return s
}

func (s *GetEndPointResponseBodyEndPointInfo) SetHost(v string) *GetEndPointResponseBodyEndPointInfo {
	s.Host = &v
	return s
}

func (s *GetEndPointResponseBodyEndPointInfo) SetStatus(v string) *GetEndPointResponseBodyEndPointInfo {
	s.Status = &v
	return s
}

func (s *GetEndPointResponseBodyEndPointInfo) SetAllowIP(v string) *GetEndPointResponseBodyEndPointInfo {
	s.AllowIP = &v
	return s
}

func (s *GetEndPointResponseBodyEndPointInfo) SetVpcID(v string) *GetEndPointResponseBodyEndPointInfo {
	s.VpcID = &v
	return s
}

func (s *GetEndPointResponseBodyEndPointInfo) SetEndPointID(v string) *GetEndPointResponseBodyEndPointInfo {
	s.EndPointID = &v
	return s
}

func (s *GetEndPointResponseBodyEndPointInfo) SetNetworkType(v string) *GetEndPointResponseBodyEndPointInfo {
	s.NetworkType = &v
	return s
}

func (s *GetEndPointResponseBodyEndPointInfo) SetPort(v string) *GetEndPointResponseBodyEndPointInfo {
	s.Port = &v
	return s
}

type GetEndPointResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetEndPointResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetEndPointResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEndPointResponse) GoString() string {
	return s.String()
}

func (s *GetEndPointResponse) SetHeaders(v map[string]*string) *GetEndPointResponse {
	s.Headers = v
	return s
}

func (s *GetEndPointResponse) SetBody(v *GetEndPointResponseBody) *GetEndPointResponse {
	s.Body = v
	return s
}

type GetEndPointByDomainRequest struct {
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DomainURL *string `json:"DomainURL,omitempty" xml:"DomainURL,omitempty"`
	RegionID  *string `json:"RegionID,omitempty" xml:"RegionID,omitempty"`
}

func (s GetEndPointByDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s GetEndPointByDomainRequest) GoString() string {
	return s.String()
}

func (s *GetEndPointByDomainRequest) SetRegionId(v string) *GetEndPointByDomainRequest {
	s.RegionId = &v
	return s
}

func (s *GetEndPointByDomainRequest) SetDomainURL(v string) *GetEndPointByDomainRequest {
	s.DomainURL = &v
	return s
}

func (s *GetEndPointByDomainRequest) SetRegionID(v string) *GetEndPointByDomainRequest {
	s.RegionID = &v
	return s
}

type GetEndPointByDomainResponseBody struct {
	RequestId    *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	EndPointInfo *GetEndPointByDomainResponseBodyEndPointInfo `json:"EndPointInfo,omitempty" xml:"EndPointInfo,omitempty" type:"Struct"`
	RegionId     *string                                      `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetEndPointByDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetEndPointByDomainResponseBody) GoString() string {
	return s.String()
}

func (s *GetEndPointByDomainResponseBody) SetRequestId(v string) *GetEndPointByDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEndPointByDomainResponseBody) SetEndPointInfo(v *GetEndPointByDomainResponseBodyEndPointInfo) *GetEndPointByDomainResponseBody {
	s.EndPointInfo = v
	return s
}

func (s *GetEndPointByDomainResponseBody) SetRegionId(v string) *GetEndPointByDomainResponseBody {
	s.RegionId = &v
	return s
}

type GetEndPointByDomainResponseBodyEndPointInfo struct {
	Product         *string `json:"product,omitempty" xml:"product,omitempty"`
	CloudInstanceID *string `json:"cloudInstanceID,omitempty" xml:"cloudInstanceID,omitempty"`
	Zone            *string `json:"zone,omitempty" xml:"zone,omitempty"`
	DomainURL       *string `json:"domainURL,omitempty" xml:"domainURL,omitempty"`
	VSwitch         *string `json:"vSwitch,omitempty" xml:"vSwitch,omitempty"`
	Host            *string `json:"host,omitempty" xml:"host,omitempty"`
	AllowIP         *string `json:"allowIP,omitempty" xml:"allowIP,omitempty"`
	VpcID           *string `json:"vpcID,omitempty" xml:"vpcID,omitempty"`
	EndPointID      *string `json:"endPointID,omitempty" xml:"endPointID,omitempty"`
	NetworkType     *string `json:"networkType,omitempty" xml:"networkType,omitempty"`
	Port            *string `json:"port,omitempty" xml:"port,omitempty"`
}

func (s GetEndPointByDomainResponseBodyEndPointInfo) String() string {
	return tea.Prettify(s)
}

func (s GetEndPointByDomainResponseBodyEndPointInfo) GoString() string {
	return s.String()
}

func (s *GetEndPointByDomainResponseBodyEndPointInfo) SetProduct(v string) *GetEndPointByDomainResponseBodyEndPointInfo {
	s.Product = &v
	return s
}

func (s *GetEndPointByDomainResponseBodyEndPointInfo) SetCloudInstanceID(v string) *GetEndPointByDomainResponseBodyEndPointInfo {
	s.CloudInstanceID = &v
	return s
}

func (s *GetEndPointByDomainResponseBodyEndPointInfo) SetZone(v string) *GetEndPointByDomainResponseBodyEndPointInfo {
	s.Zone = &v
	return s
}

func (s *GetEndPointByDomainResponseBodyEndPointInfo) SetDomainURL(v string) *GetEndPointByDomainResponseBodyEndPointInfo {
	s.DomainURL = &v
	return s
}

func (s *GetEndPointByDomainResponseBodyEndPointInfo) SetVSwitch(v string) *GetEndPointByDomainResponseBodyEndPointInfo {
	s.VSwitch = &v
	return s
}

func (s *GetEndPointByDomainResponseBodyEndPointInfo) SetHost(v string) *GetEndPointByDomainResponseBodyEndPointInfo {
	s.Host = &v
	return s
}

func (s *GetEndPointByDomainResponseBodyEndPointInfo) SetAllowIP(v string) *GetEndPointByDomainResponseBodyEndPointInfo {
	s.AllowIP = &v
	return s
}

func (s *GetEndPointByDomainResponseBodyEndPointInfo) SetVpcID(v string) *GetEndPointByDomainResponseBodyEndPointInfo {
	s.VpcID = &v
	return s
}

func (s *GetEndPointByDomainResponseBodyEndPointInfo) SetEndPointID(v string) *GetEndPointByDomainResponseBodyEndPointInfo {
	s.EndPointID = &v
	return s
}

func (s *GetEndPointByDomainResponseBodyEndPointInfo) SetNetworkType(v string) *GetEndPointByDomainResponseBodyEndPointInfo {
	s.NetworkType = &v
	return s
}

func (s *GetEndPointByDomainResponseBodyEndPointInfo) SetPort(v string) *GetEndPointByDomainResponseBodyEndPointInfo {
	s.Port = &v
	return s
}

type GetEndPointByDomainResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetEndPointByDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetEndPointByDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEndPointByDomainResponse) GoString() string {
	return s.String()
}

func (s *GetEndPointByDomainResponse) SetHeaders(v map[string]*string) *GetEndPointByDomainResponse {
	s.Headers = v
	return s
}

func (s *GetEndPointByDomainResponse) SetBody(v *GetEndPointByDomainResponseBody) *GetEndPointByDomainResponse {
	s.Body = v
	return s
}

type GetJobDetailRequest struct {
	JobId  *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	VcName *string `json:"VcName,omitempty" xml:"VcName,omitempty"`
}

func (s GetJobDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetJobDetailRequest) GoString() string {
	return s.String()
}

func (s *GetJobDetailRequest) SetJobId(v string) *GetJobDetailRequest {
	s.JobId = &v
	return s
}

func (s *GetJobDetailRequest) SetVcName(v string) *GetJobDetailRequest {
	s.VcName = &v
	return s
}

type GetJobDetailResponseBody struct {
	JobDetail *GetJobDetailResponseBodyJobDetail `json:"JobDetail,omitempty" xml:"JobDetail,omitempty" type:"Struct"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetJobDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetJobDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobDetailResponseBody) SetJobDetail(v *GetJobDetailResponseBodyJobDetail) *GetJobDetailResponseBody {
	s.JobDetail = v
	return s
}

func (s *GetJobDetailResponseBody) SetRequestId(v string) *GetJobDetailResponseBody {
	s.RequestId = &v
	return s
}

type GetJobDetailResponseBodyJobDetail struct {
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UpdateTime           *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	ExecutorResourceSpec *string `json:"ExecutorResourceSpec,omitempty" xml:"ExecutorResourceSpec,omitempty"`
	CreateTime           *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateTimeValue      *string `json:"CreateTimeValue,omitempty" xml:"CreateTimeValue,omitempty"`
	DriverResourceSpec   *string `json:"DriverResourceSpec,omitempty" xml:"DriverResourceSpec,omitempty"`
	UpdateTimeValue      *string `json:"UpdateTimeValue,omitempty" xml:"UpdateTimeValue,omitempty"`
	SparkUI              *string `json:"SparkUI,omitempty" xml:"SparkUI,omitempty"`
	SubmitTimeValue      *string `json:"SubmitTimeValue,omitempty" xml:"SubmitTimeValue,omitempty"`
	JobName              *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	JobId                *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ExecutorInstances    *string `json:"ExecutorInstances,omitempty" xml:"ExecutorInstances,omitempty"`
	VcName               *string `json:"VcName,omitempty" xml:"VcName,omitempty"`
	SubmitTime           *string `json:"SubmitTime,omitempty" xml:"SubmitTime,omitempty"`
	Detail               *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
}

func (s GetJobDetailResponseBodyJobDetail) String() string {
	return tea.Prettify(s)
}

func (s GetJobDetailResponseBodyJobDetail) GoString() string {
	return s.String()
}

func (s *GetJobDetailResponseBodyJobDetail) SetStatus(v string) *GetJobDetailResponseBodyJobDetail {
	s.Status = &v
	return s
}

func (s *GetJobDetailResponseBodyJobDetail) SetUpdateTime(v string) *GetJobDetailResponseBodyJobDetail {
	s.UpdateTime = &v
	return s
}

func (s *GetJobDetailResponseBodyJobDetail) SetExecutorResourceSpec(v string) *GetJobDetailResponseBodyJobDetail {
	s.ExecutorResourceSpec = &v
	return s
}

func (s *GetJobDetailResponseBodyJobDetail) SetCreateTime(v string) *GetJobDetailResponseBodyJobDetail {
	s.CreateTime = &v
	return s
}

func (s *GetJobDetailResponseBodyJobDetail) SetCreateTimeValue(v string) *GetJobDetailResponseBodyJobDetail {
	s.CreateTimeValue = &v
	return s
}

func (s *GetJobDetailResponseBodyJobDetail) SetDriverResourceSpec(v string) *GetJobDetailResponseBodyJobDetail {
	s.DriverResourceSpec = &v
	return s
}

func (s *GetJobDetailResponseBodyJobDetail) SetUpdateTimeValue(v string) *GetJobDetailResponseBodyJobDetail {
	s.UpdateTimeValue = &v
	return s
}

func (s *GetJobDetailResponseBodyJobDetail) SetSparkUI(v string) *GetJobDetailResponseBodyJobDetail {
	s.SparkUI = &v
	return s
}

func (s *GetJobDetailResponseBodyJobDetail) SetSubmitTimeValue(v string) *GetJobDetailResponseBodyJobDetail {
	s.SubmitTimeValue = &v
	return s
}

func (s *GetJobDetailResponseBodyJobDetail) SetJobName(v string) *GetJobDetailResponseBodyJobDetail {
	s.JobName = &v
	return s
}

func (s *GetJobDetailResponseBodyJobDetail) SetJobId(v string) *GetJobDetailResponseBodyJobDetail {
	s.JobId = &v
	return s
}

func (s *GetJobDetailResponseBodyJobDetail) SetExecutorInstances(v string) *GetJobDetailResponseBodyJobDetail {
	s.ExecutorInstances = &v
	return s
}

func (s *GetJobDetailResponseBodyJobDetail) SetVcName(v string) *GetJobDetailResponseBodyJobDetail {
	s.VcName = &v
	return s
}

func (s *GetJobDetailResponseBodyJobDetail) SetSubmitTime(v string) *GetJobDetailResponseBodyJobDetail {
	s.SubmitTime = &v
	return s
}

func (s *GetJobDetailResponseBodyJobDetail) SetDetail(v string) *GetJobDetailResponseBodyJobDetail {
	s.Detail = &v
	return s
}

type GetJobDetailResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetJobDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetJobDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetJobDetailResponse) GoString() string {
	return s.String()
}

func (s *GetJobDetailResponse) SetHeaders(v map[string]*string) *GetJobDetailResponse {
	s.Headers = v
	return s
}

func (s *GetJobDetailResponse) SetBody(v *GetJobDetailResponseBody) *GetJobDetailResponse {
	s.Body = v
	return s
}

type GetJobLogRequest struct {
	JobId  *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	VcName *string `json:"VcName,omitempty" xml:"VcName,omitempty"`
}

func (s GetJobLogRequest) String() string {
	return tea.Prettify(s)
}

func (s GetJobLogRequest) GoString() string {
	return s.String()
}

func (s *GetJobLogRequest) SetJobId(v string) *GetJobLogRequest {
	s.JobId = &v
	return s
}

func (s *GetJobLogRequest) SetVcName(v string) *GetJobLogRequest {
	s.VcName = &v
	return s
}

type GetJobLogResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s GetJobLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetJobLogResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobLogResponseBody) SetRequestId(v string) *GetJobLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetJobLogResponseBody) SetData(v string) *GetJobLogResponseBody {
	s.Data = &v
	return s
}

type GetJobLogResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetJobLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetJobLogResponse) String() string {
	return tea.Prettify(s)
}

func (s GetJobLogResponse) GoString() string {
	return s.String()
}

func (s *GetJobLogResponse) SetHeaders(v map[string]*string) *GetJobLogResponse {
	s.Headers = v
	return s
}

func (s *GetJobLogResponse) SetBody(v *GetJobLogResponseBody) *GetJobLogResponse {
	s.Body = v
	return s
}

type GetJobStatusRequest struct {
	JobId  *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	VcName *string `json:"VcName,omitempty" xml:"VcName,omitempty"`
}

func (s GetJobStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetJobStatusRequest) GoString() string {
	return s.String()
}

func (s *GetJobStatusRequest) SetJobId(v string) *GetJobStatusRequest {
	s.JobId = &v
	return s
}

func (s *GetJobStatusRequest) SetVcName(v string) *GetJobStatusRequest {
	s.VcName = &v
	return s
}

type GetJobStatusResponseBody struct {
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetJobStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetJobStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobStatusResponseBody) SetStatus(v string) *GetJobStatusResponseBody {
	s.Status = &v
	return s
}

func (s *GetJobStatusResponseBody) SetRequestId(v string) *GetJobStatusResponseBody {
	s.RequestId = &v
	return s
}

type GetJobStatusResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetJobStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetJobStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetJobStatusResponse) GoString() string {
	return s.String()
}

func (s *GetJobStatusResponse) SetHeaders(v map[string]*string) *GetJobStatusResponse {
	s.Headers = v
	return s
}

func (s *GetJobStatusResponse) SetBody(v *GetJobStatusResponseBody) *GetJobStatusResponse {
	s.Body = v
	return s
}

type KillSparkJobRequest struct {
	VcName *string `json:"VcName,omitempty" xml:"VcName,omitempty"`
	JobId  *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s KillSparkJobRequest) String() string {
	return tea.Prettify(s)
}

func (s KillSparkJobRequest) GoString() string {
	return s.String()
}

func (s *KillSparkJobRequest) SetVcName(v string) *KillSparkJobRequest {
	s.VcName = &v
	return s
}

func (s *KillSparkJobRequest) SetJobId(v string) *KillSparkJobRequest {
	s.JobId = &v
	return s
}

type KillSparkJobResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s KillSparkJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s KillSparkJobResponseBody) GoString() string {
	return s.String()
}

func (s *KillSparkJobResponseBody) SetRequestId(v string) *KillSparkJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *KillSparkJobResponseBody) SetData(v string) *KillSparkJobResponseBody {
	s.Data = &v
	return s
}

type KillSparkJobResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *KillSparkJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s KillSparkJobResponse) String() string {
	return tea.Prettify(s)
}

func (s KillSparkJobResponse) GoString() string {
	return s.String()
}

func (s *KillSparkJobResponse) SetHeaders(v map[string]*string) *KillSparkJobResponse {
	s.Headers = v
	return s
}

func (s *KillSparkJobResponse) SetBody(v *KillSparkJobResponseBody) *KillSparkJobResponse {
	s.Body = v
	return s
}

type ListSparkJobRequest struct {
	VcName     *string `json:"VcName,omitempty" xml:"VcName,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListSparkJobRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSparkJobRequest) GoString() string {
	return s.String()
}

func (s *ListSparkJobRequest) SetVcName(v string) *ListSparkJobRequest {
	s.VcName = &v
	return s
}

func (s *ListSparkJobRequest) SetPageNumber(v int32) *ListSparkJobRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSparkJobRequest) SetPageSize(v int32) *ListSparkJobRequest {
	s.PageSize = &v
	return s
}

type ListSparkJobResponseBody struct {
	DataResult *ListSparkJobResponseBodyDataResult `json:"DataResult,omitempty" xml:"DataResult,omitempty" type:"Struct"`
	RequestId  *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListSparkJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSparkJobResponseBody) GoString() string {
	return s.String()
}

func (s *ListSparkJobResponseBody) SetDataResult(v *ListSparkJobResponseBodyDataResult) *ListSparkJobResponseBody {
	s.DataResult = v
	return s
}

func (s *ListSparkJobResponseBody) SetRequestId(v string) *ListSparkJobResponseBody {
	s.RequestId = &v
	return s
}

type ListSparkJobResponseBodyDataResult struct {
	JobList    []*ListSparkJobResponseBodyDataResultJobList `json:"JobList,omitempty" xml:"JobList,omitempty" type:"Repeated"`
	PageNumber *string                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *string                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *string                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSparkJobResponseBodyDataResult) String() string {
	return tea.Prettify(s)
}

func (s ListSparkJobResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *ListSparkJobResponseBodyDataResult) SetJobList(v []*ListSparkJobResponseBodyDataResultJobList) *ListSparkJobResponseBodyDataResult {
	s.JobList = v
	return s
}

func (s *ListSparkJobResponseBodyDataResult) SetPageNumber(v string) *ListSparkJobResponseBodyDataResult {
	s.PageNumber = &v
	return s
}

func (s *ListSparkJobResponseBodyDataResult) SetPageSize(v string) *ListSparkJobResponseBodyDataResult {
	s.PageSize = &v
	return s
}

func (s *ListSparkJobResponseBodyDataResult) SetTotalCount(v string) *ListSparkJobResponseBodyDataResult {
	s.TotalCount = &v
	return s
}

type ListSparkJobResponseBodyDataResultJobList struct {
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UpdateTime           *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	ExecutorResourceSpec *string `json:"ExecutorResourceSpec,omitempty" xml:"ExecutorResourceSpec,omitempty"`
	CreateTime           *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DriverResourceSpec   *string `json:"DriverResourceSpec,omitempty" xml:"DriverResourceSpec,omitempty"`
	CreateTimeValue      *string `json:"CreateTimeValue,omitempty" xml:"CreateTimeValue,omitempty"`
	UpdateTimeValue      *string `json:"UpdateTimeValue,omitempty" xml:"UpdateTimeValue,omitempty"`
	SparkUI              *string `json:"SparkUI,omitempty" xml:"SparkUI,omitempty"`
	SubmitTimeValue      *string `json:"SubmitTimeValue,omitempty" xml:"SubmitTimeValue,omitempty"`
	JobName              *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	JobId                *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	VcName               *string `json:"VcName,omitempty" xml:"VcName,omitempty"`
	ExecutorInstances    *string `json:"ExecutorInstances,omitempty" xml:"ExecutorInstances,omitempty"`
	SubmitTime           *string `json:"SubmitTime,omitempty" xml:"SubmitTime,omitempty"`
	Detail               *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
}

func (s ListSparkJobResponseBodyDataResultJobList) String() string {
	return tea.Prettify(s)
}

func (s ListSparkJobResponseBodyDataResultJobList) GoString() string {
	return s.String()
}

func (s *ListSparkJobResponseBodyDataResultJobList) SetStatus(v string) *ListSparkJobResponseBodyDataResultJobList {
	s.Status = &v
	return s
}

func (s *ListSparkJobResponseBodyDataResultJobList) SetUpdateTime(v string) *ListSparkJobResponseBodyDataResultJobList {
	s.UpdateTime = &v
	return s
}

func (s *ListSparkJobResponseBodyDataResultJobList) SetExecutorResourceSpec(v string) *ListSparkJobResponseBodyDataResultJobList {
	s.ExecutorResourceSpec = &v
	return s
}

func (s *ListSparkJobResponseBodyDataResultJobList) SetCreateTime(v string) *ListSparkJobResponseBodyDataResultJobList {
	s.CreateTime = &v
	return s
}

func (s *ListSparkJobResponseBodyDataResultJobList) SetDriverResourceSpec(v string) *ListSparkJobResponseBodyDataResultJobList {
	s.DriverResourceSpec = &v
	return s
}

func (s *ListSparkJobResponseBodyDataResultJobList) SetCreateTimeValue(v string) *ListSparkJobResponseBodyDataResultJobList {
	s.CreateTimeValue = &v
	return s
}

func (s *ListSparkJobResponseBodyDataResultJobList) SetUpdateTimeValue(v string) *ListSparkJobResponseBodyDataResultJobList {
	s.UpdateTimeValue = &v
	return s
}

func (s *ListSparkJobResponseBodyDataResultJobList) SetSparkUI(v string) *ListSparkJobResponseBodyDataResultJobList {
	s.SparkUI = &v
	return s
}

func (s *ListSparkJobResponseBodyDataResultJobList) SetSubmitTimeValue(v string) *ListSparkJobResponseBodyDataResultJobList {
	s.SubmitTimeValue = &v
	return s
}

func (s *ListSparkJobResponseBodyDataResultJobList) SetJobName(v string) *ListSparkJobResponseBodyDataResultJobList {
	s.JobName = &v
	return s
}

func (s *ListSparkJobResponseBodyDataResultJobList) SetJobId(v string) *ListSparkJobResponseBodyDataResultJobList {
	s.JobId = &v
	return s
}

func (s *ListSparkJobResponseBodyDataResultJobList) SetVcName(v string) *ListSparkJobResponseBodyDataResultJobList {
	s.VcName = &v
	return s
}

func (s *ListSparkJobResponseBodyDataResultJobList) SetExecutorInstances(v string) *ListSparkJobResponseBodyDataResultJobList {
	s.ExecutorInstances = &v
	return s
}

func (s *ListSparkJobResponseBodyDataResultJobList) SetSubmitTime(v string) *ListSparkJobResponseBodyDataResultJobList {
	s.SubmitTime = &v
	return s
}

func (s *ListSparkJobResponseBodyDataResultJobList) SetDetail(v string) *ListSparkJobResponseBodyDataResultJobList {
	s.Detail = &v
	return s
}

type ListSparkJobResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListSparkJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSparkJobResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSparkJobResponse) GoString() string {
	return s.String()
}

func (s *ListSparkJobResponse) SetHeaders(v map[string]*string) *ListSparkJobResponse {
	s.Headers = v
	return s
}

func (s *ListSparkJobResponse) SetBody(v *ListSparkJobResponseBody) *ListSparkJobResponse {
	s.Body = v
	return s
}

type QueryAccountListRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s QueryAccountListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAccountListRequest) GoString() string {
	return s.String()
}

func (s *QueryAccountListRequest) SetRegionId(v string) *QueryAccountListRequest {
	s.RegionId = &v
	return s
}

func (s *QueryAccountListRequest) SetPageSize(v int32) *QueryAccountListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryAccountListRequest) SetPageNumber(v int32) *QueryAccountListRequest {
	s.PageNumber = &v
	return s
}

type QueryAccountListResponseBody struct {
	TotalCount *int32                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize   *int32                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int32                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Data       []*QueryAccountListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	RegionId   *string                             `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s QueryAccountListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAccountListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAccountListResponseBody) SetTotalCount(v int32) *QueryAccountListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryAccountListResponseBody) SetPageSize(v int32) *QueryAccountListResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryAccountListResponseBody) SetRequestId(v string) *QueryAccountListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAccountListResponseBody) SetPageNumber(v int32) *QueryAccountListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *QueryAccountListResponseBody) SetData(v []*QueryAccountListResponseBodyData) *QueryAccountListResponseBody {
	s.Data = v
	return s
}

func (s *QueryAccountListResponseBody) SetRegionId(v string) *QueryAccountListResponseBody {
	s.RegionId = &v
	return s
}

type QueryAccountListResponseBodyData struct {
	Remark    *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	RamUid    *string `json:"RamUid,omitempty" xml:"RamUid,omitempty"`
	ShortName *string `json:"ShortName,omitempty" xml:"ShortName,omitempty"`
	Role      *string `json:"Role,omitempty" xml:"Role,omitempty"`
	UserName  *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s QueryAccountListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryAccountListResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryAccountListResponseBodyData) SetRemark(v string) *QueryAccountListResponseBodyData {
	s.Remark = &v
	return s
}

func (s *QueryAccountListResponseBodyData) SetRamUid(v string) *QueryAccountListResponseBodyData {
	s.RamUid = &v
	return s
}

func (s *QueryAccountListResponseBodyData) SetShortName(v string) *QueryAccountListResponseBodyData {
	s.ShortName = &v
	return s
}

func (s *QueryAccountListResponseBodyData) SetRole(v string) *QueryAccountListResponseBodyData {
	s.Role = &v
	return s
}

func (s *QueryAccountListResponseBodyData) SetUserName(v string) *QueryAccountListResponseBodyData {
	s.UserName = &v
	return s
}

type QueryAccountListResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryAccountListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryAccountListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAccountListResponse) GoString() string {
	return s.String()
}

func (s *QueryAccountListResponse) SetHeaders(v map[string]*string) *QueryAccountListResponse {
	s.Headers = v
	return s
}

func (s *QueryAccountListResponse) SetBody(v *QueryAccountListResponseBody) *QueryAccountListResponse {
	s.Body = v
	return s
}

type QueryEndPointListRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s QueryEndPointListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryEndPointListRequest) GoString() string {
	return s.String()
}

func (s *QueryEndPointListRequest) SetRegionId(v string) *QueryEndPointListRequest {
	s.RegionId = &v
	return s
}

type QueryEndPointListResponseBody struct {
	RequestId    *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	EndPointList []*QueryEndPointListResponseBodyEndPointList `json:"EndPointList,omitempty" xml:"EndPointList,omitempty" type:"Repeated"`
	RegionId     *string                                      `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s QueryEndPointListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryEndPointListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryEndPointListResponseBody) SetRequestId(v string) *QueryEndPointListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryEndPointListResponseBody) SetEndPointList(v []*QueryEndPointListResponseBodyEndPointList) *QueryEndPointListResponseBody {
	s.EndPointList = v
	return s
}

func (s *QueryEndPointListResponseBody) SetRegionId(v string) *QueryEndPointListResponseBody {
	s.RegionId = &v
	return s
}

type QueryEndPointListResponseBodyEndPointList struct {
	Product     *string `json:"product,omitempty" xml:"product,omitempty"`
	DomainURL   *string `json:"domainURL,omitempty" xml:"domainURL,omitempty"`
	Zone        *string `json:"zone,omitempty" xml:"zone,omitempty"`
	VSwitch     *string `json:"vSwitch,omitempty" xml:"vSwitch,omitempty"`
	Host        *string `json:"host,omitempty" xml:"host,omitempty"`
	AllowIP     *string `json:"allowIP,omitempty" xml:"allowIP,omitempty"`
	VpcID       *string `json:"vpcID,omitempty" xml:"vpcID,omitempty"`
	Status      *string `json:"status,omitempty" xml:"status,omitempty"`
	EndPointID  *string `json:"endPointID,omitempty" xml:"endPointID,omitempty"`
	NetworkType *string `json:"networkType,omitempty" xml:"networkType,omitempty"`
	Port        *string `json:"port,omitempty" xml:"port,omitempty"`
}

func (s QueryEndPointListResponseBodyEndPointList) String() string {
	return tea.Prettify(s)
}

func (s QueryEndPointListResponseBodyEndPointList) GoString() string {
	return s.String()
}

func (s *QueryEndPointListResponseBodyEndPointList) SetProduct(v string) *QueryEndPointListResponseBodyEndPointList {
	s.Product = &v
	return s
}

func (s *QueryEndPointListResponseBodyEndPointList) SetDomainURL(v string) *QueryEndPointListResponseBodyEndPointList {
	s.DomainURL = &v
	return s
}

func (s *QueryEndPointListResponseBodyEndPointList) SetZone(v string) *QueryEndPointListResponseBodyEndPointList {
	s.Zone = &v
	return s
}

func (s *QueryEndPointListResponseBodyEndPointList) SetVSwitch(v string) *QueryEndPointListResponseBodyEndPointList {
	s.VSwitch = &v
	return s
}

func (s *QueryEndPointListResponseBodyEndPointList) SetHost(v string) *QueryEndPointListResponseBodyEndPointList {
	s.Host = &v
	return s
}

func (s *QueryEndPointListResponseBodyEndPointList) SetAllowIP(v string) *QueryEndPointListResponseBodyEndPointList {
	s.AllowIP = &v
	return s
}

func (s *QueryEndPointListResponseBodyEndPointList) SetVpcID(v string) *QueryEndPointListResponseBodyEndPointList {
	s.VpcID = &v
	return s
}

func (s *QueryEndPointListResponseBodyEndPointList) SetStatus(v string) *QueryEndPointListResponseBodyEndPointList {
	s.Status = &v
	return s
}

func (s *QueryEndPointListResponseBodyEndPointList) SetEndPointID(v string) *QueryEndPointListResponseBodyEndPointList {
	s.EndPointID = &v
	return s
}

func (s *QueryEndPointListResponseBodyEndPointList) SetNetworkType(v string) *QueryEndPointListResponseBodyEndPointList {
	s.NetworkType = &v
	return s
}

func (s *QueryEndPointListResponseBodyEndPointList) SetPort(v string) *QueryEndPointListResponseBodyEndPointList {
	s.Port = &v
	return s
}

type QueryEndPointListResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryEndPointListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryEndPointListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryEndPointListResponse) GoString() string {
	return s.String()
}

func (s *QueryEndPointListResponse) SetHeaders(v map[string]*string) *QueryEndPointListResponse {
	s.Headers = v
	return s
}

func (s *QueryEndPointListResponse) SetBody(v *QueryEndPointListResponseBody) *QueryEndPointListResponse {
	s.Body = v
	return s
}

type ReleaseInstanceRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ReleaseInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleaseInstanceRequest) GoString() string {
	return s.String()
}

func (s *ReleaseInstanceRequest) SetRegionId(v string) *ReleaseInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *ReleaseInstanceRequest) SetInstanceId(v string) *ReleaseInstanceRequest {
	s.InstanceId = &v
	return s
}

type ReleaseInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorInfo *string `json:"ErrorInfo,omitempty" xml:"ErrorInfo,omitempty"`
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Result    *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s ReleaseInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReleaseInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseInstanceResponseBody) SetRequestId(v string) *ReleaseInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReleaseInstanceResponseBody) SetErrorInfo(v string) *ReleaseInstanceResponseBody {
	s.ErrorInfo = &v
	return s
}

func (s *ReleaseInstanceResponseBody) SetErrorCode(v string) *ReleaseInstanceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ReleaseInstanceResponseBody) SetSuccess(v bool) *ReleaseInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *ReleaseInstanceResponseBody) SetResult(v string) *ReleaseInstanceResponseBody {
	s.Result = &v
	return s
}

type ReleaseInstanceResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ReleaseInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReleaseInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleaseInstanceResponse) GoString() string {
	return s.String()
}

func (s *ReleaseInstanceResponse) SetHeaders(v map[string]*string) *ReleaseInstanceResponse {
	s.Headers = v
	return s
}

func (s *ReleaseInstanceResponse) SetBody(v *ReleaseInstanceResponseBody) *ReleaseInstanceResponse {
	s.Body = v
	return s
}

type RemoveEndPointRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	EndPointID *string `json:"EndPointID,omitempty" xml:"EndPointID,omitempty"`
}

func (s RemoveEndPointRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveEndPointRequest) GoString() string {
	return s.String()
}

func (s *RemoveEndPointRequest) SetRegionId(v string) *RemoveEndPointRequest {
	s.RegionId = &v
	return s
}

func (s *RemoveEndPointRequest) SetEndPointID(v string) *RemoveEndPointRequest {
	s.EndPointID = &v
	return s
}

type RemoveEndPointResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RemoveEndPointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveEndPointResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveEndPointResponseBody) SetRequestId(v string) *RemoveEndPointResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveEndPointResponseBody) SetRegionId(v string) *RemoveEndPointResponseBody {
	s.RegionId = &v
	return s
}

type RemoveEndPointResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveEndPointResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveEndPointResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveEndPointResponse) GoString() string {
	return s.String()
}

func (s *RemoveEndPointResponse) SetHeaders(v map[string]*string) *RemoveEndPointResponse {
	s.Headers = v
	return s
}

func (s *RemoveEndPointResponse) SetBody(v *RemoveEndPointResponseBody) *RemoveEndPointResponse {
	s.Body = v
	return s
}

type ResetMainPasswordRequest struct {
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UseRandomPassword    *bool   `json:"UseRandomPassword,omitempty" xml:"UseRandomPassword,omitempty"`
	InitPassword         *string `json:"InitPassword,omitempty" xml:"InitPassword,omitempty"`
	EnableKMS            *bool   `json:"EnableKMS,omitempty" xml:"EnableKMS,omitempty"`
	ExternalUid          *string `json:"ExternalUid,omitempty" xml:"ExternalUid,omitempty"`
	ExternalAliyunUid    *string `json:"ExternalAliyunUid,omitempty" xml:"ExternalAliyunUid,omitempty"`
	ExternalBizAliyunUid *string `json:"ExternalBizAliyunUid,omitempty" xml:"ExternalBizAliyunUid,omitempty"`
}

func (s ResetMainPasswordRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetMainPasswordRequest) GoString() string {
	return s.String()
}

func (s *ResetMainPasswordRequest) SetRegionId(v string) *ResetMainPasswordRequest {
	s.RegionId = &v
	return s
}

func (s *ResetMainPasswordRequest) SetUseRandomPassword(v bool) *ResetMainPasswordRequest {
	s.UseRandomPassword = &v
	return s
}

func (s *ResetMainPasswordRequest) SetInitPassword(v string) *ResetMainPasswordRequest {
	s.InitPassword = &v
	return s
}

func (s *ResetMainPasswordRequest) SetEnableKMS(v bool) *ResetMainPasswordRequest {
	s.EnableKMS = &v
	return s
}

func (s *ResetMainPasswordRequest) SetExternalUid(v string) *ResetMainPasswordRequest {
	s.ExternalUid = &v
	return s
}

func (s *ResetMainPasswordRequest) SetExternalAliyunUid(v string) *ResetMainPasswordRequest {
	s.ExternalAliyunUid = &v
	return s
}

func (s *ResetMainPasswordRequest) SetExternalBizAliyunUid(v string) *ResetMainPasswordRequest {
	s.ExternalBizAliyunUid = &v
	return s
}

type ResetMainPasswordResponseBody struct {
	Account   *ResetMainPasswordResponseBodyAccount `json:"Account,omitempty" xml:"Account,omitempty" type:"Struct"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RegionId  *string                               `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ResetMainPasswordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetMainPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ResetMainPasswordResponseBody) SetAccount(v *ResetMainPasswordResponseBodyAccount) *ResetMainPasswordResponseBody {
	s.Account = v
	return s
}

func (s *ResetMainPasswordResponseBody) SetRequestId(v string) *ResetMainPasswordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetMainPasswordResponseBody) SetRegionId(v string) *ResetMainPasswordResponseBody {
	s.RegionId = &v
	return s
}

type ResetMainPasswordResponseBodyAccount struct {
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ResetMainPasswordResponseBodyAccount) String() string {
	return tea.Prettify(s)
}

func (s ResetMainPasswordResponseBodyAccount) GoString() string {
	return s.String()
}

func (s *ResetMainPasswordResponseBodyAccount) SetPassword(v string) *ResetMainPasswordResponseBodyAccount {
	s.Password = &v
	return s
}

func (s *ResetMainPasswordResponseBodyAccount) SetUserName(v string) *ResetMainPasswordResponseBodyAccount {
	s.UserName = &v
	return s
}

type ResetMainPasswordResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ResetMainPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResetMainPasswordResponse) String() string {
	return tea.Prettify(s)
}

func (s ResetMainPasswordResponse) GoString() string {
	return s.String()
}

func (s *ResetMainPasswordResponse) SetHeaders(v map[string]*string) *ResetMainPasswordResponse {
	s.Headers = v
	return s
}

func (s *ResetMainPasswordResponse) SetBody(v *ResetMainPasswordResponseBody) *ResetMainPasswordResponse {
	s.Body = v
	return s
}

type SetAllowIPRequest struct {
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	Product     *string `json:"Product,omitempty" xml:"Product,omitempty"`
	AllowIP     *string `json:"AllowIP,omitempty" xml:"AllowIP,omitempty"`
	Append      *bool   `json:"Append,omitempty" xml:"Append,omitempty"`
}

func (s SetAllowIPRequest) String() string {
	return tea.Prettify(s)
}

func (s SetAllowIPRequest) GoString() string {
	return s.String()
}

func (s *SetAllowIPRequest) SetRegionId(v string) *SetAllowIPRequest {
	s.RegionId = &v
	return s
}

func (s *SetAllowIPRequest) SetNetworkType(v string) *SetAllowIPRequest {
	s.NetworkType = &v
	return s
}

func (s *SetAllowIPRequest) SetProduct(v string) *SetAllowIPRequest {
	s.Product = &v
	return s
}

func (s *SetAllowIPRequest) SetAllowIP(v string) *SetAllowIPRequest {
	s.AllowIP = &v
	return s
}

func (s *SetAllowIPRequest) SetAppend(v bool) *SetAllowIPRequest {
	s.Append = &v
	return s
}

type SetAllowIPResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SetAllowIPResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetAllowIPResponseBody) GoString() string {
	return s.String()
}

func (s *SetAllowIPResponseBody) SetRequestId(v string) *SetAllowIPResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetAllowIPResponseBody) SetRegionId(v string) *SetAllowIPResponseBody {
	s.RegionId = &v
	return s
}

type SetAllowIPResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetAllowIPResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetAllowIPResponse) String() string {
	return tea.Prettify(s)
}

func (s SetAllowIPResponse) GoString() string {
	return s.String()
}

func (s *SetAllowIPResponse) SetHeaders(v map[string]*string) *SetAllowIPResponse {
	s.Headers = v
	return s
}

func (s *SetAllowIPResponse) SetBody(v *SetAllowIPResponseBody) *SetAllowIPResponse {
	s.Body = v
	return s
}

type SubmitSparkJobRequest struct {
	VcName     *string `json:"VcName,omitempty" xml:"VcName,omitempty"`
	ConfigJson *string `json:"ConfigJson,omitempty" xml:"ConfigJson,omitempty"`
}

func (s SubmitSparkJobRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitSparkJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitSparkJobRequest) SetVcName(v string) *SubmitSparkJobRequest {
	s.VcName = &v
	return s
}

func (s *SubmitSparkJobRequest) SetConfigJson(v string) *SubmitSparkJobRequest {
	s.ConfigJson = &v
	return s
}

type SubmitSparkJobResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s SubmitSparkJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitSparkJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitSparkJobResponseBody) SetRequestId(v string) *SubmitSparkJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitSparkJobResponseBody) SetJobId(v string) *SubmitSparkJobResponseBody {
	s.JobId = &v
	return s
}

type SubmitSparkJobResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SubmitSparkJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitSparkJobResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitSparkJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitSparkJobResponse) SetHeaders(v map[string]*string) *SubmitSparkJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitSparkJobResponse) SetBody(v *SubmitSparkJobResponseBody) *SubmitSparkJobResponse {
	s.Body = v
	return s
}

type UnSubscribeRegionRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UnSubscribeRegionRequest) String() string {
	return tea.Prettify(s)
}

func (s UnSubscribeRegionRequest) GoString() string {
	return s.String()
}

func (s *UnSubscribeRegionRequest) SetRegionId(v string) *UnSubscribeRegionRequest {
	s.RegionId = &v
	return s
}

type UnSubscribeRegionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UnSubscribeRegionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnSubscribeRegionResponseBody) GoString() string {
	return s.String()
}

func (s *UnSubscribeRegionResponseBody) SetRequestId(v string) *UnSubscribeRegionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnSubscribeRegionResponseBody) SetRegionId(v string) *UnSubscribeRegionResponseBody {
	s.RegionId = &v
	return s
}

type UnSubscribeRegionResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UnSubscribeRegionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnSubscribeRegionResponse) String() string {
	return tea.Prettify(s)
}

func (s UnSubscribeRegionResponse) GoString() string {
	return s.String()
}

func (s *UnSubscribeRegionResponse) SetHeaders(v map[string]*string) *UnSubscribeRegionResponse {
	s.Headers = v
	return s
}

func (s *UnSubscribeRegionResponse) SetBody(v *UnSubscribeRegionResponseBody) *UnSubscribeRegionResponse {
	s.Body = v
	return s
}

type UpdateAccountPasswordRequest struct {
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	AccountName       *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	IsShort           *bool   `json:"IsShort,omitempty" xml:"IsShort,omitempty"`
	Password          *string `json:"Password,omitempty" xml:"Password,omitempty"`
	EnableKMS         *bool   `json:"EnableKMS,omitempty" xml:"EnableKMS,omitempty"`
	UseRandomPassword *bool   `json:"UseRandomPassword,omitempty" xml:"UseRandomPassword,omitempty"`
}

func (s UpdateAccountPasswordRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAccountPasswordRequest) GoString() string {
	return s.String()
}

func (s *UpdateAccountPasswordRequest) SetRegionId(v string) *UpdateAccountPasswordRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateAccountPasswordRequest) SetAccountName(v string) *UpdateAccountPasswordRequest {
	s.AccountName = &v
	return s
}

func (s *UpdateAccountPasswordRequest) SetIsShort(v bool) *UpdateAccountPasswordRequest {
	s.IsShort = &v
	return s
}

func (s *UpdateAccountPasswordRequest) SetPassword(v string) *UpdateAccountPasswordRequest {
	s.Password = &v
	return s
}

func (s *UpdateAccountPasswordRequest) SetEnableKMS(v bool) *UpdateAccountPasswordRequest {
	s.EnableKMS = &v
	return s
}

func (s *UpdateAccountPasswordRequest) SetUseRandomPassword(v bool) *UpdateAccountPasswordRequest {
	s.UseRandomPassword = &v
	return s
}

type UpdateAccountPasswordResponseBody struct {
	Account   *UpdateAccountPasswordResponseBodyAccount `json:"Account,omitempty" xml:"Account,omitempty" type:"Struct"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RegionId  *string                                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateAccountPasswordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAccountPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAccountPasswordResponseBody) SetAccount(v *UpdateAccountPasswordResponseBodyAccount) *UpdateAccountPasswordResponseBody {
	s.Account = v
	return s
}

func (s *UpdateAccountPasswordResponseBody) SetRequestId(v string) *UpdateAccountPasswordResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAccountPasswordResponseBody) SetRegionId(v string) *UpdateAccountPasswordResponseBody {
	s.RegionId = &v
	return s
}

type UpdateAccountPasswordResponseBodyAccount struct {
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s UpdateAccountPasswordResponseBodyAccount) String() string {
	return tea.Prettify(s)
}

func (s UpdateAccountPasswordResponseBodyAccount) GoString() string {
	return s.String()
}

func (s *UpdateAccountPasswordResponseBodyAccount) SetPassword(v string) *UpdateAccountPasswordResponseBodyAccount {
	s.Password = &v
	return s
}

func (s *UpdateAccountPasswordResponseBodyAccount) SetUserName(v string) *UpdateAccountPasswordResponseBodyAccount {
	s.UserName = &v
	return s
}

type UpdateAccountPasswordResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateAccountPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAccountPasswordResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAccountPasswordResponse) GoString() string {
	return s.String()
}

func (s *UpdateAccountPasswordResponse) SetHeaders(v map[string]*string) *UpdateAccountPasswordResponse {
	s.Headers = v
	return s
}

func (s *UpdateAccountPasswordResponse) SetBody(v *UpdateAccountPasswordResponseBody) *UpdateAccountPasswordResponse {
	s.Body = v
	return s
}

type UpgradeInstanceRequest struct {
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ChargeType   *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	Component    *string `json:"Component,omitempty" xml:"Component,omitempty"`
}

func (s UpgradeInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpgradeInstanceRequest) SetRegionId(v string) *UpgradeInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *UpgradeInstanceRequest) SetInstanceId(v string) *UpgradeInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *UpgradeInstanceRequest) SetChargeType(v string) *UpgradeInstanceRequest {
	s.ChargeType = &v
	return s
}

func (s *UpgradeInstanceRequest) SetInstanceType(v string) *UpgradeInstanceRequest {
	s.InstanceType = &v
	return s
}

func (s *UpgradeInstanceRequest) SetComponent(v string) *UpgradeInstanceRequest {
	s.Component = &v
	return s
}

type UpgradeInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorInfo *string `json:"ErrorInfo,omitempty" xml:"ErrorInfo,omitempty"`
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Result    *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s UpgradeInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpgradeInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeInstanceResponseBody) SetRequestId(v string) *UpgradeInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeInstanceResponseBody) SetErrorInfo(v string) *UpgradeInstanceResponseBody {
	s.ErrorInfo = &v
	return s
}

func (s *UpgradeInstanceResponseBody) SetErrorCode(v string) *UpgradeInstanceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpgradeInstanceResponseBody) SetSuccess(v bool) *UpgradeInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *UpgradeInstanceResponseBody) SetResult(v string) *UpgradeInstanceResponseBody {
	s.Result = &v
	return s
}

type UpgradeInstanceResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpgradeInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpgradeInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpgradeInstanceResponse) GoString() string {
	return s.String()
}

func (s *UpgradeInstanceResponse) SetHeaders(v map[string]*string) *UpgradeInstanceResponse {
	s.Headers = v
	return s
}

func (s *UpgradeInstanceResponse) SetBody(v *UpgradeInstanceResponseBody) *UpgradeInstanceResponse {
	s.Body = v
	return s
}

type ValidateVirtualClusterNameRequest struct {
	VcName *string `json:"VcName,omitempty" xml:"VcName,omitempty"`
}

func (s ValidateVirtualClusterNameRequest) String() string {
	return tea.Prettify(s)
}

func (s ValidateVirtualClusterNameRequest) GoString() string {
	return s.String()
}

func (s *ValidateVirtualClusterNameRequest) SetVcName(v string) *ValidateVirtualClusterNameRequest {
	s.VcName = &v
	return s
}

type ValidateVirtualClusterNameResponseBody struct {
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *ValidateVirtualClusterNameResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s ValidateVirtualClusterNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ValidateVirtualClusterNameResponseBody) GoString() string {
	return s.String()
}

func (s *ValidateVirtualClusterNameResponseBody) SetRequestId(v string) *ValidateVirtualClusterNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *ValidateVirtualClusterNameResponseBody) SetData(v *ValidateVirtualClusterNameResponseBodyData) *ValidateVirtualClusterNameResponseBody {
	s.Data = v
	return s
}

type ValidateVirtualClusterNameResponseBodyData struct {
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Legal   *string `json:"Legal,omitempty" xml:"Legal,omitempty"`
}

func (s ValidateVirtualClusterNameResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ValidateVirtualClusterNameResponseBodyData) GoString() string {
	return s.String()
}

func (s *ValidateVirtualClusterNameResponseBodyData) SetMessage(v string) *ValidateVirtualClusterNameResponseBodyData {
	s.Message = &v
	return s
}

func (s *ValidateVirtualClusterNameResponseBodyData) SetLegal(v string) *ValidateVirtualClusterNameResponseBodyData {
	s.Legal = &v
	return s
}

type ValidateVirtualClusterNameResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ValidateVirtualClusterNameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ValidateVirtualClusterNameResponse) String() string {
	return tea.Prettify(s)
}

func (s ValidateVirtualClusterNameResponse) GoString() string {
	return s.String()
}

func (s *ValidateVirtualClusterNameResponse) SetHeaders(v map[string]*string) *ValidateVirtualClusterNameResponse {
	s.Headers = v
	return s
}

func (s *ValidateVirtualClusterNameResponse) SetBody(v *ValidateVirtualClusterNameResponseBody) *ValidateVirtualClusterNameResponse {
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
	client.EndpointMap = map[string]*string{
		"cn-beijing":                  tea.String("openanalytics.cn-beijing.aliyuncs.com"),
		"cn-zhangjiakou":              tea.String("openanalytics.cn-zhangjiakou.aliyuncs.com"),
		"cn-hangzhou":                 tea.String("openanalytics.cn-hangzhou.aliyuncs.com"),
		"cn-shanghai":                 tea.String("openanalytics.cn-shanghai.aliyuncs.com"),
		"cn-shenzhen":                 tea.String("openanalytics.cn-shenzhen.aliyuncs.com"),
		"cn-hongkong":                 tea.String("openanalytics.cn-hongkong.aliyuncs.com"),
		"ap-southeast-1":              tea.String("openanalytics.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-2":              tea.String("datalakeanalytics.ap-southeast-2.aliyuncs.com"),
		"ap-southeast-3":              tea.String("openanalytics.ap-southeast-3.aliyuncs.com"),
		"ap-northeast-1":              tea.String("datalakeanalytics.ap-northeast-1.aliyuncs.com"),
		"eu-west-1":                   tea.String("openanalytics.eu-west-1.aliyuncs.com"),
		"us-west-1":                   tea.String("openanalytics.us-west-1.aliyuncs.com"),
		"us-east-1":                   tea.String("datalakeanalytics.us-east-1.aliyuncs.com"),
		"eu-central-1":                tea.String("datalakeanalytics.eu-central-1.aliyuncs.com"),
		"ap-south-1":                  tea.String("openanalytics.ap-south-1.aliyuncs.com"),
		"ap-northeast-2-pop":          tea.String("openanalytics.ap-northeast-1.aliyuncs.com"),
		"ap-southeast-5":              tea.String("openanalytics.ap-southeast-5.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("openanalytics.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("openanalytics.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("openanalytics.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("openanalytics.aliyuncs.com"),
		"cn-chengdu":                  tea.String("openanalytics.aliyuncs.com"),
		"cn-edge-1":                   tea.String("openanalytics.aliyuncs.com"),
		"cn-fujian":                   tea.String("openanalytics.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("openanalytics.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("openanalytics.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("openanalytics.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("openanalytics.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("openanalytics.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("openanalytics.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("openanalytics.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("openanalytics.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("openanalytics.aliyuncs.com"),
		"cn-huhehaote":                tea.String("openanalytics.cn-huhehaote.aliyuncs.com"),
		"cn-north-2-gov-1":            tea.String("openanalytics.aliyuncs.com"),
		"cn-qingdao":                  tea.String("openanalytics.cn-qingdao.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("openanalytics.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("openanalytics.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("openanalytics.aliyuncs.com"),
		"cn-shanghai-finance-1":       tea.String("openanalytics.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("openanalytics.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("openanalytics.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("openanalytics.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("openanalytics.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("openanalytics.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("openanalytics.aliyuncs.com"),
		"cn-wuhan":                    tea.String("openanalytics.aliyuncs.com"),
		"cn-yushanfang":               tea.String("openanalytics.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("openanalytics.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("openanalytics.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("openanalytics.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("openanalytics.ap-northeast-1.aliyuncs.com"),
		"me-east-1":                   tea.String("openanalytics.me-east-1.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("openanalytics.ap-northeast-1.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("openanalytics-open"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AddAccountWithOptions(request *AddAccountRequest, runtime *util.RuntimeOptions) (_result *AddAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddAccountResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddAccount"), tea.String("2018-06-19"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddAccount(request *AddAccountRequest) (_result *AddAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddAccountResponse{}
	_body, _err := client.AddAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddEndPointWithOptions(request *AddEndPointRequest, runtime *util.RuntimeOptions) (_result *AddEndPointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddEndPointResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddEndPoint"), tea.String("2018-06-19"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddEndPoint(request *AddEndPointRequest) (_result *AddEndPointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddEndPointResponse{}
	_body, _err := client.AddEndPointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateInstanceWithOptions(request *CreateInstanceRequest, runtime *util.RuntimeOptions) (_result *CreateInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateInstance"), tea.String("2018-06-19"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateInstance(request *CreateInstanceRequest) (_result *CreateInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateInstanceResponse{}
	_body, _err := client.CreateInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAccountWithOptions(request *DeleteAccountRequest, runtime *util.RuntimeOptions) (_result *DeleteAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteAccountResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteAccount"), tea.String("2018-06-19"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAccount(request *DeleteAccountRequest) (_result *DeleteAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAccountResponse{}
	_body, _err := client.DeleteAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAllowIPWithOptions(request *GetAllowIPRequest, runtime *util.RuntimeOptions) (_result *GetAllowIPResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetAllowIPResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetAllowIP"), tea.String("2018-06-19"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAllowIP(request *GetAllowIPRequest) (_result *GetAllowIPResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAllowIPResponse{}
	_body, _err := client.GetAllowIPWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetEndPointWithOptions(request *GetEndPointRequest, runtime *util.RuntimeOptions) (_result *GetEndPointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetEndPointResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetEndPoint"), tea.String("2018-06-19"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetEndPoint(request *GetEndPointRequest) (_result *GetEndPointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetEndPointResponse{}
	_body, _err := client.GetEndPointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetEndPointByDomainWithOptions(request *GetEndPointByDomainRequest, runtime *util.RuntimeOptions) (_result *GetEndPointByDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetEndPointByDomainResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetEndPointByDomain"), tea.String("2018-06-19"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetEndPointByDomain(request *GetEndPointByDomainRequest) (_result *GetEndPointByDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetEndPointByDomainResponse{}
	_body, _err := client.GetEndPointByDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetJobDetailWithOptions(request *GetJobDetailRequest, runtime *util.RuntimeOptions) (_result *GetJobDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetJobDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetJobDetail"), tea.String("2018-06-19"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetJobDetail(request *GetJobDetailRequest) (_result *GetJobDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetJobDetailResponse{}
	_body, _err := client.GetJobDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetJobLogWithOptions(request *GetJobLogRequest, runtime *util.RuntimeOptions) (_result *GetJobLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetJobLogResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetJobLog"), tea.String("2018-06-19"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetJobLog(request *GetJobLogRequest) (_result *GetJobLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetJobLogResponse{}
	_body, _err := client.GetJobLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetJobStatusWithOptions(request *GetJobStatusRequest, runtime *util.RuntimeOptions) (_result *GetJobStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetJobStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetJobStatus"), tea.String("2018-06-19"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetJobStatus(request *GetJobStatusRequest) (_result *GetJobStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetJobStatusResponse{}
	_body, _err := client.GetJobStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) KillSparkJobWithOptions(request *KillSparkJobRequest, runtime *util.RuntimeOptions) (_result *KillSparkJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &KillSparkJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("KillSparkJob"), tea.String("2018-06-19"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) KillSparkJob(request *KillSparkJobRequest) (_result *KillSparkJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &KillSparkJobResponse{}
	_body, _err := client.KillSparkJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSparkJobWithOptions(request *ListSparkJobRequest, runtime *util.RuntimeOptions) (_result *ListSparkJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListSparkJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListSparkJob"), tea.String("2018-06-19"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSparkJob(request *ListSparkJobRequest) (_result *ListSparkJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSparkJobResponse{}
	_body, _err := client.ListSparkJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryAccountListWithOptions(request *QueryAccountListRequest, runtime *util.RuntimeOptions) (_result *QueryAccountListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryAccountListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryAccountList"), tea.String("2018-06-19"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryAccountList(request *QueryAccountListRequest) (_result *QueryAccountListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryAccountListResponse{}
	_body, _err := client.QueryAccountListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryEndPointListWithOptions(request *QueryEndPointListRequest, runtime *util.RuntimeOptions) (_result *QueryEndPointListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryEndPointListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryEndPointList"), tea.String("2018-06-19"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryEndPointList(request *QueryEndPointListRequest) (_result *QueryEndPointListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryEndPointListResponse{}
	_body, _err := client.QueryEndPointListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReleaseInstanceWithOptions(request *ReleaseInstanceRequest, runtime *util.RuntimeOptions) (_result *ReleaseInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ReleaseInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ReleaseInstance"), tea.String("2018-06-19"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReleaseInstance(request *ReleaseInstanceRequest) (_result *ReleaseInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReleaseInstanceResponse{}
	_body, _err := client.ReleaseInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveEndPointWithOptions(request *RemoveEndPointRequest, runtime *util.RuntimeOptions) (_result *RemoveEndPointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RemoveEndPointResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RemoveEndPoint"), tea.String("2018-06-19"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveEndPoint(request *RemoveEndPointRequest) (_result *RemoveEndPointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveEndPointResponse{}
	_body, _err := client.RemoveEndPointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResetMainPasswordWithOptions(request *ResetMainPasswordRequest, runtime *util.RuntimeOptions) (_result *ResetMainPasswordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ResetMainPasswordResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ResetMainPassword"), tea.String("2018-06-19"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResetMainPassword(request *ResetMainPasswordRequest) (_result *ResetMainPasswordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResetMainPasswordResponse{}
	_body, _err := client.ResetMainPasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetAllowIPWithOptions(request *SetAllowIPRequest, runtime *util.RuntimeOptions) (_result *SetAllowIPResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetAllowIPResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetAllowIP"), tea.String("2018-06-19"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetAllowIP(request *SetAllowIPRequest) (_result *SetAllowIPResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetAllowIPResponse{}
	_body, _err := client.SetAllowIPWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitSparkJobWithOptions(request *SubmitSparkJobRequest, runtime *util.RuntimeOptions) (_result *SubmitSparkJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SubmitSparkJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SubmitSparkJob"), tea.String("2018-06-19"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitSparkJob(request *SubmitSparkJobRequest) (_result *SubmitSparkJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitSparkJobResponse{}
	_body, _err := client.SubmitSparkJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnSubscribeRegionWithOptions(request *UnSubscribeRegionRequest, runtime *util.RuntimeOptions) (_result *UnSubscribeRegionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UnSubscribeRegionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UnSubscribeRegion"), tea.String("2018-06-19"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnSubscribeRegion(request *UnSubscribeRegionRequest) (_result *UnSubscribeRegionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnSubscribeRegionResponse{}
	_body, _err := client.UnSubscribeRegionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAccountPasswordWithOptions(request *UpdateAccountPasswordRequest, runtime *util.RuntimeOptions) (_result *UpdateAccountPasswordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateAccountPasswordResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateAccountPassword"), tea.String("2018-06-19"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAccountPassword(request *UpdateAccountPasswordRequest) (_result *UpdateAccountPasswordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAccountPasswordResponse{}
	_body, _err := client.UpdateAccountPasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpgradeInstanceWithOptions(request *UpgradeInstanceRequest, runtime *util.RuntimeOptions) (_result *UpgradeInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpgradeInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpgradeInstance"), tea.String("2018-06-19"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpgradeInstance(request *UpgradeInstanceRequest) (_result *UpgradeInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpgradeInstanceResponse{}
	_body, _err := client.UpgradeInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ValidateVirtualClusterNameWithOptions(request *ValidateVirtualClusterNameRequest, runtime *util.RuntimeOptions) (_result *ValidateVirtualClusterNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ValidateVirtualClusterNameResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ValidateVirtualClusterName"), tea.String("2018-06-19"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ValidateVirtualClusterName(request *ValidateVirtualClusterNameRequest) (_result *ValidateVirtualClusterNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ValidateVirtualClusterNameResponse{}
	_body, _err := client.ValidateVirtualClusterNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
