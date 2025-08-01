// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListServiceSourceResponseBody
	GetCode() *int32
	SetData(v []*ListServiceSourceResponseBodyData) *ListServiceSourceResponseBody
	GetData() []*ListServiceSourceResponseBodyData
	SetHttpStatusCode(v int32) *ListServiceSourceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListServiceSourceResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListServiceSourceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListServiceSourceResponseBody
	GetSuccess() *bool
}

type ListServiceSourceResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data []*ListServiceSourceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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
	// You are not authorized to perform this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// B8C4B0D8-EBB9-5F20-8295-04224FBE5529
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

func (s ListServiceSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListServiceSourceResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceSourceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListServiceSourceResponseBody) GetData() []*ListServiceSourceResponseBodyData {
	return s.Data
}

func (s *ListServiceSourceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListServiceSourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListServiceSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListServiceSourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListServiceSourceResponseBody) SetCode(v int32) *ListServiceSourceResponseBody {
	s.Code = &v
	return s
}

func (s *ListServiceSourceResponseBody) SetData(v []*ListServiceSourceResponseBodyData) *ListServiceSourceResponseBody {
	s.Data = v
	return s
}

func (s *ListServiceSourceResponseBody) SetHttpStatusCode(v int32) *ListServiceSourceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListServiceSourceResponseBody) SetMessage(v string) *ListServiceSourceResponseBody {
	s.Message = &v
	return s
}

func (s *ListServiceSourceResponseBody) SetRequestId(v string) *ListServiceSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServiceSourceResponseBody) SetSuccess(v bool) *ListServiceSourceResponseBody {
	s.Success = &v
	return s
}

func (s *ListServiceSourceResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListServiceSourceResponseBodyData struct {
	// The ID of the Container Service for Kubernetes (ACK) cluster or the endpoint of the Microservices Engine (MSE) instance.
	//
	// example:
	//
	// ***
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// Indicates whether the service source is associated with the gateway. The value 1 indicates that the service source is associated with the gateway.
	//
	// example:
	//
	// 1
	BindingWithGateway *int32 `json:"BindingWithGateway,omitempty" xml:"BindingWithGateway,omitempty"`
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
	// gw-2u9uhd9283hd92hgd39g239dg2*****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2022-01-07 18:07:57
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2022-01-07 18:07:57
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The array of service groups.
	GroupList []*string `json:"GroupList,omitempty" xml:"GroupList,omitempty" type:"Repeated"`
	// The ID.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The information about the support for Ingresses by applications.
	IngressOptions *ListServiceSourceResponseBodyDataIngressOptions `json:"IngressOptions,omitempty" xml:"IngressOptions,omitempty" type:"Struct"`
	Invalid        *bool                                            `json:"Invalid,omitempty" xml:"Invalid,omitempty"`
	// The name.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The array of root paths of service lists.
	PathList []*string `json:"PathList,omitempty" xml:"PathList,omitempty" type:"Repeated"`
	// The type of the service source.
	//
	// example:
	//
	// MSE
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The unique ID of the service source.
	//
	// example:
	//
	// mse-cn-***
	SourceUniqueId *string `json:"SourceUniqueId,omitempty" xml:"SourceUniqueId,omitempty"`
	// The type.
	//
	// example:
	//
	// NACOS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListServiceSourceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListServiceSourceResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListServiceSourceResponseBodyData) GetAddress() *string {
	return s.Address
}

func (s *ListServiceSourceResponseBodyData) GetBindingWithGateway() *int32 {
	return s.BindingWithGateway
}

func (s *ListServiceSourceResponseBodyData) GetGatewayId() *int64 {
	return s.GatewayId
}

func (s *ListServiceSourceResponseBodyData) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *ListServiceSourceResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListServiceSourceResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListServiceSourceResponseBodyData) GetGroupList() []*string {
	return s.GroupList
}

func (s *ListServiceSourceResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *ListServiceSourceResponseBodyData) GetIngressOptions() *ListServiceSourceResponseBodyDataIngressOptions {
	return s.IngressOptions
}

func (s *ListServiceSourceResponseBodyData) GetInvalid() *bool {
	return s.Invalid
}

func (s *ListServiceSourceResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListServiceSourceResponseBodyData) GetPathList() []*string {
	return s.PathList
}

func (s *ListServiceSourceResponseBodyData) GetSource() *string {
	return s.Source
}

func (s *ListServiceSourceResponseBodyData) GetSourceUniqueId() *string {
	return s.SourceUniqueId
}

func (s *ListServiceSourceResponseBodyData) GetType() *string {
	return s.Type
}

func (s *ListServiceSourceResponseBodyData) SetAddress(v string) *ListServiceSourceResponseBodyData {
	s.Address = &v
	return s
}

func (s *ListServiceSourceResponseBodyData) SetBindingWithGateway(v int32) *ListServiceSourceResponseBodyData {
	s.BindingWithGateway = &v
	return s
}

func (s *ListServiceSourceResponseBodyData) SetGatewayId(v int64) *ListServiceSourceResponseBodyData {
	s.GatewayId = &v
	return s
}

func (s *ListServiceSourceResponseBodyData) SetGatewayUniqueId(v string) *ListServiceSourceResponseBodyData {
	s.GatewayUniqueId = &v
	return s
}

func (s *ListServiceSourceResponseBodyData) SetGmtCreate(v string) *ListServiceSourceResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *ListServiceSourceResponseBodyData) SetGmtModified(v string) *ListServiceSourceResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *ListServiceSourceResponseBodyData) SetGroupList(v []*string) *ListServiceSourceResponseBodyData {
	s.GroupList = v
	return s
}

func (s *ListServiceSourceResponseBodyData) SetId(v int64) *ListServiceSourceResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListServiceSourceResponseBodyData) SetIngressOptions(v *ListServiceSourceResponseBodyDataIngressOptions) *ListServiceSourceResponseBodyData {
	s.IngressOptions = v
	return s
}

func (s *ListServiceSourceResponseBodyData) SetInvalid(v bool) *ListServiceSourceResponseBodyData {
	s.Invalid = &v
	return s
}

func (s *ListServiceSourceResponseBodyData) SetName(v string) *ListServiceSourceResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListServiceSourceResponseBodyData) SetPathList(v []*string) *ListServiceSourceResponseBodyData {
	s.PathList = v
	return s
}

func (s *ListServiceSourceResponseBodyData) SetSource(v string) *ListServiceSourceResponseBodyData {
	s.Source = &v
	return s
}

func (s *ListServiceSourceResponseBodyData) SetSourceUniqueId(v string) *ListServiceSourceResponseBodyData {
	s.SourceUniqueId = &v
	return s
}

func (s *ListServiceSourceResponseBodyData) SetType(v string) *ListServiceSourceResponseBodyData {
	s.Type = &v
	return s
}

func (s *ListServiceSourceResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListServiceSourceResponseBodyDataIngressOptions struct {
	// Indicates whether Ingresses are enabled.
	//
	// example:
	//
	// true
	EnableIngress *bool `json:"EnableIngress,omitempty" xml:"EnableIngress,omitempty"`
	// Indicates whether the Ingress status is updated.
	//
	// example:
	//
	// true
	EnableStatus *bool `json:"EnableStatus,omitempty" xml:"EnableStatus,omitempty"`
	// The Ingress class.
	//
	// example:
	//
	// com.test.xxx
	IngressClass *string `json:"IngressClass,omitempty" xml:"IngressClass,omitempty"`
	// The namespace that you want to monitor.
	//
	// example:
	//
	// default
	WatchNamespace *string `json:"WatchNamespace,omitempty" xml:"WatchNamespace,omitempty"`
}

func (s ListServiceSourceResponseBodyDataIngressOptions) String() string {
	return dara.Prettify(s)
}

func (s ListServiceSourceResponseBodyDataIngressOptions) GoString() string {
	return s.String()
}

func (s *ListServiceSourceResponseBodyDataIngressOptions) GetEnableIngress() *bool {
	return s.EnableIngress
}

func (s *ListServiceSourceResponseBodyDataIngressOptions) GetEnableStatus() *bool {
	return s.EnableStatus
}

func (s *ListServiceSourceResponseBodyDataIngressOptions) GetIngressClass() *string {
	return s.IngressClass
}

func (s *ListServiceSourceResponseBodyDataIngressOptions) GetWatchNamespace() *string {
	return s.WatchNamespace
}

func (s *ListServiceSourceResponseBodyDataIngressOptions) SetEnableIngress(v bool) *ListServiceSourceResponseBodyDataIngressOptions {
	s.EnableIngress = &v
	return s
}

func (s *ListServiceSourceResponseBodyDataIngressOptions) SetEnableStatus(v bool) *ListServiceSourceResponseBodyDataIngressOptions {
	s.EnableStatus = &v
	return s
}

func (s *ListServiceSourceResponseBodyDataIngressOptions) SetIngressClass(v string) *ListServiceSourceResponseBodyDataIngressOptions {
	s.IngressClass = &v
	return s
}

func (s *ListServiceSourceResponseBodyDataIngressOptions) SetWatchNamespace(v string) *ListServiceSourceResponseBodyDataIngressOptions {
	s.WatchNamespace = &v
	return s
}

func (s *ListServiceSourceResponseBodyDataIngressOptions) Validate() error {
	return dara.Validate(s)
}
