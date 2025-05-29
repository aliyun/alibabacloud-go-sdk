// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateChannelSignRequest struct {
	// This parameter is required.
	ChannelName *string `json:"channelName,omitempty" xml:"channelName,omitempty"`
	Contact     *string `json:"contact,omitempty" xml:"contact,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 13555555555
	Phone  *string `json:"phone,omitempty" xml:"phone,omitempty"`
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
}

func (s CreateChannelSignRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateChannelSignRequest) GoString() string {
	return s.String()
}

func (s *CreateChannelSignRequest) SetChannelName(v string) *CreateChannelSignRequest {
	s.ChannelName = &v
	return s
}

func (s *CreateChannelSignRequest) SetContact(v string) *CreateChannelSignRequest {
	s.Contact = &v
	return s
}

func (s *CreateChannelSignRequest) SetDescription(v string) *CreateChannelSignRequest {
	s.Description = &v
	return s
}

func (s *CreateChannelSignRequest) SetPhone(v string) *CreateChannelSignRequest {
	s.Phone = &v
	return s
}

func (s *CreateChannelSignRequest) SetRemark(v string) *CreateChannelSignRequest {
	s.Remark = &v
	return s
}

type CreateChannelSignResponseBody struct {
	// example:
	//
	// success
	Code *string                            `json:"code,omitempty" xml:"code,omitempty"`
	Data *CreateChannelSignResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 65857D96-A86B-5BBB-8392-0793E95DEB81
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *string `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateChannelSignResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateChannelSignResponseBody) GoString() string {
	return s.String()
}

func (s *CreateChannelSignResponseBody) SetCode(v string) *CreateChannelSignResponseBody {
	s.Code = &v
	return s
}

func (s *CreateChannelSignResponseBody) SetData(v *CreateChannelSignResponseBodyData) *CreateChannelSignResponseBody {
	s.Data = v
	return s
}

func (s *CreateChannelSignResponseBody) SetHttpStatusCode(v string) *CreateChannelSignResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateChannelSignResponseBody) SetMessage(v string) *CreateChannelSignResponseBody {
	s.Message = &v
	return s
}

func (s *CreateChannelSignResponseBody) SetRequestId(v string) *CreateChannelSignResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateChannelSignResponseBody) SetSuccess(v string) *CreateChannelSignResponseBody {
	s.Success = &v
	return s
}

type CreateChannelSignResponseBodyData struct {
	ChannelName *string `json:"channelName,omitempty" xml:"channelName,omitempty"`
	Contact     *string `json:"contact,omitempty" xml:"contact,omitempty"`
	// example:
	//
	// 2024-11-25 08:00:00
	CreateTime  *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 2024-11-25 08:00:00
	ModifyTime *string `json:"modifyTime,omitempty" xml:"modifyTime,omitempty"`
	// example:
	//
	// 13555555555
	Phone  *string `json:"phone,omitempty" xml:"phone,omitempty"`
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// example:
	//
	// review
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s CreateChannelSignResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateChannelSignResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateChannelSignResponseBodyData) SetChannelName(v string) *CreateChannelSignResponseBodyData {
	s.ChannelName = &v
	return s
}

func (s *CreateChannelSignResponseBodyData) SetContact(v string) *CreateChannelSignResponseBodyData {
	s.Contact = &v
	return s
}

func (s *CreateChannelSignResponseBodyData) SetCreateTime(v string) *CreateChannelSignResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *CreateChannelSignResponseBodyData) SetDescription(v string) *CreateChannelSignResponseBodyData {
	s.Description = &v
	return s
}

func (s *CreateChannelSignResponseBodyData) SetModifyTime(v string) *CreateChannelSignResponseBodyData {
	s.ModifyTime = &v
	return s
}

func (s *CreateChannelSignResponseBodyData) SetPhone(v string) *CreateChannelSignResponseBodyData {
	s.Phone = &v
	return s
}

func (s *CreateChannelSignResponseBodyData) SetRemark(v string) *CreateChannelSignResponseBodyData {
	s.Remark = &v
	return s
}

func (s *CreateChannelSignResponseBodyData) SetStatus(v string) *CreateChannelSignResponseBodyData {
	s.Status = &v
	return s
}

type CreateChannelSignResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateChannelSignResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateChannelSignResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateChannelSignResponse) GoString() string {
	return s.String()
}

func (s *CreateChannelSignResponse) SetHeaders(v map[string]*string) *CreateChannelSignResponse {
	s.Headers = v
	return s
}

func (s *CreateChannelSignResponse) SetStatusCode(v int32) *CreateChannelSignResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateChannelSignResponse) SetBody(v *CreateChannelSignResponseBody) *CreateChannelSignResponse {
	s.Body = v
	return s
}

type CreateProductRequest struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1221031972475999
	Euid *string `json:"euid,omitempty" xml:"euid,omitempty"`
	// This parameter is required.
	ProductName *string `json:"productName,omitempty" xml:"productName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 503041545
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CreateProductRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateProductRequest) GoString() string {
	return s.String()
}

func (s *CreateProductRequest) SetDescription(v string) *CreateProductRequest {
	s.Description = &v
	return s
}

func (s *CreateProductRequest) SetEuid(v string) *CreateProductRequest {
	s.Euid = &v
	return s
}

func (s *CreateProductRequest) SetProductName(v string) *CreateProductRequest {
	s.ProductName = &v
	return s
}

func (s *CreateProductRequest) SetTenantId(v string) *CreateProductRequest {
	s.TenantId = &v
	return s
}

func (s *CreateProductRequest) SetUserId(v string) *CreateProductRequest {
	s.UserId = &v
	return s
}

type CreateProductResponseBody struct {
	// example:
	//
	// success
	Code *string                        `json:"code,omitempty" xml:"code,omitempty"`
	Data *CreateProductResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 0abb7ee317248118358433637e749a
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *string `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateProductResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateProductResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProductResponseBody) SetCode(v string) *CreateProductResponseBody {
	s.Code = &v
	return s
}

func (s *CreateProductResponseBody) SetData(v *CreateProductResponseBodyData) *CreateProductResponseBody {
	s.Data = v
	return s
}

func (s *CreateProductResponseBody) SetHttpStatusCode(v string) *CreateProductResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateProductResponseBody) SetMessage(v string) *CreateProductResponseBody {
	s.Message = &v
	return s
}

func (s *CreateProductResponseBody) SetRequestId(v string) *CreateProductResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateProductResponseBody) SetSuccess(v string) *CreateProductResponseBody {
	s.Success = &v
	return s
}

type CreateProductResponseBodyData struct {
	// example:
	//
	// v7+7WMzYjHhyhb0c6IDL3e1rwCcphld19XzLYcwredVkHOht9et5GhPV45AqCzX7
	ApiKey      *string `json:"apiKey,omitempty" xml:"apiKey,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// el3SzmCU2p0x4RBc
	ProductKey  *string `json:"productKey,omitempty" xml:"productKey,omitempty"`
	ProductName *string `json:"productName,omitempty" xml:"productName,omitempty"`
	// example:
	//
	// 10bcac8989aed3f1047b71e6c06ef3ab
	ProductSecret *string `json:"productSecret,omitempty" xml:"productSecret,omitempty"`
	// example:
	//
	// 503041545
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// example:
	//
	// 123456
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CreateProductResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateProductResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateProductResponseBodyData) SetApiKey(v string) *CreateProductResponseBodyData {
	s.ApiKey = &v
	return s
}

func (s *CreateProductResponseBodyData) SetDescription(v string) *CreateProductResponseBodyData {
	s.Description = &v
	return s
}

func (s *CreateProductResponseBodyData) SetProductKey(v string) *CreateProductResponseBodyData {
	s.ProductKey = &v
	return s
}

func (s *CreateProductResponseBodyData) SetProductName(v string) *CreateProductResponseBodyData {
	s.ProductName = &v
	return s
}

func (s *CreateProductResponseBodyData) SetProductSecret(v string) *CreateProductResponseBodyData {
	s.ProductSecret = &v
	return s
}

func (s *CreateProductResponseBodyData) SetTenantId(v string) *CreateProductResponseBodyData {
	s.TenantId = &v
	return s
}

func (s *CreateProductResponseBodyData) SetUserId(v string) *CreateProductResponseBodyData {
	s.UserId = &v
	return s
}

type CreateProductResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateProductResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateProductResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateProductResponse) GoString() string {
	return s.String()
}

func (s *CreateProductResponse) SetHeaders(v map[string]*string) *CreateProductResponse {
	s.Headers = v
	return s
}

func (s *CreateProductResponse) SetStatusCode(v int32) *CreateProductResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateProductResponse) SetBody(v *CreateProductResponseBody) *CreateProductResponse {
	s.Body = v
	return s
}

type DeleteProductRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// el3SzmCU2p0x4RBc
	ProductKey *string `json:"productKey,omitempty" xml:"productKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 235454102432001
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s DeleteProductRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteProductRequest) GoString() string {
	return s.String()
}

func (s *DeleteProductRequest) SetProductKey(v string) *DeleteProductRequest {
	s.ProductKey = &v
	return s
}

func (s *DeleteProductRequest) SetTenantId(v string) *DeleteProductRequest {
	s.TenantId = &v
	return s
}

func (s *DeleteProductRequest) SetUserId(v string) *DeleteProductRequest {
	s.UserId = &v
	return s
}

type DeleteProductResponseBody struct {
	// example:
	//
	// success
	Code *string                        `json:"code,omitempty" xml:"code,omitempty"`
	Data *DeleteProductResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// AF54F772-60FF-56FD-A3EA-11620EF1229A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *string `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteProductResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteProductResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteProductResponseBody) SetCode(v string) *DeleteProductResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteProductResponseBody) SetData(v *DeleteProductResponseBodyData) *DeleteProductResponseBody {
	s.Data = v
	return s
}

func (s *DeleteProductResponseBody) SetHttpStatusCode(v string) *DeleteProductResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteProductResponseBody) SetMessage(v string) *DeleteProductResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteProductResponseBody) SetRequestId(v string) *DeleteProductResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteProductResponseBody) SetSuccess(v string) *DeleteProductResponseBody {
	s.Success = &v
	return s
}

type DeleteProductResponseBodyData struct {
	// example:
	//
	// el3SzmCU2p0x4RBc
	ProductKey *string `json:"productKey,omitempty" xml:"productKey,omitempty"`
	// example:
	//
	// btripOpen
	ProductName *string `json:"productName,omitempty" xml:"productName,omitempty"`
	// example:
	//
	// 355806813982786
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// example:
	//
	// 123456
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s DeleteProductResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DeleteProductResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteProductResponseBodyData) SetProductKey(v string) *DeleteProductResponseBodyData {
	s.ProductKey = &v
	return s
}

func (s *DeleteProductResponseBodyData) SetProductName(v string) *DeleteProductResponseBodyData {
	s.ProductName = &v
	return s
}

func (s *DeleteProductResponseBodyData) SetTenantId(v string) *DeleteProductResponseBodyData {
	s.TenantId = &v
	return s
}

func (s *DeleteProductResponseBodyData) SetUserId(v string) *DeleteProductResponseBodyData {
	s.UserId = &v
	return s
}

type DeleteProductResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteProductResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteProductResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteProductResponse) GoString() string {
	return s.String()
}

func (s *DeleteProductResponse) SetHeaders(v map[string]*string) *DeleteProductResponse {
	s.Headers = v
	return s
}

func (s *DeleteProductResponse) SetStatusCode(v int32) *DeleteProductResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteProductResponse) SetBody(v *DeleteProductResponseBody) *DeleteProductResponse {
	s.Body = v
	return s
}

type DeviceRegisterRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2a64edd96296880f55aa61987b
	Nonce *string `json:"nonce,omitempty" xml:"nonce,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// el3SzmCU2p0x4RBc
	ProductKey *string `json:"productKey,omitempty" xml:"productKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1748312544852
	RequestTime *string `json:"requestTime,omitempty" xml:"requestTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3spKwUgUpAGsXbbrHKnpVJPlI9wamoyhh96uqJuSyCKyJ7oscLAHRcz15dSzLG5L+ywFgYXSQNqdRtsn/Ri0j7pD0IuoKt9R7EnNo/U6viPvWD3Ldp3ehDDtOFtSrpUg6LTedvGtUWYU4x/zSD2jgCXijEdZCCMGCypcheMHRXfInYWF1xFtnCEXJfxtrBrnCk1p/pW3JSmdHJzmInnUEO3dWbNe3A==
	Signature *string `json:"signature,omitempty" xml:"signature,omitempty"`
}

func (s DeviceRegisterRequest) String() string {
	return tea.Prettify(s)
}

func (s DeviceRegisterRequest) GoString() string {
	return s.String()
}

func (s *DeviceRegisterRequest) SetNonce(v string) *DeviceRegisterRequest {
	s.Nonce = &v
	return s
}

func (s *DeviceRegisterRequest) SetProductKey(v string) *DeviceRegisterRequest {
	s.ProductKey = &v
	return s
}

func (s *DeviceRegisterRequest) SetRequestTime(v string) *DeviceRegisterRequest {
	s.RequestTime = &v
	return s
}

func (s *DeviceRegisterRequest) SetSignature(v string) *DeviceRegisterRequest {
	s.Signature = &v
	return s
}

type DeviceRegisterResponseBody struct {
	// example:
	//
	// success
	Code *string                         `json:"code,omitempty" xml:"code,omitempty"`
	Data *DeviceRegisterResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 52548360-B3AA-55EA-893F-48C16470F64A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeviceRegisterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeviceRegisterResponseBody) GoString() string {
	return s.String()
}

func (s *DeviceRegisterResponseBody) SetCode(v string) *DeviceRegisterResponseBody {
	s.Code = &v
	return s
}

func (s *DeviceRegisterResponseBody) SetData(v *DeviceRegisterResponseBodyData) *DeviceRegisterResponseBody {
	s.Data = v
	return s
}

func (s *DeviceRegisterResponseBody) SetHttpStatusCode(v int32) *DeviceRegisterResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeviceRegisterResponseBody) SetMessage(v string) *DeviceRegisterResponseBody {
	s.Message = &v
	return s
}

func (s *DeviceRegisterResponseBody) SetRequestId(v string) *DeviceRegisterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeviceRegisterResponseBody) SetSuccess(v bool) *DeviceRegisterResponseBody {
	s.Success = &v
	return s
}

type DeviceRegisterResponseBodyData struct {
	// example:
	//
	// 991fa52b7935aaa33536e05d4f4b5003
	DeviceName *string `json:"deviceName,omitempty" xml:"deviceName,omitempty"`
	// example:
	//
	// e2e928e8244f45ab30ec6ba9f9
	Nonce *string `json:"nonce,omitempty" xml:"nonce,omitempty"`
	// example:
	//
	// el3SzmCU2p0x4RBc
	ProductKey *string `json:"productKey,omitempty" xml:"productKey,omitempty"`
	// example:
	//
	// 1748312544852
	ResponseTime *string `json:"responseTime,omitempty" xml:"responseTime,omitempty"`
	// example:
	//
	// s8wPO/w79jP7sz6OaHkixAje2GmgzmZiCuCZZ+J8w77ICTyqD30lL6rUhnXwwx4MyGF62DRPFnpXtJ6c5mlmt6QdML3FkjGn+i/wR5T6QMkVDW6YRPWsx3jkIWRQ2sDnmVNBtixo2s9w3RJrnddRzVCaR/WeLOfiVLWcrLcJditzO/1YXBZ9vuRKQ4iperfhZvw372N/m8/1qtjJl+FUe2+wxK6RMxr03K7R
	Signature *string `json:"signature,omitempty" xml:"signature,omitempty"`
}

func (s DeviceRegisterResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DeviceRegisterResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeviceRegisterResponseBodyData) SetDeviceName(v string) *DeviceRegisterResponseBodyData {
	s.DeviceName = &v
	return s
}

func (s *DeviceRegisterResponseBodyData) SetNonce(v string) *DeviceRegisterResponseBodyData {
	s.Nonce = &v
	return s
}

func (s *DeviceRegisterResponseBodyData) SetProductKey(v string) *DeviceRegisterResponseBodyData {
	s.ProductKey = &v
	return s
}

func (s *DeviceRegisterResponseBodyData) SetResponseTime(v string) *DeviceRegisterResponseBodyData {
	s.ResponseTime = &v
	return s
}

func (s *DeviceRegisterResponseBodyData) SetSignature(v string) *DeviceRegisterResponseBodyData {
	s.Signature = &v
	return s
}

type DeviceRegisterResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeviceRegisterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeviceRegisterResponse) String() string {
	return tea.Prettify(s)
}

func (s DeviceRegisterResponse) GoString() string {
	return s.String()
}

func (s *DeviceRegisterResponse) SetHeaders(v map[string]*string) *DeviceRegisterResponse {
	s.Headers = v
	return s
}

func (s *DeviceRegisterResponse) SetStatusCode(v int32) *DeviceRegisterResponse {
	s.StatusCode = &v
	return s
}

func (s *DeviceRegisterResponse) SetBody(v *DeviceRegisterResponseBody) *DeviceRegisterResponse {
	s.Body = v
	return s
}

type GetChannelSignResponseBody struct {
	// example:
	//
	// success
	Code *string                         `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetChannelSignResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// D9272777-8401-5744-B059-BA21CF4BE80F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetChannelSignResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetChannelSignResponseBody) GoString() string {
	return s.String()
}

func (s *GetChannelSignResponseBody) SetCode(v string) *GetChannelSignResponseBody {
	s.Code = &v
	return s
}

func (s *GetChannelSignResponseBody) SetData(v *GetChannelSignResponseBodyData) *GetChannelSignResponseBody {
	s.Data = v
	return s
}

func (s *GetChannelSignResponseBody) SetHttpStatusCode(v string) *GetChannelSignResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetChannelSignResponseBody) SetMessage(v string) *GetChannelSignResponseBody {
	s.Message = &v
	return s
}

func (s *GetChannelSignResponseBody) SetRequestId(v string) *GetChannelSignResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetChannelSignResponseBody) SetSuccess(v string) *GetChannelSignResponseBody {
	s.Success = &v
	return s
}

type GetChannelSignResponseBodyData struct {
	ChannelName *string `json:"channelName,omitempty" xml:"channelName,omitempty"`
	Contact     *string `json:"contact,omitempty" xml:"contact,omitempty"`
	// example:
	//
	// 2025-05-24 00:00:00
	CreateTime  *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 2025-05-24 00:00:00
	ModifyTime *string `json:"modifyTime,omitempty" xml:"modifyTime,omitempty"`
	// example:
	//
	// 13555555555
	Phone  *string `json:"phone,omitempty" xml:"phone,omitempty"`
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// example:
	//
	// review
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetChannelSignResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetChannelSignResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetChannelSignResponseBodyData) SetChannelName(v string) *GetChannelSignResponseBodyData {
	s.ChannelName = &v
	return s
}

func (s *GetChannelSignResponseBodyData) SetContact(v string) *GetChannelSignResponseBodyData {
	s.Contact = &v
	return s
}

func (s *GetChannelSignResponseBodyData) SetCreateTime(v string) *GetChannelSignResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetChannelSignResponseBodyData) SetDescription(v string) *GetChannelSignResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetChannelSignResponseBodyData) SetModifyTime(v string) *GetChannelSignResponseBodyData {
	s.ModifyTime = &v
	return s
}

func (s *GetChannelSignResponseBodyData) SetPhone(v string) *GetChannelSignResponseBodyData {
	s.Phone = &v
	return s
}

func (s *GetChannelSignResponseBodyData) SetRemark(v string) *GetChannelSignResponseBodyData {
	s.Remark = &v
	return s
}

func (s *GetChannelSignResponseBodyData) SetStatus(v string) *GetChannelSignResponseBodyData {
	s.Status = &v
	return s
}

type GetChannelSignResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetChannelSignResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetChannelSignResponse) String() string {
	return tea.Prettify(s)
}

func (s GetChannelSignResponse) GoString() string {
	return s.String()
}

func (s *GetChannelSignResponse) SetHeaders(v map[string]*string) *GetChannelSignResponse {
	s.Headers = v
	return s
}

func (s *GetChannelSignResponse) SetStatusCode(v int32) *GetChannelSignResponse {
	s.StatusCode = &v
	return s
}

func (s *GetChannelSignResponse) SetBody(v *GetChannelSignResponseBody) *GetChannelSignResponse {
	s.Body = v
	return s
}

type GetQuotaInfoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 55
	RecordId *int64 `json:"recordId,omitempty" xml:"recordId,omitempty"`
}

func (s GetQuotaInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaInfoRequest) GoString() string {
	return s.String()
}

func (s *GetQuotaInfoRequest) SetRecordId(v int64) *GetQuotaInfoRequest {
	s.RecordId = &v
	return s
}

type GetQuotaInfoResponseBody struct {
	// example:
	//
	// success
	Code *string                       `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetQuotaInfoResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 7B0FC4BC-9E4B-5AD7-9D35-6559BDC0788E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetQuotaInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetQuotaInfoResponseBody) SetCode(v string) *GetQuotaInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetQuotaInfoResponseBody) SetData(v *GetQuotaInfoResponseBodyData) *GetQuotaInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetQuotaInfoResponseBody) SetHttpStatusCode(v string) *GetQuotaInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetQuotaInfoResponseBody) SetMessage(v string) *GetQuotaInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetQuotaInfoResponseBody) SetRequestId(v string) *GetQuotaInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQuotaInfoResponseBody) SetSuccess(v string) *GetQuotaInfoResponseBody {
	s.Success = &v
	return s
}

type GetQuotaInfoResponseBodyData struct {
	// example:
	//
	// 20
	ActiveLicenseCount *int64 `json:"activeLicenseCount,omitempty" xml:"activeLicenseCount,omitempty"`
	// example:
	//
	// 12
	Duration *int32 `json:"duration,omitempty" xml:"duration,omitempty"`
	// example:
	//
	// 2
	DurationType *int32 `json:"durationType,omitempty" xml:"durationType,omitempty"`
	// example:
	//
	// 100
	LicenseCount *int64 `json:"licenseCount,omitempty" xml:"licenseCount,omitempty"`
	// example:
	//
	// 100
	MaxQps *int32 `json:"maxQps,omitempty" xml:"maxQps,omitempty"`
	// example:
	//
	// el3SzmCU2p0x4RBc
	ProductKey *string `json:"productKey,omitempty" xml:"productKey,omitempty"`
	// example:
	//
	// 1
	PurchaseModel *int32 `json:"purchaseModel,omitempty" xml:"purchaseModel,omitempty"`
	// example:
	//
	// 51505222
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// example:
	//
	// 100000
	TokenNumber *int64 `json:"tokenNumber,omitempty" xml:"tokenNumber,omitempty"`
	// example:
	//
	// 123456
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetQuotaInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetQuotaInfoResponseBodyData) SetActiveLicenseCount(v int64) *GetQuotaInfoResponseBodyData {
	s.ActiveLicenseCount = &v
	return s
}

func (s *GetQuotaInfoResponseBodyData) SetDuration(v int32) *GetQuotaInfoResponseBodyData {
	s.Duration = &v
	return s
}

func (s *GetQuotaInfoResponseBodyData) SetDurationType(v int32) *GetQuotaInfoResponseBodyData {
	s.DurationType = &v
	return s
}

func (s *GetQuotaInfoResponseBodyData) SetLicenseCount(v int64) *GetQuotaInfoResponseBodyData {
	s.LicenseCount = &v
	return s
}

func (s *GetQuotaInfoResponseBodyData) SetMaxQps(v int32) *GetQuotaInfoResponseBodyData {
	s.MaxQps = &v
	return s
}

func (s *GetQuotaInfoResponseBodyData) SetProductKey(v string) *GetQuotaInfoResponseBodyData {
	s.ProductKey = &v
	return s
}

func (s *GetQuotaInfoResponseBodyData) SetPurchaseModel(v int32) *GetQuotaInfoResponseBodyData {
	s.PurchaseModel = &v
	return s
}

func (s *GetQuotaInfoResponseBodyData) SetTenantId(v string) *GetQuotaInfoResponseBodyData {
	s.TenantId = &v
	return s
}

func (s *GetQuotaInfoResponseBodyData) SetTokenNumber(v int64) *GetQuotaInfoResponseBodyData {
	s.TokenNumber = &v
	return s
}

func (s *GetQuotaInfoResponseBodyData) SetUserId(v string) *GetQuotaInfoResponseBodyData {
	s.UserId = &v
	return s
}

type GetQuotaInfoResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetQuotaInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetQuotaInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaInfoResponse) GoString() string {
	return s.String()
}

func (s *GetQuotaInfoResponse) SetHeaders(v map[string]*string) *GetQuotaInfoResponse {
	s.Headers = v
	return s
}

func (s *GetQuotaInfoResponse) SetStatusCode(v int32) *GetQuotaInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQuotaInfoResponse) SetBody(v *GetQuotaInfoResponseBody) *GetQuotaInfoResponse {
	s.Body = v
	return s
}

type GetTokenRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 5b504f84b69b9a73d3a21a2cff05e190
	DeviceName *string `json:"deviceName,omitempty" xml:"deviceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2a64edd96296880f55aa61987b
	Nonce *string `json:"nonce,omitempty" xml:"nonce,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// el3SzmCU2p0x4RBc
	ProductKey *string `json:"productKey,omitempty" xml:"productKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1748413148546
	RequestTime *string `json:"requestTime,omitempty" xml:"requestTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5/Smm8gnDSgZY2Blftq9eGYpVnpYCztoLJaJfhlH7id0lNlQxydRLtjUkGPr1qdbQq+oUn6Y1h0KJUdk0rf4am3MzdNr/Uhc47c8bbXnV8SlIC0agGo5skEQZNObzUD+sFxt8uN35/FYf7YRC8R61xY7+NPN2NLJrA/DPhewtVRRgAbb8RjergTcIG6oN1XTzLyC+76Az/3o2dPCxTfMXG3AFQc=
	Signature *string `json:"signature,omitempty" xml:"signature,omitempty"`
	// example:
	//
	// sk-4asv136677d2411b876e536bc8xxxxx
	TokenKey *string `json:"tokenKey,omitempty" xml:"tokenKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// oss
	TokenType *string `json:"tokenType,omitempty" xml:"tokenType,omitempty"`
}

func (s GetTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTokenRequest) GoString() string {
	return s.String()
}

func (s *GetTokenRequest) SetDeviceName(v string) *GetTokenRequest {
	s.DeviceName = &v
	return s
}

func (s *GetTokenRequest) SetNonce(v string) *GetTokenRequest {
	s.Nonce = &v
	return s
}

func (s *GetTokenRequest) SetProductKey(v string) *GetTokenRequest {
	s.ProductKey = &v
	return s
}

func (s *GetTokenRequest) SetRequestTime(v string) *GetTokenRequest {
	s.RequestTime = &v
	return s
}

func (s *GetTokenRequest) SetSignature(v string) *GetTokenRequest {
	s.Signature = &v
	return s
}

func (s *GetTokenRequest) SetTokenKey(v string) *GetTokenRequest {
	s.TokenKey = &v
	return s
}

func (s *GetTokenRequest) SetTokenType(v string) *GetTokenRequest {
	s.TokenType = &v
	return s
}

type GetTokenResponseBody struct {
	// example:
	//
	// success
	Code *string                   `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetTokenResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// B08AAA14-AD93-51F6-82AE-82AFAE9375B6
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetTokenResponseBody) SetCode(v string) *GetTokenResponseBody {
	s.Code = &v
	return s
}

func (s *GetTokenResponseBody) SetData(v *GetTokenResponseBodyData) *GetTokenResponseBody {
	s.Data = v
	return s
}

func (s *GetTokenResponseBody) SetHttpStatusCode(v string) *GetTokenResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetTokenResponseBody) SetMessage(v string) *GetTokenResponseBody {
	s.Message = &v
	return s
}

func (s *GetTokenResponseBody) SetRequestId(v string) *GetTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTokenResponseBody) SetSuccess(v string) *GetTokenResponseBody {
	s.Success = &v
	return s
}

type GetTokenResponseBodyData struct {
	// example:
	//
	// 5b504f84b69b9a73d3a21a2cff05e190
	DeviceName *string `json:"deviceName,omitempty" xml:"deviceName,omitempty"`
	// example:
	//
	// b79d692c315d6bfb28312edf15
	Nonce *string `json:"nonce,omitempty" xml:"nonce,omitempty"`
	// example:
	//
	// el3SzmCU2p0x4RBc
	ProductKey *string `json:"productKey,omitempty" xml:"productKey,omitempty"`
	// example:
	//
	// 127.0.0.1
	RequestIp *string `json:"requestIp,omitempty" xml:"requestIp,omitempty"`
	// example:
	//
	// 1748413248360
	ResponseTime *string `json:"responseTime,omitempty" xml:"responseTime,omitempty"`
	// example:
	//
	// N1faAjFhhaRNFaZNC8woRpQyAzEfBaIoWQEgDfds/Fwm7nIyEDLlSK3Ttx2OFebiHZ/MpHRr/3MnI/jpVWB/xNYIQxm6sccHJENHNAz6gaW+itU5wUrh+46EpqySABV8kc2pQ0HmYlbePfjjOK6lCfQjEGpekSAgQ6tDhG1lXWfKdtggq58Ut5bImMxMhk4R/PFUWrJe4CDuFu072C+foI0JlUV9TnGtVQ58oz8VRndrGXyauS/xqg8iGSZn6FyprUf5p+0ow20E
	Signature *string `json:"signature,omitempty" xml:"signature,omitempty"`
}

func (s GetTokenResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetTokenResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTokenResponseBodyData) SetDeviceName(v string) *GetTokenResponseBodyData {
	s.DeviceName = &v
	return s
}

func (s *GetTokenResponseBodyData) SetNonce(v string) *GetTokenResponseBodyData {
	s.Nonce = &v
	return s
}

func (s *GetTokenResponseBodyData) SetProductKey(v string) *GetTokenResponseBodyData {
	s.ProductKey = &v
	return s
}

func (s *GetTokenResponseBodyData) SetRequestIp(v string) *GetTokenResponseBodyData {
	s.RequestIp = &v
	return s
}

func (s *GetTokenResponseBodyData) SetResponseTime(v string) *GetTokenResponseBodyData {
	s.ResponseTime = &v
	return s
}

func (s *GetTokenResponseBodyData) SetSignature(v string) *GetTokenResponseBodyData {
	s.Signature = &v
	return s
}

type GetTokenResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTokenResponse) GoString() string {
	return s.String()
}

func (s *GetTokenResponse) SetHeaders(v map[string]*string) *GetTokenResponse {
	s.Headers = v
	return s
}

func (s *GetTokenResponse) SetStatusCode(v int32) *GetTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTokenResponse) SetBody(v *GetTokenResponseBody) *GetTokenResponse {
	s.Body = v
	return s
}

type HalfLLMAppCallRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// b883e6dcf14041fba390c1f795493c6b
	AppId    *string            `json:"appId,omitempty" xml:"appId,omitempty"`
	BizParam map[string]*string `json:"bizParam,omitempty" xml:"bizParam,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5b504f84b69b9a73d3a21a2cff05e190
	DeviceName *string `json:"deviceName,omitempty" xml:"deviceName,omitempty"`
	// This parameter is required.
	ModelTypeList []*string `json:"modelTypeList,omitempty" xml:"modelTypeList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 2oImhCz4f8XCviiM
	ProductKey *string `json:"productKey,omitempty" xml:"productKey,omitempty"`
	// example:
	//
	// 0sIRZDNncmCfBagzy534x2PH
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 678699541713794
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// This parameter is required.
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s HalfLLMAppCallRequest) String() string {
	return tea.Prettify(s)
}

func (s HalfLLMAppCallRequest) GoString() string {
	return s.String()
}

func (s *HalfLLMAppCallRequest) SetAppId(v string) *HalfLLMAppCallRequest {
	s.AppId = &v
	return s
}

func (s *HalfLLMAppCallRequest) SetBizParam(v map[string]*string) *HalfLLMAppCallRequest {
	s.BizParam = v
	return s
}

func (s *HalfLLMAppCallRequest) SetDeviceName(v string) *HalfLLMAppCallRequest {
	s.DeviceName = &v
	return s
}

func (s *HalfLLMAppCallRequest) SetModelTypeList(v []*string) *HalfLLMAppCallRequest {
	s.ModelTypeList = v
	return s
}

func (s *HalfLLMAppCallRequest) SetProductKey(v string) *HalfLLMAppCallRequest {
	s.ProductKey = &v
	return s
}

func (s *HalfLLMAppCallRequest) SetSessionId(v string) *HalfLLMAppCallRequest {
	s.SessionId = &v
	return s
}

func (s *HalfLLMAppCallRequest) SetTenantId(v string) *HalfLLMAppCallRequest {
	s.TenantId = &v
	return s
}

func (s *HalfLLMAppCallRequest) SetText(v string) *HalfLLMAppCallRequest {
	s.Text = &v
	return s
}

type HalfLLMAppCallShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// b883e6dcf14041fba390c1f795493c6b
	AppId          *string `json:"appId,omitempty" xml:"appId,omitempty"`
	BizParamShrink *string `json:"bizParam,omitempty" xml:"bizParam,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5b504f84b69b9a73d3a21a2cff05e190
	DeviceName *string `json:"deviceName,omitempty" xml:"deviceName,omitempty"`
	// This parameter is required.
	ModelTypeListShrink *string `json:"modelTypeList,omitempty" xml:"modelTypeList,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2oImhCz4f8XCviiM
	ProductKey *string `json:"productKey,omitempty" xml:"productKey,omitempty"`
	// example:
	//
	// 0sIRZDNncmCfBagzy534x2PH
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 678699541713794
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// This parameter is required.
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s HalfLLMAppCallShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s HalfLLMAppCallShrinkRequest) GoString() string {
	return s.String()
}

func (s *HalfLLMAppCallShrinkRequest) SetAppId(v string) *HalfLLMAppCallShrinkRequest {
	s.AppId = &v
	return s
}

func (s *HalfLLMAppCallShrinkRequest) SetBizParamShrink(v string) *HalfLLMAppCallShrinkRequest {
	s.BizParamShrink = &v
	return s
}

func (s *HalfLLMAppCallShrinkRequest) SetDeviceName(v string) *HalfLLMAppCallShrinkRequest {
	s.DeviceName = &v
	return s
}

func (s *HalfLLMAppCallShrinkRequest) SetModelTypeListShrink(v string) *HalfLLMAppCallShrinkRequest {
	s.ModelTypeListShrink = &v
	return s
}

func (s *HalfLLMAppCallShrinkRequest) SetProductKey(v string) *HalfLLMAppCallShrinkRequest {
	s.ProductKey = &v
	return s
}

func (s *HalfLLMAppCallShrinkRequest) SetSessionId(v string) *HalfLLMAppCallShrinkRequest {
	s.SessionId = &v
	return s
}

func (s *HalfLLMAppCallShrinkRequest) SetTenantId(v string) *HalfLLMAppCallShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *HalfLLMAppCallShrinkRequest) SetText(v string) *HalfLLMAppCallShrinkRequest {
	s.Text = &v
	return s
}

type HalfLLMAppCallResponseBody struct {
	// example:
	//
	// success
	Code *string                         `json:"code,omitempty" xml:"code,omitempty"`
	Data *HalfLLMAppCallResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// D7FEF19A-8B65-5914-9FA3-F26E820294B4
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"success,omitempty" xml:"success,omitempty"`
}

func (s HalfLLMAppCallResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HalfLLMAppCallResponseBody) GoString() string {
	return s.String()
}

func (s *HalfLLMAppCallResponseBody) SetCode(v string) *HalfLLMAppCallResponseBody {
	s.Code = &v
	return s
}

func (s *HalfLLMAppCallResponseBody) SetData(v *HalfLLMAppCallResponseBodyData) *HalfLLMAppCallResponseBody {
	s.Data = v
	return s
}

func (s *HalfLLMAppCallResponseBody) SetHttpStatusCode(v string) *HalfLLMAppCallResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *HalfLLMAppCallResponseBody) SetMessage(v string) *HalfLLMAppCallResponseBody {
	s.Message = &v
	return s
}

func (s *HalfLLMAppCallResponseBody) SetRequestId(v string) *HalfLLMAppCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *HalfLLMAppCallResponseBody) SetSuccess(v string) *HalfLLMAppCallResponseBody {
	s.Success = &v
	return s
}

type HalfLLMAppCallResponseBodyData struct {
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// null
	Message *string                               `json:"message,omitempty" xml:"message,omitempty"`
	Output  *HalfLLMAppCallResponseBodyDataOutput `json:"output,omitempty" xml:"output,omitempty" type:"Struct"`
	// example:
	//
	// 4Oy0zoqcjcclBgREcZvXF12y
	RequestId *string                           `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Rt        *HalfLLMAppCallResponseBodyDataRt `json:"rt,omitempty" xml:"rt,omitempty" type:"Struct"`
	// example:
	//
	// 0sIRZDNncmCfBagzy534x2PH
	SessionId *string                               `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	Usages    *HalfLLMAppCallResponseBodyDataUsages `json:"usages,omitempty" xml:"usages,omitempty" type:"Struct"`
}

func (s HalfLLMAppCallResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s HalfLLMAppCallResponseBodyData) GoString() string {
	return s.String()
}

func (s *HalfLLMAppCallResponseBodyData) SetCode(v string) *HalfLLMAppCallResponseBodyData {
	s.Code = &v
	return s
}

func (s *HalfLLMAppCallResponseBodyData) SetMessage(v string) *HalfLLMAppCallResponseBodyData {
	s.Message = &v
	return s
}

func (s *HalfLLMAppCallResponseBodyData) SetOutput(v *HalfLLMAppCallResponseBodyDataOutput) *HalfLLMAppCallResponseBodyData {
	s.Output = v
	return s
}

func (s *HalfLLMAppCallResponseBodyData) SetRequestId(v string) *HalfLLMAppCallResponseBodyData {
	s.RequestId = &v
	return s
}

func (s *HalfLLMAppCallResponseBodyData) SetRt(v *HalfLLMAppCallResponseBodyDataRt) *HalfLLMAppCallResponseBodyData {
	s.Rt = v
	return s
}

func (s *HalfLLMAppCallResponseBodyData) SetSessionId(v string) *HalfLLMAppCallResponseBodyData {
	s.SessionId = &v
	return s
}

func (s *HalfLLMAppCallResponseBodyData) SetUsages(v *HalfLLMAppCallResponseBodyDataUsages) *HalfLLMAppCallResponseBodyData {
	s.Usages = v
	return s
}

type HalfLLMAppCallResponseBodyDataOutput struct {
	Choices []*HalfLLMAppCallResponseBodyDataOutputChoices `json:"choices,omitempty" xml:"choices,omitempty" type:"Repeated"`
}

func (s HalfLLMAppCallResponseBodyDataOutput) String() string {
	return tea.Prettify(s)
}

func (s HalfLLMAppCallResponseBodyDataOutput) GoString() string {
	return s.String()
}

func (s *HalfLLMAppCallResponseBodyDataOutput) SetChoices(v []*HalfLLMAppCallResponseBodyDataOutputChoices) *HalfLLMAppCallResponseBodyDataOutput {
	s.Choices = v
	return s
}

type HalfLLMAppCallResponseBodyDataOutputChoices struct {
	// example:
	//
	// null
	FinishReason *string                                             `json:"finishReason,omitempty" xml:"finishReason,omitempty"`
	Message      *HalfLLMAppCallResponseBodyDataOutputChoicesMessage `json:"message,omitempty" xml:"message,omitempty" type:"Struct"`
}

func (s HalfLLMAppCallResponseBodyDataOutputChoices) String() string {
	return tea.Prettify(s)
}

func (s HalfLLMAppCallResponseBodyDataOutputChoices) GoString() string {
	return s.String()
}

func (s *HalfLLMAppCallResponseBodyDataOutputChoices) SetFinishReason(v string) *HalfLLMAppCallResponseBodyDataOutputChoices {
	s.FinishReason = &v
	return s
}

func (s *HalfLLMAppCallResponseBodyDataOutputChoices) SetMessage(v *HalfLLMAppCallResponseBodyDataOutputChoicesMessage) *HalfLLMAppCallResponseBodyDataOutputChoices {
	s.Message = v
	return s
}

type HalfLLMAppCallResponseBodyDataOutputChoicesMessage struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// assistant
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
}

func (s HalfLLMAppCallResponseBodyDataOutputChoicesMessage) String() string {
	return tea.Prettify(s)
}

func (s HalfLLMAppCallResponseBodyDataOutputChoicesMessage) GoString() string {
	return s.String()
}

func (s *HalfLLMAppCallResponseBodyDataOutputChoicesMessage) SetContent(v string) *HalfLLMAppCallResponseBodyDataOutputChoicesMessage {
	s.Content = &v
	return s
}

func (s *HalfLLMAppCallResponseBodyDataOutputChoicesMessage) SetRole(v string) *HalfLLMAppCallResponseBodyDataOutputChoicesMessage {
	s.Role = &v
	return s
}

type HalfLLMAppCallResponseBodyDataRt struct {
	// example:
	//
	// 443
	FirstRt *int64 `json:"firstRt,omitempty" xml:"firstRt,omitempty"`
	// example:
	//
	// 4432
	TotalRt *int64 `json:"totalRt,omitempty" xml:"totalRt,omitempty"`
}

func (s HalfLLMAppCallResponseBodyDataRt) String() string {
	return tea.Prettify(s)
}

func (s HalfLLMAppCallResponseBodyDataRt) GoString() string {
	return s.String()
}

func (s *HalfLLMAppCallResponseBodyDataRt) SetFirstRt(v int64) *HalfLLMAppCallResponseBodyDataRt {
	s.FirstRt = &v
	return s
}

func (s *HalfLLMAppCallResponseBodyDataRt) SetTotalRt(v int64) *HalfLLMAppCallResponseBodyDataRt {
	s.TotalRt = &v
	return s
}

type HalfLLMAppCallResponseBodyDataUsages struct {
	// example:
	//
	// 356
	InputTokens *int64 `json:"inputTokens,omitempty" xml:"inputTokens,omitempty"`
	// example:
	//
	// 698
	OutputTokens *int64 `json:"outputTokens,omitempty" xml:"outputTokens,omitempty"`
}

func (s HalfLLMAppCallResponseBodyDataUsages) String() string {
	return tea.Prettify(s)
}

func (s HalfLLMAppCallResponseBodyDataUsages) GoString() string {
	return s.String()
}

func (s *HalfLLMAppCallResponseBodyDataUsages) SetInputTokens(v int64) *HalfLLMAppCallResponseBodyDataUsages {
	s.InputTokens = &v
	return s
}

func (s *HalfLLMAppCallResponseBodyDataUsages) SetOutputTokens(v int64) *HalfLLMAppCallResponseBodyDataUsages {
	s.OutputTokens = &v
	return s
}

type HalfLLMAppCallResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HalfLLMAppCallResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HalfLLMAppCallResponse) String() string {
	return tea.Prettify(s)
}

func (s HalfLLMAppCallResponse) GoString() string {
	return s.String()
}

func (s *HalfLLMAppCallResponse) SetHeaders(v map[string]*string) *HalfLLMAppCallResponse {
	s.Headers = v
	return s
}

func (s *HalfLLMAppCallResponse) SetStatusCode(v int32) *HalfLLMAppCallResponse {
	s.StatusCode = &v
	return s
}

func (s *HalfLLMAppCallResponse) SetBody(v *HalfLLMAppCallResponseBody) *HalfLLMAppCallResponse {
	s.Body = v
	return s
}

type HalfLLMChatRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 5b504f84b69b9a73d3a21a2cff05e190
	DeviceName *string `json:"deviceName,omitempty" xml:"deviceName,omitempty"`
	// example:
	//
	// true
	EnableSearch *bool `json:"enableSearch,omitempty" xml:"enableSearch,omitempty"`
	// This parameter is required.
	InputText *string `json:"inputText,omitempty" xml:"inputText,omitempty"`
	// example:
	//
	// qwen-plus
	Model *string `json:"model,omitempty" xml:"model,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2oImhCz4f8XCviiM
	ProductKey *string `json:"productKey,omitempty" xml:"productKey,omitempty"`
	Prompt     *string `json:"prompt,omitempty" xml:"prompt,omitempty"`
	// example:
	//
	// 0sIRZDNncmCfBagzy534x2PH
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// example:
	//
	// true
	Stream *bool `json:"stream,omitempty" xml:"stream,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 520539530998273
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s HalfLLMChatRequest) String() string {
	return tea.Prettify(s)
}

func (s HalfLLMChatRequest) GoString() string {
	return s.String()
}

func (s *HalfLLMChatRequest) SetDeviceName(v string) *HalfLLMChatRequest {
	s.DeviceName = &v
	return s
}

func (s *HalfLLMChatRequest) SetEnableSearch(v bool) *HalfLLMChatRequest {
	s.EnableSearch = &v
	return s
}

func (s *HalfLLMChatRequest) SetInputText(v string) *HalfLLMChatRequest {
	s.InputText = &v
	return s
}

func (s *HalfLLMChatRequest) SetModel(v string) *HalfLLMChatRequest {
	s.Model = &v
	return s
}

func (s *HalfLLMChatRequest) SetProductKey(v string) *HalfLLMChatRequest {
	s.ProductKey = &v
	return s
}

func (s *HalfLLMChatRequest) SetPrompt(v string) *HalfLLMChatRequest {
	s.Prompt = &v
	return s
}

func (s *HalfLLMChatRequest) SetSessionId(v string) *HalfLLMChatRequest {
	s.SessionId = &v
	return s
}

func (s *HalfLLMChatRequest) SetStream(v bool) *HalfLLMChatRequest {
	s.Stream = &v
	return s
}

func (s *HalfLLMChatRequest) SetTenantId(v string) *HalfLLMChatRequest {
	s.TenantId = &v
	return s
}

type HalfLLMChatResponseBody struct {
	// example:
	//
	// success
	Code *string                      `json:"code,omitempty" xml:"code,omitempty"`
	Data *HalfLLMChatResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 5DD2E24F-93A2-551D-B192-8DBBEEFE9129
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"success,omitempty" xml:"success,omitempty"`
}

func (s HalfLLMChatResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HalfLLMChatResponseBody) GoString() string {
	return s.String()
}

func (s *HalfLLMChatResponseBody) SetCode(v string) *HalfLLMChatResponseBody {
	s.Code = &v
	return s
}

func (s *HalfLLMChatResponseBody) SetData(v *HalfLLMChatResponseBodyData) *HalfLLMChatResponseBody {
	s.Data = v
	return s
}

func (s *HalfLLMChatResponseBody) SetHttpStatusCode(v string) *HalfLLMChatResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *HalfLLMChatResponseBody) SetMessage(v string) *HalfLLMChatResponseBody {
	s.Message = &v
	return s
}

func (s *HalfLLMChatResponseBody) SetRequestId(v string) *HalfLLMChatResponseBody {
	s.RequestId = &v
	return s
}

func (s *HalfLLMChatResponseBody) SetSuccess(v string) *HalfLLMChatResponseBody {
	s.Success = &v
	return s
}

type HalfLLMChatResponseBodyData struct {
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// null
	Message *string                            `json:"message,omitempty" xml:"message,omitempty"`
	Output  *HalfLLMChatResponseBodyDataOutput `json:"output,omitempty" xml:"output,omitempty" type:"Struct"`
	// example:
	//
	// 4Oy0zoqcjcclBgREcZvXF12y
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 0sIRZDNncmCfBagzy534x2PH
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
}

func (s HalfLLMChatResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s HalfLLMChatResponseBodyData) GoString() string {
	return s.String()
}

func (s *HalfLLMChatResponseBodyData) SetCode(v string) *HalfLLMChatResponseBodyData {
	s.Code = &v
	return s
}

func (s *HalfLLMChatResponseBodyData) SetMessage(v string) *HalfLLMChatResponseBodyData {
	s.Message = &v
	return s
}

func (s *HalfLLMChatResponseBodyData) SetOutput(v *HalfLLMChatResponseBodyDataOutput) *HalfLLMChatResponseBodyData {
	s.Output = v
	return s
}

func (s *HalfLLMChatResponseBodyData) SetRequestId(v string) *HalfLLMChatResponseBodyData {
	s.RequestId = &v
	return s
}

func (s *HalfLLMChatResponseBodyData) SetSessionId(v string) *HalfLLMChatResponseBodyData {
	s.SessionId = &v
	return s
}

type HalfLLMChatResponseBodyDataOutput struct {
	Choices []*HalfLLMChatResponseBodyDataOutputChoices `json:"choices,omitempty" xml:"choices,omitempty" type:"Repeated"`
}

func (s HalfLLMChatResponseBodyDataOutput) String() string {
	return tea.Prettify(s)
}

func (s HalfLLMChatResponseBodyDataOutput) GoString() string {
	return s.String()
}

func (s *HalfLLMChatResponseBodyDataOutput) SetChoices(v []*HalfLLMChatResponseBodyDataOutputChoices) *HalfLLMChatResponseBodyDataOutput {
	s.Choices = v
	return s
}

type HalfLLMChatResponseBodyDataOutputChoices struct {
	// example:
	//
	// null
	FinishReason *string                                          `json:"finishReason,omitempty" xml:"finishReason,omitempty"`
	Message      *HalfLLMChatResponseBodyDataOutputChoicesMessage `json:"message,omitempty" xml:"message,omitempty" type:"Struct"`
}

func (s HalfLLMChatResponseBodyDataOutputChoices) String() string {
	return tea.Prettify(s)
}

func (s HalfLLMChatResponseBodyDataOutputChoices) GoString() string {
	return s.String()
}

func (s *HalfLLMChatResponseBodyDataOutputChoices) SetFinishReason(v string) *HalfLLMChatResponseBodyDataOutputChoices {
	s.FinishReason = &v
	return s
}

func (s *HalfLLMChatResponseBodyDataOutputChoices) SetMessage(v *HalfLLMChatResponseBodyDataOutputChoicesMessage) *HalfLLMChatResponseBodyDataOutputChoices {
	s.Message = v
	return s
}

type HalfLLMChatResponseBodyDataOutputChoicesMessage struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// assistant
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
}

func (s HalfLLMChatResponseBodyDataOutputChoicesMessage) String() string {
	return tea.Prettify(s)
}

func (s HalfLLMChatResponseBodyDataOutputChoicesMessage) GoString() string {
	return s.String()
}

func (s *HalfLLMChatResponseBodyDataOutputChoicesMessage) SetContent(v string) *HalfLLMChatResponseBodyDataOutputChoicesMessage {
	s.Content = &v
	return s
}

func (s *HalfLLMChatResponseBodyDataOutputChoicesMessage) SetRole(v string) *HalfLLMChatResponseBodyDataOutputChoicesMessage {
	s.Role = &v
	return s
}

type HalfLLMChatResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HalfLLMChatResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HalfLLMChatResponse) String() string {
	return tea.Prettify(s)
}

func (s HalfLLMChatResponse) GoString() string {
	return s.String()
}

func (s *HalfLLMChatResponse) SetHeaders(v map[string]*string) *HalfLLMChatResponse {
	s.Headers = v
	return s
}

func (s *HalfLLMChatResponse) SetStatusCode(v int32) *HalfLLMChatResponse {
	s.StatusCode = &v
	return s
}

func (s *HalfLLMChatResponse) SetBody(v *HalfLLMChatResponseBody) *HalfLLMChatResponse {
	s.Body = v
	return s
}

type HalfLLMImageChatRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 5b504f84b69b9a73d3a21a2cff05e190
	DeviceName *string `json:"deviceName,omitempty" xml:"deviceName,omitempty"`
	// example:
	//
	// false
	EnableSearch *bool `json:"enableSearch,omitempty" xml:"enableSearch,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://xxx/xx.jpg
	ImageUrl *string `json:"imageUrl,omitempty" xml:"imageUrl,omitempty"`
	// This parameter is required.
	InputText *string `json:"inputText,omitempty" xml:"inputText,omitempty"`
	// example:
	//
	// qwen-vl-max
	Model *string `json:"model,omitempty" xml:"model,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2oImhCz4f8XCviiM
	ProductKey *string `json:"productKey,omitempty" xml:"productKey,omitempty"`
	Prompt     *string `json:"prompt,omitempty" xml:"prompt,omitempty"`
	// example:
	//
	// 0sIRZDNncmCfBagzy534x2PH
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 520539530998273
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s HalfLLMImageChatRequest) String() string {
	return tea.Prettify(s)
}

func (s HalfLLMImageChatRequest) GoString() string {
	return s.String()
}

func (s *HalfLLMImageChatRequest) SetDeviceName(v string) *HalfLLMImageChatRequest {
	s.DeviceName = &v
	return s
}

func (s *HalfLLMImageChatRequest) SetEnableSearch(v bool) *HalfLLMImageChatRequest {
	s.EnableSearch = &v
	return s
}

func (s *HalfLLMImageChatRequest) SetImageUrl(v string) *HalfLLMImageChatRequest {
	s.ImageUrl = &v
	return s
}

func (s *HalfLLMImageChatRequest) SetInputText(v string) *HalfLLMImageChatRequest {
	s.InputText = &v
	return s
}

func (s *HalfLLMImageChatRequest) SetModel(v string) *HalfLLMImageChatRequest {
	s.Model = &v
	return s
}

func (s *HalfLLMImageChatRequest) SetProductKey(v string) *HalfLLMImageChatRequest {
	s.ProductKey = &v
	return s
}

func (s *HalfLLMImageChatRequest) SetPrompt(v string) *HalfLLMImageChatRequest {
	s.Prompt = &v
	return s
}

func (s *HalfLLMImageChatRequest) SetSessionId(v string) *HalfLLMImageChatRequest {
	s.SessionId = &v
	return s
}

func (s *HalfLLMImageChatRequest) SetTenantId(v string) *HalfLLMImageChatRequest {
	s.TenantId = &v
	return s
}

type HalfLLMImageChatResponseBody struct {
	// example:
	//
	// success
	Code *string                           `json:"code,omitempty" xml:"code,omitempty"`
	Data *HalfLLMImageChatResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 84656A01-32F0-5D12-8C72-E3AFAA0C8A29
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"success,omitempty" xml:"success,omitempty"`
}

func (s HalfLLMImageChatResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HalfLLMImageChatResponseBody) GoString() string {
	return s.String()
}

func (s *HalfLLMImageChatResponseBody) SetCode(v string) *HalfLLMImageChatResponseBody {
	s.Code = &v
	return s
}

func (s *HalfLLMImageChatResponseBody) SetData(v *HalfLLMImageChatResponseBodyData) *HalfLLMImageChatResponseBody {
	s.Data = v
	return s
}

func (s *HalfLLMImageChatResponseBody) SetHttpStatusCode(v string) *HalfLLMImageChatResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *HalfLLMImageChatResponseBody) SetMessage(v string) *HalfLLMImageChatResponseBody {
	s.Message = &v
	return s
}

func (s *HalfLLMImageChatResponseBody) SetRequestId(v string) *HalfLLMImageChatResponseBody {
	s.RequestId = &v
	return s
}

func (s *HalfLLMImageChatResponseBody) SetSuccess(v string) *HalfLLMImageChatResponseBody {
	s.Success = &v
	return s
}

type HalfLLMImageChatResponseBodyData struct {
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// null
	Message *string                                 `json:"message,omitempty" xml:"message,omitempty"`
	Output  *HalfLLMImageChatResponseBodyDataOutput `json:"output,omitempty" xml:"output,omitempty" type:"Struct"`
	// example:
	//
	// 4Oy0zoqcjcclBgREcZvXF12y
	RequestId *string                             `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Rt        *HalfLLMImageChatResponseBodyDataRt `json:"rt,omitempty" xml:"rt,omitempty" type:"Struct"`
	// example:
	//
	// 0sIRZDNncmCfBagzy534x2PH
	SessionId *string                                 `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	Usages    *HalfLLMImageChatResponseBodyDataUsages `json:"usages,omitempty" xml:"usages,omitempty" type:"Struct"`
}

func (s HalfLLMImageChatResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s HalfLLMImageChatResponseBodyData) GoString() string {
	return s.String()
}

func (s *HalfLLMImageChatResponseBodyData) SetCode(v string) *HalfLLMImageChatResponseBodyData {
	s.Code = &v
	return s
}

func (s *HalfLLMImageChatResponseBodyData) SetMessage(v string) *HalfLLMImageChatResponseBodyData {
	s.Message = &v
	return s
}

func (s *HalfLLMImageChatResponseBodyData) SetOutput(v *HalfLLMImageChatResponseBodyDataOutput) *HalfLLMImageChatResponseBodyData {
	s.Output = v
	return s
}

func (s *HalfLLMImageChatResponseBodyData) SetRequestId(v string) *HalfLLMImageChatResponseBodyData {
	s.RequestId = &v
	return s
}

func (s *HalfLLMImageChatResponseBodyData) SetRt(v *HalfLLMImageChatResponseBodyDataRt) *HalfLLMImageChatResponseBodyData {
	s.Rt = v
	return s
}

func (s *HalfLLMImageChatResponseBodyData) SetSessionId(v string) *HalfLLMImageChatResponseBodyData {
	s.SessionId = &v
	return s
}

func (s *HalfLLMImageChatResponseBodyData) SetUsages(v *HalfLLMImageChatResponseBodyDataUsages) *HalfLLMImageChatResponseBodyData {
	s.Usages = v
	return s
}

type HalfLLMImageChatResponseBodyDataOutput struct {
	Choices []*HalfLLMImageChatResponseBodyDataOutputChoices `json:"choices,omitempty" xml:"choices,omitempty" type:"Repeated"`
}

func (s HalfLLMImageChatResponseBodyDataOutput) String() string {
	return tea.Prettify(s)
}

func (s HalfLLMImageChatResponseBodyDataOutput) GoString() string {
	return s.String()
}

func (s *HalfLLMImageChatResponseBodyDataOutput) SetChoices(v []*HalfLLMImageChatResponseBodyDataOutputChoices) *HalfLLMImageChatResponseBodyDataOutput {
	s.Choices = v
	return s
}

type HalfLLMImageChatResponseBodyDataOutputChoices struct {
	// example:
	//
	// null
	FinishReason *string                                               `json:"finishReason,omitempty" xml:"finishReason,omitempty"`
	Message      *HalfLLMImageChatResponseBodyDataOutputChoicesMessage `json:"message,omitempty" xml:"message,omitempty" type:"Struct"`
}

func (s HalfLLMImageChatResponseBodyDataOutputChoices) String() string {
	return tea.Prettify(s)
}

func (s HalfLLMImageChatResponseBodyDataOutputChoices) GoString() string {
	return s.String()
}

func (s *HalfLLMImageChatResponseBodyDataOutputChoices) SetFinishReason(v string) *HalfLLMImageChatResponseBodyDataOutputChoices {
	s.FinishReason = &v
	return s
}

func (s *HalfLLMImageChatResponseBodyDataOutputChoices) SetMessage(v *HalfLLMImageChatResponseBodyDataOutputChoicesMessage) *HalfLLMImageChatResponseBodyDataOutputChoices {
	s.Message = v
	return s
}

type HalfLLMImageChatResponseBodyDataOutputChoicesMessage struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// assistant
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
}

func (s HalfLLMImageChatResponseBodyDataOutputChoicesMessage) String() string {
	return tea.Prettify(s)
}

func (s HalfLLMImageChatResponseBodyDataOutputChoicesMessage) GoString() string {
	return s.String()
}

func (s *HalfLLMImageChatResponseBodyDataOutputChoicesMessage) SetContent(v string) *HalfLLMImageChatResponseBodyDataOutputChoicesMessage {
	s.Content = &v
	return s
}

func (s *HalfLLMImageChatResponseBodyDataOutputChoicesMessage) SetRole(v string) *HalfLLMImageChatResponseBodyDataOutputChoicesMessage {
	s.Role = &v
	return s
}

type HalfLLMImageChatResponseBodyDataRt struct {
	// example:
	//
	// 1563
	FirstRt *int64 `json:"firstRt,omitempty" xml:"firstRt,omitempty"`
	// example:
	//
	// 8235
	TotalRt *int64 `json:"totalRt,omitempty" xml:"totalRt,omitempty"`
}

func (s HalfLLMImageChatResponseBodyDataRt) String() string {
	return tea.Prettify(s)
}

func (s HalfLLMImageChatResponseBodyDataRt) GoString() string {
	return s.String()
}

func (s *HalfLLMImageChatResponseBodyDataRt) SetFirstRt(v int64) *HalfLLMImageChatResponseBodyDataRt {
	s.FirstRt = &v
	return s
}

func (s *HalfLLMImageChatResponseBodyDataRt) SetTotalRt(v int64) *HalfLLMImageChatResponseBodyDataRt {
	s.TotalRt = &v
	return s
}

type HalfLLMImageChatResponseBodyDataUsages struct {
	// example:
	//
	// 136
	InputTokens *int64 `json:"inputTokens,omitempty" xml:"inputTokens,omitempty"`
	// example:
	//
	// 589
	OutputTokens *int64 `json:"outputTokens,omitempty" xml:"outputTokens,omitempty"`
}

func (s HalfLLMImageChatResponseBodyDataUsages) String() string {
	return tea.Prettify(s)
}

func (s HalfLLMImageChatResponseBodyDataUsages) GoString() string {
	return s.String()
}

func (s *HalfLLMImageChatResponseBodyDataUsages) SetInputTokens(v int64) *HalfLLMImageChatResponseBodyDataUsages {
	s.InputTokens = &v
	return s
}

func (s *HalfLLMImageChatResponseBodyDataUsages) SetOutputTokens(v int64) *HalfLLMImageChatResponseBodyDataUsages {
	s.OutputTokens = &v
	return s
}

type HalfLLMImageChatResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HalfLLMImageChatResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HalfLLMImageChatResponse) String() string {
	return tea.Prettify(s)
}

func (s HalfLLMImageChatResponse) GoString() string {
	return s.String()
}

func (s *HalfLLMImageChatResponse) SetHeaders(v map[string]*string) *HalfLLMImageChatResponse {
	s.Headers = v
	return s
}

func (s *HalfLLMImageChatResponse) SetStatusCode(v int32) *HalfLLMImageChatResponse {
	s.StatusCode = &v
	return s
}

func (s *HalfLLMImageChatResponse) SetBody(v *HalfLLMImageChatResponseBody) *HalfLLMImageChatResponse {
	s.Body = v
	return s
}

type HalfLLMImageOcrRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 5b504f84b69b9a73d3a21a2cff05e190
	DeviceName *string `json:"deviceName,omitempty" xml:"deviceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://xxx/xx.jpg
	ImageUrl *string `json:"imageUrl,omitempty" xml:"imageUrl,omitempty"`
	// example:
	//
	// qwen-vl-ocr
	Model *string `json:"model,omitempty" xml:"model,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2oImhCz4f8XCviiM
	ProductKey *string `json:"productKey,omitempty" xml:"productKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 520539530998273
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s HalfLLMImageOcrRequest) String() string {
	return tea.Prettify(s)
}

func (s HalfLLMImageOcrRequest) GoString() string {
	return s.String()
}

func (s *HalfLLMImageOcrRequest) SetDeviceName(v string) *HalfLLMImageOcrRequest {
	s.DeviceName = &v
	return s
}

func (s *HalfLLMImageOcrRequest) SetImageUrl(v string) *HalfLLMImageOcrRequest {
	s.ImageUrl = &v
	return s
}

func (s *HalfLLMImageOcrRequest) SetModel(v string) *HalfLLMImageOcrRequest {
	s.Model = &v
	return s
}

func (s *HalfLLMImageOcrRequest) SetProductKey(v string) *HalfLLMImageOcrRequest {
	s.ProductKey = &v
	return s
}

func (s *HalfLLMImageOcrRequest) SetTenantId(v string) *HalfLLMImageOcrRequest {
	s.TenantId = &v
	return s
}

type HalfLLMImageOcrResponseBody struct {
	// example:
	//
	// success
	Code *string                          `json:"code,omitempty" xml:"code,omitempty"`
	Data *HalfLLMImageOcrResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 4Oy0zoqcjcclBgREcZvXF12y
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"success,omitempty" xml:"success,omitempty"`
}

func (s HalfLLMImageOcrResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HalfLLMImageOcrResponseBody) GoString() string {
	return s.String()
}

func (s *HalfLLMImageOcrResponseBody) SetCode(v string) *HalfLLMImageOcrResponseBody {
	s.Code = &v
	return s
}

func (s *HalfLLMImageOcrResponseBody) SetData(v *HalfLLMImageOcrResponseBodyData) *HalfLLMImageOcrResponseBody {
	s.Data = v
	return s
}

func (s *HalfLLMImageOcrResponseBody) SetHttpStatusCode(v string) *HalfLLMImageOcrResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *HalfLLMImageOcrResponseBody) SetMessage(v string) *HalfLLMImageOcrResponseBody {
	s.Message = &v
	return s
}

func (s *HalfLLMImageOcrResponseBody) SetRequestId(v string) *HalfLLMImageOcrResponseBody {
	s.RequestId = &v
	return s
}

func (s *HalfLLMImageOcrResponseBody) SetSuccess(v string) *HalfLLMImageOcrResponseBody {
	s.Success = &v
	return s
}

type HalfLLMImageOcrResponseBodyData struct {
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// null
	Message *string                                `json:"message,omitempty" xml:"message,omitempty"`
	Output  *HalfLLMImageOcrResponseBodyDataOutput `json:"output,omitempty" xml:"output,omitempty" type:"Struct"`
	// example:
	//
	// 4Oy0zoqcjcclBgREcZvXF12y
	RequestId *string                            `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Rt        *HalfLLMImageOcrResponseBodyDataRt `json:"rt,omitempty" xml:"rt,omitempty" type:"Struct"`
	// example:
	//
	// 0sIRZDNncmCfBagzy534x2PH
	SessionId *string                                `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	Usages    *HalfLLMImageOcrResponseBodyDataUsages `json:"usages,omitempty" xml:"usages,omitempty" type:"Struct"`
}

func (s HalfLLMImageOcrResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s HalfLLMImageOcrResponseBodyData) GoString() string {
	return s.String()
}

func (s *HalfLLMImageOcrResponseBodyData) SetCode(v string) *HalfLLMImageOcrResponseBodyData {
	s.Code = &v
	return s
}

func (s *HalfLLMImageOcrResponseBodyData) SetMessage(v string) *HalfLLMImageOcrResponseBodyData {
	s.Message = &v
	return s
}

func (s *HalfLLMImageOcrResponseBodyData) SetOutput(v *HalfLLMImageOcrResponseBodyDataOutput) *HalfLLMImageOcrResponseBodyData {
	s.Output = v
	return s
}

func (s *HalfLLMImageOcrResponseBodyData) SetRequestId(v string) *HalfLLMImageOcrResponseBodyData {
	s.RequestId = &v
	return s
}

func (s *HalfLLMImageOcrResponseBodyData) SetRt(v *HalfLLMImageOcrResponseBodyDataRt) *HalfLLMImageOcrResponseBodyData {
	s.Rt = v
	return s
}

func (s *HalfLLMImageOcrResponseBodyData) SetSessionId(v string) *HalfLLMImageOcrResponseBodyData {
	s.SessionId = &v
	return s
}

func (s *HalfLLMImageOcrResponseBodyData) SetUsages(v *HalfLLMImageOcrResponseBodyDataUsages) *HalfLLMImageOcrResponseBodyData {
	s.Usages = v
	return s
}

type HalfLLMImageOcrResponseBodyDataOutput struct {
	Choices []*HalfLLMImageOcrResponseBodyDataOutputChoices `json:"choices,omitempty" xml:"choices,omitempty" type:"Repeated"`
}

func (s HalfLLMImageOcrResponseBodyDataOutput) String() string {
	return tea.Prettify(s)
}

func (s HalfLLMImageOcrResponseBodyDataOutput) GoString() string {
	return s.String()
}

func (s *HalfLLMImageOcrResponseBodyDataOutput) SetChoices(v []*HalfLLMImageOcrResponseBodyDataOutputChoices) *HalfLLMImageOcrResponseBodyDataOutput {
	s.Choices = v
	return s
}

type HalfLLMImageOcrResponseBodyDataOutputChoices struct {
	// example:
	//
	// null
	FinishReason *string                                              `json:"finishReason,omitempty" xml:"finishReason,omitempty"`
	Message      *HalfLLMImageOcrResponseBodyDataOutputChoicesMessage `json:"message,omitempty" xml:"message,omitempty" type:"Struct"`
}

func (s HalfLLMImageOcrResponseBodyDataOutputChoices) String() string {
	return tea.Prettify(s)
}

func (s HalfLLMImageOcrResponseBodyDataOutputChoices) GoString() string {
	return s.String()
}

func (s *HalfLLMImageOcrResponseBodyDataOutputChoices) SetFinishReason(v string) *HalfLLMImageOcrResponseBodyDataOutputChoices {
	s.FinishReason = &v
	return s
}

func (s *HalfLLMImageOcrResponseBodyDataOutputChoices) SetMessage(v *HalfLLMImageOcrResponseBodyDataOutputChoicesMessage) *HalfLLMImageOcrResponseBodyDataOutputChoices {
	s.Message = v
	return s
}

type HalfLLMImageOcrResponseBodyDataOutputChoicesMessage struct {
	// example:
	//
	// xxx
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// assistant
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
}

func (s HalfLLMImageOcrResponseBodyDataOutputChoicesMessage) String() string {
	return tea.Prettify(s)
}

func (s HalfLLMImageOcrResponseBodyDataOutputChoicesMessage) GoString() string {
	return s.String()
}

func (s *HalfLLMImageOcrResponseBodyDataOutputChoicesMessage) SetContent(v string) *HalfLLMImageOcrResponseBodyDataOutputChoicesMessage {
	s.Content = &v
	return s
}

func (s *HalfLLMImageOcrResponseBodyDataOutputChoicesMessage) SetRole(v string) *HalfLLMImageOcrResponseBodyDataOutputChoicesMessage {
	s.Role = &v
	return s
}

type HalfLLMImageOcrResponseBodyDataRt struct {
	// example:
	//
	// 635
	FirstRt *int64 `json:"firstRt,omitempty" xml:"firstRt,omitempty"`
	// example:
	//
	// 8571
	TotalRt *int64 `json:"totalRt,omitempty" xml:"totalRt,omitempty"`
}

func (s HalfLLMImageOcrResponseBodyDataRt) String() string {
	return tea.Prettify(s)
}

func (s HalfLLMImageOcrResponseBodyDataRt) GoString() string {
	return s.String()
}

func (s *HalfLLMImageOcrResponseBodyDataRt) SetFirstRt(v int64) *HalfLLMImageOcrResponseBodyDataRt {
	s.FirstRt = &v
	return s
}

func (s *HalfLLMImageOcrResponseBodyDataRt) SetTotalRt(v int64) *HalfLLMImageOcrResponseBodyDataRt {
	s.TotalRt = &v
	return s
}

type HalfLLMImageOcrResponseBodyDataUsages struct {
	// example:
	//
	// 771
	InputTokens *int64 `json:"inputTokens,omitempty" xml:"inputTokens,omitempty"`
	// example:
	//
	// 563
	OutputTokens *int64 `json:"outputTokens,omitempty" xml:"outputTokens,omitempty"`
}

func (s HalfLLMImageOcrResponseBodyDataUsages) String() string {
	return tea.Prettify(s)
}

func (s HalfLLMImageOcrResponseBodyDataUsages) GoString() string {
	return s.String()
}

func (s *HalfLLMImageOcrResponseBodyDataUsages) SetInputTokens(v int64) *HalfLLMImageOcrResponseBodyDataUsages {
	s.InputTokens = &v
	return s
}

func (s *HalfLLMImageOcrResponseBodyDataUsages) SetOutputTokens(v int64) *HalfLLMImageOcrResponseBodyDataUsages {
	s.OutputTokens = &v
	return s
}

type HalfLLMImageOcrResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HalfLLMImageOcrResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HalfLLMImageOcrResponse) String() string {
	return tea.Prettify(s)
}

func (s HalfLLMImageOcrResponse) GoString() string {
	return s.String()
}

func (s *HalfLLMImageOcrResponse) SetHeaders(v map[string]*string) *HalfLLMImageOcrResponse {
	s.Headers = v
	return s
}

func (s *HalfLLMImageOcrResponse) SetStatusCode(v int32) *HalfLLMImageOcrResponse {
	s.StatusCode = &v
	return s
}

func (s *HalfLLMImageOcrResponse) SetBody(v *HalfLLMImageOcrResponseBody) *HalfLLMImageOcrResponse {
	s.Body = v
	return s
}

type HalfLLMTTSChatRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 5b504f84b69b9a73d3a21a2cff05e190
	DeviceName *string `json:"deviceName,omitempty" xml:"deviceName,omitempty"`
	// example:
	//
	// false
	EnableSearch *bool `json:"enableSearch,omitempty" xml:"enableSearch,omitempty"`
	// example:
	//
	// .pcm。
	Format *string `json:"format,omitempty" xml:"format,omitempty"`
	// example:
	//
	// qwen-max
	Model *string `json:"model,omitempty" xml:"model,omitempty"`
	// example:
	//
	// 0
	PitchRate *int32 `json:"pitchRate,omitempty" xml:"pitchRate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2oImhCz4f8XCviiM
	ProductKey *string `json:"productKey,omitempty" xml:"productKey,omitempty"`
	Prompt     *string `json:"prompt,omitempty" xml:"prompt,omitempty"`
	// example:
	//
	// 16000
	SampleRate *int32 `json:"sampleRate,omitempty" xml:"sampleRate,omitempty"`
	// example:
	//
	// 0sIRZDNncmCfBagzy534x2PH
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// example:
	//
	// 0
	SpeechRate *int32 `json:"speechRate,omitempty" xml:"speechRate,omitempty"`
	// example:
	//
	// true
	Stream *bool `json:"stream,omitempty" xml:"stream,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 661708759700322
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// This parameter is required.
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
	// example:
	//
	// wss://nls-gateway-cn-beijing.aliyuncs.com/ws/v1
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
	// example:
	//
	// siyue
	Voice *string `json:"voice,omitempty" xml:"voice,omitempty"`
	// example:
	//
	// 50
	Volume *int32 `json:"volume,omitempty" xml:"volume,omitempty"`
}

func (s HalfLLMTTSChatRequest) String() string {
	return tea.Prettify(s)
}

func (s HalfLLMTTSChatRequest) GoString() string {
	return s.String()
}

func (s *HalfLLMTTSChatRequest) SetDeviceName(v string) *HalfLLMTTSChatRequest {
	s.DeviceName = &v
	return s
}

func (s *HalfLLMTTSChatRequest) SetEnableSearch(v bool) *HalfLLMTTSChatRequest {
	s.EnableSearch = &v
	return s
}

func (s *HalfLLMTTSChatRequest) SetFormat(v string) *HalfLLMTTSChatRequest {
	s.Format = &v
	return s
}

func (s *HalfLLMTTSChatRequest) SetModel(v string) *HalfLLMTTSChatRequest {
	s.Model = &v
	return s
}

func (s *HalfLLMTTSChatRequest) SetPitchRate(v int32) *HalfLLMTTSChatRequest {
	s.PitchRate = &v
	return s
}

func (s *HalfLLMTTSChatRequest) SetProductKey(v string) *HalfLLMTTSChatRequest {
	s.ProductKey = &v
	return s
}

func (s *HalfLLMTTSChatRequest) SetPrompt(v string) *HalfLLMTTSChatRequest {
	s.Prompt = &v
	return s
}

func (s *HalfLLMTTSChatRequest) SetSampleRate(v int32) *HalfLLMTTSChatRequest {
	s.SampleRate = &v
	return s
}

func (s *HalfLLMTTSChatRequest) SetSessionId(v string) *HalfLLMTTSChatRequest {
	s.SessionId = &v
	return s
}

func (s *HalfLLMTTSChatRequest) SetSpeechRate(v int32) *HalfLLMTTSChatRequest {
	s.SpeechRate = &v
	return s
}

func (s *HalfLLMTTSChatRequest) SetStream(v bool) *HalfLLMTTSChatRequest {
	s.Stream = &v
	return s
}

func (s *HalfLLMTTSChatRequest) SetTenantId(v string) *HalfLLMTTSChatRequest {
	s.TenantId = &v
	return s
}

func (s *HalfLLMTTSChatRequest) SetText(v string) *HalfLLMTTSChatRequest {
	s.Text = &v
	return s
}

func (s *HalfLLMTTSChatRequest) SetUrl(v string) *HalfLLMTTSChatRequest {
	s.Url = &v
	return s
}

func (s *HalfLLMTTSChatRequest) SetVoice(v string) *HalfLLMTTSChatRequest {
	s.Voice = &v
	return s
}

func (s *HalfLLMTTSChatRequest) SetVolume(v int32) *HalfLLMTTSChatRequest {
	s.Volume = &v
	return s
}

type HalfLLMTTSChatResponseBody struct {
	// example:
	//
	// success
	Code *string                         `json:"code,omitempty" xml:"code,omitempty"`
	Data *HalfLLMTTSChatResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// A9CE316B-B616-5A97-8FFC-5D0D664CA7AF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"success,omitempty" xml:"success,omitempty"`
}

func (s HalfLLMTTSChatResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HalfLLMTTSChatResponseBody) GoString() string {
	return s.String()
}

func (s *HalfLLMTTSChatResponseBody) SetCode(v string) *HalfLLMTTSChatResponseBody {
	s.Code = &v
	return s
}

func (s *HalfLLMTTSChatResponseBody) SetData(v *HalfLLMTTSChatResponseBodyData) *HalfLLMTTSChatResponseBody {
	s.Data = v
	return s
}

func (s *HalfLLMTTSChatResponseBody) SetHttpStatusCode(v string) *HalfLLMTTSChatResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *HalfLLMTTSChatResponseBody) SetMessage(v string) *HalfLLMTTSChatResponseBody {
	s.Message = &v
	return s
}

func (s *HalfLLMTTSChatResponseBody) SetRequestId(v string) *HalfLLMTTSChatResponseBody {
	s.RequestId = &v
	return s
}

func (s *HalfLLMTTSChatResponseBody) SetSuccess(v string) *HalfLLMTTSChatResponseBody {
	s.Success = &v
	return s
}

type HalfLLMTTSChatResponseBodyData struct {
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// xxxx
	Data []byte `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// null
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// sDwqZnFGwsv9x7yjVwQVKTV4
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 0sIRZDNncmCfBagzy534x2PH
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	Text      *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s HalfLLMTTSChatResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s HalfLLMTTSChatResponseBodyData) GoString() string {
	return s.String()
}

func (s *HalfLLMTTSChatResponseBodyData) SetCode(v string) *HalfLLMTTSChatResponseBodyData {
	s.Code = &v
	return s
}

func (s *HalfLLMTTSChatResponseBodyData) SetData(v []byte) *HalfLLMTTSChatResponseBodyData {
	s.Data = v
	return s
}

func (s *HalfLLMTTSChatResponseBodyData) SetMessage(v string) *HalfLLMTTSChatResponseBodyData {
	s.Message = &v
	return s
}

func (s *HalfLLMTTSChatResponseBodyData) SetRequestId(v string) *HalfLLMTTSChatResponseBodyData {
	s.RequestId = &v
	return s
}

func (s *HalfLLMTTSChatResponseBodyData) SetSessionId(v string) *HalfLLMTTSChatResponseBodyData {
	s.SessionId = &v
	return s
}

func (s *HalfLLMTTSChatResponseBodyData) SetText(v string) *HalfLLMTTSChatResponseBodyData {
	s.Text = &v
	return s
}

type HalfLLMTTSChatResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HalfLLMTTSChatResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HalfLLMTTSChatResponse) String() string {
	return tea.Prettify(s)
}

func (s HalfLLMTTSChatResponse) GoString() string {
	return s.String()
}

func (s *HalfLLMTTSChatResponse) SetHeaders(v map[string]*string) *HalfLLMTTSChatResponse {
	s.Headers = v
	return s
}

func (s *HalfLLMTTSChatResponse) SetStatusCode(v int32) *HalfLLMTTSChatResponse {
	s.StatusCode = &v
	return s
}

func (s *HalfLLMTTSChatResponse) SetBody(v *HalfLLMTTSChatResponseBody) *HalfLLMTTSChatResponse {
	s.Body = v
	return s
}

type QueryDevicePageRequest struct {
	// example:
	//
	// 5b504f84b69b9a73d3a21a2cff05e190
	DeviceName *string `json:"deviceName,omitempty" xml:"deviceName,omitempty"`
	// example:
	//
	// 1
	DisableStatus *int32 `json:"disableStatus,omitempty" xml:"disableStatus,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// el3SzmCU2p0x4RBc
	ProductKey  *string `json:"productKey,omitempty" xml:"productKey,omitempty"`
	ProductName *string `json:"productName,omitempty" xml:"productName,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
}

func (s QueryDevicePageRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDevicePageRequest) GoString() string {
	return s.String()
}

func (s *QueryDevicePageRequest) SetDeviceName(v string) *QueryDevicePageRequest {
	s.DeviceName = &v
	return s
}

func (s *QueryDevicePageRequest) SetDisableStatus(v int32) *QueryDevicePageRequest {
	s.DisableStatus = &v
	return s
}

func (s *QueryDevicePageRequest) SetPageIndex(v int32) *QueryDevicePageRequest {
	s.PageIndex = &v
	return s
}

func (s *QueryDevicePageRequest) SetPageSize(v int32) *QueryDevicePageRequest {
	s.PageSize = &v
	return s
}

func (s *QueryDevicePageRequest) SetProductKey(v string) *QueryDevicePageRequest {
	s.ProductKey = &v
	return s
}

func (s *QueryDevicePageRequest) SetProductName(v string) *QueryDevicePageRequest {
	s.ProductName = &v
	return s
}

func (s *QueryDevicePageRequest) SetStatus(v int32) *QueryDevicePageRequest {
	s.Status = &v
	return s
}

type QueryDevicePageResponseBody struct {
	// example:
	//
	// success
	Code *string                          `json:"code,omitempty" xml:"code,omitempty"`
	Data *QueryDevicePageResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 60FD351B-10C4-5C2C-ADA2-524FC39FC174
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"success,omitempty" xml:"success,omitempty"`
}

func (s QueryDevicePageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDevicePageResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDevicePageResponseBody) SetCode(v string) *QueryDevicePageResponseBody {
	s.Code = &v
	return s
}

func (s *QueryDevicePageResponseBody) SetData(v *QueryDevicePageResponseBodyData) *QueryDevicePageResponseBody {
	s.Data = v
	return s
}

func (s *QueryDevicePageResponseBody) SetHttpStatusCode(v int32) *QueryDevicePageResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryDevicePageResponseBody) SetMessage(v string) *QueryDevicePageResponseBody {
	s.Message = &v
	return s
}

func (s *QueryDevicePageResponseBody) SetRequestId(v string) *QueryDevicePageResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDevicePageResponseBody) SetSuccess(v string) *QueryDevicePageResponseBody {
	s.Success = &v
	return s
}

type QueryDevicePageResponseBodyData struct {
	Data []*QueryDevicePageResponseBodyDataData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 100
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s QueryDevicePageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryDevicePageResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryDevicePageResponseBodyData) SetData(v []*QueryDevicePageResponseBodyDataData) *QueryDevicePageResponseBodyData {
	s.Data = v
	return s
}

func (s *QueryDevicePageResponseBodyData) SetPageIndex(v int32) *QueryDevicePageResponseBodyData {
	s.PageIndex = &v
	return s
}

func (s *QueryDevicePageResponseBodyData) SetPageSize(v int32) *QueryDevicePageResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryDevicePageResponseBodyData) SetTotal(v int64) *QueryDevicePageResponseBodyData {
	s.Total = &v
	return s
}

type QueryDevicePageResponseBodyDataData struct {
	// example:
	//
	// 2025-03-15 09:44:32
	ActiveTime *string `json:"activeTime,omitempty" xml:"activeTime,omitempty"`
	// example:
	//
	// 1539704706413278
	AliyunUid *string `json:"aliyunUid,omitempty" xml:"aliyunUid,omitempty"`
	// example:
	//
	// 202504010001
	BatchNo *string `json:"batchNo,omitempty" xml:"batchNo,omitempty"`
	// example:
	//
	// 5b504f84b69b9a73d3a21a2cff05e190
	DeviceName *string `json:"deviceName,omitempty" xml:"deviceName,omitempty"`
	// example:
	//
	// 1
	DisableStatus *int32 `json:"disableStatus,omitempty" xml:"disableStatus,omitempty"`
	// example:
	//
	// 2025-04-27 09:10:31
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 2025-04-27 09:10:31
	GmtModify *string `json:"gmtModify,omitempty" xml:"gmtModify,omitempty"`
	// ID。
	//
	// example:
	//
	// 201
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// nnjNQQv3ZjyYE7H4
	ProductKey  *string `json:"productKey,omitempty" xml:"productKey,omitempty"`
	ProductName *string `json:"productName,omitempty" xml:"productName,omitempty"`
	Remark      *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 493303079000577
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s QueryDevicePageResponseBodyDataData) String() string {
	return tea.Prettify(s)
}

func (s QueryDevicePageResponseBodyDataData) GoString() string {
	return s.String()
}

func (s *QueryDevicePageResponseBodyDataData) SetActiveTime(v string) *QueryDevicePageResponseBodyDataData {
	s.ActiveTime = &v
	return s
}

func (s *QueryDevicePageResponseBodyDataData) SetAliyunUid(v string) *QueryDevicePageResponseBodyDataData {
	s.AliyunUid = &v
	return s
}

func (s *QueryDevicePageResponseBodyDataData) SetBatchNo(v string) *QueryDevicePageResponseBodyDataData {
	s.BatchNo = &v
	return s
}

func (s *QueryDevicePageResponseBodyDataData) SetDeviceName(v string) *QueryDevicePageResponseBodyDataData {
	s.DeviceName = &v
	return s
}

func (s *QueryDevicePageResponseBodyDataData) SetDisableStatus(v int32) *QueryDevicePageResponseBodyDataData {
	s.DisableStatus = &v
	return s
}

func (s *QueryDevicePageResponseBodyDataData) SetGmtCreate(v string) *QueryDevicePageResponseBodyDataData {
	s.GmtCreate = &v
	return s
}

func (s *QueryDevicePageResponseBodyDataData) SetGmtModify(v string) *QueryDevicePageResponseBodyDataData {
	s.GmtModify = &v
	return s
}

func (s *QueryDevicePageResponseBodyDataData) SetId(v int64) *QueryDevicePageResponseBodyDataData {
	s.Id = &v
	return s
}

func (s *QueryDevicePageResponseBodyDataData) SetProductKey(v string) *QueryDevicePageResponseBodyDataData {
	s.ProductKey = &v
	return s
}

func (s *QueryDevicePageResponseBodyDataData) SetProductName(v string) *QueryDevicePageResponseBodyDataData {
	s.ProductName = &v
	return s
}

func (s *QueryDevicePageResponseBodyDataData) SetRemark(v string) *QueryDevicePageResponseBodyDataData {
	s.Remark = &v
	return s
}

func (s *QueryDevicePageResponseBodyDataData) SetStatus(v int32) *QueryDevicePageResponseBodyDataData {
	s.Status = &v
	return s
}

func (s *QueryDevicePageResponseBodyDataData) SetTenantId(v string) *QueryDevicePageResponseBodyDataData {
	s.TenantId = &v
	return s
}

type QueryDevicePageResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDevicePageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDevicePageResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDevicePageResponse) GoString() string {
	return s.String()
}

func (s *QueryDevicePageResponse) SetHeaders(v map[string]*string) *QueryDevicePageResponse {
	s.Headers = v
	return s
}

func (s *QueryDevicePageResponse) SetStatusCode(v int32) *QueryDevicePageResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDevicePageResponse) SetBody(v *QueryDevicePageResponseBody) *QueryDevicePageResponse {
	s.Body = v
	return s
}

type QueryProductPageRequest struct {
	// example:
	//
	// 1
	ModelType []byte `json:"modelType,omitempty" xml:"modelType,omitempty"`
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// example:
	//
	// 20
	PageSize    *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	ProductName *string `json:"productName,omitempty" xml:"productName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 679583000646594
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 359687
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryProductPageRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryProductPageRequest) GoString() string {
	return s.String()
}

func (s *QueryProductPageRequest) SetModelType(v []byte) *QueryProductPageRequest {
	s.ModelType = v
	return s
}

func (s *QueryProductPageRequest) SetPageIndex(v int32) *QueryProductPageRequest {
	s.PageIndex = &v
	return s
}

func (s *QueryProductPageRequest) SetPageSize(v int32) *QueryProductPageRequest {
	s.PageSize = &v
	return s
}

func (s *QueryProductPageRequest) SetProductName(v string) *QueryProductPageRequest {
	s.ProductName = &v
	return s
}

func (s *QueryProductPageRequest) SetTenantId(v string) *QueryProductPageRequest {
	s.TenantId = &v
	return s
}

func (s *QueryProductPageRequest) SetUserId(v string) *QueryProductPageRequest {
	s.UserId = &v
	return s
}

type QueryProductPageResponseBody struct {
	// example:
	//
	// success
	Code *string                           `json:"code,omitempty" xml:"code,omitempty"`
	Data *QueryProductPageResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// EA4643D5-5FA8-55BA-A959-F7D3E38E0AE0
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"success,omitempty" xml:"success,omitempty"`
}

func (s QueryProductPageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryProductPageResponseBody) GoString() string {
	return s.String()
}

func (s *QueryProductPageResponseBody) SetCode(v string) *QueryProductPageResponseBody {
	s.Code = &v
	return s
}

func (s *QueryProductPageResponseBody) SetData(v *QueryProductPageResponseBodyData) *QueryProductPageResponseBody {
	s.Data = v
	return s
}

func (s *QueryProductPageResponseBody) SetHttpStatusCode(v string) *QueryProductPageResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryProductPageResponseBody) SetMessage(v string) *QueryProductPageResponseBody {
	s.Message = &v
	return s
}

func (s *QueryProductPageResponseBody) SetRequestId(v string) *QueryProductPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryProductPageResponseBody) SetSuccess(v string) *QueryProductPageResponseBody {
	s.Success = &v
	return s
}

type QueryProductPageResponseBodyData struct {
	Data []*QueryProductPageResponseBodyDataData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 100
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s QueryProductPageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryProductPageResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryProductPageResponseBodyData) SetData(v []*QueryProductPageResponseBodyDataData) *QueryProductPageResponseBodyData {
	s.Data = v
	return s
}

func (s *QueryProductPageResponseBodyData) SetPageIndex(v int32) *QueryProductPageResponseBodyData {
	s.PageIndex = &v
	return s
}

func (s *QueryProductPageResponseBodyData) SetPageSize(v int32) *QueryProductPageResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryProductPageResponseBodyData) SetTotal(v int64) *QueryProductPageResponseBodyData {
	s.Total = &v
	return s
}

type QueryProductPageResponseBodyDataData struct {
	// example:
	//
	// 100
	ActiveLicenseCount *int64 `json:"activeLicenseCount,omitempty" xml:"activeLicenseCount,omitempty"`
	// example:
	//
	// zcrzbqrF29pkgXukLaQ+6jGsohQiPhdOuIrUSVSvNO5oDntQdM76mNXj+AJ2i7eP
	ApiKey *string `json:"apiKey,omitempty" xml:"apiKey,omitempty"`
	// example:
	//
	// 2024-03-05 06:24:27
	CreateTime  *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 1000
	LicenseCount *int64 `json:"licenseCount,omitempty" xml:"licenseCount,omitempty"`
	// example:
	//
	// 100
	MaxQps *int32 `json:"maxQps,omitempty" xml:"maxQps,omitempty"`
	// example:
	//
	// nnjNQQv3ZjyYE7H4
	ProductKey  *string `json:"productKey,omitempty" xml:"productKey,omitempty"`
	ProductName *string `json:"productName,omitempty" xml:"productName,omitempty"`
	// example:
	//
	// 3dc95cac8272b86a5d10de7768d8fc41
	ProductSecret *string `json:"productSecret,omitempty" xml:"productSecret,omitempty"`
	// example:
	//
	// 383756559581697
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// example:
	//
	// 100000
	TokenCount *int64 `json:"tokenCount,omitempty" xml:"tokenCount,omitempty"`
	// example:
	//
	// 2024-03-05 06:24:27
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// example:
	//
	// 10000
	UsedToken *int64 `json:"usedToken,omitempty" xml:"usedToken,omitempty"`
	// example:
	//
	// 359687
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryProductPageResponseBodyDataData) String() string {
	return tea.Prettify(s)
}

func (s QueryProductPageResponseBodyDataData) GoString() string {
	return s.String()
}

func (s *QueryProductPageResponseBodyDataData) SetActiveLicenseCount(v int64) *QueryProductPageResponseBodyDataData {
	s.ActiveLicenseCount = &v
	return s
}

func (s *QueryProductPageResponseBodyDataData) SetApiKey(v string) *QueryProductPageResponseBodyDataData {
	s.ApiKey = &v
	return s
}

func (s *QueryProductPageResponseBodyDataData) SetCreateTime(v string) *QueryProductPageResponseBodyDataData {
	s.CreateTime = &v
	return s
}

func (s *QueryProductPageResponseBodyDataData) SetDescription(v string) *QueryProductPageResponseBodyDataData {
	s.Description = &v
	return s
}

func (s *QueryProductPageResponseBodyDataData) SetLicenseCount(v int64) *QueryProductPageResponseBodyDataData {
	s.LicenseCount = &v
	return s
}

func (s *QueryProductPageResponseBodyDataData) SetMaxQps(v int32) *QueryProductPageResponseBodyDataData {
	s.MaxQps = &v
	return s
}

func (s *QueryProductPageResponseBodyDataData) SetProductKey(v string) *QueryProductPageResponseBodyDataData {
	s.ProductKey = &v
	return s
}

func (s *QueryProductPageResponseBodyDataData) SetProductName(v string) *QueryProductPageResponseBodyDataData {
	s.ProductName = &v
	return s
}

func (s *QueryProductPageResponseBodyDataData) SetProductSecret(v string) *QueryProductPageResponseBodyDataData {
	s.ProductSecret = &v
	return s
}

func (s *QueryProductPageResponseBodyDataData) SetTenantId(v string) *QueryProductPageResponseBodyDataData {
	s.TenantId = &v
	return s
}

func (s *QueryProductPageResponseBodyDataData) SetTokenCount(v int64) *QueryProductPageResponseBodyDataData {
	s.TokenCount = &v
	return s
}

func (s *QueryProductPageResponseBodyDataData) SetUpdateTime(v string) *QueryProductPageResponseBodyDataData {
	s.UpdateTime = &v
	return s
}

func (s *QueryProductPageResponseBodyDataData) SetUsedToken(v int64) *QueryProductPageResponseBodyDataData {
	s.UsedToken = &v
	return s
}

func (s *QueryProductPageResponseBodyDataData) SetUserId(v string) *QueryProductPageResponseBodyDataData {
	s.UserId = &v
	return s
}

type QueryProductPageResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryProductPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryProductPageResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryProductPageResponse) GoString() string {
	return s.String()
}

func (s *QueryProductPageResponse) SetHeaders(v map[string]*string) *QueryProductPageResponse {
	s.Headers = v
	return s
}

func (s *QueryProductPageResponse) SetStatusCode(v int32) *QueryProductPageResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryProductPageResponse) SetBody(v *QueryProductPageResponseBody) *QueryProductPageResponse {
	s.Body = v
	return s
}

type QueryProductQuotaPageRequest struct {
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 2oImhCz4f8XCviiM
	ProductKey  *string `json:"productKey,omitempty" xml:"productKey,omitempty"`
	ProductName *string `json:"productName,omitempty" xml:"productName,omitempty"`
	// example:
	//
	// 2025-04-01 00:00:00
	PurchaseTimeEnd *string `json:"purchaseTimeEnd,omitempty" xml:"purchaseTimeEnd,omitempty"`
	// example:
	//
	// 2025-03-01 00:00:00
	PurchaseTimeStart *string `json:"purchaseTimeStart,omitempty" xml:"purchaseTimeStart,omitempty"`
	// example:
	//
	// 1
	PurchaseType *int32 `json:"purchaseType,omitempty" xml:"purchaseType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 628103740287873
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryProductQuotaPageRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryProductQuotaPageRequest) GoString() string {
	return s.String()
}

func (s *QueryProductQuotaPageRequest) SetPageIndex(v int32) *QueryProductQuotaPageRequest {
	s.PageIndex = &v
	return s
}

func (s *QueryProductQuotaPageRequest) SetPageSize(v int32) *QueryProductQuotaPageRequest {
	s.PageSize = &v
	return s
}

func (s *QueryProductQuotaPageRequest) SetProductKey(v string) *QueryProductQuotaPageRequest {
	s.ProductKey = &v
	return s
}

func (s *QueryProductQuotaPageRequest) SetProductName(v string) *QueryProductQuotaPageRequest {
	s.ProductName = &v
	return s
}

func (s *QueryProductQuotaPageRequest) SetPurchaseTimeEnd(v string) *QueryProductQuotaPageRequest {
	s.PurchaseTimeEnd = &v
	return s
}

func (s *QueryProductQuotaPageRequest) SetPurchaseTimeStart(v string) *QueryProductQuotaPageRequest {
	s.PurchaseTimeStart = &v
	return s
}

func (s *QueryProductQuotaPageRequest) SetPurchaseType(v int32) *QueryProductQuotaPageRequest {
	s.PurchaseType = &v
	return s
}

func (s *QueryProductQuotaPageRequest) SetTenantId(v string) *QueryProductQuotaPageRequest {
	s.TenantId = &v
	return s
}

func (s *QueryProductQuotaPageRequest) SetUserId(v string) *QueryProductQuotaPageRequest {
	s.UserId = &v
	return s
}

type QueryProductQuotaPageResponseBody struct {
	// example:
	//
	// success
	Code *string                                `json:"code,omitempty" xml:"code,omitempty"`
	Data *QueryProductQuotaPageResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 32B81CD6-D583-5056-A6EB-3C1107AB26C3
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"success,omitempty" xml:"success,omitempty"`
}

func (s QueryProductQuotaPageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryProductQuotaPageResponseBody) GoString() string {
	return s.String()
}

func (s *QueryProductQuotaPageResponseBody) SetCode(v string) *QueryProductQuotaPageResponseBody {
	s.Code = &v
	return s
}

func (s *QueryProductQuotaPageResponseBody) SetData(v *QueryProductQuotaPageResponseBodyData) *QueryProductQuotaPageResponseBody {
	s.Data = v
	return s
}

func (s *QueryProductQuotaPageResponseBody) SetHttpStatusCode(v string) *QueryProductQuotaPageResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryProductQuotaPageResponseBody) SetMessage(v string) *QueryProductQuotaPageResponseBody {
	s.Message = &v
	return s
}

func (s *QueryProductQuotaPageResponseBody) SetRequestId(v string) *QueryProductQuotaPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryProductQuotaPageResponseBody) SetSuccess(v string) *QueryProductQuotaPageResponseBody {
	s.Success = &v
	return s
}

type QueryProductQuotaPageResponseBodyData struct {
	Data []*QueryProductQuotaPageResponseBodyDataData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 151
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s QueryProductQuotaPageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryProductQuotaPageResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryProductQuotaPageResponseBodyData) SetData(v []*QueryProductQuotaPageResponseBodyDataData) *QueryProductQuotaPageResponseBodyData {
	s.Data = v
	return s
}

func (s *QueryProductQuotaPageResponseBodyData) SetPageIndex(v int32) *QueryProductQuotaPageResponseBodyData {
	s.PageIndex = &v
	return s
}

func (s *QueryProductQuotaPageResponseBodyData) SetPageSize(v int32) *QueryProductQuotaPageResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryProductQuotaPageResponseBodyData) SetTotal(v int64) *QueryProductQuotaPageResponseBodyData {
	s.Total = &v
	return s
}

type QueryProductQuotaPageResponseBodyDataData struct {
	// example:
	//
	// 2025-03-23 02:02:03
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 12
	Duration *int32 `json:"duration,omitempty" xml:"duration,omitempty"`
	// example:
	//
	// 2
	DurationType *int32 `json:"durationType,omitempty" xml:"durationType,omitempty"`
	// example:
	//
	// 2025-05-31 00:00:00
	ExpireTime *string `json:"expireTime,omitempty" xml:"expireTime,omitempty"`
	// ID。
	//
	// example:
	//
	// 67241348
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 0
	IfUnsubscribe *int32 `json:"ifUnsubscribe,omitempty" xml:"ifUnsubscribe,omitempty"`
	// example:
	//
	// 1
	IfUsed *int32 `json:"ifUsed,omitempty" xml:"ifUsed,omitempty"`
	// example:
	//
	// 100
	LicenseCount *int64 `json:"licenseCount,omitempty" xml:"licenseCount,omitempty"`
	// example:
	//
	// 100
	MaxQps *int32 `json:"maxQps,omitempty" xml:"maxQps,omitempty"`
	// example:
	//
	// g6RD6uvFYNZv44ky
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
	// example:
	//
	// 2oImhCz4f8XCviiM
	ProductKey  *string `json:"productKey,omitempty" xml:"productKey,omitempty"`
	ProductName *string `json:"productName,omitempty" xml:"productName,omitempty"`
	// example:
	//
	// 1
	PurchaseModel *int32 `json:"purchaseModel,omitempty" xml:"purchaseModel,omitempty"`
	// example:
	//
	// 1
	PurchaseType *int32 `json:"purchaseType,omitempty" xml:"purchaseType,omitempty"`
	// example:
	//
	// 100.0
	SettlementFee *float64 `json:"settlementFee,omitempty" xml:"settlementFee,omitempty"`
	// example:
	//
	// 217037888563265
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// example:
	//
	// 10000
	TokenNumber *int64 `json:"tokenNumber,omitempty" xml:"tokenNumber,omitempty"`
	// example:
	//
	// 1.0
	UnitPrice *float64 `json:"unitPrice,omitempty" xml:"unitPrice,omitempty"`
	// example:
	//
	// 2025-03-23 02:02:03
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// example:
	//
	// 123456
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryProductQuotaPageResponseBodyDataData) String() string {
	return tea.Prettify(s)
}

func (s QueryProductQuotaPageResponseBodyDataData) GoString() string {
	return s.String()
}

func (s *QueryProductQuotaPageResponseBodyDataData) SetCreateTime(v string) *QueryProductQuotaPageResponseBodyDataData {
	s.CreateTime = &v
	return s
}

func (s *QueryProductQuotaPageResponseBodyDataData) SetDuration(v int32) *QueryProductQuotaPageResponseBodyDataData {
	s.Duration = &v
	return s
}

func (s *QueryProductQuotaPageResponseBodyDataData) SetDurationType(v int32) *QueryProductQuotaPageResponseBodyDataData {
	s.DurationType = &v
	return s
}

func (s *QueryProductQuotaPageResponseBodyDataData) SetExpireTime(v string) *QueryProductQuotaPageResponseBodyDataData {
	s.ExpireTime = &v
	return s
}

func (s *QueryProductQuotaPageResponseBodyDataData) SetId(v int64) *QueryProductQuotaPageResponseBodyDataData {
	s.Id = &v
	return s
}

func (s *QueryProductQuotaPageResponseBodyDataData) SetIfUnsubscribe(v int32) *QueryProductQuotaPageResponseBodyDataData {
	s.IfUnsubscribe = &v
	return s
}

func (s *QueryProductQuotaPageResponseBodyDataData) SetIfUsed(v int32) *QueryProductQuotaPageResponseBodyDataData {
	s.IfUsed = &v
	return s
}

func (s *QueryProductQuotaPageResponseBodyDataData) SetLicenseCount(v int64) *QueryProductQuotaPageResponseBodyDataData {
	s.LicenseCount = &v
	return s
}

func (s *QueryProductQuotaPageResponseBodyDataData) SetMaxQps(v int32) *QueryProductQuotaPageResponseBodyDataData {
	s.MaxQps = &v
	return s
}

func (s *QueryProductQuotaPageResponseBodyDataData) SetOrderId(v string) *QueryProductQuotaPageResponseBodyDataData {
	s.OrderId = &v
	return s
}

func (s *QueryProductQuotaPageResponseBodyDataData) SetProductKey(v string) *QueryProductQuotaPageResponseBodyDataData {
	s.ProductKey = &v
	return s
}

func (s *QueryProductQuotaPageResponseBodyDataData) SetProductName(v string) *QueryProductQuotaPageResponseBodyDataData {
	s.ProductName = &v
	return s
}

func (s *QueryProductQuotaPageResponseBodyDataData) SetPurchaseModel(v int32) *QueryProductQuotaPageResponseBodyDataData {
	s.PurchaseModel = &v
	return s
}

func (s *QueryProductQuotaPageResponseBodyDataData) SetPurchaseType(v int32) *QueryProductQuotaPageResponseBodyDataData {
	s.PurchaseType = &v
	return s
}

func (s *QueryProductQuotaPageResponseBodyDataData) SetSettlementFee(v float64) *QueryProductQuotaPageResponseBodyDataData {
	s.SettlementFee = &v
	return s
}

func (s *QueryProductQuotaPageResponseBodyDataData) SetTenantId(v string) *QueryProductQuotaPageResponseBodyDataData {
	s.TenantId = &v
	return s
}

func (s *QueryProductQuotaPageResponseBodyDataData) SetTokenNumber(v int64) *QueryProductQuotaPageResponseBodyDataData {
	s.TokenNumber = &v
	return s
}

func (s *QueryProductQuotaPageResponseBodyDataData) SetUnitPrice(v float64) *QueryProductQuotaPageResponseBodyDataData {
	s.UnitPrice = &v
	return s
}

func (s *QueryProductQuotaPageResponseBodyDataData) SetUpdateTime(v string) *QueryProductQuotaPageResponseBodyDataData {
	s.UpdateTime = &v
	return s
}

func (s *QueryProductQuotaPageResponseBodyDataData) SetUserId(v string) *QueryProductQuotaPageResponseBodyDataData {
	s.UserId = &v
	return s
}

type QueryProductQuotaPageResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryProductQuotaPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryProductQuotaPageResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryProductQuotaPageResponse) GoString() string {
	return s.String()
}

func (s *QueryProductQuotaPageResponse) SetHeaders(v map[string]*string) *QueryProductQuotaPageResponse {
	s.Headers = v
	return s
}

func (s *QueryProductQuotaPageResponse) SetStatusCode(v int32) *QueryProductQuotaPageResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryProductQuotaPageResponse) SetBody(v *QueryProductQuotaPageResponseBody) *QueryProductQuotaPageResponse {
	s.Body = v
	return s
}

type QueryTokenUsageRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2025-04-02 00:00:00
	EndDate *string `json:"endDate,omitempty" xml:"endDate,omitempty"`
	// example:
	//
	// 2oImhCz4f8XCviiM
	ProductKey *string `json:"productKey,omitempty" xml:"productKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2025-04-01 00:00:00
	StartDate *string `json:"startDate,omitempty" xml:"startDate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 520539530998273
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s QueryTokenUsageRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTokenUsageRequest) GoString() string {
	return s.String()
}

func (s *QueryTokenUsageRequest) SetEndDate(v string) *QueryTokenUsageRequest {
	s.EndDate = &v
	return s
}

func (s *QueryTokenUsageRequest) SetProductKey(v string) *QueryTokenUsageRequest {
	s.ProductKey = &v
	return s
}

func (s *QueryTokenUsageRequest) SetStartDate(v string) *QueryTokenUsageRequest {
	s.StartDate = &v
	return s
}

func (s *QueryTokenUsageRequest) SetTenantId(v string) *QueryTokenUsageRequest {
	s.TenantId = &v
	return s
}

type QueryTokenUsageResponseBody struct {
	// example:
	//
	// success
	Code *string                            `json:"code,omitempty" xml:"code,omitempty"`
	Data []*QueryTokenUsageResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// B08AAA14-AD93-51F6-82AE-82AFAE9375B6
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"success,omitempty" xml:"success,omitempty"`
}

func (s QueryTokenUsageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTokenUsageResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTokenUsageResponseBody) SetCode(v string) *QueryTokenUsageResponseBody {
	s.Code = &v
	return s
}

func (s *QueryTokenUsageResponseBody) SetData(v []*QueryTokenUsageResponseBodyData) *QueryTokenUsageResponseBody {
	s.Data = v
	return s
}

func (s *QueryTokenUsageResponseBody) SetHttpStatusCode(v string) *QueryTokenUsageResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryTokenUsageResponseBody) SetMessage(v string) *QueryTokenUsageResponseBody {
	s.Message = &v
	return s
}

func (s *QueryTokenUsageResponseBody) SetRequestId(v string) *QueryTokenUsageResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTokenUsageResponseBody) SetSuccess(v string) *QueryTokenUsageResponseBody {
	s.Success = &v
	return s
}

type QueryTokenUsageResponseBodyData struct {
	// example:
	//
	// oqYVtK7DnaVj4JpbFtghJV2CZy7HwhOI0do3mf8twx9TGCMNNXYBptJ0+ULqO3xD
	ApiKey *string `json:"apiKey,omitempty" xml:"apiKey,omitempty"`
	// example:
	//
	// 1000
	InputToken *int64 `json:"inputToken,omitempty" xml:"inputToken,omitempty"`
	// example:
	//
	// 1000
	OutputToken *int64 `json:"outputToken,omitempty" xml:"outputToken,omitempty"`
	// example:
	//
	// 2oImhCz4f8XCviiM
	ProductKey  *string `json:"productKey,omitempty" xml:"productKey,omitempty"`
	ProductName *string `json:"productName,omitempty" xml:"productName,omitempty"`
	// example:
	//
	// 520539530998273
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// example:
	//
	// 2025-04-01
	UsageTime *string `json:"usageTime,omitempty" xml:"usageTime,omitempty"`
}

func (s QueryTokenUsageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryTokenUsageResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryTokenUsageResponseBodyData) SetApiKey(v string) *QueryTokenUsageResponseBodyData {
	s.ApiKey = &v
	return s
}

func (s *QueryTokenUsageResponseBodyData) SetInputToken(v int64) *QueryTokenUsageResponseBodyData {
	s.InputToken = &v
	return s
}

func (s *QueryTokenUsageResponseBodyData) SetOutputToken(v int64) *QueryTokenUsageResponseBodyData {
	s.OutputToken = &v
	return s
}

func (s *QueryTokenUsageResponseBodyData) SetProductKey(v string) *QueryTokenUsageResponseBodyData {
	s.ProductKey = &v
	return s
}

func (s *QueryTokenUsageResponseBodyData) SetProductName(v string) *QueryTokenUsageResponseBodyData {
	s.ProductName = &v
	return s
}

func (s *QueryTokenUsageResponseBodyData) SetTenantId(v string) *QueryTokenUsageResponseBodyData {
	s.TenantId = &v
	return s
}

func (s *QueryTokenUsageResponseBodyData) SetUsageTime(v string) *QueryTokenUsageResponseBodyData {
	s.UsageTime = &v
	return s
}

type QueryTokenUsageResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryTokenUsageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryTokenUsageResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTokenUsageResponse) GoString() string {
	return s.String()
}

func (s *QueryTokenUsageResponse) SetHeaders(v map[string]*string) *QueryTokenUsageResponse {
	s.Headers = v
	return s
}

func (s *QueryTokenUsageResponse) SetStatusCode(v int32) *QueryTokenUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTokenUsageResponse) SetBody(v *QueryTokenUsageResponseBody) *QueryTokenUsageResponse {
	s.Body = v
	return s
}

type RevokeChannelSignResponseBody struct {
	// example:
	//
	// success
	Code *string                            `json:"code,omitempty" xml:"code,omitempty"`
	Data *RevokeChannelSignResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 2C17015D-F916-5C2B-8C50-424DA829685E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"success,omitempty" xml:"success,omitempty"`
}

func (s RevokeChannelSignResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RevokeChannelSignResponseBody) GoString() string {
	return s.String()
}

func (s *RevokeChannelSignResponseBody) SetCode(v string) *RevokeChannelSignResponseBody {
	s.Code = &v
	return s
}

func (s *RevokeChannelSignResponseBody) SetData(v *RevokeChannelSignResponseBodyData) *RevokeChannelSignResponseBody {
	s.Data = v
	return s
}

func (s *RevokeChannelSignResponseBody) SetHttpStatusCode(v string) *RevokeChannelSignResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RevokeChannelSignResponseBody) SetMessage(v string) *RevokeChannelSignResponseBody {
	s.Message = &v
	return s
}

func (s *RevokeChannelSignResponseBody) SetRequestId(v string) *RevokeChannelSignResponseBody {
	s.RequestId = &v
	return s
}

func (s *RevokeChannelSignResponseBody) SetSuccess(v string) *RevokeChannelSignResponseBody {
	s.Success = &v
	return s
}

type RevokeChannelSignResponseBodyData struct {
	ChannelName *string `json:"channelName,omitempty" xml:"channelName,omitempty"`
	Contact     *string `json:"contact,omitempty" xml:"contact,omitempty"`
	// example:
	//
	// 2025-05-01 10:43:21
	CreateTime  *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 2025-05-01 10:43:21
	ModifyTime *string `json:"modifyTime,omitempty" xml:"modifyTime,omitempty"`
	// example:
	//
	// 13555555555
	Phone  *string `json:"phone,omitempty" xml:"phone,omitempty"`
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// example:
	//
	// revoke
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s RevokeChannelSignResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RevokeChannelSignResponseBodyData) GoString() string {
	return s.String()
}

func (s *RevokeChannelSignResponseBodyData) SetChannelName(v string) *RevokeChannelSignResponseBodyData {
	s.ChannelName = &v
	return s
}

func (s *RevokeChannelSignResponseBodyData) SetContact(v string) *RevokeChannelSignResponseBodyData {
	s.Contact = &v
	return s
}

func (s *RevokeChannelSignResponseBodyData) SetCreateTime(v string) *RevokeChannelSignResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *RevokeChannelSignResponseBodyData) SetDescription(v string) *RevokeChannelSignResponseBodyData {
	s.Description = &v
	return s
}

func (s *RevokeChannelSignResponseBodyData) SetModifyTime(v string) *RevokeChannelSignResponseBodyData {
	s.ModifyTime = &v
	return s
}

func (s *RevokeChannelSignResponseBodyData) SetPhone(v string) *RevokeChannelSignResponseBodyData {
	s.Phone = &v
	return s
}

func (s *RevokeChannelSignResponseBodyData) SetRemark(v string) *RevokeChannelSignResponseBodyData {
	s.Remark = &v
	return s
}

func (s *RevokeChannelSignResponseBodyData) SetStatus(v string) *RevokeChannelSignResponseBodyData {
	s.Status = &v
	return s
}

type RevokeChannelSignResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevokeChannelSignResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevokeChannelSignResponse) String() string {
	return tea.Prettify(s)
}

func (s RevokeChannelSignResponse) GoString() string {
	return s.String()
}

func (s *RevokeChannelSignResponse) SetHeaders(v map[string]*string) *RevokeChannelSignResponse {
	s.Headers = v
	return s
}

func (s *RevokeChannelSignResponse) SetStatusCode(v int32) *RevokeChannelSignResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeChannelSignResponse) SetBody(v *RevokeChannelSignResponseBody) *RevokeChannelSignResponse {
	s.Body = v
	return s
}

type UpdateDeviceStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 5b504f84b69b9a73d3a21a2cff05e190
	DeviceName *string `json:"deviceName,omitempty" xml:"deviceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2oImhCz4f8XCviiM
	ProductKey *string `json:"productKey,omitempty" xml:"productKey,omitempty"`
	Remark     *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
}

func (s UpdateDeviceStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeviceStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateDeviceStatusRequest) SetDeviceName(v string) *UpdateDeviceStatusRequest {
	s.DeviceName = &v
	return s
}

func (s *UpdateDeviceStatusRequest) SetProductKey(v string) *UpdateDeviceStatusRequest {
	s.ProductKey = &v
	return s
}

func (s *UpdateDeviceStatusRequest) SetRemark(v string) *UpdateDeviceStatusRequest {
	s.Remark = &v
	return s
}

func (s *UpdateDeviceStatusRequest) SetStatus(v int32) *UpdateDeviceStatusRequest {
	s.Status = &v
	return s
}

type UpdateDeviceStatusResponseBody struct {
	// example:
	//
	// success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 037E1251-1B9E-5DF5-B787-C3971A79DF89
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateDeviceStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeviceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDeviceStatusResponseBody) SetCode(v string) *UpdateDeviceStatusResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateDeviceStatusResponseBody) SetData(v bool) *UpdateDeviceStatusResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateDeviceStatusResponseBody) SetHttpStatusCode(v int32) *UpdateDeviceStatusResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateDeviceStatusResponseBody) SetMessage(v string) *UpdateDeviceStatusResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateDeviceStatusResponseBody) SetRequestId(v string) *UpdateDeviceStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDeviceStatusResponseBody) SetSuccess(v string) *UpdateDeviceStatusResponseBody {
	s.Success = &v
	return s
}

type UpdateDeviceStatusResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDeviceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDeviceStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeviceStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateDeviceStatusResponse) SetHeaders(v map[string]*string) *UpdateDeviceStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateDeviceStatusResponse) SetStatusCode(v int32) *UpdateDeviceStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDeviceStatusResponse) SetBody(v *UpdateDeviceStatusResponseBody) *UpdateDeviceStatusResponse {
	s.Body = v
	return s
}

type UpdateImageQuotaRequest struct {
	// example:
	//
	// 1
	Duration *int32 `json:"duration,omitempty" xml:"duration,omitempty"`
	// example:
	//
	// 2
	DurationType *int32 `json:"durationType,omitempty" xml:"durationType,omitempty"`
	// example:
	//
	// 100
	ImageCount *int32 `json:"imageCount,omitempty" xml:"imageCount,omitempty"`
	// example:
	//
	// 10
	LicenseCount *int64 `json:"licenseCount,omitempty" xml:"licenseCount,omitempty"`
	// example:
	//
	// 2
	PackageType *int32 `json:"packageType,omitempty" xml:"packageType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2oImhCz4f8XCviiM
	ProductKey *string `json:"productKey,omitempty" xml:"productKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PurchaseType *int32 `json:"purchaseType,omitempty" xml:"purchaseType,omitempty"`
	// example:
	//
	// 355
	RecordId *int32 `json:"recordId,omitempty" xml:"recordId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10.0
	SettlementAmount *float64 `json:"settlementAmount,omitempty" xml:"settlementAmount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 520539530998273
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1.0
	UnitPrice *float64 `json:"unitPrice,omitempty" xml:"unitPrice,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s UpdateImageQuotaRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateImageQuotaRequest) GoString() string {
	return s.String()
}

func (s *UpdateImageQuotaRequest) SetDuration(v int32) *UpdateImageQuotaRequest {
	s.Duration = &v
	return s
}

func (s *UpdateImageQuotaRequest) SetDurationType(v int32) *UpdateImageQuotaRequest {
	s.DurationType = &v
	return s
}

func (s *UpdateImageQuotaRequest) SetImageCount(v int32) *UpdateImageQuotaRequest {
	s.ImageCount = &v
	return s
}

func (s *UpdateImageQuotaRequest) SetLicenseCount(v int64) *UpdateImageQuotaRequest {
	s.LicenseCount = &v
	return s
}

func (s *UpdateImageQuotaRequest) SetPackageType(v int32) *UpdateImageQuotaRequest {
	s.PackageType = &v
	return s
}

func (s *UpdateImageQuotaRequest) SetProductKey(v string) *UpdateImageQuotaRequest {
	s.ProductKey = &v
	return s
}

func (s *UpdateImageQuotaRequest) SetPurchaseType(v int32) *UpdateImageQuotaRequest {
	s.PurchaseType = &v
	return s
}

func (s *UpdateImageQuotaRequest) SetRecordId(v int32) *UpdateImageQuotaRequest {
	s.RecordId = &v
	return s
}

func (s *UpdateImageQuotaRequest) SetSettlementAmount(v float64) *UpdateImageQuotaRequest {
	s.SettlementAmount = &v
	return s
}

func (s *UpdateImageQuotaRequest) SetTenantId(v string) *UpdateImageQuotaRequest {
	s.TenantId = &v
	return s
}

func (s *UpdateImageQuotaRequest) SetUnitPrice(v float64) *UpdateImageQuotaRequest {
	s.UnitPrice = &v
	return s
}

func (s *UpdateImageQuotaRequest) SetUserId(v string) *UpdateImageQuotaRequest {
	s.UserId = &v
	return s
}

type UpdateImageQuotaResponseBody struct {
	// example:
	//
	// success
	Code *string                           `json:"code,omitempty" xml:"code,omitempty"`
	Data *UpdateImageQuotaResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 28BD530A-C469-5CF9-9F4E-DA0AF0A1AC73
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateImageQuotaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateImageQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateImageQuotaResponseBody) SetCode(v string) *UpdateImageQuotaResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateImageQuotaResponseBody) SetData(v *UpdateImageQuotaResponseBodyData) *UpdateImageQuotaResponseBody {
	s.Data = v
	return s
}

func (s *UpdateImageQuotaResponseBody) SetHttpStatusCode(v string) *UpdateImageQuotaResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateImageQuotaResponseBody) SetMessage(v string) *UpdateImageQuotaResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateImageQuotaResponseBody) SetRequestId(v string) *UpdateImageQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateImageQuotaResponseBody) SetSuccess(v string) *UpdateImageQuotaResponseBody {
	s.Success = &v
	return s
}

type UpdateImageQuotaResponseBodyData struct {
	// example:
	//
	// FlUHDd8ol1yRmA92
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
	// example:
	//
	// 2oImhCz4f8XCviiM
	ProductKey  *string `json:"productKey,omitempty" xml:"productKey,omitempty"`
	ProductName *string `json:"productName,omitempty" xml:"productName,omitempty"`
	// example:
	//
	// 520539530998273
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// example:
	//
	// 123456
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s UpdateImageQuotaResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpdateImageQuotaResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateImageQuotaResponseBodyData) SetOrderId(v string) *UpdateImageQuotaResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *UpdateImageQuotaResponseBodyData) SetProductKey(v string) *UpdateImageQuotaResponseBodyData {
	s.ProductKey = &v
	return s
}

func (s *UpdateImageQuotaResponseBodyData) SetProductName(v string) *UpdateImageQuotaResponseBodyData {
	s.ProductName = &v
	return s
}

func (s *UpdateImageQuotaResponseBodyData) SetTenantId(v string) *UpdateImageQuotaResponseBodyData {
	s.TenantId = &v
	return s
}

func (s *UpdateImageQuotaResponseBodyData) SetUserId(v string) *UpdateImageQuotaResponseBodyData {
	s.UserId = &v
	return s
}

type UpdateImageQuotaResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateImageQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateImageQuotaResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateImageQuotaResponse) GoString() string {
	return s.String()
}

func (s *UpdateImageQuotaResponse) SetHeaders(v map[string]*string) *UpdateImageQuotaResponse {
	s.Headers = v
	return s
}

func (s *UpdateImageQuotaResponse) SetStatusCode(v int32) *UpdateImageQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateImageQuotaResponse) SetBody(v *UpdateImageQuotaResponseBody) *UpdateImageQuotaResponse {
	s.Body = v
	return s
}

type UpdateQuotaRequest struct {
	// example:
	//
	// 1
	Duration *int32 `json:"duration,omitempty" xml:"duration,omitempty"`
	// example:
	//
	// 2
	DurationType *int32 `json:"durationType,omitempty" xml:"durationType,omitempty"`
	// example:
	//
	// 10
	LicenseCount *int64 `json:"licenseCount,omitempty" xml:"licenseCount,omitempty"`
	// example:
	//
	// 100
	MaxQps *int32 `json:"maxQps,omitempty" xml:"maxQps,omitempty"`
	// example:
	//
	// 2
	PackageType *int32 `json:"packageType,omitempty" xml:"packageType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// nnjNQQv3ZjyYE7H4
	ProductKey *string `json:"productKey,omitempty" xml:"productKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PurchaseType *int32 `json:"purchaseType,omitempty" xml:"purchaseType,omitempty"`
	// example:
	//
	// 421
	RecordId *int32 `json:"recordId,omitempty" xml:"recordId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10.0
	SettlementAmount *float64 `json:"settlementAmount,omitempty" xml:"settlementAmount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 520539530998273
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// example:
	//
	// 10000
	TokenNumber *int64 `json:"tokenNumber,omitempty" xml:"tokenNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1.0
	UnitPrice *float64 `json:"unitPrice,omitempty" xml:"unitPrice,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s UpdateQuotaRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateQuotaRequest) GoString() string {
	return s.String()
}

func (s *UpdateQuotaRequest) SetDuration(v int32) *UpdateQuotaRequest {
	s.Duration = &v
	return s
}

func (s *UpdateQuotaRequest) SetDurationType(v int32) *UpdateQuotaRequest {
	s.DurationType = &v
	return s
}

func (s *UpdateQuotaRequest) SetLicenseCount(v int64) *UpdateQuotaRequest {
	s.LicenseCount = &v
	return s
}

func (s *UpdateQuotaRequest) SetMaxQps(v int32) *UpdateQuotaRequest {
	s.MaxQps = &v
	return s
}

func (s *UpdateQuotaRequest) SetPackageType(v int32) *UpdateQuotaRequest {
	s.PackageType = &v
	return s
}

func (s *UpdateQuotaRequest) SetProductKey(v string) *UpdateQuotaRequest {
	s.ProductKey = &v
	return s
}

func (s *UpdateQuotaRequest) SetPurchaseType(v int32) *UpdateQuotaRequest {
	s.PurchaseType = &v
	return s
}

func (s *UpdateQuotaRequest) SetRecordId(v int32) *UpdateQuotaRequest {
	s.RecordId = &v
	return s
}

func (s *UpdateQuotaRequest) SetSettlementAmount(v float64) *UpdateQuotaRequest {
	s.SettlementAmount = &v
	return s
}

func (s *UpdateQuotaRequest) SetTenantId(v string) *UpdateQuotaRequest {
	s.TenantId = &v
	return s
}

func (s *UpdateQuotaRequest) SetTokenNumber(v int64) *UpdateQuotaRequest {
	s.TokenNumber = &v
	return s
}

func (s *UpdateQuotaRequest) SetUnitPrice(v float64) *UpdateQuotaRequest {
	s.UnitPrice = &v
	return s
}

func (s *UpdateQuotaRequest) SetUserId(v string) *UpdateQuotaRequest {
	s.UserId = &v
	return s
}

type UpdateQuotaResponseBody struct {
	// example:
	//
	// success
	Code *string                      `json:"code,omitempty" xml:"code,omitempty"`
	Data *UpdateQuotaResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 89946BAA-E4E1-5114-8A5E-A542FEDC9B16
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateQuotaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateQuotaResponseBody) SetCode(v string) *UpdateQuotaResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateQuotaResponseBody) SetData(v *UpdateQuotaResponseBodyData) *UpdateQuotaResponseBody {
	s.Data = v
	return s
}

func (s *UpdateQuotaResponseBody) SetHttpStatusCode(v string) *UpdateQuotaResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateQuotaResponseBody) SetMessage(v string) *UpdateQuotaResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateQuotaResponseBody) SetRequestId(v string) *UpdateQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateQuotaResponseBody) SetSuccess(v string) *UpdateQuotaResponseBody {
	s.Success = &v
	return s
}

type UpdateQuotaResponseBodyData struct {
	// example:
	//
	// g6RD6uvFYNZv44ky
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
	// example:
	//
	// 2oImhCz4f8XCviiM
	ProductKey  *string `json:"productKey,omitempty" xml:"productKey,omitempty"`
	ProductName *string `json:"productName,omitempty" xml:"productName,omitempty"`
	// example:
	//
	// 520539530998273
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// example:
	//
	// 123456
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s UpdateQuotaResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpdateQuotaResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateQuotaResponseBodyData) SetOrderId(v string) *UpdateQuotaResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *UpdateQuotaResponseBodyData) SetProductKey(v string) *UpdateQuotaResponseBodyData {
	s.ProductKey = &v
	return s
}

func (s *UpdateQuotaResponseBodyData) SetProductName(v string) *UpdateQuotaResponseBodyData {
	s.ProductName = &v
	return s
}

func (s *UpdateQuotaResponseBodyData) SetTenantId(v string) *UpdateQuotaResponseBodyData {
	s.TenantId = &v
	return s
}

func (s *UpdateQuotaResponseBodyData) SetUserId(v string) *UpdateQuotaResponseBodyData {
	s.UserId = &v
	return s
}

type UpdateQuotaResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateQuotaResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateQuotaResponse) GoString() string {
	return s.String()
}

func (s *UpdateQuotaResponse) SetHeaders(v map[string]*string) *UpdateQuotaResponse {
	s.Headers = v
	return s
}

func (s *UpdateQuotaResponse) SetStatusCode(v int32) *UpdateQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateQuotaResponse) SetBody(v *UpdateQuotaResponseBody) *UpdateQuotaResponse {
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
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("bailianmodelonchip"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

// Summary:
//
// 创建渠道签约申请
//
// @param request - CreateChannelSignRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateChannelSignResponse
func (client *Client) CreateChannelSignWithOptions(request *CreateChannelSignRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateChannelSignResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChannelName)) {
		body["channelName"] = request.ChannelName
	}

	if !tea.BoolValue(util.IsUnset(request.Contact)) {
		body["contact"] = request.Contact
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Phone)) {
		body["phone"] = request.Phone
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["remark"] = request.Remark
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateChannelSign"),
		Version:     tea.String("2024-08-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/open/api/v1/channel/create"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateChannelSignResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建渠道签约申请
//
// @param request - CreateChannelSignRequest
//
// @return CreateChannelSignResponse
func (client *Client) CreateChannelSign(request *CreateChannelSignRequest) (_result *CreateChannelSignResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateChannelSignResponse{}
	_body, _err := client.CreateChannelSignWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建产品
//
// @param request - CreateProductRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateProductResponse
func (client *Client) CreateProductWithOptions(request *CreateProductRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateProductResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Euid)) {
		body["euid"] = request.Euid
	}

	if !tea.BoolValue(util.IsUnset(request.ProductName)) {
		body["productName"] = request.ProductName
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["tenantId"] = request.TenantId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateProduct"),
		Version:     tea.String("2024-08-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/open/api/v1/product/create"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateProductResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建产品
//
// @param request - CreateProductRequest
//
// @return CreateProductResponse
func (client *Client) CreateProduct(request *CreateProductRequest) (_result *CreateProductResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateProductResponse{}
	_body, _err := client.CreateProductWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除产品
//
// @param request - DeleteProductRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteProductResponse
func (client *Client) DeleteProductWithOptions(request *DeleteProductRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteProductResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		body["productKey"] = request.ProductKey
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["tenantId"] = request.TenantId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteProduct"),
		Version:     tea.String("2024-08-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/open/api/v1/product/delete"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteProductResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除产品
//
// @param request - DeleteProductRequest
//
// @return DeleteProductResponse
func (client *Client) DeleteProduct(request *DeleteProductRequest) (_result *DeleteProductResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteProductResponse{}
	_body, _err := client.DeleteProductWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 设备注册
//
// @param request - DeviceRegisterRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeviceRegisterResponse
func (client *Client) DeviceRegisterWithOptions(request *DeviceRegisterRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeviceRegisterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["productKey"] = request.ProductKey
	}

	if !tea.BoolValue(util.IsUnset(request.RequestTime)) {
		query["requestTime"] = request.RequestTime
	}

	if !tea.BoolValue(util.IsUnset(request.Signature)) {
		query["signature"] = request.Signature
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Nonce)) {
		body["nonce"] = request.Nonce
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeviceRegister"),
		Version:     tea.String("2024-08-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/open/api/device/v1/register"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeviceRegisterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设备注册
//
// @param request - DeviceRegisterRequest
//
// @return DeviceRegisterResponse
func (client *Client) DeviceRegister(request *DeviceRegisterRequest) (_result *DeviceRegisterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeviceRegisterResponse{}
	_body, _err := client.DeviceRegisterWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询渠道签约申请
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetChannelSignResponse
func (client *Client) GetChannelSignWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetChannelSignResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetChannelSign"),
		Version:     tea.String("2024-08-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/open/api/v1/channel/get"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetChannelSignResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询渠道签约申请
//
// @return GetChannelSignResponse
func (client *Client) GetChannelSign() (_result *GetChannelSignResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetChannelSignResponse{}
	_body, _err := client.GetChannelSignWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取额度信息
//
// @param request - GetQuotaInfoRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQuotaInfoResponse
func (client *Client) GetQuotaInfoWithOptions(request *GetQuotaInfoRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetQuotaInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RecordId)) {
		body["recordId"] = request.RecordId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetQuotaInfo"),
		Version:     tea.String("2024-08-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/open/api/v1/quota/get"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetQuotaInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取额度信息
//
// @param request - GetQuotaInfoRequest
//
// @return GetQuotaInfoResponse
func (client *Client) GetQuotaInfo(request *GetQuotaInfoRequest) (_result *GetQuotaInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetQuotaInfoResponse{}
	_body, _err := client.GetQuotaInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取网关校验Token
//
// @param request - GetTokenRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTokenResponse
func (client *Client) GetTokenWithOptions(request *GetTokenRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		body["deviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.Nonce)) {
		body["nonce"] = request.Nonce
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		body["productKey"] = request.ProductKey
	}

	if !tea.BoolValue(util.IsUnset(request.RequestTime)) {
		body["requestTime"] = request.RequestTime
	}

	if !tea.BoolValue(util.IsUnset(request.Signature)) {
		body["signature"] = request.Signature
	}

	if !tea.BoolValue(util.IsUnset(request.TokenKey)) {
		body["tokenKey"] = request.TokenKey
	}

	if !tea.BoolValue(util.IsUnset(request.TokenType)) {
		body["tokenType"] = request.TokenType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetToken"),
		Version:     tea.String("2024-08-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/open/api/auth/v1/token/get"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取网关校验Token
//
// @param request - GetTokenRequest
//
// @return GetTokenResponse
func (client *Client) GetToken(request *GetTokenRequest) (_result *GetTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTokenResponse{}
	_body, _err := client.GetTokenWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 半托管大模型应用请求
//
// @param tmpReq - HalfLLMAppCallRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HalfLLMAppCallResponse
func (client *Client) HalfLLMAppCallWithOptions(tmpReq *HalfLLMAppCallRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *HalfLLMAppCallResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &HalfLLMAppCallShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.BizParam)) {
		request.BizParamShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BizParam, tea.String("bizParam"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.ModelTypeList)) {
		request.ModelTypeListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ModelTypeList, tea.String("modelTypeList"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["appId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.BizParamShrink)) {
		body["bizParam"] = request.BizParamShrink
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		body["deviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.ModelTypeListShrink)) {
		body["modelTypeList"] = request.ModelTypeListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		body["productKey"] = request.ProductKey
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		body["sessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["tenantId"] = request.TenantId
	}

	if !tea.BoolValue(util.IsUnset(request.Text)) {
		body["text"] = request.Text
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("HalfLLMAppCall"),
		Version:     tea.String("2024-08-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/open/api/device/v1/half/llm/app/call"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &HalfLLMAppCallResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 半托管大模型应用请求
//
// @param request - HalfLLMAppCallRequest
//
// @return HalfLLMAppCallResponse
func (client *Client) HalfLLMAppCall(request *HalfLLMAppCallRequest) (_result *HalfLLMAppCallResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &HalfLLMAppCallResponse{}
	_body, _err := client.HalfLLMAppCallWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 半托管大模型流式文本对话
//
// @param request - HalfLLMChatRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HalfLLMChatResponse
func (client *Client) HalfLLMChatWithOptions(request *HalfLLMChatRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *HalfLLMChatResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		body["deviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.EnableSearch)) {
		body["enableSearch"] = request.EnableSearch
	}

	if !tea.BoolValue(util.IsUnset(request.InputText)) {
		body["inputText"] = request.InputText
	}

	if !tea.BoolValue(util.IsUnset(request.Model)) {
		body["model"] = request.Model
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		body["productKey"] = request.ProductKey
	}

	if !tea.BoolValue(util.IsUnset(request.Prompt)) {
		body["prompt"] = request.Prompt
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		body["sessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.Stream)) {
		body["stream"] = request.Stream
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["tenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("HalfLLMChat"),
		Version:     tea.String("2024-08-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/open/api/device/v1/half/llm/chat"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &HalfLLMChatResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 半托管大模型流式文本对话
//
// @param request - HalfLLMChatRequest
//
// @return HalfLLMChatResponse
func (client *Client) HalfLLMChat(request *HalfLLMChatRequest) (_result *HalfLLMChatResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &HalfLLMChatResponse{}
	_body, _err := client.HalfLLMChatWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 半托管大模型文本合成语音
//
// @param request - HalfLLMImageChatRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HalfLLMImageChatResponse
func (client *Client) HalfLLMImageChatWithOptions(request *HalfLLMImageChatRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *HalfLLMImageChatResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		body["deviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.EnableSearch)) {
		body["enableSearch"] = request.EnableSearch
	}

	if !tea.BoolValue(util.IsUnset(request.ImageUrl)) {
		body["imageUrl"] = request.ImageUrl
	}

	if !tea.BoolValue(util.IsUnset(request.InputText)) {
		body["inputText"] = request.InputText
	}

	if !tea.BoolValue(util.IsUnset(request.Model)) {
		body["model"] = request.Model
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		body["productKey"] = request.ProductKey
	}

	if !tea.BoolValue(util.IsUnset(request.Prompt)) {
		body["prompt"] = request.Prompt
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		body["sessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["tenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("HalfLLMImageChat"),
		Version:     tea.String("2024-08-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/open/api/device/v1/half/llm/image/chat"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &HalfLLMImageChatResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 半托管大模型文本合成语音
//
// @param request - HalfLLMImageChatRequest
//
// @return HalfLLMImageChatResponse
func (client *Client) HalfLLMImageChat(request *HalfLLMImageChatRequest) (_result *HalfLLMImageChatResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &HalfLLMImageChatResponse{}
	_body, _err := client.HalfLLMImageChatWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 半托管大模型图片识别
//
// @param request - HalfLLMImageOcrRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HalfLLMImageOcrResponse
func (client *Client) HalfLLMImageOcrWithOptions(request *HalfLLMImageOcrRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *HalfLLMImageOcrResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		body["deviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.ImageUrl)) {
		body["imageUrl"] = request.ImageUrl
	}

	if !tea.BoolValue(util.IsUnset(request.Model)) {
		body["model"] = request.Model
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		body["productKey"] = request.ProductKey
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["tenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("HalfLLMImageOcr"),
		Version:     tea.String("2024-08-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/open/api/device/v1/half/llm/image/ocr"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &HalfLLMImageOcrResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 半托管大模型图片识别
//
// @param request - HalfLLMImageOcrRequest
//
// @return HalfLLMImageOcrResponse
func (client *Client) HalfLLMImageOcr(request *HalfLLMImageOcrRequest) (_result *HalfLLMImageOcrResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &HalfLLMImageOcrResponse{}
	_body, _err := client.HalfLLMImageOcrWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 半托管大模型文本合成语音
//
// @param request - HalfLLMTTSChatRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HalfLLMTTSChatResponse
func (client *Client) HalfLLMTTSChatWithOptions(request *HalfLLMTTSChatRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *HalfLLMTTSChatResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		body["deviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.EnableSearch)) {
		body["enableSearch"] = request.EnableSearch
	}

	if !tea.BoolValue(util.IsUnset(request.Format)) {
		body["format"] = request.Format
	}

	if !tea.BoolValue(util.IsUnset(request.Model)) {
		body["model"] = request.Model
	}

	if !tea.BoolValue(util.IsUnset(request.PitchRate)) {
		body["pitchRate"] = request.PitchRate
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		body["productKey"] = request.ProductKey
	}

	if !tea.BoolValue(util.IsUnset(request.Prompt)) {
		body["prompt"] = request.Prompt
	}

	if !tea.BoolValue(util.IsUnset(request.SampleRate)) {
		body["sampleRate"] = request.SampleRate
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		body["sessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.SpeechRate)) {
		body["speechRate"] = request.SpeechRate
	}

	if !tea.BoolValue(util.IsUnset(request.Stream)) {
		body["stream"] = request.Stream
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["tenantId"] = request.TenantId
	}

	if !tea.BoolValue(util.IsUnset(request.Text)) {
		body["text"] = request.Text
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		body["url"] = request.Url
	}

	if !tea.BoolValue(util.IsUnset(request.Voice)) {
		body["voice"] = request.Voice
	}

	if !tea.BoolValue(util.IsUnset(request.Volume)) {
		body["volume"] = request.Volume
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("HalfLLMTTSChat"),
		Version:     tea.String("2024-08-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/open/api/device/v1/half/llm/tts/chat"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &HalfLLMTTSChatResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 半托管大模型文本合成语音
//
// @param request - HalfLLMTTSChatRequest
//
// @return HalfLLMTTSChatResponse
func (client *Client) HalfLLMTTSChat(request *HalfLLMTTSChatRequest) (_result *HalfLLMTTSChatResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &HalfLLMTTSChatResponse{}
	_body, _err := client.HalfLLMTTSChatWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 设备列表分页查询
//
// @param request - QueryDevicePageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDevicePageResponse
func (client *Client) QueryDevicePageWithOptions(request *QueryDevicePageRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryDevicePageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		body["deviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.DisableStatus)) {
		body["disableStatus"] = request.DisableStatus
	}

	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		body["pageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		body["productKey"] = request.ProductKey
	}

	if !tea.BoolValue(util.IsUnset(request.ProductName)) {
		body["productName"] = request.ProductName
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryDevicePage"),
		Version:     tea.String("2024-08-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/open/api/device/v1/page"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryDevicePageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设备列表分页查询
//
// @param request - QueryDevicePageRequest
//
// @return QueryDevicePageResponse
func (client *Client) QueryDevicePage(request *QueryDevicePageRequest) (_result *QueryDevicePageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryDevicePageResponse{}
	_body, _err := client.QueryDevicePageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 分页查询产品
//
// @param request - QueryProductPageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryProductPageResponse
func (client *Client) QueryProductPageWithOptions(request *QueryProductPageRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryProductPageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ModelType)) {
		body["modelType"] = request.ModelType
	}

	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		body["pageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProductName)) {
		body["productName"] = request.ProductName
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["tenantId"] = request.TenantId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryProductPage"),
		Version:     tea.String("2024-08-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/open/api/v1/product/page"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryProductPageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分页查询产品
//
// @param request - QueryProductPageRequest
//
// @return QueryProductPageResponse
func (client *Client) QueryProductPage(request *QueryProductPageRequest) (_result *QueryProductPageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryProductPageResponse{}
	_body, _err := client.QueryProductPageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 分页查询产品额度
//
// @param request - QueryProductQuotaPageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryProductQuotaPageResponse
func (client *Client) QueryProductQuotaPageWithOptions(request *QueryProductQuotaPageRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryProductQuotaPageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		body["pageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		body["productKey"] = request.ProductKey
	}

	if !tea.BoolValue(util.IsUnset(request.ProductName)) {
		body["productName"] = request.ProductName
	}

	if !tea.BoolValue(util.IsUnset(request.PurchaseTimeEnd)) {
		body["purchaseTimeEnd"] = request.PurchaseTimeEnd
	}

	if !tea.BoolValue(util.IsUnset(request.PurchaseTimeStart)) {
		body["purchaseTimeStart"] = request.PurchaseTimeStart
	}

	if !tea.BoolValue(util.IsUnset(request.PurchaseType)) {
		body["purchaseType"] = request.PurchaseType
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["tenantId"] = request.TenantId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryProductQuotaPage"),
		Version:     tea.String("2024-08-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/open/api/v1/product/quotaPage"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryProductQuotaPageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分页查询产品额度
//
// @param request - QueryProductQuotaPageRequest
//
// @return QueryProductQuotaPageResponse
func (client *Client) QueryProductQuotaPage(request *QueryProductQuotaPageRequest) (_result *QueryProductQuotaPageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryProductQuotaPageResponse{}
	_body, _err := client.QueryProductQuotaPageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询token使用量
//
// @param request - QueryTokenUsageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryTokenUsageResponse
func (client *Client) QueryTokenUsageWithOptions(request *QueryTokenUsageRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryTokenUsageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		body["endDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		body["productKey"] = request.ProductKey
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		body["startDate"] = request.StartDate
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["tenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryTokenUsage"),
		Version:     tea.String("2024-08-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/open/api/v1/token/usage/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryTokenUsageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询token使用量
//
// @param request - QueryTokenUsageRequest
//
// @return QueryTokenUsageResponse
func (client *Client) QueryTokenUsage(request *QueryTokenUsageRequest) (_result *QueryTokenUsageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryTokenUsageResponse{}
	_body, _err := client.QueryTokenUsageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 撤销渠道签约申请
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeChannelSignResponse
func (client *Client) RevokeChannelSignWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *RevokeChannelSignResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("RevokeChannelSign"),
		Version:     tea.String("2024-08-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/open/api/v1/channel/revoke"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RevokeChannelSignResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 撤销渠道签约申请
//
// @return RevokeChannelSignResponse
func (client *Client) RevokeChannelSign() (_result *RevokeChannelSignResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RevokeChannelSignResponse{}
	_body, _err := client.RevokeChannelSignWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改设备状态
//
// @param request - UpdateDeviceStatusRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDeviceStatusResponse
func (client *Client) UpdateDeviceStatusWithOptions(request *UpdateDeviceStatusRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateDeviceStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		body["deviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		body["productKey"] = request.ProductKey
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateDeviceStatus"),
		Version:     tea.String("2024-08-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/open/api/device/v1/update/status"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateDeviceStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改设备状态
//
// @param request - UpdateDeviceStatusRequest
//
// @return UpdateDeviceStatusResponse
func (client *Client) UpdateDeviceStatus(request *UpdateDeviceStatusRequest) (_result *UpdateDeviceStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateDeviceStatusResponse{}
	_body, _err := client.UpdateDeviceStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改图片模型额度
//
// @param request - UpdateImageQuotaRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateImageQuotaResponse
func (client *Client) UpdateImageQuotaWithOptions(request *UpdateImageQuotaRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateImageQuotaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Duration)) {
		body["duration"] = request.Duration
	}

	if !tea.BoolValue(util.IsUnset(request.DurationType)) {
		body["durationType"] = request.DurationType
	}

	if !tea.BoolValue(util.IsUnset(request.ImageCount)) {
		body["imageCount"] = request.ImageCount
	}

	if !tea.BoolValue(util.IsUnset(request.LicenseCount)) {
		body["licenseCount"] = request.LicenseCount
	}

	if !tea.BoolValue(util.IsUnset(request.PackageType)) {
		body["packageType"] = request.PackageType
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		body["productKey"] = request.ProductKey
	}

	if !tea.BoolValue(util.IsUnset(request.PurchaseType)) {
		body["purchaseType"] = request.PurchaseType
	}

	if !tea.BoolValue(util.IsUnset(request.RecordId)) {
		body["recordId"] = request.RecordId
	}

	if !tea.BoolValue(util.IsUnset(request.SettlementAmount)) {
		body["settlementAmount"] = request.SettlementAmount
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["tenantId"] = request.TenantId
	}

	if !tea.BoolValue(util.IsUnset(request.UnitPrice)) {
		body["unitPrice"] = request.UnitPrice
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateImageQuota"),
		Version:     tea.String("2024-08-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/open/api/v1/quota/update/image"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateImageQuotaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改图片模型额度
//
// @param request - UpdateImageQuotaRequest
//
// @return UpdateImageQuotaResponse
func (client *Client) UpdateImageQuota(request *UpdateImageQuotaRequest) (_result *UpdateImageQuotaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateImageQuotaResponse{}
	_body, _err := client.UpdateImageQuotaWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改额度
//
// @param request - UpdateQuotaRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateQuotaResponse
func (client *Client) UpdateQuotaWithOptions(request *UpdateQuotaRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateQuotaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Duration)) {
		body["duration"] = request.Duration
	}

	if !tea.BoolValue(util.IsUnset(request.DurationType)) {
		body["durationType"] = request.DurationType
	}

	if !tea.BoolValue(util.IsUnset(request.LicenseCount)) {
		body["licenseCount"] = request.LicenseCount
	}

	if !tea.BoolValue(util.IsUnset(request.MaxQps)) {
		body["maxQps"] = request.MaxQps
	}

	if !tea.BoolValue(util.IsUnset(request.PackageType)) {
		body["packageType"] = request.PackageType
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		body["productKey"] = request.ProductKey
	}

	if !tea.BoolValue(util.IsUnset(request.PurchaseType)) {
		body["purchaseType"] = request.PurchaseType
	}

	if !tea.BoolValue(util.IsUnset(request.RecordId)) {
		body["recordId"] = request.RecordId
	}

	if !tea.BoolValue(util.IsUnset(request.SettlementAmount)) {
		body["settlementAmount"] = request.SettlementAmount
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["tenantId"] = request.TenantId
	}

	if !tea.BoolValue(util.IsUnset(request.TokenNumber)) {
		body["tokenNumber"] = request.TokenNumber
	}

	if !tea.BoolValue(util.IsUnset(request.UnitPrice)) {
		body["unitPrice"] = request.UnitPrice
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateQuota"),
		Version:     tea.String("2024-08-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/open/api/v1/quota/update"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateQuotaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改额度
//
// @param request - UpdateQuotaRequest
//
// @return UpdateQuotaResponse
func (client *Client) UpdateQuota(request *UpdateQuotaRequest) (_result *UpdateQuotaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateQuotaResponse{}
	_body, _err := client.UpdateQuotaWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
