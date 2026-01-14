// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBlackWhiteListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteBlackWhiteListResponseBody
	GetCode() *int32
	SetData(v *DeleteBlackWhiteListResponseBodyData) *DeleteBlackWhiteListResponseBody
	GetData() *DeleteBlackWhiteListResponseBodyData
	SetHttpStatusCode(v int32) *DeleteBlackWhiteListResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteBlackWhiteListResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteBlackWhiteListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteBlackWhiteListResponseBody
	GetSuccess() *bool
}

type DeleteBlackWhiteListResponseBody struct {
	// example:
	//
	// 200
	Code *int32                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *DeleteBlackWhiteListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// DC34E4A3-5F1C-4E40-86EA-02EDF967****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteBlackWhiteListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteBlackWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBlackWhiteListResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteBlackWhiteListResponseBody) GetData() *DeleteBlackWhiteListResponseBodyData {
	return s.Data
}

func (s *DeleteBlackWhiteListResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteBlackWhiteListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteBlackWhiteListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteBlackWhiteListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteBlackWhiteListResponseBody) SetCode(v int32) *DeleteBlackWhiteListResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteBlackWhiteListResponseBody) SetData(v *DeleteBlackWhiteListResponseBodyData) *DeleteBlackWhiteListResponseBody {
	s.Data = v
	return s
}

func (s *DeleteBlackWhiteListResponseBody) SetHttpStatusCode(v int32) *DeleteBlackWhiteListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteBlackWhiteListResponseBody) SetMessage(v string) *DeleteBlackWhiteListResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteBlackWhiteListResponseBody) SetRequestId(v string) *DeleteBlackWhiteListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteBlackWhiteListResponseBody) SetSuccess(v bool) *DeleteBlackWhiteListResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteBlackWhiteListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteBlackWhiteListResponseBodyData struct {
	// example:
	//
	// 1.1.1.1
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 430
	GatewayId *int64 `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// example:
	//
	// gw-9cdcf8e4f58144059e73ff4c5ef9****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// example:
	//
	// 2022-08-10 20:22:34
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2022-08-10 20:22:34
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 120
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// true
	IsWhite *bool `json:"IsWhite,omitempty" xml:"IsWhite,omitempty"`
	// example:
	//
	// 1
	ResourceId *int64 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// GATEWAY
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// 1
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// IP
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DeleteBlackWhiteListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteBlackWhiteListResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteBlackWhiteListResponseBodyData) GetContent() *string {
	return s.Content
}

func (s *DeleteBlackWhiteListResponseBodyData) GetGatewayId() *int64 {
	return s.GatewayId
}

func (s *DeleteBlackWhiteListResponseBodyData) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *DeleteBlackWhiteListResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *DeleteBlackWhiteListResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DeleteBlackWhiteListResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *DeleteBlackWhiteListResponseBodyData) GetIsWhite() *bool {
	return s.IsWhite
}

func (s *DeleteBlackWhiteListResponseBodyData) GetResourceId() *int64 {
	return s.ResourceId
}

func (s *DeleteBlackWhiteListResponseBodyData) GetResourceType() *string {
	return s.ResourceType
}

func (s *DeleteBlackWhiteListResponseBodyData) GetStatus() *bool {
	return s.Status
}

func (s *DeleteBlackWhiteListResponseBodyData) GetType() *string {
	return s.Type
}

func (s *DeleteBlackWhiteListResponseBodyData) SetContent(v string) *DeleteBlackWhiteListResponseBodyData {
	s.Content = &v
	return s
}

func (s *DeleteBlackWhiteListResponseBodyData) SetGatewayId(v int64) *DeleteBlackWhiteListResponseBodyData {
	s.GatewayId = &v
	return s
}

func (s *DeleteBlackWhiteListResponseBodyData) SetGatewayUniqueId(v string) *DeleteBlackWhiteListResponseBodyData {
	s.GatewayUniqueId = &v
	return s
}

func (s *DeleteBlackWhiteListResponseBodyData) SetGmtCreate(v string) *DeleteBlackWhiteListResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *DeleteBlackWhiteListResponseBodyData) SetGmtModified(v string) *DeleteBlackWhiteListResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *DeleteBlackWhiteListResponseBodyData) SetId(v int64) *DeleteBlackWhiteListResponseBodyData {
	s.Id = &v
	return s
}

func (s *DeleteBlackWhiteListResponseBodyData) SetIsWhite(v bool) *DeleteBlackWhiteListResponseBodyData {
	s.IsWhite = &v
	return s
}

func (s *DeleteBlackWhiteListResponseBodyData) SetResourceId(v int64) *DeleteBlackWhiteListResponseBodyData {
	s.ResourceId = &v
	return s
}

func (s *DeleteBlackWhiteListResponseBodyData) SetResourceType(v string) *DeleteBlackWhiteListResponseBodyData {
	s.ResourceType = &v
	return s
}

func (s *DeleteBlackWhiteListResponseBodyData) SetStatus(v bool) *DeleteBlackWhiteListResponseBodyData {
	s.Status = &v
	return s
}

func (s *DeleteBlackWhiteListResponseBodyData) SetType(v string) *DeleteBlackWhiteListResponseBodyData {
	s.Type = &v
	return s
}

func (s *DeleteBlackWhiteListResponseBodyData) Validate() error {
	return dara.Validate(s)
}
