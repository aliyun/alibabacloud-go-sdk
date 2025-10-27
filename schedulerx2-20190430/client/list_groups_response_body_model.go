// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListGroupsResponseBody
	GetCode() *int32
	SetData(v *ListGroupsResponseBodyData) *ListGroupsResponseBody
	GetData() *ListGroupsResponseBodyData
	SetMessage(v string) *ListGroupsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListGroupsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListGroupsResponseBody
	GetSuccess() *bool
}

type ListGroupsResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The applications.
	Data *ListGroupsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 71BCC0E3-64B2-4B63-A870-AFB64EBCB58A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListGroupsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListGroupsResponseBody) GetData() *ListGroupsResponseBodyData {
	return s.Data
}

func (s *ListGroupsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListGroupsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListGroupsResponseBody) SetCode(v int32) *ListGroupsResponseBody {
	s.Code = &v
	return s
}

func (s *ListGroupsResponseBody) SetData(v *ListGroupsResponseBodyData) *ListGroupsResponseBody {
	s.Data = v
	return s
}

func (s *ListGroupsResponseBody) SetMessage(v string) *ListGroupsResponseBody {
	s.Message = &v
	return s
}

func (s *ListGroupsResponseBody) SetRequestId(v string) *ListGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGroupsResponseBody) SetSuccess(v bool) *ListGroupsResponseBody {
	s.Success = &v
	return s
}

func (s *ListGroupsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListGroupsResponseBodyData struct {
	// The applications and their details.
	AppGroups []*ListGroupsResponseBodyDataAppGroups `json:"AppGroups,omitempty" xml:"AppGroups,omitempty" type:"Repeated"`
}

func (s ListGroupsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListGroupsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListGroupsResponseBodyData) GetAppGroups() []*ListGroupsResponseBodyDataAppGroups {
	return s.AppGroups
}

func (s *ListGroupsResponseBodyData) SetAppGroups(v []*ListGroupsResponseBodyDataAppGroups) *ListGroupsResponseBodyData {
	s.AppGroups = v
	return s
}

func (s *ListGroupsResponseBodyData) Validate() error {
	if s.AppGroups != nil {
		for _, item := range s.AppGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListGroupsResponseBodyDataAppGroups struct {
	// The application group ID.
	//
	// example:
	//
	// 1
	AppGroupId *int64 `json:"AppGroupId,omitempty" xml:"AppGroupId,omitempty"`
	// The AppKey for the application.
	//
	// example:
	//
	// a3G77O6NZxq/lyo1NC****==
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// DocTest
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The application version. 1: Basic version, 2: Professional version.
	//
	// example:
	//
	// 2
	AppVersion *int32 `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// The description of the application.
	//
	// example:
	//
	// Test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The application ID.
	//
	// example:
	//
	// DocTest.Group
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the namespace.
	//
	// example:
	//
	// 1a72ecb1-b4cc-400a-a71b-20cdec9b****
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s ListGroupsResponseBodyDataAppGroups) String() string {
	return dara.Prettify(s)
}

func (s ListGroupsResponseBodyDataAppGroups) GoString() string {
	return s.String()
}

func (s *ListGroupsResponseBodyDataAppGroups) GetAppGroupId() *int64 {
	return s.AppGroupId
}

func (s *ListGroupsResponseBodyDataAppGroups) GetAppKey() *string {
	return s.AppKey
}

func (s *ListGroupsResponseBodyDataAppGroups) GetAppName() *string {
	return s.AppName
}

func (s *ListGroupsResponseBodyDataAppGroups) GetAppVersion() *int32 {
	return s.AppVersion
}

func (s *ListGroupsResponseBodyDataAppGroups) GetDescription() *string {
	return s.Description
}

func (s *ListGroupsResponseBodyDataAppGroups) GetGroupId() *string {
	return s.GroupId
}

func (s *ListGroupsResponseBodyDataAppGroups) GetNamespace() *string {
	return s.Namespace
}

func (s *ListGroupsResponseBodyDataAppGroups) SetAppGroupId(v int64) *ListGroupsResponseBodyDataAppGroups {
	s.AppGroupId = &v
	return s
}

func (s *ListGroupsResponseBodyDataAppGroups) SetAppKey(v string) *ListGroupsResponseBodyDataAppGroups {
	s.AppKey = &v
	return s
}

func (s *ListGroupsResponseBodyDataAppGroups) SetAppName(v string) *ListGroupsResponseBodyDataAppGroups {
	s.AppName = &v
	return s
}

func (s *ListGroupsResponseBodyDataAppGroups) SetAppVersion(v int32) *ListGroupsResponseBodyDataAppGroups {
	s.AppVersion = &v
	return s
}

func (s *ListGroupsResponseBodyDataAppGroups) SetDescription(v string) *ListGroupsResponseBodyDataAppGroups {
	s.Description = &v
	return s
}

func (s *ListGroupsResponseBodyDataAppGroups) SetGroupId(v string) *ListGroupsResponseBodyDataAppGroups {
	s.GroupId = &v
	return s
}

func (s *ListGroupsResponseBodyDataAppGroups) SetNamespace(v string) *ListGroupsResponseBodyDataAppGroups {
	s.Namespace = &v
	return s
}

func (s *ListGroupsResponseBodyDataAppGroups) Validate() error {
	return dara.Validate(s)
}
