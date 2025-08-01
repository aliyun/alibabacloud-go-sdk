// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAuthResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteAuthResourceResponseBody
	GetCode() *int32
	SetData(v *DeleteAuthResourceResponseBodyData) *DeleteAuthResourceResponseBody
	GetData() *DeleteAuthResourceResponseBodyData
	SetHttpStatusCode(v int32) *DeleteAuthResourceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteAuthResourceResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteAuthResourceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteAuthResourceResponseBody
	GetSuccess() *bool
}

type DeleteAuthResourceResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// 1
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *DeleteAuthResourceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// The request is successfully processed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 559412D1-BFCE-53CC-B88E-0192C331EF44
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

func (s DeleteAuthResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAuthResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAuthResourceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteAuthResourceResponseBody) GetData() *DeleteAuthResourceResponseBodyData {
	return s.Data
}

func (s *DeleteAuthResourceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteAuthResourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteAuthResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAuthResourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteAuthResourceResponseBody) SetCode(v int32) *DeleteAuthResourceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAuthResourceResponseBody) SetData(v *DeleteAuthResourceResponseBodyData) *DeleteAuthResourceResponseBody {
	s.Data = v
	return s
}

func (s *DeleteAuthResourceResponseBody) SetHttpStatusCode(v int32) *DeleteAuthResourceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteAuthResourceResponseBody) SetMessage(v string) *DeleteAuthResourceResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAuthResourceResponseBody) SetRequestId(v string) *DeleteAuthResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAuthResourceResponseBody) SetSuccess(v bool) *DeleteAuthResourceResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteAuthResourceResponseBody) Validate() error {
	return dara.Validate(s)
}

type DeleteAuthResourceResponseBodyData struct {
	// The authentication ID.
	//
	// example:
	//
	// 253
	AuthId *int64 `json:"AuthId,omitempty" xml:"AuthId,omitempty"`
	// The ID of the domain name.
	//
	// example:
	//
	// 235
	DomainId *int64 `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	// The domain name.
	//
	// example:
	//
	// name
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The ID of the gateway.
	//
	// example:
	//
	// gw-7ea3da97b96543e19f6c597cd4a9****
	GatewayId *int64 `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// The unique ID of the gateway.
	//
	// example:
	//
	// gw-7ea3da97b96543e19f6c597cd4a9****
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
	// The ID of the primary key.
	//
	// example:
	//
	// 12
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Indicates whether the whitelist mode is enabled.
	//
	// example:
	//
	// true
	IsWhite *bool `json:"IsWhite,omitempty" xml:"IsWhite,omitempty"`
	// The path.
	//
	// example:
	//
	// /zookeeper
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s DeleteAuthResourceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteAuthResourceResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteAuthResourceResponseBodyData) GetAuthId() *int64 {
	return s.AuthId
}

func (s *DeleteAuthResourceResponseBodyData) GetDomainId() *int64 {
	return s.DomainId
}

func (s *DeleteAuthResourceResponseBodyData) GetDomainName() *string {
	return s.DomainName
}

func (s *DeleteAuthResourceResponseBodyData) GetGatewayId() *int64 {
	return s.GatewayId
}

func (s *DeleteAuthResourceResponseBodyData) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *DeleteAuthResourceResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *DeleteAuthResourceResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DeleteAuthResourceResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *DeleteAuthResourceResponseBodyData) GetIsWhite() *bool {
	return s.IsWhite
}

func (s *DeleteAuthResourceResponseBodyData) GetPath() *string {
	return s.Path
}

func (s *DeleteAuthResourceResponseBodyData) SetAuthId(v int64) *DeleteAuthResourceResponseBodyData {
	s.AuthId = &v
	return s
}

func (s *DeleteAuthResourceResponseBodyData) SetDomainId(v int64) *DeleteAuthResourceResponseBodyData {
	s.DomainId = &v
	return s
}

func (s *DeleteAuthResourceResponseBodyData) SetDomainName(v string) *DeleteAuthResourceResponseBodyData {
	s.DomainName = &v
	return s
}

func (s *DeleteAuthResourceResponseBodyData) SetGatewayId(v int64) *DeleteAuthResourceResponseBodyData {
	s.GatewayId = &v
	return s
}

func (s *DeleteAuthResourceResponseBodyData) SetGatewayUniqueId(v string) *DeleteAuthResourceResponseBodyData {
	s.GatewayUniqueId = &v
	return s
}

func (s *DeleteAuthResourceResponseBodyData) SetGmtCreate(v string) *DeleteAuthResourceResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *DeleteAuthResourceResponseBodyData) SetGmtModified(v string) *DeleteAuthResourceResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *DeleteAuthResourceResponseBodyData) SetId(v int64) *DeleteAuthResourceResponseBodyData {
	s.Id = &v
	return s
}

func (s *DeleteAuthResourceResponseBodyData) SetIsWhite(v bool) *DeleteAuthResourceResponseBodyData {
	s.IsWhite = &v
	return s
}

func (s *DeleteAuthResourceResponseBodyData) SetPath(v string) *DeleteAuthResourceResponseBodyData {
	s.Path = &v
	return s
}

func (s *DeleteAuthResourceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
