// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type DescribeEmptyNumberDetectRequest struct {
	EncryptType          *string `json:"EncryptType,omitempty" xml:"EncryptType,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Phone                *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeEmptyNumberDetectRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEmptyNumberDetectRequest) GoString() string {
	return s.String()
}

func (s *DescribeEmptyNumberDetectRequest) SetEncryptType(v string) *DescribeEmptyNumberDetectRequest {
	s.EncryptType = &v
	return s
}

func (s *DescribeEmptyNumberDetectRequest) SetOwnerId(v int64) *DescribeEmptyNumberDetectRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeEmptyNumberDetectRequest) SetPhone(v string) *DescribeEmptyNumberDetectRequest {
	s.Phone = &v
	return s
}

func (s *DescribeEmptyNumberDetectRequest) SetResourceOwnerAccount(v string) *DescribeEmptyNumberDetectRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeEmptyNumberDetectRequest) SetResourceOwnerId(v int64) *DescribeEmptyNumberDetectRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeEmptyNumberDetectResponseBody struct {
	Code      *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*DescribeEmptyNumberDetectResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeEmptyNumberDetectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEmptyNumberDetectResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEmptyNumberDetectResponseBody) SetCode(v string) *DescribeEmptyNumberDetectResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeEmptyNumberDetectResponseBody) SetData(v []*DescribeEmptyNumberDetectResponseBodyData) *DescribeEmptyNumberDetectResponseBody {
	s.Data = v
	return s
}

func (s *DescribeEmptyNumberDetectResponseBody) SetMessage(v string) *DescribeEmptyNumberDetectResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeEmptyNumberDetectResponseBody) SetRequestId(v string) *DescribeEmptyNumberDetectResponseBody {
	s.RequestId = &v
	return s
}

type DescribeEmptyNumberDetectResponseBodyData struct {
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeEmptyNumberDetectResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeEmptyNumberDetectResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeEmptyNumberDetectResponseBodyData) SetNumber(v string) *DescribeEmptyNumberDetectResponseBodyData {
	s.Number = &v
	return s
}

func (s *DescribeEmptyNumberDetectResponseBodyData) SetStatus(v string) *DescribeEmptyNumberDetectResponseBodyData {
	s.Status = &v
	return s
}

type DescribeEmptyNumberDetectResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeEmptyNumberDetectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeEmptyNumberDetectResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEmptyNumberDetectResponse) GoString() string {
	return s.String()
}

func (s *DescribeEmptyNumberDetectResponse) SetHeaders(v map[string]*string) *DescribeEmptyNumberDetectResponse {
	s.Headers = v
	return s
}

func (s *DescribeEmptyNumberDetectResponse) SetStatusCode(v int32) *DescribeEmptyNumberDetectResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEmptyNumberDetectResponse) SetBody(v *DescribeEmptyNumberDetectResponseBody) *DescribeEmptyNumberDetectResponse {
	s.Body = v
	return s
}

type DescribePhoneNumberAnalysisRequest struct {
	AuthCode             *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	InputNumber          *string `json:"InputNumber,omitempty" xml:"InputNumber,omitempty"`
	Mask                 *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	NumberType           *int32  `json:"NumberType,omitempty" xml:"NumberType,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Rate                 *int64  `json:"Rate,omitempty" xml:"Rate,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribePhoneNumberAnalysisRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberAnalysisRequest) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberAnalysisRequest) SetAuthCode(v string) *DescribePhoneNumberAnalysisRequest {
	s.AuthCode = &v
	return s
}

func (s *DescribePhoneNumberAnalysisRequest) SetInputNumber(v string) *DescribePhoneNumberAnalysisRequest {
	s.InputNumber = &v
	return s
}

func (s *DescribePhoneNumberAnalysisRequest) SetMask(v string) *DescribePhoneNumberAnalysisRequest {
	s.Mask = &v
	return s
}

func (s *DescribePhoneNumberAnalysisRequest) SetNumberType(v int32) *DescribePhoneNumberAnalysisRequest {
	s.NumberType = &v
	return s
}

func (s *DescribePhoneNumberAnalysisRequest) SetOwnerId(v int64) *DescribePhoneNumberAnalysisRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribePhoneNumberAnalysisRequest) SetRate(v int64) *DescribePhoneNumberAnalysisRequest {
	s.Rate = &v
	return s
}

func (s *DescribePhoneNumberAnalysisRequest) SetResourceOwnerAccount(v string) *DescribePhoneNumberAnalysisRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribePhoneNumberAnalysisRequest) SetResourceOwnerId(v int64) *DescribePhoneNumberAnalysisRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribePhoneNumberAnalysisResponseBody struct {
	Code      *string                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*DescribePhoneNumberAnalysisResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePhoneNumberAnalysisResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberAnalysisResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberAnalysisResponseBody) SetCode(v string) *DescribePhoneNumberAnalysisResponseBody {
	s.Code = &v
	return s
}

func (s *DescribePhoneNumberAnalysisResponseBody) SetData(v []*DescribePhoneNumberAnalysisResponseBodyData) *DescribePhoneNumberAnalysisResponseBody {
	s.Data = v
	return s
}

func (s *DescribePhoneNumberAnalysisResponseBody) SetMessage(v string) *DescribePhoneNumberAnalysisResponseBody {
	s.Message = &v
	return s
}

func (s *DescribePhoneNumberAnalysisResponseBody) SetRequestId(v string) *DescribePhoneNumberAnalysisResponseBody {
	s.RequestId = &v
	return s
}

type DescribePhoneNumberAnalysisResponseBodyData struct {
	Code   *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
}

func (s DescribePhoneNumberAnalysisResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberAnalysisResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberAnalysisResponseBodyData) SetCode(v string) *DescribePhoneNumberAnalysisResponseBodyData {
	s.Code = &v
	return s
}

func (s *DescribePhoneNumberAnalysisResponseBodyData) SetNumber(v string) *DescribePhoneNumberAnalysisResponseBodyData {
	s.Number = &v
	return s
}

type DescribePhoneNumberAnalysisResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribePhoneNumberAnalysisResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePhoneNumberAnalysisResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberAnalysisResponse) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberAnalysisResponse) SetHeaders(v map[string]*string) *DescribePhoneNumberAnalysisResponse {
	s.Headers = v
	return s
}

func (s *DescribePhoneNumberAnalysisResponse) SetStatusCode(v int32) *DescribePhoneNumberAnalysisResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePhoneNumberAnalysisResponse) SetBody(v *DescribePhoneNumberAnalysisResponseBody) *DescribePhoneNumberAnalysisResponse {
	s.Body = v
	return s
}

type DescribePhoneNumberAttributeRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PhoneNumber          *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribePhoneNumberAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberAttributeRequest) SetOwnerId(v int64) *DescribePhoneNumberAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribePhoneNumberAttributeRequest) SetPhoneNumber(v string) *DescribePhoneNumberAttributeRequest {
	s.PhoneNumber = &v
	return s
}

func (s *DescribePhoneNumberAttributeRequest) SetResourceOwnerAccount(v string) *DescribePhoneNumberAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribePhoneNumberAttributeRequest) SetResourceOwnerId(v int64) *DescribePhoneNumberAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribePhoneNumberAttributeResponseBody struct {
	Code                 *string                                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Message              *string                                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	PhoneNumberAttribute *DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute `json:"PhoneNumberAttribute,omitempty" xml:"PhoneNumberAttribute,omitempty" type:"Struct"`
	RequestId            *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePhoneNumberAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberAttributeResponseBody) SetCode(v string) *DescribePhoneNumberAttributeResponseBody {
	s.Code = &v
	return s
}

func (s *DescribePhoneNumberAttributeResponseBody) SetMessage(v string) *DescribePhoneNumberAttributeResponseBody {
	s.Message = &v
	return s
}

func (s *DescribePhoneNumberAttributeResponseBody) SetPhoneNumberAttribute(v *DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute) *DescribePhoneNumberAttributeResponseBody {
	s.PhoneNumberAttribute = v
	return s
}

func (s *DescribePhoneNumberAttributeResponseBody) SetRequestId(v string) *DescribePhoneNumberAttributeResponseBody {
	s.RequestId = &v
	return s
}

type DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute struct {
	BasicCarrier        *string `json:"BasicCarrier,omitempty" xml:"BasicCarrier,omitempty"`
	Carrier             *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	City                *string `json:"City,omitempty" xml:"City,omitempty"`
	IsNumberPortability *bool   `json:"IsNumberPortability,omitempty" xml:"IsNumberPortability,omitempty"`
	NumberSegment       *int64  `json:"NumberSegment,omitempty" xml:"NumberSegment,omitempty"`
	Province            *string `json:"Province,omitempty" xml:"Province,omitempty"`
}

func (s DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute) SetBasicCarrier(v string) *DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute {
	s.BasicCarrier = &v
	return s
}

func (s *DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute) SetCarrier(v string) *DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute {
	s.Carrier = &v
	return s
}

func (s *DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute) SetCity(v string) *DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute {
	s.City = &v
	return s
}

func (s *DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute) SetIsNumberPortability(v bool) *DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute {
	s.IsNumberPortability = &v
	return s
}

func (s *DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute) SetNumberSegment(v int64) *DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute {
	s.NumberSegment = &v
	return s
}

func (s *DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute) SetProvince(v string) *DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute {
	s.Province = &v
	return s
}

type DescribePhoneNumberAttributeResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribePhoneNumberAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePhoneNumberAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberAttributeResponse) SetHeaders(v map[string]*string) *DescribePhoneNumberAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribePhoneNumberAttributeResponse) SetStatusCode(v int32) *DescribePhoneNumberAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePhoneNumberAttributeResponse) SetBody(v *DescribePhoneNumberAttributeResponseBody) *DescribePhoneNumberAttributeResponse {
	s.Body = v
	return s
}

type DescribePhoneNumberOnlineTimeRequest struct {
	AuthCode             *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	Carrier              *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	InputNumber          *string `json:"InputNumber,omitempty" xml:"InputNumber,omitempty"`
	Mask                 *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribePhoneNumberOnlineTimeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberOnlineTimeRequest) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberOnlineTimeRequest) SetAuthCode(v string) *DescribePhoneNumberOnlineTimeRequest {
	s.AuthCode = &v
	return s
}

func (s *DescribePhoneNumberOnlineTimeRequest) SetCarrier(v string) *DescribePhoneNumberOnlineTimeRequest {
	s.Carrier = &v
	return s
}

func (s *DescribePhoneNumberOnlineTimeRequest) SetInputNumber(v string) *DescribePhoneNumberOnlineTimeRequest {
	s.InputNumber = &v
	return s
}

func (s *DescribePhoneNumberOnlineTimeRequest) SetMask(v string) *DescribePhoneNumberOnlineTimeRequest {
	s.Mask = &v
	return s
}

func (s *DescribePhoneNumberOnlineTimeRequest) SetOwnerId(v int64) *DescribePhoneNumberOnlineTimeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribePhoneNumberOnlineTimeRequest) SetResourceOwnerAccount(v string) *DescribePhoneNumberOnlineTimeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribePhoneNumberOnlineTimeRequest) SetResourceOwnerId(v int64) *DescribePhoneNumberOnlineTimeRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribePhoneNumberOnlineTimeResponseBody struct {
	Code      *string                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DescribePhoneNumberOnlineTimeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePhoneNumberOnlineTimeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberOnlineTimeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberOnlineTimeResponseBody) SetCode(v string) *DescribePhoneNumberOnlineTimeResponseBody {
	s.Code = &v
	return s
}

func (s *DescribePhoneNumberOnlineTimeResponseBody) SetData(v *DescribePhoneNumberOnlineTimeResponseBodyData) *DescribePhoneNumberOnlineTimeResponseBody {
	s.Data = v
	return s
}

func (s *DescribePhoneNumberOnlineTimeResponseBody) SetMessage(v string) *DescribePhoneNumberOnlineTimeResponseBody {
	s.Message = &v
	return s
}

func (s *DescribePhoneNumberOnlineTimeResponseBody) SetRequestId(v string) *DescribePhoneNumberOnlineTimeResponseBody {
	s.RequestId = &v
	return s
}

type DescribePhoneNumberOnlineTimeResponseBodyData struct {
	VerifyResult *string `json:"VerifyResult,omitempty" xml:"VerifyResult,omitempty"`
}

func (s DescribePhoneNumberOnlineTimeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberOnlineTimeResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberOnlineTimeResponseBodyData) SetVerifyResult(v string) *DescribePhoneNumberOnlineTimeResponseBodyData {
	s.VerifyResult = &v
	return s
}

type DescribePhoneNumberOnlineTimeResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribePhoneNumberOnlineTimeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePhoneNumberOnlineTimeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberOnlineTimeResponse) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberOnlineTimeResponse) SetHeaders(v map[string]*string) *DescribePhoneNumberOnlineTimeResponse {
	s.Headers = v
	return s
}

func (s *DescribePhoneNumberOnlineTimeResponse) SetStatusCode(v int32) *DescribePhoneNumberOnlineTimeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePhoneNumberOnlineTimeResponse) SetBody(v *DescribePhoneNumberOnlineTimeResponseBody) *DescribePhoneNumberOnlineTimeResponse {
	s.Body = v
	return s
}

type DescribePhoneNumberOperatorAttributeRequest struct {
	AuthCode             *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	InputNumber          *string `json:"InputNumber,omitempty" xml:"InputNumber,omitempty"`
	Mask                 *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribePhoneNumberOperatorAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberOperatorAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberOperatorAttributeRequest) SetAuthCode(v string) *DescribePhoneNumberOperatorAttributeRequest {
	s.AuthCode = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeRequest) SetInputNumber(v string) *DescribePhoneNumberOperatorAttributeRequest {
	s.InputNumber = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeRequest) SetMask(v string) *DescribePhoneNumberOperatorAttributeRequest {
	s.Mask = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeRequest) SetOwnerId(v int64) *DescribePhoneNumberOperatorAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeRequest) SetResourceOwnerAccount(v string) *DescribePhoneNumberOperatorAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeRequest) SetResourceOwnerId(v int64) *DescribePhoneNumberOperatorAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribePhoneNumberOperatorAttributeResponseBody struct {
	Code      *string                                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DescribePhoneNumberOperatorAttributeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePhoneNumberOperatorAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberOperatorAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberOperatorAttributeResponseBody) SetCode(v string) *DescribePhoneNumberOperatorAttributeResponseBody {
	s.Code = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeResponseBody) SetData(v *DescribePhoneNumberOperatorAttributeResponseBodyData) *DescribePhoneNumberOperatorAttributeResponseBody {
	s.Data = v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeResponseBody) SetMessage(v string) *DescribePhoneNumberOperatorAttributeResponseBody {
	s.Message = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeResponseBody) SetRequestId(v string) *DescribePhoneNumberOperatorAttributeResponseBody {
	s.RequestId = &v
	return s
}

type DescribePhoneNumberOperatorAttributeResponseBodyData struct {
	BasicCarrier        *string `json:"BasicCarrier,omitempty" xml:"BasicCarrier,omitempty"`
	Carrier             *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	City                *string `json:"City,omitempty" xml:"City,omitempty"`
	IsNumberPortability *bool   `json:"IsNumberPortability,omitempty" xml:"IsNumberPortability,omitempty"`
	NumberSegment       *int64  `json:"NumberSegment,omitempty" xml:"NumberSegment,omitempty"`
	Province            *string `json:"Province,omitempty" xml:"Province,omitempty"`
}

func (s DescribePhoneNumberOperatorAttributeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberOperatorAttributeResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberOperatorAttributeResponseBodyData) SetBasicCarrier(v string) *DescribePhoneNumberOperatorAttributeResponseBodyData {
	s.BasicCarrier = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeResponseBodyData) SetCarrier(v string) *DescribePhoneNumberOperatorAttributeResponseBodyData {
	s.Carrier = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeResponseBodyData) SetCity(v string) *DescribePhoneNumberOperatorAttributeResponseBodyData {
	s.City = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeResponseBodyData) SetIsNumberPortability(v bool) *DescribePhoneNumberOperatorAttributeResponseBodyData {
	s.IsNumberPortability = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeResponseBodyData) SetNumberSegment(v int64) *DescribePhoneNumberOperatorAttributeResponseBodyData {
	s.NumberSegment = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeResponseBodyData) SetProvince(v string) *DescribePhoneNumberOperatorAttributeResponseBodyData {
	s.Province = &v
	return s
}

type DescribePhoneNumberOperatorAttributeResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribePhoneNumberOperatorAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePhoneNumberOperatorAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberOperatorAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberOperatorAttributeResponse) SetHeaders(v map[string]*string) *DescribePhoneNumberOperatorAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeResponse) SetStatusCode(v int32) *DescribePhoneNumberOperatorAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeResponse) SetBody(v *DescribePhoneNumberOperatorAttributeResponseBody) *DescribePhoneNumberOperatorAttributeResponse {
	s.Body = v
	return s
}

type DescribePhoneNumberResaleRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PhoneNumber          *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Since                *string `json:"Since,omitempty" xml:"Since,omitempty"`
}

func (s DescribePhoneNumberResaleRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberResaleRequest) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberResaleRequest) SetOwnerId(v int64) *DescribePhoneNumberResaleRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribePhoneNumberResaleRequest) SetPhoneNumber(v string) *DescribePhoneNumberResaleRequest {
	s.PhoneNumber = &v
	return s
}

func (s *DescribePhoneNumberResaleRequest) SetResourceOwnerAccount(v string) *DescribePhoneNumberResaleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribePhoneNumberResaleRequest) SetResourceOwnerId(v int64) *DescribePhoneNumberResaleRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribePhoneNumberResaleRequest) SetSince(v string) *DescribePhoneNumberResaleRequest {
	s.Since = &v
	return s
}

type DescribePhoneNumberResaleResponseBody struct {
	Code           *string                                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Message        *string                                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TwiceTelVerify *DescribePhoneNumberResaleResponseBodyTwiceTelVerify `json:"TwiceTelVerify,omitempty" xml:"TwiceTelVerify,omitempty" type:"Struct"`
}

func (s DescribePhoneNumberResaleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberResaleResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberResaleResponseBody) SetCode(v string) *DescribePhoneNumberResaleResponseBody {
	s.Code = &v
	return s
}

func (s *DescribePhoneNumberResaleResponseBody) SetMessage(v string) *DescribePhoneNumberResaleResponseBody {
	s.Message = &v
	return s
}

func (s *DescribePhoneNumberResaleResponseBody) SetRequestId(v string) *DescribePhoneNumberResaleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePhoneNumberResaleResponseBody) SetTwiceTelVerify(v *DescribePhoneNumberResaleResponseBodyTwiceTelVerify) *DescribePhoneNumberResaleResponseBody {
	s.TwiceTelVerify = v
	return s
}

type DescribePhoneNumberResaleResponseBodyTwiceTelVerify struct {
	Carrier      *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	VerifyResult *int32  `json:"VerifyResult,omitempty" xml:"VerifyResult,omitempty"`
}

func (s DescribePhoneNumberResaleResponseBodyTwiceTelVerify) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberResaleResponseBodyTwiceTelVerify) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberResaleResponseBodyTwiceTelVerify) SetCarrier(v string) *DescribePhoneNumberResaleResponseBodyTwiceTelVerify {
	s.Carrier = &v
	return s
}

func (s *DescribePhoneNumberResaleResponseBodyTwiceTelVerify) SetVerifyResult(v int32) *DescribePhoneNumberResaleResponseBodyTwiceTelVerify {
	s.VerifyResult = &v
	return s
}

type DescribePhoneNumberResaleResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribePhoneNumberResaleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePhoneNumberResaleResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberResaleResponse) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberResaleResponse) SetHeaders(v map[string]*string) *DescribePhoneNumberResaleResponse {
	s.Headers = v
	return s
}

func (s *DescribePhoneNumberResaleResponse) SetStatusCode(v int32) *DescribePhoneNumberResaleResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePhoneNumberResaleResponse) SetBody(v *DescribePhoneNumberResaleResponseBody) *DescribePhoneNumberResaleResponse {
	s.Body = v
	return s
}

type DescribePhoneNumberStatusRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PhoneNumber          *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribePhoneNumberStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberStatusRequest) SetOwnerId(v int64) *DescribePhoneNumberStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribePhoneNumberStatusRequest) SetPhoneNumber(v string) *DescribePhoneNumberStatusRequest {
	s.PhoneNumber = &v
	return s
}

func (s *DescribePhoneNumberStatusRequest) SetResourceOwnerAccount(v string) *DescribePhoneNumberStatusRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribePhoneNumberStatusRequest) SetResourceOwnerId(v int64) *DescribePhoneNumberStatusRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribePhoneNumberStatusResponseBody struct {
	Code        *string                                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Message     *string                                           `json:"Message,omitempty" xml:"Message,omitempty"`
	PhoneStatus *DescribePhoneNumberStatusResponseBodyPhoneStatus `json:"PhoneStatus,omitempty" xml:"PhoneStatus,omitempty" type:"Struct"`
	RequestId   *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePhoneNumberStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberStatusResponseBody) SetCode(v string) *DescribePhoneNumberStatusResponseBody {
	s.Code = &v
	return s
}

func (s *DescribePhoneNumberStatusResponseBody) SetMessage(v string) *DescribePhoneNumberStatusResponseBody {
	s.Message = &v
	return s
}

func (s *DescribePhoneNumberStatusResponseBody) SetPhoneStatus(v *DescribePhoneNumberStatusResponseBodyPhoneStatus) *DescribePhoneNumberStatusResponseBody {
	s.PhoneStatus = v
	return s
}

func (s *DescribePhoneNumberStatusResponseBody) SetRequestId(v string) *DescribePhoneNumberStatusResponseBody {
	s.RequestId = &v
	return s
}

type DescribePhoneNumberStatusResponseBodyPhoneStatus struct {
	Carrier  *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	SerialId *string `json:"SerialId,omitempty" xml:"SerialId,omitempty"`
	Status   *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribePhoneNumberStatusResponseBodyPhoneStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberStatusResponseBodyPhoneStatus) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberStatusResponseBodyPhoneStatus) SetCarrier(v string) *DescribePhoneNumberStatusResponseBodyPhoneStatus {
	s.Carrier = &v
	return s
}

func (s *DescribePhoneNumberStatusResponseBodyPhoneStatus) SetSerialId(v string) *DescribePhoneNumberStatusResponseBodyPhoneStatus {
	s.SerialId = &v
	return s
}

func (s *DescribePhoneNumberStatusResponseBodyPhoneStatus) SetStatus(v string) *DescribePhoneNumberStatusResponseBodyPhoneStatus {
	s.Status = &v
	return s
}

type DescribePhoneNumberStatusResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribePhoneNumberStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePhoneNumberStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberStatusResponse) SetHeaders(v map[string]*string) *DescribePhoneNumberStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribePhoneNumberStatusResponse) SetStatusCode(v int32) *DescribePhoneNumberStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePhoneNumberStatusResponse) SetBody(v *DescribePhoneNumberStatusResponseBody) *DescribePhoneNumberStatusResponse {
	s.Body = v
	return s
}

type DescribePhoneTwiceTelVerifyRequest struct {
	AuthCode             *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	InputNumber          *string `json:"InputNumber,omitempty" xml:"InputNumber,omitempty"`
	Mask                 *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribePhoneTwiceTelVerifyRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneTwiceTelVerifyRequest) GoString() string {
	return s.String()
}

func (s *DescribePhoneTwiceTelVerifyRequest) SetAuthCode(v string) *DescribePhoneTwiceTelVerifyRequest {
	s.AuthCode = &v
	return s
}

func (s *DescribePhoneTwiceTelVerifyRequest) SetInputNumber(v string) *DescribePhoneTwiceTelVerifyRequest {
	s.InputNumber = &v
	return s
}

func (s *DescribePhoneTwiceTelVerifyRequest) SetMask(v string) *DescribePhoneTwiceTelVerifyRequest {
	s.Mask = &v
	return s
}

func (s *DescribePhoneTwiceTelVerifyRequest) SetOwnerId(v int64) *DescribePhoneTwiceTelVerifyRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribePhoneTwiceTelVerifyRequest) SetResourceOwnerAccount(v string) *DescribePhoneTwiceTelVerifyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribePhoneTwiceTelVerifyRequest) SetResourceOwnerId(v int64) *DescribePhoneTwiceTelVerifyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribePhoneTwiceTelVerifyRequest) SetStartTime(v string) *DescribePhoneTwiceTelVerifyRequest {
	s.StartTime = &v
	return s
}

type DescribePhoneTwiceTelVerifyResponseBody struct {
	Code      *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DescribePhoneTwiceTelVerifyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePhoneTwiceTelVerifyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneTwiceTelVerifyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePhoneTwiceTelVerifyResponseBody) SetCode(v string) *DescribePhoneTwiceTelVerifyResponseBody {
	s.Code = &v
	return s
}

func (s *DescribePhoneTwiceTelVerifyResponseBody) SetData(v *DescribePhoneTwiceTelVerifyResponseBodyData) *DescribePhoneTwiceTelVerifyResponseBody {
	s.Data = v
	return s
}

func (s *DescribePhoneTwiceTelVerifyResponseBody) SetMessage(v string) *DescribePhoneTwiceTelVerifyResponseBody {
	s.Message = &v
	return s
}

func (s *DescribePhoneTwiceTelVerifyResponseBody) SetRequestId(v string) *DescribePhoneTwiceTelVerifyResponseBody {
	s.RequestId = &v
	return s
}

type DescribePhoneTwiceTelVerifyResponseBodyData struct {
	Carrier      *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	VerifyResult *string `json:"VerifyResult,omitempty" xml:"VerifyResult,omitempty"`
}

func (s DescribePhoneTwiceTelVerifyResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneTwiceTelVerifyResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribePhoneTwiceTelVerifyResponseBodyData) SetCarrier(v string) *DescribePhoneTwiceTelVerifyResponseBodyData {
	s.Carrier = &v
	return s
}

func (s *DescribePhoneTwiceTelVerifyResponseBodyData) SetVerifyResult(v string) *DescribePhoneTwiceTelVerifyResponseBodyData {
	s.VerifyResult = &v
	return s
}

type DescribePhoneTwiceTelVerifyResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribePhoneTwiceTelVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePhoneTwiceTelVerifyResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneTwiceTelVerifyResponse) GoString() string {
	return s.String()
}

func (s *DescribePhoneTwiceTelVerifyResponse) SetHeaders(v map[string]*string) *DescribePhoneTwiceTelVerifyResponse {
	s.Headers = v
	return s
}

func (s *DescribePhoneTwiceTelVerifyResponse) SetStatusCode(v int32) *DescribePhoneTwiceTelVerifyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePhoneTwiceTelVerifyResponse) SetBody(v *DescribePhoneTwiceTelVerifyResponseBody) *DescribePhoneTwiceTelVerifyResponse {
	s.Body = v
	return s
}

type InvalidPhoneNumberFilterRequest struct {
	AuthCode             *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	InputNumber          *string `json:"InputNumber,omitempty" xml:"InputNumber,omitempty"`
	Mask                 *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s InvalidPhoneNumberFilterRequest) String() string {
	return tea.Prettify(s)
}

func (s InvalidPhoneNumberFilterRequest) GoString() string {
	return s.String()
}

func (s *InvalidPhoneNumberFilterRequest) SetAuthCode(v string) *InvalidPhoneNumberFilterRequest {
	s.AuthCode = &v
	return s
}

func (s *InvalidPhoneNumberFilterRequest) SetInputNumber(v string) *InvalidPhoneNumberFilterRequest {
	s.InputNumber = &v
	return s
}

func (s *InvalidPhoneNumberFilterRequest) SetMask(v string) *InvalidPhoneNumberFilterRequest {
	s.Mask = &v
	return s
}

func (s *InvalidPhoneNumberFilterRequest) SetOwnerId(v int64) *InvalidPhoneNumberFilterRequest {
	s.OwnerId = &v
	return s
}

func (s *InvalidPhoneNumberFilterRequest) SetResourceOwnerAccount(v string) *InvalidPhoneNumberFilterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *InvalidPhoneNumberFilterRequest) SetResourceOwnerId(v int64) *InvalidPhoneNumberFilterRequest {
	s.ResourceOwnerId = &v
	return s
}

type InvalidPhoneNumberFilterResponseBody struct {
	Code      *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*InvalidPhoneNumberFilterResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InvalidPhoneNumberFilterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InvalidPhoneNumberFilterResponseBody) GoString() string {
	return s.String()
}

func (s *InvalidPhoneNumberFilterResponseBody) SetCode(v string) *InvalidPhoneNumberFilterResponseBody {
	s.Code = &v
	return s
}

func (s *InvalidPhoneNumberFilterResponseBody) SetData(v []*InvalidPhoneNumberFilterResponseBodyData) *InvalidPhoneNumberFilterResponseBody {
	s.Data = v
	return s
}

func (s *InvalidPhoneNumberFilterResponseBody) SetMessage(v string) *InvalidPhoneNumberFilterResponseBody {
	s.Message = &v
	return s
}

func (s *InvalidPhoneNumberFilterResponseBody) SetRequestId(v string) *InvalidPhoneNumberFilterResponseBody {
	s.RequestId = &v
	return s
}

type InvalidPhoneNumberFilterResponseBodyData struct {
	Code            *string `json:"Code,omitempty" xml:"Code,omitempty"`
	EncryptedNumber *string `json:"EncryptedNumber,omitempty" xml:"EncryptedNumber,omitempty"`
	ExpireTime      *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	OriginalNumber  *string `json:"OriginalNumber,omitempty" xml:"OriginalNumber,omitempty"`
}

func (s InvalidPhoneNumberFilterResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s InvalidPhoneNumberFilterResponseBodyData) GoString() string {
	return s.String()
}

func (s *InvalidPhoneNumberFilterResponseBodyData) SetCode(v string) *InvalidPhoneNumberFilterResponseBodyData {
	s.Code = &v
	return s
}

func (s *InvalidPhoneNumberFilterResponseBodyData) SetEncryptedNumber(v string) *InvalidPhoneNumberFilterResponseBodyData {
	s.EncryptedNumber = &v
	return s
}

func (s *InvalidPhoneNumberFilterResponseBodyData) SetExpireTime(v string) *InvalidPhoneNumberFilterResponseBodyData {
	s.ExpireTime = &v
	return s
}

func (s *InvalidPhoneNumberFilterResponseBodyData) SetOriginalNumber(v string) *InvalidPhoneNumberFilterResponseBodyData {
	s.OriginalNumber = &v
	return s
}

type InvalidPhoneNumberFilterResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *InvalidPhoneNumberFilterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InvalidPhoneNumberFilterResponse) String() string {
	return tea.Prettify(s)
}

func (s InvalidPhoneNumberFilterResponse) GoString() string {
	return s.String()
}

func (s *InvalidPhoneNumberFilterResponse) SetHeaders(v map[string]*string) *InvalidPhoneNumberFilterResponse {
	s.Headers = v
	return s
}

func (s *InvalidPhoneNumberFilterResponse) SetStatusCode(v int32) *InvalidPhoneNumberFilterResponse {
	s.StatusCode = &v
	return s
}

func (s *InvalidPhoneNumberFilterResponse) SetBody(v *InvalidPhoneNumberFilterResponseBody) *InvalidPhoneNumberFilterResponse {
	s.Body = v
	return s
}

type PhoneNumberEncryptRequest struct {
	AuthCode             *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	InputNumber          *string `json:"InputNumber,omitempty" xml:"InputNumber,omitempty"`
	Mask                 *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s PhoneNumberEncryptRequest) String() string {
	return tea.Prettify(s)
}

func (s PhoneNumberEncryptRequest) GoString() string {
	return s.String()
}

func (s *PhoneNumberEncryptRequest) SetAuthCode(v string) *PhoneNumberEncryptRequest {
	s.AuthCode = &v
	return s
}

func (s *PhoneNumberEncryptRequest) SetInputNumber(v string) *PhoneNumberEncryptRequest {
	s.InputNumber = &v
	return s
}

func (s *PhoneNumberEncryptRequest) SetMask(v string) *PhoneNumberEncryptRequest {
	s.Mask = &v
	return s
}

func (s *PhoneNumberEncryptRequest) SetOwnerId(v int64) *PhoneNumberEncryptRequest {
	s.OwnerId = &v
	return s
}

func (s *PhoneNumberEncryptRequest) SetResourceOwnerAccount(v string) *PhoneNumberEncryptRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *PhoneNumberEncryptRequest) SetResourceOwnerId(v int64) *PhoneNumberEncryptRequest {
	s.ResourceOwnerId = &v
	return s
}

type PhoneNumberEncryptResponseBody struct {
	Code      *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*PhoneNumberEncryptResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PhoneNumberEncryptResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PhoneNumberEncryptResponseBody) GoString() string {
	return s.String()
}

func (s *PhoneNumberEncryptResponseBody) SetCode(v string) *PhoneNumberEncryptResponseBody {
	s.Code = &v
	return s
}

func (s *PhoneNumberEncryptResponseBody) SetData(v []*PhoneNumberEncryptResponseBodyData) *PhoneNumberEncryptResponseBody {
	s.Data = v
	return s
}

func (s *PhoneNumberEncryptResponseBody) SetMessage(v string) *PhoneNumberEncryptResponseBody {
	s.Message = &v
	return s
}

func (s *PhoneNumberEncryptResponseBody) SetRequestId(v string) *PhoneNumberEncryptResponseBody {
	s.RequestId = &v
	return s
}

type PhoneNumberEncryptResponseBodyData struct {
	EncryptedNumber *string `json:"EncryptedNumber,omitempty" xml:"EncryptedNumber,omitempty"`
	ExpireTime      *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	OriginalNumber  *string `json:"OriginalNumber,omitempty" xml:"OriginalNumber,omitempty"`
}

func (s PhoneNumberEncryptResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PhoneNumberEncryptResponseBodyData) GoString() string {
	return s.String()
}

func (s *PhoneNumberEncryptResponseBodyData) SetEncryptedNumber(v string) *PhoneNumberEncryptResponseBodyData {
	s.EncryptedNumber = &v
	return s
}

func (s *PhoneNumberEncryptResponseBodyData) SetExpireTime(v string) *PhoneNumberEncryptResponseBodyData {
	s.ExpireTime = &v
	return s
}

func (s *PhoneNumberEncryptResponseBodyData) SetOriginalNumber(v string) *PhoneNumberEncryptResponseBodyData {
	s.OriginalNumber = &v
	return s
}

type PhoneNumberEncryptResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PhoneNumberEncryptResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PhoneNumberEncryptResponse) String() string {
	return tea.Prettify(s)
}

func (s PhoneNumberEncryptResponse) GoString() string {
	return s.String()
}

func (s *PhoneNumberEncryptResponse) SetHeaders(v map[string]*string) *PhoneNumberEncryptResponse {
	s.Headers = v
	return s
}

func (s *PhoneNumberEncryptResponse) SetStatusCode(v int32) *PhoneNumberEncryptResponse {
	s.StatusCode = &v
	return s
}

func (s *PhoneNumberEncryptResponse) SetBody(v *PhoneNumberEncryptResponseBody) *PhoneNumberEncryptResponse {
	s.Body = v
	return s
}

type PhoneNumberStatusForAccountRequest struct {
	AuthCode             *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	InputNumber          *string `json:"InputNumber,omitempty" xml:"InputNumber,omitempty"`
	Mask                 *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s PhoneNumberStatusForAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s PhoneNumberStatusForAccountRequest) GoString() string {
	return s.String()
}

func (s *PhoneNumberStatusForAccountRequest) SetAuthCode(v string) *PhoneNumberStatusForAccountRequest {
	s.AuthCode = &v
	return s
}

func (s *PhoneNumberStatusForAccountRequest) SetInputNumber(v string) *PhoneNumberStatusForAccountRequest {
	s.InputNumber = &v
	return s
}

func (s *PhoneNumberStatusForAccountRequest) SetMask(v string) *PhoneNumberStatusForAccountRequest {
	s.Mask = &v
	return s
}

func (s *PhoneNumberStatusForAccountRequest) SetOwnerId(v int64) *PhoneNumberStatusForAccountRequest {
	s.OwnerId = &v
	return s
}

func (s *PhoneNumberStatusForAccountRequest) SetResourceOwnerAccount(v string) *PhoneNumberStatusForAccountRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *PhoneNumberStatusForAccountRequest) SetResourceOwnerId(v int64) *PhoneNumberStatusForAccountRequest {
	s.ResourceOwnerId = &v
	return s
}

type PhoneNumberStatusForAccountResponseBody struct {
	Code      *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *PhoneNumberStatusForAccountResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PhoneNumberStatusForAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PhoneNumberStatusForAccountResponseBody) GoString() string {
	return s.String()
}

func (s *PhoneNumberStatusForAccountResponseBody) SetCode(v string) *PhoneNumberStatusForAccountResponseBody {
	s.Code = &v
	return s
}

func (s *PhoneNumberStatusForAccountResponseBody) SetData(v *PhoneNumberStatusForAccountResponseBodyData) *PhoneNumberStatusForAccountResponseBody {
	s.Data = v
	return s
}

func (s *PhoneNumberStatusForAccountResponseBody) SetMessage(v string) *PhoneNumberStatusForAccountResponseBody {
	s.Message = &v
	return s
}

func (s *PhoneNumberStatusForAccountResponseBody) SetRequestId(v string) *PhoneNumberStatusForAccountResponseBody {
	s.RequestId = &v
	return s
}

type PhoneNumberStatusForAccountResponseBodyData struct {
	Carrier *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	Status  *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s PhoneNumberStatusForAccountResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PhoneNumberStatusForAccountResponseBodyData) GoString() string {
	return s.String()
}

func (s *PhoneNumberStatusForAccountResponseBodyData) SetCarrier(v string) *PhoneNumberStatusForAccountResponseBodyData {
	s.Carrier = &v
	return s
}

func (s *PhoneNumberStatusForAccountResponseBodyData) SetStatus(v string) *PhoneNumberStatusForAccountResponseBodyData {
	s.Status = &v
	return s
}

type PhoneNumberStatusForAccountResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PhoneNumberStatusForAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PhoneNumberStatusForAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s PhoneNumberStatusForAccountResponse) GoString() string {
	return s.String()
}

func (s *PhoneNumberStatusForAccountResponse) SetHeaders(v map[string]*string) *PhoneNumberStatusForAccountResponse {
	s.Headers = v
	return s
}

func (s *PhoneNumberStatusForAccountResponse) SetStatusCode(v int32) *PhoneNumberStatusForAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *PhoneNumberStatusForAccountResponse) SetBody(v *PhoneNumberStatusForAccountResponseBody) *PhoneNumberStatusForAccountResponse {
	s.Body = v
	return s
}

type PhoneNumberStatusForRealRequest struct {
	AuthCode             *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	InputNumber          *string `json:"InputNumber,omitempty" xml:"InputNumber,omitempty"`
	Mask                 *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s PhoneNumberStatusForRealRequest) String() string {
	return tea.Prettify(s)
}

func (s PhoneNumberStatusForRealRequest) GoString() string {
	return s.String()
}

func (s *PhoneNumberStatusForRealRequest) SetAuthCode(v string) *PhoneNumberStatusForRealRequest {
	s.AuthCode = &v
	return s
}

func (s *PhoneNumberStatusForRealRequest) SetInputNumber(v string) *PhoneNumberStatusForRealRequest {
	s.InputNumber = &v
	return s
}

func (s *PhoneNumberStatusForRealRequest) SetMask(v string) *PhoneNumberStatusForRealRequest {
	s.Mask = &v
	return s
}

func (s *PhoneNumberStatusForRealRequest) SetOwnerId(v int64) *PhoneNumberStatusForRealRequest {
	s.OwnerId = &v
	return s
}

func (s *PhoneNumberStatusForRealRequest) SetResourceOwnerAccount(v string) *PhoneNumberStatusForRealRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *PhoneNumberStatusForRealRequest) SetResourceOwnerId(v int64) *PhoneNumberStatusForRealRequest {
	s.ResourceOwnerId = &v
	return s
}

type PhoneNumberStatusForRealResponseBody struct {
	Code      *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *PhoneNumberStatusForRealResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PhoneNumberStatusForRealResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PhoneNumberStatusForRealResponseBody) GoString() string {
	return s.String()
}

func (s *PhoneNumberStatusForRealResponseBody) SetCode(v string) *PhoneNumberStatusForRealResponseBody {
	s.Code = &v
	return s
}

func (s *PhoneNumberStatusForRealResponseBody) SetData(v *PhoneNumberStatusForRealResponseBodyData) *PhoneNumberStatusForRealResponseBody {
	s.Data = v
	return s
}

func (s *PhoneNumberStatusForRealResponseBody) SetMessage(v string) *PhoneNumberStatusForRealResponseBody {
	s.Message = &v
	return s
}

func (s *PhoneNumberStatusForRealResponseBody) SetRequestId(v string) *PhoneNumberStatusForRealResponseBody {
	s.RequestId = &v
	return s
}

type PhoneNumberStatusForRealResponseBodyData struct {
	Carrier *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	Status  *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s PhoneNumberStatusForRealResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PhoneNumberStatusForRealResponseBodyData) GoString() string {
	return s.String()
}

func (s *PhoneNumberStatusForRealResponseBodyData) SetCarrier(v string) *PhoneNumberStatusForRealResponseBodyData {
	s.Carrier = &v
	return s
}

func (s *PhoneNumberStatusForRealResponseBodyData) SetStatus(v string) *PhoneNumberStatusForRealResponseBodyData {
	s.Status = &v
	return s
}

type PhoneNumberStatusForRealResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PhoneNumberStatusForRealResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PhoneNumberStatusForRealResponse) String() string {
	return tea.Prettify(s)
}

func (s PhoneNumberStatusForRealResponse) GoString() string {
	return s.String()
}

func (s *PhoneNumberStatusForRealResponse) SetHeaders(v map[string]*string) *PhoneNumberStatusForRealResponse {
	s.Headers = v
	return s
}

func (s *PhoneNumberStatusForRealResponse) SetStatusCode(v int32) *PhoneNumberStatusForRealResponse {
	s.StatusCode = &v
	return s
}

func (s *PhoneNumberStatusForRealResponse) SetBody(v *PhoneNumberStatusForRealResponseBody) *PhoneNumberStatusForRealResponse {
	s.Body = v
	return s
}

type PhoneNumberStatusForSmsRequest struct {
	AuthCode             *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	InputNumber          *string `json:"InputNumber,omitempty" xml:"InputNumber,omitempty"`
	Mask                 *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s PhoneNumberStatusForSmsRequest) String() string {
	return tea.Prettify(s)
}

func (s PhoneNumberStatusForSmsRequest) GoString() string {
	return s.String()
}

func (s *PhoneNumberStatusForSmsRequest) SetAuthCode(v string) *PhoneNumberStatusForSmsRequest {
	s.AuthCode = &v
	return s
}

func (s *PhoneNumberStatusForSmsRequest) SetInputNumber(v string) *PhoneNumberStatusForSmsRequest {
	s.InputNumber = &v
	return s
}

func (s *PhoneNumberStatusForSmsRequest) SetMask(v string) *PhoneNumberStatusForSmsRequest {
	s.Mask = &v
	return s
}

func (s *PhoneNumberStatusForSmsRequest) SetOwnerId(v int64) *PhoneNumberStatusForSmsRequest {
	s.OwnerId = &v
	return s
}

func (s *PhoneNumberStatusForSmsRequest) SetResourceOwnerAccount(v string) *PhoneNumberStatusForSmsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *PhoneNumberStatusForSmsRequest) SetResourceOwnerId(v int64) *PhoneNumberStatusForSmsRequest {
	s.ResourceOwnerId = &v
	return s
}

type PhoneNumberStatusForSmsResponseBody struct {
	Code      *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *PhoneNumberStatusForSmsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PhoneNumberStatusForSmsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PhoneNumberStatusForSmsResponseBody) GoString() string {
	return s.String()
}

func (s *PhoneNumberStatusForSmsResponseBody) SetCode(v string) *PhoneNumberStatusForSmsResponseBody {
	s.Code = &v
	return s
}

func (s *PhoneNumberStatusForSmsResponseBody) SetData(v *PhoneNumberStatusForSmsResponseBodyData) *PhoneNumberStatusForSmsResponseBody {
	s.Data = v
	return s
}

func (s *PhoneNumberStatusForSmsResponseBody) SetMessage(v string) *PhoneNumberStatusForSmsResponseBody {
	s.Message = &v
	return s
}

func (s *PhoneNumberStatusForSmsResponseBody) SetRequestId(v string) *PhoneNumberStatusForSmsResponseBody {
	s.RequestId = &v
	return s
}

type PhoneNumberStatusForSmsResponseBodyData struct {
	Carrier *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	Status  *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s PhoneNumberStatusForSmsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PhoneNumberStatusForSmsResponseBodyData) GoString() string {
	return s.String()
}

func (s *PhoneNumberStatusForSmsResponseBodyData) SetCarrier(v string) *PhoneNumberStatusForSmsResponseBodyData {
	s.Carrier = &v
	return s
}

func (s *PhoneNumberStatusForSmsResponseBodyData) SetStatus(v string) *PhoneNumberStatusForSmsResponseBodyData {
	s.Status = &v
	return s
}

type PhoneNumberStatusForSmsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PhoneNumberStatusForSmsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PhoneNumberStatusForSmsResponse) String() string {
	return tea.Prettify(s)
}

func (s PhoneNumberStatusForSmsResponse) GoString() string {
	return s.String()
}

func (s *PhoneNumberStatusForSmsResponse) SetHeaders(v map[string]*string) *PhoneNumberStatusForSmsResponse {
	s.Headers = v
	return s
}

func (s *PhoneNumberStatusForSmsResponse) SetStatusCode(v int32) *PhoneNumberStatusForSmsResponse {
	s.StatusCode = &v
	return s
}

func (s *PhoneNumberStatusForSmsResponse) SetBody(v *PhoneNumberStatusForSmsResponseBody) *PhoneNumberStatusForSmsResponse {
	s.Body = v
	return s
}

type PhoneNumberStatusForVirtualRequest struct {
	AuthCode             *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	InputNumber          *string `json:"InputNumber,omitempty" xml:"InputNumber,omitempty"`
	Mask                 *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s PhoneNumberStatusForVirtualRequest) String() string {
	return tea.Prettify(s)
}

func (s PhoneNumberStatusForVirtualRequest) GoString() string {
	return s.String()
}

func (s *PhoneNumberStatusForVirtualRequest) SetAuthCode(v string) *PhoneNumberStatusForVirtualRequest {
	s.AuthCode = &v
	return s
}

func (s *PhoneNumberStatusForVirtualRequest) SetInputNumber(v string) *PhoneNumberStatusForVirtualRequest {
	s.InputNumber = &v
	return s
}

func (s *PhoneNumberStatusForVirtualRequest) SetMask(v string) *PhoneNumberStatusForVirtualRequest {
	s.Mask = &v
	return s
}

func (s *PhoneNumberStatusForVirtualRequest) SetOwnerId(v int64) *PhoneNumberStatusForVirtualRequest {
	s.OwnerId = &v
	return s
}

func (s *PhoneNumberStatusForVirtualRequest) SetResourceOwnerAccount(v string) *PhoneNumberStatusForVirtualRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *PhoneNumberStatusForVirtualRequest) SetResourceOwnerId(v int64) *PhoneNumberStatusForVirtualRequest {
	s.ResourceOwnerId = &v
	return s
}

type PhoneNumberStatusForVirtualResponseBody struct {
	Code      *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *PhoneNumberStatusForVirtualResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PhoneNumberStatusForVirtualResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PhoneNumberStatusForVirtualResponseBody) GoString() string {
	return s.String()
}

func (s *PhoneNumberStatusForVirtualResponseBody) SetCode(v string) *PhoneNumberStatusForVirtualResponseBody {
	s.Code = &v
	return s
}

func (s *PhoneNumberStatusForVirtualResponseBody) SetData(v *PhoneNumberStatusForVirtualResponseBodyData) *PhoneNumberStatusForVirtualResponseBody {
	s.Data = v
	return s
}

func (s *PhoneNumberStatusForVirtualResponseBody) SetMessage(v string) *PhoneNumberStatusForVirtualResponseBody {
	s.Message = &v
	return s
}

func (s *PhoneNumberStatusForVirtualResponseBody) SetRequestId(v string) *PhoneNumberStatusForVirtualResponseBody {
	s.RequestId = &v
	return s
}

type PhoneNumberStatusForVirtualResponseBodyData struct {
	IsPrivacyNumber *bool `json:"IsPrivacyNumber,omitempty" xml:"IsPrivacyNumber,omitempty"`
}

func (s PhoneNumberStatusForVirtualResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PhoneNumberStatusForVirtualResponseBodyData) GoString() string {
	return s.String()
}

func (s *PhoneNumberStatusForVirtualResponseBodyData) SetIsPrivacyNumber(v bool) *PhoneNumberStatusForVirtualResponseBodyData {
	s.IsPrivacyNumber = &v
	return s
}

type PhoneNumberStatusForVirtualResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PhoneNumberStatusForVirtualResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PhoneNumberStatusForVirtualResponse) String() string {
	return tea.Prettify(s)
}

func (s PhoneNumberStatusForVirtualResponse) GoString() string {
	return s.String()
}

func (s *PhoneNumberStatusForVirtualResponse) SetHeaders(v map[string]*string) *PhoneNumberStatusForVirtualResponse {
	s.Headers = v
	return s
}

func (s *PhoneNumberStatusForVirtualResponse) SetStatusCode(v int32) *PhoneNumberStatusForVirtualResponse {
	s.StatusCode = &v
	return s
}

func (s *PhoneNumberStatusForVirtualResponse) SetBody(v *PhoneNumberStatusForVirtualResponseBody) *PhoneNumberStatusForVirtualResponse {
	s.Body = v
	return s
}

type PhoneNumberStatusForVoiceRequest struct {
	AuthCode             *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	InputNumber          *string `json:"InputNumber,omitempty" xml:"InputNumber,omitempty"`
	Mask                 *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s PhoneNumberStatusForVoiceRequest) String() string {
	return tea.Prettify(s)
}

func (s PhoneNumberStatusForVoiceRequest) GoString() string {
	return s.String()
}

func (s *PhoneNumberStatusForVoiceRequest) SetAuthCode(v string) *PhoneNumberStatusForVoiceRequest {
	s.AuthCode = &v
	return s
}

func (s *PhoneNumberStatusForVoiceRequest) SetInputNumber(v string) *PhoneNumberStatusForVoiceRequest {
	s.InputNumber = &v
	return s
}

func (s *PhoneNumberStatusForVoiceRequest) SetMask(v string) *PhoneNumberStatusForVoiceRequest {
	s.Mask = &v
	return s
}

func (s *PhoneNumberStatusForVoiceRequest) SetOwnerId(v int64) *PhoneNumberStatusForVoiceRequest {
	s.OwnerId = &v
	return s
}

func (s *PhoneNumberStatusForVoiceRequest) SetResourceOwnerAccount(v string) *PhoneNumberStatusForVoiceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *PhoneNumberStatusForVoiceRequest) SetResourceOwnerId(v int64) *PhoneNumberStatusForVoiceRequest {
	s.ResourceOwnerId = &v
	return s
}

type PhoneNumberStatusForVoiceResponseBody struct {
	Code      *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *PhoneNumberStatusForVoiceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PhoneNumberStatusForVoiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PhoneNumberStatusForVoiceResponseBody) GoString() string {
	return s.String()
}

func (s *PhoneNumberStatusForVoiceResponseBody) SetCode(v string) *PhoneNumberStatusForVoiceResponseBody {
	s.Code = &v
	return s
}

func (s *PhoneNumberStatusForVoiceResponseBody) SetData(v *PhoneNumberStatusForVoiceResponseBodyData) *PhoneNumberStatusForVoiceResponseBody {
	s.Data = v
	return s
}

func (s *PhoneNumberStatusForVoiceResponseBody) SetMessage(v string) *PhoneNumberStatusForVoiceResponseBody {
	s.Message = &v
	return s
}

func (s *PhoneNumberStatusForVoiceResponseBody) SetRequestId(v string) *PhoneNumberStatusForVoiceResponseBody {
	s.RequestId = &v
	return s
}

type PhoneNumberStatusForVoiceResponseBodyData struct {
	Carrier *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	Status  *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s PhoneNumberStatusForVoiceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PhoneNumberStatusForVoiceResponseBodyData) GoString() string {
	return s.String()
}

func (s *PhoneNumberStatusForVoiceResponseBodyData) SetCarrier(v string) *PhoneNumberStatusForVoiceResponseBodyData {
	s.Carrier = &v
	return s
}

func (s *PhoneNumberStatusForVoiceResponseBodyData) SetStatus(v string) *PhoneNumberStatusForVoiceResponseBodyData {
	s.Status = &v
	return s
}

type PhoneNumberStatusForVoiceResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PhoneNumberStatusForVoiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PhoneNumberStatusForVoiceResponse) String() string {
	return tea.Prettify(s)
}

func (s PhoneNumberStatusForVoiceResponse) GoString() string {
	return s.String()
}

func (s *PhoneNumberStatusForVoiceResponse) SetHeaders(v map[string]*string) *PhoneNumberStatusForVoiceResponse {
	s.Headers = v
	return s
}

func (s *PhoneNumberStatusForVoiceResponse) SetStatusCode(v int32) *PhoneNumberStatusForVoiceResponse {
	s.StatusCode = &v
	return s
}

func (s *PhoneNumberStatusForVoiceResponse) SetBody(v *PhoneNumberStatusForVoiceResponseBody) *PhoneNumberStatusForVoiceResponse {
	s.Body = v
	return s
}

type ThreeElementsVerificationRequest struct {
	AuthCode             *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	CertCode             *string `json:"CertCode,omitempty" xml:"CertCode,omitempty"`
	InputNumber          *string `json:"InputNumber,omitempty" xml:"InputNumber,omitempty"`
	Mask                 *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ThreeElementsVerificationRequest) String() string {
	return tea.Prettify(s)
}

func (s ThreeElementsVerificationRequest) GoString() string {
	return s.String()
}

func (s *ThreeElementsVerificationRequest) SetAuthCode(v string) *ThreeElementsVerificationRequest {
	s.AuthCode = &v
	return s
}

func (s *ThreeElementsVerificationRequest) SetCertCode(v string) *ThreeElementsVerificationRequest {
	s.CertCode = &v
	return s
}

func (s *ThreeElementsVerificationRequest) SetInputNumber(v string) *ThreeElementsVerificationRequest {
	s.InputNumber = &v
	return s
}

func (s *ThreeElementsVerificationRequest) SetMask(v string) *ThreeElementsVerificationRequest {
	s.Mask = &v
	return s
}

func (s *ThreeElementsVerificationRequest) SetName(v string) *ThreeElementsVerificationRequest {
	s.Name = &v
	return s
}

func (s *ThreeElementsVerificationRequest) SetOwnerId(v int64) *ThreeElementsVerificationRequest {
	s.OwnerId = &v
	return s
}

func (s *ThreeElementsVerificationRequest) SetResourceOwnerAccount(v string) *ThreeElementsVerificationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ThreeElementsVerificationRequest) SetResourceOwnerId(v int64) *ThreeElementsVerificationRequest {
	s.ResourceOwnerId = &v
	return s
}

type ThreeElementsVerificationResponseBody struct {
	Code      *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ThreeElementsVerificationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ThreeElementsVerificationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ThreeElementsVerificationResponseBody) GoString() string {
	return s.String()
}

func (s *ThreeElementsVerificationResponseBody) SetCode(v string) *ThreeElementsVerificationResponseBody {
	s.Code = &v
	return s
}

func (s *ThreeElementsVerificationResponseBody) SetData(v *ThreeElementsVerificationResponseBodyData) *ThreeElementsVerificationResponseBody {
	s.Data = v
	return s
}

func (s *ThreeElementsVerificationResponseBody) SetMessage(v string) *ThreeElementsVerificationResponseBody {
	s.Message = &v
	return s
}

func (s *ThreeElementsVerificationResponseBody) SetRequestId(v string) *ThreeElementsVerificationResponseBody {
	s.RequestId = &v
	return s
}

type ThreeElementsVerificationResponseBodyData struct {
	BasicCarrier *string `json:"BasicCarrier,omitempty" xml:"BasicCarrier,omitempty"`
	IsConsistent *int32  `json:"IsConsistent,omitempty" xml:"IsConsistent,omitempty"`
}

func (s ThreeElementsVerificationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ThreeElementsVerificationResponseBodyData) GoString() string {
	return s.String()
}

func (s *ThreeElementsVerificationResponseBodyData) SetBasicCarrier(v string) *ThreeElementsVerificationResponseBodyData {
	s.BasicCarrier = &v
	return s
}

func (s *ThreeElementsVerificationResponseBodyData) SetIsConsistent(v int32) *ThreeElementsVerificationResponseBodyData {
	s.IsConsistent = &v
	return s
}

type ThreeElementsVerificationResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ThreeElementsVerificationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ThreeElementsVerificationResponse) String() string {
	return tea.Prettify(s)
}

func (s ThreeElementsVerificationResponse) GoString() string {
	return s.String()
}

func (s *ThreeElementsVerificationResponse) SetHeaders(v map[string]*string) *ThreeElementsVerificationResponse {
	s.Headers = v
	return s
}

func (s *ThreeElementsVerificationResponse) SetStatusCode(v int32) *ThreeElementsVerificationResponse {
	s.StatusCode = &v
	return s
}

func (s *ThreeElementsVerificationResponse) SetBody(v *ThreeElementsVerificationResponseBody) *ThreeElementsVerificationResponse {
	s.Body = v
	return s
}

type TwoElementsVerificationRequest struct {
	AuthCode             *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	InputNumber          *string `json:"InputNumber,omitempty" xml:"InputNumber,omitempty"`
	Mask                 *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s TwoElementsVerificationRequest) String() string {
	return tea.Prettify(s)
}

func (s TwoElementsVerificationRequest) GoString() string {
	return s.String()
}

func (s *TwoElementsVerificationRequest) SetAuthCode(v string) *TwoElementsVerificationRequest {
	s.AuthCode = &v
	return s
}

func (s *TwoElementsVerificationRequest) SetInputNumber(v string) *TwoElementsVerificationRequest {
	s.InputNumber = &v
	return s
}

func (s *TwoElementsVerificationRequest) SetMask(v string) *TwoElementsVerificationRequest {
	s.Mask = &v
	return s
}

func (s *TwoElementsVerificationRequest) SetName(v string) *TwoElementsVerificationRequest {
	s.Name = &v
	return s
}

func (s *TwoElementsVerificationRequest) SetOwnerId(v int64) *TwoElementsVerificationRequest {
	s.OwnerId = &v
	return s
}

func (s *TwoElementsVerificationRequest) SetResourceOwnerAccount(v string) *TwoElementsVerificationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *TwoElementsVerificationRequest) SetResourceOwnerId(v int64) *TwoElementsVerificationRequest {
	s.ResourceOwnerId = &v
	return s
}

type TwoElementsVerificationResponseBody struct {
	Code      *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *TwoElementsVerificationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TwoElementsVerificationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TwoElementsVerificationResponseBody) GoString() string {
	return s.String()
}

func (s *TwoElementsVerificationResponseBody) SetCode(v string) *TwoElementsVerificationResponseBody {
	s.Code = &v
	return s
}

func (s *TwoElementsVerificationResponseBody) SetData(v *TwoElementsVerificationResponseBodyData) *TwoElementsVerificationResponseBody {
	s.Data = v
	return s
}

func (s *TwoElementsVerificationResponseBody) SetMessage(v string) *TwoElementsVerificationResponseBody {
	s.Message = &v
	return s
}

func (s *TwoElementsVerificationResponseBody) SetRequestId(v string) *TwoElementsVerificationResponseBody {
	s.RequestId = &v
	return s
}

type TwoElementsVerificationResponseBodyData struct {
	BasicCarrier *string `json:"BasicCarrier,omitempty" xml:"BasicCarrier,omitempty"`
	IsConsistent *int32  `json:"IsConsistent,omitempty" xml:"IsConsistent,omitempty"`
}

func (s TwoElementsVerificationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s TwoElementsVerificationResponseBodyData) GoString() string {
	return s.String()
}

func (s *TwoElementsVerificationResponseBodyData) SetBasicCarrier(v string) *TwoElementsVerificationResponseBodyData {
	s.BasicCarrier = &v
	return s
}

func (s *TwoElementsVerificationResponseBodyData) SetIsConsistent(v int32) *TwoElementsVerificationResponseBodyData {
	s.IsConsistent = &v
	return s
}

type TwoElementsVerificationResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TwoElementsVerificationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TwoElementsVerificationResponse) String() string {
	return tea.Prettify(s)
}

func (s TwoElementsVerificationResponse) GoString() string {
	return s.String()
}

func (s *TwoElementsVerificationResponse) SetHeaders(v map[string]*string) *TwoElementsVerificationResponse {
	s.Headers = v
	return s
}

func (s *TwoElementsVerificationResponse) SetStatusCode(v int32) *TwoElementsVerificationResponse {
	s.StatusCode = &v
	return s
}

func (s *TwoElementsVerificationResponse) SetBody(v *TwoElementsVerificationResponseBody) *TwoElementsVerificationResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("dytnsapi"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) DescribeEmptyNumberDetectWithOptions(request *DescribeEmptyNumberDetectRequest, runtime *util.RuntimeOptions) (_result *DescribeEmptyNumberDetectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EncryptType)) {
		query["EncryptType"] = request.EncryptType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Phone)) {
		query["Phone"] = request.Phone
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeEmptyNumberDetect"),
		Version:     tea.String("2020-02-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeEmptyNumberDetectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEmptyNumberDetect(request *DescribeEmptyNumberDetectRequest) (_result *DescribeEmptyNumberDetectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEmptyNumberDetectResponse{}
	_body, _err := client.DescribeEmptyNumberDetectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePhoneNumberAnalysisWithOptions(request *DescribePhoneNumberAnalysisRequest, runtime *util.RuntimeOptions) (_result *DescribePhoneNumberAnalysisResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthCode)) {
		query["AuthCode"] = request.AuthCode
	}

	if !tea.BoolValue(util.IsUnset(request.InputNumber)) {
		query["InputNumber"] = request.InputNumber
	}

	if !tea.BoolValue(util.IsUnset(request.Mask)) {
		query["Mask"] = request.Mask
	}

	if !tea.BoolValue(util.IsUnset(request.NumberType)) {
		query["NumberType"] = request.NumberType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Rate)) {
		query["Rate"] = request.Rate
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePhoneNumberAnalysis"),
		Version:     tea.String("2020-02-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePhoneNumberAnalysisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePhoneNumberAnalysis(request *DescribePhoneNumberAnalysisRequest) (_result *DescribePhoneNumberAnalysisResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePhoneNumberAnalysisResponse{}
	_body, _err := client.DescribePhoneNumberAnalysisWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePhoneNumberAttributeWithOptions(request *DescribePhoneNumberAttributeRequest, runtime *util.RuntimeOptions) (_result *DescribePhoneNumberAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumber)) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePhoneNumberAttribute"),
		Version:     tea.String("2020-02-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePhoneNumberAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePhoneNumberAttribute(request *DescribePhoneNumberAttributeRequest) (_result *DescribePhoneNumberAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePhoneNumberAttributeResponse{}
	_body, _err := client.DescribePhoneNumberAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePhoneNumberOnlineTimeWithOptions(request *DescribePhoneNumberOnlineTimeRequest, runtime *util.RuntimeOptions) (_result *DescribePhoneNumberOnlineTimeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthCode)) {
		query["AuthCode"] = request.AuthCode
	}

	if !tea.BoolValue(util.IsUnset(request.Carrier)) {
		query["Carrier"] = request.Carrier
	}

	if !tea.BoolValue(util.IsUnset(request.InputNumber)) {
		query["InputNumber"] = request.InputNumber
	}

	if !tea.BoolValue(util.IsUnset(request.Mask)) {
		query["Mask"] = request.Mask
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePhoneNumberOnlineTime"),
		Version:     tea.String("2020-02-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePhoneNumberOnlineTimeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePhoneNumberOnlineTime(request *DescribePhoneNumberOnlineTimeRequest) (_result *DescribePhoneNumberOnlineTimeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePhoneNumberOnlineTimeResponse{}
	_body, _err := client.DescribePhoneNumberOnlineTimeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePhoneNumberOperatorAttributeWithOptions(request *DescribePhoneNumberOperatorAttributeRequest, runtime *util.RuntimeOptions) (_result *DescribePhoneNumberOperatorAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthCode)) {
		query["AuthCode"] = request.AuthCode
	}

	if !tea.BoolValue(util.IsUnset(request.InputNumber)) {
		query["InputNumber"] = request.InputNumber
	}

	if !tea.BoolValue(util.IsUnset(request.Mask)) {
		query["Mask"] = request.Mask
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePhoneNumberOperatorAttribute"),
		Version:     tea.String("2020-02-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePhoneNumberOperatorAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePhoneNumberOperatorAttribute(request *DescribePhoneNumberOperatorAttributeRequest) (_result *DescribePhoneNumberOperatorAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePhoneNumberOperatorAttributeResponse{}
	_body, _err := client.DescribePhoneNumberOperatorAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePhoneNumberResaleWithOptions(request *DescribePhoneNumberResaleRequest, runtime *util.RuntimeOptions) (_result *DescribePhoneNumberResaleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumber)) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Since)) {
		query["Since"] = request.Since
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePhoneNumberResale"),
		Version:     tea.String("2020-02-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePhoneNumberResaleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePhoneNumberResale(request *DescribePhoneNumberResaleRequest) (_result *DescribePhoneNumberResaleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePhoneNumberResaleResponse{}
	_body, _err := client.DescribePhoneNumberResaleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePhoneNumberStatusWithOptions(request *DescribePhoneNumberStatusRequest, runtime *util.RuntimeOptions) (_result *DescribePhoneNumberStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumber)) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePhoneNumberStatus"),
		Version:     tea.String("2020-02-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePhoneNumberStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePhoneNumberStatus(request *DescribePhoneNumberStatusRequest) (_result *DescribePhoneNumberStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePhoneNumberStatusResponse{}
	_body, _err := client.DescribePhoneNumberStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePhoneTwiceTelVerifyWithOptions(request *DescribePhoneTwiceTelVerifyRequest, runtime *util.RuntimeOptions) (_result *DescribePhoneTwiceTelVerifyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthCode)) {
		query["AuthCode"] = request.AuthCode
	}

	if !tea.BoolValue(util.IsUnset(request.InputNumber)) {
		query["InputNumber"] = request.InputNumber
	}

	if !tea.BoolValue(util.IsUnset(request.Mask)) {
		query["Mask"] = request.Mask
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePhoneTwiceTelVerify"),
		Version:     tea.String("2020-02-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePhoneTwiceTelVerifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePhoneTwiceTelVerify(request *DescribePhoneTwiceTelVerifyRequest) (_result *DescribePhoneTwiceTelVerifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePhoneTwiceTelVerifyResponse{}
	_body, _err := client.DescribePhoneTwiceTelVerifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InvalidPhoneNumberFilterWithOptions(request *InvalidPhoneNumberFilterRequest, runtime *util.RuntimeOptions) (_result *InvalidPhoneNumberFilterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthCode)) {
		query["AuthCode"] = request.AuthCode
	}

	if !tea.BoolValue(util.IsUnset(request.InputNumber)) {
		query["InputNumber"] = request.InputNumber
	}

	if !tea.BoolValue(util.IsUnset(request.Mask)) {
		query["Mask"] = request.Mask
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("InvalidPhoneNumberFilter"),
		Version:     tea.String("2020-02-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &InvalidPhoneNumberFilterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InvalidPhoneNumberFilter(request *InvalidPhoneNumberFilterRequest) (_result *InvalidPhoneNumberFilterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InvalidPhoneNumberFilterResponse{}
	_body, _err := client.InvalidPhoneNumberFilterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PhoneNumberEncryptWithOptions(request *PhoneNumberEncryptRequest, runtime *util.RuntimeOptions) (_result *PhoneNumberEncryptResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthCode)) {
		query["AuthCode"] = request.AuthCode
	}

	if !tea.BoolValue(util.IsUnset(request.InputNumber)) {
		query["InputNumber"] = request.InputNumber
	}

	if !tea.BoolValue(util.IsUnset(request.Mask)) {
		query["Mask"] = request.Mask
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PhoneNumberEncrypt"),
		Version:     tea.String("2020-02-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PhoneNumberEncryptResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PhoneNumberEncrypt(request *PhoneNumberEncryptRequest) (_result *PhoneNumberEncryptResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PhoneNumberEncryptResponse{}
	_body, _err := client.PhoneNumberEncryptWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PhoneNumberStatusForAccountWithOptions(request *PhoneNumberStatusForAccountRequest, runtime *util.RuntimeOptions) (_result *PhoneNumberStatusForAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthCode)) {
		query["AuthCode"] = request.AuthCode
	}

	if !tea.BoolValue(util.IsUnset(request.InputNumber)) {
		query["InputNumber"] = request.InputNumber
	}

	if !tea.BoolValue(util.IsUnset(request.Mask)) {
		query["Mask"] = request.Mask
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PhoneNumberStatusForAccount"),
		Version:     tea.String("2020-02-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PhoneNumberStatusForAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PhoneNumberStatusForAccount(request *PhoneNumberStatusForAccountRequest) (_result *PhoneNumberStatusForAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PhoneNumberStatusForAccountResponse{}
	_body, _err := client.PhoneNumberStatusForAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PhoneNumberStatusForRealWithOptions(request *PhoneNumberStatusForRealRequest, runtime *util.RuntimeOptions) (_result *PhoneNumberStatusForRealResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthCode)) {
		query["AuthCode"] = request.AuthCode
	}

	if !tea.BoolValue(util.IsUnset(request.InputNumber)) {
		query["InputNumber"] = request.InputNumber
	}

	if !tea.BoolValue(util.IsUnset(request.Mask)) {
		query["Mask"] = request.Mask
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PhoneNumberStatusForReal"),
		Version:     tea.String("2020-02-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PhoneNumberStatusForRealResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PhoneNumberStatusForReal(request *PhoneNumberStatusForRealRequest) (_result *PhoneNumberStatusForRealResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PhoneNumberStatusForRealResponse{}
	_body, _err := client.PhoneNumberStatusForRealWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PhoneNumberStatusForSmsWithOptions(request *PhoneNumberStatusForSmsRequest, runtime *util.RuntimeOptions) (_result *PhoneNumberStatusForSmsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthCode)) {
		query["AuthCode"] = request.AuthCode
	}

	if !tea.BoolValue(util.IsUnset(request.InputNumber)) {
		query["InputNumber"] = request.InputNumber
	}

	if !tea.BoolValue(util.IsUnset(request.Mask)) {
		query["Mask"] = request.Mask
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PhoneNumberStatusForSms"),
		Version:     tea.String("2020-02-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PhoneNumberStatusForSmsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PhoneNumberStatusForSms(request *PhoneNumberStatusForSmsRequest) (_result *PhoneNumberStatusForSmsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PhoneNumberStatusForSmsResponse{}
	_body, _err := client.PhoneNumberStatusForSmsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PhoneNumberStatusForVirtualWithOptions(request *PhoneNumberStatusForVirtualRequest, runtime *util.RuntimeOptions) (_result *PhoneNumberStatusForVirtualResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthCode)) {
		query["AuthCode"] = request.AuthCode
	}

	if !tea.BoolValue(util.IsUnset(request.InputNumber)) {
		query["InputNumber"] = request.InputNumber
	}

	if !tea.BoolValue(util.IsUnset(request.Mask)) {
		query["Mask"] = request.Mask
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PhoneNumberStatusForVirtual"),
		Version:     tea.String("2020-02-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PhoneNumberStatusForVirtualResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PhoneNumberStatusForVirtual(request *PhoneNumberStatusForVirtualRequest) (_result *PhoneNumberStatusForVirtualResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PhoneNumberStatusForVirtualResponse{}
	_body, _err := client.PhoneNumberStatusForVirtualWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PhoneNumberStatusForVoiceWithOptions(request *PhoneNumberStatusForVoiceRequest, runtime *util.RuntimeOptions) (_result *PhoneNumberStatusForVoiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthCode)) {
		query["AuthCode"] = request.AuthCode
	}

	if !tea.BoolValue(util.IsUnset(request.InputNumber)) {
		query["InputNumber"] = request.InputNumber
	}

	if !tea.BoolValue(util.IsUnset(request.Mask)) {
		query["Mask"] = request.Mask
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PhoneNumberStatusForVoice"),
		Version:     tea.String("2020-02-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PhoneNumberStatusForVoiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PhoneNumberStatusForVoice(request *PhoneNumberStatusForVoiceRequest) (_result *PhoneNumberStatusForVoiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PhoneNumberStatusForVoiceResponse{}
	_body, _err := client.PhoneNumberStatusForVoiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ThreeElementsVerificationWithOptions(request *ThreeElementsVerificationRequest, runtime *util.RuntimeOptions) (_result *ThreeElementsVerificationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthCode)) {
		query["AuthCode"] = request.AuthCode
	}

	if !tea.BoolValue(util.IsUnset(request.CertCode)) {
		query["CertCode"] = request.CertCode
	}

	if !tea.BoolValue(util.IsUnset(request.InputNumber)) {
		query["InputNumber"] = request.InputNumber
	}

	if !tea.BoolValue(util.IsUnset(request.Mask)) {
		query["Mask"] = request.Mask
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ThreeElementsVerification"),
		Version:     tea.String("2020-02-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ThreeElementsVerificationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ThreeElementsVerification(request *ThreeElementsVerificationRequest) (_result *ThreeElementsVerificationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ThreeElementsVerificationResponse{}
	_body, _err := client.ThreeElementsVerificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TwoElementsVerificationWithOptions(request *TwoElementsVerificationRequest, runtime *util.RuntimeOptions) (_result *TwoElementsVerificationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthCode)) {
		query["AuthCode"] = request.AuthCode
	}

	if !tea.BoolValue(util.IsUnset(request.InputNumber)) {
		query["InputNumber"] = request.InputNumber
	}

	if !tea.BoolValue(util.IsUnset(request.Mask)) {
		query["Mask"] = request.Mask
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TwoElementsVerification"),
		Version:     tea.String("2020-02-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TwoElementsVerificationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TwoElementsVerification(request *TwoElementsVerificationRequest) (_result *TwoElementsVerificationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TwoElementsVerificationResponse{}
	_body, _err := client.TwoElementsVerificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
