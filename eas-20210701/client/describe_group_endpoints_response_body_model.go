// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGroupEndpointsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *DescribeGroupEndpointsResponseBody
	GetAccessToken() *string
	SetEndpoints(v []*DescribeGroupEndpointsResponseBodyEndpoints) *DescribeGroupEndpointsResponseBody
	GetEndpoints() []*DescribeGroupEndpointsResponseBodyEndpoints
	SetMessage(v string) *DescribeGroupEndpointsResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeGroupEndpointsResponseBody
	GetRequestId() *string
}

type DescribeGroupEndpointsResponseBody struct {
	// The service token.
	//
	// example:
	//
	// Nzc5N2FhNTM4OTQ0YzBmYTIy****ZTUxN2NkYjg4MTJmMWQxZmY1****
	AccessToken *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	// The endpoints of the service group.
	Endpoints []*DescribeGroupEndpointsResponseBodyEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Repeated"`
	// The response message.
	//
	// example:
	//
	// Execution successful.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 890772EF-3AD6-129A-8E15-8F349C944783
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeGroupEndpointsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupEndpointsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGroupEndpointsResponseBody) GetAccessToken() *string {
	return s.AccessToken
}

func (s *DescribeGroupEndpointsResponseBody) GetEndpoints() []*DescribeGroupEndpointsResponseBodyEndpoints {
	return s.Endpoints
}

func (s *DescribeGroupEndpointsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeGroupEndpointsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeGroupEndpointsResponseBody) SetAccessToken(v string) *DescribeGroupEndpointsResponseBody {
	s.AccessToken = &v
	return s
}

func (s *DescribeGroupEndpointsResponseBody) SetEndpoints(v []*DescribeGroupEndpointsResponseBodyEndpoints) *DescribeGroupEndpointsResponseBody {
	s.Endpoints = v
	return s
}

func (s *DescribeGroupEndpointsResponseBody) SetMessage(v string) *DescribeGroupEndpointsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeGroupEndpointsResponseBody) SetRequestId(v string) *DescribeGroupEndpointsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGroupEndpointsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeGroupEndpointsResponseBodyEndpoints struct {
	BackendId         *string   `json:"BackendId,omitempty" xml:"BackendId,omitempty"`
	EndpointType      *string   `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	InternetEndpoints []*string `json:"InternetEndpoints,omitempty" xml:"InternetEndpoints,omitempty" type:"Repeated"`
	IntranetEndpoints []*string `json:"IntranetEndpoints,omitempty" xml:"IntranetEndpoints,omitempty" type:"Repeated"`
	PathType          *string   `json:"PathType,omitempty" xml:"PathType,omitempty"`
	Port              *int32    `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s DescribeGroupEndpointsResponseBodyEndpoints) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupEndpointsResponseBodyEndpoints) GoString() string {
	return s.String()
}

func (s *DescribeGroupEndpointsResponseBodyEndpoints) GetBackendId() *string {
	return s.BackendId
}

func (s *DescribeGroupEndpointsResponseBodyEndpoints) GetEndpointType() *string {
	return s.EndpointType
}

func (s *DescribeGroupEndpointsResponseBodyEndpoints) GetInternetEndpoints() []*string {
	return s.InternetEndpoints
}

func (s *DescribeGroupEndpointsResponseBodyEndpoints) GetIntranetEndpoints() []*string {
	return s.IntranetEndpoints
}

func (s *DescribeGroupEndpointsResponseBodyEndpoints) GetPathType() *string {
	return s.PathType
}

func (s *DescribeGroupEndpointsResponseBodyEndpoints) GetPort() *int32 {
	return s.Port
}

func (s *DescribeGroupEndpointsResponseBodyEndpoints) SetBackendId(v string) *DescribeGroupEndpointsResponseBodyEndpoints {
	s.BackendId = &v
	return s
}

func (s *DescribeGroupEndpointsResponseBodyEndpoints) SetEndpointType(v string) *DescribeGroupEndpointsResponseBodyEndpoints {
	s.EndpointType = &v
	return s
}

func (s *DescribeGroupEndpointsResponseBodyEndpoints) SetInternetEndpoints(v []*string) *DescribeGroupEndpointsResponseBodyEndpoints {
	s.InternetEndpoints = v
	return s
}

func (s *DescribeGroupEndpointsResponseBodyEndpoints) SetIntranetEndpoints(v []*string) *DescribeGroupEndpointsResponseBodyEndpoints {
	s.IntranetEndpoints = v
	return s
}

func (s *DescribeGroupEndpointsResponseBodyEndpoints) SetPathType(v string) *DescribeGroupEndpointsResponseBodyEndpoints {
	s.PathType = &v
	return s
}

func (s *DescribeGroupEndpointsResponseBodyEndpoints) SetPort(v int32) *DescribeGroupEndpointsResponseBodyEndpoints {
	s.Port = &v
	return s
}

func (s *DescribeGroupEndpointsResponseBodyEndpoints) Validate() error {
	return dara.Validate(s)
}
