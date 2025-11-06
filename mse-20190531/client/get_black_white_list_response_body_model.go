// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBlackWhiteListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetBlackWhiteListResponseBody
	GetCode() *int32
	SetData(v *GetBlackWhiteListResponseBodyData) *GetBlackWhiteListResponseBody
	GetData() *GetBlackWhiteListResponseBodyData
	SetHttpStatusCode(v int32) *GetBlackWhiteListResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetBlackWhiteListResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetBlackWhiteListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetBlackWhiteListResponseBody
	GetSuccess() *bool
}

type GetBlackWhiteListResponseBody struct {
	// The status code returned. A value of 200 indicates that the request is successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data of the node.
	Data *GetBlackWhiteListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message returned.
	//
	// example:
	//
	// You are not authorized to perform this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 89CBC928-4F57-51FA-A413-EE0F4CD87200
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

func (s GetBlackWhiteListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBlackWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *GetBlackWhiteListResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetBlackWhiteListResponseBody) GetData() *GetBlackWhiteListResponseBodyData {
	return s.Data
}

func (s *GetBlackWhiteListResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetBlackWhiteListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetBlackWhiteListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetBlackWhiteListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetBlackWhiteListResponseBody) SetCode(v int32) *GetBlackWhiteListResponseBody {
	s.Code = &v
	return s
}

func (s *GetBlackWhiteListResponseBody) SetData(v *GetBlackWhiteListResponseBodyData) *GetBlackWhiteListResponseBody {
	s.Data = v
	return s
}

func (s *GetBlackWhiteListResponseBody) SetHttpStatusCode(v int32) *GetBlackWhiteListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetBlackWhiteListResponseBody) SetMessage(v string) *GetBlackWhiteListResponseBody {
	s.Message = &v
	return s
}

func (s *GetBlackWhiteListResponseBody) SetRequestId(v string) *GetBlackWhiteListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBlackWhiteListResponseBody) SetSuccess(v bool) *GetBlackWhiteListResponseBody {
	s.Success = &v
	return s
}

func (s *GetBlackWhiteListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetBlackWhiteListResponseBodyData struct {
	// The content of the blacklist.
	//
	// example:
	//
	// text
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The ID of the gateway.
	//
	// example:
	//
	// 1
	GatewayId *int64 `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// The unique ID of the gateway.
	//
	// example:
	//
	// gw-7ea3da97b96543e19f6c597c****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2022-01-07T10:07:57.000+0000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 2022-01-07T10:07:57.000+0000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID.
	//
	// example:
	//
	// 275
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Indicates whether the whitelist is enabled.
	//
	// example:
	//
	// true
	IsWhite *bool `json:"IsWhite,omitempty" xml:"IsWhite,omitempty"`
	// The ID of the resource.
	//
	// example:
	//
	// 1
	ResourceId *int64 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of a resource.
	//
	// example:
	//
	// GATEWAY
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The status of the blacklist or whitelist.
	//
	// example:
	//
	// on
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type.
	//
	// example:
	//
	// IP
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetBlackWhiteListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetBlackWhiteListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetBlackWhiteListResponseBodyData) GetContent() *string {
	return s.Content
}

func (s *GetBlackWhiteListResponseBodyData) GetGatewayId() *int64 {
	return s.GatewayId
}

func (s *GetBlackWhiteListResponseBodyData) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *GetBlackWhiteListResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetBlackWhiteListResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetBlackWhiteListResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *GetBlackWhiteListResponseBodyData) GetIsWhite() *bool {
	return s.IsWhite
}

func (s *GetBlackWhiteListResponseBodyData) GetResourceId() *int64 {
	return s.ResourceId
}

func (s *GetBlackWhiteListResponseBodyData) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetBlackWhiteListResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetBlackWhiteListResponseBodyData) GetType() *string {
	return s.Type
}

func (s *GetBlackWhiteListResponseBodyData) SetContent(v string) *GetBlackWhiteListResponseBodyData {
	s.Content = &v
	return s
}

func (s *GetBlackWhiteListResponseBodyData) SetGatewayId(v int64) *GetBlackWhiteListResponseBodyData {
	s.GatewayId = &v
	return s
}

func (s *GetBlackWhiteListResponseBodyData) SetGatewayUniqueId(v string) *GetBlackWhiteListResponseBodyData {
	s.GatewayUniqueId = &v
	return s
}

func (s *GetBlackWhiteListResponseBodyData) SetGmtCreate(v string) *GetBlackWhiteListResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *GetBlackWhiteListResponseBodyData) SetGmtModified(v string) *GetBlackWhiteListResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *GetBlackWhiteListResponseBodyData) SetId(v int64) *GetBlackWhiteListResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetBlackWhiteListResponseBodyData) SetIsWhite(v bool) *GetBlackWhiteListResponseBodyData {
	s.IsWhite = &v
	return s
}

func (s *GetBlackWhiteListResponseBodyData) SetResourceId(v int64) *GetBlackWhiteListResponseBodyData {
	s.ResourceId = &v
	return s
}

func (s *GetBlackWhiteListResponseBodyData) SetResourceType(v string) *GetBlackWhiteListResponseBodyData {
	s.ResourceType = &v
	return s
}

func (s *GetBlackWhiteListResponseBodyData) SetStatus(v string) *GetBlackWhiteListResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetBlackWhiteListResponseBodyData) SetType(v string) *GetBlackWhiteListResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetBlackWhiteListResponseBodyData) Validate() error {
	return dara.Validate(s)
}
