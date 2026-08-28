// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMigrationNamespacedServicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetMigrationNamespacedServicesResponseBody
	GetCode() *string
	SetData(v *GetMigrationNamespacedServicesResponseBodyData) *GetMigrationNamespacedServicesResponseBody
	GetData() *GetMigrationNamespacedServicesResponseBodyData
	SetMessage(v string) *GetMigrationNamespacedServicesResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetMigrationNamespacedServicesResponseBody
	GetRequestId() *string
}

type GetMigrationNamespacedServicesResponseBody struct {
	// example:
	//
	// 200
	Code *string                                         `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetMigrationNamespacedServicesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 7D513911-206E-5E93-9C9E-71D63C0D68E7
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetMigrationNamespacedServicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMigrationNamespacedServicesResponseBody) GoString() string {
	return s.String()
}

func (s *GetMigrationNamespacedServicesResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetMigrationNamespacedServicesResponseBody) GetData() *GetMigrationNamespacedServicesResponseBodyData {
	return s.Data
}

func (s *GetMigrationNamespacedServicesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetMigrationNamespacedServicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMigrationNamespacedServicesResponseBody) SetCode(v string) *GetMigrationNamespacedServicesResponseBody {
	s.Code = &v
	return s
}

func (s *GetMigrationNamespacedServicesResponseBody) SetData(v *GetMigrationNamespacedServicesResponseBodyData) *GetMigrationNamespacedServicesResponseBody {
	s.Data = v
	return s
}

func (s *GetMigrationNamespacedServicesResponseBody) SetMessage(v string) *GetMigrationNamespacedServicesResponseBody {
	s.Message = &v
	return s
}

func (s *GetMigrationNamespacedServicesResponseBody) SetRequestId(v string) *GetMigrationNamespacedServicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMigrationNamespacedServicesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMigrationNamespacedServicesResponseBodyData struct {
	NamespacedServices []*GetMigrationNamespacedServicesResponseBodyDataNamespacedServices `json:"namespacedServices,omitempty" xml:"namespacedServices,omitempty" type:"Repeated"`
}

func (s GetMigrationNamespacedServicesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetMigrationNamespacedServicesResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMigrationNamespacedServicesResponseBodyData) GetNamespacedServices() []*GetMigrationNamespacedServicesResponseBodyDataNamespacedServices {
	return s.NamespacedServices
}

func (s *GetMigrationNamespacedServicesResponseBodyData) SetNamespacedServices(v []*GetMigrationNamespacedServicesResponseBodyDataNamespacedServices) *GetMigrationNamespacedServicesResponseBodyData {
	s.NamespacedServices = v
	return s
}

func (s *GetMigrationNamespacedServicesResponseBodyData) Validate() error {
	if s.NamespacedServices != nil {
		for _, item := range s.NamespacedServices {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetMigrationNamespacedServicesResponseBodyDataNamespacedServices struct {
	// example:
	//
	// default
	Namespace *string                                                                     `json:"namespace,omitempty" xml:"namespace,omitempty"`
	Services  []*GetMigrationNamespacedServicesResponseBodyDataNamespacedServicesServices `json:"services,omitempty" xml:"services,omitempty" type:"Repeated"`
}

func (s GetMigrationNamespacedServicesResponseBodyDataNamespacedServices) String() string {
	return dara.Prettify(s)
}

func (s GetMigrationNamespacedServicesResponseBodyDataNamespacedServices) GoString() string {
	return s.String()
}

func (s *GetMigrationNamespacedServicesResponseBodyDataNamespacedServices) GetNamespace() *string {
	return s.Namespace
}

func (s *GetMigrationNamespacedServicesResponseBodyDataNamespacedServices) GetServices() []*GetMigrationNamespacedServicesResponseBodyDataNamespacedServicesServices {
	return s.Services
}

func (s *GetMigrationNamespacedServicesResponseBodyDataNamespacedServices) SetNamespace(v string) *GetMigrationNamespacedServicesResponseBodyDataNamespacedServices {
	s.Namespace = &v
	return s
}

func (s *GetMigrationNamespacedServicesResponseBodyDataNamespacedServices) SetServices(v []*GetMigrationNamespacedServicesResponseBodyDataNamespacedServicesServices) *GetMigrationNamespacedServicesResponseBodyDataNamespacedServices {
	s.Services = v
	return s
}

func (s *GetMigrationNamespacedServicesResponseBodyDataNamespacedServices) Validate() error {
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

type GetMigrationNamespacedServicesResponseBodyDataNamespacedServicesServices struct {
	// example:
	//
	// nginx-ingress-lb
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// lb-xxxx / nlb-xxxx
	SlbId *string `json:"slbId,omitempty" xml:"slbId,omitempty"`
}

func (s GetMigrationNamespacedServicesResponseBodyDataNamespacedServicesServices) String() string {
	return dara.Prettify(s)
}

func (s GetMigrationNamespacedServicesResponseBodyDataNamespacedServicesServices) GoString() string {
	return s.String()
}

func (s *GetMigrationNamespacedServicesResponseBodyDataNamespacedServicesServices) GetName() *string {
	return s.Name
}

func (s *GetMigrationNamespacedServicesResponseBodyDataNamespacedServicesServices) GetSlbId() *string {
	return s.SlbId
}

func (s *GetMigrationNamespacedServicesResponseBodyDataNamespacedServicesServices) SetName(v string) *GetMigrationNamespacedServicesResponseBodyDataNamespacedServicesServices {
	s.Name = &v
	return s
}

func (s *GetMigrationNamespacedServicesResponseBodyDataNamespacedServicesServices) SetSlbId(v string) *GetMigrationNamespacedServicesResponseBodyDataNamespacedServicesServices {
	s.SlbId = &v
	return s
}

func (s *GetMigrationNamespacedServicesResponseBodyDataNamespacedServicesServices) Validate() error {
	return dara.Validate(s)
}
