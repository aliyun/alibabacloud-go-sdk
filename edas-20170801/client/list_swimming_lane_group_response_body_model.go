// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSwimmingLaneGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListSwimmingLaneGroupResponseBody
	GetCode() *int32
	SetData(v []*ListSwimmingLaneGroupResponseBodyData) *ListSwimmingLaneGroupResponseBody
	GetData() []*ListSwimmingLaneGroupResponseBodyData
	SetMessage(v string) *ListSwimmingLaneGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListSwimmingLaneGroupResponseBody
	GetRequestId() *string
}

type ListSwimmingLaneGroupResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data of the lane group.
	Data []*ListSwimmingLaneGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The additional information that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// BF238E37-671A-5045-B49A-0B29C166****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListSwimmingLaneGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSwimmingLaneGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListSwimmingLaneGroupResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListSwimmingLaneGroupResponseBody) GetData() []*ListSwimmingLaneGroupResponseBodyData {
	return s.Data
}

func (s *ListSwimmingLaneGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListSwimmingLaneGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSwimmingLaneGroupResponseBody) SetCode(v int32) *ListSwimmingLaneGroupResponseBody {
	s.Code = &v
	return s
}

func (s *ListSwimmingLaneGroupResponseBody) SetData(v []*ListSwimmingLaneGroupResponseBodyData) *ListSwimmingLaneGroupResponseBody {
	s.Data = v
	return s
}

func (s *ListSwimmingLaneGroupResponseBody) SetMessage(v string) *ListSwimmingLaneGroupResponseBody {
	s.Message = &v
	return s
}

func (s *ListSwimmingLaneGroupResponseBody) SetRequestId(v string) *ListSwimmingLaneGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSwimmingLaneGroupResponseBody) Validate() error {
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

type ListSwimmingLaneGroupResponseBodyData struct {
	// The applications that are related to the lane group.
	ApplicationList []*ListSwimmingLaneGroupResponseBodyDataApplicationList `json:"ApplicationList,omitempty" xml:"ApplicationList,omitempty" type:"Repeated"`
	// The information about the Enterprise Distributed Application Service (EDAS) ingress gateway.
	EntryApplication *ListSwimmingLaneGroupResponseBodyDataEntryApplication `json:"EntryApplication,omitempty" xml:"EntryApplication,omitempty" type:"Struct"`
	// The ID of the lane group.
	//
	// example:
	//
	// 257
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the lane group.
	//
	// example:
	//
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the microservices namespace.
	//
	// example:
	//
	// cn-shanghai:daily
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
}

func (s ListSwimmingLaneGroupResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListSwimmingLaneGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSwimmingLaneGroupResponseBodyData) GetApplicationList() []*ListSwimmingLaneGroupResponseBodyDataApplicationList {
	return s.ApplicationList
}

func (s *ListSwimmingLaneGroupResponseBodyData) GetEntryApplication() *ListSwimmingLaneGroupResponseBodyDataEntryApplication {
	return s.EntryApplication
}

func (s *ListSwimmingLaneGroupResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *ListSwimmingLaneGroupResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListSwimmingLaneGroupResponseBodyData) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *ListSwimmingLaneGroupResponseBodyData) SetApplicationList(v []*ListSwimmingLaneGroupResponseBodyDataApplicationList) *ListSwimmingLaneGroupResponseBodyData {
	s.ApplicationList = v
	return s
}

func (s *ListSwimmingLaneGroupResponseBodyData) SetEntryApplication(v *ListSwimmingLaneGroupResponseBodyDataEntryApplication) *ListSwimmingLaneGroupResponseBodyData {
	s.EntryApplication = v
	return s
}

func (s *ListSwimmingLaneGroupResponseBodyData) SetId(v int64) *ListSwimmingLaneGroupResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListSwimmingLaneGroupResponseBodyData) SetName(v string) *ListSwimmingLaneGroupResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListSwimmingLaneGroupResponseBodyData) SetNamespaceId(v string) *ListSwimmingLaneGroupResponseBodyData {
	s.NamespaceId = &v
	return s
}

func (s *ListSwimmingLaneGroupResponseBodyData) Validate() error {
	if s.ApplicationList != nil {
		for _, item := range s.ApplicationList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.EntryApplication != nil {
		if err := s.EntryApplication.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSwimmingLaneGroupResponseBodyDataApplicationList struct {
	// The ID of the application.
	//
	// example:
	//
	// 406073bf-afc2-4142-b3d7-629a0308****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// java-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
}

func (s ListSwimmingLaneGroupResponseBodyDataApplicationList) String() string {
	return dara.Prettify(s)
}

func (s ListSwimmingLaneGroupResponseBodyDataApplicationList) GoString() string {
	return s.String()
}

func (s *ListSwimmingLaneGroupResponseBodyDataApplicationList) GetAppId() *string {
	return s.AppId
}

func (s *ListSwimmingLaneGroupResponseBodyDataApplicationList) GetAppName() *string {
	return s.AppName
}

func (s *ListSwimmingLaneGroupResponseBodyDataApplicationList) SetAppId(v string) *ListSwimmingLaneGroupResponseBodyDataApplicationList {
	s.AppId = &v
	return s
}

func (s *ListSwimmingLaneGroupResponseBodyDataApplicationList) SetAppName(v string) *ListSwimmingLaneGroupResponseBodyDataApplicationList {
	s.AppName = &v
	return s
}

func (s *ListSwimmingLaneGroupResponseBodyDataApplicationList) Validate() error {
	return dara.Validate(s)
}

type ListSwimmingLaneGroupResponseBodyDataEntryApplication struct {
	// The ID of the application.
	//
	// example:
	//
	// 406073bf-afc2-4142-b3d7-629a0308****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// java-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The data source. Set the value to EDAS.
	//
	// example:
	//
	// EDAS
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s ListSwimmingLaneGroupResponseBodyDataEntryApplication) String() string {
	return dara.Prettify(s)
}

func (s ListSwimmingLaneGroupResponseBodyDataEntryApplication) GoString() string {
	return s.String()
}

func (s *ListSwimmingLaneGroupResponseBodyDataEntryApplication) GetAppId() *string {
	return s.AppId
}

func (s *ListSwimmingLaneGroupResponseBodyDataEntryApplication) GetAppName() *string {
	return s.AppName
}

func (s *ListSwimmingLaneGroupResponseBodyDataEntryApplication) GetSource() *string {
	return s.Source
}

func (s *ListSwimmingLaneGroupResponseBodyDataEntryApplication) SetAppId(v string) *ListSwimmingLaneGroupResponseBodyDataEntryApplication {
	s.AppId = &v
	return s
}

func (s *ListSwimmingLaneGroupResponseBodyDataEntryApplication) SetAppName(v string) *ListSwimmingLaneGroupResponseBodyDataEntryApplication {
	s.AppName = &v
	return s
}

func (s *ListSwimmingLaneGroupResponseBodyDataEntryApplication) SetSource(v string) *ListSwimmingLaneGroupResponseBodyDataEntryApplication {
	s.Source = &v
	return s
}

func (s *ListSwimmingLaneGroupResponseBodyDataEntryApplication) Validate() error {
	return dara.Validate(s)
}
