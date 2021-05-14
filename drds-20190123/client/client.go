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

type CheckDrdsDbNameRequest struct {
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
}

func (s CheckDrdsDbNameRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckDrdsDbNameRequest) GoString() string {
	return s.String()
}

func (s *CheckDrdsDbNameRequest) SetDrdsInstanceId(v string) *CheckDrdsDbNameRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *CheckDrdsDbNameRequest) SetDbName(v string) *CheckDrdsDbNameRequest {
	s.DbName = &v
	return s
}

type CheckDrdsDbNameResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckDrdsDbNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckDrdsDbNameResponseBody) GoString() string {
	return s.String()
}

func (s *CheckDrdsDbNameResponseBody) SetRequestId(v string) *CheckDrdsDbNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckDrdsDbNameResponseBody) SetResult(v bool) *CheckDrdsDbNameResponseBody {
	s.Result = &v
	return s
}

func (s *CheckDrdsDbNameResponseBody) SetSuccess(v bool) *CheckDrdsDbNameResponseBody {
	s.Success = &v
	return s
}

type CheckDrdsDbNameResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CheckDrdsDbNameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckDrdsDbNameResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckDrdsDbNameResponse) GoString() string {
	return s.String()
}

func (s *CheckDrdsDbNameResponse) SetHeaders(v map[string]*string) *CheckDrdsDbNameResponse {
	s.Headers = v
	return s
}

func (s *CheckDrdsDbNameResponse) SetBody(v *CheckDrdsDbNameResponseBody) *CheckDrdsDbNameResponse {
	s.Body = v
	return s
}

type CheckExpandStatusRequest struct {
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
}

func (s CheckExpandStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckExpandStatusRequest) GoString() string {
	return s.String()
}

func (s *CheckExpandStatusRequest) SetDrdsInstanceId(v string) *CheckExpandStatusRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *CheckExpandStatusRequest) SetDbName(v string) *CheckExpandStatusRequest {
	s.DbName = &v
	return s
}

type CheckExpandStatusResponseBody struct {
	Success   *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *CheckExpandStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s CheckExpandStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckExpandStatusResponseBody) GoString() string {
	return s.String()
}

func (s *CheckExpandStatusResponseBody) SetSuccess(v bool) *CheckExpandStatusResponseBody {
	s.Success = &v
	return s
}

func (s *CheckExpandStatusResponseBody) SetRequestId(v string) *CheckExpandStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckExpandStatusResponseBody) SetData(v *CheckExpandStatusResponseBodyData) *CheckExpandStatusResponseBody {
	s.Data = v
	return s
}

type CheckExpandStatusResponseBodyData struct {
	Msg      *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	IsActive *bool   `json:"IsActive,omitempty" xml:"IsActive,omitempty"`
}

func (s CheckExpandStatusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CheckExpandStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *CheckExpandStatusResponseBodyData) SetMsg(v string) *CheckExpandStatusResponseBodyData {
	s.Msg = &v
	return s
}

func (s *CheckExpandStatusResponseBodyData) SetIsActive(v bool) *CheckExpandStatusResponseBodyData {
	s.IsActive = &v
	return s
}

type CheckExpandStatusResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CheckExpandStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckExpandStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckExpandStatusResponse) GoString() string {
	return s.String()
}

func (s *CheckExpandStatusResponse) SetHeaders(v map[string]*string) *CheckExpandStatusResponse {
	s.Headers = v
	return s
}

func (s *CheckExpandStatusResponse) SetBody(v *CheckExpandStatusResponseBody) *CheckExpandStatusResponse {
	s.Body = v
	return s
}

type CheckSqlAuditEnableStatusRequest struct {
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
}

func (s CheckSqlAuditEnableStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckSqlAuditEnableStatusRequest) GoString() string {
	return s.String()
}

func (s *CheckSqlAuditEnableStatusRequest) SetDrdsInstanceId(v string) *CheckSqlAuditEnableStatusRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *CheckSqlAuditEnableStatusRequest) SetDbName(v string) *CheckSqlAuditEnableStatusRequest {
	s.DbName = &v
	return s
}

type CheckSqlAuditEnableStatusResponseBody struct {
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckSqlAuditEnableStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckSqlAuditEnableStatusResponseBody) GoString() string {
	return s.String()
}

func (s *CheckSqlAuditEnableStatusResponseBody) SetStatus(v string) *CheckSqlAuditEnableStatusResponseBody {
	s.Status = &v
	return s
}

func (s *CheckSqlAuditEnableStatusResponseBody) SetRequestId(v string) *CheckSqlAuditEnableStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckSqlAuditEnableStatusResponseBody) SetSuccess(v bool) *CheckSqlAuditEnableStatusResponseBody {
	s.Success = &v
	return s
}

type CheckSqlAuditEnableStatusResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CheckSqlAuditEnableStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckSqlAuditEnableStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckSqlAuditEnableStatusResponse) GoString() string {
	return s.String()
}

func (s *CheckSqlAuditEnableStatusResponse) SetHeaders(v map[string]*string) *CheckSqlAuditEnableStatusResponse {
	s.Headers = v
	return s
}

func (s *CheckSqlAuditEnableStatusResponse) SetBody(v *CheckSqlAuditEnableStatusResponseBody) *CheckSqlAuditEnableStatusResponse {
	s.Body = v
	return s
}

type CreateDrdsDBRequest struct {
	DrdsInstanceId       *string                               `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	DbName               *string                               `json:"DbName,omitempty" xml:"DbName,omitempty"`
	Encode               *string                               `json:"Encode,omitempty" xml:"Encode,omitempty"`
	Password             *string                               `json:"Password,omitempty" xml:"Password,omitempty"`
	Type                 *string                               `json:"Type,omitempty" xml:"Type,omitempty"`
	DbInstType           *string                               `json:"DbInstType,omitempty" xml:"DbInstType,omitempty"`
	DbInstanceIsCreating *bool                                 `json:"DbInstanceIsCreating,omitempty" xml:"DbInstanceIsCreating,omitempty"`
	AccountName          *string                               `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	RdsInstance          []*string                             `json:"RdsInstance,omitempty" xml:"RdsInstance,omitempty" type:"Repeated"`
	RdsSuperAccount      []*CreateDrdsDBRequestRdsSuperAccount `json:"RdsSuperAccount,omitempty" xml:"RdsSuperAccount,omitempty" type:"Repeated"`
	InstDbName           []*CreateDrdsDBRequestInstDbName      `json:"InstDbName,omitempty" xml:"InstDbName,omitempty" type:"Repeated"`
}

func (s CreateDrdsDBRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDrdsDBRequest) GoString() string {
	return s.String()
}

func (s *CreateDrdsDBRequest) SetDrdsInstanceId(v string) *CreateDrdsDBRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *CreateDrdsDBRequest) SetDbName(v string) *CreateDrdsDBRequest {
	s.DbName = &v
	return s
}

func (s *CreateDrdsDBRequest) SetEncode(v string) *CreateDrdsDBRequest {
	s.Encode = &v
	return s
}

func (s *CreateDrdsDBRequest) SetPassword(v string) *CreateDrdsDBRequest {
	s.Password = &v
	return s
}

func (s *CreateDrdsDBRequest) SetType(v string) *CreateDrdsDBRequest {
	s.Type = &v
	return s
}

func (s *CreateDrdsDBRequest) SetDbInstType(v string) *CreateDrdsDBRequest {
	s.DbInstType = &v
	return s
}

func (s *CreateDrdsDBRequest) SetDbInstanceIsCreating(v bool) *CreateDrdsDBRequest {
	s.DbInstanceIsCreating = &v
	return s
}

func (s *CreateDrdsDBRequest) SetAccountName(v string) *CreateDrdsDBRequest {
	s.AccountName = &v
	return s
}

func (s *CreateDrdsDBRequest) SetRdsInstance(v []*string) *CreateDrdsDBRequest {
	s.RdsInstance = v
	return s
}

func (s *CreateDrdsDBRequest) SetRdsSuperAccount(v []*CreateDrdsDBRequestRdsSuperAccount) *CreateDrdsDBRequest {
	s.RdsSuperAccount = v
	return s
}

func (s *CreateDrdsDBRequest) SetInstDbName(v []*CreateDrdsDBRequestInstDbName) *CreateDrdsDBRequest {
	s.InstDbName = v
	return s
}

type CreateDrdsDBRequestRdsSuperAccount struct {
	Password     *string `json:"Password,omitempty" xml:"Password,omitempty"`
	DbInstanceId *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
	AccountName  *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
}

func (s CreateDrdsDBRequestRdsSuperAccount) String() string {
	return tea.Prettify(s)
}

func (s CreateDrdsDBRequestRdsSuperAccount) GoString() string {
	return s.String()
}

func (s *CreateDrdsDBRequestRdsSuperAccount) SetPassword(v string) *CreateDrdsDBRequestRdsSuperAccount {
	s.Password = &v
	return s
}

func (s *CreateDrdsDBRequestRdsSuperAccount) SetDbInstanceId(v string) *CreateDrdsDBRequestRdsSuperAccount {
	s.DbInstanceId = &v
	return s
}

func (s *CreateDrdsDBRequestRdsSuperAccount) SetAccountName(v string) *CreateDrdsDBRequestRdsSuperAccount {
	s.AccountName = &v
	return s
}

type CreateDrdsDBRequestInstDbName struct {
	DbInstanceId *string   `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
	ShardDbName  []*string `json:"ShardDbName,omitempty" xml:"ShardDbName,omitempty" type:"Repeated"`
}

func (s CreateDrdsDBRequestInstDbName) String() string {
	return tea.Prettify(s)
}

func (s CreateDrdsDBRequestInstDbName) GoString() string {
	return s.String()
}

func (s *CreateDrdsDBRequestInstDbName) SetDbInstanceId(v string) *CreateDrdsDBRequestInstDbName {
	s.DbInstanceId = &v
	return s
}

func (s *CreateDrdsDBRequestInstDbName) SetShardDbName(v []*string) *CreateDrdsDBRequestInstDbName {
	s.ShardDbName = v
	return s
}

type CreateDrdsDBResponseBody struct {
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDrdsDBResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDrdsDBResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDrdsDBResponseBody) SetSuccess(v bool) *CreateDrdsDBResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDrdsDBResponseBody) SetRequestId(v string) *CreateDrdsDBResponseBody {
	s.RequestId = &v
	return s
}

type CreateDrdsDBResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDrdsDBResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDrdsDBResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDrdsDBResponse) GoString() string {
	return s.String()
}

func (s *CreateDrdsDBResponse) SetHeaders(v map[string]*string) *CreateDrdsDBResponse {
	s.Headers = v
	return s
}

func (s *CreateDrdsDBResponse) SetBody(v *CreateDrdsDBResponseBody) *CreateDrdsDBResponse {
	s.Body = v
	return s
}

type CreateDrdsInstanceRequest struct {
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ZoneId          *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	Type            *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Quantity        *int32  `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	InstanceSeries  *string `json:"InstanceSeries,omitempty" xml:"InstanceSeries,omitempty"`
	Specification   *string `json:"Specification,omitempty" xml:"Specification,omitempty"`
	ClientToken     *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	PayType         *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	VpcId           *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VswitchId       *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	IsHa            *bool   `json:"isHa,omitempty" xml:"isHa,omitempty"`
	PricingCycle    *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	Duration        *int32  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	IsAutoRenew     *bool   `json:"IsAutoRenew,omitempty" xml:"IsAutoRenew,omitempty"`
	MasterInstId    *string `json:"MasterInstId,omitempty" xml:"MasterInstId,omitempty"`
	MySQLVersion    *int32  `json:"MySQLVersion,omitempty" xml:"MySQLVersion,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s CreateDrdsInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDrdsInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateDrdsInstanceRequest) SetDescription(v string) *CreateDrdsInstanceRequest {
	s.Description = &v
	return s
}

func (s *CreateDrdsInstanceRequest) SetRegionId(v string) *CreateDrdsInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDrdsInstanceRequest) SetZoneId(v string) *CreateDrdsInstanceRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateDrdsInstanceRequest) SetType(v string) *CreateDrdsInstanceRequest {
	s.Type = &v
	return s
}

func (s *CreateDrdsInstanceRequest) SetQuantity(v int32) *CreateDrdsInstanceRequest {
	s.Quantity = &v
	return s
}

func (s *CreateDrdsInstanceRequest) SetInstanceSeries(v string) *CreateDrdsInstanceRequest {
	s.InstanceSeries = &v
	return s
}

func (s *CreateDrdsInstanceRequest) SetSpecification(v string) *CreateDrdsInstanceRequest {
	s.Specification = &v
	return s
}

func (s *CreateDrdsInstanceRequest) SetClientToken(v string) *CreateDrdsInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDrdsInstanceRequest) SetPayType(v string) *CreateDrdsInstanceRequest {
	s.PayType = &v
	return s
}

func (s *CreateDrdsInstanceRequest) SetVpcId(v string) *CreateDrdsInstanceRequest {
	s.VpcId = &v
	return s
}

func (s *CreateDrdsInstanceRequest) SetVswitchId(v string) *CreateDrdsInstanceRequest {
	s.VswitchId = &v
	return s
}

func (s *CreateDrdsInstanceRequest) SetIsHa(v bool) *CreateDrdsInstanceRequest {
	s.IsHa = &v
	return s
}

func (s *CreateDrdsInstanceRequest) SetPricingCycle(v string) *CreateDrdsInstanceRequest {
	s.PricingCycle = &v
	return s
}

func (s *CreateDrdsInstanceRequest) SetDuration(v int32) *CreateDrdsInstanceRequest {
	s.Duration = &v
	return s
}

func (s *CreateDrdsInstanceRequest) SetIsAutoRenew(v bool) *CreateDrdsInstanceRequest {
	s.IsAutoRenew = &v
	return s
}

func (s *CreateDrdsInstanceRequest) SetMasterInstId(v string) *CreateDrdsInstanceRequest {
	s.MasterInstId = &v
	return s
}

func (s *CreateDrdsInstanceRequest) SetMySQLVersion(v int32) *CreateDrdsInstanceRequest {
	s.MySQLVersion = &v
	return s
}

func (s *CreateDrdsInstanceRequest) SetResourceGroupId(v string) *CreateDrdsInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

type CreateDrdsInstanceResponseBody struct {
	Success   *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *CreateDrdsInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s CreateDrdsInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDrdsInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDrdsInstanceResponseBody) SetSuccess(v bool) *CreateDrdsInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDrdsInstanceResponseBody) SetRequestId(v string) *CreateDrdsInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDrdsInstanceResponseBody) SetData(v *CreateDrdsInstanceResponseBodyData) *CreateDrdsInstanceResponseBody {
	s.Data = v
	return s
}

type CreateDrdsInstanceResponseBodyData struct {
	OrderId            *int64                                                `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	DrdsInstanceIdList *CreateDrdsInstanceResponseBodyDataDrdsInstanceIdList `json:"DrdsInstanceIdList,omitempty" xml:"DrdsInstanceIdList,omitempty" type:"Struct"`
}

func (s CreateDrdsInstanceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateDrdsInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateDrdsInstanceResponseBodyData) SetOrderId(v int64) *CreateDrdsInstanceResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *CreateDrdsInstanceResponseBodyData) SetDrdsInstanceIdList(v *CreateDrdsInstanceResponseBodyDataDrdsInstanceIdList) *CreateDrdsInstanceResponseBodyData {
	s.DrdsInstanceIdList = v
	return s
}

type CreateDrdsInstanceResponseBodyDataDrdsInstanceIdList struct {
	DrdsInstanceIdList []*string `json:"drdsInstanceIdList,omitempty" xml:"drdsInstanceIdList,omitempty" type:"Repeated"`
}

func (s CreateDrdsInstanceResponseBodyDataDrdsInstanceIdList) String() string {
	return tea.Prettify(s)
}

func (s CreateDrdsInstanceResponseBodyDataDrdsInstanceIdList) GoString() string {
	return s.String()
}

func (s *CreateDrdsInstanceResponseBodyDataDrdsInstanceIdList) SetDrdsInstanceIdList(v []*string) *CreateDrdsInstanceResponseBodyDataDrdsInstanceIdList {
	s.DrdsInstanceIdList = v
	return s
}

type CreateDrdsInstanceResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDrdsInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDrdsInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDrdsInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateDrdsInstanceResponse) SetHeaders(v map[string]*string) *CreateDrdsInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateDrdsInstanceResponse) SetBody(v *CreateDrdsInstanceResponseBody) *CreateDrdsInstanceResponse {
	s.Body = v
	return s
}

type CreateInstanceAccountRequest struct {
	DrdsInstanceId *string                                    `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	AccountName    *string                                    `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	Password       *string                                    `json:"Password,omitempty" xml:"Password,omitempty"`
	DbPrivilege    []*CreateInstanceAccountRequestDbPrivilege `json:"DbPrivilege,omitempty" xml:"DbPrivilege,omitempty" type:"Repeated"`
}

func (s CreateInstanceAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceAccountRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceAccountRequest) SetDrdsInstanceId(v string) *CreateInstanceAccountRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *CreateInstanceAccountRequest) SetAccountName(v string) *CreateInstanceAccountRequest {
	s.AccountName = &v
	return s
}

func (s *CreateInstanceAccountRequest) SetPassword(v string) *CreateInstanceAccountRequest {
	s.Password = &v
	return s
}

func (s *CreateInstanceAccountRequest) SetDbPrivilege(v []*CreateInstanceAccountRequestDbPrivilege) *CreateInstanceAccountRequest {
	s.DbPrivilege = v
	return s
}

type CreateInstanceAccountRequestDbPrivilege struct {
	DbName    *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	Privilege *string `json:"Privilege,omitempty" xml:"Privilege,omitempty"`
}

func (s CreateInstanceAccountRequestDbPrivilege) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceAccountRequestDbPrivilege) GoString() string {
	return s.String()
}

func (s *CreateInstanceAccountRequestDbPrivilege) SetDbName(v string) *CreateInstanceAccountRequestDbPrivilege {
	s.DbName = &v
	return s
}

func (s *CreateInstanceAccountRequestDbPrivilege) SetPrivilege(v string) *CreateInstanceAccountRequestDbPrivilege {
	s.Privilege = &v
	return s
}

type CreateInstanceAccountResponseBody struct {
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateInstanceAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceAccountResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceAccountResponseBody) SetSuccess(v bool) *CreateInstanceAccountResponseBody {
	s.Success = &v
	return s
}

func (s *CreateInstanceAccountResponseBody) SetRequestId(v string) *CreateInstanceAccountResponseBody {
	s.RequestId = &v
	return s
}

type CreateInstanceAccountResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateInstanceAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateInstanceAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceAccountResponse) GoString() string {
	return s.String()
}

func (s *CreateInstanceAccountResponse) SetHeaders(v map[string]*string) *CreateInstanceAccountResponse {
	s.Headers = v
	return s
}

func (s *CreateInstanceAccountResponse) SetBody(v *CreateInstanceAccountResponseBody) *CreateInstanceAccountResponse {
	s.Body = v
	return s
}

type CreateInstanceInternetAddressRequest struct {
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateInstanceInternetAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceInternetAddressRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceInternetAddressRequest) SetDrdsInstanceId(v string) *CreateInstanceInternetAddressRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *CreateInstanceInternetAddressRequest) SetRegionId(v string) *CreateInstanceInternetAddressRequest {
	s.RegionId = &v
	return s
}

type CreateInstanceInternetAddressResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateInstanceInternetAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceInternetAddressResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceInternetAddressResponseBody) SetCode(v int32) *CreateInstanceInternetAddressResponseBody {
	s.Code = &v
	return s
}

func (s *CreateInstanceInternetAddressResponseBody) SetData(v bool) *CreateInstanceInternetAddressResponseBody {
	s.Data = &v
	return s
}

func (s *CreateInstanceInternetAddressResponseBody) SetRequestId(v string) *CreateInstanceInternetAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInstanceInternetAddressResponseBody) SetSuccess(v bool) *CreateInstanceInternetAddressResponseBody {
	s.Success = &v
	return s
}

type CreateInstanceInternetAddressResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateInstanceInternetAddressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateInstanceInternetAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceInternetAddressResponse) GoString() string {
	return s.String()
}

func (s *CreateInstanceInternetAddressResponse) SetHeaders(v map[string]*string) *CreateInstanceInternetAddressResponse {
	s.Headers = v
	return s
}

func (s *CreateInstanceInternetAddressResponse) SetBody(v *CreateInstanceInternetAddressResponseBody) *CreateInstanceInternetAddressResponse {
	s.Body = v
	return s
}

type CreateOrderForRdsRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Params   *string `json:"Params,omitempty" xml:"Params,omitempty"`
}

func (s CreateOrderForRdsRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateOrderForRdsRequest) GoString() string {
	return s.String()
}

func (s *CreateOrderForRdsRequest) SetRegionId(v string) *CreateOrderForRdsRequest {
	s.RegionId = &v
	return s
}

func (s *CreateOrderForRdsRequest) SetParams(v string) *CreateOrderForRdsRequest {
	s.Params = &v
	return s
}

type CreateOrderForRdsResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateOrderForRdsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateOrderForRdsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOrderForRdsResponseBody) SetData(v string) *CreateOrderForRdsResponseBody {
	s.Data = &v
	return s
}

func (s *CreateOrderForRdsResponseBody) SetRequestId(v string) *CreateOrderForRdsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOrderForRdsResponseBody) SetSuccess(v bool) *CreateOrderForRdsResponseBody {
	s.Success = &v
	return s
}

type CreateOrderForRdsResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateOrderForRdsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateOrderForRdsResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateOrderForRdsResponse) GoString() string {
	return s.String()
}

func (s *CreateOrderForRdsResponse) SetHeaders(v map[string]*string) *CreateOrderForRdsResponse {
	s.Headers = v
	return s
}

func (s *CreateOrderForRdsResponse) SetBody(v *CreateOrderForRdsResponseBody) *CreateOrderForRdsResponse {
	s.Body = v
	return s
}

type CreateShardTaskRequest struct {
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DrdsInstanceId  *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	DbName          *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	SourceTableName *string `json:"SourceTableName,omitempty" xml:"SourceTableName,omitempty"`
	TargetTableName *string `json:"TargetTableName,omitempty" xml:"TargetTableName,omitempty"`
	TaskType        *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s CreateShardTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateShardTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateShardTaskRequest) SetRegionId(v string) *CreateShardTaskRequest {
	s.RegionId = &v
	return s
}

func (s *CreateShardTaskRequest) SetDrdsInstanceId(v string) *CreateShardTaskRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *CreateShardTaskRequest) SetDbName(v string) *CreateShardTaskRequest {
	s.DbName = &v
	return s
}

func (s *CreateShardTaskRequest) SetSourceTableName(v string) *CreateShardTaskRequest {
	s.SourceTableName = &v
	return s
}

func (s *CreateShardTaskRequest) SetTargetTableName(v string) *CreateShardTaskRequest {
	s.TargetTableName = &v
	return s
}

func (s *CreateShardTaskRequest) SetTaskType(v string) *CreateShardTaskRequest {
	s.TaskType = &v
	return s
}

type CreateShardTaskResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateShardTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateShardTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateShardTaskResponseBody) SetData(v bool) *CreateShardTaskResponseBody {
	s.Data = &v
	return s
}

func (s *CreateShardTaskResponseBody) SetRequestId(v string) *CreateShardTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateShardTaskResponseBody) SetSuccess(v bool) *CreateShardTaskResponseBody {
	s.Success = &v
	return s
}

type CreateShardTaskResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateShardTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateShardTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateShardTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateShardTaskResponse) SetHeaders(v map[string]*string) *CreateShardTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateShardTaskResponse) SetBody(v *CreateShardTaskResponseBody) *CreateShardTaskResponse {
	s.Body = v
	return s
}

type DescribeBackMenuRequest struct {
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s DescribeBackMenuRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackMenuRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackMenuRequest) SetDrdsInstanceId(v string) *DescribeBackMenuRequest {
	s.DrdsInstanceId = &v
	return s
}

type DescribeBackMenuResponseBody struct {
	Success   *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	List      *DescribeBackMenuResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Struct"`
}

func (s DescribeBackMenuResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackMenuResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackMenuResponseBody) SetSuccess(v bool) *DescribeBackMenuResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeBackMenuResponseBody) SetRequestId(v string) *DescribeBackMenuResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackMenuResponseBody) SetList(v *DescribeBackMenuResponseBodyList) *DescribeBackMenuResponseBody {
	s.List = v
	return s
}

type DescribeBackMenuResponseBodyList struct {
	List []*DescribeBackMenuResponseBodyListList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
}

func (s DescribeBackMenuResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackMenuResponseBodyList) GoString() string {
	return s.String()
}

func (s *DescribeBackMenuResponseBodyList) SetList(v []*DescribeBackMenuResponseBodyListList) *DescribeBackMenuResponseBodyList {
	s.List = v
	return s
}

type DescribeBackMenuResponseBodyListList struct {
	Support  *bool   `json:"Support,omitempty" xml:"Support,omitempty"`
	MenuName *string `json:"MenuName,omitempty" xml:"MenuName,omitempty"`
}

func (s DescribeBackMenuResponseBodyListList) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackMenuResponseBodyListList) GoString() string {
	return s.String()
}

func (s *DescribeBackMenuResponseBodyListList) SetSupport(v bool) *DescribeBackMenuResponseBodyListList {
	s.Support = &v
	return s
}

func (s *DescribeBackMenuResponseBodyListList) SetMenuName(v string) *DescribeBackMenuResponseBodyListList {
	s.MenuName = &v
	return s
}

type DescribeBackMenuResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeBackMenuResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeBackMenuResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackMenuResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackMenuResponse) SetHeaders(v map[string]*string) *DescribeBackMenuResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackMenuResponse) SetBody(v *DescribeBackMenuResponseBody) *DescribeBackMenuResponse {
	s.Body = v
	return s
}

type DescribeBackupDbsRequest struct {
	DrdsInstanceId       *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	PreferredRestoreTime *string `json:"PreferredRestoreTime,omitempty" xml:"PreferredRestoreTime,omitempty"`
	BackupId             *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
}

func (s DescribeBackupDbsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupDbsRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupDbsRequest) SetDrdsInstanceId(v string) *DescribeBackupDbsRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeBackupDbsRequest) SetPreferredRestoreTime(v string) *DescribeBackupDbsRequest {
	s.PreferredRestoreTime = &v
	return s
}

func (s *DescribeBackupDbsRequest) SetBackupId(v string) *DescribeBackupDbsRequest {
	s.BackupId = &v
	return s
}

type DescribeBackupDbsResponseBody struct {
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DbNames   *DescribeBackupDbsResponseBodyDbNames `json:"DbNames,omitempty" xml:"DbNames,omitempty" type:"Struct"`
}

func (s DescribeBackupDbsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupDbsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupDbsResponseBody) SetSuccess(v bool) *DescribeBackupDbsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeBackupDbsResponseBody) SetRequestId(v string) *DescribeBackupDbsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupDbsResponseBody) SetDbNames(v *DescribeBackupDbsResponseBodyDbNames) *DescribeBackupDbsResponseBody {
	s.DbNames = v
	return s
}

type DescribeBackupDbsResponseBodyDbNames struct {
	DbName []*string `json:"dbName,omitempty" xml:"dbName,omitempty" type:"Repeated"`
}

func (s DescribeBackupDbsResponseBodyDbNames) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupDbsResponseBodyDbNames) GoString() string {
	return s.String()
}

func (s *DescribeBackupDbsResponseBodyDbNames) SetDbName(v []*string) *DescribeBackupDbsResponseBodyDbNames {
	s.DbName = v
	return s
}

type DescribeBackupDbsResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeBackupDbsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeBackupDbsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupDbsResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupDbsResponse) SetHeaders(v map[string]*string) *DescribeBackupDbsResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackupDbsResponse) SetBody(v *DescribeBackupDbsResponseBody) *DescribeBackupDbsResponse {
	s.Body = v
	return s
}

type DescribeBackupLocalRequest struct {
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s DescribeBackupLocalRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupLocalRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupLocalRequest) SetDrdsInstanceId(v string) *DescribeBackupLocalRequest {
	s.DrdsInstanceId = &v
	return s
}

type DescribeBackupLocalResponseBody struct {
	Success        *bool                                          `json:"Success,omitempty" xml:"Success,omitempty"`
	RequestId      *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	BackupPolicyDO *DescribeBackupLocalResponseBodyBackupPolicyDO `json:"BackupPolicyDO,omitempty" xml:"BackupPolicyDO,omitempty" type:"Struct"`
}

func (s DescribeBackupLocalResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupLocalResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupLocalResponseBody) SetSuccess(v bool) *DescribeBackupLocalResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeBackupLocalResponseBody) SetRequestId(v string) *DescribeBackupLocalResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupLocalResponseBody) SetBackupPolicyDO(v *DescribeBackupLocalResponseBodyBackupPolicyDO) *DescribeBackupLocalResponseBody {
	s.BackupPolicyDO = v
	return s
}

type DescribeBackupLocalResponseBodyBackupPolicyDO struct {
	BackupDbName              *string `json:"BackupDbName,omitempty" xml:"BackupDbName,omitempty"`
	LogBackupRetentionPeriod  *int64  `json:"LogBackupRetentionPeriod,omitempty" xml:"LogBackupRetentionPeriod,omitempty"`
	DataBackupRetentionPeriod *int64  `json:"DataBackupRetentionPeriod,omitempty" xml:"DataBackupRetentionPeriod,omitempty"`
	BackupType                *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	BackupLevel               *string `json:"BackupLevel,omitempty" xml:"BackupLevel,omitempty"`
	LocalLogRetentionHours    *int64  `json:"LocalLogRetentionHours,omitempty" xml:"LocalLogRetentionHours,omitempty"`
	HighSpaceUsageProtection  *int64  `json:"HighSpaceUsageProtection,omitempty" xml:"HighSpaceUsageProtection,omitempty"`
	GmtModified               *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	BackupRetentionPeriod     *int64  `json:"BackupRetentionPeriod,omitempty" xml:"BackupRetentionPeriod,omitempty"`
	BackupPolicyMode          *string `json:"BackupPolicyMode,omitempty" xml:"BackupPolicyMode,omitempty"`
	PreferredBackupPeriod     *string `json:"PreferredBackupPeriod,omitempty" xml:"PreferredBackupPeriod,omitempty"`
	LocalLogRetentionSpace    *int64  `json:"LocalLogRetentionSpace,omitempty" xml:"LocalLogRetentionSpace,omitempty"`
	BackupAppName             *string `json:"BackupAppName,omitempty" xml:"BackupAppName,omitempty"`
	PreferredBackupTime       *string `json:"PreferredBackupTime,omitempty" xml:"PreferredBackupTime,omitempty"`
	GmtCreate                 *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	BackupMode                *string `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	BackupLog                 *string `json:"BackupLog,omitempty" xml:"BackupLog,omitempty"`
	NextBackupActuallyTime    *string `json:"NextBackupActuallyTime,omitempty" xml:"NextBackupActuallyTime,omitempty"`
}

func (s DescribeBackupLocalResponseBodyBackupPolicyDO) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupLocalResponseBodyBackupPolicyDO) GoString() string {
	return s.String()
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) SetBackupDbName(v string) *DescribeBackupLocalResponseBodyBackupPolicyDO {
	s.BackupDbName = &v
	return s
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) SetLogBackupRetentionPeriod(v int64) *DescribeBackupLocalResponseBodyBackupPolicyDO {
	s.LogBackupRetentionPeriod = &v
	return s
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) SetDataBackupRetentionPeriod(v int64) *DescribeBackupLocalResponseBodyBackupPolicyDO {
	s.DataBackupRetentionPeriod = &v
	return s
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) SetBackupType(v string) *DescribeBackupLocalResponseBodyBackupPolicyDO {
	s.BackupType = &v
	return s
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) SetBackupLevel(v string) *DescribeBackupLocalResponseBodyBackupPolicyDO {
	s.BackupLevel = &v
	return s
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) SetLocalLogRetentionHours(v int64) *DescribeBackupLocalResponseBodyBackupPolicyDO {
	s.LocalLogRetentionHours = &v
	return s
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) SetHighSpaceUsageProtection(v int64) *DescribeBackupLocalResponseBodyBackupPolicyDO {
	s.HighSpaceUsageProtection = &v
	return s
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) SetGmtModified(v int64) *DescribeBackupLocalResponseBodyBackupPolicyDO {
	s.GmtModified = &v
	return s
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) SetBackupRetentionPeriod(v int64) *DescribeBackupLocalResponseBodyBackupPolicyDO {
	s.BackupRetentionPeriod = &v
	return s
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) SetBackupPolicyMode(v string) *DescribeBackupLocalResponseBodyBackupPolicyDO {
	s.BackupPolicyMode = &v
	return s
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) SetPreferredBackupPeriod(v string) *DescribeBackupLocalResponseBodyBackupPolicyDO {
	s.PreferredBackupPeriod = &v
	return s
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) SetLocalLogRetentionSpace(v int64) *DescribeBackupLocalResponseBodyBackupPolicyDO {
	s.LocalLogRetentionSpace = &v
	return s
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) SetBackupAppName(v string) *DescribeBackupLocalResponseBodyBackupPolicyDO {
	s.BackupAppName = &v
	return s
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) SetPreferredBackupTime(v string) *DescribeBackupLocalResponseBodyBackupPolicyDO {
	s.PreferredBackupTime = &v
	return s
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) SetGmtCreate(v int64) *DescribeBackupLocalResponseBodyBackupPolicyDO {
	s.GmtCreate = &v
	return s
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) SetBackupMode(v string) *DescribeBackupLocalResponseBodyBackupPolicyDO {
	s.BackupMode = &v
	return s
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) SetBackupLog(v string) *DescribeBackupLocalResponseBodyBackupPolicyDO {
	s.BackupLog = &v
	return s
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) SetNextBackupActuallyTime(v string) *DescribeBackupLocalResponseBodyBackupPolicyDO {
	s.NextBackupActuallyTime = &v
	return s
}

type DescribeBackupLocalResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeBackupLocalResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeBackupLocalResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupLocalResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupLocalResponse) SetHeaders(v map[string]*string) *DescribeBackupLocalResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackupLocalResponse) SetBody(v *DescribeBackupLocalResponseBody) *DescribeBackupLocalResponse {
	s.Body = v
	return s
}

type DescribeBackupPolicyRequest struct {
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s DescribeBackupPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyRequest) SetDrdsInstanceId(v string) *DescribeBackupPolicyRequest {
	s.DrdsInstanceId = &v
	return s
}

type DescribeBackupPolicyResponseBody struct {
	Success        *bool                                           `json:"Success,omitempty" xml:"Success,omitempty"`
	RequestId      *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	BackupPolicyDO *DescribeBackupPolicyResponseBodyBackupPolicyDO `json:"BackupPolicyDO,omitempty" xml:"BackupPolicyDO,omitempty" type:"Struct"`
}

func (s DescribeBackupPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyResponseBody) SetSuccess(v bool) *DescribeBackupPolicyResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetRequestId(v string) *DescribeBackupPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetBackupPolicyDO(v *DescribeBackupPolicyResponseBodyBackupPolicyDO) *DescribeBackupPolicyResponseBody {
	s.BackupPolicyDO = v
	return s
}

type DescribeBackupPolicyResponseBodyBackupPolicyDO struct {
	BackupDbName              *string `json:"BackupDbName,omitempty" xml:"BackupDbName,omitempty"`
	LogBackupRetentionPeriod  *int64  `json:"LogBackupRetentionPeriod,omitempty" xml:"LogBackupRetentionPeriod,omitempty"`
	DataBackupRetentionPeriod *int64  `json:"DataBackupRetentionPeriod,omitempty" xml:"DataBackupRetentionPeriod,omitempty"`
	BackupType                *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	BackupLevel               *string `json:"BackupLevel,omitempty" xml:"BackupLevel,omitempty"`
	LocalLogRetentionHours    *int64  `json:"LocalLogRetentionHours,omitempty" xml:"LocalLogRetentionHours,omitempty"`
	HighSpaceUsageProtection  *int64  `json:"HighSpaceUsageProtection,omitempty" xml:"HighSpaceUsageProtection,omitempty"`
	GmtModified               *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	BackupRetentionPeriod     *int64  `json:"BackupRetentionPeriod,omitempty" xml:"BackupRetentionPeriod,omitempty"`
	BackupPolicyMode          *string `json:"BackupPolicyMode,omitempty" xml:"BackupPolicyMode,omitempty"`
	PreferredBackupPeriod     *string `json:"PreferredBackupPeriod,omitempty" xml:"PreferredBackupPeriod,omitempty"`
	LocalLogRetentionSpace    *int64  `json:"LocalLogRetentionSpace,omitempty" xml:"LocalLogRetentionSpace,omitempty"`
	BackupAppName             *string `json:"BackupAppName,omitempty" xml:"BackupAppName,omitempty"`
	PreferredBackupTime       *string `json:"PreferredBackupTime,omitempty" xml:"PreferredBackupTime,omitempty"`
	GmtCreate                 *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	BackupMode                *string `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	BackupLog                 *string `json:"BackupLog,omitempty" xml:"BackupLog,omitempty"`
	NextBackupActuallyTime    *string `json:"NextBackupActuallyTime,omitempty" xml:"NextBackupActuallyTime,omitempty"`
}

func (s DescribeBackupPolicyResponseBodyBackupPolicyDO) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPolicyResponseBodyBackupPolicyDO) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) SetBackupDbName(v string) *DescribeBackupPolicyResponseBodyBackupPolicyDO {
	s.BackupDbName = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) SetLogBackupRetentionPeriod(v int64) *DescribeBackupPolicyResponseBodyBackupPolicyDO {
	s.LogBackupRetentionPeriod = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) SetDataBackupRetentionPeriod(v int64) *DescribeBackupPolicyResponseBodyBackupPolicyDO {
	s.DataBackupRetentionPeriod = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) SetBackupType(v string) *DescribeBackupPolicyResponseBodyBackupPolicyDO {
	s.BackupType = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) SetBackupLevel(v string) *DescribeBackupPolicyResponseBodyBackupPolicyDO {
	s.BackupLevel = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) SetLocalLogRetentionHours(v int64) *DescribeBackupPolicyResponseBodyBackupPolicyDO {
	s.LocalLogRetentionHours = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) SetHighSpaceUsageProtection(v int64) *DescribeBackupPolicyResponseBodyBackupPolicyDO {
	s.HighSpaceUsageProtection = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) SetGmtModified(v int64) *DescribeBackupPolicyResponseBodyBackupPolicyDO {
	s.GmtModified = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) SetBackupRetentionPeriod(v int64) *DescribeBackupPolicyResponseBodyBackupPolicyDO {
	s.BackupRetentionPeriod = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) SetBackupPolicyMode(v string) *DescribeBackupPolicyResponseBodyBackupPolicyDO {
	s.BackupPolicyMode = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) SetPreferredBackupPeriod(v string) *DescribeBackupPolicyResponseBodyBackupPolicyDO {
	s.PreferredBackupPeriod = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) SetLocalLogRetentionSpace(v int64) *DescribeBackupPolicyResponseBodyBackupPolicyDO {
	s.LocalLogRetentionSpace = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) SetBackupAppName(v string) *DescribeBackupPolicyResponseBodyBackupPolicyDO {
	s.BackupAppName = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) SetPreferredBackupTime(v string) *DescribeBackupPolicyResponseBodyBackupPolicyDO {
	s.PreferredBackupTime = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) SetGmtCreate(v int64) *DescribeBackupPolicyResponseBodyBackupPolicyDO {
	s.GmtCreate = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) SetBackupMode(v string) *DescribeBackupPolicyResponseBodyBackupPolicyDO {
	s.BackupMode = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) SetBackupLog(v string) *DescribeBackupPolicyResponseBodyBackupPolicyDO {
	s.BackupLog = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) SetNextBackupActuallyTime(v string) *DescribeBackupPolicyResponseBodyBackupPolicyDO {
	s.NextBackupActuallyTime = &v
	return s
}

type DescribeBackupPolicyResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeBackupPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeBackupPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPolicyResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyResponse) SetHeaders(v map[string]*string) *DescribeBackupPolicyResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackupPolicyResponse) SetBody(v *DescribeBackupPolicyResponseBody) *DescribeBackupPolicyResponse {
	s.Body = v
	return s
}

type DescribeBackupSetsRequest struct {
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s DescribeBackupSetsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupSetsRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupSetsRequest) SetDrdsInstanceId(v string) *DescribeBackupSetsRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeBackupSetsRequest) SetStartTime(v string) *DescribeBackupSetsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeBackupSetsRequest) SetEndTime(v string) *DescribeBackupSetsRequest {
	s.EndTime = &v
	return s
}

type DescribeBackupSetsResponseBody struct {
	Success    *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
	RequestId  *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	BackupSets *DescribeBackupSetsResponseBodyBackupSets `json:"BackupSets,omitempty" xml:"BackupSets,omitempty" type:"Struct"`
}

func (s DescribeBackupSetsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupSetsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupSetsResponseBody) SetSuccess(v bool) *DescribeBackupSetsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeBackupSetsResponseBody) SetRequestId(v string) *DescribeBackupSetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupSetsResponseBody) SetBackupSets(v *DescribeBackupSetsResponseBodyBackupSets) *DescribeBackupSetsResponseBody {
	s.BackupSets = v
	return s
}

type DescribeBackupSetsResponseBodyBackupSets struct {
	BackupSet []*DescribeBackupSetsResponseBodyBackupSetsBackupSet `json:"backupSet,omitempty" xml:"backupSet,omitempty" type:"Repeated"`
}

func (s DescribeBackupSetsResponseBodyBackupSets) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupSetsResponseBodyBackupSets) GoString() string {
	return s.String()
}

func (s *DescribeBackupSetsResponseBodyBackupSets) SetBackupSet(v []*DescribeBackupSetsResponseBodyBackupSetsBackupSet) *DescribeBackupSetsResponseBodyBackupSets {
	s.BackupSet = v
	return s
}

type DescribeBackupSetsResponseBodyBackupSetsBackupSet struct {
	EnableRecovery      *bool                                                       `json:"EnableRecovery,omitempty" xml:"EnableRecovery,omitempty"`
	Status              *int64                                                      `json:"Status,omitempty" xml:"Status,omitempty"`
	BackupConsitentTime *string                                                     `json:"BackupConsitentTime,omitempty" xml:"BackupConsitentTime,omitempty"`
	BackupType          *string                                                     `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	BackupStartTime     *int64                                                      `json:"BackupStartTime,omitempty" xml:"BackupStartTime,omitempty"`
	BackupLevel         *string                                                     `json:"BackupLevel,omitempty" xml:"BackupLevel,omitempty"`
	BackupMode          *string                                                     `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	BackupEndTime       *int64                                                      `json:"BackupEndTime,omitempty" xml:"BackupEndTime,omitempty"`
	Id                  *string                                                     `json:"Id,omitempty" xml:"Id,omitempty"`
	BackupTotalSize     *string                                                     `json:"BackupTotalSize,omitempty" xml:"BackupTotalSize,omitempty"`
	BackupDbs           *DescribeBackupSetsResponseBodyBackupSetsBackupSetBackupDbs `json:"BackupDbs,omitempty" xml:"BackupDbs,omitempty" type:"Struct"`
}

func (s DescribeBackupSetsResponseBodyBackupSetsBackupSet) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupSetsResponseBodyBackupSetsBackupSet) GoString() string {
	return s.String()
}

func (s *DescribeBackupSetsResponseBodyBackupSetsBackupSet) SetEnableRecovery(v bool) *DescribeBackupSetsResponseBodyBackupSetsBackupSet {
	s.EnableRecovery = &v
	return s
}

func (s *DescribeBackupSetsResponseBodyBackupSetsBackupSet) SetStatus(v int64) *DescribeBackupSetsResponseBodyBackupSetsBackupSet {
	s.Status = &v
	return s
}

func (s *DescribeBackupSetsResponseBodyBackupSetsBackupSet) SetBackupConsitentTime(v string) *DescribeBackupSetsResponseBodyBackupSetsBackupSet {
	s.BackupConsitentTime = &v
	return s
}

func (s *DescribeBackupSetsResponseBodyBackupSetsBackupSet) SetBackupType(v string) *DescribeBackupSetsResponseBodyBackupSetsBackupSet {
	s.BackupType = &v
	return s
}

func (s *DescribeBackupSetsResponseBodyBackupSetsBackupSet) SetBackupStartTime(v int64) *DescribeBackupSetsResponseBodyBackupSetsBackupSet {
	s.BackupStartTime = &v
	return s
}

func (s *DescribeBackupSetsResponseBodyBackupSetsBackupSet) SetBackupLevel(v string) *DescribeBackupSetsResponseBodyBackupSetsBackupSet {
	s.BackupLevel = &v
	return s
}

func (s *DescribeBackupSetsResponseBodyBackupSetsBackupSet) SetBackupMode(v string) *DescribeBackupSetsResponseBodyBackupSetsBackupSet {
	s.BackupMode = &v
	return s
}

func (s *DescribeBackupSetsResponseBodyBackupSetsBackupSet) SetBackupEndTime(v int64) *DescribeBackupSetsResponseBodyBackupSetsBackupSet {
	s.BackupEndTime = &v
	return s
}

func (s *DescribeBackupSetsResponseBodyBackupSetsBackupSet) SetId(v string) *DescribeBackupSetsResponseBodyBackupSetsBackupSet {
	s.Id = &v
	return s
}

func (s *DescribeBackupSetsResponseBodyBackupSetsBackupSet) SetBackupTotalSize(v string) *DescribeBackupSetsResponseBodyBackupSetsBackupSet {
	s.BackupTotalSize = &v
	return s
}

func (s *DescribeBackupSetsResponseBodyBackupSetsBackupSet) SetBackupDbs(v *DescribeBackupSetsResponseBodyBackupSetsBackupSetBackupDbs) *DescribeBackupSetsResponseBodyBackupSetsBackupSet {
	s.BackupDbs = v
	return s
}

type DescribeBackupSetsResponseBodyBackupSetsBackupSetBackupDbs struct {
	BackupDb []*string `json:"backupDb,omitempty" xml:"backupDb,omitempty" type:"Repeated"`
}

func (s DescribeBackupSetsResponseBodyBackupSetsBackupSetBackupDbs) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupSetsResponseBodyBackupSetsBackupSetBackupDbs) GoString() string {
	return s.String()
}

func (s *DescribeBackupSetsResponseBodyBackupSetsBackupSetBackupDbs) SetBackupDb(v []*string) *DescribeBackupSetsResponseBodyBackupSetsBackupSetBackupDbs {
	s.BackupDb = v
	return s
}

type DescribeBackupSetsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeBackupSetsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeBackupSetsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupSetsResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupSetsResponse) SetHeaders(v map[string]*string) *DescribeBackupSetsResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackupSetsResponse) SetBody(v *DescribeBackupSetsResponseBody) *DescribeBackupSetsResponse {
	s.Body = v
	return s
}

type DescribeBackupTimesRequest struct {
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s DescribeBackupTimesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupTimesRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupTimesRequest) SetDrdsInstanceId(v string) *DescribeBackupTimesRequest {
	s.DrdsInstanceId = &v
	return s
}

type DescribeBackupTimesResponseBody struct {
	Success     *bool                                       `json:"Success,omitempty" xml:"Success,omitempty"`
	RequestId   *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RestoreTime *DescribeBackupTimesResponseBodyRestoreTime `json:"RestoreTime,omitempty" xml:"RestoreTime,omitempty" type:"Struct"`
}

func (s DescribeBackupTimesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupTimesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupTimesResponseBody) SetSuccess(v bool) *DescribeBackupTimesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeBackupTimesResponseBody) SetRequestId(v string) *DescribeBackupTimesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupTimesResponseBody) SetRestoreTime(v *DescribeBackupTimesResponseBodyRestoreTime) *DescribeBackupTimesResponseBody {
	s.RestoreTime = v
	return s
}

type DescribeBackupTimesResponseBodyRestoreTime struct {
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeBackupTimesResponseBodyRestoreTime) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupTimesResponseBodyRestoreTime) GoString() string {
	return s.String()
}

func (s *DescribeBackupTimesResponseBodyRestoreTime) SetEndTime(v string) *DescribeBackupTimesResponseBodyRestoreTime {
	s.EndTime = &v
	return s
}

func (s *DescribeBackupTimesResponseBodyRestoreTime) SetStartTime(v string) *DescribeBackupTimesResponseBodyRestoreTime {
	s.StartTime = &v
	return s
}

type DescribeBackupTimesResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeBackupTimesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeBackupTimesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupTimesResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupTimesResponse) SetHeaders(v map[string]*string) *DescribeBackupTimesResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackupTimesResponse) SetBody(v *DescribeBackupTimesResponseBody) *DescribeBackupTimesResponse {
	s.Body = v
	return s
}

type DescribeBroadcastTablesRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	Query          *string `json:"Query,omitempty" xml:"Query,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentPage    *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
}

func (s DescribeBroadcastTablesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBroadcastTablesRequest) GoString() string {
	return s.String()
}

func (s *DescribeBroadcastTablesRequest) SetRegionId(v string) *DescribeBroadcastTablesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeBroadcastTablesRequest) SetDrdsInstanceId(v string) *DescribeBroadcastTablesRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeBroadcastTablesRequest) SetDbName(v string) *DescribeBroadcastTablesRequest {
	s.DbName = &v
	return s
}

func (s *DescribeBroadcastTablesRequest) SetQuery(v string) *DescribeBroadcastTablesRequest {
	s.Query = &v
	return s
}

func (s *DescribeBroadcastTablesRequest) SetPageSize(v int32) *DescribeBroadcastTablesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeBroadcastTablesRequest) SetCurrentPage(v int32) *DescribeBroadcastTablesRequest {
	s.CurrentPage = &v
	return s
}

type DescribeBroadcastTablesResponseBody struct {
	RequestId  *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
	IsShard    *bool                                      `json:"IsShard,omitempty" xml:"IsShard,omitempty"`
	PageNumber *int32                                     `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total      *int32                                     `json:"Total,omitempty" xml:"Total,omitempty"`
	List       []*DescribeBroadcastTablesResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s DescribeBroadcastTablesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBroadcastTablesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBroadcastTablesResponseBody) SetRequestId(v string) *DescribeBroadcastTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBroadcastTablesResponseBody) SetSuccess(v bool) *DescribeBroadcastTablesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeBroadcastTablesResponseBody) SetIsShard(v bool) *DescribeBroadcastTablesResponseBody {
	s.IsShard = &v
	return s
}

func (s *DescribeBroadcastTablesResponseBody) SetPageNumber(v int32) *DescribeBroadcastTablesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeBroadcastTablesResponseBody) SetPageSize(v int32) *DescribeBroadcastTablesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeBroadcastTablesResponseBody) SetTotal(v int32) *DescribeBroadcastTablesResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeBroadcastTablesResponseBody) SetList(v []*DescribeBroadcastTablesResponseBodyList) *DescribeBroadcastTablesResponseBody {
	s.List = v
	return s
}

type DescribeBroadcastTablesResponseBodyList struct {
	Status        *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	IsShard       *bool   `json:"IsShard,omitempty" xml:"IsShard,omitempty"`
	Broadcast     *bool   `json:"Broadcast,omitempty" xml:"Broadcast,omitempty"`
	Table         *string `json:"Table,omitempty" xml:"Table,omitempty"`
	DbInstType    *int32  `json:"DbInstType,omitempty" xml:"DbInstType,omitempty"`
	BroadcastType *string `json:"BroadcastType,omitempty" xml:"BroadcastType,omitempty"`
}

func (s DescribeBroadcastTablesResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s DescribeBroadcastTablesResponseBodyList) GoString() string {
	return s.String()
}

func (s *DescribeBroadcastTablesResponseBodyList) SetStatus(v int32) *DescribeBroadcastTablesResponseBodyList {
	s.Status = &v
	return s
}

func (s *DescribeBroadcastTablesResponseBodyList) SetIsShard(v bool) *DescribeBroadcastTablesResponseBodyList {
	s.IsShard = &v
	return s
}

func (s *DescribeBroadcastTablesResponseBodyList) SetBroadcast(v bool) *DescribeBroadcastTablesResponseBodyList {
	s.Broadcast = &v
	return s
}

func (s *DescribeBroadcastTablesResponseBodyList) SetTable(v string) *DescribeBroadcastTablesResponseBodyList {
	s.Table = &v
	return s
}

func (s *DescribeBroadcastTablesResponseBodyList) SetDbInstType(v int32) *DescribeBroadcastTablesResponseBodyList {
	s.DbInstType = &v
	return s
}

func (s *DescribeBroadcastTablesResponseBodyList) SetBroadcastType(v string) *DescribeBroadcastTablesResponseBodyList {
	s.BroadcastType = &v
	return s
}

type DescribeBroadcastTablesResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeBroadcastTablesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeBroadcastTablesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBroadcastTablesResponse) GoString() string {
	return s.String()
}

func (s *DescribeBroadcastTablesResponse) SetHeaders(v map[string]*string) *DescribeBroadcastTablesResponse {
	s.Headers = v
	return s
}

func (s *DescribeBroadcastTablesResponse) SetBody(v *DescribeBroadcastTablesResponseBody) *DescribeBroadcastTablesResponse {
	s.Body = v
	return s
}

type DescribeDbInstanceDbsRequest struct {
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	DbInstanceId   *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
	AccountName    *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	Password       *string `json:"Password,omitempty" xml:"Password,omitempty"`
	DbInstType     *string `json:"DbInstType,omitempty" xml:"DbInstType,omitempty"`
}

func (s DescribeDbInstanceDbsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDbInstanceDbsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDbInstanceDbsRequest) SetDrdsInstanceId(v string) *DescribeDbInstanceDbsRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeDbInstanceDbsRequest) SetDbInstanceId(v string) *DescribeDbInstanceDbsRequest {
	s.DbInstanceId = &v
	return s
}

func (s *DescribeDbInstanceDbsRequest) SetAccountName(v string) *DescribeDbInstanceDbsRequest {
	s.AccountName = &v
	return s
}

func (s *DescribeDbInstanceDbsRequest) SetPassword(v string) *DescribeDbInstanceDbsRequest {
	s.Password = &v
	return s
}

func (s *DescribeDbInstanceDbsRequest) SetDbInstType(v string) *DescribeDbInstanceDbsRequest {
	s.DbInstType = &v
	return s
}

type DescribeDbInstanceDbsResponseBody struct {
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                       `json:"Success,omitempty" xml:"Success,omitempty"`
	Total     *string                                     `json:"Total,omitempty" xml:"Total,omitempty"`
	Databases *DescribeDbInstanceDbsResponseBodyDatabases `json:"Databases,omitempty" xml:"Databases,omitempty" type:"Struct"`
}

func (s DescribeDbInstanceDbsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDbInstanceDbsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDbInstanceDbsResponseBody) SetRequestId(v string) *DescribeDbInstanceDbsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDbInstanceDbsResponseBody) SetSuccess(v bool) *DescribeDbInstanceDbsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDbInstanceDbsResponseBody) SetTotal(v string) *DescribeDbInstanceDbsResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeDbInstanceDbsResponseBody) SetDatabases(v *DescribeDbInstanceDbsResponseBodyDatabases) *DescribeDbInstanceDbsResponseBody {
	s.Databases = v
	return s
}

type DescribeDbInstanceDbsResponseBodyDatabases struct {
	Database []*DescribeDbInstanceDbsResponseBodyDatabasesDatabase `json:"Database,omitempty" xml:"Database,omitempty" type:"Repeated"`
}

func (s DescribeDbInstanceDbsResponseBodyDatabases) String() string {
	return tea.Prettify(s)
}

func (s DescribeDbInstanceDbsResponseBodyDatabases) GoString() string {
	return s.String()
}

func (s *DescribeDbInstanceDbsResponseBodyDatabases) SetDatabase(v []*DescribeDbInstanceDbsResponseBodyDatabasesDatabase) *DescribeDbInstanceDbsResponseBodyDatabases {
	s.Database = v
	return s
}

type DescribeDbInstanceDbsResponseBodyDatabasesDatabase struct {
	Status      *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	DbName      *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s DescribeDbInstanceDbsResponseBodyDatabasesDatabase) String() string {
	return tea.Prettify(s)
}

func (s DescribeDbInstanceDbsResponseBodyDatabasesDatabase) GoString() string {
	return s.String()
}

func (s *DescribeDbInstanceDbsResponseBodyDatabasesDatabase) SetStatus(v int32) *DescribeDbInstanceDbsResponseBodyDatabasesDatabase {
	s.Status = &v
	return s
}

func (s *DescribeDbInstanceDbsResponseBodyDatabasesDatabase) SetDbName(v string) *DescribeDbInstanceDbsResponseBodyDatabasesDatabase {
	s.DbName = &v
	return s
}

func (s *DescribeDbInstanceDbsResponseBodyDatabasesDatabase) SetDescription(v string) *DescribeDbInstanceDbsResponseBodyDatabasesDatabase {
	s.Description = &v
	return s
}

type DescribeDbInstanceDbsResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDbInstanceDbsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDbInstanceDbsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDbInstanceDbsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDbInstanceDbsResponse) SetHeaders(v map[string]*string) *DescribeDbInstanceDbsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDbInstanceDbsResponse) SetBody(v *DescribeDbInstanceDbsResponseBody) *DescribeDbInstanceDbsResponse {
	s.Body = v
	return s
}

type DescribeDbInstancesRequest struct {
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	Search         *string `json:"Search,omitempty" xml:"Search,omitempty"`
	DbInstType     *string `json:"DbInstType,omitempty" xml:"DbInstType,omitempty"`
	PageNumber     *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeDbInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDbInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDbInstancesRequest) SetDrdsInstanceId(v string) *DescribeDbInstancesRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeDbInstancesRequest) SetSearch(v string) *DescribeDbInstancesRequest {
	s.Search = &v
	return s
}

func (s *DescribeDbInstancesRequest) SetDbInstType(v string) *DescribeDbInstancesRequest {
	s.DbInstType = &v
	return s
}

func (s *DescribeDbInstancesRequest) SetPageNumber(v int32) *DescribeDbInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDbInstancesRequest) SetPageSize(v int32) *DescribeDbInstancesRequest {
	s.PageSize = &v
	return s
}

type DescribeDbInstancesResponseBody struct {
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Items     *DescribeDbInstancesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
}

func (s DescribeDbInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDbInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDbInstancesResponseBody) SetRequestId(v string) *DescribeDbInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDbInstancesResponseBody) SetItems(v *DescribeDbInstancesResponseBodyItems) *DescribeDbInstancesResponseBody {
	s.Items = v
	return s
}

type DescribeDbInstancesResponseBodyItems struct {
	DBInstance []*DescribeDbInstancesResponseBodyItemsDBInstance `json:"DBInstance,omitempty" xml:"DBInstance,omitempty" type:"Repeated"`
}

func (s DescribeDbInstancesResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDbInstancesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDbInstancesResponseBodyItems) SetDBInstance(v []*DescribeDbInstancesResponseBodyItemsDBInstance) *DescribeDbInstancesResponseBodyItems {
	s.DBInstance = v
	return s
}

type DescribeDbInstancesResponseBodyItemsDBInstance struct {
	InstanceNetworkType   *string                                                             `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	DBInstanceType        *string                                                             `json:"DBInstanceType,omitempty" xml:"DBInstanceType,omitempty"`
	ZoneId                *string                                                             `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	DBInstanceStatus      *int32                                                              `json:"DBInstanceStatus,omitempty" xml:"DBInstanceStatus,omitempty"`
	DBInstanceId          *string                                                             `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	Engine                *string                                                             `json:"Engine,omitempty" xml:"Engine,omitempty"`
	DBInstanceDescription *string                                                             `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	EngineVersion         *string                                                             `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	RegionId              *string                                                             `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ReadOnlyDBInstanceId  *DescribeDbInstancesResponseBodyItemsDBInstanceReadOnlyDBInstanceId `json:"ReadOnlyDBInstanceId,omitempty" xml:"ReadOnlyDBInstanceId,omitempty" type:"Struct"`
}

func (s DescribeDbInstancesResponseBodyItemsDBInstance) String() string {
	return tea.Prettify(s)
}

func (s DescribeDbInstancesResponseBodyItemsDBInstance) GoString() string {
	return s.String()
}

func (s *DescribeDbInstancesResponseBodyItemsDBInstance) SetInstanceNetworkType(v string) *DescribeDbInstancesResponseBodyItemsDBInstance {
	s.InstanceNetworkType = &v
	return s
}

func (s *DescribeDbInstancesResponseBodyItemsDBInstance) SetDBInstanceType(v string) *DescribeDbInstancesResponseBodyItemsDBInstance {
	s.DBInstanceType = &v
	return s
}

func (s *DescribeDbInstancesResponseBodyItemsDBInstance) SetZoneId(v string) *DescribeDbInstancesResponseBodyItemsDBInstance {
	s.ZoneId = &v
	return s
}

func (s *DescribeDbInstancesResponseBodyItemsDBInstance) SetDBInstanceStatus(v int32) *DescribeDbInstancesResponseBodyItemsDBInstance {
	s.DBInstanceStatus = &v
	return s
}

func (s *DescribeDbInstancesResponseBodyItemsDBInstance) SetDBInstanceId(v string) *DescribeDbInstancesResponseBodyItemsDBInstance {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDbInstancesResponseBodyItemsDBInstance) SetEngine(v string) *DescribeDbInstancesResponseBodyItemsDBInstance {
	s.Engine = &v
	return s
}

func (s *DescribeDbInstancesResponseBodyItemsDBInstance) SetDBInstanceDescription(v string) *DescribeDbInstancesResponseBodyItemsDBInstance {
	s.DBInstanceDescription = &v
	return s
}

func (s *DescribeDbInstancesResponseBodyItemsDBInstance) SetEngineVersion(v string) *DescribeDbInstancesResponseBodyItemsDBInstance {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDbInstancesResponseBodyItemsDBInstance) SetRegionId(v string) *DescribeDbInstancesResponseBodyItemsDBInstance {
	s.RegionId = &v
	return s
}

func (s *DescribeDbInstancesResponseBodyItemsDBInstance) SetReadOnlyDBInstanceId(v *DescribeDbInstancesResponseBodyItemsDBInstanceReadOnlyDBInstanceId) *DescribeDbInstancesResponseBodyItemsDBInstance {
	s.ReadOnlyDBInstanceId = v
	return s
}

type DescribeDbInstancesResponseBodyItemsDBInstanceReadOnlyDBInstanceId struct {
	ReadOnlyDBInstanceId []*string `json:"ReadOnlyDBInstanceId,omitempty" xml:"ReadOnlyDBInstanceId,omitempty" type:"Repeated"`
}

func (s DescribeDbInstancesResponseBodyItemsDBInstanceReadOnlyDBInstanceId) String() string {
	return tea.Prettify(s)
}

func (s DescribeDbInstancesResponseBodyItemsDBInstanceReadOnlyDBInstanceId) GoString() string {
	return s.String()
}

func (s *DescribeDbInstancesResponseBodyItemsDBInstanceReadOnlyDBInstanceId) SetReadOnlyDBInstanceId(v []*string) *DescribeDbInstancesResponseBodyItemsDBInstanceReadOnlyDBInstanceId {
	s.ReadOnlyDBInstanceId = v
	return s
}

type DescribeDbInstancesResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDbInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDbInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDbInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDbInstancesResponse) SetHeaders(v map[string]*string) *DescribeDbInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDbInstancesResponse) SetBody(v *DescribeDbInstancesResponseBody) *DescribeDbInstancesResponse {
	s.Body = v
	return s
}

type DescribeDrdsDBRequest struct {
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
}

func (s DescribeDrdsDBRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDBRequest) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDBRequest) SetDrdsInstanceId(v string) *DescribeDrdsDBRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeDrdsDBRequest) SetDbName(v string) *DescribeDrdsDBRequest {
	s.DbName = &v
	return s
}

type DescribeDrdsDBResponseBody struct {
	Success   *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *DescribeDrdsDBResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s DescribeDrdsDBResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDBResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDBResponseBody) SetSuccess(v bool) *DescribeDrdsDBResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDrdsDBResponseBody) SetRequestId(v string) *DescribeDrdsDBResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrdsDBResponseBody) SetData(v *DescribeDrdsDBResponseBodyData) *DescribeDrdsDBResponseBody {
	s.Data = v
	return s
}

type DescribeDrdsDBResponseBodyData struct {
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
	DbName     *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	Schema     *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Mode       *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	DbInstType *string `json:"DbInstType,omitempty" xml:"DbInstType,omitempty"`
	InstRole   *string `json:"InstRole,omitempty" xml:"InstRole,omitempty"`
}

func (s DescribeDrdsDBResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDBResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDBResponseBodyData) SetStatus(v string) *DescribeDrdsDBResponseBodyData {
	s.Status = &v
	return s
}

func (s *DescribeDrdsDBResponseBodyData) SetDbName(v string) *DescribeDrdsDBResponseBodyData {
	s.DbName = &v
	return s
}

func (s *DescribeDrdsDBResponseBodyData) SetSchema(v string) *DescribeDrdsDBResponseBodyData {
	s.Schema = &v
	return s
}

func (s *DescribeDrdsDBResponseBodyData) SetCreateTime(v string) *DescribeDrdsDBResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *DescribeDrdsDBResponseBodyData) SetMode(v string) *DescribeDrdsDBResponseBodyData {
	s.Mode = &v
	return s
}

func (s *DescribeDrdsDBResponseBodyData) SetDbInstType(v string) *DescribeDrdsDBResponseBodyData {
	s.DbInstType = &v
	return s
}

func (s *DescribeDrdsDBResponseBodyData) SetInstRole(v string) *DescribeDrdsDBResponseBodyData {
	s.InstRole = &v
	return s
}

type DescribeDrdsDBResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDrdsDBResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDrdsDBResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDBResponse) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDBResponse) SetHeaders(v map[string]*string) *DescribeDrdsDBResponse {
	s.Headers = v
	return s
}

func (s *DescribeDrdsDBResponse) SetBody(v *DescribeDrdsDBResponseBody) *DescribeDrdsDBResponse {
	s.Body = v
	return s
}

type DescribeDrdsDBClusterRequest struct {
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DbInstanceId   *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
}

func (s DescribeDrdsDBClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDBClusterRequest) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDBClusterRequest) SetDrdsInstanceId(v string) *DescribeDrdsDBClusterRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeDrdsDBClusterRequest) SetDbName(v string) *DescribeDrdsDBClusterRequest {
	s.DbName = &v
	return s
}

func (s *DescribeDrdsDBClusterRequest) SetDbInstanceId(v string) *DescribeDrdsDBClusterRequest {
	s.DbInstanceId = &v
	return s
}

type DescribeDrdsDBClusterResponseBody struct {
	Success    *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
	RequestId  *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DbInstance *DescribeDrdsDBClusterResponseBodyDbInstance `json:"DbInstance,omitempty" xml:"DbInstance,omitempty" type:"Struct"`
}

func (s DescribeDrdsDBClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDBClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDBClusterResponseBody) SetSuccess(v bool) *DescribeDrdsDBClusterResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDrdsDBClusterResponseBody) SetRequestId(v string) *DescribeDrdsDBClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrdsDBClusterResponseBody) SetDbInstance(v *DescribeDrdsDBClusterResponseBodyDbInstance) *DescribeDrdsDBClusterResponseBody {
	s.DbInstance = v
	return s
}

type DescribeDrdsDBClusterResponseBodyDbInstance struct {
	ExpireTime       *string                                               `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	PayType          *string                                               `json:"PayType,omitempty" xml:"PayType,omitempty"`
	DBInstanceStatus *string                                               `json:"DBInstanceStatus,omitempty" xml:"DBInstanceStatus,omitempty"`
	NetworkType      *string                                               `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	Port             *int32                                                `json:"Port,omitempty" xml:"Port,omitempty"`
	EngineVersion    *string                                               `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	RdsInstType      *string                                               `json:"RdsInstType,omitempty" xml:"RdsInstType,omitempty"`
	RemainDays       *string                                               `json:"RemainDays,omitempty" xml:"RemainDays,omitempty"`
	DBInstanceId     *string                                               `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DbInstType       *string                                               `json:"DbInstType,omitempty" xml:"DbInstType,omitempty"`
	Engine           *string                                               `json:"Engine,omitempty" xml:"Engine,omitempty"`
	ReadMode         *string                                               `json:"ReadMode,omitempty" xml:"ReadMode,omitempty"`
	Endpoints        *DescribeDrdsDBClusterResponseBodyDbInstanceEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Struct"`
	DBNodes          *DescribeDrdsDBClusterResponseBodyDbInstanceDBNodes   `json:"DBNodes,omitempty" xml:"DBNodes,omitempty" type:"Struct"`
}

func (s DescribeDrdsDBClusterResponseBodyDbInstance) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDBClusterResponseBodyDbInstance) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstance) SetExpireTime(v string) *DescribeDrdsDBClusterResponseBodyDbInstance {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstance) SetPayType(v string) *DescribeDrdsDBClusterResponseBodyDbInstance {
	s.PayType = &v
	return s
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstance) SetDBInstanceStatus(v string) *DescribeDrdsDBClusterResponseBodyDbInstance {
	s.DBInstanceStatus = &v
	return s
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstance) SetNetworkType(v string) *DescribeDrdsDBClusterResponseBodyDbInstance {
	s.NetworkType = &v
	return s
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstance) SetPort(v int32) *DescribeDrdsDBClusterResponseBodyDbInstance {
	s.Port = &v
	return s
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstance) SetEngineVersion(v string) *DescribeDrdsDBClusterResponseBodyDbInstance {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstance) SetRdsInstType(v string) *DescribeDrdsDBClusterResponseBodyDbInstance {
	s.RdsInstType = &v
	return s
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstance) SetRemainDays(v string) *DescribeDrdsDBClusterResponseBodyDbInstance {
	s.RemainDays = &v
	return s
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstance) SetDBInstanceId(v string) *DescribeDrdsDBClusterResponseBodyDbInstance {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstance) SetDbInstType(v string) *DescribeDrdsDBClusterResponseBodyDbInstance {
	s.DbInstType = &v
	return s
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstance) SetEngine(v string) *DescribeDrdsDBClusterResponseBodyDbInstance {
	s.Engine = &v
	return s
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstance) SetReadMode(v string) *DescribeDrdsDBClusterResponseBodyDbInstance {
	s.ReadMode = &v
	return s
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstance) SetEndpoints(v *DescribeDrdsDBClusterResponseBodyDbInstanceEndpoints) *DescribeDrdsDBClusterResponseBodyDbInstance {
	s.Endpoints = v
	return s
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstance) SetDBNodes(v *DescribeDrdsDBClusterResponseBodyDbInstanceDBNodes) *DescribeDrdsDBClusterResponseBodyDbInstance {
	s.DBNodes = v
	return s
}

type DescribeDrdsDBClusterResponseBodyDbInstanceEndpoints struct {
	Endpoint []*DescribeDrdsDBClusterResponseBodyDbInstanceEndpointsEndpoint `json:"Endpoint,omitempty" xml:"Endpoint,omitempty" type:"Repeated"`
}

func (s DescribeDrdsDBClusterResponseBodyDbInstanceEndpoints) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDBClusterResponseBodyDbInstanceEndpoints) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstanceEndpoints) SetEndpoint(v []*DescribeDrdsDBClusterResponseBodyDbInstanceEndpointsEndpoint) *DescribeDrdsDBClusterResponseBodyDbInstanceEndpoints {
	s.Endpoint = v
	return s
}

type DescribeDrdsDBClusterResponseBodyDbInstanceEndpointsEndpoint struct {
	ReadWeight *int32  `json:"ReadWeight,omitempty" xml:"ReadWeight,omitempty"`
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	NodeIds    *string `json:"NodeIds,omitempty" xml:"NodeIds,omitempty"`
}

func (s DescribeDrdsDBClusterResponseBodyDbInstanceEndpointsEndpoint) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDBClusterResponseBodyDbInstanceEndpointsEndpoint) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstanceEndpointsEndpoint) SetReadWeight(v int32) *DescribeDrdsDBClusterResponseBodyDbInstanceEndpointsEndpoint {
	s.ReadWeight = &v
	return s
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstanceEndpointsEndpoint) SetEndpointId(v string) *DescribeDrdsDBClusterResponseBodyDbInstanceEndpointsEndpoint {
	s.EndpointId = &v
	return s
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstanceEndpointsEndpoint) SetNodeIds(v string) *DescribeDrdsDBClusterResponseBodyDbInstanceEndpointsEndpoint {
	s.NodeIds = &v
	return s
}

type DescribeDrdsDBClusterResponseBodyDbInstanceDBNodes struct {
	DBNode []*DescribeDrdsDBClusterResponseBodyDbInstanceDBNodesDBNode `json:"DBNode,omitempty" xml:"DBNode,omitempty" type:"Repeated"`
}

func (s DescribeDrdsDBClusterResponseBodyDbInstanceDBNodes) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDBClusterResponseBodyDbInstanceDBNodes) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstanceDBNodes) SetDBNode(v []*DescribeDrdsDBClusterResponseBodyDbInstanceDBNodesDBNode) *DescribeDrdsDBClusterResponseBodyDbInstanceDBNodes {
	s.DBNode = v
	return s
}

type DescribeDrdsDBClusterResponseBodyDbInstanceDBNodesDBNode struct {
	DBNodeRole   *string `json:"DBNodeRole,omitempty" xml:"DBNodeRole,omitempty"`
	ZoneId       *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	DBNodeId     *string `json:"DBNodeId,omitempty" xml:"DBNodeId,omitempty"`
	DBNodeStatus *string `json:"DBNodeStatus,omitempty" xml:"DBNodeStatus,omitempty"`
}

func (s DescribeDrdsDBClusterResponseBodyDbInstanceDBNodesDBNode) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDBClusterResponseBodyDbInstanceDBNodesDBNode) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstanceDBNodesDBNode) SetDBNodeRole(v string) *DescribeDrdsDBClusterResponseBodyDbInstanceDBNodesDBNode {
	s.DBNodeRole = &v
	return s
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstanceDBNodesDBNode) SetZoneId(v string) *DescribeDrdsDBClusterResponseBodyDbInstanceDBNodesDBNode {
	s.ZoneId = &v
	return s
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstanceDBNodesDBNode) SetDBNodeId(v string) *DescribeDrdsDBClusterResponseBodyDbInstanceDBNodesDBNode {
	s.DBNodeId = &v
	return s
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstanceDBNodesDBNode) SetDBNodeStatus(v string) *DescribeDrdsDBClusterResponseBodyDbInstanceDBNodesDBNode {
	s.DBNodeStatus = &v
	return s
}

type DescribeDrdsDBClusterResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDrdsDBClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDrdsDBClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDBClusterResponse) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDBClusterResponse) SetHeaders(v map[string]*string) *DescribeDrdsDBClusterResponse {
	s.Headers = v
	return s
}

func (s *DescribeDrdsDBClusterResponse) SetBody(v *DescribeDrdsDBClusterResponseBody) *DescribeDrdsDBClusterResponse {
	s.Body = v
	return s
}

type DescribeDrdsDbInstanceRequest struct {
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DbInstanceId   *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
}

func (s DescribeDrdsDbInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDbInstanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDbInstanceRequest) SetDrdsInstanceId(v string) *DescribeDrdsDbInstanceRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeDrdsDbInstanceRequest) SetDbName(v string) *DescribeDrdsDbInstanceRequest {
	s.DbName = &v
	return s
}

func (s *DescribeDrdsDbInstanceRequest) SetDbInstanceId(v string) *DescribeDrdsDbInstanceRequest {
	s.DbInstanceId = &v
	return s
}

type DescribeDrdsDbInstanceResponseBody struct {
	Success    *bool                                         `json:"Success,omitempty" xml:"Success,omitempty"`
	RequestId  *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DbInstance *DescribeDrdsDbInstanceResponseBodyDbInstance `json:"DbInstance,omitempty" xml:"DbInstance,omitempty" type:"Struct"`
}

func (s DescribeDrdsDbInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDbInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDbInstanceResponseBody) SetSuccess(v bool) *DescribeDrdsDbInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBody) SetRequestId(v string) *DescribeDrdsDbInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBody) SetDbInstance(v *DescribeDrdsDbInstanceResponseBodyDbInstance) *DescribeDrdsDbInstanceResponseBody {
	s.DbInstance = v
	return s
}

type DescribeDrdsDbInstanceResponseBodyDbInstance struct {
	ExpireTime        *string                                                        `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	PayType           *string                                                        `json:"PayType,omitempty" xml:"PayType,omitempty"`
	DBInstanceStatus  *string                                                        `json:"DBInstanceStatus,omitempty" xml:"DBInstanceStatus,omitempty"`
	NetworkType       *string                                                        `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	Port              *int32                                                         `json:"Port,omitempty" xml:"Port,omitempty"`
	EngineVersion     *string                                                        `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	DmInstanceId      *string                                                        `json:"DmInstanceId,omitempty" xml:"DmInstanceId,omitempty"`
	ConnectUrl        *string                                                        `json:"ConnectUrl,omitempty" xml:"ConnectUrl,omitempty"`
	ReadWeight        *int32                                                         `json:"ReadWeight,omitempty" xml:"ReadWeight,omitempty"`
	RdsInstType       *string                                                        `json:"RdsInstType,omitempty" xml:"RdsInstType,omitempty"`
	RemainDays        *string                                                        `json:"RemainDays,omitempty" xml:"RemainDays,omitempty"`
	DBInstanceId      *string                                                        `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DbInstType        *string                                                        `json:"DbInstType,omitempty" xml:"DbInstType,omitempty"`
	Engine            *string                                                        `json:"Engine,omitempty" xml:"Engine,omitempty"`
	ReadOnlyInstances *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstances `json:"ReadOnlyInstances,omitempty" xml:"ReadOnlyInstances,omitempty" type:"Struct"`
}

func (s DescribeDrdsDbInstanceResponseBodyDbInstance) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDbInstanceResponseBodyDbInstance) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstance) SetExpireTime(v string) *DescribeDrdsDbInstanceResponseBodyDbInstance {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstance) SetPayType(v string) *DescribeDrdsDbInstanceResponseBodyDbInstance {
	s.PayType = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstance) SetDBInstanceStatus(v string) *DescribeDrdsDbInstanceResponseBodyDbInstance {
	s.DBInstanceStatus = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstance) SetNetworkType(v string) *DescribeDrdsDbInstanceResponseBodyDbInstance {
	s.NetworkType = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstance) SetPort(v int32) *DescribeDrdsDbInstanceResponseBodyDbInstance {
	s.Port = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstance) SetEngineVersion(v string) *DescribeDrdsDbInstanceResponseBodyDbInstance {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstance) SetDmInstanceId(v string) *DescribeDrdsDbInstanceResponseBodyDbInstance {
	s.DmInstanceId = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstance) SetConnectUrl(v string) *DescribeDrdsDbInstanceResponseBodyDbInstance {
	s.ConnectUrl = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstance) SetReadWeight(v int32) *DescribeDrdsDbInstanceResponseBodyDbInstance {
	s.ReadWeight = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstance) SetRdsInstType(v string) *DescribeDrdsDbInstanceResponseBodyDbInstance {
	s.RdsInstType = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstance) SetRemainDays(v string) *DescribeDrdsDbInstanceResponseBodyDbInstance {
	s.RemainDays = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstance) SetDBInstanceId(v string) *DescribeDrdsDbInstanceResponseBodyDbInstance {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstance) SetDbInstType(v string) *DescribeDrdsDbInstanceResponseBodyDbInstance {
	s.DbInstType = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstance) SetEngine(v string) *DescribeDrdsDbInstanceResponseBodyDbInstance {
	s.Engine = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstance) SetReadOnlyInstances(v *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstances) *DescribeDrdsDbInstanceResponseBodyDbInstance {
	s.ReadOnlyInstances = v
	return s
}

type DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstances struct {
	ReadOnlyInstance []*DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance `json:"ReadOnlyInstance,omitempty" xml:"ReadOnlyInstance,omitempty" type:"Repeated"`
}

func (s DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstances) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstances) SetReadOnlyInstance(v []*DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstances {
	s.ReadOnlyInstance = v
	return s
}

type DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance struct {
	ExpireTime       *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	PayType          *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	VersionAction    *int32  `json:"VersionAction,omitempty" xml:"VersionAction,omitempty"`
	DBInstanceStatus *string `json:"DBInstanceStatus,omitempty" xml:"DBInstanceStatus,omitempty"`
	NetworkType      *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	Port             *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	EngineVersion    *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	DmInstanceId     *string `json:"DmInstanceId,omitempty" xml:"DmInstanceId,omitempty"`
	ConnectUrl       *string `json:"ConnectUrl,omitempty" xml:"ConnectUrl,omitempty"`
	ReadWeight       *int32  `json:"ReadWeight,omitempty" xml:"ReadWeight,omitempty"`
	RdsInstType      *string `json:"RdsInstType,omitempty" xml:"RdsInstType,omitempty"`
	RemainDays       *string `json:"RemainDays,omitempty" xml:"RemainDays,omitempty"`
	DBInstanceId     *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DbInstType       *string `json:"DbInstType,omitempty" xml:"DbInstType,omitempty"`
	Engine           *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
}

func (s DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) SetExpireTime(v string) *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) SetPayType(v string) *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.PayType = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) SetVersionAction(v int32) *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.VersionAction = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) SetDBInstanceStatus(v string) *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.DBInstanceStatus = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) SetNetworkType(v string) *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.NetworkType = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) SetPort(v int32) *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.Port = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) SetEngineVersion(v string) *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) SetDmInstanceId(v string) *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.DmInstanceId = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) SetConnectUrl(v string) *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.ConnectUrl = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) SetReadWeight(v int32) *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.ReadWeight = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) SetRdsInstType(v string) *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.RdsInstType = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) SetRemainDays(v string) *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.RemainDays = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) SetDBInstanceId(v string) *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) SetDbInstType(v string) *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.DbInstType = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) SetEngine(v string) *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.Engine = &v
	return s
}

type DescribeDrdsDbInstanceResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDrdsDbInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDrdsDbInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDbInstanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDbInstanceResponse) SetHeaders(v map[string]*string) *DescribeDrdsDbInstanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeDrdsDbInstanceResponse) SetBody(v *DescribeDrdsDbInstanceResponseBody) *DescribeDrdsDbInstanceResponse {
	s.Body = v
	return s
}

type DescribeDrdsDbInstancesRequest struct {
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	PageNumber     *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeDrdsDbInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDbInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDbInstancesRequest) SetDrdsInstanceId(v string) *DescribeDrdsDbInstancesRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeDrdsDbInstancesRequest) SetDbName(v string) *DescribeDrdsDbInstancesRequest {
	s.DbName = &v
	return s
}

func (s *DescribeDrdsDbInstancesRequest) SetPageNumber(v int32) *DescribeDrdsDbInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDrdsDbInstancesRequest) SetPageSize(v int32) *DescribeDrdsDbInstancesRequest {
	s.PageSize = &v
	return s
}

type DescribeDrdsDbInstancesResponseBody struct {
	PageSize    *string                                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber  *string                                         `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	RequestId   *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total       *string                                         `json:"Total,omitempty" xml:"Total,omitempty"`
	Success     *bool                                           `json:"Success,omitempty" xml:"Success,omitempty"`
	DbInstances *DescribeDrdsDbInstancesResponseBodyDbInstances `json:"DbInstances,omitempty" xml:"DbInstances,omitempty" type:"Struct"`
}

func (s DescribeDrdsDbInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDbInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDbInstancesResponseBody) SetPageSize(v string) *DescribeDrdsDbInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBody) SetPageNumber(v string) *DescribeDrdsDbInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBody) SetRequestId(v string) *DescribeDrdsDbInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBody) SetTotal(v string) *DescribeDrdsDbInstancesResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBody) SetSuccess(v bool) *DescribeDrdsDbInstancesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBody) SetDbInstances(v *DescribeDrdsDbInstancesResponseBodyDbInstances) *DescribeDrdsDbInstancesResponseBody {
	s.DbInstances = v
	return s
}

type DescribeDrdsDbInstancesResponseBodyDbInstances struct {
	DbInstance []*DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance `json:"DbInstance,omitempty" xml:"DbInstance,omitempty" type:"Repeated"`
}

func (s DescribeDrdsDbInstancesResponseBodyDbInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDbInstancesResponseBodyDbInstances) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstances) SetDbInstance(v []*DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) *DescribeDrdsDbInstancesResponseBodyDbInstances {
	s.DbInstance = v
	return s
}

type DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance struct {
	ExpireTime        *string                                                                    `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	PayType           *string                                                                    `json:"PayType,omitempty" xml:"PayType,omitempty"`
	DBInstanceStatus  *string                                                                    `json:"DBInstanceStatus,omitempty" xml:"DBInstanceStatus,omitempty"`
	NetworkType       *string                                                                    `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	Port              *int32                                                                     `json:"Port,omitempty" xml:"Port,omitempty"`
	EngineVersion     *string                                                                    `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	DmInstanceId      *string                                                                    `json:"DmInstanceId,omitempty" xml:"DmInstanceId,omitempty"`
	ConnectUrl        *string                                                                    `json:"ConnectUrl,omitempty" xml:"ConnectUrl,omitempty"`
	ReadWeight        *int32                                                                     `json:"ReadWeight,omitempty" xml:"ReadWeight,omitempty"`
	RdsInstType       *string                                                                    `json:"RdsInstType,omitempty" xml:"RdsInstType,omitempty"`
	RemainDays        *int32                                                                     `json:"RemainDays,omitempty" xml:"RemainDays,omitempty"`
	DBInstanceId      *string                                                                    `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DbInstType        *string                                                                    `json:"DbInstType,omitempty" xml:"DbInstType,omitempty"`
	Engine            *string                                                                    `json:"Engine,omitempty" xml:"Engine,omitempty"`
	ReadOnlyInstances *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstances `json:"ReadOnlyInstances,omitempty" xml:"ReadOnlyInstances,omitempty" type:"Struct"`
}

func (s DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) SetExpireTime(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) SetPayType(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance {
	s.PayType = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) SetDBInstanceStatus(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance {
	s.DBInstanceStatus = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) SetNetworkType(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance {
	s.NetworkType = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) SetPort(v int32) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance {
	s.Port = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) SetEngineVersion(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) SetDmInstanceId(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance {
	s.DmInstanceId = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) SetConnectUrl(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance {
	s.ConnectUrl = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) SetReadWeight(v int32) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance {
	s.ReadWeight = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) SetRdsInstType(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance {
	s.RdsInstType = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) SetRemainDays(v int32) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance {
	s.RemainDays = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) SetDBInstanceId(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) SetDbInstType(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance {
	s.DbInstType = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) SetEngine(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance {
	s.Engine = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) SetReadOnlyInstances(v *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstances) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance {
	s.ReadOnlyInstances = v
	return s
}

type DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstances struct {
	ReadOnlyInstance []*DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance `json:"ReadOnlyInstance,omitempty" xml:"ReadOnlyInstance,omitempty" type:"Repeated"`
}

func (s DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstances) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstances) SetReadOnlyInstance(v []*DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstances {
	s.ReadOnlyInstance = v
	return s
}

type DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance struct {
	ExpireTime       *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	PayType          *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	DBInstanceStatus *string `json:"DBInstanceStatus,omitempty" xml:"DBInstanceStatus,omitempty"`
	NetworkType      *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	Port             *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	EngineVersion    *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	DmInstanceId     *string `json:"DmInstanceId,omitempty" xml:"DmInstanceId,omitempty"`
	ConnectUrl       *string `json:"ConnectUrl,omitempty" xml:"ConnectUrl,omitempty"`
	ReadWeight       *int32  `json:"ReadWeight,omitempty" xml:"ReadWeight,omitempty"`
	RdsInstType      *string `json:"RdsInstType,omitempty" xml:"RdsInstType,omitempty"`
	InstanceName     *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	RemainDays       *int32  `json:"RemainDays,omitempty" xml:"RemainDays,omitempty"`
	DbInstType       *string `json:"DbInstType,omitempty" xml:"DbInstType,omitempty"`
	Engine           *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
}

func (s DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) SetExpireTime(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) SetPayType(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.PayType = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) SetDBInstanceStatus(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.DBInstanceStatus = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) SetNetworkType(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.NetworkType = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) SetPort(v int32) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.Port = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) SetEngineVersion(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) SetDmInstanceId(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.DmInstanceId = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) SetConnectUrl(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.ConnectUrl = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) SetReadWeight(v int32) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.ReadWeight = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) SetRdsInstType(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.RdsInstType = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) SetInstanceName(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.InstanceName = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) SetRemainDays(v int32) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.RemainDays = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) SetDbInstType(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.DbInstType = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) SetEngine(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.Engine = &v
	return s
}

type DescribeDrdsDbInstancesResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDrdsDbInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDrdsDbInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDbInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDbInstancesResponse) SetHeaders(v map[string]*string) *DescribeDrdsDbInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDrdsDbInstancesResponse) SetBody(v *DescribeDrdsDbInstancesResponseBody) *DescribeDrdsDbInstancesResponse {
	s.Body = v
	return s
}

type DescribeDrdsDBIpWhiteListRequest struct {
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	GroupName      *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s DescribeDrdsDBIpWhiteListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDBIpWhiteListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDBIpWhiteListRequest) SetDrdsInstanceId(v string) *DescribeDrdsDBIpWhiteListRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeDrdsDBIpWhiteListRequest) SetDbName(v string) *DescribeDrdsDBIpWhiteListRequest {
	s.DbName = &v
	return s
}

func (s *DescribeDrdsDBIpWhiteListRequest) SetGroupName(v string) *DescribeDrdsDBIpWhiteListRequest {
	s.GroupName = &v
	return s
}

type DescribeDrdsDBIpWhiteListResponseBody struct {
	Success     *bool                                             `json:"Success,omitempty" xml:"Success,omitempty"`
	RequestId   *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	IpWhiteList *DescribeDrdsDBIpWhiteListResponseBodyIpWhiteList `json:"IpWhiteList,omitempty" xml:"IpWhiteList,omitempty" type:"Struct"`
}

func (s DescribeDrdsDBIpWhiteListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDBIpWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDBIpWhiteListResponseBody) SetSuccess(v bool) *DescribeDrdsDBIpWhiteListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDrdsDBIpWhiteListResponseBody) SetRequestId(v string) *DescribeDrdsDBIpWhiteListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrdsDBIpWhiteListResponseBody) SetIpWhiteList(v *DescribeDrdsDBIpWhiteListResponseBodyIpWhiteList) *DescribeDrdsDBIpWhiteListResponseBody {
	s.IpWhiteList = v
	return s
}

type DescribeDrdsDBIpWhiteListResponseBodyIpWhiteList struct {
	Ip []*string `json:"Ip,omitempty" xml:"Ip,omitempty" type:"Repeated"`
}

func (s DescribeDrdsDBIpWhiteListResponseBodyIpWhiteList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDBIpWhiteListResponseBodyIpWhiteList) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDBIpWhiteListResponseBodyIpWhiteList) SetIp(v []*string) *DescribeDrdsDBIpWhiteListResponseBodyIpWhiteList {
	s.Ip = v
	return s
}

type DescribeDrdsDBIpWhiteListResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDrdsDBIpWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDrdsDBIpWhiteListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDBIpWhiteListResponse) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDBIpWhiteListResponse) SetHeaders(v map[string]*string) *DescribeDrdsDBIpWhiteListResponse {
	s.Headers = v
	return s
}

func (s *DescribeDrdsDBIpWhiteListResponse) SetBody(v *DescribeDrdsDBIpWhiteListResponseBody) *DescribeDrdsDBIpWhiteListResponse {
	s.Body = v
	return s
}

type DescribeDrdsDbRdsNameListRequest struct {
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
}

func (s DescribeDrdsDbRdsNameListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDbRdsNameListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDbRdsNameListRequest) SetDrdsInstanceId(v string) *DescribeDrdsDbRdsNameListRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeDrdsDbRdsNameListRequest) SetDbName(v string) *DescribeDrdsDbRdsNameListRequest {
	s.DbName = &v
	return s
}

type DescribeDrdsDbRdsNameListResponseBody struct {
	Success          *bool                                                  `json:"Success,omitempty" xml:"Success,omitempty"`
	RequestId        *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	InstanceNameList *DescribeDrdsDbRdsNameListResponseBodyInstanceNameList `json:"InstanceNameList,omitempty" xml:"InstanceNameList,omitempty" type:"Struct"`
}

func (s DescribeDrdsDbRdsNameListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDbRdsNameListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDbRdsNameListResponseBody) SetSuccess(v bool) *DescribeDrdsDbRdsNameListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDrdsDbRdsNameListResponseBody) SetRequestId(v string) *DescribeDrdsDbRdsNameListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrdsDbRdsNameListResponseBody) SetInstanceNameList(v *DescribeDrdsDbRdsNameListResponseBodyInstanceNameList) *DescribeDrdsDbRdsNameListResponseBody {
	s.InstanceNameList = v
	return s
}

type DescribeDrdsDbRdsNameListResponseBodyInstanceNameList struct {
	InstanceName []*string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty" type:"Repeated"`
}

func (s DescribeDrdsDbRdsNameListResponseBodyInstanceNameList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDbRdsNameListResponseBodyInstanceNameList) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDbRdsNameListResponseBodyInstanceNameList) SetInstanceName(v []*string) *DescribeDrdsDbRdsNameListResponseBodyInstanceNameList {
	s.InstanceName = v
	return s
}

type DescribeDrdsDbRdsNameListResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDrdsDbRdsNameListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDrdsDbRdsNameListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDbRdsNameListResponse) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDbRdsNameListResponse) SetHeaders(v map[string]*string) *DescribeDrdsDbRdsNameListResponse {
	s.Headers = v
	return s
}

func (s *DescribeDrdsDbRdsNameListResponse) SetBody(v *DescribeDrdsDbRdsNameListResponseBody) *DescribeDrdsDbRdsNameListResponse {
	s.Body = v
	return s
}

type DescribeDrdsDBsRequest struct {
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	PageNumber     *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDrdsDBsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDBsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDBsRequest) SetDrdsInstanceId(v string) *DescribeDrdsDBsRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeDrdsDBsRequest) SetPageNumber(v int32) *DescribeDrdsDBsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDrdsDBsRequest) SetPageSize(v int32) *DescribeDrdsDBsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDrdsDBsRequest) SetRegionId(v string) *DescribeDrdsDBsRequest {
	s.RegionId = &v
	return s
}

type DescribeDrdsDBsResponseBody struct {
	PageSize   *string                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *string                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	RequestId  *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total      *string                          `json:"Total,omitempty" xml:"Total,omitempty"`
	Success    *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
	Data       *DescribeDrdsDBsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s DescribeDrdsDBsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDBsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDBsResponseBody) SetPageSize(v string) *DescribeDrdsDBsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDrdsDBsResponseBody) SetPageNumber(v string) *DescribeDrdsDBsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDrdsDBsResponseBody) SetRequestId(v string) *DescribeDrdsDBsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrdsDBsResponseBody) SetTotal(v string) *DescribeDrdsDBsResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeDrdsDBsResponseBody) SetSuccess(v bool) *DescribeDrdsDBsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDrdsDBsResponseBody) SetData(v *DescribeDrdsDBsResponseBodyData) *DescribeDrdsDBsResponseBody {
	s.Data = v
	return s
}

type DescribeDrdsDBsResponseBodyData struct {
	Db []*DescribeDrdsDBsResponseBodyDataDb `json:"Db,omitempty" xml:"Db,omitempty" type:"Repeated"`
}

func (s DescribeDrdsDBsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDBsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDBsResponseBodyData) SetDb(v []*DescribeDrdsDBsResponseBodyDataDb) *DescribeDrdsDBsResponseBodyData {
	s.Db = v
	return s
}

type DescribeDrdsDBsResponseBodyDataDb struct {
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
	DbName     *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	Schema     *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Mode       *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	DbInstType *string `json:"DbInstType,omitempty" xml:"DbInstType,omitempty"`
}

func (s DescribeDrdsDBsResponseBodyDataDb) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDBsResponseBodyDataDb) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDBsResponseBodyDataDb) SetStatus(v string) *DescribeDrdsDBsResponseBodyDataDb {
	s.Status = &v
	return s
}

func (s *DescribeDrdsDBsResponseBodyDataDb) SetDbName(v string) *DescribeDrdsDBsResponseBodyDataDb {
	s.DbName = &v
	return s
}

func (s *DescribeDrdsDBsResponseBodyDataDb) SetSchema(v string) *DescribeDrdsDBsResponseBodyDataDb {
	s.Schema = &v
	return s
}

func (s *DescribeDrdsDBsResponseBodyDataDb) SetCreateTime(v string) *DescribeDrdsDBsResponseBodyDataDb {
	s.CreateTime = &v
	return s
}

func (s *DescribeDrdsDBsResponseBodyDataDb) SetMode(v string) *DescribeDrdsDBsResponseBodyDataDb {
	s.Mode = &v
	return s
}

func (s *DescribeDrdsDBsResponseBodyDataDb) SetDbInstType(v string) *DescribeDrdsDBsResponseBodyDataDb {
	s.DbInstType = &v
	return s
}

type DescribeDrdsDBsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDrdsDBsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDrdsDBsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDBsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDBsResponse) SetHeaders(v map[string]*string) *DescribeDrdsDBsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDrdsDBsResponse) SetBody(v *DescribeDrdsDBsResponseBody) *DescribeDrdsDBsResponse {
	s.Body = v
	return s
}

type DescribeDrdsDbTasksRequest struct {
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	TaskType       *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DescribeDrdsDbTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDbTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDbTasksRequest) SetDrdsInstanceId(v string) *DescribeDrdsDbTasksRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeDrdsDbTasksRequest) SetDbName(v string) *DescribeDrdsDbTasksRequest {
	s.DbName = &v
	return s
}

func (s *DescribeDrdsDbTasksRequest) SetTaskType(v string) *DescribeDrdsDbTasksRequest {
	s.TaskType = &v
	return s
}

type DescribeDrdsDbTasksResponseBody struct {
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tasks     *DescribeDrdsDbTasksResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Struct"`
}

func (s DescribeDrdsDbTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDbTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDbTasksResponseBody) SetSuccess(v bool) *DescribeDrdsDbTasksResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDrdsDbTasksResponseBody) SetRequestId(v string) *DescribeDrdsDbTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrdsDbTasksResponseBody) SetTasks(v *DescribeDrdsDbTasksResponseBodyTasks) *DescribeDrdsDbTasksResponseBody {
	s.Tasks = v
	return s
}

type DescribeDrdsDbTasksResponseBodyTasks struct {
	Task []*DescribeDrdsDbTasksResponseBodyTasksTask `json:"Task,omitempty" xml:"Task,omitempty" type:"Repeated"`
}

func (s DescribeDrdsDbTasksResponseBodyTasks) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDbTasksResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDbTasksResponseBodyTasks) SetTask(v []*DescribeDrdsDbTasksResponseBodyTasksTask) *DescribeDrdsDbTasksResponseBodyTasks {
	s.Task = v
	return s
}

type DescribeDrdsDbTasksResponseBodyTasksTask struct {
	Progress        *int32  `json:"Progress,omitempty" xml:"Progress,omitempty"`
	TaskPhase       *string `json:"TaskPhase,omitempty" xml:"TaskPhase,omitempty"`
	TbComputeLength *int32  `json:"TbComputeLength,omitempty" xml:"TbComputeLength,omitempty"`
	TaskName        *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	ParentJobId     *string `json:"ParentJobId,omitempty" xml:"ParentJobId,omitempty"`
	Label           *string `json:"Label,omitempty" xml:"Label,omitempty"`
	TaskType        *int32  `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	DbComputeLength *int32  `json:"DbComputeLength,omitempty" xml:"DbComputeLength,omitempty"`
	TaskStatus      *int32  `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	AllowCancel     *bool   `json:"AllowCancel,omitempty" xml:"AllowCancel,omitempty"`
	ShowProgress    *bool   `json:"ShowProgress,omitempty" xml:"ShowProgress,omitempty"`
	TaskDetail      *string `json:"TaskDetail,omitempty" xml:"TaskDetail,omitempty"`
	GmtCreate       *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	DetailTaskId    *string `json:"DetailTaskId,omitempty" xml:"DetailTaskId,omitempty"`
	TargetId        *int64  `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	ExpandType      *string `json:"ExpandType,omitempty" xml:"ExpandType,omitempty"`
}

func (s DescribeDrdsDbTasksResponseBodyTasksTask) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDbTasksResponseBodyTasksTask) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDbTasksResponseBodyTasksTask) SetProgress(v int32) *DescribeDrdsDbTasksResponseBodyTasksTask {
	s.Progress = &v
	return s
}

func (s *DescribeDrdsDbTasksResponseBodyTasksTask) SetTaskPhase(v string) *DescribeDrdsDbTasksResponseBodyTasksTask {
	s.TaskPhase = &v
	return s
}

func (s *DescribeDrdsDbTasksResponseBodyTasksTask) SetTbComputeLength(v int32) *DescribeDrdsDbTasksResponseBodyTasksTask {
	s.TbComputeLength = &v
	return s
}

func (s *DescribeDrdsDbTasksResponseBodyTasksTask) SetTaskName(v string) *DescribeDrdsDbTasksResponseBodyTasksTask {
	s.TaskName = &v
	return s
}

func (s *DescribeDrdsDbTasksResponseBodyTasksTask) SetParentJobId(v string) *DescribeDrdsDbTasksResponseBodyTasksTask {
	s.ParentJobId = &v
	return s
}

func (s *DescribeDrdsDbTasksResponseBodyTasksTask) SetLabel(v string) *DescribeDrdsDbTasksResponseBodyTasksTask {
	s.Label = &v
	return s
}

func (s *DescribeDrdsDbTasksResponseBodyTasksTask) SetTaskType(v int32) *DescribeDrdsDbTasksResponseBodyTasksTask {
	s.TaskType = &v
	return s
}

func (s *DescribeDrdsDbTasksResponseBodyTasksTask) SetDbComputeLength(v int32) *DescribeDrdsDbTasksResponseBodyTasksTask {
	s.DbComputeLength = &v
	return s
}

func (s *DescribeDrdsDbTasksResponseBodyTasksTask) SetTaskStatus(v int32) *DescribeDrdsDbTasksResponseBodyTasksTask {
	s.TaskStatus = &v
	return s
}

func (s *DescribeDrdsDbTasksResponseBodyTasksTask) SetAllowCancel(v bool) *DescribeDrdsDbTasksResponseBodyTasksTask {
	s.AllowCancel = &v
	return s
}

func (s *DescribeDrdsDbTasksResponseBodyTasksTask) SetShowProgress(v bool) *DescribeDrdsDbTasksResponseBodyTasksTask {
	s.ShowProgress = &v
	return s
}

func (s *DescribeDrdsDbTasksResponseBodyTasksTask) SetTaskDetail(v string) *DescribeDrdsDbTasksResponseBodyTasksTask {
	s.TaskDetail = &v
	return s
}

func (s *DescribeDrdsDbTasksResponseBodyTasksTask) SetGmtCreate(v int64) *DescribeDrdsDbTasksResponseBodyTasksTask {
	s.GmtCreate = &v
	return s
}

func (s *DescribeDrdsDbTasksResponseBodyTasksTask) SetDetailTaskId(v string) *DescribeDrdsDbTasksResponseBodyTasksTask {
	s.DetailTaskId = &v
	return s
}

func (s *DescribeDrdsDbTasksResponseBodyTasksTask) SetTargetId(v int64) *DescribeDrdsDbTasksResponseBodyTasksTask {
	s.TargetId = &v
	return s
}

func (s *DescribeDrdsDbTasksResponseBodyTasksTask) SetExpandType(v string) *DescribeDrdsDbTasksResponseBodyTasksTask {
	s.ExpandType = &v
	return s
}

type DescribeDrdsDbTasksResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDrdsDbTasksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDrdsDbTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDbTasksResponse) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDbTasksResponse) SetHeaders(v map[string]*string) *DescribeDrdsDbTasksResponse {
	s.Headers = v
	return s
}

func (s *DescribeDrdsDbTasksResponse) SetBody(v *DescribeDrdsDbTasksResponseBody) *DescribeDrdsDbTasksResponse {
	s.Body = v
	return s
}

type DescribeDrdsInstanceRequest struct {
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDrdsInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceRequest) SetDrdsInstanceId(v string) *DescribeDrdsInstanceRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeDrdsInstanceRequest) SetRegionId(v string) *DescribeDrdsInstanceRequest {
	s.RegionId = &v
	return s
}

type DescribeDrdsInstanceResponseBody struct {
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *DescribeDrdsInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s DescribeDrdsInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceResponseBody) SetSuccess(v bool) *DescribeDrdsInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBody) SetRequestId(v string) *DescribeDrdsInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBody) SetData(v *DescribeDrdsInstanceResponseBodyData) *DescribeDrdsInstanceResponseBody {
	s.Data = v
	return s
}

type DescribeDrdsInstanceResponseBodyData struct {
	Type                  *string                                                    `json:"Type,omitempty" xml:"Type,omitempty"`
	Status                *string                                                    `json:"Status,omitempty" xml:"Status,omitempty"`
	CreateTime            *int64                                                     `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	VersionAction         *string                                                    `json:"VersionAction,omitempty" xml:"VersionAction,omitempty"`
	StorageType           *string                                                    `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	NetworkType           *string                                                    `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	Label                 *string                                                    `json:"Label,omitempty" xml:"Label,omitempty"`
	MysqlVersion          *int32                                                     `json:"MysqlVersion,omitempty" xml:"MysqlVersion,omitempty"`
	InstanceSpec          *string                                                    `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty"`
	VpcCloudInstanceId    *string                                                    `json:"VpcCloudInstanceId,omitempty" xml:"VpcCloudInstanceId,omitempty"`
	Description           *string                                                    `json:"Description,omitempty" xml:"Description,omitempty"`
	Version               *int64                                                     `json:"Version,omitempty" xml:"Version,omitempty"`
	ExpireDate            *int64                                                     `json:"ExpireDate,omitempty" xml:"ExpireDate,omitempty"`
	MasterInstanceId      *string                                                    `json:"MasterInstanceId,omitempty" xml:"MasterInstanceId,omitempty"`
	CommodityCode         *string                                                    `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	MachineType           *string                                                    `json:"MachineType,omitempty" xml:"MachineType,omitempty"`
	InstanceSeries        *string                                                    `json:"InstanceSeries,omitempty" xml:"InstanceSeries,omitempty"`
	ProductVersion        *string                                                    `json:"ProductVersion,omitempty" xml:"ProductVersion,omitempty"`
	RegionId              *string                                                    `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId       *string                                                    `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	DrdsInstanceId        *string                                                    `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	ZoneId                *string                                                    `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	InstRole              *string                                                    `json:"InstRole,omitempty" xml:"InstRole,omitempty"`
	OrderInstanceId       *string                                                    `json:"OrderInstanceId,omitempty" xml:"OrderInstanceId,omitempty"`
	Vips                  *DescribeDrdsInstanceResponseBodyDataVips                  `json:"Vips,omitempty" xml:"Vips,omitempty" type:"Struct"`
	ReadOnlyDBInstanceIds *DescribeDrdsInstanceResponseBodyDataReadOnlyDBInstanceIds `json:"ReadOnlyDBInstanceIds,omitempty" xml:"ReadOnlyDBInstanceIds,omitempty" type:"Struct"`
}

func (s DescribeDrdsInstanceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceResponseBodyData) SetType(v string) *DescribeDrdsInstanceResponseBodyData {
	s.Type = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetStatus(v string) *DescribeDrdsInstanceResponseBodyData {
	s.Status = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetCreateTime(v int64) *DescribeDrdsInstanceResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetVersionAction(v string) *DescribeDrdsInstanceResponseBodyData {
	s.VersionAction = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetStorageType(v string) *DescribeDrdsInstanceResponseBodyData {
	s.StorageType = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetNetworkType(v string) *DescribeDrdsInstanceResponseBodyData {
	s.NetworkType = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetLabel(v string) *DescribeDrdsInstanceResponseBodyData {
	s.Label = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetMysqlVersion(v int32) *DescribeDrdsInstanceResponseBodyData {
	s.MysqlVersion = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetInstanceSpec(v string) *DescribeDrdsInstanceResponseBodyData {
	s.InstanceSpec = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetVpcCloudInstanceId(v string) *DescribeDrdsInstanceResponseBodyData {
	s.VpcCloudInstanceId = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetDescription(v string) *DescribeDrdsInstanceResponseBodyData {
	s.Description = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetVersion(v int64) *DescribeDrdsInstanceResponseBodyData {
	s.Version = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetExpireDate(v int64) *DescribeDrdsInstanceResponseBodyData {
	s.ExpireDate = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetMasterInstanceId(v string) *DescribeDrdsInstanceResponseBodyData {
	s.MasterInstanceId = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetCommodityCode(v string) *DescribeDrdsInstanceResponseBodyData {
	s.CommodityCode = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetMachineType(v string) *DescribeDrdsInstanceResponseBodyData {
	s.MachineType = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetInstanceSeries(v string) *DescribeDrdsInstanceResponseBodyData {
	s.InstanceSeries = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetProductVersion(v string) *DescribeDrdsInstanceResponseBodyData {
	s.ProductVersion = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetRegionId(v string) *DescribeDrdsInstanceResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetResourceGroupId(v string) *DescribeDrdsInstanceResponseBodyData {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetDrdsInstanceId(v string) *DescribeDrdsInstanceResponseBodyData {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetZoneId(v string) *DescribeDrdsInstanceResponseBodyData {
	s.ZoneId = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetInstRole(v string) *DescribeDrdsInstanceResponseBodyData {
	s.InstRole = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetOrderInstanceId(v string) *DescribeDrdsInstanceResponseBodyData {
	s.OrderInstanceId = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetVips(v *DescribeDrdsInstanceResponseBodyDataVips) *DescribeDrdsInstanceResponseBodyData {
	s.Vips = v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetReadOnlyDBInstanceIds(v *DescribeDrdsInstanceResponseBodyDataReadOnlyDBInstanceIds) *DescribeDrdsInstanceResponseBodyData {
	s.ReadOnlyDBInstanceIds = v
	return s
}

type DescribeDrdsInstanceResponseBodyDataVips struct {
	Vip []*DescribeDrdsInstanceResponseBodyDataVipsVip `json:"Vip,omitempty" xml:"Vip,omitempty" type:"Repeated"`
}

func (s DescribeDrdsInstanceResponseBodyDataVips) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceResponseBodyDataVips) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceResponseBodyDataVips) SetVip(v []*DescribeDrdsInstanceResponseBodyDataVipsVip) *DescribeDrdsInstanceResponseBodyDataVips {
	s.Vip = v
	return s
}

type DescribeDrdsInstanceResponseBodyDataVipsVip struct {
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
	VpcId      *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VswitchId  *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	Dns        *string `json:"Dns,omitempty" xml:"Dns,omitempty"`
	Port       *string `json:"Port,omitempty" xml:"Port,omitempty"`
	ExpireDays *int64  `json:"ExpireDays,omitempty" xml:"ExpireDays,omitempty"`
}

func (s DescribeDrdsInstanceResponseBodyDataVipsVip) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceResponseBodyDataVipsVip) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceResponseBodyDataVipsVip) SetType(v string) *DescribeDrdsInstanceResponseBodyDataVipsVip {
	s.Type = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyDataVipsVip) SetVpcId(v string) *DescribeDrdsInstanceResponseBodyDataVipsVip {
	s.VpcId = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyDataVipsVip) SetVswitchId(v string) *DescribeDrdsInstanceResponseBodyDataVipsVip {
	s.VswitchId = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyDataVipsVip) SetDns(v string) *DescribeDrdsInstanceResponseBodyDataVipsVip {
	s.Dns = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyDataVipsVip) SetPort(v string) *DescribeDrdsInstanceResponseBodyDataVipsVip {
	s.Port = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyDataVipsVip) SetExpireDays(v int64) *DescribeDrdsInstanceResponseBodyDataVipsVip {
	s.ExpireDays = &v
	return s
}

type DescribeDrdsInstanceResponseBodyDataReadOnlyDBInstanceIds struct {
	ReadOnlyDBInstanceId []*string `json:"ReadOnlyDBInstanceId,omitempty" xml:"ReadOnlyDBInstanceId,omitempty" type:"Repeated"`
}

func (s DescribeDrdsInstanceResponseBodyDataReadOnlyDBInstanceIds) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceResponseBodyDataReadOnlyDBInstanceIds) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceResponseBodyDataReadOnlyDBInstanceIds) SetReadOnlyDBInstanceId(v []*string) *DescribeDrdsInstanceResponseBodyDataReadOnlyDBInstanceIds {
	s.ReadOnlyDBInstanceId = v
	return s
}

type DescribeDrdsInstanceResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDrdsInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDrdsInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceResponse) SetHeaders(v map[string]*string) *DescribeDrdsInstanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeDrdsInstanceResponse) SetBody(v *DescribeDrdsInstanceResponseBody) *DescribeDrdsInstanceResponse {
	s.Body = v
	return s
}

type DescribeDrdsInstanceDbMonitorRequest struct {
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	Key            *string `json:"Key,omitempty" xml:"Key,omitempty"`
	StartTime      *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime        *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDrdsInstanceDbMonitorRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceDbMonitorRequest) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceDbMonitorRequest) SetDrdsInstanceId(v string) *DescribeDrdsInstanceDbMonitorRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeDrdsInstanceDbMonitorRequest) SetDbName(v string) *DescribeDrdsInstanceDbMonitorRequest {
	s.DbName = &v
	return s
}

func (s *DescribeDrdsInstanceDbMonitorRequest) SetKey(v string) *DescribeDrdsInstanceDbMonitorRequest {
	s.Key = &v
	return s
}

func (s *DescribeDrdsInstanceDbMonitorRequest) SetStartTime(v int64) *DescribeDrdsInstanceDbMonitorRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDrdsInstanceDbMonitorRequest) SetEndTime(v int64) *DescribeDrdsInstanceDbMonitorRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDrdsInstanceDbMonitorRequest) SetRegionId(v string) *DescribeDrdsInstanceDbMonitorRequest {
	s.RegionId = &v
	return s
}

type DescribeDrdsInstanceDbMonitorResponseBody struct {
	Success   *bool                                            `json:"Success,omitempty" xml:"Success,omitempty"`
	RequestId *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      []*DescribeDrdsInstanceDbMonitorResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s DescribeDrdsInstanceDbMonitorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceDbMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceDbMonitorResponseBody) SetSuccess(v bool) *DescribeDrdsInstanceDbMonitorResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDrdsInstanceDbMonitorResponseBody) SetRequestId(v string) *DescribeDrdsInstanceDbMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrdsInstanceDbMonitorResponseBody) SetData(v []*DescribeDrdsInstanceDbMonitorResponseBodyData) *DescribeDrdsInstanceDbMonitorResponseBody {
	s.Data = v
	return s
}

type DescribeDrdsInstanceDbMonitorResponseBodyData struct {
	Key    *string                                                `json:"Key,omitempty" xml:"Key,omitempty"`
	Unit   *string                                                `json:"Unit,omitempty" xml:"Unit,omitempty"`
	Values []*DescribeDrdsInstanceDbMonitorResponseBodyDataValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s DescribeDrdsInstanceDbMonitorResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceDbMonitorResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceDbMonitorResponseBodyData) SetKey(v string) *DescribeDrdsInstanceDbMonitorResponseBodyData {
	s.Key = &v
	return s
}

func (s *DescribeDrdsInstanceDbMonitorResponseBodyData) SetUnit(v string) *DescribeDrdsInstanceDbMonitorResponseBodyData {
	s.Unit = &v
	return s
}

func (s *DescribeDrdsInstanceDbMonitorResponseBodyData) SetValues(v []*DescribeDrdsInstanceDbMonitorResponseBodyDataValues) *DescribeDrdsInstanceDbMonitorResponseBodyData {
	s.Values = v
	return s
}

type DescribeDrdsInstanceDbMonitorResponseBodyDataValues struct {
	Date  *int64  `json:"Date,omitempty" xml:"Date,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDrdsInstanceDbMonitorResponseBodyDataValues) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceDbMonitorResponseBodyDataValues) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceDbMonitorResponseBodyDataValues) SetDate(v int64) *DescribeDrdsInstanceDbMonitorResponseBodyDataValues {
	s.Date = &v
	return s
}

func (s *DescribeDrdsInstanceDbMonitorResponseBodyDataValues) SetValue(v string) *DescribeDrdsInstanceDbMonitorResponseBodyDataValues {
	s.Value = &v
	return s
}

type DescribeDrdsInstanceDbMonitorResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDrdsInstanceDbMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDrdsInstanceDbMonitorResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceDbMonitorResponse) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceDbMonitorResponse) SetHeaders(v map[string]*string) *DescribeDrdsInstanceDbMonitorResponse {
	s.Headers = v
	return s
}

func (s *DescribeDrdsInstanceDbMonitorResponse) SetBody(v *DescribeDrdsInstanceDbMonitorResponseBody) *DescribeDrdsInstanceDbMonitorResponse {
	s.Body = v
	return s
}

type DescribeDrdsInstanceLevelTasksRequest struct {
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s DescribeDrdsInstanceLevelTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceLevelTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceLevelTasksRequest) SetDrdsInstanceId(v string) *DescribeDrdsInstanceLevelTasksRequest {
	s.DrdsInstanceId = &v
	return s
}

type DescribeDrdsInstanceLevelTasksResponseBody struct {
	Success   *bool                                            `json:"Success,omitempty" xml:"Success,omitempty"`
	RequestId *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tasks     *DescribeDrdsInstanceLevelTasksResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Struct"`
}

func (s DescribeDrdsInstanceLevelTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceLevelTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceLevelTasksResponseBody) SetSuccess(v bool) *DescribeDrdsInstanceLevelTasksResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDrdsInstanceLevelTasksResponseBody) SetRequestId(v string) *DescribeDrdsInstanceLevelTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrdsInstanceLevelTasksResponseBody) SetTasks(v *DescribeDrdsInstanceLevelTasksResponseBodyTasks) *DescribeDrdsInstanceLevelTasksResponseBody {
	s.Tasks = v
	return s
}

type DescribeDrdsInstanceLevelTasksResponseBodyTasks struct {
	Task []*DescribeDrdsInstanceLevelTasksResponseBodyTasksTask `json:"Task,omitempty" xml:"Task,omitempty" type:"Repeated"`
}

func (s DescribeDrdsInstanceLevelTasksResponseBodyTasks) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceLevelTasksResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceLevelTasksResponseBodyTasks) SetTask(v []*DescribeDrdsInstanceLevelTasksResponseBodyTasksTask) *DescribeDrdsInstanceLevelTasksResponseBodyTasks {
	s.Task = v
	return s
}

type DescribeDrdsInstanceLevelTasksResponseBodyTasksTask struct {
	TaskType            *int32  `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	TaskPhase           *string `json:"TaskPhase,omitempty" xml:"TaskPhase,omitempty"`
	Progress            *int32  `json:"Progress,omitempty" xml:"Progress,omitempty"`
	AllowCancel         *bool   `json:"AllowCancel,omitempty" xml:"AllowCancel,omitempty"`
	TaskStatus          *int32  `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	ShowProgress        *bool   `json:"ShowProgress,omitempty" xml:"ShowProgress,omitempty"`
	TaskName            *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	ProgressDescription *string `json:"ProgressDescription,omitempty" xml:"ProgressDescription,omitempty"`
	GmtCreate           *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	TargetId            *int64  `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	ErrMsg              *string `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
}

func (s DescribeDrdsInstanceLevelTasksResponseBodyTasksTask) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceLevelTasksResponseBodyTasksTask) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask) SetTaskType(v int32) *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask {
	s.TaskType = &v
	return s
}

func (s *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask) SetTaskPhase(v string) *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask {
	s.TaskPhase = &v
	return s
}

func (s *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask) SetProgress(v int32) *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask {
	s.Progress = &v
	return s
}

func (s *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask) SetAllowCancel(v bool) *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask {
	s.AllowCancel = &v
	return s
}

func (s *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask) SetTaskStatus(v int32) *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask {
	s.TaskStatus = &v
	return s
}

func (s *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask) SetShowProgress(v bool) *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask {
	s.ShowProgress = &v
	return s
}

func (s *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask) SetTaskName(v string) *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask {
	s.TaskName = &v
	return s
}

func (s *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask) SetProgressDescription(v string) *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask {
	s.ProgressDescription = &v
	return s
}

func (s *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask) SetGmtCreate(v int64) *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask {
	s.GmtCreate = &v
	return s
}

func (s *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask) SetTargetId(v int64) *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask {
	s.TargetId = &v
	return s
}

func (s *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask) SetErrMsg(v string) *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask {
	s.ErrMsg = &v
	return s
}

type DescribeDrdsInstanceLevelTasksResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDrdsInstanceLevelTasksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDrdsInstanceLevelTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceLevelTasksResponse) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceLevelTasksResponse) SetHeaders(v map[string]*string) *DescribeDrdsInstanceLevelTasksResponse {
	s.Headers = v
	return s
}

func (s *DescribeDrdsInstanceLevelTasksResponse) SetBody(v *DescribeDrdsInstanceLevelTasksResponseBody) *DescribeDrdsInstanceLevelTasksResponse {
	s.Body = v
	return s
}

type DescribeDrdsInstanceMonitorRequest struct {
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	Key            *string `json:"Key,omitempty" xml:"Key,omitempty"`
	StartTime      *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime        *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PeriodMultiple *int32  `json:"PeriodMultiple,omitempty" xml:"PeriodMultiple,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDrdsInstanceMonitorRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceMonitorRequest) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceMonitorRequest) SetDrdsInstanceId(v string) *DescribeDrdsInstanceMonitorRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeDrdsInstanceMonitorRequest) SetKey(v string) *DescribeDrdsInstanceMonitorRequest {
	s.Key = &v
	return s
}

func (s *DescribeDrdsInstanceMonitorRequest) SetStartTime(v int64) *DescribeDrdsInstanceMonitorRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDrdsInstanceMonitorRequest) SetEndTime(v int64) *DescribeDrdsInstanceMonitorRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDrdsInstanceMonitorRequest) SetPeriodMultiple(v int32) *DescribeDrdsInstanceMonitorRequest {
	s.PeriodMultiple = &v
	return s
}

func (s *DescribeDrdsInstanceMonitorRequest) SetRegionId(v string) *DescribeDrdsInstanceMonitorRequest {
	s.RegionId = &v
	return s
}

type DescribeDrdsInstanceMonitorResponseBody struct {
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      []*DescribeDrdsInstanceMonitorResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s DescribeDrdsInstanceMonitorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceMonitorResponseBody) SetRequestId(v string) *DescribeDrdsInstanceMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrdsInstanceMonitorResponseBody) SetData(v []*DescribeDrdsInstanceMonitorResponseBodyData) *DescribeDrdsInstanceMonitorResponseBody {
	s.Data = v
	return s
}

type DescribeDrdsInstanceMonitorResponseBodyData struct {
	Key     *string                                              `json:"Key,omitempty" xml:"Key,omitempty"`
	Unit    *string                                              `json:"Unit,omitempty" xml:"Unit,omitempty"`
	NodeNum *int32                                               `json:"NodeNum,omitempty" xml:"NodeNum,omitempty"`
	Values  []*DescribeDrdsInstanceMonitorResponseBodyDataValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s DescribeDrdsInstanceMonitorResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceMonitorResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceMonitorResponseBodyData) SetKey(v string) *DescribeDrdsInstanceMonitorResponseBodyData {
	s.Key = &v
	return s
}

func (s *DescribeDrdsInstanceMonitorResponseBodyData) SetUnit(v string) *DescribeDrdsInstanceMonitorResponseBodyData {
	s.Unit = &v
	return s
}

func (s *DescribeDrdsInstanceMonitorResponseBodyData) SetNodeNum(v int32) *DescribeDrdsInstanceMonitorResponseBodyData {
	s.NodeNum = &v
	return s
}

func (s *DescribeDrdsInstanceMonitorResponseBodyData) SetValues(v []*DescribeDrdsInstanceMonitorResponseBodyDataValues) *DescribeDrdsInstanceMonitorResponseBodyData {
	s.Values = v
	return s
}

type DescribeDrdsInstanceMonitorResponseBodyDataValues struct {
	Date  *int64  `json:"Date,omitempty" xml:"Date,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDrdsInstanceMonitorResponseBodyDataValues) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceMonitorResponseBodyDataValues) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceMonitorResponseBodyDataValues) SetDate(v int64) *DescribeDrdsInstanceMonitorResponseBodyDataValues {
	s.Date = &v
	return s
}

func (s *DescribeDrdsInstanceMonitorResponseBodyDataValues) SetValue(v string) *DescribeDrdsInstanceMonitorResponseBodyDataValues {
	s.Value = &v
	return s
}

type DescribeDrdsInstanceMonitorResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDrdsInstanceMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDrdsInstanceMonitorResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceMonitorResponse) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceMonitorResponse) SetHeaders(v map[string]*string) *DescribeDrdsInstanceMonitorResponse {
	s.Headers = v
	return s
}

func (s *DescribeDrdsInstanceMonitorResponse) SetBody(v *DescribeDrdsInstanceMonitorResponseBody) *DescribeDrdsInstanceMonitorResponse {
	s.Body = v
	return s
}

type DescribeDrdsInstancesRequest struct {
	Type            *string                            `json:"Type,omitempty" xml:"Type,omitempty"`
	Description     *string                            `json:"Description,omitempty" xml:"Description,omitempty"`
	Expired         *bool                              `json:"Expired,omitempty" xml:"Expired,omitempty"`
	PageNumber      *int32                             `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceGroupId *string                            `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	RegionId        *string                            `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Mix             *bool                              `json:"Mix,omitempty" xml:"Mix,omitempty"`
	ProductVersion  *string                            `json:"ProductVersion,omitempty" xml:"ProductVersion,omitempty"`
	Tag             []*DescribeDrdsInstancesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeDrdsInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstancesRequest) SetType(v string) *DescribeDrdsInstancesRequest {
	s.Type = &v
	return s
}

func (s *DescribeDrdsInstancesRequest) SetDescription(v string) *DescribeDrdsInstancesRequest {
	s.Description = &v
	return s
}

func (s *DescribeDrdsInstancesRequest) SetExpired(v bool) *DescribeDrdsInstancesRequest {
	s.Expired = &v
	return s
}

func (s *DescribeDrdsInstancesRequest) SetPageNumber(v int32) *DescribeDrdsInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDrdsInstancesRequest) SetPageSize(v int32) *DescribeDrdsInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDrdsInstancesRequest) SetResourceGroupId(v string) *DescribeDrdsInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDrdsInstancesRequest) SetRegionId(v string) *DescribeDrdsInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDrdsInstancesRequest) SetMix(v bool) *DescribeDrdsInstancesRequest {
	s.Mix = &v
	return s
}

func (s *DescribeDrdsInstancesRequest) SetProductVersion(v string) *DescribeDrdsInstancesRequest {
	s.ProductVersion = &v
	return s
}

func (s *DescribeDrdsInstancesRequest) SetTag(v []*DescribeDrdsInstancesRequestTag) *DescribeDrdsInstancesRequest {
	s.Tag = v
	return s
}

type DescribeDrdsInstancesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDrdsInstancesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstancesRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstancesRequestTag) SetKey(v string) *DescribeDrdsInstancesRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeDrdsInstancesRequestTag) SetValue(v string) *DescribeDrdsInstancesRequestTag {
	s.Value = &v
	return s
}

type DescribeDrdsInstancesResponseBody struct {
	RequestId  *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int32                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total      *int32                                      `json:"Total,omitempty" xml:"Total,omitempty"`
	Instances  *DescribeDrdsInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Struct"`
}

func (s DescribeDrdsInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstancesResponseBody) SetRequestId(v string) *DescribeDrdsInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBody) SetPageNumber(v int32) *DescribeDrdsInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBody) SetPageSize(v int32) *DescribeDrdsInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBody) SetTotal(v int32) *DescribeDrdsInstancesResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBody) SetInstances(v *DescribeDrdsInstancesResponseBodyInstances) *DescribeDrdsInstancesResponseBody {
	s.Instances = v
	return s
}

type DescribeDrdsInstancesResponseBodyInstances struct {
	Instance []*DescribeDrdsInstancesResponseBodyInstancesInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Repeated"`
}

func (s DescribeDrdsInstancesResponseBodyInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstancesResponseBodyInstances) SetInstance(v []*DescribeDrdsInstancesResponseBodyInstancesInstance) *DescribeDrdsInstancesResponseBodyInstances {
	s.Instance = v
	return s
}

type DescribeDrdsInstancesResponseBodyInstancesInstance struct {
	VpcId                 *string                                                                  `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	Status                *string                                                                  `json:"Status,omitempty" xml:"Status,omitempty"`
	Type                  *string                                                                  `json:"Type,omitempty" xml:"Type,omitempty"`
	CommodityCode         *string                                                                  `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	MachineType           *string                                                                  `json:"MachineType,omitempty" xml:"MachineType,omitempty"`
	CreateTime            *int64                                                                   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	VersionAction         *string                                                                  `json:"VersionAction,omitempty" xml:"VersionAction,omitempty"`
	InstanceSeries        *string                                                                  `json:"InstanceSeries,omitempty" xml:"InstanceSeries,omitempty"`
	NetworkType           *string                                                                  `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	Label                 *string                                                                  `json:"Label,omitempty" xml:"Label,omitempty"`
	ProductVersion        *string                                                                  `json:"ProductVersion,omitempty" xml:"ProductVersion,omitempty"`
	InstanceSpec          *string                                                                  `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty"`
	RegionId              *string                                                                  `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	VpcCloudInstanceId    *string                                                                  `json:"VpcCloudInstanceId,omitempty" xml:"VpcCloudInstanceId,omitempty"`
	Description           *string                                                                  `json:"Description,omitempty" xml:"Description,omitempty"`
	Version               *int64                                                                   `json:"Version,omitempty" xml:"Version,omitempty"`
	ResourceGroupId       *string                                                                  `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ZoneId                *string                                                                  `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	DrdsInstanceId        *string                                                                  `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	ExpireDate            *int64                                                                   `json:"ExpireDate,omitempty" xml:"ExpireDate,omitempty"`
	MasterInstanceId      *string                                                                  `json:"MasterInstanceId,omitempty" xml:"MasterInstanceId,omitempty"`
	OrderInstanceId       *string                                                                  `json:"OrderInstanceId,omitempty" xml:"OrderInstanceId,omitempty"`
	InstRole              *string                                                                  `json:"InstRole,omitempty" xml:"InstRole,omitempty"`
	Vips                  *DescribeDrdsInstancesResponseBodyInstancesInstanceVips                  `json:"Vips,omitempty" xml:"Vips,omitempty" type:"Struct"`
	ReadOnlyDBInstanceIds *DescribeDrdsInstancesResponseBodyInstancesInstanceReadOnlyDBInstanceIds `json:"ReadOnlyDBInstanceIds,omitempty" xml:"ReadOnlyDBInstanceIds,omitempty" type:"Struct"`
}

func (s DescribeDrdsInstancesResponseBodyInstancesInstance) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstancesResponseBodyInstancesInstance) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetVpcId(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.VpcId = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetStatus(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.Status = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetType(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.Type = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetCommodityCode(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.CommodityCode = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetMachineType(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.MachineType = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetCreateTime(v int64) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.CreateTime = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetVersionAction(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.VersionAction = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetInstanceSeries(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.InstanceSeries = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetNetworkType(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.NetworkType = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetLabel(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.Label = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetProductVersion(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.ProductVersion = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetInstanceSpec(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.InstanceSpec = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetRegionId(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.RegionId = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetVpcCloudInstanceId(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.VpcCloudInstanceId = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetDescription(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.Description = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetVersion(v int64) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.Version = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetResourceGroupId(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetZoneId(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.ZoneId = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetDrdsInstanceId(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetExpireDate(v int64) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.ExpireDate = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetMasterInstanceId(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.MasterInstanceId = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetOrderInstanceId(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.OrderInstanceId = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetInstRole(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.InstRole = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetVips(v *DescribeDrdsInstancesResponseBodyInstancesInstanceVips) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.Vips = v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetReadOnlyDBInstanceIds(v *DescribeDrdsInstancesResponseBodyInstancesInstanceReadOnlyDBInstanceIds) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.ReadOnlyDBInstanceIds = v
	return s
}

type DescribeDrdsInstancesResponseBodyInstancesInstanceVips struct {
	Vip []*DescribeDrdsInstancesResponseBodyInstancesInstanceVipsVip `json:"Vip,omitempty" xml:"Vip,omitempty" type:"Repeated"`
}

func (s DescribeDrdsInstancesResponseBodyInstancesInstanceVips) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstancesResponseBodyInstancesInstanceVips) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstanceVips) SetVip(v []*DescribeDrdsInstancesResponseBodyInstancesInstanceVipsVip) *DescribeDrdsInstancesResponseBodyInstancesInstanceVips {
	s.Vip = v
	return s
}

type DescribeDrdsInstancesResponseBodyInstancesInstanceVipsVip struct {
	VpcId     *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty"`
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	Port      *string `json:"Port,omitempty" xml:"Port,omitempty"`
	IP        *string `json:"IP,omitempty" xml:"IP,omitempty"`
}

func (s DescribeDrdsInstancesResponseBodyInstancesInstanceVipsVip) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstancesResponseBodyInstancesInstanceVipsVip) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstanceVipsVip) SetVpcId(v string) *DescribeDrdsInstancesResponseBodyInstancesInstanceVipsVip {
	s.VpcId = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstanceVipsVip) SetType(v string) *DescribeDrdsInstancesResponseBodyInstancesInstanceVipsVip {
	s.Type = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstanceVipsVip) SetVswitchId(v string) *DescribeDrdsInstancesResponseBodyInstancesInstanceVipsVip {
	s.VswitchId = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstanceVipsVip) SetPort(v string) *DescribeDrdsInstancesResponseBodyInstancesInstanceVipsVip {
	s.Port = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstanceVipsVip) SetIP(v string) *DescribeDrdsInstancesResponseBodyInstancesInstanceVipsVip {
	s.IP = &v
	return s
}

type DescribeDrdsInstancesResponseBodyInstancesInstanceReadOnlyDBInstanceIds struct {
	ReadOnlyDBInstanceId []*string `json:"ReadOnlyDBInstanceId,omitempty" xml:"ReadOnlyDBInstanceId,omitempty" type:"Repeated"`
}

func (s DescribeDrdsInstancesResponseBodyInstancesInstanceReadOnlyDBInstanceIds) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstancesResponseBodyInstancesInstanceReadOnlyDBInstanceIds) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstanceReadOnlyDBInstanceIds) SetReadOnlyDBInstanceId(v []*string) *DescribeDrdsInstancesResponseBodyInstancesInstanceReadOnlyDBInstanceIds {
	s.ReadOnlyDBInstanceId = v
	return s
}

type DescribeDrdsInstancesResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDrdsInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDrdsInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstancesResponse) SetHeaders(v map[string]*string) *DescribeDrdsInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDrdsInstancesResponse) SetBody(v *DescribeDrdsInstancesResponseBody) *DescribeDrdsInstancesResponse {
	s.Body = v
	return s
}

type DescribeDrdsInstanceVersionRequest struct {
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDrdsInstanceVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceVersionRequest) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceVersionRequest) SetDrdsInstanceId(v string) *DescribeDrdsInstanceVersionRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeDrdsInstanceVersionRequest) SetRegionId(v string) *DescribeDrdsInstanceVersionRequest {
	s.RegionId = &v
	return s
}

type DescribeDrdsInstanceVersionResponseBody struct {
	Success   *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *DescribeDrdsInstanceVersionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s DescribeDrdsInstanceVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceVersionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceVersionResponseBody) SetSuccess(v bool) *DescribeDrdsInstanceVersionResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDrdsInstanceVersionResponseBody) SetRequestId(v string) *DescribeDrdsInstanceVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrdsInstanceVersionResponseBody) SetData(v *DescribeDrdsInstanceVersionResponseBodyData) *DescribeDrdsInstanceVersionResponseBody {
	s.Data = v
	return s
}

type DescribeDrdsInstanceVersionResponseBodyData struct {
	NewestVersion   *string `json:"NewestVersion,omitempty" xml:"NewestVersion,omitempty"`
	InstanceVersion *string `json:"InstanceVersion,omitempty" xml:"InstanceVersion,omitempty"`
}

func (s DescribeDrdsInstanceVersionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceVersionResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceVersionResponseBodyData) SetNewestVersion(v string) *DescribeDrdsInstanceVersionResponseBodyData {
	s.NewestVersion = &v
	return s
}

func (s *DescribeDrdsInstanceVersionResponseBodyData) SetInstanceVersion(v string) *DescribeDrdsInstanceVersionResponseBodyData {
	s.InstanceVersion = &v
	return s
}

type DescribeDrdsInstanceVersionResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDrdsInstanceVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDrdsInstanceVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceVersionResponse) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceVersionResponse) SetHeaders(v map[string]*string) *DescribeDrdsInstanceVersionResponse {
	s.Headers = v
	return s
}

func (s *DescribeDrdsInstanceVersionResponse) SetBody(v *DescribeDrdsInstanceVersionResponseBody) *DescribeDrdsInstanceVersionResponse {
	s.Body = v
	return s
}

type DescribeDrdsParamsRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	ParamLevel     *string `json:"ParamLevel,omitempty" xml:"ParamLevel,omitempty"`
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
}

func (s DescribeDrdsParamsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsParamsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDrdsParamsRequest) SetRegionId(v string) *DescribeDrdsParamsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDrdsParamsRequest) SetDrdsInstanceId(v string) *DescribeDrdsParamsRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeDrdsParamsRequest) SetParamLevel(v string) *DescribeDrdsParamsRequest {
	s.ParamLevel = &v
	return s
}

func (s *DescribeDrdsParamsRequest) SetDbName(v string) *DescribeDrdsParamsRequest {
	s.DbName = &v
	return s
}

type DescribeDrdsParamsResponseBody struct {
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	List      []*DescribeDrdsParamsResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s DescribeDrdsParamsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsParamsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsParamsResponseBody) SetSuccess(v bool) *DescribeDrdsParamsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDrdsParamsResponseBody) SetRequestId(v string) *DescribeDrdsParamsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrdsParamsResponseBody) SetList(v []*DescribeDrdsParamsResponseBodyList) *DescribeDrdsParamsResponseBody {
	s.List = v
	return s
}

type DescribeDrdsParamsResponseBodyList struct {
	ParamDefaultValue *string `json:"ParamDefaultValue,omitempty" xml:"ParamDefaultValue,omitempty"`
	ParamLevel        *string `json:"ParamLevel,omitempty" xml:"ParamLevel,omitempty"`
	ParamName         *string `json:"ParamName,omitempty" xml:"ParamName,omitempty"`
	ParamType         *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
	ParamValue        *string `json:"ParamValue,omitempty" xml:"ParamValue,omitempty"`
	NeedRestart       *bool   `json:"NeedRestart,omitempty" xml:"NeedRestart,omitempty"`
	ParamRanges       *string `json:"ParamRanges,omitempty" xml:"ParamRanges,omitempty"`
	DbName            *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	ParamEnglishName  *string `json:"ParamEnglishName,omitempty" xml:"ParamEnglishName,omitempty"`
	ParamDesc         *string `json:"ParamDesc,omitempty" xml:"ParamDesc,omitempty"`
	ParamVariableName *string `json:"ParamVariableName,omitempty" xml:"ParamVariableName,omitempty"`
}

func (s DescribeDrdsParamsResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsParamsResponseBodyList) GoString() string {
	return s.String()
}

func (s *DescribeDrdsParamsResponseBodyList) SetParamDefaultValue(v string) *DescribeDrdsParamsResponseBodyList {
	s.ParamDefaultValue = &v
	return s
}

func (s *DescribeDrdsParamsResponseBodyList) SetParamLevel(v string) *DescribeDrdsParamsResponseBodyList {
	s.ParamLevel = &v
	return s
}

func (s *DescribeDrdsParamsResponseBodyList) SetParamName(v string) *DescribeDrdsParamsResponseBodyList {
	s.ParamName = &v
	return s
}

func (s *DescribeDrdsParamsResponseBodyList) SetParamType(v string) *DescribeDrdsParamsResponseBodyList {
	s.ParamType = &v
	return s
}

func (s *DescribeDrdsParamsResponseBodyList) SetParamValue(v string) *DescribeDrdsParamsResponseBodyList {
	s.ParamValue = &v
	return s
}

func (s *DescribeDrdsParamsResponseBodyList) SetNeedRestart(v bool) *DescribeDrdsParamsResponseBodyList {
	s.NeedRestart = &v
	return s
}

func (s *DescribeDrdsParamsResponseBodyList) SetParamRanges(v string) *DescribeDrdsParamsResponseBodyList {
	s.ParamRanges = &v
	return s
}

func (s *DescribeDrdsParamsResponseBodyList) SetDbName(v string) *DescribeDrdsParamsResponseBodyList {
	s.DbName = &v
	return s
}

func (s *DescribeDrdsParamsResponseBodyList) SetParamEnglishName(v string) *DescribeDrdsParamsResponseBodyList {
	s.ParamEnglishName = &v
	return s
}

func (s *DescribeDrdsParamsResponseBodyList) SetParamDesc(v string) *DescribeDrdsParamsResponseBodyList {
	s.ParamDesc = &v
	return s
}

func (s *DescribeDrdsParamsResponseBodyList) SetParamVariableName(v string) *DescribeDrdsParamsResponseBodyList {
	s.ParamVariableName = &v
	return s
}

type DescribeDrdsParamsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDrdsParamsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDrdsParamsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsParamsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDrdsParamsResponse) SetHeaders(v map[string]*string) *DescribeDrdsParamsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDrdsParamsResponse) SetBody(v *DescribeDrdsParamsResponseBody) *DescribeDrdsParamsResponse {
	s.Body = v
	return s
}

type DescribeDrdsRdsInstancesRequest struct {
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s DescribeDrdsRdsInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsRdsInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDrdsRdsInstancesRequest) SetDrdsInstanceId(v string) *DescribeDrdsRdsInstancesRequest {
	s.DrdsInstanceId = &v
	return s
}

type DescribeDrdsRdsInstancesResponseBody struct {
	RequestId   *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success     *bool                                            `json:"Success,omitempty" xml:"Success,omitempty"`
	DbInstances *DescribeDrdsRdsInstancesResponseBodyDbInstances `json:"DbInstances,omitempty" xml:"DbInstances,omitempty" type:"Struct"`
}

func (s DescribeDrdsRdsInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsRdsInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsRdsInstancesResponseBody) SetRequestId(v string) *DescribeDrdsRdsInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrdsRdsInstancesResponseBody) SetSuccess(v bool) *DescribeDrdsRdsInstancesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDrdsRdsInstancesResponseBody) SetDbInstances(v *DescribeDrdsRdsInstancesResponseBodyDbInstances) *DescribeDrdsRdsInstancesResponseBody {
	s.DbInstances = v
	return s
}

type DescribeDrdsRdsInstancesResponseBodyDbInstances struct {
	DbInstance []*DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance `json:"DbInstance,omitempty" xml:"DbInstance,omitempty" type:"Repeated"`
}

func (s DescribeDrdsRdsInstancesResponseBodyDbInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsRdsInstancesResponseBodyDbInstances) GoString() string {
	return s.String()
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstances) SetDbInstance(v []*DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) *DescribeDrdsRdsInstancesResponseBodyDbInstances {
	s.DbInstance = v
	return s
}

type DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance struct {
	DBInstanceCPU       *string `json:"DBInstanceCPU,omitempty" xml:"DBInstanceCPU,omitempty"`
	DBInstanceMemory    *int64  `json:"DBInstanceMemory,omitempty" xml:"DBInstanceMemory,omitempty"`
	PayType             *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	DBInstanceStatus    *string `json:"DBInstanceStatus,omitempty" xml:"DBInstanceStatus,omitempty"`
	NetworkType         *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	Port                *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	EngineVersion       *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	DmInstanceId        *string `json:"DmInstanceId,omitempty" xml:"DmInstanceId,omitempty"`
	DBInstanceStorage   *int64  `json:"DBInstanceStorage,omitempty" xml:"DBInstanceStorage,omitempty"`
	ConnectUrl          *string `json:"ConnectUrl,omitempty" xml:"ConnectUrl,omitempty"`
	ReadWeight          *int32  `json:"ReadWeight,omitempty" xml:"ReadWeight,omitempty"`
	RdsInstType         *string `json:"RdsInstType,omitempty" xml:"RdsInstType,omitempty"`
	DBInstanceClassType *string `json:"DBInstanceClassType,omitempty" xml:"DBInstanceClassType,omitempty"`
	DBInstanceId        *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	Engine              *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	DbInstType          *string `json:"DbInstType,omitempty" xml:"DbInstType,omitempty"`
}

func (s DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) GoString() string {
	return s.String()
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) SetDBInstanceCPU(v string) *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance {
	s.DBInstanceCPU = &v
	return s
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) SetDBInstanceMemory(v int64) *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance {
	s.DBInstanceMemory = &v
	return s
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) SetPayType(v string) *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance {
	s.PayType = &v
	return s
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) SetDBInstanceStatus(v string) *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance {
	s.DBInstanceStatus = &v
	return s
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) SetNetworkType(v string) *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance {
	s.NetworkType = &v
	return s
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) SetPort(v int32) *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance {
	s.Port = &v
	return s
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) SetEngineVersion(v string) *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) SetDmInstanceId(v string) *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance {
	s.DmInstanceId = &v
	return s
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) SetDBInstanceStorage(v int64) *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance {
	s.DBInstanceStorage = &v
	return s
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) SetConnectUrl(v string) *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance {
	s.ConnectUrl = &v
	return s
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) SetReadWeight(v int32) *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance {
	s.ReadWeight = &v
	return s
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) SetRdsInstType(v string) *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance {
	s.RdsInstType = &v
	return s
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) SetDBInstanceClassType(v string) *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance {
	s.DBInstanceClassType = &v
	return s
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) SetDBInstanceId(v string) *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) SetEngine(v string) *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance {
	s.Engine = &v
	return s
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) SetDbInstType(v string) *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance {
	s.DbInstType = &v
	return s
}

type DescribeDrdsRdsInstancesResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDrdsRdsInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDrdsRdsInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsRdsInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDrdsRdsInstancesResponse) SetHeaders(v map[string]*string) *DescribeDrdsRdsInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDrdsRdsInstancesResponse) SetBody(v *DescribeDrdsRdsInstancesResponseBody) *DescribeDrdsRdsInstancesResponse {
	s.Body = v
	return s
}

type DescribeDrdsShardingDbsRequest struct {
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DbNamePattern  *string `json:"DbNamePattern,omitempty" xml:"DbNamePattern,omitempty"`
}

func (s DescribeDrdsShardingDbsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsShardingDbsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDrdsShardingDbsRequest) SetDrdsInstanceId(v string) *DescribeDrdsShardingDbsRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeDrdsShardingDbsRequest) SetDbName(v string) *DescribeDrdsShardingDbsRequest {
	s.DbName = &v
	return s
}

func (s *DescribeDrdsShardingDbsRequest) SetDbNamePattern(v string) *DescribeDrdsShardingDbsRequest {
	s.DbNamePattern = &v
	return s
}

type DescribeDrdsShardingDbsResponseBody struct {
	Success     *bool                                           `json:"Success,omitempty" xml:"Success,omitempty"`
	RequestId   *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ShardingDbs *DescribeDrdsShardingDbsResponseBodyShardingDbs `json:"ShardingDbs,omitempty" xml:"ShardingDbs,omitempty" type:"Struct"`
}

func (s DescribeDrdsShardingDbsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsShardingDbsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsShardingDbsResponseBody) SetSuccess(v bool) *DescribeDrdsShardingDbsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDrdsShardingDbsResponseBody) SetRequestId(v string) *DescribeDrdsShardingDbsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrdsShardingDbsResponseBody) SetShardingDbs(v *DescribeDrdsShardingDbsResponseBodyShardingDbs) *DescribeDrdsShardingDbsResponseBody {
	s.ShardingDbs = v
	return s
}

type DescribeDrdsShardingDbsResponseBodyShardingDbs struct {
	ShardingDb []*DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb `json:"ShardingDb,omitempty" xml:"ShardingDb,omitempty" type:"Repeated"`
}

func (s DescribeDrdsShardingDbsResponseBodyShardingDbs) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsShardingDbsResponseBodyShardingDbs) GoString() string {
	return s.String()
}

func (s *DescribeDrdsShardingDbsResponseBodyShardingDbs) SetShardingDb(v []*DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb) *DescribeDrdsShardingDbsResponseBodyShardingDbs {
	s.ShardingDb = v
	return s
}

type DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb struct {
	MinPoolSize                *int32  `json:"MinPoolSize,omitempty" xml:"MinPoolSize,omitempty"`
	MaxPoolSize                *int32  `json:"MaxPoolSize,omitempty" xml:"MaxPoolSize,omitempty"`
	DbInstanceId               *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
	ConnectUrl                 *string `json:"ConnectUrl,omitempty" xml:"ConnectUrl,omitempty"`
	GroupName                  *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	DbType                     *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	IdleTimeOut                *int32  `json:"IdleTimeOut,omitempty" xml:"IdleTimeOut,omitempty"`
	ShardingDbName             *string `json:"ShardingDbName,omitempty" xml:"ShardingDbName,omitempty"`
	BlockingTimeout            *int32  `json:"BlockingTimeout,omitempty" xml:"BlockingTimeout,omitempty"`
	PreparedStatementCacheSize *int32  `json:"PreparedStatementCacheSize,omitempty" xml:"PreparedStatementCacheSize,omitempty"`
	ConnectionProperties       *string `json:"ConnectionProperties,omitempty" xml:"ConnectionProperties,omitempty"`
	UserName                   *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	DbStatus                   *string `json:"DbStatus,omitempty" xml:"DbStatus,omitempty"`
}

func (s DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb) GoString() string {
	return s.String()
}

func (s *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb) SetMinPoolSize(v int32) *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb {
	s.MinPoolSize = &v
	return s
}

func (s *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb) SetMaxPoolSize(v int32) *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb {
	s.MaxPoolSize = &v
	return s
}

func (s *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb) SetDbInstanceId(v string) *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb {
	s.DbInstanceId = &v
	return s
}

func (s *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb) SetConnectUrl(v string) *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb {
	s.ConnectUrl = &v
	return s
}

func (s *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb) SetGroupName(v string) *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb {
	s.GroupName = &v
	return s
}

func (s *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb) SetDbType(v string) *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb {
	s.DbType = &v
	return s
}

func (s *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb) SetIdleTimeOut(v int32) *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb {
	s.IdleTimeOut = &v
	return s
}

func (s *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb) SetShardingDbName(v string) *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb {
	s.ShardingDbName = &v
	return s
}

func (s *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb) SetBlockingTimeout(v int32) *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb {
	s.BlockingTimeout = &v
	return s
}

func (s *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb) SetPreparedStatementCacheSize(v int32) *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb {
	s.PreparedStatementCacheSize = &v
	return s
}

func (s *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb) SetConnectionProperties(v string) *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb {
	s.ConnectionProperties = &v
	return s
}

func (s *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb) SetUserName(v string) *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb {
	s.UserName = &v
	return s
}

func (s *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb) SetDbStatus(v string) *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb {
	s.DbStatus = &v
	return s
}

type DescribeDrdsShardingDbsResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDrdsShardingDbsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDrdsShardingDbsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsShardingDbsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDrdsShardingDbsResponse) SetHeaders(v map[string]*string) *DescribeDrdsShardingDbsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDrdsShardingDbsResponse) SetBody(v *DescribeDrdsShardingDbsResponseBody) *DescribeDrdsShardingDbsResponse {
	s.Body = v
	return s
}

type DescribeDrdsSlowSqlsRequest struct {
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	ExeTime        *int64  `json:"ExeTime,omitempty" xml:"ExeTime,omitempty"`
	StartTime      *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime        *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageNumber     *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeDrdsSlowSqlsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsSlowSqlsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDrdsSlowSqlsRequest) SetDrdsInstanceId(v string) *DescribeDrdsSlowSqlsRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeDrdsSlowSqlsRequest) SetDbName(v string) *DescribeDrdsSlowSqlsRequest {
	s.DbName = &v
	return s
}

func (s *DescribeDrdsSlowSqlsRequest) SetExeTime(v int64) *DescribeDrdsSlowSqlsRequest {
	s.ExeTime = &v
	return s
}

func (s *DescribeDrdsSlowSqlsRequest) SetStartTime(v int64) *DescribeDrdsSlowSqlsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDrdsSlowSqlsRequest) SetEndTime(v int64) *DescribeDrdsSlowSqlsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDrdsSlowSqlsRequest) SetPageNumber(v int32) *DescribeDrdsSlowSqlsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDrdsSlowSqlsRequest) SetPageSize(v int32) *DescribeDrdsSlowSqlsRequest {
	s.PageSize = &v
	return s
}

type DescribeDrdsSlowSqlsResponseBody struct {
	PageSize   *int32                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int32                                 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	RequestId  *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total      *int32                                 `json:"Total,omitempty" xml:"Total,omitempty"`
	Success    *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
	Items      *DescribeDrdsSlowSqlsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
}

func (s DescribeDrdsSlowSqlsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsSlowSqlsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsSlowSqlsResponseBody) SetPageSize(v int32) *DescribeDrdsSlowSqlsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDrdsSlowSqlsResponseBody) SetPageNumber(v int32) *DescribeDrdsSlowSqlsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDrdsSlowSqlsResponseBody) SetRequestId(v string) *DescribeDrdsSlowSqlsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrdsSlowSqlsResponseBody) SetTotal(v int32) *DescribeDrdsSlowSqlsResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeDrdsSlowSqlsResponseBody) SetSuccess(v bool) *DescribeDrdsSlowSqlsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDrdsSlowSqlsResponseBody) SetItems(v *DescribeDrdsSlowSqlsResponseBodyItems) *DescribeDrdsSlowSqlsResponseBody {
	s.Items = v
	return s
}

type DescribeDrdsSlowSqlsResponseBodyItems struct {
	Item []*DescribeDrdsSlowSqlsResponseBodyItemsItem `json:"Item,omitempty" xml:"Item,omitempty" type:"Repeated"`
}

func (s DescribeDrdsSlowSqlsResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsSlowSqlsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDrdsSlowSqlsResponseBodyItems) SetItem(v []*DescribeDrdsSlowSqlsResponseBodyItemsItem) *DescribeDrdsSlowSqlsResponseBodyItems {
	s.Item = v
	return s
}

type DescribeDrdsSlowSqlsResponseBodyItemsItem struct {
	SendTime     *int64  `json:"SendTime,omitempty" xml:"SendTime,omitempty"`
	Host         *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Sql          *string `json:"Sql,omitempty" xml:"Sql,omitempty"`
	ResponseTime *int64  `json:"ResponseTime,omitempty" xml:"ResponseTime,omitempty"`
	Schema       *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
}

func (s DescribeDrdsSlowSqlsResponseBodyItemsItem) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsSlowSqlsResponseBodyItemsItem) GoString() string {
	return s.String()
}

func (s *DescribeDrdsSlowSqlsResponseBodyItemsItem) SetSendTime(v int64) *DescribeDrdsSlowSqlsResponseBodyItemsItem {
	s.SendTime = &v
	return s
}

func (s *DescribeDrdsSlowSqlsResponseBodyItemsItem) SetHost(v string) *DescribeDrdsSlowSqlsResponseBodyItemsItem {
	s.Host = &v
	return s
}

func (s *DescribeDrdsSlowSqlsResponseBodyItemsItem) SetSql(v string) *DescribeDrdsSlowSqlsResponseBodyItemsItem {
	s.Sql = &v
	return s
}

func (s *DescribeDrdsSlowSqlsResponseBodyItemsItem) SetResponseTime(v int64) *DescribeDrdsSlowSqlsResponseBodyItemsItem {
	s.ResponseTime = &v
	return s
}

func (s *DescribeDrdsSlowSqlsResponseBodyItemsItem) SetSchema(v string) *DescribeDrdsSlowSqlsResponseBodyItemsItem {
	s.Schema = &v
	return s
}

type DescribeDrdsSlowSqlsResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDrdsSlowSqlsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDrdsSlowSqlsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsSlowSqlsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDrdsSlowSqlsResponse) SetHeaders(v map[string]*string) *DescribeDrdsSlowSqlsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDrdsSlowSqlsResponse) SetBody(v *DescribeDrdsSlowSqlsResponseBody) *DescribeDrdsSlowSqlsResponse {
	s.Body = v
	return s
}

type DescribeDrdsSqlAuditStatusRequest struct {
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s DescribeDrdsSqlAuditStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsSqlAuditStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeDrdsSqlAuditStatusRequest) SetDrdsInstanceId(v string) *DescribeDrdsSqlAuditStatusRequest {
	s.DrdsInstanceId = &v
	return s
}

type DescribeDrdsSqlAuditStatusResponseBody struct {
	Success   *bool                                       `json:"Success,omitempty" xml:"Success,omitempty"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *DescribeDrdsSqlAuditStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s DescribeDrdsSqlAuditStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsSqlAuditStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsSqlAuditStatusResponseBody) SetSuccess(v bool) *DescribeDrdsSqlAuditStatusResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDrdsSqlAuditStatusResponseBody) SetRequestId(v string) *DescribeDrdsSqlAuditStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrdsSqlAuditStatusResponseBody) SetData(v *DescribeDrdsSqlAuditStatusResponseBodyData) *DescribeDrdsSqlAuditStatusResponseBody {
	s.Data = v
	return s
}

type DescribeDrdsSqlAuditStatusResponseBodyData struct {
	Data []*DescribeDrdsSqlAuditStatusResponseBodyDataData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s DescribeDrdsSqlAuditStatusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsSqlAuditStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDrdsSqlAuditStatusResponseBodyData) SetData(v []*DescribeDrdsSqlAuditStatusResponseBodyDataData) *DescribeDrdsSqlAuditStatusResponseBodyData {
	s.Data = v
	return s
}

type DescribeDrdsSqlAuditStatusResponseBodyDataData struct {
	ExtraSlsLogStore  *string `json:"ExtraSlsLogStore,omitempty" xml:"ExtraSlsLogStore,omitempty"`
	DbName            *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	Detailed          *string `json:"Detailed,omitempty" xml:"Detailed,omitempty"`
	ExtraWriteEnabled *bool   `json:"ExtraWriteEnabled,omitempty" xml:"ExtraWriteEnabled,omitempty"`
	Enabled           *string `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	ExtraAliUid       *int64  `json:"ExtraAliUid,omitempty" xml:"ExtraAliUid,omitempty"`
	ExtraSlsProject   *string `json:"ExtraSlsProject,omitempty" xml:"ExtraSlsProject,omitempty"`
}

func (s DescribeDrdsSqlAuditStatusResponseBodyDataData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsSqlAuditStatusResponseBodyDataData) GoString() string {
	return s.String()
}

func (s *DescribeDrdsSqlAuditStatusResponseBodyDataData) SetExtraSlsLogStore(v string) *DescribeDrdsSqlAuditStatusResponseBodyDataData {
	s.ExtraSlsLogStore = &v
	return s
}

func (s *DescribeDrdsSqlAuditStatusResponseBodyDataData) SetDbName(v string) *DescribeDrdsSqlAuditStatusResponseBodyDataData {
	s.DbName = &v
	return s
}

func (s *DescribeDrdsSqlAuditStatusResponseBodyDataData) SetDetailed(v string) *DescribeDrdsSqlAuditStatusResponseBodyDataData {
	s.Detailed = &v
	return s
}

func (s *DescribeDrdsSqlAuditStatusResponseBodyDataData) SetExtraWriteEnabled(v bool) *DescribeDrdsSqlAuditStatusResponseBodyDataData {
	s.ExtraWriteEnabled = &v
	return s
}

func (s *DescribeDrdsSqlAuditStatusResponseBodyDataData) SetEnabled(v string) *DescribeDrdsSqlAuditStatusResponseBodyDataData {
	s.Enabled = &v
	return s
}

func (s *DescribeDrdsSqlAuditStatusResponseBodyDataData) SetExtraAliUid(v int64) *DescribeDrdsSqlAuditStatusResponseBodyDataData {
	s.ExtraAliUid = &v
	return s
}

func (s *DescribeDrdsSqlAuditStatusResponseBodyDataData) SetExtraSlsProject(v string) *DescribeDrdsSqlAuditStatusResponseBodyDataData {
	s.ExtraSlsProject = &v
	return s
}

type DescribeDrdsSqlAuditStatusResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDrdsSqlAuditStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDrdsSqlAuditStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsSqlAuditStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeDrdsSqlAuditStatusResponse) SetHeaders(v map[string]*string) *DescribeDrdsSqlAuditStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeDrdsSqlAuditStatusResponse) SetBody(v *DescribeDrdsSqlAuditStatusResponseBody) *DescribeDrdsSqlAuditStatusResponse {
	s.Body = v
	return s
}

type DescribeDrdsTasksRequest struct {
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	TaskType       *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DescribeDrdsTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeDrdsTasksRequest) SetDrdsInstanceId(v string) *DescribeDrdsTasksRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeDrdsTasksRequest) SetDbName(v string) *DescribeDrdsTasksRequest {
	s.DbName = &v
	return s
}

func (s *DescribeDrdsTasksRequest) SetTaskType(v string) *DescribeDrdsTasksRequest {
	s.TaskType = &v
	return s
}

type DescribeDrdsTasksResponseBody struct {
	Success   *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tasks     *DescribeDrdsTasksResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Struct"`
}

func (s DescribeDrdsTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsTasksResponseBody) SetSuccess(v bool) *DescribeDrdsTasksResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDrdsTasksResponseBody) SetRequestId(v string) *DescribeDrdsTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrdsTasksResponseBody) SetTasks(v *DescribeDrdsTasksResponseBodyTasks) *DescribeDrdsTasksResponseBody {
	s.Tasks = v
	return s
}

type DescribeDrdsTasksResponseBodyTasks struct {
	Task []*DescribeDrdsTasksResponseBodyTasksTask `json:"Task,omitempty" xml:"Task,omitempty" type:"Repeated"`
}

func (s DescribeDrdsTasksResponseBodyTasks) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsTasksResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *DescribeDrdsTasksResponseBodyTasks) SetTask(v []*DescribeDrdsTasksResponseBodyTasksTask) *DescribeDrdsTasksResponseBodyTasks {
	s.Task = v
	return s
}

type DescribeDrdsTasksResponseBodyTasksTask struct {
	State   *string `json:"State,omitempty" xml:"State,omitempty"`
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Id      *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeDrdsTasksResponseBodyTasksTask) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsTasksResponseBodyTasksTask) GoString() string {
	return s.String()
}

func (s *DescribeDrdsTasksResponseBodyTasksTask) SetState(v string) *DescribeDrdsTasksResponseBodyTasksTask {
	s.State = &v
	return s
}

func (s *DescribeDrdsTasksResponseBodyTasksTask) SetContent(v string) *DescribeDrdsTasksResponseBodyTasksTask {
	s.Content = &v
	return s
}

func (s *DescribeDrdsTasksResponseBodyTasksTask) SetId(v int64) *DescribeDrdsTasksResponseBodyTasksTask {
	s.Id = &v
	return s
}

type DescribeDrdsTasksResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDrdsTasksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDrdsTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsTasksResponse) GoString() string {
	return s.String()
}

func (s *DescribeDrdsTasksResponse) SetHeaders(v map[string]*string) *DescribeDrdsTasksResponse {
	s.Headers = v
	return s
}

func (s *DescribeDrdsTasksResponse) SetBody(v *DescribeDrdsTasksResponseBody) *DescribeDrdsTasksResponse {
	s.Body = v
	return s
}

type DescribeExpandLogicTableInfoListRequest struct {
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
}

func (s DescribeExpandLogicTableInfoListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpandLogicTableInfoListRequest) GoString() string {
	return s.String()
}

func (s *DescribeExpandLogicTableInfoListRequest) SetDrdsInstanceId(v string) *DescribeExpandLogicTableInfoListRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeExpandLogicTableInfoListRequest) SetDbName(v string) *DescribeExpandLogicTableInfoListRequest {
	s.DbName = &v
	return s
}

type DescribeExpandLogicTableInfoListResponseBody struct {
	Success   *bool                                             `json:"Success,omitempty" xml:"Success,omitempty"`
	RequestId *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *DescribeExpandLogicTableInfoListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s DescribeExpandLogicTableInfoListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpandLogicTableInfoListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExpandLogicTableInfoListResponseBody) SetSuccess(v bool) *DescribeExpandLogicTableInfoListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeExpandLogicTableInfoListResponseBody) SetRequestId(v string) *DescribeExpandLogicTableInfoListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeExpandLogicTableInfoListResponseBody) SetData(v *DescribeExpandLogicTableInfoListResponseBodyData) *DescribeExpandLogicTableInfoListResponseBody {
	s.Data = v
	return s
}

type DescribeExpandLogicTableInfoListResponseBodyData struct {
	Data []*DescribeExpandLogicTableInfoListResponseBodyDataData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
}

func (s DescribeExpandLogicTableInfoListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpandLogicTableInfoListResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeExpandLogicTableInfoListResponseBodyData) SetData(v []*DescribeExpandLogicTableInfoListResponseBodyDataData) *DescribeExpandLogicTableInfoListResponseBodyData {
	s.Data = v
	return s
}

type DescribeExpandLogicTableInfoListResponseBodyDataData struct {
	ShardTbKey *string `json:"ShardTbKey,omitempty" xml:"ShardTbKey,omitempty"`
	TableName  *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	ShardDbKey *string `json:"ShardDbKey,omitempty" xml:"ShardDbKey,omitempty"`
}

func (s DescribeExpandLogicTableInfoListResponseBodyDataData) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpandLogicTableInfoListResponseBodyDataData) GoString() string {
	return s.String()
}

func (s *DescribeExpandLogicTableInfoListResponseBodyDataData) SetShardTbKey(v string) *DescribeExpandLogicTableInfoListResponseBodyDataData {
	s.ShardTbKey = &v
	return s
}

func (s *DescribeExpandLogicTableInfoListResponseBodyDataData) SetTableName(v string) *DescribeExpandLogicTableInfoListResponseBodyDataData {
	s.TableName = &v
	return s
}

func (s *DescribeExpandLogicTableInfoListResponseBodyDataData) SetShardDbKey(v string) *DescribeExpandLogicTableInfoListResponseBodyDataData {
	s.ShardDbKey = &v
	return s
}

type DescribeExpandLogicTableInfoListResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeExpandLogicTableInfoListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeExpandLogicTableInfoListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpandLogicTableInfoListResponse) GoString() string {
	return s.String()
}

func (s *DescribeExpandLogicTableInfoListResponse) SetHeaders(v map[string]*string) *DescribeExpandLogicTableInfoListResponse {
	s.Headers = v
	return s
}

func (s *DescribeExpandLogicTableInfoListResponse) SetBody(v *DescribeExpandLogicTableInfoListResponseBody) *DescribeExpandLogicTableInfoListResponse {
	s.Body = v
	return s
}

type DescribeHiStoreInstanceInfoRequest struct {
	DrdsInstanceId    *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	HistoreInstanceId *string `json:"HistoreInstanceId,omitempty" xml:"HistoreInstanceId,omitempty"`
}

func (s DescribeHiStoreInstanceInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeHiStoreInstanceInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeHiStoreInstanceInfoRequest) SetDrdsInstanceId(v string) *DescribeHiStoreInstanceInfoRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeHiStoreInstanceInfoRequest) SetHistoreInstanceId(v string) *DescribeHiStoreInstanceInfoRequest {
	s.HistoreInstanceId = &v
	return s
}

type DescribeHiStoreInstanceInfoResponseBody struct {
	Success             *bool                                                       `json:"Success,omitempty" xml:"Success,omitempty"`
	RequestId           *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HiStoreInstanceInfo *DescribeHiStoreInstanceInfoResponseBodyHiStoreInstanceInfo `json:"HiStoreInstanceInfo,omitempty" xml:"HiStoreInstanceInfo,omitempty" type:"Struct"`
}

func (s DescribeHiStoreInstanceInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeHiStoreInstanceInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHiStoreInstanceInfoResponseBody) SetSuccess(v bool) *DescribeHiStoreInstanceInfoResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeHiStoreInstanceInfoResponseBody) SetRequestId(v string) *DescribeHiStoreInstanceInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHiStoreInstanceInfoResponseBody) SetHiStoreInstanceInfo(v *DescribeHiStoreInstanceInfoResponseBodyHiStoreInstanceInfo) *DescribeHiStoreInstanceInfoResponseBody {
	s.HiStoreInstanceInfo = v
	return s
}

type DescribeHiStoreInstanceInfoResponseBodyHiStoreInstanceInfo struct {
	GmtCreate         *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	MachineSpec       *string `json:"MachineSpec,omitempty" xml:"MachineSpec,omitempty"`
	DiskSize          *int32  `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	RpmVersion        *string `json:"RpmVersion,omitempty" xml:"RpmVersion,omitempty"`
	HistoreInstanceId *string `json:"HistoreInstanceId,omitempty" xml:"HistoreInstanceId,omitempty"`
}

func (s DescribeHiStoreInstanceInfoResponseBodyHiStoreInstanceInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeHiStoreInstanceInfoResponseBodyHiStoreInstanceInfo) GoString() string {
	return s.String()
}

func (s *DescribeHiStoreInstanceInfoResponseBodyHiStoreInstanceInfo) SetGmtCreate(v int64) *DescribeHiStoreInstanceInfoResponseBodyHiStoreInstanceInfo {
	s.GmtCreate = &v
	return s
}

func (s *DescribeHiStoreInstanceInfoResponseBodyHiStoreInstanceInfo) SetMachineSpec(v string) *DescribeHiStoreInstanceInfoResponseBodyHiStoreInstanceInfo {
	s.MachineSpec = &v
	return s
}

func (s *DescribeHiStoreInstanceInfoResponseBodyHiStoreInstanceInfo) SetDiskSize(v int32) *DescribeHiStoreInstanceInfoResponseBodyHiStoreInstanceInfo {
	s.DiskSize = &v
	return s
}

func (s *DescribeHiStoreInstanceInfoResponseBodyHiStoreInstanceInfo) SetRpmVersion(v string) *DescribeHiStoreInstanceInfoResponseBodyHiStoreInstanceInfo {
	s.RpmVersion = &v
	return s
}

func (s *DescribeHiStoreInstanceInfoResponseBodyHiStoreInstanceInfo) SetHistoreInstanceId(v string) *DescribeHiStoreInstanceInfoResponseBodyHiStoreInstanceInfo {
	s.HistoreInstanceId = &v
	return s
}

type DescribeHiStoreInstanceInfoResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeHiStoreInstanceInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeHiStoreInstanceInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeHiStoreInstanceInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeHiStoreInstanceInfoResponse) SetHeaders(v map[string]*string) *DescribeHiStoreInstanceInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeHiStoreInstanceInfoResponse) SetBody(v *DescribeHiStoreInstanceInfoResponseBody) *DescribeHiStoreInstanceInfoResponse {
	s.Body = v
	return s
}

type DescribeHotDbListRequest struct {
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
}

func (s DescribeHotDbListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeHotDbListRequest) GoString() string {
	return s.String()
}

func (s *DescribeHotDbListRequest) SetDrdsInstanceId(v string) *DescribeHotDbListRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeHotDbListRequest) SetDbName(v string) *DescribeHotDbListRequest {
	s.DbName = &v
	return s
}

type DescribeHotDbListResponseBody struct {
	Msg       *string                            `json:"Msg,omitempty" xml:"Msg,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
	Data      *DescribeHotDbListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s DescribeHotDbListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeHotDbListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHotDbListResponseBody) SetMsg(v string) *DescribeHotDbListResponseBody {
	s.Msg = &v
	return s
}

func (s *DescribeHotDbListResponseBody) SetRequestId(v string) *DescribeHotDbListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHotDbListResponseBody) SetSuccess(v bool) *DescribeHotDbListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeHotDbListResponseBody) SetData(v *DescribeHotDbListResponseBodyData) *DescribeHotDbListResponseBody {
	s.Data = v
	return s
}

type DescribeHotDbListResponseBodyData struct {
	RandomCode *string                                `json:"RandomCode,omitempty" xml:"RandomCode,omitempty"`
	List       *DescribeHotDbListResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Struct"`
}

func (s DescribeHotDbListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeHotDbListResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeHotDbListResponseBodyData) SetRandomCode(v string) *DescribeHotDbListResponseBodyData {
	s.RandomCode = &v
	return s
}

func (s *DescribeHotDbListResponseBodyData) SetList(v *DescribeHotDbListResponseBodyDataList) *DescribeHotDbListResponseBodyData {
	s.List = v
	return s
}

type DescribeHotDbListResponseBodyDataList struct {
	InstanceDb []*DescribeHotDbListResponseBodyDataListInstanceDb `json:"InstanceDb,omitempty" xml:"InstanceDb,omitempty" type:"Repeated"`
}

func (s DescribeHotDbListResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s DescribeHotDbListResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *DescribeHotDbListResponseBodyDataList) SetInstanceDb(v []*DescribeHotDbListResponseBodyDataListInstanceDb) *DescribeHotDbListResponseBodyDataList {
	s.InstanceDb = v
	return s
}

type DescribeHotDbListResponseBodyDataListInstanceDb struct {
	InstanceName *string                                                   `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	HotDbList    *DescribeHotDbListResponseBodyDataListInstanceDbHotDbList `json:"HotDbList,omitempty" xml:"HotDbList,omitempty" type:"Struct"`
}

func (s DescribeHotDbListResponseBodyDataListInstanceDb) String() string {
	return tea.Prettify(s)
}

func (s DescribeHotDbListResponseBodyDataListInstanceDb) GoString() string {
	return s.String()
}

func (s *DescribeHotDbListResponseBodyDataListInstanceDb) SetInstanceName(v string) *DescribeHotDbListResponseBodyDataListInstanceDb {
	s.InstanceName = &v
	return s
}

func (s *DescribeHotDbListResponseBodyDataListInstanceDb) SetHotDbList(v *DescribeHotDbListResponseBodyDataListInstanceDbHotDbList) *DescribeHotDbListResponseBodyDataListInstanceDb {
	s.HotDbList = v
	return s
}

type DescribeHotDbListResponseBodyDataListInstanceDbHotDbList struct {
	Data []*string `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
}

func (s DescribeHotDbListResponseBodyDataListInstanceDbHotDbList) String() string {
	return tea.Prettify(s)
}

func (s DescribeHotDbListResponseBodyDataListInstanceDbHotDbList) GoString() string {
	return s.String()
}

func (s *DescribeHotDbListResponseBodyDataListInstanceDbHotDbList) SetData(v []*string) *DescribeHotDbListResponseBodyDataListInstanceDbHotDbList {
	s.Data = v
	return s
}

type DescribeHotDbListResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeHotDbListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeHotDbListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeHotDbListResponse) GoString() string {
	return s.String()
}

func (s *DescribeHotDbListResponse) SetHeaders(v map[string]*string) *DescribeHotDbListResponse {
	s.Headers = v
	return s
}

func (s *DescribeHotDbListResponse) SetBody(v *DescribeHotDbListResponseBody) *DescribeHotDbListResponse {
	s.Body = v
	return s
}

type DescribeInstanceAccountsRequest struct {
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s DescribeInstanceAccountsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceAccountsRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAccountsRequest) SetDrdsInstanceId(v string) *DescribeInstanceAccountsRequest {
	s.DrdsInstanceId = &v
	return s
}

type DescribeInstanceAccountsResponseBody struct {
	Success          *bool                                                 `json:"Success,omitempty" xml:"Success,omitempty"`
	RequestId        *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	InstanceAccounts *DescribeInstanceAccountsResponseBodyInstanceAccounts `json:"InstanceAccounts,omitempty" xml:"InstanceAccounts,omitempty" type:"Struct"`
}

func (s DescribeInstanceAccountsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAccountsResponseBody) SetSuccess(v bool) *DescribeInstanceAccountsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeInstanceAccountsResponseBody) SetRequestId(v string) *DescribeInstanceAccountsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceAccountsResponseBody) SetInstanceAccounts(v *DescribeInstanceAccountsResponseBodyInstanceAccounts) *DescribeInstanceAccountsResponseBody {
	s.InstanceAccounts = v
	return s
}

type DescribeInstanceAccountsResponseBodyInstanceAccounts struct {
	InstanceAccount []*DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccount `json:"InstanceAccount,omitempty" xml:"InstanceAccount,omitempty" type:"Repeated"`
}

func (s DescribeInstanceAccountsResponseBodyInstanceAccounts) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceAccountsResponseBodyInstanceAccounts) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAccountsResponseBodyInstanceAccounts) SetInstanceAccount(v []*DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccount) *DescribeInstanceAccountsResponseBodyInstanceAccounts {
	s.InstanceAccount = v
	return s
}

type DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccount struct {
	Host         *string                                                                          `json:"Host,omitempty" xml:"Host,omitempty"`
	Description  *string                                                                          `json:"Description,omitempty" xml:"Description,omitempty"`
	AccountType  *int32                                                                           `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	AccountName  *string                                                                          `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	DbPrivileges *DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccountDbPrivileges `json:"DbPrivileges,omitempty" xml:"DbPrivileges,omitempty" type:"Struct"`
}

func (s DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccount) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccount) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccount) SetHost(v string) *DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccount {
	s.Host = &v
	return s
}

func (s *DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccount) SetDescription(v string) *DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccount {
	s.Description = &v
	return s
}

func (s *DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccount) SetAccountType(v int32) *DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccount {
	s.AccountType = &v
	return s
}

func (s *DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccount) SetAccountName(v string) *DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccount {
	s.AccountName = &v
	return s
}

func (s *DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccount) SetDbPrivileges(v *DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccountDbPrivileges) *DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccount {
	s.DbPrivileges = v
	return s
}

type DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccountDbPrivileges struct {
	DbPrivilege []*DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccountDbPrivilegesDbPrivilege `json:"DbPrivilege,omitempty" xml:"DbPrivilege,omitempty" type:"Repeated"`
}

func (s DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccountDbPrivileges) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccountDbPrivileges) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccountDbPrivileges) SetDbPrivilege(v []*DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccountDbPrivilegesDbPrivilege) *DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccountDbPrivileges {
	s.DbPrivilege = v
	return s
}

type DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccountDbPrivilegesDbPrivilege struct {
	DbName    *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	Privilege *string `json:"Privilege,omitempty" xml:"Privilege,omitempty"`
}

func (s DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccountDbPrivilegesDbPrivilege) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccountDbPrivilegesDbPrivilege) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccountDbPrivilegesDbPrivilege) SetDbName(v string) *DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccountDbPrivilegesDbPrivilege {
	s.DbName = &v
	return s
}

func (s *DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccountDbPrivilegesDbPrivilege) SetPrivilege(v string) *DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccountDbPrivilegesDbPrivilege {
	s.Privilege = &v
	return s
}

type DescribeInstanceAccountsResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeInstanceAccountsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstanceAccountsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceAccountsResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAccountsResponse) SetHeaders(v map[string]*string) *DescribeInstanceAccountsResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceAccountsResponse) SetBody(v *DescribeInstanceAccountsResponseBody) *DescribeInstanceAccountsResponse {
	s.Body = v
	return s
}

type DescribeInstanceMenuSwitchRequest struct {
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s DescribeInstanceMenuSwitchRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceMenuSwitchRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceMenuSwitchRequest) SetDrdsInstanceId(v string) *DescribeInstanceMenuSwitchRequest {
	s.DrdsInstanceId = &v
	return s
}

type DescribeInstanceMenuSwitchResponseBody struct {
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                  `json:"Success,omitempty" xml:"Success,omitempty"`
	Config    map[string]interface{} `json:"Config,omitempty" xml:"Config,omitempty"`
}

func (s DescribeInstanceMenuSwitchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceMenuSwitchResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceMenuSwitchResponseBody) SetRequestId(v string) *DescribeInstanceMenuSwitchResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceMenuSwitchResponseBody) SetSuccess(v bool) *DescribeInstanceMenuSwitchResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeInstanceMenuSwitchResponseBody) SetConfig(v map[string]interface{}) *DescribeInstanceMenuSwitchResponseBody {
	s.Config = v
	return s
}

type DescribeInstanceMenuSwitchResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeInstanceMenuSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstanceMenuSwitchResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceMenuSwitchResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceMenuSwitchResponse) SetHeaders(v map[string]*string) *DescribeInstanceMenuSwitchResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceMenuSwitchResponse) SetBody(v *DescribeInstanceMenuSwitchResponseBody) *DescribeInstanceMenuSwitchResponse {
	s.Body = v
	return s
}

type DescribeInstanceSwitchAzoneRequest struct {
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s DescribeInstanceSwitchAzoneRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSwitchAzoneRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSwitchAzoneRequest) SetDrdsInstanceId(v string) *DescribeInstanceSwitchAzoneRequest {
	s.DrdsInstanceId = &v
	return s
}

type DescribeInstanceSwitchAzoneResponseBody struct {
	Success   *bool                                          `json:"Success,omitempty" xml:"Success,omitempty"`
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeInstanceSwitchAzoneResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribeInstanceSwitchAzoneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSwitchAzoneResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSwitchAzoneResponseBody) SetSuccess(v bool) *DescribeInstanceSwitchAzoneResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeInstanceSwitchAzoneResponseBody) SetRequestId(v string) *DescribeInstanceSwitchAzoneResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceSwitchAzoneResponseBody) SetResult(v *DescribeInstanceSwitchAzoneResponseBodyResult) *DescribeInstanceSwitchAzoneResponseBody {
	s.Result = v
	return s
}

type DescribeInstanceSwitchAzoneResponseBodyResult struct {
	OriginAzoneId *string                                                    `json:"OriginAzoneId,omitempty" xml:"OriginAzoneId,omitempty"`
	SwitchAble    *bool                                                      `json:"SwitchAble,omitempty" xml:"SwitchAble,omitempty"`
	RegionId      *string                                                    `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TargetAzones  *DescribeInstanceSwitchAzoneResponseBodyResultTargetAzones `json:"TargetAzones,omitempty" xml:"TargetAzones,omitempty" type:"Struct"`
}

func (s DescribeInstanceSwitchAzoneResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSwitchAzoneResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSwitchAzoneResponseBodyResult) SetOriginAzoneId(v string) *DescribeInstanceSwitchAzoneResponseBodyResult {
	s.OriginAzoneId = &v
	return s
}

func (s *DescribeInstanceSwitchAzoneResponseBodyResult) SetSwitchAble(v bool) *DescribeInstanceSwitchAzoneResponseBodyResult {
	s.SwitchAble = &v
	return s
}

func (s *DescribeInstanceSwitchAzoneResponseBodyResult) SetRegionId(v string) *DescribeInstanceSwitchAzoneResponseBodyResult {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceSwitchAzoneResponseBodyResult) SetTargetAzones(v *DescribeInstanceSwitchAzoneResponseBodyResultTargetAzones) *DescribeInstanceSwitchAzoneResponseBodyResult {
	s.TargetAzones = v
	return s
}

type DescribeInstanceSwitchAzoneResponseBodyResultTargetAzones struct {
	TargetAzone []*string `json:"TargetAzone,omitempty" xml:"TargetAzone,omitempty" type:"Repeated"`
}

func (s DescribeInstanceSwitchAzoneResponseBodyResultTargetAzones) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSwitchAzoneResponseBodyResultTargetAzones) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSwitchAzoneResponseBodyResultTargetAzones) SetTargetAzone(v []*string) *DescribeInstanceSwitchAzoneResponseBodyResultTargetAzones {
	s.TargetAzone = v
	return s
}

type DescribeInstanceSwitchAzoneResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeInstanceSwitchAzoneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstanceSwitchAzoneResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSwitchAzoneResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSwitchAzoneResponse) SetHeaders(v map[string]*string) *DescribeInstanceSwitchAzoneResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceSwitchAzoneResponse) SetBody(v *DescribeInstanceSwitchAzoneResponseBody) *DescribeInstanceSwitchAzoneResponse {
	s.Body = v
	return s
}

type DescribeInstanceSwitchNetworkRequest struct {
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s DescribeInstanceSwitchNetworkRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSwitchNetworkRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSwitchNetworkRequest) SetDrdsInstanceId(v string) *DescribeInstanceSwitchNetworkRequest {
	s.DrdsInstanceId = &v
	return s
}

type DescribeInstanceSwitchNetworkResponseBody struct {
	Success   *bool                                              `json:"Success,omitempty" xml:"Success,omitempty"`
	RequestId *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	VpcInfos  *DescribeInstanceSwitchNetworkResponseBodyVpcInfos `json:"VpcInfos,omitempty" xml:"VpcInfos,omitempty" type:"Struct"`
}

func (s DescribeInstanceSwitchNetworkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSwitchNetworkResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSwitchNetworkResponseBody) SetSuccess(v bool) *DescribeInstanceSwitchNetworkResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeInstanceSwitchNetworkResponseBody) SetRequestId(v string) *DescribeInstanceSwitchNetworkResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceSwitchNetworkResponseBody) SetVpcInfos(v *DescribeInstanceSwitchNetworkResponseBodyVpcInfos) *DescribeInstanceSwitchNetworkResponseBody {
	s.VpcInfos = v
	return s
}

type DescribeInstanceSwitchNetworkResponseBodyVpcInfos struct {
	VpcInfo []*DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfo `json:"VpcInfo,omitempty" xml:"VpcInfo,omitempty" type:"Repeated"`
}

func (s DescribeInstanceSwitchNetworkResponseBodyVpcInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSwitchNetworkResponseBodyVpcInfos) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSwitchNetworkResponseBodyVpcInfos) SetVpcInfo(v []*DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfo) *DescribeInstanceSwitchNetworkResponseBodyVpcInfos {
	s.VpcInfo = v
	return s
}

type DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfo struct {
	VpcId        *string                                                               `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VpcName      *string                                                               `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
	RegionId     *string                                                               `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	VswitchInfos *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfos `json:"VswitchInfos,omitempty" xml:"VswitchInfos,omitempty" type:"Struct"`
}

func (s DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfo) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfo) SetVpcId(v string) *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfo {
	s.VpcId = &v
	return s
}

func (s *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfo) SetVpcName(v string) *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfo {
	s.VpcName = &v
	return s
}

func (s *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfo) SetRegionId(v string) *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfo {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfo) SetVswitchInfos(v *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfos) *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfo {
	s.VswitchInfos = v
	return s
}

type DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfos struct {
	VswitchInfo []*DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfosVswitchInfo `json:"VswitchInfo,omitempty" xml:"VswitchInfo,omitempty" type:"Repeated"`
}

func (s DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfos) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfos) SetVswitchInfo(v []*DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfosVswitchInfo) *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfos {
	s.VswitchInfo = v
	return s
}

type DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfosVswitchInfo struct {
	VpcId         *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	DrdsSupported *bool   `json:"DrdsSupported,omitempty" xml:"DrdsSupported,omitempty"`
	VswitchId     *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	AzoneId       *string `json:"AzoneId,omitempty" xml:"AzoneId,omitempty"`
	VswitchName   *string `json:"VswitchName,omitempty" xml:"VswitchName,omitempty"`
}

func (s DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfosVswitchInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfosVswitchInfo) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfosVswitchInfo) SetVpcId(v string) *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfosVswitchInfo {
	s.VpcId = &v
	return s
}

func (s *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfosVswitchInfo) SetDrdsSupported(v bool) *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfosVswitchInfo {
	s.DrdsSupported = &v
	return s
}

func (s *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfosVswitchInfo) SetVswitchId(v string) *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfosVswitchInfo {
	s.VswitchId = &v
	return s
}

func (s *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfosVswitchInfo) SetAzoneId(v string) *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfosVswitchInfo {
	s.AzoneId = &v
	return s
}

func (s *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfosVswitchInfo) SetVswitchName(v string) *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfosVswitchInfo {
	s.VswitchName = &v
	return s
}

type DescribeInstanceSwitchNetworkResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeInstanceSwitchNetworkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstanceSwitchNetworkResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSwitchNetworkResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSwitchNetworkResponse) SetHeaders(v map[string]*string) *DescribeInstanceSwitchNetworkResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceSwitchNetworkResponse) SetBody(v *DescribeInstanceSwitchNetworkResponseBody) *DescribeInstanceSwitchNetworkResponse {
	s.Body = v
	return s
}

type DescribeInstDbLogInfoRequest struct {
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
}

func (s DescribeInstDbLogInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstDbLogInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstDbLogInfoRequest) SetDrdsInstanceId(v string) *DescribeInstDbLogInfoRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeInstDbLogInfoRequest) SetDbName(v string) *DescribeInstDbLogInfoRequest {
	s.DbName = &v
	return s
}

type DescribeInstDbLogInfoResponseBody struct {
	Success      *bool                                          `json:"Success,omitempty" xml:"Success,omitempty"`
	RequestId    *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	LogTimeRange *DescribeInstDbLogInfoResponseBodyLogTimeRange `json:"LogTimeRange,omitempty" xml:"LogTimeRange,omitempty" type:"Struct"`
}

func (s DescribeInstDbLogInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstDbLogInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstDbLogInfoResponseBody) SetSuccess(v bool) *DescribeInstDbLogInfoResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeInstDbLogInfoResponseBody) SetRequestId(v string) *DescribeInstDbLogInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstDbLogInfoResponseBody) SetLogTimeRange(v *DescribeInstDbLogInfoResponseBodyLogTimeRange) *DescribeInstDbLogInfoResponseBody {
	s.LogTimeRange = v
	return s
}

type DescribeInstDbLogInfoResponseBodyLogTimeRange struct {
	SupportOldestTime *int64 `json:"SupportOldestTime,omitempty" xml:"SupportOldestTime,omitempty"`
	SupportLatestTime *int64 `json:"SupportLatestTime,omitempty" xml:"SupportLatestTime,omitempty"`
}

func (s DescribeInstDbLogInfoResponseBodyLogTimeRange) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstDbLogInfoResponseBodyLogTimeRange) GoString() string {
	return s.String()
}

func (s *DescribeInstDbLogInfoResponseBodyLogTimeRange) SetSupportOldestTime(v int64) *DescribeInstDbLogInfoResponseBodyLogTimeRange {
	s.SupportOldestTime = &v
	return s
}

func (s *DescribeInstDbLogInfoResponseBodyLogTimeRange) SetSupportLatestTime(v int64) *DescribeInstDbLogInfoResponseBodyLogTimeRange {
	s.SupportLatestTime = &v
	return s
}

type DescribeInstDbLogInfoResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeInstDbLogInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstDbLogInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstDbLogInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstDbLogInfoResponse) SetHeaders(v map[string]*string) *DescribeInstDbLogInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstDbLogInfoResponse) SetBody(v *DescribeInstDbLogInfoResponseBody) *DescribeInstDbLogInfoResponse {
	s.Body = v
	return s
}

type DescribeInstDbSlsInfoRequest struct {
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
}

func (s DescribeInstDbSlsInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstDbSlsInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstDbSlsInfoRequest) SetDrdsInstanceId(v string) *DescribeInstDbSlsInfoRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeInstDbSlsInfoRequest) SetDbName(v string) *DescribeInstDbSlsInfoRequest {
	s.DbName = &v
	return s
}

type DescribeInstDbSlsInfoResponseBody struct {
	Success   *bool                                       `json:"Success,omitempty" xml:"Success,omitempty"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	AuditInfo *DescribeInstDbSlsInfoResponseBodyAuditInfo `json:"AuditInfo,omitempty" xml:"AuditInfo,omitempty" type:"Struct"`
}

func (s DescribeInstDbSlsInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstDbSlsInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstDbSlsInfoResponseBody) SetSuccess(v bool) *DescribeInstDbSlsInfoResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeInstDbSlsInfoResponseBody) SetRequestId(v string) *DescribeInstDbSlsInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstDbSlsInfoResponseBody) SetAuditInfo(v *DescribeInstDbSlsInfoResponseBodyAuditInfo) *DescribeInstDbSlsInfoResponseBody {
	s.AuditInfo = v
	return s
}

type DescribeInstDbSlsInfoResponseBodyAuditInfo struct {
	LogStore *string `json:"LogStore,omitempty" xml:"LogStore,omitempty"`
	Project  *string `json:"Project,omitempty" xml:"Project,omitempty"`
}

func (s DescribeInstDbSlsInfoResponseBodyAuditInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstDbSlsInfoResponseBodyAuditInfo) GoString() string {
	return s.String()
}

func (s *DescribeInstDbSlsInfoResponseBodyAuditInfo) SetLogStore(v string) *DescribeInstDbSlsInfoResponseBodyAuditInfo {
	s.LogStore = &v
	return s
}

func (s *DescribeInstDbSlsInfoResponseBodyAuditInfo) SetProject(v string) *DescribeInstDbSlsInfoResponseBodyAuditInfo {
	s.Project = &v
	return s
}

type DescribeInstDbSlsInfoResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeInstDbSlsInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstDbSlsInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstDbSlsInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstDbSlsInfoResponse) SetHeaders(v map[string]*string) *DescribeInstDbSlsInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstDbSlsInfoResponse) SetBody(v *DescribeInstDbSlsInfoResponseBody) *DescribeInstDbSlsInfoResponse {
	s.Body = v
	return s
}

type DescribePreCheckResultRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	TaskId         *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribePreCheckResultRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePreCheckResultRequest) GoString() string {
	return s.String()
}

func (s *DescribePreCheckResultRequest) SetRegionId(v string) *DescribePreCheckResultRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePreCheckResultRequest) SetDrdsInstanceId(v string) *DescribePreCheckResultRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribePreCheckResultRequest) SetTaskId(v string) *DescribePreCheckResultRequest {
	s.TaskId = &v
	return s
}

type DescribePreCheckResultResponseBody struct {
	Success        *bool                                             `json:"Success,omitempty" xml:"Success,omitempty"`
	RequestId      *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PreCheckResult *DescribePreCheckResultResponseBodyPreCheckResult `json:"PreCheckResult,omitempty" xml:"PreCheckResult,omitempty" type:"Struct"`
}

func (s DescribePreCheckResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePreCheckResultResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePreCheckResultResponseBody) SetSuccess(v bool) *DescribePreCheckResultResponseBody {
	s.Success = &v
	return s
}

func (s *DescribePreCheckResultResponseBody) SetRequestId(v string) *DescribePreCheckResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePreCheckResultResponseBody) SetPreCheckResult(v *DescribePreCheckResultResponseBodyPreCheckResult) *DescribePreCheckResultResponseBody {
	s.PreCheckResult = v
	return s
}

type DescribePreCheckResultResponseBodyPreCheckResult struct {
	PreCheckName  *string                                                          `json:"PreCheckName,omitempty" xml:"PreCheckName,omitempty"`
	State         *string                                                          `json:"State,omitempty" xml:"State,omitempty"`
	SubCheckItems []*DescribePreCheckResultResponseBodyPreCheckResultSubCheckItems `json:"SubCheckItems,omitempty" xml:"SubCheckItems,omitempty" type:"Repeated"`
}

func (s DescribePreCheckResultResponseBodyPreCheckResult) String() string {
	return tea.Prettify(s)
}

func (s DescribePreCheckResultResponseBodyPreCheckResult) GoString() string {
	return s.String()
}

func (s *DescribePreCheckResultResponseBodyPreCheckResult) SetPreCheckName(v string) *DescribePreCheckResultResponseBodyPreCheckResult {
	s.PreCheckName = &v
	return s
}

func (s *DescribePreCheckResultResponseBodyPreCheckResult) SetState(v string) *DescribePreCheckResultResponseBodyPreCheckResult {
	s.State = &v
	return s
}

func (s *DescribePreCheckResultResponseBodyPreCheckResult) SetSubCheckItems(v []*DescribePreCheckResultResponseBodyPreCheckResultSubCheckItems) *DescribePreCheckResultResponseBodyPreCheckResult {
	s.SubCheckItems = v
	return s
}

type DescribePreCheckResultResponseBodyPreCheckResultSubCheckItems struct {
	ErrorMsgCode     *string   `json:"ErrorMsgCode,omitempty" xml:"ErrorMsgCode,omitempty"`
	PreCheckItemName *string   `json:"PreCheckItemName,omitempty" xml:"PreCheckItemName,omitempty"`
	State            *string   `json:"State,omitempty" xml:"State,omitempty"`
	ErrorMsgParams   []*string `json:"ErrorMsgParams,omitempty" xml:"ErrorMsgParams,omitempty" type:"Repeated"`
}

func (s DescribePreCheckResultResponseBodyPreCheckResultSubCheckItems) String() string {
	return tea.Prettify(s)
}

func (s DescribePreCheckResultResponseBodyPreCheckResultSubCheckItems) GoString() string {
	return s.String()
}

func (s *DescribePreCheckResultResponseBodyPreCheckResultSubCheckItems) SetErrorMsgCode(v string) *DescribePreCheckResultResponseBodyPreCheckResultSubCheckItems {
	s.ErrorMsgCode = &v
	return s
}

func (s *DescribePreCheckResultResponseBodyPreCheckResultSubCheckItems) SetPreCheckItemName(v string) *DescribePreCheckResultResponseBodyPreCheckResultSubCheckItems {
	s.PreCheckItemName = &v
	return s
}

func (s *DescribePreCheckResultResponseBodyPreCheckResultSubCheckItems) SetState(v string) *DescribePreCheckResultResponseBodyPreCheckResultSubCheckItems {
	s.State = &v
	return s
}

func (s *DescribePreCheckResultResponseBodyPreCheckResultSubCheckItems) SetErrorMsgParams(v []*string) *DescribePreCheckResultResponseBodyPreCheckResultSubCheckItems {
	s.ErrorMsgParams = v
	return s
}

type DescribePreCheckResultResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribePreCheckResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePreCheckResultResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePreCheckResultResponse) GoString() string {
	return s.String()
}

func (s *DescribePreCheckResultResponse) SetHeaders(v map[string]*string) *DescribePreCheckResultResponse {
	s.Headers = v
	return s
}

func (s *DescribePreCheckResultResponse) SetBody(v *DescribePreCheckResultResponseBody) *DescribePreCheckResultResponse {
	s.Body = v
	return s
}

type DescribeRdsCommodityRequest struct {
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	CommodityCode  *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	OrderType      *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
}

func (s DescribeRdsCommodityRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsCommodityRequest) GoString() string {
	return s.String()
}

func (s *DescribeRdsCommodityRequest) SetDrdsInstanceId(v string) *DescribeRdsCommodityRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeRdsCommodityRequest) SetCommodityCode(v string) *DescribeRdsCommodityRequest {
	s.CommodityCode = &v
	return s
}

func (s *DescribeRdsCommodityRequest) SetOrderType(v string) *DescribeRdsCommodityRequest {
	s.OrderType = &v
	return s
}

type DescribeRdsCommodityResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeRdsCommodityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsCommodityResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRdsCommodityResponseBody) SetData(v string) *DescribeRdsCommodityResponseBody {
	s.Data = &v
	return s
}

func (s *DescribeRdsCommodityResponseBody) SetRequestId(v string) *DescribeRdsCommodityResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRdsCommodityResponseBody) SetSuccess(v bool) *DescribeRdsCommodityResponseBody {
	s.Success = &v
	return s
}

type DescribeRdsCommodityResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRdsCommodityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRdsCommodityResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsCommodityResponse) GoString() string {
	return s.String()
}

func (s *DescribeRdsCommodityResponse) SetHeaders(v map[string]*string) *DescribeRdsCommodityResponse {
	s.Headers = v
	return s
}

func (s *DescribeRdsCommodityResponse) SetBody(v *DescribeRdsCommodityResponseBody) *DescribeRdsCommodityResponse {
	s.Body = v
	return s
}

type DescribeRDSPerformanceRequest struct {
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	RdsInstanceId  *string `json:"RdsInstanceId,omitempty" xml:"RdsInstanceId,omitempty"`
	Keys           *string `json:"Keys,omitempty" xml:"Keys,omitempty"`
	StartTime      *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime        *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	DbInstType     *string `json:"DbInstType,omitempty" xml:"DbInstType,omitempty"`
}

func (s DescribeRDSPerformanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRDSPerformanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeRDSPerformanceRequest) SetDrdsInstanceId(v string) *DescribeRDSPerformanceRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeRDSPerformanceRequest) SetRdsInstanceId(v string) *DescribeRDSPerformanceRequest {
	s.RdsInstanceId = &v
	return s
}

func (s *DescribeRDSPerformanceRequest) SetKeys(v string) *DescribeRDSPerformanceRequest {
	s.Keys = &v
	return s
}

func (s *DescribeRDSPerformanceRequest) SetStartTime(v int64) *DescribeRDSPerformanceRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeRDSPerformanceRequest) SetEndTime(v int64) *DescribeRDSPerformanceRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeRDSPerformanceRequest) SetDbInstType(v string) *DescribeRDSPerformanceRequest {
	s.DbInstType = &v
	return s
}

type DescribeRDSPerformanceResponseBody struct {
	Success   *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      []*DescribeRDSPerformanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s DescribeRDSPerformanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRDSPerformanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRDSPerformanceResponseBody) SetSuccess(v bool) *DescribeRDSPerformanceResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeRDSPerformanceResponseBody) SetRequestId(v string) *DescribeRDSPerformanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRDSPerformanceResponseBody) SetData(v []*DescribeRDSPerformanceResponseBodyData) *DescribeRDSPerformanceResponseBody {
	s.Data = v
	return s
}

type DescribeRDSPerformanceResponseBodyData struct {
	Key      *string                                         `json:"Key,omitempty" xml:"Key,omitempty"`
	NodeName *string                                         `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	Unit     *string                                         `json:"Unit,omitempty" xml:"Unit,omitempty"`
	NodeNum  *int32                                          `json:"NodeNum,omitempty" xml:"NodeNum,omitempty"`
	Values   []*DescribeRDSPerformanceResponseBodyDataValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s DescribeRDSPerformanceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeRDSPerformanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeRDSPerformanceResponseBodyData) SetKey(v string) *DescribeRDSPerformanceResponseBodyData {
	s.Key = &v
	return s
}

func (s *DescribeRDSPerformanceResponseBodyData) SetNodeName(v string) *DescribeRDSPerformanceResponseBodyData {
	s.NodeName = &v
	return s
}

func (s *DescribeRDSPerformanceResponseBodyData) SetUnit(v string) *DescribeRDSPerformanceResponseBodyData {
	s.Unit = &v
	return s
}

func (s *DescribeRDSPerformanceResponseBodyData) SetNodeNum(v int32) *DescribeRDSPerformanceResponseBodyData {
	s.NodeNum = &v
	return s
}

func (s *DescribeRDSPerformanceResponseBodyData) SetValues(v []*DescribeRDSPerformanceResponseBodyDataValues) *DescribeRDSPerformanceResponseBodyData {
	s.Values = v
	return s
}

type DescribeRDSPerformanceResponseBodyDataValues struct {
	Date  *int64  `json:"Date,omitempty" xml:"Date,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeRDSPerformanceResponseBodyDataValues) String() string {
	return tea.Prettify(s)
}

func (s DescribeRDSPerformanceResponseBodyDataValues) GoString() string {
	return s.String()
}

func (s *DescribeRDSPerformanceResponseBodyDataValues) SetDate(v int64) *DescribeRDSPerformanceResponseBodyDataValues {
	s.Date = &v
	return s
}

func (s *DescribeRDSPerformanceResponseBodyDataValues) SetValue(v string) *DescribeRDSPerformanceResponseBodyDataValues {
	s.Value = &v
	return s
}

type DescribeRDSPerformanceResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRDSPerformanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRDSPerformanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRDSPerformanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeRDSPerformanceResponse) SetHeaders(v map[string]*string) *DescribeRDSPerformanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeRDSPerformanceResponse) SetBody(v *DescribeRDSPerformanceResponseBody) *DescribeRDSPerformanceResponse {
	s.Body = v
	return s
}

type DescribeRdsPerformanceSummaryRequest struct {
	RegionId       *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DrdsInstanceId *string   `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	RdsInstanceId  []*string `json:"RdsInstanceId,omitempty" xml:"RdsInstanceId,omitempty" type:"Repeated"`
}

func (s DescribeRdsPerformanceSummaryRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsPerformanceSummaryRequest) GoString() string {
	return s.String()
}

func (s *DescribeRdsPerformanceSummaryRequest) SetRegionId(v string) *DescribeRdsPerformanceSummaryRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRdsPerformanceSummaryRequest) SetDrdsInstanceId(v string) *DescribeRdsPerformanceSummaryRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeRdsPerformanceSummaryRequest) SetRdsInstanceId(v []*string) *DescribeRdsPerformanceSummaryRequest {
	s.RdsInstanceId = v
	return s
}

type DescribeRdsPerformanceSummaryResponseBody struct {
	Success             *bool                                                           `json:"Success,omitempty" xml:"Success,omitempty"`
	RequestId           *string                                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RdsPerformanceInfos []*DescribeRdsPerformanceSummaryResponseBodyRdsPerformanceInfos `json:"RdsPerformanceInfos,omitempty" xml:"RdsPerformanceInfos,omitempty" type:"Repeated"`
}

func (s DescribeRdsPerformanceSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsPerformanceSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRdsPerformanceSummaryResponseBody) SetSuccess(v bool) *DescribeRdsPerformanceSummaryResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeRdsPerformanceSummaryResponseBody) SetRequestId(v string) *DescribeRdsPerformanceSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRdsPerformanceSummaryResponseBody) SetRdsPerformanceInfos(v []*DescribeRdsPerformanceSummaryResponseBodyRdsPerformanceInfos) *DescribeRdsPerformanceSummaryResponseBody {
	s.RdsPerformanceInfos = v
	return s
}

type DescribeRdsPerformanceSummaryResponseBodyRdsPerformanceInfos struct {
	Cpu            *float32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	ActiveSessions *int32   `json:"ActiveSessions,omitempty" xml:"ActiveSessions,omitempty"`
	TotalSessions  *int32   `json:"TotalSessions,omitempty" xml:"TotalSessions,omitempty"`
	RdsId          *string  `json:"RdsId,omitempty" xml:"RdsId,omitempty"`
	Iops           *float32 `json:"Iops,omitempty" xml:"Iops,omitempty"`
	SpaceUsage     *int64   `json:"SpaceUsage,omitempty" xml:"SpaceUsage,omitempty"`
}

func (s DescribeRdsPerformanceSummaryResponseBodyRdsPerformanceInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsPerformanceSummaryResponseBodyRdsPerformanceInfos) GoString() string {
	return s.String()
}

func (s *DescribeRdsPerformanceSummaryResponseBodyRdsPerformanceInfos) SetCpu(v float32) *DescribeRdsPerformanceSummaryResponseBodyRdsPerformanceInfos {
	s.Cpu = &v
	return s
}

func (s *DescribeRdsPerformanceSummaryResponseBodyRdsPerformanceInfos) SetActiveSessions(v int32) *DescribeRdsPerformanceSummaryResponseBodyRdsPerformanceInfos {
	s.ActiveSessions = &v
	return s
}

func (s *DescribeRdsPerformanceSummaryResponseBodyRdsPerformanceInfos) SetTotalSessions(v int32) *DescribeRdsPerformanceSummaryResponseBodyRdsPerformanceInfos {
	s.TotalSessions = &v
	return s
}

func (s *DescribeRdsPerformanceSummaryResponseBodyRdsPerformanceInfos) SetRdsId(v string) *DescribeRdsPerformanceSummaryResponseBodyRdsPerformanceInfos {
	s.RdsId = &v
	return s
}

func (s *DescribeRdsPerformanceSummaryResponseBodyRdsPerformanceInfos) SetIops(v float32) *DescribeRdsPerformanceSummaryResponseBodyRdsPerformanceInfos {
	s.Iops = &v
	return s
}

func (s *DescribeRdsPerformanceSummaryResponseBodyRdsPerformanceInfos) SetSpaceUsage(v int64) *DescribeRdsPerformanceSummaryResponseBodyRdsPerformanceInfos {
	s.SpaceUsage = &v
	return s
}

type DescribeRdsPerformanceSummaryResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRdsPerformanceSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRdsPerformanceSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsPerformanceSummaryResponse) GoString() string {
	return s.String()
}

func (s *DescribeRdsPerformanceSummaryResponse) SetHeaders(v map[string]*string) *DescribeRdsPerformanceSummaryResponse {
	s.Headers = v
	return s
}

func (s *DescribeRdsPerformanceSummaryResponse) SetBody(v *DescribeRdsPerformanceSummaryResponseBody) *DescribeRdsPerformanceSummaryResponse {
	s.Body = v
	return s
}

type DescribeRdsSuperAccountInstancesRequest struct {
	DrdsInstanceId *string   `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	DbInstType     *string   `json:"DbInstType,omitempty" xml:"DbInstType,omitempty"`
	RdsInstance    []*string `json:"RdsInstance,omitempty" xml:"RdsInstance,omitempty" type:"Repeated"`
}

func (s DescribeRdsSuperAccountInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsSuperAccountInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeRdsSuperAccountInstancesRequest) SetDrdsInstanceId(v string) *DescribeRdsSuperAccountInstancesRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeRdsSuperAccountInstancesRequest) SetDbInstType(v string) *DescribeRdsSuperAccountInstancesRequest {
	s.DbInstType = &v
	return s
}

func (s *DescribeRdsSuperAccountInstancesRequest) SetRdsInstance(v []*string) *DescribeRdsSuperAccountInstancesRequest {
	s.RdsInstance = v
	return s
}

type DescribeRdsSuperAccountInstancesResponseBody struct {
	RequestId   *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DbInstances *DescribeRdsSuperAccountInstancesResponseBodyDbInstances `json:"DbInstances,omitempty" xml:"DbInstances,omitempty" type:"Struct"`
}

func (s DescribeRdsSuperAccountInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsSuperAccountInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRdsSuperAccountInstancesResponseBody) SetRequestId(v string) *DescribeRdsSuperAccountInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRdsSuperAccountInstancesResponseBody) SetDbInstances(v *DescribeRdsSuperAccountInstancesResponseBodyDbInstances) *DescribeRdsSuperAccountInstancesResponseBody {
	s.DbInstances = v
	return s
}

type DescribeRdsSuperAccountInstancesResponseBodyDbInstances struct {
	DbInstance []*string `json:"DbInstance,omitempty" xml:"DbInstance,omitempty" type:"Repeated"`
}

func (s DescribeRdsSuperAccountInstancesResponseBodyDbInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsSuperAccountInstancesResponseBodyDbInstances) GoString() string {
	return s.String()
}

func (s *DescribeRdsSuperAccountInstancesResponseBodyDbInstances) SetDbInstance(v []*string) *DescribeRdsSuperAccountInstancesResponseBodyDbInstances {
	s.DbInstance = v
	return s
}

type DescribeRdsSuperAccountInstancesResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRdsSuperAccountInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRdsSuperAccountInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsSuperAccountInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeRdsSuperAccountInstancesResponse) SetHeaders(v map[string]*string) *DescribeRdsSuperAccountInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeRdsSuperAccountInstancesResponse) SetBody(v *DescribeRdsSuperAccountInstancesResponseBody) *DescribeRdsSuperAccountInstancesResponse {
	s.Body = v
	return s
}

type DescribeRestoreOrderRequest struct {
	DrdsInstanceId      *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	PreferredBackupTime *string `json:"PreferredBackupTime,omitempty" xml:"PreferredBackupTime,omitempty"`
	BackupMode          *string `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	BackupLevel         *string `json:"BackupLevel,omitempty" xml:"BackupLevel,omitempty"`
	BackupDbNames       *string `json:"BackupDbNames,omitempty" xml:"BackupDbNames,omitempty"`
	BackupId            *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
}

func (s DescribeRestoreOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreOrderRequest) GoString() string {
	return s.String()
}

func (s *DescribeRestoreOrderRequest) SetDrdsInstanceId(v string) *DescribeRestoreOrderRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeRestoreOrderRequest) SetPreferredBackupTime(v string) *DescribeRestoreOrderRequest {
	s.PreferredBackupTime = &v
	return s
}

func (s *DescribeRestoreOrderRequest) SetBackupMode(v string) *DescribeRestoreOrderRequest {
	s.BackupMode = &v
	return s
}

func (s *DescribeRestoreOrderRequest) SetBackupLevel(v string) *DescribeRestoreOrderRequest {
	s.BackupLevel = &v
	return s
}

func (s *DescribeRestoreOrderRequest) SetBackupDbNames(v string) *DescribeRestoreOrderRequest {
	s.BackupDbNames = &v
	return s
}

func (s *DescribeRestoreOrderRequest) SetBackupId(v string) *DescribeRestoreOrderRequest {
	s.BackupId = &v
	return s
}

type DescribeRestoreOrderResponseBody struct {
	Success        *bool                                           `json:"Success,omitempty" xml:"Success,omitempty"`
	RequestId      *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RestoreOrderDO *DescribeRestoreOrderResponseBodyRestoreOrderDO `json:"RestoreOrderDO,omitempty" xml:"RestoreOrderDO,omitempty" type:"Struct"`
}

func (s DescribeRestoreOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreOrderResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRestoreOrderResponseBody) SetSuccess(v bool) *DescribeRestoreOrderResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeRestoreOrderResponseBody) SetRequestId(v string) *DescribeRestoreOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRestoreOrderResponseBody) SetRestoreOrderDO(v *DescribeRestoreOrderResponseBodyRestoreOrderDO) *DescribeRestoreOrderResponseBody {
	s.RestoreOrderDO = v
	return s
}

type DescribeRestoreOrderResponseBodyRestoreOrderDO struct {
	DrdsOrderDOList  *DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOList  `json:"DrdsOrderDOList,omitempty" xml:"DrdsOrderDOList,omitempty" type:"Struct"`
	RdsOrderDOList   *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOList   `json:"RdsOrderDOList,omitempty" xml:"RdsOrderDOList,omitempty" type:"Struct"`
	PolarOrderDOList *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOList `json:"PolarOrderDOList,omitempty" xml:"PolarOrderDOList,omitempty" type:"Struct"`
}

func (s DescribeRestoreOrderResponseBodyRestoreOrderDO) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreOrderResponseBodyRestoreOrderDO) GoString() string {
	return s.String()
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDO) SetDrdsOrderDOList(v *DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOList) *DescribeRestoreOrderResponseBodyRestoreOrderDO {
	s.DrdsOrderDOList = v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDO) SetRdsOrderDOList(v *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOList) *DescribeRestoreOrderResponseBodyRestoreOrderDO {
	s.RdsOrderDOList = v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDO) SetPolarOrderDOList(v *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOList) *DescribeRestoreOrderResponseBodyRestoreOrderDO {
	s.PolarOrderDOList = v
	return s
}

type DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOList struct {
	DrdsOrderDOList []*DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOListDrdsOrderDOList `json:"DrdsOrderDOList,omitempty" xml:"DrdsOrderDOList,omitempty" type:"Repeated"`
}

func (s DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOList) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOList) GoString() string {
	return s.String()
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOList) SetDrdsOrderDOList(v []*DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOListDrdsOrderDOList) *DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOList {
	s.DrdsOrderDOList = v
	return s
}

type DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOListDrdsOrderDOList struct {
	VpcId     *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	Network   *string `json:"Network,omitempty" xml:"Network,omitempty"`
	VSwtichId *string `json:"VSwtichId,omitempty" xml:"VSwtichId,omitempty"`
	InstSpec  *string `json:"InstSpec,omitempty" xml:"InstSpec,omitempty"`
	AzoneId   *string `json:"AzoneId,omitempty" xml:"AzoneId,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOListDrdsOrderDOList) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOListDrdsOrderDOList) GoString() string {
	return s.String()
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOListDrdsOrderDOList) SetVpcId(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOListDrdsOrderDOList {
	s.VpcId = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOListDrdsOrderDOList) SetNetwork(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOListDrdsOrderDOList {
	s.Network = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOListDrdsOrderDOList) SetVSwtichId(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOListDrdsOrderDOList {
	s.VSwtichId = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOListDrdsOrderDOList) SetInstSpec(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOListDrdsOrderDOList {
	s.InstSpec = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOListDrdsOrderDOList) SetAzoneId(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOListDrdsOrderDOList {
	s.AzoneId = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOListDrdsOrderDOList) SetRegionId(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOListDrdsOrderDOList {
	s.RegionId = &v
	return s
}

type DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOList struct {
	RdsOrderDOList []*DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList `json:"RdsOrderDOList,omitempty" xml:"RdsOrderDOList,omitempty" type:"Repeated"`
}

func (s DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOList) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOList) GoString() string {
	return s.String()
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOList) SetRdsOrderDOList(v []*DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList) *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOList {
	s.RdsOrderDOList = v
	return s
}

type DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList struct {
	Network           *string `json:"Network,omitempty" xml:"Network,omitempty"`
	Version           *string `json:"Version,omitempty" xml:"Version,omitempty"`
	InstanceClass     *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	DbInstanceStorage *string `json:"DbInstanceStorage,omitempty" xml:"DbInstanceStorage,omitempty"`
	Num               *int64  `json:"Num,omitempty" xml:"Num,omitempty"`
	Engine            *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	AzoneId           *string `json:"AzoneId,omitempty" xml:"AzoneId,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList) GoString() string {
	return s.String()
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList) SetNetwork(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList {
	s.Network = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList) SetVersion(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList {
	s.Version = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList) SetInstanceClass(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList {
	s.InstanceClass = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList) SetDbInstanceStorage(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList {
	s.DbInstanceStorage = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList) SetNum(v int64) *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList {
	s.Num = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList) SetEngine(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList {
	s.Engine = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList) SetAzoneId(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList {
	s.AzoneId = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList) SetRegionId(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList {
	s.RegionId = &v
	return s
}

type DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOList struct {
	PolarOrderDOList []*DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList `json:"PolarOrderDOList,omitempty" xml:"PolarOrderDOList,omitempty" type:"Repeated"`
}

func (s DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOList) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOList) GoString() string {
	return s.String()
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOList) SetPolarOrderDOList(v []*DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList) *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOList {
	s.PolarOrderDOList = v
	return s
}

type DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList struct {
	Network           *string `json:"Network,omitempty" xml:"Network,omitempty"`
	Version           *string `json:"Version,omitempty" xml:"Version,omitempty"`
	InstanceClass     *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	DbInstanceStorage *string `json:"DbInstanceStorage,omitempty" xml:"DbInstanceStorage,omitempty"`
	Num               *int64  `json:"Num,omitempty" xml:"Num,omitempty"`
	Engine            *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	AzoneId           *string `json:"AzoneId,omitempty" xml:"AzoneId,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList) GoString() string {
	return s.String()
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList) SetNetwork(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList {
	s.Network = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList) SetVersion(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList {
	s.Version = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList) SetInstanceClass(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList {
	s.InstanceClass = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList) SetDbInstanceStorage(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList {
	s.DbInstanceStorage = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList) SetNum(v int64) *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList {
	s.Num = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList) SetEngine(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList {
	s.Engine = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList) SetAzoneId(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList {
	s.AzoneId = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList) SetRegionId(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList {
	s.RegionId = &v
	return s
}

type DescribeRestoreOrderResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRestoreOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRestoreOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreOrderResponse) GoString() string {
	return s.String()
}

func (s *DescribeRestoreOrderResponse) SetHeaders(v map[string]*string) *DescribeRestoreOrderResponse {
	s.Headers = v
	return s
}

func (s *DescribeRestoreOrderResponse) SetBody(v *DescribeRestoreOrderResponseBody) *DescribeRestoreOrderResponse {
	s.Body = v
	return s
}

type DescribeShardTaskInfoRequest struct {
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DrdsInstanceId  *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	DbName          *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	SourceTableName *string `json:"SourceTableName,omitempty" xml:"SourceTableName,omitempty"`
	TargetTableName *string `json:"TargetTableName,omitempty" xml:"TargetTableName,omitempty"`
}

func (s DescribeShardTaskInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeShardTaskInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeShardTaskInfoRequest) SetRegionId(v string) *DescribeShardTaskInfoRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeShardTaskInfoRequest) SetDrdsInstanceId(v string) *DescribeShardTaskInfoRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeShardTaskInfoRequest) SetDbName(v string) *DescribeShardTaskInfoRequest {
	s.DbName = &v
	return s
}

func (s *DescribeShardTaskInfoRequest) SetSourceTableName(v string) *DescribeShardTaskInfoRequest {
	s.SourceTableName = &v
	return s
}

func (s *DescribeShardTaskInfoRequest) SetTargetTableName(v string) *DescribeShardTaskInfoRequest {
	s.TargetTableName = &v
	return s
}

type DescribeShardTaskInfoResponseBody struct {
	Success   *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *DescribeShardTaskInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s DescribeShardTaskInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeShardTaskInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeShardTaskInfoResponseBody) SetSuccess(v bool) *DescribeShardTaskInfoResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBody) SetRequestId(v string) *DescribeShardTaskInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBody) SetData(v *DescribeShardTaskInfoResponseBodyData) *DescribeShardTaskInfoResponseBody {
	s.Data = v
	return s
}

type DescribeShardTaskInfoResponseBodyData struct {
	Status          *string                                          `json:"Status,omitempty" xml:"Status,omitempty"`
	Stage           *string                                          `json:"Stage,omitempty" xml:"Stage,omitempty"`
	Progress        *string                                          `json:"Progress,omitempty" xml:"Progress,omitempty"`
	Expired         *string                                          `json:"Expired,omitempty" xml:"Expired,omitempty"`
	TargetTableName *string                                          `json:"TargetTableName,omitempty" xml:"TargetTableName,omitempty"`
	SourceTableName *string                                          `json:"SourceTableName,omitempty" xml:"SourceTableName,omitempty"`
	Full            *DescribeShardTaskInfoResponseBodyDataFull       `json:"Full,omitempty" xml:"Full,omitempty" type:"Struct"`
	FullCheck       *DescribeShardTaskInfoResponseBodyDataFullCheck  `json:"FullCheck,omitempty" xml:"FullCheck,omitempty" type:"Struct"`
	FullRevise      *DescribeShardTaskInfoResponseBodyDataFullRevise `json:"FullRevise,omitempty" xml:"FullRevise,omitempty" type:"Struct"`
	Review          *DescribeShardTaskInfoResponseBodyDataReview     `json:"Review,omitempty" xml:"Review,omitempty" type:"Struct"`
	Increment       *DescribeShardTaskInfoResponseBodyDataIncrement  `json:"Increment,omitempty" xml:"Increment,omitempty" type:"Struct"`
}

func (s DescribeShardTaskInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeShardTaskInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeShardTaskInfoResponseBodyData) SetStatus(v string) *DescribeShardTaskInfoResponseBodyData {
	s.Status = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyData) SetStage(v string) *DescribeShardTaskInfoResponseBodyData {
	s.Stage = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyData) SetProgress(v string) *DescribeShardTaskInfoResponseBodyData {
	s.Progress = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyData) SetExpired(v string) *DescribeShardTaskInfoResponseBodyData {
	s.Expired = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyData) SetTargetTableName(v string) *DescribeShardTaskInfoResponseBodyData {
	s.TargetTableName = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyData) SetSourceTableName(v string) *DescribeShardTaskInfoResponseBodyData {
	s.SourceTableName = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyData) SetFull(v *DescribeShardTaskInfoResponseBodyDataFull) *DescribeShardTaskInfoResponseBodyData {
	s.Full = v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyData) SetFullCheck(v *DescribeShardTaskInfoResponseBodyDataFullCheck) *DescribeShardTaskInfoResponseBodyData {
	s.FullCheck = v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyData) SetFullRevise(v *DescribeShardTaskInfoResponseBodyDataFullRevise) *DescribeShardTaskInfoResponseBodyData {
	s.FullRevise = v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyData) SetReview(v *DescribeShardTaskInfoResponseBodyDataReview) *DescribeShardTaskInfoResponseBodyData {
	s.Review = v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyData) SetIncrement(v *DescribeShardTaskInfoResponseBodyDataIncrement) *DescribeShardTaskInfoResponseBodyData {
	s.Increment = v
	return s
}

type DescribeShardTaskInfoResponseBodyDataFull struct {
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Progress  *int32  `json:"Progress,omitempty" xml:"Progress,omitempty"`
	Tps       *int32  `json:"Tps,omitempty" xml:"Tps,omitempty"`
	Total     *int32  `json:"Total,omitempty" xml:"Total,omitempty"`
	Expired   *int32  `json:"Expired,omitempty" xml:"Expired,omitempty"`
}

func (s DescribeShardTaskInfoResponseBodyDataFull) String() string {
	return tea.Prettify(s)
}

func (s DescribeShardTaskInfoResponseBodyDataFull) GoString() string {
	return s.String()
}

func (s *DescribeShardTaskInfoResponseBodyDataFull) SetStartTime(v string) *DescribeShardTaskInfoResponseBodyDataFull {
	s.StartTime = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyDataFull) SetProgress(v int32) *DescribeShardTaskInfoResponseBodyDataFull {
	s.Progress = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyDataFull) SetTps(v int32) *DescribeShardTaskInfoResponseBodyDataFull {
	s.Tps = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyDataFull) SetTotal(v int32) *DescribeShardTaskInfoResponseBodyDataFull {
	s.Total = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyDataFull) SetExpired(v int32) *DescribeShardTaskInfoResponseBodyDataFull {
	s.Expired = &v
	return s
}

type DescribeShardTaskInfoResponseBodyDataFullCheck struct {
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Progress  *int32  `json:"Progress,omitempty" xml:"Progress,omitempty"`
	Tps       *int32  `json:"Tps,omitempty" xml:"Tps,omitempty"`
	Total     *int32  `json:"Total,omitempty" xml:"Total,omitempty"`
	Expired   *int32  `json:"Expired,omitempty" xml:"Expired,omitempty"`
}

func (s DescribeShardTaskInfoResponseBodyDataFullCheck) String() string {
	return tea.Prettify(s)
}

func (s DescribeShardTaskInfoResponseBodyDataFullCheck) GoString() string {
	return s.String()
}

func (s *DescribeShardTaskInfoResponseBodyDataFullCheck) SetStartTime(v string) *DescribeShardTaskInfoResponseBodyDataFullCheck {
	s.StartTime = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyDataFullCheck) SetProgress(v int32) *DescribeShardTaskInfoResponseBodyDataFullCheck {
	s.Progress = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyDataFullCheck) SetTps(v int32) *DescribeShardTaskInfoResponseBodyDataFullCheck {
	s.Tps = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyDataFullCheck) SetTotal(v int32) *DescribeShardTaskInfoResponseBodyDataFullCheck {
	s.Total = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyDataFullCheck) SetExpired(v int32) *DescribeShardTaskInfoResponseBodyDataFullCheck {
	s.Expired = &v
	return s
}

type DescribeShardTaskInfoResponseBodyDataFullRevise struct {
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Progress  *int32  `json:"Progress,omitempty" xml:"Progress,omitempty"`
	Tps       *int32  `json:"Tps,omitempty" xml:"Tps,omitempty"`
	Total     *int32  `json:"Total,omitempty" xml:"Total,omitempty"`
	Expired   *int32  `json:"Expired,omitempty" xml:"Expired,omitempty"`
}

func (s DescribeShardTaskInfoResponseBodyDataFullRevise) String() string {
	return tea.Prettify(s)
}

func (s DescribeShardTaskInfoResponseBodyDataFullRevise) GoString() string {
	return s.String()
}

func (s *DescribeShardTaskInfoResponseBodyDataFullRevise) SetStartTime(v string) *DescribeShardTaskInfoResponseBodyDataFullRevise {
	s.StartTime = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyDataFullRevise) SetProgress(v int32) *DescribeShardTaskInfoResponseBodyDataFullRevise {
	s.Progress = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyDataFullRevise) SetTps(v int32) *DescribeShardTaskInfoResponseBodyDataFullRevise {
	s.Tps = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyDataFullRevise) SetTotal(v int32) *DescribeShardTaskInfoResponseBodyDataFullRevise {
	s.Total = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyDataFullRevise) SetExpired(v int32) *DescribeShardTaskInfoResponseBodyDataFullRevise {
	s.Expired = &v
	return s
}

type DescribeShardTaskInfoResponseBodyDataReview struct {
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Progress  *int32  `json:"Progress,omitempty" xml:"Progress,omitempty"`
	Tps       *int32  `json:"Tps,omitempty" xml:"Tps,omitempty"`
	Total     *int32  `json:"Total,omitempty" xml:"Total,omitempty"`
	Expired   *int32  `json:"Expired,omitempty" xml:"Expired,omitempty"`
}

func (s DescribeShardTaskInfoResponseBodyDataReview) String() string {
	return tea.Prettify(s)
}

func (s DescribeShardTaskInfoResponseBodyDataReview) GoString() string {
	return s.String()
}

func (s *DescribeShardTaskInfoResponseBodyDataReview) SetStartTime(v string) *DescribeShardTaskInfoResponseBodyDataReview {
	s.StartTime = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyDataReview) SetProgress(v int32) *DescribeShardTaskInfoResponseBodyDataReview {
	s.Progress = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyDataReview) SetTps(v int32) *DescribeShardTaskInfoResponseBodyDataReview {
	s.Tps = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyDataReview) SetTotal(v int32) *DescribeShardTaskInfoResponseBodyDataReview {
	s.Total = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyDataReview) SetExpired(v int32) *DescribeShardTaskInfoResponseBodyDataReview {
	s.Expired = &v
	return s
}

type DescribeShardTaskInfoResponseBodyDataIncrement struct {
	Delay     *int32  `json:"Delay,omitempty" xml:"Delay,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Tps       *int32  `json:"Tps,omitempty" xml:"Tps,omitempty"`
}

func (s DescribeShardTaskInfoResponseBodyDataIncrement) String() string {
	return tea.Prettify(s)
}

func (s DescribeShardTaskInfoResponseBodyDataIncrement) GoString() string {
	return s.String()
}

func (s *DescribeShardTaskInfoResponseBodyDataIncrement) SetDelay(v int32) *DescribeShardTaskInfoResponseBodyDataIncrement {
	s.Delay = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyDataIncrement) SetStartTime(v string) *DescribeShardTaskInfoResponseBodyDataIncrement {
	s.StartTime = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyDataIncrement) SetTps(v int32) *DescribeShardTaskInfoResponseBodyDataIncrement {
	s.Tps = &v
	return s
}

type DescribeShardTaskInfoResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeShardTaskInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeShardTaskInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeShardTaskInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeShardTaskInfoResponse) SetHeaders(v map[string]*string) *DescribeShardTaskInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeShardTaskInfoResponse) SetBody(v *DescribeShardTaskInfoResponseBody) *DescribeShardTaskInfoResponse {
	s.Body = v
	return s
}

type DescribeShardTaskListRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	Query          *string `json:"Query,omitempty" xml:"Query,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentPage    *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	TaskType       *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DescribeShardTaskListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeShardTaskListRequest) GoString() string {
	return s.String()
}

func (s *DescribeShardTaskListRequest) SetRegionId(v string) *DescribeShardTaskListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeShardTaskListRequest) SetDrdsInstanceId(v string) *DescribeShardTaskListRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeShardTaskListRequest) SetDbName(v string) *DescribeShardTaskListRequest {
	s.DbName = &v
	return s
}

func (s *DescribeShardTaskListRequest) SetQuery(v string) *DescribeShardTaskListRequest {
	s.Query = &v
	return s
}

func (s *DescribeShardTaskListRequest) SetPageSize(v int32) *DescribeShardTaskListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeShardTaskListRequest) SetCurrentPage(v int32) *DescribeShardTaskListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeShardTaskListRequest) SetTaskType(v string) *DescribeShardTaskListRequest {
	s.TaskType = &v
	return s
}

type DescribeShardTaskListResponseBody struct {
	RequestId  *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
	PageNumber *int32                                   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total      *int32                                   `json:"Total,omitempty" xml:"Total,omitempty"`
	List       []*DescribeShardTaskListResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s DescribeShardTaskListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeShardTaskListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeShardTaskListResponseBody) SetRequestId(v string) *DescribeShardTaskListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeShardTaskListResponseBody) SetSuccess(v bool) *DescribeShardTaskListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeShardTaskListResponseBody) SetPageNumber(v int32) *DescribeShardTaskListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeShardTaskListResponseBody) SetPageSize(v int32) *DescribeShardTaskListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeShardTaskListResponseBody) SetTotal(v int32) *DescribeShardTaskListResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeShardTaskListResponseBody) SetList(v []*DescribeShardTaskListResponseBodyList) *DescribeShardTaskListResponseBody {
	s.List = v
	return s
}

type DescribeShardTaskListResponseBodyList struct {
	SourceTableName *string `json:"SourceTableName,omitempty" xml:"SourceTableName,omitempty"`
	TargetTableName *string `json:"TargetTableName,omitempty" xml:"TargetTableName,omitempty"`
	Expired         *int64  `json:"Expired,omitempty" xml:"Expired,omitempty"`
	Stage           *int32  `json:"Stage,omitempty" xml:"Stage,omitempty"`
	Progress        *int32  `json:"Progress,omitempty" xml:"Progress,omitempty"`
	Delay           *int32  `json:"Delay,omitempty" xml:"Delay,omitempty"`
}

func (s DescribeShardTaskListResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s DescribeShardTaskListResponseBodyList) GoString() string {
	return s.String()
}

func (s *DescribeShardTaskListResponseBodyList) SetSourceTableName(v string) *DescribeShardTaskListResponseBodyList {
	s.SourceTableName = &v
	return s
}

func (s *DescribeShardTaskListResponseBodyList) SetTargetTableName(v string) *DescribeShardTaskListResponseBodyList {
	s.TargetTableName = &v
	return s
}

func (s *DescribeShardTaskListResponseBodyList) SetExpired(v int64) *DescribeShardTaskListResponseBodyList {
	s.Expired = &v
	return s
}

func (s *DescribeShardTaskListResponseBodyList) SetStage(v int32) *DescribeShardTaskListResponseBodyList {
	s.Stage = &v
	return s
}

func (s *DescribeShardTaskListResponseBodyList) SetProgress(v int32) *DescribeShardTaskListResponseBodyList {
	s.Progress = &v
	return s
}

func (s *DescribeShardTaskListResponseBodyList) SetDelay(v int32) *DescribeShardTaskListResponseBodyList {
	s.Delay = &v
	return s
}

type DescribeShardTaskListResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeShardTaskListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeShardTaskListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeShardTaskListResponse) GoString() string {
	return s.String()
}

func (s *DescribeShardTaskListResponse) SetHeaders(v map[string]*string) *DescribeShardTaskListResponse {
	s.Headers = v
	return s
}

func (s *DescribeShardTaskListResponse) SetBody(v *DescribeShardTaskListResponseBody) *DescribeShardTaskListResponse {
	s.Body = v
	return s
}

type DescribeSqlFlashbakTaskRequest struct {
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s DescribeSqlFlashbakTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSqlFlashbakTaskRequest) GoString() string {
	return s.String()
}

func (s *DescribeSqlFlashbakTaskRequest) SetDrdsInstanceId(v string) *DescribeSqlFlashbakTaskRequest {
	s.DrdsInstanceId = &v
	return s
}

type DescribeSqlFlashbakTaskResponseBody struct {
	Success           *bool                                                 `json:"Success,omitempty" xml:"Success,omitempty"`
	RequestId         *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SqlFlashbackTasks *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasks `json:"SqlFlashbackTasks,omitempty" xml:"SqlFlashbackTasks,omitempty" type:"Struct"`
}

func (s DescribeSqlFlashbakTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSqlFlashbakTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSqlFlashbakTaskResponseBody) SetSuccess(v bool) *DescribeSqlFlashbakTaskResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBody) SetRequestId(v string) *DescribeSqlFlashbakTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBody) SetSqlFlashbackTasks(v *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasks) *DescribeSqlFlashbakTaskResponseBody {
	s.SqlFlashbackTasks = v
	return s
}

type DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasks struct {
	SqlFlashbackTask []*DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask `json:"SqlFlashbackTask,omitempty" xml:"SqlFlashbackTask,omitempty" type:"Repeated"`
}

func (s DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasks) String() string {
	return tea.Prettify(s)
}

func (s DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasks) GoString() string {
	return s.String()
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasks) SetSqlFlashbackTask(v []*DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasks {
	s.SqlFlashbackTask = v
	return s
}

type DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask struct {
	TableName         *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	SearchEndTime     *int64  `json:"SearchEndTime,omitempty" xml:"SearchEndTime,omitempty"`
	ExpireTime        *int64  `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	DownloadUrl       *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	RecallProgress    *int32  `json:"RecallProgress,omitempty" xml:"RecallProgress,omitempty"`
	SqlPk             *string `json:"SqlPk,omitempty" xml:"SqlPk,omitempty"`
	InstId            *string `json:"InstId,omitempty" xml:"InstId,omitempty"`
	RecallType        *int32  `json:"RecallType,omitempty" xml:"RecallType,omitempty"`
	SearchStartTime   *int64  `json:"SearchStartTime,omitempty" xml:"SearchStartTime,omitempty"`
	GmtModified       *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	SqlCounter        *int64  `json:"SqlCounter,omitempty" xml:"SqlCounter,omitempty"`
	DbName            *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	RecallRestoreType *int32  `json:"RecallRestoreType,omitempty" xml:"RecallRestoreType,omitempty"`
	GmtCreate         *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	TraceId           *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
	Id                *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	RecallStatus      *int32  `json:"RecallStatus,omitempty" xml:"RecallStatus,omitempty"`
	SqlType           *string `json:"SqlType,omitempty" xml:"SqlType,omitempty"`
}

func (s DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) String() string {
	return tea.Prettify(s)
}

func (s DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) GoString() string {
	return s.String()
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) SetTableName(v string) *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask {
	s.TableName = &v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) SetSearchEndTime(v int64) *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask {
	s.SearchEndTime = &v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) SetExpireTime(v int64) *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask {
	s.ExpireTime = &v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) SetDownloadUrl(v string) *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask {
	s.DownloadUrl = &v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) SetRecallProgress(v int32) *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask {
	s.RecallProgress = &v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) SetSqlPk(v string) *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask {
	s.SqlPk = &v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) SetInstId(v string) *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask {
	s.InstId = &v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) SetRecallType(v int32) *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask {
	s.RecallType = &v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) SetSearchStartTime(v int64) *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask {
	s.SearchStartTime = &v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) SetGmtModified(v int64) *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask {
	s.GmtModified = &v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) SetSqlCounter(v int64) *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask {
	s.SqlCounter = &v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) SetDbName(v string) *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask {
	s.DbName = &v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) SetRecallRestoreType(v int32) *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask {
	s.RecallRestoreType = &v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) SetGmtCreate(v int64) *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask {
	s.GmtCreate = &v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) SetTraceId(v string) *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask {
	s.TraceId = &v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) SetId(v int64) *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask {
	s.Id = &v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) SetRecallStatus(v int32) *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask {
	s.RecallStatus = &v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) SetSqlType(v string) *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask {
	s.SqlType = &v
	return s
}

type DescribeSqlFlashbakTaskResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSqlFlashbakTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSqlFlashbakTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSqlFlashbakTaskResponse) GoString() string {
	return s.String()
}

func (s *DescribeSqlFlashbakTaskResponse) SetHeaders(v map[string]*string) *DescribeSqlFlashbakTaskResponse {
	s.Headers = v
	return s
}

func (s *DescribeSqlFlashbakTaskResponse) SetBody(v *DescribeSqlFlashbakTaskResponseBody) *DescribeSqlFlashbakTaskResponse {
	s.Body = v
	return s
}

type DescribeTableRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	TableName      *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeTableRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTableRequest) GoString() string {
	return s.String()
}

func (s *DescribeTableRequest) SetRegionId(v string) *DescribeTableRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeTableRequest) SetDrdsInstanceId(v string) *DescribeTableRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeTableRequest) SetDbName(v string) *DescribeTableRequest {
	s.DbName = &v
	return s
}

func (s *DescribeTableRequest) SetTableName(v string) *DescribeTableRequest {
	s.TableName = &v
	return s
}

type DescribeTableResponseBody struct {
	Success   *bool                          `json:"Success,omitempty" xml:"Success,omitempty"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *DescribeTableResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s DescribeTableResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTableResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTableResponseBody) SetSuccess(v bool) *DescribeTableResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeTableResponseBody) SetRequestId(v string) *DescribeTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTableResponseBody) SetData(v *DescribeTableResponseBodyData) *DescribeTableResponseBody {
	s.Data = v
	return s
}

type DescribeTableResponseBodyData struct {
	List []*DescribeTableResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s DescribeTableResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeTableResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeTableResponseBodyData) SetList(v []*DescribeTableResponseBodyDataList) *DescribeTableResponseBodyData {
	s.List = v
	return s
}

type DescribeTableResponseBodyDataList struct {
	Index       *string `json:"Index,omitempty" xml:"Index,omitempty"`
	IsAllowNull *string `json:"IsAllowNull,omitempty" xml:"IsAllowNull,omitempty"`
	ColumnName  *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	IsPk        *string `json:"IsPk,omitempty" xml:"IsPk,omitempty"`
	ColumnType  *string `json:"ColumnType,omitempty" xml:"ColumnType,omitempty"`
	Extra       *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
}

func (s DescribeTableResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s DescribeTableResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *DescribeTableResponseBodyDataList) SetIndex(v string) *DescribeTableResponseBodyDataList {
	s.Index = &v
	return s
}

func (s *DescribeTableResponseBodyDataList) SetIsAllowNull(v string) *DescribeTableResponseBodyDataList {
	s.IsAllowNull = &v
	return s
}

func (s *DescribeTableResponseBodyDataList) SetColumnName(v string) *DescribeTableResponseBodyDataList {
	s.ColumnName = &v
	return s
}

func (s *DescribeTableResponseBodyDataList) SetIsPk(v string) *DescribeTableResponseBodyDataList {
	s.IsPk = &v
	return s
}

func (s *DescribeTableResponseBodyDataList) SetColumnType(v string) *DescribeTableResponseBodyDataList {
	s.ColumnType = &v
	return s
}

func (s *DescribeTableResponseBodyDataList) SetExtra(v string) *DescribeTableResponseBodyDataList {
	s.Extra = &v
	return s
}

type DescribeTableResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeTableResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeTableResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTableResponse) GoString() string {
	return s.String()
}

func (s *DescribeTableResponse) SetHeaders(v map[string]*string) *DescribeTableResponse {
	s.Headers = v
	return s
}

func (s *DescribeTableResponse) SetBody(v *DescribeTableResponseBody) *DescribeTableResponse {
	s.Body = v
	return s
}

type DescribeTableListByTypeRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	Query          *string `json:"Query,omitempty" xml:"Query,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentPage    *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	TableType      *string `json:"TableType,omitempty" xml:"TableType,omitempty"`
}

func (s DescribeTableListByTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTableListByTypeRequest) GoString() string {
	return s.String()
}

func (s *DescribeTableListByTypeRequest) SetRegionId(v string) *DescribeTableListByTypeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeTableListByTypeRequest) SetDrdsInstanceId(v string) *DescribeTableListByTypeRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeTableListByTypeRequest) SetDbName(v string) *DescribeTableListByTypeRequest {
	s.DbName = &v
	return s
}

func (s *DescribeTableListByTypeRequest) SetQuery(v string) *DescribeTableListByTypeRequest {
	s.Query = &v
	return s
}

func (s *DescribeTableListByTypeRequest) SetPageSize(v int32) *DescribeTableListByTypeRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTableListByTypeRequest) SetCurrentPage(v int32) *DescribeTableListByTypeRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeTableListByTypeRequest) SetTableType(v string) *DescribeTableListByTypeRequest {
	s.TableType = &v
	return s
}

type DescribeTableListByTypeResponseBody struct {
	PageSize   *int32                                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int32                                     `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	RequestId  *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total      *int32                                     `json:"Total,omitempty" xml:"Total,omitempty"`
	Success    *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
	List       []*DescribeTableListByTypeResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s DescribeTableListByTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTableListByTypeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTableListByTypeResponseBody) SetPageSize(v int32) *DescribeTableListByTypeResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeTableListByTypeResponseBody) SetPageNumber(v int32) *DescribeTableListByTypeResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeTableListByTypeResponseBody) SetRequestId(v string) *DescribeTableListByTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTableListByTypeResponseBody) SetTotal(v int32) *DescribeTableListByTypeResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeTableListByTypeResponseBody) SetSuccess(v bool) *DescribeTableListByTypeResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeTableListByTypeResponseBody) SetList(v []*DescribeTableListByTypeResponseBodyList) *DescribeTableListByTypeResponseBody {
	s.List = v
	return s
}

type DescribeTableListByTypeResponseBodyList struct {
	Property  *string `json:"Property,omitempty" xml:"Property,omitempty"`
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeTableListByTypeResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s DescribeTableListByTypeResponseBodyList) GoString() string {
	return s.String()
}

func (s *DescribeTableListByTypeResponseBodyList) SetProperty(v string) *DescribeTableListByTypeResponseBodyList {
	s.Property = &v
	return s
}

func (s *DescribeTableListByTypeResponseBodyList) SetTableName(v string) *DescribeTableListByTypeResponseBodyList {
	s.TableName = &v
	return s
}

type DescribeTableListByTypeResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeTableListByTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeTableListByTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTableListByTypeResponse) GoString() string {
	return s.String()
}

func (s *DescribeTableListByTypeResponse) SetHeaders(v map[string]*string) *DescribeTableListByTypeResponse {
	s.Headers = v
	return s
}

func (s *DescribeTableListByTypeResponse) SetBody(v *DescribeTableListByTypeResponseBody) *DescribeTableListByTypeResponse {
	s.Body = v
	return s
}

type DescribeTablesRequest struct {
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	Query          *string `json:"Query,omitempty" xml:"Query,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentPage    *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeTablesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTablesRequest) GoString() string {
	return s.String()
}

func (s *DescribeTablesRequest) SetDrdsInstanceId(v string) *DescribeTablesRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeTablesRequest) SetDbName(v string) *DescribeTablesRequest {
	s.DbName = &v
	return s
}

func (s *DescribeTablesRequest) SetQuery(v string) *DescribeTablesRequest {
	s.Query = &v
	return s
}

func (s *DescribeTablesRequest) SetPageSize(v int32) *DescribeTablesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTablesRequest) SetCurrentPage(v int32) *DescribeTablesRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeTablesRequest) SetRegionId(v string) *DescribeTablesRequest {
	s.RegionId = &v
	return s
}

type DescribeTablesResponseBody struct {
	PageSize   *int32                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int32                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	RequestId  *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total      *int32                            `json:"Total,omitempty" xml:"Total,omitempty"`
	Success    *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
	List       []*DescribeTablesResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s DescribeTablesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTablesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTablesResponseBody) SetPageSize(v int32) *DescribeTablesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeTablesResponseBody) SetPageNumber(v int32) *DescribeTablesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeTablesResponseBody) SetRequestId(v string) *DescribeTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTablesResponseBody) SetTotal(v int32) *DescribeTablesResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeTablesResponseBody) SetSuccess(v bool) *DescribeTablesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeTablesResponseBody) SetList(v []*DescribeTablesResponseBodyList) *DescribeTablesResponseBody {
	s.List = v
	return s
}

type DescribeTablesResponseBodyList struct {
	Status             *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	IsLocked           *bool   `json:"IsLocked,omitempty" xml:"IsLocked,omitempty"`
	ShardKey           *string `json:"ShardKey,omitempty" xml:"ShardKey,omitempty"`
	IsShard            *bool   `json:"IsShard,omitempty" xml:"IsShard,omitempty"`
	Broadcast          *bool   `json:"Broadcast,omitempty" xml:"Broadcast,omitempty"`
	AllowFullTableScan *bool   `json:"AllowFullTableScan,omitempty" xml:"AllowFullTableScan,omitempty"`
	Table              *string `json:"Table,omitempty" xml:"Table,omitempty"`
	DbInstType         *int32  `json:"DbInstType,omitempty" xml:"DbInstType,omitempty"`
}

func (s DescribeTablesResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s DescribeTablesResponseBodyList) GoString() string {
	return s.String()
}

func (s *DescribeTablesResponseBodyList) SetStatus(v int32) *DescribeTablesResponseBodyList {
	s.Status = &v
	return s
}

func (s *DescribeTablesResponseBodyList) SetIsLocked(v bool) *DescribeTablesResponseBodyList {
	s.IsLocked = &v
	return s
}

func (s *DescribeTablesResponseBodyList) SetShardKey(v string) *DescribeTablesResponseBodyList {
	s.ShardKey = &v
	return s
}

func (s *DescribeTablesResponseBodyList) SetIsShard(v bool) *DescribeTablesResponseBodyList {
	s.IsShard = &v
	return s
}

func (s *DescribeTablesResponseBodyList) SetBroadcast(v bool) *DescribeTablesResponseBodyList {
	s.Broadcast = &v
	return s
}

func (s *DescribeTablesResponseBodyList) SetAllowFullTableScan(v bool) *DescribeTablesResponseBodyList {
	s.AllowFullTableScan = &v
	return s
}

func (s *DescribeTablesResponseBodyList) SetTable(v string) *DescribeTablesResponseBodyList {
	s.Table = &v
	return s
}

func (s *DescribeTablesResponseBodyList) SetDbInstType(v int32) *DescribeTablesResponseBodyList {
	s.DbInstType = &v
	return s
}

type DescribeTablesResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeTablesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeTablesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTablesResponse) GoString() string {
	return s.String()
}

func (s *DescribeTablesResponse) SetHeaders(v map[string]*string) *DescribeTablesResponse {
	s.Headers = v
	return s
}

func (s *DescribeTablesResponse) SetBody(v *DescribeTablesResponseBody) *DescribeTablesResponse {
	s.Body = v
	return s
}

type DisableSqlAuditRequest struct {
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
}

func (s DisableSqlAuditRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableSqlAuditRequest) GoString() string {
	return s.String()
}

func (s *DisableSqlAuditRequest) SetDrdsInstanceId(v string) *DisableSqlAuditRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DisableSqlAuditRequest) SetDbName(v string) *DisableSqlAuditRequest {
	s.DbName = &v
	return s
}

type DisableSqlAuditResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DisableSqlAuditResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableSqlAuditResponseBody) GoString() string {
	return s.String()
}

func (s *DisableSqlAuditResponseBody) SetRequestId(v string) *DisableSqlAuditResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableSqlAuditResponseBody) SetResult(v bool) *DisableSqlAuditResponseBody {
	s.Result = &v
	return s
}

func (s *DisableSqlAuditResponseBody) SetSuccess(v bool) *DisableSqlAuditResponseBody {
	s.Success = &v
	return s
}

type DisableSqlAuditResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DisableSqlAuditResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableSqlAuditResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableSqlAuditResponse) GoString() string {
	return s.String()
}

func (s *DisableSqlAuditResponse) SetHeaders(v map[string]*string) *DisableSqlAuditResponse {
	s.Headers = v
	return s
}

func (s *DisableSqlAuditResponse) SetBody(v *DisableSqlAuditResponseBody) *DisableSqlAuditResponse {
	s.Body = v
	return s
}

type EnableSqlAuditRequest struct {
	DrdsInstanceId       *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	DbName               *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	IsRecall             *bool   `json:"IsRecall,omitempty" xml:"IsRecall,omitempty"`
	RecallStartTimestamp *string `json:"RecallStartTimestamp,omitempty" xml:"RecallStartTimestamp,omitempty"`
	RecallEndTimestamp   *string `json:"RecallEndTimestamp,omitempty" xml:"RecallEndTimestamp,omitempty"`
}

func (s EnableSqlAuditRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableSqlAuditRequest) GoString() string {
	return s.String()
}

func (s *EnableSqlAuditRequest) SetDrdsInstanceId(v string) *EnableSqlAuditRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *EnableSqlAuditRequest) SetDbName(v string) *EnableSqlAuditRequest {
	s.DbName = &v
	return s
}

func (s *EnableSqlAuditRequest) SetIsRecall(v bool) *EnableSqlAuditRequest {
	s.IsRecall = &v
	return s
}

func (s *EnableSqlAuditRequest) SetRecallStartTimestamp(v string) *EnableSqlAuditRequest {
	s.RecallStartTimestamp = &v
	return s
}

func (s *EnableSqlAuditRequest) SetRecallEndTimestamp(v string) *EnableSqlAuditRequest {
	s.RecallEndTimestamp = &v
	return s
}

type EnableSqlAuditResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnableSqlAuditResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableSqlAuditResponseBody) GoString() string {
	return s.String()
}

func (s *EnableSqlAuditResponseBody) SetRequestId(v string) *EnableSqlAuditResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnableSqlAuditResponseBody) SetResult(v bool) *EnableSqlAuditResponseBody {
	s.Result = &v
	return s
}

func (s *EnableSqlAuditResponseBody) SetSuccess(v bool) *EnableSqlAuditResponseBody {
	s.Success = &v
	return s
}

type EnableSqlAuditResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EnableSqlAuditResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableSqlAuditResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableSqlAuditResponse) GoString() string {
	return s.String()
}

func (s *EnableSqlAuditResponse) SetHeaders(v map[string]*string) *EnableSqlAuditResponse {
	s.Headers = v
	return s
}

func (s *EnableSqlAuditResponse) SetBody(v *EnableSqlAuditResponseBody) *EnableSqlAuditResponse {
	s.Body = v
	return s
}

type EnableSqlFlashbackMatchSwitchRequest struct {
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
}

func (s EnableSqlFlashbackMatchSwitchRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableSqlFlashbackMatchSwitchRequest) GoString() string {
	return s.String()
}

func (s *EnableSqlFlashbackMatchSwitchRequest) SetDrdsInstanceId(v string) *EnableSqlFlashbackMatchSwitchRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *EnableSqlFlashbackMatchSwitchRequest) SetDbName(v string) *EnableSqlFlashbackMatchSwitchRequest {
	s.DbName = &v
	return s
}

type EnableSqlFlashbackMatchSwitchResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnableSqlFlashbackMatchSwitchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableSqlFlashbackMatchSwitchResponseBody) GoString() string {
	return s.String()
}

func (s *EnableSqlFlashbackMatchSwitchResponseBody) SetRequestId(v string) *EnableSqlFlashbackMatchSwitchResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnableSqlFlashbackMatchSwitchResponseBody) SetResult(v bool) *EnableSqlFlashbackMatchSwitchResponseBody {
	s.Result = &v
	return s
}

func (s *EnableSqlFlashbackMatchSwitchResponseBody) SetSuccess(v bool) *EnableSqlFlashbackMatchSwitchResponseBody {
	s.Success = &v
	return s
}

type EnableSqlFlashbackMatchSwitchResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EnableSqlFlashbackMatchSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableSqlFlashbackMatchSwitchResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableSqlFlashbackMatchSwitchResponse) GoString() string {
	return s.String()
}

func (s *EnableSqlFlashbackMatchSwitchResponse) SetHeaders(v map[string]*string) *EnableSqlFlashbackMatchSwitchResponse {
	s.Headers = v
	return s
}

func (s *EnableSqlFlashbackMatchSwitchResponse) SetBody(v *EnableSqlFlashbackMatchSwitchResponseBody) *EnableSqlFlashbackMatchSwitchResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	RegionId     *string                       `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceType *string                       `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	NextToken    *string                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Tag          []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	ResourceId   []*string                     `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetRegionId(v string) *ListTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetTag(v []*ListTagResourcesRequestTag) *ListTagResourcesRequest {
	s.Tag = v
	return s
}

func (s *ListTagResourcesRequest) SetResourceId(v []*string) *ListTagResourcesRequest {
	s.ResourceId = v
	return s
}

type ListTagResourcesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequestTag) SetKey(v string) *ListTagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *ListTagResourcesRequestTag) SetValue(v string) *ListTagResourcesRequestTag {
	s.Value = &v
	return s
}

type ListTagResourcesResponseBody struct {
	NextToken    *string                                   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId    *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
	TagResources *ListTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Struct"`
}

func (s ListTagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBody) SetNextToken(v string) *ListTagResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetRequestId(v string) *ListTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetSuccess(v bool) *ListTagResourcesResponseBody {
	s.Success = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetTagResources(v *ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody {
	s.TagResources = v
	return s
}

type ListTagResourcesResponseBodyTagResources struct {
	TagResource []*ListTagResourcesResponseBodyTagResourcesTagResource `json:"TagResource,omitempty" xml:"TagResource,omitempty" type:"Repeated"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagResource(v []*ListTagResourcesResponseBodyTagResourcesTagResource) *ListTagResourcesResponseBodyTagResources {
	s.TagResource = v
	return s
}

type ListTagResourcesResponseBodyTagResourcesTagResource struct {
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagValue     *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	TagKey       *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceType(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagValue(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagValue = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceId(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagKey(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagKey = &v
	return s
}

type ListTagResourcesResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponse) SetHeaders(v map[string]*string) *ListTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListTagResourcesResponse) SetBody(v *ListTagResourcesResponseBody) *ListTagResourcesResponse {
	s.Body = v
	return s
}

type ManagePrivateRdsRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	DBInstanceId   *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	RdsAction      *string `json:"RdsAction,omitempty" xml:"RdsAction,omitempty"`
	Params         *string `json:"Params,omitempty" xml:"Params,omitempty"`
}

func (s ManagePrivateRdsRequest) String() string {
	return tea.Prettify(s)
}

func (s ManagePrivateRdsRequest) GoString() string {
	return s.String()
}

func (s *ManagePrivateRdsRequest) SetRegionId(v string) *ManagePrivateRdsRequest {
	s.RegionId = &v
	return s
}

func (s *ManagePrivateRdsRequest) SetDrdsInstanceId(v string) *ManagePrivateRdsRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *ManagePrivateRdsRequest) SetDBInstanceId(v string) *ManagePrivateRdsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ManagePrivateRdsRequest) SetRdsAction(v string) *ManagePrivateRdsRequest {
	s.RdsAction = &v
	return s
}

func (s *ManagePrivateRdsRequest) SetParams(v string) *ManagePrivateRdsRequest {
	s.Params = &v
	return s
}

type ManagePrivateRdsResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ManagePrivateRdsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ManagePrivateRdsResponseBody) GoString() string {
	return s.String()
}

func (s *ManagePrivateRdsResponseBody) SetData(v string) *ManagePrivateRdsResponseBody {
	s.Data = &v
	return s
}

func (s *ManagePrivateRdsResponseBody) SetRequestId(v string) *ManagePrivateRdsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ManagePrivateRdsResponseBody) SetSuccess(v bool) *ManagePrivateRdsResponseBody {
	s.Success = &v
	return s
}

type ManagePrivateRdsResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ManagePrivateRdsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ManagePrivateRdsResponse) String() string {
	return tea.Prettify(s)
}

func (s ManagePrivateRdsResponse) GoString() string {
	return s.String()
}

func (s *ManagePrivateRdsResponse) SetHeaders(v map[string]*string) *ManagePrivateRdsResponse {
	s.Headers = v
	return s
}

func (s *ManagePrivateRdsResponse) SetBody(v *ManagePrivateRdsResponseBody) *ManagePrivateRdsResponse {
	s.Body = v
	return s
}

type ModifyDrdsInstanceDescriptionRequest struct {
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s ModifyDrdsInstanceDescriptionRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDrdsInstanceDescriptionRequest) GoString() string {
	return s.String()
}

func (s *ModifyDrdsInstanceDescriptionRequest) SetDrdsInstanceId(v string) *ModifyDrdsInstanceDescriptionRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *ModifyDrdsInstanceDescriptionRequest) SetDescription(v string) *ModifyDrdsInstanceDescriptionRequest {
	s.Description = &v
	return s
}

type ModifyDrdsInstanceDescriptionResponseBody struct {
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDrdsInstanceDescriptionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDrdsInstanceDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDrdsInstanceDescriptionResponseBody) SetSuccess(v bool) *ModifyDrdsInstanceDescriptionResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyDrdsInstanceDescriptionResponseBody) SetRequestId(v string) *ModifyDrdsInstanceDescriptionResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDrdsInstanceDescriptionResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDrdsInstanceDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDrdsInstanceDescriptionResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDrdsInstanceDescriptionResponse) GoString() string {
	return s.String()
}

func (s *ModifyDrdsInstanceDescriptionResponse) SetHeaders(v map[string]*string) *ModifyDrdsInstanceDescriptionResponse {
	s.Headers = v
	return s
}

func (s *ModifyDrdsInstanceDescriptionResponse) SetBody(v *ModifyDrdsInstanceDescriptionResponseBody) *ModifyDrdsInstanceDescriptionResponse {
	s.Body = v
	return s
}

type ModifyDrdsIpWhiteListRequest struct {
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	IpWhiteList    *string `json:"IpWhiteList,omitempty" xml:"IpWhiteList,omitempty"`
	Mode           *bool   `json:"Mode,omitempty" xml:"Mode,omitempty"`
	GroupName      *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	GroupAttribute *string `json:"GroupAttribute,omitempty" xml:"GroupAttribute,omitempty"`
}

func (s ModifyDrdsIpWhiteListRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDrdsIpWhiteListRequest) GoString() string {
	return s.String()
}

func (s *ModifyDrdsIpWhiteListRequest) SetDrdsInstanceId(v string) *ModifyDrdsIpWhiteListRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *ModifyDrdsIpWhiteListRequest) SetDbName(v string) *ModifyDrdsIpWhiteListRequest {
	s.DbName = &v
	return s
}

func (s *ModifyDrdsIpWhiteListRequest) SetIpWhiteList(v string) *ModifyDrdsIpWhiteListRequest {
	s.IpWhiteList = &v
	return s
}

func (s *ModifyDrdsIpWhiteListRequest) SetMode(v bool) *ModifyDrdsIpWhiteListRequest {
	s.Mode = &v
	return s
}

func (s *ModifyDrdsIpWhiteListRequest) SetGroupName(v string) *ModifyDrdsIpWhiteListRequest {
	s.GroupName = &v
	return s
}

func (s *ModifyDrdsIpWhiteListRequest) SetGroupAttribute(v string) *ModifyDrdsIpWhiteListRequest {
	s.GroupAttribute = &v
	return s
}

type ModifyDrdsIpWhiteListResponseBody struct {
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDrdsIpWhiteListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDrdsIpWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDrdsIpWhiteListResponseBody) SetSuccess(v bool) *ModifyDrdsIpWhiteListResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyDrdsIpWhiteListResponseBody) SetRequestId(v string) *ModifyDrdsIpWhiteListResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDrdsIpWhiteListResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDrdsIpWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDrdsIpWhiteListResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDrdsIpWhiteListResponse) GoString() string {
	return s.String()
}

func (s *ModifyDrdsIpWhiteListResponse) SetHeaders(v map[string]*string) *ModifyDrdsIpWhiteListResponse {
	s.Headers = v
	return s
}

func (s *ModifyDrdsIpWhiteListResponse) SetBody(v *ModifyDrdsIpWhiteListResponseBody) *ModifyDrdsIpWhiteListResponse {
	s.Body = v
	return s
}

type ModifyRdsReadWeightRequest struct {
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	InstanceNames  *string `json:"InstanceNames,omitempty" xml:"InstanceNames,omitempty"`
	Weights        *string `json:"Weights,omitempty" xml:"Weights,omitempty"`
}

func (s ModifyRdsReadWeightRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyRdsReadWeightRequest) GoString() string {
	return s.String()
}

func (s *ModifyRdsReadWeightRequest) SetDrdsInstanceId(v string) *ModifyRdsReadWeightRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *ModifyRdsReadWeightRequest) SetDbName(v string) *ModifyRdsReadWeightRequest {
	s.DbName = &v
	return s
}

func (s *ModifyRdsReadWeightRequest) SetInstanceNames(v string) *ModifyRdsReadWeightRequest {
	s.InstanceNames = &v
	return s
}

func (s *ModifyRdsReadWeightRequest) SetWeights(v string) *ModifyRdsReadWeightRequest {
	s.Weights = &v
	return s
}

type ModifyRdsReadWeightResponseBody struct {
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyRdsReadWeightResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyRdsReadWeightResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyRdsReadWeightResponseBody) SetSuccess(v bool) *ModifyRdsReadWeightResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyRdsReadWeightResponseBody) SetRequestId(v string) *ModifyRdsReadWeightResponseBody {
	s.RequestId = &v
	return s
}

type ModifyRdsReadWeightResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyRdsReadWeightResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyRdsReadWeightResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyRdsReadWeightResponse) GoString() string {
	return s.String()
}

func (s *ModifyRdsReadWeightResponse) SetHeaders(v map[string]*string) *ModifyRdsReadWeightResponse {
	s.Headers = v
	return s
}

func (s *ModifyRdsReadWeightResponse) SetBody(v *ModifyRdsReadWeightResponseBody) *ModifyRdsReadWeightResponse {
	s.Body = v
	return s
}

type PutStartBackupRequest struct {
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	BackupMode     *string `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	BackupLevel    *string `json:"BackupLevel,omitempty" xml:"BackupLevel,omitempty"`
	BackupDbNames  *string `json:"BackupDbNames,omitempty" xml:"BackupDbNames,omitempty"`
}

func (s PutStartBackupRequest) String() string {
	return tea.Prettify(s)
}

func (s PutStartBackupRequest) GoString() string {
	return s.String()
}

func (s *PutStartBackupRequest) SetDrdsInstanceId(v string) *PutStartBackupRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *PutStartBackupRequest) SetBackupMode(v string) *PutStartBackupRequest {
	s.BackupMode = &v
	return s
}

func (s *PutStartBackupRequest) SetBackupLevel(v string) *PutStartBackupRequest {
	s.BackupLevel = &v
	return s
}

func (s *PutStartBackupRequest) SetBackupDbNames(v string) *PutStartBackupRequest {
	s.BackupDbNames = &v
	return s
}

type PutStartBackupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *string `json:"Result,omitempty" xml:"Result,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PutStartBackupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PutStartBackupResponseBody) GoString() string {
	return s.String()
}

func (s *PutStartBackupResponseBody) SetRequestId(v string) *PutStartBackupResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutStartBackupResponseBody) SetResult(v string) *PutStartBackupResponseBody {
	s.Result = &v
	return s
}

func (s *PutStartBackupResponseBody) SetSuccess(v bool) *PutStartBackupResponseBody {
	s.Success = &v
	return s
}

type PutStartBackupResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PutStartBackupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PutStartBackupResponse) String() string {
	return tea.Prettify(s)
}

func (s PutStartBackupResponse) GoString() string {
	return s.String()
}

func (s *PutStartBackupResponse) SetHeaders(v map[string]*string) *PutStartBackupResponse {
	s.Headers = v
	return s
}

func (s *PutStartBackupResponse) SetBody(v *PutStartBackupResponseBody) *PutStartBackupResponse {
	s.Body = v
	return s
}

type ReleaseInstanceInternetAddressRequest struct {
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ReleaseInstanceInternetAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleaseInstanceInternetAddressRequest) GoString() string {
	return s.String()
}

func (s *ReleaseInstanceInternetAddressRequest) SetDrdsInstanceId(v string) *ReleaseInstanceInternetAddressRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *ReleaseInstanceInternetAddressRequest) SetRegionId(v string) *ReleaseInstanceInternetAddressRequest {
	s.RegionId = &v
	return s
}

type ReleaseInstanceInternetAddressResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleaseInstanceInternetAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReleaseInstanceInternetAddressResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseInstanceInternetAddressResponseBody) SetData(v bool) *ReleaseInstanceInternetAddressResponseBody {
	s.Data = &v
	return s
}

func (s *ReleaseInstanceInternetAddressResponseBody) SetRequestId(v string) *ReleaseInstanceInternetAddressResponseBody {
	s.RequestId = &v
	return s
}

type ReleaseInstanceInternetAddressResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ReleaseInstanceInternetAddressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReleaseInstanceInternetAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleaseInstanceInternetAddressResponse) GoString() string {
	return s.String()
}

func (s *ReleaseInstanceInternetAddressResponse) SetHeaders(v map[string]*string) *ReleaseInstanceInternetAddressResponse {
	s.Headers = v
	return s
}

func (s *ReleaseInstanceInternetAddressResponse) SetBody(v *ReleaseInstanceInternetAddressResponseBody) *ReleaseInstanceInternetAddressResponse {
	s.Body = v
	return s
}

type RemoveBackupsSetRequest struct {
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	BackupId       *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
}

func (s RemoveBackupsSetRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveBackupsSetRequest) GoString() string {
	return s.String()
}

func (s *RemoveBackupsSetRequest) SetDrdsInstanceId(v string) *RemoveBackupsSetRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *RemoveBackupsSetRequest) SetBackupId(v string) *RemoveBackupsSetRequest {
	s.BackupId = &v
	return s
}

type RemoveBackupsSetResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *string `json:"Result,omitempty" xml:"Result,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RemoveBackupsSetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveBackupsSetResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveBackupsSetResponseBody) SetRequestId(v string) *RemoveBackupsSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveBackupsSetResponseBody) SetResult(v string) *RemoveBackupsSetResponseBody {
	s.Result = &v
	return s
}

func (s *RemoveBackupsSetResponseBody) SetSuccess(v bool) *RemoveBackupsSetResponseBody {
	s.Success = &v
	return s
}

type RemoveBackupsSetResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveBackupsSetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveBackupsSetResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveBackupsSetResponse) GoString() string {
	return s.String()
}

func (s *RemoveBackupsSetResponse) SetHeaders(v map[string]*string) *RemoveBackupsSetResponse {
	s.Headers = v
	return s
}

func (s *RemoveBackupsSetResponse) SetBody(v *RemoveBackupsSetResponseBody) *RemoveBackupsSetResponse {
	s.Body = v
	return s
}

type RemoveDrdsDbRequest struct {
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
}

func (s RemoveDrdsDbRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveDrdsDbRequest) GoString() string {
	return s.String()
}

func (s *RemoveDrdsDbRequest) SetDrdsInstanceId(v string) *RemoveDrdsDbRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *RemoveDrdsDbRequest) SetDbName(v string) *RemoveDrdsDbRequest {
	s.DbName = &v
	return s
}

type RemoveDrdsDbResponseBody struct {
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveDrdsDbResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveDrdsDbResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveDrdsDbResponseBody) SetSuccess(v bool) *RemoveDrdsDbResponseBody {
	s.Success = &v
	return s
}

func (s *RemoveDrdsDbResponseBody) SetRequestId(v string) *RemoveDrdsDbResponseBody {
	s.RequestId = &v
	return s
}

type RemoveDrdsDbResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveDrdsDbResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveDrdsDbResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveDrdsDbResponse) GoString() string {
	return s.String()
}

func (s *RemoveDrdsDbResponse) SetHeaders(v map[string]*string) *RemoveDrdsDbResponse {
	s.Headers = v
	return s
}

func (s *RemoveDrdsDbResponse) SetBody(v *RemoveDrdsDbResponseBody) *RemoveDrdsDbResponse {
	s.Body = v
	return s
}

type RemoveDrdsDbFailedRecordRequest struct {
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
}

func (s RemoveDrdsDbFailedRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveDrdsDbFailedRecordRequest) GoString() string {
	return s.String()
}

func (s *RemoveDrdsDbFailedRecordRequest) SetDrdsInstanceId(v string) *RemoveDrdsDbFailedRecordRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *RemoveDrdsDbFailedRecordRequest) SetDbName(v string) *RemoveDrdsDbFailedRecordRequest {
	s.DbName = &v
	return s
}

type RemoveDrdsDbFailedRecordResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RemoveDrdsDbFailedRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveDrdsDbFailedRecordResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveDrdsDbFailedRecordResponseBody) SetRequestId(v string) *RemoveDrdsDbFailedRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveDrdsDbFailedRecordResponseBody) SetResult(v bool) *RemoveDrdsDbFailedRecordResponseBody {
	s.Result = &v
	return s
}

func (s *RemoveDrdsDbFailedRecordResponseBody) SetSuccess(v bool) *RemoveDrdsDbFailedRecordResponseBody {
	s.Success = &v
	return s
}

type RemoveDrdsDbFailedRecordResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveDrdsDbFailedRecordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveDrdsDbFailedRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveDrdsDbFailedRecordResponse) GoString() string {
	return s.String()
}

func (s *RemoveDrdsDbFailedRecordResponse) SetHeaders(v map[string]*string) *RemoveDrdsDbFailedRecordResponse {
	s.Headers = v
	return s
}

func (s *RemoveDrdsDbFailedRecordResponse) SetBody(v *RemoveDrdsDbFailedRecordResponseBody) *RemoveDrdsDbFailedRecordResponse {
	s.Body = v
	return s
}

type RemoveDrdsInstanceRequest struct {
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s RemoveDrdsInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveDrdsInstanceRequest) GoString() string {
	return s.String()
}

func (s *RemoveDrdsInstanceRequest) SetDrdsInstanceId(v string) *RemoveDrdsInstanceRequest {
	s.DrdsInstanceId = &v
	return s
}

type RemoveDrdsInstanceResponseBody struct {
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveDrdsInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveDrdsInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveDrdsInstanceResponseBody) SetSuccess(v bool) *RemoveDrdsInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *RemoveDrdsInstanceResponseBody) SetRequestId(v string) *RemoveDrdsInstanceResponseBody {
	s.RequestId = &v
	return s
}

type RemoveDrdsInstanceResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveDrdsInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveDrdsInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveDrdsInstanceResponse) GoString() string {
	return s.String()
}

func (s *RemoveDrdsInstanceResponse) SetHeaders(v map[string]*string) *RemoveDrdsInstanceResponse {
	s.Headers = v
	return s
}

func (s *RemoveDrdsInstanceResponse) SetBody(v *RemoveDrdsInstanceResponseBody) *RemoveDrdsInstanceResponse {
	s.Body = v
	return s
}

type RemoveInstanceAccountRequest struct {
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	AccountName    *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
}

func (s RemoveInstanceAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveInstanceAccountRequest) GoString() string {
	return s.String()
}

func (s *RemoveInstanceAccountRequest) SetDrdsInstanceId(v string) *RemoveInstanceAccountRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *RemoveInstanceAccountRequest) SetAccountName(v string) *RemoveInstanceAccountRequest {
	s.AccountName = &v
	return s
}

type RemoveInstanceAccountResponseBody struct {
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveInstanceAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveInstanceAccountResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveInstanceAccountResponseBody) SetSuccess(v bool) *RemoveInstanceAccountResponseBody {
	s.Success = &v
	return s
}

func (s *RemoveInstanceAccountResponseBody) SetRequestId(v string) *RemoveInstanceAccountResponseBody {
	s.RequestId = &v
	return s
}

type RemoveInstanceAccountResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveInstanceAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveInstanceAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveInstanceAccountResponse) GoString() string {
	return s.String()
}

func (s *RemoveInstanceAccountResponse) SetHeaders(v map[string]*string) *RemoveInstanceAccountResponse {
	s.Headers = v
	return s
}

func (s *RemoveInstanceAccountResponse) SetBody(v *RemoveInstanceAccountResponseBody) *RemoveInstanceAccountResponse {
	s.Body = v
	return s
}

type SetBackupLocalRequest struct {
	DrdsInstanceId           *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	LocalLogRetentionHours   *string `json:"LocalLogRetentionHours,omitempty" xml:"LocalLogRetentionHours,omitempty"`
	LocalLogRetentionSpace   *string `json:"LocalLogRetentionSpace,omitempty" xml:"LocalLogRetentionSpace,omitempty"`
	HighSpaceUsageProtection *string `json:"HighSpaceUsageProtection,omitempty" xml:"HighSpaceUsageProtection,omitempty"`
}

func (s SetBackupLocalRequest) String() string {
	return tea.Prettify(s)
}

func (s SetBackupLocalRequest) GoString() string {
	return s.String()
}

func (s *SetBackupLocalRequest) SetDrdsInstanceId(v string) *SetBackupLocalRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *SetBackupLocalRequest) SetLocalLogRetentionHours(v string) *SetBackupLocalRequest {
	s.LocalLogRetentionHours = &v
	return s
}

func (s *SetBackupLocalRequest) SetLocalLogRetentionSpace(v string) *SetBackupLocalRequest {
	s.LocalLogRetentionSpace = &v
	return s
}

func (s *SetBackupLocalRequest) SetHighSpaceUsageProtection(v string) *SetBackupLocalRequest {
	s.HighSpaceUsageProtection = &v
	return s
}

type SetBackupLocalResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *string `json:"Result,omitempty" xml:"Result,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetBackupLocalResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetBackupLocalResponseBody) GoString() string {
	return s.String()
}

func (s *SetBackupLocalResponseBody) SetRequestId(v string) *SetBackupLocalResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetBackupLocalResponseBody) SetResult(v string) *SetBackupLocalResponseBody {
	s.Result = &v
	return s
}

func (s *SetBackupLocalResponseBody) SetSuccess(v bool) *SetBackupLocalResponseBody {
	s.Success = &v
	return s
}

type SetBackupLocalResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetBackupLocalResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetBackupLocalResponse) String() string {
	return tea.Prettify(s)
}

func (s SetBackupLocalResponse) GoString() string {
	return s.String()
}

func (s *SetBackupLocalResponse) SetHeaders(v map[string]*string) *SetBackupLocalResponse {
	s.Headers = v
	return s
}

func (s *SetBackupLocalResponse) SetBody(v *SetBackupLocalResponseBody) *SetBackupLocalResponse {
	s.Body = v
	return s
}

type SetBackupPolicyRequest struct {
	DrdsInstanceId            *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	PreferredBackupPeriod     *string `json:"PreferredBackupPeriod,omitempty" xml:"PreferredBackupPeriod,omitempty"`
	PreferredBackupStartTime  *string `json:"PreferredBackupStartTime,omitempty" xml:"PreferredBackupStartTime,omitempty"`
	PreferredBackupEndTime    *string `json:"PreferredBackupEndTime,omitempty" xml:"PreferredBackupEndTime,omitempty"`
	BackupMode                *string `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	BackupLevel               *string `json:"BackupLevel,omitempty" xml:"BackupLevel,omitempty"`
	BackupDbNames             *string `json:"BackupDbNames,omitempty" xml:"BackupDbNames,omitempty"`
	BackupLog                 *string `json:"BackupLog,omitempty" xml:"BackupLog,omitempty"`
	DataBackupRetentionPeriod *string `json:"DataBackupRetentionPeriod,omitempty" xml:"DataBackupRetentionPeriod,omitempty"`
	LogBackupRetentionPeriod  *string `json:"LogBackupRetentionPeriod,omitempty" xml:"LogBackupRetentionPeriod,omitempty"`
}

func (s SetBackupPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s SetBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *SetBackupPolicyRequest) SetDrdsInstanceId(v string) *SetBackupPolicyRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *SetBackupPolicyRequest) SetPreferredBackupPeriod(v string) *SetBackupPolicyRequest {
	s.PreferredBackupPeriod = &v
	return s
}

func (s *SetBackupPolicyRequest) SetPreferredBackupStartTime(v string) *SetBackupPolicyRequest {
	s.PreferredBackupStartTime = &v
	return s
}

func (s *SetBackupPolicyRequest) SetPreferredBackupEndTime(v string) *SetBackupPolicyRequest {
	s.PreferredBackupEndTime = &v
	return s
}

func (s *SetBackupPolicyRequest) SetBackupMode(v string) *SetBackupPolicyRequest {
	s.BackupMode = &v
	return s
}

func (s *SetBackupPolicyRequest) SetBackupLevel(v string) *SetBackupPolicyRequest {
	s.BackupLevel = &v
	return s
}

func (s *SetBackupPolicyRequest) SetBackupDbNames(v string) *SetBackupPolicyRequest {
	s.BackupDbNames = &v
	return s
}

func (s *SetBackupPolicyRequest) SetBackupLog(v string) *SetBackupPolicyRequest {
	s.BackupLog = &v
	return s
}

func (s *SetBackupPolicyRequest) SetDataBackupRetentionPeriod(v string) *SetBackupPolicyRequest {
	s.DataBackupRetentionPeriod = &v
	return s
}

func (s *SetBackupPolicyRequest) SetLogBackupRetentionPeriod(v string) *SetBackupPolicyRequest {
	s.LogBackupRetentionPeriod = &v
	return s
}

type SetBackupPolicyResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *string `json:"Result,omitempty" xml:"Result,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetBackupPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetBackupPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *SetBackupPolicyResponseBody) SetRequestId(v string) *SetBackupPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetBackupPolicyResponseBody) SetResult(v string) *SetBackupPolicyResponseBody {
	s.Result = &v
	return s
}

func (s *SetBackupPolicyResponseBody) SetSuccess(v bool) *SetBackupPolicyResponseBody {
	s.Success = &v
	return s
}

type SetBackupPolicyResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetBackupPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetBackupPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s SetBackupPolicyResponse) GoString() string {
	return s.String()
}

func (s *SetBackupPolicyResponse) SetHeaders(v map[string]*string) *SetBackupPolicyResponse {
	s.Headers = v
	return s
}

func (s *SetBackupPolicyResponse) SetBody(v *SetBackupPolicyResponseBody) *SetBackupPolicyResponse {
	s.Body = v
	return s
}

type SetupBroadcastTablesRequest struct {
	RegionId       *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DrdsInstanceId *string   `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	DbName         *string   `json:"DbName,omitempty" xml:"DbName,omitempty"`
	Active         *bool     `json:"Active,omitempty" xml:"Active,omitempty"`
	TableName      []*string `json:"TableName,omitempty" xml:"TableName,omitempty" type:"Repeated"`
}

func (s SetupBroadcastTablesRequest) String() string {
	return tea.Prettify(s)
}

func (s SetupBroadcastTablesRequest) GoString() string {
	return s.String()
}

func (s *SetupBroadcastTablesRequest) SetRegionId(v string) *SetupBroadcastTablesRequest {
	s.RegionId = &v
	return s
}

func (s *SetupBroadcastTablesRequest) SetDrdsInstanceId(v string) *SetupBroadcastTablesRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *SetupBroadcastTablesRequest) SetDbName(v string) *SetupBroadcastTablesRequest {
	s.DbName = &v
	return s
}

func (s *SetupBroadcastTablesRequest) SetActive(v bool) *SetupBroadcastTablesRequest {
	s.Active = &v
	return s
}

func (s *SetupBroadcastTablesRequest) SetTableName(v []*string) *SetupBroadcastTablesRequest {
	s.TableName = v
	return s
}

type SetupBroadcastTablesResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetupBroadcastTablesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetupBroadcastTablesResponseBody) GoString() string {
	return s.String()
}

func (s *SetupBroadcastTablesResponseBody) SetData(v bool) *SetupBroadcastTablesResponseBody {
	s.Data = &v
	return s
}

func (s *SetupBroadcastTablesResponseBody) SetRequestId(v string) *SetupBroadcastTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetupBroadcastTablesResponseBody) SetSuccess(v bool) *SetupBroadcastTablesResponseBody {
	s.Success = &v
	return s
}

type SetupBroadcastTablesResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetupBroadcastTablesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetupBroadcastTablesResponse) String() string {
	return tea.Prettify(s)
}

func (s SetupBroadcastTablesResponse) GoString() string {
	return s.String()
}

func (s *SetupBroadcastTablesResponse) SetHeaders(v map[string]*string) *SetupBroadcastTablesResponse {
	s.Headers = v
	return s
}

func (s *SetupBroadcastTablesResponse) SetBody(v *SetupBroadcastTablesResponseBody) *SetupBroadcastTablesResponse {
	s.Body = v
	return s
}

type SetupDrdsParamsRequest struct {
	RegionId       *string                       `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DrdsInstanceId *string                       `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	ParamLevel     *string                       `json:"ParamLevel,omitempty" xml:"ParamLevel,omitempty"`
	Data           []*SetupDrdsParamsRequestData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s SetupDrdsParamsRequest) String() string {
	return tea.Prettify(s)
}

func (s SetupDrdsParamsRequest) GoString() string {
	return s.String()
}

func (s *SetupDrdsParamsRequest) SetRegionId(v string) *SetupDrdsParamsRequest {
	s.RegionId = &v
	return s
}

func (s *SetupDrdsParamsRequest) SetDrdsInstanceId(v string) *SetupDrdsParamsRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *SetupDrdsParamsRequest) SetParamLevel(v string) *SetupDrdsParamsRequest {
	s.ParamLevel = &v
	return s
}

func (s *SetupDrdsParamsRequest) SetData(v []*SetupDrdsParamsRequestData) *SetupDrdsParamsRequest {
	s.Data = v
	return s
}

type SetupDrdsParamsRequestData struct {
	DbName            *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	ParamType         *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
	ParamValue        *string `json:"ParamValue,omitempty" xml:"ParamValue,omitempty"`
	ParamVariableName *string `json:"ParamVariableName,omitempty" xml:"ParamVariableName,omitempty"`
	ParamRanges       *string `json:"ParamRanges,omitempty" xml:"ParamRanges,omitempty"`
}

func (s SetupDrdsParamsRequestData) String() string {
	return tea.Prettify(s)
}

func (s SetupDrdsParamsRequestData) GoString() string {
	return s.String()
}

func (s *SetupDrdsParamsRequestData) SetDbName(v string) *SetupDrdsParamsRequestData {
	s.DbName = &v
	return s
}

func (s *SetupDrdsParamsRequestData) SetParamType(v string) *SetupDrdsParamsRequestData {
	s.ParamType = &v
	return s
}

func (s *SetupDrdsParamsRequestData) SetParamValue(v string) *SetupDrdsParamsRequestData {
	s.ParamValue = &v
	return s
}

func (s *SetupDrdsParamsRequestData) SetParamVariableName(v string) *SetupDrdsParamsRequestData {
	s.ParamVariableName = &v
	return s
}

func (s *SetupDrdsParamsRequestData) SetParamRanges(v string) *SetupDrdsParamsRequestData {
	s.ParamRanges = &v
	return s
}

type SetupDrdsParamsResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetupDrdsParamsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetupDrdsParamsResponseBody) GoString() string {
	return s.String()
}

func (s *SetupDrdsParamsResponseBody) SetData(v bool) *SetupDrdsParamsResponseBody {
	s.Data = &v
	return s
}

func (s *SetupDrdsParamsResponseBody) SetRequestId(v string) *SetupDrdsParamsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetupDrdsParamsResponseBody) SetSuccess(v bool) *SetupDrdsParamsResponseBody {
	s.Success = &v
	return s
}

type SetupDrdsParamsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetupDrdsParamsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetupDrdsParamsResponse) String() string {
	return tea.Prettify(s)
}

func (s SetupDrdsParamsResponse) GoString() string {
	return s.String()
}

func (s *SetupDrdsParamsResponse) SetHeaders(v map[string]*string) *SetupDrdsParamsResponse {
	s.Headers = v
	return s
}

func (s *SetupDrdsParamsResponse) SetBody(v *SetupDrdsParamsResponseBody) *SetupDrdsParamsResponse {
	s.Body = v
	return s
}

type SetupTableRequest struct {
	RegionId           *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DrdsInstanceId     *string   `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	DbName             *string   `json:"DbName,omitempty" xml:"DbName,omitempty"`
	AllowFullTableScan *bool     `json:"AllowFullTableScan,omitempty" xml:"AllowFullTableScan,omitempty"`
	TableName          []*string `json:"TableName,omitempty" xml:"TableName,omitempty" type:"Repeated"`
}

func (s SetupTableRequest) String() string {
	return tea.Prettify(s)
}

func (s SetupTableRequest) GoString() string {
	return s.String()
}

func (s *SetupTableRequest) SetRegionId(v string) *SetupTableRequest {
	s.RegionId = &v
	return s
}

func (s *SetupTableRequest) SetDrdsInstanceId(v string) *SetupTableRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *SetupTableRequest) SetDbName(v string) *SetupTableRequest {
	s.DbName = &v
	return s
}

func (s *SetupTableRequest) SetAllowFullTableScan(v bool) *SetupTableRequest {
	s.AllowFullTableScan = &v
	return s
}

func (s *SetupTableRequest) SetTableName(v []*string) *SetupTableRequest {
	s.TableName = v
	return s
}

type SetupTableResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetupTableResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetupTableResponseBody) GoString() string {
	return s.String()
}

func (s *SetupTableResponseBody) SetData(v bool) *SetupTableResponseBody {
	s.Data = &v
	return s
}

func (s *SetupTableResponseBody) SetRequestId(v string) *SetupTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetupTableResponseBody) SetSuccess(v bool) *SetupTableResponseBody {
	s.Success = &v
	return s
}

type SetupTableResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetupTableResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetupTableResponse) String() string {
	return tea.Prettify(s)
}

func (s SetupTableResponse) GoString() string {
	return s.String()
}

func (s *SetupTableResponse) SetHeaders(v map[string]*string) *SetupTableResponse {
	s.Headers = v
	return s
}

func (s *SetupTableResponse) SetBody(v *SetupTableResponseBody) *SetupTableResponse {
	s.Body = v
	return s
}

type StartRestoreRequest struct {
	DrdsInstanceId      *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	PreferredBackupTime *string `json:"PreferredBackupTime,omitempty" xml:"PreferredBackupTime,omitempty"`
	BackupMode          *string `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	BackupLevel         *string `json:"BackupLevel,omitempty" xml:"BackupLevel,omitempty"`
	BackupDbNames       *string `json:"BackupDbNames,omitempty" xml:"BackupDbNames,omitempty"`
	BackupId            *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
}

func (s StartRestoreRequest) String() string {
	return tea.Prettify(s)
}

func (s StartRestoreRequest) GoString() string {
	return s.String()
}

func (s *StartRestoreRequest) SetDrdsInstanceId(v string) *StartRestoreRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *StartRestoreRequest) SetPreferredBackupTime(v string) *StartRestoreRequest {
	s.PreferredBackupTime = &v
	return s
}

func (s *StartRestoreRequest) SetBackupMode(v string) *StartRestoreRequest {
	s.BackupMode = &v
	return s
}

func (s *StartRestoreRequest) SetBackupLevel(v string) *StartRestoreRequest {
	s.BackupLevel = &v
	return s
}

func (s *StartRestoreRequest) SetBackupDbNames(v string) *StartRestoreRequest {
	s.BackupDbNames = &v
	return s
}

func (s *StartRestoreRequest) SetBackupId(v string) *StartRestoreRequest {
	s.BackupId = &v
	return s
}

type StartRestoreResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *string `json:"Result,omitempty" xml:"Result,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StartRestoreResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartRestoreResponseBody) GoString() string {
	return s.String()
}

func (s *StartRestoreResponseBody) SetRequestId(v string) *StartRestoreResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartRestoreResponseBody) SetResult(v string) *StartRestoreResponseBody {
	s.Result = &v
	return s
}

func (s *StartRestoreResponseBody) SetSuccess(v bool) *StartRestoreResponseBody {
	s.Success = &v
	return s
}

type StartRestoreResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartRestoreResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartRestoreResponse) String() string {
	return tea.Prettify(s)
}

func (s StartRestoreResponse) GoString() string {
	return s.String()
}

func (s *StartRestoreResponse) SetHeaders(v map[string]*string) *StartRestoreResponse {
	s.Headers = v
	return s
}

func (s *StartRestoreResponse) SetBody(v *StartRestoreResponseBody) *StartRestoreResponse {
	s.Body = v
	return s
}

type SubmitCleanTaskRequest struct {
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	ParentJobId    *string `json:"ParentJobId,omitempty" xml:"ParentJobId,omitempty"`
	JobId          *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ExpandType     *string `json:"ExpandType,omitempty" xml:"ExpandType,omitempty"`
}

func (s SubmitCleanTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitCleanTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitCleanTaskRequest) SetDrdsInstanceId(v string) *SubmitCleanTaskRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *SubmitCleanTaskRequest) SetDbName(v string) *SubmitCleanTaskRequest {
	s.DbName = &v
	return s
}

func (s *SubmitCleanTaskRequest) SetParentJobId(v string) *SubmitCleanTaskRequest {
	s.ParentJobId = &v
	return s
}

func (s *SubmitCleanTaskRequest) SetJobId(v string) *SubmitCleanTaskRequest {
	s.JobId = &v
	return s
}

func (s *SubmitCleanTaskRequest) SetExpandType(v string) *SubmitCleanTaskRequest {
	s.ExpandType = &v
	return s
}

type SubmitCleanTaskResponseBody struct {
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitCleanTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitCleanTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitCleanTaskResponseBody) SetSuccess(v bool) *SubmitCleanTaskResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitCleanTaskResponseBody) SetRequestId(v string) *SubmitCleanTaskResponseBody {
	s.RequestId = &v
	return s
}

type SubmitCleanTaskResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SubmitCleanTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitCleanTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitCleanTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitCleanTaskResponse) SetHeaders(v map[string]*string) *SubmitCleanTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitCleanTaskResponse) SetBody(v *SubmitCleanTaskResponseBody) *SubmitCleanTaskResponse {
	s.Body = v
	return s
}

type SubmitHotExpandPreCheckTaskRequest struct {
	DrdsInstanceId *string   `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	DbName         *string   `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DbInstType     *string   `json:"DbInstType,omitempty" xml:"DbInstType,omitempty"`
	TableList      []*string `json:"TableList,omitempty" xml:"TableList,omitempty" type:"Repeated"`
}

func (s SubmitHotExpandPreCheckTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitHotExpandPreCheckTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitHotExpandPreCheckTaskRequest) SetDrdsInstanceId(v string) *SubmitHotExpandPreCheckTaskRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *SubmitHotExpandPreCheckTaskRequest) SetDbName(v string) *SubmitHotExpandPreCheckTaskRequest {
	s.DbName = &v
	return s
}

func (s *SubmitHotExpandPreCheckTaskRequest) SetDbInstType(v string) *SubmitHotExpandPreCheckTaskRequest {
	s.DbInstType = &v
	return s
}

func (s *SubmitHotExpandPreCheckTaskRequest) SetTableList(v []*string) *SubmitHotExpandPreCheckTaskRequest {
	s.TableList = v
	return s
}

type SubmitHotExpandPreCheckTaskResponseBody struct {
	Msg       *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitHotExpandPreCheckTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitHotExpandPreCheckTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitHotExpandPreCheckTaskResponseBody) SetMsg(v string) *SubmitHotExpandPreCheckTaskResponseBody {
	s.Msg = &v
	return s
}

func (s *SubmitHotExpandPreCheckTaskResponseBody) SetRequestId(v string) *SubmitHotExpandPreCheckTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitHotExpandPreCheckTaskResponseBody) SetTaskId(v int64) *SubmitHotExpandPreCheckTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *SubmitHotExpandPreCheckTaskResponseBody) SetSuccess(v bool) *SubmitHotExpandPreCheckTaskResponseBody {
	s.Success = &v
	return s
}

type SubmitHotExpandPreCheckTaskResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SubmitHotExpandPreCheckTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitHotExpandPreCheckTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitHotExpandPreCheckTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitHotExpandPreCheckTaskResponse) SetHeaders(v map[string]*string) *SubmitHotExpandPreCheckTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitHotExpandPreCheckTaskResponse) SetBody(v *SubmitHotExpandPreCheckTaskResponseBody) *SubmitHotExpandPreCheckTaskResponse {
	s.Body = v
	return s
}

type SubmitHotExpandTaskRequest struct {
	DrdsInstanceId       *string                                           `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	DbName               *string                                           `json:"DbName,omitempty" xml:"DbName,omitempty"`
	TaskName             *string                                           `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	TaskDesc             *string                                           `json:"TaskDesc,omitempty" xml:"TaskDesc,omitempty"`
	InstanceDbMapping    []*SubmitHotExpandTaskRequestInstanceDbMapping    `json:"InstanceDbMapping,omitempty" xml:"InstanceDbMapping,omitempty" type:"Repeated"`
	Mapping              []*SubmitHotExpandTaskRequestMapping              `json:"Mapping,omitempty" xml:"Mapping,omitempty" type:"Repeated"`
	SupperAccountMapping []*SubmitHotExpandTaskRequestSupperAccountMapping `json:"SupperAccountMapping,omitempty" xml:"SupperAccountMapping,omitempty" type:"Repeated"`
	ExtendedMapping      []*SubmitHotExpandTaskRequestExtendedMapping      `json:"ExtendedMapping,omitempty" xml:"ExtendedMapping,omitempty" type:"Repeated"`
}

func (s SubmitHotExpandTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitHotExpandTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitHotExpandTaskRequest) SetDrdsInstanceId(v string) *SubmitHotExpandTaskRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *SubmitHotExpandTaskRequest) SetDbName(v string) *SubmitHotExpandTaskRequest {
	s.DbName = &v
	return s
}

func (s *SubmitHotExpandTaskRequest) SetTaskName(v string) *SubmitHotExpandTaskRequest {
	s.TaskName = &v
	return s
}

func (s *SubmitHotExpandTaskRequest) SetTaskDesc(v string) *SubmitHotExpandTaskRequest {
	s.TaskDesc = &v
	return s
}

func (s *SubmitHotExpandTaskRequest) SetInstanceDbMapping(v []*SubmitHotExpandTaskRequestInstanceDbMapping) *SubmitHotExpandTaskRequest {
	s.InstanceDbMapping = v
	return s
}

func (s *SubmitHotExpandTaskRequest) SetMapping(v []*SubmitHotExpandTaskRequestMapping) *SubmitHotExpandTaskRequest {
	s.Mapping = v
	return s
}

func (s *SubmitHotExpandTaskRequest) SetSupperAccountMapping(v []*SubmitHotExpandTaskRequestSupperAccountMapping) *SubmitHotExpandTaskRequest {
	s.SupperAccountMapping = v
	return s
}

func (s *SubmitHotExpandTaskRequest) SetExtendedMapping(v []*SubmitHotExpandTaskRequestExtendedMapping) *SubmitHotExpandTaskRequest {
	s.ExtendedMapping = v
	return s
}

type SubmitHotExpandTaskRequestInstanceDbMapping struct {
	DbList       *string `json:"DbList,omitempty" xml:"DbList,omitempty"`
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
}

func (s SubmitHotExpandTaskRequestInstanceDbMapping) String() string {
	return tea.Prettify(s)
}

func (s SubmitHotExpandTaskRequestInstanceDbMapping) GoString() string {
	return s.String()
}

func (s *SubmitHotExpandTaskRequestInstanceDbMapping) SetDbList(v string) *SubmitHotExpandTaskRequestInstanceDbMapping {
	s.DbList = &v
	return s
}

func (s *SubmitHotExpandTaskRequestInstanceDbMapping) SetInstanceName(v string) *SubmitHotExpandTaskRequestInstanceDbMapping {
	s.InstanceName = &v
	return s
}

type SubmitHotExpandTaskRequestMapping struct {
	HotDbName     *string `json:"HotDbName,omitempty" xml:"HotDbName,omitempty"`
	ShardTbValue  *string `json:"ShardTbValue,omitempty" xml:"ShardTbValue,omitempty"`
	HotTableName  *string `json:"HotTableName,omitempty" xml:"HotTableName,omitempty"`
	ShardDbValue  *string `json:"ShardDbValue,omitempty" xml:"ShardDbValue,omitempty"`
	TbShardColumn *string `json:"TbShardColumn,omitempty" xml:"TbShardColumn,omitempty"`
	DbShardColumn *string `json:"DbShardColumn,omitempty" xml:"DbShardColumn,omitempty"`
	LogicTable    *string `json:"LogicTable,omitempty" xml:"LogicTable,omitempty"`
}

func (s SubmitHotExpandTaskRequestMapping) String() string {
	return tea.Prettify(s)
}

func (s SubmitHotExpandTaskRequestMapping) GoString() string {
	return s.String()
}

func (s *SubmitHotExpandTaskRequestMapping) SetHotDbName(v string) *SubmitHotExpandTaskRequestMapping {
	s.HotDbName = &v
	return s
}

func (s *SubmitHotExpandTaskRequestMapping) SetShardTbValue(v string) *SubmitHotExpandTaskRequestMapping {
	s.ShardTbValue = &v
	return s
}

func (s *SubmitHotExpandTaskRequestMapping) SetHotTableName(v string) *SubmitHotExpandTaskRequestMapping {
	s.HotTableName = &v
	return s
}

func (s *SubmitHotExpandTaskRequestMapping) SetShardDbValue(v string) *SubmitHotExpandTaskRequestMapping {
	s.ShardDbValue = &v
	return s
}

func (s *SubmitHotExpandTaskRequestMapping) SetTbShardColumn(v string) *SubmitHotExpandTaskRequestMapping {
	s.TbShardColumn = &v
	return s
}

func (s *SubmitHotExpandTaskRequestMapping) SetDbShardColumn(v string) *SubmitHotExpandTaskRequestMapping {
	s.DbShardColumn = &v
	return s
}

func (s *SubmitHotExpandTaskRequestMapping) SetLogicTable(v string) *SubmitHotExpandTaskRequestMapping {
	s.LogicTable = &v
	return s
}

type SubmitHotExpandTaskRequestSupperAccountMapping struct {
	InstanceName   *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	SupperAccount  *string `json:"SupperAccount,omitempty" xml:"SupperAccount,omitempty"`
	SupperPassword *string `json:"SupperPassword,omitempty" xml:"SupperPassword,omitempty"`
}

func (s SubmitHotExpandTaskRequestSupperAccountMapping) String() string {
	return tea.Prettify(s)
}

func (s SubmitHotExpandTaskRequestSupperAccountMapping) GoString() string {
	return s.String()
}

func (s *SubmitHotExpandTaskRequestSupperAccountMapping) SetInstanceName(v string) *SubmitHotExpandTaskRequestSupperAccountMapping {
	s.InstanceName = &v
	return s
}

func (s *SubmitHotExpandTaskRequestSupperAccountMapping) SetSupperAccount(v string) *SubmitHotExpandTaskRequestSupperAccountMapping {
	s.SupperAccount = &v
	return s
}

func (s *SubmitHotExpandTaskRequestSupperAccountMapping) SetSupperPassword(v string) *SubmitHotExpandTaskRequestSupperAccountMapping {
	s.SupperPassword = &v
	return s
}

type SubmitHotExpandTaskRequestExtendedMapping struct {
	SrcDb         *string `json:"SrcDb,omitempty" xml:"SrcDb,omitempty"`
	SrcInstanceId *string `json:"SrcInstanceId,omitempty" xml:"SrcInstanceId,omitempty"`
}

func (s SubmitHotExpandTaskRequestExtendedMapping) String() string {
	return tea.Prettify(s)
}

func (s SubmitHotExpandTaskRequestExtendedMapping) GoString() string {
	return s.String()
}

func (s *SubmitHotExpandTaskRequestExtendedMapping) SetSrcDb(v string) *SubmitHotExpandTaskRequestExtendedMapping {
	s.SrcDb = &v
	return s
}

func (s *SubmitHotExpandTaskRequestExtendedMapping) SetSrcInstanceId(v string) *SubmitHotExpandTaskRequestExtendedMapping {
	s.SrcInstanceId = &v
	return s
}

type SubmitHotExpandTaskResponseBody struct {
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitHotExpandTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitHotExpandTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitHotExpandTaskResponseBody) SetSuccess(v bool) *SubmitHotExpandTaskResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitHotExpandTaskResponseBody) SetRequestId(v string) *SubmitHotExpandTaskResponseBody {
	s.RequestId = &v
	return s
}

type SubmitHotExpandTaskResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SubmitHotExpandTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitHotExpandTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitHotExpandTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitHotExpandTaskResponse) SetHeaders(v map[string]*string) *SubmitHotExpandTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitHotExpandTaskResponse) SetBody(v *SubmitHotExpandTaskResponseBody) *SubmitHotExpandTaskResponse {
	s.Body = v
	return s
}

type SubmitSmoothExpandPreCheckRequest struct {
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DbInstType     *string `json:"DbInstType,omitempty" xml:"DbInstType,omitempty"`
}

func (s SubmitSmoothExpandPreCheckRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitSmoothExpandPreCheckRequest) GoString() string {
	return s.String()
}

func (s *SubmitSmoothExpandPreCheckRequest) SetDrdsInstanceId(v string) *SubmitSmoothExpandPreCheckRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *SubmitSmoothExpandPreCheckRequest) SetDbName(v string) *SubmitSmoothExpandPreCheckRequest {
	s.DbName = &v
	return s
}

func (s *SubmitSmoothExpandPreCheckRequest) SetDbInstType(v string) *SubmitSmoothExpandPreCheckRequest {
	s.DbInstType = &v
	return s
}

type SubmitSmoothExpandPreCheckResponseBody struct {
	Msg       *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitSmoothExpandPreCheckResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitSmoothExpandPreCheckResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitSmoothExpandPreCheckResponseBody) SetMsg(v string) *SubmitSmoothExpandPreCheckResponseBody {
	s.Msg = &v
	return s
}

func (s *SubmitSmoothExpandPreCheckResponseBody) SetRequestId(v string) *SubmitSmoothExpandPreCheckResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitSmoothExpandPreCheckResponseBody) SetTaskId(v int64) *SubmitSmoothExpandPreCheckResponseBody {
	s.TaskId = &v
	return s
}

func (s *SubmitSmoothExpandPreCheckResponseBody) SetSuccess(v bool) *SubmitSmoothExpandPreCheckResponseBody {
	s.Success = &v
	return s
}

type SubmitSmoothExpandPreCheckResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SubmitSmoothExpandPreCheckResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitSmoothExpandPreCheckResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitSmoothExpandPreCheckResponse) GoString() string {
	return s.String()
}

func (s *SubmitSmoothExpandPreCheckResponse) SetHeaders(v map[string]*string) *SubmitSmoothExpandPreCheckResponse {
	s.Headers = v
	return s
}

func (s *SubmitSmoothExpandPreCheckResponse) SetBody(v *SubmitSmoothExpandPreCheckResponseBody) *SubmitSmoothExpandPreCheckResponse {
	s.Body = v
	return s
}

type SubmitSmoothExpandPreCheckTaskRequest struct {
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
}

func (s SubmitSmoothExpandPreCheckTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitSmoothExpandPreCheckTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitSmoothExpandPreCheckTaskRequest) SetDrdsInstanceId(v string) *SubmitSmoothExpandPreCheckTaskRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *SubmitSmoothExpandPreCheckTaskRequest) SetDbName(v string) *SubmitSmoothExpandPreCheckTaskRequest {
	s.DbName = &v
	return s
}

type SubmitSmoothExpandPreCheckTaskResponseBody struct {
	Msg       *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitSmoothExpandPreCheckTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitSmoothExpandPreCheckTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitSmoothExpandPreCheckTaskResponseBody) SetMsg(v string) *SubmitSmoothExpandPreCheckTaskResponseBody {
	s.Msg = &v
	return s
}

func (s *SubmitSmoothExpandPreCheckTaskResponseBody) SetRequestId(v string) *SubmitSmoothExpandPreCheckTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitSmoothExpandPreCheckTaskResponseBody) SetTaskId(v int64) *SubmitSmoothExpandPreCheckTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *SubmitSmoothExpandPreCheckTaskResponseBody) SetSuccess(v bool) *SubmitSmoothExpandPreCheckTaskResponseBody {
	s.Success = &v
	return s
}

type SubmitSmoothExpandPreCheckTaskResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SubmitSmoothExpandPreCheckTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitSmoothExpandPreCheckTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitSmoothExpandPreCheckTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitSmoothExpandPreCheckTaskResponse) SetHeaders(v map[string]*string) *SubmitSmoothExpandPreCheckTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitSmoothExpandPreCheckTaskResponse) SetBody(v *SubmitSmoothExpandPreCheckTaskResponseBody) *SubmitSmoothExpandPreCheckTaskResponse {
	s.Body = v
	return s
}

type SubmitSmoothExpandTaskRequest struct {
	DrdsInstanceId       *string                                           `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	DbName               *string                                           `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DbInstanceIsCreating *bool                                             `json:"DbInstanceIsCreating,omitempty" xml:"DbInstanceIsCreating,omitempty"`
	TransferTaskInfos    []*SubmitSmoothExpandTaskRequestTransferTaskInfos `json:"TransferTaskInfos,omitempty" xml:"TransferTaskInfos,omitempty" type:"Repeated"`
	RdsSuperInstances    []*SubmitSmoothExpandTaskRequestRdsSuperInstances `json:"RdsSuperInstances,omitempty" xml:"RdsSuperInstances,omitempty" type:"Repeated"`
}

func (s SubmitSmoothExpandTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitSmoothExpandTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitSmoothExpandTaskRequest) SetDrdsInstanceId(v string) *SubmitSmoothExpandTaskRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *SubmitSmoothExpandTaskRequest) SetDbName(v string) *SubmitSmoothExpandTaskRequest {
	s.DbName = &v
	return s
}

func (s *SubmitSmoothExpandTaskRequest) SetDbInstanceIsCreating(v bool) *SubmitSmoothExpandTaskRequest {
	s.DbInstanceIsCreating = &v
	return s
}

func (s *SubmitSmoothExpandTaskRequest) SetTransferTaskInfos(v []*SubmitSmoothExpandTaskRequestTransferTaskInfos) *SubmitSmoothExpandTaskRequest {
	s.TransferTaskInfos = v
	return s
}

func (s *SubmitSmoothExpandTaskRequest) SetRdsSuperInstances(v []*SubmitSmoothExpandTaskRequestRdsSuperInstances) *SubmitSmoothExpandTaskRequest {
	s.RdsSuperInstances = v
	return s
}

type SubmitSmoothExpandTaskRequestTransferTaskInfos struct {
	DbName          *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	InstanceType    *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	DstInstanceName *string `json:"DstInstanceName,omitempty" xml:"DstInstanceName,omitempty"`
	SrcInstanceName *string `json:"SrcInstanceName,omitempty" xml:"SrcInstanceName,omitempty"`
}

func (s SubmitSmoothExpandTaskRequestTransferTaskInfos) String() string {
	return tea.Prettify(s)
}

func (s SubmitSmoothExpandTaskRequestTransferTaskInfos) GoString() string {
	return s.String()
}

func (s *SubmitSmoothExpandTaskRequestTransferTaskInfos) SetDbName(v string) *SubmitSmoothExpandTaskRequestTransferTaskInfos {
	s.DbName = &v
	return s
}

func (s *SubmitSmoothExpandTaskRequestTransferTaskInfos) SetInstanceType(v string) *SubmitSmoothExpandTaskRequestTransferTaskInfos {
	s.InstanceType = &v
	return s
}

func (s *SubmitSmoothExpandTaskRequestTransferTaskInfos) SetDstInstanceName(v string) *SubmitSmoothExpandTaskRequestTransferTaskInfos {
	s.DstInstanceName = &v
	return s
}

func (s *SubmitSmoothExpandTaskRequestTransferTaskInfos) SetSrcInstanceName(v string) *SubmitSmoothExpandTaskRequestTransferTaskInfos {
	s.SrcInstanceName = &v
	return s
}

type SubmitSmoothExpandTaskRequestRdsSuperInstances struct {
	Password    *string `json:"Password,omitempty" xml:"Password,omitempty"`
	RdsName     *string `json:"RdsName,omitempty" xml:"RdsName,omitempty"`
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
}

func (s SubmitSmoothExpandTaskRequestRdsSuperInstances) String() string {
	return tea.Prettify(s)
}

func (s SubmitSmoothExpandTaskRequestRdsSuperInstances) GoString() string {
	return s.String()
}

func (s *SubmitSmoothExpandTaskRequestRdsSuperInstances) SetPassword(v string) *SubmitSmoothExpandTaskRequestRdsSuperInstances {
	s.Password = &v
	return s
}

func (s *SubmitSmoothExpandTaskRequestRdsSuperInstances) SetRdsName(v string) *SubmitSmoothExpandTaskRequestRdsSuperInstances {
	s.RdsName = &v
	return s
}

func (s *SubmitSmoothExpandTaskRequestRdsSuperInstances) SetAccountName(v string) *SubmitSmoothExpandTaskRequestRdsSuperInstances {
	s.AccountName = &v
	return s
}

type SubmitSmoothExpandTaskResponseBody struct {
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitSmoothExpandTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitSmoothExpandTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitSmoothExpandTaskResponseBody) SetSuccess(v bool) *SubmitSmoothExpandTaskResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitSmoothExpandTaskResponseBody) SetRequestId(v string) *SubmitSmoothExpandTaskResponseBody {
	s.RequestId = &v
	return s
}

type SubmitSmoothExpandTaskResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SubmitSmoothExpandTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitSmoothExpandTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitSmoothExpandTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitSmoothExpandTaskResponse) SetHeaders(v map[string]*string) *SubmitSmoothExpandTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitSmoothExpandTaskResponse) SetBody(v *SubmitSmoothExpandTaskResponseBody) *SubmitSmoothExpandTaskResponse {
	s.Body = v
	return s
}

type SubmitSqlFlashbackTaskRequest struct {
	DrdsInstanceId    *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	DbName            *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	StartTime         *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime           *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	TraceId           *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
	TableName         *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	SqlType           *string `json:"SqlType,omitempty" xml:"SqlType,omitempty"`
	SqlPk             *string `json:"SqlPk,omitempty" xml:"SqlPk,omitempty"`
	RecallType        *int32  `json:"RecallType,omitempty" xml:"RecallType,omitempty"`
	RecallRestoreType *int32  `json:"RecallRestoreType,omitempty" xml:"RecallRestoreType,omitempty"`
}

func (s SubmitSqlFlashbackTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitSqlFlashbackTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitSqlFlashbackTaskRequest) SetDrdsInstanceId(v string) *SubmitSqlFlashbackTaskRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *SubmitSqlFlashbackTaskRequest) SetDbName(v string) *SubmitSqlFlashbackTaskRequest {
	s.DbName = &v
	return s
}

func (s *SubmitSqlFlashbackTaskRequest) SetStartTime(v string) *SubmitSqlFlashbackTaskRequest {
	s.StartTime = &v
	return s
}

func (s *SubmitSqlFlashbackTaskRequest) SetEndTime(v string) *SubmitSqlFlashbackTaskRequest {
	s.EndTime = &v
	return s
}

func (s *SubmitSqlFlashbackTaskRequest) SetTraceId(v string) *SubmitSqlFlashbackTaskRequest {
	s.TraceId = &v
	return s
}

func (s *SubmitSqlFlashbackTaskRequest) SetTableName(v string) *SubmitSqlFlashbackTaskRequest {
	s.TableName = &v
	return s
}

func (s *SubmitSqlFlashbackTaskRequest) SetSqlType(v string) *SubmitSqlFlashbackTaskRequest {
	s.SqlType = &v
	return s
}

func (s *SubmitSqlFlashbackTaskRequest) SetSqlPk(v string) *SubmitSqlFlashbackTaskRequest {
	s.SqlPk = &v
	return s
}

func (s *SubmitSqlFlashbackTaskRequest) SetRecallType(v int32) *SubmitSqlFlashbackTaskRequest {
	s.RecallType = &v
	return s
}

func (s *SubmitSqlFlashbackTaskRequest) SetRecallRestoreType(v int32) *SubmitSqlFlashbackTaskRequest {
	s.RecallRestoreType = &v
	return s
}

type SubmitSqlFlashbackTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId    *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s SubmitSqlFlashbackTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitSqlFlashbackTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitSqlFlashbackTaskResponseBody) SetRequestId(v string) *SubmitSqlFlashbackTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitSqlFlashbackTaskResponseBody) SetSuccess(v bool) *SubmitSqlFlashbackTaskResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitSqlFlashbackTaskResponseBody) SetTaskId(v int64) *SubmitSqlFlashbackTaskResponseBody {
	s.TaskId = &v
	return s
}

type SubmitSqlFlashbackTaskResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SubmitSqlFlashbackTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitSqlFlashbackTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitSqlFlashbackTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitSqlFlashbackTaskResponse) SetHeaders(v map[string]*string) *SubmitSqlFlashbackTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitSqlFlashbackTaskResponse) SetBody(v *SubmitSqlFlashbackTaskResponseBody) *SubmitSqlFlashbackTaskResponse {
	s.Body = v
	return s
}

type SubmitSwitchTaskRequest struct {
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	ParentJobId    *string `json:"ParentJobId,omitempty" xml:"ParentJobId,omitempty"`
	JobId          *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ExpandType     *string `json:"ExpandType,omitempty" xml:"ExpandType,omitempty"`
}

func (s SubmitSwitchTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitSwitchTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitSwitchTaskRequest) SetDrdsInstanceId(v string) *SubmitSwitchTaskRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *SubmitSwitchTaskRequest) SetDbName(v string) *SubmitSwitchTaskRequest {
	s.DbName = &v
	return s
}

func (s *SubmitSwitchTaskRequest) SetParentJobId(v string) *SubmitSwitchTaskRequest {
	s.ParentJobId = &v
	return s
}

func (s *SubmitSwitchTaskRequest) SetJobId(v string) *SubmitSwitchTaskRequest {
	s.JobId = &v
	return s
}

func (s *SubmitSwitchTaskRequest) SetExpandType(v string) *SubmitSwitchTaskRequest {
	s.ExpandType = &v
	return s
}

type SubmitSwitchTaskResponseBody struct {
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitSwitchTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitSwitchTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitSwitchTaskResponseBody) SetSuccess(v bool) *SubmitSwitchTaskResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitSwitchTaskResponseBody) SetRequestId(v string) *SubmitSwitchTaskResponseBody {
	s.RequestId = &v
	return s
}

type SubmitSwitchTaskResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SubmitSwitchTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitSwitchTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitSwitchTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitSwitchTaskResponse) SetHeaders(v map[string]*string) *SubmitSwitchTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitSwitchTaskResponse) SetBody(v *SubmitSwitchTaskResponseBody) *SubmitSwitchTaskResponse {
	s.Body = v
	return s
}

type SwitchGlobalBroadcastTypeRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
}

func (s SwitchGlobalBroadcastTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s SwitchGlobalBroadcastTypeRequest) GoString() string {
	return s.String()
}

func (s *SwitchGlobalBroadcastTypeRequest) SetRegionId(v string) *SwitchGlobalBroadcastTypeRequest {
	s.RegionId = &v
	return s
}

func (s *SwitchGlobalBroadcastTypeRequest) SetDrdsInstanceId(v string) *SwitchGlobalBroadcastTypeRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *SwitchGlobalBroadcastTypeRequest) SetDbName(v string) *SwitchGlobalBroadcastTypeRequest {
	s.DbName = &v
	return s
}

type SwitchGlobalBroadcastTypeResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SwitchGlobalBroadcastTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SwitchGlobalBroadcastTypeResponseBody) GoString() string {
	return s.String()
}

func (s *SwitchGlobalBroadcastTypeResponseBody) SetData(v bool) *SwitchGlobalBroadcastTypeResponseBody {
	s.Data = &v
	return s
}

func (s *SwitchGlobalBroadcastTypeResponseBody) SetRequestId(v string) *SwitchGlobalBroadcastTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *SwitchGlobalBroadcastTypeResponseBody) SetSuccess(v bool) *SwitchGlobalBroadcastTypeResponseBody {
	s.Success = &v
	return s
}

type SwitchGlobalBroadcastTypeResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SwitchGlobalBroadcastTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SwitchGlobalBroadcastTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s SwitchGlobalBroadcastTypeResponse) GoString() string {
	return s.String()
}

func (s *SwitchGlobalBroadcastTypeResponse) SetHeaders(v map[string]*string) *SwitchGlobalBroadcastTypeResponse {
	s.Headers = v
	return s
}

func (s *SwitchGlobalBroadcastTypeResponse) SetBody(v *SwitchGlobalBroadcastTypeResponseBody) *SwitchGlobalBroadcastTypeResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	RegionId     *string                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceType *string                   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tag          []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	ResourceId   []*string                 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetRegionId(v string) *TagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetTag(v []*TagResourcesRequestTag) *TagResourcesRequest {
	s.Tag = v
	return s
}

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
	return s
}

type TagResourcesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s TagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *TagResourcesRequestTag) SetKey(v string) *TagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *TagResourcesRequestTag) SetValue(v string) *TagResourcesRequestTag {
	s.Value = &v
	return s
}

type TagResourcesResponseBody struct {
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *TagResourcesResponseBody) SetSuccess(v bool) *TagResourcesResponseBody {
	s.Success = &v
	return s
}

func (s *TagResourcesResponseBody) SetRequestId(v string) *TagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type TagResourcesResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponse) GoString() string {
	return s.String()
}

func (s *TagResourcesResponse) SetHeaders(v map[string]*string) *TagResourcesResponse {
	s.Headers = v
	return s
}

func (s *TagResourcesResponse) SetBody(v *TagResourcesResponseBody) *TagResourcesResponse {
	s.Body = v
	return s
}

type UntagResourcesRequest struct {
	RegionId     *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceType *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	All          *bool     `json:"All,omitempty" xml:"All,omitempty"`
	ResourceId   []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	TagKey       []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetRegionId(v string) *UntagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceId(v []*string) *UntagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UntagResourcesRequest) SetTagKey(v []*string) *UntagResourcesRequest {
	s.TagKey = v
	return s
}

type UntagResourcesResponseBody struct {
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UntagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponseBody) SetSuccess(v bool) *UntagResourcesResponseBody {
	s.Success = &v
	return s
}

func (s *UntagResourcesResponseBody) SetRequestId(v string) *UntagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type UntagResourcesResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UntagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponse) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponse) SetHeaders(v map[string]*string) *UntagResourcesResponse {
	s.Headers = v
	return s
}

func (s *UntagResourcesResponse) SetBody(v *UntagResourcesResponseBody) *UntagResourcesResponse {
	s.Body = v
	return s
}

type UpdateInstanceNetworkRequest struct {
	DrdsInstanceId         *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	SrcInstanceNetworkType *string `json:"SrcInstanceNetworkType,omitempty" xml:"SrcInstanceNetworkType,omitempty"`
	RetainClassic          *bool   `json:"RetainClassic,omitempty" xml:"RetainClassic,omitempty"`
	ClassicExpiredDays     *int32  `json:"ClassicExpiredDays,omitempty" xml:"ClassicExpiredDays,omitempty"`
}

func (s UpdateInstanceNetworkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceNetworkRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceNetworkRequest) SetDrdsInstanceId(v string) *UpdateInstanceNetworkRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *UpdateInstanceNetworkRequest) SetSrcInstanceNetworkType(v string) *UpdateInstanceNetworkRequest {
	s.SrcInstanceNetworkType = &v
	return s
}

func (s *UpdateInstanceNetworkRequest) SetRetainClassic(v bool) *UpdateInstanceNetworkRequest {
	s.RetainClassic = &v
	return s
}

func (s *UpdateInstanceNetworkRequest) SetClassicExpiredDays(v int32) *UpdateInstanceNetworkRequest {
	s.ClassicExpiredDays = &v
	return s
}

type UpdateInstanceNetworkResponseBody struct {
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateInstanceNetworkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceNetworkResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceNetworkResponseBody) SetSuccess(v bool) *UpdateInstanceNetworkResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateInstanceNetworkResponseBody) SetRequestId(v string) *UpdateInstanceNetworkResponseBody {
	s.RequestId = &v
	return s
}

type UpdateInstanceNetworkResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateInstanceNetworkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateInstanceNetworkResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceNetworkResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstanceNetworkResponse) SetHeaders(v map[string]*string) *UpdateInstanceNetworkResponse {
	s.Headers = v
	return s
}

func (s *UpdateInstanceNetworkResponse) SetBody(v *UpdateInstanceNetworkResponseBody) *UpdateInstanceNetworkResponse {
	s.Body = v
	return s
}

type UpdatePrivateRdsClassRequest struct {
	RdsClass       *string `json:"RdsClass,omitempty" xml:"RdsClass,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	DBInstanceId   *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	Storage        *string `json:"Storage,omitempty" xml:"Storage,omitempty"`
	AutoUseCoupon  *bool   `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	PrePayDuration *int32  `json:"PrePayDuration,omitempty" xml:"PrePayDuration,omitempty"`
}

func (s UpdatePrivateRdsClassRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdatePrivateRdsClassRequest) GoString() string {
	return s.String()
}

func (s *UpdatePrivateRdsClassRequest) SetRdsClass(v string) *UpdatePrivateRdsClassRequest {
	s.RdsClass = &v
	return s
}

func (s *UpdatePrivateRdsClassRequest) SetDrdsInstanceId(v string) *UpdatePrivateRdsClassRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *UpdatePrivateRdsClassRequest) SetDBInstanceId(v string) *UpdatePrivateRdsClassRequest {
	s.DBInstanceId = &v
	return s
}

func (s *UpdatePrivateRdsClassRequest) SetStorage(v string) *UpdatePrivateRdsClassRequest {
	s.Storage = &v
	return s
}

func (s *UpdatePrivateRdsClassRequest) SetAutoUseCoupon(v bool) *UpdatePrivateRdsClassRequest {
	s.AutoUseCoupon = &v
	return s
}

func (s *UpdatePrivateRdsClassRequest) SetPrePayDuration(v int32) *UpdatePrivateRdsClassRequest {
	s.PrePayDuration = &v
	return s
}

type UpdatePrivateRdsClassResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdatePrivateRdsClassResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdatePrivateRdsClassResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePrivateRdsClassResponseBody) SetData(v string) *UpdatePrivateRdsClassResponseBody {
	s.Data = &v
	return s
}

func (s *UpdatePrivateRdsClassResponseBody) SetRequestId(v string) *UpdatePrivateRdsClassResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePrivateRdsClassResponseBody) SetSuccess(v bool) *UpdatePrivateRdsClassResponseBody {
	s.Success = &v
	return s
}

type UpdatePrivateRdsClassResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdatePrivateRdsClassResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdatePrivateRdsClassResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdatePrivateRdsClassResponse) GoString() string {
	return s.String()
}

func (s *UpdatePrivateRdsClassResponse) SetHeaders(v map[string]*string) *UpdatePrivateRdsClassResponse {
	s.Headers = v
	return s
}

func (s *UpdatePrivateRdsClassResponse) SetBody(v *UpdatePrivateRdsClassResponseBody) *UpdatePrivateRdsClassResponse {
	s.Body = v
	return s
}

type UpdateResourceGroupAttributeRequest struct {
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DrdsInstanceId     *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	NewResourceGroupId *string `json:"NewResourceGroupId,omitempty" xml:"NewResourceGroupId,omitempty"`
}

func (s UpdateResourceGroupAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceGroupAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateResourceGroupAttributeRequest) SetRegionId(v string) *UpdateResourceGroupAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateResourceGroupAttributeRequest) SetDrdsInstanceId(v string) *UpdateResourceGroupAttributeRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *UpdateResourceGroupAttributeRequest) SetNewResourceGroupId(v string) *UpdateResourceGroupAttributeRequest {
	s.NewResourceGroupId = &v
	return s
}

type UpdateResourceGroupAttributeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateResourceGroupAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceGroupAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateResourceGroupAttributeResponseBody) SetRequestId(v string) *UpdateResourceGroupAttributeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateResourceGroupAttributeResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateResourceGroupAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateResourceGroupAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceGroupAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateResourceGroupAttributeResponse) SetHeaders(v map[string]*string) *UpdateResourceGroupAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateResourceGroupAttributeResponse) SetBody(v *UpdateResourceGroupAttributeResponseBody) *UpdateResourceGroupAttributeResponse {
	s.Body = v
	return s
}

type UpgradeHiStoreInstanceRequest struct {
	DrdsInstanceId    *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	HistoreInstanceId *string `json:"HistoreInstanceId,omitempty" xml:"HistoreInstanceId,omitempty"`
}

func (s UpgradeHiStoreInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeHiStoreInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpgradeHiStoreInstanceRequest) SetDrdsInstanceId(v string) *UpgradeHiStoreInstanceRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *UpgradeHiStoreInstanceRequest) SetRegionId(v string) *UpgradeHiStoreInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *UpgradeHiStoreInstanceRequest) SetHistoreInstanceId(v string) *UpgradeHiStoreInstanceRequest {
	s.HistoreInstanceId = &v
	return s
}

type UpgradeHiStoreInstanceResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpgradeHiStoreInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpgradeHiStoreInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeHiStoreInstanceResponseBody) SetData(v string) *UpgradeHiStoreInstanceResponseBody {
	s.Data = &v
	return s
}

func (s *UpgradeHiStoreInstanceResponseBody) SetRequestId(v string) *UpgradeHiStoreInstanceResponseBody {
	s.RequestId = &v
	return s
}

type UpgradeHiStoreInstanceResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpgradeHiStoreInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpgradeHiStoreInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpgradeHiStoreInstanceResponse) GoString() string {
	return s.String()
}

func (s *UpgradeHiStoreInstanceResponse) SetHeaders(v map[string]*string) *UpgradeHiStoreInstanceResponse {
	s.Headers = v
	return s
}

func (s *UpgradeHiStoreInstanceResponse) SetBody(v *UpgradeHiStoreInstanceResponseBody) *UpgradeHiStoreInstanceResponse {
	s.Body = v
	return s
}

type UpgradeInstanceVersionRequest struct {
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Rpm            *string `json:"Rpm,omitempty" xml:"Rpm,omitempty"`
}

func (s UpgradeInstanceVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeInstanceVersionRequest) GoString() string {
	return s.String()
}

func (s *UpgradeInstanceVersionRequest) SetDrdsInstanceId(v string) *UpgradeInstanceVersionRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *UpgradeInstanceVersionRequest) SetRegionId(v string) *UpgradeInstanceVersionRequest {
	s.RegionId = &v
	return s
}

func (s *UpgradeInstanceVersionRequest) SetRpm(v string) *UpgradeInstanceVersionRequest {
	s.Rpm = &v
	return s
}

type UpgradeInstanceVersionResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpgradeInstanceVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpgradeInstanceVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeInstanceVersionResponseBody) SetData(v string) *UpgradeInstanceVersionResponseBody {
	s.Data = &v
	return s
}

func (s *UpgradeInstanceVersionResponseBody) SetRequestId(v string) *UpgradeInstanceVersionResponseBody {
	s.RequestId = &v
	return s
}

type UpgradeInstanceVersionResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpgradeInstanceVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpgradeInstanceVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpgradeInstanceVersionResponse) GoString() string {
	return s.String()
}

func (s *UpgradeInstanceVersionResponse) SetHeaders(v map[string]*string) *UpgradeInstanceVersionResponse {
	s.Headers = v
	return s
}

func (s *UpgradeInstanceVersionResponse) SetBody(v *UpgradeInstanceVersionResponseBody) *UpgradeInstanceVersionResponse {
	s.Body = v
	return s
}

type ValidateShardTaskRequest struct {
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DrdsInstanceId  *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	DbName          *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	SourceTableName *string `json:"SourceTableName,omitempty" xml:"SourceTableName,omitempty"`
	TargetTableName *string `json:"TargetTableName,omitempty" xml:"TargetTableName,omitempty"`
	TaskType        *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s ValidateShardTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s ValidateShardTaskRequest) GoString() string {
	return s.String()
}

func (s *ValidateShardTaskRequest) SetRegionId(v string) *ValidateShardTaskRequest {
	s.RegionId = &v
	return s
}

func (s *ValidateShardTaskRequest) SetDrdsInstanceId(v string) *ValidateShardTaskRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *ValidateShardTaskRequest) SetDbName(v string) *ValidateShardTaskRequest {
	s.DbName = &v
	return s
}

func (s *ValidateShardTaskRequest) SetSourceTableName(v string) *ValidateShardTaskRequest {
	s.SourceTableName = &v
	return s
}

func (s *ValidateShardTaskRequest) SetTargetTableName(v string) *ValidateShardTaskRequest {
	s.TargetTableName = &v
	return s
}

func (s *ValidateShardTaskRequest) SetTaskType(v string) *ValidateShardTaskRequest {
	s.TaskType = &v
	return s
}

type ValidateShardTaskResponseBody struct {
	Success   *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	List      []*ValidateShardTaskResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s ValidateShardTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ValidateShardTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ValidateShardTaskResponseBody) SetSuccess(v bool) *ValidateShardTaskResponseBody {
	s.Success = &v
	return s
}

func (s *ValidateShardTaskResponseBody) SetRequestId(v string) *ValidateShardTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ValidateShardTaskResponseBody) SetList(v []*ValidateShardTaskResponseBodyList) *ValidateShardTaskResponseBody {
	s.List = v
	return s
}

type ValidateShardTaskResponseBodyList struct {
	Result *int32  `json:"Result,omitempty" xml:"Result,omitempty"`
	Item   *string `json:"Item,omitempty" xml:"Item,omitempty"`
}

func (s ValidateShardTaskResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s ValidateShardTaskResponseBodyList) GoString() string {
	return s.String()
}

func (s *ValidateShardTaskResponseBodyList) SetResult(v int32) *ValidateShardTaskResponseBodyList {
	s.Result = &v
	return s
}

func (s *ValidateShardTaskResponseBodyList) SetItem(v string) *ValidateShardTaskResponseBodyList {
	s.Item = &v
	return s
}

type ValidateShardTaskResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ValidateShardTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ValidateShardTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s ValidateShardTaskResponse) GoString() string {
	return s.String()
}

func (s *ValidateShardTaskResponse) SetHeaders(v map[string]*string) *ValidateShardTaskResponse {
	s.Headers = v
	return s
}

func (s *ValidateShardTaskResponse) SetBody(v *ValidateShardTaskResponseBody) *ValidateShardTaskResponse {
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
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"ap-northeast-1":              tea.String("drds.ap-southeast-1.aliyuncs.com"),
		"ap-northeast-2-pop":          tea.String("drds.ap-southeast-1.aliyuncs.com"),
		"ap-south-1":                  tea.String("drds.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-2":              tea.String("drds.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-3":              tea.String("drds.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-5":              tea.String("drds.ap-southeast-1.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("drds.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("drds.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("drds.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("drds.aliyuncs.com"),
		"cn-chengdu":                  tea.String("drds.aliyuncs.com"),
		"cn-edge-1":                   tea.String("drds.aliyuncs.com"),
		"cn-fujian":                   tea.String("drds.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("drds.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("drds.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("drds.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("drds.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("drds.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("drds.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("drds.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("drds.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("drds.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("drds.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("drds.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("drds.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("drds.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("drds.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("drds.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("drds.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("drds.aliyuncs.com"),
		"cn-wuhan":                    tea.String("drds.aliyuncs.com"),
		"cn-yushanfang":               tea.String("drds.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("drds.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("drds.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("drds.aliyuncs.com"),
		"eu-central-1":                tea.String("drds.ap-southeast-1.aliyuncs.com"),
		"eu-west-1":                   tea.String("drds.ap-southeast-1.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("drds.ap-southeast-1.aliyuncs.com"),
		"me-east-1":                   tea.String("drds.ap-southeast-1.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("drds.ap-southeast-1.aliyuncs.com"),
		"us-west-1":                   tea.String("drds.ap-southeast-1.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("drds"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) CheckDrdsDbNameWithOptions(request *CheckDrdsDbNameRequest, runtime *util.RuntimeOptions) (_result *CheckDrdsDbNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CheckDrdsDbNameResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CheckDrdsDbName"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckDrdsDbName(request *CheckDrdsDbNameRequest) (_result *CheckDrdsDbNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckDrdsDbNameResponse{}
	_body, _err := client.CheckDrdsDbNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckExpandStatusWithOptions(request *CheckExpandStatusRequest, runtime *util.RuntimeOptions) (_result *CheckExpandStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CheckExpandStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CheckExpandStatus"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckExpandStatus(request *CheckExpandStatusRequest) (_result *CheckExpandStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckExpandStatusResponse{}
	_body, _err := client.CheckExpandStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckSqlAuditEnableStatusWithOptions(request *CheckSqlAuditEnableStatusRequest, runtime *util.RuntimeOptions) (_result *CheckSqlAuditEnableStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CheckSqlAuditEnableStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CheckSqlAuditEnableStatus"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckSqlAuditEnableStatus(request *CheckSqlAuditEnableStatusRequest) (_result *CheckSqlAuditEnableStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckSqlAuditEnableStatusResponse{}
	_body, _err := client.CheckSqlAuditEnableStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDrdsDBWithOptions(request *CreateDrdsDBRequest, runtime *util.RuntimeOptions) (_result *CreateDrdsDBResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateDrdsDBResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateDrdsDB"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDrdsDB(request *CreateDrdsDBRequest) (_result *CreateDrdsDBResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDrdsDBResponse{}
	_body, _err := client.CreateDrdsDBWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDrdsInstanceWithOptions(request *CreateDrdsInstanceRequest, runtime *util.RuntimeOptions) (_result *CreateDrdsInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateDrdsInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateDrdsInstance"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDrdsInstance(request *CreateDrdsInstanceRequest) (_result *CreateDrdsInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDrdsInstanceResponse{}
	_body, _err := client.CreateDrdsInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateInstanceAccountWithOptions(request *CreateInstanceAccountRequest, runtime *util.RuntimeOptions) (_result *CreateInstanceAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateInstanceAccountResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateInstanceAccount"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateInstanceAccount(request *CreateInstanceAccountRequest) (_result *CreateInstanceAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateInstanceAccountResponse{}
	_body, _err := client.CreateInstanceAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateInstanceInternetAddressWithOptions(request *CreateInstanceInternetAddressRequest, runtime *util.RuntimeOptions) (_result *CreateInstanceInternetAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateInstanceInternetAddressResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateInstanceInternetAddress"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateInstanceInternetAddress(request *CreateInstanceInternetAddressRequest) (_result *CreateInstanceInternetAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateInstanceInternetAddressResponse{}
	_body, _err := client.CreateInstanceInternetAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateOrderForRdsWithOptions(request *CreateOrderForRdsRequest, runtime *util.RuntimeOptions) (_result *CreateOrderForRdsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateOrderForRdsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateOrderForRds"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateOrderForRds(request *CreateOrderForRdsRequest) (_result *CreateOrderForRdsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateOrderForRdsResponse{}
	_body, _err := client.CreateOrderForRdsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateShardTaskWithOptions(request *CreateShardTaskRequest, runtime *util.RuntimeOptions) (_result *CreateShardTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateShardTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateShardTask"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateShardTask(request *CreateShardTaskRequest) (_result *CreateShardTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateShardTaskResponse{}
	_body, _err := client.CreateShardTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBackMenuWithOptions(request *DescribeBackMenuRequest, runtime *util.RuntimeOptions) (_result *DescribeBackMenuResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeBackMenuResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeBackMenu"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBackMenu(request *DescribeBackMenuRequest) (_result *DescribeBackMenuResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBackMenuResponse{}
	_body, _err := client.DescribeBackMenuWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBackupDbsWithOptions(request *DescribeBackupDbsRequest, runtime *util.RuntimeOptions) (_result *DescribeBackupDbsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeBackupDbsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeBackupDbs"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBackupDbs(request *DescribeBackupDbsRequest) (_result *DescribeBackupDbsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBackupDbsResponse{}
	_body, _err := client.DescribeBackupDbsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBackupLocalWithOptions(request *DescribeBackupLocalRequest, runtime *util.RuntimeOptions) (_result *DescribeBackupLocalResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeBackupLocalResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeBackupLocal"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBackupLocal(request *DescribeBackupLocalRequest) (_result *DescribeBackupLocalResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBackupLocalResponse{}
	_body, _err := client.DescribeBackupLocalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBackupPolicyWithOptions(request *DescribeBackupPolicyRequest, runtime *util.RuntimeOptions) (_result *DescribeBackupPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeBackupPolicyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeBackupPolicy"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBackupPolicy(request *DescribeBackupPolicyRequest) (_result *DescribeBackupPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBackupPolicyResponse{}
	_body, _err := client.DescribeBackupPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBackupSetsWithOptions(request *DescribeBackupSetsRequest, runtime *util.RuntimeOptions) (_result *DescribeBackupSetsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeBackupSetsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeBackupSets"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBackupSets(request *DescribeBackupSetsRequest) (_result *DescribeBackupSetsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBackupSetsResponse{}
	_body, _err := client.DescribeBackupSetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBackupTimesWithOptions(request *DescribeBackupTimesRequest, runtime *util.RuntimeOptions) (_result *DescribeBackupTimesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeBackupTimesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeBackupTimes"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBackupTimes(request *DescribeBackupTimesRequest) (_result *DescribeBackupTimesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBackupTimesResponse{}
	_body, _err := client.DescribeBackupTimesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBroadcastTablesWithOptions(request *DescribeBroadcastTablesRequest, runtime *util.RuntimeOptions) (_result *DescribeBroadcastTablesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeBroadcastTablesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeBroadcastTables"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBroadcastTables(request *DescribeBroadcastTablesRequest) (_result *DescribeBroadcastTablesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBroadcastTablesResponse{}
	_body, _err := client.DescribeBroadcastTablesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDbInstanceDbsWithOptions(request *DescribeDbInstanceDbsRequest, runtime *util.RuntimeOptions) (_result *DescribeDbInstanceDbsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDbInstanceDbsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDbInstanceDbs"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDbInstanceDbs(request *DescribeDbInstanceDbsRequest) (_result *DescribeDbInstanceDbsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDbInstanceDbsResponse{}
	_body, _err := client.DescribeDbInstanceDbsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDbInstancesWithOptions(request *DescribeDbInstancesRequest, runtime *util.RuntimeOptions) (_result *DescribeDbInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDbInstancesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDbInstances"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDbInstances(request *DescribeDbInstancesRequest) (_result *DescribeDbInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDbInstancesResponse{}
	_body, _err := client.DescribeDbInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDrdsDBWithOptions(request *DescribeDrdsDBRequest, runtime *util.RuntimeOptions) (_result *DescribeDrdsDBResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDrdsDBResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDrdsDB"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDrdsDB(request *DescribeDrdsDBRequest) (_result *DescribeDrdsDBResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDrdsDBResponse{}
	_body, _err := client.DescribeDrdsDBWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDrdsDBClusterWithOptions(request *DescribeDrdsDBClusterRequest, runtime *util.RuntimeOptions) (_result *DescribeDrdsDBClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDrdsDBClusterResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDrdsDBCluster"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDrdsDBCluster(request *DescribeDrdsDBClusterRequest) (_result *DescribeDrdsDBClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDrdsDBClusterResponse{}
	_body, _err := client.DescribeDrdsDBClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDrdsDbInstanceWithOptions(request *DescribeDrdsDbInstanceRequest, runtime *util.RuntimeOptions) (_result *DescribeDrdsDbInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDrdsDbInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDrdsDbInstance"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDrdsDbInstance(request *DescribeDrdsDbInstanceRequest) (_result *DescribeDrdsDbInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDrdsDbInstanceResponse{}
	_body, _err := client.DescribeDrdsDbInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDrdsDbInstancesWithOptions(request *DescribeDrdsDbInstancesRequest, runtime *util.RuntimeOptions) (_result *DescribeDrdsDbInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDrdsDbInstancesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDrdsDbInstances"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDrdsDbInstances(request *DescribeDrdsDbInstancesRequest) (_result *DescribeDrdsDbInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDrdsDbInstancesResponse{}
	_body, _err := client.DescribeDrdsDbInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDrdsDBIpWhiteListWithOptions(request *DescribeDrdsDBIpWhiteListRequest, runtime *util.RuntimeOptions) (_result *DescribeDrdsDBIpWhiteListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDrdsDBIpWhiteListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDrdsDBIpWhiteList"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDrdsDBIpWhiteList(request *DescribeDrdsDBIpWhiteListRequest) (_result *DescribeDrdsDBIpWhiteListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDrdsDBIpWhiteListResponse{}
	_body, _err := client.DescribeDrdsDBIpWhiteListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDrdsDbRdsNameListWithOptions(request *DescribeDrdsDbRdsNameListRequest, runtime *util.RuntimeOptions) (_result *DescribeDrdsDbRdsNameListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDrdsDbRdsNameListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDrdsDbRdsNameList"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDrdsDbRdsNameList(request *DescribeDrdsDbRdsNameListRequest) (_result *DescribeDrdsDbRdsNameListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDrdsDbRdsNameListResponse{}
	_body, _err := client.DescribeDrdsDbRdsNameListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDrdsDBsWithOptions(request *DescribeDrdsDBsRequest, runtime *util.RuntimeOptions) (_result *DescribeDrdsDBsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDrdsDBsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDrdsDBs"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDrdsDBs(request *DescribeDrdsDBsRequest) (_result *DescribeDrdsDBsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDrdsDBsResponse{}
	_body, _err := client.DescribeDrdsDBsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDrdsDbTasksWithOptions(request *DescribeDrdsDbTasksRequest, runtime *util.RuntimeOptions) (_result *DescribeDrdsDbTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDrdsDbTasksResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDrdsDbTasks"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDrdsDbTasks(request *DescribeDrdsDbTasksRequest) (_result *DescribeDrdsDbTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDrdsDbTasksResponse{}
	_body, _err := client.DescribeDrdsDbTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDrdsInstanceWithOptions(request *DescribeDrdsInstanceRequest, runtime *util.RuntimeOptions) (_result *DescribeDrdsInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDrdsInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDrdsInstance"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDrdsInstance(request *DescribeDrdsInstanceRequest) (_result *DescribeDrdsInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDrdsInstanceResponse{}
	_body, _err := client.DescribeDrdsInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDrdsInstanceDbMonitorWithOptions(request *DescribeDrdsInstanceDbMonitorRequest, runtime *util.RuntimeOptions) (_result *DescribeDrdsInstanceDbMonitorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDrdsInstanceDbMonitorResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDrdsInstanceDbMonitor"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDrdsInstanceDbMonitor(request *DescribeDrdsInstanceDbMonitorRequest) (_result *DescribeDrdsInstanceDbMonitorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDrdsInstanceDbMonitorResponse{}
	_body, _err := client.DescribeDrdsInstanceDbMonitorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDrdsInstanceLevelTasksWithOptions(request *DescribeDrdsInstanceLevelTasksRequest, runtime *util.RuntimeOptions) (_result *DescribeDrdsInstanceLevelTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDrdsInstanceLevelTasksResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDrdsInstanceLevelTasks"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDrdsInstanceLevelTasks(request *DescribeDrdsInstanceLevelTasksRequest) (_result *DescribeDrdsInstanceLevelTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDrdsInstanceLevelTasksResponse{}
	_body, _err := client.DescribeDrdsInstanceLevelTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDrdsInstanceMonitorWithOptions(request *DescribeDrdsInstanceMonitorRequest, runtime *util.RuntimeOptions) (_result *DescribeDrdsInstanceMonitorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDrdsInstanceMonitorResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDrdsInstanceMonitor"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDrdsInstanceMonitor(request *DescribeDrdsInstanceMonitorRequest) (_result *DescribeDrdsInstanceMonitorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDrdsInstanceMonitorResponse{}
	_body, _err := client.DescribeDrdsInstanceMonitorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDrdsInstancesWithOptions(request *DescribeDrdsInstancesRequest, runtime *util.RuntimeOptions) (_result *DescribeDrdsInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDrdsInstancesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDrdsInstances"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDrdsInstances(request *DescribeDrdsInstancesRequest) (_result *DescribeDrdsInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDrdsInstancesResponse{}
	_body, _err := client.DescribeDrdsInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDrdsInstanceVersionWithOptions(request *DescribeDrdsInstanceVersionRequest, runtime *util.RuntimeOptions) (_result *DescribeDrdsInstanceVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDrdsInstanceVersionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDrdsInstanceVersion"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDrdsInstanceVersion(request *DescribeDrdsInstanceVersionRequest) (_result *DescribeDrdsInstanceVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDrdsInstanceVersionResponse{}
	_body, _err := client.DescribeDrdsInstanceVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDrdsParamsWithOptions(request *DescribeDrdsParamsRequest, runtime *util.RuntimeOptions) (_result *DescribeDrdsParamsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDrdsParamsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDrdsParams"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDrdsParams(request *DescribeDrdsParamsRequest) (_result *DescribeDrdsParamsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDrdsParamsResponse{}
	_body, _err := client.DescribeDrdsParamsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDrdsRdsInstancesWithOptions(request *DescribeDrdsRdsInstancesRequest, runtime *util.RuntimeOptions) (_result *DescribeDrdsRdsInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDrdsRdsInstancesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDrdsRdsInstances"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDrdsRdsInstances(request *DescribeDrdsRdsInstancesRequest) (_result *DescribeDrdsRdsInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDrdsRdsInstancesResponse{}
	_body, _err := client.DescribeDrdsRdsInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDrdsShardingDbsWithOptions(request *DescribeDrdsShardingDbsRequest, runtime *util.RuntimeOptions) (_result *DescribeDrdsShardingDbsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDrdsShardingDbsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDrdsShardingDbs"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDrdsShardingDbs(request *DescribeDrdsShardingDbsRequest) (_result *DescribeDrdsShardingDbsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDrdsShardingDbsResponse{}
	_body, _err := client.DescribeDrdsShardingDbsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDrdsSlowSqlsWithOptions(request *DescribeDrdsSlowSqlsRequest, runtime *util.RuntimeOptions) (_result *DescribeDrdsSlowSqlsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDrdsSlowSqlsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDrdsSlowSqls"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDrdsSlowSqls(request *DescribeDrdsSlowSqlsRequest) (_result *DescribeDrdsSlowSqlsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDrdsSlowSqlsResponse{}
	_body, _err := client.DescribeDrdsSlowSqlsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDrdsSqlAuditStatusWithOptions(request *DescribeDrdsSqlAuditStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeDrdsSqlAuditStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDrdsSqlAuditStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDrdsSqlAuditStatus"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDrdsSqlAuditStatus(request *DescribeDrdsSqlAuditStatusRequest) (_result *DescribeDrdsSqlAuditStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDrdsSqlAuditStatusResponse{}
	_body, _err := client.DescribeDrdsSqlAuditStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDrdsTasksWithOptions(request *DescribeDrdsTasksRequest, runtime *util.RuntimeOptions) (_result *DescribeDrdsTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDrdsTasksResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDrdsTasks"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDrdsTasks(request *DescribeDrdsTasksRequest) (_result *DescribeDrdsTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDrdsTasksResponse{}
	_body, _err := client.DescribeDrdsTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeExpandLogicTableInfoListWithOptions(request *DescribeExpandLogicTableInfoListRequest, runtime *util.RuntimeOptions) (_result *DescribeExpandLogicTableInfoListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeExpandLogicTableInfoListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeExpandLogicTableInfoList"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeExpandLogicTableInfoList(request *DescribeExpandLogicTableInfoListRequest) (_result *DescribeExpandLogicTableInfoListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeExpandLogicTableInfoListResponse{}
	_body, _err := client.DescribeExpandLogicTableInfoListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeHiStoreInstanceInfoWithOptions(request *DescribeHiStoreInstanceInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeHiStoreInstanceInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeHiStoreInstanceInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeHiStoreInstanceInfo"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeHiStoreInstanceInfo(request *DescribeHiStoreInstanceInfoRequest) (_result *DescribeHiStoreInstanceInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeHiStoreInstanceInfoResponse{}
	_body, _err := client.DescribeHiStoreInstanceInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeHotDbListWithOptions(request *DescribeHotDbListRequest, runtime *util.RuntimeOptions) (_result *DescribeHotDbListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeHotDbListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeHotDbList"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeHotDbList(request *DescribeHotDbListRequest) (_result *DescribeHotDbListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeHotDbListResponse{}
	_body, _err := client.DescribeHotDbListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceAccountsWithOptions(request *DescribeInstanceAccountsRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceAccountsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeInstanceAccountsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeInstanceAccounts"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstanceAccounts(request *DescribeInstanceAccountsRequest) (_result *DescribeInstanceAccountsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceAccountsResponse{}
	_body, _err := client.DescribeInstanceAccountsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceMenuSwitchWithOptions(request *DescribeInstanceMenuSwitchRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceMenuSwitchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeInstanceMenuSwitchResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeInstanceMenuSwitch"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstanceMenuSwitch(request *DescribeInstanceMenuSwitchRequest) (_result *DescribeInstanceMenuSwitchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceMenuSwitchResponse{}
	_body, _err := client.DescribeInstanceMenuSwitchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceSwitchAzoneWithOptions(request *DescribeInstanceSwitchAzoneRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceSwitchAzoneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeInstanceSwitchAzoneResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeInstanceSwitchAzone"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstanceSwitchAzone(request *DescribeInstanceSwitchAzoneRequest) (_result *DescribeInstanceSwitchAzoneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceSwitchAzoneResponse{}
	_body, _err := client.DescribeInstanceSwitchAzoneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceSwitchNetworkWithOptions(request *DescribeInstanceSwitchNetworkRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceSwitchNetworkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeInstanceSwitchNetworkResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeInstanceSwitchNetwork"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstanceSwitchNetwork(request *DescribeInstanceSwitchNetworkRequest) (_result *DescribeInstanceSwitchNetworkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceSwitchNetworkResponse{}
	_body, _err := client.DescribeInstanceSwitchNetworkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstDbLogInfoWithOptions(request *DescribeInstDbLogInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeInstDbLogInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeInstDbLogInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeInstDbLogInfo"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstDbLogInfo(request *DescribeInstDbLogInfoRequest) (_result *DescribeInstDbLogInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstDbLogInfoResponse{}
	_body, _err := client.DescribeInstDbLogInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstDbSlsInfoWithOptions(request *DescribeInstDbSlsInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeInstDbSlsInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeInstDbSlsInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeInstDbSlsInfo"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstDbSlsInfo(request *DescribeInstDbSlsInfoRequest) (_result *DescribeInstDbSlsInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstDbSlsInfoResponse{}
	_body, _err := client.DescribeInstDbSlsInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePreCheckResultWithOptions(request *DescribePreCheckResultRequest, runtime *util.RuntimeOptions) (_result *DescribePreCheckResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribePreCheckResultResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribePreCheckResult"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePreCheckResult(request *DescribePreCheckResultRequest) (_result *DescribePreCheckResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePreCheckResultResponse{}
	_body, _err := client.DescribePreCheckResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRdsCommodityWithOptions(request *DescribeRdsCommodityRequest, runtime *util.RuntimeOptions) (_result *DescribeRdsCommodityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRdsCommodityResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRdsCommodity"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRdsCommodity(request *DescribeRdsCommodityRequest) (_result *DescribeRdsCommodityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRdsCommodityResponse{}
	_body, _err := client.DescribeRdsCommodityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRDSPerformanceWithOptions(request *DescribeRDSPerformanceRequest, runtime *util.RuntimeOptions) (_result *DescribeRDSPerformanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRDSPerformanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRDSPerformance"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRDSPerformance(request *DescribeRDSPerformanceRequest) (_result *DescribeRDSPerformanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRDSPerformanceResponse{}
	_body, _err := client.DescribeRDSPerformanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRdsPerformanceSummaryWithOptions(request *DescribeRdsPerformanceSummaryRequest, runtime *util.RuntimeOptions) (_result *DescribeRdsPerformanceSummaryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRdsPerformanceSummaryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRdsPerformanceSummary"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRdsPerformanceSummary(request *DescribeRdsPerformanceSummaryRequest) (_result *DescribeRdsPerformanceSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRdsPerformanceSummaryResponse{}
	_body, _err := client.DescribeRdsPerformanceSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRdsSuperAccountInstancesWithOptions(request *DescribeRdsSuperAccountInstancesRequest, runtime *util.RuntimeOptions) (_result *DescribeRdsSuperAccountInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRdsSuperAccountInstancesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRdsSuperAccountInstances"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRdsSuperAccountInstances(request *DescribeRdsSuperAccountInstancesRequest) (_result *DescribeRdsSuperAccountInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRdsSuperAccountInstancesResponse{}
	_body, _err := client.DescribeRdsSuperAccountInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRestoreOrderWithOptions(request *DescribeRestoreOrderRequest, runtime *util.RuntimeOptions) (_result *DescribeRestoreOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRestoreOrderResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRestoreOrder"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRestoreOrder(request *DescribeRestoreOrderRequest) (_result *DescribeRestoreOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRestoreOrderResponse{}
	_body, _err := client.DescribeRestoreOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeShardTaskInfoWithOptions(request *DescribeShardTaskInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeShardTaskInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeShardTaskInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeShardTaskInfo"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeShardTaskInfo(request *DescribeShardTaskInfoRequest) (_result *DescribeShardTaskInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeShardTaskInfoResponse{}
	_body, _err := client.DescribeShardTaskInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeShardTaskListWithOptions(request *DescribeShardTaskListRequest, runtime *util.RuntimeOptions) (_result *DescribeShardTaskListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeShardTaskListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeShardTaskList"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeShardTaskList(request *DescribeShardTaskListRequest) (_result *DescribeShardTaskListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeShardTaskListResponse{}
	_body, _err := client.DescribeShardTaskListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSqlFlashbakTaskWithOptions(request *DescribeSqlFlashbakTaskRequest, runtime *util.RuntimeOptions) (_result *DescribeSqlFlashbakTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSqlFlashbakTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSqlFlashbakTask"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSqlFlashbakTask(request *DescribeSqlFlashbakTaskRequest) (_result *DescribeSqlFlashbakTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSqlFlashbakTaskResponse{}
	_body, _err := client.DescribeSqlFlashbakTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTableWithOptions(request *DescribeTableRequest, runtime *util.RuntimeOptions) (_result *DescribeTableResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeTableResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeTable"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTable(request *DescribeTableRequest) (_result *DescribeTableResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTableResponse{}
	_body, _err := client.DescribeTableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTableListByTypeWithOptions(request *DescribeTableListByTypeRequest, runtime *util.RuntimeOptions) (_result *DescribeTableListByTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeTableListByTypeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeTableListByType"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTableListByType(request *DescribeTableListByTypeRequest) (_result *DescribeTableListByTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTableListByTypeResponse{}
	_body, _err := client.DescribeTableListByTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTablesWithOptions(request *DescribeTablesRequest, runtime *util.RuntimeOptions) (_result *DescribeTablesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeTablesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeTables"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTables(request *DescribeTablesRequest) (_result *DescribeTablesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTablesResponse{}
	_body, _err := client.DescribeTablesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableSqlAuditWithOptions(request *DisableSqlAuditRequest, runtime *util.RuntimeOptions) (_result *DisableSqlAuditResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DisableSqlAuditResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DisableSqlAudit"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableSqlAudit(request *DisableSqlAuditRequest) (_result *DisableSqlAuditResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableSqlAuditResponse{}
	_body, _err := client.DisableSqlAuditWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableSqlAuditWithOptions(request *EnableSqlAuditRequest, runtime *util.RuntimeOptions) (_result *EnableSqlAuditResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &EnableSqlAuditResponse{}
	_body, _err := client.DoRPCRequest(tea.String("EnableSqlAudit"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableSqlAudit(request *EnableSqlAuditRequest) (_result *EnableSqlAuditResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableSqlAuditResponse{}
	_body, _err := client.EnableSqlAuditWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableSqlFlashbackMatchSwitchWithOptions(request *EnableSqlFlashbackMatchSwitchRequest, runtime *util.RuntimeOptions) (_result *EnableSqlFlashbackMatchSwitchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &EnableSqlFlashbackMatchSwitchResponse{}
	_body, _err := client.DoRPCRequest(tea.String("EnableSqlFlashbackMatchSwitch"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableSqlFlashbackMatchSwitch(request *EnableSqlFlashbackMatchSwitchRequest) (_result *EnableSqlFlashbackMatchSwitchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableSqlFlashbackMatchSwitchResponse{}
	_body, _err := client.EnableSqlFlashbackMatchSwitchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListTagResources"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.ListTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ManagePrivateRdsWithOptions(request *ManagePrivateRdsRequest, runtime *util.RuntimeOptions) (_result *ManagePrivateRdsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ManagePrivateRdsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ManagePrivateRds"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ManagePrivateRds(request *ManagePrivateRdsRequest) (_result *ManagePrivateRdsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ManagePrivateRdsResponse{}
	_body, _err := client.ManagePrivateRdsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDrdsInstanceDescriptionWithOptions(request *ModifyDrdsInstanceDescriptionRequest, runtime *util.RuntimeOptions) (_result *ModifyDrdsInstanceDescriptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyDrdsInstanceDescriptionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyDrdsInstanceDescription"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDrdsInstanceDescription(request *ModifyDrdsInstanceDescriptionRequest) (_result *ModifyDrdsInstanceDescriptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDrdsInstanceDescriptionResponse{}
	_body, _err := client.ModifyDrdsInstanceDescriptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDrdsIpWhiteListWithOptions(request *ModifyDrdsIpWhiteListRequest, runtime *util.RuntimeOptions) (_result *ModifyDrdsIpWhiteListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyDrdsIpWhiteListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyDrdsIpWhiteList"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDrdsIpWhiteList(request *ModifyDrdsIpWhiteListRequest) (_result *ModifyDrdsIpWhiteListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDrdsIpWhiteListResponse{}
	_body, _err := client.ModifyDrdsIpWhiteListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyRdsReadWeightWithOptions(request *ModifyRdsReadWeightRequest, runtime *util.RuntimeOptions) (_result *ModifyRdsReadWeightResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyRdsReadWeightResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyRdsReadWeight"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyRdsReadWeight(request *ModifyRdsReadWeightRequest) (_result *ModifyRdsReadWeightResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyRdsReadWeightResponse{}
	_body, _err := client.ModifyRdsReadWeightWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PutStartBackupWithOptions(request *PutStartBackupRequest, runtime *util.RuntimeOptions) (_result *PutStartBackupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &PutStartBackupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("PutStartBackup"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PutStartBackup(request *PutStartBackupRequest) (_result *PutStartBackupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PutStartBackupResponse{}
	_body, _err := client.PutStartBackupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReleaseInstanceInternetAddressWithOptions(request *ReleaseInstanceInternetAddressRequest, runtime *util.RuntimeOptions) (_result *ReleaseInstanceInternetAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ReleaseInstanceInternetAddressResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ReleaseInstanceInternetAddress"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReleaseInstanceInternetAddress(request *ReleaseInstanceInternetAddressRequest) (_result *ReleaseInstanceInternetAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReleaseInstanceInternetAddressResponse{}
	_body, _err := client.ReleaseInstanceInternetAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveBackupsSetWithOptions(request *RemoveBackupsSetRequest, runtime *util.RuntimeOptions) (_result *RemoveBackupsSetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RemoveBackupsSetResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RemoveBackupsSet"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveBackupsSet(request *RemoveBackupsSetRequest) (_result *RemoveBackupsSetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveBackupsSetResponse{}
	_body, _err := client.RemoveBackupsSetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveDrdsDbWithOptions(request *RemoveDrdsDbRequest, runtime *util.RuntimeOptions) (_result *RemoveDrdsDbResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RemoveDrdsDbResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RemoveDrdsDb"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveDrdsDb(request *RemoveDrdsDbRequest) (_result *RemoveDrdsDbResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveDrdsDbResponse{}
	_body, _err := client.RemoveDrdsDbWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveDrdsDbFailedRecordWithOptions(request *RemoveDrdsDbFailedRecordRequest, runtime *util.RuntimeOptions) (_result *RemoveDrdsDbFailedRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RemoveDrdsDbFailedRecordResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RemoveDrdsDbFailedRecord"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveDrdsDbFailedRecord(request *RemoveDrdsDbFailedRecordRequest) (_result *RemoveDrdsDbFailedRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveDrdsDbFailedRecordResponse{}
	_body, _err := client.RemoveDrdsDbFailedRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveDrdsInstanceWithOptions(request *RemoveDrdsInstanceRequest, runtime *util.RuntimeOptions) (_result *RemoveDrdsInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RemoveDrdsInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RemoveDrdsInstance"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveDrdsInstance(request *RemoveDrdsInstanceRequest) (_result *RemoveDrdsInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveDrdsInstanceResponse{}
	_body, _err := client.RemoveDrdsInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveInstanceAccountWithOptions(request *RemoveInstanceAccountRequest, runtime *util.RuntimeOptions) (_result *RemoveInstanceAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RemoveInstanceAccountResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RemoveInstanceAccount"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveInstanceAccount(request *RemoveInstanceAccountRequest) (_result *RemoveInstanceAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveInstanceAccountResponse{}
	_body, _err := client.RemoveInstanceAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetBackupLocalWithOptions(request *SetBackupLocalRequest, runtime *util.RuntimeOptions) (_result *SetBackupLocalResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetBackupLocalResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetBackupLocal"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetBackupLocal(request *SetBackupLocalRequest) (_result *SetBackupLocalResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetBackupLocalResponse{}
	_body, _err := client.SetBackupLocalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetBackupPolicyWithOptions(request *SetBackupPolicyRequest, runtime *util.RuntimeOptions) (_result *SetBackupPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetBackupPolicyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetBackupPolicy"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetBackupPolicy(request *SetBackupPolicyRequest) (_result *SetBackupPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetBackupPolicyResponse{}
	_body, _err := client.SetBackupPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetupBroadcastTablesWithOptions(request *SetupBroadcastTablesRequest, runtime *util.RuntimeOptions) (_result *SetupBroadcastTablesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetupBroadcastTablesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetupBroadcastTables"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetupBroadcastTables(request *SetupBroadcastTablesRequest) (_result *SetupBroadcastTablesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetupBroadcastTablesResponse{}
	_body, _err := client.SetupBroadcastTablesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetupDrdsParamsWithOptions(request *SetupDrdsParamsRequest, runtime *util.RuntimeOptions) (_result *SetupDrdsParamsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetupDrdsParamsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetupDrdsParams"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetupDrdsParams(request *SetupDrdsParamsRequest) (_result *SetupDrdsParamsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetupDrdsParamsResponse{}
	_body, _err := client.SetupDrdsParamsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetupTableWithOptions(request *SetupTableRequest, runtime *util.RuntimeOptions) (_result *SetupTableResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetupTableResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetupTable"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetupTable(request *SetupTableRequest) (_result *SetupTableResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetupTableResponse{}
	_body, _err := client.SetupTableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartRestoreWithOptions(request *StartRestoreRequest, runtime *util.RuntimeOptions) (_result *StartRestoreResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StartRestoreResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StartRestore"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartRestore(request *StartRestoreRequest) (_result *StartRestoreResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartRestoreResponse{}
	_body, _err := client.StartRestoreWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitCleanTaskWithOptions(request *SubmitCleanTaskRequest, runtime *util.RuntimeOptions) (_result *SubmitCleanTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SubmitCleanTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SubmitCleanTask"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitCleanTask(request *SubmitCleanTaskRequest) (_result *SubmitCleanTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitCleanTaskResponse{}
	_body, _err := client.SubmitCleanTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitHotExpandPreCheckTaskWithOptions(request *SubmitHotExpandPreCheckTaskRequest, runtime *util.RuntimeOptions) (_result *SubmitHotExpandPreCheckTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SubmitHotExpandPreCheckTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SubmitHotExpandPreCheckTask"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitHotExpandPreCheckTask(request *SubmitHotExpandPreCheckTaskRequest) (_result *SubmitHotExpandPreCheckTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitHotExpandPreCheckTaskResponse{}
	_body, _err := client.SubmitHotExpandPreCheckTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitHotExpandTaskWithOptions(request *SubmitHotExpandTaskRequest, runtime *util.RuntimeOptions) (_result *SubmitHotExpandTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SubmitHotExpandTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SubmitHotExpandTask"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitHotExpandTask(request *SubmitHotExpandTaskRequest) (_result *SubmitHotExpandTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitHotExpandTaskResponse{}
	_body, _err := client.SubmitHotExpandTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitSmoothExpandPreCheckWithOptions(request *SubmitSmoothExpandPreCheckRequest, runtime *util.RuntimeOptions) (_result *SubmitSmoothExpandPreCheckResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SubmitSmoothExpandPreCheckResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SubmitSmoothExpandPreCheck"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitSmoothExpandPreCheck(request *SubmitSmoothExpandPreCheckRequest) (_result *SubmitSmoothExpandPreCheckResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitSmoothExpandPreCheckResponse{}
	_body, _err := client.SubmitSmoothExpandPreCheckWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitSmoothExpandPreCheckTaskWithOptions(request *SubmitSmoothExpandPreCheckTaskRequest, runtime *util.RuntimeOptions) (_result *SubmitSmoothExpandPreCheckTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SubmitSmoothExpandPreCheckTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SubmitSmoothExpandPreCheckTask"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitSmoothExpandPreCheckTask(request *SubmitSmoothExpandPreCheckTaskRequest) (_result *SubmitSmoothExpandPreCheckTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitSmoothExpandPreCheckTaskResponse{}
	_body, _err := client.SubmitSmoothExpandPreCheckTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitSmoothExpandTaskWithOptions(request *SubmitSmoothExpandTaskRequest, runtime *util.RuntimeOptions) (_result *SubmitSmoothExpandTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SubmitSmoothExpandTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SubmitSmoothExpandTask"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitSmoothExpandTask(request *SubmitSmoothExpandTaskRequest) (_result *SubmitSmoothExpandTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitSmoothExpandTaskResponse{}
	_body, _err := client.SubmitSmoothExpandTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitSqlFlashbackTaskWithOptions(request *SubmitSqlFlashbackTaskRequest, runtime *util.RuntimeOptions) (_result *SubmitSqlFlashbackTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SubmitSqlFlashbackTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SubmitSqlFlashbackTask"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitSqlFlashbackTask(request *SubmitSqlFlashbackTaskRequest) (_result *SubmitSqlFlashbackTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitSqlFlashbackTaskResponse{}
	_body, _err := client.SubmitSqlFlashbackTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitSwitchTaskWithOptions(request *SubmitSwitchTaskRequest, runtime *util.RuntimeOptions) (_result *SubmitSwitchTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SubmitSwitchTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SubmitSwitchTask"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitSwitchTask(request *SubmitSwitchTaskRequest) (_result *SubmitSwitchTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitSwitchTaskResponse{}
	_body, _err := client.SubmitSwitchTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SwitchGlobalBroadcastTypeWithOptions(request *SwitchGlobalBroadcastTypeRequest, runtime *util.RuntimeOptions) (_result *SwitchGlobalBroadcastTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SwitchGlobalBroadcastTypeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SwitchGlobalBroadcastType"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SwitchGlobalBroadcastType(request *SwitchGlobalBroadcastTypeRequest) (_result *SwitchGlobalBroadcastTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SwitchGlobalBroadcastTypeResponse{}
	_body, _err := client.SwitchGlobalBroadcastTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *util.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("TagResources"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TagResources(request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TagResourcesResponse{}
	_body, _err := client.TagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, runtime *util.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UntagResources"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UntagResources(request *UntagResourcesRequest) (_result *UntagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UntagResourcesResponse{}
	_body, _err := client.UntagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateInstanceNetworkWithOptions(request *UpdateInstanceNetworkRequest, runtime *util.RuntimeOptions) (_result *UpdateInstanceNetworkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateInstanceNetworkResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateInstanceNetwork"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateInstanceNetwork(request *UpdateInstanceNetworkRequest) (_result *UpdateInstanceNetworkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateInstanceNetworkResponse{}
	_body, _err := client.UpdateInstanceNetworkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdatePrivateRdsClassWithOptions(request *UpdatePrivateRdsClassRequest, runtime *util.RuntimeOptions) (_result *UpdatePrivateRdsClassResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdatePrivateRdsClassResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdatePrivateRdsClass"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdatePrivateRdsClass(request *UpdatePrivateRdsClassRequest) (_result *UpdatePrivateRdsClassResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdatePrivateRdsClassResponse{}
	_body, _err := client.UpdatePrivateRdsClassWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateResourceGroupAttributeWithOptions(request *UpdateResourceGroupAttributeRequest, runtime *util.RuntimeOptions) (_result *UpdateResourceGroupAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateResourceGroupAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateResourceGroupAttribute"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateResourceGroupAttribute(request *UpdateResourceGroupAttributeRequest) (_result *UpdateResourceGroupAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateResourceGroupAttributeResponse{}
	_body, _err := client.UpdateResourceGroupAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpgradeHiStoreInstanceWithOptions(request *UpgradeHiStoreInstanceRequest, runtime *util.RuntimeOptions) (_result *UpgradeHiStoreInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpgradeHiStoreInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpgradeHiStoreInstance"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpgradeHiStoreInstance(request *UpgradeHiStoreInstanceRequest) (_result *UpgradeHiStoreInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpgradeHiStoreInstanceResponse{}
	_body, _err := client.UpgradeHiStoreInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpgradeInstanceVersionWithOptions(request *UpgradeInstanceVersionRequest, runtime *util.RuntimeOptions) (_result *UpgradeInstanceVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpgradeInstanceVersionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpgradeInstanceVersion"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpgradeInstanceVersion(request *UpgradeInstanceVersionRequest) (_result *UpgradeInstanceVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpgradeInstanceVersionResponse{}
	_body, _err := client.UpgradeInstanceVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ValidateShardTaskWithOptions(request *ValidateShardTaskRequest, runtime *util.RuntimeOptions) (_result *ValidateShardTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ValidateShardTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ValidateShardTask"), tea.String("2019-01-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ValidateShardTask(request *ValidateShardTaskRequest) (_result *ValidateShardTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ValidateShardTaskResponse{}
	_body, _err := client.ValidateShardTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
