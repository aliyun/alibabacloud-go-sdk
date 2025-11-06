// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGatewayAuthConsumerDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetGatewayAuthConsumerDetailResponseBody
	GetCode() *int32
	SetData(v *GetGatewayAuthConsumerDetailResponseBodyData) *GetGatewayAuthConsumerDetailResponseBody
	GetData() *GetGatewayAuthConsumerDetailResponseBodyData
	SetDynamicCode(v string) *GetGatewayAuthConsumerDetailResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *GetGatewayAuthConsumerDetailResponseBody
	GetDynamicMessage() *string
	SetErrorCode(v string) *GetGatewayAuthConsumerDetailResponseBody
	GetErrorCode() *string
	SetHttpStatusCode(v int32) *GetGatewayAuthConsumerDetailResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetGatewayAuthConsumerDetailResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetGatewayAuthConsumerDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetGatewayAuthConsumerDetailResponseBody
	GetSuccess() *bool
}

type GetGatewayAuthConsumerDetailResponseBody struct {
	// The status code. A value of 200 is returned if the request is successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *GetGatewayAuthConsumerDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The dynamic part in the error message.
	//
	// example:
	//
	// code
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic part in the error message. This parameter is used to replace the \\*\\*%s\\*\\	- variable in the **ErrMessage*	- parameter.
	//
	// >  If the return value of the **ErrMessage*	- parameter is **The Value of Input Parameter %s is not valid*	- and the return value of the **DynamicMessage*	- parameter is **DtsJobId**, the specified **DtsJobId*	- parameter is invalid.
	//
	// example:
	//
	// The specified parameter is invalid.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error code that is returned.
	//
	// example:
	//
	// mse-100-000
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// The request is successfully processed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 316F5F64-F73D-42DC-8632-01E308B6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`: The request was successful.
	//
	// 	- `false`: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetGatewayAuthConsumerDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayAuthConsumerDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetGatewayAuthConsumerDetailResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetGatewayAuthConsumerDetailResponseBody) GetData() *GetGatewayAuthConsumerDetailResponseBodyData {
	return s.Data
}

func (s *GetGatewayAuthConsumerDetailResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *GetGatewayAuthConsumerDetailResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *GetGatewayAuthConsumerDetailResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetGatewayAuthConsumerDetailResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetGatewayAuthConsumerDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetGatewayAuthConsumerDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetGatewayAuthConsumerDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetGatewayAuthConsumerDetailResponseBody) SetCode(v int32) *GetGatewayAuthConsumerDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetGatewayAuthConsumerDetailResponseBody) SetData(v *GetGatewayAuthConsumerDetailResponseBodyData) *GetGatewayAuthConsumerDetailResponseBody {
	s.Data = v
	return s
}

func (s *GetGatewayAuthConsumerDetailResponseBody) SetDynamicCode(v string) *GetGatewayAuthConsumerDetailResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetGatewayAuthConsumerDetailResponseBody) SetDynamicMessage(v string) *GetGatewayAuthConsumerDetailResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetGatewayAuthConsumerDetailResponseBody) SetErrorCode(v string) *GetGatewayAuthConsumerDetailResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetGatewayAuthConsumerDetailResponseBody) SetHttpStatusCode(v int32) *GetGatewayAuthConsumerDetailResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetGatewayAuthConsumerDetailResponseBody) SetMessage(v string) *GetGatewayAuthConsumerDetailResponseBody {
	s.Message = &v
	return s
}

func (s *GetGatewayAuthConsumerDetailResponseBody) SetRequestId(v string) *GetGatewayAuthConsumerDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGatewayAuthConsumerDetailResponseBody) SetSuccess(v bool) *GetGatewayAuthConsumerDetailResponseBody {
	s.Success = &v
	return s
}

func (s *GetGatewayAuthConsumerDetailResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetGatewayAuthConsumerDetailResponseBodyData struct {
	// The status of the consumer. Valid values:
	//
	// 	- true: enabled
	//
	// 	- false: disabled
	//
	// example:
	//
	// true
	ConsumerStatus *bool `json:"ConsumerStatus,omitempty" xml:"ConsumerStatus,omitempty"`
	// The description.
	//
	// example:
	//
	// Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The encryption type. Valid values:
	//
	// 	- RSA
	//
	// 	- OCT
	//
	// example:
	//
	// RSA
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// The unique ID of the gateway.
	//
	// example:
	//
	// gw-7ea3da97b96543e19f6c597cd4a9****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The time when the consumer authentication record was created.
	//
	// example:
	//
	// 2031-03-30 02:35:12
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the consumer authentication record was modified.
	//
	// example:
	//
	// 2023-02-01 14:17:57
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the consumer.
	//
	// example:
	//
	// 12
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The JWT public key. The JSON format is supported.
	//
	// example:
	//
	// {
	//
	//       "keys": [
	//
	//             {
	//
	//                   "e": "AQAB",
	//
	//                   "kid": "DHFbpoIUqrY8t2zpA2qXfCmr5VO5ZEr4RzHU_-envvQ",
	//
	//                   "kty": "RSA",
	//
	//                   "n": "xAE7eB6qugXyCAG3yhh7pkDkT65pHymX-P7KfIupjf59vsdo91bSP9C8H07pSAGQO1MV_xFj9VswgsCg4R6otmg5PV2He95lZdHtOcU5DXIg_pbhLdKXbi66GlVeK6ABZOUW3WYtnNHD-91gVuoeJT_DwtGGcp4ignkgXfkiEm4sw-4sfb4qdt5oLbyVpmW6x9cfa7vs2WTfURiCrBoUqgBo_-4WTiULmmHSGZHOjzwa8WtrtOQGsAFjIbno85jp6MnGGGZPYZbDAa_b3y5u-YpW7ypZrvD8BgtKVjgtQgZhLAGezMt0ua3DRrWnKqTZ0BJ_EyxOGuHJrLsn00fnMQ"
	//
	//             }
	//
	//       ]
	//
	// }
	Jwks *string `json:"Jwks,omitempty" xml:"Jwks,omitempty"`
	// The name of the key used for JWT-based identity authentication.
	//
	// example:
	//
	// iss
	KeyName *string `json:"KeyName,omitempty" xml:"KeyName,omitempty"`
	// The value of the key used for JWT-based identity authentication.
	//
	// example:
	//
	// abcd
	KeyValue *string `json:"KeyValue,omitempty" xml:"KeyValue,omitempty"`
	// The name of the consumer.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The creator.
	//
	// example:
	//
	// 123
	PrimaryUser *string `json:"PrimaryUser,omitempty" xml:"PrimaryUser,omitempty"`
	// The resource list.
	ResourceList []*GetGatewayAuthConsumerDetailResponseBodyDataResourceList `json:"ResourceList,omitempty" xml:"ResourceList,omitempty" type:"Repeated"`
	// The names of the parameters that are required to verify each token. By default, each token is prefixed with Bearer and stored in the Authorization header, such as `Authorization: Bearer <Content of a token>`.
	//
	// example:
	//
	// Authorization
	TokenName *string `json:"TokenName,omitempty" xml:"TokenName,omitempty"`
	// Specifies whether to enable pass-through.
	//
	// example:
	//
	// true
	TokenPass *bool `json:"TokenPass,omitempty" xml:"TokenPass,omitempty"`
	// The positions of the parameters that are required to verify each token. By default, each token is prefixed with Bearer and stored in the Authorization header, such as `Authorization: Bearer <Content of a token>`.
	//
	// example:
	//
	// HEADER
	TokenPosition *string `json:"TokenPosition,omitempty" xml:"TokenPosition,omitempty"`
	// The prefixes of the parameters that are required to verify each token. By default, each token is prefixed with Bearer and stored in the Authorization header, such as `Authorization: Bearer <Content of a token>`.
	//
	// example:
	//
	// Bearer
	TokenPrefix *string `json:"TokenPrefix,omitempty" xml:"TokenPrefix,omitempty"`
	// The authentication type. Valid values:
	//
	// 	- JWT
	//
	// example:
	//
	// JWT
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetGatewayAuthConsumerDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayAuthConsumerDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetGatewayAuthConsumerDetailResponseBodyData) GetConsumerStatus() *bool {
	return s.ConsumerStatus
}

func (s *GetGatewayAuthConsumerDetailResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetGatewayAuthConsumerDetailResponseBodyData) GetEncodeType() *string {
	return s.EncodeType
}

func (s *GetGatewayAuthConsumerDetailResponseBodyData) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *GetGatewayAuthConsumerDetailResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetGatewayAuthConsumerDetailResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetGatewayAuthConsumerDetailResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *GetGatewayAuthConsumerDetailResponseBodyData) GetJwks() *string {
	return s.Jwks
}

func (s *GetGatewayAuthConsumerDetailResponseBodyData) GetKeyName() *string {
	return s.KeyName
}

func (s *GetGatewayAuthConsumerDetailResponseBodyData) GetKeyValue() *string {
	return s.KeyValue
}

func (s *GetGatewayAuthConsumerDetailResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetGatewayAuthConsumerDetailResponseBodyData) GetPrimaryUser() *string {
	return s.PrimaryUser
}

func (s *GetGatewayAuthConsumerDetailResponseBodyData) GetResourceList() []*GetGatewayAuthConsumerDetailResponseBodyDataResourceList {
	return s.ResourceList
}

func (s *GetGatewayAuthConsumerDetailResponseBodyData) GetTokenName() *string {
	return s.TokenName
}

func (s *GetGatewayAuthConsumerDetailResponseBodyData) GetTokenPass() *bool {
	return s.TokenPass
}

func (s *GetGatewayAuthConsumerDetailResponseBodyData) GetTokenPosition() *string {
	return s.TokenPosition
}

func (s *GetGatewayAuthConsumerDetailResponseBodyData) GetTokenPrefix() *string {
	return s.TokenPrefix
}

func (s *GetGatewayAuthConsumerDetailResponseBodyData) GetType() *string {
	return s.Type
}

func (s *GetGatewayAuthConsumerDetailResponseBodyData) SetConsumerStatus(v bool) *GetGatewayAuthConsumerDetailResponseBodyData {
	s.ConsumerStatus = &v
	return s
}

func (s *GetGatewayAuthConsumerDetailResponseBodyData) SetDescription(v string) *GetGatewayAuthConsumerDetailResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetGatewayAuthConsumerDetailResponseBodyData) SetEncodeType(v string) *GetGatewayAuthConsumerDetailResponseBodyData {
	s.EncodeType = &v
	return s
}

func (s *GetGatewayAuthConsumerDetailResponseBodyData) SetGatewayUniqueId(v string) *GetGatewayAuthConsumerDetailResponseBodyData {
	s.GatewayUniqueId = &v
	return s
}

func (s *GetGatewayAuthConsumerDetailResponseBodyData) SetGmtCreate(v string) *GetGatewayAuthConsumerDetailResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *GetGatewayAuthConsumerDetailResponseBodyData) SetGmtModified(v string) *GetGatewayAuthConsumerDetailResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *GetGatewayAuthConsumerDetailResponseBodyData) SetId(v int64) *GetGatewayAuthConsumerDetailResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetGatewayAuthConsumerDetailResponseBodyData) SetJwks(v string) *GetGatewayAuthConsumerDetailResponseBodyData {
	s.Jwks = &v
	return s
}

func (s *GetGatewayAuthConsumerDetailResponseBodyData) SetKeyName(v string) *GetGatewayAuthConsumerDetailResponseBodyData {
	s.KeyName = &v
	return s
}

func (s *GetGatewayAuthConsumerDetailResponseBodyData) SetKeyValue(v string) *GetGatewayAuthConsumerDetailResponseBodyData {
	s.KeyValue = &v
	return s
}

func (s *GetGatewayAuthConsumerDetailResponseBodyData) SetName(v string) *GetGatewayAuthConsumerDetailResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetGatewayAuthConsumerDetailResponseBodyData) SetPrimaryUser(v string) *GetGatewayAuthConsumerDetailResponseBodyData {
	s.PrimaryUser = &v
	return s
}

func (s *GetGatewayAuthConsumerDetailResponseBodyData) SetResourceList(v []*GetGatewayAuthConsumerDetailResponseBodyDataResourceList) *GetGatewayAuthConsumerDetailResponseBodyData {
	s.ResourceList = v
	return s
}

func (s *GetGatewayAuthConsumerDetailResponseBodyData) SetTokenName(v string) *GetGatewayAuthConsumerDetailResponseBodyData {
	s.TokenName = &v
	return s
}

func (s *GetGatewayAuthConsumerDetailResponseBodyData) SetTokenPass(v bool) *GetGatewayAuthConsumerDetailResponseBodyData {
	s.TokenPass = &v
	return s
}

func (s *GetGatewayAuthConsumerDetailResponseBodyData) SetTokenPosition(v string) *GetGatewayAuthConsumerDetailResponseBodyData {
	s.TokenPosition = &v
	return s
}

func (s *GetGatewayAuthConsumerDetailResponseBodyData) SetTokenPrefix(v string) *GetGatewayAuthConsumerDetailResponseBodyData {
	s.TokenPrefix = &v
	return s
}

func (s *GetGatewayAuthConsumerDetailResponseBodyData) SetType(v string) *GetGatewayAuthConsumerDetailResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetGatewayAuthConsumerDetailResponseBodyData) Validate() error {
	if s.ResourceList != nil {
		for _, item := range s.ResourceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetGatewayAuthConsumerDetailResponseBodyDataResourceList struct {
	// The consumer ID.
	//
	// example:
	//
	// 123
	ConsumerId *int64 `json:"ConsumerId,omitempty" xml:"ConsumerId,omitempty"`
	// The unique ID of the gateway.
	//
	// example:
	//
	// gw-4822033a8513496fa10f05c934f*****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The time when the resource associated with the consumer authentication record was created.
	//
	// example:
	//
	// 2022-12-06 01:38:03
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the resource associated with the consumer authentication record was modified.
	//
	// example:
	//
	// 2022-12-06 01:38:03
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the authorized resource for the consumer.
	//
	// example:
	//
	// 16
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The resource authorization state. Valid values:
	//
	// 	- true: Resource authorization is enabled.
	//
	// 	- false: Resource authorization is disabled.
	//
	// example:
	//
	// true
	ResourceStatus *bool `json:"ResourceStatus,omitempty" xml:"ResourceStatus,omitempty"`
	// The ID of the route.
	//
	// example:
	//
	// 3458
	RouteId *int64 `json:"RouteId,omitempty" xml:"RouteId,omitempty"`
	// The name of the route.
	//
	// example:
	//
	// test
	RouteName *string `json:"RouteName,omitempty" xml:"RouteName,omitempty"`
}

func (s GetGatewayAuthConsumerDetailResponseBodyDataResourceList) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayAuthConsumerDetailResponseBodyDataResourceList) GoString() string {
	return s.String()
}

func (s *GetGatewayAuthConsumerDetailResponseBodyDataResourceList) GetConsumerId() *int64 {
	return s.ConsumerId
}

func (s *GetGatewayAuthConsumerDetailResponseBodyDataResourceList) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *GetGatewayAuthConsumerDetailResponseBodyDataResourceList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetGatewayAuthConsumerDetailResponseBodyDataResourceList) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetGatewayAuthConsumerDetailResponseBodyDataResourceList) GetId() *int64 {
	return s.Id
}

func (s *GetGatewayAuthConsumerDetailResponseBodyDataResourceList) GetResourceStatus() *bool {
	return s.ResourceStatus
}

func (s *GetGatewayAuthConsumerDetailResponseBodyDataResourceList) GetRouteId() *int64 {
	return s.RouteId
}

func (s *GetGatewayAuthConsumerDetailResponseBodyDataResourceList) GetRouteName() *string {
	return s.RouteName
}

func (s *GetGatewayAuthConsumerDetailResponseBodyDataResourceList) SetConsumerId(v int64) *GetGatewayAuthConsumerDetailResponseBodyDataResourceList {
	s.ConsumerId = &v
	return s
}

func (s *GetGatewayAuthConsumerDetailResponseBodyDataResourceList) SetGatewayUniqueId(v string) *GetGatewayAuthConsumerDetailResponseBodyDataResourceList {
	s.GatewayUniqueId = &v
	return s
}

func (s *GetGatewayAuthConsumerDetailResponseBodyDataResourceList) SetGmtCreate(v string) *GetGatewayAuthConsumerDetailResponseBodyDataResourceList {
	s.GmtCreate = &v
	return s
}

func (s *GetGatewayAuthConsumerDetailResponseBodyDataResourceList) SetGmtModified(v string) *GetGatewayAuthConsumerDetailResponseBodyDataResourceList {
	s.GmtModified = &v
	return s
}

func (s *GetGatewayAuthConsumerDetailResponseBodyDataResourceList) SetId(v int64) *GetGatewayAuthConsumerDetailResponseBodyDataResourceList {
	s.Id = &v
	return s
}

func (s *GetGatewayAuthConsumerDetailResponseBodyDataResourceList) SetResourceStatus(v bool) *GetGatewayAuthConsumerDetailResponseBodyDataResourceList {
	s.ResourceStatus = &v
	return s
}

func (s *GetGatewayAuthConsumerDetailResponseBodyDataResourceList) SetRouteId(v int64) *GetGatewayAuthConsumerDetailResponseBodyDataResourceList {
	s.RouteId = &v
	return s
}

func (s *GetGatewayAuthConsumerDetailResponseBodyDataResourceList) SetRouteName(v string) *GetGatewayAuthConsumerDetailResponseBodyDataResourceList {
	s.RouteName = &v
	return s
}

func (s *GetGatewayAuthConsumerDetailResponseBodyDataResourceList) Validate() error {
	return dara.Validate(s)
}
