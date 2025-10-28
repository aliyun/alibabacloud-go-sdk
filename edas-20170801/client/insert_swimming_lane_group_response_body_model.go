// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertSwimmingLaneGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *InsertSwimmingLaneGroupResponseBody
	GetCode() *int32
	SetData(v *InsertSwimmingLaneGroupResponseBodyData) *InsertSwimmingLaneGroupResponseBody
	GetData() *InsertSwimmingLaneGroupResponseBodyData
	SetMessage(v string) *InsertSwimmingLaneGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *InsertSwimmingLaneGroupResponseBody
	GetRequestId() *string
}

type InsertSwimmingLaneGroupResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data that is returned.
	Data *InsertSwimmingLaneGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// D5268CAC-D356-5C8D-BC7C-FBE0D13B****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InsertSwimmingLaneGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InsertSwimmingLaneGroupResponseBody) GoString() string {
	return s.String()
}

func (s *InsertSwimmingLaneGroupResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *InsertSwimmingLaneGroupResponseBody) GetData() *InsertSwimmingLaneGroupResponseBodyData {
	return s.Data
}

func (s *InsertSwimmingLaneGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InsertSwimmingLaneGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InsertSwimmingLaneGroupResponseBody) SetCode(v int32) *InsertSwimmingLaneGroupResponseBody {
	s.Code = &v
	return s
}

func (s *InsertSwimmingLaneGroupResponseBody) SetData(v *InsertSwimmingLaneGroupResponseBodyData) *InsertSwimmingLaneGroupResponseBody {
	s.Data = v
	return s
}

func (s *InsertSwimmingLaneGroupResponseBody) SetMessage(v string) *InsertSwimmingLaneGroupResponseBody {
	s.Message = &v
	return s
}

func (s *InsertSwimmingLaneGroupResponseBody) SetRequestId(v string) *InsertSwimmingLaneGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *InsertSwimmingLaneGroupResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InsertSwimmingLaneGroupResponseBodyData struct {
	// The list of all applications that are related to the lane group.
	ApplicationList *InsertSwimmingLaneGroupResponseBodyDataApplicationList `json:"ApplicationList,omitempty" xml:"ApplicationList,omitempty" type:"Struct"`
	// The information about the Enterprise Distributed Application Service (EDAS) ingress gateway.
	EntryApplication *InsertSwimmingLaneGroupResponseBodyDataEntryApplication `json:"EntryApplication,omitempty" xml:"EntryApplication,omitempty" type:"Struct"`
	// The ID of the lane group.
	//
	// example:
	//
	// 64
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the lane group.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the namespace.
	//
	// example:
	//
	// cn-hangzhou:test
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
}

func (s InsertSwimmingLaneGroupResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s InsertSwimmingLaneGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *InsertSwimmingLaneGroupResponseBodyData) GetApplicationList() *InsertSwimmingLaneGroupResponseBodyDataApplicationList {
	return s.ApplicationList
}

func (s *InsertSwimmingLaneGroupResponseBodyData) GetEntryApplication() *InsertSwimmingLaneGroupResponseBodyDataEntryApplication {
	return s.EntryApplication
}

func (s *InsertSwimmingLaneGroupResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *InsertSwimmingLaneGroupResponseBodyData) GetName() *string {
	return s.Name
}

func (s *InsertSwimmingLaneGroupResponseBodyData) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *InsertSwimmingLaneGroupResponseBodyData) SetApplicationList(v *InsertSwimmingLaneGroupResponseBodyDataApplicationList) *InsertSwimmingLaneGroupResponseBodyData {
	s.ApplicationList = v
	return s
}

func (s *InsertSwimmingLaneGroupResponseBodyData) SetEntryApplication(v *InsertSwimmingLaneGroupResponseBodyDataEntryApplication) *InsertSwimmingLaneGroupResponseBodyData {
	s.EntryApplication = v
	return s
}

func (s *InsertSwimmingLaneGroupResponseBodyData) SetId(v int64) *InsertSwimmingLaneGroupResponseBodyData {
	s.Id = &v
	return s
}

func (s *InsertSwimmingLaneGroupResponseBodyData) SetName(v string) *InsertSwimmingLaneGroupResponseBodyData {
	s.Name = &v
	return s
}

func (s *InsertSwimmingLaneGroupResponseBodyData) SetNamespaceId(v string) *InsertSwimmingLaneGroupResponseBodyData {
	s.NamespaceId = &v
	return s
}

func (s *InsertSwimmingLaneGroupResponseBodyData) Validate() error {
	if s.ApplicationList != nil {
		if err := s.ApplicationList.Validate(); err != nil {
			return err
		}
	}
	if s.EntryApplication != nil {
		if err := s.EntryApplication.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InsertSwimmingLaneGroupResponseBodyDataApplicationList struct {
	Application []*InsertSwimmingLaneGroupResponseBodyDataApplicationListApplication `json:"Application,omitempty" xml:"Application,omitempty" type:"Repeated"`
}

func (s InsertSwimmingLaneGroupResponseBodyDataApplicationList) String() string {
	return dara.Prettify(s)
}

func (s InsertSwimmingLaneGroupResponseBodyDataApplicationList) GoString() string {
	return s.String()
}

func (s *InsertSwimmingLaneGroupResponseBodyDataApplicationList) GetApplication() []*InsertSwimmingLaneGroupResponseBodyDataApplicationListApplication {
	return s.Application
}

func (s *InsertSwimmingLaneGroupResponseBodyDataApplicationList) SetApplication(v []*InsertSwimmingLaneGroupResponseBodyDataApplicationListApplication) *InsertSwimmingLaneGroupResponseBodyDataApplicationList {
	s.Application = v
	return s
}

func (s *InsertSwimmingLaneGroupResponseBodyDataApplicationList) Validate() error {
	if s.Application != nil {
		for _, item := range s.Application {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type InsertSwimmingLaneGroupResponseBodyDataApplicationListApplication struct {
	// The ID of the application.
	//
	// example:
	//
	// bdb251cc-02a6-48dd-891b-2ab21b25****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
}

func (s InsertSwimmingLaneGroupResponseBodyDataApplicationListApplication) String() string {
	return dara.Prettify(s)
}

func (s InsertSwimmingLaneGroupResponseBodyDataApplicationListApplication) GoString() string {
	return s.String()
}

func (s *InsertSwimmingLaneGroupResponseBodyDataApplicationListApplication) GetAppId() *string {
	return s.AppId
}

func (s *InsertSwimmingLaneGroupResponseBodyDataApplicationListApplication) GetAppName() *string {
	return s.AppName
}

func (s *InsertSwimmingLaneGroupResponseBodyDataApplicationListApplication) SetAppId(v string) *InsertSwimmingLaneGroupResponseBodyDataApplicationListApplication {
	s.AppId = &v
	return s
}

func (s *InsertSwimmingLaneGroupResponseBodyDataApplicationListApplication) SetAppName(v string) *InsertSwimmingLaneGroupResponseBodyDataApplicationListApplication {
	s.AppName = &v
	return s
}

func (s *InsertSwimmingLaneGroupResponseBodyDataApplicationListApplication) Validate() error {
	return dara.Validate(s)
}

type InsertSwimmingLaneGroupResponseBodyDataEntryApplication struct {
	// The ID of the application.
	//
	// example:
	//
	// bdb251cc-02a6-48dd-891b-2ab21b25c****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
}

func (s InsertSwimmingLaneGroupResponseBodyDataEntryApplication) String() string {
	return dara.Prettify(s)
}

func (s InsertSwimmingLaneGroupResponseBodyDataEntryApplication) GoString() string {
	return s.String()
}

func (s *InsertSwimmingLaneGroupResponseBodyDataEntryApplication) GetAppId() *string {
	return s.AppId
}

func (s *InsertSwimmingLaneGroupResponseBodyDataEntryApplication) GetAppName() *string {
	return s.AppName
}

func (s *InsertSwimmingLaneGroupResponseBodyDataEntryApplication) SetAppId(v string) *InsertSwimmingLaneGroupResponseBodyDataEntryApplication {
	s.AppId = &v
	return s
}

func (s *InsertSwimmingLaneGroupResponseBodyDataEntryApplication) SetAppName(v string) *InsertSwimmingLaneGroupResponseBodyDataEntryApplication {
	s.AppName = &v
	return s
}

func (s *InsertSwimmingLaneGroupResponseBodyDataEntryApplication) Validate() error {
	return dara.Validate(s)
}
