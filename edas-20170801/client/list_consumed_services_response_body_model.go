// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConsumedServicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListConsumedServicesResponseBody
	GetCode() *int32
	SetConsumedServicesList(v *ListConsumedServicesResponseBodyConsumedServicesList) *ListConsumedServicesResponseBody
	GetConsumedServicesList() *ListConsumedServicesResponseBodyConsumedServicesList
	SetMessage(v string) *ListConsumedServicesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListConsumedServicesResponseBody
	GetRequestId() *string
}

type ListConsumedServicesResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code                 *int32                                                `json:"Code,omitempty" xml:"Code,omitempty"`
	ConsumedServicesList *ListConsumedServicesResponseBodyConsumedServicesList `json:"ConsumedServicesList,omitempty" xml:"ConsumedServicesList,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique request ID.
	//
	// example:
	//
	// a5281053-08e4-47a5-b2ab-5c0323de7b5a
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListConsumedServicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListConsumedServicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListConsumedServicesResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListConsumedServicesResponseBody) GetConsumedServicesList() *ListConsumedServicesResponseBodyConsumedServicesList {
	return s.ConsumedServicesList
}

func (s *ListConsumedServicesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListConsumedServicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListConsumedServicesResponseBody) SetCode(v int32) *ListConsumedServicesResponseBody {
	s.Code = &v
	return s
}

func (s *ListConsumedServicesResponseBody) SetConsumedServicesList(v *ListConsumedServicesResponseBodyConsumedServicesList) *ListConsumedServicesResponseBody {
	s.ConsumedServicesList = v
	return s
}

func (s *ListConsumedServicesResponseBody) SetMessage(v string) *ListConsumedServicesResponseBody {
	s.Message = &v
	return s
}

func (s *ListConsumedServicesResponseBody) SetRequestId(v string) *ListConsumedServicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListConsumedServicesResponseBody) Validate() error {
	if s.ConsumedServicesList != nil {
		if err := s.ConsumedServicesList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListConsumedServicesResponseBodyConsumedServicesList struct {
	ListConsumedServices []*ListConsumedServicesResponseBodyConsumedServicesListListConsumedServices `json:"ListConsumedServices,omitempty" xml:"ListConsumedServices,omitempty" type:"Repeated"`
}

func (s ListConsumedServicesResponseBodyConsumedServicesList) String() string {
	return dara.Prettify(s)
}

func (s ListConsumedServicesResponseBodyConsumedServicesList) GoString() string {
	return s.String()
}

func (s *ListConsumedServicesResponseBodyConsumedServicesList) GetListConsumedServices() []*ListConsumedServicesResponseBodyConsumedServicesListListConsumedServices {
	return s.ListConsumedServices
}

func (s *ListConsumedServicesResponseBodyConsumedServicesList) SetListConsumedServices(v []*ListConsumedServicesResponseBodyConsumedServicesListListConsumedServices) *ListConsumedServicesResponseBodyConsumedServicesList {
	s.ListConsumedServices = v
	return s
}

func (s *ListConsumedServicesResponseBodyConsumedServicesList) Validate() error {
	if s.ListConsumedServices != nil {
		for _, item := range s.ListConsumedServices {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListConsumedServicesResponseBodyConsumedServicesListListConsumedServices struct {
	AppId             *string                                                                         `json:"AppId,omitempty" xml:"AppId,omitempty"`
	DockerApplication *bool                                                                           `json:"DockerApplication,omitempty" xml:"DockerApplication,omitempty"`
	Group2Ip          *string                                                                         `json:"Group2Ip,omitempty" xml:"Group2Ip,omitempty"`
	Groups            *ListConsumedServicesResponseBodyConsumedServicesListListConsumedServicesGroups `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Struct"`
	Ips               *ListConsumedServicesResponseBodyConsumedServicesListListConsumedServicesIps    `json:"Ips,omitempty" xml:"Ips,omitempty" type:"Struct"`
	Name              *string                                                                         `json:"Name,omitempty" xml:"Name,omitempty"`
	Type              *string                                                                         `json:"Type,omitempty" xml:"Type,omitempty"`
	Version           *string                                                                         `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListConsumedServicesResponseBodyConsumedServicesListListConsumedServices) String() string {
	return dara.Prettify(s)
}

func (s ListConsumedServicesResponseBodyConsumedServicesListListConsumedServices) GoString() string {
	return s.String()
}

func (s *ListConsumedServicesResponseBodyConsumedServicesListListConsumedServices) GetAppId() *string {
	return s.AppId
}

func (s *ListConsumedServicesResponseBodyConsumedServicesListListConsumedServices) GetDockerApplication() *bool {
	return s.DockerApplication
}

func (s *ListConsumedServicesResponseBodyConsumedServicesListListConsumedServices) GetGroup2Ip() *string {
	return s.Group2Ip
}

func (s *ListConsumedServicesResponseBodyConsumedServicesListListConsumedServices) GetGroups() *ListConsumedServicesResponseBodyConsumedServicesListListConsumedServicesGroups {
	return s.Groups
}

func (s *ListConsumedServicesResponseBodyConsumedServicesListListConsumedServices) GetIps() *ListConsumedServicesResponseBodyConsumedServicesListListConsumedServicesIps {
	return s.Ips
}

func (s *ListConsumedServicesResponseBodyConsumedServicesListListConsumedServices) GetName() *string {
	return s.Name
}

func (s *ListConsumedServicesResponseBodyConsumedServicesListListConsumedServices) GetType() *string {
	return s.Type
}

func (s *ListConsumedServicesResponseBodyConsumedServicesListListConsumedServices) GetVersion() *string {
	return s.Version
}

func (s *ListConsumedServicesResponseBodyConsumedServicesListListConsumedServices) SetAppId(v string) *ListConsumedServicesResponseBodyConsumedServicesListListConsumedServices {
	s.AppId = &v
	return s
}

func (s *ListConsumedServicesResponseBodyConsumedServicesListListConsumedServices) SetDockerApplication(v bool) *ListConsumedServicesResponseBodyConsumedServicesListListConsumedServices {
	s.DockerApplication = &v
	return s
}

func (s *ListConsumedServicesResponseBodyConsumedServicesListListConsumedServices) SetGroup2Ip(v string) *ListConsumedServicesResponseBodyConsumedServicesListListConsumedServices {
	s.Group2Ip = &v
	return s
}

func (s *ListConsumedServicesResponseBodyConsumedServicesListListConsumedServices) SetGroups(v *ListConsumedServicesResponseBodyConsumedServicesListListConsumedServicesGroups) *ListConsumedServicesResponseBodyConsumedServicesListListConsumedServices {
	s.Groups = v
	return s
}

func (s *ListConsumedServicesResponseBodyConsumedServicesListListConsumedServices) SetIps(v *ListConsumedServicesResponseBodyConsumedServicesListListConsumedServicesIps) *ListConsumedServicesResponseBodyConsumedServicesListListConsumedServices {
	s.Ips = v
	return s
}

func (s *ListConsumedServicesResponseBodyConsumedServicesListListConsumedServices) SetName(v string) *ListConsumedServicesResponseBodyConsumedServicesListListConsumedServices {
	s.Name = &v
	return s
}

func (s *ListConsumedServicesResponseBodyConsumedServicesListListConsumedServices) SetType(v string) *ListConsumedServicesResponseBodyConsumedServicesListListConsumedServices {
	s.Type = &v
	return s
}

func (s *ListConsumedServicesResponseBodyConsumedServicesListListConsumedServices) SetVersion(v string) *ListConsumedServicesResponseBodyConsumedServicesListListConsumedServices {
	s.Version = &v
	return s
}

func (s *ListConsumedServicesResponseBodyConsumedServicesListListConsumedServices) Validate() error {
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

type ListConsumedServicesResponseBodyConsumedServicesListListConsumedServicesGroups struct {
	Group []*string `json:"group,omitempty" xml:"group,omitempty" type:"Repeated"`
}

func (s ListConsumedServicesResponseBodyConsumedServicesListListConsumedServicesGroups) String() string {
	return dara.Prettify(s)
}

func (s ListConsumedServicesResponseBodyConsumedServicesListListConsumedServicesGroups) GoString() string {
	return s.String()
}

func (s *ListConsumedServicesResponseBodyConsumedServicesListListConsumedServicesGroups) GetGroup() []*string {
	return s.Group
}

func (s *ListConsumedServicesResponseBodyConsumedServicesListListConsumedServicesGroups) SetGroup(v []*string) *ListConsumedServicesResponseBodyConsumedServicesListListConsumedServicesGroups {
	s.Group = v
	return s
}

func (s *ListConsumedServicesResponseBodyConsumedServicesListListConsumedServicesGroups) Validate() error {
	return dara.Validate(s)
}

type ListConsumedServicesResponseBodyConsumedServicesListListConsumedServicesIps struct {
	Ip []*string `json:"ip,omitempty" xml:"ip,omitempty" type:"Repeated"`
}

func (s ListConsumedServicesResponseBodyConsumedServicesListListConsumedServicesIps) String() string {
	return dara.Prettify(s)
}

func (s ListConsumedServicesResponseBodyConsumedServicesListListConsumedServicesIps) GoString() string {
	return s.String()
}

func (s *ListConsumedServicesResponseBodyConsumedServicesListListConsumedServicesIps) GetIp() []*string {
	return s.Ip
}

func (s *ListConsumedServicesResponseBodyConsumedServicesListListConsumedServicesIps) SetIp(v []*string) *ListConsumedServicesResponseBodyConsumedServicesListListConsumedServicesIps {
	s.Ip = v
	return s
}

func (s *ListConsumedServicesResponseBodyConsumedServicesListListConsumedServicesIps) Validate() error {
	return dara.Validate(s)
}
