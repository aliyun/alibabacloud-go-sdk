// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInternetEndpoint(v string) *CreateServiceResponseBody
	GetInternetEndpoint() *string
	SetIntranetEndpoint(v string) *CreateServiceResponseBody
	GetIntranetEndpoint() *string
	SetRegion(v string) *CreateServiceResponseBody
	GetRegion() *string
	SetRequestId(v string) *CreateServiceResponseBody
	GetRequestId() *string
	SetServiceId(v string) *CreateServiceResponseBody
	GetServiceId() *string
	SetServiceName(v string) *CreateServiceResponseBody
	GetServiceName() *string
	SetStatus(v string) *CreateServiceResponseBody
	GetStatus() *string
}

type CreateServiceResponseBody struct {
	// The public endpoint of the created service.
	//
	// example:
	//
	// http://pai-eas.vpc.cn-shanghai.****
	InternetEndpoint *string `json:"InternetEndpoint,omitempty" xml:"InternetEndpoint,omitempty"`
	// The internal endpoint of the created service.
	//
	// example:
	//
	// http://pai-eas.cn-shanghai.****
	IntranetEndpoint *string `json:"IntranetEndpoint,omitempty" xml:"IntranetEndpoint,omitempty"`
	// The region ID of the created service.
	//
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the created service.
	//
	// example:
	//
	// eas-m-aaxxxddf
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The name of the created service.
	//
	// example:
	//
	// yourname
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The service state.
	//
	// example:
	//
	// Creating
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceResponseBody) GetInternetEndpoint() *string {
	return s.InternetEndpoint
}

func (s *CreateServiceResponseBody) GetIntranetEndpoint() *string {
	return s.IntranetEndpoint
}

func (s *CreateServiceResponseBody) GetRegion() *string {
	return s.Region
}

func (s *CreateServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateServiceResponseBody) GetServiceId() *string {
	return s.ServiceId
}

func (s *CreateServiceResponseBody) GetServiceName() *string {
	return s.ServiceName
}

func (s *CreateServiceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateServiceResponseBody) SetInternetEndpoint(v string) *CreateServiceResponseBody {
	s.InternetEndpoint = &v
	return s
}

func (s *CreateServiceResponseBody) SetIntranetEndpoint(v string) *CreateServiceResponseBody {
	s.IntranetEndpoint = &v
	return s
}

func (s *CreateServiceResponseBody) SetRegion(v string) *CreateServiceResponseBody {
	s.Region = &v
	return s
}

func (s *CreateServiceResponseBody) SetRequestId(v string) *CreateServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServiceResponseBody) SetServiceId(v string) *CreateServiceResponseBody {
	s.ServiceId = &v
	return s
}

func (s *CreateServiceResponseBody) SetServiceName(v string) *CreateServiceResponseBody {
	s.ServiceName = &v
	return s
}

func (s *CreateServiceResponseBody) SetStatus(v string) *CreateServiceResponseBody {
	s.Status = &v
	return s
}

func (s *CreateServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
