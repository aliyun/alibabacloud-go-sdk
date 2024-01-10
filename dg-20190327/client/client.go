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

type AddDatabaseRequest struct {
	ClientToken   *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DbDescription *string `json:"DbDescription,omitempty" xml:"DbDescription,omitempty"`
	DbName        *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DbPassword    *string `json:"DbPassword,omitempty" xml:"DbPassword,omitempty"`
	DbType        *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	DbUserName    *string `json:"DbUserName,omitempty" xml:"DbUserName,omitempty"`
	GatewayId     *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	Host          *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Port          *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AddDatabaseRequest) String() string {
	return tea.Prettify(s)
}

func (s AddDatabaseRequest) GoString() string {
	return s.String()
}

func (s *AddDatabaseRequest) SetClientToken(v string) *AddDatabaseRequest {
	s.ClientToken = &v
	return s
}

func (s *AddDatabaseRequest) SetDbDescription(v string) *AddDatabaseRequest {
	s.DbDescription = &v
	return s
}

func (s *AddDatabaseRequest) SetDbName(v string) *AddDatabaseRequest {
	s.DbName = &v
	return s
}

func (s *AddDatabaseRequest) SetDbPassword(v string) *AddDatabaseRequest {
	s.DbPassword = &v
	return s
}

func (s *AddDatabaseRequest) SetDbType(v string) *AddDatabaseRequest {
	s.DbType = &v
	return s
}

func (s *AddDatabaseRequest) SetDbUserName(v string) *AddDatabaseRequest {
	s.DbUserName = &v
	return s
}

func (s *AddDatabaseRequest) SetGatewayId(v string) *AddDatabaseRequest {
	s.GatewayId = &v
	return s
}

func (s *AddDatabaseRequest) SetHost(v string) *AddDatabaseRequest {
	s.Host = &v
	return s
}

func (s *AddDatabaseRequest) SetPort(v int32) *AddDatabaseRequest {
	s.Port = &v
	return s
}

func (s *AddDatabaseRequest) SetRegionId(v string) *AddDatabaseRequest {
	s.RegionId = &v
	return s
}

type AddDatabaseResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddDatabaseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddDatabaseResponseBody) GoString() string {
	return s.String()
}

func (s *AddDatabaseResponseBody) SetCode(v string) *AddDatabaseResponseBody {
	s.Code = &v
	return s
}

func (s *AddDatabaseResponseBody) SetData(v string) *AddDatabaseResponseBody {
	s.Data = &v
	return s
}

func (s *AddDatabaseResponseBody) SetErrorMsg(v string) *AddDatabaseResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *AddDatabaseResponseBody) SetRequestId(v string) *AddDatabaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddDatabaseResponseBody) SetSuccess(v bool) *AddDatabaseResponseBody {
	s.Success = &v
	return s
}

type AddDatabaseResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddDatabaseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddDatabaseResponse) String() string {
	return tea.Prettify(s)
}

func (s AddDatabaseResponse) GoString() string {
	return s.String()
}

func (s *AddDatabaseResponse) SetHeaders(v map[string]*string) *AddDatabaseResponse {
	s.Headers = v
	return s
}

func (s *AddDatabaseResponse) SetStatusCode(v int32) *AddDatabaseResponse {
	s.StatusCode = &v
	return s
}

func (s *AddDatabaseResponse) SetBody(v *AddDatabaseResponseBody) *AddDatabaseResponse {
	s.Body = v
	return s
}

type AddDatabaseListRequest struct {
	ClientToken    *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DatabaseString *string `json:"DatabaseString,omitempty" xml:"DatabaseString,omitempty"`
}

func (s AddDatabaseListRequest) String() string {
	return tea.Prettify(s)
}

func (s AddDatabaseListRequest) GoString() string {
	return s.String()
}

func (s *AddDatabaseListRequest) SetClientToken(v string) *AddDatabaseListRequest {
	s.ClientToken = &v
	return s
}

func (s *AddDatabaseListRequest) SetDatabaseString(v string) *AddDatabaseListRequest {
	s.DatabaseString = &v
	return s
}

type AddDatabaseListResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddDatabaseListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddDatabaseListResponseBody) GoString() string {
	return s.String()
}

func (s *AddDatabaseListResponseBody) SetCode(v string) *AddDatabaseListResponseBody {
	s.Code = &v
	return s
}

func (s *AddDatabaseListResponseBody) SetData(v string) *AddDatabaseListResponseBody {
	s.Data = &v
	return s
}

func (s *AddDatabaseListResponseBody) SetErrorMsg(v string) *AddDatabaseListResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *AddDatabaseListResponseBody) SetRequestId(v string) *AddDatabaseListResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddDatabaseListResponseBody) SetSuccess(v bool) *AddDatabaseListResponseBody {
	s.Success = &v
	return s
}

type AddDatabaseListResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddDatabaseListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddDatabaseListResponse) String() string {
	return tea.Prettify(s)
}

func (s AddDatabaseListResponse) GoString() string {
	return s.String()
}

func (s *AddDatabaseListResponse) SetHeaders(v map[string]*string) *AddDatabaseListResponse {
	s.Headers = v
	return s
}

func (s *AddDatabaseListResponse) SetStatusCode(v int32) *AddDatabaseListResponse {
	s.StatusCode = &v
	return s
}

func (s *AddDatabaseListResponse) SetBody(v *AddDatabaseListResponseBody) *AddDatabaseListResponse {
	s.Body = v
	return s
}

type ConnectDatabaseRequest struct {
	DbName     *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DbPassword *string `json:"DbPassword,omitempty" xml:"DbPassword,omitempty"`
	DbType     *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	DbUserName *string `json:"DbUserName,omitempty" xml:"DbUserName,omitempty"`
	GatewayId  *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	Host       *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Port       *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s ConnectDatabaseRequest) String() string {
	return tea.Prettify(s)
}

func (s ConnectDatabaseRequest) GoString() string {
	return s.String()
}

func (s *ConnectDatabaseRequest) SetDbName(v string) *ConnectDatabaseRequest {
	s.DbName = &v
	return s
}

func (s *ConnectDatabaseRequest) SetDbPassword(v string) *ConnectDatabaseRequest {
	s.DbPassword = &v
	return s
}

func (s *ConnectDatabaseRequest) SetDbType(v string) *ConnectDatabaseRequest {
	s.DbType = &v
	return s
}

func (s *ConnectDatabaseRequest) SetDbUserName(v string) *ConnectDatabaseRequest {
	s.DbUserName = &v
	return s
}

func (s *ConnectDatabaseRequest) SetGatewayId(v string) *ConnectDatabaseRequest {
	s.GatewayId = &v
	return s
}

func (s *ConnectDatabaseRequest) SetHost(v string) *ConnectDatabaseRequest {
	s.Host = &v
	return s
}

func (s *ConnectDatabaseRequest) SetPort(v int32) *ConnectDatabaseRequest {
	s.Port = &v
	return s
}

type ConnectDatabaseResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ConnectDatabaseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConnectDatabaseResponseBody) GoString() string {
	return s.String()
}

func (s *ConnectDatabaseResponseBody) SetCode(v string) *ConnectDatabaseResponseBody {
	s.Code = &v
	return s
}

func (s *ConnectDatabaseResponseBody) SetData(v string) *ConnectDatabaseResponseBody {
	s.Data = &v
	return s
}

func (s *ConnectDatabaseResponseBody) SetErrorMsg(v string) *ConnectDatabaseResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ConnectDatabaseResponseBody) SetRequestId(v string) *ConnectDatabaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConnectDatabaseResponseBody) SetSuccess(v bool) *ConnectDatabaseResponseBody {
	s.Success = &v
	return s
}

type ConnectDatabaseResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ConnectDatabaseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ConnectDatabaseResponse) String() string {
	return tea.Prettify(s)
}

func (s ConnectDatabaseResponse) GoString() string {
	return s.String()
}

func (s *ConnectDatabaseResponse) SetHeaders(v map[string]*string) *ConnectDatabaseResponse {
	s.Headers = v
	return s
}

func (s *ConnectDatabaseResponse) SetStatusCode(v int32) *ConnectDatabaseResponse {
	s.StatusCode = &v
	return s
}

func (s *ConnectDatabaseResponse) SetBody(v *ConnectDatabaseResponseBody) *ConnectDatabaseResponse {
	s.Body = v
	return s
}

type CreateDatabaseAccessPointRequest struct {
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DbInstanceId *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	VSwitchId    *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcAZone     *string `json:"VpcAZone,omitempty" xml:"VpcAZone,omitempty"`
	VpcId        *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VpcRegionId  *string `json:"VpcRegionId,omitempty" xml:"VpcRegionId,omitempty"`
}

func (s CreateDatabaseAccessPointRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDatabaseAccessPointRequest) GoString() string {
	return s.String()
}

func (s *CreateDatabaseAccessPointRequest) SetClientToken(v string) *CreateDatabaseAccessPointRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDatabaseAccessPointRequest) SetDbInstanceId(v string) *CreateDatabaseAccessPointRequest {
	s.DbInstanceId = &v
	return s
}

func (s *CreateDatabaseAccessPointRequest) SetRegionId(v string) *CreateDatabaseAccessPointRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDatabaseAccessPointRequest) SetVSwitchId(v string) *CreateDatabaseAccessPointRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateDatabaseAccessPointRequest) SetVpcAZone(v string) *CreateDatabaseAccessPointRequest {
	s.VpcAZone = &v
	return s
}

func (s *CreateDatabaseAccessPointRequest) SetVpcId(v string) *CreateDatabaseAccessPointRequest {
	s.VpcId = &v
	return s
}

func (s *CreateDatabaseAccessPointRequest) SetVpcRegionId(v string) *CreateDatabaseAccessPointRequest {
	s.VpcRegionId = &v
	return s
}

type CreateDatabaseAccessPointResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDatabaseAccessPointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDatabaseAccessPointResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDatabaseAccessPointResponseBody) SetCode(v string) *CreateDatabaseAccessPointResponseBody {
	s.Code = &v
	return s
}

func (s *CreateDatabaseAccessPointResponseBody) SetData(v string) *CreateDatabaseAccessPointResponseBody {
	s.Data = &v
	return s
}

func (s *CreateDatabaseAccessPointResponseBody) SetErrorMsg(v string) *CreateDatabaseAccessPointResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *CreateDatabaseAccessPointResponseBody) SetRequestId(v string) *CreateDatabaseAccessPointResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDatabaseAccessPointResponseBody) SetSuccess(v bool) *CreateDatabaseAccessPointResponseBody {
	s.Success = &v
	return s
}

type CreateDatabaseAccessPointResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateDatabaseAccessPointResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDatabaseAccessPointResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDatabaseAccessPointResponse) GoString() string {
	return s.String()
}

func (s *CreateDatabaseAccessPointResponse) SetHeaders(v map[string]*string) *CreateDatabaseAccessPointResponse {
	s.Headers = v
	return s
}

func (s *CreateDatabaseAccessPointResponse) SetStatusCode(v int32) *CreateDatabaseAccessPointResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDatabaseAccessPointResponse) SetBody(v *CreateDatabaseAccessPointResponseBody) *CreateDatabaseAccessPointResponse {
	s.Body = v
	return s
}

type CreateGatewayRequest struct {
	GatewayDesc *string `json:"GatewayDesc,omitempty" xml:"GatewayDesc,omitempty"`
	GatewayName *string `json:"GatewayName,omitempty" xml:"GatewayName,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateGatewayRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGatewayRequest) GoString() string {
	return s.String()
}

func (s *CreateGatewayRequest) SetGatewayDesc(v string) *CreateGatewayRequest {
	s.GatewayDesc = &v
	return s
}

func (s *CreateGatewayRequest) SetGatewayName(v string) *CreateGatewayRequest {
	s.GatewayName = &v
	return s
}

func (s *CreateGatewayRequest) SetRegionId(v string) *CreateGatewayRequest {
	s.RegionId = &v
	return s
}

type CreateGatewayResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateGatewayResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGatewayResponseBody) SetCode(v string) *CreateGatewayResponseBody {
	s.Code = &v
	return s
}

func (s *CreateGatewayResponseBody) SetData(v string) *CreateGatewayResponseBody {
	s.Data = &v
	return s
}

func (s *CreateGatewayResponseBody) SetErrorMsg(v string) *CreateGatewayResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *CreateGatewayResponseBody) SetRequestId(v string) *CreateGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGatewayResponseBody) SetSuccess(v bool) *CreateGatewayResponseBody {
	s.Success = &v
	return s
}

type CreateGatewayResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateGatewayResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGatewayResponse) GoString() string {
	return s.String()
}

func (s *CreateGatewayResponse) SetHeaders(v map[string]*string) *CreateGatewayResponse {
	s.Headers = v
	return s
}

func (s *CreateGatewayResponse) SetStatusCode(v int32) *CreateGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGatewayResponse) SetBody(v *CreateGatewayResponseBody) *CreateGatewayResponse {
	s.Body = v
	return s
}

type CreateGatewayVerifyCodeRequest struct {
	GatewayId *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
}

func (s CreateGatewayVerifyCodeRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGatewayVerifyCodeRequest) GoString() string {
	return s.String()
}

func (s *CreateGatewayVerifyCodeRequest) SetGatewayId(v string) *CreateGatewayVerifyCodeRequest {
	s.GatewayId = &v
	return s
}

type CreateGatewayVerifyCodeResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateGatewayVerifyCodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGatewayVerifyCodeResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGatewayVerifyCodeResponseBody) SetCode(v string) *CreateGatewayVerifyCodeResponseBody {
	s.Code = &v
	return s
}

func (s *CreateGatewayVerifyCodeResponseBody) SetData(v string) *CreateGatewayVerifyCodeResponseBody {
	s.Data = &v
	return s
}

func (s *CreateGatewayVerifyCodeResponseBody) SetErrorMsg(v string) *CreateGatewayVerifyCodeResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *CreateGatewayVerifyCodeResponseBody) SetRequestId(v string) *CreateGatewayVerifyCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGatewayVerifyCodeResponseBody) SetSuccess(v bool) *CreateGatewayVerifyCodeResponseBody {
	s.Success = &v
	return s
}

type CreateGatewayVerifyCodeResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateGatewayVerifyCodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateGatewayVerifyCodeResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGatewayVerifyCodeResponse) GoString() string {
	return s.String()
}

func (s *CreateGatewayVerifyCodeResponse) SetHeaders(v map[string]*string) *CreateGatewayVerifyCodeResponse {
	s.Headers = v
	return s
}

func (s *CreateGatewayVerifyCodeResponse) SetStatusCode(v int32) *CreateGatewayVerifyCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGatewayVerifyCodeResponse) SetBody(v *CreateGatewayVerifyCodeResponseBody) *CreateGatewayVerifyCodeResponse {
	s.Body = v
	return s
}

type DeleteDatabaseRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteDatabaseRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDatabaseRequest) GoString() string {
	return s.String()
}

func (s *DeleteDatabaseRequest) SetInstanceId(v string) *DeleteDatabaseRequest {
	s.InstanceId = &v
	return s
}

type DeleteDatabaseResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteDatabaseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDatabaseResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDatabaseResponseBody) SetCode(v string) *DeleteDatabaseResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteDatabaseResponseBody) SetData(v string) *DeleteDatabaseResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteDatabaseResponseBody) SetErrorMsg(v string) *DeleteDatabaseResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *DeleteDatabaseResponseBody) SetRequestId(v string) *DeleteDatabaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDatabaseResponseBody) SetSuccess(v bool) *DeleteDatabaseResponseBody {
	s.Success = &v
	return s
}

type DeleteDatabaseResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteDatabaseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDatabaseResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDatabaseResponse) GoString() string {
	return s.String()
}

func (s *DeleteDatabaseResponse) SetHeaders(v map[string]*string) *DeleteDatabaseResponse {
	s.Headers = v
	return s
}

func (s *DeleteDatabaseResponse) SetStatusCode(v int32) *DeleteDatabaseResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDatabaseResponse) SetBody(v *DeleteDatabaseResponseBody) *DeleteDatabaseResponse {
	s.Body = v
	return s
}

type DeleteDatabaseAccessPointRequest struct {
	DbInstanceId *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	VSwitchId    *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcAZone     *string `json:"VpcAZone,omitempty" xml:"VpcAZone,omitempty"`
	VpcId        *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VpcRegionId  *string `json:"VpcRegionId,omitempty" xml:"VpcRegionId,omitempty"`
}

func (s DeleteDatabaseAccessPointRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDatabaseAccessPointRequest) GoString() string {
	return s.String()
}

func (s *DeleteDatabaseAccessPointRequest) SetDbInstanceId(v string) *DeleteDatabaseAccessPointRequest {
	s.DbInstanceId = &v
	return s
}

func (s *DeleteDatabaseAccessPointRequest) SetRegionId(v string) *DeleteDatabaseAccessPointRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteDatabaseAccessPointRequest) SetVSwitchId(v string) *DeleteDatabaseAccessPointRequest {
	s.VSwitchId = &v
	return s
}

func (s *DeleteDatabaseAccessPointRequest) SetVpcAZone(v string) *DeleteDatabaseAccessPointRequest {
	s.VpcAZone = &v
	return s
}

func (s *DeleteDatabaseAccessPointRequest) SetVpcId(v string) *DeleteDatabaseAccessPointRequest {
	s.VpcId = &v
	return s
}

func (s *DeleteDatabaseAccessPointRequest) SetVpcRegionId(v string) *DeleteDatabaseAccessPointRequest {
	s.VpcRegionId = &v
	return s
}

type DeleteDatabaseAccessPointResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteDatabaseAccessPointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDatabaseAccessPointResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDatabaseAccessPointResponseBody) SetCode(v string) *DeleteDatabaseAccessPointResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteDatabaseAccessPointResponseBody) SetData(v string) *DeleteDatabaseAccessPointResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteDatabaseAccessPointResponseBody) SetErrorMsg(v string) *DeleteDatabaseAccessPointResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *DeleteDatabaseAccessPointResponseBody) SetRequestId(v string) *DeleteDatabaseAccessPointResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDatabaseAccessPointResponseBody) SetSuccess(v bool) *DeleteDatabaseAccessPointResponseBody {
	s.Success = &v
	return s
}

type DeleteDatabaseAccessPointResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteDatabaseAccessPointResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDatabaseAccessPointResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDatabaseAccessPointResponse) GoString() string {
	return s.String()
}

func (s *DeleteDatabaseAccessPointResponse) SetHeaders(v map[string]*string) *DeleteDatabaseAccessPointResponse {
	s.Headers = v
	return s
}

func (s *DeleteDatabaseAccessPointResponse) SetStatusCode(v int32) *DeleteDatabaseAccessPointResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDatabaseAccessPointResponse) SetBody(v *DeleteDatabaseAccessPointResponseBody) *DeleteDatabaseAccessPointResponse {
	s.Body = v
	return s
}

type DeleteGatewayRequest struct {
	GatewayId *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
}

func (s DeleteGatewayRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteGatewayRequest) GoString() string {
	return s.String()
}

func (s *DeleteGatewayRequest) SetGatewayId(v string) *DeleteGatewayRequest {
	s.GatewayId = &v
	return s
}

type DeleteGatewayResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteGatewayResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGatewayResponseBody) SetCode(v string) *DeleteGatewayResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteGatewayResponseBody) SetData(v string) *DeleteGatewayResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteGatewayResponseBody) SetErrorMsg(v string) *DeleteGatewayResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *DeleteGatewayResponseBody) SetRequestId(v string) *DeleteGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteGatewayResponseBody) SetSuccess(v bool) *DeleteGatewayResponseBody {
	s.Success = &v
	return s
}

type DeleteGatewayResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteGatewayResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteGatewayResponse) GoString() string {
	return s.String()
}

func (s *DeleteGatewayResponse) SetHeaders(v map[string]*string) *DeleteGatewayResponse {
	s.Headers = v
	return s
}

func (s *DeleteGatewayResponse) SetStatusCode(v int32) *DeleteGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGatewayResponse) SetBody(v *DeleteGatewayResponseBody) *DeleteGatewayResponse {
	s.Body = v
	return s
}

type DeleteGatewayInstanceRequest struct {
	GatewayId         *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	GatewayInstanceId *string `json:"GatewayInstanceId,omitempty" xml:"GatewayInstanceId,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteGatewayInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteGatewayInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteGatewayInstanceRequest) SetGatewayId(v string) *DeleteGatewayInstanceRequest {
	s.GatewayId = &v
	return s
}

func (s *DeleteGatewayInstanceRequest) SetGatewayInstanceId(v string) *DeleteGatewayInstanceRequest {
	s.GatewayInstanceId = &v
	return s
}

func (s *DeleteGatewayInstanceRequest) SetRegionId(v string) *DeleteGatewayInstanceRequest {
	s.RegionId = &v
	return s
}

type DeleteGatewayInstanceResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteGatewayInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteGatewayInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGatewayInstanceResponseBody) SetCode(v string) *DeleteGatewayInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteGatewayInstanceResponseBody) SetData(v string) *DeleteGatewayInstanceResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteGatewayInstanceResponseBody) SetErrorMsg(v string) *DeleteGatewayInstanceResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *DeleteGatewayInstanceResponseBody) SetRequestId(v string) *DeleteGatewayInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteGatewayInstanceResponseBody) SetSuccess(v bool) *DeleteGatewayInstanceResponseBody {
	s.Success = &v
	return s
}

type DeleteGatewayInstanceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteGatewayInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteGatewayInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteGatewayInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeleteGatewayInstanceResponse) SetHeaders(v map[string]*string) *DeleteGatewayInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeleteGatewayInstanceResponse) SetStatusCode(v int32) *DeleteGatewayInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGatewayInstanceResponse) SetBody(v *DeleteGatewayInstanceResponseBody) *DeleteGatewayInstanceResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) SetRegionId(v string) *DescribeRegionsRequest {
	s.RegionId = &v
	return s
}

type DescribeRegionsResponseBody struct {
	Code     *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// Region
	Regions   *DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetCode(v string) *DescribeRegionsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetErrorMsg(v string) *DescribeRegionsResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetRegions(v *DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
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

type DescribeRegionsResponseBodyRegions struct {
	Region []*DescribeRegionsResponseBodyRegionsRegion `json:"Region,omitempty" xml:"Region,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) SetRegion(v []*DescribeRegionsResponseBodyRegionsRegion) *DescribeRegionsResponseBodyRegions {
	s.Region = v
	return s
}

type DescribeRegionsResponseBodyRegionsRegion struct {
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// Region Endpoint。
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsResponseBodyRegionsRegion) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsRegion) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetLocalName(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.LocalName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetRegionEndpoint(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.RegionEndpoint = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetRegionId(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.RegionId = &v
	return s
}

type DescribeRegionsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeRegionsResponse) SetStatusCode(v int32) *DescribeRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRegionsResponse) SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse {
	s.Body = v
	return s
}

type DownloadGatewayProgramRequest struct {
	DgVersion *string `json:"DgVersion,omitempty" xml:"DgVersion,omitempty"`
	UserOS    *string `json:"UserOS,omitempty" xml:"UserOS,omitempty"`
}

func (s DownloadGatewayProgramRequest) String() string {
	return tea.Prettify(s)
}

func (s DownloadGatewayProgramRequest) GoString() string {
	return s.String()
}

func (s *DownloadGatewayProgramRequest) SetDgVersion(v string) *DownloadGatewayProgramRequest {
	s.DgVersion = &v
	return s
}

func (s *DownloadGatewayProgramRequest) SetUserOS(v string) *DownloadGatewayProgramRequest {
	s.UserOS = &v
	return s
}

type DownloadGatewayProgramResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DownloadGatewayProgramResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DownloadGatewayProgramResponseBody) GoString() string {
	return s.String()
}

func (s *DownloadGatewayProgramResponseBody) SetCode(v string) *DownloadGatewayProgramResponseBody {
	s.Code = &v
	return s
}

func (s *DownloadGatewayProgramResponseBody) SetData(v string) *DownloadGatewayProgramResponseBody {
	s.Data = &v
	return s
}

func (s *DownloadGatewayProgramResponseBody) SetErrorMsg(v string) *DownloadGatewayProgramResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *DownloadGatewayProgramResponseBody) SetRequestId(v string) *DownloadGatewayProgramResponseBody {
	s.RequestId = &v
	return s
}

func (s *DownloadGatewayProgramResponseBody) SetSuccess(v bool) *DownloadGatewayProgramResponseBody {
	s.Success = &v
	return s
}

type DownloadGatewayProgramResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DownloadGatewayProgramResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DownloadGatewayProgramResponse) String() string {
	return tea.Prettify(s)
}

func (s DownloadGatewayProgramResponse) GoString() string {
	return s.String()
}

func (s *DownloadGatewayProgramResponse) SetHeaders(v map[string]*string) *DownloadGatewayProgramResponse {
	s.Headers = v
	return s
}

func (s *DownloadGatewayProgramResponse) SetStatusCode(v int32) *DownloadGatewayProgramResponse {
	s.StatusCode = &v
	return s
}

func (s *DownloadGatewayProgramResponse) SetBody(v *DownloadGatewayProgramResponseBody) *DownloadGatewayProgramResponse {
	s.Body = v
	return s
}

type FindUserGatewayByIdRequest struct {
	GatewayId *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
}

func (s FindUserGatewayByIdRequest) String() string {
	return tea.Prettify(s)
}

func (s FindUserGatewayByIdRequest) GoString() string {
	return s.String()
}

func (s *FindUserGatewayByIdRequest) SetGatewayId(v string) *FindUserGatewayByIdRequest {
	s.GatewayId = &v
	return s
}

type FindUserGatewayByIdResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s FindUserGatewayByIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FindUserGatewayByIdResponseBody) GoString() string {
	return s.String()
}

func (s *FindUserGatewayByIdResponseBody) SetCode(v string) *FindUserGatewayByIdResponseBody {
	s.Code = &v
	return s
}

func (s *FindUserGatewayByIdResponseBody) SetData(v string) *FindUserGatewayByIdResponseBody {
	s.Data = &v
	return s
}

func (s *FindUserGatewayByIdResponseBody) SetErrorMsg(v string) *FindUserGatewayByIdResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *FindUserGatewayByIdResponseBody) SetRequestId(v string) *FindUserGatewayByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *FindUserGatewayByIdResponseBody) SetSuccess(v bool) *FindUserGatewayByIdResponseBody {
	s.Success = &v
	return s
}

type FindUserGatewayByIdResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *FindUserGatewayByIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FindUserGatewayByIdResponse) String() string {
	return tea.Prettify(s)
}

func (s FindUserGatewayByIdResponse) GoString() string {
	return s.String()
}

func (s *FindUserGatewayByIdResponse) SetHeaders(v map[string]*string) *FindUserGatewayByIdResponse {
	s.Headers = v
	return s
}

func (s *FindUserGatewayByIdResponse) SetStatusCode(v int32) *FindUserGatewayByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *FindUserGatewayByIdResponse) SetBody(v *FindUserGatewayByIdResponseBody) *FindUserGatewayByIdResponse {
	s.Body = v
	return s
}

type GetUserDatabasesRequest struct {
	DbType     *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	GatewayId  *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	Host       *string `json:"Host,omitempty" xml:"Host,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Port       *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SearchKey  *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
}

func (s GetUserDatabasesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserDatabasesRequest) GoString() string {
	return s.String()
}

func (s *GetUserDatabasesRequest) SetDbType(v string) *GetUserDatabasesRequest {
	s.DbType = &v
	return s
}

func (s *GetUserDatabasesRequest) SetGatewayId(v string) *GetUserDatabasesRequest {
	s.GatewayId = &v
	return s
}

func (s *GetUserDatabasesRequest) SetHost(v string) *GetUserDatabasesRequest {
	s.Host = &v
	return s
}

func (s *GetUserDatabasesRequest) SetInstanceId(v string) *GetUserDatabasesRequest {
	s.InstanceId = &v
	return s
}

func (s *GetUserDatabasesRequest) SetPageNumber(v string) *GetUserDatabasesRequest {
	s.PageNumber = &v
	return s
}

func (s *GetUserDatabasesRequest) SetPageSize(v string) *GetUserDatabasesRequest {
	s.PageSize = &v
	return s
}

func (s *GetUserDatabasesRequest) SetPort(v int32) *GetUserDatabasesRequest {
	s.Port = &v
	return s
}

func (s *GetUserDatabasesRequest) SetRegionId(v string) *GetUserDatabasesRequest {
	s.RegionId = &v
	return s
}

func (s *GetUserDatabasesRequest) SetSearchKey(v string) *GetUserDatabasesRequest {
	s.SearchKey = &v
	return s
}

type GetUserDatabasesResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Count     *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetUserDatabasesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserDatabasesResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserDatabasesResponseBody) SetCode(v string) *GetUserDatabasesResponseBody {
	s.Code = &v
	return s
}

func (s *GetUserDatabasesResponseBody) SetCount(v int32) *GetUserDatabasesResponseBody {
	s.Count = &v
	return s
}

func (s *GetUserDatabasesResponseBody) SetData(v string) *GetUserDatabasesResponseBody {
	s.Data = &v
	return s
}

func (s *GetUserDatabasesResponseBody) SetErrorMsg(v string) *GetUserDatabasesResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetUserDatabasesResponseBody) SetRequestId(v string) *GetUserDatabasesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserDatabasesResponseBody) SetSuccess(v string) *GetUserDatabasesResponseBody {
	s.Success = &v
	return s
}

type GetUserDatabasesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetUserDatabasesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUserDatabasesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserDatabasesResponse) GoString() string {
	return s.String()
}

func (s *GetUserDatabasesResponse) SetHeaders(v map[string]*string) *GetUserDatabasesResponse {
	s.Headers = v
	return s
}

func (s *GetUserDatabasesResponse) SetStatusCode(v int32) *GetUserDatabasesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserDatabasesResponse) SetBody(v *GetUserDatabasesResponseBody) *GetUserDatabasesResponse {
	s.Body = v
	return s
}

type GetUserGatewayInstancesRequest struct {
	GatewayId *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
}

func (s GetUserGatewayInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserGatewayInstancesRequest) GoString() string {
	return s.String()
}

func (s *GetUserGatewayInstancesRequest) SetGatewayId(v string) *GetUserGatewayInstancesRequest {
	s.GatewayId = &v
	return s
}

type GetUserGatewayInstancesResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetUserGatewayInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserGatewayInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserGatewayInstancesResponseBody) SetCode(v string) *GetUserGatewayInstancesResponseBody {
	s.Code = &v
	return s
}

func (s *GetUserGatewayInstancesResponseBody) SetData(v string) *GetUserGatewayInstancesResponseBody {
	s.Data = &v
	return s
}

func (s *GetUserGatewayInstancesResponseBody) SetErrorMsg(v string) *GetUserGatewayInstancesResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetUserGatewayInstancesResponseBody) SetRequestId(v string) *GetUserGatewayInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserGatewayInstancesResponseBody) SetSuccess(v bool) *GetUserGatewayInstancesResponseBody {
	s.Success = &v
	return s
}

type GetUserGatewayInstancesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetUserGatewayInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUserGatewayInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserGatewayInstancesResponse) GoString() string {
	return s.String()
}

func (s *GetUserGatewayInstancesResponse) SetHeaders(v map[string]*string) *GetUserGatewayInstancesResponse {
	s.Headers = v
	return s
}

func (s *GetUserGatewayInstancesResponse) SetStatusCode(v int32) *GetUserGatewayInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserGatewayInstancesResponse) SetBody(v *GetUserGatewayInstancesResponseBody) *GetUserGatewayInstancesResponse {
	s.Body = v
	return s
}

type GetUserGatewaysRequest struct {
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SearchKey  *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
}

func (s GetUserGatewaysRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserGatewaysRequest) GoString() string {
	return s.String()
}

func (s *GetUserGatewaysRequest) SetPageNumber(v int32) *GetUserGatewaysRequest {
	s.PageNumber = &v
	return s
}

func (s *GetUserGatewaysRequest) SetPageSize(v int32) *GetUserGatewaysRequest {
	s.PageSize = &v
	return s
}

func (s *GetUserGatewaysRequest) SetRegionId(v string) *GetUserGatewaysRequest {
	s.RegionId = &v
	return s
}

func (s *GetUserGatewaysRequest) SetSearchKey(v string) *GetUserGatewaysRequest {
	s.SearchKey = &v
	return s
}

type GetUserGatewaysResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Count     *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetUserGatewaysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserGatewaysResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserGatewaysResponseBody) SetCode(v string) *GetUserGatewaysResponseBody {
	s.Code = &v
	return s
}

func (s *GetUserGatewaysResponseBody) SetCount(v int32) *GetUserGatewaysResponseBody {
	s.Count = &v
	return s
}

func (s *GetUserGatewaysResponseBody) SetData(v string) *GetUserGatewaysResponseBody {
	s.Data = &v
	return s
}

func (s *GetUserGatewaysResponseBody) SetErrorMsg(v string) *GetUserGatewaysResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetUserGatewaysResponseBody) SetRequestId(v string) *GetUserGatewaysResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserGatewaysResponseBody) SetSuccess(v bool) *GetUserGatewaysResponseBody {
	s.Success = &v
	return s
}

type GetUserGatewaysResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetUserGatewaysResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUserGatewaysResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserGatewaysResponse) GoString() string {
	return s.String()
}

func (s *GetUserGatewaysResponse) SetHeaders(v map[string]*string) *GetUserGatewaysResponse {
	s.Headers = v
	return s
}

func (s *GetUserGatewaysResponse) SetStatusCode(v int32) *GetUserGatewaysResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserGatewaysResponse) SetBody(v *GetUserGatewaysResponseBody) *GetUserGatewaysResponse {
	s.Body = v
	return s
}

type ListDatabaseAccessPointRequest struct {
	DbInstanceId *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
	GatewayId    *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	Host         *string `json:"Host,omitempty" xml:"Host,omitempty"`
	PageNumber   *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Port         *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SearchKey    *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	VpcId        *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListDatabaseAccessPointRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDatabaseAccessPointRequest) GoString() string {
	return s.String()
}

func (s *ListDatabaseAccessPointRequest) SetDbInstanceId(v string) *ListDatabaseAccessPointRequest {
	s.DbInstanceId = &v
	return s
}

func (s *ListDatabaseAccessPointRequest) SetGatewayId(v string) *ListDatabaseAccessPointRequest {
	s.GatewayId = &v
	return s
}

func (s *ListDatabaseAccessPointRequest) SetHost(v string) *ListDatabaseAccessPointRequest {
	s.Host = &v
	return s
}

func (s *ListDatabaseAccessPointRequest) SetPageNumber(v string) *ListDatabaseAccessPointRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDatabaseAccessPointRequest) SetPageSize(v string) *ListDatabaseAccessPointRequest {
	s.PageSize = &v
	return s
}

func (s *ListDatabaseAccessPointRequest) SetPort(v int32) *ListDatabaseAccessPointRequest {
	s.Port = &v
	return s
}

func (s *ListDatabaseAccessPointRequest) SetRegionId(v string) *ListDatabaseAccessPointRequest {
	s.RegionId = &v
	return s
}

func (s *ListDatabaseAccessPointRequest) SetSearchKey(v string) *ListDatabaseAccessPointRequest {
	s.SearchKey = &v
	return s
}

func (s *ListDatabaseAccessPointRequest) SetVpcId(v string) *ListDatabaseAccessPointRequest {
	s.VpcId = &v
	return s
}

type ListDatabaseAccessPointResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Count     *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDatabaseAccessPointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDatabaseAccessPointResponseBody) GoString() string {
	return s.String()
}

func (s *ListDatabaseAccessPointResponseBody) SetCode(v string) *ListDatabaseAccessPointResponseBody {
	s.Code = &v
	return s
}

func (s *ListDatabaseAccessPointResponseBody) SetCount(v int32) *ListDatabaseAccessPointResponseBody {
	s.Count = &v
	return s
}

func (s *ListDatabaseAccessPointResponseBody) SetData(v string) *ListDatabaseAccessPointResponseBody {
	s.Data = &v
	return s
}

func (s *ListDatabaseAccessPointResponseBody) SetErrorMsg(v string) *ListDatabaseAccessPointResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListDatabaseAccessPointResponseBody) SetRequestId(v string) *ListDatabaseAccessPointResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDatabaseAccessPointResponseBody) SetSuccess(v string) *ListDatabaseAccessPointResponseBody {
	s.Success = &v
	return s
}

type ListDatabaseAccessPointResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDatabaseAccessPointResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDatabaseAccessPointResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDatabaseAccessPointResponse) GoString() string {
	return s.String()
}

func (s *ListDatabaseAccessPointResponse) SetHeaders(v map[string]*string) *ListDatabaseAccessPointResponse {
	s.Headers = v
	return s
}

func (s *ListDatabaseAccessPointResponse) SetStatusCode(v int32) *ListDatabaseAccessPointResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDatabaseAccessPointResponse) SetBody(v *ListDatabaseAccessPointResponseBody) *ListDatabaseAccessPointResponse {
	s.Body = v
	return s
}

type ModifyDatabaseRequest struct {
	DbDescription *string `json:"DbDescription,omitempty" xml:"DbDescription,omitempty"`
	DbName        *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DbPassword    *string `json:"DbPassword,omitempty" xml:"DbPassword,omitempty"`
	DbType        *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	DbUserName    *string `json:"DbUserName,omitempty" xml:"DbUserName,omitempty"`
	Host          *string `json:"Host,omitempty" xml:"Host,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Port          *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s ModifyDatabaseRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDatabaseRequest) GoString() string {
	return s.String()
}

func (s *ModifyDatabaseRequest) SetDbDescription(v string) *ModifyDatabaseRequest {
	s.DbDescription = &v
	return s
}

func (s *ModifyDatabaseRequest) SetDbName(v string) *ModifyDatabaseRequest {
	s.DbName = &v
	return s
}

func (s *ModifyDatabaseRequest) SetDbPassword(v string) *ModifyDatabaseRequest {
	s.DbPassword = &v
	return s
}

func (s *ModifyDatabaseRequest) SetDbType(v string) *ModifyDatabaseRequest {
	s.DbType = &v
	return s
}

func (s *ModifyDatabaseRequest) SetDbUserName(v string) *ModifyDatabaseRequest {
	s.DbUserName = &v
	return s
}

func (s *ModifyDatabaseRequest) SetHost(v string) *ModifyDatabaseRequest {
	s.Host = &v
	return s
}

func (s *ModifyDatabaseRequest) SetInstanceId(v string) *ModifyDatabaseRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyDatabaseRequest) SetPort(v int32) *ModifyDatabaseRequest {
	s.Port = &v
	return s
}

type ModifyDatabaseResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyDatabaseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDatabaseResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDatabaseResponseBody) SetCode(v string) *ModifyDatabaseResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyDatabaseResponseBody) SetData(v string) *ModifyDatabaseResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyDatabaseResponseBody) SetErrorMsg(v string) *ModifyDatabaseResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ModifyDatabaseResponseBody) SetRequestId(v string) *ModifyDatabaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDatabaseResponseBody) SetSuccess(v bool) *ModifyDatabaseResponseBody {
	s.Success = &v
	return s
}

type ModifyDatabaseResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyDatabaseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDatabaseResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDatabaseResponse) GoString() string {
	return s.String()
}

func (s *ModifyDatabaseResponse) SetHeaders(v map[string]*string) *ModifyDatabaseResponse {
	s.Headers = v
	return s
}

func (s *ModifyDatabaseResponse) SetStatusCode(v int32) *ModifyDatabaseResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDatabaseResponse) SetBody(v *ModifyDatabaseResponseBody) *ModifyDatabaseResponse {
	s.Body = v
	return s
}

type ModifyGatewayRequest struct {
	GatewayDesc *string `json:"GatewayDesc,omitempty" xml:"GatewayDesc,omitempty"`
	GatewayId   *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	GatewayName *string `json:"GatewayName,omitempty" xml:"GatewayName,omitempty"`
}

func (s ModifyGatewayRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyGatewayRequest) GoString() string {
	return s.String()
}

func (s *ModifyGatewayRequest) SetGatewayDesc(v string) *ModifyGatewayRequest {
	s.GatewayDesc = &v
	return s
}

func (s *ModifyGatewayRequest) SetGatewayId(v string) *ModifyGatewayRequest {
	s.GatewayId = &v
	return s
}

func (s *ModifyGatewayRequest) SetGatewayName(v string) *ModifyGatewayRequest {
	s.GatewayName = &v
	return s
}

type ModifyGatewayResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyGatewayResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyGatewayResponseBody) SetCode(v string) *ModifyGatewayResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyGatewayResponseBody) SetData(v string) *ModifyGatewayResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyGatewayResponseBody) SetErrorMsg(v string) *ModifyGatewayResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ModifyGatewayResponseBody) SetRequestId(v string) *ModifyGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyGatewayResponseBody) SetSuccess(v bool) *ModifyGatewayResponseBody {
	s.Success = &v
	return s
}

type ModifyGatewayResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyGatewayResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyGatewayResponse) GoString() string {
	return s.String()
}

func (s *ModifyGatewayResponse) SetHeaders(v map[string]*string) *ModifyGatewayResponse {
	s.Headers = v
	return s
}

func (s *ModifyGatewayResponse) SetStatusCode(v int32) *ModifyGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyGatewayResponse) SetBody(v *ModifyGatewayResponseBody) *ModifyGatewayResponse {
	s.Body = v
	return s
}

type RetryDatabaseRequest struct {
	ClientToken   *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DbDescription *string `json:"DbDescription,omitempty" xml:"DbDescription,omitempty"`
	DbName        *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DbPassword    *string `json:"DbPassword,omitempty" xml:"DbPassword,omitempty"`
	DbType        *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	DbUserName    *string `json:"DbUserName,omitempty" xml:"DbUserName,omitempty"`
	GatewayId     *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	Host          *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Port          *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RetryDatabaseRequest) String() string {
	return tea.Prettify(s)
}

func (s RetryDatabaseRequest) GoString() string {
	return s.String()
}

func (s *RetryDatabaseRequest) SetClientToken(v string) *RetryDatabaseRequest {
	s.ClientToken = &v
	return s
}

func (s *RetryDatabaseRequest) SetDbDescription(v string) *RetryDatabaseRequest {
	s.DbDescription = &v
	return s
}

func (s *RetryDatabaseRequest) SetDbName(v string) *RetryDatabaseRequest {
	s.DbName = &v
	return s
}

func (s *RetryDatabaseRequest) SetDbPassword(v string) *RetryDatabaseRequest {
	s.DbPassword = &v
	return s
}

func (s *RetryDatabaseRequest) SetDbType(v string) *RetryDatabaseRequest {
	s.DbType = &v
	return s
}

func (s *RetryDatabaseRequest) SetDbUserName(v string) *RetryDatabaseRequest {
	s.DbUserName = &v
	return s
}

func (s *RetryDatabaseRequest) SetGatewayId(v string) *RetryDatabaseRequest {
	s.GatewayId = &v
	return s
}

func (s *RetryDatabaseRequest) SetHost(v string) *RetryDatabaseRequest {
	s.Host = &v
	return s
}

func (s *RetryDatabaseRequest) SetPort(v int32) *RetryDatabaseRequest {
	s.Port = &v
	return s
}

func (s *RetryDatabaseRequest) SetRegionId(v string) *RetryDatabaseRequest {
	s.RegionId = &v
	return s
}

type RetryDatabaseResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RetryDatabaseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RetryDatabaseResponseBody) GoString() string {
	return s.String()
}

func (s *RetryDatabaseResponseBody) SetCode(v string) *RetryDatabaseResponseBody {
	s.Code = &v
	return s
}

func (s *RetryDatabaseResponseBody) SetData(v string) *RetryDatabaseResponseBody {
	s.Data = &v
	return s
}

func (s *RetryDatabaseResponseBody) SetErrorMsg(v string) *RetryDatabaseResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *RetryDatabaseResponseBody) SetRequestId(v string) *RetryDatabaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *RetryDatabaseResponseBody) SetSuccess(v bool) *RetryDatabaseResponseBody {
	s.Success = &v
	return s
}

type RetryDatabaseResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RetryDatabaseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RetryDatabaseResponse) String() string {
	return tea.Prettify(s)
}

func (s RetryDatabaseResponse) GoString() string {
	return s.String()
}

func (s *RetryDatabaseResponse) SetHeaders(v map[string]*string) *RetryDatabaseResponse {
	s.Headers = v
	return s
}

func (s *RetryDatabaseResponse) SetStatusCode(v int32) *RetryDatabaseResponse {
	s.StatusCode = &v
	return s
}

func (s *RetryDatabaseResponse) SetBody(v *RetryDatabaseResponseBody) *RetryDatabaseResponse {
	s.Body = v
	return s
}

type StopGatewayRequest struct {
	GatewayId         *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	GatewayInstanceId *string `json:"GatewayInstanceId,omitempty" xml:"GatewayInstanceId,omitempty"`
}

func (s StopGatewayRequest) String() string {
	return tea.Prettify(s)
}

func (s StopGatewayRequest) GoString() string {
	return s.String()
}

func (s *StopGatewayRequest) SetGatewayId(v string) *StopGatewayRequest {
	s.GatewayId = &v
	return s
}

func (s *StopGatewayRequest) SetGatewayInstanceId(v string) *StopGatewayRequest {
	s.GatewayInstanceId = &v
	return s
}

type StopGatewayResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StopGatewayResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *StopGatewayResponseBody) SetCode(v string) *StopGatewayResponseBody {
	s.Code = &v
	return s
}

func (s *StopGatewayResponseBody) SetData(v string) *StopGatewayResponseBody {
	s.Data = &v
	return s
}

func (s *StopGatewayResponseBody) SetErrorMsg(v string) *StopGatewayResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *StopGatewayResponseBody) SetRequestId(v string) *StopGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopGatewayResponseBody) SetSuccess(v bool) *StopGatewayResponseBody {
	s.Success = &v
	return s
}

type StopGatewayResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StopGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopGatewayResponse) String() string {
	return tea.Prettify(s)
}

func (s StopGatewayResponse) GoString() string {
	return s.String()
}

func (s *StopGatewayResponse) SetHeaders(v map[string]*string) *StopGatewayResponse {
	s.Headers = v
	return s
}

func (s *StopGatewayResponse) SetStatusCode(v int32) *StopGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *StopGatewayResponse) SetBody(v *StopGatewayResponseBody) *StopGatewayResponse {
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
		"ap-northeast-1":              tea.String("dg.aliyuncs.com"),
		"ap-northeast-2-pop":          tea.String("dg.aliyuncs.com"),
		"ap-south-1":                  tea.String("dg.aliyuncs.com"),
		"ap-southeast-1":              tea.String("dg.aliyuncs.com"),
		"ap-southeast-2":              tea.String("dg.aliyuncs.com"),
		"ap-southeast-3":              tea.String("dg.aliyuncs.com"),
		"ap-southeast-5":              tea.String("dg.aliyuncs.com"),
		"cn-beijing":                  tea.String("dg.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("dg.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("dg.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("dg.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("dg.aliyuncs.com"),
		"cn-chengdu":                  tea.String("dg.aliyuncs.com"),
		"cn-edge-1":                   tea.String("dg.aliyuncs.com"),
		"cn-fujian":                   tea.String("dg.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("dg.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("dg.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("dg.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("dg.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("dg.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("dg.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("dg.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("dg.aliyuncs.com"),
		"cn-hongkong":                 tea.String("dg.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("dg.aliyuncs.com"),
		"cn-huhehaote":                tea.String("dg.aliyuncs.com"),
		"cn-north-2-gov-1":            tea.String("dg.aliyuncs.com"),
		"cn-qingdao":                  tea.String("dg.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("dg.aliyuncs.com"),
		"cn-shanghai":                 tea.String("dg.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("dg.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("dg.aliyuncs.com"),
		"cn-shanghai-finance-1":       tea.String("dg.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("dg.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("dg.aliyuncs.com"),
		"cn-shenzhen":                 tea.String("dg.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("dg.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("dg.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("dg.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("dg.aliyuncs.com"),
		"cn-wuhan":                    tea.String("dg.aliyuncs.com"),
		"cn-yushanfang":               tea.String("dg.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("dg.aliyuncs.com"),
		"cn-zhangjiakou":              tea.String("dg.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("dg.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("dg.aliyuncs.com"),
		"eu-central-1":                tea.String("dg.aliyuncs.com"),
		"eu-west-1":                   tea.String("dg.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("dg.aliyuncs.com"),
		"me-east-1":                   tea.String("dg.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("dg.aliyuncs.com"),
		"us-east-1":                   tea.String("dg.aliyuncs.com"),
		"us-west-1":                   tea.String("dg.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("dg"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AddDatabaseWithOptions(request *AddDatabaseRequest, runtime *util.RuntimeOptions) (_result *AddDatabaseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DbDescription)) {
		body["DbDescription"] = request.DbDescription
	}

	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		body["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DbPassword)) {
		body["DbPassword"] = request.DbPassword
	}

	if !tea.BoolValue(util.IsUnset(request.DbType)) {
		body["DbType"] = request.DbType
	}

	if !tea.BoolValue(util.IsUnset(request.DbUserName)) {
		body["DbUserName"] = request.DbUserName
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		body["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.Host)) {
		body["Host"] = request.Host
	}

	if !tea.BoolValue(util.IsUnset(request.Port)) {
		body["Port"] = request.Port
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddDatabase"),
		Version:     tea.String("2019-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddDatabaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddDatabase(request *AddDatabaseRequest) (_result *AddDatabaseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddDatabaseResponse{}
	_body, _err := client.AddDatabaseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddDatabaseListWithOptions(request *AddDatabaseListRequest, runtime *util.RuntimeOptions) (_result *AddDatabaseListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseString)) {
		body["DatabaseString"] = request.DatabaseString
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddDatabaseList"),
		Version:     tea.String("2019-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddDatabaseListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddDatabaseList(request *AddDatabaseListRequest) (_result *AddDatabaseListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddDatabaseListResponse{}
	_body, _err := client.AddDatabaseListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ConnectDatabaseWithOptions(request *ConnectDatabaseRequest, runtime *util.RuntimeOptions) (_result *ConnectDatabaseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		body["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DbPassword)) {
		body["DbPassword"] = request.DbPassword
	}

	if !tea.BoolValue(util.IsUnset(request.DbType)) {
		body["DbType"] = request.DbType
	}

	if !tea.BoolValue(util.IsUnset(request.DbUserName)) {
		body["DbUserName"] = request.DbUserName
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		body["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.Host)) {
		body["Host"] = request.Host
	}

	if !tea.BoolValue(util.IsUnset(request.Port)) {
		body["Port"] = request.Port
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ConnectDatabase"),
		Version:     tea.String("2019-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ConnectDatabaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConnectDatabase(request *ConnectDatabaseRequest) (_result *ConnectDatabaseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConnectDatabaseResponse{}
	_body, _err := client.ConnectDatabaseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDatabaseAccessPointWithOptions(request *CreateDatabaseAccessPointRequest, runtime *util.RuntimeOptions) (_result *CreateDatabaseAccessPointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DbInstanceId)) {
		body["DbInstanceId"] = request.DbInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		body["VSwitchId"] = request.VSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcAZone)) {
		body["VpcAZone"] = request.VpcAZone
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		body["VpcId"] = request.VpcId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcRegionId)) {
		body["VpcRegionId"] = request.VpcRegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDatabaseAccessPoint"),
		Version:     tea.String("2019-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDatabaseAccessPointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDatabaseAccessPoint(request *CreateDatabaseAccessPointRequest) (_result *CreateDatabaseAccessPointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDatabaseAccessPointResponse{}
	_body, _err := client.CreateDatabaseAccessPointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateGatewayWithOptions(request *CreateGatewayRequest, runtime *util.RuntimeOptions) (_result *CreateGatewayResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayDesc)) {
		body["GatewayDesc"] = request.GatewayDesc
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayName)) {
		body["GatewayName"] = request.GatewayName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateGateway"),
		Version:     tea.String("2019-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateGatewayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateGateway(request *CreateGatewayRequest) (_result *CreateGatewayResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateGatewayResponse{}
	_body, _err := client.CreateGatewayWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateGatewayVerifyCodeWithOptions(request *CreateGatewayVerifyCodeRequest, runtime *util.RuntimeOptions) (_result *CreateGatewayVerifyCodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		body["GatewayId"] = request.GatewayId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateGatewayVerifyCode"),
		Version:     tea.String("2019-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateGatewayVerifyCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateGatewayVerifyCode(request *CreateGatewayVerifyCodeRequest) (_result *CreateGatewayVerifyCodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateGatewayVerifyCodeResponse{}
	_body, _err := client.CreateGatewayVerifyCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDatabaseWithOptions(request *DeleteDatabaseRequest, runtime *util.RuntimeOptions) (_result *DeleteDatabaseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDatabase"),
		Version:     tea.String("2019-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDatabaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDatabase(request *DeleteDatabaseRequest) (_result *DeleteDatabaseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDatabaseResponse{}
	_body, _err := client.DeleteDatabaseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDatabaseAccessPointWithOptions(request *DeleteDatabaseAccessPointRequest, runtime *util.RuntimeOptions) (_result *DeleteDatabaseAccessPointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbInstanceId)) {
		body["DbInstanceId"] = request.DbInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		body["VSwitchId"] = request.VSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcAZone)) {
		body["VpcAZone"] = request.VpcAZone
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		body["VpcId"] = request.VpcId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcRegionId)) {
		body["VpcRegionId"] = request.VpcRegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDatabaseAccessPoint"),
		Version:     tea.String("2019-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDatabaseAccessPointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDatabaseAccessPoint(request *DeleteDatabaseAccessPointRequest) (_result *DeleteDatabaseAccessPointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDatabaseAccessPointResponse{}
	_body, _err := client.DeleteDatabaseAccessPointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteGatewayWithOptions(request *DeleteGatewayRequest, runtime *util.RuntimeOptions) (_result *DeleteGatewayResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		body["GatewayId"] = request.GatewayId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteGateway"),
		Version:     tea.String("2019-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteGatewayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteGateway(request *DeleteGatewayRequest) (_result *DeleteGatewayResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteGatewayResponse{}
	_body, _err := client.DeleteGatewayWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteGatewayInstanceWithOptions(request *DeleteGatewayInstanceRequest, runtime *util.RuntimeOptions) (_result *DeleteGatewayInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		body["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayInstanceId)) {
		body["GatewayInstanceId"] = request.GatewayInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteGatewayInstance"),
		Version:     tea.String("2019-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteGatewayInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteGatewayInstance(request *DeleteGatewayInstanceRequest) (_result *DeleteGatewayInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteGatewayInstanceResponse{}
	_body, _err := client.DeleteGatewayInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRegions"),
		Version:     tea.String("2019-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

func (client *Client) DescribeRegions(request *DescribeRegionsRequest) (_result *DescribeRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DownloadGatewayProgramWithOptions(request *DownloadGatewayProgramRequest, runtime *util.RuntimeOptions) (_result *DownloadGatewayProgramResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DgVersion)) {
		body["DgVersion"] = request.DgVersion
	}

	if !tea.BoolValue(util.IsUnset(request.UserOS)) {
		body["UserOS"] = request.UserOS
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DownloadGatewayProgram"),
		Version:     tea.String("2019-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DownloadGatewayProgramResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DownloadGatewayProgram(request *DownloadGatewayProgramRequest) (_result *DownloadGatewayProgramResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DownloadGatewayProgramResponse{}
	_body, _err := client.DownloadGatewayProgramWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FindUserGatewayByIdWithOptions(request *FindUserGatewayByIdRequest, runtime *util.RuntimeOptions) (_result *FindUserGatewayByIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		body["GatewayId"] = request.GatewayId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("FindUserGatewayById"),
		Version:     tea.String("2019-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &FindUserGatewayByIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FindUserGatewayById(request *FindUserGatewayByIdRequest) (_result *FindUserGatewayByIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FindUserGatewayByIdResponse{}
	_body, _err := client.FindUserGatewayByIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUserDatabasesWithOptions(request *GetUserDatabasesRequest, runtime *util.RuntimeOptions) (_result *GetUserDatabasesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbType)) {
		body["DbType"] = request.DbType
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		body["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.Host)) {
		body["Host"] = request.Host
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Port)) {
		body["Port"] = request.Port
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		body["SearchKey"] = request.SearchKey
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUserDatabases"),
		Version:     tea.String("2019-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserDatabasesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUserDatabases(request *GetUserDatabasesRequest) (_result *GetUserDatabasesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUserDatabasesResponse{}
	_body, _err := client.GetUserDatabasesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUserGatewayInstancesWithOptions(request *GetUserGatewayInstancesRequest, runtime *util.RuntimeOptions) (_result *GetUserGatewayInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		body["GatewayId"] = request.GatewayId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUserGatewayInstances"),
		Version:     tea.String("2019-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserGatewayInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUserGatewayInstances(request *GetUserGatewayInstancesRequest) (_result *GetUserGatewayInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUserGatewayInstancesResponse{}
	_body, _err := client.GetUserGatewayInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUserGatewaysWithOptions(request *GetUserGatewaysRequest, runtime *util.RuntimeOptions) (_result *GetUserGatewaysResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		body["SearchKey"] = request.SearchKey
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUserGateways"),
		Version:     tea.String("2019-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserGatewaysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUserGateways(request *GetUserGatewaysRequest) (_result *GetUserGatewaysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUserGatewaysResponse{}
	_body, _err := client.GetUserGatewaysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDatabaseAccessPointWithOptions(request *ListDatabaseAccessPointRequest, runtime *util.RuntimeOptions) (_result *ListDatabaseAccessPointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbInstanceId)) {
		body["DbInstanceId"] = request.DbInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		body["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.Host)) {
		body["Host"] = request.Host
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Port)) {
		body["Port"] = request.Port
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		body["SearchKey"] = request.SearchKey
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		body["VpcId"] = request.VpcId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDatabaseAccessPoint"),
		Version:     tea.String("2019-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDatabaseAccessPointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDatabaseAccessPoint(request *ListDatabaseAccessPointRequest) (_result *ListDatabaseAccessPointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDatabaseAccessPointResponse{}
	_body, _err := client.ListDatabaseAccessPointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDatabaseWithOptions(request *ModifyDatabaseRequest, runtime *util.RuntimeOptions) (_result *ModifyDatabaseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbDescription)) {
		body["DbDescription"] = request.DbDescription
	}

	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		body["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DbPassword)) {
		body["DbPassword"] = request.DbPassword
	}

	if !tea.BoolValue(util.IsUnset(request.DbType)) {
		body["DbType"] = request.DbType
	}

	if !tea.BoolValue(util.IsUnset(request.DbUserName)) {
		body["DbUserName"] = request.DbUserName
	}

	if !tea.BoolValue(util.IsUnset(request.Host)) {
		body["Host"] = request.Host
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Port)) {
		body["Port"] = request.Port
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDatabase"),
		Version:     tea.String("2019-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDatabaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDatabase(request *ModifyDatabaseRequest) (_result *ModifyDatabaseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDatabaseResponse{}
	_body, _err := client.ModifyDatabaseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyGatewayWithOptions(request *ModifyGatewayRequest, runtime *util.RuntimeOptions) (_result *ModifyGatewayResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayDesc)) {
		body["GatewayDesc"] = request.GatewayDesc
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		body["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayName)) {
		body["GatewayName"] = request.GatewayName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyGateway"),
		Version:     tea.String("2019-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyGatewayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyGateway(request *ModifyGatewayRequest) (_result *ModifyGatewayResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyGatewayResponse{}
	_body, _err := client.ModifyGatewayWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RetryDatabaseWithOptions(request *RetryDatabaseRequest, runtime *util.RuntimeOptions) (_result *RetryDatabaseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DbDescription)) {
		body["DbDescription"] = request.DbDescription
	}

	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		body["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DbPassword)) {
		body["DbPassword"] = request.DbPassword
	}

	if !tea.BoolValue(util.IsUnset(request.DbType)) {
		body["DbType"] = request.DbType
	}

	if !tea.BoolValue(util.IsUnset(request.DbUserName)) {
		body["DbUserName"] = request.DbUserName
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		body["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.Host)) {
		body["Host"] = request.Host
	}

	if !tea.BoolValue(util.IsUnset(request.Port)) {
		body["Port"] = request.Port
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RetryDatabase"),
		Version:     tea.String("2019-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RetryDatabaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RetryDatabase(request *RetryDatabaseRequest) (_result *RetryDatabaseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RetryDatabaseResponse{}
	_body, _err := client.RetryDatabaseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopGatewayWithOptions(request *StopGatewayRequest, runtime *util.RuntimeOptions) (_result *StopGatewayResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		body["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayInstanceId)) {
		body["GatewayInstanceId"] = request.GatewayInstanceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("StopGateway"),
		Version:     tea.String("2019-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopGatewayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopGateway(request *StopGatewayRequest) (_result *StopGatewayResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopGatewayResponse{}
	_body, _err := client.StopGatewayWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
