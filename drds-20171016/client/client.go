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

type CreateDrdsAccountRequest struct {
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	Password       *string `json:"Password,omitempty" xml:"Password,omitempty"`
	UserName       *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s CreateDrdsAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDrdsAccountRequest) GoString() string {
	return s.String()
}

func (s *CreateDrdsAccountRequest) SetDbName(v string) *CreateDrdsAccountRequest {
	s.DbName = &v
	return s
}

func (s *CreateDrdsAccountRequest) SetDrdsInstanceId(v string) *CreateDrdsAccountRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *CreateDrdsAccountRequest) SetPassword(v string) *CreateDrdsAccountRequest {
	s.Password = &v
	return s
}

func (s *CreateDrdsAccountRequest) SetUserName(v string) *CreateDrdsAccountRequest {
	s.UserName = &v
	return s
}

type CreateDrdsAccountResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDrdsAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDrdsAccountResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDrdsAccountResponseBody) SetRequestId(v string) *CreateDrdsAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDrdsAccountResponseBody) SetSuccess(v bool) *CreateDrdsAccountResponseBody {
	s.Success = &v
	return s
}

type CreateDrdsAccountResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDrdsAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDrdsAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDrdsAccountResponse) GoString() string {
	return s.String()
}

func (s *CreateDrdsAccountResponse) SetHeaders(v map[string]*string) *CreateDrdsAccountResponse {
	s.Headers = v
	return s
}

func (s *CreateDrdsAccountResponse) SetBody(v *CreateDrdsAccountResponseBody) *CreateDrdsAccountResponse {
	s.Body = v
	return s
}

type CreateDrdsDBRequest struct {
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	Encode         *string `json:"Encode,omitempty" xml:"Encode,omitempty"`
	Password       *string `json:"Password,omitempty" xml:"Password,omitempty"`
	RdsInstances   *string `json:"RdsInstances,omitempty" xml:"RdsInstances,omitempty"`
}

func (s CreateDrdsDBRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDrdsDBRequest) GoString() string {
	return s.String()
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

func (s *CreateDrdsDBRequest) SetPassword(v string) *CreateDrdsDBRequest {
	s.Password = &v
	return s
}

func (s *CreateDrdsDBRequest) SetRdsInstances(v string) *CreateDrdsDBRequest {
	s.RdsInstances = &v
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
	ClientToken    *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Duration       *int32  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	InstanceSeries *string `json:"InstanceSeries,omitempty" xml:"InstanceSeries,omitempty"`
	IsAutoRenew    *bool   `json:"IsAutoRenew,omitempty" xml:"IsAutoRenew,omitempty"`
	IsHa           *bool   `json:"IsHa,omitempty" xml:"IsHa,omitempty"`
	PayType        *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	PricingCycle   *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	Quantity       *int32  `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Specification  *string `json:"Specification,omitempty" xml:"Specification,omitempty"`
	Type           *string `json:"Type,omitempty" xml:"Type,omitempty"`
	VpcId          *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VswitchId      *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	ZoneId         *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
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

func (s *CreateDrdsInstanceRequest) SetIsHa(v bool) *CreateDrdsInstanceRequest {
	s.IsHa = &v
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
	DrdsInstanceId []*string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty" type:"Repeated"`
}

func (s CreateDrdsInstanceResponseBodyDataDrdsInstanceIdList) String() string {
	return tea.Prettify(s)
}

func (s CreateDrdsInstanceResponseBodyDataDrdsInstanceIdList) GoString() string {
	return s.String()
}

func (s *CreateDrdsInstanceResponseBodyDataDrdsInstanceIdList) SetDrdsInstanceId(v []*string) *CreateDrdsInstanceResponseBodyDataDrdsInstanceIdList {
	s.DrdsInstanceId = v
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

type CreateReadOnlyAccountRequest struct {
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	Password       *string `json:"password,omitempty" xml:"password,omitempty"`
}

func (s CreateReadOnlyAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateReadOnlyAccountRequest) GoString() string {
	return s.String()
}

func (s *CreateReadOnlyAccountRequest) SetDbName(v string) *CreateReadOnlyAccountRequest {
	s.DbName = &v
	return s
}

func (s *CreateReadOnlyAccountRequest) SetDrdsInstanceId(v string) *CreateReadOnlyAccountRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *CreateReadOnlyAccountRequest) SetPassword(v string) *CreateReadOnlyAccountRequest {
	s.Password = &v
	return s
}

type CreateReadOnlyAccountResponseBody struct {
	Data      *CreateReadOnlyAccountResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateReadOnlyAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateReadOnlyAccountResponseBody) GoString() string {
	return s.String()
}

func (s *CreateReadOnlyAccountResponseBody) SetData(v *CreateReadOnlyAccountResponseBodyData) *CreateReadOnlyAccountResponseBody {
	s.Data = v
	return s
}

func (s *CreateReadOnlyAccountResponseBody) SetRequestId(v string) *CreateReadOnlyAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateReadOnlyAccountResponseBody) SetSuccess(v bool) *CreateReadOnlyAccountResponseBody {
	s.Success = &v
	return s
}

type CreateReadOnlyAccountResponseBodyData struct {
	AccountName    *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s CreateReadOnlyAccountResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateReadOnlyAccountResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateReadOnlyAccountResponseBodyData) SetAccountName(v string) *CreateReadOnlyAccountResponseBodyData {
	s.AccountName = &v
	return s
}

func (s *CreateReadOnlyAccountResponseBodyData) SetDbName(v string) *CreateReadOnlyAccountResponseBodyData {
	s.DbName = &v
	return s
}

func (s *CreateReadOnlyAccountResponseBodyData) SetDrdsInstanceId(v string) *CreateReadOnlyAccountResponseBodyData {
	s.DrdsInstanceId = &v
	return s
}

type CreateReadOnlyAccountResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateReadOnlyAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateReadOnlyAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateReadOnlyAccountResponse) GoString() string {
	return s.String()
}

func (s *CreateReadOnlyAccountResponse) SetHeaders(v map[string]*string) *CreateReadOnlyAccountResponse {
	s.Headers = v
	return s
}

func (s *CreateReadOnlyAccountResponse) SetBody(v *CreateReadOnlyAccountResponseBody) *CreateReadOnlyAccountResponse {
	s.Body = v
	return s
}

type DeleteDrdsDBRequest struct {
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s DeleteDrdsDBRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDrdsDBRequest) GoString() string {
	return s.String()
}

func (s *DeleteDrdsDBRequest) SetDbName(v string) *DeleteDrdsDBRequest {
	s.DbName = &v
	return s
}

func (s *DeleteDrdsDBRequest) SetDrdsInstanceId(v string) *DeleteDrdsDBRequest {
	s.DrdsInstanceId = &v
	return s
}

type DeleteDrdsDBResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteDrdsDBResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDrdsDBResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDrdsDBResponseBody) SetRequestId(v string) *DeleteDrdsDBResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDrdsDBResponseBody) SetSuccess(v bool) *DeleteDrdsDBResponseBody {
	s.Success = &v
	return s
}

type DeleteDrdsDBResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDrdsDBResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDrdsDBResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDrdsDBResponse) GoString() string {
	return s.String()
}

func (s *DeleteDrdsDBResponse) SetHeaders(v map[string]*string) *DeleteDrdsDBResponse {
	s.Headers = v
	return s
}

func (s *DeleteDrdsDBResponse) SetBody(v *DeleteDrdsDBResponseBody) *DeleteDrdsDBResponse {
	s.Body = v
	return s
}

type DeleteFailedDrdsDBRequest struct {
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s DeleteFailedDrdsDBRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFailedDrdsDBRequest) GoString() string {
	return s.String()
}

func (s *DeleteFailedDrdsDBRequest) SetDbName(v string) *DeleteFailedDrdsDBRequest {
	s.DbName = &v
	return s
}

func (s *DeleteFailedDrdsDBRequest) SetDrdsInstanceId(v string) *DeleteFailedDrdsDBRequest {
	s.DrdsInstanceId = &v
	return s
}

type DeleteFailedDrdsDBResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteFailedDrdsDBResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFailedDrdsDBResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFailedDrdsDBResponseBody) SetRequestId(v string) *DeleteFailedDrdsDBResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFailedDrdsDBResponseBody) SetSuccess(v bool) *DeleteFailedDrdsDBResponseBody {
	s.Success = &v
	return s
}

type DeleteFailedDrdsDBResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteFailedDrdsDBResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteFailedDrdsDBResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFailedDrdsDBResponse) GoString() string {
	return s.String()
}

func (s *DeleteFailedDrdsDBResponse) SetHeaders(v map[string]*string) *DeleteFailedDrdsDBResponse {
	s.Headers = v
	return s
}

func (s *DeleteFailedDrdsDBResponse) SetBody(v *DeleteFailedDrdsDBResponseBody) *DeleteFailedDrdsDBResponse {
	s.Body = v
	return s
}

type DescribeCreateDrdsInstanceStatusRequest struct {
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s DescribeCreateDrdsInstanceStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCreateDrdsInstanceStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeCreateDrdsInstanceStatusRequest) SetDrdsInstanceId(v string) *DescribeCreateDrdsInstanceStatusRequest {
	s.DrdsInstanceId = &v
	return s
}

type DescribeCreateDrdsInstanceStatusResponseBody struct {
	Data      *DescribeCreateDrdsInstanceStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeCreateDrdsInstanceStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCreateDrdsInstanceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCreateDrdsInstanceStatusResponseBody) SetData(v *DescribeCreateDrdsInstanceStatusResponseBodyData) *DescribeCreateDrdsInstanceStatusResponseBody {
	s.Data = v
	return s
}

func (s *DescribeCreateDrdsInstanceStatusResponseBody) SetRequestId(v string) *DescribeCreateDrdsInstanceStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCreateDrdsInstanceStatusResponseBody) SetSuccess(v bool) *DescribeCreateDrdsInstanceStatusResponseBody {
	s.Success = &v
	return s
}

type DescribeCreateDrdsInstanceStatusResponseBodyData struct {
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeCreateDrdsInstanceStatusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeCreateDrdsInstanceStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeCreateDrdsInstanceStatusResponseBodyData) SetStatus(v string) *DescribeCreateDrdsInstanceStatusResponseBodyData {
	s.Status = &v
	return s
}

type DescribeCreateDrdsInstanceStatusResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCreateDrdsInstanceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCreateDrdsInstanceStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCreateDrdsInstanceStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeCreateDrdsInstanceStatusResponse) SetHeaders(v map[string]*string) *DescribeCreateDrdsInstanceStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeCreateDrdsInstanceStatusResponse) SetBody(v *DescribeCreateDrdsInstanceStatusResponseBody) *DescribeCreateDrdsInstanceStatusResponse {
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
	DbName     *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	Mode       *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	Msg        *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	Status     *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
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

func (s *DescribeDrdsDBResponseBodyData) SetDbName(v string) *DescribeDrdsDBResponseBodyData {
	s.DbName = &v
	return s
}

func (s *DescribeDrdsDBResponseBodyData) SetMode(v string) *DescribeDrdsDBResponseBodyData {
	s.Mode = &v
	return s
}

func (s *DescribeDrdsDBResponseBodyData) SetMsg(v string) *DescribeDrdsDBResponseBodyData {
	s.Msg = &v
	return s
}

func (s *DescribeDrdsDBResponseBodyData) SetStatus(v int32) *DescribeDrdsDBResponseBodyData {
	s.Status = &v
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

type DescribeDrdsDBIpWhiteListRequest struct {
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	GroupName      *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
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

type DescribeDrdsDBIpWhiteListResponseBody struct {
	Data      *DescribeDrdsDBIpWhiteListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDrdsDBIpWhiteListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDBIpWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDBIpWhiteListResponseBody) SetData(v *DescribeDrdsDBIpWhiteListResponseBodyData) *DescribeDrdsDBIpWhiteListResponseBody {
	s.Data = v
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

type DescribeDrdsDBIpWhiteListResponseBodyData struct {
	IpWhiteList *DescribeDrdsDBIpWhiteListResponseBodyDataIpWhiteList `json:"IpWhiteList,omitempty" xml:"IpWhiteList,omitempty" type:"Struct"`
}

func (s DescribeDrdsDBIpWhiteListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDBIpWhiteListResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDBIpWhiteListResponseBodyData) SetIpWhiteList(v *DescribeDrdsDBIpWhiteListResponseBodyDataIpWhiteList) *DescribeDrdsDBIpWhiteListResponseBodyData {
	s.IpWhiteList = v
	return s
}

type DescribeDrdsDBIpWhiteListResponseBodyDataIpWhiteList struct {
	Ip []*string `json:"Ip,omitempty" xml:"Ip,omitempty" type:"Repeated"`
}

func (s DescribeDrdsDBIpWhiteListResponseBodyDataIpWhiteList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsDBIpWhiteListResponseBodyDataIpWhiteList) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDBIpWhiteListResponseBodyDataIpWhiteList) SetIp(v []*string) *DescribeDrdsDBIpWhiteListResponseBodyDataIpWhiteList {
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

type DescribeDrdsDBsRequest struct {
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
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

type DescribeDrdsDBsResponseBody struct {
	Data      *DescribeDrdsDBsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
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

func (s *DescribeDrdsDBsResponseBody) SetRequestId(v string) *DescribeDrdsDBsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrdsDBsResponseBody) SetSuccess(v bool) *DescribeDrdsDBsResponseBody {
	s.Success = &v
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
	DbName     *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	Mode       *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	Msg        *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	Status     *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
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

func (s *DescribeDrdsDBsResponseBodyDataDb) SetDbName(v string) *DescribeDrdsDBsResponseBodyDataDb {
	s.DbName = &v
	return s
}

func (s *DescribeDrdsDBsResponseBodyDataDb) SetMode(v string) *DescribeDrdsDBsResponseBodyDataDb {
	s.Mode = &v
	return s
}

func (s *DescribeDrdsDBsResponseBodyDataDb) SetMsg(v string) *DescribeDrdsDBsResponseBodyDataDb {
	s.Msg = &v
	return s
}

func (s *DescribeDrdsDBsResponseBodyDataDb) SetStatus(v int32) *DescribeDrdsDBsResponseBodyDataDb {
	s.Status = &v
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

type DescribeDrdsInstanceRequest struct {
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
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
	CreateTime         *int64                                    `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description        *string                                   `json:"Description,omitempty" xml:"Description,omitempty"`
	DrdsInstanceId     *string                                   `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	NetworkType        *string                                   `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	RegionId           *string                                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Specification      *string                                   `json:"Specification,omitempty" xml:"Specification,omitempty"`
	Status             *string                                   `json:"Status,omitempty" xml:"Status,omitempty"`
	Type               *string                                   `json:"Type,omitempty" xml:"Type,omitempty"`
	Version            *int64                                    `json:"Version,omitempty" xml:"Version,omitempty"`
	Vips               *DescribeDrdsInstanceResponseBodyDataVips `json:"Vips,omitempty" xml:"Vips,omitempty" type:"Struct"`
	VpcCloudInstanceId *string                                   `json:"VpcCloudInstanceId,omitempty" xml:"VpcCloudInstanceId,omitempty"`
	ZoneId             *string                                   `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDrdsInstanceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceResponseBodyData) GoString() string {
	return s.String()
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

func (s *DescribeDrdsInstanceResponseBodyData) SetNetworkType(v string) *DescribeDrdsInstanceResponseBodyData {
	s.NetworkType = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetRegionId(v string) *DescribeDrdsInstanceResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetSpecification(v string) *DescribeDrdsInstanceResponseBodyData {
	s.Specification = &v
	return s
}

func (s *DescribeDrdsInstanceResponseBodyData) SetStatus(v string) *DescribeDrdsInstanceResponseBodyData {
	s.Status = &v
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
	IP        *string `json:"IP,omitempty" xml:"IP,omitempty"`
	Port      *string `json:"Port,omitempty" xml:"Port,omitempty"`
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty"`
	VpcId     *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
}

func (s DescribeDrdsInstanceResponseBodyDataVipsVip) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceResponseBodyDataVipsVip) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceResponseBodyDataVipsVip) SetIP(v string) *DescribeDrdsInstanceResponseBodyDataVipsVip {
	s.IP = &v
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
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	EndTime        *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Key            *string `json:"Key,omitempty" xml:"Key,omitempty"`
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

func (s *DescribeDrdsInstanceDbMonitorRequest) SetStartTime(v int64) *DescribeDrdsInstanceDbMonitorRequest {
	s.StartTime = &v
	return s
}

type DescribeDrdsInstanceDbMonitorResponseBody struct {
	Data      *DescribeDrdsInstanceDbMonitorResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDrdsInstanceDbMonitorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceDbMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceDbMonitorResponseBody) SetData(v *DescribeDrdsInstanceDbMonitorResponseBodyData) *DescribeDrdsInstanceDbMonitorResponseBody {
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
	PartialPerformanceData []*DescribeDrdsInstanceDbMonitorResponseBodyDataPartialPerformanceData `json:"PartialPerformanceData,omitempty" xml:"PartialPerformanceData,omitempty" type:"Repeated"`
}

func (s DescribeDrdsInstanceDbMonitorResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceDbMonitorResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceDbMonitorResponseBodyData) SetPartialPerformanceData(v []*DescribeDrdsInstanceDbMonitorResponseBodyDataPartialPerformanceData) *DescribeDrdsInstanceDbMonitorResponseBodyData {
	s.PartialPerformanceData = v
	return s
}

type DescribeDrdsInstanceDbMonitorResponseBodyDataPartialPerformanceData struct {
	Key    *string                                                                    `json:"Key,omitempty" xml:"Key,omitempty"`
	Unit   *string                                                                    `json:"Unit,omitempty" xml:"Unit,omitempty"`
	Values *DescribeDrdsInstanceDbMonitorResponseBodyDataPartialPerformanceDataValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Struct"`
}

func (s DescribeDrdsInstanceDbMonitorResponseBodyDataPartialPerformanceData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceDbMonitorResponseBodyDataPartialPerformanceData) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceDbMonitorResponseBodyDataPartialPerformanceData) SetKey(v string) *DescribeDrdsInstanceDbMonitorResponseBodyDataPartialPerformanceData {
	s.Key = &v
	return s
}

func (s *DescribeDrdsInstanceDbMonitorResponseBodyDataPartialPerformanceData) SetUnit(v string) *DescribeDrdsInstanceDbMonitorResponseBodyDataPartialPerformanceData {
	s.Unit = &v
	return s
}

func (s *DescribeDrdsInstanceDbMonitorResponseBodyDataPartialPerformanceData) SetValues(v *DescribeDrdsInstanceDbMonitorResponseBodyDataPartialPerformanceDataValues) *DescribeDrdsInstanceDbMonitorResponseBodyDataPartialPerformanceData {
	s.Values = v
	return s
}

type DescribeDrdsInstanceDbMonitorResponseBodyDataPartialPerformanceDataValues struct {
	PerformanceValue []*DescribeDrdsInstanceDbMonitorResponseBodyDataPartialPerformanceDataValuesPerformanceValue `json:"PerformanceValue,omitempty" xml:"PerformanceValue,omitempty" type:"Repeated"`
}

func (s DescribeDrdsInstanceDbMonitorResponseBodyDataPartialPerformanceDataValues) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceDbMonitorResponseBodyDataPartialPerformanceDataValues) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceDbMonitorResponseBodyDataPartialPerformanceDataValues) SetPerformanceValue(v []*DescribeDrdsInstanceDbMonitorResponseBodyDataPartialPerformanceDataValuesPerformanceValue) *DescribeDrdsInstanceDbMonitorResponseBodyDataPartialPerformanceDataValues {
	s.PerformanceValue = v
	return s
}

type DescribeDrdsInstanceDbMonitorResponseBodyDataPartialPerformanceDataValuesPerformanceValue struct {
	Date  *int64  `json:"Date,omitempty" xml:"Date,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDrdsInstanceDbMonitorResponseBodyDataPartialPerformanceDataValuesPerformanceValue) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceDbMonitorResponseBodyDataPartialPerformanceDataValuesPerformanceValue) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceDbMonitorResponseBodyDataPartialPerformanceDataValuesPerformanceValue) SetDate(v int64) *DescribeDrdsInstanceDbMonitorResponseBodyDataPartialPerformanceDataValuesPerformanceValue {
	s.Date = &v
	return s
}

func (s *DescribeDrdsInstanceDbMonitorResponseBodyDataPartialPerformanceDataValuesPerformanceValue) SetValue(v string) *DescribeDrdsInstanceDbMonitorResponseBodyDataPartialPerformanceDataValuesPerformanceValue {
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

type DescribeDrdsInstanceMonitorRequest struct {
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	EndTime        *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Key            *string `json:"Key,omitempty" xml:"Key,omitempty"`
	PeriodMultiple *int32  `json:"PeriodMultiple,omitempty" xml:"PeriodMultiple,omitempty"`
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

func (s *DescribeDrdsInstanceMonitorRequest) SetStartTime(v int64) *DescribeDrdsInstanceMonitorRequest {
	s.StartTime = &v
	return s
}

type DescribeDrdsInstanceMonitorResponseBody struct {
	Data      *DescribeDrdsInstanceMonitorResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDrdsInstanceMonitorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceMonitorResponseBody) SetData(v *DescribeDrdsInstanceMonitorResponseBodyData) *DescribeDrdsInstanceMonitorResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDrdsInstanceMonitorResponseBody) SetRequestId(v string) *DescribeDrdsInstanceMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrdsInstanceMonitorResponseBody) SetSuccess(v bool) *DescribeDrdsInstanceMonitorResponseBody {
	s.Success = &v
	return s
}

type DescribeDrdsInstanceMonitorResponseBodyData struct {
	PartialPerformanceData []*DescribeDrdsInstanceMonitorResponseBodyDataPartialPerformanceData `json:"PartialPerformanceData,omitempty" xml:"PartialPerformanceData,omitempty" type:"Repeated"`
}

func (s DescribeDrdsInstanceMonitorResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceMonitorResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceMonitorResponseBodyData) SetPartialPerformanceData(v []*DescribeDrdsInstanceMonitorResponseBodyDataPartialPerformanceData) *DescribeDrdsInstanceMonitorResponseBodyData {
	s.PartialPerformanceData = v
	return s
}

type DescribeDrdsInstanceMonitorResponseBodyDataPartialPerformanceData struct {
	Key    *string                                                                  `json:"Key,omitempty" xml:"Key,omitempty"`
	Unit   *string                                                                  `json:"Unit,omitempty" xml:"Unit,omitempty"`
	Values *DescribeDrdsInstanceMonitorResponseBodyDataPartialPerformanceDataValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Struct"`
}

func (s DescribeDrdsInstanceMonitorResponseBodyDataPartialPerformanceData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceMonitorResponseBodyDataPartialPerformanceData) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceMonitorResponseBodyDataPartialPerformanceData) SetKey(v string) *DescribeDrdsInstanceMonitorResponseBodyDataPartialPerformanceData {
	s.Key = &v
	return s
}

func (s *DescribeDrdsInstanceMonitorResponseBodyDataPartialPerformanceData) SetUnit(v string) *DescribeDrdsInstanceMonitorResponseBodyDataPartialPerformanceData {
	s.Unit = &v
	return s
}

func (s *DescribeDrdsInstanceMonitorResponseBodyDataPartialPerformanceData) SetValues(v *DescribeDrdsInstanceMonitorResponseBodyDataPartialPerformanceDataValues) *DescribeDrdsInstanceMonitorResponseBodyDataPartialPerformanceData {
	s.Values = v
	return s
}

type DescribeDrdsInstanceMonitorResponseBodyDataPartialPerformanceDataValues struct {
	PerformanceValue []*DescribeDrdsInstanceMonitorResponseBodyDataPartialPerformanceDataValuesPerformanceValue `json:"PerformanceValue,omitempty" xml:"PerformanceValue,omitempty" type:"Repeated"`
}

func (s DescribeDrdsInstanceMonitorResponseBodyDataPartialPerformanceDataValues) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceMonitorResponseBodyDataPartialPerformanceDataValues) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceMonitorResponseBodyDataPartialPerformanceDataValues) SetPerformanceValue(v []*DescribeDrdsInstanceMonitorResponseBodyDataPartialPerformanceDataValuesPerformanceValue) *DescribeDrdsInstanceMonitorResponseBodyDataPartialPerformanceDataValues {
	s.PerformanceValue = v
	return s
}

type DescribeDrdsInstanceMonitorResponseBodyDataPartialPerformanceDataValuesPerformanceValue struct {
	Date  *int64  `json:"Date,omitempty" xml:"Date,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDrdsInstanceMonitorResponseBodyDataPartialPerformanceDataValuesPerformanceValue) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceMonitorResponseBodyDataPartialPerformanceDataValuesPerformanceValue) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceMonitorResponseBodyDataPartialPerformanceDataValuesPerformanceValue) SetDate(v int64) *DescribeDrdsInstanceMonitorResponseBodyDataPartialPerformanceDataValuesPerformanceValue {
	s.Date = &v
	return s
}

func (s *DescribeDrdsInstanceMonitorResponseBodyDataPartialPerformanceDataValuesPerformanceValue) SetValue(v string) *DescribeDrdsInstanceMonitorResponseBodyDataPartialPerformanceDataValuesPerformanceValue {
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

type DescribeDrdsInstanceNetInfoForInnerRequest struct {
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s DescribeDrdsInstanceNetInfoForInnerRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceNetInfoForInnerRequest) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceNetInfoForInnerRequest) SetDrdsInstanceId(v string) *DescribeDrdsInstanceNetInfoForInnerRequest {
	s.DrdsInstanceId = &v
	return s
}

type DescribeDrdsInstanceNetInfoForInnerResponseBody struct {
	DrdsInstanceId *string                                                  `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	NetInfos       *DescribeDrdsInstanceNetInfoForInnerResponseBodyNetInfos `json:"NetInfos,omitempty" xml:"NetInfos,omitempty" type:"Struct"`
	NetworkType    *string                                                  `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	RequestId      *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDrdsInstanceNetInfoForInnerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceNetInfoForInnerResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceNetInfoForInnerResponseBody) SetDrdsInstanceId(v string) *DescribeDrdsInstanceNetInfoForInnerResponseBody {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeDrdsInstanceNetInfoForInnerResponseBody) SetNetInfos(v *DescribeDrdsInstanceNetInfoForInnerResponseBodyNetInfos) *DescribeDrdsInstanceNetInfoForInnerResponseBody {
	s.NetInfos = v
	return s
}

func (s *DescribeDrdsInstanceNetInfoForInnerResponseBody) SetNetworkType(v string) *DescribeDrdsInstanceNetInfoForInnerResponseBody {
	s.NetworkType = &v
	return s
}

func (s *DescribeDrdsInstanceNetInfoForInnerResponseBody) SetRequestId(v string) *DescribeDrdsInstanceNetInfoForInnerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrdsInstanceNetInfoForInnerResponseBody) SetSuccess(v bool) *DescribeDrdsInstanceNetInfoForInnerResponseBody {
	s.Success = &v
	return s
}

type DescribeDrdsInstanceNetInfoForInnerResponseBodyNetInfos struct {
	NetInfo []*DescribeDrdsInstanceNetInfoForInnerResponseBodyNetInfosNetInfo `json:"NetInfo,omitempty" xml:"NetInfo,omitempty" type:"Repeated"`
}

func (s DescribeDrdsInstanceNetInfoForInnerResponseBodyNetInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceNetInfoForInnerResponseBodyNetInfos) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceNetInfoForInnerResponseBodyNetInfos) SetNetInfo(v []*DescribeDrdsInstanceNetInfoForInnerResponseBodyNetInfosNetInfo) *DescribeDrdsInstanceNetInfoForInnerResponseBodyNetInfos {
	s.NetInfo = v
	return s
}

type DescribeDrdsInstanceNetInfoForInnerResponseBodyNetInfosNetInfo struct {
	IP       *string `json:"IP,omitempty" xml:"IP,omitempty"`
	IsForVpc *bool   `json:"IsForVpc,omitempty" xml:"IsForVpc,omitempty"`
	Port     *string `json:"Port,omitempty" xml:"Port,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeDrdsInstanceNetInfoForInnerResponseBodyNetInfosNetInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceNetInfoForInnerResponseBodyNetInfosNetInfo) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceNetInfoForInnerResponseBodyNetInfosNetInfo) SetIP(v string) *DescribeDrdsInstanceNetInfoForInnerResponseBodyNetInfosNetInfo {
	s.IP = &v
	return s
}

func (s *DescribeDrdsInstanceNetInfoForInnerResponseBodyNetInfosNetInfo) SetIsForVpc(v bool) *DescribeDrdsInstanceNetInfoForInnerResponseBodyNetInfosNetInfo {
	s.IsForVpc = &v
	return s
}

func (s *DescribeDrdsInstanceNetInfoForInnerResponseBodyNetInfosNetInfo) SetPort(v string) *DescribeDrdsInstanceNetInfoForInnerResponseBodyNetInfosNetInfo {
	s.Port = &v
	return s
}

func (s *DescribeDrdsInstanceNetInfoForInnerResponseBodyNetInfosNetInfo) SetType(v string) *DescribeDrdsInstanceNetInfoForInnerResponseBodyNetInfosNetInfo {
	s.Type = &v
	return s
}

type DescribeDrdsInstanceNetInfoForInnerResponse struct {
	Headers map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDrdsInstanceNetInfoForInnerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDrdsInstanceNetInfoForInnerResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstanceNetInfoForInnerResponse) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceNetInfoForInnerResponse) SetHeaders(v map[string]*string) *DescribeDrdsInstanceNetInfoForInnerResponse {
	s.Headers = v
	return s
}

func (s *DescribeDrdsInstanceNetInfoForInnerResponse) SetBody(v *DescribeDrdsInstanceNetInfoForInnerResponseBody) *DescribeDrdsInstanceNetInfoForInnerResponse {
	s.Body = v
	return s
}

type DescribeDrdsInstancesRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Tags     *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeDrdsInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstancesRequest) SetRegionId(v string) *DescribeDrdsInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDrdsInstancesRequest) SetTags(v string) *DescribeDrdsInstancesRequest {
	s.Tags = &v
	return s
}

func (s *DescribeDrdsInstancesRequest) SetType(v string) *DescribeDrdsInstancesRequest {
	s.Type = &v
	return s
}

type DescribeDrdsInstancesResponseBody struct {
	Data      *DescribeDrdsInstancesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDrdsInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstancesResponseBody) SetData(v *DescribeDrdsInstancesResponseBodyData) *DescribeDrdsInstancesResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDrdsInstancesResponseBody) SetRequestId(v string) *DescribeDrdsInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBody) SetSuccess(v bool) *DescribeDrdsInstancesResponseBody {
	s.Success = &v
	return s
}

type DescribeDrdsInstancesResponseBodyData struct {
	Instance []*DescribeDrdsInstancesResponseBodyDataInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Repeated"`
}

func (s DescribeDrdsInstancesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstancesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstancesResponseBodyData) SetInstance(v []*DescribeDrdsInstancesResponseBodyDataInstance) *DescribeDrdsInstancesResponseBodyData {
	s.Instance = v
	return s
}

type DescribeDrdsInstancesResponseBodyDataInstance struct {
	CreateTime         *int64                                                    `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description        *string                                                   `json:"Description,omitempty" xml:"Description,omitempty"`
	DrdsInstanceId     *string                                                   `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	InstRole           *string                                                   `json:"InstRole,omitempty" xml:"InstRole,omitempty"`
	MasterInstId       *string                                                   `json:"MasterInstId,omitempty" xml:"MasterInstId,omitempty"`
	NetworkType        *string                                                   `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	RegionId           *string                                                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SlaveInstId        *DescribeDrdsInstancesResponseBodyDataInstanceSlaveInstId `json:"SlaveInstId,omitempty" xml:"SlaveInstId,omitempty" type:"Struct"`
	Status             *string                                                   `json:"Status,omitempty" xml:"Status,omitempty"`
	Type               *string                                                   `json:"Type,omitempty" xml:"Type,omitempty"`
	Version            *int64                                                    `json:"Version,omitempty" xml:"Version,omitempty"`
	Vips               *DescribeDrdsInstancesResponseBodyDataInstanceVips        `json:"Vips,omitempty" xml:"Vips,omitempty" type:"Struct"`
	VpcCloudInstanceId *string                                                   `json:"VpcCloudInstanceId,omitempty" xml:"VpcCloudInstanceId,omitempty"`
	ZoneId             *string                                                   `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDrdsInstancesResponseBodyDataInstance) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstancesResponseBodyDataInstance) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstancesResponseBodyDataInstance) SetCreateTime(v int64) *DescribeDrdsInstancesResponseBodyDataInstance {
	s.CreateTime = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyDataInstance) SetDescription(v string) *DescribeDrdsInstancesResponseBodyDataInstance {
	s.Description = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyDataInstance) SetDrdsInstanceId(v string) *DescribeDrdsInstancesResponseBodyDataInstance {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyDataInstance) SetInstRole(v string) *DescribeDrdsInstancesResponseBodyDataInstance {
	s.InstRole = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyDataInstance) SetMasterInstId(v string) *DescribeDrdsInstancesResponseBodyDataInstance {
	s.MasterInstId = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyDataInstance) SetNetworkType(v string) *DescribeDrdsInstancesResponseBodyDataInstance {
	s.NetworkType = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyDataInstance) SetRegionId(v string) *DescribeDrdsInstancesResponseBodyDataInstance {
	s.RegionId = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyDataInstance) SetSlaveInstId(v *DescribeDrdsInstancesResponseBodyDataInstanceSlaveInstId) *DescribeDrdsInstancesResponseBodyDataInstance {
	s.SlaveInstId = v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyDataInstance) SetStatus(v string) *DescribeDrdsInstancesResponseBodyDataInstance {
	s.Status = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyDataInstance) SetType(v string) *DescribeDrdsInstancesResponseBodyDataInstance {
	s.Type = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyDataInstance) SetVersion(v int64) *DescribeDrdsInstancesResponseBodyDataInstance {
	s.Version = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyDataInstance) SetVips(v *DescribeDrdsInstancesResponseBodyDataInstanceVips) *DescribeDrdsInstancesResponseBodyDataInstance {
	s.Vips = v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyDataInstance) SetVpcCloudInstanceId(v string) *DescribeDrdsInstancesResponseBodyDataInstance {
	s.VpcCloudInstanceId = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyDataInstance) SetZoneId(v string) *DescribeDrdsInstancesResponseBodyDataInstance {
	s.ZoneId = &v
	return s
}

type DescribeDrdsInstancesResponseBodyDataInstanceSlaveInstId struct {
	InstId []*string `json:"instId,omitempty" xml:"instId,omitempty" type:"Repeated"`
}

func (s DescribeDrdsInstancesResponseBodyDataInstanceSlaveInstId) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstancesResponseBodyDataInstanceSlaveInstId) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstancesResponseBodyDataInstanceSlaveInstId) SetInstId(v []*string) *DescribeDrdsInstancesResponseBodyDataInstanceSlaveInstId {
	s.InstId = v
	return s
}

type DescribeDrdsInstancesResponseBodyDataInstanceVips struct {
	Vip []*DescribeDrdsInstancesResponseBodyDataInstanceVipsVip `json:"Vip,omitempty" xml:"Vip,omitempty" type:"Repeated"`
}

func (s DescribeDrdsInstancesResponseBodyDataInstanceVips) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstancesResponseBodyDataInstanceVips) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstancesResponseBodyDataInstanceVips) SetVip(v []*DescribeDrdsInstancesResponseBodyDataInstanceVipsVip) *DescribeDrdsInstancesResponseBodyDataInstanceVips {
	s.Vip = v
	return s
}

type DescribeDrdsInstancesResponseBodyDataInstanceVipsVip struct {
	IP        *string `json:"IP,omitempty" xml:"IP,omitempty"`
	Port      *string `json:"Port,omitempty" xml:"Port,omitempty"`
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty"`
	VpcId     *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
}

func (s DescribeDrdsInstancesResponseBodyDataInstanceVipsVip) String() string {
	return tea.Prettify(s)
}

func (s DescribeDrdsInstancesResponseBodyDataInstanceVipsVip) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstancesResponseBodyDataInstanceVipsVip) SetIP(v string) *DescribeDrdsInstancesResponseBodyDataInstanceVipsVip {
	s.IP = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyDataInstanceVipsVip) SetPort(v string) *DescribeDrdsInstancesResponseBodyDataInstanceVipsVip {
	s.Port = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyDataInstanceVipsVip) SetType(v string) *DescribeDrdsInstancesResponseBodyDataInstanceVipsVip {
	s.Type = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyDataInstanceVipsVip) SetVpcId(v string) *DescribeDrdsInstancesResponseBodyDataInstanceVipsVip {
	s.VpcId = &v
	return s
}

func (s *DescribeDrdsInstancesResponseBodyDataInstanceVipsVip) SetVswitchId(v string) *DescribeDrdsInstancesResponseBodyDataInstanceVipsVip {
	s.VswitchId = &v
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

type DescribeRdsListRequest struct {
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRdsListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsListRequest) GoString() string {
	return s.String()
}

func (s *DescribeRdsListRequest) SetDbName(v string) *DescribeRdsListRequest {
	s.DbName = &v
	return s
}

func (s *DescribeRdsListRequest) SetDrdsInstanceId(v string) *DescribeRdsListRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeRdsListRequest) SetRegionId(v string) *DescribeRdsListRequest {
	s.RegionId = &v
	return s
}

type DescribeRdsListResponseBody struct {
	Data      *DescribeRdsListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeRdsListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRdsListResponseBody) SetData(v *DescribeRdsListResponseBodyData) *DescribeRdsListResponseBody {
	s.Data = v
	return s
}

func (s *DescribeRdsListResponseBody) SetRequestId(v string) *DescribeRdsListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRdsListResponseBody) SetSuccess(v bool) *DescribeRdsListResponseBody {
	s.Success = &v
	return s
}

type DescribeRdsListResponseBodyData struct {
	RdsInstance []*DescribeRdsListResponseBodyDataRdsInstance `json:"RdsInstance,omitempty" xml:"RdsInstance,omitempty" type:"Repeated"`
}

func (s DescribeRdsListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsListResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeRdsListResponseBodyData) SetRdsInstance(v []*DescribeRdsListResponseBodyDataRdsInstance) *DescribeRdsListResponseBodyData {
	s.RdsInstance = v
	return s
}

type DescribeRdsListResponseBodyDataRdsInstance struct {
	ConnectUrl       *string                                                     `json:"ConnectUrl,omitempty" xml:"ConnectUrl,omitempty"`
	DbType           *string                                                     `json:"DbType,omitempty" xml:"DbType,omitempty"`
	InstanceId       *int32                                                      `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName     *string                                                     `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	InstanceStatus   *string                                                     `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	Port             *int32                                                      `json:"Port,omitempty" xml:"Port,omitempty"`
	ReadOnlyChildren *DescribeRdsListResponseBodyDataRdsInstanceReadOnlyChildren `json:"ReadOnlyChildren,omitempty" xml:"ReadOnlyChildren,omitempty" type:"Struct"`
	ReadWeight       *int32                                                      `json:"ReadWeight,omitempty" xml:"ReadWeight,omitempty"`
}

func (s DescribeRdsListResponseBodyDataRdsInstance) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsListResponseBodyDataRdsInstance) GoString() string {
	return s.String()
}

func (s *DescribeRdsListResponseBodyDataRdsInstance) SetConnectUrl(v string) *DescribeRdsListResponseBodyDataRdsInstance {
	s.ConnectUrl = &v
	return s
}

func (s *DescribeRdsListResponseBodyDataRdsInstance) SetDbType(v string) *DescribeRdsListResponseBodyDataRdsInstance {
	s.DbType = &v
	return s
}

func (s *DescribeRdsListResponseBodyDataRdsInstance) SetInstanceId(v int32) *DescribeRdsListResponseBodyDataRdsInstance {
	s.InstanceId = &v
	return s
}

func (s *DescribeRdsListResponseBodyDataRdsInstance) SetInstanceName(v string) *DescribeRdsListResponseBodyDataRdsInstance {
	s.InstanceName = &v
	return s
}

func (s *DescribeRdsListResponseBodyDataRdsInstance) SetInstanceStatus(v string) *DescribeRdsListResponseBodyDataRdsInstance {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeRdsListResponseBodyDataRdsInstance) SetPort(v int32) *DescribeRdsListResponseBodyDataRdsInstance {
	s.Port = &v
	return s
}

func (s *DescribeRdsListResponseBodyDataRdsInstance) SetReadOnlyChildren(v *DescribeRdsListResponseBodyDataRdsInstanceReadOnlyChildren) *DescribeRdsListResponseBodyDataRdsInstance {
	s.ReadOnlyChildren = v
	return s
}

func (s *DescribeRdsListResponseBodyDataRdsInstance) SetReadWeight(v int32) *DescribeRdsListResponseBodyDataRdsInstance {
	s.ReadWeight = &v
	return s
}

type DescribeRdsListResponseBodyDataRdsInstanceReadOnlyChildren struct {
	Child []*DescribeRdsListResponseBodyDataRdsInstanceReadOnlyChildrenChild `json:"Child,omitempty" xml:"Child,omitempty" type:"Repeated"`
}

func (s DescribeRdsListResponseBodyDataRdsInstanceReadOnlyChildren) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsListResponseBodyDataRdsInstanceReadOnlyChildren) GoString() string {
	return s.String()
}

func (s *DescribeRdsListResponseBodyDataRdsInstanceReadOnlyChildren) SetChild(v []*DescribeRdsListResponseBodyDataRdsInstanceReadOnlyChildrenChild) *DescribeRdsListResponseBodyDataRdsInstanceReadOnlyChildren {
	s.Child = v
	return s
}

type DescribeRdsListResponseBodyDataRdsInstanceReadOnlyChildrenChild struct {
	ConnectUrl     *string `json:"ConnectUrl,omitempty" xml:"ConnectUrl,omitempty"`
	DbType         *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName   *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	ReadWeight     *int32  `json:"ReadWeight,omitempty" xml:"ReadWeight,omitempty"`
	Port           *int32  `json:"port,omitempty" xml:"port,omitempty"`
}

func (s DescribeRdsListResponseBodyDataRdsInstanceReadOnlyChildrenChild) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsListResponseBodyDataRdsInstanceReadOnlyChildrenChild) GoString() string {
	return s.String()
}

func (s *DescribeRdsListResponseBodyDataRdsInstanceReadOnlyChildrenChild) SetConnectUrl(v string) *DescribeRdsListResponseBodyDataRdsInstanceReadOnlyChildrenChild {
	s.ConnectUrl = &v
	return s
}

func (s *DescribeRdsListResponseBodyDataRdsInstanceReadOnlyChildrenChild) SetDbType(v string) *DescribeRdsListResponseBodyDataRdsInstanceReadOnlyChildrenChild {
	s.DbType = &v
	return s
}

func (s *DescribeRdsListResponseBodyDataRdsInstanceReadOnlyChildrenChild) SetInstanceId(v string) *DescribeRdsListResponseBodyDataRdsInstanceReadOnlyChildrenChild {
	s.InstanceId = &v
	return s
}

func (s *DescribeRdsListResponseBodyDataRdsInstanceReadOnlyChildrenChild) SetInstanceName(v string) *DescribeRdsListResponseBodyDataRdsInstanceReadOnlyChildrenChild {
	s.InstanceName = &v
	return s
}

func (s *DescribeRdsListResponseBodyDataRdsInstanceReadOnlyChildrenChild) SetInstanceStatus(v string) *DescribeRdsListResponseBodyDataRdsInstanceReadOnlyChildrenChild {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeRdsListResponseBodyDataRdsInstanceReadOnlyChildrenChild) SetReadWeight(v int32) *DescribeRdsListResponseBodyDataRdsInstanceReadOnlyChildrenChild {
	s.ReadWeight = &v
	return s
}

func (s *DescribeRdsListResponseBodyDataRdsInstanceReadOnlyChildrenChild) SetPort(v int32) *DescribeRdsListResponseBodyDataRdsInstanceReadOnlyChildrenChild {
	s.Port = &v
	return s
}

type DescribeRdsListResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRdsListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRdsListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsListResponse) GoString() string {
	return s.String()
}

func (s *DescribeRdsListResponse) SetHeaders(v map[string]*string) *DescribeRdsListResponse {
	s.Headers = v
	return s
}

func (s *DescribeRdsListResponse) SetBody(v *DescribeRdsListResponseBody) *DescribeRdsListResponse {
	s.Body = v
	return s
}

type DescribeReadOnlyAccountRequest struct {
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s DescribeReadOnlyAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeReadOnlyAccountRequest) GoString() string {
	return s.String()
}

func (s *DescribeReadOnlyAccountRequest) SetDbName(v string) *DescribeReadOnlyAccountRequest {
	s.DbName = &v
	return s
}

func (s *DescribeReadOnlyAccountRequest) SetDrdsInstanceId(v string) *DescribeReadOnlyAccountRequest {
	s.DrdsInstanceId = &v
	return s
}

type DescribeReadOnlyAccountResponseBody struct {
	Data      *DescribeReadOnlyAccountResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeReadOnlyAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeReadOnlyAccountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeReadOnlyAccountResponseBody) SetData(v *DescribeReadOnlyAccountResponseBodyData) *DescribeReadOnlyAccountResponseBody {
	s.Data = v
	return s
}

func (s *DescribeReadOnlyAccountResponseBody) SetRequestId(v string) *DescribeReadOnlyAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeReadOnlyAccountResponseBody) SetSuccess(v bool) *DescribeReadOnlyAccountResponseBody {
	s.Success = &v
	return s
}

type DescribeReadOnlyAccountResponseBodyData struct {
	AccountName    *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s DescribeReadOnlyAccountResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeReadOnlyAccountResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeReadOnlyAccountResponseBodyData) SetAccountName(v string) *DescribeReadOnlyAccountResponseBodyData {
	s.AccountName = &v
	return s
}

func (s *DescribeReadOnlyAccountResponseBodyData) SetDbName(v string) *DescribeReadOnlyAccountResponseBodyData {
	s.DbName = &v
	return s
}

func (s *DescribeReadOnlyAccountResponseBodyData) SetDrdsInstanceId(v string) *DescribeReadOnlyAccountResponseBodyData {
	s.DrdsInstanceId = &v
	return s
}

type DescribeReadOnlyAccountResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeReadOnlyAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeReadOnlyAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeReadOnlyAccountResponse) GoString() string {
	return s.String()
}

func (s *DescribeReadOnlyAccountResponse) SetHeaders(v map[string]*string) *DescribeReadOnlyAccountResponse {
	s.Headers = v
	return s
}

func (s *DescribeReadOnlyAccountResponse) SetBody(v *DescribeReadOnlyAccountResponseBody) *DescribeReadOnlyAccountResponse {
	s.Body = v
	return s
}

type DescribeRegionsResponseBody struct {
	DrdsRegions *DescribeRegionsResponseBodyDrdsRegions `json:"DrdsRegions,omitempty" xml:"DrdsRegions,omitempty" type:"Struct"`
	RequestId   *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success     *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetDrdsRegions(v *DescribeRegionsResponseBodyDrdsRegions) *DescribeRegionsResponseBody {
	s.DrdsRegions = v
	return s
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetSuccess(v bool) *DescribeRegionsResponseBody {
	s.Success = &v
	return s
}

type DescribeRegionsResponseBodyDrdsRegions struct {
	DrdsRegion []*DescribeRegionsResponseBodyDrdsRegionsDrdsRegion `json:"DrdsRegion,omitempty" xml:"DrdsRegion,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBodyDrdsRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyDrdsRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyDrdsRegions) SetDrdsRegion(v []*DescribeRegionsResponseBodyDrdsRegionsDrdsRegion) *DescribeRegionsResponseBodyDrdsRegions {
	s.DrdsRegion = v
	return s
}

type DescribeRegionsResponseBodyDrdsRegionsDrdsRegion struct {
	InstanceSeriesList *DescribeRegionsResponseBodyDrdsRegionsDrdsRegionInstanceSeriesList `json:"InstanceSeriesList,omitempty" xml:"InstanceSeriesList,omitempty" type:"Struct"`
	RegionId           *string                                                             `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RegionName         *string                                                             `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	ZoneId             *string                                                             `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	ZoneName           *string                                                             `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
}

func (s DescribeRegionsResponseBodyDrdsRegionsDrdsRegion) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyDrdsRegionsDrdsRegion) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyDrdsRegionsDrdsRegion) SetInstanceSeriesList(v *DescribeRegionsResponseBodyDrdsRegionsDrdsRegionInstanceSeriesList) *DescribeRegionsResponseBodyDrdsRegionsDrdsRegion {
	s.InstanceSeriesList = v
	return s
}

func (s *DescribeRegionsResponseBodyDrdsRegionsDrdsRegion) SetRegionId(v string) *DescribeRegionsResponseBodyDrdsRegionsDrdsRegion {
	s.RegionId = &v
	return s
}

func (s *DescribeRegionsResponseBodyDrdsRegionsDrdsRegion) SetRegionName(v string) *DescribeRegionsResponseBodyDrdsRegionsDrdsRegion {
	s.RegionName = &v
	return s
}

func (s *DescribeRegionsResponseBodyDrdsRegionsDrdsRegion) SetZoneId(v string) *DescribeRegionsResponseBodyDrdsRegionsDrdsRegion {
	s.ZoneId = &v
	return s
}

func (s *DescribeRegionsResponseBodyDrdsRegionsDrdsRegion) SetZoneName(v string) *DescribeRegionsResponseBodyDrdsRegionsDrdsRegion {
	s.ZoneName = &v
	return s
}

type DescribeRegionsResponseBodyDrdsRegionsDrdsRegionInstanceSeriesList struct {
	InstanceSeries []*DescribeRegionsResponseBodyDrdsRegionsDrdsRegionInstanceSeriesListInstanceSeries `json:"InstanceSeries,omitempty" xml:"InstanceSeries,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBodyDrdsRegionsDrdsRegionInstanceSeriesList) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyDrdsRegionsDrdsRegionInstanceSeriesList) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyDrdsRegionsDrdsRegionInstanceSeriesList) SetInstanceSeries(v []*DescribeRegionsResponseBodyDrdsRegionsDrdsRegionInstanceSeriesListInstanceSeries) *DescribeRegionsResponseBodyDrdsRegionsDrdsRegionInstanceSeriesList {
	s.InstanceSeries = v
	return s
}

type DescribeRegionsResponseBodyDrdsRegionsDrdsRegionInstanceSeriesListInstanceSeries struct {
	SeriesId   *string                                                                                   `json:"SeriesId,omitempty" xml:"SeriesId,omitempty"`
	SeriesName *string                                                                                   `json:"SeriesName,omitempty" xml:"SeriesName,omitempty"`
	SpecList   *DescribeRegionsResponseBodyDrdsRegionsDrdsRegionInstanceSeriesListInstanceSeriesSpecList `json:"SpecList,omitempty" xml:"SpecList,omitempty" type:"Struct"`
}

func (s DescribeRegionsResponseBodyDrdsRegionsDrdsRegionInstanceSeriesListInstanceSeries) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyDrdsRegionsDrdsRegionInstanceSeriesListInstanceSeries) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyDrdsRegionsDrdsRegionInstanceSeriesListInstanceSeries) SetSeriesId(v string) *DescribeRegionsResponseBodyDrdsRegionsDrdsRegionInstanceSeriesListInstanceSeries {
	s.SeriesId = &v
	return s
}

func (s *DescribeRegionsResponseBodyDrdsRegionsDrdsRegionInstanceSeriesListInstanceSeries) SetSeriesName(v string) *DescribeRegionsResponseBodyDrdsRegionsDrdsRegionInstanceSeriesListInstanceSeries {
	s.SeriesName = &v
	return s
}

func (s *DescribeRegionsResponseBodyDrdsRegionsDrdsRegionInstanceSeriesListInstanceSeries) SetSpecList(v *DescribeRegionsResponseBodyDrdsRegionsDrdsRegionInstanceSeriesListInstanceSeriesSpecList) *DescribeRegionsResponseBodyDrdsRegionsDrdsRegionInstanceSeriesListInstanceSeries {
	s.SpecList = v
	return s
}

type DescribeRegionsResponseBodyDrdsRegionsDrdsRegionInstanceSeriesListInstanceSeriesSpecList struct {
	Spec []*DescribeRegionsResponseBodyDrdsRegionsDrdsRegionInstanceSeriesListInstanceSeriesSpecListSpec `json:"Spec,omitempty" xml:"Spec,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBodyDrdsRegionsDrdsRegionInstanceSeriesListInstanceSeriesSpecList) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyDrdsRegionsDrdsRegionInstanceSeriesListInstanceSeriesSpecList) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyDrdsRegionsDrdsRegionInstanceSeriesListInstanceSeriesSpecList) SetSpec(v []*DescribeRegionsResponseBodyDrdsRegionsDrdsRegionInstanceSeriesListInstanceSeriesSpecListSpec) *DescribeRegionsResponseBodyDrdsRegionsDrdsRegionInstanceSeriesListInstanceSeriesSpecList {
	s.Spec = v
	return s
}

type DescribeRegionsResponseBodyDrdsRegionsDrdsRegionInstanceSeriesListInstanceSeriesSpecListSpec struct {
	SpecId   *string `json:"SpecId,omitempty" xml:"SpecId,omitempty"`
	SpecName *string `json:"SpecName,omitempty" xml:"SpecName,omitempty"`
}

func (s DescribeRegionsResponseBodyDrdsRegionsDrdsRegionInstanceSeriesListInstanceSeriesSpecListSpec) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyDrdsRegionsDrdsRegionInstanceSeriesListInstanceSeriesSpecListSpec) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyDrdsRegionsDrdsRegionInstanceSeriesListInstanceSeriesSpecListSpec) SetSpecId(v string) *DescribeRegionsResponseBodyDrdsRegionsDrdsRegionInstanceSeriesListInstanceSeriesSpecListSpec {
	s.SpecId = &v
	return s
}

func (s *DescribeRegionsResponseBodyDrdsRegionsDrdsRegionInstanceSeriesListInstanceSeriesSpecListSpec) SetSpecName(v string) *DescribeRegionsResponseBodyDrdsRegionsDrdsRegionInstanceSeriesListInstanceSeriesSpecListSpec {
	s.SpecName = &v
	return s
}

type DescribeRegionsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponse) SetHeaders(v map[string]*string) *DescribeRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRegionsResponse) SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse {
	s.Body = v
	return s
}

type DescribeShardDBsRequest struct {
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s DescribeShardDBsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeShardDBsRequest) GoString() string {
	return s.String()
}

func (s *DescribeShardDBsRequest) SetDbName(v string) *DescribeShardDBsRequest {
	s.DbName = &v
	return s
}

func (s *DescribeShardDBsRequest) SetDrdsInstanceId(v string) *DescribeShardDBsRequest {
	s.DrdsInstanceId = &v
	return s
}

type DescribeShardDBsResponseBody struct {
	Data      *DescribeShardDBsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeShardDBsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeShardDBsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeShardDBsResponseBody) SetData(v *DescribeShardDBsResponseBodyData) *DescribeShardDBsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeShardDBsResponseBody) SetRequestId(v string) *DescribeShardDBsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeShardDBsResponseBody) SetSuccess(v bool) *DescribeShardDBsResponseBody {
	s.Success = &v
	return s
}

type DescribeShardDBsResponseBodyData struct {
	DbIntancePair []*DescribeShardDBsResponseBodyDataDbIntancePair `json:"DbIntancePair,omitempty" xml:"DbIntancePair,omitempty" type:"Repeated"`
}

func (s DescribeShardDBsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeShardDBsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeShardDBsResponseBodyData) SetDbIntancePair(v []*DescribeShardDBsResponseBodyDataDbIntancePair) *DescribeShardDBsResponseBodyData {
	s.DbIntancePair = v
	return s
}

type DescribeShardDBsResponseBodyDataDbIntancePair struct {
	GroupName    *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	SubDbName    *string `json:"SubDbName,omitempty" xml:"SubDbName,omitempty"`
}

func (s DescribeShardDBsResponseBodyDataDbIntancePair) String() string {
	return tea.Prettify(s)
}

func (s DescribeShardDBsResponseBodyDataDbIntancePair) GoString() string {
	return s.String()
}

func (s *DescribeShardDBsResponseBodyDataDbIntancePair) SetGroupName(v string) *DescribeShardDBsResponseBodyDataDbIntancePair {
	s.GroupName = &v
	return s
}

func (s *DescribeShardDBsResponseBodyDataDbIntancePair) SetInstanceName(v string) *DescribeShardDBsResponseBodyDataDbIntancePair {
	s.InstanceName = &v
	return s
}

func (s *DescribeShardDBsResponseBodyDataDbIntancePair) SetSubDbName(v string) *DescribeShardDBsResponseBodyDataDbIntancePair {
	s.SubDbName = &v
	return s
}

type DescribeShardDBsResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeShardDBsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeShardDBsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeShardDBsResponse) GoString() string {
	return s.String()
}

func (s *DescribeShardDBsResponse) SetHeaders(v map[string]*string) *DescribeShardDBsResponse {
	s.Headers = v
	return s
}

func (s *DescribeShardDBsResponse) SetBody(v *DescribeShardDBsResponseBody) *DescribeShardDBsResponse {
	s.Body = v
	return s
}

type DescribeShardDbConnectionInfoRequest struct {
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	SubDbName      *string `json:"SubDbName,omitempty" xml:"SubDbName,omitempty"`
}

func (s DescribeShardDbConnectionInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeShardDbConnectionInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeShardDbConnectionInfoRequest) SetDbName(v string) *DescribeShardDbConnectionInfoRequest {
	s.DbName = &v
	return s
}

func (s *DescribeShardDbConnectionInfoRequest) SetDrdsInstanceId(v string) *DescribeShardDbConnectionInfoRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeShardDbConnectionInfoRequest) SetSubDbName(v string) *DescribeShardDbConnectionInfoRequest {
	s.SubDbName = &v
	return s
}

type DescribeShardDbConnectionInfoResponseBody struct {
	ConnectionInfo *DescribeShardDbConnectionInfoResponseBodyConnectionInfo `json:"ConnectionInfo,omitempty" xml:"ConnectionInfo,omitempty" type:"Struct"`
	RequestId      *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeShardDbConnectionInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeShardDbConnectionInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeShardDbConnectionInfoResponseBody) SetConnectionInfo(v *DescribeShardDbConnectionInfoResponseBodyConnectionInfo) *DescribeShardDbConnectionInfoResponseBody {
	s.ConnectionInfo = v
	return s
}

func (s *DescribeShardDbConnectionInfoResponseBody) SetRequestId(v string) *DescribeShardDbConnectionInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeShardDbConnectionInfoResponseBody) SetSuccess(v bool) *DescribeShardDbConnectionInfoResponseBody {
	s.Success = &v
	return s
}

type DescribeShardDbConnectionInfoResponseBodyConnectionInfo struct {
	InstanceName               *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	InstanceUrl                *string `json:"InstanceUrl,omitempty" xml:"InstanceUrl,omitempty"`
	BlockingTimeout            *int32  `json:"blockingTimeout,omitempty" xml:"blockingTimeout,omitempty"`
	ConnectionProperties       *string `json:"connectionProperties,omitempty" xml:"connectionProperties,omitempty"`
	DbStatus                   *string `json:"dbStatus,omitempty" xml:"dbStatus,omitempty"`
	DbType                     *string `json:"dbType,omitempty" xml:"dbType,omitempty"`
	IdleTimeOut                *int32  `json:"idleTimeOut,omitempty" xml:"idleTimeOut,omitempty"`
	MaxPoolSize                *int32  `json:"maxPoolSize,omitempty" xml:"maxPoolSize,omitempty"`
	MinPoolSize                *int32  `json:"minPoolSize,omitempty" xml:"minPoolSize,omitempty"`
	PreparedStatementCacheSize *int32  `json:"preparedStatementCacheSize,omitempty" xml:"preparedStatementCacheSize,omitempty"`
	SubDbName                  *string `json:"subDbName,omitempty" xml:"subDbName,omitempty"`
	UserName                   *string `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s DescribeShardDbConnectionInfoResponseBodyConnectionInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeShardDbConnectionInfoResponseBodyConnectionInfo) GoString() string {
	return s.String()
}

func (s *DescribeShardDbConnectionInfoResponseBodyConnectionInfo) SetInstanceName(v string) *DescribeShardDbConnectionInfoResponseBodyConnectionInfo {
	s.InstanceName = &v
	return s
}

func (s *DescribeShardDbConnectionInfoResponseBodyConnectionInfo) SetInstanceUrl(v string) *DescribeShardDbConnectionInfoResponseBodyConnectionInfo {
	s.InstanceUrl = &v
	return s
}

func (s *DescribeShardDbConnectionInfoResponseBodyConnectionInfo) SetBlockingTimeout(v int32) *DescribeShardDbConnectionInfoResponseBodyConnectionInfo {
	s.BlockingTimeout = &v
	return s
}

func (s *DescribeShardDbConnectionInfoResponseBodyConnectionInfo) SetConnectionProperties(v string) *DescribeShardDbConnectionInfoResponseBodyConnectionInfo {
	s.ConnectionProperties = &v
	return s
}

func (s *DescribeShardDbConnectionInfoResponseBodyConnectionInfo) SetDbStatus(v string) *DescribeShardDbConnectionInfoResponseBodyConnectionInfo {
	s.DbStatus = &v
	return s
}

func (s *DescribeShardDbConnectionInfoResponseBodyConnectionInfo) SetDbType(v string) *DescribeShardDbConnectionInfoResponseBodyConnectionInfo {
	s.DbType = &v
	return s
}

func (s *DescribeShardDbConnectionInfoResponseBodyConnectionInfo) SetIdleTimeOut(v int32) *DescribeShardDbConnectionInfoResponseBodyConnectionInfo {
	s.IdleTimeOut = &v
	return s
}

func (s *DescribeShardDbConnectionInfoResponseBodyConnectionInfo) SetMaxPoolSize(v int32) *DescribeShardDbConnectionInfoResponseBodyConnectionInfo {
	s.MaxPoolSize = &v
	return s
}

func (s *DescribeShardDbConnectionInfoResponseBodyConnectionInfo) SetMinPoolSize(v int32) *DescribeShardDbConnectionInfoResponseBodyConnectionInfo {
	s.MinPoolSize = &v
	return s
}

func (s *DescribeShardDbConnectionInfoResponseBodyConnectionInfo) SetPreparedStatementCacheSize(v int32) *DescribeShardDbConnectionInfoResponseBodyConnectionInfo {
	s.PreparedStatementCacheSize = &v
	return s
}

func (s *DescribeShardDbConnectionInfoResponseBodyConnectionInfo) SetSubDbName(v string) *DescribeShardDbConnectionInfoResponseBodyConnectionInfo {
	s.SubDbName = &v
	return s
}

func (s *DescribeShardDbConnectionInfoResponseBodyConnectionInfo) SetUserName(v string) *DescribeShardDbConnectionInfoResponseBodyConnectionInfo {
	s.UserName = &v
	return s
}

type DescribeShardDbConnectionInfoResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeShardDbConnectionInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeShardDbConnectionInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeShardDbConnectionInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeShardDbConnectionInfoResponse) SetHeaders(v map[string]*string) *DescribeShardDbConnectionInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeShardDbConnectionInfoResponse) SetBody(v *DescribeShardDbConnectionInfoResponseBody) *DescribeShardDbConnectionInfoResponse {
	s.Body = v
	return s
}

type EnableInstanceRequest struct {
	BackupId        *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	ClientToken     *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DbInstanceClass *string `json:"DbInstanceClass,omitempty" xml:"DbInstanceClass,omitempty"`
	DrdsInstanceId  *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	EngineVersion   *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	RestoreTime     *string `json:"RestoreTime,omitempty" xml:"RestoreTime,omitempty"`
	SourceDbInstId  *string `json:"SourceDbInstId,omitempty" xml:"SourceDbInstId,omitempty"`
	SwitchId        *string `json:"SwitchId,omitempty" xml:"SwitchId,omitempty"`
	VpcId           *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZoneId          *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s EnableInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableInstanceRequest) GoString() string {
	return s.String()
}

func (s *EnableInstanceRequest) SetBackupId(v string) *EnableInstanceRequest {
	s.BackupId = &v
	return s
}

func (s *EnableInstanceRequest) SetClientToken(v string) *EnableInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *EnableInstanceRequest) SetDbInstanceClass(v string) *EnableInstanceRequest {
	s.DbInstanceClass = &v
	return s
}

func (s *EnableInstanceRequest) SetDrdsInstanceId(v string) *EnableInstanceRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *EnableInstanceRequest) SetEngineVersion(v string) *EnableInstanceRequest {
	s.EngineVersion = &v
	return s
}

func (s *EnableInstanceRequest) SetRestoreTime(v string) *EnableInstanceRequest {
	s.RestoreTime = &v
	return s
}

func (s *EnableInstanceRequest) SetSourceDbInstId(v string) *EnableInstanceRequest {
	s.SourceDbInstId = &v
	return s
}

func (s *EnableInstanceRequest) SetSwitchId(v string) *EnableInstanceRequest {
	s.SwitchId = &v
	return s
}

func (s *EnableInstanceRequest) SetVpcId(v string) *EnableInstanceRequest {
	s.VpcId = &v
	return s
}

func (s *EnableInstanceRequest) SetZoneId(v string) *EnableInstanceRequest {
	s.ZoneId = &v
	return s
}

type EnableInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *string `json:"Result,omitempty" xml:"Result,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnableInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *EnableInstanceResponseBody) SetRequestId(v string) *EnableInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnableInstanceResponseBody) SetResult(v string) *EnableInstanceResponseBody {
	s.Result = &v
	return s
}

func (s *EnableInstanceResponseBody) SetSuccess(v bool) *EnableInstanceResponseBody {
	s.Success = &v
	return s
}

type EnableInstanceResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EnableInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableInstanceResponse) GoString() string {
	return s.String()
}

func (s *EnableInstanceResponse) SetHeaders(v map[string]*string) *EnableInstanceResponse {
	s.Headers = v
	return s
}

func (s *EnableInstanceResponse) SetBody(v *EnableInstanceResponseBody) *EnableInstanceResponse {
	s.Body = v
	return s
}

type ModifyDrdsDBPasswdRequest struct {
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	NewPasswd      *string `json:"NewPasswd,omitempty" xml:"NewPasswd,omitempty"`
}

func (s ModifyDrdsDBPasswdRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDrdsDBPasswdRequest) GoString() string {
	return s.String()
}

func (s *ModifyDrdsDBPasswdRequest) SetDbName(v string) *ModifyDrdsDBPasswdRequest {
	s.DbName = &v
	return s
}

func (s *ModifyDrdsDBPasswdRequest) SetDrdsInstanceId(v string) *ModifyDrdsDBPasswdRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *ModifyDrdsDBPasswdRequest) SetNewPasswd(v string) *ModifyDrdsDBPasswdRequest {
	s.NewPasswd = &v
	return s
}

type ModifyDrdsDBPasswdResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyDrdsDBPasswdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDrdsDBPasswdResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDrdsDBPasswdResponseBody) SetRequestId(v string) *ModifyDrdsDBPasswdResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDrdsDBPasswdResponseBody) SetSuccess(v bool) *ModifyDrdsDBPasswdResponseBody {
	s.Success = &v
	return s
}

type ModifyDrdsDBPasswdResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDrdsDBPasswdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDrdsDBPasswdResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDrdsDBPasswdResponse) GoString() string {
	return s.String()
}

func (s *ModifyDrdsDBPasswdResponse) SetHeaders(v map[string]*string) *ModifyDrdsDBPasswdResponse {
	s.Headers = v
	return s
}

func (s *ModifyDrdsDBPasswdResponse) SetBody(v *ModifyDrdsDBPasswdResponseBody) *ModifyDrdsDBPasswdResponse {
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

type ModifyFullTableScanRequest struct {
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	FullTableScan  *bool   `json:"FullTableScan,omitempty" xml:"FullTableScan,omitempty"`
	TableNames     *string `json:"TableNames,omitempty" xml:"TableNames,omitempty"`
}

func (s ModifyFullTableScanRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyFullTableScanRequest) GoString() string {
	return s.String()
}

func (s *ModifyFullTableScanRequest) SetDbName(v string) *ModifyFullTableScanRequest {
	s.DbName = &v
	return s
}

func (s *ModifyFullTableScanRequest) SetDrdsInstanceId(v string) *ModifyFullTableScanRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *ModifyFullTableScanRequest) SetFullTableScan(v bool) *ModifyFullTableScanRequest {
	s.FullTableScan = &v
	return s
}

func (s *ModifyFullTableScanRequest) SetTableNames(v string) *ModifyFullTableScanRequest {
	s.TableNames = &v
	return s
}

type ModifyFullTableScanResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyFullTableScanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyFullTableScanResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyFullTableScanResponseBody) SetRequestId(v string) *ModifyFullTableScanResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyFullTableScanResponseBody) SetSuccess(v bool) *ModifyFullTableScanResponseBody {
	s.Success = &v
	return s
}

type ModifyFullTableScanResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyFullTableScanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyFullTableScanResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyFullTableScanResponse) GoString() string {
	return s.String()
}

func (s *ModifyFullTableScanResponse) SetHeaders(v map[string]*string) *ModifyFullTableScanResponse {
	s.Headers = v
	return s
}

func (s *ModifyFullTableScanResponse) SetBody(v *ModifyFullTableScanResponseBody) *ModifyFullTableScanResponse {
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

type ModifyReadOnlyAccountPasswordRequest struct {
	AccountName    *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	NewPasswd      *string `json:"NewPasswd,omitempty" xml:"NewPasswd,omitempty"`
	OriginPassword *string `json:"OriginPassword,omitempty" xml:"OriginPassword,omitempty"`
}

func (s ModifyReadOnlyAccountPasswordRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyReadOnlyAccountPasswordRequest) GoString() string {
	return s.String()
}

func (s *ModifyReadOnlyAccountPasswordRequest) SetAccountName(v string) *ModifyReadOnlyAccountPasswordRequest {
	s.AccountName = &v
	return s
}

func (s *ModifyReadOnlyAccountPasswordRequest) SetDbName(v string) *ModifyReadOnlyAccountPasswordRequest {
	s.DbName = &v
	return s
}

func (s *ModifyReadOnlyAccountPasswordRequest) SetDrdsInstanceId(v string) *ModifyReadOnlyAccountPasswordRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *ModifyReadOnlyAccountPasswordRequest) SetNewPasswd(v string) *ModifyReadOnlyAccountPasswordRequest {
	s.NewPasswd = &v
	return s
}

func (s *ModifyReadOnlyAccountPasswordRequest) SetOriginPassword(v string) *ModifyReadOnlyAccountPasswordRequest {
	s.OriginPassword = &v
	return s
}

type ModifyReadOnlyAccountPasswordResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyReadOnlyAccountPasswordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyReadOnlyAccountPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyReadOnlyAccountPasswordResponseBody) SetRequestId(v string) *ModifyReadOnlyAccountPasswordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyReadOnlyAccountPasswordResponseBody) SetSuccess(v bool) *ModifyReadOnlyAccountPasswordResponseBody {
	s.Success = &v
	return s
}

type ModifyReadOnlyAccountPasswordResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyReadOnlyAccountPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyReadOnlyAccountPasswordResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyReadOnlyAccountPasswordResponse) GoString() string {
	return s.String()
}

func (s *ModifyReadOnlyAccountPasswordResponse) SetHeaders(v map[string]*string) *ModifyReadOnlyAccountPasswordResponse {
	s.Headers = v
	return s
}

func (s *ModifyReadOnlyAccountPasswordResponse) SetBody(v *ModifyReadOnlyAccountPasswordResponseBody) *ModifyReadOnlyAccountPasswordResponse {
	s.Body = v
	return s
}

type QueryInstanceInfoByConnRequest struct {
	Host     *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Port     *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s QueryInstanceInfoByConnRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryInstanceInfoByConnRequest) GoString() string {
	return s.String()
}

func (s *QueryInstanceInfoByConnRequest) SetHost(v string) *QueryInstanceInfoByConnRequest {
	s.Host = &v
	return s
}

func (s *QueryInstanceInfoByConnRequest) SetPort(v int32) *QueryInstanceInfoByConnRequest {
	s.Port = &v
	return s
}

func (s *QueryInstanceInfoByConnRequest) SetUserName(v string) *QueryInstanceInfoByConnRequest {
	s.UserName = &v
	return s
}

type QueryInstanceInfoByConnResponseBody struct {
	Data      *QueryInstanceInfoByConnResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryInstanceInfoByConnResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryInstanceInfoByConnResponseBody) GoString() string {
	return s.String()
}

func (s *QueryInstanceInfoByConnResponseBody) SetData(v *QueryInstanceInfoByConnResponseBodyData) *QueryInstanceInfoByConnResponseBody {
	s.Data = v
	return s
}

func (s *QueryInstanceInfoByConnResponseBody) SetRequestId(v string) *QueryInstanceInfoByConnResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryInstanceInfoByConnResponseBody) SetSuccess(v bool) *QueryInstanceInfoByConnResponseBody {
	s.Success = &v
	return s
}

type QueryInstanceInfoByConnResponseBodyData struct {
	CreateTime         *int64                                       `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description        *string                                      `json:"Description,omitempty" xml:"Description,omitempty"`
	DrdsInstanceId     *string                                      `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	NetworkType        *string                                      `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	RegionId           *string                                      `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SpecTypeId         *string                                      `json:"SpecTypeId,omitempty" xml:"SpecTypeId,omitempty"`
	SpecTypeName       *string                                      `json:"SpecTypeName,omitempty" xml:"SpecTypeName,omitempty"`
	Specification      *string                                      `json:"Specification,omitempty" xml:"Specification,omitempty"`
	Status             *string                                      `json:"Status,omitempty" xml:"Status,omitempty"`
	Type               *string                                      `json:"Type,omitempty" xml:"Type,omitempty"`
	Version            *int64                                       `json:"Version,omitempty" xml:"Version,omitempty"`
	Vips               *QueryInstanceInfoByConnResponseBodyDataVips `json:"Vips,omitempty" xml:"Vips,omitempty" type:"Struct"`
	VpcCloudInstanceId *string                                      `json:"VpcCloudInstanceId,omitempty" xml:"VpcCloudInstanceId,omitempty"`
	ZoneId             *string                                      `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s QueryInstanceInfoByConnResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryInstanceInfoByConnResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryInstanceInfoByConnResponseBodyData) SetCreateTime(v int64) *QueryInstanceInfoByConnResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *QueryInstanceInfoByConnResponseBodyData) SetDescription(v string) *QueryInstanceInfoByConnResponseBodyData {
	s.Description = &v
	return s
}

func (s *QueryInstanceInfoByConnResponseBodyData) SetDrdsInstanceId(v string) *QueryInstanceInfoByConnResponseBodyData {
	s.DrdsInstanceId = &v
	return s
}

func (s *QueryInstanceInfoByConnResponseBodyData) SetNetworkType(v string) *QueryInstanceInfoByConnResponseBodyData {
	s.NetworkType = &v
	return s
}

func (s *QueryInstanceInfoByConnResponseBodyData) SetRegionId(v string) *QueryInstanceInfoByConnResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *QueryInstanceInfoByConnResponseBodyData) SetSpecTypeId(v string) *QueryInstanceInfoByConnResponseBodyData {
	s.SpecTypeId = &v
	return s
}

func (s *QueryInstanceInfoByConnResponseBodyData) SetSpecTypeName(v string) *QueryInstanceInfoByConnResponseBodyData {
	s.SpecTypeName = &v
	return s
}

func (s *QueryInstanceInfoByConnResponseBodyData) SetSpecification(v string) *QueryInstanceInfoByConnResponseBodyData {
	s.Specification = &v
	return s
}

func (s *QueryInstanceInfoByConnResponseBodyData) SetStatus(v string) *QueryInstanceInfoByConnResponseBodyData {
	s.Status = &v
	return s
}

func (s *QueryInstanceInfoByConnResponseBodyData) SetType(v string) *QueryInstanceInfoByConnResponseBodyData {
	s.Type = &v
	return s
}

func (s *QueryInstanceInfoByConnResponseBodyData) SetVersion(v int64) *QueryInstanceInfoByConnResponseBodyData {
	s.Version = &v
	return s
}

func (s *QueryInstanceInfoByConnResponseBodyData) SetVips(v *QueryInstanceInfoByConnResponseBodyDataVips) *QueryInstanceInfoByConnResponseBodyData {
	s.Vips = v
	return s
}

func (s *QueryInstanceInfoByConnResponseBodyData) SetVpcCloudInstanceId(v string) *QueryInstanceInfoByConnResponseBodyData {
	s.VpcCloudInstanceId = &v
	return s
}

func (s *QueryInstanceInfoByConnResponseBodyData) SetZoneId(v string) *QueryInstanceInfoByConnResponseBodyData {
	s.ZoneId = &v
	return s
}

type QueryInstanceInfoByConnResponseBodyDataVips struct {
	Vip []*QueryInstanceInfoByConnResponseBodyDataVipsVip `json:"Vip,omitempty" xml:"Vip,omitempty" type:"Repeated"`
}

func (s QueryInstanceInfoByConnResponseBodyDataVips) String() string {
	return tea.Prettify(s)
}

func (s QueryInstanceInfoByConnResponseBodyDataVips) GoString() string {
	return s.String()
}

func (s *QueryInstanceInfoByConnResponseBodyDataVips) SetVip(v []*QueryInstanceInfoByConnResponseBodyDataVipsVip) *QueryInstanceInfoByConnResponseBodyDataVips {
	s.Vip = v
	return s
}

type QueryInstanceInfoByConnResponseBodyDataVipsVip struct {
	IP        *string `json:"IP,omitempty" xml:"IP,omitempty"`
	Port      *string `json:"Port,omitempty" xml:"Port,omitempty"`
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty"`
	VpcId     *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
}

func (s QueryInstanceInfoByConnResponseBodyDataVipsVip) String() string {
	return tea.Prettify(s)
}

func (s QueryInstanceInfoByConnResponseBodyDataVipsVip) GoString() string {
	return s.String()
}

func (s *QueryInstanceInfoByConnResponseBodyDataVipsVip) SetIP(v string) *QueryInstanceInfoByConnResponseBodyDataVipsVip {
	s.IP = &v
	return s
}

func (s *QueryInstanceInfoByConnResponseBodyDataVipsVip) SetPort(v string) *QueryInstanceInfoByConnResponseBodyDataVipsVip {
	s.Port = &v
	return s
}

func (s *QueryInstanceInfoByConnResponseBodyDataVipsVip) SetType(v string) *QueryInstanceInfoByConnResponseBodyDataVipsVip {
	s.Type = &v
	return s
}

func (s *QueryInstanceInfoByConnResponseBodyDataVipsVip) SetVpcId(v string) *QueryInstanceInfoByConnResponseBodyDataVipsVip {
	s.VpcId = &v
	return s
}

func (s *QueryInstanceInfoByConnResponseBodyDataVipsVip) SetVswitchId(v string) *QueryInstanceInfoByConnResponseBodyDataVipsVip {
	s.VswitchId = &v
	return s
}

type QueryInstanceInfoByConnResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryInstanceInfoByConnResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryInstanceInfoByConnResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryInstanceInfoByConnResponse) GoString() string {
	return s.String()
}

func (s *QueryInstanceInfoByConnResponse) SetHeaders(v map[string]*string) *QueryInstanceInfoByConnResponse {
	s.Headers = v
	return s
}

func (s *QueryInstanceInfoByConnResponse) SetBody(v *QueryInstanceInfoByConnResponseBody) *QueryInstanceInfoByConnResponse {
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

type RemoveReadOnlyAccountRequest struct {
	AccountName    *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s RemoveReadOnlyAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveReadOnlyAccountRequest) GoString() string {
	return s.String()
}

func (s *RemoveReadOnlyAccountRequest) SetAccountName(v string) *RemoveReadOnlyAccountRequest {
	s.AccountName = &v
	return s
}

func (s *RemoveReadOnlyAccountRequest) SetDbName(v string) *RemoveReadOnlyAccountRequest {
	s.DbName = &v
	return s
}

func (s *RemoveReadOnlyAccountRequest) SetDrdsInstanceId(v string) *RemoveReadOnlyAccountRequest {
	s.DrdsInstanceId = &v
	return s
}

type RemoveReadOnlyAccountResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RemoveReadOnlyAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveReadOnlyAccountResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveReadOnlyAccountResponseBody) SetRequestId(v string) *RemoveReadOnlyAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveReadOnlyAccountResponseBody) SetSuccess(v bool) *RemoveReadOnlyAccountResponseBody {
	s.Success = &v
	return s
}

type RemoveReadOnlyAccountResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveReadOnlyAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveReadOnlyAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveReadOnlyAccountResponse) GoString() string {
	return s.String()
}

func (s *RemoveReadOnlyAccountResponse) SetHeaders(v map[string]*string) *RemoveReadOnlyAccountResponse {
	s.Headers = v
	return s
}

func (s *RemoveReadOnlyAccountResponse) SetBody(v *RemoveReadOnlyAccountResponseBody) *RemoveReadOnlyAccountResponse {
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

func (client *Client) CreateDrdsAccountWithOptions(request *CreateDrdsAccountRequest, runtime *util.RuntimeOptions) (_result *CreateDrdsAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DbName"] = request.DbName
	query["DrdsInstanceId"] = request.DrdsInstanceId
	query["Password"] = request.Password
	query["UserName"] = request.UserName
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDrdsAccount"),
		Version:     tea.String("2017-10-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDrdsAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDrdsAccount(request *CreateDrdsAccountRequest) (_result *CreateDrdsAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDrdsAccountResponse{}
	_body, _err := client.CreateDrdsAccountWithOptions(request, runtime)
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
	query["DbName"] = request.DbName
	query["DrdsInstanceId"] = request.DrdsInstanceId
	query["Encode"] = request.Encode
	query["Password"] = request.Password
	query["RdsInstances"] = request.RdsInstances
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDrdsDB"),
		Version:     tea.String("2017-10-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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
	query["ClientToken"] = request.ClientToken
	query["Description"] = request.Description
	query["Duration"] = request.Duration
	query["InstanceSeries"] = request.InstanceSeries
	query["IsAutoRenew"] = request.IsAutoRenew
	query["IsHa"] = request.IsHa
	query["PayType"] = request.PayType
	query["PricingCycle"] = request.PricingCycle
	query["Quantity"] = request.Quantity
	query["RegionId"] = request.RegionId
	query["Specification"] = request.Specification
	query["Type"] = request.Type
	query["VpcId"] = request.VpcId
	query["VswitchId"] = request.VswitchId
	query["ZoneId"] = request.ZoneId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDrdsInstance"),
		Version:     tea.String("2017-10-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) CreateReadOnlyAccountWithOptions(request *CreateReadOnlyAccountRequest, runtime *util.RuntimeOptions) (_result *CreateReadOnlyAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DbName"] = request.DbName
	query["DrdsInstanceId"] = request.DrdsInstanceId
	query["password"] = request.Password
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateReadOnlyAccount"),
		Version:     tea.String("2017-10-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateReadOnlyAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateReadOnlyAccount(request *CreateReadOnlyAccountRequest) (_result *CreateReadOnlyAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateReadOnlyAccountResponse{}
	_body, _err := client.CreateReadOnlyAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDrdsDBWithOptions(request *DeleteDrdsDBRequest, runtime *util.RuntimeOptions) (_result *DeleteDrdsDBResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DbName"] = request.DbName
	query["DrdsInstanceId"] = request.DrdsInstanceId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDrdsDB"),
		Version:     tea.String("2017-10-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDrdsDBResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDrdsDB(request *DeleteDrdsDBRequest) (_result *DeleteDrdsDBResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDrdsDBResponse{}
	_body, _err := client.DeleteDrdsDBWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteFailedDrdsDBWithOptions(request *DeleteFailedDrdsDBRequest, runtime *util.RuntimeOptions) (_result *DeleteFailedDrdsDBResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DbName"] = request.DbName
	query["DrdsInstanceId"] = request.DrdsInstanceId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteFailedDrdsDB"),
		Version:     tea.String("2017-10-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteFailedDrdsDBResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteFailedDrdsDB(request *DeleteFailedDrdsDBRequest) (_result *DeleteFailedDrdsDBResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteFailedDrdsDBResponse{}
	_body, _err := client.DeleteFailedDrdsDBWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCreateDrdsInstanceStatusWithOptions(request *DescribeCreateDrdsInstanceStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeCreateDrdsInstanceStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DrdsInstanceId"] = request.DrdsInstanceId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCreateDrdsInstanceStatus"),
		Version:     tea.String("2017-10-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCreateDrdsInstanceStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCreateDrdsInstanceStatus(request *DescribeCreateDrdsInstanceStatusRequest) (_result *DescribeCreateDrdsInstanceStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCreateDrdsInstanceStatusResponse{}
	_body, _err := client.DescribeCreateDrdsInstanceStatusWithOptions(request, runtime)
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
	query["DbName"] = request.DbName
	query["DrdsInstanceId"] = request.DrdsInstanceId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDrdsDB"),
		Version:     tea.String("2017-10-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) DescribeDrdsDBIpWhiteListWithOptions(request *DescribeDrdsDBIpWhiteListRequest, runtime *util.RuntimeOptions) (_result *DescribeDrdsDBIpWhiteListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DbName"] = request.DbName
	query["DrdsInstanceId"] = request.DrdsInstanceId
	query["GroupName"] = request.GroupName
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDrdsDBIpWhiteList"),
		Version:     tea.String("2017-10-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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
	query["DrdsInstanceId"] = request.DrdsInstanceId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDrdsDBs"),
		Version:     tea.String("2017-10-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) DescribeDrdsInstanceWithOptions(request *DescribeDrdsInstanceRequest, runtime *util.RuntimeOptions) (_result *DescribeDrdsInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DrdsInstanceId"] = request.DrdsInstanceId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDrdsInstance"),
		Version:     tea.String("2017-10-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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
	query["DbName"] = request.DbName
	query["DrdsInstanceId"] = request.DrdsInstanceId
	query["EndTime"] = request.EndTime
	query["Key"] = request.Key
	query["StartTime"] = request.StartTime
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDrdsInstanceDbMonitor"),
		Version:     tea.String("2017-10-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) DescribeDrdsInstanceMonitorWithOptions(request *DescribeDrdsInstanceMonitorRequest, runtime *util.RuntimeOptions) (_result *DescribeDrdsInstanceMonitorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DrdsInstanceId"] = request.DrdsInstanceId
	query["EndTime"] = request.EndTime
	query["Key"] = request.Key
	query["PeriodMultiple"] = request.PeriodMultiple
	query["StartTime"] = request.StartTime
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDrdsInstanceMonitor"),
		Version:     tea.String("2017-10-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) DescribeDrdsInstanceNetInfoForInnerWithOptions(request *DescribeDrdsInstanceNetInfoForInnerRequest, runtime *util.RuntimeOptions) (_result *DescribeDrdsInstanceNetInfoForInnerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DrdsInstanceId"] = request.DrdsInstanceId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDrdsInstanceNetInfoForInner"),
		Version:     tea.String("2017-10-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDrdsInstanceNetInfoForInnerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDrdsInstanceNetInfoForInner(request *DescribeDrdsInstanceNetInfoForInnerRequest) (_result *DescribeDrdsInstanceNetInfoForInnerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDrdsInstanceNetInfoForInnerResponse{}
	_body, _err := client.DescribeDrdsInstanceNetInfoForInnerWithOptions(request, runtime)
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
	query["RegionId"] = request.RegionId
	query["Tags"] = request.Tags
	query["Type"] = request.Type
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDrdsInstances"),
		Version:     tea.String("2017-10-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) DescribeRdsListWithOptions(request *DescribeRdsListRequest, runtime *util.RuntimeOptions) (_result *DescribeRdsListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DbName"] = request.DbName
	query["DrdsInstanceId"] = request.DrdsInstanceId
	query["RegionId"] = request.RegionId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRdsList"),
		Version:     tea.String("2017-10-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRdsListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRdsList(request *DescribeRdsListRequest) (_result *DescribeRdsListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRdsListResponse{}
	_body, _err := client.DescribeRdsListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeReadOnlyAccountWithOptions(request *DescribeReadOnlyAccountRequest, runtime *util.RuntimeOptions) (_result *DescribeReadOnlyAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DbName"] = request.DbName
	query["DrdsInstanceId"] = request.DrdsInstanceId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeReadOnlyAccount"),
		Version:     tea.String("2017-10-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeReadOnlyAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeReadOnlyAccount(request *DescribeReadOnlyAccountRequest) (_result *DescribeReadOnlyAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeReadOnlyAccountResponse{}
	_body, _err := client.DescribeReadOnlyAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRegionsWithOptions(runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("DescribeRegions"),
		Version:     tea.String("2017-10-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRegions() (_result *DescribeRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeShardDBsWithOptions(request *DescribeShardDBsRequest, runtime *util.RuntimeOptions) (_result *DescribeShardDBsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DbName"] = request.DbName
	query["DrdsInstanceId"] = request.DrdsInstanceId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeShardDBs"),
		Version:     tea.String("2017-10-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeShardDBsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeShardDBs(request *DescribeShardDBsRequest) (_result *DescribeShardDBsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeShardDBsResponse{}
	_body, _err := client.DescribeShardDBsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeShardDbConnectionInfoWithOptions(request *DescribeShardDbConnectionInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeShardDbConnectionInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DbName"] = request.DbName
	query["DrdsInstanceId"] = request.DrdsInstanceId
	query["SubDbName"] = request.SubDbName
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeShardDbConnectionInfo"),
		Version:     tea.String("2017-10-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeShardDbConnectionInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeShardDbConnectionInfo(request *DescribeShardDbConnectionInfoRequest) (_result *DescribeShardDbConnectionInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeShardDbConnectionInfoResponse{}
	_body, _err := client.DescribeShardDbConnectionInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableInstanceWithOptions(request *EnableInstanceRequest, runtime *util.RuntimeOptions) (_result *EnableInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["BackupId"] = request.BackupId
	query["ClientToken"] = request.ClientToken
	query["DbInstanceClass"] = request.DbInstanceClass
	query["DrdsInstanceId"] = request.DrdsInstanceId
	query["EngineVersion"] = request.EngineVersion
	query["RestoreTime"] = request.RestoreTime
	query["SourceDbInstId"] = request.SourceDbInstId
	query["SwitchId"] = request.SwitchId
	query["VpcId"] = request.VpcId
	query["ZoneId"] = request.ZoneId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableInstance"),
		Version:     tea.String("2017-10-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableInstance(request *EnableInstanceRequest) (_result *EnableInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableInstanceResponse{}
	_body, _err := client.EnableInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDrdsDBPasswdWithOptions(request *ModifyDrdsDBPasswdRequest, runtime *util.RuntimeOptions) (_result *ModifyDrdsDBPasswdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DbName"] = request.DbName
	query["DrdsInstanceId"] = request.DrdsInstanceId
	query["NewPasswd"] = request.NewPasswd
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDrdsDBPasswd"),
		Version:     tea.String("2017-10-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDrdsDBPasswdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDrdsDBPasswd(request *ModifyDrdsDBPasswdRequest) (_result *ModifyDrdsDBPasswdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDrdsDBPasswdResponse{}
	_body, _err := client.ModifyDrdsDBPasswdWithOptions(request, runtime)
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
	query["Description"] = request.Description
	query["DrdsInstanceId"] = request.DrdsInstanceId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDrdsInstanceDescription"),
		Version:     tea.String("2017-10-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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
	query["DbName"] = request.DbName
	query["DrdsInstanceId"] = request.DrdsInstanceId
	query["GroupAttribute"] = request.GroupAttribute
	query["GroupName"] = request.GroupName
	query["IpWhiteList"] = request.IpWhiteList
	query["Mode"] = request.Mode
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDrdsIpWhiteList"),
		Version:     tea.String("2017-10-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) ModifyFullTableScanWithOptions(request *ModifyFullTableScanRequest, runtime *util.RuntimeOptions) (_result *ModifyFullTableScanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DbName"] = request.DbName
	query["DrdsInstanceId"] = request.DrdsInstanceId
	query["FullTableScan"] = request.FullTableScan
	query["TableNames"] = request.TableNames
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyFullTableScan"),
		Version:     tea.String("2017-10-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyFullTableScanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyFullTableScan(request *ModifyFullTableScanRequest) (_result *ModifyFullTableScanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyFullTableScanResponse{}
	_body, _err := client.ModifyFullTableScanWithOptions(request, runtime)
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
	query["DbName"] = request.DbName
	query["DrdsInstanceId"] = request.DrdsInstanceId
	query["InstanceNames"] = request.InstanceNames
	query["Weights"] = request.Weights
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyRdsReadWeight"),
		Version:     tea.String("2017-10-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) ModifyReadOnlyAccountPasswordWithOptions(request *ModifyReadOnlyAccountPasswordRequest, runtime *util.RuntimeOptions) (_result *ModifyReadOnlyAccountPasswordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["AccountName"] = request.AccountName
	query["DbName"] = request.DbName
	query["DrdsInstanceId"] = request.DrdsInstanceId
	query["NewPasswd"] = request.NewPasswd
	query["OriginPassword"] = request.OriginPassword
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyReadOnlyAccountPassword"),
		Version:     tea.String("2017-10-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyReadOnlyAccountPasswordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyReadOnlyAccountPassword(request *ModifyReadOnlyAccountPasswordRequest) (_result *ModifyReadOnlyAccountPasswordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyReadOnlyAccountPasswordResponse{}
	_body, _err := client.ModifyReadOnlyAccountPasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryInstanceInfoByConnWithOptions(request *QueryInstanceInfoByConnRequest, runtime *util.RuntimeOptions) (_result *QueryInstanceInfoByConnResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["Host"] = request.Host
	query["Port"] = request.Port
	query["UserName"] = request.UserName
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryInstanceInfoByConn"),
		Version:     tea.String("2017-10-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryInstanceInfoByConnResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryInstanceInfoByConn(request *QueryInstanceInfoByConnRequest) (_result *QueryInstanceInfoByConnResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryInstanceInfoByConnResponse{}
	_body, _err := client.QueryInstanceInfoByConnWithOptions(request, runtime)
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
	query["DrdsInstanceId"] = request.DrdsInstanceId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveDrdsInstance"),
		Version:     tea.String("2017-10-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) RemoveReadOnlyAccountWithOptions(request *RemoveReadOnlyAccountRequest, runtime *util.RuntimeOptions) (_result *RemoveReadOnlyAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["AccountName"] = request.AccountName
	query["DbName"] = request.DbName
	query["DrdsInstanceId"] = request.DrdsInstanceId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveReadOnlyAccount"),
		Version:     tea.String("2017-10-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveReadOnlyAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveReadOnlyAccount(request *RemoveReadOnlyAccountRequest) (_result *RemoveReadOnlyAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveReadOnlyAccountResponse{}
	_body, _err := client.RemoveReadOnlyAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
