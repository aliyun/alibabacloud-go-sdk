// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPublishedServicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListPublishedServicesResponseBody
	GetCode() *int32
	SetMessage(v string) *ListPublishedServicesResponseBody
	GetMessage() *string
	SetPublishedServicesList(v *ListPublishedServicesResponseBodyPublishedServicesList) *ListPublishedServicesResponseBody
	GetPublishedServicesList() *ListPublishedServicesResponseBodyPublishedServicesList
	SetRequestId(v string) *ListPublishedServicesResponseBody
	GetRequestId() *string
}

type ListPublishedServicesResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The published services.
	PublishedServicesList *ListPublishedServicesResponseBodyPublishedServicesList `json:"PublishedServicesList,omitempty" xml:"PublishedServicesList,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 1D6FC-4307-4583-BA6F-215F3857E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPublishedServicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPublishedServicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPublishedServicesResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListPublishedServicesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListPublishedServicesResponseBody) GetPublishedServicesList() *ListPublishedServicesResponseBodyPublishedServicesList {
	return s.PublishedServicesList
}

func (s *ListPublishedServicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPublishedServicesResponseBody) SetCode(v int32) *ListPublishedServicesResponseBody {
	s.Code = &v
	return s
}

func (s *ListPublishedServicesResponseBody) SetMessage(v string) *ListPublishedServicesResponseBody {
	s.Message = &v
	return s
}

func (s *ListPublishedServicesResponseBody) SetPublishedServicesList(v *ListPublishedServicesResponseBodyPublishedServicesList) *ListPublishedServicesResponseBody {
	s.PublishedServicesList = v
	return s
}

func (s *ListPublishedServicesResponseBody) SetRequestId(v string) *ListPublishedServicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPublishedServicesResponseBody) Validate() error {
	if s.PublishedServicesList != nil {
		if err := s.PublishedServicesList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListPublishedServicesResponseBodyPublishedServicesList struct {
	ListPublishedServices []*ListPublishedServicesResponseBodyPublishedServicesListListPublishedServices `json:"ListPublishedServices,omitempty" xml:"ListPublishedServices,omitempty" type:"Repeated"`
}

func (s ListPublishedServicesResponseBodyPublishedServicesList) String() string {
	return dara.Prettify(s)
}

func (s ListPublishedServicesResponseBodyPublishedServicesList) GoString() string {
	return s.String()
}

func (s *ListPublishedServicesResponseBodyPublishedServicesList) GetListPublishedServices() []*ListPublishedServicesResponseBodyPublishedServicesListListPublishedServices {
	return s.ListPublishedServices
}

func (s *ListPublishedServicesResponseBodyPublishedServicesList) SetListPublishedServices(v []*ListPublishedServicesResponseBodyPublishedServicesListListPublishedServices) *ListPublishedServicesResponseBodyPublishedServicesList {
	s.ListPublishedServices = v
	return s
}

func (s *ListPublishedServicesResponseBodyPublishedServicesList) Validate() error {
	if s.ListPublishedServices != nil {
		for _, item := range s.ListPublishedServices {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPublishedServicesResponseBodyPublishedServicesListListPublishedServices struct {
	// The ID of the application.
	//
	// example:
	//
	// ECD1D6FC-4307-4583-BA6F-215F3857E****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// Indicates whether the application runs in a Docker container. Valid values:
	//
	// 	- true: The application runs in a Docker container.
	//
	// 	- false: The application does not run in a Docker container.
	//
	// example:
	//
	// false
	DockerApplication *bool `json:"DockerApplication,omitempty" xml:"DockerApplication,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// ""
	Group2Ip *string                                                                            `json:"Group2Ip,omitempty" xml:"Group2Ip,omitempty"`
	Groups   *ListPublishedServicesResponseBodyPublishedServicesListListPublishedServicesGroups `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Struct"`
	Ips      *ListPublishedServicesResponseBodyPublishedServicesListListPublishedServicesIps    `json:"Ips,omitempty" xml:"Ips,omitempty" type:"Struct"`
	// The name of the published service.
	//
	// example:
	//
	// providers:com.****
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the published service.
	//
	// example:
	//
	// RESTful
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The version of the published services.
	//
	// example:
	//
	// --
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListPublishedServicesResponseBodyPublishedServicesListListPublishedServices) String() string {
	return dara.Prettify(s)
}

func (s ListPublishedServicesResponseBodyPublishedServicesListListPublishedServices) GoString() string {
	return s.String()
}

func (s *ListPublishedServicesResponseBodyPublishedServicesListListPublishedServices) GetAppId() *string {
	return s.AppId
}

func (s *ListPublishedServicesResponseBodyPublishedServicesListListPublishedServices) GetDockerApplication() *bool {
	return s.DockerApplication
}

func (s *ListPublishedServicesResponseBodyPublishedServicesListListPublishedServices) GetGroup2Ip() *string {
	return s.Group2Ip
}

func (s *ListPublishedServicesResponseBodyPublishedServicesListListPublishedServices) GetGroups() *ListPublishedServicesResponseBodyPublishedServicesListListPublishedServicesGroups {
	return s.Groups
}

func (s *ListPublishedServicesResponseBodyPublishedServicesListListPublishedServices) GetIps() *ListPublishedServicesResponseBodyPublishedServicesListListPublishedServicesIps {
	return s.Ips
}

func (s *ListPublishedServicesResponseBodyPublishedServicesListListPublishedServices) GetName() *string {
	return s.Name
}

func (s *ListPublishedServicesResponseBodyPublishedServicesListListPublishedServices) GetType() *string {
	return s.Type
}

func (s *ListPublishedServicesResponseBodyPublishedServicesListListPublishedServices) GetVersion() *string {
	return s.Version
}

func (s *ListPublishedServicesResponseBodyPublishedServicesListListPublishedServices) SetAppId(v string) *ListPublishedServicesResponseBodyPublishedServicesListListPublishedServices {
	s.AppId = &v
	return s
}

func (s *ListPublishedServicesResponseBodyPublishedServicesListListPublishedServices) SetDockerApplication(v bool) *ListPublishedServicesResponseBodyPublishedServicesListListPublishedServices {
	s.DockerApplication = &v
	return s
}

func (s *ListPublishedServicesResponseBodyPublishedServicesListListPublishedServices) SetGroup2Ip(v string) *ListPublishedServicesResponseBodyPublishedServicesListListPublishedServices {
	s.Group2Ip = &v
	return s
}

func (s *ListPublishedServicesResponseBodyPublishedServicesListListPublishedServices) SetGroups(v *ListPublishedServicesResponseBodyPublishedServicesListListPublishedServicesGroups) *ListPublishedServicesResponseBodyPublishedServicesListListPublishedServices {
	s.Groups = v
	return s
}

func (s *ListPublishedServicesResponseBodyPublishedServicesListListPublishedServices) SetIps(v *ListPublishedServicesResponseBodyPublishedServicesListListPublishedServicesIps) *ListPublishedServicesResponseBodyPublishedServicesListListPublishedServices {
	s.Ips = v
	return s
}

func (s *ListPublishedServicesResponseBodyPublishedServicesListListPublishedServices) SetName(v string) *ListPublishedServicesResponseBodyPublishedServicesListListPublishedServices {
	s.Name = &v
	return s
}

func (s *ListPublishedServicesResponseBodyPublishedServicesListListPublishedServices) SetType(v string) *ListPublishedServicesResponseBodyPublishedServicesListListPublishedServices {
	s.Type = &v
	return s
}

func (s *ListPublishedServicesResponseBodyPublishedServicesListListPublishedServices) SetVersion(v string) *ListPublishedServicesResponseBodyPublishedServicesListListPublishedServices {
	s.Version = &v
	return s
}

func (s *ListPublishedServicesResponseBodyPublishedServicesListListPublishedServices) Validate() error {
	if s.Groups != nil {
		if err := s.Groups.Validate(); err != nil {
			return err
		}
	}
	if s.Ips != nil {
		if err := s.Ips.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListPublishedServicesResponseBodyPublishedServicesListListPublishedServicesGroups struct {
	Group []*string `json:"group,omitempty" xml:"group,omitempty" type:"Repeated"`
}

func (s ListPublishedServicesResponseBodyPublishedServicesListListPublishedServicesGroups) String() string {
	return dara.Prettify(s)
}

func (s ListPublishedServicesResponseBodyPublishedServicesListListPublishedServicesGroups) GoString() string {
	return s.String()
}

func (s *ListPublishedServicesResponseBodyPublishedServicesListListPublishedServicesGroups) GetGroup() []*string {
	return s.Group
}

func (s *ListPublishedServicesResponseBodyPublishedServicesListListPublishedServicesGroups) SetGroup(v []*string) *ListPublishedServicesResponseBodyPublishedServicesListListPublishedServicesGroups {
	s.Group = v
	return s
}

func (s *ListPublishedServicesResponseBodyPublishedServicesListListPublishedServicesGroups) Validate() error {
	return dara.Validate(s)
}

type ListPublishedServicesResponseBodyPublishedServicesListListPublishedServicesIps struct {
	Ip []*string `json:"ip,omitempty" xml:"ip,omitempty" type:"Repeated"`
}

func (s ListPublishedServicesResponseBodyPublishedServicesListListPublishedServicesIps) String() string {
	return dara.Prettify(s)
}

func (s ListPublishedServicesResponseBodyPublishedServicesListListPublishedServicesIps) GoString() string {
	return s.String()
}

func (s *ListPublishedServicesResponseBodyPublishedServicesListListPublishedServicesIps) GetIp() []*string {
	return s.Ip
}

func (s *ListPublishedServicesResponseBodyPublishedServicesListListPublishedServicesIps) SetIp(v []*string) *ListPublishedServicesResponseBodyPublishedServicesListListPublishedServicesIps {
	s.Ip = v
	return s
}

func (s *ListPublishedServicesResponseBodyPublishedServicesListListPublishedServicesIps) Validate() error {
	return dara.Validate(s)
}
