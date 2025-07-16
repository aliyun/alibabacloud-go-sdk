// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInternetEndpoint(v string) *CreateAppServiceResponseBody
	GetInternetEndpoint() *string
	SetIntranetEndpoint(v string) *CreateAppServiceResponseBody
	GetIntranetEndpoint() *string
	SetRegion(v string) *CreateAppServiceResponseBody
	GetRegion() *string
	SetRequestId(v string) *CreateAppServiceResponseBody
	GetRequestId() *string
	SetServiceId(v string) *CreateAppServiceResponseBody
	GetServiceId() *string
	SetServiceName(v string) *CreateAppServiceResponseBody
	GetServiceName() *string
	SetStatus(v string) *CreateAppServiceResponseBody
	GetStatus() *string
}

type CreateAppServiceResponseBody struct {
	// The public endpoint of the service.
	//
	// example:
	//
	// http://pai-eas.cn-shanghai.****
	InternetEndpoint *string `json:"InternetEndpoint,omitempty" xml:"InternetEndpoint,omitempty"`
	// The internal endpoint of the service.
	//
	// example:
	//
	// http://pai-eas-vpc.cn-shanghai.****
	IntranetEndpoint *string `json:"IntranetEndpoint,omitempty" xml:"IntranetEndpoint,omitempty"`
	// The region ID of the service.
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
	// The service ID.
	//
	// example:
	//
	// eas-m-aaxxxddf
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The service name.
	//
	// example:
	//
	// foo
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The service state.
	//
	// example:
	//
	// Creating
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateAppServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAppServiceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppServiceResponseBody) GetInternetEndpoint() *string {
	return s.InternetEndpoint
}

func (s *CreateAppServiceResponseBody) GetIntranetEndpoint() *string {
	return s.IntranetEndpoint
}

func (s *CreateAppServiceResponseBody) GetRegion() *string {
	return s.Region
}

func (s *CreateAppServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAppServiceResponseBody) GetServiceId() *string {
	return s.ServiceId
}

func (s *CreateAppServiceResponseBody) GetServiceName() *string {
	return s.ServiceName
}

func (s *CreateAppServiceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateAppServiceResponseBody) SetInternetEndpoint(v string) *CreateAppServiceResponseBody {
	s.InternetEndpoint = &v
	return s
}

func (s *CreateAppServiceResponseBody) SetIntranetEndpoint(v string) *CreateAppServiceResponseBody {
	s.IntranetEndpoint = &v
	return s
}

func (s *CreateAppServiceResponseBody) SetRegion(v string) *CreateAppServiceResponseBody {
	s.Region = &v
	return s
}

func (s *CreateAppServiceResponseBody) SetRequestId(v string) *CreateAppServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAppServiceResponseBody) SetServiceId(v string) *CreateAppServiceResponseBody {
	s.ServiceId = &v
	return s
}

func (s *CreateAppServiceResponseBody) SetServiceName(v string) *CreateAppServiceResponseBody {
	s.ServiceName = &v
	return s
}

func (s *CreateAppServiceResponseBody) SetStatus(v string) *CreateAppServiceResponseBody {
	s.Status = &v
	return s
}

func (s *CreateAppServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
