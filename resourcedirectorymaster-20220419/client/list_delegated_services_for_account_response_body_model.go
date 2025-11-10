// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDelegatedServicesForAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDelegatedServices(v *ListDelegatedServicesForAccountResponseBodyDelegatedServices) *ListDelegatedServicesForAccountResponseBody
	GetDelegatedServices() *ListDelegatedServicesForAccountResponseBodyDelegatedServices
	SetRequestId(v string) *ListDelegatedServicesForAccountResponseBody
	GetRequestId() *string
}

type ListDelegatedServicesForAccountResponseBody struct {
	// The information about the trusted services.
	//
	// > If the value of this parameter is empty, the member is not specified as a delegated administrator account.
	DelegatedServices *ListDelegatedServicesForAccountResponseBodyDelegatedServices `json:"DelegatedServices,omitempty" xml:"DelegatedServices,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// D9C03B94-9396-4794-A74B-13DC437556A6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDelegatedServicesForAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDelegatedServicesForAccountResponseBody) GoString() string {
	return s.String()
}

func (s *ListDelegatedServicesForAccountResponseBody) GetDelegatedServices() *ListDelegatedServicesForAccountResponseBodyDelegatedServices {
	return s.DelegatedServices
}

func (s *ListDelegatedServicesForAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDelegatedServicesForAccountResponseBody) SetDelegatedServices(v *ListDelegatedServicesForAccountResponseBodyDelegatedServices) *ListDelegatedServicesForAccountResponseBody {
	s.DelegatedServices = v
	return s
}

func (s *ListDelegatedServicesForAccountResponseBody) SetRequestId(v string) *ListDelegatedServicesForAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDelegatedServicesForAccountResponseBody) Validate() error {
	if s.DelegatedServices != nil {
		if err := s.DelegatedServices.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDelegatedServicesForAccountResponseBodyDelegatedServices struct {
	DelegatedService []*ListDelegatedServicesForAccountResponseBodyDelegatedServicesDelegatedService `json:"DelegatedService,omitempty" xml:"DelegatedService,omitempty" type:"Repeated"`
}

func (s ListDelegatedServicesForAccountResponseBodyDelegatedServices) String() string {
	return dara.Prettify(s)
}

func (s ListDelegatedServicesForAccountResponseBodyDelegatedServices) GoString() string {
	return s.String()
}

func (s *ListDelegatedServicesForAccountResponseBodyDelegatedServices) GetDelegatedService() []*ListDelegatedServicesForAccountResponseBodyDelegatedServicesDelegatedService {
	return s.DelegatedService
}

func (s *ListDelegatedServicesForAccountResponseBodyDelegatedServices) SetDelegatedService(v []*ListDelegatedServicesForAccountResponseBodyDelegatedServicesDelegatedService) *ListDelegatedServicesForAccountResponseBodyDelegatedServices {
	s.DelegatedService = v
	return s
}

func (s *ListDelegatedServicesForAccountResponseBodyDelegatedServices) Validate() error {
	if s.DelegatedService != nil {
		for _, item := range s.DelegatedService {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDelegatedServicesForAccountResponseBodyDelegatedServicesDelegatedService struct {
	// The time when the member was specified as a delegated administrator account.
	//
	// example:
	//
	// 1616652684164
	DelegationEnabledTime *string `json:"DelegationEnabledTime,omitempty" xml:"DelegationEnabledTime,omitempty"`
	// The identifier of the trusted service.
	//
	// example:
	//
	// cloudfw.aliyuncs.com
	ServicePrincipal *string `json:"ServicePrincipal,omitempty" xml:"ServicePrincipal,omitempty"`
	// The status of the trusted service. Valid values:
	//
	// 	- ENABLED: enabled
	//
	// 	- DISABLED: disabled
	//
	// example:
	//
	// ENABLED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListDelegatedServicesForAccountResponseBodyDelegatedServicesDelegatedService) String() string {
	return dara.Prettify(s)
}

func (s ListDelegatedServicesForAccountResponseBodyDelegatedServicesDelegatedService) GoString() string {
	return s.String()
}

func (s *ListDelegatedServicesForAccountResponseBodyDelegatedServicesDelegatedService) GetDelegationEnabledTime() *string {
	return s.DelegationEnabledTime
}

func (s *ListDelegatedServicesForAccountResponseBodyDelegatedServicesDelegatedService) GetServicePrincipal() *string {
	return s.ServicePrincipal
}

func (s *ListDelegatedServicesForAccountResponseBodyDelegatedServicesDelegatedService) GetStatus() *string {
	return s.Status
}

func (s *ListDelegatedServicesForAccountResponseBodyDelegatedServicesDelegatedService) SetDelegationEnabledTime(v string) *ListDelegatedServicesForAccountResponseBodyDelegatedServicesDelegatedService {
	s.DelegationEnabledTime = &v
	return s
}

func (s *ListDelegatedServicesForAccountResponseBodyDelegatedServicesDelegatedService) SetServicePrincipal(v string) *ListDelegatedServicesForAccountResponseBodyDelegatedServicesDelegatedService {
	s.ServicePrincipal = &v
	return s
}

func (s *ListDelegatedServicesForAccountResponseBodyDelegatedServicesDelegatedService) SetStatus(v string) *ListDelegatedServicesForAccountResponseBodyDelegatedServicesDelegatedService {
	s.Status = &v
	return s
}

func (s *ListDelegatedServicesForAccountResponseBodyDelegatedServicesDelegatedService) Validate() error {
	return dara.Validate(s)
}
