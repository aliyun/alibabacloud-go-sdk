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

type DataValue struct {
	ServiceId *int64  `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s DataValue) String() string {
	return tea.Prettify(s)
}

func (s DataValue) GoString() string {
	return s.String()
}

func (s *DataValue) SetServiceId(v int64) *DataValue {
	s.ServiceId = &v
	return s
}

func (s *DataValue) SetStatus(v string) *DataValue {
	s.Status = &v
	return s
}

func (s *DataValue) SetCode(v int32) *DataValue {
	s.Code = &v
	return s
}

func (s *DataValue) SetMessage(v string) *DataValue {
	s.Message = &v
	return s
}

type ADClockRequest struct {
	Params      *string `json:"Params,omitempty" xml:"Params,omitempty"`
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
}

func (s ADClockRequest) String() string {
	return tea.Prettify(s)
}

func (s ADClockRequest) GoString() string {
	return s.String()
}

func (s *ADClockRequest) SetParams(v string) *ADClockRequest {
	s.Params = &v
	return s
}

func (s *ADClockRequest) SetServiceCode(v string) *ADClockRequest {
	s.ServiceCode = &v
	return s
}

type ADClockResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ADClockResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ADClockResponseBody) GoString() string {
	return s.String()
}

func (s *ADClockResponseBody) SetData(v string) *ADClockResponseBody {
	s.Data = &v
	return s
}

func (s *ADClockResponseBody) SetRequestId(v string) *ADClockResponseBody {
	s.RequestId = &v
	return s
}

type ADClockResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ADClockResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ADClockResponse) String() string {
	return tea.Prettify(s)
}

func (s ADClockResponse) GoString() string {
	return s.String()
}

func (s *ADClockResponse) SetHeaders(v map[string]*string) *ADClockResponse {
	s.Headers = v
	return s
}

func (s *ADClockResponse) SetStatusCode(v int32) *ADClockResponse {
	s.StatusCode = &v
	return s
}

func (s *ADClockResponse) SetBody(v *ADClockResponseBody) *ADClockResponse {
	s.Body = v
	return s
}

type ADMMURequest struct {
	Params      *string `json:"Params,omitempty" xml:"Params,omitempty"`
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
}

func (s ADMMURequest) String() string {
	return tea.Prettify(s)
}

func (s ADMMURequest) GoString() string {
	return s.String()
}

func (s *ADMMURequest) SetParams(v string) *ADMMURequest {
	s.Params = &v
	return s
}

func (s *ADMMURequest) SetServiceCode(v string) *ADMMURequest {
	s.ServiceCode = &v
	return s
}

type ADMMUResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ADMMUResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ADMMUResponseBody) GoString() string {
	return s.String()
}

func (s *ADMMUResponseBody) SetData(v string) *ADMMUResponseBody {
	s.Data = &v
	return s
}

func (s *ADMMUResponseBody) SetRequestId(v string) *ADMMUResponseBody {
	s.RequestId = &v
	return s
}

type ADMMUResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ADMMUResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ADMMUResponse) String() string {
	return tea.Prettify(s)
}

func (s ADMMUResponse) GoString() string {
	return s.String()
}

func (s *ADMMUResponse) SetHeaders(v map[string]*string) *ADMMUResponse {
	s.Headers = v
	return s
}

func (s *ADMMUResponse) SetStatusCode(v int32) *ADMMUResponse {
	s.StatusCode = &v
	return s
}

func (s *ADMMUResponse) SetBody(v *ADMMUResponseBody) *ADMMUResponse {
	s.Body = v
	return s
}

type ADMiniCogRequest struct {
	Params      *string `json:"Params,omitempty" xml:"Params,omitempty"`
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
}

func (s ADMiniCogRequest) String() string {
	return tea.Prettify(s)
}

func (s ADMiniCogRequest) GoString() string {
	return s.String()
}

func (s *ADMiniCogRequest) SetParams(v string) *ADMiniCogRequest {
	s.Params = &v
	return s
}

func (s *ADMiniCogRequest) SetServiceCode(v string) *ADMiniCogRequest {
	s.ServiceCode = &v
	return s
}

type ADMiniCogResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ADMiniCogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ADMiniCogResponseBody) GoString() string {
	return s.String()
}

func (s *ADMiniCogResponseBody) SetData(v string) *ADMiniCogResponseBody {
	s.Data = &v
	return s
}

func (s *ADMiniCogResponseBody) SetRequestId(v string) *ADMiniCogResponseBody {
	s.RequestId = &v
	return s
}

type ADMiniCogResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ADMiniCogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ADMiniCogResponse) String() string {
	return tea.Prettify(s)
}

func (s ADMiniCogResponse) GoString() string {
	return s.String()
}

func (s *ADMiniCogResponse) SetHeaders(v map[string]*string) *ADMiniCogResponse {
	s.Headers = v
	return s
}

func (s *ADMiniCogResponse) SetStatusCode(v int32) *ADMiniCogResponse {
	s.StatusCode = &v
	return s
}

func (s *ADMiniCogResponse) SetBody(v *ADMiniCogResponseBody) *ADMiniCogResponse {
	s.Body = v
	return s
}

type ADMiniCogResultRequest struct {
	Params      *string `json:"Params,omitempty" xml:"Params,omitempty"`
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
}

func (s ADMiniCogResultRequest) String() string {
	return tea.Prettify(s)
}

func (s ADMiniCogResultRequest) GoString() string {
	return s.String()
}

func (s *ADMiniCogResultRequest) SetParams(v string) *ADMiniCogResultRequest {
	s.Params = &v
	return s
}

func (s *ADMiniCogResultRequest) SetServiceCode(v string) *ADMiniCogResultRequest {
	s.ServiceCode = &v
	return s
}

type ADMiniCogResultResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ADMiniCogResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ADMiniCogResultResponseBody) GoString() string {
	return s.String()
}

func (s *ADMiniCogResultResponseBody) SetData(v string) *ADMiniCogResultResponseBody {
	s.Data = &v
	return s
}

func (s *ADMiniCogResultResponseBody) SetRequestId(v string) *ADMiniCogResultResponseBody {
	s.RequestId = &v
	return s
}

type ADMiniCogResultResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ADMiniCogResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ADMiniCogResultResponse) String() string {
	return tea.Prettify(s)
}

func (s ADMiniCogResultResponse) GoString() string {
	return s.String()
}

func (s *ADMiniCogResultResponse) SetHeaders(v map[string]*string) *ADMiniCogResultResponse {
	s.Headers = v
	return s
}

func (s *ADMiniCogResultResponse) SetStatusCode(v int32) *ADMiniCogResultResponse {
	s.StatusCode = &v
	return s
}

func (s *ADMiniCogResultResponse) SetBody(v *ADMiniCogResultResponseBody) *ADMiniCogResultResponse {
	s.Body = v
	return s
}

type DeleteServiceDataByConditionsRequest struct {
	Conditions map[string]interface{} `json:"Conditions,omitempty" xml:"Conditions,omitempty"`
	ServiceId  *int64                 `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s DeleteServiceDataByConditionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceDataByConditionsRequest) GoString() string {
	return s.String()
}

func (s *DeleteServiceDataByConditionsRequest) SetConditions(v map[string]interface{}) *DeleteServiceDataByConditionsRequest {
	s.Conditions = v
	return s
}

func (s *DeleteServiceDataByConditionsRequest) SetServiceId(v int64) *DeleteServiceDataByConditionsRequest {
	s.ServiceId = &v
	return s
}

type DeleteServiceDataByConditionsShrinkRequest struct {
	ConditionsShrink *string `json:"Conditions,omitempty" xml:"Conditions,omitempty"`
	ServiceId        *int64  `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s DeleteServiceDataByConditionsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceDataByConditionsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteServiceDataByConditionsShrinkRequest) SetConditionsShrink(v string) *DeleteServiceDataByConditionsShrinkRequest {
	s.ConditionsShrink = &v
	return s
}

func (s *DeleteServiceDataByConditionsShrinkRequest) SetServiceId(v int64) *DeleteServiceDataByConditionsShrinkRequest {
	s.ServiceId = &v
	return s
}

type DeleteServiceDataByConditionsResponseBody struct {
	Code      *int32      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Msg       *string     `json:"Msg,omitempty" xml:"Msg,omitempty"`
	RequestId *string     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteServiceDataByConditionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceDataByConditionsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteServiceDataByConditionsResponseBody) SetCode(v int32) *DeleteServiceDataByConditionsResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteServiceDataByConditionsResponseBody) SetData(v interface{}) *DeleteServiceDataByConditionsResponseBody {
	s.Data = v
	return s
}

func (s *DeleteServiceDataByConditionsResponseBody) SetMsg(v string) *DeleteServiceDataByConditionsResponseBody {
	s.Msg = &v
	return s
}

func (s *DeleteServiceDataByConditionsResponseBody) SetRequestId(v string) *DeleteServiceDataByConditionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteServiceDataByConditionsResponseBody) SetSuccess(v bool) *DeleteServiceDataByConditionsResponseBody {
	s.Success = &v
	return s
}

type DeleteServiceDataByConditionsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteServiceDataByConditionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteServiceDataByConditionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceDataByConditionsResponse) GoString() string {
	return s.String()
}

func (s *DeleteServiceDataByConditionsResponse) SetHeaders(v map[string]*string) *DeleteServiceDataByConditionsResponse {
	s.Headers = v
	return s
}

func (s *DeleteServiceDataByConditionsResponse) SetStatusCode(v int32) *DeleteServiceDataByConditionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteServiceDataByConditionsResponse) SetBody(v *DeleteServiceDataByConditionsResponseBody) *DeleteServiceDataByConditionsResponse {
	s.Body = v
	return s
}

type DeleteServiceDataByIdsRequest struct {
	Ids       []*string `json:"Ids,omitempty" xml:"Ids,omitempty" type:"Repeated"`
	ServiceId *int64    `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s DeleteServiceDataByIdsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceDataByIdsRequest) GoString() string {
	return s.String()
}

func (s *DeleteServiceDataByIdsRequest) SetIds(v []*string) *DeleteServiceDataByIdsRequest {
	s.Ids = v
	return s
}

func (s *DeleteServiceDataByIdsRequest) SetServiceId(v int64) *DeleteServiceDataByIdsRequest {
	s.ServiceId = &v
	return s
}

type DeleteServiceDataByIdsShrinkRequest struct {
	IdsShrink *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	ServiceId *int64  `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s DeleteServiceDataByIdsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceDataByIdsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteServiceDataByIdsShrinkRequest) SetIdsShrink(v string) *DeleteServiceDataByIdsShrinkRequest {
	s.IdsShrink = &v
	return s
}

func (s *DeleteServiceDataByIdsShrinkRequest) SetServiceId(v int64) *DeleteServiceDataByIdsShrinkRequest {
	s.ServiceId = &v
	return s
}

type DeleteServiceDataByIdsResponseBody struct {
	Code      *int32      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Msg       *string     `json:"Msg,omitempty" xml:"Msg,omitempty"`
	RequestId *string     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteServiceDataByIdsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceDataByIdsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteServiceDataByIdsResponseBody) SetCode(v int32) *DeleteServiceDataByIdsResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteServiceDataByIdsResponseBody) SetData(v interface{}) *DeleteServiceDataByIdsResponseBody {
	s.Data = v
	return s
}

func (s *DeleteServiceDataByIdsResponseBody) SetMsg(v string) *DeleteServiceDataByIdsResponseBody {
	s.Msg = &v
	return s
}

func (s *DeleteServiceDataByIdsResponseBody) SetRequestId(v string) *DeleteServiceDataByIdsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteServiceDataByIdsResponseBody) SetSuccess(v bool) *DeleteServiceDataByIdsResponseBody {
	s.Success = &v
	return s
}

type DeleteServiceDataByIdsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteServiceDataByIdsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteServiceDataByIdsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceDataByIdsResponse) GoString() string {
	return s.String()
}

func (s *DeleteServiceDataByIdsResponse) SetHeaders(v map[string]*string) *DeleteServiceDataByIdsResponse {
	s.Headers = v
	return s
}

func (s *DeleteServiceDataByIdsResponse) SetStatusCode(v int32) *DeleteServiceDataByIdsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteServiceDataByIdsResponse) SetBody(v *DeleteServiceDataByIdsResponseBody) *DeleteServiceDataByIdsResponse {
	s.Body = v
	return s
}

type GetBrandChEcomRequest struct {
	ImageUrl    *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	Text        *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s GetBrandChEcomRequest) String() string {
	return tea.Prettify(s)
}

func (s GetBrandChEcomRequest) GoString() string {
	return s.String()
}

func (s *GetBrandChEcomRequest) SetImageUrl(v string) *GetBrandChEcomRequest {
	s.ImageUrl = &v
	return s
}

func (s *GetBrandChEcomRequest) SetServiceCode(v string) *GetBrandChEcomRequest {
	s.ServiceCode = &v
	return s
}

func (s *GetBrandChEcomRequest) SetText(v string) *GetBrandChEcomRequest {
	s.Text = &v
	return s
}

type GetBrandChEcomResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetBrandChEcomResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetBrandChEcomResponseBody) GoString() string {
	return s.String()
}

func (s *GetBrandChEcomResponseBody) SetData(v string) *GetBrandChEcomResponseBody {
	s.Data = &v
	return s
}

func (s *GetBrandChEcomResponseBody) SetRequestId(v string) *GetBrandChEcomResponseBody {
	s.RequestId = &v
	return s
}

type GetBrandChEcomResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBrandChEcomResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBrandChEcomResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBrandChEcomResponse) GoString() string {
	return s.String()
}

func (s *GetBrandChEcomResponse) SetHeaders(v map[string]*string) *GetBrandChEcomResponse {
	s.Headers = v
	return s
}

func (s *GetBrandChEcomResponse) SetStatusCode(v int32) *GetBrandChEcomResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBrandChEcomResponse) SetBody(v *GetBrandChEcomResponseBody) *GetBrandChEcomResponse {
	s.Body = v
	return s
}

type GetCateChEcomRequest struct {
	ImageUrl    *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	Text        *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s GetCateChEcomRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCateChEcomRequest) GoString() string {
	return s.String()
}

func (s *GetCateChEcomRequest) SetImageUrl(v string) *GetCateChEcomRequest {
	s.ImageUrl = &v
	return s
}

func (s *GetCateChEcomRequest) SetServiceCode(v string) *GetCateChEcomRequest {
	s.ServiceCode = &v
	return s
}

func (s *GetCateChEcomRequest) SetText(v string) *GetCateChEcomRequest {
	s.Text = &v
	return s
}

type GetCateChEcomResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetCateChEcomResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCateChEcomResponseBody) GoString() string {
	return s.String()
}

func (s *GetCateChEcomResponseBody) SetData(v string) *GetCateChEcomResponseBody {
	s.Data = &v
	return s
}

func (s *GetCateChEcomResponseBody) SetRequestId(v string) *GetCateChEcomResponseBody {
	s.RequestId = &v
	return s
}

type GetCateChEcomResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCateChEcomResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCateChEcomResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCateChEcomResponse) GoString() string {
	return s.String()
}

func (s *GetCateChEcomResponse) SetHeaders(v map[string]*string) *GetCateChEcomResponse {
	s.Headers = v
	return s
}

func (s *GetCateChEcomResponse) SetStatusCode(v int32) *GetCateChEcomResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCateChEcomResponse) SetBody(v *GetCateChEcomResponseBody) *GetCateChEcomResponse {
	s.Body = v
	return s
}

type GetCheckDuplicationChMedicalRequest struct {
	OriginQ     *string `json:"OriginQ,omitempty" xml:"OriginQ,omitempty"`
	OriginT     *string `json:"OriginT,omitempty" xml:"OriginT,omitempty"`
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
}

func (s GetCheckDuplicationChMedicalRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCheckDuplicationChMedicalRequest) GoString() string {
	return s.String()
}

func (s *GetCheckDuplicationChMedicalRequest) SetOriginQ(v string) *GetCheckDuplicationChMedicalRequest {
	s.OriginQ = &v
	return s
}

func (s *GetCheckDuplicationChMedicalRequest) SetOriginT(v string) *GetCheckDuplicationChMedicalRequest {
	s.OriginT = &v
	return s
}

func (s *GetCheckDuplicationChMedicalRequest) SetServiceCode(v string) *GetCheckDuplicationChMedicalRequest {
	s.ServiceCode = &v
	return s
}

type GetCheckDuplicationChMedicalResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetCheckDuplicationChMedicalResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCheckDuplicationChMedicalResponseBody) GoString() string {
	return s.String()
}

func (s *GetCheckDuplicationChMedicalResponseBody) SetData(v string) *GetCheckDuplicationChMedicalResponseBody {
	s.Data = &v
	return s
}

func (s *GetCheckDuplicationChMedicalResponseBody) SetRequestId(v string) *GetCheckDuplicationChMedicalResponseBody {
	s.RequestId = &v
	return s
}

type GetCheckDuplicationChMedicalResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCheckDuplicationChMedicalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCheckDuplicationChMedicalResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCheckDuplicationChMedicalResponse) GoString() string {
	return s.String()
}

func (s *GetCheckDuplicationChMedicalResponse) SetHeaders(v map[string]*string) *GetCheckDuplicationChMedicalResponse {
	s.Headers = v
	return s
}

func (s *GetCheckDuplicationChMedicalResponse) SetStatusCode(v int32) *GetCheckDuplicationChMedicalResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCheckDuplicationChMedicalResponse) SetBody(v *GetCheckDuplicationChMedicalResponseBody) *GetCheckDuplicationChMedicalResponse {
	s.Body = v
	return s
}

type GetDiagnosisChMedicalRequest struct {
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
}

func (s GetDiagnosisChMedicalRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDiagnosisChMedicalRequest) GoString() string {
	return s.String()
}

func (s *GetDiagnosisChMedicalRequest) SetName(v string) *GetDiagnosisChMedicalRequest {
	s.Name = &v
	return s
}

func (s *GetDiagnosisChMedicalRequest) SetServiceCode(v string) *GetDiagnosisChMedicalRequest {
	s.ServiceCode = &v
	return s
}

type GetDiagnosisChMedicalResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDiagnosisChMedicalResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDiagnosisChMedicalResponseBody) GoString() string {
	return s.String()
}

func (s *GetDiagnosisChMedicalResponseBody) SetData(v string) *GetDiagnosisChMedicalResponseBody {
	s.Data = &v
	return s
}

func (s *GetDiagnosisChMedicalResponseBody) SetRequestId(v string) *GetDiagnosisChMedicalResponseBody {
	s.RequestId = &v
	return s
}

type GetDiagnosisChMedicalResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDiagnosisChMedicalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDiagnosisChMedicalResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDiagnosisChMedicalResponse) GoString() string {
	return s.String()
}

func (s *GetDiagnosisChMedicalResponse) SetHeaders(v map[string]*string) *GetDiagnosisChMedicalResponse {
	s.Headers = v
	return s
}

func (s *GetDiagnosisChMedicalResponse) SetStatusCode(v int32) *GetDiagnosisChMedicalResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDiagnosisChMedicalResponse) SetBody(v *GetDiagnosisChMedicalResponseBody) *GetDiagnosisChMedicalResponse {
	s.Body = v
	return s
}

type GetDpChEcomRequest struct {
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	Text        *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s GetDpChEcomRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDpChEcomRequest) GoString() string {
	return s.String()
}

func (s *GetDpChEcomRequest) SetServiceCode(v string) *GetDpChEcomRequest {
	s.ServiceCode = &v
	return s
}

func (s *GetDpChEcomRequest) SetText(v string) *GetDpChEcomRequest {
	s.Text = &v
	return s
}

type GetDpChEcomResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDpChEcomResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDpChEcomResponseBody) GoString() string {
	return s.String()
}

func (s *GetDpChEcomResponseBody) SetData(v string) *GetDpChEcomResponseBody {
	s.Data = &v
	return s
}

func (s *GetDpChEcomResponseBody) SetRequestId(v string) *GetDpChEcomResponseBody {
	s.RequestId = &v
	return s
}

type GetDpChEcomResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDpChEcomResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDpChEcomResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDpChEcomResponse) GoString() string {
	return s.String()
}

func (s *GetDpChEcomResponse) SetHeaders(v map[string]*string) *GetDpChEcomResponse {
	s.Headers = v
	return s
}

func (s *GetDpChEcomResponse) SetStatusCode(v int32) *GetDpChEcomResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDpChEcomResponse) SetBody(v *GetDpChEcomResponseBody) *GetDpChEcomResponse {
	s.Body = v
	return s
}

type GetDpChGeneralCTBRequest struct {
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	Text        *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s GetDpChGeneralCTBRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDpChGeneralCTBRequest) GoString() string {
	return s.String()
}

func (s *GetDpChGeneralCTBRequest) SetServiceCode(v string) *GetDpChGeneralCTBRequest {
	s.ServiceCode = &v
	return s
}

func (s *GetDpChGeneralCTBRequest) SetText(v string) *GetDpChGeneralCTBRequest {
	s.Text = &v
	return s
}

type GetDpChGeneralCTBResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDpChGeneralCTBResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDpChGeneralCTBResponseBody) GoString() string {
	return s.String()
}

func (s *GetDpChGeneralCTBResponseBody) SetData(v string) *GetDpChGeneralCTBResponseBody {
	s.Data = &v
	return s
}

func (s *GetDpChGeneralCTBResponseBody) SetRequestId(v string) *GetDpChGeneralCTBResponseBody {
	s.RequestId = &v
	return s
}

type GetDpChGeneralCTBResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDpChGeneralCTBResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDpChGeneralCTBResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDpChGeneralCTBResponse) GoString() string {
	return s.String()
}

func (s *GetDpChGeneralCTBResponse) SetHeaders(v map[string]*string) *GetDpChGeneralCTBResponse {
	s.Headers = v
	return s
}

func (s *GetDpChGeneralCTBResponse) SetStatusCode(v int32) *GetDpChGeneralCTBResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDpChGeneralCTBResponse) SetBody(v *GetDpChGeneralCTBResponseBody) *GetDpChGeneralCTBResponse {
	s.Body = v
	return s
}

type GetDpChGeneralStanfordRequest struct {
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	Text        *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s GetDpChGeneralStanfordRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDpChGeneralStanfordRequest) GoString() string {
	return s.String()
}

func (s *GetDpChGeneralStanfordRequest) SetServiceCode(v string) *GetDpChGeneralStanfordRequest {
	s.ServiceCode = &v
	return s
}

func (s *GetDpChGeneralStanfordRequest) SetText(v string) *GetDpChGeneralStanfordRequest {
	s.Text = &v
	return s
}

type GetDpChGeneralStanfordResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDpChGeneralStanfordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDpChGeneralStanfordResponseBody) GoString() string {
	return s.String()
}

func (s *GetDpChGeneralStanfordResponseBody) SetData(v string) *GetDpChGeneralStanfordResponseBody {
	s.Data = &v
	return s
}

func (s *GetDpChGeneralStanfordResponseBody) SetRequestId(v string) *GetDpChGeneralStanfordResponseBody {
	s.RequestId = &v
	return s
}

type GetDpChGeneralStanfordResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDpChGeneralStanfordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDpChGeneralStanfordResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDpChGeneralStanfordResponse) GoString() string {
	return s.String()
}

func (s *GetDpChGeneralStanfordResponse) SetHeaders(v map[string]*string) *GetDpChGeneralStanfordResponse {
	s.Headers = v
	return s
}

func (s *GetDpChGeneralStanfordResponse) SetStatusCode(v int32) *GetDpChGeneralStanfordResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDpChGeneralStanfordResponse) SetBody(v *GetDpChGeneralStanfordResponseBody) *GetDpChGeneralStanfordResponse {
	s.Body = v
	return s
}

type GetEcChGeneralRequest struct {
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	Text        *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s GetEcChGeneralRequest) String() string {
	return tea.Prettify(s)
}

func (s GetEcChGeneralRequest) GoString() string {
	return s.String()
}

func (s *GetEcChGeneralRequest) SetServiceCode(v string) *GetEcChGeneralRequest {
	s.ServiceCode = &v
	return s
}

func (s *GetEcChGeneralRequest) SetText(v string) *GetEcChGeneralRequest {
	s.Text = &v
	return s
}

type GetEcChGeneralResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetEcChGeneralResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetEcChGeneralResponseBody) GoString() string {
	return s.String()
}

func (s *GetEcChGeneralResponseBody) SetData(v string) *GetEcChGeneralResponseBody {
	s.Data = &v
	return s
}

func (s *GetEcChGeneralResponseBody) SetRequestId(v string) *GetEcChGeneralResponseBody {
	s.RequestId = &v
	return s
}

type GetEcChGeneralResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEcChGeneralResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEcChGeneralResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEcChGeneralResponse) GoString() string {
	return s.String()
}

func (s *GetEcChGeneralResponse) SetHeaders(v map[string]*string) *GetEcChGeneralResponse {
	s.Headers = v
	return s
}

func (s *GetEcChGeneralResponse) SetStatusCode(v int32) *GetEcChGeneralResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEcChGeneralResponse) SetBody(v *GetEcChGeneralResponseBody) *GetEcChGeneralResponse {
	s.Body = v
	return s
}

type GetEcEnGeneralRequest struct {
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	Text        *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s GetEcEnGeneralRequest) String() string {
	return tea.Prettify(s)
}

func (s GetEcEnGeneralRequest) GoString() string {
	return s.String()
}

func (s *GetEcEnGeneralRequest) SetServiceCode(v string) *GetEcEnGeneralRequest {
	s.ServiceCode = &v
	return s
}

func (s *GetEcEnGeneralRequest) SetText(v string) *GetEcEnGeneralRequest {
	s.Text = &v
	return s
}

type GetEcEnGeneralResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetEcEnGeneralResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetEcEnGeneralResponseBody) GoString() string {
	return s.String()
}

func (s *GetEcEnGeneralResponseBody) SetData(v string) *GetEcEnGeneralResponseBody {
	s.Data = &v
	return s
}

func (s *GetEcEnGeneralResponseBody) SetRequestId(v string) *GetEcEnGeneralResponseBody {
	s.RequestId = &v
	return s
}

type GetEcEnGeneralResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEcEnGeneralResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEcEnGeneralResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEcEnGeneralResponse) GoString() string {
	return s.String()
}

func (s *GetEcEnGeneralResponse) SetHeaders(v map[string]*string) *GetEcEnGeneralResponse {
	s.Headers = v
	return s
}

func (s *GetEcEnGeneralResponse) SetStatusCode(v int32) *GetEcEnGeneralResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEcEnGeneralResponse) SetBody(v *GetEcEnGeneralResponseBody) *GetEcEnGeneralResponse {
	s.Body = v
	return s
}

type GetEmbeddingRequest struct {
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	Text        *string `json:"Text,omitempty" xml:"Text,omitempty"`
	TextType    *string `json:"TextType,omitempty" xml:"TextType,omitempty"`
}

func (s GetEmbeddingRequest) String() string {
	return tea.Prettify(s)
}

func (s GetEmbeddingRequest) GoString() string {
	return s.String()
}

func (s *GetEmbeddingRequest) SetServiceCode(v string) *GetEmbeddingRequest {
	s.ServiceCode = &v
	return s
}

func (s *GetEmbeddingRequest) SetText(v string) *GetEmbeddingRequest {
	s.Text = &v
	return s
}

func (s *GetEmbeddingRequest) SetTextType(v string) *GetEmbeddingRequest {
	s.TextType = &v
	return s
}

type GetEmbeddingResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetEmbeddingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetEmbeddingResponseBody) GoString() string {
	return s.String()
}

func (s *GetEmbeddingResponseBody) SetData(v string) *GetEmbeddingResponseBody {
	s.Data = &v
	return s
}

func (s *GetEmbeddingResponseBody) SetRequestId(v string) *GetEmbeddingResponseBody {
	s.RequestId = &v
	return s
}

type GetEmbeddingResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEmbeddingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEmbeddingResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEmbeddingResponse) GoString() string {
	return s.String()
}

func (s *GetEmbeddingResponse) SetHeaders(v map[string]*string) *GetEmbeddingResponse {
	s.Headers = v
	return s
}

func (s *GetEmbeddingResponse) SetStatusCode(v int32) *GetEmbeddingResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEmbeddingResponse) SetBody(v *GetEmbeddingResponseBody) *GetEmbeddingResponse {
	s.Body = v
	return s
}

type GetItemPubChEcomRequest struct {
	ImageUrl    *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	Text        *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s GetItemPubChEcomRequest) String() string {
	return tea.Prettify(s)
}

func (s GetItemPubChEcomRequest) GoString() string {
	return s.String()
}

func (s *GetItemPubChEcomRequest) SetImageUrl(v string) *GetItemPubChEcomRequest {
	s.ImageUrl = &v
	return s
}

func (s *GetItemPubChEcomRequest) SetServiceCode(v string) *GetItemPubChEcomRequest {
	s.ServiceCode = &v
	return s
}

func (s *GetItemPubChEcomRequest) SetText(v string) *GetItemPubChEcomRequest {
	s.Text = &v
	return s
}

type GetItemPubChEcomResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetItemPubChEcomResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetItemPubChEcomResponseBody) GoString() string {
	return s.String()
}

func (s *GetItemPubChEcomResponseBody) SetData(v string) *GetItemPubChEcomResponseBody {
	s.Data = &v
	return s
}

func (s *GetItemPubChEcomResponseBody) SetRequestId(v string) *GetItemPubChEcomResponseBody {
	s.RequestId = &v
	return s
}

type GetItemPubChEcomResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetItemPubChEcomResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetItemPubChEcomResponse) String() string {
	return tea.Prettify(s)
}

func (s GetItemPubChEcomResponse) GoString() string {
	return s.String()
}

func (s *GetItemPubChEcomResponse) SetHeaders(v map[string]*string) *GetItemPubChEcomResponse {
	s.Headers = v
	return s
}

func (s *GetItemPubChEcomResponse) SetStatusCode(v int32) *GetItemPubChEcomResponse {
	s.StatusCode = &v
	return s
}

func (s *GetItemPubChEcomResponse) SetBody(v *GetItemPubChEcomResponseBody) *GetItemPubChEcomResponse {
	s.Body = v
	return s
}

type GetKeywordChEcomRequest struct {
	ApiVersion  *string `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	Text        *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s GetKeywordChEcomRequest) String() string {
	return tea.Prettify(s)
}

func (s GetKeywordChEcomRequest) GoString() string {
	return s.String()
}

func (s *GetKeywordChEcomRequest) SetApiVersion(v string) *GetKeywordChEcomRequest {
	s.ApiVersion = &v
	return s
}

func (s *GetKeywordChEcomRequest) SetServiceCode(v string) *GetKeywordChEcomRequest {
	s.ServiceCode = &v
	return s
}

func (s *GetKeywordChEcomRequest) SetText(v string) *GetKeywordChEcomRequest {
	s.Text = &v
	return s
}

type GetKeywordChEcomResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetKeywordChEcomResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetKeywordChEcomResponseBody) GoString() string {
	return s.String()
}

func (s *GetKeywordChEcomResponseBody) SetData(v string) *GetKeywordChEcomResponseBody {
	s.Data = &v
	return s
}

func (s *GetKeywordChEcomResponseBody) SetRequestId(v string) *GetKeywordChEcomResponseBody {
	s.RequestId = &v
	return s
}

type GetKeywordChEcomResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetKeywordChEcomResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetKeywordChEcomResponse) String() string {
	return tea.Prettify(s)
}

func (s GetKeywordChEcomResponse) GoString() string {
	return s.String()
}

func (s *GetKeywordChEcomResponse) SetHeaders(v map[string]*string) *GetKeywordChEcomResponse {
	s.Headers = v
	return s
}

func (s *GetKeywordChEcomResponse) SetStatusCode(v int32) *GetKeywordChEcomResponse {
	s.StatusCode = &v
	return s
}

func (s *GetKeywordChEcomResponse) SetBody(v *GetKeywordChEcomResponseBody) *GetKeywordChEcomResponse {
	s.Body = v
	return s
}

type GetKeywordEnEcomRequest struct {
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	Text        *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s GetKeywordEnEcomRequest) String() string {
	return tea.Prettify(s)
}

func (s GetKeywordEnEcomRequest) GoString() string {
	return s.String()
}

func (s *GetKeywordEnEcomRequest) SetServiceCode(v string) *GetKeywordEnEcomRequest {
	s.ServiceCode = &v
	return s
}

func (s *GetKeywordEnEcomRequest) SetText(v string) *GetKeywordEnEcomRequest {
	s.Text = &v
	return s
}

type GetKeywordEnEcomResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetKeywordEnEcomResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetKeywordEnEcomResponseBody) GoString() string {
	return s.String()
}

func (s *GetKeywordEnEcomResponseBody) SetData(v string) *GetKeywordEnEcomResponseBody {
	s.Data = &v
	return s
}

func (s *GetKeywordEnEcomResponseBody) SetRequestId(v string) *GetKeywordEnEcomResponseBody {
	s.RequestId = &v
	return s
}

type GetKeywordEnEcomResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetKeywordEnEcomResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetKeywordEnEcomResponse) String() string {
	return tea.Prettify(s)
}

func (s GetKeywordEnEcomResponse) GoString() string {
	return s.String()
}

func (s *GetKeywordEnEcomResponse) SetHeaders(v map[string]*string) *GetKeywordEnEcomResponse {
	s.Headers = v
	return s
}

func (s *GetKeywordEnEcomResponse) SetStatusCode(v int32) *GetKeywordEnEcomResponse {
	s.StatusCode = &v
	return s
}

func (s *GetKeywordEnEcomResponse) SetBody(v *GetKeywordEnEcomResponseBody) *GetKeywordEnEcomResponse {
	s.Body = v
	return s
}

type GetMedicineChMedicalRequest struct {
	Factory       *string `json:"Factory,omitempty" xml:"Factory,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ServiceCode   *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	Specification *string `json:"Specification,omitempty" xml:"Specification,omitempty"`
	Unit          *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
}

func (s GetMedicineChMedicalRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMedicineChMedicalRequest) GoString() string {
	return s.String()
}

func (s *GetMedicineChMedicalRequest) SetFactory(v string) *GetMedicineChMedicalRequest {
	s.Factory = &v
	return s
}

func (s *GetMedicineChMedicalRequest) SetName(v string) *GetMedicineChMedicalRequest {
	s.Name = &v
	return s
}

func (s *GetMedicineChMedicalRequest) SetServiceCode(v string) *GetMedicineChMedicalRequest {
	s.ServiceCode = &v
	return s
}

func (s *GetMedicineChMedicalRequest) SetSpecification(v string) *GetMedicineChMedicalRequest {
	s.Specification = &v
	return s
}

func (s *GetMedicineChMedicalRequest) SetUnit(v string) *GetMedicineChMedicalRequest {
	s.Unit = &v
	return s
}

type GetMedicineChMedicalResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMedicineChMedicalResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMedicineChMedicalResponseBody) GoString() string {
	return s.String()
}

func (s *GetMedicineChMedicalResponseBody) SetData(v string) *GetMedicineChMedicalResponseBody {
	s.Data = &v
	return s
}

func (s *GetMedicineChMedicalResponseBody) SetRequestId(v string) *GetMedicineChMedicalResponseBody {
	s.RequestId = &v
	return s
}

type GetMedicineChMedicalResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMedicineChMedicalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMedicineChMedicalResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMedicineChMedicalResponse) GoString() string {
	return s.String()
}

func (s *GetMedicineChMedicalResponse) SetHeaders(v map[string]*string) *GetMedicineChMedicalResponse {
	s.Headers = v
	return s
}

func (s *GetMedicineChMedicalResponse) SetStatusCode(v int32) *GetMedicineChMedicalResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMedicineChMedicalResponse) SetBody(v *GetMedicineChMedicalResponseBody) *GetMedicineChMedicalResponse {
	s.Body = v
	return s
}

type GetNerChEcomRequest struct {
	LexerId     *string `json:"LexerId,omitempty" xml:"LexerId,omitempty"`
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	Text        *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s GetNerChEcomRequest) String() string {
	return tea.Prettify(s)
}

func (s GetNerChEcomRequest) GoString() string {
	return s.String()
}

func (s *GetNerChEcomRequest) SetLexerId(v string) *GetNerChEcomRequest {
	s.LexerId = &v
	return s
}

func (s *GetNerChEcomRequest) SetServiceCode(v string) *GetNerChEcomRequest {
	s.ServiceCode = &v
	return s
}

func (s *GetNerChEcomRequest) SetText(v string) *GetNerChEcomRequest {
	s.Text = &v
	return s
}

type GetNerChEcomResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetNerChEcomResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetNerChEcomResponseBody) GoString() string {
	return s.String()
}

func (s *GetNerChEcomResponseBody) SetData(v string) *GetNerChEcomResponseBody {
	s.Data = &v
	return s
}

func (s *GetNerChEcomResponseBody) SetRequestId(v string) *GetNerChEcomResponseBody {
	s.RequestId = &v
	return s
}

type GetNerChEcomResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNerChEcomResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNerChEcomResponse) String() string {
	return tea.Prettify(s)
}

func (s GetNerChEcomResponse) GoString() string {
	return s.String()
}

func (s *GetNerChEcomResponse) SetHeaders(v map[string]*string) *GetNerChEcomResponse {
	s.Headers = v
	return s
}

func (s *GetNerChEcomResponse) SetStatusCode(v int32) *GetNerChEcomResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNerChEcomResponse) SetBody(v *GetNerChEcomResponseBody) *GetNerChEcomResponse {
	s.Body = v
	return s
}

type GetNerChMedicalRequest struct {
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	Text        *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s GetNerChMedicalRequest) String() string {
	return tea.Prettify(s)
}

func (s GetNerChMedicalRequest) GoString() string {
	return s.String()
}

func (s *GetNerChMedicalRequest) SetServiceCode(v string) *GetNerChMedicalRequest {
	s.ServiceCode = &v
	return s
}

func (s *GetNerChMedicalRequest) SetText(v string) *GetNerChMedicalRequest {
	s.Text = &v
	return s
}

type GetNerChMedicalResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetNerChMedicalResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetNerChMedicalResponseBody) GoString() string {
	return s.String()
}

func (s *GetNerChMedicalResponseBody) SetData(v string) *GetNerChMedicalResponseBody {
	s.Data = &v
	return s
}

func (s *GetNerChMedicalResponseBody) SetRequestId(v string) *GetNerChMedicalResponseBody {
	s.RequestId = &v
	return s
}

type GetNerChMedicalResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNerChMedicalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNerChMedicalResponse) String() string {
	return tea.Prettify(s)
}

func (s GetNerChMedicalResponse) GoString() string {
	return s.String()
}

func (s *GetNerChMedicalResponse) SetHeaders(v map[string]*string) *GetNerChMedicalResponse {
	s.Headers = v
	return s
}

func (s *GetNerChMedicalResponse) SetStatusCode(v int32) *GetNerChMedicalResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNerChMedicalResponse) SetBody(v *GetNerChMedicalResponseBody) *GetNerChMedicalResponse {
	s.Body = v
	return s
}

type GetNerCustomizedChEcomRequest struct {
	LexerId     *string `json:"LexerId,omitempty" xml:"LexerId,omitempty"`
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	Text        *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s GetNerCustomizedChEcomRequest) String() string {
	return tea.Prettify(s)
}

func (s GetNerCustomizedChEcomRequest) GoString() string {
	return s.String()
}

func (s *GetNerCustomizedChEcomRequest) SetLexerId(v string) *GetNerCustomizedChEcomRequest {
	s.LexerId = &v
	return s
}

func (s *GetNerCustomizedChEcomRequest) SetServiceCode(v string) *GetNerCustomizedChEcomRequest {
	s.ServiceCode = &v
	return s
}

func (s *GetNerCustomizedChEcomRequest) SetText(v string) *GetNerCustomizedChEcomRequest {
	s.Text = &v
	return s
}

type GetNerCustomizedChEcomResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetNerCustomizedChEcomResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetNerCustomizedChEcomResponseBody) GoString() string {
	return s.String()
}

func (s *GetNerCustomizedChEcomResponseBody) SetData(v string) *GetNerCustomizedChEcomResponseBody {
	s.Data = &v
	return s
}

func (s *GetNerCustomizedChEcomResponseBody) SetRequestId(v string) *GetNerCustomizedChEcomResponseBody {
	s.RequestId = &v
	return s
}

type GetNerCustomizedChEcomResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNerCustomizedChEcomResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNerCustomizedChEcomResponse) String() string {
	return tea.Prettify(s)
}

func (s GetNerCustomizedChEcomResponse) GoString() string {
	return s.String()
}

func (s *GetNerCustomizedChEcomResponse) SetHeaders(v map[string]*string) *GetNerCustomizedChEcomResponse {
	s.Headers = v
	return s
}

func (s *GetNerCustomizedChEcomResponse) SetStatusCode(v int32) *GetNerCustomizedChEcomResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNerCustomizedChEcomResponse) SetBody(v *GetNerCustomizedChEcomResponseBody) *GetNerCustomizedChEcomResponse {
	s.Body = v
	return s
}

type GetNerCustomizedSeaEcomRequest struct {
	Language    *string `json:"Language,omitempty" xml:"Language,omitempty"`
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	Text        *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s GetNerCustomizedSeaEcomRequest) String() string {
	return tea.Prettify(s)
}

func (s GetNerCustomizedSeaEcomRequest) GoString() string {
	return s.String()
}

func (s *GetNerCustomizedSeaEcomRequest) SetLanguage(v string) *GetNerCustomizedSeaEcomRequest {
	s.Language = &v
	return s
}

func (s *GetNerCustomizedSeaEcomRequest) SetServiceCode(v string) *GetNerCustomizedSeaEcomRequest {
	s.ServiceCode = &v
	return s
}

func (s *GetNerCustomizedSeaEcomRequest) SetText(v string) *GetNerCustomizedSeaEcomRequest {
	s.Text = &v
	return s
}

type GetNerCustomizedSeaEcomResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetNerCustomizedSeaEcomResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetNerCustomizedSeaEcomResponseBody) GoString() string {
	return s.String()
}

func (s *GetNerCustomizedSeaEcomResponseBody) SetData(v string) *GetNerCustomizedSeaEcomResponseBody {
	s.Data = &v
	return s
}

func (s *GetNerCustomizedSeaEcomResponseBody) SetRequestId(v string) *GetNerCustomizedSeaEcomResponseBody {
	s.RequestId = &v
	return s
}

type GetNerCustomizedSeaEcomResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNerCustomizedSeaEcomResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNerCustomizedSeaEcomResponse) String() string {
	return tea.Prettify(s)
}

func (s GetNerCustomizedSeaEcomResponse) GoString() string {
	return s.String()
}

func (s *GetNerCustomizedSeaEcomResponse) SetHeaders(v map[string]*string) *GetNerCustomizedSeaEcomResponse {
	s.Headers = v
	return s
}

func (s *GetNerCustomizedSeaEcomResponse) SetStatusCode(v int32) *GetNerCustomizedSeaEcomResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNerCustomizedSeaEcomResponse) SetBody(v *GetNerCustomizedSeaEcomResponseBody) *GetNerCustomizedSeaEcomResponse {
	s.Body = v
	return s
}

type GetOpenNLURequest struct {
	Examples    *string `json:"Examples,omitempty" xml:"Examples,omitempty"`
	Labels      *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	Sentence    *string `json:"Sentence,omitempty" xml:"Sentence,omitempty"`
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	Task        *string `json:"Task,omitempty" xml:"Task,omitempty"`
}

func (s GetOpenNLURequest) String() string {
	return tea.Prettify(s)
}

func (s GetOpenNLURequest) GoString() string {
	return s.String()
}

func (s *GetOpenNLURequest) SetExamples(v string) *GetOpenNLURequest {
	s.Examples = &v
	return s
}

func (s *GetOpenNLURequest) SetLabels(v string) *GetOpenNLURequest {
	s.Labels = &v
	return s
}

func (s *GetOpenNLURequest) SetSentence(v string) *GetOpenNLURequest {
	s.Sentence = &v
	return s
}

func (s *GetOpenNLURequest) SetServiceCode(v string) *GetOpenNLURequest {
	s.ServiceCode = &v
	return s
}

func (s *GetOpenNLURequest) SetTask(v string) *GetOpenNLURequest {
	s.Task = &v
	return s
}

type GetOpenNLUResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetOpenNLUResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOpenNLUResponseBody) GoString() string {
	return s.String()
}

func (s *GetOpenNLUResponseBody) SetData(v string) *GetOpenNLUResponseBody {
	s.Data = &v
	return s
}

func (s *GetOpenNLUResponseBody) SetRequestId(v string) *GetOpenNLUResponseBody {
	s.RequestId = &v
	return s
}

type GetOpenNLUResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOpenNLUResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOpenNLUResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOpenNLUResponse) GoString() string {
	return s.String()
}

func (s *GetOpenNLUResponse) SetHeaders(v map[string]*string) *GetOpenNLUResponse {
	s.Headers = v
	return s
}

func (s *GetOpenNLUResponse) SetStatusCode(v int32) *GetOpenNLUResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOpenNLUResponse) SetBody(v *GetOpenNLUResponseBody) *GetOpenNLUResponse {
	s.Body = v
	return s
}

type GetOpenNLUHighRecallRequest struct {
	Examples    *string `json:"Examples,omitempty" xml:"Examples,omitempty"`
	Labels      *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	Sentence    *string `json:"Sentence,omitempty" xml:"Sentence,omitempty"`
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	Task        *string `json:"Task,omitempty" xml:"Task,omitempty"`
}

func (s GetOpenNLUHighRecallRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOpenNLUHighRecallRequest) GoString() string {
	return s.String()
}

func (s *GetOpenNLUHighRecallRequest) SetExamples(v string) *GetOpenNLUHighRecallRequest {
	s.Examples = &v
	return s
}

func (s *GetOpenNLUHighRecallRequest) SetLabels(v string) *GetOpenNLUHighRecallRequest {
	s.Labels = &v
	return s
}

func (s *GetOpenNLUHighRecallRequest) SetSentence(v string) *GetOpenNLUHighRecallRequest {
	s.Sentence = &v
	return s
}

func (s *GetOpenNLUHighRecallRequest) SetServiceCode(v string) *GetOpenNLUHighRecallRequest {
	s.ServiceCode = &v
	return s
}

func (s *GetOpenNLUHighRecallRequest) SetTask(v string) *GetOpenNLUHighRecallRequest {
	s.Task = &v
	return s
}

type GetOpenNLUHighRecallResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetOpenNLUHighRecallResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOpenNLUHighRecallResponseBody) GoString() string {
	return s.String()
}

func (s *GetOpenNLUHighRecallResponseBody) SetData(v string) *GetOpenNLUHighRecallResponseBody {
	s.Data = &v
	return s
}

func (s *GetOpenNLUHighRecallResponseBody) SetRequestId(v string) *GetOpenNLUHighRecallResponseBody {
	s.RequestId = &v
	return s
}

type GetOpenNLUHighRecallResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOpenNLUHighRecallResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOpenNLUHighRecallResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOpenNLUHighRecallResponse) GoString() string {
	return s.String()
}

func (s *GetOpenNLUHighRecallResponse) SetHeaders(v map[string]*string) *GetOpenNLUHighRecallResponse {
	s.Headers = v
	return s
}

func (s *GetOpenNLUHighRecallResponse) SetStatusCode(v int32) *GetOpenNLUHighRecallResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOpenNLUHighRecallResponse) SetBody(v *GetOpenNLUHighRecallResponseBody) *GetOpenNLUHighRecallResponse {
	s.Body = v
	return s
}

type GetOperationChMedicalRequest struct {
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
}

func (s GetOperationChMedicalRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOperationChMedicalRequest) GoString() string {
	return s.String()
}

func (s *GetOperationChMedicalRequest) SetName(v string) *GetOperationChMedicalRequest {
	s.Name = &v
	return s
}

func (s *GetOperationChMedicalRequest) SetServiceCode(v string) *GetOperationChMedicalRequest {
	s.ServiceCode = &v
	return s
}

type GetOperationChMedicalResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetOperationChMedicalResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOperationChMedicalResponseBody) GoString() string {
	return s.String()
}

func (s *GetOperationChMedicalResponseBody) SetData(v string) *GetOperationChMedicalResponseBody {
	s.Data = &v
	return s
}

func (s *GetOperationChMedicalResponseBody) SetRequestId(v string) *GetOperationChMedicalResponseBody {
	s.RequestId = &v
	return s
}

type GetOperationChMedicalResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOperationChMedicalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOperationChMedicalResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOperationChMedicalResponse) GoString() string {
	return s.String()
}

func (s *GetOperationChMedicalResponse) SetHeaders(v map[string]*string) *GetOperationChMedicalResponse {
	s.Headers = v
	return s
}

func (s *GetOperationChMedicalResponse) SetStatusCode(v int32) *GetOperationChMedicalResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOperationChMedicalResponse) SetBody(v *GetOperationChMedicalResponseBody) *GetOperationChMedicalResponse {
	s.Body = v
	return s
}

type GetPosChEcomRequest struct {
	OutType     *string `json:"OutType,omitempty" xml:"OutType,omitempty"`
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	Text        *string `json:"Text,omitempty" xml:"Text,omitempty"`
	TokenizerId *string `json:"TokenizerId,omitempty" xml:"TokenizerId,omitempty"`
}

func (s GetPosChEcomRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPosChEcomRequest) GoString() string {
	return s.String()
}

func (s *GetPosChEcomRequest) SetOutType(v string) *GetPosChEcomRequest {
	s.OutType = &v
	return s
}

func (s *GetPosChEcomRequest) SetServiceCode(v string) *GetPosChEcomRequest {
	s.ServiceCode = &v
	return s
}

func (s *GetPosChEcomRequest) SetText(v string) *GetPosChEcomRequest {
	s.Text = &v
	return s
}

func (s *GetPosChEcomRequest) SetTokenizerId(v string) *GetPosChEcomRequest {
	s.TokenizerId = &v
	return s
}

type GetPosChEcomResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPosChEcomResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPosChEcomResponseBody) GoString() string {
	return s.String()
}

func (s *GetPosChEcomResponseBody) SetData(v string) *GetPosChEcomResponseBody {
	s.Data = &v
	return s
}

func (s *GetPosChEcomResponseBody) SetRequestId(v string) *GetPosChEcomResponseBody {
	s.RequestId = &v
	return s
}

type GetPosChEcomResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPosChEcomResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPosChEcomResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPosChEcomResponse) GoString() string {
	return s.String()
}

func (s *GetPosChEcomResponse) SetHeaders(v map[string]*string) *GetPosChEcomResponse {
	s.Headers = v
	return s
}

func (s *GetPosChEcomResponse) SetStatusCode(v int32) *GetPosChEcomResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPosChEcomResponse) SetBody(v *GetPosChEcomResponseBody) *GetPosChEcomResponse {
	s.Body = v
	return s
}

type GetPosChGeneralRequest struct {
	OutType     *string `json:"OutType,omitempty" xml:"OutType,omitempty"`
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	Text        *string `json:"Text,omitempty" xml:"Text,omitempty"`
	TokenizerId *string `json:"TokenizerId,omitempty" xml:"TokenizerId,omitempty"`
}

func (s GetPosChGeneralRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPosChGeneralRequest) GoString() string {
	return s.String()
}

func (s *GetPosChGeneralRequest) SetOutType(v string) *GetPosChGeneralRequest {
	s.OutType = &v
	return s
}

func (s *GetPosChGeneralRequest) SetServiceCode(v string) *GetPosChGeneralRequest {
	s.ServiceCode = &v
	return s
}

func (s *GetPosChGeneralRequest) SetText(v string) *GetPosChGeneralRequest {
	s.Text = &v
	return s
}

func (s *GetPosChGeneralRequest) SetTokenizerId(v string) *GetPosChGeneralRequest {
	s.TokenizerId = &v
	return s
}

type GetPosChGeneralResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPosChGeneralResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPosChGeneralResponseBody) GoString() string {
	return s.String()
}

func (s *GetPosChGeneralResponseBody) SetData(v string) *GetPosChGeneralResponseBody {
	s.Data = &v
	return s
}

func (s *GetPosChGeneralResponseBody) SetRequestId(v string) *GetPosChGeneralResponseBody {
	s.RequestId = &v
	return s
}

type GetPosChGeneralResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPosChGeneralResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPosChGeneralResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPosChGeneralResponse) GoString() string {
	return s.String()
}

func (s *GetPosChGeneralResponse) SetHeaders(v map[string]*string) *GetPosChGeneralResponse {
	s.Headers = v
	return s
}

func (s *GetPosChGeneralResponse) SetStatusCode(v int32) *GetPosChGeneralResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPosChGeneralResponse) SetBody(v *GetPosChGeneralResponseBody) *GetPosChGeneralResponse {
	s.Body = v
	return s
}

type GetPriceChEcomRequest struct {
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	Text        *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s GetPriceChEcomRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPriceChEcomRequest) GoString() string {
	return s.String()
}

func (s *GetPriceChEcomRequest) SetServiceCode(v string) *GetPriceChEcomRequest {
	s.ServiceCode = &v
	return s
}

func (s *GetPriceChEcomRequest) SetText(v string) *GetPriceChEcomRequest {
	s.Text = &v
	return s
}

type GetPriceChEcomResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPriceChEcomResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPriceChEcomResponseBody) GoString() string {
	return s.String()
}

func (s *GetPriceChEcomResponseBody) SetData(v string) *GetPriceChEcomResponseBody {
	s.Data = &v
	return s
}

func (s *GetPriceChEcomResponseBody) SetRequestId(v string) *GetPriceChEcomResponseBody {
	s.RequestId = &v
	return s
}

type GetPriceChEcomResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPriceChEcomResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPriceChEcomResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPriceChEcomResponse) GoString() string {
	return s.String()
}

func (s *GetPriceChEcomResponse) SetHeaders(v map[string]*string) *GetPriceChEcomResponse {
	s.Headers = v
	return s
}

func (s *GetPriceChEcomResponse) SetStatusCode(v int32) *GetPriceChEcomResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPriceChEcomResponse) SetBody(v *GetPriceChEcomResponseBody) *GetPriceChEcomResponse {
	s.Body = v
	return s
}

type GetSSETestRequest struct {
	Params      *string `json:"Params,omitempty" xml:"Params,omitempty"`
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
}

func (s GetSSETestRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSSETestRequest) GoString() string {
	return s.String()
}

func (s *GetSSETestRequest) SetParams(v string) *GetSSETestRequest {
	s.Params = &v
	return s
}

func (s *GetSSETestRequest) SetServiceCode(v string) *GetSSETestRequest {
	s.ServiceCode = &v
	return s
}

type GetSSETestResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSSETestResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSSETestResponseBody) GoString() string {
	return s.String()
}

func (s *GetSSETestResponseBody) SetData(v string) *GetSSETestResponseBody {
	s.Data = &v
	return s
}

func (s *GetSSETestResponseBody) SetRequestId(v string) *GetSSETestResponseBody {
	s.RequestId = &v
	return s
}

type GetSSETestResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSSETestResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSSETestResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSSETestResponse) GoString() string {
	return s.String()
}

func (s *GetSSETestResponse) SetHeaders(v map[string]*string) *GetSSETestResponse {
	s.Headers = v
	return s
}

func (s *GetSSETestResponse) SetStatusCode(v int32) *GetSSETestResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSSETestResponse) SetBody(v *GetSSETestResponseBody) *GetSSETestResponse {
	s.Body = v
	return s
}

type GetSaChGeneralRequest struct {
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	Text        *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s GetSaChGeneralRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSaChGeneralRequest) GoString() string {
	return s.String()
}

func (s *GetSaChGeneralRequest) SetServiceCode(v string) *GetSaChGeneralRequest {
	s.ServiceCode = &v
	return s
}

func (s *GetSaChGeneralRequest) SetText(v string) *GetSaChGeneralRequest {
	s.Text = &v
	return s
}

type GetSaChGeneralResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSaChGeneralResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSaChGeneralResponseBody) GoString() string {
	return s.String()
}

func (s *GetSaChGeneralResponseBody) SetData(v string) *GetSaChGeneralResponseBody {
	s.Data = &v
	return s
}

func (s *GetSaChGeneralResponseBody) SetRequestId(v string) *GetSaChGeneralResponseBody {
	s.RequestId = &v
	return s
}

type GetSaChGeneralResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSaChGeneralResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSaChGeneralResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSaChGeneralResponse) GoString() string {
	return s.String()
}

func (s *GetSaChGeneralResponse) SetHeaders(v map[string]*string) *GetSaChGeneralResponse {
	s.Headers = v
	return s
}

func (s *GetSaChGeneralResponse) SetStatusCode(v int32) *GetSaChGeneralResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSaChGeneralResponse) SetBody(v *GetSaChGeneralResponseBody) *GetSaChGeneralResponse {
	s.Body = v
	return s
}

type GetSaSeaEcomRequest struct {
	Language    *string `json:"Language,omitempty" xml:"Language,omitempty"`
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	Text        *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s GetSaSeaEcomRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSaSeaEcomRequest) GoString() string {
	return s.String()
}

func (s *GetSaSeaEcomRequest) SetLanguage(v string) *GetSaSeaEcomRequest {
	s.Language = &v
	return s
}

func (s *GetSaSeaEcomRequest) SetServiceCode(v string) *GetSaSeaEcomRequest {
	s.ServiceCode = &v
	return s
}

func (s *GetSaSeaEcomRequest) SetText(v string) *GetSaSeaEcomRequest {
	s.Text = &v
	return s
}

type GetSaSeaEcomResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSaSeaEcomResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSaSeaEcomResponseBody) GoString() string {
	return s.String()
}

func (s *GetSaSeaEcomResponseBody) SetData(v string) *GetSaSeaEcomResponseBody {
	s.Data = &v
	return s
}

func (s *GetSaSeaEcomResponseBody) SetRequestId(v string) *GetSaSeaEcomResponseBody {
	s.RequestId = &v
	return s
}

type GetSaSeaEcomResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSaSeaEcomResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSaSeaEcomResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSaSeaEcomResponse) GoString() string {
	return s.String()
}

func (s *GetSaSeaEcomResponse) SetHeaders(v map[string]*string) *GetSaSeaEcomResponse {
	s.Headers = v
	return s
}

func (s *GetSaSeaEcomResponse) SetStatusCode(v int32) *GetSaSeaEcomResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSaSeaEcomResponse) SetBody(v *GetSaSeaEcomResponseBody) *GetSaSeaEcomResponse {
	s.Body = v
	return s
}

type GetServiceDataImportStatusRequest struct {
	DataImportIds []*int64 `json:"DataImportIds,omitempty" xml:"DataImportIds,omitempty" type:"Repeated"`
}

func (s GetServiceDataImportStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetServiceDataImportStatusRequest) GoString() string {
	return s.String()
}

func (s *GetServiceDataImportStatusRequest) SetDataImportIds(v []*int64) *GetServiceDataImportStatusRequest {
	s.DataImportIds = v
	return s
}

type GetServiceDataImportStatusShrinkRequest struct {
	DataImportIdsShrink *string `json:"DataImportIds,omitempty" xml:"DataImportIds,omitempty"`
}

func (s GetServiceDataImportStatusShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetServiceDataImportStatusShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetServiceDataImportStatusShrinkRequest) SetDataImportIdsShrink(v string) *GetServiceDataImportStatusShrinkRequest {
	s.DataImportIdsShrink = &v
	return s
}

type GetServiceDataImportStatusResponseBody struct {
	Code      *int32                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      map[string]*DataValue `json:"Data,omitempty" xml:"Data,omitempty"`
	Msg       *string               `json:"Msg,omitempty" xml:"Msg,omitempty"`
	RequestId *string               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetServiceDataImportStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetServiceDataImportStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceDataImportStatusResponseBody) SetCode(v int32) *GetServiceDataImportStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetServiceDataImportStatusResponseBody) SetData(v map[string]*DataValue) *GetServiceDataImportStatusResponseBody {
	s.Data = v
	return s
}

func (s *GetServiceDataImportStatusResponseBody) SetMsg(v string) *GetServiceDataImportStatusResponseBody {
	s.Msg = &v
	return s
}

func (s *GetServiceDataImportStatusResponseBody) SetRequestId(v string) *GetServiceDataImportStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetServiceDataImportStatusResponseBody) SetSuccess(v bool) *GetServiceDataImportStatusResponseBody {
	s.Success = &v
	return s
}

type GetServiceDataImportStatusResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetServiceDataImportStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServiceDataImportStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetServiceDataImportStatusResponse) GoString() string {
	return s.String()
}

func (s *GetServiceDataImportStatusResponse) SetHeaders(v map[string]*string) *GetServiceDataImportStatusResponse {
	s.Headers = v
	return s
}

func (s *GetServiceDataImportStatusResponse) SetStatusCode(v int32) *GetServiceDataImportStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceDataImportStatusResponse) SetBody(v *GetServiceDataImportStatusResponseBody) *GetServiceDataImportStatusResponse {
	s.Body = v
	return s
}

type GetSimilarityChMedicalRequest struct {
	OriginQ     *string `json:"OriginQ,omitempty" xml:"OriginQ,omitempty"`
	OriginT     *string `json:"OriginT,omitempty" xml:"OriginT,omitempty"`
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
}

func (s GetSimilarityChMedicalRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSimilarityChMedicalRequest) GoString() string {
	return s.String()
}

func (s *GetSimilarityChMedicalRequest) SetOriginQ(v string) *GetSimilarityChMedicalRequest {
	s.OriginQ = &v
	return s
}

func (s *GetSimilarityChMedicalRequest) SetOriginT(v string) *GetSimilarityChMedicalRequest {
	s.OriginT = &v
	return s
}

func (s *GetSimilarityChMedicalRequest) SetServiceCode(v string) *GetSimilarityChMedicalRequest {
	s.ServiceCode = &v
	return s
}

type GetSimilarityChMedicalResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSimilarityChMedicalResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSimilarityChMedicalResponseBody) GoString() string {
	return s.String()
}

func (s *GetSimilarityChMedicalResponseBody) SetData(v string) *GetSimilarityChMedicalResponseBody {
	s.Data = &v
	return s
}

func (s *GetSimilarityChMedicalResponseBody) SetRequestId(v string) *GetSimilarityChMedicalResponseBody {
	s.RequestId = &v
	return s
}

type GetSimilarityChMedicalResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSimilarityChMedicalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSimilarityChMedicalResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSimilarityChMedicalResponse) GoString() string {
	return s.String()
}

func (s *GetSimilarityChMedicalResponse) SetHeaders(v map[string]*string) *GetSimilarityChMedicalResponse {
	s.Headers = v
	return s
}

func (s *GetSimilarityChMedicalResponse) SetStatusCode(v int32) *GetSimilarityChMedicalResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSimilarityChMedicalResponse) SetBody(v *GetSimilarityChMedicalResponseBody) *GetSimilarityChMedicalResponse {
	s.Body = v
	return s
}

type GetSummaryChEcomRequest struct {
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	Text        *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s GetSummaryChEcomRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSummaryChEcomRequest) GoString() string {
	return s.String()
}

func (s *GetSummaryChEcomRequest) SetServiceCode(v string) *GetSummaryChEcomRequest {
	s.ServiceCode = &v
	return s
}

func (s *GetSummaryChEcomRequest) SetText(v string) *GetSummaryChEcomRequest {
	s.Text = &v
	return s
}

type GetSummaryChEcomResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSummaryChEcomResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSummaryChEcomResponseBody) GoString() string {
	return s.String()
}

func (s *GetSummaryChEcomResponseBody) SetData(v string) *GetSummaryChEcomResponseBody {
	s.Data = &v
	return s
}

func (s *GetSummaryChEcomResponseBody) SetRequestId(v string) *GetSummaryChEcomResponseBody {
	s.RequestId = &v
	return s
}

type GetSummaryChEcomResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSummaryChEcomResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSummaryChEcomResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSummaryChEcomResponse) GoString() string {
	return s.String()
}

func (s *GetSummaryChEcomResponse) SetHeaders(v map[string]*string) *GetSummaryChEcomResponse {
	s.Headers = v
	return s
}

func (s *GetSummaryChEcomResponse) SetStatusCode(v int32) *GetSummaryChEcomResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSummaryChEcomResponse) SetBody(v *GetSummaryChEcomResponseBody) *GetSummaryChEcomResponse {
	s.Body = v
	return s
}

type GetTableQAServiceInfoByIdRequest struct {
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	ServiceId   *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s GetTableQAServiceInfoByIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTableQAServiceInfoByIdRequest) GoString() string {
	return s.String()
}

func (s *GetTableQAServiceInfoByIdRequest) SetServiceCode(v string) *GetTableQAServiceInfoByIdRequest {
	s.ServiceCode = &v
	return s
}

func (s *GetTableQAServiceInfoByIdRequest) SetServiceId(v string) *GetTableQAServiceInfoByIdRequest {
	s.ServiceId = &v
	return s
}

type GetTableQAServiceInfoByIdResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetTableQAServiceInfoByIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTableQAServiceInfoByIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetTableQAServiceInfoByIdResponseBody) SetData(v string) *GetTableQAServiceInfoByIdResponseBody {
	s.Data = &v
	return s
}

func (s *GetTableQAServiceInfoByIdResponseBody) SetRequestId(v string) *GetTableQAServiceInfoByIdResponseBody {
	s.RequestId = &v
	return s
}

type GetTableQAServiceInfoByIdResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTableQAServiceInfoByIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTableQAServiceInfoByIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTableQAServiceInfoByIdResponse) GoString() string {
	return s.String()
}

func (s *GetTableQAServiceInfoByIdResponse) SetHeaders(v map[string]*string) *GetTableQAServiceInfoByIdResponse {
	s.Headers = v
	return s
}

func (s *GetTableQAServiceInfoByIdResponse) SetStatusCode(v int32) *GetTableQAServiceInfoByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTableQAServiceInfoByIdResponse) SetBody(v *GetTableQAServiceInfoByIdResponseBody) *GetTableQAServiceInfoByIdResponse {
	s.Body = v
	return s
}

type GetTcChEcomRequest struct {
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	Text        *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s GetTcChEcomRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTcChEcomRequest) GoString() string {
	return s.String()
}

func (s *GetTcChEcomRequest) SetServiceCode(v string) *GetTcChEcomRequest {
	s.ServiceCode = &v
	return s
}

func (s *GetTcChEcomRequest) SetText(v string) *GetTcChEcomRequest {
	s.Text = &v
	return s
}

type GetTcChEcomResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetTcChEcomResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTcChEcomResponseBody) GoString() string {
	return s.String()
}

func (s *GetTcChEcomResponseBody) SetData(v string) *GetTcChEcomResponseBody {
	s.Data = &v
	return s
}

func (s *GetTcChEcomResponseBody) SetRequestId(v string) *GetTcChEcomResponseBody {
	s.RequestId = &v
	return s
}

type GetTcChEcomResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTcChEcomResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTcChEcomResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTcChEcomResponse) GoString() string {
	return s.String()
}

func (s *GetTcChEcomResponse) SetHeaders(v map[string]*string) *GetTcChEcomResponse {
	s.Headers = v
	return s
}

func (s *GetTcChEcomResponse) SetStatusCode(v int32) *GetTcChEcomResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTcChEcomResponse) SetBody(v *GetTcChEcomResponseBody) *GetTcChEcomResponse {
	s.Body = v
	return s
}

type GetTcChGeneralRequest struct {
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	Text        *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s GetTcChGeneralRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTcChGeneralRequest) GoString() string {
	return s.String()
}

func (s *GetTcChGeneralRequest) SetServiceCode(v string) *GetTcChGeneralRequest {
	s.ServiceCode = &v
	return s
}

func (s *GetTcChGeneralRequest) SetText(v string) *GetTcChGeneralRequest {
	s.Text = &v
	return s
}

type GetTcChGeneralResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetTcChGeneralResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTcChGeneralResponseBody) GoString() string {
	return s.String()
}

func (s *GetTcChGeneralResponseBody) SetData(v string) *GetTcChGeneralResponseBody {
	s.Data = &v
	return s
}

func (s *GetTcChGeneralResponseBody) SetRequestId(v string) *GetTcChGeneralResponseBody {
	s.RequestId = &v
	return s
}

type GetTcChGeneralResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTcChGeneralResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTcChGeneralResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTcChGeneralResponse) GoString() string {
	return s.String()
}

func (s *GetTcChGeneralResponse) SetHeaders(v map[string]*string) *GetTcChGeneralResponse {
	s.Headers = v
	return s
}

func (s *GetTcChGeneralResponse) SetStatusCode(v int32) *GetTcChGeneralResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTcChGeneralResponse) SetBody(v *GetTcChGeneralResponseBody) *GetTcChGeneralResponse {
	s.Body = v
	return s
}

type GetTsChEcomRequest struct {
	OriginQ     *string `json:"OriginQ,omitempty" xml:"OriginQ,omitempty"`
	OriginT     *string `json:"OriginT,omitempty" xml:"OriginT,omitempty"`
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetTsChEcomRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTsChEcomRequest) GoString() string {
	return s.String()
}

func (s *GetTsChEcomRequest) SetOriginQ(v string) *GetTsChEcomRequest {
	s.OriginQ = &v
	return s
}

func (s *GetTsChEcomRequest) SetOriginT(v string) *GetTsChEcomRequest {
	s.OriginT = &v
	return s
}

func (s *GetTsChEcomRequest) SetServiceCode(v string) *GetTsChEcomRequest {
	s.ServiceCode = &v
	return s
}

func (s *GetTsChEcomRequest) SetType(v string) *GetTsChEcomRequest {
	s.Type = &v
	return s
}

type GetTsChEcomResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetTsChEcomResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTsChEcomResponseBody) GoString() string {
	return s.String()
}

func (s *GetTsChEcomResponseBody) SetData(v string) *GetTsChEcomResponseBody {
	s.Data = &v
	return s
}

func (s *GetTsChEcomResponseBody) SetRequestId(v string) *GetTsChEcomResponseBody {
	s.RequestId = &v
	return s
}

type GetTsChEcomResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTsChEcomResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTsChEcomResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTsChEcomResponse) GoString() string {
	return s.String()
}

func (s *GetTsChEcomResponse) SetHeaders(v map[string]*string) *GetTsChEcomResponse {
	s.Headers = v
	return s
}

func (s *GetTsChEcomResponse) SetStatusCode(v int32) *GetTsChEcomResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTsChEcomResponse) SetBody(v *GetTsChEcomResponseBody) *GetTsChEcomResponse {
	s.Body = v
	return s
}

type GetUserUploadSignRequest struct {
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
}

func (s GetUserUploadSignRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserUploadSignRequest) GoString() string {
	return s.String()
}

func (s *GetUserUploadSignRequest) SetServiceCode(v string) *GetUserUploadSignRequest {
	s.ServiceCode = &v
	return s
}

type GetUserUploadSignResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetUserUploadSignResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserUploadSignResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserUploadSignResponseBody) SetData(v string) *GetUserUploadSignResponseBody {
	s.Data = &v
	return s
}

func (s *GetUserUploadSignResponseBody) SetRequestId(v string) *GetUserUploadSignResponseBody {
	s.RequestId = &v
	return s
}

type GetUserUploadSignResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserUploadSignResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserUploadSignResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserUploadSignResponse) GoString() string {
	return s.String()
}

func (s *GetUserUploadSignResponse) SetHeaders(v map[string]*string) *GetUserUploadSignResponse {
	s.Headers = v
	return s
}

func (s *GetUserUploadSignResponse) SetStatusCode(v int32) *GetUserUploadSignResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserUploadSignResponse) SetBody(v *GetUserUploadSignResponseBody) *GetUserUploadSignResponse {
	s.Body = v
	return s
}

type GetWeChCommentRequest struct {
	Operation   *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	Size        *string `json:"Size,omitempty" xml:"Size,omitempty"`
	Text        *string `json:"Text,omitempty" xml:"Text,omitempty"`
	TokenizerId *string `json:"TokenizerId,omitempty" xml:"TokenizerId,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetWeChCommentRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWeChCommentRequest) GoString() string {
	return s.String()
}

func (s *GetWeChCommentRequest) SetOperation(v string) *GetWeChCommentRequest {
	s.Operation = &v
	return s
}

func (s *GetWeChCommentRequest) SetServiceCode(v string) *GetWeChCommentRequest {
	s.ServiceCode = &v
	return s
}

func (s *GetWeChCommentRequest) SetSize(v string) *GetWeChCommentRequest {
	s.Size = &v
	return s
}

func (s *GetWeChCommentRequest) SetText(v string) *GetWeChCommentRequest {
	s.Text = &v
	return s
}

func (s *GetWeChCommentRequest) SetTokenizerId(v string) *GetWeChCommentRequest {
	s.TokenizerId = &v
	return s
}

func (s *GetWeChCommentRequest) SetType(v string) *GetWeChCommentRequest {
	s.Type = &v
	return s
}

type GetWeChCommentResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetWeChCommentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWeChCommentResponseBody) GoString() string {
	return s.String()
}

func (s *GetWeChCommentResponseBody) SetData(v string) *GetWeChCommentResponseBody {
	s.Data = &v
	return s
}

func (s *GetWeChCommentResponseBody) SetRequestId(v string) *GetWeChCommentResponseBody {
	s.RequestId = &v
	return s
}

type GetWeChCommentResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWeChCommentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWeChCommentResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWeChCommentResponse) GoString() string {
	return s.String()
}

func (s *GetWeChCommentResponse) SetHeaders(v map[string]*string) *GetWeChCommentResponse {
	s.Headers = v
	return s
}

func (s *GetWeChCommentResponse) SetStatusCode(v int32) *GetWeChCommentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWeChCommentResponse) SetBody(v *GetWeChCommentResponseBody) *GetWeChCommentResponse {
	s.Body = v
	return s
}

type GetWeChEcomRequest struct {
	Operation   *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	Size        *string `json:"Size,omitempty" xml:"Size,omitempty"`
	Text        *string `json:"Text,omitempty" xml:"Text,omitempty"`
	TokenizerId *string `json:"TokenizerId,omitempty" xml:"TokenizerId,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetWeChEcomRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWeChEcomRequest) GoString() string {
	return s.String()
}

func (s *GetWeChEcomRequest) SetOperation(v string) *GetWeChEcomRequest {
	s.Operation = &v
	return s
}

func (s *GetWeChEcomRequest) SetServiceCode(v string) *GetWeChEcomRequest {
	s.ServiceCode = &v
	return s
}

func (s *GetWeChEcomRequest) SetSize(v string) *GetWeChEcomRequest {
	s.Size = &v
	return s
}

func (s *GetWeChEcomRequest) SetText(v string) *GetWeChEcomRequest {
	s.Text = &v
	return s
}

func (s *GetWeChEcomRequest) SetTokenizerId(v string) *GetWeChEcomRequest {
	s.TokenizerId = &v
	return s
}

func (s *GetWeChEcomRequest) SetType(v string) *GetWeChEcomRequest {
	s.Type = &v
	return s
}

type GetWeChEcomResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetWeChEcomResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWeChEcomResponseBody) GoString() string {
	return s.String()
}

func (s *GetWeChEcomResponseBody) SetData(v string) *GetWeChEcomResponseBody {
	s.Data = &v
	return s
}

func (s *GetWeChEcomResponseBody) SetRequestId(v string) *GetWeChEcomResponseBody {
	s.RequestId = &v
	return s
}

type GetWeChEcomResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWeChEcomResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWeChEcomResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWeChEcomResponse) GoString() string {
	return s.String()
}

func (s *GetWeChEcomResponse) SetHeaders(v map[string]*string) *GetWeChEcomResponse {
	s.Headers = v
	return s
}

func (s *GetWeChEcomResponse) SetStatusCode(v int32) *GetWeChEcomResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWeChEcomResponse) SetBody(v *GetWeChEcomResponseBody) *GetWeChEcomResponse {
	s.Body = v
	return s
}

type GetWeChEntertainmentRequest struct {
	Operation   *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	Size        *string `json:"Size,omitempty" xml:"Size,omitempty"`
	Text        *string `json:"Text,omitempty" xml:"Text,omitempty"`
	TokenizerId *string `json:"TokenizerId,omitempty" xml:"TokenizerId,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetWeChEntertainmentRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWeChEntertainmentRequest) GoString() string {
	return s.String()
}

func (s *GetWeChEntertainmentRequest) SetOperation(v string) *GetWeChEntertainmentRequest {
	s.Operation = &v
	return s
}

func (s *GetWeChEntertainmentRequest) SetServiceCode(v string) *GetWeChEntertainmentRequest {
	s.ServiceCode = &v
	return s
}

func (s *GetWeChEntertainmentRequest) SetSize(v string) *GetWeChEntertainmentRequest {
	s.Size = &v
	return s
}

func (s *GetWeChEntertainmentRequest) SetText(v string) *GetWeChEntertainmentRequest {
	s.Text = &v
	return s
}

func (s *GetWeChEntertainmentRequest) SetTokenizerId(v string) *GetWeChEntertainmentRequest {
	s.TokenizerId = &v
	return s
}

func (s *GetWeChEntertainmentRequest) SetType(v string) *GetWeChEntertainmentRequest {
	s.Type = &v
	return s
}

type GetWeChEntertainmentResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetWeChEntertainmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWeChEntertainmentResponseBody) GoString() string {
	return s.String()
}

func (s *GetWeChEntertainmentResponseBody) SetData(v string) *GetWeChEntertainmentResponseBody {
	s.Data = &v
	return s
}

func (s *GetWeChEntertainmentResponseBody) SetRequestId(v string) *GetWeChEntertainmentResponseBody {
	s.RequestId = &v
	return s
}

type GetWeChEntertainmentResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWeChEntertainmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWeChEntertainmentResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWeChEntertainmentResponse) GoString() string {
	return s.String()
}

func (s *GetWeChEntertainmentResponse) SetHeaders(v map[string]*string) *GetWeChEntertainmentResponse {
	s.Headers = v
	return s
}

func (s *GetWeChEntertainmentResponse) SetStatusCode(v int32) *GetWeChEntertainmentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWeChEntertainmentResponse) SetBody(v *GetWeChEntertainmentResponseBody) *GetWeChEntertainmentResponse {
	s.Body = v
	return s
}

type GetWeChGeneralRequest struct {
	Operation   *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	Size        *string `json:"Size,omitempty" xml:"Size,omitempty"`
	Text        *string `json:"Text,omitempty" xml:"Text,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetWeChGeneralRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWeChGeneralRequest) GoString() string {
	return s.String()
}

func (s *GetWeChGeneralRequest) SetOperation(v string) *GetWeChGeneralRequest {
	s.Operation = &v
	return s
}

func (s *GetWeChGeneralRequest) SetServiceCode(v string) *GetWeChGeneralRequest {
	s.ServiceCode = &v
	return s
}

func (s *GetWeChGeneralRequest) SetSize(v string) *GetWeChGeneralRequest {
	s.Size = &v
	return s
}

func (s *GetWeChGeneralRequest) SetText(v string) *GetWeChGeneralRequest {
	s.Text = &v
	return s
}

func (s *GetWeChGeneralRequest) SetType(v string) *GetWeChGeneralRequest {
	s.Type = &v
	return s
}

type GetWeChGeneralResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetWeChGeneralResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWeChGeneralResponseBody) GoString() string {
	return s.String()
}

func (s *GetWeChGeneralResponseBody) SetData(v string) *GetWeChGeneralResponseBody {
	s.Data = &v
	return s
}

func (s *GetWeChGeneralResponseBody) SetRequestId(v string) *GetWeChGeneralResponseBody {
	s.RequestId = &v
	return s
}

type GetWeChGeneralResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWeChGeneralResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWeChGeneralResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWeChGeneralResponse) GoString() string {
	return s.String()
}

func (s *GetWeChGeneralResponse) SetHeaders(v map[string]*string) *GetWeChGeneralResponse {
	s.Headers = v
	return s
}

func (s *GetWeChGeneralResponse) SetStatusCode(v int32) *GetWeChGeneralResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWeChGeneralResponse) SetBody(v *GetWeChGeneralResponseBody) *GetWeChGeneralResponse {
	s.Body = v
	return s
}

type GetWeChSearchRequest struct {
	Operation   *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	Size        *string `json:"Size,omitempty" xml:"Size,omitempty"`
	Text        *string `json:"Text,omitempty" xml:"Text,omitempty"`
	TokenizerId *string `json:"TokenizerId,omitempty" xml:"TokenizerId,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetWeChSearchRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWeChSearchRequest) GoString() string {
	return s.String()
}

func (s *GetWeChSearchRequest) SetOperation(v string) *GetWeChSearchRequest {
	s.Operation = &v
	return s
}

func (s *GetWeChSearchRequest) SetServiceCode(v string) *GetWeChSearchRequest {
	s.ServiceCode = &v
	return s
}

func (s *GetWeChSearchRequest) SetSize(v string) *GetWeChSearchRequest {
	s.Size = &v
	return s
}

func (s *GetWeChSearchRequest) SetText(v string) *GetWeChSearchRequest {
	s.Text = &v
	return s
}

func (s *GetWeChSearchRequest) SetTokenizerId(v string) *GetWeChSearchRequest {
	s.TokenizerId = &v
	return s
}

func (s *GetWeChSearchRequest) SetType(v string) *GetWeChSearchRequest {
	s.Type = &v
	return s
}

type GetWeChSearchResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetWeChSearchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWeChSearchResponseBody) GoString() string {
	return s.String()
}

func (s *GetWeChSearchResponseBody) SetData(v string) *GetWeChSearchResponseBody {
	s.Data = &v
	return s
}

func (s *GetWeChSearchResponseBody) SetRequestId(v string) *GetWeChSearchResponseBody {
	s.RequestId = &v
	return s
}

type GetWeChSearchResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWeChSearchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWeChSearchResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWeChSearchResponse) GoString() string {
	return s.String()
}

func (s *GetWeChSearchResponse) SetHeaders(v map[string]*string) *GetWeChSearchResponse {
	s.Headers = v
	return s
}

func (s *GetWeChSearchResponse) SetStatusCode(v int32) *GetWeChSearchResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWeChSearchResponse) SetBody(v *GetWeChSearchResponseBody) *GetWeChSearchResponse {
	s.Body = v
	return s
}

type GetWsChGeneralRequest struct {
	OutType     *string `json:"OutType,omitempty" xml:"OutType,omitempty"`
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	Text        *string `json:"Text,omitempty" xml:"Text,omitempty"`
	TokenizerId *string `json:"TokenizerId,omitempty" xml:"TokenizerId,omitempty"`
}

func (s GetWsChGeneralRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWsChGeneralRequest) GoString() string {
	return s.String()
}

func (s *GetWsChGeneralRequest) SetOutType(v string) *GetWsChGeneralRequest {
	s.OutType = &v
	return s
}

func (s *GetWsChGeneralRequest) SetServiceCode(v string) *GetWsChGeneralRequest {
	s.ServiceCode = &v
	return s
}

func (s *GetWsChGeneralRequest) SetText(v string) *GetWsChGeneralRequest {
	s.Text = &v
	return s
}

func (s *GetWsChGeneralRequest) SetTokenizerId(v string) *GetWsChGeneralRequest {
	s.TokenizerId = &v
	return s
}

type GetWsChGeneralResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetWsChGeneralResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWsChGeneralResponseBody) GoString() string {
	return s.String()
}

func (s *GetWsChGeneralResponseBody) SetData(v string) *GetWsChGeneralResponseBody {
	s.Data = &v
	return s
}

func (s *GetWsChGeneralResponseBody) SetRequestId(v string) *GetWsChGeneralResponseBody {
	s.RequestId = &v
	return s
}

type GetWsChGeneralResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWsChGeneralResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWsChGeneralResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWsChGeneralResponse) GoString() string {
	return s.String()
}

func (s *GetWsChGeneralResponse) SetHeaders(v map[string]*string) *GetWsChGeneralResponse {
	s.Headers = v
	return s
}

func (s *GetWsChGeneralResponse) SetStatusCode(v int32) *GetWsChGeneralResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWsChGeneralResponse) SetBody(v *GetWsChGeneralResponseBody) *GetWsChGeneralResponse {
	s.Body = v
	return s
}

type GetWsCustomizedChEcomCommentRequest struct {
	OutType     *string `json:"OutType,omitempty" xml:"OutType,omitempty"`
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	Text        *string `json:"Text,omitempty" xml:"Text,omitempty"`
	TokenizerId *string `json:"TokenizerId,omitempty" xml:"TokenizerId,omitempty"`
}

func (s GetWsCustomizedChEcomCommentRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWsCustomizedChEcomCommentRequest) GoString() string {
	return s.String()
}

func (s *GetWsCustomizedChEcomCommentRequest) SetOutType(v string) *GetWsCustomizedChEcomCommentRequest {
	s.OutType = &v
	return s
}

func (s *GetWsCustomizedChEcomCommentRequest) SetServiceCode(v string) *GetWsCustomizedChEcomCommentRequest {
	s.ServiceCode = &v
	return s
}

func (s *GetWsCustomizedChEcomCommentRequest) SetText(v string) *GetWsCustomizedChEcomCommentRequest {
	s.Text = &v
	return s
}

func (s *GetWsCustomizedChEcomCommentRequest) SetTokenizerId(v string) *GetWsCustomizedChEcomCommentRequest {
	s.TokenizerId = &v
	return s
}

type GetWsCustomizedChEcomCommentResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetWsCustomizedChEcomCommentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWsCustomizedChEcomCommentResponseBody) GoString() string {
	return s.String()
}

func (s *GetWsCustomizedChEcomCommentResponseBody) SetData(v string) *GetWsCustomizedChEcomCommentResponseBody {
	s.Data = &v
	return s
}

func (s *GetWsCustomizedChEcomCommentResponseBody) SetRequestId(v string) *GetWsCustomizedChEcomCommentResponseBody {
	s.RequestId = &v
	return s
}

type GetWsCustomizedChEcomCommentResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWsCustomizedChEcomCommentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWsCustomizedChEcomCommentResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWsCustomizedChEcomCommentResponse) GoString() string {
	return s.String()
}

func (s *GetWsCustomizedChEcomCommentResponse) SetHeaders(v map[string]*string) *GetWsCustomizedChEcomCommentResponse {
	s.Headers = v
	return s
}

func (s *GetWsCustomizedChEcomCommentResponse) SetStatusCode(v int32) *GetWsCustomizedChEcomCommentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWsCustomizedChEcomCommentResponse) SetBody(v *GetWsCustomizedChEcomCommentResponseBody) *GetWsCustomizedChEcomCommentResponse {
	s.Body = v
	return s
}

type GetWsCustomizedChEcomContentRequest struct {
	OutType     *string `json:"OutType,omitempty" xml:"OutType,omitempty"`
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	Text        *string `json:"Text,omitempty" xml:"Text,omitempty"`
	TokenizerId *string `json:"TokenizerId,omitempty" xml:"TokenizerId,omitempty"`
}

func (s GetWsCustomizedChEcomContentRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWsCustomizedChEcomContentRequest) GoString() string {
	return s.String()
}

func (s *GetWsCustomizedChEcomContentRequest) SetOutType(v string) *GetWsCustomizedChEcomContentRequest {
	s.OutType = &v
	return s
}

func (s *GetWsCustomizedChEcomContentRequest) SetServiceCode(v string) *GetWsCustomizedChEcomContentRequest {
	s.ServiceCode = &v
	return s
}

func (s *GetWsCustomizedChEcomContentRequest) SetText(v string) *GetWsCustomizedChEcomContentRequest {
	s.Text = &v
	return s
}

func (s *GetWsCustomizedChEcomContentRequest) SetTokenizerId(v string) *GetWsCustomizedChEcomContentRequest {
	s.TokenizerId = &v
	return s
}

type GetWsCustomizedChEcomContentResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetWsCustomizedChEcomContentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWsCustomizedChEcomContentResponseBody) GoString() string {
	return s.String()
}

func (s *GetWsCustomizedChEcomContentResponseBody) SetData(v string) *GetWsCustomizedChEcomContentResponseBody {
	s.Data = &v
	return s
}

func (s *GetWsCustomizedChEcomContentResponseBody) SetRequestId(v string) *GetWsCustomizedChEcomContentResponseBody {
	s.RequestId = &v
	return s
}

type GetWsCustomizedChEcomContentResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWsCustomizedChEcomContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWsCustomizedChEcomContentResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWsCustomizedChEcomContentResponse) GoString() string {
	return s.String()
}

func (s *GetWsCustomizedChEcomContentResponse) SetHeaders(v map[string]*string) *GetWsCustomizedChEcomContentResponse {
	s.Headers = v
	return s
}

func (s *GetWsCustomizedChEcomContentResponse) SetStatusCode(v int32) *GetWsCustomizedChEcomContentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWsCustomizedChEcomContentResponse) SetBody(v *GetWsCustomizedChEcomContentResponseBody) *GetWsCustomizedChEcomContentResponse {
	s.Body = v
	return s
}

type GetWsCustomizedChEcomTitleRequest struct {
	OutType     *string `json:"OutType,omitempty" xml:"OutType,omitempty"`
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	Text        *string `json:"Text,omitempty" xml:"Text,omitempty"`
	TokenizerId *string `json:"TokenizerId,omitempty" xml:"TokenizerId,omitempty"`
}

func (s GetWsCustomizedChEcomTitleRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWsCustomizedChEcomTitleRequest) GoString() string {
	return s.String()
}

func (s *GetWsCustomizedChEcomTitleRequest) SetOutType(v string) *GetWsCustomizedChEcomTitleRequest {
	s.OutType = &v
	return s
}

func (s *GetWsCustomizedChEcomTitleRequest) SetServiceCode(v string) *GetWsCustomizedChEcomTitleRequest {
	s.ServiceCode = &v
	return s
}

func (s *GetWsCustomizedChEcomTitleRequest) SetText(v string) *GetWsCustomizedChEcomTitleRequest {
	s.Text = &v
	return s
}

func (s *GetWsCustomizedChEcomTitleRequest) SetTokenizerId(v string) *GetWsCustomizedChEcomTitleRequest {
	s.TokenizerId = &v
	return s
}

type GetWsCustomizedChEcomTitleResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetWsCustomizedChEcomTitleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWsCustomizedChEcomTitleResponseBody) GoString() string {
	return s.String()
}

func (s *GetWsCustomizedChEcomTitleResponseBody) SetData(v string) *GetWsCustomizedChEcomTitleResponseBody {
	s.Data = &v
	return s
}

func (s *GetWsCustomizedChEcomTitleResponseBody) SetRequestId(v string) *GetWsCustomizedChEcomTitleResponseBody {
	s.RequestId = &v
	return s
}

type GetWsCustomizedChEcomTitleResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWsCustomizedChEcomTitleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWsCustomizedChEcomTitleResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWsCustomizedChEcomTitleResponse) GoString() string {
	return s.String()
}

func (s *GetWsCustomizedChEcomTitleResponse) SetHeaders(v map[string]*string) *GetWsCustomizedChEcomTitleResponse {
	s.Headers = v
	return s
}

func (s *GetWsCustomizedChEcomTitleResponse) SetStatusCode(v int32) *GetWsCustomizedChEcomTitleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWsCustomizedChEcomTitleResponse) SetBody(v *GetWsCustomizedChEcomTitleResponseBody) *GetWsCustomizedChEcomTitleResponse {
	s.Body = v
	return s
}

type GetWsCustomizedChEntertainmentRequest struct {
	OutType     *string `json:"OutType,omitempty" xml:"OutType,omitempty"`
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	Text        *string `json:"Text,omitempty" xml:"Text,omitempty"`
	TokenizerId *string `json:"TokenizerId,omitempty" xml:"TokenizerId,omitempty"`
}

func (s GetWsCustomizedChEntertainmentRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWsCustomizedChEntertainmentRequest) GoString() string {
	return s.String()
}

func (s *GetWsCustomizedChEntertainmentRequest) SetOutType(v string) *GetWsCustomizedChEntertainmentRequest {
	s.OutType = &v
	return s
}

func (s *GetWsCustomizedChEntertainmentRequest) SetServiceCode(v string) *GetWsCustomizedChEntertainmentRequest {
	s.ServiceCode = &v
	return s
}

func (s *GetWsCustomizedChEntertainmentRequest) SetText(v string) *GetWsCustomizedChEntertainmentRequest {
	s.Text = &v
	return s
}

func (s *GetWsCustomizedChEntertainmentRequest) SetTokenizerId(v string) *GetWsCustomizedChEntertainmentRequest {
	s.TokenizerId = &v
	return s
}

type GetWsCustomizedChEntertainmentResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetWsCustomizedChEntertainmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWsCustomizedChEntertainmentResponseBody) GoString() string {
	return s.String()
}

func (s *GetWsCustomizedChEntertainmentResponseBody) SetData(v string) *GetWsCustomizedChEntertainmentResponseBody {
	s.Data = &v
	return s
}

func (s *GetWsCustomizedChEntertainmentResponseBody) SetRequestId(v string) *GetWsCustomizedChEntertainmentResponseBody {
	s.RequestId = &v
	return s
}

type GetWsCustomizedChEntertainmentResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWsCustomizedChEntertainmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWsCustomizedChEntertainmentResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWsCustomizedChEntertainmentResponse) GoString() string {
	return s.String()
}

func (s *GetWsCustomizedChEntertainmentResponse) SetHeaders(v map[string]*string) *GetWsCustomizedChEntertainmentResponse {
	s.Headers = v
	return s
}

func (s *GetWsCustomizedChEntertainmentResponse) SetStatusCode(v int32) *GetWsCustomizedChEntertainmentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWsCustomizedChEntertainmentResponse) SetBody(v *GetWsCustomizedChEntertainmentResponseBody) *GetWsCustomizedChEntertainmentResponse {
	s.Body = v
	return s
}

type GetWsCustomizedChGeneralRequest struct {
	OutType     *string `json:"OutType,omitempty" xml:"OutType,omitempty"`
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	Text        *string `json:"Text,omitempty" xml:"Text,omitempty"`
	TokenizerId *string `json:"TokenizerId,omitempty" xml:"TokenizerId,omitempty"`
}

func (s GetWsCustomizedChGeneralRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWsCustomizedChGeneralRequest) GoString() string {
	return s.String()
}

func (s *GetWsCustomizedChGeneralRequest) SetOutType(v string) *GetWsCustomizedChGeneralRequest {
	s.OutType = &v
	return s
}

func (s *GetWsCustomizedChGeneralRequest) SetServiceCode(v string) *GetWsCustomizedChGeneralRequest {
	s.ServiceCode = &v
	return s
}

func (s *GetWsCustomizedChGeneralRequest) SetText(v string) *GetWsCustomizedChGeneralRequest {
	s.Text = &v
	return s
}

func (s *GetWsCustomizedChGeneralRequest) SetTokenizerId(v string) *GetWsCustomizedChGeneralRequest {
	s.TokenizerId = &v
	return s
}

type GetWsCustomizedChGeneralResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetWsCustomizedChGeneralResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWsCustomizedChGeneralResponseBody) GoString() string {
	return s.String()
}

func (s *GetWsCustomizedChGeneralResponseBody) SetData(v string) *GetWsCustomizedChGeneralResponseBody {
	s.Data = &v
	return s
}

func (s *GetWsCustomizedChGeneralResponseBody) SetRequestId(v string) *GetWsCustomizedChGeneralResponseBody {
	s.RequestId = &v
	return s
}

type GetWsCustomizedChGeneralResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWsCustomizedChGeneralResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWsCustomizedChGeneralResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWsCustomizedChGeneralResponse) GoString() string {
	return s.String()
}

func (s *GetWsCustomizedChGeneralResponse) SetHeaders(v map[string]*string) *GetWsCustomizedChGeneralResponse {
	s.Headers = v
	return s
}

func (s *GetWsCustomizedChGeneralResponse) SetStatusCode(v int32) *GetWsCustomizedChGeneralResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWsCustomizedChGeneralResponse) SetBody(v *GetWsCustomizedChGeneralResponseBody) *GetWsCustomizedChGeneralResponse {
	s.Body = v
	return s
}

type GetWsCustomizedChO2ORequest struct {
	OutType     *string `json:"OutType,omitempty" xml:"OutType,omitempty"`
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	Text        *string `json:"Text,omitempty" xml:"Text,omitempty"`
	TokenizerId *string `json:"TokenizerId,omitempty" xml:"TokenizerId,omitempty"`
}

func (s GetWsCustomizedChO2ORequest) String() string {
	return tea.Prettify(s)
}

func (s GetWsCustomizedChO2ORequest) GoString() string {
	return s.String()
}

func (s *GetWsCustomizedChO2ORequest) SetOutType(v string) *GetWsCustomizedChO2ORequest {
	s.OutType = &v
	return s
}

func (s *GetWsCustomizedChO2ORequest) SetServiceCode(v string) *GetWsCustomizedChO2ORequest {
	s.ServiceCode = &v
	return s
}

func (s *GetWsCustomizedChO2ORequest) SetText(v string) *GetWsCustomizedChO2ORequest {
	s.Text = &v
	return s
}

func (s *GetWsCustomizedChO2ORequest) SetTokenizerId(v string) *GetWsCustomizedChO2ORequest {
	s.TokenizerId = &v
	return s
}

type GetWsCustomizedChO2OResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetWsCustomizedChO2OResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWsCustomizedChO2OResponseBody) GoString() string {
	return s.String()
}

func (s *GetWsCustomizedChO2OResponseBody) SetData(v string) *GetWsCustomizedChO2OResponseBody {
	s.Data = &v
	return s
}

func (s *GetWsCustomizedChO2OResponseBody) SetRequestId(v string) *GetWsCustomizedChO2OResponseBody {
	s.RequestId = &v
	return s
}

type GetWsCustomizedChO2OResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWsCustomizedChO2OResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWsCustomizedChO2OResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWsCustomizedChO2OResponse) GoString() string {
	return s.String()
}

func (s *GetWsCustomizedChO2OResponse) SetHeaders(v map[string]*string) *GetWsCustomizedChO2OResponse {
	s.Headers = v
	return s
}

func (s *GetWsCustomizedChO2OResponse) SetStatusCode(v int32) *GetWsCustomizedChO2OResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWsCustomizedChO2OResponse) SetBody(v *GetWsCustomizedChO2OResponseBody) *GetWsCustomizedChO2OResponse {
	s.Body = v
	return s
}

type GetWsCustomizedSeaEcomRequest struct {
	Language    *string `json:"Language,omitempty" xml:"Language,omitempty"`
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	Text        *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s GetWsCustomizedSeaEcomRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWsCustomizedSeaEcomRequest) GoString() string {
	return s.String()
}

func (s *GetWsCustomizedSeaEcomRequest) SetLanguage(v string) *GetWsCustomizedSeaEcomRequest {
	s.Language = &v
	return s
}

func (s *GetWsCustomizedSeaEcomRequest) SetServiceCode(v string) *GetWsCustomizedSeaEcomRequest {
	s.ServiceCode = &v
	return s
}

func (s *GetWsCustomizedSeaEcomRequest) SetText(v string) *GetWsCustomizedSeaEcomRequest {
	s.Text = &v
	return s
}

type GetWsCustomizedSeaEcomResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetWsCustomizedSeaEcomResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWsCustomizedSeaEcomResponseBody) GoString() string {
	return s.String()
}

func (s *GetWsCustomizedSeaEcomResponseBody) SetData(v string) *GetWsCustomizedSeaEcomResponseBody {
	s.Data = &v
	return s
}

func (s *GetWsCustomizedSeaEcomResponseBody) SetRequestId(v string) *GetWsCustomizedSeaEcomResponseBody {
	s.RequestId = &v
	return s
}

type GetWsCustomizedSeaEcomResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWsCustomizedSeaEcomResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWsCustomizedSeaEcomResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWsCustomizedSeaEcomResponse) GoString() string {
	return s.String()
}

func (s *GetWsCustomizedSeaEcomResponse) SetHeaders(v map[string]*string) *GetWsCustomizedSeaEcomResponse {
	s.Headers = v
	return s
}

func (s *GetWsCustomizedSeaEcomResponse) SetStatusCode(v int32) *GetWsCustomizedSeaEcomResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWsCustomizedSeaEcomResponse) SetBody(v *GetWsCustomizedSeaEcomResponseBody) *GetWsCustomizedSeaEcomResponse {
	s.Body = v
	return s
}

type GetWsCustomizedSeaGeneralRequest struct {
	Language    *string `json:"Language,omitempty" xml:"Language,omitempty"`
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	Text        *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s GetWsCustomizedSeaGeneralRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWsCustomizedSeaGeneralRequest) GoString() string {
	return s.String()
}

func (s *GetWsCustomizedSeaGeneralRequest) SetLanguage(v string) *GetWsCustomizedSeaGeneralRequest {
	s.Language = &v
	return s
}

func (s *GetWsCustomizedSeaGeneralRequest) SetServiceCode(v string) *GetWsCustomizedSeaGeneralRequest {
	s.ServiceCode = &v
	return s
}

func (s *GetWsCustomizedSeaGeneralRequest) SetText(v string) *GetWsCustomizedSeaGeneralRequest {
	s.Text = &v
	return s
}

type GetWsCustomizedSeaGeneralResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetWsCustomizedSeaGeneralResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWsCustomizedSeaGeneralResponseBody) GoString() string {
	return s.String()
}

func (s *GetWsCustomizedSeaGeneralResponseBody) SetData(v string) *GetWsCustomizedSeaGeneralResponseBody {
	s.Data = &v
	return s
}

func (s *GetWsCustomizedSeaGeneralResponseBody) SetRequestId(v string) *GetWsCustomizedSeaGeneralResponseBody {
	s.RequestId = &v
	return s
}

type GetWsCustomizedSeaGeneralResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWsCustomizedSeaGeneralResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWsCustomizedSeaGeneralResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWsCustomizedSeaGeneralResponse) GoString() string {
	return s.String()
}

func (s *GetWsCustomizedSeaGeneralResponse) SetHeaders(v map[string]*string) *GetWsCustomizedSeaGeneralResponse {
	s.Headers = v
	return s
}

func (s *GetWsCustomizedSeaGeneralResponse) SetStatusCode(v int32) *GetWsCustomizedSeaGeneralResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWsCustomizedSeaGeneralResponse) SetBody(v *GetWsCustomizedSeaGeneralResponseBody) *GetWsCustomizedSeaGeneralResponse {
	s.Body = v
	return s
}

type ImportServiceDataRequest struct {
	Partition []map[string]*string `json:"Partition,omitempty" xml:"Partition,omitempty" type:"Repeated"`
	ServiceId *int64               `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	SubPath   *string              `json:"SubPath,omitempty" xml:"SubPath,omitempty"`
	Url       *string              `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ImportServiceDataRequest) String() string {
	return tea.Prettify(s)
}

func (s ImportServiceDataRequest) GoString() string {
	return s.String()
}

func (s *ImportServiceDataRequest) SetPartition(v []map[string]*string) *ImportServiceDataRequest {
	s.Partition = v
	return s
}

func (s *ImportServiceDataRequest) SetServiceId(v int64) *ImportServiceDataRequest {
	s.ServiceId = &v
	return s
}

func (s *ImportServiceDataRequest) SetSubPath(v string) *ImportServiceDataRequest {
	s.SubPath = &v
	return s
}

func (s *ImportServiceDataRequest) SetUrl(v string) *ImportServiceDataRequest {
	s.Url = &v
	return s
}

type ImportServiceDataShrinkRequest struct {
	PartitionShrink *string `json:"Partition,omitempty" xml:"Partition,omitempty"`
	ServiceId       *int64  `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	SubPath         *string `json:"SubPath,omitempty" xml:"SubPath,omitempty"`
	Url             *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ImportServiceDataShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ImportServiceDataShrinkRequest) GoString() string {
	return s.String()
}

func (s *ImportServiceDataShrinkRequest) SetPartitionShrink(v string) *ImportServiceDataShrinkRequest {
	s.PartitionShrink = &v
	return s
}

func (s *ImportServiceDataShrinkRequest) SetServiceId(v int64) *ImportServiceDataShrinkRequest {
	s.ServiceId = &v
	return s
}

func (s *ImportServiceDataShrinkRequest) SetSubPath(v string) *ImportServiceDataShrinkRequest {
	s.SubPath = &v
	return s
}

func (s *ImportServiceDataShrinkRequest) SetUrl(v string) *ImportServiceDataShrinkRequest {
	s.Url = &v
	return s
}

type ImportServiceDataResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	Msg       *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ImportServiceDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ImportServiceDataResponseBody) GoString() string {
	return s.String()
}

func (s *ImportServiceDataResponseBody) SetCode(v int32) *ImportServiceDataResponseBody {
	s.Code = &v
	return s
}

func (s *ImportServiceDataResponseBody) SetData(v int64) *ImportServiceDataResponseBody {
	s.Data = &v
	return s
}

func (s *ImportServiceDataResponseBody) SetMsg(v string) *ImportServiceDataResponseBody {
	s.Msg = &v
	return s
}

func (s *ImportServiceDataResponseBody) SetRequestId(v string) *ImportServiceDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportServiceDataResponseBody) SetSuccess(v bool) *ImportServiceDataResponseBody {
	s.Success = &v
	return s
}

type ImportServiceDataResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImportServiceDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImportServiceDataResponse) String() string {
	return tea.Prettify(s)
}

func (s ImportServiceDataResponse) GoString() string {
	return s.String()
}

func (s *ImportServiceDataResponse) SetHeaders(v map[string]*string) *ImportServiceDataResponse {
	s.Headers = v
	return s
}

func (s *ImportServiceDataResponse) SetStatusCode(v int32) *ImportServiceDataResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportServiceDataResponse) SetBody(v *ImportServiceDataResponseBody) *ImportServiceDataResponse {
	s.Body = v
	return s
}

type ImportServiceDataV2Request struct {
	DataType  *string                                `json:"DataType,omitempty" xml:"DataType,omitempty"`
	Documents []*ImportServiceDataV2RequestDocuments `json:"Documents,omitempty" xml:"Documents,omitempty" type:"Repeated"`
	ServiceId *int64                                 `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s ImportServiceDataV2Request) String() string {
	return tea.Prettify(s)
}

func (s ImportServiceDataV2Request) GoString() string {
	return s.String()
}

func (s *ImportServiceDataV2Request) SetDataType(v string) *ImportServiceDataV2Request {
	s.DataType = &v
	return s
}

func (s *ImportServiceDataV2Request) SetDocuments(v []*ImportServiceDataV2RequestDocuments) *ImportServiceDataV2Request {
	s.Documents = v
	return s
}

func (s *ImportServiceDataV2Request) SetServiceId(v int64) *ImportServiceDataV2Request {
	s.ServiceId = &v
	return s
}

type ImportServiceDataV2RequestDocuments struct {
	BizParams     map[string]*string `json:"BizParams,omitempty" xml:"BizParams,omitempty"`
	DocId         *string            `json:"DocId,omitempty" xml:"DocId,omitempty"`
	FileExtension *string            `json:"FileExtension,omitempty" xml:"FileExtension,omitempty"`
	FileName      *string            `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FilePath      *string            `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	Version       *string            `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ImportServiceDataV2RequestDocuments) String() string {
	return tea.Prettify(s)
}

func (s ImportServiceDataV2RequestDocuments) GoString() string {
	return s.String()
}

func (s *ImportServiceDataV2RequestDocuments) SetBizParams(v map[string]*string) *ImportServiceDataV2RequestDocuments {
	s.BizParams = v
	return s
}

func (s *ImportServiceDataV2RequestDocuments) SetDocId(v string) *ImportServiceDataV2RequestDocuments {
	s.DocId = &v
	return s
}

func (s *ImportServiceDataV2RequestDocuments) SetFileExtension(v string) *ImportServiceDataV2RequestDocuments {
	s.FileExtension = &v
	return s
}

func (s *ImportServiceDataV2RequestDocuments) SetFileName(v string) *ImportServiceDataV2RequestDocuments {
	s.FileName = &v
	return s
}

func (s *ImportServiceDataV2RequestDocuments) SetFilePath(v string) *ImportServiceDataV2RequestDocuments {
	s.FilePath = &v
	return s
}

func (s *ImportServiceDataV2RequestDocuments) SetVersion(v string) *ImportServiceDataV2RequestDocuments {
	s.Version = &v
	return s
}

type ImportServiceDataV2ShrinkRequest struct {
	DataType        *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	DocumentsShrink *string `json:"Documents,omitempty" xml:"Documents,omitempty"`
	ServiceId       *int64  `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s ImportServiceDataV2ShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ImportServiceDataV2ShrinkRequest) GoString() string {
	return s.String()
}

func (s *ImportServiceDataV2ShrinkRequest) SetDataType(v string) *ImportServiceDataV2ShrinkRequest {
	s.DataType = &v
	return s
}

func (s *ImportServiceDataV2ShrinkRequest) SetDocumentsShrink(v string) *ImportServiceDataV2ShrinkRequest {
	s.DocumentsShrink = &v
	return s
}

func (s *ImportServiceDataV2ShrinkRequest) SetServiceId(v int64) *ImportServiceDataV2ShrinkRequest {
	s.ServiceId = &v
	return s
}

type ImportServiceDataV2ResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	Msg       *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ImportServiceDataV2ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ImportServiceDataV2ResponseBody) GoString() string {
	return s.String()
}

func (s *ImportServiceDataV2ResponseBody) SetCode(v int32) *ImportServiceDataV2ResponseBody {
	s.Code = &v
	return s
}

func (s *ImportServiceDataV2ResponseBody) SetData(v int64) *ImportServiceDataV2ResponseBody {
	s.Data = &v
	return s
}

func (s *ImportServiceDataV2ResponseBody) SetMsg(v string) *ImportServiceDataV2ResponseBody {
	s.Msg = &v
	return s
}

func (s *ImportServiceDataV2ResponseBody) SetRequestId(v string) *ImportServiceDataV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportServiceDataV2ResponseBody) SetSuccess(v bool) *ImportServiceDataV2ResponseBody {
	s.Success = &v
	return s
}

type ImportServiceDataV2Response struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImportServiceDataV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImportServiceDataV2Response) String() string {
	return tea.Prettify(s)
}

func (s ImportServiceDataV2Response) GoString() string {
	return s.String()
}

func (s *ImportServiceDataV2Response) SetHeaders(v map[string]*string) *ImportServiceDataV2Response {
	s.Headers = v
	return s
}

func (s *ImportServiceDataV2Response) SetStatusCode(v int32) *ImportServiceDataV2Response {
	s.StatusCode = &v
	return s
}

func (s *ImportServiceDataV2Response) SetBody(v *ImportServiceDataV2ResponseBody) *ImportServiceDataV2Response {
	s.Body = v
	return s
}

type InsertCustomRequest struct {
	ApiId          *int32  `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	CustomFileName *string `json:"CustomFileName,omitempty" xml:"CustomFileName,omitempty"`
	CustomUrl      *string `json:"CustomUrl,omitempty" xml:"CustomUrl,omitempty"`
	RegFileName    *string `json:"RegFileName,omitempty" xml:"RegFileName,omitempty"`
	RegUrl         *string `json:"RegUrl,omitempty" xml:"RegUrl,omitempty"`
	ServiceCode    *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
}

func (s InsertCustomRequest) String() string {
	return tea.Prettify(s)
}

func (s InsertCustomRequest) GoString() string {
	return s.String()
}

func (s *InsertCustomRequest) SetApiId(v int32) *InsertCustomRequest {
	s.ApiId = &v
	return s
}

func (s *InsertCustomRequest) SetCustomFileName(v string) *InsertCustomRequest {
	s.CustomFileName = &v
	return s
}

func (s *InsertCustomRequest) SetCustomUrl(v string) *InsertCustomRequest {
	s.CustomUrl = &v
	return s
}

func (s *InsertCustomRequest) SetRegFileName(v string) *InsertCustomRequest {
	s.RegFileName = &v
	return s
}

func (s *InsertCustomRequest) SetRegUrl(v string) *InsertCustomRequest {
	s.RegUrl = &v
	return s
}

func (s *InsertCustomRequest) SetServiceCode(v string) *InsertCustomRequest {
	s.ServiceCode = &v
	return s
}

type InsertCustomResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InsertCustomResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InsertCustomResponseBody) GoString() string {
	return s.String()
}

func (s *InsertCustomResponseBody) SetData(v string) *InsertCustomResponseBody {
	s.Data = &v
	return s
}

func (s *InsertCustomResponseBody) SetRequestId(v string) *InsertCustomResponseBody {
	s.RequestId = &v
	return s
}

type InsertCustomResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InsertCustomResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InsertCustomResponse) String() string {
	return tea.Prettify(s)
}

func (s InsertCustomResponse) GoString() string {
	return s.String()
}

func (s *InsertCustomResponse) SetHeaders(v map[string]*string) *InsertCustomResponse {
	s.Headers = v
	return s
}

func (s *InsertCustomResponse) SetStatusCode(v int32) *InsertCustomResponse {
	s.StatusCode = &v
	return s
}

func (s *InsertCustomResponse) SetBody(v *InsertCustomResponseBody) *InsertCustomResponse {
	s.Body = v
	return s
}

type OpenAlinlpServiceResponseBody struct {
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenAlinlpServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpenAlinlpServiceResponseBody) GoString() string {
	return s.String()
}

func (s *OpenAlinlpServiceResponseBody) SetOrderId(v string) *OpenAlinlpServiceResponseBody {
	s.OrderId = &v
	return s
}

func (s *OpenAlinlpServiceResponseBody) SetRequestId(v string) *OpenAlinlpServiceResponseBody {
	s.RequestId = &v
	return s
}

type OpenAlinlpServiceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenAlinlpServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenAlinlpServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenAlinlpServiceResponse) GoString() string {
	return s.String()
}

func (s *OpenAlinlpServiceResponse) SetHeaders(v map[string]*string) *OpenAlinlpServiceResponse {
	s.Headers = v
	return s
}

func (s *OpenAlinlpServiceResponse) SetStatusCode(v int32) *OpenAlinlpServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenAlinlpServiceResponse) SetBody(v *OpenAlinlpServiceResponseBody) *OpenAlinlpServiceResponse {
	s.Body = v
	return s
}

type PostISConvRewriterRequest struct {
	Algorithm  *string                `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	Debug      *bool                  `json:"Debug,omitempty" xml:"Debug,omitempty"`
	Input      map[string]interface{} `json:"Input,omitempty" xml:"Input,omitempty"`
	Parameters map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	Version    *string                `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s PostISConvRewriterRequest) String() string {
	return tea.Prettify(s)
}

func (s PostISConvRewriterRequest) GoString() string {
	return s.String()
}

func (s *PostISConvRewriterRequest) SetAlgorithm(v string) *PostISConvRewriterRequest {
	s.Algorithm = &v
	return s
}

func (s *PostISConvRewriterRequest) SetDebug(v bool) *PostISConvRewriterRequest {
	s.Debug = &v
	return s
}

func (s *PostISConvRewriterRequest) SetInput(v map[string]interface{}) *PostISConvRewriterRequest {
	s.Input = v
	return s
}

func (s *PostISConvRewriterRequest) SetParameters(v map[string]interface{}) *PostISConvRewriterRequest {
	s.Parameters = v
	return s
}

func (s *PostISConvRewriterRequest) SetVersion(v string) *PostISConvRewriterRequest {
	s.Version = &v
	return s
}

type PostISConvRewriterShrinkRequest struct {
	Algorithm        *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	Debug            *bool   `json:"Debug,omitempty" xml:"Debug,omitempty"`
	InputShrink      *string `json:"Input,omitempty" xml:"Input,omitempty"`
	ParametersShrink *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	Version          *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s PostISConvRewriterShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s PostISConvRewriterShrinkRequest) GoString() string {
	return s.String()
}

func (s *PostISConvRewriterShrinkRequest) SetAlgorithm(v string) *PostISConvRewriterShrinkRequest {
	s.Algorithm = &v
	return s
}

func (s *PostISConvRewriterShrinkRequest) SetDebug(v bool) *PostISConvRewriterShrinkRequest {
	s.Debug = &v
	return s
}

func (s *PostISConvRewriterShrinkRequest) SetInputShrink(v string) *PostISConvRewriterShrinkRequest {
	s.InputShrink = &v
	return s
}

func (s *PostISConvRewriterShrinkRequest) SetParametersShrink(v string) *PostISConvRewriterShrinkRequest {
	s.ParametersShrink = &v
	return s
}

func (s *PostISConvRewriterShrinkRequest) SetVersion(v string) *PostISConvRewriterShrinkRequest {
	s.Version = &v
	return s
}

type PostISConvRewriterResponseBody struct {
	Data      map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	DebugInfo map[string]interface{} `json:"DebugInfo,omitempty" xml:"DebugInfo,omitempty"`
	Message   *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *int32                 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s PostISConvRewriterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PostISConvRewriterResponseBody) GoString() string {
	return s.String()
}

func (s *PostISConvRewriterResponseBody) SetData(v map[string]interface{}) *PostISConvRewriterResponseBody {
	s.Data = v
	return s
}

func (s *PostISConvRewriterResponseBody) SetDebugInfo(v map[string]interface{}) *PostISConvRewriterResponseBody {
	s.DebugInfo = v
	return s
}

func (s *PostISConvRewriterResponseBody) SetMessage(v string) *PostISConvRewriterResponseBody {
	s.Message = &v
	return s
}

func (s *PostISConvRewriterResponseBody) SetRequestId(v string) *PostISConvRewriterResponseBody {
	s.RequestId = &v
	return s
}

func (s *PostISConvRewriterResponseBody) SetStatus(v int32) *PostISConvRewriterResponseBody {
	s.Status = &v
	return s
}

type PostISConvRewriterResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PostISConvRewriterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PostISConvRewriterResponse) String() string {
	return tea.Prettify(s)
}

func (s PostISConvRewriterResponse) GoString() string {
	return s.String()
}

func (s *PostISConvRewriterResponse) SetHeaders(v map[string]*string) *PostISConvRewriterResponse {
	s.Headers = v
	return s
}

func (s *PostISConvRewriterResponse) SetStatusCode(v int32) *PostISConvRewriterResponse {
	s.StatusCode = &v
	return s
}

func (s *PostISConvRewriterResponse) SetBody(v *PostISConvRewriterResponseBody) *PostISConvRewriterResponse {
	s.Body = v
	return s
}

type PostISRetrieveRouterRequest struct {
	Algorithm  *string                `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	Debug      *bool                  `json:"Debug,omitempty" xml:"Debug,omitempty"`
	Input      map[string]interface{} `json:"Input,omitempty" xml:"Input,omitempty"`
	Parameters map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	Version    *string                `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s PostISRetrieveRouterRequest) String() string {
	return tea.Prettify(s)
}

func (s PostISRetrieveRouterRequest) GoString() string {
	return s.String()
}

func (s *PostISRetrieveRouterRequest) SetAlgorithm(v string) *PostISRetrieveRouterRequest {
	s.Algorithm = &v
	return s
}

func (s *PostISRetrieveRouterRequest) SetDebug(v bool) *PostISRetrieveRouterRequest {
	s.Debug = &v
	return s
}

func (s *PostISRetrieveRouterRequest) SetInput(v map[string]interface{}) *PostISRetrieveRouterRequest {
	s.Input = v
	return s
}

func (s *PostISRetrieveRouterRequest) SetParameters(v map[string]interface{}) *PostISRetrieveRouterRequest {
	s.Parameters = v
	return s
}

func (s *PostISRetrieveRouterRequest) SetVersion(v string) *PostISRetrieveRouterRequest {
	s.Version = &v
	return s
}

type PostISRetrieveRouterShrinkRequest struct {
	Algorithm        *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	Debug            *bool   `json:"Debug,omitempty" xml:"Debug,omitempty"`
	InputShrink      *string `json:"Input,omitempty" xml:"Input,omitempty"`
	ParametersShrink *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	Version          *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s PostISRetrieveRouterShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s PostISRetrieveRouterShrinkRequest) GoString() string {
	return s.String()
}

func (s *PostISRetrieveRouterShrinkRequest) SetAlgorithm(v string) *PostISRetrieveRouterShrinkRequest {
	s.Algorithm = &v
	return s
}

func (s *PostISRetrieveRouterShrinkRequest) SetDebug(v bool) *PostISRetrieveRouterShrinkRequest {
	s.Debug = &v
	return s
}

func (s *PostISRetrieveRouterShrinkRequest) SetInputShrink(v string) *PostISRetrieveRouterShrinkRequest {
	s.InputShrink = &v
	return s
}

func (s *PostISRetrieveRouterShrinkRequest) SetParametersShrink(v string) *PostISRetrieveRouterShrinkRequest {
	s.ParametersShrink = &v
	return s
}

func (s *PostISRetrieveRouterShrinkRequest) SetVersion(v string) *PostISRetrieveRouterShrinkRequest {
	s.Version = &v
	return s
}

type PostISRetrieveRouterResponseBody struct {
	Data      map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	DebugInfo map[string]interface{} `json:"DebugInfo,omitempty" xml:"DebugInfo,omitempty"`
	Message   *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *int32                 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s PostISRetrieveRouterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PostISRetrieveRouterResponseBody) GoString() string {
	return s.String()
}

func (s *PostISRetrieveRouterResponseBody) SetData(v map[string]interface{}) *PostISRetrieveRouterResponseBody {
	s.Data = v
	return s
}

func (s *PostISRetrieveRouterResponseBody) SetDebugInfo(v map[string]interface{}) *PostISRetrieveRouterResponseBody {
	s.DebugInfo = v
	return s
}

func (s *PostISRetrieveRouterResponseBody) SetMessage(v string) *PostISRetrieveRouterResponseBody {
	s.Message = &v
	return s
}

func (s *PostISRetrieveRouterResponseBody) SetRequestId(v string) *PostISRetrieveRouterResponseBody {
	s.RequestId = &v
	return s
}

func (s *PostISRetrieveRouterResponseBody) SetStatus(v int32) *PostISRetrieveRouterResponseBody {
	s.Status = &v
	return s
}

type PostISRetrieveRouterResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PostISRetrieveRouterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PostISRetrieveRouterResponse) String() string {
	return tea.Prettify(s)
}

func (s PostISRetrieveRouterResponse) GoString() string {
	return s.String()
}

func (s *PostISRetrieveRouterResponse) SetHeaders(v map[string]*string) *PostISRetrieveRouterResponse {
	s.Headers = v
	return s
}

func (s *PostISRetrieveRouterResponse) SetStatusCode(v int32) *PostISRetrieveRouterResponse {
	s.StatusCode = &v
	return s
}

func (s *PostISRetrieveRouterResponse) SetBody(v *PostISRetrieveRouterResponseBody) *PostISRetrieveRouterResponse {
	s.Body = v
	return s
}

type PostMSConvSearchTokenGeneratedResponseBody struct {
	Code           *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *string `json:"Data,omitempty" xml:"Data,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Msg            *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PostMSConvSearchTokenGeneratedResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PostMSConvSearchTokenGeneratedResponseBody) GoString() string {
	return s.String()
}

func (s *PostMSConvSearchTokenGeneratedResponseBody) SetCode(v int32) *PostMSConvSearchTokenGeneratedResponseBody {
	s.Code = &v
	return s
}

func (s *PostMSConvSearchTokenGeneratedResponseBody) SetData(v string) *PostMSConvSearchTokenGeneratedResponseBody {
	s.Data = &v
	return s
}

func (s *PostMSConvSearchTokenGeneratedResponseBody) SetHttpStatusCode(v int32) *PostMSConvSearchTokenGeneratedResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *PostMSConvSearchTokenGeneratedResponseBody) SetMsg(v string) *PostMSConvSearchTokenGeneratedResponseBody {
	s.Msg = &v
	return s
}

func (s *PostMSConvSearchTokenGeneratedResponseBody) SetRequestId(v string) *PostMSConvSearchTokenGeneratedResponseBody {
	s.RequestId = &v
	return s
}

func (s *PostMSConvSearchTokenGeneratedResponseBody) SetSuccess(v bool) *PostMSConvSearchTokenGeneratedResponseBody {
	s.Success = &v
	return s
}

type PostMSConvSearchTokenGeneratedResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PostMSConvSearchTokenGeneratedResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PostMSConvSearchTokenGeneratedResponse) String() string {
	return tea.Prettify(s)
}

func (s PostMSConvSearchTokenGeneratedResponse) GoString() string {
	return s.String()
}

func (s *PostMSConvSearchTokenGeneratedResponse) SetHeaders(v map[string]*string) *PostMSConvSearchTokenGeneratedResponse {
	s.Headers = v
	return s
}

func (s *PostMSConvSearchTokenGeneratedResponse) SetStatusCode(v int32) *PostMSConvSearchTokenGeneratedResponse {
	s.StatusCode = &v
	return s
}

func (s *PostMSConvSearchTokenGeneratedResponse) SetBody(v *PostMSConvSearchTokenGeneratedResponseBody) *PostMSConvSearchTokenGeneratedResponse {
	s.Body = v
	return s
}

type PostMSDataProcessingCountRequest struct {
	DataIds      []*string `json:"DataIds,omitempty" xml:"DataIds,omitempty" type:"Repeated"`
	DataImportId *int64    `json:"DataImportId,omitempty" xml:"DataImportId,omitempty"`
	ServiceId    *int64    `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s PostMSDataProcessingCountRequest) String() string {
	return tea.Prettify(s)
}

func (s PostMSDataProcessingCountRequest) GoString() string {
	return s.String()
}

func (s *PostMSDataProcessingCountRequest) SetDataIds(v []*string) *PostMSDataProcessingCountRequest {
	s.DataIds = v
	return s
}

func (s *PostMSDataProcessingCountRequest) SetDataImportId(v int64) *PostMSDataProcessingCountRequest {
	s.DataImportId = &v
	return s
}

func (s *PostMSDataProcessingCountRequest) SetServiceId(v int64) *PostMSDataProcessingCountRequest {
	s.ServiceId = &v
	return s
}

type PostMSDataProcessingCountShrinkRequest struct {
	DataIdsShrink *string `json:"DataIds,omitempty" xml:"DataIds,omitempty"`
	DataImportId  *int64  `json:"DataImportId,omitempty" xml:"DataImportId,omitempty"`
	ServiceId     *int64  `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s PostMSDataProcessingCountShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s PostMSDataProcessingCountShrinkRequest) GoString() string {
	return s.String()
}

func (s *PostMSDataProcessingCountShrinkRequest) SetDataIdsShrink(v string) *PostMSDataProcessingCountShrinkRequest {
	s.DataIdsShrink = &v
	return s
}

func (s *PostMSDataProcessingCountShrinkRequest) SetDataImportId(v int64) *PostMSDataProcessingCountShrinkRequest {
	s.DataImportId = &v
	return s
}

func (s *PostMSDataProcessingCountShrinkRequest) SetServiceId(v int64) *PostMSDataProcessingCountShrinkRequest {
	s.ServiceId = &v
	return s
}

type PostMSDataProcessingCountResponseBody struct {
	Code           *int32                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *PostMSDataProcessingCountResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                                     `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Msg            *string                                    `json:"Msg,omitempty" xml:"Msg,omitempty"`
	RequestId      *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PostMSDataProcessingCountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PostMSDataProcessingCountResponseBody) GoString() string {
	return s.String()
}

func (s *PostMSDataProcessingCountResponseBody) SetCode(v int32) *PostMSDataProcessingCountResponseBody {
	s.Code = &v
	return s
}

func (s *PostMSDataProcessingCountResponseBody) SetData(v *PostMSDataProcessingCountResponseBodyData) *PostMSDataProcessingCountResponseBody {
	s.Data = v
	return s
}

func (s *PostMSDataProcessingCountResponseBody) SetHttpStatusCode(v int32) *PostMSDataProcessingCountResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *PostMSDataProcessingCountResponseBody) SetMsg(v string) *PostMSDataProcessingCountResponseBody {
	s.Msg = &v
	return s
}

func (s *PostMSDataProcessingCountResponseBody) SetRequestId(v string) *PostMSDataProcessingCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *PostMSDataProcessingCountResponseBody) SetSuccess(v bool) *PostMSDataProcessingCountResponseBody {
	s.Success = &v
	return s
}

type PostMSDataProcessingCountResponseBodyData struct {
	DataProcessedStatuses []*PostMSDataProcessingCountResponseBodyDataDataProcessedStatuses `json:"DataProcessedStatuses,omitempty" xml:"DataProcessedStatuses,omitempty" type:"Repeated"`
	Status                *string                                                           `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s PostMSDataProcessingCountResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PostMSDataProcessingCountResponseBodyData) GoString() string {
	return s.String()
}

func (s *PostMSDataProcessingCountResponseBodyData) SetDataProcessedStatuses(v []*PostMSDataProcessingCountResponseBodyDataDataProcessedStatuses) *PostMSDataProcessingCountResponseBodyData {
	s.DataProcessedStatuses = v
	return s
}

func (s *PostMSDataProcessingCountResponseBodyData) SetStatus(v string) *PostMSDataProcessingCountResponseBodyData {
	s.Status = &v
	return s
}

type PostMSDataProcessingCountResponseBodyDataDataProcessedStatuses struct {
	ChunkNum      *string                                                                        `json:"ChunkNum,omitempty" xml:"ChunkNum,omitempty"`
	DataId        *string                                                                        `json:"DataId,omitempty" xml:"DataId,omitempty"`
	ErrorDataList []*PostMSDataProcessingCountResponseBodyDataDataProcessedStatusesErrorDataList `json:"ErrorDataList,omitempty" xml:"ErrorDataList,omitempty" type:"Repeated"`
	OpStatus      map[string]*int32                                                              `json:"OpStatus,omitempty" xml:"OpStatus,omitempty"`
	Status        *string                                                                        `json:"Status,omitempty" xml:"Status,omitempty"`
	VersionValue  *string                                                                        `json:"VersionValue,omitempty" xml:"VersionValue,omitempty"`
}

func (s PostMSDataProcessingCountResponseBodyDataDataProcessedStatuses) String() string {
	return tea.Prettify(s)
}

func (s PostMSDataProcessingCountResponseBodyDataDataProcessedStatuses) GoString() string {
	return s.String()
}

func (s *PostMSDataProcessingCountResponseBodyDataDataProcessedStatuses) SetChunkNum(v string) *PostMSDataProcessingCountResponseBodyDataDataProcessedStatuses {
	s.ChunkNum = &v
	return s
}

func (s *PostMSDataProcessingCountResponseBodyDataDataProcessedStatuses) SetDataId(v string) *PostMSDataProcessingCountResponseBodyDataDataProcessedStatuses {
	s.DataId = &v
	return s
}

func (s *PostMSDataProcessingCountResponseBodyDataDataProcessedStatuses) SetErrorDataList(v []*PostMSDataProcessingCountResponseBodyDataDataProcessedStatusesErrorDataList) *PostMSDataProcessingCountResponseBodyDataDataProcessedStatuses {
	s.ErrorDataList = v
	return s
}

func (s *PostMSDataProcessingCountResponseBodyDataDataProcessedStatuses) SetOpStatus(v map[string]*int32) *PostMSDataProcessingCountResponseBodyDataDataProcessedStatuses {
	s.OpStatus = v
	return s
}

func (s *PostMSDataProcessingCountResponseBodyDataDataProcessedStatuses) SetStatus(v string) *PostMSDataProcessingCountResponseBodyDataDataProcessedStatuses {
	s.Status = &v
	return s
}

func (s *PostMSDataProcessingCountResponseBodyDataDataProcessedStatuses) SetVersionValue(v string) *PostMSDataProcessingCountResponseBodyDataDataProcessedStatuses {
	s.VersionValue = &v
	return s
}

type PostMSDataProcessingCountResponseBodyDataDataProcessedStatusesErrorDataList struct {
	Count     *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	OpType    *string `json:"OpType,omitempty" xml:"OpType,omitempty"`
}

func (s PostMSDataProcessingCountResponseBodyDataDataProcessedStatusesErrorDataList) String() string {
	return tea.Prettify(s)
}

func (s PostMSDataProcessingCountResponseBodyDataDataProcessedStatusesErrorDataList) GoString() string {
	return s.String()
}

func (s *PostMSDataProcessingCountResponseBodyDataDataProcessedStatusesErrorDataList) SetCount(v int32) *PostMSDataProcessingCountResponseBodyDataDataProcessedStatusesErrorDataList {
	s.Count = &v
	return s
}

func (s *PostMSDataProcessingCountResponseBodyDataDataProcessedStatusesErrorDataList) SetErrorCode(v string) *PostMSDataProcessingCountResponseBodyDataDataProcessedStatusesErrorDataList {
	s.ErrorCode = &v
	return s
}

func (s *PostMSDataProcessingCountResponseBodyDataDataProcessedStatusesErrorDataList) SetOpType(v string) *PostMSDataProcessingCountResponseBodyDataDataProcessedStatusesErrorDataList {
	s.OpType = &v
	return s
}

type PostMSDataProcessingCountResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PostMSDataProcessingCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PostMSDataProcessingCountResponse) String() string {
	return tea.Prettify(s)
}

func (s PostMSDataProcessingCountResponse) GoString() string {
	return s.String()
}

func (s *PostMSDataProcessingCountResponse) SetHeaders(v map[string]*string) *PostMSDataProcessingCountResponse {
	s.Headers = v
	return s
}

func (s *PostMSDataProcessingCountResponse) SetStatusCode(v int32) *PostMSDataProcessingCountResponse {
	s.StatusCode = &v
	return s
}

func (s *PostMSDataProcessingCountResponse) SetBody(v *PostMSDataProcessingCountResponseBody) *PostMSDataProcessingCountResponse {
	s.Body = v
	return s
}

type PostMSSearchEnhanceRequest struct {
	Body             *string                `json:"Body,omitempty" xml:"Body,omitempty"`
	CustomConfigInfo map[string]interface{} `json:"CustomConfigInfo,omitempty" xml:"CustomConfigInfo,omitempty"`
	Debug            *bool                  `json:"Debug,omitempty" xml:"Debug,omitempty"`
	Fields           []*string              `json:"Fields,omitempty" xml:"Fields,omitempty" type:"Repeated"`
	Filters          *string                `json:"Filters,omitempty" xml:"Filters,omitempty"`
	MinScore         *float64               `json:"MinScore,omitempty" xml:"MinScore,omitempty"`
	Page             *int32                 `json:"Page,omitempty" xml:"Page,omitempty"`
	Queries          *string                `json:"Queries,omitempty" xml:"Queries,omitempty"`
	RankModelInfo    map[string]interface{} `json:"RankModelInfo,omitempty" xml:"RankModelInfo,omitempty"`
	Rows             *int32                 `json:"Rows,omitempty" xml:"Rows,omitempty"`
	ServiceId        *int64                 `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	Sort             []*string              `json:"Sort,omitempty" xml:"Sort,omitempty" type:"Repeated"`
	Type             *string                `json:"Type,omitempty" xml:"Type,omitempty"`
	Uq               *string                `json:"Uq,omitempty" xml:"Uq,omitempty"`
}

func (s PostMSSearchEnhanceRequest) String() string {
	return tea.Prettify(s)
}

func (s PostMSSearchEnhanceRequest) GoString() string {
	return s.String()
}

func (s *PostMSSearchEnhanceRequest) SetBody(v string) *PostMSSearchEnhanceRequest {
	s.Body = &v
	return s
}

func (s *PostMSSearchEnhanceRequest) SetCustomConfigInfo(v map[string]interface{}) *PostMSSearchEnhanceRequest {
	s.CustomConfigInfo = v
	return s
}

func (s *PostMSSearchEnhanceRequest) SetDebug(v bool) *PostMSSearchEnhanceRequest {
	s.Debug = &v
	return s
}

func (s *PostMSSearchEnhanceRequest) SetFields(v []*string) *PostMSSearchEnhanceRequest {
	s.Fields = v
	return s
}

func (s *PostMSSearchEnhanceRequest) SetFilters(v string) *PostMSSearchEnhanceRequest {
	s.Filters = &v
	return s
}

func (s *PostMSSearchEnhanceRequest) SetMinScore(v float64) *PostMSSearchEnhanceRequest {
	s.MinScore = &v
	return s
}

func (s *PostMSSearchEnhanceRequest) SetPage(v int32) *PostMSSearchEnhanceRequest {
	s.Page = &v
	return s
}

func (s *PostMSSearchEnhanceRequest) SetQueries(v string) *PostMSSearchEnhanceRequest {
	s.Queries = &v
	return s
}

func (s *PostMSSearchEnhanceRequest) SetRankModelInfo(v map[string]interface{}) *PostMSSearchEnhanceRequest {
	s.RankModelInfo = v
	return s
}

func (s *PostMSSearchEnhanceRequest) SetRows(v int32) *PostMSSearchEnhanceRequest {
	s.Rows = &v
	return s
}

func (s *PostMSSearchEnhanceRequest) SetServiceId(v int64) *PostMSSearchEnhanceRequest {
	s.ServiceId = &v
	return s
}

func (s *PostMSSearchEnhanceRequest) SetSort(v []*string) *PostMSSearchEnhanceRequest {
	s.Sort = v
	return s
}

func (s *PostMSSearchEnhanceRequest) SetType(v string) *PostMSSearchEnhanceRequest {
	s.Type = &v
	return s
}

func (s *PostMSSearchEnhanceRequest) SetUq(v string) *PostMSSearchEnhanceRequest {
	s.Uq = &v
	return s
}

type PostMSSearchEnhanceShrinkRequest struct {
	Body                   *string  `json:"Body,omitempty" xml:"Body,omitempty"`
	CustomConfigInfoShrink *string  `json:"CustomConfigInfo,omitempty" xml:"CustomConfigInfo,omitempty"`
	Debug                  *bool    `json:"Debug,omitempty" xml:"Debug,omitempty"`
	FieldsShrink           *string  `json:"Fields,omitempty" xml:"Fields,omitempty"`
	Filters                *string  `json:"Filters,omitempty" xml:"Filters,omitempty"`
	MinScore               *float64 `json:"MinScore,omitempty" xml:"MinScore,omitempty"`
	Page                   *int32   `json:"Page,omitempty" xml:"Page,omitempty"`
	Queries                *string  `json:"Queries,omitempty" xml:"Queries,omitempty"`
	RankModelInfoShrink    *string  `json:"RankModelInfo,omitempty" xml:"RankModelInfo,omitempty"`
	Rows                   *int32   `json:"Rows,omitempty" xml:"Rows,omitempty"`
	ServiceId              *int64   `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	SortShrink             *string  `json:"Sort,omitempty" xml:"Sort,omitempty"`
	Type                   *string  `json:"Type,omitempty" xml:"Type,omitempty"`
	Uq                     *string  `json:"Uq,omitempty" xml:"Uq,omitempty"`
}

func (s PostMSSearchEnhanceShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s PostMSSearchEnhanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *PostMSSearchEnhanceShrinkRequest) SetBody(v string) *PostMSSearchEnhanceShrinkRequest {
	s.Body = &v
	return s
}

func (s *PostMSSearchEnhanceShrinkRequest) SetCustomConfigInfoShrink(v string) *PostMSSearchEnhanceShrinkRequest {
	s.CustomConfigInfoShrink = &v
	return s
}

func (s *PostMSSearchEnhanceShrinkRequest) SetDebug(v bool) *PostMSSearchEnhanceShrinkRequest {
	s.Debug = &v
	return s
}

func (s *PostMSSearchEnhanceShrinkRequest) SetFieldsShrink(v string) *PostMSSearchEnhanceShrinkRequest {
	s.FieldsShrink = &v
	return s
}

func (s *PostMSSearchEnhanceShrinkRequest) SetFilters(v string) *PostMSSearchEnhanceShrinkRequest {
	s.Filters = &v
	return s
}

func (s *PostMSSearchEnhanceShrinkRequest) SetMinScore(v float64) *PostMSSearchEnhanceShrinkRequest {
	s.MinScore = &v
	return s
}

func (s *PostMSSearchEnhanceShrinkRequest) SetPage(v int32) *PostMSSearchEnhanceShrinkRequest {
	s.Page = &v
	return s
}

func (s *PostMSSearchEnhanceShrinkRequest) SetQueries(v string) *PostMSSearchEnhanceShrinkRequest {
	s.Queries = &v
	return s
}

func (s *PostMSSearchEnhanceShrinkRequest) SetRankModelInfoShrink(v string) *PostMSSearchEnhanceShrinkRequest {
	s.RankModelInfoShrink = &v
	return s
}

func (s *PostMSSearchEnhanceShrinkRequest) SetRows(v int32) *PostMSSearchEnhanceShrinkRequest {
	s.Rows = &v
	return s
}

func (s *PostMSSearchEnhanceShrinkRequest) SetServiceId(v int64) *PostMSSearchEnhanceShrinkRequest {
	s.ServiceId = &v
	return s
}

func (s *PostMSSearchEnhanceShrinkRequest) SetSortShrink(v string) *PostMSSearchEnhanceShrinkRequest {
	s.SortShrink = &v
	return s
}

func (s *PostMSSearchEnhanceShrinkRequest) SetType(v string) *PostMSSearchEnhanceShrinkRequest {
	s.Type = &v
	return s
}

func (s *PostMSSearchEnhanceShrinkRequest) SetUq(v string) *PostMSSearchEnhanceShrinkRequest {
	s.Uq = &v
	return s
}

type PostMSSearchEnhanceResponseBody struct {
	Code           *int32      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	HttpStatusCode *int32      `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Msg            *string     `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PostMSSearchEnhanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PostMSSearchEnhanceResponseBody) GoString() string {
	return s.String()
}

func (s *PostMSSearchEnhanceResponseBody) SetCode(v int32) *PostMSSearchEnhanceResponseBody {
	s.Code = &v
	return s
}

func (s *PostMSSearchEnhanceResponseBody) SetData(v interface{}) *PostMSSearchEnhanceResponseBody {
	s.Data = v
	return s
}

func (s *PostMSSearchEnhanceResponseBody) SetHttpStatusCode(v int32) *PostMSSearchEnhanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *PostMSSearchEnhanceResponseBody) SetMsg(v string) *PostMSSearchEnhanceResponseBody {
	s.Msg = &v
	return s
}

func (s *PostMSSearchEnhanceResponseBody) SetRequestId(v string) *PostMSSearchEnhanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *PostMSSearchEnhanceResponseBody) SetSuccess(v bool) *PostMSSearchEnhanceResponseBody {
	s.Success = &v
	return s
}

type PostMSSearchEnhanceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PostMSSearchEnhanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PostMSSearchEnhanceResponse) String() string {
	return tea.Prettify(s)
}

func (s PostMSSearchEnhanceResponse) GoString() string {
	return s.String()
}

func (s *PostMSSearchEnhanceResponse) SetHeaders(v map[string]*string) *PostMSSearchEnhanceResponse {
	s.Headers = v
	return s
}

func (s *PostMSSearchEnhanceResponse) SetStatusCode(v int32) *PostMSSearchEnhanceResponse {
	s.StatusCode = &v
	return s
}

func (s *PostMSSearchEnhanceResponse) SetBody(v *PostMSSearchEnhanceResponseBody) *PostMSSearchEnhanceResponse {
	s.Body = v
	return s
}

type PostMSServiceDataImportRequest struct {
	DataType  *string                                    `json:"DataType,omitempty" xml:"DataType,omitempty"`
	Documents []*PostMSServiceDataImportRequestDocuments `json:"Documents,omitempty" xml:"Documents,omitempty" type:"Repeated"`
	ServiceId *int64                                     `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s PostMSServiceDataImportRequest) String() string {
	return tea.Prettify(s)
}

func (s PostMSServiceDataImportRequest) GoString() string {
	return s.String()
}

func (s *PostMSServiceDataImportRequest) SetDataType(v string) *PostMSServiceDataImportRequest {
	s.DataType = &v
	return s
}

func (s *PostMSServiceDataImportRequest) SetDocuments(v []*PostMSServiceDataImportRequestDocuments) *PostMSServiceDataImportRequest {
	s.Documents = v
	return s
}

func (s *PostMSServiceDataImportRequest) SetServiceId(v int64) *PostMSServiceDataImportRequest {
	s.ServiceId = &v
	return s
}

type PostMSServiceDataImportRequestDocuments struct {
	BizParams     map[string]interface{} `json:"BizParams,omitempty" xml:"BizParams,omitempty"`
	DocId         *string                `json:"DocId,omitempty" xml:"DocId,omitempty"`
	FileExtension *string                `json:"FileExtension,omitempty" xml:"FileExtension,omitempty"`
	FileName      *string                `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FilePath      *string                `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	Version       *string                `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s PostMSServiceDataImportRequestDocuments) String() string {
	return tea.Prettify(s)
}

func (s PostMSServiceDataImportRequestDocuments) GoString() string {
	return s.String()
}

func (s *PostMSServiceDataImportRequestDocuments) SetBizParams(v map[string]interface{}) *PostMSServiceDataImportRequestDocuments {
	s.BizParams = v
	return s
}

func (s *PostMSServiceDataImportRequestDocuments) SetDocId(v string) *PostMSServiceDataImportRequestDocuments {
	s.DocId = &v
	return s
}

func (s *PostMSServiceDataImportRequestDocuments) SetFileExtension(v string) *PostMSServiceDataImportRequestDocuments {
	s.FileExtension = &v
	return s
}

func (s *PostMSServiceDataImportRequestDocuments) SetFileName(v string) *PostMSServiceDataImportRequestDocuments {
	s.FileName = &v
	return s
}

func (s *PostMSServiceDataImportRequestDocuments) SetFilePath(v string) *PostMSServiceDataImportRequestDocuments {
	s.FilePath = &v
	return s
}

func (s *PostMSServiceDataImportRequestDocuments) SetVersion(v string) *PostMSServiceDataImportRequestDocuments {
	s.Version = &v
	return s
}

type PostMSServiceDataImportShrinkRequest struct {
	DataType        *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	DocumentsShrink *string `json:"Documents,omitempty" xml:"Documents,omitempty"`
	ServiceId       *int64  `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s PostMSServiceDataImportShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s PostMSServiceDataImportShrinkRequest) GoString() string {
	return s.String()
}

func (s *PostMSServiceDataImportShrinkRequest) SetDataType(v string) *PostMSServiceDataImportShrinkRequest {
	s.DataType = &v
	return s
}

func (s *PostMSServiceDataImportShrinkRequest) SetDocumentsShrink(v string) *PostMSServiceDataImportShrinkRequest {
	s.DocumentsShrink = &v
	return s
}

func (s *PostMSServiceDataImportShrinkRequest) SetServiceId(v int64) *PostMSServiceDataImportShrinkRequest {
	s.ServiceId = &v
	return s
}

type PostMSServiceDataImportResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	Msg       *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PostMSServiceDataImportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PostMSServiceDataImportResponseBody) GoString() string {
	return s.String()
}

func (s *PostMSServiceDataImportResponseBody) SetCode(v int32) *PostMSServiceDataImportResponseBody {
	s.Code = &v
	return s
}

func (s *PostMSServiceDataImportResponseBody) SetData(v int64) *PostMSServiceDataImportResponseBody {
	s.Data = &v
	return s
}

func (s *PostMSServiceDataImportResponseBody) SetMsg(v string) *PostMSServiceDataImportResponseBody {
	s.Msg = &v
	return s
}

func (s *PostMSServiceDataImportResponseBody) SetRequestId(v string) *PostMSServiceDataImportResponseBody {
	s.RequestId = &v
	return s
}

func (s *PostMSServiceDataImportResponseBody) SetSuccess(v bool) *PostMSServiceDataImportResponseBody {
	s.Success = &v
	return s
}

type PostMSServiceDataImportResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PostMSServiceDataImportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PostMSServiceDataImportResponse) String() string {
	return tea.Prettify(s)
}

func (s PostMSServiceDataImportResponse) GoString() string {
	return s.String()
}

func (s *PostMSServiceDataImportResponse) SetHeaders(v map[string]*string) *PostMSServiceDataImportResponse {
	s.Headers = v
	return s
}

func (s *PostMSServiceDataImportResponse) SetStatusCode(v int32) *PostMSServiceDataImportResponse {
	s.StatusCode = &v
	return s
}

func (s *PostMSServiceDataImportResponse) SetBody(v *PostMSServiceDataImportResponseBody) *PostMSServiceDataImportResponse {
	s.Body = v
	return s
}

type RequestTableQARequest struct {
	Params      *string `json:"Params,omitempty" xml:"Params,omitempty"`
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
}

func (s RequestTableQARequest) String() string {
	return tea.Prettify(s)
}

func (s RequestTableQARequest) GoString() string {
	return s.String()
}

func (s *RequestTableQARequest) SetParams(v string) *RequestTableQARequest {
	s.Params = &v
	return s
}

func (s *RequestTableQARequest) SetServiceCode(v string) *RequestTableQARequest {
	s.ServiceCode = &v
	return s
}

type RequestTableQAResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RequestTableQAResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RequestTableQAResponseBody) GoString() string {
	return s.String()
}

func (s *RequestTableQAResponseBody) SetData(v string) *RequestTableQAResponseBody {
	s.Data = &v
	return s
}

func (s *RequestTableQAResponseBody) SetRequestId(v string) *RequestTableQAResponseBody {
	s.RequestId = &v
	return s
}

type RequestTableQAResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RequestTableQAResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RequestTableQAResponse) String() string {
	return tea.Prettify(s)
}

func (s RequestTableQAResponse) GoString() string {
	return s.String()
}

func (s *RequestTableQAResponse) SetHeaders(v map[string]*string) *RequestTableQAResponse {
	s.Headers = v
	return s
}

func (s *RequestTableQAResponse) SetStatusCode(v int32) *RequestTableQAResponse {
	s.StatusCode = &v
	return s
}

func (s *RequestTableQAResponse) SetBody(v *RequestTableQAResponseBody) *RequestTableQAResponse {
	s.Body = v
	return s
}

type RequestTableQAOnlineRequest struct {
	BotId       *string `json:"BotId,omitempty" xml:"BotId,omitempty"`
	Params      *string `json:"Params,omitempty" xml:"Params,omitempty"`
	Question    *string `json:"Question,omitempty" xml:"Question,omitempty"`
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
}

func (s RequestTableQAOnlineRequest) String() string {
	return tea.Prettify(s)
}

func (s RequestTableQAOnlineRequest) GoString() string {
	return s.String()
}

func (s *RequestTableQAOnlineRequest) SetBotId(v string) *RequestTableQAOnlineRequest {
	s.BotId = &v
	return s
}

func (s *RequestTableQAOnlineRequest) SetParams(v string) *RequestTableQAOnlineRequest {
	s.Params = &v
	return s
}

func (s *RequestTableQAOnlineRequest) SetQuestion(v string) *RequestTableQAOnlineRequest {
	s.Question = &v
	return s
}

func (s *RequestTableQAOnlineRequest) SetServiceCode(v string) *RequestTableQAOnlineRequest {
	s.ServiceCode = &v
	return s
}

type RequestTableQAOnlineResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RequestTableQAOnlineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RequestTableQAOnlineResponseBody) GoString() string {
	return s.String()
}

func (s *RequestTableQAOnlineResponseBody) SetData(v string) *RequestTableQAOnlineResponseBody {
	s.Data = &v
	return s
}

func (s *RequestTableQAOnlineResponseBody) SetRequestId(v string) *RequestTableQAOnlineResponseBody {
	s.RequestId = &v
	return s
}

type RequestTableQAOnlineResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RequestTableQAOnlineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RequestTableQAOnlineResponse) String() string {
	return tea.Prettify(s)
}

func (s RequestTableQAOnlineResponse) GoString() string {
	return s.String()
}

func (s *RequestTableQAOnlineResponse) SetHeaders(v map[string]*string) *RequestTableQAOnlineResponse {
	s.Headers = v
	return s
}

func (s *RequestTableQAOnlineResponse) SetStatusCode(v int32) *RequestTableQAOnlineResponse {
	s.StatusCode = &v
	return s
}

func (s *RequestTableQAOnlineResponse) SetBody(v *RequestTableQAOnlineResponseBody) *RequestTableQAOnlineResponse {
	s.Body = v
	return s
}

type UpdateServiceDataRequest struct {
	Conditions map[string]interface{} `json:"Conditions,omitempty" xml:"Conditions,omitempty"`
	ServiceId  *int64                 `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s UpdateServiceDataRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceDataRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceDataRequest) SetConditions(v map[string]interface{}) *UpdateServiceDataRequest {
	s.Conditions = v
	return s
}

func (s *UpdateServiceDataRequest) SetServiceId(v int64) *UpdateServiceDataRequest {
	s.ServiceId = &v
	return s
}

type UpdateServiceDataShrinkRequest struct {
	ConditionsShrink *string `json:"Conditions,omitempty" xml:"Conditions,omitempty"`
	ServiceId        *int64  `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s UpdateServiceDataShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceDataShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceDataShrinkRequest) SetConditionsShrink(v string) *UpdateServiceDataShrinkRequest {
	s.ConditionsShrink = &v
	return s
}

func (s *UpdateServiceDataShrinkRequest) SetServiceId(v int64) *UpdateServiceDataShrinkRequest {
	s.ServiceId = &v
	return s
}

type UpdateServiceDataResponseBody struct {
	Code      *int32      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Msg       *string     `json:"Msg,omitempty" xml:"Msg,omitempty"`
	RequestId *string     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateServiceDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceDataResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServiceDataResponseBody) SetCode(v int32) *UpdateServiceDataResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateServiceDataResponseBody) SetData(v interface{}) *UpdateServiceDataResponseBody {
	s.Data = v
	return s
}

func (s *UpdateServiceDataResponseBody) SetMsg(v string) *UpdateServiceDataResponseBody {
	s.Msg = &v
	return s
}

func (s *UpdateServiceDataResponseBody) SetRequestId(v string) *UpdateServiceDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateServiceDataResponseBody) SetSuccess(v bool) *UpdateServiceDataResponseBody {
	s.Success = &v
	return s
}

type UpdateServiceDataResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateServiceDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateServiceDataResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceDataResponse) GoString() string {
	return s.String()
}

func (s *UpdateServiceDataResponse) SetHeaders(v map[string]*string) *UpdateServiceDataResponse {
	s.Headers = v
	return s
}

func (s *UpdateServiceDataResponse) SetStatusCode(v int32) *UpdateServiceDataResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateServiceDataResponse) SetBody(v *UpdateServiceDataResponseBody) *UpdateServiceDataResponse {
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
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("alinlp"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) ADClockWithOptions(request *ADClockRequest, runtime *util.RuntimeOptions) (_result *ADClockResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Params)) {
		body["Params"] = request.Params
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		body["ServiceCode"] = request.ServiceCode
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ADClock"),
		Version:     tea.String("2020-06-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ADClockResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ADClock(request *ADClockRequest) (_result *ADClockResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ADClockResponse{}
	_body, _err := client.ADClockWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ADMMUWithOptions(request *ADMMURequest, runtime *util.RuntimeOptions) (_result *ADMMUResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Params)) {
		body["Params"] = request.Params
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		body["ServiceCode"] = request.ServiceCode
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ADMMU"),
		Version:     tea.String("2020-06-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ADMMUResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ADMMU(request *ADMMURequest) (_result *ADMMUResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ADMMUResponse{}
	_body, _err := client.ADMMUWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ADMiniCogWithOptions(request *ADMiniCogRequest, runtime *util.RuntimeOptions) (_result *ADMiniCogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Params)) {
		body["Params"] = request.Params
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		body["ServiceCode"] = request.ServiceCode
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ADMiniCog"),
		Version:     tea.String("2020-06-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ADMiniCogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ADMiniCog(request *ADMiniCogRequest) (_result *ADMiniCogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ADMiniCogResponse{}
	_body, _err := client.ADMiniCogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ADMiniCogResultWithOptions(request *ADMiniCogResultRequest, runtime *util.RuntimeOptions) (_result *ADMiniCogResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Params)) {
		body["Params"] = request.Params
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		body["ServiceCode"] = request.ServiceCode
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ADMiniCogResult"),
		Version:     tea.String("2020-06-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ADMiniCogResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ADMiniCogResult(request *ADMiniCogResultRequest) (_result *ADMiniCogResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ADMiniCogResultResponse{}
	_body, _err := client.ADMiniCogResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteServiceDataByConditionsWithOptions(tmpReq *DeleteServiceDataByConditionsRequest, runtime *util.RuntimeOptions) (_result *DeleteServiceDataByConditionsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DeleteServiceDataByConditionsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Conditions)) {
		request.ConditionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Conditions, tea.String("Conditions"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConditionsShrink)) {
		body["Conditions"] = request.ConditionsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		body["ServiceId"] = request.ServiceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteServiceDataByConditions"),
		Version:     tea.String("2020-06-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteServiceDataByConditionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteServiceDataByConditions(request *DeleteServiceDataByConditionsRequest) (_result *DeleteServiceDataByConditionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteServiceDataByConditionsResponse{}
	_body, _err := client.DeleteServiceDataByConditionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteServiceDataByIdsWithOptions(tmpReq *DeleteServiceDataByIdsRequest, runtime *util.RuntimeOptions) (_result *DeleteServiceDataByIdsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DeleteServiceDataByIdsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Ids)) {
		request.IdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Ids, tea.String("Ids"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IdsShrink)) {
		body["Ids"] = request.IdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		body["ServiceId"] = request.ServiceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteServiceDataByIds"),
		Version:     tea.String("2020-06-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteServiceDataByIdsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteServiceDataByIds(request *DeleteServiceDataByIdsRequest) (_result *DeleteServiceDataByIdsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteServiceDataByIdsResponse{}
	_body, _err := client.DeleteServiceDataByIdsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetBrandChEcomWithOptions(request *GetBrandChEcomRequest, runtime *util.RuntimeOptions) (_result *GetBrandChEcomResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageUrl)) {
		body["ImageUrl"] = request.ImageUrl
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		body["ServiceCode"] = request.ServiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.Text)) {
		body["Text"] = request.Text
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetBrandChEcom"),
		Version:     tea.String("2020-06-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetBrandChEcomResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetBrandChEcom(request *GetBrandChEcomRequest) (_result *GetBrandChEcomResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetBrandChEcomResponse{}
	_body, _err := client.GetBrandChEcomWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCateChEcomWithOptions(request *GetCateChEcomRequest, runtime *util.RuntimeOptions) (_result *GetCateChEcomResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageUrl)) {
		body["ImageUrl"] = request.ImageUrl
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		body["ServiceCode"] = request.ServiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.Text)) {
		body["Text"] = request.Text
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCateChEcom"),
		Version:     tea.String("2020-06-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCateChEcomResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCateChEcom(request *GetCateChEcomRequest) (_result *GetCateChEcomResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCateChEcomResponse{}
	_body, _err := client.GetCateChEcomWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCheckDuplicationChMedicalWithOptions(request *GetCheckDuplicationChMedicalRequest, runtime *util.RuntimeOptions) (_result *GetCheckDuplicationChMedicalResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OriginQ)) {
		body["OriginQ"] = request.OriginQ
	}

	if !tea.BoolValue(util.IsUnset(request.OriginT)) {
		body["OriginT"] = request.OriginT
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		body["ServiceCode"] = request.ServiceCode
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCheckDuplicationChMedical"),
		Version:     tea.String("2020-06-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCheckDuplicationChMedicalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCheckDuplicationChMedical(request *GetCheckDuplicationChMedicalRequest) (_result *GetCheckDuplicationChMedicalResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCheckDuplicationChMedicalResponse{}
	_body, _err := client.GetCheckDuplicationChMedicalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDiagnosisChMedicalWithOptions(request *GetDiagnosisChMedicalRequest, runtime *util.RuntimeOptions) (_result *GetDiagnosisChMedicalResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		body["ServiceCode"] = request.ServiceCode
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDiagnosisChMedical"),
		Version:     tea.String("2020-06-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDiagnosisChMedicalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDiagnosisChMedical(request *GetDiagnosisChMedicalRequest) (_result *GetDiagnosisChMedicalResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDiagnosisChMedicalResponse{}
	_body, _err := client.GetDiagnosisChMedicalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDpChEcomWithOptions(request *GetDpChEcomRequest, runtime *util.RuntimeOptions) (_result *GetDpChEcomResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		body["ServiceCode"] = request.ServiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.Text)) {
		body["Text"] = request.Text
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDpChEcom"),
		Version:     tea.String("2020-06-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDpChEcomResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDpChEcom(request *GetDpChEcomRequest) (_result *GetDpChEcomResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDpChEcomResponse{}
	_body, _err := client.GetDpChEcomWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDpChGeneralCTBWithOptions(request *GetDpChGeneralCTBRequest, runtime *util.RuntimeOptions) (_result *GetDpChGeneralCTBResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		body["ServiceCode"] = request.ServiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.Text)) {
		body["Text"] = request.Text
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDpChGeneralCTB"),
		Version:     tea.String("2020-06-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDpChGeneralCTBResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDpChGeneralCTB(request *GetDpChGeneralCTBRequest) (_result *GetDpChGeneralCTBResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDpChGeneralCTBResponse{}
	_body, _err := client.GetDpChGeneralCTBWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDpChGeneralStanfordWithOptions(request *GetDpChGeneralStanfordRequest, runtime *util.RuntimeOptions) (_result *GetDpChGeneralStanfordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		body["ServiceCode"] = request.ServiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.Text)) {
		body["Text"] = request.Text
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDpChGeneralStanford"),
		Version:     tea.String("2020-06-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDpChGeneralStanfordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDpChGeneralStanford(request *GetDpChGeneralStanfordRequest) (_result *GetDpChGeneralStanfordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDpChGeneralStanfordResponse{}
	_body, _err := client.GetDpChGeneralStanfordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetEcChGeneralWithOptions(request *GetEcChGeneralRequest, runtime *util.RuntimeOptions) (_result *GetEcChGeneralResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		body["ServiceCode"] = request.ServiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.Text)) {
		body["Text"] = request.Text
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetEcChGeneral"),
		Version:     tea.String("2020-06-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetEcChGeneralResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetEcChGeneral(request *GetEcChGeneralRequest) (_result *GetEcChGeneralResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetEcChGeneralResponse{}
	_body, _err := client.GetEcChGeneralWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetEcEnGeneralWithOptions(request *GetEcEnGeneralRequest, runtime *util.RuntimeOptions) (_result *GetEcEnGeneralResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		body["ServiceCode"] = request.ServiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.Text)) {
		body["Text"] = request.Text
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetEcEnGeneral"),
		Version:     tea.String("2020-06-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetEcEnGeneralResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetEcEnGeneral(request *GetEcEnGeneralRequest) (_result *GetEcEnGeneralResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetEcEnGeneralResponse{}
	_body, _err := client.GetEcEnGeneralWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetEmbeddingWithOptions(request *GetEmbeddingRequest, runtime *util.RuntimeOptions) (_result *GetEmbeddingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		body["ServiceCode"] = request.ServiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.Text)) {
		body["Text"] = request.Text
	}

	if !tea.BoolValue(util.IsUnset(request.TextType)) {
		body["TextType"] = request.TextType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetEmbedding"),
		Version:     tea.String("2020-06-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetEmbeddingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetEmbedding(request *GetEmbeddingRequest) (_result *GetEmbeddingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetEmbeddingResponse{}
	_body, _err := client.GetEmbeddingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetItemPubChEcomWithOptions(request *GetItemPubChEcomRequest, runtime *util.RuntimeOptions) (_result *GetItemPubChEcomResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageUrl)) {
		body["ImageUrl"] = request.ImageUrl
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		body["ServiceCode"] = request.ServiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.Text)) {
		body["Text"] = request.Text
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetItemPubChEcom"),
		Version:     tea.String("2020-06-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetItemPubChEcomResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetItemPubChEcom(request *GetItemPubChEcomRequest) (_result *GetItemPubChEcomResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetItemPubChEcomResponse{}
	_body, _err := client.GetItemPubChEcomWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetKeywordChEcomWithOptions(request *GetKeywordChEcomRequest, runtime *util.RuntimeOptions) (_result *GetKeywordChEcomResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiVersion)) {
		body["ApiVersion"] = request.ApiVersion
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		body["ServiceCode"] = request.ServiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.Text)) {
		body["Text"] = request.Text
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetKeywordChEcom"),
		Version:     tea.String("2020-06-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetKeywordChEcomResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetKeywordChEcom(request *GetKeywordChEcomRequest) (_result *GetKeywordChEcomResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetKeywordChEcomResponse{}
	_body, _err := client.GetKeywordChEcomWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetKeywordEnEcomWithOptions(request *GetKeywordEnEcomRequest, runtime *util.RuntimeOptions) (_result *GetKeywordEnEcomResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		body["ServiceCode"] = request.ServiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.Text)) {
		body["Text"] = request.Text
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetKeywordEnEcom"),
		Version:     tea.String("2020-06-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetKeywordEnEcomResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetKeywordEnEcom(request *GetKeywordEnEcomRequest) (_result *GetKeywordEnEcomResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetKeywordEnEcomResponse{}
	_body, _err := client.GetKeywordEnEcomWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMedicineChMedicalWithOptions(request *GetMedicineChMedicalRequest, runtime *util.RuntimeOptions) (_result *GetMedicineChMedicalResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Factory)) {
		body["Factory"] = request.Factory
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		body["ServiceCode"] = request.ServiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.Specification)) {
		body["Specification"] = request.Specification
	}

	if !tea.BoolValue(util.IsUnset(request.Unit)) {
		body["Unit"] = request.Unit
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetMedicineChMedical"),
		Version:     tea.String("2020-06-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMedicineChMedicalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMedicineChMedical(request *GetMedicineChMedicalRequest) (_result *GetMedicineChMedicalResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMedicineChMedicalResponse{}
	_body, _err := client.GetMedicineChMedicalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetNerChEcomWithOptions(request *GetNerChEcomRequest, runtime *util.RuntimeOptions) (_result *GetNerChEcomResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LexerId)) {
		body["LexerId"] = request.LexerId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		body["ServiceCode"] = request.ServiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.Text)) {
		body["Text"] = request.Text
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetNerChEcom"),
		Version:     tea.String("2020-06-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetNerChEcomResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetNerChEcom(request *GetNerChEcomRequest) (_result *GetNerChEcomResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetNerChEcomResponse{}
	_body, _err := client.GetNerChEcomWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetNerChMedicalWithOptions(request *GetNerChMedicalRequest, runtime *util.RuntimeOptions) (_result *GetNerChMedicalResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		body["ServiceCode"] = request.ServiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.Text)) {
		body["Text"] = request.Text
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetNerChMedical"),
		Version:     tea.String("2020-06-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetNerChMedicalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetNerChMedical(request *GetNerChMedicalRequest) (_result *GetNerChMedicalResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetNerChMedicalResponse{}
	_body, _err := client.GetNerChMedicalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetNerCustomizedChEcomWithOptions(request *GetNerCustomizedChEcomRequest, runtime *util.RuntimeOptions) (_result *GetNerCustomizedChEcomResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LexerId)) {
		body["LexerId"] = request.LexerId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		body["ServiceCode"] = request.ServiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.Text)) {
		body["Text"] = request.Text
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetNerCustomizedChEcom"),
		Version:     tea.String("2020-06-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetNerCustomizedChEcomResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetNerCustomizedChEcom(request *GetNerCustomizedChEcomRequest) (_result *GetNerCustomizedChEcomResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetNerCustomizedChEcomResponse{}
	_body, _err := client.GetNerCustomizedChEcomWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetNerCustomizedSeaEcomWithOptions(request *GetNerCustomizedSeaEcomRequest, runtime *util.RuntimeOptions) (_result *GetNerCustomizedSeaEcomResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Language)) {
		body["Language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		body["ServiceCode"] = request.ServiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.Text)) {
		body["Text"] = request.Text
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetNerCustomizedSeaEcom"),
		Version:     tea.String("2020-06-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetNerCustomizedSeaEcomResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetNerCustomizedSeaEcom(request *GetNerCustomizedSeaEcomRequest) (_result *GetNerCustomizedSeaEcomResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetNerCustomizedSeaEcomResponse{}
	_body, _err := client.GetNerCustomizedSeaEcomWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOpenNLUWithOptions(request *GetOpenNLURequest, runtime *util.RuntimeOptions) (_result *GetOpenNLUResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Examples)) {
		body["Examples"] = request.Examples
	}

	if !tea.BoolValue(util.IsUnset(request.Labels)) {
		body["Labels"] = request.Labels
	}

	if !tea.BoolValue(util.IsUnset(request.Sentence)) {
		body["Sentence"] = request.Sentence
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		body["ServiceCode"] = request.ServiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.Task)) {
		body["Task"] = request.Task
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOpenNLU"),
		Version:     tea.String("2020-06-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOpenNLUResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOpenNLU(request *GetOpenNLURequest) (_result *GetOpenNLUResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOpenNLUResponse{}
	_body, _err := client.GetOpenNLUWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOpenNLUHighRecallWithOptions(request *GetOpenNLUHighRecallRequest, runtime *util.RuntimeOptions) (_result *GetOpenNLUHighRecallResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Examples)) {
		body["Examples"] = request.Examples
	}

	if !tea.BoolValue(util.IsUnset(request.Labels)) {
		body["Labels"] = request.Labels
	}

	if !tea.BoolValue(util.IsUnset(request.Sentence)) {
		body["Sentence"] = request.Sentence
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		body["ServiceCode"] = request.ServiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.Task)) {
		body["Task"] = request.Task
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOpenNLUHighRecall"),
		Version:     tea.String("2020-06-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOpenNLUHighRecallResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOpenNLUHighRecall(request *GetOpenNLUHighRecallRequest) (_result *GetOpenNLUHighRecallResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOpenNLUHighRecallResponse{}
	_body, _err := client.GetOpenNLUHighRecallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOperationChMedicalWithOptions(request *GetOperationChMedicalRequest, runtime *util.RuntimeOptions) (_result *GetOperationChMedicalResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		body["ServiceCode"] = request.ServiceCode
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOperationChMedical"),
		Version:     tea.String("2020-06-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOperationChMedicalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOperationChMedical(request *GetOperationChMedicalRequest) (_result *GetOperationChMedicalResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOperationChMedicalResponse{}
	_body, _err := client.GetOperationChMedicalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPosChEcomWithOptions(request *GetPosChEcomRequest, runtime *util.RuntimeOptions) (_result *GetPosChEcomResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OutType)) {
		body["OutType"] = request.OutType
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		body["ServiceCode"] = request.ServiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.Text)) {
		body["Text"] = request.Text
	}

	if !tea.BoolValue(util.IsUnset(request.TokenizerId)) {
		body["TokenizerId"] = request.TokenizerId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPosChEcom"),
		Version:     tea.String("2020-06-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPosChEcomResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPosChEcom(request *GetPosChEcomRequest) (_result *GetPosChEcomResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPosChEcomResponse{}
	_body, _err := client.GetPosChEcomWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPosChGeneralWithOptions(request *GetPosChGeneralRequest, runtime *util.RuntimeOptions) (_result *GetPosChGeneralResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OutType)) {
		body["OutType"] = request.OutType
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		body["ServiceCode"] = request.ServiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.Text)) {
		body["Text"] = request.Text
	}

	if !tea.BoolValue(util.IsUnset(request.TokenizerId)) {
		body["TokenizerId"] = request.TokenizerId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPosChGeneral"),
		Version:     tea.String("2020-06-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPosChGeneralResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPosChGeneral(request *GetPosChGeneralRequest) (_result *GetPosChGeneralResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPosChGeneralResponse{}
	_body, _err := client.GetPosChGeneralWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPriceChEcomWithOptions(request *GetPriceChEcomRequest, runtime *util.RuntimeOptions) (_result *GetPriceChEcomResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		body["ServiceCode"] = request.ServiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.Text)) {
		body["Text"] = request.Text
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPriceChEcom"),
		Version:     tea.String("2020-06-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPriceChEcomResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPriceChEcom(request *GetPriceChEcomRequest) (_result *GetPriceChEcomResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPriceChEcomResponse{}
	_body, _err := client.GetPriceChEcomWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSSETestWithOptions(request *GetSSETestRequest, runtime *util.RuntimeOptions) (_result *GetSSETestResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Params)) {
		body["Params"] = request.Params
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		body["ServiceCode"] = request.ServiceCode
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSSETest"),
		Version:     tea.String("2020-06-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSSETestResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSSETest(request *GetSSETestRequest) (_result *GetSSETestResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSSETestResponse{}
	_body, _err := client.GetSSETestWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSaChGeneralWithOptions(request *GetSaChGeneralRequest, runtime *util.RuntimeOptions) (_result *GetSaChGeneralResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		body["ServiceCode"] = request.ServiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.Text)) {
		body["Text"] = request.Text
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSaChGeneral"),
		Version:     tea.String("2020-06-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSaChGeneralResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSaChGeneral(request *GetSaChGeneralRequest) (_result *GetSaChGeneralResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSaChGeneralResponse{}
	_body, _err := client.GetSaChGeneralWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSaSeaEcomWithOptions(request *GetSaSeaEcomRequest, runtime *util.RuntimeOptions) (_result *GetSaSeaEcomResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Language)) {
		body["Language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		body["ServiceCode"] = request.ServiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.Text)) {
		body["Text"] = request.Text
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSaSeaEcom"),
		Version:     tea.String("2020-06-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSaSeaEcomResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSaSeaEcom(request *GetSaSeaEcomRequest) (_result *GetSaSeaEcomResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSaSeaEcomResponse{}
	_body, _err := client.GetSaSeaEcomWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetServiceDataImportStatusWithOptions(tmpReq *GetServiceDataImportStatusRequest, runtime *util.RuntimeOptions) (_result *GetServiceDataImportStatusResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetServiceDataImportStatusShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.DataImportIds)) {
		request.DataImportIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DataImportIds, tea.String("DataImportIds"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DataImportIdsShrink)) {
		body["DataImportIds"] = request.DataImportIdsShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetServiceDataImportStatus"),
		Version:     tea.String("2020-06-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetServiceDataImportStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetServiceDataImportStatus(request *GetServiceDataImportStatusRequest) (_result *GetServiceDataImportStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetServiceDataImportStatusResponse{}
	_body, _err := client.GetServiceDataImportStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSimilarityChMedicalWithOptions(request *GetSimilarityChMedicalRequest, runtime *util.RuntimeOptions) (_result *GetSimilarityChMedicalResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OriginQ)) {
		body["OriginQ"] = request.OriginQ
	}

	if !tea.BoolValue(util.IsUnset(request.OriginT)) {
		body["OriginT"] = request.OriginT
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		body["ServiceCode"] = request.ServiceCode
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSimilarityChMedical"),
		Version:     tea.String("2020-06-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSimilarityChMedicalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSimilarityChMedical(request *GetSimilarityChMedicalRequest) (_result *GetSimilarityChMedicalResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSimilarityChMedicalResponse{}
	_body, _err := client.GetSimilarityChMedicalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSummaryChEcomWithOptions(request *GetSummaryChEcomRequest, runtime *util.RuntimeOptions) (_result *GetSummaryChEcomResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		body["ServiceCode"] = request.ServiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.Text)) {
		body["Text"] = request.Text
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSummaryChEcom"),
		Version:     tea.String("2020-06-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSummaryChEcomResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSummaryChEcom(request *GetSummaryChEcomRequest) (_result *GetSummaryChEcomResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSummaryChEcomResponse{}
	_body, _err := client.GetSummaryChEcomWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTableQAServiceInfoByIdWithOptions(request *GetTableQAServiceInfoByIdRequest, runtime *util.RuntimeOptions) (_result *GetTableQAServiceInfoByIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		body["ServiceCode"] = request.ServiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		body["ServiceId"] = request.ServiceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTableQAServiceInfoById"),
		Version:     tea.String("2020-06-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTableQAServiceInfoByIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTableQAServiceInfoById(request *GetTableQAServiceInfoByIdRequest) (_result *GetTableQAServiceInfoByIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTableQAServiceInfoByIdResponse{}
	_body, _err := client.GetTableQAServiceInfoByIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTcChEcomWithOptions(request *GetTcChEcomRequest, runtime *util.RuntimeOptions) (_result *GetTcChEcomResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		body["ServiceCode"] = request.ServiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.Text)) {
		body["Text"] = request.Text
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTcChEcom"),
		Version:     tea.String("2020-06-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTcChEcomResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTcChEcom(request *GetTcChEcomRequest) (_result *GetTcChEcomResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTcChEcomResponse{}
	_body, _err := client.GetTcChEcomWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTcChGeneralWithOptions(request *GetTcChGeneralRequest, runtime *util.RuntimeOptions) (_result *GetTcChGeneralResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		body["ServiceCode"] = request.ServiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.Text)) {
		body["Text"] = request.Text
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTcChGeneral"),
		Version:     tea.String("2020-06-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTcChGeneralResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTcChGeneral(request *GetTcChGeneralRequest) (_result *GetTcChGeneralResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTcChGeneralResponse{}
	_body, _err := client.GetTcChGeneralWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTsChEcomWithOptions(request *GetTsChEcomRequest, runtime *util.RuntimeOptions) (_result *GetTsChEcomResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OriginQ)) {
		body["OriginQ"] = request.OriginQ
	}

	if !tea.BoolValue(util.IsUnset(request.OriginT)) {
		body["OriginT"] = request.OriginT
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		body["ServiceCode"] = request.ServiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTsChEcom"),
		Version:     tea.String("2020-06-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTsChEcomResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTsChEcom(request *GetTsChEcomRequest) (_result *GetTsChEcomResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTsChEcomResponse{}
	_body, _err := client.GetTsChEcomWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUserUploadSignWithOptions(request *GetUserUploadSignRequest, runtime *util.RuntimeOptions) (_result *GetUserUploadSignResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		body["ServiceCode"] = request.ServiceCode
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUserUploadSign"),
		Version:     tea.String("2020-06-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserUploadSignResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUserUploadSign(request *GetUserUploadSignRequest) (_result *GetUserUploadSignResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUserUploadSignResponse{}
	_body, _err := client.GetUserUploadSignWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetWeChCommentWithOptions(request *GetWeChCommentRequest, runtime *util.RuntimeOptions) (_result *GetWeChCommentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Operation)) {
		body["Operation"] = request.Operation
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		body["ServiceCode"] = request.ServiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		body["Size"] = request.Size
	}

	if !tea.BoolValue(util.IsUnset(request.Text)) {
		body["Text"] = request.Text
	}

	if !tea.BoolValue(util.IsUnset(request.TokenizerId)) {
		body["TokenizerId"] = request.TokenizerId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetWeChComment"),
		Version:     tea.String("2020-06-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetWeChCommentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetWeChComment(request *GetWeChCommentRequest) (_result *GetWeChCommentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetWeChCommentResponse{}
	_body, _err := client.GetWeChCommentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetWeChEcomWithOptions(request *GetWeChEcomRequest, runtime *util.RuntimeOptions) (_result *GetWeChEcomResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Operation)) {
		body["Operation"] = request.Operation
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		body["ServiceCode"] = request.ServiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		body["Size"] = request.Size
	}

	if !tea.BoolValue(util.IsUnset(request.Text)) {
		body["Text"] = request.Text
	}

	if !tea.BoolValue(util.IsUnset(request.TokenizerId)) {
		body["TokenizerId"] = request.TokenizerId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetWeChEcom"),
		Version:     tea.String("2020-06-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetWeChEcomResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetWeChEcom(request *GetWeChEcomRequest) (_result *GetWeChEcomResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetWeChEcomResponse{}
	_body, _err := client.GetWeChEcomWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetWeChEntertainmentWithOptions(request *GetWeChEntertainmentRequest, runtime *util.RuntimeOptions) (_result *GetWeChEntertainmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Operation)) {
		body["Operation"] = request.Operation
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		body["ServiceCode"] = request.ServiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		body["Size"] = request.Size
	}

	if !tea.BoolValue(util.IsUnset(request.Text)) {
		body["Text"] = request.Text
	}

	if !tea.BoolValue(util.IsUnset(request.TokenizerId)) {
		body["TokenizerId"] = request.TokenizerId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetWeChEntertainment"),
		Version:     tea.String("2020-06-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetWeChEntertainmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetWeChEntertainment(request *GetWeChEntertainmentRequest) (_result *GetWeChEntertainmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetWeChEntertainmentResponse{}
	_body, _err := client.GetWeChEntertainmentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetWeChGeneralWithOptions(request *GetWeChGeneralRequest, runtime *util.RuntimeOptions) (_result *GetWeChGeneralResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Operation)) {
		body["Operation"] = request.Operation
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		body["ServiceCode"] = request.ServiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		body["Size"] = request.Size
	}

	if !tea.BoolValue(util.IsUnset(request.Text)) {
		body["Text"] = request.Text
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetWeChGeneral"),
		Version:     tea.String("2020-06-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetWeChGeneralResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetWeChGeneral(request *GetWeChGeneralRequest) (_result *GetWeChGeneralResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetWeChGeneralResponse{}
	_body, _err := client.GetWeChGeneralWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetWeChSearchWithOptions(request *GetWeChSearchRequest, runtime *util.RuntimeOptions) (_result *GetWeChSearchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Operation)) {
		body["Operation"] = request.Operation
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		body["ServiceCode"] = request.ServiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		body["Size"] = request.Size
	}

	if !tea.BoolValue(util.IsUnset(request.Text)) {
		body["Text"] = request.Text
	}

	if !tea.BoolValue(util.IsUnset(request.TokenizerId)) {
		body["TokenizerId"] = request.TokenizerId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetWeChSearch"),
		Version:     tea.String("2020-06-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetWeChSearchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetWeChSearch(request *GetWeChSearchRequest) (_result *GetWeChSearchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetWeChSearchResponse{}
	_body, _err := client.GetWeChSearchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetWsChGeneralWithOptions(request *GetWsChGeneralRequest, runtime *util.RuntimeOptions) (_result *GetWsChGeneralResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OutType)) {
		body["OutType"] = request.OutType
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		body["ServiceCode"] = request.ServiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.Text)) {
		body["Text"] = request.Text
	}

	if !tea.BoolValue(util.IsUnset(request.TokenizerId)) {
		body["TokenizerId"] = request.TokenizerId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetWsChGeneral"),
		Version:     tea.String("2020-06-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetWsChGeneralResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetWsChGeneral(request *GetWsChGeneralRequest) (_result *GetWsChGeneralResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetWsChGeneralResponse{}
	_body, _err := client.GetWsChGeneralWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetWsCustomizedChEcomCommentWithOptions(request *GetWsCustomizedChEcomCommentRequest, runtime *util.RuntimeOptions) (_result *GetWsCustomizedChEcomCommentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OutType)) {
		body["OutType"] = request.OutType
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		body["ServiceCode"] = request.ServiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.Text)) {
		body["Text"] = request.Text
	}

	if !tea.BoolValue(util.IsUnset(request.TokenizerId)) {
		body["TokenizerId"] = request.TokenizerId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetWsCustomizedChEcomComment"),
		Version:     tea.String("2020-06-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetWsCustomizedChEcomCommentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetWsCustomizedChEcomComment(request *GetWsCustomizedChEcomCommentRequest) (_result *GetWsCustomizedChEcomCommentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetWsCustomizedChEcomCommentResponse{}
	_body, _err := client.GetWsCustomizedChEcomCommentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetWsCustomizedChEcomContentWithOptions(request *GetWsCustomizedChEcomContentRequest, runtime *util.RuntimeOptions) (_result *GetWsCustomizedChEcomContentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OutType)) {
		body["OutType"] = request.OutType
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		body["ServiceCode"] = request.ServiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.Text)) {
		body["Text"] = request.Text
	}

	if !tea.BoolValue(util.IsUnset(request.TokenizerId)) {
		body["TokenizerId"] = request.TokenizerId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetWsCustomizedChEcomContent"),
		Version:     tea.String("2020-06-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetWsCustomizedChEcomContentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetWsCustomizedChEcomContent(request *GetWsCustomizedChEcomContentRequest) (_result *GetWsCustomizedChEcomContentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetWsCustomizedChEcomContentResponse{}
	_body, _err := client.GetWsCustomizedChEcomContentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetWsCustomizedChEcomTitleWithOptions(request *GetWsCustomizedChEcomTitleRequest, runtime *util.RuntimeOptions) (_result *GetWsCustomizedChEcomTitleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OutType)) {
		body["OutType"] = request.OutType
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		body["ServiceCode"] = request.ServiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.Text)) {
		body["Text"] = request.Text
	}

	if !tea.BoolValue(util.IsUnset(request.TokenizerId)) {
		body["TokenizerId"] = request.TokenizerId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetWsCustomizedChEcomTitle"),
		Version:     tea.String("2020-06-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetWsCustomizedChEcomTitleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetWsCustomizedChEcomTitle(request *GetWsCustomizedChEcomTitleRequest) (_result *GetWsCustomizedChEcomTitleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetWsCustomizedChEcomTitleResponse{}
	_body, _err := client.GetWsCustomizedChEcomTitleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetWsCustomizedChEntertainmentWithOptions(request *GetWsCustomizedChEntertainmentRequest, runtime *util.RuntimeOptions) (_result *GetWsCustomizedChEntertainmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OutType)) {
		body["OutType"] = request.OutType
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		body["ServiceCode"] = request.ServiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.Text)) {
		body["Text"] = request.Text
	}

	if !tea.BoolValue(util.IsUnset(request.TokenizerId)) {
		body["TokenizerId"] = request.TokenizerId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetWsCustomizedChEntertainment"),
		Version:     tea.String("2020-06-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetWsCustomizedChEntertainmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetWsCustomizedChEntertainment(request *GetWsCustomizedChEntertainmentRequest) (_result *GetWsCustomizedChEntertainmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetWsCustomizedChEntertainmentResponse{}
	_body, _err := client.GetWsCustomizedChEntertainmentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetWsCustomizedChGeneralWithOptions(request *GetWsCustomizedChGeneralRequest, runtime *util.RuntimeOptions) (_result *GetWsCustomizedChGeneralResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OutType)) {
		body["OutType"] = request.OutType
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		body["ServiceCode"] = request.ServiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.Text)) {
		body["Text"] = request.Text
	}

	if !tea.BoolValue(util.IsUnset(request.TokenizerId)) {
		body["TokenizerId"] = request.TokenizerId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetWsCustomizedChGeneral"),
		Version:     tea.String("2020-06-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetWsCustomizedChGeneralResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetWsCustomizedChGeneral(request *GetWsCustomizedChGeneralRequest) (_result *GetWsCustomizedChGeneralResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetWsCustomizedChGeneralResponse{}
	_body, _err := client.GetWsCustomizedChGeneralWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetWsCustomizedChO2OWithOptions(request *GetWsCustomizedChO2ORequest, runtime *util.RuntimeOptions) (_result *GetWsCustomizedChO2OResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OutType)) {
		body["OutType"] = request.OutType
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		body["ServiceCode"] = request.ServiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.Text)) {
		body["Text"] = request.Text
	}

	if !tea.BoolValue(util.IsUnset(request.TokenizerId)) {
		body["TokenizerId"] = request.TokenizerId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetWsCustomizedChO2O"),
		Version:     tea.String("2020-06-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetWsCustomizedChO2OResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetWsCustomizedChO2O(request *GetWsCustomizedChO2ORequest) (_result *GetWsCustomizedChO2OResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetWsCustomizedChO2OResponse{}
	_body, _err := client.GetWsCustomizedChO2OWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetWsCustomizedSeaEcomWithOptions(request *GetWsCustomizedSeaEcomRequest, runtime *util.RuntimeOptions) (_result *GetWsCustomizedSeaEcomResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Language)) {
		body["Language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		body["ServiceCode"] = request.ServiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.Text)) {
		body["Text"] = request.Text
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetWsCustomizedSeaEcom"),
		Version:     tea.String("2020-06-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetWsCustomizedSeaEcomResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetWsCustomizedSeaEcom(request *GetWsCustomizedSeaEcomRequest) (_result *GetWsCustomizedSeaEcomResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetWsCustomizedSeaEcomResponse{}
	_body, _err := client.GetWsCustomizedSeaEcomWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetWsCustomizedSeaGeneralWithOptions(request *GetWsCustomizedSeaGeneralRequest, runtime *util.RuntimeOptions) (_result *GetWsCustomizedSeaGeneralResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Language)) {
		body["Language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		body["ServiceCode"] = request.ServiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.Text)) {
		body["Text"] = request.Text
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetWsCustomizedSeaGeneral"),
		Version:     tea.String("2020-06-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetWsCustomizedSeaGeneralResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetWsCustomizedSeaGeneral(request *GetWsCustomizedSeaGeneralRequest) (_result *GetWsCustomizedSeaGeneralResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetWsCustomizedSeaGeneralResponse{}
	_body, _err := client.GetWsCustomizedSeaGeneralWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ImportServiceDataWithOptions(tmpReq *ImportServiceDataRequest, runtime *util.RuntimeOptions) (_result *ImportServiceDataResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ImportServiceDataShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Partition)) {
		request.PartitionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Partition, tea.String("Partition"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PartitionShrink)) {
		body["Partition"] = request.PartitionShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		body["ServiceId"] = request.ServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.SubPath)) {
		body["SubPath"] = request.SubPath
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		body["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ImportServiceData"),
		Version:     tea.String("2020-06-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ImportServiceDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ImportServiceData(request *ImportServiceDataRequest) (_result *ImportServiceDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ImportServiceDataResponse{}
	_body, _err := client.ImportServiceDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ImportServiceDataV2WithOptions(tmpReq *ImportServiceDataV2Request, runtime *util.RuntimeOptions) (_result *ImportServiceDataV2Response, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ImportServiceDataV2ShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Documents)) {
		request.DocumentsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Documents, tea.String("Documents"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DataType)) {
		body["DataType"] = request.DataType
	}

	if !tea.BoolValue(util.IsUnset(request.DocumentsShrink)) {
		body["Documents"] = request.DocumentsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		body["ServiceId"] = request.ServiceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ImportServiceDataV2"),
		Version:     tea.String("2020-06-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ImportServiceDataV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ImportServiceDataV2(request *ImportServiceDataV2Request) (_result *ImportServiceDataV2Response, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ImportServiceDataV2Response{}
	_body, _err := client.ImportServiceDataV2WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InsertCustomWithOptions(request *InsertCustomRequest, runtime *util.RuntimeOptions) (_result *InsertCustomResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiId)) {
		body["ApiId"] = request.ApiId
	}

	if !tea.BoolValue(util.IsUnset(request.CustomFileName)) {
		body["CustomFileName"] = request.CustomFileName
	}

	if !tea.BoolValue(util.IsUnset(request.CustomUrl)) {
		body["CustomUrl"] = request.CustomUrl
	}

	if !tea.BoolValue(util.IsUnset(request.RegFileName)) {
		body["RegFileName"] = request.RegFileName
	}

	if !tea.BoolValue(util.IsUnset(request.RegUrl)) {
		body["RegUrl"] = request.RegUrl
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		body["ServiceCode"] = request.ServiceCode
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("InsertCustom"),
		Version:     tea.String("2020-06-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &InsertCustomResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InsertCustom(request *InsertCustomRequest) (_result *InsertCustomResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InsertCustomResponse{}
	_body, _err := client.InsertCustomWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OpenAlinlpServiceWithOptions(runtime *util.RuntimeOptions) (_result *OpenAlinlpServiceResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("OpenAlinlpService"),
		Version:     tea.String("2020-06-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &OpenAlinlpServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OpenAlinlpService() (_result *OpenAlinlpServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OpenAlinlpServiceResponse{}
	_body, _err := client.OpenAlinlpServiceWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PostISConvRewriterWithOptions(tmpReq *PostISConvRewriterRequest, runtime *util.RuntimeOptions) (_result *PostISConvRewriterResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &PostISConvRewriterShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Input)) {
		request.InputShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Input, tea.String("Input"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Parameters)) {
		request.ParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Parameters, tea.String("Parameters"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Algorithm)) {
		body["Algorithm"] = request.Algorithm
	}

	if !tea.BoolValue(util.IsUnset(request.Debug)) {
		body["Debug"] = request.Debug
	}

	if !tea.BoolValue(util.IsUnset(request.InputShrink)) {
		body["Input"] = request.InputShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ParametersShrink)) {
		body["Parameters"] = request.ParametersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Version)) {
		body["Version"] = request.Version
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PostISConvRewriter"),
		Version:     tea.String("2020-06-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PostISConvRewriterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PostISConvRewriter(request *PostISConvRewriterRequest) (_result *PostISConvRewriterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PostISConvRewriterResponse{}
	_body, _err := client.PostISConvRewriterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PostISRetrieveRouterWithOptions(tmpReq *PostISRetrieveRouterRequest, runtime *util.RuntimeOptions) (_result *PostISRetrieveRouterResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &PostISRetrieveRouterShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Input)) {
		request.InputShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Input, tea.String("Input"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Parameters)) {
		request.ParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Parameters, tea.String("Parameters"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Debug)) {
		query["Debug"] = request.Debug
	}

	if !tea.BoolValue(util.IsUnset(request.InputShrink)) {
		query["Input"] = request.InputShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ParametersShrink)) {
		query["Parameters"] = request.ParametersShrink
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Algorithm)) {
		body["Algorithm"] = request.Algorithm
	}

	if !tea.BoolValue(util.IsUnset(request.Version)) {
		body["Version"] = request.Version
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PostISRetrieveRouter"),
		Version:     tea.String("2020-06-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PostISRetrieveRouterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PostISRetrieveRouter(request *PostISRetrieveRouterRequest) (_result *PostISRetrieveRouterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PostISRetrieveRouterResponse{}
	_body, _err := client.PostISRetrieveRouterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PostMSConvSearchTokenGeneratedWithOptions(runtime *util.RuntimeOptions) (_result *PostMSConvSearchTokenGeneratedResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("PostMSConvSearchTokenGenerated"),
		Version:     tea.String("2020-06-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PostMSConvSearchTokenGeneratedResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PostMSConvSearchTokenGenerated() (_result *PostMSConvSearchTokenGeneratedResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PostMSConvSearchTokenGeneratedResponse{}
	_body, _err := client.PostMSConvSearchTokenGeneratedWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PostMSDataProcessingCountWithOptions(tmpReq *PostMSDataProcessingCountRequest, runtime *util.RuntimeOptions) (_result *PostMSDataProcessingCountResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &PostMSDataProcessingCountShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.DataIds)) {
		request.DataIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DataIds, tea.String("DataIds"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DataIdsShrink)) {
		body["DataIds"] = request.DataIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.DataImportId)) {
		body["DataImportId"] = request.DataImportId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		body["ServiceId"] = request.ServiceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PostMSDataProcessingCount"),
		Version:     tea.String("2020-06-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PostMSDataProcessingCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PostMSDataProcessingCount(request *PostMSDataProcessingCountRequest) (_result *PostMSDataProcessingCountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PostMSDataProcessingCountResponse{}
	_body, _err := client.PostMSDataProcessingCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PostMSSearchEnhanceWithOptions(tmpReq *PostMSSearchEnhanceRequest, runtime *util.RuntimeOptions) (_result *PostMSSearchEnhanceResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &PostMSSearchEnhanceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.CustomConfigInfo)) {
		request.CustomConfigInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CustomConfigInfo, tea.String("CustomConfigInfo"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Fields)) {
		request.FieldsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Fields, tea.String("Fields"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.RankModelInfo)) {
		request.RankModelInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RankModelInfo, tea.String("RankModelInfo"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Sort)) {
		request.SortShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Sort, tea.String("Sort"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Body)) {
		body["Body"] = request.Body
	}

	if !tea.BoolValue(util.IsUnset(request.CustomConfigInfoShrink)) {
		body["CustomConfigInfo"] = request.CustomConfigInfoShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Debug)) {
		body["Debug"] = request.Debug
	}

	if !tea.BoolValue(util.IsUnset(request.FieldsShrink)) {
		body["Fields"] = request.FieldsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Filters)) {
		body["Filters"] = request.Filters
	}

	if !tea.BoolValue(util.IsUnset(request.MinScore)) {
		body["MinScore"] = request.MinScore
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		body["Page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.Queries)) {
		body["Queries"] = request.Queries
	}

	if !tea.BoolValue(util.IsUnset(request.RankModelInfoShrink)) {
		body["RankModelInfo"] = request.RankModelInfoShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Rows)) {
		body["Rows"] = request.Rows
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		body["ServiceId"] = request.ServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.SortShrink)) {
		body["Sort"] = request.SortShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.Uq)) {
		body["Uq"] = request.Uq
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PostMSSearchEnhance"),
		Version:     tea.String("2020-06-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PostMSSearchEnhanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PostMSSearchEnhance(request *PostMSSearchEnhanceRequest) (_result *PostMSSearchEnhanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PostMSSearchEnhanceResponse{}
	_body, _err := client.PostMSSearchEnhanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PostMSServiceDataImportWithOptions(tmpReq *PostMSServiceDataImportRequest, runtime *util.RuntimeOptions) (_result *PostMSServiceDataImportResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &PostMSServiceDataImportShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Documents)) {
		request.DocumentsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Documents, tea.String("Documents"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DataType)) {
		body["DataType"] = request.DataType
	}

	if !tea.BoolValue(util.IsUnset(request.DocumentsShrink)) {
		body["Documents"] = request.DocumentsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		body["ServiceId"] = request.ServiceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PostMSServiceDataImport"),
		Version:     tea.String("2020-06-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PostMSServiceDataImportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PostMSServiceDataImport(request *PostMSServiceDataImportRequest) (_result *PostMSServiceDataImportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PostMSServiceDataImportResponse{}
	_body, _err := client.PostMSServiceDataImportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RequestTableQAWithOptions(request *RequestTableQARequest, runtime *util.RuntimeOptions) (_result *RequestTableQAResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Params)) {
		body["Params"] = request.Params
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		body["ServiceCode"] = request.ServiceCode
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RequestTableQA"),
		Version:     tea.String("2020-06-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RequestTableQAResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RequestTableQA(request *RequestTableQARequest) (_result *RequestTableQAResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RequestTableQAResponse{}
	_body, _err := client.RequestTableQAWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RequestTableQAOnlineWithOptions(request *RequestTableQAOnlineRequest, runtime *util.RuntimeOptions) (_result *RequestTableQAOnlineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BotId)) {
		body["BotId"] = request.BotId
	}

	if !tea.BoolValue(util.IsUnset(request.Params)) {
		body["Params"] = request.Params
	}

	if !tea.BoolValue(util.IsUnset(request.Question)) {
		body["Question"] = request.Question
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		body["ServiceCode"] = request.ServiceCode
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RequestTableQAOnline"),
		Version:     tea.String("2020-06-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RequestTableQAOnlineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RequestTableQAOnline(request *RequestTableQAOnlineRequest) (_result *RequestTableQAOnlineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RequestTableQAOnlineResponse{}
	_body, _err := client.RequestTableQAOnlineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateServiceDataWithOptions(tmpReq *UpdateServiceDataRequest, runtime *util.RuntimeOptions) (_result *UpdateServiceDataResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateServiceDataShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Conditions)) {
		request.ConditionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Conditions, tea.String("Conditions"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConditionsShrink)) {
		body["Conditions"] = request.ConditionsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		body["ServiceId"] = request.ServiceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateServiceData"),
		Version:     tea.String("2020-06-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateServiceDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateServiceData(request *UpdateServiceDataRequest) (_result *UpdateServiceDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateServiceDataResponse{}
	_body, _err := client.UpdateServiceDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
