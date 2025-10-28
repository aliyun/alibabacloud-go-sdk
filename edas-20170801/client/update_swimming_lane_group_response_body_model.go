// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSwimmingLaneGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateSwimmingLaneGroupResponseBody
	GetCode() *int32
	SetData(v *UpdateSwimmingLaneGroupResponseBodyData) *UpdateSwimmingLaneGroupResponseBody
	GetData() *UpdateSwimmingLaneGroupResponseBodyData
	SetMessage(v string) *UpdateSwimmingLaneGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateSwimmingLaneGroupResponseBody
	GetRequestId() *string
}

type UpdateSwimmingLaneGroupResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data that is returned.
	Data *UpdateSwimmingLaneGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 7580ED24-A2F0-5ECC-9F2B-B92E2509****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSwimmingLaneGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSwimmingLaneGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSwimmingLaneGroupResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateSwimmingLaneGroupResponseBody) GetData() *UpdateSwimmingLaneGroupResponseBodyData {
	return s.Data
}

func (s *UpdateSwimmingLaneGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateSwimmingLaneGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSwimmingLaneGroupResponseBody) SetCode(v int32) *UpdateSwimmingLaneGroupResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateSwimmingLaneGroupResponseBody) SetData(v *UpdateSwimmingLaneGroupResponseBodyData) *UpdateSwimmingLaneGroupResponseBody {
	s.Data = v
	return s
}

func (s *UpdateSwimmingLaneGroupResponseBody) SetMessage(v string) *UpdateSwimmingLaneGroupResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateSwimmingLaneGroupResponseBody) SetRequestId(v string) *UpdateSwimmingLaneGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSwimmingLaneGroupResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateSwimmingLaneGroupResponseBodyData struct {
	// The list of applications related to the lane group.
	ApplicationList []*UpdateSwimmingLaneGroupResponseBodyDataApplicationList `json:"ApplicationList,omitempty" xml:"ApplicationList,omitempty" type:"Repeated"`
	// The EDAS ingress gateway information.
	EntryApplication *UpdateSwimmingLaneGroupResponseBodyDataEntryApplication `json:"EntryApplication,omitempty" xml:"EntryApplication,omitempty" type:"Struct"`
	// The ID of the lane group.
	//
	// example:
	//
	// 98
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the lane group.
	//
	// example:
	//
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the namespace.
	//
	// example:
	//
	// cn-hangzhou:test
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
}

func (s UpdateSwimmingLaneGroupResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateSwimmingLaneGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateSwimmingLaneGroupResponseBodyData) GetApplicationList() []*UpdateSwimmingLaneGroupResponseBodyDataApplicationList {
	return s.ApplicationList
}

func (s *UpdateSwimmingLaneGroupResponseBodyData) GetEntryApplication() *UpdateSwimmingLaneGroupResponseBodyDataEntryApplication {
	return s.EntryApplication
}

func (s *UpdateSwimmingLaneGroupResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *UpdateSwimmingLaneGroupResponseBodyData) GetName() *string {
	return s.Name
}

func (s *UpdateSwimmingLaneGroupResponseBodyData) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *UpdateSwimmingLaneGroupResponseBodyData) SetApplicationList(v []*UpdateSwimmingLaneGroupResponseBodyDataApplicationList) *UpdateSwimmingLaneGroupResponseBodyData {
	s.ApplicationList = v
	return s
}

func (s *UpdateSwimmingLaneGroupResponseBodyData) SetEntryApplication(v *UpdateSwimmingLaneGroupResponseBodyDataEntryApplication) *UpdateSwimmingLaneGroupResponseBodyData {
	s.EntryApplication = v
	return s
}

func (s *UpdateSwimmingLaneGroupResponseBodyData) SetId(v int64) *UpdateSwimmingLaneGroupResponseBodyData {
	s.Id = &v
	return s
}

func (s *UpdateSwimmingLaneGroupResponseBodyData) SetName(v string) *UpdateSwimmingLaneGroupResponseBodyData {
	s.Name = &v
	return s
}

func (s *UpdateSwimmingLaneGroupResponseBodyData) SetNamespaceId(v string) *UpdateSwimmingLaneGroupResponseBodyData {
	s.NamespaceId = &v
	return s
}

func (s *UpdateSwimmingLaneGroupResponseBodyData) Validate() error {
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

type UpdateSwimmingLaneGroupResponseBodyDataApplicationList struct {
	// The ID of the application.
	//
	// example:
	//
	// 476d26d9-b54c-40b8-8af9-d898cdc2****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
}

func (s UpdateSwimmingLaneGroupResponseBodyDataApplicationList) String() string {
	return dara.Prettify(s)
}

func (s UpdateSwimmingLaneGroupResponseBodyDataApplicationList) GoString() string {
	return s.String()
}

func (s *UpdateSwimmingLaneGroupResponseBodyDataApplicationList) GetAppId() *string {
	return s.AppId
}

func (s *UpdateSwimmingLaneGroupResponseBodyDataApplicationList) GetAppName() *string {
	return s.AppName
}

func (s *UpdateSwimmingLaneGroupResponseBodyDataApplicationList) SetAppId(v string) *UpdateSwimmingLaneGroupResponseBodyDataApplicationList {
	s.AppId = &v
	return s
}

func (s *UpdateSwimmingLaneGroupResponseBodyDataApplicationList) SetAppName(v string) *UpdateSwimmingLaneGroupResponseBodyDataApplicationList {
	s.AppName = &v
	return s
}

func (s *UpdateSwimmingLaneGroupResponseBodyDataApplicationList) Validate() error {
	return dara.Validate(s)
}

type UpdateSwimmingLaneGroupResponseBodyDataEntryApplication struct {
	// The ID of the application.
	//
	// example:
	//
	// d52c9de9-53d0-4191-aa72-88974a6f****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// test-gateway
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
}

func (s UpdateSwimmingLaneGroupResponseBodyDataEntryApplication) String() string {
	return dara.Prettify(s)
}

func (s UpdateSwimmingLaneGroupResponseBodyDataEntryApplication) GoString() string {
	return s.String()
}

func (s *UpdateSwimmingLaneGroupResponseBodyDataEntryApplication) GetAppId() *string {
	return s.AppId
}

func (s *UpdateSwimmingLaneGroupResponseBodyDataEntryApplication) GetAppName() *string {
	return s.AppName
}

func (s *UpdateSwimmingLaneGroupResponseBodyDataEntryApplication) SetAppId(v string) *UpdateSwimmingLaneGroupResponseBodyDataEntryApplication {
	s.AppId = &v
	return s
}

func (s *UpdateSwimmingLaneGroupResponseBodyDataEntryApplication) SetAppName(v string) *UpdateSwimmingLaneGroupResponseBodyDataEntryApplication {
	s.AppName = &v
	return s
}

func (s *UpdateSwimmingLaneGroupResponseBodyDataEntryApplication) Validate() error {
	return dara.Validate(s)
}
