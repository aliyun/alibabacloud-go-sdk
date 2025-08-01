// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGatewayServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteGatewayServiceResponseBody
	GetCode() *int32
	SetData(v *DeleteGatewayServiceResponseBodyData) *DeleteGatewayServiceResponseBody
	GetData() *DeleteGatewayServiceResponseBodyData
	SetHttpStatusCode(v int32) *DeleteGatewayServiceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteGatewayServiceResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteGatewayServiceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteGatewayServiceResponseBody
	GetSuccess() *bool
}

type DeleteGatewayServiceResponseBody struct {
	// The response code returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *DeleteGatewayServiceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// The unique ID of the request.
	//
	// example:
	//
	// B3545F76-6ED1-586F-8DB9-ECE07985F381
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

func (s DeleteGatewayServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteGatewayServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGatewayServiceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteGatewayServiceResponseBody) GetData() *DeleteGatewayServiceResponseBodyData {
	return s.Data
}

func (s *DeleteGatewayServiceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteGatewayServiceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteGatewayServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteGatewayServiceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteGatewayServiceResponseBody) SetCode(v int32) *DeleteGatewayServiceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteGatewayServiceResponseBody) SetData(v *DeleteGatewayServiceResponseBodyData) *DeleteGatewayServiceResponseBody {
	s.Data = v
	return s
}

func (s *DeleteGatewayServiceResponseBody) SetHttpStatusCode(v int32) *DeleteGatewayServiceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteGatewayServiceResponseBody) SetMessage(v string) *DeleteGatewayServiceResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteGatewayServiceResponseBody) SetRequestId(v string) *DeleteGatewayServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteGatewayServiceResponseBody) SetSuccess(v bool) *DeleteGatewayServiceResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteGatewayServiceResponseBody) Validate() error {
	return dara.Validate(s)
}

type DeleteGatewayServiceResponseBodyData struct {
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
	// The time when the service was created.
	//
	// example:
	//
	// 2022-01-14 14:39:16
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The last modification time.
	//
	// example:
	//
	// 2022-01-07T10:07:57.000+0000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The name of the group.
	//
	// example:
	//
	// DEFAULT_GROUP
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The ID of the service.
	//
	// example:
	//
	// 190
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// IP
	Ips []*string `json:"Ips,omitempty" xml:"Ips,omitempty" type:"Repeated"`
	// The basic information about the service.
	//
	// example:
	//
	// {}
	MetaInfo *string `json:"MetaInfo,omitempty" xml:"MetaInfo,omitempty"`
	// The name of the service.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The namespace.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The name of the service registered with the service registry.
	//
	// example:
	//
	// test
	ServiceNameInRegistry *string `json:"ServiceNameInRegistry,omitempty" xml:"ServiceNameInRegistry,omitempty"`
	// The ID of the service source.
	//
	// example:
	//
	// 1
	SourceId *int64 `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// The source type of the service.
	//
	// example:
	//
	// MSE
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s DeleteGatewayServiceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteGatewayServiceResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteGatewayServiceResponseBodyData) GetGatewayId() *int64 {
	return s.GatewayId
}

func (s *DeleteGatewayServiceResponseBodyData) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *DeleteGatewayServiceResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *DeleteGatewayServiceResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DeleteGatewayServiceResponseBodyData) GetGroupName() *string {
	return s.GroupName
}

func (s *DeleteGatewayServiceResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *DeleteGatewayServiceResponseBodyData) GetIps() []*string {
	return s.Ips
}

func (s *DeleteGatewayServiceResponseBodyData) GetMetaInfo() *string {
	return s.MetaInfo
}

func (s *DeleteGatewayServiceResponseBodyData) GetName() *string {
	return s.Name
}

func (s *DeleteGatewayServiceResponseBodyData) GetNamespace() *string {
	return s.Namespace
}

func (s *DeleteGatewayServiceResponseBodyData) GetServiceNameInRegistry() *string {
	return s.ServiceNameInRegistry
}

func (s *DeleteGatewayServiceResponseBodyData) GetSourceId() *int64 {
	return s.SourceId
}

func (s *DeleteGatewayServiceResponseBodyData) GetSourceType() *string {
	return s.SourceType
}

func (s *DeleteGatewayServiceResponseBodyData) SetGatewayId(v int64) *DeleteGatewayServiceResponseBodyData {
	s.GatewayId = &v
	return s
}

func (s *DeleteGatewayServiceResponseBodyData) SetGatewayUniqueId(v string) *DeleteGatewayServiceResponseBodyData {
	s.GatewayUniqueId = &v
	return s
}

func (s *DeleteGatewayServiceResponseBodyData) SetGmtCreate(v string) *DeleteGatewayServiceResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *DeleteGatewayServiceResponseBodyData) SetGmtModified(v string) *DeleteGatewayServiceResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *DeleteGatewayServiceResponseBodyData) SetGroupName(v string) *DeleteGatewayServiceResponseBodyData {
	s.GroupName = &v
	return s
}

func (s *DeleteGatewayServiceResponseBodyData) SetId(v int64) *DeleteGatewayServiceResponseBodyData {
	s.Id = &v
	return s
}

func (s *DeleteGatewayServiceResponseBodyData) SetIps(v []*string) *DeleteGatewayServiceResponseBodyData {
	s.Ips = v
	return s
}

func (s *DeleteGatewayServiceResponseBodyData) SetMetaInfo(v string) *DeleteGatewayServiceResponseBodyData {
	s.MetaInfo = &v
	return s
}

func (s *DeleteGatewayServiceResponseBodyData) SetName(v string) *DeleteGatewayServiceResponseBodyData {
	s.Name = &v
	return s
}

func (s *DeleteGatewayServiceResponseBodyData) SetNamespace(v string) *DeleteGatewayServiceResponseBodyData {
	s.Namespace = &v
	return s
}

func (s *DeleteGatewayServiceResponseBodyData) SetServiceNameInRegistry(v string) *DeleteGatewayServiceResponseBodyData {
	s.ServiceNameInRegistry = &v
	return s
}

func (s *DeleteGatewayServiceResponseBodyData) SetSourceId(v int64) *DeleteGatewayServiceResponseBodyData {
	s.SourceId = &v
	return s
}

func (s *DeleteGatewayServiceResponseBodyData) SetSourceType(v string) *DeleteGatewayServiceResponseBodyData {
	s.SourceType = &v
	return s
}

func (s *DeleteGatewayServiceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
