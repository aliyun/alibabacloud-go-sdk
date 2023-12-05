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

type CheckReadyFlagRequest struct {
	Cycles     *int32  `json:"Cycles,omitempty" xml:"Cycles,omitempty"`
	DataType   *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	DomainCode *string `json:"DomainCode,omitempty" xml:"DomainCode,omitempty"`
	EndTime    *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Period     *string `json:"Period,omitempty" xml:"Period,omitempty"`
	StartTime  *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s CheckReadyFlagRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckReadyFlagRequest) GoString() string {
	return s.String()
}

func (s *CheckReadyFlagRequest) SetCycles(v int32) *CheckReadyFlagRequest {
	s.Cycles = &v
	return s
}

func (s *CheckReadyFlagRequest) SetDataType(v string) *CheckReadyFlagRequest {
	s.DataType = &v
	return s
}

func (s *CheckReadyFlagRequest) SetDomainCode(v string) *CheckReadyFlagRequest {
	s.DomainCode = &v
	return s
}

func (s *CheckReadyFlagRequest) SetEndTime(v int64) *CheckReadyFlagRequest {
	s.EndTime = &v
	return s
}

func (s *CheckReadyFlagRequest) SetPeriod(v string) *CheckReadyFlagRequest {
	s.Period = &v
	return s
}

func (s *CheckReadyFlagRequest) SetStartTime(v int64) *CheckReadyFlagRequest {
	s.StartTime = &v
	return s
}

type CheckReadyFlagResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckReadyFlagResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckReadyFlagResponseBody) GoString() string {
	return s.String()
}

func (s *CheckReadyFlagResponseBody) SetData(v string) *CheckReadyFlagResponseBody {
	s.Data = &v
	return s
}

func (s *CheckReadyFlagResponseBody) SetRequestId(v string) *CheckReadyFlagResponseBody {
	s.RequestId = &v
	return s
}

type CheckReadyFlagResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CheckReadyFlagResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckReadyFlagResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckReadyFlagResponse) GoString() string {
	return s.String()
}

func (s *CheckReadyFlagResponse) SetHeaders(v map[string]*string) *CheckReadyFlagResponse {
	s.Headers = v
	return s
}

func (s *CheckReadyFlagResponse) SetStatusCode(v int32) *CheckReadyFlagResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckReadyFlagResponse) SetBody(v *CheckReadyFlagResponseBody) *CheckReadyFlagResponse {
	s.Body = v
	return s
}

type DeleteDomainPartRequest struct {
	DataType   *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	DomainCode *string `json:"DomainCode,omitempty" xml:"DomainCode,omitempty"`
	Part       *string `json:"Part,omitempty" xml:"Part,omitempty"`
}

func (s DeleteDomainPartRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDomainPartRequest) GoString() string {
	return s.String()
}

func (s *DeleteDomainPartRequest) SetDataType(v string) *DeleteDomainPartRequest {
	s.DataType = &v
	return s
}

func (s *DeleteDomainPartRequest) SetDomainCode(v string) *DeleteDomainPartRequest {
	s.DomainCode = &v
	return s
}

func (s *DeleteDomainPartRequest) SetPart(v string) *DeleteDomainPartRequest {
	s.Part = &v
	return s
}

type DeleteDomainPartResponseBody struct {
	DataType   *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	DomainCode *string `json:"DomainCode,omitempty" xml:"DomainCode,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDomainPartResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDomainPartResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDomainPartResponseBody) SetDataType(v string) *DeleteDomainPartResponseBody {
	s.DataType = &v
	return s
}

func (s *DeleteDomainPartResponseBody) SetDomainCode(v string) *DeleteDomainPartResponseBody {
	s.DomainCode = &v
	return s
}

func (s *DeleteDomainPartResponseBody) SetRequestId(v string) *DeleteDomainPartResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDomainPartResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteDomainPartResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDomainPartResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDomainPartResponse) GoString() string {
	return s.String()
}

func (s *DeleteDomainPartResponse) SetHeaders(v map[string]*string) *DeleteDomainPartResponse {
	s.Headers = v
	return s
}

func (s *DeleteDomainPartResponse) SetStatusCode(v int32) *DeleteDomainPartResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDomainPartResponse) SetBody(v *DeleteDomainPartResponseBody) *DeleteDomainPartResponse {
	s.Body = v
	return s
}

type DeleteDomainPartByProxyRequest struct {
	DataType   *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	DomainCode *string `json:"DomainCode,omitempty" xml:"DomainCode,omitempty"`
	Part       *string `json:"Part,omitempty" xml:"Part,omitempty"`
}

func (s DeleteDomainPartByProxyRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDomainPartByProxyRequest) GoString() string {
	return s.String()
}

func (s *DeleteDomainPartByProxyRequest) SetDataType(v string) *DeleteDomainPartByProxyRequest {
	s.DataType = &v
	return s
}

func (s *DeleteDomainPartByProxyRequest) SetDomainCode(v string) *DeleteDomainPartByProxyRequest {
	s.DomainCode = &v
	return s
}

func (s *DeleteDomainPartByProxyRequest) SetPart(v string) *DeleteDomainPartByProxyRequest {
	s.Part = &v
	return s
}

type DeleteDomainPartByProxyResponseBody struct {
	DataType   *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	DomainCode *string `json:"DomainCode,omitempty" xml:"DomainCode,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDomainPartByProxyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDomainPartByProxyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDomainPartByProxyResponseBody) SetDataType(v string) *DeleteDomainPartByProxyResponseBody {
	s.DataType = &v
	return s
}

func (s *DeleteDomainPartByProxyResponseBody) SetDomainCode(v string) *DeleteDomainPartByProxyResponseBody {
	s.DomainCode = &v
	return s
}

func (s *DeleteDomainPartByProxyResponseBody) SetRequestId(v string) *DeleteDomainPartByProxyResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDomainPartByProxyResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteDomainPartByProxyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDomainPartByProxyResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDomainPartByProxyResponse) GoString() string {
	return s.String()
}

func (s *DeleteDomainPartByProxyResponse) SetHeaders(v map[string]*string) *DeleteDomainPartByProxyResponse {
	s.Headers = v
	return s
}

func (s *DeleteDomainPartByProxyResponse) SetStatusCode(v int32) *DeleteDomainPartByProxyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDomainPartByProxyResponse) SetBody(v *DeleteDomainPartByProxyResponseBody) *DeleteDomainPartByProxyResponse {
	s.Body = v
	return s
}

type DeleteMeasureDataRequest struct {
	ApiType    *string `json:"ApiType,omitempty" xml:"ApiType,omitempty"`
	Compressed *bool   `json:"Compressed,omitempty" xml:"Compressed,omitempty"`
	Data       *string `json:"Data,omitempty" xml:"Data,omitempty"`
	DataType   *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// OMS Domain
	DomainCode *string `json:"DomainCode,omitempty" xml:"DomainCode,omitempty"`
	Filter     *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
}

func (s DeleteMeasureDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteMeasureDataRequest) GoString() string {
	return s.String()
}

func (s *DeleteMeasureDataRequest) SetApiType(v string) *DeleteMeasureDataRequest {
	s.ApiType = &v
	return s
}

func (s *DeleteMeasureDataRequest) SetCompressed(v bool) *DeleteMeasureDataRequest {
	s.Compressed = &v
	return s
}

func (s *DeleteMeasureDataRequest) SetData(v string) *DeleteMeasureDataRequest {
	s.Data = &v
	return s
}

func (s *DeleteMeasureDataRequest) SetDataType(v string) *DeleteMeasureDataRequest {
	s.DataType = &v
	return s
}

func (s *DeleteMeasureDataRequest) SetDomainCode(v string) *DeleteMeasureDataRequest {
	s.DomainCode = &v
	return s
}

func (s *DeleteMeasureDataRequest) SetFilter(v string) *DeleteMeasureDataRequest {
	s.Filter = &v
	return s
}

type DeleteMeasureDataResponseBody struct {
	ApiType  *string `json:"ApiType,omitempty" xml:"ApiType,omitempty"`
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// OMS Domain
	DomainCode *string `json:"DomainCode,omitempty" xml:"DomainCode,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total     *int64  `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DeleteMeasureDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteMeasureDataResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMeasureDataResponseBody) SetApiType(v string) *DeleteMeasureDataResponseBody {
	s.ApiType = &v
	return s
}

func (s *DeleteMeasureDataResponseBody) SetDataType(v string) *DeleteMeasureDataResponseBody {
	s.DataType = &v
	return s
}

func (s *DeleteMeasureDataResponseBody) SetDomainCode(v string) *DeleteMeasureDataResponseBody {
	s.DomainCode = &v
	return s
}

func (s *DeleteMeasureDataResponseBody) SetRequestId(v string) *DeleteMeasureDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMeasureDataResponseBody) SetTotal(v int64) *DeleteMeasureDataResponseBody {
	s.Total = &v
	return s
}

type DeleteMeasureDataResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteMeasureDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteMeasureDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteMeasureDataResponse) GoString() string {
	return s.String()
}

func (s *DeleteMeasureDataResponse) SetHeaders(v map[string]*string) *DeleteMeasureDataResponse {
	s.Headers = v
	return s
}

func (s *DeleteMeasureDataResponse) SetStatusCode(v int32) *DeleteMeasureDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMeasureDataResponse) SetBody(v *DeleteMeasureDataResponseBody) *DeleteMeasureDataResponse {
	s.Body = v
	return s
}

type GetAccessPolicyConfigRequest struct {
	AliUid         *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	CompressEnable *bool   `json:"CompressEnable,omitempty" xml:"CompressEnable,omitempty"`
}

func (s GetAccessPolicyConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAccessPolicyConfigRequest) GoString() string {
	return s.String()
}

func (s *GetAccessPolicyConfigRequest) SetAliUid(v string) *GetAccessPolicyConfigRequest {
	s.AliUid = &v
	return s
}

func (s *GetAccessPolicyConfigRequest) SetCompressEnable(v bool) *GetAccessPolicyConfigRequest {
	s.CompressEnable = &v
	return s
}

type GetAccessPolicyConfigResponseBody struct {
	AliUid     *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	Compressed *bool   `json:"Compressed,omitempty" xml:"Compressed,omitempty"`
	Data       *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAccessPolicyConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAccessPolicyConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccessPolicyConfigResponseBody) SetAliUid(v string) *GetAccessPolicyConfigResponseBody {
	s.AliUid = &v
	return s
}

func (s *GetAccessPolicyConfigResponseBody) SetCompressed(v bool) *GetAccessPolicyConfigResponseBody {
	s.Compressed = &v
	return s
}

func (s *GetAccessPolicyConfigResponseBody) SetData(v string) *GetAccessPolicyConfigResponseBody {
	s.Data = &v
	return s
}

func (s *GetAccessPolicyConfigResponseBody) SetRequestId(v string) *GetAccessPolicyConfigResponseBody {
	s.RequestId = &v
	return s
}

type GetAccessPolicyConfigResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAccessPolicyConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAccessPolicyConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAccessPolicyConfigResponse) GoString() string {
	return s.String()
}

func (s *GetAccessPolicyConfigResponse) SetHeaders(v map[string]*string) *GetAccessPolicyConfigResponse {
	s.Headers = v
	return s
}

func (s *GetAccessPolicyConfigResponse) SetStatusCode(v int32) *GetAccessPolicyConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAccessPolicyConfigResponse) SetBody(v *GetAccessPolicyConfigResponseBody) *GetAccessPolicyConfigResponse {
	s.Body = v
	return s
}

type GetDomainConfigRequest struct {
	CompressEnable *bool   `json:"CompressEnable,omitempty" xml:"CompressEnable,omitempty"`
	DomainCode     *string `json:"DomainCode,omitempty" xml:"DomainCode,omitempty"`
}

func (s GetDomainConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDomainConfigRequest) GoString() string {
	return s.String()
}

func (s *GetDomainConfigRequest) SetCompressEnable(v bool) *GetDomainConfigRequest {
	s.CompressEnable = &v
	return s
}

func (s *GetDomainConfigRequest) SetDomainCode(v string) *GetDomainConfigRequest {
	s.DomainCode = &v
	return s
}

type GetDomainConfigResponseBody struct {
	Compressed *bool   `json:"Compressed,omitempty" xml:"Compressed,omitempty"`
	Data       *string `json:"Data,omitempty" xml:"Data,omitempty"`
	DomainCode *string `json:"DomainCode,omitempty" xml:"DomainCode,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDomainConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDomainConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetDomainConfigResponseBody) SetCompressed(v bool) *GetDomainConfigResponseBody {
	s.Compressed = &v
	return s
}

func (s *GetDomainConfigResponseBody) SetData(v string) *GetDomainConfigResponseBody {
	s.Data = &v
	return s
}

func (s *GetDomainConfigResponseBody) SetDomainCode(v string) *GetDomainConfigResponseBody {
	s.DomainCode = &v
	return s
}

func (s *GetDomainConfigResponseBody) SetRequestId(v string) *GetDomainConfigResponseBody {
	s.RequestId = &v
	return s
}

type GetDomainConfigResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDomainConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDomainConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDomainConfigResponse) GoString() string {
	return s.String()
}

func (s *GetDomainConfigResponse) SetHeaders(v map[string]*string) *GetDomainConfigResponse {
	s.Headers = v
	return s
}

func (s *GetDomainConfigResponse) SetStatusCode(v int32) *GetDomainConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDomainConfigResponse) SetBody(v *GetDomainConfigResponseBody) *GetDomainConfigResponse {
	s.Body = v
	return s
}

type GetDomainPartRequest struct {
	CompressEnable *bool   `json:"CompressEnable,omitempty" xml:"CompressEnable,omitempty"`
	DataType       *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	DomainCode     *string `json:"DomainCode,omitempty" xml:"DomainCode,omitempty"`
	Part           *string `json:"Part,omitempty" xml:"Part,omitempty"`
}

func (s GetDomainPartRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDomainPartRequest) GoString() string {
	return s.String()
}

func (s *GetDomainPartRequest) SetCompressEnable(v bool) *GetDomainPartRequest {
	s.CompressEnable = &v
	return s
}

func (s *GetDomainPartRequest) SetDataType(v string) *GetDomainPartRequest {
	s.DataType = &v
	return s
}

func (s *GetDomainPartRequest) SetDomainCode(v string) *GetDomainPartRequest {
	s.DomainCode = &v
	return s
}

func (s *GetDomainPartRequest) SetPart(v string) *GetDomainPartRequest {
	s.Part = &v
	return s
}

type GetDomainPartResponseBody struct {
	Compressed *bool   `json:"Compressed,omitempty" xml:"Compressed,omitempty"`
	Data       *string `json:"Data,omitempty" xml:"Data,omitempty"`
	DataType   *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	DomainCode *string `json:"DomainCode,omitempty" xml:"DomainCode,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDomainPartResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDomainPartResponseBody) GoString() string {
	return s.String()
}

func (s *GetDomainPartResponseBody) SetCompressed(v bool) *GetDomainPartResponseBody {
	s.Compressed = &v
	return s
}

func (s *GetDomainPartResponseBody) SetData(v string) *GetDomainPartResponseBody {
	s.Data = &v
	return s
}

func (s *GetDomainPartResponseBody) SetDataType(v string) *GetDomainPartResponseBody {
	s.DataType = &v
	return s
}

func (s *GetDomainPartResponseBody) SetDomainCode(v string) *GetDomainPartResponseBody {
	s.DomainCode = &v
	return s
}

func (s *GetDomainPartResponseBody) SetRequestId(v string) *GetDomainPartResponseBody {
	s.RequestId = &v
	return s
}

type GetDomainPartResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDomainPartResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDomainPartResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDomainPartResponse) GoString() string {
	return s.String()
}

func (s *GetDomainPartResponse) SetHeaders(v map[string]*string) *GetDomainPartResponse {
	s.Headers = v
	return s
}

func (s *GetDomainPartResponse) SetStatusCode(v int32) *GetDomainPartResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDomainPartResponse) SetBody(v *GetDomainPartResponseBody) *GetDomainPartResponse {
	s.Body = v
	return s
}

type GetDomainPartByProxyRequest struct {
	CompressEnable *bool   `json:"CompressEnable,omitempty" xml:"CompressEnable,omitempty"`
	DataType       *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	DomainCode     *string `json:"DomainCode,omitempty" xml:"DomainCode,omitempty"`
	Part           *string `json:"Part,omitempty" xml:"Part,omitempty"`
}

func (s GetDomainPartByProxyRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDomainPartByProxyRequest) GoString() string {
	return s.String()
}

func (s *GetDomainPartByProxyRequest) SetCompressEnable(v bool) *GetDomainPartByProxyRequest {
	s.CompressEnable = &v
	return s
}

func (s *GetDomainPartByProxyRequest) SetDataType(v string) *GetDomainPartByProxyRequest {
	s.DataType = &v
	return s
}

func (s *GetDomainPartByProxyRequest) SetDomainCode(v string) *GetDomainPartByProxyRequest {
	s.DomainCode = &v
	return s
}

func (s *GetDomainPartByProxyRequest) SetPart(v string) *GetDomainPartByProxyRequest {
	s.Part = &v
	return s
}

type GetDomainPartByProxyResponseBody struct {
	Compressed *bool   `json:"Compressed,omitempty" xml:"Compressed,omitempty"`
	Data       *string `json:"Data,omitempty" xml:"Data,omitempty"`
	DataType   *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	DomainCode *string `json:"DomainCode,omitempty" xml:"DomainCode,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDomainPartByProxyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDomainPartByProxyResponseBody) GoString() string {
	return s.String()
}

func (s *GetDomainPartByProxyResponseBody) SetCompressed(v bool) *GetDomainPartByProxyResponseBody {
	s.Compressed = &v
	return s
}

func (s *GetDomainPartByProxyResponseBody) SetData(v string) *GetDomainPartByProxyResponseBody {
	s.Data = &v
	return s
}

func (s *GetDomainPartByProxyResponseBody) SetDataType(v string) *GetDomainPartByProxyResponseBody {
	s.DataType = &v
	return s
}

func (s *GetDomainPartByProxyResponseBody) SetDomainCode(v string) *GetDomainPartByProxyResponseBody {
	s.DomainCode = &v
	return s
}

func (s *GetDomainPartByProxyResponseBody) SetRequestId(v string) *GetDomainPartByProxyResponseBody {
	s.RequestId = &v
	return s
}

type GetDomainPartByProxyResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDomainPartByProxyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDomainPartByProxyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDomainPartByProxyResponse) GoString() string {
	return s.String()
}

func (s *GetDomainPartByProxyResponse) SetHeaders(v map[string]*string) *GetDomainPartByProxyResponse {
	s.Headers = v
	return s
}

func (s *GetDomainPartByProxyResponse) SetStatusCode(v int32) *GetDomainPartByProxyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDomainPartByProxyResponse) SetBody(v *GetDomainPartByProxyResponseBody) *GetDomainPartByProxyResponse {
	s.Body = v
	return s
}

type GetDomainPartConfigRequest struct {
	CompressEnable *bool   `json:"CompressEnable,omitempty" xml:"CompressEnable,omitempty"`
	DataType       *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	DomainCode     *string `json:"DomainCode,omitempty" xml:"DomainCode,omitempty"`
	Part           *string `json:"Part,omitempty" xml:"Part,omitempty"`
}

func (s GetDomainPartConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDomainPartConfigRequest) GoString() string {
	return s.String()
}

func (s *GetDomainPartConfigRequest) SetCompressEnable(v bool) *GetDomainPartConfigRequest {
	s.CompressEnable = &v
	return s
}

func (s *GetDomainPartConfigRequest) SetDataType(v string) *GetDomainPartConfigRequest {
	s.DataType = &v
	return s
}

func (s *GetDomainPartConfigRequest) SetDomainCode(v string) *GetDomainPartConfigRequest {
	s.DomainCode = &v
	return s
}

func (s *GetDomainPartConfigRequest) SetPart(v string) *GetDomainPartConfigRequest {
	s.Part = &v
	return s
}

type GetDomainPartConfigResponseBody struct {
	Compressed *bool   `json:"Compressed,omitempty" xml:"Compressed,omitempty"`
	Data       *string `json:"Data,omitempty" xml:"Data,omitempty"`
	DataType   *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	DomainCode *string `json:"DomainCode,omitempty" xml:"DomainCode,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDomainPartConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDomainPartConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetDomainPartConfigResponseBody) SetCompressed(v bool) *GetDomainPartConfigResponseBody {
	s.Compressed = &v
	return s
}

func (s *GetDomainPartConfigResponseBody) SetData(v string) *GetDomainPartConfigResponseBody {
	s.Data = &v
	return s
}

func (s *GetDomainPartConfigResponseBody) SetDataType(v string) *GetDomainPartConfigResponseBody {
	s.DataType = &v
	return s
}

func (s *GetDomainPartConfigResponseBody) SetDomainCode(v string) *GetDomainPartConfigResponseBody {
	s.DomainCode = &v
	return s
}

func (s *GetDomainPartConfigResponseBody) SetRequestId(v string) *GetDomainPartConfigResponseBody {
	s.RequestId = &v
	return s
}

type GetDomainPartConfigResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDomainPartConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDomainPartConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDomainPartConfigResponse) GoString() string {
	return s.String()
}

func (s *GetDomainPartConfigResponse) SetHeaders(v map[string]*string) *GetDomainPartConfigResponse {
	s.Headers = v
	return s
}

func (s *GetDomainPartConfigResponse) SetStatusCode(v int32) *GetDomainPartConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDomainPartConfigResponse) SetBody(v *GetDomainPartConfigResponseBody) *GetDomainPartConfigResponse {
	s.Body = v
	return s
}

type GetFileConfigRequest struct {
	DimensionType *string `json:"DimensionType,omitempty" xml:"DimensionType,omitempty"`
	DomainCode    *string `json:"DomainCode,omitempty" xml:"DomainCode,omitempty"`
}

func (s GetFileConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFileConfigRequest) GoString() string {
	return s.String()
}

func (s *GetFileConfigRequest) SetDimensionType(v string) *GetFileConfigRequest {
	s.DimensionType = &v
	return s
}

func (s *GetFileConfigRequest) SetDomainCode(v string) *GetFileConfigRequest {
	s.DomainCode = &v
	return s
}

type GetFileConfigResponseBody struct {
	Data         *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetFileConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFileConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetFileConfigResponseBody) SetData(v string) *GetFileConfigResponseBody {
	s.Data = &v
	return s
}

func (s *GetFileConfigResponseBody) SetErrorCode(v string) *GetFileConfigResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetFileConfigResponseBody) SetErrorMessage(v string) *GetFileConfigResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetFileConfigResponseBody) SetRequestId(v string) *GetFileConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFileConfigResponseBody) SetSuccess(v bool) *GetFileConfigResponseBody {
	s.Success = &v
	return s
}

type GetFileConfigResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetFileConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetFileConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFileConfigResponse) GoString() string {
	return s.String()
}

func (s *GetFileConfigResponse) SetHeaders(v map[string]*string) *GetFileConfigResponse {
	s.Headers = v
	return s
}

func (s *GetFileConfigResponse) SetStatusCode(v int32) *GetFileConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFileConfigResponse) SetBody(v *GetFileConfigResponseBody) *GetFileConfigResponse {
	s.Body = v
	return s
}

type GetIncrementMeasureDataByProxyRequest struct {
	CompressEnable  *bool   `json:"CompressEnable,omitempty" xml:"CompressEnable,omitempty"`
	DataType        *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	DomainCode      *string `json:"DomainCode,omitempty" xml:"DomainCode,omitempty"`
	ModifyEndTime   *int64  `json:"ModifyEndTime,omitempty" xml:"ModifyEndTime,omitempty"`
	ModifyStartTime *int64  `json:"ModifyStartTime,omitempty" xml:"ModifyStartTime,omitempty"`
	RowKeyMapStr    *string `json:"RowKeyMapStr,omitempty" xml:"RowKeyMapStr,omitempty"`
}

func (s GetIncrementMeasureDataByProxyRequest) String() string {
	return tea.Prettify(s)
}

func (s GetIncrementMeasureDataByProxyRequest) GoString() string {
	return s.String()
}

func (s *GetIncrementMeasureDataByProxyRequest) SetCompressEnable(v bool) *GetIncrementMeasureDataByProxyRequest {
	s.CompressEnable = &v
	return s
}

func (s *GetIncrementMeasureDataByProxyRequest) SetDataType(v string) *GetIncrementMeasureDataByProxyRequest {
	s.DataType = &v
	return s
}

func (s *GetIncrementMeasureDataByProxyRequest) SetDomainCode(v string) *GetIncrementMeasureDataByProxyRequest {
	s.DomainCode = &v
	return s
}

func (s *GetIncrementMeasureDataByProxyRequest) SetModifyEndTime(v int64) *GetIncrementMeasureDataByProxyRequest {
	s.ModifyEndTime = &v
	return s
}

func (s *GetIncrementMeasureDataByProxyRequest) SetModifyStartTime(v int64) *GetIncrementMeasureDataByProxyRequest {
	s.ModifyStartTime = &v
	return s
}

func (s *GetIncrementMeasureDataByProxyRequest) SetRowKeyMapStr(v string) *GetIncrementMeasureDataByProxyRequest {
	s.RowKeyMapStr = &v
	return s
}

type GetIncrementMeasureDataByProxyResponseBody struct {
	Compressed *string `json:"Compressed,omitempty" xml:"Compressed,omitempty"`
	Data       *string `json:"Data,omitempty" xml:"Data,omitempty"`
	DataType   *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	DomainCode *string `json:"DomainCode,omitempty" xml:"DomainCode,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetIncrementMeasureDataByProxyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetIncrementMeasureDataByProxyResponseBody) GoString() string {
	return s.String()
}

func (s *GetIncrementMeasureDataByProxyResponseBody) SetCompressed(v string) *GetIncrementMeasureDataByProxyResponseBody {
	s.Compressed = &v
	return s
}

func (s *GetIncrementMeasureDataByProxyResponseBody) SetData(v string) *GetIncrementMeasureDataByProxyResponseBody {
	s.Data = &v
	return s
}

func (s *GetIncrementMeasureDataByProxyResponseBody) SetDataType(v string) *GetIncrementMeasureDataByProxyResponseBody {
	s.DataType = &v
	return s
}

func (s *GetIncrementMeasureDataByProxyResponseBody) SetDomainCode(v string) *GetIncrementMeasureDataByProxyResponseBody {
	s.DomainCode = &v
	return s
}

func (s *GetIncrementMeasureDataByProxyResponseBody) SetRequestId(v string) *GetIncrementMeasureDataByProxyResponseBody {
	s.RequestId = &v
	return s
}

type GetIncrementMeasureDataByProxyResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetIncrementMeasureDataByProxyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetIncrementMeasureDataByProxyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetIncrementMeasureDataByProxyResponse) GoString() string {
	return s.String()
}

func (s *GetIncrementMeasureDataByProxyResponse) SetHeaders(v map[string]*string) *GetIncrementMeasureDataByProxyResponse {
	s.Headers = v
	return s
}

func (s *GetIncrementMeasureDataByProxyResponse) SetStatusCode(v int32) *GetIncrementMeasureDataByProxyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIncrementMeasureDataByProxyResponse) SetBody(v *GetIncrementMeasureDataByProxyResponseBody) *GetIncrementMeasureDataByProxyResponse {
	s.Body = v
	return s
}

type GetMeasureDataRequest struct {
	ApiType        *string `json:"ApiType,omitempty" xml:"ApiType,omitempty"`
	CompressEnable *bool   `json:"CompressEnable,omitempty" xml:"CompressEnable,omitempty"`
	DataType       *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	DomainCode     *string `json:"DomainCode,omitempty" xml:"DomainCode,omitempty"`
	Filter         *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	MaxResult      *int32  `json:"MaxResult,omitempty" xml:"MaxResult,omitempty"`
	NextToken      *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	QueryField     *string `json:"QueryField,omitempty" xml:"QueryField,omitempty"`
}

func (s GetMeasureDataRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMeasureDataRequest) GoString() string {
	return s.String()
}

func (s *GetMeasureDataRequest) SetApiType(v string) *GetMeasureDataRequest {
	s.ApiType = &v
	return s
}

func (s *GetMeasureDataRequest) SetCompressEnable(v bool) *GetMeasureDataRequest {
	s.CompressEnable = &v
	return s
}

func (s *GetMeasureDataRequest) SetDataType(v string) *GetMeasureDataRequest {
	s.DataType = &v
	return s
}

func (s *GetMeasureDataRequest) SetDomainCode(v string) *GetMeasureDataRequest {
	s.DomainCode = &v
	return s
}

func (s *GetMeasureDataRequest) SetFilter(v string) *GetMeasureDataRequest {
	s.Filter = &v
	return s
}

func (s *GetMeasureDataRequest) SetMaxResult(v int32) *GetMeasureDataRequest {
	s.MaxResult = &v
	return s
}

func (s *GetMeasureDataRequest) SetNextToken(v string) *GetMeasureDataRequest {
	s.NextToken = &v
	return s
}

func (s *GetMeasureDataRequest) SetQueryField(v string) *GetMeasureDataRequest {
	s.QueryField = &v
	return s
}

type GetMeasureDataResponseBody struct {
	ApiType    *string `json:"ApiType,omitempty" xml:"ApiType,omitempty"`
	Compressed *bool   `json:"Compressed,omitempty" xml:"Compressed,omitempty"`
	Data       *string `json:"Data,omitempty" xml:"Data,omitempty"`
	DataType   *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	DomainCode *string `json:"DomainCode,omitempty" xml:"DomainCode,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMeasureDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMeasureDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetMeasureDataResponseBody) SetApiType(v string) *GetMeasureDataResponseBody {
	s.ApiType = &v
	return s
}

func (s *GetMeasureDataResponseBody) SetCompressed(v bool) *GetMeasureDataResponseBody {
	s.Compressed = &v
	return s
}

func (s *GetMeasureDataResponseBody) SetData(v string) *GetMeasureDataResponseBody {
	s.Data = &v
	return s
}

func (s *GetMeasureDataResponseBody) SetDataType(v string) *GetMeasureDataResponseBody {
	s.DataType = &v
	return s
}

func (s *GetMeasureDataResponseBody) SetDomainCode(v string) *GetMeasureDataResponseBody {
	s.DomainCode = &v
	return s
}

func (s *GetMeasureDataResponseBody) SetNextToken(v string) *GetMeasureDataResponseBody {
	s.NextToken = &v
	return s
}

func (s *GetMeasureDataResponseBody) SetRequestId(v string) *GetMeasureDataResponseBody {
	s.RequestId = &v
	return s
}

type GetMeasureDataResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetMeasureDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMeasureDataResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMeasureDataResponse) GoString() string {
	return s.String()
}

func (s *GetMeasureDataResponse) SetHeaders(v map[string]*string) *GetMeasureDataResponse {
	s.Headers = v
	return s
}

func (s *GetMeasureDataResponse) SetStatusCode(v int32) *GetMeasureDataResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMeasureDataResponse) SetBody(v *GetMeasureDataResponseBody) *GetMeasureDataResponse {
	s.Body = v
	return s
}

type GetOpenApiConfigRequest struct {
	CompressEnable *bool   `json:"CompressEnable,omitempty" xml:"CompressEnable,omitempty"`
	DataType       *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	ProductName    *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	TableName      *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	SiteBid        *string `json:"siteBid,omitempty" xml:"siteBid,omitempty"`
}

func (s GetOpenApiConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOpenApiConfigRequest) GoString() string {
	return s.String()
}

func (s *GetOpenApiConfigRequest) SetCompressEnable(v bool) *GetOpenApiConfigRequest {
	s.CompressEnable = &v
	return s
}

func (s *GetOpenApiConfigRequest) SetDataType(v string) *GetOpenApiConfigRequest {
	s.DataType = &v
	return s
}

func (s *GetOpenApiConfigRequest) SetProductName(v string) *GetOpenApiConfigRequest {
	s.ProductName = &v
	return s
}

func (s *GetOpenApiConfigRequest) SetTableName(v string) *GetOpenApiConfigRequest {
	s.TableName = &v
	return s
}

func (s *GetOpenApiConfigRequest) SetSiteBid(v string) *GetOpenApiConfigRequest {
	s.SiteBid = &v
	return s
}

type GetOpenApiConfigResponseBody struct {
	Compressed  *bool   `json:"Compressed,omitempty" xml:"Compressed,omitempty"`
	Data        *string `json:"Data,omitempty" xml:"Data,omitempty"`
	DataType    *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SiteBid     *string `json:"SiteBid,omitempty" xml:"SiteBid,omitempty"`
	TableName   *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s GetOpenApiConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOpenApiConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetOpenApiConfigResponseBody) SetCompressed(v bool) *GetOpenApiConfigResponseBody {
	s.Compressed = &v
	return s
}

func (s *GetOpenApiConfigResponseBody) SetData(v string) *GetOpenApiConfigResponseBody {
	s.Data = &v
	return s
}

func (s *GetOpenApiConfigResponseBody) SetDataType(v string) *GetOpenApiConfigResponseBody {
	s.DataType = &v
	return s
}

func (s *GetOpenApiConfigResponseBody) SetProductName(v string) *GetOpenApiConfigResponseBody {
	s.ProductName = &v
	return s
}

func (s *GetOpenApiConfigResponseBody) SetRequestId(v string) *GetOpenApiConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOpenApiConfigResponseBody) SetSiteBid(v string) *GetOpenApiConfigResponseBody {
	s.SiteBid = &v
	return s
}

func (s *GetOpenApiConfigResponseBody) SetTableName(v string) *GetOpenApiConfigResponseBody {
	s.TableName = &v
	return s
}

type GetOpenApiConfigResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOpenApiConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOpenApiConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOpenApiConfigResponse) GoString() string {
	return s.String()
}

func (s *GetOpenApiConfigResponse) SetHeaders(v map[string]*string) *GetOpenApiConfigResponse {
	s.Headers = v
	return s
}

func (s *GetOpenApiConfigResponse) SetStatusCode(v int32) *GetOpenApiConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOpenApiConfigResponse) SetBody(v *GetOpenApiConfigResponseBody) *GetOpenApiConfigResponse {
	s.Body = v
	return s
}

type GetReadyFlagRequest struct {
	ApiType        *string `json:"ApiType,omitempty" xml:"ApiType,omitempty"`
	CompressEnable *bool   `json:"CompressEnable,omitempty" xml:"CompressEnable,omitempty"`
	DataType       *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	DomainCode     *string `json:"DomainCode,omitempty" xml:"DomainCode,omitempty"`
	Filter         *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	MaxResult      *int32  `json:"MaxResult,omitempty" xml:"MaxResult,omitempty"`
	NextToken      *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s GetReadyFlagRequest) String() string {
	return tea.Prettify(s)
}

func (s GetReadyFlagRequest) GoString() string {
	return s.String()
}

func (s *GetReadyFlagRequest) SetApiType(v string) *GetReadyFlagRequest {
	s.ApiType = &v
	return s
}

func (s *GetReadyFlagRequest) SetCompressEnable(v bool) *GetReadyFlagRequest {
	s.CompressEnable = &v
	return s
}

func (s *GetReadyFlagRequest) SetDataType(v string) *GetReadyFlagRequest {
	s.DataType = &v
	return s
}

func (s *GetReadyFlagRequest) SetDomainCode(v string) *GetReadyFlagRequest {
	s.DomainCode = &v
	return s
}

func (s *GetReadyFlagRequest) SetFilter(v string) *GetReadyFlagRequest {
	s.Filter = &v
	return s
}

func (s *GetReadyFlagRequest) SetMaxResult(v int32) *GetReadyFlagRequest {
	s.MaxResult = &v
	return s
}

func (s *GetReadyFlagRequest) SetNextToken(v string) *GetReadyFlagRequest {
	s.NextToken = &v
	return s
}

type GetReadyFlagResponseBody struct {
	ApiType    *string `json:"ApiType,omitempty" xml:"ApiType,omitempty"`
	Compressed *bool   `json:"Compressed,omitempty" xml:"Compressed,omitempty"`
	Data       *string `json:"Data,omitempty" xml:"Data,omitempty"`
	DataType   *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	DomainCode *string `json:"DomainCode,omitempty" xml:"DomainCode,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetReadyFlagResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetReadyFlagResponseBody) GoString() string {
	return s.String()
}

func (s *GetReadyFlagResponseBody) SetApiType(v string) *GetReadyFlagResponseBody {
	s.ApiType = &v
	return s
}

func (s *GetReadyFlagResponseBody) SetCompressed(v bool) *GetReadyFlagResponseBody {
	s.Compressed = &v
	return s
}

func (s *GetReadyFlagResponseBody) SetData(v string) *GetReadyFlagResponseBody {
	s.Data = &v
	return s
}

func (s *GetReadyFlagResponseBody) SetDataType(v string) *GetReadyFlagResponseBody {
	s.DataType = &v
	return s
}

func (s *GetReadyFlagResponseBody) SetDomainCode(v string) *GetReadyFlagResponseBody {
	s.DomainCode = &v
	return s
}

func (s *GetReadyFlagResponseBody) SetNextToken(v string) *GetReadyFlagResponseBody {
	s.NextToken = &v
	return s
}

func (s *GetReadyFlagResponseBody) SetRequestId(v string) *GetReadyFlagResponseBody {
	s.RequestId = &v
	return s
}

type GetReadyFlagResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetReadyFlagResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetReadyFlagResponse) String() string {
	return tea.Prettify(s)
}

func (s GetReadyFlagResponse) GoString() string {
	return s.String()
}

func (s *GetReadyFlagResponse) SetHeaders(v map[string]*string) *GetReadyFlagResponse {
	s.Headers = v
	return s
}

func (s *GetReadyFlagResponse) SetStatusCode(v int32) *GetReadyFlagResponse {
	s.StatusCode = &v
	return s
}

func (s *GetReadyFlagResponse) SetBody(v *GetReadyFlagResponseBody) *GetReadyFlagResponse {
	s.Body = v
	return s
}

type GetReadyFlagAlertConfigRequest struct {
	CompressEnable *bool   `json:"CompressEnable,omitempty" xml:"CompressEnable,omitempty"`
	DataType       *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	DomainCode     *string `json:"DomainCode,omitempty" xml:"DomainCode,omitempty"`
}

func (s GetReadyFlagAlertConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s GetReadyFlagAlertConfigRequest) GoString() string {
	return s.String()
}

func (s *GetReadyFlagAlertConfigRequest) SetCompressEnable(v bool) *GetReadyFlagAlertConfigRequest {
	s.CompressEnable = &v
	return s
}

func (s *GetReadyFlagAlertConfigRequest) SetDataType(v string) *GetReadyFlagAlertConfigRequest {
	s.DataType = &v
	return s
}

func (s *GetReadyFlagAlertConfigRequest) SetDomainCode(v string) *GetReadyFlagAlertConfigRequest {
	s.DomainCode = &v
	return s
}

type GetReadyFlagAlertConfigResponseBody struct {
	Compressed *bool   `json:"Compressed,omitempty" xml:"Compressed,omitempty"`
	Data       *string `json:"Data,omitempty" xml:"Data,omitempty"`
	DataType   *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	DomainCode *string `json:"DomainCode,omitempty" xml:"DomainCode,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetReadyFlagAlertConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetReadyFlagAlertConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetReadyFlagAlertConfigResponseBody) SetCompressed(v bool) *GetReadyFlagAlertConfigResponseBody {
	s.Compressed = &v
	return s
}

func (s *GetReadyFlagAlertConfigResponseBody) SetData(v string) *GetReadyFlagAlertConfigResponseBody {
	s.Data = &v
	return s
}

func (s *GetReadyFlagAlertConfigResponseBody) SetDataType(v string) *GetReadyFlagAlertConfigResponseBody {
	s.DataType = &v
	return s
}

func (s *GetReadyFlagAlertConfigResponseBody) SetDomainCode(v string) *GetReadyFlagAlertConfigResponseBody {
	s.DomainCode = &v
	return s
}

func (s *GetReadyFlagAlertConfigResponseBody) SetRequestId(v string) *GetReadyFlagAlertConfigResponseBody {
	s.RequestId = &v
	return s
}

type GetReadyFlagAlertConfigResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetReadyFlagAlertConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetReadyFlagAlertConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetReadyFlagAlertConfigResponse) GoString() string {
	return s.String()
}

func (s *GetReadyFlagAlertConfigResponse) SetHeaders(v map[string]*string) *GetReadyFlagAlertConfigResponse {
	s.Headers = v
	return s
}

func (s *GetReadyFlagAlertConfigResponse) SetStatusCode(v int32) *GetReadyFlagAlertConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetReadyFlagAlertConfigResponse) SetBody(v *GetReadyFlagAlertConfigResponseBody) *GetReadyFlagAlertConfigResponse {
	s.Body = v
	return s
}

type GetReadyFlagByProxyRequest struct {
	ApiType        *string `json:"ApiType,omitempty" xml:"ApiType,omitempty"`
	CompressEnable *bool   `json:"CompressEnable,omitempty" xml:"CompressEnable,omitempty"`
	DataType       *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	DomainCode     *string `json:"DomainCode,omitempty" xml:"DomainCode,omitempty"`
	Filter         *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	MaxResult      *int32  `json:"MaxResult,omitempty" xml:"MaxResult,omitempty"`
	NextToken      *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s GetReadyFlagByProxyRequest) String() string {
	return tea.Prettify(s)
}

func (s GetReadyFlagByProxyRequest) GoString() string {
	return s.String()
}

func (s *GetReadyFlagByProxyRequest) SetApiType(v string) *GetReadyFlagByProxyRequest {
	s.ApiType = &v
	return s
}

func (s *GetReadyFlagByProxyRequest) SetCompressEnable(v bool) *GetReadyFlagByProxyRequest {
	s.CompressEnable = &v
	return s
}

func (s *GetReadyFlagByProxyRequest) SetDataType(v string) *GetReadyFlagByProxyRequest {
	s.DataType = &v
	return s
}

func (s *GetReadyFlagByProxyRequest) SetDomainCode(v string) *GetReadyFlagByProxyRequest {
	s.DomainCode = &v
	return s
}

func (s *GetReadyFlagByProxyRequest) SetFilter(v string) *GetReadyFlagByProxyRequest {
	s.Filter = &v
	return s
}

func (s *GetReadyFlagByProxyRequest) SetMaxResult(v int32) *GetReadyFlagByProxyRequest {
	s.MaxResult = &v
	return s
}

func (s *GetReadyFlagByProxyRequest) SetNextToken(v string) *GetReadyFlagByProxyRequest {
	s.NextToken = &v
	return s
}

type GetReadyFlagByProxyResponseBody struct {
	ApiType    *string `json:"ApiType,omitempty" xml:"ApiType,omitempty"`
	Compressed *bool   `json:"Compressed,omitempty" xml:"Compressed,omitempty"`
	Data       *string `json:"Data,omitempty" xml:"Data,omitempty"`
	DataType   *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	DomainCode *string `json:"DomainCode,omitempty" xml:"DomainCode,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetReadyFlagByProxyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetReadyFlagByProxyResponseBody) GoString() string {
	return s.String()
}

func (s *GetReadyFlagByProxyResponseBody) SetApiType(v string) *GetReadyFlagByProxyResponseBody {
	s.ApiType = &v
	return s
}

func (s *GetReadyFlagByProxyResponseBody) SetCompressed(v bool) *GetReadyFlagByProxyResponseBody {
	s.Compressed = &v
	return s
}

func (s *GetReadyFlagByProxyResponseBody) SetData(v string) *GetReadyFlagByProxyResponseBody {
	s.Data = &v
	return s
}

func (s *GetReadyFlagByProxyResponseBody) SetDataType(v string) *GetReadyFlagByProxyResponseBody {
	s.DataType = &v
	return s
}

func (s *GetReadyFlagByProxyResponseBody) SetDomainCode(v string) *GetReadyFlagByProxyResponseBody {
	s.DomainCode = &v
	return s
}

func (s *GetReadyFlagByProxyResponseBody) SetNextToken(v string) *GetReadyFlagByProxyResponseBody {
	s.NextToken = &v
	return s
}

func (s *GetReadyFlagByProxyResponseBody) SetRequestId(v string) *GetReadyFlagByProxyResponseBody {
	s.RequestId = &v
	return s
}

type GetReadyFlagByProxyResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetReadyFlagByProxyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetReadyFlagByProxyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetReadyFlagByProxyResponse) GoString() string {
	return s.String()
}

func (s *GetReadyFlagByProxyResponse) SetHeaders(v map[string]*string) *GetReadyFlagByProxyResponse {
	s.Headers = v
	return s
}

func (s *GetReadyFlagByProxyResponse) SetStatusCode(v int32) *GetReadyFlagByProxyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetReadyFlagByProxyResponse) SetBody(v *GetReadyFlagByProxyResponseBody) *GetReadyFlagByProxyResponse {
	s.Body = v
	return s
}

type PutDomainPartRequest struct {
	Compressed *bool   `json:"Compressed,omitempty" xml:"Compressed,omitempty"`
	Data       *string `json:"Data,omitempty" xml:"Data,omitempty"`
	DataType   *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	DomainCode *string `json:"DomainCode,omitempty" xml:"DomainCode,omitempty"`
}

func (s PutDomainPartRequest) String() string {
	return tea.Prettify(s)
}

func (s PutDomainPartRequest) GoString() string {
	return s.String()
}

func (s *PutDomainPartRequest) SetCompressed(v bool) *PutDomainPartRequest {
	s.Compressed = &v
	return s
}

func (s *PutDomainPartRequest) SetData(v string) *PutDomainPartRequest {
	s.Data = &v
	return s
}

func (s *PutDomainPartRequest) SetDataType(v string) *PutDomainPartRequest {
	s.DataType = &v
	return s
}

func (s *PutDomainPartRequest) SetDomainCode(v string) *PutDomainPartRequest {
	s.DomainCode = &v
	return s
}

type PutDomainPartResponseBody struct {
	DataType   *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	DomainCode *string `json:"DomainCode,omitempty" xml:"DomainCode,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PutDomainPartResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PutDomainPartResponseBody) GoString() string {
	return s.String()
}

func (s *PutDomainPartResponseBody) SetDataType(v string) *PutDomainPartResponseBody {
	s.DataType = &v
	return s
}

func (s *PutDomainPartResponseBody) SetDomainCode(v string) *PutDomainPartResponseBody {
	s.DomainCode = &v
	return s
}

func (s *PutDomainPartResponseBody) SetRequestId(v string) *PutDomainPartResponseBody {
	s.RequestId = &v
	return s
}

type PutDomainPartResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PutDomainPartResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PutDomainPartResponse) String() string {
	return tea.Prettify(s)
}

func (s PutDomainPartResponse) GoString() string {
	return s.String()
}

func (s *PutDomainPartResponse) SetHeaders(v map[string]*string) *PutDomainPartResponse {
	s.Headers = v
	return s
}

func (s *PutDomainPartResponse) SetStatusCode(v int32) *PutDomainPartResponse {
	s.StatusCode = &v
	return s
}

func (s *PutDomainPartResponse) SetBody(v *PutDomainPartResponseBody) *PutDomainPartResponse {
	s.Body = v
	return s
}

type PutDomainPartByProxyRequest struct {
	Compressed *bool   `json:"Compressed,omitempty" xml:"Compressed,omitempty"`
	Data       *string `json:"Data,omitempty" xml:"Data,omitempty"`
	DataType   *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	DomainCode *string `json:"DomainCode,omitempty" xml:"DomainCode,omitempty"`
}

func (s PutDomainPartByProxyRequest) String() string {
	return tea.Prettify(s)
}

func (s PutDomainPartByProxyRequest) GoString() string {
	return s.String()
}

func (s *PutDomainPartByProxyRequest) SetCompressed(v bool) *PutDomainPartByProxyRequest {
	s.Compressed = &v
	return s
}

func (s *PutDomainPartByProxyRequest) SetData(v string) *PutDomainPartByProxyRequest {
	s.Data = &v
	return s
}

func (s *PutDomainPartByProxyRequest) SetDataType(v string) *PutDomainPartByProxyRequest {
	s.DataType = &v
	return s
}

func (s *PutDomainPartByProxyRequest) SetDomainCode(v string) *PutDomainPartByProxyRequest {
	s.DomainCode = &v
	return s
}

type PutDomainPartByProxyResponseBody struct {
	DataType   *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	DomainCode *string `json:"DomainCode,omitempty" xml:"DomainCode,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PutDomainPartByProxyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PutDomainPartByProxyResponseBody) GoString() string {
	return s.String()
}

func (s *PutDomainPartByProxyResponseBody) SetDataType(v string) *PutDomainPartByProxyResponseBody {
	s.DataType = &v
	return s
}

func (s *PutDomainPartByProxyResponseBody) SetDomainCode(v string) *PutDomainPartByProxyResponseBody {
	s.DomainCode = &v
	return s
}

func (s *PutDomainPartByProxyResponseBody) SetRequestId(v string) *PutDomainPartByProxyResponseBody {
	s.RequestId = &v
	return s
}

type PutDomainPartByProxyResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PutDomainPartByProxyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PutDomainPartByProxyResponse) String() string {
	return tea.Prettify(s)
}

func (s PutDomainPartByProxyResponse) GoString() string {
	return s.String()
}

func (s *PutDomainPartByProxyResponse) SetHeaders(v map[string]*string) *PutDomainPartByProxyResponse {
	s.Headers = v
	return s
}

func (s *PutDomainPartByProxyResponse) SetStatusCode(v int32) *PutDomainPartByProxyResponse {
	s.StatusCode = &v
	return s
}

func (s *PutDomainPartByProxyResponse) SetBody(v *PutDomainPartByProxyResponseBody) *PutDomainPartByProxyResponse {
	s.Body = v
	return s
}

type PutMeasureDataRequest struct {
	ApiType         *string `json:"ApiType,omitempty" xml:"ApiType,omitempty"`
	Compressed      *bool   `json:"Compressed,omitempty" xml:"Compressed,omitempty"`
	Data            *string `json:"Data,omitempty" xml:"Data,omitempty"`
	DataType        *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	DomainCode      *string `json:"DomainCode,omitempty" xml:"DomainCode,omitempty"`
	Filter          *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	SourceRequestId *string `json:"SourceRequestId,omitempty" xml:"SourceRequestId,omitempty"`
}

func (s PutMeasureDataRequest) String() string {
	return tea.Prettify(s)
}

func (s PutMeasureDataRequest) GoString() string {
	return s.String()
}

func (s *PutMeasureDataRequest) SetApiType(v string) *PutMeasureDataRequest {
	s.ApiType = &v
	return s
}

func (s *PutMeasureDataRequest) SetCompressed(v bool) *PutMeasureDataRequest {
	s.Compressed = &v
	return s
}

func (s *PutMeasureDataRequest) SetData(v string) *PutMeasureDataRequest {
	s.Data = &v
	return s
}

func (s *PutMeasureDataRequest) SetDataType(v string) *PutMeasureDataRequest {
	s.DataType = &v
	return s
}

func (s *PutMeasureDataRequest) SetDomainCode(v string) *PutMeasureDataRequest {
	s.DomainCode = &v
	return s
}

func (s *PutMeasureDataRequest) SetFilter(v string) *PutMeasureDataRequest {
	s.Filter = &v
	return s
}

func (s *PutMeasureDataRequest) SetSourceRequestId(v string) *PutMeasureDataRequest {
	s.SourceRequestId = &v
	return s
}

type PutMeasureDataResponseBody struct {
	ApiType         *string `json:"ApiType,omitempty" xml:"ApiType,omitempty"`
	DataType        *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	DomainCode      *string `json:"DomainCode,omitempty" xml:"DomainCode,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SourceRequestId *string `json:"SourceRequestId,omitempty" xml:"SourceRequestId,omitempty"`
}

func (s PutMeasureDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PutMeasureDataResponseBody) GoString() string {
	return s.String()
}

func (s *PutMeasureDataResponseBody) SetApiType(v string) *PutMeasureDataResponseBody {
	s.ApiType = &v
	return s
}

func (s *PutMeasureDataResponseBody) SetDataType(v string) *PutMeasureDataResponseBody {
	s.DataType = &v
	return s
}

func (s *PutMeasureDataResponseBody) SetDomainCode(v string) *PutMeasureDataResponseBody {
	s.DomainCode = &v
	return s
}

func (s *PutMeasureDataResponseBody) SetRequestId(v string) *PutMeasureDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutMeasureDataResponseBody) SetSourceRequestId(v string) *PutMeasureDataResponseBody {
	s.SourceRequestId = &v
	return s
}

type PutMeasureDataResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PutMeasureDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PutMeasureDataResponse) String() string {
	return tea.Prettify(s)
}

func (s PutMeasureDataResponse) GoString() string {
	return s.String()
}

func (s *PutMeasureDataResponse) SetHeaders(v map[string]*string) *PutMeasureDataResponse {
	s.Headers = v
	return s
}

func (s *PutMeasureDataResponse) SetStatusCode(v int32) *PutMeasureDataResponse {
	s.StatusCode = &v
	return s
}

func (s *PutMeasureDataResponse) SetBody(v *PutMeasureDataResponseBody) *PutMeasureDataResponse {
	s.Body = v
	return s
}

type PutMeasureDataByProxyRequest struct {
	ApiType    *string `json:"ApiType,omitempty" xml:"ApiType,omitempty"`
	Compressed *bool   `json:"Compressed,omitempty" xml:"Compressed,omitempty"`
	Data       *string `json:"Data,omitempty" xml:"Data,omitempty"`
	DataType   *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	DomainCode *string `json:"DomainCode,omitempty" xml:"DomainCode,omitempty"`
	Filter     *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
}

func (s PutMeasureDataByProxyRequest) String() string {
	return tea.Prettify(s)
}

func (s PutMeasureDataByProxyRequest) GoString() string {
	return s.String()
}

func (s *PutMeasureDataByProxyRequest) SetApiType(v string) *PutMeasureDataByProxyRequest {
	s.ApiType = &v
	return s
}

func (s *PutMeasureDataByProxyRequest) SetCompressed(v bool) *PutMeasureDataByProxyRequest {
	s.Compressed = &v
	return s
}

func (s *PutMeasureDataByProxyRequest) SetData(v string) *PutMeasureDataByProxyRequest {
	s.Data = &v
	return s
}

func (s *PutMeasureDataByProxyRequest) SetDataType(v string) *PutMeasureDataByProxyRequest {
	s.DataType = &v
	return s
}

func (s *PutMeasureDataByProxyRequest) SetDomainCode(v string) *PutMeasureDataByProxyRequest {
	s.DomainCode = &v
	return s
}

func (s *PutMeasureDataByProxyRequest) SetFilter(v string) *PutMeasureDataByProxyRequest {
	s.Filter = &v
	return s
}

type PutMeasureDataByProxyResponseBody struct {
	ApiType    *string `json:"ApiType,omitempty" xml:"ApiType,omitempty"`
	DataType   *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	DomainCode *string `json:"DomainCode,omitempty" xml:"DomainCode,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PutMeasureDataByProxyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PutMeasureDataByProxyResponseBody) GoString() string {
	return s.String()
}

func (s *PutMeasureDataByProxyResponseBody) SetApiType(v string) *PutMeasureDataByProxyResponseBody {
	s.ApiType = &v
	return s
}

func (s *PutMeasureDataByProxyResponseBody) SetDataType(v string) *PutMeasureDataByProxyResponseBody {
	s.DataType = &v
	return s
}

func (s *PutMeasureDataByProxyResponseBody) SetDomainCode(v string) *PutMeasureDataByProxyResponseBody {
	s.DomainCode = &v
	return s
}

func (s *PutMeasureDataByProxyResponseBody) SetRequestId(v string) *PutMeasureDataByProxyResponseBody {
	s.RequestId = &v
	return s
}

type PutMeasureDataByProxyResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PutMeasureDataByProxyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PutMeasureDataByProxyResponse) String() string {
	return tea.Prettify(s)
}

func (s PutMeasureDataByProxyResponse) GoString() string {
	return s.String()
}

func (s *PutMeasureDataByProxyResponse) SetHeaders(v map[string]*string) *PutMeasureDataByProxyResponse {
	s.Headers = v
	return s
}

func (s *PutMeasureDataByProxyResponse) SetStatusCode(v int32) *PutMeasureDataByProxyResponse {
	s.StatusCode = &v
	return s
}

func (s *PutMeasureDataByProxyResponse) SetBody(v *PutMeasureDataByProxyResponseBody) *PutMeasureDataByProxyResponse {
	s.Body = v
	return s
}

type PutReadyFlagRequest struct {
	ApiType         *string `json:"ApiType,omitempty" xml:"ApiType,omitempty"`
	Compressed      *bool   `json:"Compressed,omitempty" xml:"Compressed,omitempty"`
	Data            *string `json:"Data,omitempty" xml:"Data,omitempty"`
	DataType        *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	DomainCode      *string `json:"DomainCode,omitempty" xml:"DomainCode,omitempty"`
	SourceRequestId *string `json:"SourceRequestId,omitempty" xml:"SourceRequestId,omitempty"`
}

func (s PutReadyFlagRequest) String() string {
	return tea.Prettify(s)
}

func (s PutReadyFlagRequest) GoString() string {
	return s.String()
}

func (s *PutReadyFlagRequest) SetApiType(v string) *PutReadyFlagRequest {
	s.ApiType = &v
	return s
}

func (s *PutReadyFlagRequest) SetCompressed(v bool) *PutReadyFlagRequest {
	s.Compressed = &v
	return s
}

func (s *PutReadyFlagRequest) SetData(v string) *PutReadyFlagRequest {
	s.Data = &v
	return s
}

func (s *PutReadyFlagRequest) SetDataType(v string) *PutReadyFlagRequest {
	s.DataType = &v
	return s
}

func (s *PutReadyFlagRequest) SetDomainCode(v string) *PutReadyFlagRequest {
	s.DomainCode = &v
	return s
}

func (s *PutReadyFlagRequest) SetSourceRequestId(v string) *PutReadyFlagRequest {
	s.SourceRequestId = &v
	return s
}

type PutReadyFlagResponseBody struct {
	ApiType         *string `json:"ApiType,omitempty" xml:"ApiType,omitempty"`
	DataType        *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	DomainCode      *string `json:"DomainCode,omitempty" xml:"DomainCode,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SourceRequestId *string `json:"SourceRequestId,omitempty" xml:"SourceRequestId,omitempty"`
}

func (s PutReadyFlagResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PutReadyFlagResponseBody) GoString() string {
	return s.String()
}

func (s *PutReadyFlagResponseBody) SetApiType(v string) *PutReadyFlagResponseBody {
	s.ApiType = &v
	return s
}

func (s *PutReadyFlagResponseBody) SetDataType(v string) *PutReadyFlagResponseBody {
	s.DataType = &v
	return s
}

func (s *PutReadyFlagResponseBody) SetDomainCode(v string) *PutReadyFlagResponseBody {
	s.DomainCode = &v
	return s
}

func (s *PutReadyFlagResponseBody) SetRequestId(v string) *PutReadyFlagResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutReadyFlagResponseBody) SetSourceRequestId(v string) *PutReadyFlagResponseBody {
	s.SourceRequestId = &v
	return s
}

type PutReadyFlagResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PutReadyFlagResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PutReadyFlagResponse) String() string {
	return tea.Prettify(s)
}

func (s PutReadyFlagResponse) GoString() string {
	return s.String()
}

func (s *PutReadyFlagResponse) SetHeaders(v map[string]*string) *PutReadyFlagResponse {
	s.Headers = v
	return s
}

func (s *PutReadyFlagResponse) SetStatusCode(v int32) *PutReadyFlagResponse {
	s.StatusCode = &v
	return s
}

func (s *PutReadyFlagResponse) SetBody(v *PutReadyFlagResponseBody) *PutReadyFlagResponse {
	s.Body = v
	return s
}

type PutReadyFlagByProxyRequest struct {
	ApiType    *string `json:"ApiType,omitempty" xml:"ApiType,omitempty"`
	Compressed *bool   `json:"Compressed,omitempty" xml:"Compressed,omitempty"`
	Data       *string `json:"Data,omitempty" xml:"Data,omitempty"`
	DataType   *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	DomainCode *string `json:"DomainCode,omitempty" xml:"DomainCode,omitempty"`
}

func (s PutReadyFlagByProxyRequest) String() string {
	return tea.Prettify(s)
}

func (s PutReadyFlagByProxyRequest) GoString() string {
	return s.String()
}

func (s *PutReadyFlagByProxyRequest) SetApiType(v string) *PutReadyFlagByProxyRequest {
	s.ApiType = &v
	return s
}

func (s *PutReadyFlagByProxyRequest) SetCompressed(v bool) *PutReadyFlagByProxyRequest {
	s.Compressed = &v
	return s
}

func (s *PutReadyFlagByProxyRequest) SetData(v string) *PutReadyFlagByProxyRequest {
	s.Data = &v
	return s
}

func (s *PutReadyFlagByProxyRequest) SetDataType(v string) *PutReadyFlagByProxyRequest {
	s.DataType = &v
	return s
}

func (s *PutReadyFlagByProxyRequest) SetDomainCode(v string) *PutReadyFlagByProxyRequest {
	s.DomainCode = &v
	return s
}

type PutReadyFlagByProxyResponseBody struct {
	ApiType    *string `json:"ApiType,omitempty" xml:"ApiType,omitempty"`
	DataType   *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	DomainCode *string `json:"DomainCode,omitempty" xml:"DomainCode,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PutReadyFlagByProxyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PutReadyFlagByProxyResponseBody) GoString() string {
	return s.String()
}

func (s *PutReadyFlagByProxyResponseBody) SetApiType(v string) *PutReadyFlagByProxyResponseBody {
	s.ApiType = &v
	return s
}

func (s *PutReadyFlagByProxyResponseBody) SetDataType(v string) *PutReadyFlagByProxyResponseBody {
	s.DataType = &v
	return s
}

func (s *PutReadyFlagByProxyResponseBody) SetDomainCode(v string) *PutReadyFlagByProxyResponseBody {
	s.DomainCode = &v
	return s
}

func (s *PutReadyFlagByProxyResponseBody) SetRequestId(v string) *PutReadyFlagByProxyResponseBody {
	s.RequestId = &v
	return s
}

type PutReadyFlagByProxyResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PutReadyFlagByProxyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PutReadyFlagByProxyResponse) String() string {
	return tea.Prettify(s)
}

func (s PutReadyFlagByProxyResponse) GoString() string {
	return s.String()
}

func (s *PutReadyFlagByProxyResponse) SetHeaders(v map[string]*string) *PutReadyFlagByProxyResponse {
	s.Headers = v
	return s
}

func (s *PutReadyFlagByProxyResponse) SetStatusCode(v int32) *PutReadyFlagByProxyResponse {
	s.StatusCode = &v
	return s
}

func (s *PutReadyFlagByProxyResponse) SetBody(v *PutReadyFlagByProxyResponseBody) *PutReadyFlagByProxyResponse {
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
		"ap-southeast-1": tea.String("oms.ap-southeast-1.aliyuncs.com"),
		"cn-hangzhou":    tea.String("pre-oms.cn-hangzhou.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("oms"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) CheckReadyFlagWithOptions(request *CheckReadyFlagRequest, runtime *util.RuntimeOptions) (_result *CheckReadyFlagResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Cycles)) {
		query["Cycles"] = request.Cycles
	}

	if !tea.BoolValue(util.IsUnset(request.DataType)) {
		query["DataType"] = request.DataType
	}

	if !tea.BoolValue(util.IsUnset(request.DomainCode)) {
		query["DomainCode"] = request.DomainCode
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckReadyFlag"),
		Version:     tea.String("2016-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckReadyFlagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckReadyFlag(request *CheckReadyFlagRequest) (_result *CheckReadyFlagResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckReadyFlagResponse{}
	_body, _err := client.CheckReadyFlagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDomainPartWithOptions(request *DeleteDomainPartRequest, runtime *util.RuntimeOptions) (_result *DeleteDomainPartResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DataType)) {
		query["DataType"] = request.DataType
	}

	if !tea.BoolValue(util.IsUnset(request.DomainCode)) {
		query["DomainCode"] = request.DomainCode
	}

	if !tea.BoolValue(util.IsUnset(request.Part)) {
		query["Part"] = request.Part
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDomainPart"),
		Version:     tea.String("2016-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDomainPartResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDomainPart(request *DeleteDomainPartRequest) (_result *DeleteDomainPartResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDomainPartResponse{}
	_body, _err := client.DeleteDomainPartWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDomainPartByProxyWithOptions(request *DeleteDomainPartByProxyRequest, runtime *util.RuntimeOptions) (_result *DeleteDomainPartByProxyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DataType)) {
		query["DataType"] = request.DataType
	}

	if !tea.BoolValue(util.IsUnset(request.DomainCode)) {
		query["DomainCode"] = request.DomainCode
	}

	if !tea.BoolValue(util.IsUnset(request.Part)) {
		query["Part"] = request.Part
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDomainPartByProxy"),
		Version:     tea.String("2016-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDomainPartByProxyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDomainPartByProxy(request *DeleteDomainPartByProxyRequest) (_result *DeleteDomainPartByProxyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDomainPartByProxyResponse{}
	_body, _err := client.DeleteDomainPartByProxyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteMeasureDataWithOptions(request *DeleteMeasureDataRequest, runtime *util.RuntimeOptions) (_result *DeleteMeasureDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiType)) {
		query["ApiType"] = request.ApiType
	}

	if !tea.BoolValue(util.IsUnset(request.Compressed)) {
		query["Compressed"] = request.Compressed
	}

	if !tea.BoolValue(util.IsUnset(request.Data)) {
		query["Data"] = request.Data
	}

	if !tea.BoolValue(util.IsUnset(request.DataType)) {
		query["DataType"] = request.DataType
	}

	if !tea.BoolValue(util.IsUnset(request.DomainCode)) {
		query["DomainCode"] = request.DomainCode
	}

	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		query["Filter"] = request.Filter
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteMeasureData"),
		Version:     tea.String("2016-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteMeasureDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteMeasureData(request *DeleteMeasureDataRequest) (_result *DeleteMeasureDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteMeasureDataResponse{}
	_body, _err := client.DeleteMeasureDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAccessPolicyConfigWithOptions(request *GetAccessPolicyConfigRequest, runtime *util.RuntimeOptions) (_result *GetAccessPolicyConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAccessPolicyConfig"),
		Version:     tea.String("2016-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAccessPolicyConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAccessPolicyConfig(request *GetAccessPolicyConfigRequest) (_result *GetAccessPolicyConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAccessPolicyConfigResponse{}
	_body, _err := client.GetAccessPolicyConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDomainConfigWithOptions(request *GetDomainConfigRequest, runtime *util.RuntimeOptions) (_result *GetDomainConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDomainConfig"),
		Version:     tea.String("2016-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDomainConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDomainConfig(request *GetDomainConfigRequest) (_result *GetDomainConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDomainConfigResponse{}
	_body, _err := client.GetDomainConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDomainPartWithOptions(request *GetDomainPartRequest, runtime *util.RuntimeOptions) (_result *GetDomainPartResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDomainPart"),
		Version:     tea.String("2016-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDomainPartResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDomainPart(request *GetDomainPartRequest) (_result *GetDomainPartResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDomainPartResponse{}
	_body, _err := client.GetDomainPartWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDomainPartByProxyWithOptions(request *GetDomainPartByProxyRequest, runtime *util.RuntimeOptions) (_result *GetDomainPartByProxyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CompressEnable)) {
		query["CompressEnable"] = request.CompressEnable
	}

	if !tea.BoolValue(util.IsUnset(request.DataType)) {
		query["DataType"] = request.DataType
	}

	if !tea.BoolValue(util.IsUnset(request.DomainCode)) {
		query["DomainCode"] = request.DomainCode
	}

	if !tea.BoolValue(util.IsUnset(request.Part)) {
		query["Part"] = request.Part
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDomainPartByProxy"),
		Version:     tea.String("2016-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDomainPartByProxyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDomainPartByProxy(request *GetDomainPartByProxyRequest) (_result *GetDomainPartByProxyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDomainPartByProxyResponse{}
	_body, _err := client.GetDomainPartByProxyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDomainPartConfigWithOptions(request *GetDomainPartConfigRequest, runtime *util.RuntimeOptions) (_result *GetDomainPartConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDomainPartConfig"),
		Version:     tea.String("2016-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDomainPartConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDomainPartConfig(request *GetDomainPartConfigRequest) (_result *GetDomainPartConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDomainPartConfigResponse{}
	_body, _err := client.GetDomainPartConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetFileConfigWithOptions(request *GetFileConfigRequest, runtime *util.RuntimeOptions) (_result *GetFileConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetFileConfig"),
		Version:     tea.String("2016-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFileConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetFileConfig(request *GetFileConfigRequest) (_result *GetFileConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetFileConfigResponse{}
	_body, _err := client.GetFileConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetIncrementMeasureDataByProxyWithOptions(request *GetIncrementMeasureDataByProxyRequest, runtime *util.RuntimeOptions) (_result *GetIncrementMeasureDataByProxyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetIncrementMeasureDataByProxy"),
		Version:     tea.String("2016-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetIncrementMeasureDataByProxyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetIncrementMeasureDataByProxy(request *GetIncrementMeasureDataByProxyRequest) (_result *GetIncrementMeasureDataByProxyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetIncrementMeasureDataByProxyResponse{}
	_body, _err := client.GetIncrementMeasureDataByProxyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMeasureDataWithOptions(request *GetMeasureDataRequest, runtime *util.RuntimeOptions) (_result *GetMeasureDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiType)) {
		query["ApiType"] = request.ApiType
	}

	if !tea.BoolValue(util.IsUnset(request.CompressEnable)) {
		query["CompressEnable"] = request.CompressEnable
	}

	if !tea.BoolValue(util.IsUnset(request.DataType)) {
		query["DataType"] = request.DataType
	}

	if !tea.BoolValue(util.IsUnset(request.DomainCode)) {
		query["DomainCode"] = request.DomainCode
	}

	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		query["Filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResult)) {
		query["MaxResult"] = request.MaxResult
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.QueryField)) {
		query["QueryField"] = request.QueryField
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetMeasureData"),
		Version:     tea.String("2016-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMeasureDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMeasureData(request *GetMeasureDataRequest) (_result *GetMeasureDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMeasureDataResponse{}
	_body, _err := client.GetMeasureDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOpenApiConfigWithOptions(request *GetOpenApiConfigRequest, runtime *util.RuntimeOptions) (_result *GetOpenApiConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOpenApiConfig"),
		Version:     tea.String("2016-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOpenApiConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOpenApiConfig(request *GetOpenApiConfigRequest) (_result *GetOpenApiConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOpenApiConfigResponse{}
	_body, _err := client.GetOpenApiConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetReadyFlagWithOptions(request *GetReadyFlagRequest, runtime *util.RuntimeOptions) (_result *GetReadyFlagResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetReadyFlag"),
		Version:     tea.String("2016-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetReadyFlagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetReadyFlag(request *GetReadyFlagRequest) (_result *GetReadyFlagResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetReadyFlagResponse{}
	_body, _err := client.GetReadyFlagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetReadyFlagAlertConfigWithOptions(request *GetReadyFlagAlertConfigRequest, runtime *util.RuntimeOptions) (_result *GetReadyFlagAlertConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetReadyFlagAlertConfig"),
		Version:     tea.String("2016-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetReadyFlagAlertConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetReadyFlagAlertConfig(request *GetReadyFlagAlertConfigRequest) (_result *GetReadyFlagAlertConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetReadyFlagAlertConfigResponse{}
	_body, _err := client.GetReadyFlagAlertConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetReadyFlagByProxyWithOptions(request *GetReadyFlagByProxyRequest, runtime *util.RuntimeOptions) (_result *GetReadyFlagByProxyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiType)) {
		query["ApiType"] = request.ApiType
	}

	if !tea.BoolValue(util.IsUnset(request.CompressEnable)) {
		query["CompressEnable"] = request.CompressEnable
	}

	if !tea.BoolValue(util.IsUnset(request.DataType)) {
		query["DataType"] = request.DataType
	}

	if !tea.BoolValue(util.IsUnset(request.DomainCode)) {
		query["DomainCode"] = request.DomainCode
	}

	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		query["Filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResult)) {
		query["MaxResult"] = request.MaxResult
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetReadyFlagByProxy"),
		Version:     tea.String("2016-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetReadyFlagByProxyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetReadyFlagByProxy(request *GetReadyFlagByProxyRequest) (_result *GetReadyFlagByProxyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetReadyFlagByProxyResponse{}
	_body, _err := client.GetReadyFlagByProxyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PutDomainPartWithOptions(request *PutDomainPartRequest, runtime *util.RuntimeOptions) (_result *PutDomainPartResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Compressed)) {
		query["Compressed"] = request.Compressed
	}

	if !tea.BoolValue(util.IsUnset(request.Data)) {
		query["Data"] = request.Data
	}

	if !tea.BoolValue(util.IsUnset(request.DataType)) {
		query["DataType"] = request.DataType
	}

	if !tea.BoolValue(util.IsUnset(request.DomainCode)) {
		query["DomainCode"] = request.DomainCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PutDomainPart"),
		Version:     tea.String("2016-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PutDomainPartResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PutDomainPart(request *PutDomainPartRequest) (_result *PutDomainPartResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PutDomainPartResponse{}
	_body, _err := client.PutDomainPartWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PutDomainPartByProxyWithOptions(request *PutDomainPartByProxyRequest, runtime *util.RuntimeOptions) (_result *PutDomainPartByProxyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Compressed)) {
		query["Compressed"] = request.Compressed
	}

	if !tea.BoolValue(util.IsUnset(request.Data)) {
		query["Data"] = request.Data
	}

	if !tea.BoolValue(util.IsUnset(request.DataType)) {
		query["DataType"] = request.DataType
	}

	if !tea.BoolValue(util.IsUnset(request.DomainCode)) {
		query["DomainCode"] = request.DomainCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PutDomainPartByProxy"),
		Version:     tea.String("2016-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PutDomainPartByProxyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PutDomainPartByProxy(request *PutDomainPartByProxyRequest) (_result *PutDomainPartByProxyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PutDomainPartByProxyResponse{}
	_body, _err := client.PutDomainPartByProxyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PutMeasureDataWithOptions(request *PutMeasureDataRequest, runtime *util.RuntimeOptions) (_result *PutMeasureDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiType)) {
		query["ApiType"] = request.ApiType
	}

	if !tea.BoolValue(util.IsUnset(request.Compressed)) {
		query["Compressed"] = request.Compressed
	}

	if !tea.BoolValue(util.IsUnset(request.Data)) {
		query["Data"] = request.Data
	}

	if !tea.BoolValue(util.IsUnset(request.DataType)) {
		query["DataType"] = request.DataType
	}

	if !tea.BoolValue(util.IsUnset(request.DomainCode)) {
		query["DomainCode"] = request.DomainCode
	}

	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		query["Filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.SourceRequestId)) {
		query["SourceRequestId"] = request.SourceRequestId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PutMeasureData"),
		Version:     tea.String("2016-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PutMeasureDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PutMeasureData(request *PutMeasureDataRequest) (_result *PutMeasureDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PutMeasureDataResponse{}
	_body, _err := client.PutMeasureDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PutMeasureDataByProxyWithOptions(request *PutMeasureDataByProxyRequest, runtime *util.RuntimeOptions) (_result *PutMeasureDataByProxyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiType)) {
		query["ApiType"] = request.ApiType
	}

	if !tea.BoolValue(util.IsUnset(request.Compressed)) {
		query["Compressed"] = request.Compressed
	}

	if !tea.BoolValue(util.IsUnset(request.Data)) {
		query["Data"] = request.Data
	}

	if !tea.BoolValue(util.IsUnset(request.DataType)) {
		query["DataType"] = request.DataType
	}

	if !tea.BoolValue(util.IsUnset(request.DomainCode)) {
		query["DomainCode"] = request.DomainCode
	}

	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		query["Filter"] = request.Filter
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PutMeasureDataByProxy"),
		Version:     tea.String("2016-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PutMeasureDataByProxyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PutMeasureDataByProxy(request *PutMeasureDataByProxyRequest) (_result *PutMeasureDataByProxyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PutMeasureDataByProxyResponse{}
	_body, _err := client.PutMeasureDataByProxyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PutReadyFlagWithOptions(request *PutReadyFlagRequest, runtime *util.RuntimeOptions) (_result *PutReadyFlagResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiType)) {
		query["ApiType"] = request.ApiType
	}

	if !tea.BoolValue(util.IsUnset(request.Compressed)) {
		query["Compressed"] = request.Compressed
	}

	if !tea.BoolValue(util.IsUnset(request.Data)) {
		query["Data"] = request.Data
	}

	if !tea.BoolValue(util.IsUnset(request.DataType)) {
		query["DataType"] = request.DataType
	}

	if !tea.BoolValue(util.IsUnset(request.DomainCode)) {
		query["DomainCode"] = request.DomainCode
	}

	if !tea.BoolValue(util.IsUnset(request.SourceRequestId)) {
		query["SourceRequestId"] = request.SourceRequestId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PutReadyFlag"),
		Version:     tea.String("2016-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PutReadyFlagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PutReadyFlag(request *PutReadyFlagRequest) (_result *PutReadyFlagResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PutReadyFlagResponse{}
	_body, _err := client.PutReadyFlagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PutReadyFlagByProxyWithOptions(request *PutReadyFlagByProxyRequest, runtime *util.RuntimeOptions) (_result *PutReadyFlagByProxyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiType)) {
		query["ApiType"] = request.ApiType
	}

	if !tea.BoolValue(util.IsUnset(request.Compressed)) {
		query["Compressed"] = request.Compressed
	}

	if !tea.BoolValue(util.IsUnset(request.Data)) {
		query["Data"] = request.Data
	}

	if !tea.BoolValue(util.IsUnset(request.DataType)) {
		query["DataType"] = request.DataType
	}

	if !tea.BoolValue(util.IsUnset(request.DomainCode)) {
		query["DomainCode"] = request.DomainCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PutReadyFlagByProxy"),
		Version:     tea.String("2016-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PutReadyFlagByProxyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PutReadyFlagByProxy(request *PutReadyFlagByProxyRequest) (_result *PutReadyFlagByProxyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PutReadyFlagByProxyResponse{}
	_body, _err := client.PutReadyFlagByProxyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
