// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ChangeAccountPasswordRequest struct {
	AccountName    *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	Password       *string `json:"Password,omitempty" xml:"Password,omitempty"`
}

func (s ChangeAccountPasswordRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangeAccountPasswordRequest) GoString() string {
	return s.String()
}

func (s *ChangeAccountPasswordRequest) SetAccountName(v string) *ChangeAccountPasswordRequest {
	s.AccountName = &v
	return s
}

func (s *ChangeAccountPasswordRequest) SetDrdsInstanceId(v string) *ChangeAccountPasswordRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *ChangeAccountPasswordRequest) SetPassword(v string) *ChangeAccountPasswordRequest {
	s.Password = &v
	return s
}

type ChangeAccountPasswordResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ChangeAccountPasswordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChangeAccountPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeAccountPasswordResponseBody) SetRequestId(v string) *ChangeAccountPasswordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeAccountPasswordResponseBody) SetSuccess(v bool) *ChangeAccountPasswordResponseBody {
	s.Success = &v
	return s
}

type ChangeAccountPasswordResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ChangeAccountPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ChangeAccountPasswordResponse) String() string {
	return tea.Prettify(s)
}

func (s ChangeAccountPasswordResponse) GoString() string {
	return s.String()
}

func (s *ChangeAccountPasswordResponse) SetHeaders(v map[string]*string) *ChangeAccountPasswordResponse {
	s.Headers = v
	return s
}

func (s *ChangeAccountPasswordResponse) SetStatusCode(v int32) *ChangeAccountPasswordResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeAccountPasswordResponse) SetBody(v *ChangeAccountPasswordResponseBody) *ChangeAccountPasswordResponse {
	s.Body = v
	return s
}

type ChangeInstanceAzoneRequest struct {
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	DrdsRegionId   *string `json:"DrdsRegionId,omitempty" xml:"DrdsRegionId,omitempty"`
	OriginAzoneId  *string `json:"OriginAzoneId,omitempty" xml:"OriginAzoneId,omitempty"`
	TargetAzoneId  *string `json:"TargetAzoneId,omitempty" xml:"TargetAzoneId,omitempty"`
}

func (s ChangeInstanceAzoneRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangeInstanceAzoneRequest) GoString() string {
	return s.String()
}

func (s *ChangeInstanceAzoneRequest) SetDrdsInstanceId(v string) *ChangeInstanceAzoneRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *ChangeInstanceAzoneRequest) SetDrdsRegionId(v string) *ChangeInstanceAzoneRequest {
	s.DrdsRegionId = &v
	return s
}

func (s *ChangeInstanceAzoneRequest) SetOriginAzoneId(v string) *ChangeInstanceAzoneRequest {
	s.OriginAzoneId = &v
	return s
}

func (s *ChangeInstanceAzoneRequest) SetTargetAzoneId(v string) *ChangeInstanceAzoneRequest {
	s.TargetAzoneId = &v
	return s
}

type ChangeInstanceAzoneResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ChangeInstanceAzoneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChangeInstanceAzoneResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeInstanceAzoneResponseBody) SetRequestId(v string) *ChangeInstanceAzoneResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeInstanceAzoneResponseBody) SetSuccess(v bool) *ChangeInstanceAzoneResponseBody {
	s.Success = &v
	return s
}

type ChangeInstanceAzoneResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ChangeInstanceAzoneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ChangeInstanceAzoneResponse) String() string {
	return tea.Prettify(s)
}

func (s ChangeInstanceAzoneResponse) GoString() string {
	return s.String()
}

func (s *ChangeInstanceAzoneResponse) SetHeaders(v map[string]*string) *ChangeInstanceAzoneResponse {
	s.Headers = v
	return s
}

func (s *ChangeInstanceAzoneResponse) SetStatusCode(v int32) *ChangeInstanceAzoneResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeInstanceAzoneResponse) SetBody(v *ChangeInstanceAzoneResponseBody) *ChangeInstanceAzoneResponse {
	s.Body = v
	return s
}

type CheckDrdsDbNameRequest struct {
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s CheckDrdsDbNameRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckDrdsDbNameRequest) GoString() string {
	return s.String()
}

func (s *CheckDrdsDbNameRequest) SetDbName(v string) *CheckDrdsDbNameRequest {
	s.DbName = &v
	return s
}

func (s *CheckDrdsDbNameRequest) SetDrdsInstanceId(v string) *CheckDrdsDbNameRequest {
	s.DrdsInstanceId = &v
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CheckDrdsDbNameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CheckDrdsDbNameResponse) SetStatusCode(v int32) *CheckDrdsDbNameResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckDrdsDbNameResponse) SetBody(v *CheckDrdsDbNameResponseBody) *CheckDrdsDbNameResponse {
	s.Body = v
	return s
}

type CheckExpandStatusRequest struct {
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s CheckExpandStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckExpandStatusRequest) GoString() string {
	return s.String()
}

func (s *CheckExpandStatusRequest) SetDbName(v string) *CheckExpandStatusRequest {
	s.DbName = &v
	return s
}

func (s *CheckExpandStatusRequest) SetDrdsInstanceId(v string) *CheckExpandStatusRequest {
	s.DrdsInstanceId = &v
	return s
}

type CheckExpandStatusResponseBody struct {
	Data      *CheckExpandStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckExpandStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckExpandStatusResponseBody) GoString() string {
	return s.String()
}

func (s *CheckExpandStatusResponseBody) SetData(v *CheckExpandStatusResponseBodyData) *CheckExpandStatusResponseBody {
	s.Data = v
	return s
}

func (s *CheckExpandStatusResponseBody) SetRequestId(v string) *CheckExpandStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckExpandStatusResponseBody) SetSuccess(v bool) *CheckExpandStatusResponseBody {
	s.Success = &v
	return s
}

type CheckExpandStatusResponseBodyData struct {
	IsActive *bool   `json:"IsActive,omitempty" xml:"IsActive,omitempty"`
	Msg      *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
}

func (s CheckExpandStatusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CheckExpandStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *CheckExpandStatusResponseBodyData) SetIsActive(v bool) *CheckExpandStatusResponseBodyData {
	s.IsActive = &v
	return s
}

func (s *CheckExpandStatusResponseBodyData) SetMsg(v string) *CheckExpandStatusResponseBodyData {
	s.Msg = &v
	return s
}

type CheckExpandStatusResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CheckExpandStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CheckExpandStatusResponse) SetStatusCode(v int32) *CheckExpandStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckExpandStatusResponse) SetBody(v *CheckExpandStatusResponseBody) *CheckExpandStatusResponse {
	s.Body = v
	return s
}

type CheckSqlAuditEnableStatusRequest struct {
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s CheckSqlAuditEnableStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckSqlAuditEnableStatusRequest) GoString() string {
	return s.String()
}

func (s *CheckSqlAuditEnableStatusRequest) SetDbName(v string) *CheckSqlAuditEnableStatusRequest {
	s.DbName = &v
	return s
}

func (s *CheckSqlAuditEnableStatusRequest) SetDrdsInstanceId(v string) *CheckSqlAuditEnableStatusRequest {
	s.DrdsInstanceId = &v
	return s
}

type CheckSqlAuditEnableStatusResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckSqlAuditEnableStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckSqlAuditEnableStatusResponseBody) GoString() string {
	return s.String()
}

func (s *CheckSqlAuditEnableStatusResponseBody) SetRequestId(v string) *CheckSqlAuditEnableStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckSqlAuditEnableStatusResponseBody) SetStatus(v string) *CheckSqlAuditEnableStatusResponseBody {
	s.Status = &v
	return s
}

func (s *CheckSqlAuditEnableStatusResponseBody) SetSuccess(v bool) *CheckSqlAuditEnableStatusResponseBody {
	s.Success = &v
	return s
}

type CheckSqlAuditEnableStatusResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CheckSqlAuditEnableStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CheckSqlAuditEnableStatusResponse) SetStatusCode(v int32) *CheckSqlAuditEnableStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckSqlAuditEnableStatusResponse) SetBody(v *CheckSqlAuditEnableStatusResponseBody) *CheckSqlAuditEnableStatusResponse {
	s.Body = v
	return s
}

type CreateDrdsDBRequest struct {
	AccountName          *string                               `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	DbInstType           *string                               `json:"DbInstType,omitempty" xml:"DbInstType,omitempty"`
	DbInstanceIsCreating *bool                                 `json:"DbInstanceIsCreating,omitempty" xml:"DbInstanceIsCreating,omitempty"`
	DbName               *string                               `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DrdsInstanceId       *string                               `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	Encode               *string                               `json:"Encode,omitempty" xml:"Encode,omitempty"`
	InstDbName           []*CreateDrdsDBRequestInstDbName      `json:"InstDbName,omitempty" xml:"InstDbName,omitempty" type:"Repeated"`
	Password             *string                               `json:"Password,omitempty" xml:"Password,omitempty"`
	RdsInstance          []*string                             `json:"RdsInstance,omitempty" xml:"RdsInstance,omitempty" type:"Repeated"`
	RdsSuperAccount      []*CreateDrdsDBRequestRdsSuperAccount `json:"RdsSuperAccount,omitempty" xml:"RdsSuperAccount,omitempty" type:"Repeated"`
	Type                 *string                               `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateDrdsDBRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDrdsDBRequest) GoString() string {
	return s.String()
}

func (s *CreateDrdsDBRequest) SetAccountName(v string) *CreateDrdsDBRequest {
	s.AccountName = &v
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

func (s *CreateDrdsDBRequest) SetDbName(v string) *CreateDrdsDBRequest {
	s.DbName = &v
	return s
}

func (s *CreateDrdsDBRequest) SetDrdsInstanceId(v string) *CreateDrdsDBRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *CreateDrdsDBRequest) SetEncode(v string) *CreateDrdsDBRequest {
	s.Encode = &v
	return s
}

func (s *CreateDrdsDBRequest) SetInstDbName(v []*CreateDrdsDBRequestInstDbName) *CreateDrdsDBRequest {
	s.InstDbName = v
	return s
}

func (s *CreateDrdsDBRequest) SetPassword(v string) *CreateDrdsDBRequest {
	s.Password = &v
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

func (s *CreateDrdsDBRequest) SetType(v string) *CreateDrdsDBRequest {
	s.Type = &v
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

type CreateDrdsDBRequestRdsSuperAccount struct {
	AccountName  *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	DbInstanceId *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
	Password     *string `json:"Password,omitempty" xml:"Password,omitempty"`
}

func (s CreateDrdsDBRequestRdsSuperAccount) String() string {
	return tea.Prettify(s)
}

func (s CreateDrdsDBRequestRdsSuperAccount) GoString() string {
	return s.String()
}

func (s *CreateDrdsDBRequestRdsSuperAccount) SetAccountName(v string) *CreateDrdsDBRequestRdsSuperAccount {
	s.AccountName = &v
	return s
}

func (s *CreateDrdsDBRequestRdsSuperAccount) SetDbInstanceId(v string) *CreateDrdsDBRequestRdsSuperAccount {
	s.DbInstanceId = &v
	return s
}

func (s *CreateDrdsDBRequestRdsSuperAccount) SetPassword(v string) *CreateDrdsDBRequestRdsSuperAccount {
	s.Password = &v
	return s
}

type CreateDrdsDBResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDrdsDBResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDrdsDBResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDrdsDBResponseBody) SetRequestId(v string) *CreateDrdsDBResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDrdsDBResponseBody) SetSuccess(v bool) *CreateDrdsDBResponseBody {
	s.Success = &v
	return s
}

type CreateDrdsDBResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateDrdsDBResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateDrdsDBResponse) SetStatusCode(v int32) *CreateDrdsDBResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDrdsDBResponse) SetBody(v *CreateDrdsDBResponseBody) *CreateDrdsDBResponse {
	s.Body = v
	return s
}

type CreateDrdsInstanceRequest struct {
	ClientToken     *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Duration        *int32  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	InstanceSeries  *string `json:"InstanceSeries,omitempty" xml:"InstanceSeries,omitempty"`
	IsAutoRenew     *bool   `json:"IsAutoRenew,omitempty" xml:"IsAutoRenew,omitempty"`
	MasterInstId    *string `json:"MasterInstId,omitempty" xml:"MasterInstId,omitempty"`
	MySQLVersion    *int32  `json:"MySQLVersion,omitempty" xml:"MySQLVersion,omitempty"`
	PayType         *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	PricingCycle    *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	Quantity        *int32  `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Specification   *string `json:"Specification,omitempty" xml:"Specification,omitempty"`
	Type            *string `json:"Type,omitempty" xml:"Type,omitempty"`
	VpcId           *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VswitchId       *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	ZoneId          *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	IsHa            *bool   `json:"isHa,omitempty" xml:"isHa,omitempty"`
}

func (s CreateDrdsInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDrdsInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateDrdsInstanceRequest) SetClientToken(v string) *CreateDrdsInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDrdsInstanceRequest) SetDescription(v string) *CreateDrdsInstanceRequest {
	s.Description = &v
	return s
}

func (s *CreateDrdsInstanceRequest) SetDuration(v int32) *CreateDrdsInstanceRequest {
	s.Duration = &v
	return s
}

func (s *CreateDrdsInstanceRequest) SetInstanceSeries(v string) *CreateDrdsInstanceRequest {
	s.InstanceSeries = &v
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

func (s *CreateDrdsInstanceRequest) SetPayType(v string) *CreateDrdsInstanceRequest {
	s.PayType = &v
	return s
}

func (s *CreateDrdsInstanceRequest) SetPricingCycle(v string) *CreateDrdsInstanceRequest {
	s.PricingCycle = &v
	return s
}

func (s *CreateDrdsInstanceRequest) SetQuantity(v int32) *CreateDrdsInstanceRequest {
	s.Quantity = &v
	return s
}

func (s *CreateDrdsInstanceRequest) SetRegionId(v string) *CreateDrdsInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDrdsInstanceRequest) SetResourceGroupId(v string) *CreateDrdsInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateDrdsInstanceRequest) SetSpecification(v string) *CreateDrdsInstanceRequest {
	s.Specification = &v
	return s
}

func (s *CreateDrdsInstanceRequest) SetType(v string) *CreateDrdsInstanceRequest {
	s.Type = &v
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

func (s *CreateDrdsInstanceRequest) SetZoneId(v string) *CreateDrdsInstanceRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateDrdsInstanceRequest) SetIsHa(v bool) *CreateDrdsInstanceRequest {
	s.IsHa = &v
	return s
}

type CreateDrdsInstanceResponseBody struct {
	Data      *CreateDrdsInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDrdsInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDrdsInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDrdsInstanceResponseBody) SetData(v *CreateDrdsInstanceResponseBodyData) *CreateDrdsInstanceResponseBody {
	s.Data = v
	return s
}

func (s *CreateDrdsInstanceResponseBody) SetRequestId(v string) *CreateDrdsInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDrdsInstanceResponseBody) SetSuccess(v bool) *CreateDrdsInstanceResponseBody {
	s.Success = &v
	return s
}

type CreateDrdsInstanceResponseBodyData struct {
	DrdsInstanceIdList *CreateDrdsInstanceResponseBodyDataDrdsInstanceIdList `json:"DrdsInstanceIdList,omitempty" xml:"DrdsInstanceIdList,omitempty" type:"Struct"`
	OrderId            *int64                                                `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s CreateDrdsInstanceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateDrdsInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateDrdsInstanceResponseBodyData) SetDrdsInstanceIdList(v *CreateDrdsInstanceResponseBodyDataDrdsInstanceIdList) *CreateDrdsInstanceResponseBodyData {
	s.DrdsInstanceIdList = v
	return s
}

func (s *CreateDrdsInstanceResponseBodyData) SetOrderId(v int64) *CreateDrdsInstanceResponseBodyData {
	s.OrderId = &v
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateDrdsInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateDrdsInstanceResponse) SetStatusCode(v int32) *CreateDrdsInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDrdsInstanceResponse) SetBody(v *CreateDrdsInstanceResponseBody) *CreateDrdsInstanceResponse {
	s.Body = v
	return s
}

type CreateInstanceAccountRequest struct {
	AccountName    *string                                    `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	DbPrivilege    []*CreateInstanceAccountRequestDbPrivilege `json:"DbPrivilege,omitempty" xml:"DbPrivilege,omitempty" type:"Repeated"`
	DrdsInstanceId *string                                    `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	Password       *string                                    `json:"Password,omitempty" xml:"Password,omitempty"`
}

func (s CreateInstanceAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceAccountRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceAccountRequest) SetAccountName(v string) *CreateInstanceAccountRequest {
	s.AccountName = &v
	return s
}

func (s *CreateInstanceAccountRequest) SetDbPrivilege(v []*CreateInstanceAccountRequestDbPrivilege) *CreateInstanceAccountRequest {
	s.DbPrivilege = v
	return s
}

func (s *CreateInstanceAccountRequest) SetDrdsInstanceId(v string) *CreateInstanceAccountRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *CreateInstanceAccountRequest) SetPassword(v string) *CreateInstanceAccountRequest {
	s.Password = &v
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
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateInstanceAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceAccountResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceAccountResponseBody) SetRequestId(v string) *CreateInstanceAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInstanceAccountResponseBody) SetSuccess(v bool) *CreateInstanceAccountResponseBody {
	s.Success = &v
	return s
}

type CreateInstanceAccountResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateInstanceAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateInstanceAccountResponse) SetStatusCode(v int32) *CreateInstanceAccountResponse {
	s.StatusCode = &v
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
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateInstanceInternetAddressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateInstanceInternetAddressResponse) SetStatusCode(v int32) *CreateInstanceInternetAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateInstanceInternetAddressResponse) SetBody(v *CreateInstanceInternetAddressResponseBody) *CreateInstanceInternetAddressResponse {
	s.Body = v
	return s
}

type CreateOrderForRdsRequest struct {
	Params   *string `json:"Params,omitempty" xml:"Params,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateOrderForRdsRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateOrderForRdsRequest) GoString() string {
	return s.String()
}

func (s *CreateOrderForRdsRequest) SetParams(v string) *CreateOrderForRdsRequest {
	s.Params = &v
	return s
}

func (s *CreateOrderForRdsRequest) SetRegionId(v string) *CreateOrderForRdsRequest {
	s.RegionId = &v
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateOrderForRdsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateOrderForRdsResponse) SetStatusCode(v int32) *CreateOrderForRdsResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOrderForRdsResponse) SetBody(v *CreateOrderForRdsResponseBody) *CreateOrderForRdsResponse {
	s.Body = v
	return s
}

type CreateShardTaskRequest struct {
	DbName          *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DrdsInstanceId  *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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

func (s *CreateShardTaskRequest) SetDbName(v string) *CreateShardTaskRequest {
	s.DbName = &v
	return s
}

func (s *CreateShardTaskRequest) SetDrdsInstanceId(v string) *CreateShardTaskRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *CreateShardTaskRequest) SetRegionId(v string) *CreateShardTaskRequest {
	s.RegionId = &v
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateShardTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateShardTaskResponse) SetStatusCode(v int32) *CreateShardTaskResponse {
	s.StatusCode = &v
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
	List      *DescribeBackMenuResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Struct"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeBackMenuResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackMenuResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackMenuResponseBody) SetList(v *DescribeBackMenuResponseBodyList) *DescribeBackMenuResponseBody {
	s.List = v
	return s
}

func (s *DescribeBackMenuResponseBody) SetRequestId(v string) *DescribeBackMenuResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackMenuResponseBody) SetSuccess(v bool) *DescribeBackMenuResponseBody {
	s.Success = &v
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
	MenuName *string `json:"MenuName,omitempty" xml:"MenuName,omitempty"`
	Support  *bool   `json:"Support,omitempty" xml:"Support,omitempty"`
}

func (s DescribeBackMenuResponseBodyListList) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackMenuResponseBodyListList) GoString() string {
	return s.String()
}

func (s *DescribeBackMenuResponseBodyListList) SetMenuName(v string) *DescribeBackMenuResponseBodyListList {
	s.MenuName = &v
	return s
}

func (s *DescribeBackMenuResponseBodyListList) SetSupport(v bool) *DescribeBackMenuResponseBodyListList {
	s.Support = &v
	return s
}

type DescribeBackMenuResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeBackMenuResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeBackMenuResponse) SetStatusCode(v int32) *DescribeBackMenuResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBackMenuResponse) SetBody(v *DescribeBackMenuResponseBody) *DescribeBackMenuResponse {
	s.Body = v
	return s
}

type DescribeBackupDbsRequest struct {
	BackupId             *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	DrdsInstanceId       *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	PreferredRestoreTime *string `json:"PreferredRestoreTime,omitempty" xml:"PreferredRestoreTime,omitempty"`
}

func (s DescribeBackupDbsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupDbsRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupDbsRequest) SetBackupId(v string) *DescribeBackupDbsRequest {
	s.BackupId = &v
	return s
}

func (s *DescribeBackupDbsRequest) SetDrdsInstanceId(v string) *DescribeBackupDbsRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeBackupDbsRequest) SetPreferredRestoreTime(v string) *DescribeBackupDbsRequest {
	s.PreferredRestoreTime = &v
	return s
}

type DescribeBackupDbsResponseBody struct {
	DbNames   *DescribeBackupDbsResponseBodyDbNames `json:"DbNames,omitempty" xml:"DbNames,omitempty" type:"Struct"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeBackupDbsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupDbsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupDbsResponseBody) SetDbNames(v *DescribeBackupDbsResponseBodyDbNames) *DescribeBackupDbsResponseBody {
	s.DbNames = v
	return s
}

func (s *DescribeBackupDbsResponseBody) SetRequestId(v string) *DescribeBackupDbsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupDbsResponseBody) SetSuccess(v bool) *DescribeBackupDbsResponseBody {
	s.Success = &v
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeBackupDbsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeBackupDbsResponse) SetStatusCode(v int32) *DescribeBackupDbsResponse {
	s.StatusCode = &v
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
	BackupPolicyDO *DescribeBackupLocalResponseBodyBackupPolicyDO `json:"BackupPolicyDO,omitempty" xml:"BackupPolicyDO,omitempty" type:"Struct"`
	RequestId      *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeBackupLocalResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupLocalResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupLocalResponseBody) SetBackupPolicyDO(v *DescribeBackupLocalResponseBodyBackupPolicyDO) *DescribeBackupLocalResponseBody {
	s.BackupPolicyDO = v
	return s
}

func (s *DescribeBackupLocalResponseBody) SetRequestId(v string) *DescribeBackupLocalResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupLocalResponseBody) SetSuccess(v bool) *DescribeBackupLocalResponseBody {
	s.Success = &v
	return s
}

type DescribeBackupLocalResponseBodyBackupPolicyDO struct {
	BackupAppName             *string `json:"BackupAppName,omitempty" xml:"BackupAppName,omitempty"`
	BackupDbName              *string `json:"BackupDbName,omitempty" xml:"BackupDbName,omitempty"`
	BackupLevel               *string `json:"BackupLevel,omitempty" xml:"BackupLevel,omitempty"`
	BackupLog                 *string `json:"BackupLog,omitempty" xml:"BackupLog,omitempty"`
	BackupMode                *string `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	BackupPolicyMode          *string `json:"BackupPolicyMode,omitempty" xml:"BackupPolicyMode,omitempty"`
	BackupRetentionPeriod     *int64  `json:"BackupRetentionPeriod,omitempty" xml:"BackupRetentionPeriod,omitempty"`
	BackupType                *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	DataBackupRetentionPeriod *int64  `json:"DataBackupRetentionPeriod,omitempty" xml:"DataBackupRetentionPeriod,omitempty"`
	GmtCreate                 *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified               *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	HighSpaceUsageProtection  *int64  `json:"HighSpaceUsageProtection,omitempty" xml:"HighSpaceUsageProtection,omitempty"`
	LocalLogRetentionHours    *int64  `json:"LocalLogRetentionHours,omitempty" xml:"LocalLogRetentionHours,omitempty"`
	LocalLogRetentionSpace    *int64  `json:"LocalLogRetentionSpace,omitempty" xml:"LocalLogRetentionSpace,omitempty"`
	LogBackupRetentionPeriod  *int64  `json:"LogBackupRetentionPeriod,omitempty" xml:"LogBackupRetentionPeriod,omitempty"`
	NextBackupActuallyTime    *string `json:"NextBackupActuallyTime,omitempty" xml:"NextBackupActuallyTime,omitempty"`
	PreferredBackupPeriod     *string `json:"PreferredBackupPeriod,omitempty" xml:"PreferredBackupPeriod,omitempty"`
	PreferredBackupTime       *string `json:"PreferredBackupTime,omitempty" xml:"PreferredBackupTime,omitempty"`
}

func (s DescribeBackupLocalResponseBodyBackupPolicyDO) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupLocalResponseBodyBackupPolicyDO) GoString() string {
	return s.String()
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) SetBackupAppName(v string) *DescribeBackupLocalResponseBodyBackupPolicyDO {
	s.BackupAppName = &v
	return s
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) SetBackupDbName(v string) *DescribeBackupLocalResponseBodyBackupPolicyDO {
	s.BackupDbName = &v
	return s
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) SetBackupLevel(v string) *DescribeBackupLocalResponseBodyBackupPolicyDO {
	s.BackupLevel = &v
	return s
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) SetBackupLog(v string) *DescribeBackupLocalResponseBodyBackupPolicyDO {
	s.BackupLog = &v
	return s
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) SetBackupMode(v string) *DescribeBackupLocalResponseBodyBackupPolicyDO {
	s.BackupMode = &v
	return s
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) SetBackupPolicyMode(v string) *DescribeBackupLocalResponseBodyBackupPolicyDO {
	s.BackupPolicyMode = &v
	return s
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) SetBackupRetentionPeriod(v int64) *DescribeBackupLocalResponseBodyBackupPolicyDO {
	s.BackupRetentionPeriod = &v
	return s
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) SetBackupType(v string) *DescribeBackupLocalResponseBodyBackupPolicyDO {
	s.BackupType = &v
	return s
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) SetDataBackupRetentionPeriod(v int64) *DescribeBackupLocalResponseBodyBackupPolicyDO {
	s.DataBackupRetentionPeriod = &v
	return s
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) SetGmtCreate(v int64) *DescribeBackupLocalResponseBodyBackupPolicyDO {
	s.GmtCreate = &v
	return s
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) SetGmtModified(v int64) *DescribeBackupLocalResponseBodyBackupPolicyDO {
	s.GmtModified = &v
	return s
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) SetHighSpaceUsageProtection(v int64) *DescribeBackupLocalResponseBodyBackupPolicyDO {
	s.HighSpaceUsageProtection = &v
	return s
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) SetLocalLogRetentionHours(v int64) *DescribeBackupLocalResponseBodyBackupPolicyDO {
	s.LocalLogRetentionHours = &v
	return s
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) SetLocalLogRetentionSpace(v int64) *DescribeBackupLocalResponseBodyBackupPolicyDO {
	s.LocalLogRetentionSpace = &v
	return s
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) SetLogBackupRetentionPeriod(v int64) *DescribeBackupLocalResponseBodyBackupPolicyDO {
	s.LogBackupRetentionPeriod = &v
	return s
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) SetNextBackupActuallyTime(v string) *DescribeBackupLocalResponseBodyBackupPolicyDO {
	s.NextBackupActuallyTime = &v
	return s
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) SetPreferredBackupPeriod(v string) *DescribeBackupLocalResponseBodyBackupPolicyDO {
	s.PreferredBackupPeriod = &v
	return s
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) SetPreferredBackupTime(v string) *DescribeBackupLocalResponseBodyBackupPolicyDO {
	s.PreferredBackupTime = &v
	return s
}

type DescribeBackupLocalResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeBackupLocalResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeBackupLocalResponse) SetStatusCode(v int32) *DescribeBackupLocalResponse {
	s.StatusCode = &v
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
	BackupPolicyDO *DescribeBackupPolicyResponseBodyBackupPolicyDO `json:"BackupPolicyDO,omitempty" xml:"BackupPolicyDO,omitempty" type:"Struct"`
	RequestId      *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeBackupPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyResponseBody) SetBackupPolicyDO(v *DescribeBackupPolicyResponseBodyBackupPolicyDO) *DescribeBackupPolicyResponseBody {
	s.BackupPolicyDO = v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetRequestId(v string) *DescribeBackupPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetSuccess(v bool) *DescribeBackupPolicyResponseBody {
	s.Success = &v
	return s
}

type DescribeBackupPolicyResponseBodyBackupPolicyDO struct {
	BackupAppName             *string `json:"BackupAppName,omitempty" xml:"BackupAppName,omitempty"`
	BackupDbName              *string `json:"BackupDbName,omitempty" xml:"BackupDbName,omitempty"`
	BackupLevel               *string `json:"BackupLevel,omitempty" xml:"BackupLevel,omitempty"`
	BackupLog                 *string `json:"BackupLog,omitempty" xml:"BackupLog,omitempty"`
	BackupMode                *string `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	BackupPolicyMode          *string `json:"BackupPolicyMode,omitempty" xml:"BackupPolicyMode,omitempty"`
	BackupRetentionPeriod     *int64  `json:"BackupRetentionPeriod,omitempty" xml:"BackupRetentionPeriod,omitempty"`
	BackupType                *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	DataBackupRetentionPeriod *int64  `json:"DataBackupRetentionPeriod,omitempty" xml:"DataBackupRetentionPeriod,omitempty"`
	GmtCreate                 *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified               *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	HighSpaceUsageProtection  *int64  `json:"HighSpaceUsageProtection,omitempty" xml:"HighSpaceUsageProtection,omitempty"`
	LocalLogRetentionHours    *int64  `json:"LocalLogRetentionHours,omitempty" xml:"LocalLogRetentionHours,omitempty"`
	LocalLogRetentionSpace    *int64  `json:"LocalLogRetentionSpace,omitempty" xml:"LocalLogRetentionSpace,omitempty"`
	LogBackupRetentionPeriod  *int64  `json:"LogBackupRetentionPeriod,omitempty" xml:"LogBackupRetentionPeriod,omitempty"`
	NextBackupActuallyTime    *string `json:"NextBackupActuallyTime,omitempty" xml:"NextBackupActuallyTime,omitempty"`
	PreferredBackupPeriod     *string `json:"PreferredBackupPeriod,omitempty" xml:"PreferredBackupPeriod,omitempty"`
	PreferredBackupTime       *string `json:"PreferredBackupTime,omitempty" xml:"PreferredBackupTime,omitempty"`
}

func (s DescribeBackupPolicyResponseBodyBackupPolicyDO) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPolicyResponseBodyBackupPolicyDO) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) SetBackupAppName(v string) *DescribeBackupPolicyResponseBodyBackupPolicyDO {
	s.BackupAppName = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) SetBackupDbName(v string) *DescribeBackupPolicyResponseBodyBackupPolicyDO {
	s.BackupDbName = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) SetBackupLevel(v string) *DescribeBackupPolicyResponseBodyBackupPolicyDO {
	s.BackupLevel = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) SetBackupLog(v string) *DescribeBackupPolicyResponseBodyBackupPolicyDO {
	s.BackupLog = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) SetBackupMode(v string) *DescribeBackupPolicyResponseBodyBackupPolicyDO {
	s.BackupMode = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) SetBackupPolicyMode(v string) *DescribeBackupPolicyResponseBodyBackupPolicyDO {
	s.BackupPolicyMode = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) SetBackupRetentionPeriod(v int64) *DescribeBackupPolicyResponseBodyBackupPolicyDO {
	s.BackupRetentionPeriod = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) SetBackupType(v string) *DescribeBackupPolicyResponseBodyBackupPolicyDO {
	s.BackupType = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) SetDataBackupRetentionPeriod(v int64) *DescribeBackupPolicyResponseBodyBackupPolicyDO {
	s.DataBackupRetentionPeriod = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) SetGmtCreate(v int64) *DescribeBackupPolicyResponseBodyBackupPolicyDO {
	s.GmtCreate = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) SetGmtModified(v int64) *DescribeBackupPolicyResponseBodyBackupPolicyDO {
	s.GmtModified = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) SetHighSpaceUsageProtection(v int64) *DescribeBackupPolicyResponseBodyBackupPolicyDO {
	s.HighSpaceUsageProtection = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) SetLocalLogRetentionHours(v int64) *DescribeBackupPolicyResponseBodyBackupPolicyDO {
	s.LocalLogRetentionHours = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) SetLocalLogRetentionSpace(v int64) *DescribeBackupPolicyResponseBodyBackupPolicyDO {
	s.LocalLogRetentionSpace = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) SetLogBackupRetentionPeriod(v int64) *DescribeBackupPolicyResponseBodyBackupPolicyDO {
	s.LogBackupRetentionPeriod = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) SetNextBackupActuallyTime(v string) *DescribeBackupPolicyResponseBodyBackupPolicyDO {
	s.NextBackupActuallyTime = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) SetPreferredBackupPeriod(v string) *DescribeBackupPolicyResponseBodyBackupPolicyDO {
	s.PreferredBackupPeriod = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) SetPreferredBackupTime(v string) *DescribeBackupPolicyResponseBodyBackupPolicyDO {
	s.PreferredBackupTime = &v
	return s
}

type DescribeBackupPolicyResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeBackupPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeBackupPolicyResponse) SetStatusCode(v int32) *DescribeBackupPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBackupPolicyResponse) SetBody(v *DescribeBackupPolicyResponseBody) *DescribeBackupPolicyResponse {
	s.Body = v
	return s
}

type DescribeBackupSetsRequest struct {
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
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

func (s *DescribeBackupSetsRequest) SetEndTime(v string) *DescribeBackupSetsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeBackupSetsRequest) SetStartTime(v string) *DescribeBackupSetsRequest {
	s.StartTime = &v
	return s
}

type DescribeBackupSetsResponseBody struct {
	BackupSets *DescribeBackupSetsResponseBodyBackupSets `json:"BackupSets,omitempty" xml:"BackupSets,omitempty" type:"Struct"`
	RequestId  *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeBackupSetsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupSetsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupSetsResponseBody) SetBackupSets(v *DescribeBackupSetsResponseBodyBackupSets) *DescribeBackupSetsResponseBody {
	s.BackupSets = v
	return s
}

func (s *DescribeBackupSetsResponseBody) SetRequestId(v string) *DescribeBackupSetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupSetsResponseBody) SetSuccess(v bool) *DescribeBackupSetsResponseBody {
	s.Success = &v
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
	BackupConsitentTime *string                                                     `json:"BackupConsitentTime,omitempty" xml:"BackupConsitentTime,omitempty"`
	BackupDbs           *DescribeBackupSetsResponseBodyBackupSetsBackupSetBackupDbs `json:"BackupDbs,omitempty" xml:"BackupDbs,omitempty" type:"Struct"`
	BackupEndTime       *int64                                                      `json:"BackupEndTime,omitempty" xml:"BackupEndTime,omitempty"`
	BackupLevel         *string                                                     `json:"BackupLevel,omitempty" xml:"BackupLevel,omitempty"`
	BackupMode          *string                                                     `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	BackupStartTime     *int64                                                      `json:"BackupStartTime,omitempty" xml:"BackupStartTime,omitempty"`
	BackupTotalSize     *string                                                     `json:"BackupTotalSize,omitempty" xml:"BackupTotalSize,omitempty"`
	BackupType          *string                                                     `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	EnableRecovery      *bool                                                       `json:"EnableRecovery,omitempty" xml:"EnableRecovery,omitempty"`
	Id                  *string                                                     `json:"Id,omitempty" xml:"Id,omitempty"`
	Status              *int64                                                      `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeBackupSetsResponseBodyBackupSetsBackupSet) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupSetsResponseBodyBackupSetsBackupSet) GoString() string {
	return s.String()
}

func (s *DescribeBackupSetsResponseBodyBackupSetsBackupSet) SetBackupConsitentTime(v string) *DescribeBackupSetsResponseBodyBackupSetsBackupSet {
	s.BackupConsitentTime = &v
	return s
}

func (s *DescribeBackupSetsResponseBodyBackupSetsBackupSet) SetBackupDbs(v *DescribeBackupSetsResponseBodyBackupSetsBackupSetBackupDbs) *DescribeBackupSetsResponseBodyBackupSetsBackupSet {
	s.BackupDbs = v
	return s
}

func (s *DescribeBackupSetsResponseBodyBackupSetsBackupSet) SetBackupEndTime(v int64) *DescribeBackupSetsResponseBodyBackupSetsBackupSet {
	s.BackupEndTime = &v
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

func (s *DescribeBackupSetsResponseBodyBackupSetsBackupSet) SetBackupStartTime(v int64) *DescribeBackupSetsResponseBodyBackupSetsBackupSet {
	s.BackupStartTime = &v
	return s
}

func (s *DescribeBackupSetsResponseBodyBackupSetsBackupSet) SetBackupTotalSize(v string) *DescribeBackupSetsResponseBodyBackupSetsBackupSet {
	s.BackupTotalSize = &v
	return s
}

func (s *DescribeBackupSetsResponseBodyBackupSetsBackupSet) SetBackupType(v string) *DescribeBackupSetsResponseBodyBackupSetsBackupSet {
	s.BackupType = &v
	return s
}

func (s *DescribeBackupSetsResponseBodyBackupSetsBackupSet) SetEnableRecovery(v bool) *DescribeBackupSetsResponseBodyBackupSetsBackupSet {
	s.EnableRecovery = &v
	return s
}

func (s *DescribeBackupSetsResponseBodyBackupSetsBackupSet) SetId(v string) *DescribeBackupSetsResponseBodyBackupSetsBackupSet {
	s.Id = &v
	return s
}

func (s *DescribeBackupSetsResponseBodyBackupSetsBackupSet) SetStatus(v int64) *DescribeBackupSetsResponseBodyBackupSetsBackupSet {
	s.Status = &v
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeBackupSetsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeBackupSetsResponse) SetStatusCode(v int32) *DescribeBackupSetsResponse {
	s.StatusCode = &v
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
	RequestId   *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RestoreTime *DescribeBackupTimesResponseBodyRestoreTime `json:"RestoreTime,omitempty" xml:"RestoreTime,omitempty" type:"Struct"`
	Success     *bool                                       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeBackupTimesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupTimesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupTimesResponseBody) SetRequestId(v string) *DescribeBackupTimesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupTimesResponseBody) SetRestoreTime(v *DescribeBackupTimesResponseBodyRestoreTime) *DescribeBackupTimesResponseBody {
	s.RestoreTime = v
	return s
}

func (s *DescribeBackupTimesResponseBody) SetSuccess(v bool) *DescribeBackupTimesResponseBody {
	s.Success = &v
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeBackupTimesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeBackupTimesResponse) SetStatusCode(v int32) *DescribeBackupTimesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBackupTimesResponse) SetBody(v *DescribeBackupTimesResponseBody) *DescribeBackupTimesResponse {
	s.Body = v
	return s
}

type DescribeBroadcastTablesRequest struct {
	CurrentPage    *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Query          *string `json:"Query,omitempty" xml:"Query,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeBroadcastTablesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBroadcastTablesRequest) GoString() string {
	return s.String()
}

func (s *DescribeBroadcastTablesRequest) SetCurrentPage(v int32) *DescribeBroadcastTablesRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeBroadcastTablesRequest) SetDbName(v string) *DescribeBroadcastTablesRequest {
	s.DbName = &v
	return s
}

func (s *DescribeBroadcastTablesRequest) SetDrdsInstanceId(v string) *DescribeBroadcastTablesRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeBroadcastTablesRequest) SetPageSize(v int32) *DescribeBroadcastTablesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeBroadcastTablesRequest) SetQuery(v string) *DescribeBroadcastTablesRequest {
	s.Query = &v
	return s
}

func (s *DescribeBroadcastTablesRequest) SetRegionId(v string) *DescribeBroadcastTablesRequest {
	s.RegionId = &v
	return s
}

type DescribeBroadcastTablesResponseBody struct {
	IsShard    *bool                                      `json:"IsShard,omitempty" xml:"IsShard,omitempty"`
	List       []*DescribeBroadcastTablesResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	PageNumber *int32                                     `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
	Total      *int32                                     `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeBroadcastTablesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBroadcastTablesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBroadcastTablesResponseBody) SetIsShard(v bool) *DescribeBroadcastTablesResponseBody {
	s.IsShard = &v
	return s
}

func (s *DescribeBroadcastTablesResponseBody) SetList(v []*DescribeBroadcastTablesResponseBodyList) *DescribeBroadcastTablesResponseBody {
	s.List = v
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

func (s *DescribeBroadcastTablesResponseBody) SetRequestId(v string) *DescribeBroadcastTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBroadcastTablesResponseBody) SetSuccess(v bool) *DescribeBroadcastTablesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeBroadcastTablesResponseBody) SetTotal(v int32) *DescribeBroadcastTablesResponseBody {
	s.Total = &v
	return s
}

type DescribeBroadcastTablesResponseBodyList struct {
	Broadcast     *bool   `json:"Broadcast,omitempty" xml:"Broadcast,omitempty"`
	BroadcastType *string `json:"BroadcastType,omitempty" xml:"BroadcastType,omitempty"`
	DbInstType    *int32  `json:"DbInstType,omitempty" xml:"DbInstType,omitempty"`
	IsShard       *bool   `json:"IsShard,omitempty" xml:"IsShard,omitempty"`
	Status        *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	Table         *string `json:"Table,omitempty" xml:"Table,omitempty"`
}

func (s DescribeBroadcastTablesResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s DescribeBroadcastTablesResponseBodyList) GoString() string {
	return s.String()
}

func (s *DescribeBroadcastTablesResponseBodyList) SetBroadcast(v bool) *DescribeBroadcastTablesResponseBodyList {
	s.Broadcast = &v
	return s
}

func (s *DescribeBroadcastTablesResponseBodyList) SetBroadcastType(v string) *DescribeBroadcastTablesResponseBodyList {
	s.BroadcastType = &v
	return s
}

func (s *DescribeBroadcastTablesResponseBodyList) SetDbInstType(v int32) *DescribeBroadcastTablesResponseBodyList {
	s.DbInstType = &v
	return s
}

func (s *DescribeBroadcastTablesResponseBodyList) SetIsShard(v bool) *DescribeBroadcastTablesResponseBodyList {
	s.IsShard = &v
	return s
}

func (s *DescribeBroadcastTablesResponseBodyList) SetStatus(v int32) *DescribeBroadcastTablesResponseBodyList {
	s.Status = &v
	return s
}

func (s *DescribeBroadcastTablesResponseBodyList) SetTable(v string) *DescribeBroadcastTablesResponseBodyList {
	s.Table = &v
	return s
}

type DescribeBroadcastTablesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeBroadcastTablesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeBroadcastTablesResponse) SetStatusCode(v int32) *DescribeBroadcastTablesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBroadcastTablesResponse) SetBody(v *DescribeBroadcastTablesResponseBody) *DescribeBroadcastTablesResponse {
	s.Body = v
	return s
}

type DescribeDbInstanceDbsRequest struct {
	AccountName    *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	DbInstType     *string `json:"DbInstType,omitempty" xml:"DbInstType,omitempty"`
	DbInstanceId   *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	Password       *string `json:"Password,omitempty" xml:"Password,omitempty"`
}

func (s DescribeDbInstanceDbsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDbInstanceDbsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDbInstanceDbsRequest) SetAccountName(v string) *DescribeDbInstanceDbsRequest {
	s.AccountName = &v
	return s
}

func (s *DescribeDbInstanceDbsRequest) SetDbInstType(v string) *DescribeDbInstanceDbsRequest {
	s.DbInstType = &v
	return s
}

func (s *DescribeDbInstanceDbsRequest) SetDbInstanceId(v string) *DescribeDbInstanceDbsRequest {
	s.DbInstanceId = &v
	return s
}

func (s *DescribeDbInstanceDbsRequest) SetDrdsInstanceId(v string) *DescribeDbInstanceDbsRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeDbInstanceDbsRequest) SetPassword(v string) *DescribeDbInstanceDbsRequest {
	s.Password = &v
	return s
}

type DescribeDbInstanceDbsResponseBody struct {
	Databases *DescribeDbInstanceDbsResponseBodyDatabases `json:"Databases,omitempty" xml:"Databases,omitempty" type:"Struct"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                       `json:"Success,omitempty" xml:"Success,omitempty"`
	Total     *string                                     `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeDbInstanceDbsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDbInstanceDbsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDbInstanceDbsResponseBody) SetDatabases(v *DescribeDbInstanceDbsResponseBodyDatabases) *DescribeDbInstanceDbsResponseBody {
	s.Databases = v
	return s
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
	DbName      *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Status      *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDbInstanceDbsResponseBodyDatabasesDatabase) String() string {
	return tea.Prettify(s)
}

func (s DescribeDbInstanceDbsResponseBodyDatabasesDatabase) GoString() string {
	return s.String()
}

func (s *DescribeDbInstanceDbsResponseBodyDatabasesDatabase) SetDbName(v string) *DescribeDbInstanceDbsResponseBodyDatabasesDatabase {
	s.DbName = &v
	return s
}

func (s *DescribeDbInstanceDbsResponseBodyDatabasesDatabase) SetDescription(v string) *DescribeDbInstanceDbsResponseBodyDatabasesDatabase {
	s.Description = &v
	return s
}

func (s *DescribeDbInstanceDbsResponseBodyDatabasesDatabase) SetStatus(v int32) *DescribeDbInstanceDbsResponseBodyDatabasesDatabase {
	s.Status = &v
	return s
}

type DescribeDbInstanceDbsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDbInstanceDbsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeDbInstanceDbsResponse) SetStatusCode(v int32) *DescribeDbInstanceDbsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDbInstanceDbsResponse) SetBody(v *DescribeDbInstanceDbsResponseBody) *DescribeDbInstanceDbsResponse {
	s.Body = v
	return s
}

type DescribeDbInstancesRequest struct {
	DbInstType     *string `json:"DbInstType,omitempty" xml:"DbInstType,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	PageNumber     *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Search         *string `json:"Search,omitempty" xml:"Search,omitempty"`
}

func (s DescribeDbInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDbInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDbInstancesRequest) SetDbInstType(v string) *DescribeDbInstancesRequest {
	s.DbInstType = &v
	return s
}

func (s *DescribeDbInstancesRequest) SetDrdsInstanceId(v string) *DescribeDbInstancesRequest {
	s.DrdsInstanceId = &v
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

func (s *DescribeDbInstancesRequest) SetSearch(v string) *DescribeDbInstancesRequest {
	s.Search = &v
	return s
}

type DescribeDbInstancesResponseBody struct {
	Items     *DescribeDbInstancesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDbInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDbInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDbInstancesResponseBody) SetItems(v *DescribeDbInstancesResponseBodyItems) *DescribeDbInstancesResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDbInstancesResponseBody) SetRequestId(v string) *DescribeDbInstancesResponseBody {
	s.RequestId = &v
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
	DBInstanceDescription *string                                                             `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	DBInstanceId          *string                                                             `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DBInstanceStatus      *int32                                                              `json:"DBInstanceStatus,omitempty" xml:"DBInstanceStatus,omitempty"`
	DBInstanceType        *string                                                             `json:"DBInstanceType,omitempty" xml:"DBInstanceType,omitempty"`
	Engine                *string                                                             `json:"Engine,omitempty" xml:"Engine,omitempty"`
	EngineVersion         *string                                                             `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	InstanceNetworkType   *string                                                             `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	ReadOnlyDBInstanceId  *DescribeDbInstancesResponseBodyItemsDBInstanceReadOnlyDBInstanceId `json:"ReadOnlyDBInstanceId,omitempty" xml:"ReadOnlyDBInstanceId,omitempty" type:"Struct"`
	RegionId              *string                                                             `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ZoneId                *string                                                             `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDbInstancesResponseBodyItemsDBInstance) String() string {
	return tea.Prettify(s)
}

func (s DescribeDbInstancesResponseBodyItemsDBInstance) GoString() string {
	return s.String()
}

func (s *DescribeDbInstancesResponseBodyItemsDBInstance) SetDBInstanceDescription(v string) *DescribeDbInstancesResponseBodyItemsDBInstance {
	s.DBInstanceDescription = &v
	return s
}

func (s *DescribeDbInstancesResponseBodyItemsDBInstance) SetDBInstanceId(v string) *DescribeDbInstancesResponseBodyItemsDBInstance {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDbInstancesResponseBodyItemsDBInstance) SetDBInstanceStatus(v int32) *DescribeDbInstancesResponseBodyItemsDBInstance {
	s.DBInstanceStatus = &v
	return s
}

func (s *DescribeDbInstancesResponseBodyItemsDBInstance) SetDBInstanceType(v string) *DescribeDbInstancesResponseBodyItemsDBInstance {
	s.DBInstanceType = &v
	return s
}

func (s *DescribeDbInstancesResponseBodyItemsDBInstance) SetEngine(v string) *DescribeDbInstancesResponseBodyItemsDBInstance {
	s.Engine = &v
	return s
}

func (s *DescribeDbInstancesResponseBodyItemsDBInstance) SetEngineVersion(v string) *DescribeDbInstancesResponseBodyItemsDBInstance {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDbInstancesResponseBodyItemsDBInstance) SetInstanceNetworkType(v string) *DescribeDbInstancesResponseBodyItemsDBInstance {
	s.InstanceNetworkType = &v
	return s
}

func (s *DescribeDbInstancesResponseBodyItemsDBInstance) SetReadOnlyDBInstanceId(v *DescribeDbInstancesResponseBodyItemsDBInstanceReadOnlyDBInstanceId) *DescribeDbInstancesResponseBodyItemsDBInstance {
	s.ReadOnlyDBInstanceId = v
	return s
}

func (s *DescribeDbInstancesResponseBodyItemsDBInstance) SetRegionId(v string) *DescribeDbInstancesResponseBodyItemsDBInstance {
	s.RegionId = &v
	return s
}

func (s *DescribeDbInstancesResponseBodyItemsDBInstance) SetZoneId(v string) *DescribeDbInstancesResponseBodyItemsDBInstance {
	s.ZoneId = &v
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDbInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeDbInstancesResponse) SetStatusCode(v int32) *DescribeDbInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDbInstancesResponse) SetBody(v *DescribeDbInstancesResponseBody) *DescribeDbInstancesResponse {
	s.Body = v
	return s
}

type DescribeDrdsDBRequest struct {
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s DescribeDrdsDBRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDBRequest) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDBRequest) SetDbName(v string) *DescribeDrdsDBRequest {
	s.DbName = &v
	return s
}

func (s *DescribeDrdsDBRequest) SetDrdsInstanceId(v string) *DescribeDrdsDBRequest {
	s.DrdsInstanceId = &v
	return s
}

type DescribeDrdsDBResponseBody struct {
	Data      *DescribeDrdsDBResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDrdsDBResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDBResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDBResponseBody) SetData(v *DescribeDrdsDBResponseBodyData) *DescribeDrdsDBResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDrdsDBResponseBody) SetRequestId(v string) *DescribeDrdsDBResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrdsDBResponseBody) SetSuccess(v bool) *DescribeDrdsDBResponseBody {
	s.Success = &v
	return s
}

type DescribeDrdsDBResponseBodyData struct {
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DbInstType *string `json:"DbInstType,omitempty" xml:"DbInstType,omitempty"`
	DbName     *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	InstRole   *string `json:"InstRole,omitempty" xml:"InstRole,omitempty"`
	Mode       *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	Schema     *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDrdsDBResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDBResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDBResponseBodyData) SetCreateTime(v string) *DescribeDrdsDBResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *DescribeDrdsDBResponseBodyData) SetDbInstType(v string) *DescribeDrdsDBResponseBodyData {
	s.DbInstType = &v
	return s
}

func (s *DescribeDrdsDBResponseBodyData) SetDbName(v string) *DescribeDrdsDBResponseBodyData {
	s.DbName = &v
	return s
}

func (s *DescribeDrdsDBResponseBodyData) SetInstRole(v string) *DescribeDrdsDBResponseBodyData {
	s.InstRole = &v
	return s
}

func (s *DescribeDrdsDBResponseBodyData) SetMode(v string) *DescribeDrdsDBResponseBodyData {
	s.Mode = &v
	return s
}

func (s *DescribeDrdsDBResponseBodyData) SetSchema(v string) *DescribeDrdsDBResponseBodyData {
	s.Schema = &v
	return s
}

func (s *DescribeDrdsDBResponseBodyData) SetStatus(v string) *DescribeDrdsDBResponseBodyData {
	s.Status = &v
	return s
}

type DescribeDrdsDBResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDrdsDBResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeDrdsDBResponse) SetStatusCode(v int32) *DescribeDrdsDBResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDrdsDBResponse) SetBody(v *DescribeDrdsDBResponseBody) *DescribeDrdsDBResponse {
	s.Body = v
	return s
}

type DescribeDrdsDBClusterRequest struct {
	DbInstanceId   *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s DescribeDrdsDBClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDBClusterRequest) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDBClusterRequest) SetDbInstanceId(v string) *DescribeDrdsDBClusterRequest {
	s.DbInstanceId = &v
	return s
}

func (s *DescribeDrdsDBClusterRequest) SetDbName(v string) *DescribeDrdsDBClusterRequest {
	s.DbName = &v
	return s
}

func (s *DescribeDrdsDBClusterRequest) SetDrdsInstanceId(v string) *DescribeDrdsDBClusterRequest {
	s.DrdsInstanceId = &v
	return s
}

type DescribeDrdsDBClusterResponseBody struct {
	DbInstance *DescribeDrdsDBClusterResponseBodyDbInstance `json:"DbInstance,omitempty" xml:"DbInstance,omitempty" type:"Struct"`
	RequestId  *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDrdsDBClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDBClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDBClusterResponseBody) SetDbInstance(v *DescribeDrdsDBClusterResponseBodyDbInstance) *DescribeDrdsDBClusterResponseBody {
	s.DbInstance = v
	return s
}

func (s *DescribeDrdsDBClusterResponseBody) SetRequestId(v string) *DescribeDrdsDBClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrdsDBClusterResponseBody) SetSuccess(v bool) *DescribeDrdsDBClusterResponseBody {
	s.Success = &v
	return s
}

type DescribeDrdsDBClusterResponseBodyDbInstance struct {
	DBInstanceId     *string                                               `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DBInstanceStatus *string                                               `json:"DBInstanceStatus,omitempty" xml:"DBInstanceStatus,omitempty"`
	DBNodes          *DescribeDrdsDBClusterResponseBodyDbInstanceDBNodes   `json:"DBNodes,omitempty" xml:"DBNodes,omitempty" type:"Struct"`
	DbInstType       *string                                               `json:"DbInstType,omitempty" xml:"DbInstType,omitempty"`
	Endpoints        *DescribeDrdsDBClusterResponseBodyDbInstanceEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Struct"`
	Engine           *string                                               `json:"Engine,omitempty" xml:"Engine,omitempty"`
	EngineVersion    *string                                               `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	ExpireTime       *string                                               `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	NetworkType      *string                                               `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	PayType          *string                                               `json:"PayType,omitempty" xml:"PayType,omitempty"`
	Port             *int32                                                `json:"Port,omitempty" xml:"Port,omitempty"`
	RdsInstType      *string                                               `json:"RdsInstType,omitempty" xml:"RdsInstType,omitempty"`
	ReadMode         *string                                               `json:"ReadMode,omitempty" xml:"ReadMode,omitempty"`
	RemainDays       *string                                               `json:"RemainDays,omitempty" xml:"RemainDays,omitempty"`
}

func (s DescribeDrdsDBClusterResponseBodyDbInstance) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDBClusterResponseBodyDbInstance) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstance) SetDBInstanceId(v string) *DescribeDrdsDBClusterResponseBodyDbInstance {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstance) SetDBInstanceStatus(v string) *DescribeDrdsDBClusterResponseBodyDbInstance {
	s.DBInstanceStatus = &v
	return s
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstance) SetDBNodes(v *DescribeDrdsDBClusterResponseBodyDbInstanceDBNodes) *DescribeDrdsDBClusterResponseBodyDbInstance {
	s.DBNodes = v
	return s
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstance) SetDbInstType(v string) *DescribeDrdsDBClusterResponseBodyDbInstance {
	s.DbInstType = &v
	return s
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstance) SetEndpoints(v *DescribeDrdsDBClusterResponseBodyDbInstanceEndpoints) *DescribeDrdsDBClusterResponseBodyDbInstance {
	s.Endpoints = v
	return s
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstance) SetEngine(v string) *DescribeDrdsDBClusterResponseBodyDbInstance {
	s.Engine = &v
	return s
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstance) SetEngineVersion(v string) *DescribeDrdsDBClusterResponseBodyDbInstance {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstance) SetExpireTime(v string) *DescribeDrdsDBClusterResponseBodyDbInstance {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstance) SetNetworkType(v string) *DescribeDrdsDBClusterResponseBodyDbInstance {
	s.NetworkType = &v
	return s
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstance) SetPayType(v string) *DescribeDrdsDBClusterResponseBodyDbInstance {
	s.PayType = &v
	return s
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstance) SetPort(v int32) *DescribeDrdsDBClusterResponseBodyDbInstance {
	s.Port = &v
	return s
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstance) SetRdsInstType(v string) *DescribeDrdsDBClusterResponseBodyDbInstance {
	s.RdsInstType = &v
	return s
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstance) SetReadMode(v string) *DescribeDrdsDBClusterResponseBodyDbInstance {
	s.ReadMode = &v
	return s
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstance) SetRemainDays(v string) *DescribeDrdsDBClusterResponseBodyDbInstance {
	s.RemainDays = &v
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
	DBNodeId     *string `json:"DBNodeId,omitempty" xml:"DBNodeId,omitempty"`
	DBNodeRole   *string `json:"DBNodeRole,omitempty" xml:"DBNodeRole,omitempty"`
	DBNodeStatus *string `json:"DBNodeStatus,omitempty" xml:"DBNodeStatus,omitempty"`
	ZoneId       *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDrdsDBClusterResponseBodyDbInstanceDBNodesDBNode) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDBClusterResponseBodyDbInstanceDBNodesDBNode) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstanceDBNodesDBNode) SetDBNodeId(v string) *DescribeDrdsDBClusterResponseBodyDbInstanceDBNodesDBNode {
	s.DBNodeId = &v
	return s
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstanceDBNodesDBNode) SetDBNodeRole(v string) *DescribeDrdsDBClusterResponseBodyDbInstanceDBNodesDBNode {
	s.DBNodeRole = &v
	return s
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstanceDBNodesDBNode) SetDBNodeStatus(v string) *DescribeDrdsDBClusterResponseBodyDbInstanceDBNodesDBNode {
	s.DBNodeStatus = &v
	return s
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstanceDBNodesDBNode) SetZoneId(v string) *DescribeDrdsDBClusterResponseBodyDbInstanceDBNodesDBNode {
	s.ZoneId = &v
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
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	NodeIds    *string `json:"NodeIds,omitempty" xml:"NodeIds,omitempty"`
	ReadWeight *int32  `json:"ReadWeight,omitempty" xml:"ReadWeight,omitempty"`
}

func (s DescribeDrdsDBClusterResponseBodyDbInstanceEndpointsEndpoint) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDBClusterResponseBodyDbInstanceEndpointsEndpoint) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstanceEndpointsEndpoint) SetEndpointId(v string) *DescribeDrdsDBClusterResponseBodyDbInstanceEndpointsEndpoint {
	s.EndpointId = &v
	return s
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstanceEndpointsEndpoint) SetNodeIds(v string) *DescribeDrdsDBClusterResponseBodyDbInstanceEndpointsEndpoint {
	s.NodeIds = &v
	return s
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstanceEndpointsEndpoint) SetReadWeight(v int32) *DescribeDrdsDBClusterResponseBodyDbInstanceEndpointsEndpoint {
	s.ReadWeight = &v
	return s
}

type DescribeDrdsDBClusterResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDrdsDBClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeDrdsDBClusterResponse) SetStatusCode(v int32) *DescribeDrdsDBClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDrdsDBClusterResponse) SetBody(v *DescribeDrdsDBClusterResponseBody) *DescribeDrdsDBClusterResponse {
	s.Body = v
	return s
}

type DescribeDrdsDBIpWhiteListRequest struct {
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	GroupName      *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDrdsDBIpWhiteListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDBIpWhiteListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDBIpWhiteListRequest) SetDbName(v string) *DescribeDrdsDBIpWhiteListRequest {
	s.DbName = &v
	return s
}

func (s *DescribeDrdsDBIpWhiteListRequest) SetDrdsInstanceId(v string) *DescribeDrdsDBIpWhiteListRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeDrdsDBIpWhiteListRequest) SetGroupName(v string) *DescribeDrdsDBIpWhiteListRequest {
	s.GroupName = &v
	return s
}

func (s *DescribeDrdsDBIpWhiteListRequest) SetRegionId(v string) *DescribeDrdsDBIpWhiteListRequest {
	s.RegionId = &v
	return s
}

type DescribeDrdsDBIpWhiteListResponseBody struct {
	IpWhiteList *DescribeDrdsDBIpWhiteListResponseBodyIpWhiteList `json:"IpWhiteList,omitempty" xml:"IpWhiteList,omitempty" type:"Struct"`
	RequestId   *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success     *bool                                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDrdsDBIpWhiteListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDBIpWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDBIpWhiteListResponseBody) SetIpWhiteList(v *DescribeDrdsDBIpWhiteListResponseBodyIpWhiteList) *DescribeDrdsDBIpWhiteListResponseBody {
	s.IpWhiteList = v
	return s
}

func (s *DescribeDrdsDBIpWhiteListResponseBody) SetRequestId(v string) *DescribeDrdsDBIpWhiteListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrdsDBIpWhiteListResponseBody) SetSuccess(v bool) *DescribeDrdsDBIpWhiteListResponseBody {
	s.Success = &v
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
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDrdsDBIpWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeDrdsDBIpWhiteListResponse) SetStatusCode(v int32) *DescribeDrdsDBIpWhiteListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDrdsDBIpWhiteListResponse) SetBody(v *DescribeDrdsDBIpWhiteListResponseBody) *DescribeDrdsDBIpWhiteListResponse {
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
	Data       *DescribeDrdsDBsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	PageNumber *string                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *string                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
	Total      *string                          `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeDrdsDBsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDBsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDBsResponseBody) SetData(v *DescribeDrdsDBsResponseBodyData) *DescribeDrdsDBsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDrdsDBsResponseBody) SetPageNumber(v string) *DescribeDrdsDBsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDrdsDBsResponseBody) SetPageSize(v string) *DescribeDrdsDBsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDrdsDBsResponseBody) SetRequestId(v string) *DescribeDrdsDBsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrdsDBsResponseBody) SetSuccess(v bool) *DescribeDrdsDBsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDrdsDBsResponseBody) SetTotal(v string) *DescribeDrdsDBsResponseBody {
	s.Total = &v
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
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DbInstType *string `json:"DbInstType,omitempty" xml:"DbInstType,omitempty"`
	DbName     *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	Mode       *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	Schema     *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDrdsDBsResponseBodyDataDb) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDBsResponseBodyDataDb) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDBsResponseBodyDataDb) SetCreateTime(v string) *DescribeDrdsDBsResponseBodyDataDb {
	s.CreateTime = &v
	return s
}

func (s *DescribeDrdsDBsResponseBodyDataDb) SetDbInstType(v string) *DescribeDrdsDBsResponseBodyDataDb {
	s.DbInstType = &v
	return s
}

func (s *DescribeDrdsDBsResponseBodyDataDb) SetDbName(v string) *DescribeDrdsDBsResponseBodyDataDb {
	s.DbName = &v
	return s
}

func (s *DescribeDrdsDBsResponseBodyDataDb) SetMode(v string) *DescribeDrdsDBsResponseBodyDataDb {
	s.Mode = &v
	return s
}

func (s *DescribeDrdsDBsResponseBodyDataDb) SetSchema(v string) *DescribeDrdsDBsResponseBodyDataDb {
	s.Schema = &v
	return s
}

func (s *DescribeDrdsDBsResponseBodyDataDb) SetStatus(v string) *DescribeDrdsDBsResponseBodyDataDb {
	s.Status = &v
	return s
}

type DescribeDrdsDBsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDrdsDBsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeDrdsDBsResponse) SetStatusCode(v int32) *DescribeDrdsDBsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDrdsDBsResponse) SetBody(v *DescribeDrdsDBsResponseBody) *DescribeDrdsDBsResponse {
	s.Body = v
	return s
}

type DescribeDrdsDbInstanceRequest struct {
	DbInstanceId   *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s DescribeDrdsDbInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDbInstanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDbInstanceRequest) SetDbInstanceId(v string) *DescribeDrdsDbInstanceRequest {
	s.DbInstanceId = &v
	return s
}

func (s *DescribeDrdsDbInstanceRequest) SetDbName(v string) *DescribeDrdsDbInstanceRequest {
	s.DbName = &v
	return s
}

func (s *DescribeDrdsDbInstanceRequest) SetDrdsInstanceId(v string) *DescribeDrdsDbInstanceRequest {
	s.DrdsInstanceId = &v
	return s
}

type DescribeDrdsDbInstanceResponseBody struct {
	DbInstance *DescribeDrdsDbInstanceResponseBodyDbInstance `json:"DbInstance,omitempty" xml:"DbInstance,omitempty" type:"Struct"`
	RequestId  *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDrdsDbInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDbInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDbInstanceResponseBody) SetDbInstance(v *DescribeDrdsDbInstanceResponseBodyDbInstance) *DescribeDrdsDbInstanceResponseBody {
	s.DbInstance = v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBody) SetRequestId(v string) *DescribeDrdsDbInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBody) SetSuccess(v bool) *DescribeDrdsDbInstanceResponseBody {
	s.Success = &v
	return s
}

type DescribeDrdsDbInstanceResponseBodyDbInstance struct {
	ConnectUrl        *string                                                        `json:"ConnectUrl,omitempty" xml:"ConnectUrl,omitempty"`
	DBInstanceId      *string                                                        `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DBInstanceStatus  *string                                                        `json:"DBInstanceStatus,omitempty" xml:"DBInstanceStatus,omitempty"`
	DbInstType        *string                                                        `json:"DbInstType,omitempty" xml:"DbInstType,omitempty"`
	DmInstanceId      *string                                                        `json:"DmInstanceId,omitempty" xml:"DmInstanceId,omitempty"`
	Engine            *string                                                        `json:"Engine,omitempty" xml:"Engine,omitempty"`
	EngineVersion     *string                                                        `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	ExpireTime        *string                                                        `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	NetworkType       *string                                                        `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	PayType           *string                                                        `json:"PayType,omitempty" xml:"PayType,omitempty"`
	Port              *int32                                                         `json:"Port,omitempty" xml:"Port,omitempty"`
	RdsInstType       *string                                                        `json:"RdsInstType,omitempty" xml:"RdsInstType,omitempty"`
	ReadOnlyInstances *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstances `json:"ReadOnlyInstances,omitempty" xml:"ReadOnlyInstances,omitempty" type:"Struct"`
	ReadWeight        *int32                                                         `json:"ReadWeight,omitempty" xml:"ReadWeight,omitempty"`
	RemainDays        *string                                                        `json:"RemainDays,omitempty" xml:"RemainDays,omitempty"`
}

func (s DescribeDrdsDbInstanceResponseBodyDbInstance) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDbInstanceResponseBodyDbInstance) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstance) SetConnectUrl(v string) *DescribeDrdsDbInstanceResponseBodyDbInstance {
	s.ConnectUrl = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstance) SetDBInstanceId(v string) *DescribeDrdsDbInstanceResponseBodyDbInstance {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstance) SetDBInstanceStatus(v string) *DescribeDrdsDbInstanceResponseBodyDbInstance {
	s.DBInstanceStatus = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstance) SetDbInstType(v string) *DescribeDrdsDbInstanceResponseBodyDbInstance {
	s.DbInstType = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstance) SetDmInstanceId(v string) *DescribeDrdsDbInstanceResponseBodyDbInstance {
	s.DmInstanceId = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstance) SetEngine(v string) *DescribeDrdsDbInstanceResponseBodyDbInstance {
	s.Engine = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstance) SetEngineVersion(v string) *DescribeDrdsDbInstanceResponseBodyDbInstance {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstance) SetExpireTime(v string) *DescribeDrdsDbInstanceResponseBodyDbInstance {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstance) SetNetworkType(v string) *DescribeDrdsDbInstanceResponseBodyDbInstance {
	s.NetworkType = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstance) SetPayType(v string) *DescribeDrdsDbInstanceResponseBodyDbInstance {
	s.PayType = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstance) SetPort(v int32) *DescribeDrdsDbInstanceResponseBodyDbInstance {
	s.Port = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstance) SetRdsInstType(v string) *DescribeDrdsDbInstanceResponseBodyDbInstance {
	s.RdsInstType = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstance) SetReadOnlyInstances(v *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstances) *DescribeDrdsDbInstanceResponseBodyDbInstance {
	s.ReadOnlyInstances = v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstance) SetReadWeight(v int32) *DescribeDrdsDbInstanceResponseBodyDbInstance {
	s.ReadWeight = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstance) SetRemainDays(v string) *DescribeDrdsDbInstanceResponseBodyDbInstance {
	s.RemainDays = &v
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
	ConnectUrl       *string `json:"ConnectUrl,omitempty" xml:"ConnectUrl,omitempty"`
	DBInstanceId     *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DBInstanceStatus *string `json:"DBInstanceStatus,omitempty" xml:"DBInstanceStatus,omitempty"`
	DbInstType       *string `json:"DbInstType,omitempty" xml:"DbInstType,omitempty"`
	DmInstanceId     *string `json:"DmInstanceId,omitempty" xml:"DmInstanceId,omitempty"`
	Engine           *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	EngineVersion    *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	ExpireTime       *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	NetworkType      *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	PayType          *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	Port             *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	RdsInstType      *string `json:"RdsInstType,omitempty" xml:"RdsInstType,omitempty"`
	ReadWeight       *int32  `json:"ReadWeight,omitempty" xml:"ReadWeight,omitempty"`
	RemainDays       *string `json:"RemainDays,omitempty" xml:"RemainDays,omitempty"`
	VersionAction    *int32  `json:"VersionAction,omitempty" xml:"VersionAction,omitempty"`
}

func (s DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) SetConnectUrl(v string) *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.ConnectUrl = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) SetDBInstanceId(v string) *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) SetDBInstanceStatus(v string) *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.DBInstanceStatus = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) SetDbInstType(v string) *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.DbInstType = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) SetDmInstanceId(v string) *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.DmInstanceId = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) SetEngine(v string) *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.Engine = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) SetEngineVersion(v string) *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) SetExpireTime(v string) *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) SetNetworkType(v string) *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.NetworkType = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) SetPayType(v string) *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.PayType = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) SetPort(v int32) *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.Port = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) SetRdsInstType(v string) *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.RdsInstType = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) SetReadWeight(v int32) *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.ReadWeight = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) SetRemainDays(v string) *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.RemainDays = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) SetVersionAction(v int32) *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.VersionAction = &v
	return s
}

type DescribeDrdsDbInstanceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDrdsDbInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeDrdsDbInstanceResponse) SetStatusCode(v int32) *DescribeDrdsDbInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponse) SetBody(v *DescribeDrdsDbInstanceResponseBody) *DescribeDrdsDbInstanceResponse {
	s.Body = v
	return s
}

type DescribeDrdsDbInstancesRequest struct {
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	PageNumber     *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeDrdsDbInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDbInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDbInstancesRequest) SetDbName(v string) *DescribeDrdsDbInstancesRequest {
	s.DbName = &v
	return s
}

func (s *DescribeDrdsDbInstancesRequest) SetDrdsInstanceId(v string) *DescribeDrdsDbInstancesRequest {
	s.DrdsInstanceId = &v
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
	DbInstances *DescribeDrdsDbInstancesResponseBodyDbInstances `json:"DbInstances,omitempty" xml:"DbInstances,omitempty" type:"Struct"`
	PageNumber  *string                                         `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *string                                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success     *bool                                           `json:"Success,omitempty" xml:"Success,omitempty"`
	Total       *string                                         `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeDrdsDbInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDbInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDbInstancesResponseBody) SetDbInstances(v *DescribeDrdsDbInstancesResponseBodyDbInstances) *DescribeDrdsDbInstancesResponseBody {
	s.DbInstances = v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBody) SetPageNumber(v string) *DescribeDrdsDbInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBody) SetPageSize(v string) *DescribeDrdsDbInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBody) SetRequestId(v string) *DescribeDrdsDbInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBody) SetSuccess(v bool) *DescribeDrdsDbInstancesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBody) SetTotal(v string) *DescribeDrdsDbInstancesResponseBody {
	s.Total = &v
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
	ConnectUrl        *string                                                                    `json:"ConnectUrl,omitempty" xml:"ConnectUrl,omitempty"`
	DBInstanceId      *string                                                                    `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DBInstanceStatus  *string                                                                    `json:"DBInstanceStatus,omitempty" xml:"DBInstanceStatus,omitempty"`
	DbInstType        *string                                                                    `json:"DbInstType,omitempty" xml:"DbInstType,omitempty"`
	DmInstanceId      *string                                                                    `json:"DmInstanceId,omitempty" xml:"DmInstanceId,omitempty"`
	Engine            *string                                                                    `json:"Engine,omitempty" xml:"Engine,omitempty"`
	EngineVersion     *string                                                                    `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	ExpireTime        *string                                                                    `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	NetworkType       *string                                                                    `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	PayType           *string                                                                    `json:"PayType,omitempty" xml:"PayType,omitempty"`
	Port              *int32                                                                     `json:"Port,omitempty" xml:"Port,omitempty"`
	RdsInstType       *string                                                                    `json:"RdsInstType,omitempty" xml:"RdsInstType,omitempty"`
	ReadOnlyInstances *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstances `json:"ReadOnlyInstances,omitempty" xml:"ReadOnlyInstances,omitempty" type:"Struct"`
	ReadWeight        *int32                                                                     `json:"ReadWeight,omitempty" xml:"ReadWeight,omitempty"`
	RemainDays        *int32                                                                     `json:"RemainDays,omitempty" xml:"RemainDays,omitempty"`
}

func (s DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) SetConnectUrl(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance {
	s.ConnectUrl = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) SetDBInstanceId(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) SetDBInstanceStatus(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance {
	s.DBInstanceStatus = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) SetDbInstType(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance {
	s.DbInstType = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) SetDmInstanceId(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance {
	s.DmInstanceId = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) SetEngine(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance {
	s.Engine = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) SetEngineVersion(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) SetExpireTime(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) SetNetworkType(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance {
	s.NetworkType = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) SetPayType(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance {
	s.PayType = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) SetPort(v int32) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance {
	s.Port = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) SetRdsInstType(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance {
	s.RdsInstType = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) SetReadOnlyInstances(v *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstances) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance {
	s.ReadOnlyInstances = v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) SetReadWeight(v int32) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance {
	s.ReadWeight = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) SetRemainDays(v int32) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance {
	s.RemainDays = &v
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
	ConnectUrl       *string `json:"ConnectUrl,omitempty" xml:"ConnectUrl,omitempty"`
	DBInstanceStatus *string `json:"DBInstanceStatus,omitempty" xml:"DBInstanceStatus,omitempty"`
	DbInstType       *string `json:"DbInstType,omitempty" xml:"DbInstType,omitempty"`
	DmInstanceId     *string `json:"DmInstanceId,omitempty" xml:"DmInstanceId,omitempty"`
	Engine           *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	EngineVersion    *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	ExpireTime       *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	InstanceName     *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	NetworkType      *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	PayType          *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	Port             *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	RdsInstType      *string `json:"RdsInstType,omitempty" xml:"RdsInstType,omitempty"`
	ReadWeight       *int32  `json:"ReadWeight,omitempty" xml:"ReadWeight,omitempty"`
	RemainDays       *int32  `json:"RemainDays,omitempty" xml:"RemainDays,omitempty"`
}

func (s DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) SetConnectUrl(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.ConnectUrl = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) SetDBInstanceStatus(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.DBInstanceStatus = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) SetDbInstType(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.DbInstType = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) SetDmInstanceId(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.DmInstanceId = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) SetEngine(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.Engine = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) SetEngineVersion(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) SetExpireTime(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) SetInstanceName(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.InstanceName = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) SetNetworkType(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.NetworkType = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) SetPayType(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.PayType = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) SetPort(v int32) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.Port = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) SetRdsInstType(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.RdsInstType = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) SetReadWeight(v int32) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.ReadWeight = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) SetRemainDays(v int32) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.RemainDays = &v
	return s
}

type DescribeDrdsDbInstancesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDrdsDbInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeDrdsDbInstancesResponse) SetStatusCode(v int32) *DescribeDrdsDbInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponse) SetBody(v *DescribeDrdsDbInstancesResponseBody) *DescribeDrdsDbInstancesResponse {
	s.Body = v
	return s
}

type DescribeDrdsDbRdsNameListRequest struct {
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s DescribeDrdsDbRdsNameListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDbRdsNameListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDbRdsNameListRequest) SetDbName(v string) *DescribeDrdsDbRdsNameListRequest {
	s.DbName = &v
	return s
}

func (s *DescribeDrdsDbRdsNameListRequest) SetDrdsInstanceId(v string) *DescribeDrdsDbRdsNameListRequest {
	s.DrdsInstanceId = &v
	return s
}

type DescribeDrdsDbRdsNameListResponseBody struct {
	InstanceNameList *DescribeDrdsDbRdsNameListResponseBodyInstanceNameList `json:"InstanceNameList,omitempty" xml:"InstanceNameList,omitempty" type:"Struct"`
	RequestId        *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success          *bool                                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDrdsDbRdsNameListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDbRdsNameListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDbRdsNameListResponseBody) SetInstanceNameList(v *DescribeDrdsDbRdsNameListResponseBodyInstanceNameList) *DescribeDrdsDbRdsNameListResponseBody {
	s.InstanceNameList = v
	return s
}

func (s *DescribeDrdsDbRdsNameListResponseBody) SetRequestId(v string) *DescribeDrdsDbRdsNameListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrdsDbRdsNameListResponseBody) SetSuccess(v bool) *DescribeDrdsDbRdsNameListResponseBody {
	s.Success = &v
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
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDrdsDbRdsNameListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeDrdsDbRdsNameListResponse) SetStatusCode(v int32) *DescribeDrdsDbRdsNameListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDrdsDbRdsNameListResponse) SetBody(v *DescribeDrdsDbRdsNameListResponseBody) *DescribeDrdsDbRdsNameListResponse {
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
	Data      *DescribeDrdsInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDrdsInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceResponseBody) SetData(v *DescribeDrdsInstanceResponseBodyData) *DescribeDrdsInstanceResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDrdsInstanceResponseBody) SetRequestId(v string) *DescribeDrdsInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBody) SetSuccess(v bool) *DescribeDrdsInstanceResponseBody {
	s.Success = &v
	return s
}

type DescribeDrdsInstanceResponseBodyData struct {
	CommodityCode         *string                                                    `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	CreateTime            *int64                                                     `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description           *string                                                    `json:"Description,omitempty" xml:"Description,omitempty"`
	DrdsInstanceId        *string                                                    `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	ExpireDate            *int64                                                     `json:"ExpireDate,omitempty" xml:"ExpireDate,omitempty"`
	InstRole              *string                                                    `json:"InstRole,omitempty" xml:"InstRole,omitempty"`
	InstanceSeries        *string                                                    `json:"InstanceSeries,omitempty" xml:"InstanceSeries,omitempty"`
	InstanceSpec          *string                                                    `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty"`
	Label                 *string                                                    `json:"Label,omitempty" xml:"Label,omitempty"`
	MachineType           *string                                                    `json:"MachineType,omitempty" xml:"MachineType,omitempty"`
	MasterInstanceId      *string                                                    `json:"MasterInstanceId,omitempty" xml:"MasterInstanceId,omitempty"`
	MysqlVersion          *int32                                                     `json:"MysqlVersion,omitempty" xml:"MysqlVersion,omitempty"`
	NetworkType           *string                                                    `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	OrderInstanceId       *string                                                    `json:"OrderInstanceId,omitempty" xml:"OrderInstanceId,omitempty"`
	ProductVersion        *string                                                    `json:"ProductVersion,omitempty" xml:"ProductVersion,omitempty"`
	ReadOnlyDBInstanceIds *DescribeDrdsInstanceResponseBodyDataReadOnlyDBInstanceIds `json:"ReadOnlyDBInstanceIds,omitempty" xml:"ReadOnlyDBInstanceIds,omitempty" type:"Struct"`
	RegionId              *string                                                    `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId       *string                                                    `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Status                *string                                                    `json:"Status,omitempty" xml:"Status,omitempty"`
	StorageType           *string                                                    `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	Type                  *string                                                    `json:"Type,omitempty" xml:"Type,omitempty"`
	Version               *int64                                                     `json:"Version,omitempty" xml:"Version,omitempty"`
	VersionAction         *string                                                    `json:"VersionAction,omitempty" xml:"VersionAction,omitempty"`
	Vips                  *DescribeDrdsInstanceResponseBodyDataVips                  `json:"Vips,omitempty" xml:"Vips,omitempty" type:"Struct"`
	VpcCloudInstanceId    *string                                                    `json:"VpcCloudInstanceId,omitempty" xml:"VpcCloudInstanceId,omitempty"`
	ZoneId                *string                                                    `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDrdsInstanceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceResponseBodyData) SetCommodityCode(v string) *DescribeDrdsInstanceResponseBodyData {
	s.CommodityCode = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetCreateTime(v int64) *DescribeDrdsInstanceResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetDescription(v string) *DescribeDrdsInstanceResponseBodyData {
	s.Description = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetDrdsInstanceId(v string) *DescribeDrdsInstanceResponseBodyData {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetExpireDate(v int64) *DescribeDrdsInstanceResponseBodyData {
	s.ExpireDate = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetInstRole(v string) *DescribeDrdsInstanceResponseBodyData {
	s.InstRole = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetInstanceSeries(v string) *DescribeDrdsInstanceResponseBodyData {
	s.InstanceSeries = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetInstanceSpec(v string) *DescribeDrdsInstanceResponseBodyData {
	s.InstanceSpec = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetLabel(v string) *DescribeDrdsInstanceResponseBodyData {
	s.Label = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetMachineType(v string) *DescribeDrdsInstanceResponseBodyData {
	s.MachineType = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetMasterInstanceId(v string) *DescribeDrdsInstanceResponseBodyData {
	s.MasterInstanceId = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetMysqlVersion(v int32) *DescribeDrdsInstanceResponseBodyData {
	s.MysqlVersion = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetNetworkType(v string) *DescribeDrdsInstanceResponseBodyData {
	s.NetworkType = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetOrderInstanceId(v string) *DescribeDrdsInstanceResponseBodyData {
	s.OrderInstanceId = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetProductVersion(v string) *DescribeDrdsInstanceResponseBodyData {
	s.ProductVersion = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetReadOnlyDBInstanceIds(v *DescribeDrdsInstanceResponseBodyDataReadOnlyDBInstanceIds) *DescribeDrdsInstanceResponseBodyData {
	s.ReadOnlyDBInstanceIds = v
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

func (s *DescribeDrdsInstanceResponseBodyData) SetStatus(v string) *DescribeDrdsInstanceResponseBodyData {
	s.Status = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetStorageType(v string) *DescribeDrdsInstanceResponseBodyData {
	s.StorageType = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetType(v string) *DescribeDrdsInstanceResponseBodyData {
	s.Type = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetVersion(v int64) *DescribeDrdsInstanceResponseBodyData {
	s.Version = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetVersionAction(v string) *DescribeDrdsInstanceResponseBodyData {
	s.VersionAction = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetVips(v *DescribeDrdsInstanceResponseBodyDataVips) *DescribeDrdsInstanceResponseBodyData {
	s.Vips = v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetVpcCloudInstanceId(v string) *DescribeDrdsInstanceResponseBodyData {
	s.VpcCloudInstanceId = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetZoneId(v string) *DescribeDrdsInstanceResponseBodyData {
	s.ZoneId = &v
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
	Dns        *string `json:"Dns,omitempty" xml:"Dns,omitempty"`
	ExpireDays *int64  `json:"ExpireDays,omitempty" xml:"ExpireDays,omitempty"`
	Port       *string `json:"Port,omitempty" xml:"Port,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
	VpcId      *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VswitchId  *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
}

func (s DescribeDrdsInstanceResponseBodyDataVipsVip) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceResponseBodyDataVipsVip) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceResponseBodyDataVipsVip) SetDns(v string) *DescribeDrdsInstanceResponseBodyDataVipsVip {
	s.Dns = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyDataVipsVip) SetExpireDays(v int64) *DescribeDrdsInstanceResponseBodyDataVipsVip {
	s.ExpireDays = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyDataVipsVip) SetPort(v string) *DescribeDrdsInstanceResponseBodyDataVipsVip {
	s.Port = &v
	return s
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

type DescribeDrdsInstanceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDrdsInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeDrdsInstanceResponse) SetStatusCode(v int32) *DescribeDrdsInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDrdsInstanceResponse) SetBody(v *DescribeDrdsInstanceResponseBody) *DescribeDrdsInstanceResponse {
	s.Body = v
	return s
}

type DescribeDrdsInstanceDbMonitorRequest struct {
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	EndTime        *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Key            *string `json:"Key,omitempty" xml:"Key,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StartTime      *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDrdsInstanceDbMonitorRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceDbMonitorRequest) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceDbMonitorRequest) SetDbName(v string) *DescribeDrdsInstanceDbMonitorRequest {
	s.DbName = &v
	return s
}

func (s *DescribeDrdsInstanceDbMonitorRequest) SetDrdsInstanceId(v string) *DescribeDrdsInstanceDbMonitorRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeDrdsInstanceDbMonitorRequest) SetEndTime(v int64) *DescribeDrdsInstanceDbMonitorRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDrdsInstanceDbMonitorRequest) SetKey(v string) *DescribeDrdsInstanceDbMonitorRequest {
	s.Key = &v
	return s
}

func (s *DescribeDrdsInstanceDbMonitorRequest) SetRegionId(v string) *DescribeDrdsInstanceDbMonitorRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDrdsInstanceDbMonitorRequest) SetStartTime(v int64) *DescribeDrdsInstanceDbMonitorRequest {
	s.StartTime = &v
	return s
}

type DescribeDrdsInstanceDbMonitorResponseBody struct {
	Data      []*DescribeDrdsInstanceDbMonitorResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	RequestId *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDrdsInstanceDbMonitorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceDbMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceDbMonitorResponseBody) SetData(v []*DescribeDrdsInstanceDbMonitorResponseBodyData) *DescribeDrdsInstanceDbMonitorResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDrdsInstanceDbMonitorResponseBody) SetRequestId(v string) *DescribeDrdsInstanceDbMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrdsInstanceDbMonitorResponseBody) SetSuccess(v bool) *DescribeDrdsInstanceDbMonitorResponseBody {
	s.Success = &v
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
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDrdsInstanceDbMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeDrdsInstanceDbMonitorResponse) SetStatusCode(v int32) *DescribeDrdsInstanceDbMonitorResponse {
	s.StatusCode = &v
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
	RequestId *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                            `json:"Success,omitempty" xml:"Success,omitempty"`
	Tasks     *DescribeDrdsInstanceLevelTasksResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Struct"`
}

func (s DescribeDrdsInstanceLevelTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceLevelTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceLevelTasksResponseBody) SetRequestId(v string) *DescribeDrdsInstanceLevelTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrdsInstanceLevelTasksResponseBody) SetSuccess(v bool) *DescribeDrdsInstanceLevelTasksResponseBody {
	s.Success = &v
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
	AllowCancel         *bool   `json:"AllowCancel,omitempty" xml:"AllowCancel,omitempty"`
	ErrMsg              *string `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	GmtCreate           *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Progress            *int32  `json:"Progress,omitempty" xml:"Progress,omitempty"`
	ProgressDescription *string `json:"ProgressDescription,omitempty" xml:"ProgressDescription,omitempty"`
	ShowProgress        *bool   `json:"ShowProgress,omitempty" xml:"ShowProgress,omitempty"`
	TargetId            *int64  `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	TaskName            *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	TaskPhase           *string `json:"TaskPhase,omitempty" xml:"TaskPhase,omitempty"`
	TaskStatus          *int32  `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	TaskType            *int32  `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DescribeDrdsInstanceLevelTasksResponseBodyTasksTask) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceLevelTasksResponseBodyTasksTask) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask) SetAllowCancel(v bool) *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask {
	s.AllowCancel = &v
	return s
}

func (s *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask) SetErrMsg(v string) *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask {
	s.ErrMsg = &v
	return s
}

func (s *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask) SetGmtCreate(v int64) *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask {
	s.GmtCreate = &v
	return s
}

func (s *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask) SetProgress(v int32) *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask {
	s.Progress = &v
	return s
}

func (s *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask) SetProgressDescription(v string) *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask {
	s.ProgressDescription = &v
	return s
}

func (s *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask) SetShowProgress(v bool) *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask {
	s.ShowProgress = &v
	return s
}

func (s *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask) SetTargetId(v int64) *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask {
	s.TargetId = &v
	return s
}

func (s *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask) SetTaskName(v string) *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask {
	s.TaskName = &v
	return s
}

func (s *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask) SetTaskPhase(v string) *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask {
	s.TaskPhase = &v
	return s
}

func (s *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask) SetTaskStatus(v int32) *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask {
	s.TaskStatus = &v
	return s
}

func (s *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask) SetTaskType(v int32) *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask {
	s.TaskType = &v
	return s
}

type DescribeDrdsInstanceLevelTasksResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDrdsInstanceLevelTasksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeDrdsInstanceLevelTasksResponse) SetStatusCode(v int32) *DescribeDrdsInstanceLevelTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDrdsInstanceLevelTasksResponse) SetBody(v *DescribeDrdsInstanceLevelTasksResponseBody) *DescribeDrdsInstanceLevelTasksResponse {
	s.Body = v
	return s
}

type DescribeDrdsInstanceMonitorRequest struct {
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	EndTime        *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Key            *string `json:"Key,omitempty" xml:"Key,omitempty"`
	PeriodMultiple *int32  `json:"PeriodMultiple,omitempty" xml:"PeriodMultiple,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StartTime      *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
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

func (s *DescribeDrdsInstanceMonitorRequest) SetEndTime(v int64) *DescribeDrdsInstanceMonitorRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDrdsInstanceMonitorRequest) SetKey(v string) *DescribeDrdsInstanceMonitorRequest {
	s.Key = &v
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

func (s *DescribeDrdsInstanceMonitorRequest) SetStartTime(v int64) *DescribeDrdsInstanceMonitorRequest {
	s.StartTime = &v
	return s
}

type DescribeDrdsInstanceMonitorResponseBody struct {
	Data      []*DescribeDrdsInstanceMonitorResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDrdsInstanceMonitorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceMonitorResponseBody) SetData(v []*DescribeDrdsInstanceMonitorResponseBodyData) *DescribeDrdsInstanceMonitorResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDrdsInstanceMonitorResponseBody) SetRequestId(v string) *DescribeDrdsInstanceMonitorResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDrdsInstanceMonitorResponseBodyData struct {
	Key     *string                                              `json:"Key,omitempty" xml:"Key,omitempty"`
	NodeNum *int32                                               `json:"NodeNum,omitempty" xml:"NodeNum,omitempty"`
	Unit    *string                                              `json:"Unit,omitempty" xml:"Unit,omitempty"`
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

func (s *DescribeDrdsInstanceMonitorResponseBodyData) SetNodeNum(v int32) *DescribeDrdsInstanceMonitorResponseBodyData {
	s.NodeNum = &v
	return s
}

func (s *DescribeDrdsInstanceMonitorResponseBodyData) SetUnit(v string) *DescribeDrdsInstanceMonitorResponseBodyData {
	s.Unit = &v
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
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDrdsInstanceMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeDrdsInstanceMonitorResponse) SetStatusCode(v int32) *DescribeDrdsInstanceMonitorResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDrdsInstanceMonitorResponse) SetBody(v *DescribeDrdsInstanceMonitorResponseBody) *DescribeDrdsInstanceMonitorResponse {
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
	Data      *DescribeDrdsInstanceVersionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDrdsInstanceVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceVersionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceVersionResponseBody) SetData(v *DescribeDrdsInstanceVersionResponseBodyData) *DescribeDrdsInstanceVersionResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDrdsInstanceVersionResponseBody) SetRequestId(v string) *DescribeDrdsInstanceVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrdsInstanceVersionResponseBody) SetSuccess(v bool) *DescribeDrdsInstanceVersionResponseBody {
	s.Success = &v
	return s
}

type DescribeDrdsInstanceVersionResponseBodyData struct {
	InstanceVersion *string `json:"InstanceVersion,omitempty" xml:"InstanceVersion,omitempty"`
	NewestVersion   *string `json:"NewestVersion,omitempty" xml:"NewestVersion,omitempty"`
}

func (s DescribeDrdsInstanceVersionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceVersionResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceVersionResponseBodyData) SetInstanceVersion(v string) *DescribeDrdsInstanceVersionResponseBodyData {
	s.InstanceVersion = &v
	return s
}

func (s *DescribeDrdsInstanceVersionResponseBodyData) SetNewestVersion(v string) *DescribeDrdsInstanceVersionResponseBodyData {
	s.NewestVersion = &v
	return s
}

type DescribeDrdsInstanceVersionResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDrdsInstanceVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeDrdsInstanceVersionResponse) SetStatusCode(v int32) *DescribeDrdsInstanceVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDrdsInstanceVersionResponse) SetBody(v *DescribeDrdsInstanceVersionResponseBody) *DescribeDrdsInstanceVersionResponse {
	s.Body = v
	return s
}

type DescribeDrdsInstancesRequest struct {
	Description     *string                            `json:"Description,omitempty" xml:"Description,omitempty"`
	Expired         *bool                              `json:"Expired,omitempty" xml:"Expired,omitempty"`
	Mix             *bool                              `json:"Mix,omitempty" xml:"Mix,omitempty"`
	PageNumber      *int32                             `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProductVersion  *string                            `json:"ProductVersion,omitempty" xml:"ProductVersion,omitempty"`
	RegionId        *string                            `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string                            `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Tag             []*DescribeDrdsInstancesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	Type            *string                            `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeDrdsInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstancesRequest) SetDescription(v string) *DescribeDrdsInstancesRequest {
	s.Description = &v
	return s
}

func (s *DescribeDrdsInstancesRequest) SetExpired(v bool) *DescribeDrdsInstancesRequest {
	s.Expired = &v
	return s
}

func (s *DescribeDrdsInstancesRequest) SetMix(v bool) *DescribeDrdsInstancesRequest {
	s.Mix = &v
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

func (s *DescribeDrdsInstancesRequest) SetProductVersion(v string) *DescribeDrdsInstancesRequest {
	s.ProductVersion = &v
	return s
}

func (s *DescribeDrdsInstancesRequest) SetRegionId(v string) *DescribeDrdsInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDrdsInstancesRequest) SetResourceGroupId(v string) *DescribeDrdsInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDrdsInstancesRequest) SetTag(v []*DescribeDrdsInstancesRequestTag) *DescribeDrdsInstancesRequest {
	s.Tag = v
	return s
}

func (s *DescribeDrdsInstancesRequest) SetType(v string) *DescribeDrdsInstancesRequest {
	s.Type = &v
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
	Instances  *DescribeDrdsInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Struct"`
	PageNumber *int32                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total      *int32                                      `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeDrdsInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstancesResponseBody) SetInstances(v *DescribeDrdsInstancesResponseBodyInstances) *DescribeDrdsInstancesResponseBody {
	s.Instances = v
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

func (s *DescribeDrdsInstancesResponseBody) SetRequestId(v string) *DescribeDrdsInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBody) SetTotal(v int32) *DescribeDrdsInstancesResponseBody {
	s.Total = &v
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
	CommodityCode         *string                                                                  `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	CreateTime            *int64                                                                   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description           *string                                                                  `json:"Description,omitempty" xml:"Description,omitempty"`
	DrdsInstanceId        *string                                                                  `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	ExpireDate            *int64                                                                   `json:"ExpireDate,omitempty" xml:"ExpireDate,omitempty"`
	InstRole              *string                                                                  `json:"InstRole,omitempty" xml:"InstRole,omitempty"`
	InstanceSeries        *string                                                                  `json:"InstanceSeries,omitempty" xml:"InstanceSeries,omitempty"`
	InstanceSpec          *string                                                                  `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty"`
	Label                 *string                                                                  `json:"Label,omitempty" xml:"Label,omitempty"`
	MachineType           *string                                                                  `json:"MachineType,omitempty" xml:"MachineType,omitempty"`
	MasterInstanceId      *string                                                                  `json:"MasterInstanceId,omitempty" xml:"MasterInstanceId,omitempty"`
	NetworkType           *string                                                                  `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	OrderInstanceId       *string                                                                  `json:"OrderInstanceId,omitempty" xml:"OrderInstanceId,omitempty"`
	ProductVersion        *string                                                                  `json:"ProductVersion,omitempty" xml:"ProductVersion,omitempty"`
	ReadOnlyDBInstanceIds *DescribeDrdsInstancesResponseBodyInstancesInstanceReadOnlyDBInstanceIds `json:"ReadOnlyDBInstanceIds,omitempty" xml:"ReadOnlyDBInstanceIds,omitempty" type:"Struct"`
	RegionId              *string                                                                  `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId       *string                                                                  `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Status                *string                                                                  `json:"Status,omitempty" xml:"Status,omitempty"`
	Type                  *string                                                                  `json:"Type,omitempty" xml:"Type,omitempty"`
	Version               *int64                                                                   `json:"Version,omitempty" xml:"Version,omitempty"`
	VersionAction         *string                                                                  `json:"VersionAction,omitempty" xml:"VersionAction,omitempty"`
	Vips                  *DescribeDrdsInstancesResponseBodyInstancesInstanceVips                  `json:"Vips,omitempty" xml:"Vips,omitempty" type:"Struct"`
	VpcCloudInstanceId    *string                                                                  `json:"VpcCloudInstanceId,omitempty" xml:"VpcCloudInstanceId,omitempty"`
	VpcId                 *string                                                                  `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZoneId                *string                                                                  `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	Series                *string                                                                  `json:"series,omitempty" xml:"series,omitempty"`
}

func (s DescribeDrdsInstancesResponseBodyInstancesInstance) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstancesResponseBodyInstancesInstance) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetCommodityCode(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.CommodityCode = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetCreateTime(v int64) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.CreateTime = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetDescription(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.Description = &v
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

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetInstRole(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.InstRole = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetInstanceSeries(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.InstanceSeries = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetInstanceSpec(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.InstanceSpec = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetLabel(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.Label = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetMachineType(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.MachineType = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetMasterInstanceId(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.MasterInstanceId = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetNetworkType(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.NetworkType = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetOrderInstanceId(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.OrderInstanceId = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetProductVersion(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.ProductVersion = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetReadOnlyDBInstanceIds(v *DescribeDrdsInstancesResponseBodyInstancesInstanceReadOnlyDBInstanceIds) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.ReadOnlyDBInstanceIds = v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetRegionId(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.RegionId = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetResourceGroupId(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.ResourceGroupId = &v
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

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetVersion(v int64) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.Version = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetVersionAction(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.VersionAction = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetVips(v *DescribeDrdsInstancesResponseBodyInstancesInstanceVips) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.Vips = v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetVpcCloudInstanceId(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.VpcCloudInstanceId = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetVpcId(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.VpcId = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetZoneId(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.ZoneId = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstance) SetSeries(v string) *DescribeDrdsInstancesResponseBodyInstancesInstance {
	s.Series = &v
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
	IP        *string `json:"IP,omitempty" xml:"IP,omitempty"`
	Port      *string `json:"Port,omitempty" xml:"Port,omitempty"`
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty"`
	VpcId     *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	Dns       *string `json:"dns,omitempty" xml:"dns,omitempty"`
}

func (s DescribeDrdsInstancesResponseBodyInstancesInstanceVipsVip) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstancesResponseBodyInstancesInstanceVipsVip) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstanceVipsVip) SetIP(v string) *DescribeDrdsInstancesResponseBodyInstancesInstanceVipsVip {
	s.IP = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstanceVipsVip) SetPort(v string) *DescribeDrdsInstancesResponseBodyInstancesInstanceVipsVip {
	s.Port = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstanceVipsVip) SetType(v string) *DescribeDrdsInstancesResponseBodyInstancesInstanceVipsVip {
	s.Type = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstanceVipsVip) SetVpcId(v string) *DescribeDrdsInstancesResponseBodyInstancesInstanceVipsVip {
	s.VpcId = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstanceVipsVip) SetVswitchId(v string) *DescribeDrdsInstancesResponseBodyInstancesInstanceVipsVip {
	s.VswitchId = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyInstancesInstanceVipsVip) SetDns(v string) *DescribeDrdsInstancesResponseBodyInstancesInstanceVipsVip {
	s.Dns = &v
	return s
}

type DescribeDrdsInstancesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDrdsInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeDrdsInstancesResponse) SetStatusCode(v int32) *DescribeDrdsInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDrdsInstancesResponse) SetBody(v *DescribeDrdsInstancesResponseBody) *DescribeDrdsInstancesResponse {
	s.Body = v
	return s
}

type DescribeDrdsParamsRequest struct {
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	ParamLevel     *string `json:"ParamLevel,omitempty" xml:"ParamLevel,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDrdsParamsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsParamsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDrdsParamsRequest) SetDbName(v string) *DescribeDrdsParamsRequest {
	s.DbName = &v
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

func (s *DescribeDrdsParamsRequest) SetRegionId(v string) *DescribeDrdsParamsRequest {
	s.RegionId = &v
	return s
}

type DescribeDrdsParamsResponseBody struct {
	List      []*DescribeDrdsParamsResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDrdsParamsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsParamsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsParamsResponseBody) SetList(v []*DescribeDrdsParamsResponseBodyList) *DescribeDrdsParamsResponseBody {
	s.List = v
	return s
}

func (s *DescribeDrdsParamsResponseBody) SetRequestId(v string) *DescribeDrdsParamsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrdsParamsResponseBody) SetSuccess(v bool) *DescribeDrdsParamsResponseBody {
	s.Success = &v
	return s
}

type DescribeDrdsParamsResponseBodyList struct {
	DbName            *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	NeedRestart       *bool   `json:"NeedRestart,omitempty" xml:"NeedRestart,omitempty"`
	ParamDefaultValue *string `json:"ParamDefaultValue,omitempty" xml:"ParamDefaultValue,omitempty"`
	ParamDesc         *string `json:"ParamDesc,omitempty" xml:"ParamDesc,omitempty"`
	ParamEnglishName  *string `json:"ParamEnglishName,omitempty" xml:"ParamEnglishName,omitempty"`
	ParamLevel        *string `json:"ParamLevel,omitempty" xml:"ParamLevel,omitempty"`
	ParamName         *string `json:"ParamName,omitempty" xml:"ParamName,omitempty"`
	ParamRanges       *string `json:"ParamRanges,omitempty" xml:"ParamRanges,omitempty"`
	ParamType         *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
	ParamValue        *string `json:"ParamValue,omitempty" xml:"ParamValue,omitempty"`
	ParamVariableName *string `json:"ParamVariableName,omitempty" xml:"ParamVariableName,omitempty"`
}

func (s DescribeDrdsParamsResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsParamsResponseBodyList) GoString() string {
	return s.String()
}

func (s *DescribeDrdsParamsResponseBodyList) SetDbName(v string) *DescribeDrdsParamsResponseBodyList {
	s.DbName = &v
	return s
}

func (s *DescribeDrdsParamsResponseBodyList) SetNeedRestart(v bool) *DescribeDrdsParamsResponseBodyList {
	s.NeedRestart = &v
	return s
}

func (s *DescribeDrdsParamsResponseBodyList) SetParamDefaultValue(v string) *DescribeDrdsParamsResponseBodyList {
	s.ParamDefaultValue = &v
	return s
}

func (s *DescribeDrdsParamsResponseBodyList) SetParamDesc(v string) *DescribeDrdsParamsResponseBodyList {
	s.ParamDesc = &v
	return s
}

func (s *DescribeDrdsParamsResponseBodyList) SetParamEnglishName(v string) *DescribeDrdsParamsResponseBodyList {
	s.ParamEnglishName = &v
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

func (s *DescribeDrdsParamsResponseBodyList) SetParamRanges(v string) *DescribeDrdsParamsResponseBodyList {
	s.ParamRanges = &v
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

func (s *DescribeDrdsParamsResponseBodyList) SetParamVariableName(v string) *DescribeDrdsParamsResponseBodyList {
	s.ParamVariableName = &v
	return s
}

type DescribeDrdsParamsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDrdsParamsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeDrdsParamsResponse) SetStatusCode(v int32) *DescribeDrdsParamsResponse {
	s.StatusCode = &v
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
	DbInstances *DescribeDrdsRdsInstancesResponseBodyDbInstances `json:"DbInstances,omitempty" xml:"DbInstances,omitempty" type:"Struct"`
	RequestId   *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success     *bool                                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDrdsRdsInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsRdsInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsRdsInstancesResponseBody) SetDbInstances(v *DescribeDrdsRdsInstancesResponseBodyDbInstances) *DescribeDrdsRdsInstancesResponseBody {
	s.DbInstances = v
	return s
}

func (s *DescribeDrdsRdsInstancesResponseBody) SetRequestId(v string) *DescribeDrdsRdsInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrdsRdsInstancesResponseBody) SetSuccess(v bool) *DescribeDrdsRdsInstancesResponseBody {
	s.Success = &v
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
	ConnectUrl          *string `json:"ConnectUrl,omitempty" xml:"ConnectUrl,omitempty"`
	DBInstanceCPU       *string `json:"DBInstanceCPU,omitempty" xml:"DBInstanceCPU,omitempty"`
	DBInstanceClassType *string `json:"DBInstanceClassType,omitempty" xml:"DBInstanceClassType,omitempty"`
	DBInstanceId        *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DBInstanceMemory    *int64  `json:"DBInstanceMemory,omitempty" xml:"DBInstanceMemory,omitempty"`
	DBInstanceStatus    *string `json:"DBInstanceStatus,omitempty" xml:"DBInstanceStatus,omitempty"`
	DBInstanceStorage   *int64  `json:"DBInstanceStorage,omitempty" xml:"DBInstanceStorage,omitempty"`
	DbInstType          *string `json:"DbInstType,omitempty" xml:"DbInstType,omitempty"`
	DmInstanceId        *string `json:"DmInstanceId,omitempty" xml:"DmInstanceId,omitempty"`
	Engine              *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	EngineVersion       *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	NetworkType         *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	PayType             *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	Port                *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	RdsInstType         *string `json:"RdsInstType,omitempty" xml:"RdsInstType,omitempty"`
	ReadWeight          *int32  `json:"ReadWeight,omitempty" xml:"ReadWeight,omitempty"`
}

func (s DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) GoString() string {
	return s.String()
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) SetConnectUrl(v string) *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance {
	s.ConnectUrl = &v
	return s
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) SetDBInstanceCPU(v string) *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance {
	s.DBInstanceCPU = &v
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

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) SetDBInstanceMemory(v int64) *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance {
	s.DBInstanceMemory = &v
	return s
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) SetDBInstanceStatus(v string) *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance {
	s.DBInstanceStatus = &v
	return s
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) SetDBInstanceStorage(v int64) *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance {
	s.DBInstanceStorage = &v
	return s
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) SetDbInstType(v string) *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance {
	s.DbInstType = &v
	return s
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) SetDmInstanceId(v string) *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance {
	s.DmInstanceId = &v
	return s
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) SetEngine(v string) *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance {
	s.Engine = &v
	return s
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) SetEngineVersion(v string) *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) SetNetworkType(v string) *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance {
	s.NetworkType = &v
	return s
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) SetPayType(v string) *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance {
	s.PayType = &v
	return s
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) SetPort(v int32) *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance {
	s.Port = &v
	return s
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) SetRdsInstType(v string) *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance {
	s.RdsInstType = &v
	return s
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) SetReadWeight(v int32) *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance {
	s.ReadWeight = &v
	return s
}

type DescribeDrdsRdsInstancesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDrdsRdsInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeDrdsRdsInstancesResponse) SetStatusCode(v int32) *DescribeDrdsRdsInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDrdsRdsInstancesResponse) SetBody(v *DescribeDrdsRdsInstancesResponseBody) *DescribeDrdsRdsInstancesResponse {
	s.Body = v
	return s
}

type DescribeDrdsShardingDbsRequest struct {
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DbNamePattern  *string `json:"DbNamePattern,omitempty" xml:"DbNamePattern,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	PageNumber     *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeDrdsShardingDbsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsShardingDbsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDrdsShardingDbsRequest) SetDbName(v string) *DescribeDrdsShardingDbsRequest {
	s.DbName = &v
	return s
}

func (s *DescribeDrdsShardingDbsRequest) SetDbNamePattern(v string) *DescribeDrdsShardingDbsRequest {
	s.DbNamePattern = &v
	return s
}

func (s *DescribeDrdsShardingDbsRequest) SetDrdsInstanceId(v string) *DescribeDrdsShardingDbsRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeDrdsShardingDbsRequest) SetPageNumber(v int64) *DescribeDrdsShardingDbsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDrdsShardingDbsRequest) SetPageSize(v int64) *DescribeDrdsShardingDbsRequest {
	s.PageSize = &v
	return s
}

type DescribeDrdsShardingDbsResponseBody struct {
	PageNumber  *string                                         `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *string                                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ShardingDbs *DescribeDrdsShardingDbsResponseBodyShardingDbs `json:"ShardingDbs,omitempty" xml:"ShardingDbs,omitempty" type:"Struct"`
	Success     *bool                                           `json:"Success,omitempty" xml:"Success,omitempty"`
	Total       *string                                         `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeDrdsShardingDbsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsShardingDbsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsShardingDbsResponseBody) SetPageNumber(v string) *DescribeDrdsShardingDbsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDrdsShardingDbsResponseBody) SetPageSize(v string) *DescribeDrdsShardingDbsResponseBody {
	s.PageSize = &v
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

func (s *DescribeDrdsShardingDbsResponseBody) SetSuccess(v bool) *DescribeDrdsShardingDbsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDrdsShardingDbsResponseBody) SetTotal(v string) *DescribeDrdsShardingDbsResponseBody {
	s.Total = &v
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
	BlockingTimeout            *int32  `json:"BlockingTimeout,omitempty" xml:"BlockingTimeout,omitempty"`
	ConnectUrl                 *string `json:"ConnectUrl,omitempty" xml:"ConnectUrl,omitempty"`
	ConnectionProperties       *string `json:"ConnectionProperties,omitempty" xml:"ConnectionProperties,omitempty"`
	DbInstanceId               *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
	DbStatus                   *string `json:"DbStatus,omitempty" xml:"DbStatus,omitempty"`
	DbType                     *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	GroupName                  *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	IdleTimeOut                *int32  `json:"IdleTimeOut,omitempty" xml:"IdleTimeOut,omitempty"`
	MaxPoolSize                *int32  `json:"MaxPoolSize,omitempty" xml:"MaxPoolSize,omitempty"`
	MinPoolSize                *int32  `json:"MinPoolSize,omitempty" xml:"MinPoolSize,omitempty"`
	PreparedStatementCacheSize *int32  `json:"PreparedStatementCacheSize,omitempty" xml:"PreparedStatementCacheSize,omitempty"`
	ShardingDbName             *string `json:"ShardingDbName,omitempty" xml:"ShardingDbName,omitempty"`
	UserName                   *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb) GoString() string {
	return s.String()
}

func (s *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb) SetBlockingTimeout(v int32) *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb {
	s.BlockingTimeout = &v
	return s
}

func (s *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb) SetConnectUrl(v string) *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb {
	s.ConnectUrl = &v
	return s
}

func (s *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb) SetConnectionProperties(v string) *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb {
	s.ConnectionProperties = &v
	return s
}

func (s *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb) SetDbInstanceId(v string) *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb {
	s.DbInstanceId = &v
	return s
}

func (s *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb) SetDbStatus(v string) *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb {
	s.DbStatus = &v
	return s
}

func (s *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb) SetDbType(v string) *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb {
	s.DbType = &v
	return s
}

func (s *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb) SetGroupName(v string) *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb {
	s.GroupName = &v
	return s
}

func (s *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb) SetIdleTimeOut(v int32) *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb {
	s.IdleTimeOut = &v
	return s
}

func (s *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb) SetMaxPoolSize(v int32) *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb {
	s.MaxPoolSize = &v
	return s
}

func (s *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb) SetMinPoolSize(v int32) *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb {
	s.MinPoolSize = &v
	return s
}

func (s *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb) SetPreparedStatementCacheSize(v int32) *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb {
	s.PreparedStatementCacheSize = &v
	return s
}

func (s *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb) SetShardingDbName(v string) *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb {
	s.ShardingDbName = &v
	return s
}

func (s *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb) SetUserName(v string) *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb {
	s.UserName = &v
	return s
}

type DescribeDrdsShardingDbsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDrdsShardingDbsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeDrdsShardingDbsResponse) SetStatusCode(v int32) *DescribeDrdsShardingDbsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDrdsShardingDbsResponse) SetBody(v *DescribeDrdsShardingDbsResponseBody) *DescribeDrdsShardingDbsResponse {
	s.Body = v
	return s
}

type DescribeDrdsSlowSqlsRequest struct {
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	EndTime        *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ExeTime        *int64  `json:"ExeTime,omitempty" xml:"ExeTime,omitempty"`
	PageNumber     *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartTime      *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDrdsSlowSqlsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsSlowSqlsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDrdsSlowSqlsRequest) SetDbName(v string) *DescribeDrdsSlowSqlsRequest {
	s.DbName = &v
	return s
}

func (s *DescribeDrdsSlowSqlsRequest) SetDrdsInstanceId(v string) *DescribeDrdsSlowSqlsRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeDrdsSlowSqlsRequest) SetEndTime(v int64) *DescribeDrdsSlowSqlsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDrdsSlowSqlsRequest) SetExeTime(v int64) *DescribeDrdsSlowSqlsRequest {
	s.ExeTime = &v
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

func (s *DescribeDrdsSlowSqlsRequest) SetStartTime(v int64) *DescribeDrdsSlowSqlsRequest {
	s.StartTime = &v
	return s
}

type DescribeDrdsSlowSqlsResponseBody struct {
	Items      *DescribeDrdsSlowSqlsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	PageNumber *int32                                 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
	Total      *int32                                 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeDrdsSlowSqlsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsSlowSqlsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsSlowSqlsResponseBody) SetItems(v *DescribeDrdsSlowSqlsResponseBodyItems) *DescribeDrdsSlowSqlsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDrdsSlowSqlsResponseBody) SetPageNumber(v int32) *DescribeDrdsSlowSqlsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDrdsSlowSqlsResponseBody) SetPageSize(v int32) *DescribeDrdsSlowSqlsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDrdsSlowSqlsResponseBody) SetRequestId(v string) *DescribeDrdsSlowSqlsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrdsSlowSqlsResponseBody) SetSuccess(v bool) *DescribeDrdsSlowSqlsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDrdsSlowSqlsResponseBody) SetTotal(v int32) *DescribeDrdsSlowSqlsResponseBody {
	s.Total = &v
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
	Host         *string `json:"Host,omitempty" xml:"Host,omitempty"`
	ResponseTime *int64  `json:"ResponseTime,omitempty" xml:"ResponseTime,omitempty"`
	Schema       *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
	SendTime     *int64  `json:"SendTime,omitempty" xml:"SendTime,omitempty"`
	Sql          *string `json:"Sql,omitempty" xml:"Sql,omitempty"`
}

func (s DescribeDrdsSlowSqlsResponseBodyItemsItem) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsSlowSqlsResponseBodyItemsItem) GoString() string {
	return s.String()
}

func (s *DescribeDrdsSlowSqlsResponseBodyItemsItem) SetHost(v string) *DescribeDrdsSlowSqlsResponseBodyItemsItem {
	s.Host = &v
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

func (s *DescribeDrdsSlowSqlsResponseBodyItemsItem) SetSendTime(v int64) *DescribeDrdsSlowSqlsResponseBodyItemsItem {
	s.SendTime = &v
	return s
}

func (s *DescribeDrdsSlowSqlsResponseBodyItemsItem) SetSql(v string) *DescribeDrdsSlowSqlsResponseBodyItemsItem {
	s.Sql = &v
	return s
}

type DescribeDrdsSlowSqlsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDrdsSlowSqlsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeDrdsSlowSqlsResponse) SetStatusCode(v int32) *DescribeDrdsSlowSqlsResponse {
	s.StatusCode = &v
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
	Data      *DescribeDrdsSqlAuditStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDrdsSqlAuditStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsSqlAuditStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsSqlAuditStatusResponseBody) SetData(v *DescribeDrdsSqlAuditStatusResponseBodyData) *DescribeDrdsSqlAuditStatusResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDrdsSqlAuditStatusResponseBody) SetRequestId(v string) *DescribeDrdsSqlAuditStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrdsSqlAuditStatusResponseBody) SetSuccess(v bool) *DescribeDrdsSqlAuditStatusResponseBody {
	s.Success = &v
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
	DbName            *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	Detailed          *string `json:"Detailed,omitempty" xml:"Detailed,omitempty"`
	Enabled           *string `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	ExtraAliUid       *int64  `json:"ExtraAliUid,omitempty" xml:"ExtraAliUid,omitempty"`
	ExtraSlsLogStore  *string `json:"ExtraSlsLogStore,omitempty" xml:"ExtraSlsLogStore,omitempty"`
	ExtraSlsProject   *string `json:"ExtraSlsProject,omitempty" xml:"ExtraSlsProject,omitempty"`
	ExtraWriteEnabled *bool   `json:"ExtraWriteEnabled,omitempty" xml:"ExtraWriteEnabled,omitempty"`
}

func (s DescribeDrdsSqlAuditStatusResponseBodyDataData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsSqlAuditStatusResponseBodyDataData) GoString() string {
	return s.String()
}

func (s *DescribeDrdsSqlAuditStatusResponseBodyDataData) SetDbName(v string) *DescribeDrdsSqlAuditStatusResponseBodyDataData {
	s.DbName = &v
	return s
}

func (s *DescribeDrdsSqlAuditStatusResponseBodyDataData) SetDetailed(v string) *DescribeDrdsSqlAuditStatusResponseBodyDataData {
	s.Detailed = &v
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

func (s *DescribeDrdsSqlAuditStatusResponseBodyDataData) SetExtraSlsLogStore(v string) *DescribeDrdsSqlAuditStatusResponseBodyDataData {
	s.ExtraSlsLogStore = &v
	return s
}

func (s *DescribeDrdsSqlAuditStatusResponseBodyDataData) SetExtraSlsProject(v string) *DescribeDrdsSqlAuditStatusResponseBodyDataData {
	s.ExtraSlsProject = &v
	return s
}

func (s *DescribeDrdsSqlAuditStatusResponseBodyDataData) SetExtraWriteEnabled(v bool) *DescribeDrdsSqlAuditStatusResponseBodyDataData {
	s.ExtraWriteEnabled = &v
	return s
}

type DescribeDrdsSqlAuditStatusResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDrdsSqlAuditStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeDrdsSqlAuditStatusResponse) SetStatusCode(v int32) *DescribeDrdsSqlAuditStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDrdsSqlAuditStatusResponse) SetBody(v *DescribeDrdsSqlAuditStatusResponseBody) *DescribeDrdsSqlAuditStatusResponse {
	s.Body = v
	return s
}

type DescribeDrdsTasksRequest struct {
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	TaskType       *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DescribeDrdsTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeDrdsTasksRequest) SetDbName(v string) *DescribeDrdsTasksRequest {
	s.DbName = &v
	return s
}

func (s *DescribeDrdsTasksRequest) SetDrdsInstanceId(v string) *DescribeDrdsTasksRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeDrdsTasksRequest) SetTaskType(v string) *DescribeDrdsTasksRequest {
	s.TaskType = &v
	return s
}

type DescribeDrdsTasksResponseBody struct {
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
	Tasks     *DescribeDrdsTasksResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Struct"`
}

func (s DescribeDrdsTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsTasksResponseBody) SetRequestId(v string) *DescribeDrdsTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrdsTasksResponseBody) SetSuccess(v bool) *DescribeDrdsTasksResponseBody {
	s.Success = &v
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
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Id      *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	State   *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s DescribeDrdsTasksResponseBodyTasksTask) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsTasksResponseBodyTasksTask) GoString() string {
	return s.String()
}

func (s *DescribeDrdsTasksResponseBodyTasksTask) SetContent(v string) *DescribeDrdsTasksResponseBodyTasksTask {
	s.Content = &v
	return s
}

func (s *DescribeDrdsTasksResponseBodyTasksTask) SetId(v int64) *DescribeDrdsTasksResponseBodyTasksTask {
	s.Id = &v
	return s
}

func (s *DescribeDrdsTasksResponseBodyTasksTask) SetState(v string) *DescribeDrdsTasksResponseBodyTasksTask {
	s.State = &v
	return s
}

type DescribeDrdsTasksResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDrdsTasksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeDrdsTasksResponse) SetStatusCode(v int32) *DescribeDrdsTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDrdsTasksResponse) SetBody(v *DescribeDrdsTasksResponseBody) *DescribeDrdsTasksResponse {
	s.Body = v
	return s
}

type DescribeExpandLogicTableInfoListRequest struct {
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s DescribeExpandLogicTableInfoListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpandLogicTableInfoListRequest) GoString() string {
	return s.String()
}

func (s *DescribeExpandLogicTableInfoListRequest) SetDbName(v string) *DescribeExpandLogicTableInfoListRequest {
	s.DbName = &v
	return s
}

func (s *DescribeExpandLogicTableInfoListRequest) SetDrdsInstanceId(v string) *DescribeExpandLogicTableInfoListRequest {
	s.DrdsInstanceId = &v
	return s
}

type DescribeExpandLogicTableInfoListResponseBody struct {
	Data      *DescribeExpandLogicTableInfoListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeExpandLogicTableInfoListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpandLogicTableInfoListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExpandLogicTableInfoListResponseBody) SetData(v *DescribeExpandLogicTableInfoListResponseBodyData) *DescribeExpandLogicTableInfoListResponseBody {
	s.Data = v
	return s
}

func (s *DescribeExpandLogicTableInfoListResponseBody) SetRequestId(v string) *DescribeExpandLogicTableInfoListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeExpandLogicTableInfoListResponseBody) SetSuccess(v bool) *DescribeExpandLogicTableInfoListResponseBody {
	s.Success = &v
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
	ShardDbKey *string `json:"ShardDbKey,omitempty" xml:"ShardDbKey,omitempty"`
	ShardTbKey *string `json:"ShardTbKey,omitempty" xml:"ShardTbKey,omitempty"`
	TableName  *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeExpandLogicTableInfoListResponseBodyDataData) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpandLogicTableInfoListResponseBodyDataData) GoString() string {
	return s.String()
}

func (s *DescribeExpandLogicTableInfoListResponseBodyDataData) SetShardDbKey(v string) *DescribeExpandLogicTableInfoListResponseBodyDataData {
	s.ShardDbKey = &v
	return s
}

func (s *DescribeExpandLogicTableInfoListResponseBodyDataData) SetShardTbKey(v string) *DescribeExpandLogicTableInfoListResponseBodyDataData {
	s.ShardTbKey = &v
	return s
}

func (s *DescribeExpandLogicTableInfoListResponseBodyDataData) SetTableName(v string) *DescribeExpandLogicTableInfoListResponseBodyDataData {
	s.TableName = &v
	return s
}

type DescribeExpandLogicTableInfoListResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeExpandLogicTableInfoListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeExpandLogicTableInfoListResponse) SetStatusCode(v int32) *DescribeExpandLogicTableInfoListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeExpandLogicTableInfoListResponse) SetBody(v *DescribeExpandLogicTableInfoListResponseBody) *DescribeExpandLogicTableInfoListResponse {
	s.Body = v
	return s
}

type DescribeHotDbListRequest struct {
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s DescribeHotDbListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeHotDbListRequest) GoString() string {
	return s.String()
}

func (s *DescribeHotDbListRequest) SetDbName(v string) *DescribeHotDbListRequest {
	s.DbName = &v
	return s
}

func (s *DescribeHotDbListRequest) SetDrdsInstanceId(v string) *DescribeHotDbListRequest {
	s.DrdsInstanceId = &v
	return s
}

type DescribeHotDbListResponseBody struct {
	Data      *DescribeHotDbListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Msg       *string                            `json:"Msg,omitempty" xml:"Msg,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeHotDbListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeHotDbListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHotDbListResponseBody) SetData(v *DescribeHotDbListResponseBodyData) *DescribeHotDbListResponseBody {
	s.Data = v
	return s
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

type DescribeHotDbListResponseBodyData struct {
	List       *DescribeHotDbListResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Struct"`
	RandomCode *string                                `json:"RandomCode,omitempty" xml:"RandomCode,omitempty"`
}

func (s DescribeHotDbListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeHotDbListResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeHotDbListResponseBodyData) SetList(v *DescribeHotDbListResponseBodyDataList) *DescribeHotDbListResponseBodyData {
	s.List = v
	return s
}

func (s *DescribeHotDbListResponseBodyData) SetRandomCode(v string) *DescribeHotDbListResponseBodyData {
	s.RandomCode = &v
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
	HotDbList    *DescribeHotDbListResponseBodyDataListInstanceDbHotDbList `json:"HotDbList,omitempty" xml:"HotDbList,omitempty" type:"Struct"`
	InstanceName *string                                                   `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
}

func (s DescribeHotDbListResponseBodyDataListInstanceDb) String() string {
	return tea.Prettify(s)
}

func (s DescribeHotDbListResponseBodyDataListInstanceDb) GoString() string {
	return s.String()
}

func (s *DescribeHotDbListResponseBodyDataListInstanceDb) SetHotDbList(v *DescribeHotDbListResponseBodyDataListInstanceDbHotDbList) *DescribeHotDbListResponseBodyDataListInstanceDb {
	s.HotDbList = v
	return s
}

func (s *DescribeHotDbListResponseBodyDataListInstanceDb) SetInstanceName(v string) *DescribeHotDbListResponseBodyDataListInstanceDb {
	s.InstanceName = &v
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeHotDbListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeHotDbListResponse) SetStatusCode(v int32) *DescribeHotDbListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHotDbListResponse) SetBody(v *DescribeHotDbListResponseBody) *DescribeHotDbListResponse {
	s.Body = v
	return s
}

type DescribeInstDbLogInfoRequest struct {
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s DescribeInstDbLogInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstDbLogInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstDbLogInfoRequest) SetDbName(v string) *DescribeInstDbLogInfoRequest {
	s.DbName = &v
	return s
}

func (s *DescribeInstDbLogInfoRequest) SetDrdsInstanceId(v string) *DescribeInstDbLogInfoRequest {
	s.DrdsInstanceId = &v
	return s
}

type DescribeInstDbLogInfoResponseBody struct {
	LogTimeRange *DescribeInstDbLogInfoResponseBodyLogTimeRange `json:"LogTimeRange,omitempty" xml:"LogTimeRange,omitempty" type:"Struct"`
	RequestId    *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeInstDbLogInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstDbLogInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstDbLogInfoResponseBody) SetLogTimeRange(v *DescribeInstDbLogInfoResponseBodyLogTimeRange) *DescribeInstDbLogInfoResponseBody {
	s.LogTimeRange = v
	return s
}

func (s *DescribeInstDbLogInfoResponseBody) SetRequestId(v string) *DescribeInstDbLogInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstDbLogInfoResponseBody) SetSuccess(v bool) *DescribeInstDbLogInfoResponseBody {
	s.Success = &v
	return s
}

type DescribeInstDbLogInfoResponseBodyLogTimeRange struct {
	SupportLatestTime *int64 `json:"SupportLatestTime,omitempty" xml:"SupportLatestTime,omitempty"`
	SupportOldestTime *int64 `json:"SupportOldestTime,omitempty" xml:"SupportOldestTime,omitempty"`
}

func (s DescribeInstDbLogInfoResponseBodyLogTimeRange) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstDbLogInfoResponseBodyLogTimeRange) GoString() string {
	return s.String()
}

func (s *DescribeInstDbLogInfoResponseBodyLogTimeRange) SetSupportLatestTime(v int64) *DescribeInstDbLogInfoResponseBodyLogTimeRange {
	s.SupportLatestTime = &v
	return s
}

func (s *DescribeInstDbLogInfoResponseBodyLogTimeRange) SetSupportOldestTime(v int64) *DescribeInstDbLogInfoResponseBodyLogTimeRange {
	s.SupportOldestTime = &v
	return s
}

type DescribeInstDbLogInfoResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeInstDbLogInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeInstDbLogInfoResponse) SetStatusCode(v int32) *DescribeInstDbLogInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstDbLogInfoResponse) SetBody(v *DescribeInstDbLogInfoResponseBody) *DescribeInstDbLogInfoResponse {
	s.Body = v
	return s
}

type DescribeInstDbSlsInfoRequest struct {
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s DescribeInstDbSlsInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstDbSlsInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstDbSlsInfoRequest) SetDbName(v string) *DescribeInstDbSlsInfoRequest {
	s.DbName = &v
	return s
}

func (s *DescribeInstDbSlsInfoRequest) SetDrdsInstanceId(v string) *DescribeInstDbSlsInfoRequest {
	s.DrdsInstanceId = &v
	return s
}

type DescribeInstDbSlsInfoResponseBody struct {
	AuditInfo *DescribeInstDbSlsInfoResponseBodyAuditInfo `json:"AuditInfo,omitempty" xml:"AuditInfo,omitempty" type:"Struct"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeInstDbSlsInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstDbSlsInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstDbSlsInfoResponseBody) SetAuditInfo(v *DescribeInstDbSlsInfoResponseBodyAuditInfo) *DescribeInstDbSlsInfoResponseBody {
	s.AuditInfo = v
	return s
}

func (s *DescribeInstDbSlsInfoResponseBody) SetRequestId(v string) *DescribeInstDbSlsInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstDbSlsInfoResponseBody) SetSuccess(v bool) *DescribeInstDbSlsInfoResponseBody {
	s.Success = &v
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeInstDbSlsInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeInstDbSlsInfoResponse) SetStatusCode(v int32) *DescribeInstDbSlsInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstDbSlsInfoResponse) SetBody(v *DescribeInstDbSlsInfoResponseBody) *DescribeInstDbSlsInfoResponse {
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
	InstanceAccounts *DescribeInstanceAccountsResponseBodyInstanceAccounts `json:"InstanceAccounts,omitempty" xml:"InstanceAccounts,omitempty" type:"Struct"`
	RequestId        *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success          *bool                                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeInstanceAccountsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAccountsResponseBody) SetInstanceAccounts(v *DescribeInstanceAccountsResponseBodyInstanceAccounts) *DescribeInstanceAccountsResponseBody {
	s.InstanceAccounts = v
	return s
}

func (s *DescribeInstanceAccountsResponseBody) SetRequestId(v string) *DescribeInstanceAccountsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceAccountsResponseBody) SetSuccess(v bool) *DescribeInstanceAccountsResponseBody {
	s.Success = &v
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
	AccountName  *string                                                                          `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AccountType  *int32                                                                           `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	DbPrivileges *DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccountDbPrivileges `json:"DbPrivileges,omitempty" xml:"DbPrivileges,omitempty" type:"Struct"`
	Description  *string                                                                          `json:"Description,omitempty" xml:"Description,omitempty"`
	Host         *string                                                                          `json:"Host,omitempty" xml:"Host,omitempty"`
}

func (s DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccount) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccount) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccount) SetAccountName(v string) *DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccount {
	s.AccountName = &v
	return s
}

func (s *DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccount) SetAccountType(v int32) *DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccount {
	s.AccountType = &v
	return s
}

func (s *DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccount) SetDbPrivileges(v *DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccountDbPrivileges) *DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccount {
	s.DbPrivileges = v
	return s
}

func (s *DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccount) SetDescription(v string) *DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccount {
	s.Description = &v
	return s
}

func (s *DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccount) SetHost(v string) *DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccount {
	s.Host = &v
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeInstanceAccountsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeInstanceAccountsResponse) SetStatusCode(v int32) *DescribeInstanceAccountsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceAccountsResponse) SetBody(v *DescribeInstanceAccountsResponseBody) *DescribeInstanceAccountsResponse {
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
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeInstanceSwitchAzoneResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success   *bool                                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeInstanceSwitchAzoneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSwitchAzoneResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSwitchAzoneResponseBody) SetRequestId(v string) *DescribeInstanceSwitchAzoneResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceSwitchAzoneResponseBody) SetResult(v *DescribeInstanceSwitchAzoneResponseBodyResult) *DescribeInstanceSwitchAzoneResponseBody {
	s.Result = v
	return s
}

func (s *DescribeInstanceSwitchAzoneResponseBody) SetSuccess(v bool) *DescribeInstanceSwitchAzoneResponseBody {
	s.Success = &v
	return s
}

type DescribeInstanceSwitchAzoneResponseBodyResult struct {
	OriginAzoneId *string                                                    `json:"OriginAzoneId,omitempty" xml:"OriginAzoneId,omitempty"`
	RegionId      *string                                                    `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SwitchAble    *bool                                                      `json:"SwitchAble,omitempty" xml:"SwitchAble,omitempty"`
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

func (s *DescribeInstanceSwitchAzoneResponseBodyResult) SetRegionId(v string) *DescribeInstanceSwitchAzoneResponseBodyResult {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceSwitchAzoneResponseBodyResult) SetSwitchAble(v bool) *DescribeInstanceSwitchAzoneResponseBodyResult {
	s.SwitchAble = &v
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
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeInstanceSwitchAzoneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeInstanceSwitchAzoneResponse) SetStatusCode(v int32) *DescribeInstanceSwitchAzoneResponse {
	s.StatusCode = &v
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
	RequestId *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                              `json:"Success,omitempty" xml:"Success,omitempty"`
	VpcInfos  *DescribeInstanceSwitchNetworkResponseBodyVpcInfos `json:"VpcInfos,omitempty" xml:"VpcInfos,omitempty" type:"Struct"`
}

func (s DescribeInstanceSwitchNetworkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSwitchNetworkResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSwitchNetworkResponseBody) SetRequestId(v string) *DescribeInstanceSwitchNetworkResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceSwitchNetworkResponseBody) SetSuccess(v bool) *DescribeInstanceSwitchNetworkResponseBody {
	s.Success = &v
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
	RegionId     *string                                                               `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	VpcId        *string                                                               `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VpcName      *string                                                               `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
	VswitchInfos *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfos `json:"VswitchInfos,omitempty" xml:"VswitchInfos,omitempty" type:"Struct"`
}

func (s DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfo) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfo) SetRegionId(v string) *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfo {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfo) SetVpcId(v string) *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfo {
	s.VpcId = &v
	return s
}

func (s *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfo) SetVpcName(v string) *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfo {
	s.VpcName = &v
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
	AzoneId       *string `json:"AzoneId,omitempty" xml:"AzoneId,omitempty"`
	DrdsSupported *bool   `json:"DrdsSupported,omitempty" xml:"DrdsSupported,omitempty"`
	VpcId         *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VswitchId     *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	VswitchName   *string `json:"VswitchName,omitempty" xml:"VswitchName,omitempty"`
}

func (s DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfosVswitchInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfosVswitchInfo) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfosVswitchInfo) SetAzoneId(v string) *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfosVswitchInfo {
	s.AzoneId = &v
	return s
}

func (s *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfosVswitchInfo) SetDrdsSupported(v bool) *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfosVswitchInfo {
	s.DrdsSupported = &v
	return s
}

func (s *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfosVswitchInfo) SetVpcId(v string) *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfosVswitchInfo {
	s.VpcId = &v
	return s
}

func (s *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfosVswitchInfo) SetVswitchId(v string) *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfosVswitchInfo {
	s.VswitchId = &v
	return s
}

func (s *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfosVswitchInfo) SetVswitchName(v string) *DescribeInstanceSwitchNetworkResponseBodyVpcInfosVpcInfoVswitchInfosVswitchInfo {
	s.VswitchName = &v
	return s
}

type DescribeInstanceSwitchNetworkResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeInstanceSwitchNetworkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeInstanceSwitchNetworkResponse) SetStatusCode(v int32) *DescribeInstanceSwitchNetworkResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceSwitchNetworkResponse) SetBody(v *DescribeInstanceSwitchNetworkResponseBody) *DescribeInstanceSwitchNetworkResponse {
	s.Body = v
	return s
}

type DescribePreCheckResultRequest struct {
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TaskId         *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribePreCheckResultRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePreCheckResultRequest) GoString() string {
	return s.String()
}

func (s *DescribePreCheckResultRequest) SetDrdsInstanceId(v string) *DescribePreCheckResultRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribePreCheckResultRequest) SetRegionId(v string) *DescribePreCheckResultRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePreCheckResultRequest) SetTaskId(v string) *DescribePreCheckResultRequest {
	s.TaskId = &v
	return s
}

type DescribePreCheckResultResponseBody struct {
	PreCheckResult *DescribePreCheckResultResponseBodyPreCheckResult `json:"PreCheckResult,omitempty" xml:"PreCheckResult,omitempty" type:"Struct"`
	RequestId      *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribePreCheckResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePreCheckResultResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePreCheckResultResponseBody) SetPreCheckResult(v *DescribePreCheckResultResponseBodyPreCheckResult) *DescribePreCheckResultResponseBody {
	s.PreCheckResult = v
	return s
}

func (s *DescribePreCheckResultResponseBody) SetRequestId(v string) *DescribePreCheckResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePreCheckResultResponseBody) SetSuccess(v bool) *DescribePreCheckResultResponseBody {
	s.Success = &v
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
	ErrorMsgParams   []*string `json:"ErrorMsgParams,omitempty" xml:"ErrorMsgParams,omitempty" type:"Repeated"`
	PreCheckItemName *string   `json:"PreCheckItemName,omitempty" xml:"PreCheckItemName,omitempty"`
	State            *string   `json:"State,omitempty" xml:"State,omitempty"`
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

func (s *DescribePreCheckResultResponseBodyPreCheckResultSubCheckItems) SetErrorMsgParams(v []*string) *DescribePreCheckResultResponseBodyPreCheckResultSubCheckItems {
	s.ErrorMsgParams = v
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

type DescribePreCheckResultResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribePreCheckResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribePreCheckResultResponse) SetStatusCode(v int32) *DescribePreCheckResultResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePreCheckResultResponse) SetBody(v *DescribePreCheckResultResponseBody) *DescribePreCheckResultResponse {
	s.Body = v
	return s
}

type DescribeRDSPerformanceRequest struct {
	DbInstType     *string `json:"DbInstType,omitempty" xml:"DbInstType,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	EndTime        *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Keys           *string `json:"Keys,omitempty" xml:"Keys,omitempty"`
	RdsInstanceId  *string `json:"RdsInstanceId,omitempty" xml:"RdsInstanceId,omitempty"`
	StartTime      *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeRDSPerformanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRDSPerformanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeRDSPerformanceRequest) SetDbInstType(v string) *DescribeRDSPerformanceRequest {
	s.DbInstType = &v
	return s
}

func (s *DescribeRDSPerformanceRequest) SetDrdsInstanceId(v string) *DescribeRDSPerformanceRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeRDSPerformanceRequest) SetEndTime(v int64) *DescribeRDSPerformanceRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeRDSPerformanceRequest) SetKeys(v string) *DescribeRDSPerformanceRequest {
	s.Keys = &v
	return s
}

func (s *DescribeRDSPerformanceRequest) SetRdsInstanceId(v string) *DescribeRDSPerformanceRequest {
	s.RdsInstanceId = &v
	return s
}

func (s *DescribeRDSPerformanceRequest) SetStartTime(v int64) *DescribeRDSPerformanceRequest {
	s.StartTime = &v
	return s
}

type DescribeRDSPerformanceResponseBody struct {
	Data      []*DescribeRDSPerformanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeRDSPerformanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRDSPerformanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRDSPerformanceResponseBody) SetData(v []*DescribeRDSPerformanceResponseBodyData) *DescribeRDSPerformanceResponseBody {
	s.Data = v
	return s
}

func (s *DescribeRDSPerformanceResponseBody) SetRequestId(v string) *DescribeRDSPerformanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRDSPerformanceResponseBody) SetSuccess(v bool) *DescribeRDSPerformanceResponseBody {
	s.Success = &v
	return s
}

type DescribeRDSPerformanceResponseBodyData struct {
	Key      *string                                         `json:"Key,omitempty" xml:"Key,omitempty"`
	NodeName *string                                         `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	NodeNum  *int32                                          `json:"NodeNum,omitempty" xml:"NodeNum,omitempty"`
	Unit     *string                                         `json:"Unit,omitempty" xml:"Unit,omitempty"`
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

func (s *DescribeRDSPerformanceResponseBodyData) SetNodeNum(v int32) *DescribeRDSPerformanceResponseBodyData {
	s.NodeNum = &v
	return s
}

func (s *DescribeRDSPerformanceResponseBodyData) SetUnit(v string) *DescribeRDSPerformanceResponseBodyData {
	s.Unit = &v
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeRDSPerformanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeRDSPerformanceResponse) SetStatusCode(v int32) *DescribeRDSPerformanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRDSPerformanceResponse) SetBody(v *DescribeRDSPerformanceResponseBody) *DescribeRDSPerformanceResponse {
	s.Body = v
	return s
}

type DescribeRdsCommodityRequest struct {
	CommodityCode  *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	OrderType      *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
}

func (s DescribeRdsCommodityRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsCommodityRequest) GoString() string {
	return s.String()
}

func (s *DescribeRdsCommodityRequest) SetCommodityCode(v string) *DescribeRdsCommodityRequest {
	s.CommodityCode = &v
	return s
}

func (s *DescribeRdsCommodityRequest) SetDrdsInstanceId(v string) *DescribeRdsCommodityRequest {
	s.DrdsInstanceId = &v
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeRdsCommodityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeRdsCommodityResponse) SetStatusCode(v int32) *DescribeRdsCommodityResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRdsCommodityResponse) SetBody(v *DescribeRdsCommodityResponseBody) *DescribeRdsCommodityResponse {
	s.Body = v
	return s
}

type DescribeRdsPerformanceSummaryRequest struct {
	DrdsInstanceId *string   `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	RdsInstanceId  []*string `json:"RdsInstanceId,omitempty" xml:"RdsInstanceId,omitempty" type:"Repeated"`
	RegionId       *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRdsPerformanceSummaryRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsPerformanceSummaryRequest) GoString() string {
	return s.String()
}

func (s *DescribeRdsPerformanceSummaryRequest) SetDrdsInstanceId(v string) *DescribeRdsPerformanceSummaryRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeRdsPerformanceSummaryRequest) SetRdsInstanceId(v []*string) *DescribeRdsPerformanceSummaryRequest {
	s.RdsInstanceId = v
	return s
}

func (s *DescribeRdsPerformanceSummaryRequest) SetRegionId(v string) *DescribeRdsPerformanceSummaryRequest {
	s.RegionId = &v
	return s
}

type DescribeRdsPerformanceSummaryResponseBody struct {
	RdsPerformanceInfos []*DescribeRdsPerformanceSummaryResponseBodyRdsPerformanceInfos `json:"RdsPerformanceInfos,omitempty" xml:"RdsPerformanceInfos,omitempty" type:"Repeated"`
	RequestId           *string                                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success             *bool                                                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeRdsPerformanceSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsPerformanceSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRdsPerformanceSummaryResponseBody) SetRdsPerformanceInfos(v []*DescribeRdsPerformanceSummaryResponseBodyRdsPerformanceInfos) *DescribeRdsPerformanceSummaryResponseBody {
	s.RdsPerformanceInfos = v
	return s
}

func (s *DescribeRdsPerformanceSummaryResponseBody) SetRequestId(v string) *DescribeRdsPerformanceSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRdsPerformanceSummaryResponseBody) SetSuccess(v bool) *DescribeRdsPerformanceSummaryResponseBody {
	s.Success = &v
	return s
}

type DescribeRdsPerformanceSummaryResponseBodyRdsPerformanceInfos struct {
	ActiveSessions *int32   `json:"ActiveSessions,omitempty" xml:"ActiveSessions,omitempty"`
	Cpu            *float32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	Iops           *float32 `json:"Iops,omitempty" xml:"Iops,omitempty"`
	RdsId          *string  `json:"RdsId,omitempty" xml:"RdsId,omitempty"`
	SpaceUsage     *int64   `json:"SpaceUsage,omitempty" xml:"SpaceUsage,omitempty"`
	TotalSessions  *int32   `json:"TotalSessions,omitempty" xml:"TotalSessions,omitempty"`
}

func (s DescribeRdsPerformanceSummaryResponseBodyRdsPerformanceInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsPerformanceSummaryResponseBodyRdsPerformanceInfos) GoString() string {
	return s.String()
}

func (s *DescribeRdsPerformanceSummaryResponseBodyRdsPerformanceInfos) SetActiveSessions(v int32) *DescribeRdsPerformanceSummaryResponseBodyRdsPerformanceInfos {
	s.ActiveSessions = &v
	return s
}

func (s *DescribeRdsPerformanceSummaryResponseBodyRdsPerformanceInfos) SetCpu(v float32) *DescribeRdsPerformanceSummaryResponseBodyRdsPerformanceInfos {
	s.Cpu = &v
	return s
}

func (s *DescribeRdsPerformanceSummaryResponseBodyRdsPerformanceInfos) SetIops(v float32) *DescribeRdsPerformanceSummaryResponseBodyRdsPerformanceInfos {
	s.Iops = &v
	return s
}

func (s *DescribeRdsPerformanceSummaryResponseBodyRdsPerformanceInfos) SetRdsId(v string) *DescribeRdsPerformanceSummaryResponseBodyRdsPerformanceInfos {
	s.RdsId = &v
	return s
}

func (s *DescribeRdsPerformanceSummaryResponseBodyRdsPerformanceInfos) SetSpaceUsage(v int64) *DescribeRdsPerformanceSummaryResponseBodyRdsPerformanceInfos {
	s.SpaceUsage = &v
	return s
}

func (s *DescribeRdsPerformanceSummaryResponseBodyRdsPerformanceInfos) SetTotalSessions(v int32) *DescribeRdsPerformanceSummaryResponseBodyRdsPerformanceInfos {
	s.TotalSessions = &v
	return s
}

type DescribeRdsPerformanceSummaryResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeRdsPerformanceSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeRdsPerformanceSummaryResponse) SetStatusCode(v int32) *DescribeRdsPerformanceSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRdsPerformanceSummaryResponse) SetBody(v *DescribeRdsPerformanceSummaryResponseBody) *DescribeRdsPerformanceSummaryResponse {
	s.Body = v
	return s
}

type DescribeRdsSuperAccountInstancesRequest struct {
	DbInstType     *string   `json:"DbInstType,omitempty" xml:"DbInstType,omitempty"`
	DrdsInstanceId *string   `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	RdsInstance    []*string `json:"RdsInstance,omitempty" xml:"RdsInstance,omitempty" type:"Repeated"`
}

func (s DescribeRdsSuperAccountInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsSuperAccountInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeRdsSuperAccountInstancesRequest) SetDbInstType(v string) *DescribeRdsSuperAccountInstancesRequest {
	s.DbInstType = &v
	return s
}

func (s *DescribeRdsSuperAccountInstancesRequest) SetDrdsInstanceId(v string) *DescribeRdsSuperAccountInstancesRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeRdsSuperAccountInstancesRequest) SetRdsInstance(v []*string) *DescribeRdsSuperAccountInstancesRequest {
	s.RdsInstance = v
	return s
}

type DescribeRdsSuperAccountInstancesResponseBody struct {
	DbInstances *DescribeRdsSuperAccountInstancesResponseBodyDbInstances `json:"DbInstances,omitempty" xml:"DbInstances,omitempty" type:"Struct"`
	RequestId   *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRdsSuperAccountInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsSuperAccountInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRdsSuperAccountInstancesResponseBody) SetDbInstances(v *DescribeRdsSuperAccountInstancesResponseBodyDbInstances) *DescribeRdsSuperAccountInstancesResponseBody {
	s.DbInstances = v
	return s
}

func (s *DescribeRdsSuperAccountInstancesResponseBody) SetRequestId(v string) *DescribeRdsSuperAccountInstancesResponseBody {
	s.RequestId = &v
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
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeRdsSuperAccountInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeRdsSuperAccountInstancesResponse) SetStatusCode(v int32) *DescribeRdsSuperAccountInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRdsSuperAccountInstancesResponse) SetBody(v *DescribeRdsSuperAccountInstancesResponseBody) *DescribeRdsSuperAccountInstancesResponse {
	s.Body = v
	return s
}

type DescribeRecycleBinStatusRequest struct {
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRecycleBinStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecycleBinStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeRecycleBinStatusRequest) SetDbName(v string) *DescribeRecycleBinStatusRequest {
	s.DbName = &v
	return s
}

func (s *DescribeRecycleBinStatusRequest) SetDrdsInstanceId(v string) *DescribeRecycleBinStatusRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeRecycleBinStatusRequest) SetRegionId(v string) *DescribeRecycleBinStatusRequest {
	s.RegionId = &v
	return s
}

type DescribeRecycleBinStatusResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeRecycleBinStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecycleBinStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRecycleBinStatusResponseBody) SetRequestId(v string) *DescribeRecycleBinStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRecycleBinStatusResponseBody) SetStatus(v string) *DescribeRecycleBinStatusResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeRecycleBinStatusResponseBody) SetSuccess(v bool) *DescribeRecycleBinStatusResponseBody {
	s.Success = &v
	return s
}

type DescribeRecycleBinStatusResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeRecycleBinStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRecycleBinStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecycleBinStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeRecycleBinStatusResponse) SetHeaders(v map[string]*string) *DescribeRecycleBinStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeRecycleBinStatusResponse) SetStatusCode(v int32) *DescribeRecycleBinStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRecycleBinStatusResponse) SetBody(v *DescribeRecycleBinStatusResponseBody) *DescribeRecycleBinStatusResponse {
	s.Body = v
	return s
}

type DescribeRecycleBinTablesRequest struct {
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRecycleBinTablesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecycleBinTablesRequest) GoString() string {
	return s.String()
}

func (s *DescribeRecycleBinTablesRequest) SetDbName(v string) *DescribeRecycleBinTablesRequest {
	s.DbName = &v
	return s
}

func (s *DescribeRecycleBinTablesRequest) SetDrdsInstanceId(v string) *DescribeRecycleBinTablesRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeRecycleBinTablesRequest) SetRegionId(v string) *DescribeRecycleBinTablesRequest {
	s.RegionId = &v
	return s
}

type DescribeRecycleBinTablesResponseBody struct {
	Data      []*DescribeRecycleBinTablesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeRecycleBinTablesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecycleBinTablesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRecycleBinTablesResponseBody) SetData(v []*DescribeRecycleBinTablesResponseBodyData) *DescribeRecycleBinTablesResponseBody {
	s.Data = v
	return s
}

func (s *DescribeRecycleBinTablesResponseBody) SetRequestId(v string) *DescribeRecycleBinTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRecycleBinTablesResponseBody) SetSuccess(v bool) *DescribeRecycleBinTablesResponseBody {
	s.Success = &v
	return s
}

type DescribeRecycleBinTablesResponseBodyData struct {
	CreateTime        *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	OriginalTableName *string `json:"OriginalTableName,omitempty" xml:"OriginalTableName,omitempty"`
	TableName         *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeRecycleBinTablesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecycleBinTablesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeRecycleBinTablesResponseBodyData) SetCreateTime(v string) *DescribeRecycleBinTablesResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *DescribeRecycleBinTablesResponseBodyData) SetOriginalTableName(v string) *DescribeRecycleBinTablesResponseBodyData {
	s.OriginalTableName = &v
	return s
}

func (s *DescribeRecycleBinTablesResponseBodyData) SetTableName(v string) *DescribeRecycleBinTablesResponseBodyData {
	s.TableName = &v
	return s
}

type DescribeRecycleBinTablesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeRecycleBinTablesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRecycleBinTablesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecycleBinTablesResponse) GoString() string {
	return s.String()
}

func (s *DescribeRecycleBinTablesResponse) SetHeaders(v map[string]*string) *DescribeRecycleBinTablesResponse {
	s.Headers = v
	return s
}

func (s *DescribeRecycleBinTablesResponse) SetStatusCode(v int32) *DescribeRecycleBinTablesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRecycleBinTablesResponse) SetBody(v *DescribeRecycleBinTablesResponseBody) *DescribeRecycleBinTablesResponse {
	s.Body = v
	return s
}

type DescribeRestoreOrderRequest struct {
	BackupDbNames       *string `json:"BackupDbNames,omitempty" xml:"BackupDbNames,omitempty"`
	BackupId            *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	BackupLevel         *string `json:"BackupLevel,omitempty" xml:"BackupLevel,omitempty"`
	BackupMode          *string `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	DrdsInstanceId      *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	PreferredBackupTime *string `json:"PreferredBackupTime,omitempty" xml:"PreferredBackupTime,omitempty"`
}

func (s DescribeRestoreOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreOrderRequest) GoString() string {
	return s.String()
}

func (s *DescribeRestoreOrderRequest) SetBackupDbNames(v string) *DescribeRestoreOrderRequest {
	s.BackupDbNames = &v
	return s
}

func (s *DescribeRestoreOrderRequest) SetBackupId(v string) *DescribeRestoreOrderRequest {
	s.BackupId = &v
	return s
}

func (s *DescribeRestoreOrderRequest) SetBackupLevel(v string) *DescribeRestoreOrderRequest {
	s.BackupLevel = &v
	return s
}

func (s *DescribeRestoreOrderRequest) SetBackupMode(v string) *DescribeRestoreOrderRequest {
	s.BackupMode = &v
	return s
}

func (s *DescribeRestoreOrderRequest) SetDrdsInstanceId(v string) *DescribeRestoreOrderRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeRestoreOrderRequest) SetPreferredBackupTime(v string) *DescribeRestoreOrderRequest {
	s.PreferredBackupTime = &v
	return s
}

type DescribeRestoreOrderResponseBody struct {
	RequestId      *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RestoreOrderDO *DescribeRestoreOrderResponseBodyRestoreOrderDO `json:"RestoreOrderDO,omitempty" xml:"RestoreOrderDO,omitempty" type:"Struct"`
	Success        *bool                                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeRestoreOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreOrderResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRestoreOrderResponseBody) SetRequestId(v string) *DescribeRestoreOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRestoreOrderResponseBody) SetRestoreOrderDO(v *DescribeRestoreOrderResponseBodyRestoreOrderDO) *DescribeRestoreOrderResponseBody {
	s.RestoreOrderDO = v
	return s
}

func (s *DescribeRestoreOrderResponseBody) SetSuccess(v bool) *DescribeRestoreOrderResponseBody {
	s.Success = &v
	return s
}

type DescribeRestoreOrderResponseBodyRestoreOrderDO struct {
	DrdsOrderDOList  *DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOList  `json:"DrdsOrderDOList,omitempty" xml:"DrdsOrderDOList,omitempty" type:"Struct"`
	PolarOrderDOList *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOList `json:"PolarOrderDOList,omitempty" xml:"PolarOrderDOList,omitempty" type:"Struct"`
	RdsOrderDOList   *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOList   `json:"RdsOrderDOList,omitempty" xml:"RdsOrderDOList,omitempty" type:"Struct"`
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

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDO) SetPolarOrderDOList(v *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOList) *DescribeRestoreOrderResponseBodyRestoreOrderDO {
	s.PolarOrderDOList = v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDO) SetRdsOrderDOList(v *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOList) *DescribeRestoreOrderResponseBodyRestoreOrderDO {
	s.RdsOrderDOList = v
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
	AzoneId   *string `json:"AzoneId,omitempty" xml:"AzoneId,omitempty"`
	InstSpec  *string `json:"InstSpec,omitempty" xml:"InstSpec,omitempty"`
	Network   *string `json:"Network,omitempty" xml:"Network,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	VSwtichId *string `json:"VSwtichId,omitempty" xml:"VSwtichId,omitempty"`
	VpcId     *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOListDrdsOrderDOList) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOListDrdsOrderDOList) GoString() string {
	return s.String()
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOListDrdsOrderDOList) SetAzoneId(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOListDrdsOrderDOList {
	s.AzoneId = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOListDrdsOrderDOList) SetInstSpec(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOListDrdsOrderDOList {
	s.InstSpec = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOListDrdsOrderDOList) SetNetwork(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOListDrdsOrderDOList {
	s.Network = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOListDrdsOrderDOList) SetRegionId(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOListDrdsOrderDOList {
	s.RegionId = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOListDrdsOrderDOList) SetVSwtichId(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOListDrdsOrderDOList {
	s.VSwtichId = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOListDrdsOrderDOList) SetVpcId(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDODrdsOrderDOListDrdsOrderDOList {
	s.VpcId = &v
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
	AzoneId           *string `json:"AzoneId,omitempty" xml:"AzoneId,omitempty"`
	DbInstanceStorage *string `json:"DbInstanceStorage,omitempty" xml:"DbInstanceStorage,omitempty"`
	Engine            *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	InstanceClass     *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	Network           *string `json:"Network,omitempty" xml:"Network,omitempty"`
	Num               *int64  `json:"Num,omitempty" xml:"Num,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Version           *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList) GoString() string {
	return s.String()
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList) SetAzoneId(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList {
	s.AzoneId = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList) SetDbInstanceStorage(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList {
	s.DbInstanceStorage = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList) SetEngine(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList {
	s.Engine = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList) SetInstanceClass(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList {
	s.InstanceClass = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList) SetNetwork(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList {
	s.Network = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList) SetNum(v int64) *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList {
	s.Num = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList) SetRegionId(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList {
	s.RegionId = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList) SetVersion(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDOPolarOrderDOListPolarOrderDOList {
	s.Version = &v
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
	AzoneId           *string `json:"AzoneId,omitempty" xml:"AzoneId,omitempty"`
	DbInstanceStorage *string `json:"DbInstanceStorage,omitempty" xml:"DbInstanceStorage,omitempty"`
	Engine            *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	InstanceClass     *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	Network           *string `json:"Network,omitempty" xml:"Network,omitempty"`
	Num               *int64  `json:"Num,omitempty" xml:"Num,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Version           *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList) GoString() string {
	return s.String()
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList) SetAzoneId(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList {
	s.AzoneId = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList) SetDbInstanceStorage(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList {
	s.DbInstanceStorage = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList) SetEngine(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList {
	s.Engine = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList) SetInstanceClass(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList {
	s.InstanceClass = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList) SetNetwork(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList {
	s.Network = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList) SetNum(v int64) *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList {
	s.Num = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList) SetRegionId(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList {
	s.RegionId = &v
	return s
}

func (s *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList) SetVersion(v string) *DescribeRestoreOrderResponseBodyRestoreOrderDORdsOrderDOListRdsOrderDOList {
	s.Version = &v
	return s
}

type DescribeRestoreOrderResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeRestoreOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeRestoreOrderResponse) SetStatusCode(v int32) *DescribeRestoreOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRestoreOrderResponse) SetBody(v *DescribeRestoreOrderResponseBody) *DescribeRestoreOrderResponse {
	s.Body = v
	return s
}

type DescribeShardTaskInfoRequest struct {
	DbName          *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DrdsInstanceId  *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SourceTableName *string `json:"SourceTableName,omitempty" xml:"SourceTableName,omitempty"`
	TargetTableName *string `json:"TargetTableName,omitempty" xml:"TargetTableName,omitempty"`
}

func (s DescribeShardTaskInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeShardTaskInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeShardTaskInfoRequest) SetDbName(v string) *DescribeShardTaskInfoRequest {
	s.DbName = &v
	return s
}

func (s *DescribeShardTaskInfoRequest) SetDrdsInstanceId(v string) *DescribeShardTaskInfoRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeShardTaskInfoRequest) SetRegionId(v string) *DescribeShardTaskInfoRequest {
	s.RegionId = &v
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
	Data      *DescribeShardTaskInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeShardTaskInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeShardTaskInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeShardTaskInfoResponseBody) SetData(v *DescribeShardTaskInfoResponseBodyData) *DescribeShardTaskInfoResponseBody {
	s.Data = v
	return s
}

func (s *DescribeShardTaskInfoResponseBody) SetRequestId(v string) *DescribeShardTaskInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBody) SetSuccess(v bool) *DescribeShardTaskInfoResponseBody {
	s.Success = &v
	return s
}

type DescribeShardTaskInfoResponseBodyData struct {
	Expired         *string                                          `json:"Expired,omitempty" xml:"Expired,omitempty"`
	Full            *DescribeShardTaskInfoResponseBodyDataFull       `json:"Full,omitempty" xml:"Full,omitempty" type:"Struct"`
	FullCheck       *DescribeShardTaskInfoResponseBodyDataFullCheck  `json:"FullCheck,omitempty" xml:"FullCheck,omitempty" type:"Struct"`
	FullRevise      *DescribeShardTaskInfoResponseBodyDataFullRevise `json:"FullRevise,omitempty" xml:"FullRevise,omitempty" type:"Struct"`
	Increment       *DescribeShardTaskInfoResponseBodyDataIncrement  `json:"Increment,omitempty" xml:"Increment,omitempty" type:"Struct"`
	Progress        *string                                          `json:"Progress,omitempty" xml:"Progress,omitempty"`
	Review          *DescribeShardTaskInfoResponseBodyDataReview     `json:"Review,omitempty" xml:"Review,omitempty" type:"Struct"`
	SourceTableName *string                                          `json:"SourceTableName,omitempty" xml:"SourceTableName,omitempty"`
	Stage           *string                                          `json:"Stage,omitempty" xml:"Stage,omitempty"`
	Status          *string                                          `json:"Status,omitempty" xml:"Status,omitempty"`
	TargetTableName *string                                          `json:"TargetTableName,omitempty" xml:"TargetTableName,omitempty"`
}

func (s DescribeShardTaskInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeShardTaskInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeShardTaskInfoResponseBodyData) SetExpired(v string) *DescribeShardTaskInfoResponseBodyData {
	s.Expired = &v
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

func (s *DescribeShardTaskInfoResponseBodyData) SetIncrement(v *DescribeShardTaskInfoResponseBodyDataIncrement) *DescribeShardTaskInfoResponseBodyData {
	s.Increment = v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyData) SetProgress(v string) *DescribeShardTaskInfoResponseBodyData {
	s.Progress = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyData) SetReview(v *DescribeShardTaskInfoResponseBodyDataReview) *DescribeShardTaskInfoResponseBodyData {
	s.Review = v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyData) SetSourceTableName(v string) *DescribeShardTaskInfoResponseBodyData {
	s.SourceTableName = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyData) SetStage(v string) *DescribeShardTaskInfoResponseBodyData {
	s.Stage = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyData) SetStatus(v string) *DescribeShardTaskInfoResponseBodyData {
	s.Status = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyData) SetTargetTableName(v string) *DescribeShardTaskInfoResponseBodyData {
	s.TargetTableName = &v
	return s
}

type DescribeShardTaskInfoResponseBodyDataFull struct {
	Expired   *int32  `json:"Expired,omitempty" xml:"Expired,omitempty"`
	Progress  *int32  `json:"Progress,omitempty" xml:"Progress,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Total     *int32  `json:"Total,omitempty" xml:"Total,omitempty"`
	Tps       *int32  `json:"Tps,omitempty" xml:"Tps,omitempty"`
}

func (s DescribeShardTaskInfoResponseBodyDataFull) String() string {
	return tea.Prettify(s)
}

func (s DescribeShardTaskInfoResponseBodyDataFull) GoString() string {
	return s.String()
}

func (s *DescribeShardTaskInfoResponseBodyDataFull) SetExpired(v int32) *DescribeShardTaskInfoResponseBodyDataFull {
	s.Expired = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyDataFull) SetProgress(v int32) *DescribeShardTaskInfoResponseBodyDataFull {
	s.Progress = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyDataFull) SetStartTime(v string) *DescribeShardTaskInfoResponseBodyDataFull {
	s.StartTime = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyDataFull) SetTotal(v int32) *DescribeShardTaskInfoResponseBodyDataFull {
	s.Total = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyDataFull) SetTps(v int32) *DescribeShardTaskInfoResponseBodyDataFull {
	s.Tps = &v
	return s
}

type DescribeShardTaskInfoResponseBodyDataFullCheck struct {
	Expired   *int32  `json:"Expired,omitempty" xml:"Expired,omitempty"`
	Progress  *int32  `json:"Progress,omitempty" xml:"Progress,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Total     *int32  `json:"Total,omitempty" xml:"Total,omitempty"`
	Tps       *int32  `json:"Tps,omitempty" xml:"Tps,omitempty"`
}

func (s DescribeShardTaskInfoResponseBodyDataFullCheck) String() string {
	return tea.Prettify(s)
}

func (s DescribeShardTaskInfoResponseBodyDataFullCheck) GoString() string {
	return s.String()
}

func (s *DescribeShardTaskInfoResponseBodyDataFullCheck) SetExpired(v int32) *DescribeShardTaskInfoResponseBodyDataFullCheck {
	s.Expired = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyDataFullCheck) SetProgress(v int32) *DescribeShardTaskInfoResponseBodyDataFullCheck {
	s.Progress = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyDataFullCheck) SetStartTime(v string) *DescribeShardTaskInfoResponseBodyDataFullCheck {
	s.StartTime = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyDataFullCheck) SetTotal(v int32) *DescribeShardTaskInfoResponseBodyDataFullCheck {
	s.Total = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyDataFullCheck) SetTps(v int32) *DescribeShardTaskInfoResponseBodyDataFullCheck {
	s.Tps = &v
	return s
}

type DescribeShardTaskInfoResponseBodyDataFullRevise struct {
	Expired   *int32  `json:"Expired,omitempty" xml:"Expired,omitempty"`
	Progress  *int32  `json:"Progress,omitempty" xml:"Progress,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Total     *int32  `json:"Total,omitempty" xml:"Total,omitempty"`
	Tps       *int32  `json:"Tps,omitempty" xml:"Tps,omitempty"`
}

func (s DescribeShardTaskInfoResponseBodyDataFullRevise) String() string {
	return tea.Prettify(s)
}

func (s DescribeShardTaskInfoResponseBodyDataFullRevise) GoString() string {
	return s.String()
}

func (s *DescribeShardTaskInfoResponseBodyDataFullRevise) SetExpired(v int32) *DescribeShardTaskInfoResponseBodyDataFullRevise {
	s.Expired = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyDataFullRevise) SetProgress(v int32) *DescribeShardTaskInfoResponseBodyDataFullRevise {
	s.Progress = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyDataFullRevise) SetStartTime(v string) *DescribeShardTaskInfoResponseBodyDataFullRevise {
	s.StartTime = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyDataFullRevise) SetTotal(v int32) *DescribeShardTaskInfoResponseBodyDataFullRevise {
	s.Total = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyDataFullRevise) SetTps(v int32) *DescribeShardTaskInfoResponseBodyDataFullRevise {
	s.Tps = &v
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

type DescribeShardTaskInfoResponseBodyDataReview struct {
	Expired   *int32  `json:"Expired,omitempty" xml:"Expired,omitempty"`
	Progress  *int32  `json:"Progress,omitempty" xml:"Progress,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Total     *int32  `json:"Total,omitempty" xml:"Total,omitempty"`
	Tps       *int32  `json:"Tps,omitempty" xml:"Tps,omitempty"`
}

func (s DescribeShardTaskInfoResponseBodyDataReview) String() string {
	return tea.Prettify(s)
}

func (s DescribeShardTaskInfoResponseBodyDataReview) GoString() string {
	return s.String()
}

func (s *DescribeShardTaskInfoResponseBodyDataReview) SetExpired(v int32) *DescribeShardTaskInfoResponseBodyDataReview {
	s.Expired = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyDataReview) SetProgress(v int32) *DescribeShardTaskInfoResponseBodyDataReview {
	s.Progress = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyDataReview) SetStartTime(v string) *DescribeShardTaskInfoResponseBodyDataReview {
	s.StartTime = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyDataReview) SetTotal(v int32) *DescribeShardTaskInfoResponseBodyDataReview {
	s.Total = &v
	return s
}

func (s *DescribeShardTaskInfoResponseBodyDataReview) SetTps(v int32) *DescribeShardTaskInfoResponseBodyDataReview {
	s.Tps = &v
	return s
}

type DescribeShardTaskInfoResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeShardTaskInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeShardTaskInfoResponse) SetStatusCode(v int32) *DescribeShardTaskInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeShardTaskInfoResponse) SetBody(v *DescribeShardTaskInfoResponseBody) *DescribeShardTaskInfoResponse {
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
	RequestId         *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SqlFlashbackTasks *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasks `json:"SqlFlashbackTasks,omitempty" xml:"SqlFlashbackTasks,omitempty" type:"Struct"`
	Success           *bool                                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeSqlFlashbakTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSqlFlashbakTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSqlFlashbakTaskResponseBody) SetRequestId(v string) *DescribeSqlFlashbakTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBody) SetSqlFlashbackTasks(v *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasks) *DescribeSqlFlashbakTaskResponseBody {
	s.SqlFlashbackTasks = v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBody) SetSuccess(v bool) *DescribeSqlFlashbakTaskResponseBody {
	s.Success = &v
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
	DbName            *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DownloadUrl       *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	ExpireTime        *int64  `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	GmtCreate         *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified       *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id                *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	InstId            *string `json:"InstId,omitempty" xml:"InstId,omitempty"`
	RecallProgress    *int32  `json:"RecallProgress,omitempty" xml:"RecallProgress,omitempty"`
	RecallRestoreType *int32  `json:"RecallRestoreType,omitempty" xml:"RecallRestoreType,omitempty"`
	RecallStatus      *int32  `json:"RecallStatus,omitempty" xml:"RecallStatus,omitempty"`
	RecallType        *int32  `json:"RecallType,omitempty" xml:"RecallType,omitempty"`
	SearchEndTime     *int64  `json:"SearchEndTime,omitempty" xml:"SearchEndTime,omitempty"`
	SearchStartTime   *int64  `json:"SearchStartTime,omitempty" xml:"SearchStartTime,omitempty"`
	SqlCounter        *int64  `json:"SqlCounter,omitempty" xml:"SqlCounter,omitempty"`
	SqlPk             *string `json:"SqlPk,omitempty" xml:"SqlPk,omitempty"`
	SqlType           *string `json:"SqlType,omitempty" xml:"SqlType,omitempty"`
	TableName         *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	TraceId           *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) String() string {
	return tea.Prettify(s)
}

func (s DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) GoString() string {
	return s.String()
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) SetDbName(v string) *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask {
	s.DbName = &v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) SetDownloadUrl(v string) *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask {
	s.DownloadUrl = &v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) SetExpireTime(v int64) *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask {
	s.ExpireTime = &v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) SetGmtCreate(v int64) *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask {
	s.GmtCreate = &v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) SetGmtModified(v int64) *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask {
	s.GmtModified = &v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) SetId(v int64) *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask {
	s.Id = &v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) SetInstId(v string) *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask {
	s.InstId = &v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) SetRecallProgress(v int32) *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask {
	s.RecallProgress = &v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) SetRecallRestoreType(v int32) *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask {
	s.RecallRestoreType = &v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) SetRecallStatus(v int32) *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask {
	s.RecallStatus = &v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) SetRecallType(v int32) *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask {
	s.RecallType = &v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) SetSearchEndTime(v int64) *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask {
	s.SearchEndTime = &v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) SetSearchStartTime(v int64) *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask {
	s.SearchStartTime = &v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) SetSqlCounter(v int64) *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask {
	s.SqlCounter = &v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) SetSqlPk(v string) *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask {
	s.SqlPk = &v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) SetSqlType(v string) *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask {
	s.SqlType = &v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) SetTableName(v string) *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask {
	s.TableName = &v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) SetTraceId(v string) *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask {
	s.TraceId = &v
	return s
}

type DescribeSqlFlashbakTaskResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeSqlFlashbakTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeSqlFlashbakTaskResponse) SetStatusCode(v int32) *DescribeSqlFlashbakTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSqlFlashbakTaskResponse) SetBody(v *DescribeSqlFlashbakTaskResponseBody) *DescribeSqlFlashbakTaskResponse {
	s.Body = v
	return s
}

type DescribeTableRequest struct {
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TableName      *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeTableRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTableRequest) GoString() string {
	return s.String()
}

func (s *DescribeTableRequest) SetDbName(v string) *DescribeTableRequest {
	s.DbName = &v
	return s
}

func (s *DescribeTableRequest) SetDrdsInstanceId(v string) *DescribeTableRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeTableRequest) SetRegionId(v string) *DescribeTableRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeTableRequest) SetTableName(v string) *DescribeTableRequest {
	s.TableName = &v
	return s
}

type DescribeTableResponseBody struct {
	Data      *DescribeTableResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeTableResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTableResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTableResponseBody) SetData(v *DescribeTableResponseBodyData) *DescribeTableResponseBody {
	s.Data = v
	return s
}

func (s *DescribeTableResponseBody) SetRequestId(v string) *DescribeTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTableResponseBody) SetSuccess(v bool) *DescribeTableResponseBody {
	s.Success = &v
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
	ColumnName  *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	ColumnType  *string `json:"ColumnType,omitempty" xml:"ColumnType,omitempty"`
	Extra       *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	Index       *string `json:"Index,omitempty" xml:"Index,omitempty"`
	IsAllowNull *string `json:"IsAllowNull,omitempty" xml:"IsAllowNull,omitempty"`
	IsPk        *string `json:"IsPk,omitempty" xml:"IsPk,omitempty"`
}

func (s DescribeTableResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s DescribeTableResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *DescribeTableResponseBodyDataList) SetColumnName(v string) *DescribeTableResponseBodyDataList {
	s.ColumnName = &v
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

func (s *DescribeTableResponseBodyDataList) SetIndex(v string) *DescribeTableResponseBodyDataList {
	s.Index = &v
	return s
}

func (s *DescribeTableResponseBodyDataList) SetIsAllowNull(v string) *DescribeTableResponseBodyDataList {
	s.IsAllowNull = &v
	return s
}

func (s *DescribeTableResponseBodyDataList) SetIsPk(v string) *DescribeTableResponseBodyDataList {
	s.IsPk = &v
	return s
}

type DescribeTableResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeTableResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeTableResponse) SetStatusCode(v int32) *DescribeTableResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTableResponse) SetBody(v *DescribeTableResponseBody) *DescribeTableResponse {
	s.Body = v
	return s
}

type DescribeTableListByTypeRequest struct {
	CurrentPage    *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Query          *string `json:"Query,omitempty" xml:"Query,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TableType      *string `json:"TableType,omitempty" xml:"TableType,omitempty"`
}

func (s DescribeTableListByTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTableListByTypeRequest) GoString() string {
	return s.String()
}

func (s *DescribeTableListByTypeRequest) SetCurrentPage(v int32) *DescribeTableListByTypeRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeTableListByTypeRequest) SetDbName(v string) *DescribeTableListByTypeRequest {
	s.DbName = &v
	return s
}

func (s *DescribeTableListByTypeRequest) SetDrdsInstanceId(v string) *DescribeTableListByTypeRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeTableListByTypeRequest) SetPageSize(v int32) *DescribeTableListByTypeRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTableListByTypeRequest) SetQuery(v string) *DescribeTableListByTypeRequest {
	s.Query = &v
	return s
}

func (s *DescribeTableListByTypeRequest) SetRegionId(v string) *DescribeTableListByTypeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeTableListByTypeRequest) SetTableType(v string) *DescribeTableListByTypeRequest {
	s.TableType = &v
	return s
}

type DescribeTableListByTypeResponseBody struct {
	List       []*DescribeTableListByTypeResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	PageNumber *int32                                     `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
	Total      *int32                                     `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeTableListByTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTableListByTypeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTableListByTypeResponseBody) SetList(v []*DescribeTableListByTypeResponseBodyList) *DescribeTableListByTypeResponseBody {
	s.List = v
	return s
}

func (s *DescribeTableListByTypeResponseBody) SetPageNumber(v int32) *DescribeTableListByTypeResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeTableListByTypeResponseBody) SetPageSize(v int32) *DescribeTableListByTypeResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeTableListByTypeResponseBody) SetRequestId(v string) *DescribeTableListByTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTableListByTypeResponseBody) SetSuccess(v bool) *DescribeTableListByTypeResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeTableListByTypeResponseBody) SetTotal(v int32) *DescribeTableListByTypeResponseBody {
	s.Total = &v
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeTableListByTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeTableListByTypeResponse) SetStatusCode(v int32) *DescribeTableListByTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTableListByTypeResponse) SetBody(v *DescribeTableListByTypeResponseBody) *DescribeTableListByTypeResponse {
	s.Body = v
	return s
}

type DescribeTablesRequest struct {
	CurrentPage    *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Query          *string `json:"Query,omitempty" xml:"Query,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeTablesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTablesRequest) GoString() string {
	return s.String()
}

func (s *DescribeTablesRequest) SetCurrentPage(v int32) *DescribeTablesRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeTablesRequest) SetDbName(v string) *DescribeTablesRequest {
	s.DbName = &v
	return s
}

func (s *DescribeTablesRequest) SetDrdsInstanceId(v string) *DescribeTablesRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeTablesRequest) SetPageSize(v int32) *DescribeTablesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTablesRequest) SetQuery(v string) *DescribeTablesRequest {
	s.Query = &v
	return s
}

func (s *DescribeTablesRequest) SetRegionId(v string) *DescribeTablesRequest {
	s.RegionId = &v
	return s
}

type DescribeTablesResponseBody struct {
	List       []*DescribeTablesResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	PageNumber *int32                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
	Total      *int32                            `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeTablesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTablesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTablesResponseBody) SetList(v []*DescribeTablesResponseBodyList) *DescribeTablesResponseBody {
	s.List = v
	return s
}

func (s *DescribeTablesResponseBody) SetPageNumber(v int32) *DescribeTablesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeTablesResponseBody) SetPageSize(v int32) *DescribeTablesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeTablesResponseBody) SetRequestId(v string) *DescribeTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTablesResponseBody) SetSuccess(v bool) *DescribeTablesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeTablesResponseBody) SetTotal(v int32) *DescribeTablesResponseBody {
	s.Total = &v
	return s
}

type DescribeTablesResponseBodyList struct {
	AllowFullTableScan *bool   `json:"AllowFullTableScan,omitempty" xml:"AllowFullTableScan,omitempty"`
	Broadcast          *bool   `json:"Broadcast,omitempty" xml:"Broadcast,omitempty"`
	DbInstType         *int32  `json:"DbInstType,omitempty" xml:"DbInstType,omitempty"`
	IsLocked           *bool   `json:"IsLocked,omitempty" xml:"IsLocked,omitempty"`
	IsShard            *bool   `json:"IsShard,omitempty" xml:"IsShard,omitempty"`
	ShardKey           *string `json:"ShardKey,omitempty" xml:"ShardKey,omitempty"`
	Status             *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	Table              *string `json:"Table,omitempty" xml:"Table,omitempty"`
}

func (s DescribeTablesResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s DescribeTablesResponseBodyList) GoString() string {
	return s.String()
}

func (s *DescribeTablesResponseBodyList) SetAllowFullTableScan(v bool) *DescribeTablesResponseBodyList {
	s.AllowFullTableScan = &v
	return s
}

func (s *DescribeTablesResponseBodyList) SetBroadcast(v bool) *DescribeTablesResponseBodyList {
	s.Broadcast = &v
	return s
}

func (s *DescribeTablesResponseBodyList) SetDbInstType(v int32) *DescribeTablesResponseBodyList {
	s.DbInstType = &v
	return s
}

func (s *DescribeTablesResponseBodyList) SetIsLocked(v bool) *DescribeTablesResponseBodyList {
	s.IsLocked = &v
	return s
}

func (s *DescribeTablesResponseBodyList) SetIsShard(v bool) *DescribeTablesResponseBodyList {
	s.IsShard = &v
	return s
}

func (s *DescribeTablesResponseBodyList) SetShardKey(v string) *DescribeTablesResponseBodyList {
	s.ShardKey = &v
	return s
}

func (s *DescribeTablesResponseBodyList) SetStatus(v int32) *DescribeTablesResponseBodyList {
	s.Status = &v
	return s
}

func (s *DescribeTablesResponseBodyList) SetTable(v string) *DescribeTablesResponseBodyList {
	s.Table = &v
	return s
}

type DescribeTablesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeTablesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeTablesResponse) SetStatusCode(v int32) *DescribeTablesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTablesResponse) SetBody(v *DescribeTablesResponseBody) *DescribeTablesResponse {
	s.Body = v
	return s
}

type DisableSqlAuditRequest struct {
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s DisableSqlAuditRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableSqlAuditRequest) GoString() string {
	return s.String()
}

func (s *DisableSqlAuditRequest) SetDbName(v string) *DisableSqlAuditRequest {
	s.DbName = &v
	return s
}

func (s *DisableSqlAuditRequest) SetDrdsInstanceId(v string) *DisableSqlAuditRequest {
	s.DrdsInstanceId = &v
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DisableSqlAuditResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DisableSqlAuditResponse) SetStatusCode(v int32) *DisableSqlAuditResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableSqlAuditResponse) SetBody(v *DisableSqlAuditResponseBody) *DisableSqlAuditResponse {
	s.Body = v
	return s
}

type EnableInstanceIpv6AddressRequest struct {
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s EnableInstanceIpv6AddressRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableInstanceIpv6AddressRequest) GoString() string {
	return s.String()
}

func (s *EnableInstanceIpv6AddressRequest) SetDrdsInstanceId(v string) *EnableInstanceIpv6AddressRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *EnableInstanceIpv6AddressRequest) SetRegionId(v string) *EnableInstanceIpv6AddressRequest {
	s.RegionId = &v
	return s
}

type EnableInstanceIpv6AddressResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableInstanceIpv6AddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableInstanceIpv6AddressResponseBody) GoString() string {
	return s.String()
}

func (s *EnableInstanceIpv6AddressResponseBody) SetData(v bool) *EnableInstanceIpv6AddressResponseBody {
	s.Data = &v
	return s
}

func (s *EnableInstanceIpv6AddressResponseBody) SetRequestId(v string) *EnableInstanceIpv6AddressResponseBody {
	s.RequestId = &v
	return s
}

type EnableInstanceIpv6AddressResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EnableInstanceIpv6AddressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableInstanceIpv6AddressResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableInstanceIpv6AddressResponse) GoString() string {
	return s.String()
}

func (s *EnableInstanceIpv6AddressResponse) SetHeaders(v map[string]*string) *EnableInstanceIpv6AddressResponse {
	s.Headers = v
	return s
}

func (s *EnableInstanceIpv6AddressResponse) SetStatusCode(v int32) *EnableInstanceIpv6AddressResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableInstanceIpv6AddressResponse) SetBody(v *EnableInstanceIpv6AddressResponseBody) *EnableInstanceIpv6AddressResponse {
	s.Body = v
	return s
}

type EnableSqlAuditRequest struct {
	DbName               *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DrdsInstanceId       *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	IsRecall             *bool   `json:"IsRecall,omitempty" xml:"IsRecall,omitempty"`
	RecallEndTimestamp   *string `json:"RecallEndTimestamp,omitempty" xml:"RecallEndTimestamp,omitempty"`
	RecallStartTimestamp *string `json:"RecallStartTimestamp,omitempty" xml:"RecallStartTimestamp,omitempty"`
}

func (s EnableSqlAuditRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableSqlAuditRequest) GoString() string {
	return s.String()
}

func (s *EnableSqlAuditRequest) SetDbName(v string) *EnableSqlAuditRequest {
	s.DbName = &v
	return s
}

func (s *EnableSqlAuditRequest) SetDrdsInstanceId(v string) *EnableSqlAuditRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *EnableSqlAuditRequest) SetIsRecall(v bool) *EnableSqlAuditRequest {
	s.IsRecall = &v
	return s
}

func (s *EnableSqlAuditRequest) SetRecallEndTimestamp(v string) *EnableSqlAuditRequest {
	s.RecallEndTimestamp = &v
	return s
}

func (s *EnableSqlAuditRequest) SetRecallStartTimestamp(v string) *EnableSqlAuditRequest {
	s.RecallStartTimestamp = &v
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EnableSqlAuditResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *EnableSqlAuditResponse) SetStatusCode(v int32) *EnableSqlAuditResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableSqlAuditResponse) SetBody(v *EnableSqlAuditResponseBody) *EnableSqlAuditResponse {
	s.Body = v
	return s
}

type EnableSqlFlashbackMatchSwitchRequest struct {
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s EnableSqlFlashbackMatchSwitchRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableSqlFlashbackMatchSwitchRequest) GoString() string {
	return s.String()
}

func (s *EnableSqlFlashbackMatchSwitchRequest) SetDbName(v string) *EnableSqlFlashbackMatchSwitchRequest {
	s.DbName = &v
	return s
}

func (s *EnableSqlFlashbackMatchSwitchRequest) SetDrdsInstanceId(v string) *EnableSqlFlashbackMatchSwitchRequest {
	s.DrdsInstanceId = &v
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
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EnableSqlFlashbackMatchSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *EnableSqlFlashbackMatchSwitchResponse) SetStatusCode(v int32) *EnableSqlFlashbackMatchSwitchResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableSqlFlashbackMatchSwitchResponse) SetBody(v *EnableSqlFlashbackMatchSwitchResponseBody) *EnableSqlFlashbackMatchSwitchResponse {
	s.Body = v
	return s
}

type FlashbackRecycleBinTableRequest struct {
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TableName      *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s FlashbackRecycleBinTableRequest) String() string {
	return tea.Prettify(s)
}

func (s FlashbackRecycleBinTableRequest) GoString() string {
	return s.String()
}

func (s *FlashbackRecycleBinTableRequest) SetDbName(v string) *FlashbackRecycleBinTableRequest {
	s.DbName = &v
	return s
}

func (s *FlashbackRecycleBinTableRequest) SetDrdsInstanceId(v string) *FlashbackRecycleBinTableRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *FlashbackRecycleBinTableRequest) SetRegionId(v string) *FlashbackRecycleBinTableRequest {
	s.RegionId = &v
	return s
}

func (s *FlashbackRecycleBinTableRequest) SetTableName(v string) *FlashbackRecycleBinTableRequest {
	s.TableName = &v
	return s
}

type FlashbackRecycleBinTableResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s FlashbackRecycleBinTableResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FlashbackRecycleBinTableResponseBody) GoString() string {
	return s.String()
}

func (s *FlashbackRecycleBinTableResponseBody) SetData(v bool) *FlashbackRecycleBinTableResponseBody {
	s.Data = &v
	return s
}

func (s *FlashbackRecycleBinTableResponseBody) SetRequestId(v string) *FlashbackRecycleBinTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *FlashbackRecycleBinTableResponseBody) SetSuccess(v bool) *FlashbackRecycleBinTableResponseBody {
	s.Success = &v
	return s
}

type FlashbackRecycleBinTableResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *FlashbackRecycleBinTableResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FlashbackRecycleBinTableResponse) String() string {
	return tea.Prettify(s)
}

func (s FlashbackRecycleBinTableResponse) GoString() string {
	return s.String()
}

func (s *FlashbackRecycleBinTableResponse) SetHeaders(v map[string]*string) *FlashbackRecycleBinTableResponse {
	s.Headers = v
	return s
}

func (s *FlashbackRecycleBinTableResponse) SetStatusCode(v int32) *FlashbackRecycleBinTableResponse {
	s.StatusCode = &v
	return s
}

func (s *FlashbackRecycleBinTableResponse) SetBody(v *FlashbackRecycleBinTableResponseBody) *FlashbackRecycleBinTableResponse {
	s.Body = v
	return s
}

type GetDrdsDbRdsRelationInfoRequest struct {
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s GetDrdsDbRdsRelationInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDrdsDbRdsRelationInfoRequest) GoString() string {
	return s.String()
}

func (s *GetDrdsDbRdsRelationInfoRequest) SetDbName(v string) *GetDrdsDbRdsRelationInfoRequest {
	s.DbName = &v
	return s
}

func (s *GetDrdsDbRdsRelationInfoRequest) SetDrdsInstanceId(v string) *GetDrdsDbRdsRelationInfoRequest {
	s.DrdsInstanceId = &v
	return s
}

type GetDrdsDbRdsRelationInfoResponseBody struct {
	Data      []*GetDrdsDbRdsRelationInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDrdsDbRdsRelationInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDrdsDbRdsRelationInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetDrdsDbRdsRelationInfoResponseBody) SetData(v []*GetDrdsDbRdsRelationInfoResponseBodyData) *GetDrdsDbRdsRelationInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetDrdsDbRdsRelationInfoResponseBody) SetRequestId(v string) *GetDrdsDbRdsRelationInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDrdsDbRdsRelationInfoResponseBody) SetSuccess(v bool) *GetDrdsDbRdsRelationInfoResponseBody {
	s.Success = &v
	return s
}

type GetDrdsDbRdsRelationInfoResponseBodyData struct {
	RdsInstanceId        *string   `json:"RdsInstanceId,omitempty" xml:"RdsInstanceId,omitempty"`
	ReadOnlyInstanceInfo []*string `json:"ReadOnlyInstanceInfo,omitempty" xml:"ReadOnlyInstanceInfo,omitempty" type:"Repeated"`
	UsedInstanceId       *string   `json:"UsedInstanceId,omitempty" xml:"UsedInstanceId,omitempty"`
	UsedInstanceType     *string   `json:"UsedInstanceType,omitempty" xml:"UsedInstanceType,omitempty"`
}

func (s GetDrdsDbRdsRelationInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetDrdsDbRdsRelationInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDrdsDbRdsRelationInfoResponseBodyData) SetRdsInstanceId(v string) *GetDrdsDbRdsRelationInfoResponseBodyData {
	s.RdsInstanceId = &v
	return s
}

func (s *GetDrdsDbRdsRelationInfoResponseBodyData) SetReadOnlyInstanceInfo(v []*string) *GetDrdsDbRdsRelationInfoResponseBodyData {
	s.ReadOnlyInstanceInfo = v
	return s
}

func (s *GetDrdsDbRdsRelationInfoResponseBodyData) SetUsedInstanceId(v string) *GetDrdsDbRdsRelationInfoResponseBodyData {
	s.UsedInstanceId = &v
	return s
}

func (s *GetDrdsDbRdsRelationInfoResponseBodyData) SetUsedInstanceType(v string) *GetDrdsDbRdsRelationInfoResponseBodyData {
	s.UsedInstanceType = &v
	return s
}

type GetDrdsDbRdsRelationInfoResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDrdsDbRdsRelationInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDrdsDbRdsRelationInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDrdsDbRdsRelationInfoResponse) GoString() string {
	return s.String()
}

func (s *GetDrdsDbRdsRelationInfoResponse) SetHeaders(v map[string]*string) *GetDrdsDbRdsRelationInfoResponse {
	s.Headers = v
	return s
}

func (s *GetDrdsDbRdsRelationInfoResponse) SetStatusCode(v int32) *GetDrdsDbRdsRelationInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDrdsDbRdsRelationInfoResponse) SetBody(v *GetDrdsDbRdsRelationInfoResponseBody) *GetDrdsDbRdsRelationInfoResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	NextToken    *string                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId     *string                       `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId   []*string                     `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceType *string                       `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tag          []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetRegionId(v string) *ListTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceId(v []*string) *ListTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetTag(v []*ListTagResourcesRequestTag) *ListTagResourcesRequest {
	s.Tag = v
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
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKey       *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue     *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceId(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceType(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagKey(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagKey = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagValue(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagValue = &v
	return s
}

type ListTagResourcesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListTagResourcesResponse) SetStatusCode(v int32) *ListTagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagResourcesResponse) SetBody(v *ListTagResourcesResponseBody) *ListTagResourcesResponse {
	s.Body = v
	return s
}

type ManagePrivateRdsRequest struct {
	DBInstanceId   *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	Params         *string `json:"Params,omitempty" xml:"Params,omitempty"`
	RdsAction      *string `json:"RdsAction,omitempty" xml:"RdsAction,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ManagePrivateRdsRequest) String() string {
	return tea.Prettify(s)
}

func (s ManagePrivateRdsRequest) GoString() string {
	return s.String()
}

func (s *ManagePrivateRdsRequest) SetDBInstanceId(v string) *ManagePrivateRdsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ManagePrivateRdsRequest) SetDrdsInstanceId(v string) *ManagePrivateRdsRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *ManagePrivateRdsRequest) SetParams(v string) *ManagePrivateRdsRequest {
	s.Params = &v
	return s
}

func (s *ManagePrivateRdsRequest) SetRdsAction(v string) *ManagePrivateRdsRequest {
	s.RdsAction = &v
	return s
}

func (s *ManagePrivateRdsRequest) SetRegionId(v string) *ManagePrivateRdsRequest {
	s.RegionId = &v
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ManagePrivateRdsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ManagePrivateRdsResponse) SetStatusCode(v int32) *ManagePrivateRdsResponse {
	s.StatusCode = &v
	return s
}

func (s *ManagePrivateRdsResponse) SetBody(v *ManagePrivateRdsResponseBody) *ManagePrivateRdsResponse {
	s.Body = v
	return s
}

type ModifyAccountDescriptionRequest struct {
	AccountName    *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s ModifyAccountDescriptionRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccountDescriptionRequest) GoString() string {
	return s.String()
}

func (s *ModifyAccountDescriptionRequest) SetAccountName(v string) *ModifyAccountDescriptionRequest {
	s.AccountName = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetDescription(v string) *ModifyAccountDescriptionRequest {
	s.Description = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetDrdsInstanceId(v string) *ModifyAccountDescriptionRequest {
	s.DrdsInstanceId = &v
	return s
}

type ModifyAccountDescriptionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyAccountDescriptionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccountDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAccountDescriptionResponseBody) SetRequestId(v string) *ModifyAccountDescriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAccountDescriptionResponseBody) SetSuccess(v bool) *ModifyAccountDescriptionResponseBody {
	s.Success = &v
	return s
}

type ModifyAccountDescriptionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyAccountDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyAccountDescriptionResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccountDescriptionResponse) GoString() string {
	return s.String()
}

func (s *ModifyAccountDescriptionResponse) SetHeaders(v map[string]*string) *ModifyAccountDescriptionResponse {
	s.Headers = v
	return s
}

func (s *ModifyAccountDescriptionResponse) SetStatusCode(v int32) *ModifyAccountDescriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAccountDescriptionResponse) SetBody(v *ModifyAccountDescriptionResponseBody) *ModifyAccountDescriptionResponse {
	s.Body = v
	return s
}

type ModifyAccountPrivilegeRequest struct {
	AccountName    *string                                     `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	DbPrivilege    []*ModifyAccountPrivilegeRequestDbPrivilege `json:"DbPrivilege,omitempty" xml:"DbPrivilege,omitempty" type:"Repeated"`
	DrdsInstanceId *string                                     `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	RegionId       *string                                     `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyAccountPrivilegeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccountPrivilegeRequest) GoString() string {
	return s.String()
}

func (s *ModifyAccountPrivilegeRequest) SetAccountName(v string) *ModifyAccountPrivilegeRequest {
	s.AccountName = &v
	return s
}

func (s *ModifyAccountPrivilegeRequest) SetDbPrivilege(v []*ModifyAccountPrivilegeRequestDbPrivilege) *ModifyAccountPrivilegeRequest {
	s.DbPrivilege = v
	return s
}

func (s *ModifyAccountPrivilegeRequest) SetDrdsInstanceId(v string) *ModifyAccountPrivilegeRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *ModifyAccountPrivilegeRequest) SetRegionId(v string) *ModifyAccountPrivilegeRequest {
	s.RegionId = &v
	return s
}

type ModifyAccountPrivilegeRequestDbPrivilege struct {
	DbName    *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	Privilege *string `json:"Privilege,omitempty" xml:"Privilege,omitempty"`
}

func (s ModifyAccountPrivilegeRequestDbPrivilege) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccountPrivilegeRequestDbPrivilege) GoString() string {
	return s.String()
}

func (s *ModifyAccountPrivilegeRequestDbPrivilege) SetDbName(v string) *ModifyAccountPrivilegeRequestDbPrivilege {
	s.DbName = &v
	return s
}

func (s *ModifyAccountPrivilegeRequestDbPrivilege) SetPrivilege(v string) *ModifyAccountPrivilegeRequestDbPrivilege {
	s.Privilege = &v
	return s
}

type ModifyAccountPrivilegeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyAccountPrivilegeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccountPrivilegeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAccountPrivilegeResponseBody) SetRequestId(v string) *ModifyAccountPrivilegeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAccountPrivilegeResponseBody) SetSuccess(v bool) *ModifyAccountPrivilegeResponseBody {
	s.Success = &v
	return s
}

type ModifyAccountPrivilegeResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyAccountPrivilegeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyAccountPrivilegeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccountPrivilegeResponse) GoString() string {
	return s.String()
}

func (s *ModifyAccountPrivilegeResponse) SetHeaders(v map[string]*string) *ModifyAccountPrivilegeResponse {
	s.Headers = v
	return s
}

func (s *ModifyAccountPrivilegeResponse) SetStatusCode(v int32) *ModifyAccountPrivilegeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAccountPrivilegeResponse) SetBody(v *ModifyAccountPrivilegeResponseBody) *ModifyAccountPrivilegeResponse {
	s.Body = v
	return s
}

type ModifyDrdsInstanceDescriptionRequest struct {
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s ModifyDrdsInstanceDescriptionRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDrdsInstanceDescriptionRequest) GoString() string {
	return s.String()
}

func (s *ModifyDrdsInstanceDescriptionRequest) SetDescription(v string) *ModifyDrdsInstanceDescriptionRequest {
	s.Description = &v
	return s
}

func (s *ModifyDrdsInstanceDescriptionRequest) SetDrdsInstanceId(v string) *ModifyDrdsInstanceDescriptionRequest {
	s.DrdsInstanceId = &v
	return s
}

type ModifyDrdsInstanceDescriptionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyDrdsInstanceDescriptionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDrdsInstanceDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDrdsInstanceDescriptionResponseBody) SetRequestId(v string) *ModifyDrdsInstanceDescriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDrdsInstanceDescriptionResponseBody) SetSuccess(v bool) *ModifyDrdsInstanceDescriptionResponseBody {
	s.Success = &v
	return s
}

type ModifyDrdsInstanceDescriptionResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyDrdsInstanceDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifyDrdsInstanceDescriptionResponse) SetStatusCode(v int32) *ModifyDrdsInstanceDescriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDrdsInstanceDescriptionResponse) SetBody(v *ModifyDrdsInstanceDescriptionResponseBody) *ModifyDrdsInstanceDescriptionResponse {
	s.Body = v
	return s
}

type ModifyDrdsIpWhiteListRequest struct {
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	GroupAttribute *string `json:"GroupAttribute,omitempty" xml:"GroupAttribute,omitempty"`
	GroupName      *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	IpWhiteList    *string `json:"IpWhiteList,omitempty" xml:"IpWhiteList,omitempty"`
	Mode           *bool   `json:"Mode,omitempty" xml:"Mode,omitempty"`
}

func (s ModifyDrdsIpWhiteListRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDrdsIpWhiteListRequest) GoString() string {
	return s.String()
}

func (s *ModifyDrdsIpWhiteListRequest) SetDbName(v string) *ModifyDrdsIpWhiteListRequest {
	s.DbName = &v
	return s
}

func (s *ModifyDrdsIpWhiteListRequest) SetDrdsInstanceId(v string) *ModifyDrdsIpWhiteListRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *ModifyDrdsIpWhiteListRequest) SetGroupAttribute(v string) *ModifyDrdsIpWhiteListRequest {
	s.GroupAttribute = &v
	return s
}

func (s *ModifyDrdsIpWhiteListRequest) SetGroupName(v string) *ModifyDrdsIpWhiteListRequest {
	s.GroupName = &v
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

type ModifyDrdsIpWhiteListResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyDrdsIpWhiteListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDrdsIpWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDrdsIpWhiteListResponseBody) SetRequestId(v string) *ModifyDrdsIpWhiteListResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDrdsIpWhiteListResponseBody) SetSuccess(v bool) *ModifyDrdsIpWhiteListResponseBody {
	s.Success = &v
	return s
}

type ModifyDrdsIpWhiteListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyDrdsIpWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifyDrdsIpWhiteListResponse) SetStatusCode(v int32) *ModifyDrdsIpWhiteListResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDrdsIpWhiteListResponse) SetBody(v *ModifyDrdsIpWhiteListResponseBody) *ModifyDrdsIpWhiteListResponse {
	s.Body = v
	return s
}

type ModifyPolarDbReadWeightRequest struct {
	DbInstanceId   *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DbNodeIds      *string `json:"DbNodeIds,omitempty" xml:"DbNodeIds,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	Weights        *string `json:"Weights,omitempty" xml:"Weights,omitempty"`
}

func (s ModifyPolarDbReadWeightRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyPolarDbReadWeightRequest) GoString() string {
	return s.String()
}

func (s *ModifyPolarDbReadWeightRequest) SetDbInstanceId(v string) *ModifyPolarDbReadWeightRequest {
	s.DbInstanceId = &v
	return s
}

func (s *ModifyPolarDbReadWeightRequest) SetDbName(v string) *ModifyPolarDbReadWeightRequest {
	s.DbName = &v
	return s
}

func (s *ModifyPolarDbReadWeightRequest) SetDbNodeIds(v string) *ModifyPolarDbReadWeightRequest {
	s.DbNodeIds = &v
	return s
}

func (s *ModifyPolarDbReadWeightRequest) SetDrdsInstanceId(v string) *ModifyPolarDbReadWeightRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *ModifyPolarDbReadWeightRequest) SetWeights(v string) *ModifyPolarDbReadWeightRequest {
	s.Weights = &v
	return s
}

type ModifyPolarDbReadWeightResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyPolarDbReadWeightResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyPolarDbReadWeightResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyPolarDbReadWeightResponseBody) SetRequestId(v string) *ModifyPolarDbReadWeightResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyPolarDbReadWeightResponseBody) SetSuccess(v bool) *ModifyPolarDbReadWeightResponseBody {
	s.Success = &v
	return s
}

type ModifyPolarDbReadWeightResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyPolarDbReadWeightResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyPolarDbReadWeightResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyPolarDbReadWeightResponse) GoString() string {
	return s.String()
}

func (s *ModifyPolarDbReadWeightResponse) SetHeaders(v map[string]*string) *ModifyPolarDbReadWeightResponse {
	s.Headers = v
	return s
}

func (s *ModifyPolarDbReadWeightResponse) SetStatusCode(v int32) *ModifyPolarDbReadWeightResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyPolarDbReadWeightResponse) SetBody(v *ModifyPolarDbReadWeightResponseBody) *ModifyPolarDbReadWeightResponse {
	s.Body = v
	return s
}

type ModifyRdsReadWeightRequest struct {
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	InstanceNames  *string `json:"InstanceNames,omitempty" xml:"InstanceNames,omitempty"`
	Weights        *string `json:"Weights,omitempty" xml:"Weights,omitempty"`
}

func (s ModifyRdsReadWeightRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyRdsReadWeightRequest) GoString() string {
	return s.String()
}

func (s *ModifyRdsReadWeightRequest) SetDbName(v string) *ModifyRdsReadWeightRequest {
	s.DbName = &v
	return s
}

func (s *ModifyRdsReadWeightRequest) SetDrdsInstanceId(v string) *ModifyRdsReadWeightRequest {
	s.DrdsInstanceId = &v
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
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyRdsReadWeightResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyRdsReadWeightResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyRdsReadWeightResponseBody) SetRequestId(v string) *ModifyRdsReadWeightResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyRdsReadWeightResponseBody) SetSuccess(v bool) *ModifyRdsReadWeightResponseBody {
	s.Success = &v
	return s
}

type ModifyRdsReadWeightResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyRdsReadWeightResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifyRdsReadWeightResponse) SetStatusCode(v int32) *ModifyRdsReadWeightResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyRdsReadWeightResponse) SetBody(v *ModifyRdsReadWeightResponseBody) *ModifyRdsReadWeightResponse {
	s.Body = v
	return s
}

type PutStartBackupRequest struct {
	BackupDbNames  *string `json:"BackupDbNames,omitempty" xml:"BackupDbNames,omitempty"`
	BackupLevel    *string `json:"BackupLevel,omitempty" xml:"BackupLevel,omitempty"`
	BackupMode     *string `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s PutStartBackupRequest) String() string {
	return tea.Prettify(s)
}

func (s PutStartBackupRequest) GoString() string {
	return s.String()
}

func (s *PutStartBackupRequest) SetBackupDbNames(v string) *PutStartBackupRequest {
	s.BackupDbNames = &v
	return s
}

func (s *PutStartBackupRequest) SetBackupLevel(v string) *PutStartBackupRequest {
	s.BackupLevel = &v
	return s
}

func (s *PutStartBackupRequest) SetBackupMode(v string) *PutStartBackupRequest {
	s.BackupMode = &v
	return s
}

func (s *PutStartBackupRequest) SetDrdsInstanceId(v string) *PutStartBackupRequest {
	s.DrdsInstanceId = &v
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PutStartBackupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *PutStartBackupResponse) SetStatusCode(v int32) *PutStartBackupResponse {
	s.StatusCode = &v
	return s
}

func (s *PutStartBackupResponse) SetBody(v *PutStartBackupResponseBody) *PutStartBackupResponse {
	s.Body = v
	return s
}

type RefreshDrdsAtomUrlRequest struct {
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s RefreshDrdsAtomUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s RefreshDrdsAtomUrlRequest) GoString() string {
	return s.String()
}

func (s *RefreshDrdsAtomUrlRequest) SetDbName(v string) *RefreshDrdsAtomUrlRequest {
	s.DbName = &v
	return s
}

func (s *RefreshDrdsAtomUrlRequest) SetDrdsInstanceId(v string) *RefreshDrdsAtomUrlRequest {
	s.DrdsInstanceId = &v
	return s
}

type RefreshDrdsAtomUrlResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RefreshDrdsAtomUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RefreshDrdsAtomUrlResponseBody) GoString() string {
	return s.String()
}

func (s *RefreshDrdsAtomUrlResponseBody) SetRequestId(v string) *RefreshDrdsAtomUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *RefreshDrdsAtomUrlResponseBody) SetResult(v bool) *RefreshDrdsAtomUrlResponseBody {
	s.Result = &v
	return s
}

func (s *RefreshDrdsAtomUrlResponseBody) SetSuccess(v bool) *RefreshDrdsAtomUrlResponseBody {
	s.Success = &v
	return s
}

type RefreshDrdsAtomUrlResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RefreshDrdsAtomUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RefreshDrdsAtomUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s RefreshDrdsAtomUrlResponse) GoString() string {
	return s.String()
}

func (s *RefreshDrdsAtomUrlResponse) SetHeaders(v map[string]*string) *RefreshDrdsAtomUrlResponse {
	s.Headers = v
	return s
}

func (s *RefreshDrdsAtomUrlResponse) SetStatusCode(v int32) *RefreshDrdsAtomUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *RefreshDrdsAtomUrlResponse) SetBody(v *RefreshDrdsAtomUrlResponseBody) *RefreshDrdsAtomUrlResponse {
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
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ReleaseInstanceInternetAddressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ReleaseInstanceInternetAddressResponse) SetStatusCode(v int32) *ReleaseInstanceInternetAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseInstanceInternetAddressResponse) SetBody(v *ReleaseInstanceInternetAddressResponseBody) *ReleaseInstanceInternetAddressResponse {
	s.Body = v
	return s
}

type RemoveBackupsSetRequest struct {
	BackupId       *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s RemoveBackupsSetRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveBackupsSetRequest) GoString() string {
	return s.String()
}

func (s *RemoveBackupsSetRequest) SetBackupId(v string) *RemoveBackupsSetRequest {
	s.BackupId = &v
	return s
}

func (s *RemoveBackupsSetRequest) SetDrdsInstanceId(v string) *RemoveBackupsSetRequest {
	s.DrdsInstanceId = &v
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RemoveBackupsSetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RemoveBackupsSetResponse) SetStatusCode(v int32) *RemoveBackupsSetResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveBackupsSetResponse) SetBody(v *RemoveBackupsSetResponseBody) *RemoveBackupsSetResponse {
	s.Body = v
	return s
}

type RemoveDrdsDbRequest struct {
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s RemoveDrdsDbRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveDrdsDbRequest) GoString() string {
	return s.String()
}

func (s *RemoveDrdsDbRequest) SetDbName(v string) *RemoveDrdsDbRequest {
	s.DbName = &v
	return s
}

func (s *RemoveDrdsDbRequest) SetDrdsInstanceId(v string) *RemoveDrdsDbRequest {
	s.DrdsInstanceId = &v
	return s
}

type RemoveDrdsDbResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RemoveDrdsDbResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveDrdsDbResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveDrdsDbResponseBody) SetRequestId(v string) *RemoveDrdsDbResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveDrdsDbResponseBody) SetSuccess(v bool) *RemoveDrdsDbResponseBody {
	s.Success = &v
	return s
}

type RemoveDrdsDbResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RemoveDrdsDbResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RemoveDrdsDbResponse) SetStatusCode(v int32) *RemoveDrdsDbResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveDrdsDbResponse) SetBody(v *RemoveDrdsDbResponseBody) *RemoveDrdsDbResponse {
	s.Body = v
	return s
}

type RemoveDrdsDbFailedRecordRequest struct {
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s RemoveDrdsDbFailedRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveDrdsDbFailedRecordRequest) GoString() string {
	return s.String()
}

func (s *RemoveDrdsDbFailedRecordRequest) SetDbName(v string) *RemoveDrdsDbFailedRecordRequest {
	s.DbName = &v
	return s
}

func (s *RemoveDrdsDbFailedRecordRequest) SetDrdsInstanceId(v string) *RemoveDrdsDbFailedRecordRequest {
	s.DrdsInstanceId = &v
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RemoveDrdsDbFailedRecordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RemoveDrdsDbFailedRecordResponse) SetStatusCode(v int32) *RemoveDrdsDbFailedRecordResponse {
	s.StatusCode = &v
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
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RemoveDrdsInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveDrdsInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveDrdsInstanceResponseBody) SetRequestId(v string) *RemoveDrdsInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveDrdsInstanceResponseBody) SetSuccess(v bool) *RemoveDrdsInstanceResponseBody {
	s.Success = &v
	return s
}

type RemoveDrdsInstanceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RemoveDrdsInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RemoveDrdsInstanceResponse) SetStatusCode(v int32) *RemoveDrdsInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveDrdsInstanceResponse) SetBody(v *RemoveDrdsInstanceResponseBody) *RemoveDrdsInstanceResponse {
	s.Body = v
	return s
}

type RemoveInstanceAccountRequest struct {
	AccountName    *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s RemoveInstanceAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveInstanceAccountRequest) GoString() string {
	return s.String()
}

func (s *RemoveInstanceAccountRequest) SetAccountName(v string) *RemoveInstanceAccountRequest {
	s.AccountName = &v
	return s
}

func (s *RemoveInstanceAccountRequest) SetDrdsInstanceId(v string) *RemoveInstanceAccountRequest {
	s.DrdsInstanceId = &v
	return s
}

type RemoveInstanceAccountResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RemoveInstanceAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveInstanceAccountResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveInstanceAccountResponseBody) SetRequestId(v string) *RemoveInstanceAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveInstanceAccountResponseBody) SetSuccess(v bool) *RemoveInstanceAccountResponseBody {
	s.Success = &v
	return s
}

type RemoveInstanceAccountResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RemoveInstanceAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RemoveInstanceAccountResponse) SetStatusCode(v int32) *RemoveInstanceAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveInstanceAccountResponse) SetBody(v *RemoveInstanceAccountResponseBody) *RemoveInstanceAccountResponse {
	s.Body = v
	return s
}

type RemoveRecycleBinTableRequest struct {
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TableName      *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s RemoveRecycleBinTableRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveRecycleBinTableRequest) GoString() string {
	return s.String()
}

func (s *RemoveRecycleBinTableRequest) SetDbName(v string) *RemoveRecycleBinTableRequest {
	s.DbName = &v
	return s
}

func (s *RemoveRecycleBinTableRequest) SetDrdsInstanceId(v string) *RemoveRecycleBinTableRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *RemoveRecycleBinTableRequest) SetRegionId(v string) *RemoveRecycleBinTableRequest {
	s.RegionId = &v
	return s
}

func (s *RemoveRecycleBinTableRequest) SetTableName(v string) *RemoveRecycleBinTableRequest {
	s.TableName = &v
	return s
}

type RemoveRecycleBinTableResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RemoveRecycleBinTableResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveRecycleBinTableResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveRecycleBinTableResponseBody) SetData(v bool) *RemoveRecycleBinTableResponseBody {
	s.Data = &v
	return s
}

func (s *RemoveRecycleBinTableResponseBody) SetRequestId(v string) *RemoveRecycleBinTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveRecycleBinTableResponseBody) SetSuccess(v bool) *RemoveRecycleBinTableResponseBody {
	s.Success = &v
	return s
}

type RemoveRecycleBinTableResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RemoveRecycleBinTableResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveRecycleBinTableResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveRecycleBinTableResponse) GoString() string {
	return s.String()
}

func (s *RemoveRecycleBinTableResponse) SetHeaders(v map[string]*string) *RemoveRecycleBinTableResponse {
	s.Headers = v
	return s
}

func (s *RemoveRecycleBinTableResponse) SetStatusCode(v int32) *RemoveRecycleBinTableResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveRecycleBinTableResponse) SetBody(v *RemoveRecycleBinTableResponseBody) *RemoveRecycleBinTableResponse {
	s.Body = v
	return s
}

type RestartDrdsInstanceRequest struct {
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s RestartDrdsInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RestartDrdsInstanceRequest) GoString() string {
	return s.String()
}

func (s *RestartDrdsInstanceRequest) SetDrdsInstanceId(v string) *RestartDrdsInstanceRequest {
	s.DrdsInstanceId = &v
	return s
}

type RestartDrdsInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId    *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s RestartDrdsInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RestartDrdsInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RestartDrdsInstanceResponseBody) SetRequestId(v string) *RestartDrdsInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestartDrdsInstanceResponseBody) SetSuccess(v bool) *RestartDrdsInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *RestartDrdsInstanceResponseBody) SetTaskId(v int64) *RestartDrdsInstanceResponseBody {
	s.TaskId = &v
	return s
}

type RestartDrdsInstanceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RestartDrdsInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RestartDrdsInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s RestartDrdsInstanceResponse) GoString() string {
	return s.String()
}

func (s *RestartDrdsInstanceResponse) SetHeaders(v map[string]*string) *RestartDrdsInstanceResponse {
	s.Headers = v
	return s
}

func (s *RestartDrdsInstanceResponse) SetStatusCode(v int32) *RestartDrdsInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RestartDrdsInstanceResponse) SetBody(v *RestartDrdsInstanceResponseBody) *RestartDrdsInstanceResponse {
	s.Body = v
	return s
}

type RollbackInstanceVersionRequest struct {
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RollbackInstanceVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s RollbackInstanceVersionRequest) GoString() string {
	return s.String()
}

func (s *RollbackInstanceVersionRequest) SetDrdsInstanceId(v string) *RollbackInstanceVersionRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *RollbackInstanceVersionRequest) SetRegionId(v string) *RollbackInstanceVersionRequest {
	s.RegionId = &v
	return s
}

type RollbackInstanceVersionResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RollbackInstanceVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RollbackInstanceVersionResponseBody) GoString() string {
	return s.String()
}

func (s *RollbackInstanceVersionResponseBody) SetData(v string) *RollbackInstanceVersionResponseBody {
	s.Data = &v
	return s
}

func (s *RollbackInstanceVersionResponseBody) SetRequestId(v string) *RollbackInstanceVersionResponseBody {
	s.RequestId = &v
	return s
}

type RollbackInstanceVersionResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RollbackInstanceVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RollbackInstanceVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s RollbackInstanceVersionResponse) GoString() string {
	return s.String()
}

func (s *RollbackInstanceVersionResponse) SetHeaders(v map[string]*string) *RollbackInstanceVersionResponse {
	s.Headers = v
	return s
}

func (s *RollbackInstanceVersionResponse) SetStatusCode(v int32) *RollbackInstanceVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *RollbackInstanceVersionResponse) SetBody(v *RollbackInstanceVersionResponseBody) *RollbackInstanceVersionResponse {
	s.Body = v
	return s
}

type SetBackupLocalRequest struct {
	DrdsInstanceId           *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	HighSpaceUsageProtection *string `json:"HighSpaceUsageProtection,omitempty" xml:"HighSpaceUsageProtection,omitempty"`
	LocalLogRetentionHours   *string `json:"LocalLogRetentionHours,omitempty" xml:"LocalLogRetentionHours,omitempty"`
	LocalLogRetentionSpace   *string `json:"LocalLogRetentionSpace,omitempty" xml:"LocalLogRetentionSpace,omitempty"`
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

func (s *SetBackupLocalRequest) SetHighSpaceUsageProtection(v string) *SetBackupLocalRequest {
	s.HighSpaceUsageProtection = &v
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetBackupLocalResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SetBackupLocalResponse) SetStatusCode(v int32) *SetBackupLocalResponse {
	s.StatusCode = &v
	return s
}

func (s *SetBackupLocalResponse) SetBody(v *SetBackupLocalResponseBody) *SetBackupLocalResponse {
	s.Body = v
	return s
}

type SetBackupPolicyRequest struct {
	BackupDbNames             *string `json:"BackupDbNames,omitempty" xml:"BackupDbNames,omitempty"`
	BackupLevel               *string `json:"BackupLevel,omitempty" xml:"BackupLevel,omitempty"`
	BackupLog                 *string `json:"BackupLog,omitempty" xml:"BackupLog,omitempty"`
	BackupMode                *string `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	DataBackupRetentionPeriod *string `json:"DataBackupRetentionPeriod,omitempty" xml:"DataBackupRetentionPeriod,omitempty"`
	DrdsInstanceId            *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	LogBackupRetentionPeriod  *string `json:"LogBackupRetentionPeriod,omitempty" xml:"LogBackupRetentionPeriod,omitempty"`
	PreferredBackupEndTime    *string `json:"PreferredBackupEndTime,omitempty" xml:"PreferredBackupEndTime,omitempty"`
	PreferredBackupPeriod     *string `json:"PreferredBackupPeriod,omitempty" xml:"PreferredBackupPeriod,omitempty"`
	PreferredBackupStartTime  *string `json:"PreferredBackupStartTime,omitempty" xml:"PreferredBackupStartTime,omitempty"`
}

func (s SetBackupPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s SetBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *SetBackupPolicyRequest) SetBackupDbNames(v string) *SetBackupPolicyRequest {
	s.BackupDbNames = &v
	return s
}

func (s *SetBackupPolicyRequest) SetBackupLevel(v string) *SetBackupPolicyRequest {
	s.BackupLevel = &v
	return s
}

func (s *SetBackupPolicyRequest) SetBackupLog(v string) *SetBackupPolicyRequest {
	s.BackupLog = &v
	return s
}

func (s *SetBackupPolicyRequest) SetBackupMode(v string) *SetBackupPolicyRequest {
	s.BackupMode = &v
	return s
}

func (s *SetBackupPolicyRequest) SetDataBackupRetentionPeriod(v string) *SetBackupPolicyRequest {
	s.DataBackupRetentionPeriod = &v
	return s
}

func (s *SetBackupPolicyRequest) SetDrdsInstanceId(v string) *SetBackupPolicyRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *SetBackupPolicyRequest) SetLogBackupRetentionPeriod(v string) *SetBackupPolicyRequest {
	s.LogBackupRetentionPeriod = &v
	return s
}

func (s *SetBackupPolicyRequest) SetPreferredBackupEndTime(v string) *SetBackupPolicyRequest {
	s.PreferredBackupEndTime = &v
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetBackupPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SetBackupPolicyResponse) SetStatusCode(v int32) *SetBackupPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *SetBackupPolicyResponse) SetBody(v *SetBackupPolicyResponseBody) *SetBackupPolicyResponse {
	s.Body = v
	return s
}

type SetupBroadcastTablesRequest struct {
	Active         *bool     `json:"Active,omitempty" xml:"Active,omitempty"`
	DbName         *string   `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DrdsInstanceId *string   `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	RegionId       *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TableName      []*string `json:"TableName,omitempty" xml:"TableName,omitempty" type:"Repeated"`
}

func (s SetupBroadcastTablesRequest) String() string {
	return tea.Prettify(s)
}

func (s SetupBroadcastTablesRequest) GoString() string {
	return s.String()
}

func (s *SetupBroadcastTablesRequest) SetActive(v bool) *SetupBroadcastTablesRequest {
	s.Active = &v
	return s
}

func (s *SetupBroadcastTablesRequest) SetDbName(v string) *SetupBroadcastTablesRequest {
	s.DbName = &v
	return s
}

func (s *SetupBroadcastTablesRequest) SetDrdsInstanceId(v string) *SetupBroadcastTablesRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *SetupBroadcastTablesRequest) SetRegionId(v string) *SetupBroadcastTablesRequest {
	s.RegionId = &v
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetupBroadcastTablesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SetupBroadcastTablesResponse) SetStatusCode(v int32) *SetupBroadcastTablesResponse {
	s.StatusCode = &v
	return s
}

func (s *SetupBroadcastTablesResponse) SetBody(v *SetupBroadcastTablesResponseBody) *SetupBroadcastTablesResponse {
	s.Body = v
	return s
}

type SetupDrdsParamsRequest struct {
	Data           []*SetupDrdsParamsRequestData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	DrdsInstanceId *string                       `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	ParamLevel     *string                       `json:"ParamLevel,omitempty" xml:"ParamLevel,omitempty"`
	RegionId       *string                       `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SetupDrdsParamsRequest) String() string {
	return tea.Prettify(s)
}

func (s SetupDrdsParamsRequest) GoString() string {
	return s.String()
}

func (s *SetupDrdsParamsRequest) SetData(v []*SetupDrdsParamsRequestData) *SetupDrdsParamsRequest {
	s.Data = v
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

func (s *SetupDrdsParamsRequest) SetRegionId(v string) *SetupDrdsParamsRequest {
	s.RegionId = &v
	return s
}

type SetupDrdsParamsRequestData struct {
	DbName            *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	ParamRanges       *string `json:"ParamRanges,omitempty" xml:"ParamRanges,omitempty"`
	ParamType         *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
	ParamValue        *string `json:"ParamValue,omitempty" xml:"ParamValue,omitempty"`
	ParamVariableName *string `json:"ParamVariableName,omitempty" xml:"ParamVariableName,omitempty"`
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

func (s *SetupDrdsParamsRequestData) SetParamRanges(v string) *SetupDrdsParamsRequestData {
	s.ParamRanges = &v
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetupDrdsParamsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SetupDrdsParamsResponse) SetStatusCode(v int32) *SetupDrdsParamsResponse {
	s.StatusCode = &v
	return s
}

func (s *SetupDrdsParamsResponse) SetBody(v *SetupDrdsParamsResponseBody) *SetupDrdsParamsResponse {
	s.Body = v
	return s
}

type SetupRecycleBinStatusRequest struct {
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StatusAction   *string `json:"StatusAction,omitempty" xml:"StatusAction,omitempty"`
}

func (s SetupRecycleBinStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s SetupRecycleBinStatusRequest) GoString() string {
	return s.String()
}

func (s *SetupRecycleBinStatusRequest) SetDbName(v string) *SetupRecycleBinStatusRequest {
	s.DbName = &v
	return s
}

func (s *SetupRecycleBinStatusRequest) SetDrdsInstanceId(v string) *SetupRecycleBinStatusRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *SetupRecycleBinStatusRequest) SetRegionId(v string) *SetupRecycleBinStatusRequest {
	s.RegionId = &v
	return s
}

func (s *SetupRecycleBinStatusRequest) SetStatusAction(v string) *SetupRecycleBinStatusRequest {
	s.StatusAction = &v
	return s
}

type SetupRecycleBinStatusResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetupRecycleBinStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetupRecycleBinStatusResponseBody) GoString() string {
	return s.String()
}

func (s *SetupRecycleBinStatusResponseBody) SetData(v bool) *SetupRecycleBinStatusResponseBody {
	s.Data = &v
	return s
}

func (s *SetupRecycleBinStatusResponseBody) SetRequestId(v string) *SetupRecycleBinStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetupRecycleBinStatusResponseBody) SetSuccess(v bool) *SetupRecycleBinStatusResponseBody {
	s.Success = &v
	return s
}

type SetupRecycleBinStatusResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetupRecycleBinStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetupRecycleBinStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s SetupRecycleBinStatusResponse) GoString() string {
	return s.String()
}

func (s *SetupRecycleBinStatusResponse) SetHeaders(v map[string]*string) *SetupRecycleBinStatusResponse {
	s.Headers = v
	return s
}

func (s *SetupRecycleBinStatusResponse) SetStatusCode(v int32) *SetupRecycleBinStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *SetupRecycleBinStatusResponse) SetBody(v *SetupRecycleBinStatusResponseBody) *SetupRecycleBinStatusResponse {
	s.Body = v
	return s
}

type SetupTableRequest struct {
	AllowFullTableScan *bool     `json:"AllowFullTableScan,omitempty" xml:"AllowFullTableScan,omitempty"`
	DbName             *string   `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DrdsInstanceId     *string   `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	RegionId           *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TableName          []*string `json:"TableName,omitempty" xml:"TableName,omitempty" type:"Repeated"`
}

func (s SetupTableRequest) String() string {
	return tea.Prettify(s)
}

func (s SetupTableRequest) GoString() string {
	return s.String()
}

func (s *SetupTableRequest) SetAllowFullTableScan(v bool) *SetupTableRequest {
	s.AllowFullTableScan = &v
	return s
}

func (s *SetupTableRequest) SetDbName(v string) *SetupTableRequest {
	s.DbName = &v
	return s
}

func (s *SetupTableRequest) SetDrdsInstanceId(v string) *SetupTableRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *SetupTableRequest) SetRegionId(v string) *SetupTableRequest {
	s.RegionId = &v
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetupTableResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SetupTableResponse) SetStatusCode(v int32) *SetupTableResponse {
	s.StatusCode = &v
	return s
}

func (s *SetupTableResponse) SetBody(v *SetupTableResponseBody) *SetupTableResponse {
	s.Body = v
	return s
}

type StartRestoreRequest struct {
	BackupDbNames       *string `json:"BackupDbNames,omitempty" xml:"BackupDbNames,omitempty"`
	BackupId            *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	BackupLevel         *string `json:"BackupLevel,omitempty" xml:"BackupLevel,omitempty"`
	BackupMode          *string `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	DrdsInstanceId      *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	PreferredBackupTime *string `json:"PreferredBackupTime,omitempty" xml:"PreferredBackupTime,omitempty"`
}

func (s StartRestoreRequest) String() string {
	return tea.Prettify(s)
}

func (s StartRestoreRequest) GoString() string {
	return s.String()
}

func (s *StartRestoreRequest) SetBackupDbNames(v string) *StartRestoreRequest {
	s.BackupDbNames = &v
	return s
}

func (s *StartRestoreRequest) SetBackupId(v string) *StartRestoreRequest {
	s.BackupId = &v
	return s
}

func (s *StartRestoreRequest) SetBackupLevel(v string) *StartRestoreRequest {
	s.BackupLevel = &v
	return s
}

func (s *StartRestoreRequest) SetBackupMode(v string) *StartRestoreRequest {
	s.BackupMode = &v
	return s
}

func (s *StartRestoreRequest) SetDrdsInstanceId(v string) *StartRestoreRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *StartRestoreRequest) SetPreferredBackupTime(v string) *StartRestoreRequest {
	s.PreferredBackupTime = &v
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StartRestoreResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *StartRestoreResponse) SetStatusCode(v int32) *StartRestoreResponse {
	s.StatusCode = &v
	return s
}

func (s *StartRestoreResponse) SetBody(v *StartRestoreResponseBody) *StartRestoreResponse {
	s.Body = v
	return s
}

type SubmitCleanTaskRequest struct {
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	ExpandType     *string `json:"ExpandType,omitempty" xml:"ExpandType,omitempty"`
	JobId          *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ParentJobId    *string `json:"ParentJobId,omitempty" xml:"ParentJobId,omitempty"`
}

func (s SubmitCleanTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitCleanTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitCleanTaskRequest) SetDbName(v string) *SubmitCleanTaskRequest {
	s.DbName = &v
	return s
}

func (s *SubmitCleanTaskRequest) SetDrdsInstanceId(v string) *SubmitCleanTaskRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *SubmitCleanTaskRequest) SetExpandType(v string) *SubmitCleanTaskRequest {
	s.ExpandType = &v
	return s
}

func (s *SubmitCleanTaskRequest) SetJobId(v string) *SubmitCleanTaskRequest {
	s.JobId = &v
	return s
}

func (s *SubmitCleanTaskRequest) SetParentJobId(v string) *SubmitCleanTaskRequest {
	s.ParentJobId = &v
	return s
}

type SubmitCleanTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitCleanTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitCleanTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitCleanTaskResponseBody) SetRequestId(v string) *SubmitCleanTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitCleanTaskResponseBody) SetSuccess(v bool) *SubmitCleanTaskResponseBody {
	s.Success = &v
	return s
}

type SubmitCleanTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SubmitCleanTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SubmitCleanTaskResponse) SetStatusCode(v int32) *SubmitCleanTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitCleanTaskResponse) SetBody(v *SubmitCleanTaskResponseBody) *SubmitCleanTaskResponse {
	s.Body = v
	return s
}

type SubmitHotExpandPreCheckTaskRequest struct {
	DbInstType     *string   `json:"DbInstType,omitempty" xml:"DbInstType,omitempty"`
	DbName         *string   `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DrdsInstanceId *string   `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	TableList      []*string `json:"TableList,omitempty" xml:"TableList,omitempty" type:"Repeated"`
}

func (s SubmitHotExpandPreCheckTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitHotExpandPreCheckTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitHotExpandPreCheckTaskRequest) SetDbInstType(v string) *SubmitHotExpandPreCheckTaskRequest {
	s.DbInstType = &v
	return s
}

func (s *SubmitHotExpandPreCheckTaskRequest) SetDbName(v string) *SubmitHotExpandPreCheckTaskRequest {
	s.DbName = &v
	return s
}

func (s *SubmitHotExpandPreCheckTaskRequest) SetDrdsInstanceId(v string) *SubmitHotExpandPreCheckTaskRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *SubmitHotExpandPreCheckTaskRequest) SetTableList(v []*string) *SubmitHotExpandPreCheckTaskRequest {
	s.TableList = v
	return s
}

type SubmitHotExpandPreCheckTaskResponseBody struct {
	Msg       *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId    *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
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

func (s *SubmitHotExpandPreCheckTaskResponseBody) SetSuccess(v bool) *SubmitHotExpandPreCheckTaskResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitHotExpandPreCheckTaskResponseBody) SetTaskId(v int64) *SubmitHotExpandPreCheckTaskResponseBody {
	s.TaskId = &v
	return s
}

type SubmitHotExpandPreCheckTaskResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SubmitHotExpandPreCheckTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SubmitHotExpandPreCheckTaskResponse) SetStatusCode(v int32) *SubmitHotExpandPreCheckTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitHotExpandPreCheckTaskResponse) SetBody(v *SubmitHotExpandPreCheckTaskResponseBody) *SubmitHotExpandPreCheckTaskResponse {
	s.Body = v
	return s
}

type SubmitHotExpandTaskRequest struct {
	DbName               *string                                           `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DrdsInstanceId       *string                                           `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	ExtendedMapping      []*SubmitHotExpandTaskRequestExtendedMapping      `json:"ExtendedMapping,omitempty" xml:"ExtendedMapping,omitempty" type:"Repeated"`
	InstanceDbMapping    []*SubmitHotExpandTaskRequestInstanceDbMapping    `json:"InstanceDbMapping,omitempty" xml:"InstanceDbMapping,omitempty" type:"Repeated"`
	Mapping              []*SubmitHotExpandTaskRequestMapping              `json:"Mapping,omitempty" xml:"Mapping,omitempty" type:"Repeated"`
	SupperAccountMapping []*SubmitHotExpandTaskRequestSupperAccountMapping `json:"SupperAccountMapping,omitempty" xml:"SupperAccountMapping,omitempty" type:"Repeated"`
	TaskDesc             *string                                           `json:"TaskDesc,omitempty" xml:"TaskDesc,omitempty"`
	TaskName             *string                                           `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s SubmitHotExpandTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitHotExpandTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitHotExpandTaskRequest) SetDbName(v string) *SubmitHotExpandTaskRequest {
	s.DbName = &v
	return s
}

func (s *SubmitHotExpandTaskRequest) SetDrdsInstanceId(v string) *SubmitHotExpandTaskRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *SubmitHotExpandTaskRequest) SetExtendedMapping(v []*SubmitHotExpandTaskRequestExtendedMapping) *SubmitHotExpandTaskRequest {
	s.ExtendedMapping = v
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

func (s *SubmitHotExpandTaskRequest) SetTaskDesc(v string) *SubmitHotExpandTaskRequest {
	s.TaskDesc = &v
	return s
}

func (s *SubmitHotExpandTaskRequest) SetTaskName(v string) *SubmitHotExpandTaskRequest {
	s.TaskName = &v
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
	DbShardColumn *string `json:"DbShardColumn,omitempty" xml:"DbShardColumn,omitempty"`
	HotDbName     *string `json:"HotDbName,omitempty" xml:"HotDbName,omitempty"`
	HotTableName  *string `json:"HotTableName,omitempty" xml:"HotTableName,omitempty"`
	LogicTable    *string `json:"LogicTable,omitempty" xml:"LogicTable,omitempty"`
	ShardDbValue  *string `json:"ShardDbValue,omitempty" xml:"ShardDbValue,omitempty"`
	ShardTbValue  *string `json:"ShardTbValue,omitempty" xml:"ShardTbValue,omitempty"`
	TbShardColumn *string `json:"TbShardColumn,omitempty" xml:"TbShardColumn,omitempty"`
}

func (s SubmitHotExpandTaskRequestMapping) String() string {
	return tea.Prettify(s)
}

func (s SubmitHotExpandTaskRequestMapping) GoString() string {
	return s.String()
}

func (s *SubmitHotExpandTaskRequestMapping) SetDbShardColumn(v string) *SubmitHotExpandTaskRequestMapping {
	s.DbShardColumn = &v
	return s
}

func (s *SubmitHotExpandTaskRequestMapping) SetHotDbName(v string) *SubmitHotExpandTaskRequestMapping {
	s.HotDbName = &v
	return s
}

func (s *SubmitHotExpandTaskRequestMapping) SetHotTableName(v string) *SubmitHotExpandTaskRequestMapping {
	s.HotTableName = &v
	return s
}

func (s *SubmitHotExpandTaskRequestMapping) SetLogicTable(v string) *SubmitHotExpandTaskRequestMapping {
	s.LogicTable = &v
	return s
}

func (s *SubmitHotExpandTaskRequestMapping) SetShardDbValue(v string) *SubmitHotExpandTaskRequestMapping {
	s.ShardDbValue = &v
	return s
}

func (s *SubmitHotExpandTaskRequestMapping) SetShardTbValue(v string) *SubmitHotExpandTaskRequestMapping {
	s.ShardTbValue = &v
	return s
}

func (s *SubmitHotExpandTaskRequestMapping) SetTbShardColumn(v string) *SubmitHotExpandTaskRequestMapping {
	s.TbShardColumn = &v
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

type SubmitHotExpandTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitHotExpandTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitHotExpandTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitHotExpandTaskResponseBody) SetRequestId(v string) *SubmitHotExpandTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitHotExpandTaskResponseBody) SetSuccess(v bool) *SubmitHotExpandTaskResponseBody {
	s.Success = &v
	return s
}

type SubmitHotExpandTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SubmitHotExpandTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SubmitHotExpandTaskResponse) SetStatusCode(v int32) *SubmitHotExpandTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitHotExpandTaskResponse) SetBody(v *SubmitHotExpandTaskResponseBody) *SubmitHotExpandTaskResponse {
	s.Body = v
	return s
}

type SubmitSmoothExpandPreCheckRequest struct {
	DbInstType     *string `json:"DbInstType,omitempty" xml:"DbInstType,omitempty"`
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s SubmitSmoothExpandPreCheckRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitSmoothExpandPreCheckRequest) GoString() string {
	return s.String()
}

func (s *SubmitSmoothExpandPreCheckRequest) SetDbInstType(v string) *SubmitSmoothExpandPreCheckRequest {
	s.DbInstType = &v
	return s
}

func (s *SubmitSmoothExpandPreCheckRequest) SetDbName(v string) *SubmitSmoothExpandPreCheckRequest {
	s.DbName = &v
	return s
}

func (s *SubmitSmoothExpandPreCheckRequest) SetDrdsInstanceId(v string) *SubmitSmoothExpandPreCheckRequest {
	s.DrdsInstanceId = &v
	return s
}

type SubmitSmoothExpandPreCheckResponseBody struct {
	Msg       *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId    *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
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

func (s *SubmitSmoothExpandPreCheckResponseBody) SetSuccess(v bool) *SubmitSmoothExpandPreCheckResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitSmoothExpandPreCheckResponseBody) SetTaskId(v int64) *SubmitSmoothExpandPreCheckResponseBody {
	s.TaskId = &v
	return s
}

type SubmitSmoothExpandPreCheckResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SubmitSmoothExpandPreCheckResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SubmitSmoothExpandPreCheckResponse) SetStatusCode(v int32) *SubmitSmoothExpandPreCheckResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitSmoothExpandPreCheckResponse) SetBody(v *SubmitSmoothExpandPreCheckResponseBody) *SubmitSmoothExpandPreCheckResponse {
	s.Body = v
	return s
}

type SubmitSmoothExpandPreCheckTaskRequest struct {
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s SubmitSmoothExpandPreCheckTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitSmoothExpandPreCheckTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitSmoothExpandPreCheckTaskRequest) SetDbName(v string) *SubmitSmoothExpandPreCheckTaskRequest {
	s.DbName = &v
	return s
}

func (s *SubmitSmoothExpandPreCheckTaskRequest) SetDrdsInstanceId(v string) *SubmitSmoothExpandPreCheckTaskRequest {
	s.DrdsInstanceId = &v
	return s
}

type SubmitSmoothExpandPreCheckTaskResponseBody struct {
	Msg       *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId    *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
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

func (s *SubmitSmoothExpandPreCheckTaskResponseBody) SetSuccess(v bool) *SubmitSmoothExpandPreCheckTaskResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitSmoothExpandPreCheckTaskResponseBody) SetTaskId(v int64) *SubmitSmoothExpandPreCheckTaskResponseBody {
	s.TaskId = &v
	return s
}

type SubmitSmoothExpandPreCheckTaskResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SubmitSmoothExpandPreCheckTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SubmitSmoothExpandPreCheckTaskResponse) SetStatusCode(v int32) *SubmitSmoothExpandPreCheckTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitSmoothExpandPreCheckTaskResponse) SetBody(v *SubmitSmoothExpandPreCheckTaskResponseBody) *SubmitSmoothExpandPreCheckTaskResponse {
	s.Body = v
	return s
}

type SubmitSqlFlashbackTaskRequest struct {
	DbName            *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DrdsInstanceId    *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	EndTime           *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RecallRestoreType *int32  `json:"RecallRestoreType,omitempty" xml:"RecallRestoreType,omitempty"`
	RecallType        *int32  `json:"RecallType,omitempty" xml:"RecallType,omitempty"`
	SqlPk             *string `json:"SqlPk,omitempty" xml:"SqlPk,omitempty"`
	SqlType           *string `json:"SqlType,omitempty" xml:"SqlType,omitempty"`
	StartTime         *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TableName         *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	TraceId           *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s SubmitSqlFlashbackTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitSqlFlashbackTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitSqlFlashbackTaskRequest) SetDbName(v string) *SubmitSqlFlashbackTaskRequest {
	s.DbName = &v
	return s
}

func (s *SubmitSqlFlashbackTaskRequest) SetDrdsInstanceId(v string) *SubmitSqlFlashbackTaskRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *SubmitSqlFlashbackTaskRequest) SetEndTime(v string) *SubmitSqlFlashbackTaskRequest {
	s.EndTime = &v
	return s
}

func (s *SubmitSqlFlashbackTaskRequest) SetRecallRestoreType(v int32) *SubmitSqlFlashbackTaskRequest {
	s.RecallRestoreType = &v
	return s
}

func (s *SubmitSqlFlashbackTaskRequest) SetRecallType(v int32) *SubmitSqlFlashbackTaskRequest {
	s.RecallType = &v
	return s
}

func (s *SubmitSqlFlashbackTaskRequest) SetSqlPk(v string) *SubmitSqlFlashbackTaskRequest {
	s.SqlPk = &v
	return s
}

func (s *SubmitSqlFlashbackTaskRequest) SetSqlType(v string) *SubmitSqlFlashbackTaskRequest {
	s.SqlType = &v
	return s
}

func (s *SubmitSqlFlashbackTaskRequest) SetStartTime(v string) *SubmitSqlFlashbackTaskRequest {
	s.StartTime = &v
	return s
}

func (s *SubmitSqlFlashbackTaskRequest) SetTableName(v string) *SubmitSqlFlashbackTaskRequest {
	s.TableName = &v
	return s
}

func (s *SubmitSqlFlashbackTaskRequest) SetTraceId(v string) *SubmitSqlFlashbackTaskRequest {
	s.TraceId = &v
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SubmitSqlFlashbackTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SubmitSqlFlashbackTaskResponse) SetStatusCode(v int32) *SubmitSqlFlashbackTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitSqlFlashbackTaskResponse) SetBody(v *SubmitSqlFlashbackTaskResponseBody) *SubmitSqlFlashbackTaskResponse {
	s.Body = v
	return s
}

type SwitchGlobalBroadcastTypeRequest struct {
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SwitchGlobalBroadcastTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s SwitchGlobalBroadcastTypeRequest) GoString() string {
	return s.String()
}

func (s *SwitchGlobalBroadcastTypeRequest) SetDbName(v string) *SwitchGlobalBroadcastTypeRequest {
	s.DbName = &v
	return s
}

func (s *SwitchGlobalBroadcastTypeRequest) SetDrdsInstanceId(v string) *SwitchGlobalBroadcastTypeRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *SwitchGlobalBroadcastTypeRequest) SetRegionId(v string) *SwitchGlobalBroadcastTypeRequest {
	s.RegionId = &v
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
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SwitchGlobalBroadcastTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SwitchGlobalBroadcastTypeResponse) SetStatusCode(v int32) *SwitchGlobalBroadcastTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *SwitchGlobalBroadcastTypeResponse) SetBody(v *SwitchGlobalBroadcastTypeResponseBody) *SwitchGlobalBroadcastTypeResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	RegionId     *string                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId   []*string                 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceType *string                   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tag          []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
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
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s TagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *TagResourcesResponseBody) SetRequestId(v string) *TagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *TagResourcesResponseBody) SetSuccess(v bool) *TagResourcesResponseBody {
	s.Success = &v
	return s
}

type TagResourcesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *TagResourcesResponse) SetStatusCode(v int32) *TagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *TagResourcesResponse) SetBody(v *TagResourcesResponseBody) *TagResourcesResponse {
	s.Body = v
	return s
}

type UntagResourcesRequest struct {
	All          *bool     `json:"All,omitempty" xml:"All,omitempty"`
	RegionId     *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId   []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceType *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKey       []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesRequest) SetRegionId(v string) *UntagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceId(v []*string) *UntagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetTagKey(v []*string) *UntagResourcesRequest {
	s.TagKey = v
	return s
}

type UntagResourcesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UntagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponseBody) SetRequestId(v string) *UntagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UntagResourcesResponseBody) SetSuccess(v bool) *UntagResourcesResponseBody {
	s.Success = &v
	return s
}

type UntagResourcesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UntagResourcesResponse) SetStatusCode(v int32) *UntagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *UntagResourcesResponse) SetBody(v *UntagResourcesResponseBody) *UntagResourcesResponse {
	s.Body = v
	return s
}

type UpdateInstanceNetworkRequest struct {
	ClassicExpiredDays     *int32  `json:"ClassicExpiredDays,omitempty" xml:"ClassicExpiredDays,omitempty"`
	DrdsInstanceId         *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	RetainClassic          *bool   `json:"RetainClassic,omitempty" xml:"RetainClassic,omitempty"`
	SrcInstanceNetworkType *string `json:"SrcInstanceNetworkType,omitempty" xml:"SrcInstanceNetworkType,omitempty"`
}

func (s UpdateInstanceNetworkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceNetworkRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceNetworkRequest) SetClassicExpiredDays(v int32) *UpdateInstanceNetworkRequest {
	s.ClassicExpiredDays = &v
	return s
}

func (s *UpdateInstanceNetworkRequest) SetDrdsInstanceId(v string) *UpdateInstanceNetworkRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *UpdateInstanceNetworkRequest) SetRetainClassic(v bool) *UpdateInstanceNetworkRequest {
	s.RetainClassic = &v
	return s
}

func (s *UpdateInstanceNetworkRequest) SetSrcInstanceNetworkType(v string) *UpdateInstanceNetworkRequest {
	s.SrcInstanceNetworkType = &v
	return s
}

type UpdateInstanceNetworkResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateInstanceNetworkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceNetworkResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceNetworkResponseBody) SetRequestId(v string) *UpdateInstanceNetworkResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateInstanceNetworkResponseBody) SetSuccess(v bool) *UpdateInstanceNetworkResponseBody {
	s.Success = &v
	return s
}

type UpdateInstanceNetworkResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateInstanceNetworkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateInstanceNetworkResponse) SetStatusCode(v int32) *UpdateInstanceNetworkResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateInstanceNetworkResponse) SetBody(v *UpdateInstanceNetworkResponseBody) *UpdateInstanceNetworkResponse {
	s.Body = v
	return s
}

type UpdatePrivateRdsClassRequest struct {
	AutoUseCoupon  *bool   `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	DBInstanceId   *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	PrePayDuration *int32  `json:"PrePayDuration,omitempty" xml:"PrePayDuration,omitempty"`
	RdsClass       *string `json:"RdsClass,omitempty" xml:"RdsClass,omitempty"`
	Storage        *string `json:"Storage,omitempty" xml:"Storage,omitempty"`
}

func (s UpdatePrivateRdsClassRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdatePrivateRdsClassRequest) GoString() string {
	return s.String()
}

func (s *UpdatePrivateRdsClassRequest) SetAutoUseCoupon(v bool) *UpdatePrivateRdsClassRequest {
	s.AutoUseCoupon = &v
	return s
}

func (s *UpdatePrivateRdsClassRequest) SetDBInstanceId(v string) *UpdatePrivateRdsClassRequest {
	s.DBInstanceId = &v
	return s
}

func (s *UpdatePrivateRdsClassRequest) SetDrdsInstanceId(v string) *UpdatePrivateRdsClassRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *UpdatePrivateRdsClassRequest) SetPrePayDuration(v int32) *UpdatePrivateRdsClassRequest {
	s.PrePayDuration = &v
	return s
}

func (s *UpdatePrivateRdsClassRequest) SetRdsClass(v string) *UpdatePrivateRdsClassRequest {
	s.RdsClass = &v
	return s
}

func (s *UpdatePrivateRdsClassRequest) SetStorage(v string) *UpdatePrivateRdsClassRequest {
	s.Storage = &v
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdatePrivateRdsClassResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdatePrivateRdsClassResponse) SetStatusCode(v int32) *UpdatePrivateRdsClassResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePrivateRdsClassResponse) SetBody(v *UpdatePrivateRdsClassResponseBody) *UpdatePrivateRdsClassResponse {
	s.Body = v
	return s
}

type UpdateResourceGroupAttributeRequest struct {
	DrdsInstanceId     *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	NewResourceGroupId *string `json:"NewResourceGroupId,omitempty" xml:"NewResourceGroupId,omitempty"`
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateResourceGroupAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceGroupAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateResourceGroupAttributeRequest) SetDrdsInstanceId(v string) *UpdateResourceGroupAttributeRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *UpdateResourceGroupAttributeRequest) SetNewResourceGroupId(v string) *UpdateResourceGroupAttributeRequest {
	s.NewResourceGroupId = &v
	return s
}

func (s *UpdateResourceGroupAttributeRequest) SetRegionId(v string) *UpdateResourceGroupAttributeRequest {
	s.RegionId = &v
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
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateResourceGroupAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateResourceGroupAttributeResponse) SetStatusCode(v int32) *UpdateResourceGroupAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateResourceGroupAttributeResponse) SetBody(v *UpdateResourceGroupAttributeResponseBody) *UpdateResourceGroupAttributeResponse {
	s.Body = v
	return s
}

type UpgradeHiStoreInstanceRequest struct {
	DrdsInstanceId    *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	HistoreInstanceId *string `json:"HistoreInstanceId,omitempty" xml:"HistoreInstanceId,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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

func (s *UpgradeHiStoreInstanceRequest) SetHistoreInstanceId(v string) *UpgradeHiStoreInstanceRequest {
	s.HistoreInstanceId = &v
	return s
}

func (s *UpgradeHiStoreInstanceRequest) SetRegionId(v string) *UpgradeHiStoreInstanceRequest {
	s.RegionId = &v
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpgradeHiStoreInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpgradeHiStoreInstanceResponse) SetStatusCode(v int32) *UpgradeHiStoreInstanceResponse {
	s.StatusCode = &v
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpgradeInstanceVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpgradeInstanceVersionResponse) SetStatusCode(v int32) *UpgradeInstanceVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeInstanceVersionResponse) SetBody(v *UpgradeInstanceVersionResponseBody) *UpgradeInstanceVersionResponse {
	s.Body = v
	return s
}

type ValidateShardTaskRequest struct {
	DbName          *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DrdsInstanceId  *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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

func (s *ValidateShardTaskRequest) SetDbName(v string) *ValidateShardTaskRequest {
	s.DbName = &v
	return s
}

func (s *ValidateShardTaskRequest) SetDrdsInstanceId(v string) *ValidateShardTaskRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *ValidateShardTaskRequest) SetRegionId(v string) *ValidateShardTaskRequest {
	s.RegionId = &v
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
	List      []*ValidateShardTaskResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ValidateShardTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ValidateShardTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ValidateShardTaskResponseBody) SetList(v []*ValidateShardTaskResponseBodyList) *ValidateShardTaskResponseBody {
	s.List = v
	return s
}

func (s *ValidateShardTaskResponseBody) SetRequestId(v string) *ValidateShardTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ValidateShardTaskResponseBody) SetSuccess(v bool) *ValidateShardTaskResponseBody {
	s.Success = &v
	return s
}

type ValidateShardTaskResponseBodyList struct {
	Item   *string `json:"Item,omitempty" xml:"Item,omitempty"`
	Result *int32  `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s ValidateShardTaskResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s ValidateShardTaskResponseBodyList) GoString() string {
	return s.String()
}

func (s *ValidateShardTaskResponseBodyList) SetItem(v string) *ValidateShardTaskResponseBodyList {
	s.Item = &v
	return s
}

func (s *ValidateShardTaskResponseBodyList) SetResult(v int32) *ValidateShardTaskResponseBodyList {
	s.Result = &v
	return s
}

type ValidateShardTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ValidateShardTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ValidateShardTaskResponse) SetStatusCode(v int32) *ValidateShardTaskResponse {
	s.StatusCode = &v
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

func (client *Client) ChangeAccountPasswordWithOptions(request *ChangeAccountPasswordRequest, runtime *util.RuntimeOptions) (_result *ChangeAccountPasswordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		query["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		query["Password"] = request.Password
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ChangeAccountPassword"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ChangeAccountPasswordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ChangeAccountPassword(request *ChangeAccountPasswordRequest) (_result *ChangeAccountPasswordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ChangeAccountPasswordResponse{}
	_body, _err := client.ChangeAccountPasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ChangeInstanceAzoneWithOptions(request *ChangeInstanceAzoneRequest, runtime *util.RuntimeOptions) (_result *ChangeInstanceAzoneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsRegionId)) {
		query["DrdsRegionId"] = request.DrdsRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.OriginAzoneId)) {
		query["OriginAzoneId"] = request.OriginAzoneId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetAzoneId)) {
		query["TargetAzoneId"] = request.TargetAzoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ChangeInstanceAzone"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ChangeInstanceAzoneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ChangeInstanceAzone(request *ChangeInstanceAzoneRequest) (_result *ChangeInstanceAzoneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ChangeInstanceAzoneResponse{}
	_body, _err := client.ChangeInstanceAzoneWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckDrdsDbName"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckDrdsDbNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckExpandStatus"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckExpandStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckSqlAuditEnableStatus"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckSqlAuditEnableStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		query["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.DbInstType)) {
		query["DbInstType"] = request.DbInstType
	}

	if !tea.BoolValue(util.IsUnset(request.DbInstanceIsCreating)) {
		query["DbInstanceIsCreating"] = request.DbInstanceIsCreating
	}

	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Encode)) {
		query["Encode"] = request.Encode
	}

	if !tea.BoolValue(util.IsUnset(request.InstDbName)) {
		query["InstDbName"] = request.InstDbName
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		query["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.RdsInstance)) {
		query["RdsInstance"] = request.RdsInstance
	}

	if !tea.BoolValue(util.IsUnset(request.RdsSuperAccount)) {
		query["RdsSuperAccount"] = request.RdsSuperAccount
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDrdsDB"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDrdsDBResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Duration)) {
		query["Duration"] = request.Duration
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceSeries)) {
		query["InstanceSeries"] = request.InstanceSeries
	}

	if !tea.BoolValue(util.IsUnset(request.IsAutoRenew)) {
		query["IsAutoRenew"] = request.IsAutoRenew
	}

	if !tea.BoolValue(util.IsUnset(request.MasterInstId)) {
		query["MasterInstId"] = request.MasterInstId
	}

	if !tea.BoolValue(util.IsUnset(request.MySQLVersion)) {
		query["MySQLVersion"] = request.MySQLVersion
	}

	if !tea.BoolValue(util.IsUnset(request.PayType)) {
		query["PayType"] = request.PayType
	}

	if !tea.BoolValue(util.IsUnset(request.PricingCycle)) {
		query["PricingCycle"] = request.PricingCycle
	}

	if !tea.BoolValue(util.IsUnset(request.Quantity)) {
		query["Quantity"] = request.Quantity
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Specification)) {
		query["Specification"] = request.Specification
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	if !tea.BoolValue(util.IsUnset(request.VswitchId)) {
		query["VswitchId"] = request.VswitchId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	if !tea.BoolValue(util.IsUnset(request.IsHa)) {
		query["isHa"] = request.IsHa
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDrdsInstance"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDrdsInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		query["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.DbPrivilege)) {
		query["DbPrivilege"] = request.DbPrivilege
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		query["Password"] = request.Password
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateInstanceAccount"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateInstanceAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateInstanceInternetAddress"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateInstanceInternetAddressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Params)) {
		query["Params"] = request.Params
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateOrderForRds"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateOrderForRdsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceTableName)) {
		query["SourceTableName"] = request.SourceTableName
	}

	if !tea.BoolValue(util.IsUnset(request.TargetTableName)) {
		query["TargetTableName"] = request.TargetTableName
	}

	if !tea.BoolValue(util.IsUnset(request.TaskType)) {
		query["TaskType"] = request.TaskType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateShardTask"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateShardTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeBackMenu"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeBackMenuResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackupId)) {
		query["BackupId"] = request.BackupId
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PreferredRestoreTime)) {
		query["PreferredRestoreTime"] = request.PreferredRestoreTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeBackupDbs"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeBackupDbsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeBackupLocal"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeBackupLocalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeBackupPolicy"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeBackupPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeBackupSets"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeBackupSetsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeBackupTimes"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeBackupTimesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		query["Query"] = request.Query
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeBroadcastTables"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeBroadcastTablesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		query["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.DbInstType)) {
		query["DbInstType"] = request.DbInstType
	}

	if !tea.BoolValue(util.IsUnset(request.DbInstanceId)) {
		query["DbInstanceId"] = request.DbInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		query["Password"] = request.Password
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDbInstanceDbs"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDbInstanceDbsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbInstType)) {
		query["DbInstType"] = request.DbInstType
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Search)) {
		query["Search"] = request.Search
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDbInstances"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDbInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDrdsDB"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDrdsDBResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbInstanceId)) {
		query["DbInstanceId"] = request.DbInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDrdsDBCluster"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDrdsDBClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DescribeDrdsDBIpWhiteListWithOptions(request *DescribeDrdsDBIpWhiteListRequest, runtime *util.RuntimeOptions) (_result *DescribeDrdsDBIpWhiteListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDrdsDBIpWhiteList"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDrdsDBIpWhiteListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DescribeDrdsDBsWithOptions(request *DescribeDrdsDBsRequest, runtime *util.RuntimeOptions) (_result *DescribeDrdsDBsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDrdsDBs"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDrdsDBsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DescribeDrdsDbInstanceWithOptions(request *DescribeDrdsDbInstanceRequest, runtime *util.RuntimeOptions) (_result *DescribeDrdsDbInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbInstanceId)) {
		query["DbInstanceId"] = request.DbInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDrdsDbInstance"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDrdsDbInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDrdsDbInstances"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDrdsDbInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DescribeDrdsDbRdsNameListWithOptions(request *DescribeDrdsDbRdsNameListRequest, runtime *util.RuntimeOptions) (_result *DescribeDrdsDbRdsNameListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDrdsDbRdsNameList"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDrdsDbRdsNameListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DescribeDrdsInstanceWithOptions(request *DescribeDrdsInstanceRequest, runtime *util.RuntimeOptions) (_result *DescribeDrdsInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDrdsInstance"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDrdsInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Key)) {
		query["Key"] = request.Key
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDrdsInstanceDbMonitor"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDrdsInstanceDbMonitorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDrdsInstanceLevelTasks"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDrdsInstanceLevelTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Key)) {
		query["Key"] = request.Key
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodMultiple)) {
		query["PeriodMultiple"] = request.PeriodMultiple
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDrdsInstanceMonitor"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDrdsInstanceMonitorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DescribeDrdsInstanceVersionWithOptions(request *DescribeDrdsInstanceVersionRequest, runtime *util.RuntimeOptions) (_result *DescribeDrdsInstanceVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDrdsInstanceVersion"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDrdsInstanceVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DescribeDrdsInstancesWithOptions(request *DescribeDrdsInstancesRequest, runtime *util.RuntimeOptions) (_result *DescribeDrdsInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Expired)) {
		query["Expired"] = request.Expired
	}

	if !tea.BoolValue(util.IsUnset(request.Mix)) {
		query["Mix"] = request.Mix
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProductVersion)) {
		query["ProductVersion"] = request.ProductVersion
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDrdsInstances"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDrdsInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DescribeDrdsParamsWithOptions(request *DescribeDrdsParamsRequest, runtime *util.RuntimeOptions) (_result *DescribeDrdsParamsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ParamLevel)) {
		query["ParamLevel"] = request.ParamLevel
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDrdsParams"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDrdsParamsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDrdsRdsInstances"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDrdsRdsInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DbNamePattern)) {
		query["DbNamePattern"] = request.DbNamePattern
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDrdsShardingDbs"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDrdsShardingDbsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.ExeTime)) {
		query["ExeTime"] = request.ExeTime
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDrdsSlowSqls"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDrdsSlowSqlsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDrdsSqlAuditStatus"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDrdsSqlAuditStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskType)) {
		query["TaskType"] = request.TaskType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDrdsTasks"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDrdsTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeExpandLogicTableInfoList"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeExpandLogicTableInfoListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DescribeHotDbListWithOptions(request *DescribeHotDbListRequest, runtime *util.RuntimeOptions) (_result *DescribeHotDbListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeHotDbList"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeHotDbListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DescribeInstDbLogInfoWithOptions(request *DescribeInstDbLogInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeInstDbLogInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInstDbLogInfo"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInstDbLogInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInstDbSlsInfo"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInstDbSlsInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DescribeInstanceAccountsWithOptions(request *DescribeInstanceAccountsRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceAccountsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInstanceAccounts"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInstanceAccountsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DescribeInstanceSwitchAzoneWithOptions(request *DescribeInstanceSwitchAzoneRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceSwitchAzoneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInstanceSwitchAzone"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInstanceSwitchAzoneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInstanceSwitchNetwork"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInstanceSwitchNetworkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DescribePreCheckResultWithOptions(request *DescribePreCheckResultRequest, runtime *util.RuntimeOptions) (_result *DescribePreCheckResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePreCheckResult"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePreCheckResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DescribeRDSPerformanceWithOptions(request *DescribeRDSPerformanceRequest, runtime *util.RuntimeOptions) (_result *DescribeRDSPerformanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbInstType)) {
		query["DbInstType"] = request.DbInstType
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Keys)) {
		query["Keys"] = request.Keys
	}

	if !tea.BoolValue(util.IsUnset(request.RdsInstanceId)) {
		query["RdsInstanceId"] = request.RdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRDSPerformance"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRDSPerformanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DescribeRdsCommodityWithOptions(request *DescribeRdsCommodityRequest, runtime *util.RuntimeOptions) (_result *DescribeRdsCommodityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CommodityCode)) {
		query["CommodityCode"] = request.CommodityCode
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OrderType)) {
		query["OrderType"] = request.OrderType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRdsCommodity"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRdsCommodityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DescribeRdsPerformanceSummaryWithOptions(request *DescribeRdsPerformanceSummaryRequest, runtime *util.RuntimeOptions) (_result *DescribeRdsPerformanceSummaryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RdsInstanceId)) {
		query["RdsInstanceId"] = request.RdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRdsPerformanceSummary"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRdsPerformanceSummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbInstType)) {
		query["DbInstType"] = request.DbInstType
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RdsInstance)) {
		query["RdsInstance"] = request.RdsInstance
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRdsSuperAccountInstances"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRdsSuperAccountInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DescribeRecycleBinStatusWithOptions(request *DescribeRecycleBinStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeRecycleBinStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRecycleBinStatus"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRecycleBinStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRecycleBinStatus(request *DescribeRecycleBinStatusRequest) (_result *DescribeRecycleBinStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRecycleBinStatusResponse{}
	_body, _err := client.DescribeRecycleBinStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRecycleBinTablesWithOptions(request *DescribeRecycleBinTablesRequest, runtime *util.RuntimeOptions) (_result *DescribeRecycleBinTablesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRecycleBinTables"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRecycleBinTablesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRecycleBinTables(request *DescribeRecycleBinTablesRequest) (_result *DescribeRecycleBinTablesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRecycleBinTablesResponse{}
	_body, _err := client.DescribeRecycleBinTablesWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackupDbNames)) {
		query["BackupDbNames"] = request.BackupDbNames
	}

	if !tea.BoolValue(util.IsUnset(request.BackupId)) {
		query["BackupId"] = request.BackupId
	}

	if !tea.BoolValue(util.IsUnset(request.BackupLevel)) {
		query["BackupLevel"] = request.BackupLevel
	}

	if !tea.BoolValue(util.IsUnset(request.BackupMode)) {
		query["BackupMode"] = request.BackupMode
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PreferredBackupTime)) {
		query["PreferredBackupTime"] = request.PreferredBackupTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRestoreOrder"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRestoreOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceTableName)) {
		query["SourceTableName"] = request.SourceTableName
	}

	if !tea.BoolValue(util.IsUnset(request.TargetTableName)) {
		query["TargetTableName"] = request.TargetTableName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeShardTaskInfo"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeShardTaskInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DescribeSqlFlashbakTaskWithOptions(request *DescribeSqlFlashbakTaskRequest, runtime *util.RuntimeOptions) (_result *DescribeSqlFlashbakTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSqlFlashbakTask"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSqlFlashbakTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TableName)) {
		query["TableName"] = request.TableName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTable"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		query["Query"] = request.Query
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TableType)) {
		query["TableType"] = request.TableType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTableListByType"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTableListByTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		query["Query"] = request.Query
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTables"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTablesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DisableSqlAudit"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableSqlAuditResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) EnableInstanceIpv6AddressWithOptions(request *EnableInstanceIpv6AddressRequest, runtime *util.RuntimeOptions) (_result *EnableInstanceIpv6AddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableInstanceIpv6Address"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableInstanceIpv6AddressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableInstanceIpv6Address(request *EnableInstanceIpv6AddressRequest) (_result *EnableInstanceIpv6AddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableInstanceIpv6AddressResponse{}
	_body, _err := client.EnableInstanceIpv6AddressWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.IsRecall)) {
		query["IsRecall"] = request.IsRecall
	}

	if !tea.BoolValue(util.IsUnset(request.RecallEndTimestamp)) {
		query["RecallEndTimestamp"] = request.RecallEndTimestamp
	}

	if !tea.BoolValue(util.IsUnset(request.RecallStartTimestamp)) {
		query["RecallStartTimestamp"] = request.RecallStartTimestamp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableSqlAudit"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableSqlAuditResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableSqlFlashbackMatchSwitch"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableSqlFlashbackMatchSwitchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) FlashbackRecycleBinTableWithOptions(request *FlashbackRecycleBinTableRequest, runtime *util.RuntimeOptions) (_result *FlashbackRecycleBinTableResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TableName)) {
		query["TableName"] = request.TableName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("FlashbackRecycleBinTable"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &FlashbackRecycleBinTableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FlashbackRecycleBinTable(request *FlashbackRecycleBinTableRequest) (_result *FlashbackRecycleBinTableResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FlashbackRecycleBinTableResponse{}
	_body, _err := client.FlashbackRecycleBinTableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDrdsDbRdsRelationInfoWithOptions(request *GetDrdsDbRdsRelationInfoRequest, runtime *util.RuntimeOptions) (_result *GetDrdsDbRdsRelationInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDrdsDbRdsRelationInfo"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDrdsDbRdsRelationInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDrdsDbRdsRelationInfo(request *GetDrdsDbRdsRelationInfoRequest) (_result *GetDrdsDbRdsRelationInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDrdsDbRdsRelationInfoResponse{}
	_body, _err := client.GetDrdsDbRdsRelationInfoWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagResources"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Params)) {
		query["Params"] = request.Params
	}

	if !tea.BoolValue(util.IsUnset(request.RdsAction)) {
		query["RdsAction"] = request.RdsAction
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ManagePrivateRds"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ManagePrivateRdsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) ModifyAccountDescriptionWithOptions(request *ModifyAccountDescriptionRequest, runtime *util.RuntimeOptions) (_result *ModifyAccountDescriptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		query["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyAccountDescription"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyAccountDescriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyAccountDescription(request *ModifyAccountDescriptionRequest) (_result *ModifyAccountDescriptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyAccountDescriptionResponse{}
	_body, _err := client.ModifyAccountDescriptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyAccountPrivilegeWithOptions(request *ModifyAccountPrivilegeRequest, runtime *util.RuntimeOptions) (_result *ModifyAccountPrivilegeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		query["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.DbPrivilege)) {
		query["DbPrivilege"] = request.DbPrivilege
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyAccountPrivilege"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyAccountPrivilegeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyAccountPrivilege(request *ModifyAccountPrivilegeRequest) (_result *ModifyAccountPrivilegeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyAccountPrivilegeResponse{}
	_body, _err := client.ModifyAccountPrivilegeWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDrdsInstanceDescription"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDrdsInstanceDescriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupAttribute)) {
		query["GroupAttribute"] = request.GroupAttribute
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.IpWhiteList)) {
		query["IpWhiteList"] = request.IpWhiteList
	}

	if !tea.BoolValue(util.IsUnset(request.Mode)) {
		query["Mode"] = request.Mode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDrdsIpWhiteList"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDrdsIpWhiteListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) ModifyPolarDbReadWeightWithOptions(request *ModifyPolarDbReadWeightRequest, runtime *util.RuntimeOptions) (_result *ModifyPolarDbReadWeightResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbInstanceId)) {
		query["DbInstanceId"] = request.DbInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DbNodeIds)) {
		query["DbNodeIds"] = request.DbNodeIds
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Weights)) {
		query["Weights"] = request.Weights
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyPolarDbReadWeight"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyPolarDbReadWeightResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyPolarDbReadWeight(request *ModifyPolarDbReadWeightRequest) (_result *ModifyPolarDbReadWeightResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyPolarDbReadWeightResponse{}
	_body, _err := client.ModifyPolarDbReadWeightWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceNames)) {
		query["InstanceNames"] = request.InstanceNames
	}

	if !tea.BoolValue(util.IsUnset(request.Weights)) {
		query["Weights"] = request.Weights
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyRdsReadWeight"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyRdsReadWeightResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackupDbNames)) {
		query["BackupDbNames"] = request.BackupDbNames
	}

	if !tea.BoolValue(util.IsUnset(request.BackupLevel)) {
		query["BackupLevel"] = request.BackupLevel
	}

	if !tea.BoolValue(util.IsUnset(request.BackupMode)) {
		query["BackupMode"] = request.BackupMode
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PutStartBackup"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PutStartBackupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) RefreshDrdsAtomUrlWithOptions(request *RefreshDrdsAtomUrlRequest, runtime *util.RuntimeOptions) (_result *RefreshDrdsAtomUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RefreshDrdsAtomUrl"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RefreshDrdsAtomUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RefreshDrdsAtomUrl(request *RefreshDrdsAtomUrlRequest) (_result *RefreshDrdsAtomUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RefreshDrdsAtomUrlResponse{}
	_body, _err := client.RefreshDrdsAtomUrlWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ReleaseInstanceInternetAddress"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReleaseInstanceInternetAddressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackupId)) {
		query["BackupId"] = request.BackupId
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveBackupsSet"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveBackupsSetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveDrdsDb"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveDrdsDbResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveDrdsDbFailedRecord"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveDrdsDbFailedRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveDrdsInstance"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveDrdsInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		query["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveInstanceAccount"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveInstanceAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) RemoveRecycleBinTableWithOptions(request *RemoveRecycleBinTableRequest, runtime *util.RuntimeOptions) (_result *RemoveRecycleBinTableResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TableName)) {
		query["TableName"] = request.TableName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveRecycleBinTable"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveRecycleBinTableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveRecycleBinTable(request *RemoveRecycleBinTableRequest) (_result *RemoveRecycleBinTableResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveRecycleBinTableResponse{}
	_body, _err := client.RemoveRecycleBinTableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RestartDrdsInstanceWithOptions(request *RestartDrdsInstanceRequest, runtime *util.RuntimeOptions) (_result *RestartDrdsInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RestartDrdsInstance"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RestartDrdsInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RestartDrdsInstance(request *RestartDrdsInstanceRequest) (_result *RestartDrdsInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RestartDrdsInstanceResponse{}
	_body, _err := client.RestartDrdsInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RollbackInstanceVersionWithOptions(request *RollbackInstanceVersionRequest, runtime *util.RuntimeOptions) (_result *RollbackInstanceVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RollbackInstanceVersion"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RollbackInstanceVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RollbackInstanceVersion(request *RollbackInstanceVersionRequest) (_result *RollbackInstanceVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RollbackInstanceVersionResponse{}
	_body, _err := client.RollbackInstanceVersionWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.HighSpaceUsageProtection)) {
		query["HighSpaceUsageProtection"] = request.HighSpaceUsageProtection
	}

	if !tea.BoolValue(util.IsUnset(request.LocalLogRetentionHours)) {
		query["LocalLogRetentionHours"] = request.LocalLogRetentionHours
	}

	if !tea.BoolValue(util.IsUnset(request.LocalLogRetentionSpace)) {
		query["LocalLogRetentionSpace"] = request.LocalLogRetentionSpace
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetBackupLocal"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetBackupLocalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackupDbNames)) {
		query["BackupDbNames"] = request.BackupDbNames
	}

	if !tea.BoolValue(util.IsUnset(request.BackupLevel)) {
		query["BackupLevel"] = request.BackupLevel
	}

	if !tea.BoolValue(util.IsUnset(request.BackupLog)) {
		query["BackupLog"] = request.BackupLog
	}

	if !tea.BoolValue(util.IsUnset(request.BackupMode)) {
		query["BackupMode"] = request.BackupMode
	}

	if !tea.BoolValue(util.IsUnset(request.DataBackupRetentionPeriod)) {
		query["DataBackupRetentionPeriod"] = request.DataBackupRetentionPeriod
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.LogBackupRetentionPeriod)) {
		query["LogBackupRetentionPeriod"] = request.LogBackupRetentionPeriod
	}

	if !tea.BoolValue(util.IsUnset(request.PreferredBackupEndTime)) {
		query["PreferredBackupEndTime"] = request.PreferredBackupEndTime
	}

	if !tea.BoolValue(util.IsUnset(request.PreferredBackupPeriod)) {
		query["PreferredBackupPeriod"] = request.PreferredBackupPeriod
	}

	if !tea.BoolValue(util.IsUnset(request.PreferredBackupStartTime)) {
		query["PreferredBackupStartTime"] = request.PreferredBackupStartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetBackupPolicy"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetBackupPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Active)) {
		query["Active"] = request.Active
	}

	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TableName)) {
		query["TableName"] = request.TableName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetupBroadcastTables"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetupBroadcastTablesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Data)) {
		query["Data"] = request.Data
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ParamLevel)) {
		query["ParamLevel"] = request.ParamLevel
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetupDrdsParams"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetupDrdsParamsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) SetupRecycleBinStatusWithOptions(request *SetupRecycleBinStatusRequest, runtime *util.RuntimeOptions) (_result *SetupRecycleBinStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StatusAction)) {
		query["StatusAction"] = request.StatusAction
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetupRecycleBinStatus"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetupRecycleBinStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetupRecycleBinStatus(request *SetupRecycleBinStatusRequest) (_result *SetupRecycleBinStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetupRecycleBinStatusResponse{}
	_body, _err := client.SetupRecycleBinStatusWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AllowFullTableScan)) {
		query["AllowFullTableScan"] = request.AllowFullTableScan
	}

	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TableName)) {
		query["TableName"] = request.TableName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetupTable"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetupTableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackupDbNames)) {
		query["BackupDbNames"] = request.BackupDbNames
	}

	if !tea.BoolValue(util.IsUnset(request.BackupId)) {
		query["BackupId"] = request.BackupId
	}

	if !tea.BoolValue(util.IsUnset(request.BackupLevel)) {
		query["BackupLevel"] = request.BackupLevel
	}

	if !tea.BoolValue(util.IsUnset(request.BackupMode)) {
		query["BackupMode"] = request.BackupMode
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PreferredBackupTime)) {
		query["PreferredBackupTime"] = request.PreferredBackupTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartRestore"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartRestoreResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ExpandType)) {
		query["ExpandType"] = request.ExpandType
	}

	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		query["JobId"] = request.JobId
	}

	if !tea.BoolValue(util.IsUnset(request.ParentJobId)) {
		query["ParentJobId"] = request.ParentJobId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitCleanTask"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitCleanTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbInstType)) {
		query["DbInstType"] = request.DbInstType
	}

	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.TableList)) {
		query["TableList"] = request.TableList
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitHotExpandPreCheckTask"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitHotExpandPreCheckTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ExtendedMapping)) {
		query["ExtendedMapping"] = request.ExtendedMapping
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceDbMapping)) {
		query["InstanceDbMapping"] = request.InstanceDbMapping
	}

	if !tea.BoolValue(util.IsUnset(request.Mapping)) {
		query["Mapping"] = request.Mapping
	}

	if !tea.BoolValue(util.IsUnset(request.SupperAccountMapping)) {
		query["SupperAccountMapping"] = request.SupperAccountMapping
	}

	if !tea.BoolValue(util.IsUnset(request.TaskDesc)) {
		query["TaskDesc"] = request.TaskDesc
	}

	if !tea.BoolValue(util.IsUnset(request.TaskName)) {
		query["TaskName"] = request.TaskName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitHotExpandTask"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitHotExpandTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbInstType)) {
		query["DbInstType"] = request.DbInstType
	}

	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitSmoothExpandPreCheck"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitSmoothExpandPreCheckResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitSmoothExpandPreCheckTask"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitSmoothExpandPreCheckTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) SubmitSqlFlashbackTaskWithOptions(request *SubmitSqlFlashbackTaskRequest, runtime *util.RuntimeOptions) (_result *SubmitSqlFlashbackTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.RecallRestoreType)) {
		query["RecallRestoreType"] = request.RecallRestoreType
	}

	if !tea.BoolValue(util.IsUnset(request.RecallType)) {
		query["RecallType"] = request.RecallType
	}

	if !tea.BoolValue(util.IsUnset(request.SqlPk)) {
		query["SqlPk"] = request.SqlPk
	}

	if !tea.BoolValue(util.IsUnset(request.SqlType)) {
		query["SqlType"] = request.SqlType
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TableName)) {
		query["TableName"] = request.TableName
	}

	if !tea.BoolValue(util.IsUnset(request.TraceId)) {
		query["TraceId"] = request.TraceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitSqlFlashbackTask"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitSqlFlashbackTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) SwitchGlobalBroadcastTypeWithOptions(request *SwitchGlobalBroadcastTypeRequest, runtime *util.RuntimeOptions) (_result *SwitchGlobalBroadcastTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SwitchGlobalBroadcastType"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SwitchGlobalBroadcastTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TagResources"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.All)) {
		query["All"] = request.All
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TagKey)) {
		query["TagKey"] = request.TagKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UntagResources"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClassicExpiredDays)) {
		query["ClassicExpiredDays"] = request.ClassicExpiredDays
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RetainClassic)) {
		query["RetainClassic"] = request.RetainClassic
	}

	if !tea.BoolValue(util.IsUnset(request.SrcInstanceNetworkType)) {
		query["SrcInstanceNetworkType"] = request.SrcInstanceNetworkType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateInstanceNetwork"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateInstanceNetworkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoUseCoupon)) {
		query["AutoUseCoupon"] = request.AutoUseCoupon
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PrePayDuration)) {
		query["PrePayDuration"] = request.PrePayDuration
	}

	if !tea.BoolValue(util.IsUnset(request.RdsClass)) {
		query["RdsClass"] = request.RdsClass
	}

	if !tea.BoolValue(util.IsUnset(request.Storage)) {
		query["Storage"] = request.Storage
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdatePrivateRdsClass"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdatePrivateRdsClassResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.NewResourceGroupId)) {
		query["NewResourceGroupId"] = request.NewResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateResourceGroupAttribute"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateResourceGroupAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.HistoreInstanceId)) {
		query["HistoreInstanceId"] = request.HistoreInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpgradeHiStoreInstance"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpgradeHiStoreInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Rpm)) {
		query["Rpm"] = request.Rpm
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpgradeInstanceVersion"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpgradeInstanceVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DrdsInstanceId)) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceTableName)) {
		query["SourceTableName"] = request.SourceTableName
	}

	if !tea.BoolValue(util.IsUnset(request.TargetTableName)) {
		query["TargetTableName"] = request.TargetTableName
	}

	if !tea.BoolValue(util.IsUnset(request.TaskType)) {
		query["TaskType"] = request.TaskType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ValidateShardTask"),
		Version:     tea.String("2019-01-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ValidateShardTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
