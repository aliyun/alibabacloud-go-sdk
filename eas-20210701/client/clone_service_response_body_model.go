// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloneServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInternetEndpoint(v string) *CloneServiceResponseBody
	GetInternetEndpoint() *string
	SetIntranetEndpoint(v string) *CloneServiceResponseBody
	GetIntranetEndpoint() *string
	SetRequestId(v string) *CloneServiceResponseBody
	GetRequestId() *string
	SetServiceId(v string) *CloneServiceResponseBody
	GetServiceId() *string
	SetServiceName(v string) *CloneServiceResponseBody
	GetServiceName() *string
	SetStatus(v string) *CloneServiceResponseBody
	GetStatus() *string
}

type CloneServiceResponseBody struct {
	// The public endpoint of the service.
	//
	// example:
	//
	// http://10123*****.cn-shanghai.aliyuncs.com/api/predict/echo
	InternetEndpoint *string `json:"InternetEndpoint,omitempty" xml:"InternetEndpoint,omitempty"`
	// The private endpoint of the service.
	//
	// example:
	//
	// http://10123*****.vpc.cn-shanghai.aliyuncs.com/api/predict/echo
	IntranetEndpoint *string `json:"IntranetEndpoint,omitempty" xml:"IntranetEndpoint,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The service ID.
	//
	// example:
	//
	// eas-m-r9knx7n9guf2p*****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The service name.
	//
	// example:
	//
	// foo
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The service status.
	//
	// example:
	//
	// Creating
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CloneServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloneServiceResponseBody) GoString() string {
	return s.String()
}

func (s *CloneServiceResponseBody) GetInternetEndpoint() *string {
	return s.InternetEndpoint
}

func (s *CloneServiceResponseBody) GetIntranetEndpoint() *string {
	return s.IntranetEndpoint
}

func (s *CloneServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloneServiceResponseBody) GetServiceId() *string {
	return s.ServiceId
}

func (s *CloneServiceResponseBody) GetServiceName() *string {
	return s.ServiceName
}

func (s *CloneServiceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CloneServiceResponseBody) SetInternetEndpoint(v string) *CloneServiceResponseBody {
	s.InternetEndpoint = &v
	return s
}

func (s *CloneServiceResponseBody) SetIntranetEndpoint(v string) *CloneServiceResponseBody {
	s.IntranetEndpoint = &v
	return s
}

func (s *CloneServiceResponseBody) SetRequestId(v string) *CloneServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloneServiceResponseBody) SetServiceId(v string) *CloneServiceResponseBody {
	s.ServiceId = &v
	return s
}

func (s *CloneServiceResponseBody) SetServiceName(v string) *CloneServiceResponseBody {
	s.ServiceName = &v
	return s
}

func (s *CloneServiceResponseBody) SetStatus(v string) *CloneServiceResponseBody {
	s.Status = &v
	return s
}

func (s *CloneServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
