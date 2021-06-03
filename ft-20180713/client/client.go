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

type BatchAuditTest01Request struct {
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	BatchAuditTest01 *string `json:"BatchAuditTest01,omitempty" xml:"BatchAuditTest01,omitempty"`
	Demo01           *string `json:"Demo01,omitempty" xml:"Demo01,omitempty"`
	Test010101       *bool   `json:"Test010101,omitempty" xml:"Test010101,omitempty"`
}

func (s BatchAuditTest01Request) String() string {
	return tea.Prettify(s)
}

func (s BatchAuditTest01Request) GoString() string {
	return s.String()
}

func (s *BatchAuditTest01Request) SetName(v string) *BatchAuditTest01Request {
	s.Name = &v
	return s
}

func (s *BatchAuditTest01Request) SetBatchAuditTest01(v string) *BatchAuditTest01Request {
	s.BatchAuditTest01 = &v
	return s
}

func (s *BatchAuditTest01Request) SetDemo01(v string) *BatchAuditTest01Request {
	s.Demo01 = &v
	return s
}

func (s *BatchAuditTest01Request) SetTest010101(v bool) *BatchAuditTest01Request {
	s.Test010101 = &v
	return s
}

type BatchAuditTest01ResponseBody struct {
	Name      *string                             `json:"Name,omitempty" xml:"Name,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Demo01    *BatchAuditTest01ResponseBodyDemo01 `json:"Demo01,omitempty" xml:"Demo01,omitempty" type:"Struct"`
}

func (s BatchAuditTest01ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchAuditTest01ResponseBody) GoString() string {
	return s.String()
}

func (s *BatchAuditTest01ResponseBody) SetName(v string) *BatchAuditTest01ResponseBody {
	s.Name = &v
	return s
}

func (s *BatchAuditTest01ResponseBody) SetRequestId(v string) *BatchAuditTest01ResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchAuditTest01ResponseBody) SetDemo01(v *BatchAuditTest01ResponseBodyDemo01) *BatchAuditTest01ResponseBody {
	s.Demo01 = v
	return s
}

type BatchAuditTest01ResponseBodyDemo01 struct {
	Demo011 *BatchAuditTest01ResponseBodyDemo01Demo011 `json:"Demo011,omitempty" xml:"Demo011,omitempty" type:"Struct"`
}

func (s BatchAuditTest01ResponseBodyDemo01) String() string {
	return tea.Prettify(s)
}

func (s BatchAuditTest01ResponseBodyDemo01) GoString() string {
	return s.String()
}

func (s *BatchAuditTest01ResponseBodyDemo01) SetDemo011(v *BatchAuditTest01ResponseBodyDemo01Demo011) *BatchAuditTest01ResponseBodyDemo01 {
	s.Demo011 = v
	return s
}

type BatchAuditTest01ResponseBodyDemo01Demo011 struct {
	Demo011 []*BatchAuditTest01ResponseBodyDemo01Demo011Demo011 `json:"Demo011,omitempty" xml:"Demo011,omitempty" type:"Repeated"`
}

func (s BatchAuditTest01ResponseBodyDemo01Demo011) String() string {
	return tea.Prettify(s)
}

func (s BatchAuditTest01ResponseBodyDemo01Demo011) GoString() string {
	return s.String()
}

func (s *BatchAuditTest01ResponseBodyDemo01Demo011) SetDemo011(v []*BatchAuditTest01ResponseBodyDemo01Demo011Demo011) *BatchAuditTest01ResponseBodyDemo01Demo011 {
	s.Demo011 = v
	return s
}

type BatchAuditTest01ResponseBodyDemo01Demo011Demo011 struct {
	Demo0111 *string `json:"Demo0111,omitempty" xml:"Demo0111,omitempty"`
}

func (s BatchAuditTest01ResponseBodyDemo01Demo011Demo011) String() string {
	return tea.Prettify(s)
}

func (s BatchAuditTest01ResponseBodyDemo01Demo011Demo011) GoString() string {
	return s.String()
}

func (s *BatchAuditTest01ResponseBodyDemo01Demo011Demo011) SetDemo0111(v string) *BatchAuditTest01ResponseBodyDemo01Demo011Demo011 {
	s.Demo0111 = &v
	return s
}

type BatchAuditTest01Response struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchAuditTest01ResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchAuditTest01Response) String() string {
	return tea.Prettify(s)
}

func (s BatchAuditTest01Response) GoString() string {
	return s.String()
}

func (s *BatchAuditTest01Response) SetHeaders(v map[string]*string) *BatchAuditTest01Response {
	s.Headers = v
	return s
}

func (s *BatchAuditTest01Response) SetBody(v *BatchAuditTest01ResponseBody) *BatchAuditTest01Response {
	s.Body = v
	return s
}

type FTApiAliasApiRequest struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s FTApiAliasApiRequest) String() string {
	return tea.Prettify(s)
}

func (s FTApiAliasApiRequest) GoString() string {
	return s.String()
}

func (s *FTApiAliasApiRequest) SetName(v string) *FTApiAliasApiRequest {
	s.Name = &v
	return s
}

type FTApiAliasApiResponseBody struct {
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s FTApiAliasApiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FTApiAliasApiResponseBody) GoString() string {
	return s.String()
}

func (s *FTApiAliasApiResponseBody) SetName(v string) *FTApiAliasApiResponseBody {
	s.Name = &v
	return s
}

func (s *FTApiAliasApiResponseBody) SetRequestId(v string) *FTApiAliasApiResponseBody {
	s.RequestId = &v
	return s
}

type FTApiAliasApiResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *FTApiAliasApiResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FTApiAliasApiResponse) String() string {
	return tea.Prettify(s)
}

func (s FTApiAliasApiResponse) GoString() string {
	return s.String()
}

func (s *FTApiAliasApiResponse) SetHeaders(v map[string]*string) *FTApiAliasApiResponse {
	s.Headers = v
	return s
}

func (s *FTApiAliasApiResponse) SetBody(v *FTApiAliasApiResponseBody) *FTApiAliasApiResponse {
	s.Body = v
	return s
}

type FtDynamicAddressDubboRequest struct {
	IntValue    *int32  `json:"IntValue,omitempty" xml:"IntValue,omitempty"`
	StringValue *string `json:"StringValue,omitempty" xml:"StringValue,omitempty"`
}

func (s FtDynamicAddressDubboRequest) String() string {
	return tea.Prettify(s)
}

func (s FtDynamicAddressDubboRequest) GoString() string {
	return s.String()
}

func (s *FtDynamicAddressDubboRequest) SetIntValue(v int32) *FtDynamicAddressDubboRequest {
	s.IntValue = &v
	return s
}

func (s *FtDynamicAddressDubboRequest) SetStringValue(v string) *FtDynamicAddressDubboRequest {
	s.StringValue = &v
	return s
}

type FtDynamicAddressDubboResponseBody struct {
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StringValue *string `json:"StringValue,omitempty" xml:"StringValue,omitempty"`
	IntValue    *int32  `json:"IntValue,omitempty" xml:"IntValue,omitempty"`
}

func (s FtDynamicAddressDubboResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FtDynamicAddressDubboResponseBody) GoString() string {
	return s.String()
}

func (s *FtDynamicAddressDubboResponseBody) SetRequestId(v string) *FtDynamicAddressDubboResponseBody {
	s.RequestId = &v
	return s
}

func (s *FtDynamicAddressDubboResponseBody) SetStringValue(v string) *FtDynamicAddressDubboResponseBody {
	s.StringValue = &v
	return s
}

func (s *FtDynamicAddressDubboResponseBody) SetIntValue(v int32) *FtDynamicAddressDubboResponseBody {
	s.IntValue = &v
	return s
}

type FtDynamicAddressDubboResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *FtDynamicAddressDubboResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FtDynamicAddressDubboResponse) String() string {
	return tea.Prettify(s)
}

func (s FtDynamicAddressDubboResponse) GoString() string {
	return s.String()
}

func (s *FtDynamicAddressDubboResponse) SetHeaders(v map[string]*string) *FtDynamicAddressDubboResponse {
	s.Headers = v
	return s
}

func (s *FtDynamicAddressDubboResponse) SetBody(v *FtDynamicAddressDubboResponseBody) *FtDynamicAddressDubboResponse {
	s.Body = v
	return s
}

type FtDynamicAddressHsfResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s FtDynamicAddressHsfResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FtDynamicAddressHsfResponseBody) GoString() string {
	return s.String()
}

func (s *FtDynamicAddressHsfResponseBody) SetRequestId(v string) *FtDynamicAddressHsfResponseBody {
	s.RequestId = &v
	return s
}

type FtDynamicAddressHsfResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *FtDynamicAddressHsfResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FtDynamicAddressHsfResponse) String() string {
	return tea.Prettify(s)
}

func (s FtDynamicAddressHsfResponse) GoString() string {
	return s.String()
}

func (s *FtDynamicAddressHsfResponse) SetHeaders(v map[string]*string) *FtDynamicAddressHsfResponse {
	s.Headers = v
	return s
}

func (s *FtDynamicAddressHsfResponse) SetBody(v *FtDynamicAddressHsfResponseBody) *FtDynamicAddressHsfResponse {
	s.Body = v
	return s
}

type FtDynamicAddressHttpVpcRequest struct {
	StringValue  map[string]interface{} `json:"StringValue,omitempty" xml:"StringValue,omitempty"`
	DefaultValue map[string]interface{} `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	OtherParam   map[string]interface{} `json:"OtherParam,omitempty" xml:"OtherParam,omitempty"`
	BooleanParam *bool                  `json:"BooleanParam,omitempty" xml:"BooleanParam,omitempty"`
	P1           *string                `json:"P1,omitempty" xml:"P1,omitempty"`
}

func (s FtDynamicAddressHttpVpcRequest) String() string {
	return tea.Prettify(s)
}

func (s FtDynamicAddressHttpVpcRequest) GoString() string {
	return s.String()
}

func (s *FtDynamicAddressHttpVpcRequest) SetStringValue(v map[string]interface{}) *FtDynamicAddressHttpVpcRequest {
	s.StringValue = v
	return s
}

func (s *FtDynamicAddressHttpVpcRequest) SetDefaultValue(v map[string]interface{}) *FtDynamicAddressHttpVpcRequest {
	s.DefaultValue = v
	return s
}

func (s *FtDynamicAddressHttpVpcRequest) SetOtherParam(v map[string]interface{}) *FtDynamicAddressHttpVpcRequest {
	s.OtherParam = v
	return s
}

func (s *FtDynamicAddressHttpVpcRequest) SetBooleanParam(v bool) *FtDynamicAddressHttpVpcRequest {
	s.BooleanParam = &v
	return s
}

func (s *FtDynamicAddressHttpVpcRequest) SetP1(v string) *FtDynamicAddressHttpVpcRequest {
	s.P1 = &v
	return s
}

type FtDynamicAddressHttpVpcShrinkRequest struct {
	StringValueShrink  *string `json:"StringValue,omitempty" xml:"StringValue,omitempty"`
	DefaultValueShrink *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	OtherParamShrink   *string `json:"OtherParam,omitempty" xml:"OtherParam,omitempty"`
	BooleanParam       *bool   `json:"BooleanParam,omitempty" xml:"BooleanParam,omitempty"`
	P1                 *string `json:"P1,omitempty" xml:"P1,omitempty"`
}

func (s FtDynamicAddressHttpVpcShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s FtDynamicAddressHttpVpcShrinkRequest) GoString() string {
	return s.String()
}

func (s *FtDynamicAddressHttpVpcShrinkRequest) SetStringValueShrink(v string) *FtDynamicAddressHttpVpcShrinkRequest {
	s.StringValueShrink = &v
	return s
}

func (s *FtDynamicAddressHttpVpcShrinkRequest) SetDefaultValueShrink(v string) *FtDynamicAddressHttpVpcShrinkRequest {
	s.DefaultValueShrink = &v
	return s
}

func (s *FtDynamicAddressHttpVpcShrinkRequest) SetOtherParamShrink(v string) *FtDynamicAddressHttpVpcShrinkRequest {
	s.OtherParamShrink = &v
	return s
}

func (s *FtDynamicAddressHttpVpcShrinkRequest) SetBooleanParam(v bool) *FtDynamicAddressHttpVpcShrinkRequest {
	s.BooleanParam = &v
	return s
}

func (s *FtDynamicAddressHttpVpcShrinkRequest) SetP1(v string) *FtDynamicAddressHttpVpcShrinkRequest {
	s.P1 = &v
	return s
}

type FtDynamicAddressHttpVpcResponseBody struct {
	ServiceRpcSign *string `json:"ServiceRpcSign,omitempty" xml:"ServiceRpcSign,omitempty"`
	Params         *string `json:"Params,omitempty" xml:"Params,omitempty"`
}

func (s FtDynamicAddressHttpVpcResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FtDynamicAddressHttpVpcResponseBody) GoString() string {
	return s.String()
}

func (s *FtDynamicAddressHttpVpcResponseBody) SetServiceRpcSign(v string) *FtDynamicAddressHttpVpcResponseBody {
	s.ServiceRpcSign = &v
	return s
}

func (s *FtDynamicAddressHttpVpcResponseBody) SetParams(v string) *FtDynamicAddressHttpVpcResponseBody {
	s.Params = &v
	return s
}

type FtDynamicAddressHttpVpcResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *FtDynamicAddressHttpVpcResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FtDynamicAddressHttpVpcResponse) String() string {
	return tea.Prettify(s)
}

func (s FtDynamicAddressHttpVpcResponse) GoString() string {
	return s.String()
}

func (s *FtDynamicAddressHttpVpcResponse) SetHeaders(v map[string]*string) *FtDynamicAddressHttpVpcResponse {
	s.Headers = v
	return s
}

func (s *FtDynamicAddressHttpVpcResponse) SetBody(v *FtDynamicAddressHttpVpcResponseBody) *FtDynamicAddressHttpVpcResponse {
	s.Body = v
	return s
}

type FtEagleEyeRequest struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s FtEagleEyeRequest) String() string {
	return tea.Prettify(s)
}

func (s FtEagleEyeRequest) GoString() string {
	return s.String()
}

func (s *FtEagleEyeRequest) SetName(v string) *FtEagleEyeRequest {
	s.Name = &v
	return s
}

type FtEagleEyeResponseBody struct {
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	EagleEyeTraceId *string `json:"eagleEyeTraceId,omitempty" xml:"eagleEyeTraceId,omitempty"`
}

func (s FtEagleEyeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FtEagleEyeResponseBody) GoString() string {
	return s.String()
}

func (s *FtEagleEyeResponseBody) SetRequestId(v string) *FtEagleEyeResponseBody {
	s.RequestId = &v
	return s
}

func (s *FtEagleEyeResponseBody) SetName(v string) *FtEagleEyeResponseBody {
	s.Name = &v
	return s
}

func (s *FtEagleEyeResponseBody) SetEagleEyeTraceId(v string) *FtEagleEyeResponseBody {
	s.EagleEyeTraceId = &v
	return s
}

type FtEagleEyeResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *FtEagleEyeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FtEagleEyeResponse) String() string {
	return tea.Prettify(s)
}

func (s FtEagleEyeResponse) GoString() string {
	return s.String()
}

func (s *FtEagleEyeResponse) SetHeaders(v map[string]*string) *FtEagleEyeResponse {
	s.Headers = v
	return s
}

func (s *FtEagleEyeResponse) SetBody(v *FtEagleEyeResponseBody) *FtEagleEyeResponse {
	s.Body = v
	return s
}

type FtFlowSpecialRequest struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s FtFlowSpecialRequest) String() string {
	return tea.Prettify(s)
}

func (s FtFlowSpecialRequest) GoString() string {
	return s.String()
}

func (s *FtFlowSpecialRequest) SetName(v string) *FtFlowSpecialRequest {
	s.Name = &v
	return s
}

type FtFlowSpecialResponseBody struct {
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s FtFlowSpecialResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FtFlowSpecialResponseBody) GoString() string {
	return s.String()
}

func (s *FtFlowSpecialResponseBody) SetName(v string) *FtFlowSpecialResponseBody {
	s.Name = &v
	return s
}

func (s *FtFlowSpecialResponseBody) SetRequestId(v string) *FtFlowSpecialResponseBody {
	s.RequestId = &v
	return s
}

type FtFlowSpecialResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *FtFlowSpecialResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FtFlowSpecialResponse) String() string {
	return tea.Prettify(s)
}

func (s FtFlowSpecialResponse) GoString() string {
	return s.String()
}

func (s *FtFlowSpecialResponse) SetHeaders(v map[string]*string) *FtFlowSpecialResponse {
	s.Headers = v
	return s
}

func (s *FtFlowSpecialResponse) SetBody(v *FtFlowSpecialResponseBody) *FtFlowSpecialResponse {
	s.Body = v
	return s
}

type FtGatedLaunchPolicy4Request struct {
	IsGatedLaunch *string `json:"IsGatedLaunch,omitempty" xml:"IsGatedLaunch,omitempty"`
}

func (s FtGatedLaunchPolicy4Request) String() string {
	return tea.Prettify(s)
}

func (s FtGatedLaunchPolicy4Request) GoString() string {
	return s.String()
}

func (s *FtGatedLaunchPolicy4Request) SetIsGatedLaunch(v string) *FtGatedLaunchPolicy4Request {
	s.IsGatedLaunch = &v
	return s
}

type FtGatedLaunchPolicy4ResponseBody struct {
	IsGatedLaunch *string `json:"IsGatedLaunch,omitempty" xml:"IsGatedLaunch,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s FtGatedLaunchPolicy4ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FtGatedLaunchPolicy4ResponseBody) GoString() string {
	return s.String()
}

func (s *FtGatedLaunchPolicy4ResponseBody) SetIsGatedLaunch(v string) *FtGatedLaunchPolicy4ResponseBody {
	s.IsGatedLaunch = &v
	return s
}

func (s *FtGatedLaunchPolicy4ResponseBody) SetRequestId(v string) *FtGatedLaunchPolicy4ResponseBody {
	s.RequestId = &v
	return s
}

type FtGatedLaunchPolicy4Response struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *FtGatedLaunchPolicy4ResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FtGatedLaunchPolicy4Response) String() string {
	return tea.Prettify(s)
}

func (s FtGatedLaunchPolicy4Response) GoString() string {
	return s.String()
}

func (s *FtGatedLaunchPolicy4Response) SetHeaders(v map[string]*string) *FtGatedLaunchPolicy4Response {
	s.Headers = v
	return s
}

func (s *FtGatedLaunchPolicy4Response) SetBody(v *FtGatedLaunchPolicy4ResponseBody) *FtGatedLaunchPolicy4Response {
	s.Body = v
	return s
}

type FtIpFlowControlRequest struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s FtIpFlowControlRequest) String() string {
	return tea.Prettify(s)
}

func (s FtIpFlowControlRequest) GoString() string {
	return s.String()
}

func (s *FtIpFlowControlRequest) SetName(v string) *FtIpFlowControlRequest {
	s.Name = &v
	return s
}

type FtIpFlowControlResponseBody struct {
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s FtIpFlowControlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FtIpFlowControlResponseBody) GoString() string {
	return s.String()
}

func (s *FtIpFlowControlResponseBody) SetName(v string) *FtIpFlowControlResponseBody {
	s.Name = &v
	return s
}

func (s *FtIpFlowControlResponseBody) SetRequestId(v string) *FtIpFlowControlResponseBody {
	s.RequestId = &v
	return s
}

type FtIpFlowControlResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *FtIpFlowControlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FtIpFlowControlResponse) String() string {
	return tea.Prettify(s)
}

func (s FtIpFlowControlResponse) GoString() string {
	return s.String()
}

func (s *FtIpFlowControlResponse) SetHeaders(v map[string]*string) *FtIpFlowControlResponse {
	s.Headers = v
	return s
}

func (s *FtIpFlowControlResponse) SetBody(v *FtIpFlowControlResponseBody) *FtIpFlowControlResponse {
	s.Body = v
	return s
}

type FtParamListRequest struct {
	Name *string                   `json:"Name,omitempty" xml:"Name,omitempty"`
	Disk []*FtParamListRequestDisk `json:"Disk,omitempty" xml:"Disk,omitempty" type:"Repeated"`
}

func (s FtParamListRequest) String() string {
	return tea.Prettify(s)
}

func (s FtParamListRequest) GoString() string {
	return s.String()
}

func (s *FtParamListRequest) SetName(v string) *FtParamListRequest {
	s.Name = &v
	return s
}

func (s *FtParamListRequest) SetDisk(v []*FtParamListRequestDisk) *FtParamListRequest {
	s.Disk = v
	return s
}

type FtParamListRequestDisk struct {
	Type []*string `json:"Type,omitempty" xml:"Type,omitempty" type:"Repeated"`
	Size []*string `json:"Size,omitempty" xml:"Size,omitempty" type:"Repeated"`
}

func (s FtParamListRequestDisk) String() string {
	return tea.Prettify(s)
}

func (s FtParamListRequestDisk) GoString() string {
	return s.String()
}

func (s *FtParamListRequestDisk) SetType(v []*string) *FtParamListRequestDisk {
	s.Type = v
	return s
}

func (s *FtParamListRequestDisk) SetSize(v []*string) *FtParamListRequestDisk {
	s.Size = v
	return s
}

type FtParamListResponseBody struct {
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s FtParamListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FtParamListResponseBody) GoString() string {
	return s.String()
}

func (s *FtParamListResponseBody) SetName(v string) *FtParamListResponseBody {
	s.Name = &v
	return s
}

func (s *FtParamListResponseBody) SetRequestId(v string) *FtParamListResponseBody {
	s.RequestId = &v
	return s
}

type FtParamListResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *FtParamListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FtParamListResponse) String() string {
	return tea.Prettify(s)
}

func (s FtParamListResponse) GoString() string {
	return s.String()
}

func (s *FtParamListResponse) SetHeaders(v map[string]*string) *FtParamListResponse {
	s.Headers = v
	return s
}

func (s *FtParamListResponse) SetBody(v *FtParamListResponseBody) *FtParamListResponse {
	s.Body = v
	return s
}

type TestFlowStrategy01Request struct {
	Names map[string]interface{} `json:"Names,omitempty" xml:"Names,omitempty"`
}

func (s TestFlowStrategy01Request) String() string {
	return tea.Prettify(s)
}

func (s TestFlowStrategy01Request) GoString() string {
	return s.String()
}

func (s *TestFlowStrategy01Request) SetNames(v map[string]interface{}) *TestFlowStrategy01Request {
	s.Names = v
	return s
}

type TestFlowStrategy01ShrinkRequest struct {
	NamesShrink *string `json:"Names,omitempty" xml:"Names,omitempty"`
}

func (s TestFlowStrategy01ShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s TestFlowStrategy01ShrinkRequest) GoString() string {
	return s.String()
}

func (s *TestFlowStrategy01ShrinkRequest) SetNamesShrink(v string) *TestFlowStrategy01ShrinkRequest {
	s.NamesShrink = &v
	return s
}

type TestFlowStrategy01ResponseBody struct {
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	List      []*string `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	Names     []*string `json:"Names,omitempty" xml:"Names,omitempty" type:"Repeated"`
}

func (s TestFlowStrategy01ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TestFlowStrategy01ResponseBody) GoString() string {
	return s.String()
}

func (s *TestFlowStrategy01ResponseBody) SetRequestId(v string) *TestFlowStrategy01ResponseBody {
	s.RequestId = &v
	return s
}

func (s *TestFlowStrategy01ResponseBody) SetList(v []*string) *TestFlowStrategy01ResponseBody {
	s.List = v
	return s
}

func (s *TestFlowStrategy01ResponseBody) SetNames(v []*string) *TestFlowStrategy01ResponseBody {
	s.Names = v
	return s
}

type TestFlowStrategy01Response struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *TestFlowStrategy01ResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TestFlowStrategy01Response) String() string {
	return tea.Prettify(s)
}

func (s TestFlowStrategy01Response) GoString() string {
	return s.String()
}

func (s *TestFlowStrategy01Response) SetHeaders(v map[string]*string) *TestFlowStrategy01Response {
	s.Headers = v
	return s
}

func (s *TestFlowStrategy01Response) SetBody(v *TestFlowStrategy01ResponseBody) *TestFlowStrategy01Response {
	s.Body = v
	return s
}

type TestHttpApiRequest struct {
	StringValue  map[string]interface{} `json:"StringValue,omitempty" xml:"StringValue,omitempty"`
	DefaultValue map[string]interface{} `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	OtherParam   map[string]interface{} `json:"OtherParam,omitempty" xml:"OtherParam,omitempty"`
	BooleanParam *bool                  `json:"BooleanParam,omitempty" xml:"BooleanParam,omitempty"`
}

func (s TestHttpApiRequest) String() string {
	return tea.Prettify(s)
}

func (s TestHttpApiRequest) GoString() string {
	return s.String()
}

func (s *TestHttpApiRequest) SetStringValue(v map[string]interface{}) *TestHttpApiRequest {
	s.StringValue = v
	return s
}

func (s *TestHttpApiRequest) SetDefaultValue(v map[string]interface{}) *TestHttpApiRequest {
	s.DefaultValue = v
	return s
}

func (s *TestHttpApiRequest) SetOtherParam(v map[string]interface{}) *TestHttpApiRequest {
	s.OtherParam = v
	return s
}

func (s *TestHttpApiRequest) SetBooleanParam(v bool) *TestHttpApiRequest {
	s.BooleanParam = &v
	return s
}

type TestHttpApiShrinkRequest struct {
	StringValueShrink  *string `json:"StringValue,omitempty" xml:"StringValue,omitempty"`
	DefaultValueShrink *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	OtherParamShrink   *string `json:"OtherParam,omitempty" xml:"OtherParam,omitempty"`
	BooleanParam       *bool   `json:"BooleanParam,omitempty" xml:"BooleanParam,omitempty"`
}

func (s TestHttpApiShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s TestHttpApiShrinkRequest) GoString() string {
	return s.String()
}

func (s *TestHttpApiShrinkRequest) SetStringValueShrink(v string) *TestHttpApiShrinkRequest {
	s.StringValueShrink = &v
	return s
}

func (s *TestHttpApiShrinkRequest) SetDefaultValueShrink(v string) *TestHttpApiShrinkRequest {
	s.DefaultValueShrink = &v
	return s
}

func (s *TestHttpApiShrinkRequest) SetOtherParamShrink(v string) *TestHttpApiShrinkRequest {
	s.OtherParamShrink = &v
	return s
}

func (s *TestHttpApiShrinkRequest) SetBooleanParam(v bool) *TestHttpApiShrinkRequest {
	s.BooleanParam = &v
	return s
}

type TestHttpApiResponseBody struct {
	ServiceRpcSign *string `json:"ServiceRpcSign,omitempty" xml:"ServiceRpcSign,omitempty"`
	Params         *string `json:"Params,omitempty" xml:"Params,omitempty"`
}

func (s TestHttpApiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TestHttpApiResponseBody) GoString() string {
	return s.String()
}

func (s *TestHttpApiResponseBody) SetServiceRpcSign(v string) *TestHttpApiResponseBody {
	s.ServiceRpcSign = &v
	return s
}

func (s *TestHttpApiResponseBody) SetParams(v string) *TestHttpApiResponseBody {
	s.Params = &v
	return s
}

type TestHttpApiResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *TestHttpApiResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TestHttpApiResponse) String() string {
	return tea.Prettify(s)
}

func (s TestHttpApiResponse) GoString() string {
	return s.String()
}

func (s *TestHttpApiResponse) SetHeaders(v map[string]*string) *TestHttpApiResponse {
	s.Headers = v
	return s
}

func (s *TestHttpApiResponse) SetBody(v *TestHttpApiResponseBody) *TestHttpApiResponse {
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
		"ap-northeast-2-pop":          tea.String("ft.aliyuncs.com"),
		"ap-south-1":                  tea.String("ft.aliyuncs.com"),
		"ap-southeast-1":              tea.String("ft.aliyuncs.com"),
		"ap-southeast-2":              tea.String("ft.aliyuncs.com"),
		"ap-southeast-3":              tea.String("ft.aliyuncs.com"),
		"ap-southeast-5":              tea.String("ft.aliyuncs.com"),
		"cn-beijing":                  tea.String("ft.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("ft.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("ft.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("ft.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("ft.aliyuncs.com"),
		"cn-chengdu":                  tea.String("ft.aliyuncs.com"),
		"cn-edge-1":                   tea.String("ft.aliyuncs.com"),
		"cn-fujian":                   tea.String("ft.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("ft.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("ft.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("ft.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("ft.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("ft.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("ft.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("ft.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("ft.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("ft.aliyuncs.com"),
		"cn-huhehaote":                tea.String("ft.aliyuncs.com"),
		"cn-huhehaote-nebula-1":       tea.String("ft.aliyuncs.com"),
		"cn-qingdao":                  tea.String("ft.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("ft.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("ft.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("ft.aliyuncs.com"),
		"cn-shanghai-finance-1":       tea.String("ft.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("ft.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("ft.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("ft.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("ft.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("ft.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("ft.aliyuncs.com"),
		"cn-wuhan":                    tea.String("ft.aliyuncs.com"),
		"cn-wulanchabu":               tea.String("ft.aliyuncs.com"),
		"cn-yushanfang":               tea.String("ft.aliyuncs.com"),
		"cn-zhangbei":                 tea.String("ft.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("ft.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("ft.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("ft.aliyuncs.com"),
		"eu-central-1":                tea.String("ft.aliyuncs.com"),
		"eu-west-1":                   tea.String("ft.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("ft.aliyuncs.com"),
		"me-east-1":                   tea.String("ft.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("ft.aliyuncs.com"),
		"us-west-1":                   tea.String("ft.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("ft"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) BatchAuditTest01WithOptions(request *BatchAuditTest01Request, runtime *util.RuntimeOptions) (_result *BatchAuditTest01Response, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &BatchAuditTest01Response{}
	_body, _err := client.DoRPCRequest(tea.String("BatchAuditTest01"), tea.String("2018-07-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchAuditTest01(request *BatchAuditTest01Request) (_result *BatchAuditTest01Response, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchAuditTest01Response{}
	_body, _err := client.BatchAuditTest01WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FTApiAliasApiWithOptions(request *FTApiAliasApiRequest, runtime *util.RuntimeOptions) (_result *FTApiAliasApiResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &FTApiAliasApiResponse{}
	_body, _err := client.DoRPCRequest(tea.String("FTApiAliasApi"), tea.String("2018-07-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FTApiAliasApi(request *FTApiAliasApiRequest) (_result *FTApiAliasApiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FTApiAliasApiResponse{}
	_body, _err := client.FTApiAliasApiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FtDynamicAddressDubboWithOptions(request *FtDynamicAddressDubboRequest, runtime *util.RuntimeOptions) (_result *FtDynamicAddressDubboResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &FtDynamicAddressDubboResponse{}
	_body, _err := client.DoRPCRequest(tea.String("FtDynamicAddressDubbo"), tea.String("2018-07-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FtDynamicAddressDubbo(request *FtDynamicAddressDubboRequest) (_result *FtDynamicAddressDubboResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FtDynamicAddressDubboResponse{}
	_body, _err := client.FtDynamicAddressDubboWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FtDynamicAddressHsfWithOptions(runtime *util.RuntimeOptions) (_result *FtDynamicAddressHsfResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	_result = &FtDynamicAddressHsfResponse{}
	_body, _err := client.DoRPCRequest(tea.String("FtDynamicAddressHsf"), tea.String("2018-07-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FtDynamicAddressHsf() (_result *FtDynamicAddressHsfResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FtDynamicAddressHsfResponse{}
	_body, _err := client.FtDynamicAddressHsfWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FtDynamicAddressHttpVpcWithOptions(tmpReq *FtDynamicAddressHttpVpcRequest, runtime *util.RuntimeOptions) (_result *FtDynamicAddressHttpVpcResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &FtDynamicAddressHttpVpcShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.StringValue)) {
		request.StringValueShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StringValue, tea.String("StringValue"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.DefaultValue)) {
		request.DefaultValueShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DefaultValue, tea.String("DefaultValue"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.OtherParam)) {
		request.OtherParamShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OtherParam, tea.String("OtherParam"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &FtDynamicAddressHttpVpcResponse{}
	_body, _err := client.DoRPCRequest(tea.String("FtDynamicAddressHttpVpc"), tea.String("2018-07-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FtDynamicAddressHttpVpc(request *FtDynamicAddressHttpVpcRequest) (_result *FtDynamicAddressHttpVpcResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FtDynamicAddressHttpVpcResponse{}
	_body, _err := client.FtDynamicAddressHttpVpcWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FtEagleEyeWithOptions(request *FtEagleEyeRequest, runtime *util.RuntimeOptions) (_result *FtEagleEyeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &FtEagleEyeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("FtEagleEye"), tea.String("2018-07-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FtEagleEye(request *FtEagleEyeRequest) (_result *FtEagleEyeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FtEagleEyeResponse{}
	_body, _err := client.FtEagleEyeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FtFlowSpecialWithOptions(request *FtFlowSpecialRequest, runtime *util.RuntimeOptions) (_result *FtFlowSpecialResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &FtFlowSpecialResponse{}
	_body, _err := client.DoRPCRequest(tea.String("FtFlowSpecial"), tea.String("2018-07-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FtFlowSpecial(request *FtFlowSpecialRequest) (_result *FtFlowSpecialResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FtFlowSpecialResponse{}
	_body, _err := client.FtFlowSpecialWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FtGatedLaunchPolicy4WithOptions(request *FtGatedLaunchPolicy4Request, runtime *util.RuntimeOptions) (_result *FtGatedLaunchPolicy4Response, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &FtGatedLaunchPolicy4Response{}
	_body, _err := client.DoRPCRequest(tea.String("FtGatedLaunchPolicy4"), tea.String("2018-07-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FtGatedLaunchPolicy4(request *FtGatedLaunchPolicy4Request) (_result *FtGatedLaunchPolicy4Response, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FtGatedLaunchPolicy4Response{}
	_body, _err := client.FtGatedLaunchPolicy4WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FtIpFlowControlWithOptions(request *FtIpFlowControlRequest, runtime *util.RuntimeOptions) (_result *FtIpFlowControlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &FtIpFlowControlResponse{}
	_body, _err := client.DoRPCRequest(tea.String("FtIpFlowControl"), tea.String("2018-07-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FtIpFlowControl(request *FtIpFlowControlRequest) (_result *FtIpFlowControlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FtIpFlowControlResponse{}
	_body, _err := client.FtIpFlowControlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FtParamListWithOptions(request *FtParamListRequest, runtime *util.RuntimeOptions) (_result *FtParamListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &FtParamListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("FtParamList"), tea.String("2018-07-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FtParamList(request *FtParamListRequest) (_result *FtParamListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FtParamListResponse{}
	_body, _err := client.FtParamListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TestFlowStrategy01WithOptions(tmpReq *TestFlowStrategy01Request, runtime *util.RuntimeOptions) (_result *TestFlowStrategy01Response, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &TestFlowStrategy01ShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Names)) {
		request.NamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Names, tea.String("Names"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &TestFlowStrategy01Response{}
	_body, _err := client.DoRPCRequest(tea.String("TestFlowStrategy01"), tea.String("2018-07-13"), tea.String("HTTPS"), tea.String("PUT"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TestFlowStrategy01(request *TestFlowStrategy01Request) (_result *TestFlowStrategy01Response, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TestFlowStrategy01Response{}
	_body, _err := client.TestFlowStrategy01WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TestHttpApiWithOptions(tmpReq *TestHttpApiRequest, runtime *util.RuntimeOptions) (_result *TestHttpApiResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &TestHttpApiShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.StringValue)) {
		request.StringValueShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StringValue, tea.String("StringValue"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.DefaultValue)) {
		request.DefaultValueShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DefaultValue, tea.String("DefaultValue"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.OtherParam)) {
		request.OtherParamShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OtherParam, tea.String("OtherParam"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &TestHttpApiResponse{}
	_body, _err := client.DoRPCRequest(tea.String("TestHttpApi"), tea.String("2018-07-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TestHttpApi(request *TestHttpApiRequest) (_result *TestHttpApiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TestHttpApiResponse{}
	_body, _err := client.TestHttpApiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
