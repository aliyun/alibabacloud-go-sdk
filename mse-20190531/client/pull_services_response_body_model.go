// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPullServicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *PullServicesResponseBody
	GetCode() *int32
	SetData(v []*PullServicesResponseBodyData) *PullServicesResponseBody
	GetData() []*PullServicesResponseBodyData
	SetHttpStatusCode(v int32) *PullServicesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *PullServicesResponseBody
	GetMessage() *string
	SetRequestId(v string) *PullServicesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PullServicesResponseBody
	GetSuccess() *bool
}

type PullServicesResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data []*PullServicesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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
	// You are not authorized to perform this operation. Action: mse:PullServices, Resource: acs:mse:cn-shenzhen:1228932054837788:*
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// EC0A99B9-8BA3-5FE3-8FE7-D7C719CF7BD2
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

func (s PullServicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PullServicesResponseBody) GoString() string {
	return s.String()
}

func (s *PullServicesResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *PullServicesResponseBody) GetData() []*PullServicesResponseBodyData {
	return s.Data
}

func (s *PullServicesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *PullServicesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PullServicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PullServicesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PullServicesResponseBody) SetCode(v int32) *PullServicesResponseBody {
	s.Code = &v
	return s
}

func (s *PullServicesResponseBody) SetData(v []*PullServicesResponseBodyData) *PullServicesResponseBody {
	s.Data = v
	return s
}

func (s *PullServicesResponseBody) SetHttpStatusCode(v int32) *PullServicesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *PullServicesResponseBody) SetMessage(v string) *PullServicesResponseBody {
	s.Message = &v
	return s
}

func (s *PullServicesResponseBody) SetRequestId(v string) *PullServicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *PullServicesResponseBody) SetSuccess(v bool) *PullServicesResponseBody {
	s.Success = &v
	return s
}

func (s *PullServicesResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type PullServicesResponseBodyData struct {
	// The name of the group.
	//
	// example:
	//
	// test
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The namespace.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The alias of the namespace.
	//
	// example:
	//
	// public
	NamespaceShowName *string `json:"NamespaceShowName,omitempty" xml:"NamespaceShowName,omitempty"`
	// The information about services.
	Services []*PullServicesResponseBodyDataServices `json:"Services,omitempty" xml:"Services,omitempty" type:"Repeated"`
}

func (s PullServicesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s PullServicesResponseBodyData) GoString() string {
	return s.String()
}

func (s *PullServicesResponseBodyData) GetGroupName() *string {
	return s.GroupName
}

func (s *PullServicesResponseBodyData) GetNamespace() *string {
	return s.Namespace
}

func (s *PullServicesResponseBodyData) GetNamespaceShowName() *string {
	return s.NamespaceShowName
}

func (s *PullServicesResponseBodyData) GetServices() []*PullServicesResponseBodyDataServices {
	return s.Services
}

func (s *PullServicesResponseBodyData) SetGroupName(v string) *PullServicesResponseBodyData {
	s.GroupName = &v
	return s
}

func (s *PullServicesResponseBodyData) SetNamespace(v string) *PullServicesResponseBodyData {
	s.Namespace = &v
	return s
}

func (s *PullServicesResponseBodyData) SetNamespaceShowName(v string) *PullServicesResponseBodyData {
	s.NamespaceShowName = &v
	return s
}

func (s *PullServicesResponseBodyData) SetServices(v []*PullServicesResponseBodyDataServices) *PullServicesResponseBodyData {
	s.Services = v
	return s
}

func (s *PullServicesResponseBodyData) Validate() error {
	if s.Services != nil {
		for _, item := range s.Services {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type PullServicesResponseBodyDataServices struct {
	// The name of the group.
	//
	// example:
	//
	// test
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
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
	// public
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	SaeAppId  *string `json:"SaeAppId,omitempty" xml:"SaeAppId,omitempty"`
	// The ID of the service source.
	//
	// example:
	//
	// 1
	SourceId     *string  `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	SourceIdList []*int64 `json:"SourceIdList,omitempty" xml:"SourceIdList,omitempty" type:"Repeated"`
	// The type of the service source.
	//
	// example:
	//
	// MSE
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s PullServicesResponseBodyDataServices) String() string {
	return dara.Prettify(s)
}

func (s PullServicesResponseBodyDataServices) GoString() string {
	return s.String()
}

func (s *PullServicesResponseBodyDataServices) GetGroupName() *string {
	return s.GroupName
}

func (s *PullServicesResponseBodyDataServices) GetName() *string {
	return s.Name
}

func (s *PullServicesResponseBodyDataServices) GetNamespace() *string {
	return s.Namespace
}

func (s *PullServicesResponseBodyDataServices) GetSaeAppId() *string {
	return s.SaeAppId
}

func (s *PullServicesResponseBodyDataServices) GetSourceId() *string {
	return s.SourceId
}

func (s *PullServicesResponseBodyDataServices) GetSourceIdList() []*int64 {
	return s.SourceIdList
}

func (s *PullServicesResponseBodyDataServices) GetSourceType() *string {
	return s.SourceType
}

func (s *PullServicesResponseBodyDataServices) SetGroupName(v string) *PullServicesResponseBodyDataServices {
	s.GroupName = &v
	return s
}

func (s *PullServicesResponseBodyDataServices) SetName(v string) *PullServicesResponseBodyDataServices {
	s.Name = &v
	return s
}

func (s *PullServicesResponseBodyDataServices) SetNamespace(v string) *PullServicesResponseBodyDataServices {
	s.Namespace = &v
	return s
}

func (s *PullServicesResponseBodyDataServices) SetSaeAppId(v string) *PullServicesResponseBodyDataServices {
	s.SaeAppId = &v
	return s
}

func (s *PullServicesResponseBodyDataServices) SetSourceId(v string) *PullServicesResponseBodyDataServices {
	s.SourceId = &v
	return s
}

func (s *PullServicesResponseBodyDataServices) SetSourceIdList(v []*int64) *PullServicesResponseBodyDataServices {
	s.SourceIdList = v
	return s
}

func (s *PullServicesResponseBodyDataServices) SetSourceType(v string) *PullServicesResponseBodyDataServices {
	s.SourceType = &v
	return s
}

func (s *PullServicesResponseBodyDataServices) Validate() error {
	return dara.Validate(s)
}
