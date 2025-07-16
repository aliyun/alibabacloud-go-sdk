// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceEndpointsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *DescribeServiceEndpointsResponseBody
	GetAccessToken() *string
	SetEndpoints(v []*DescribeServiceEndpointsResponseBodyEndpoints) *DescribeServiceEndpointsResponseBody
	GetEndpoints() []*DescribeServiceEndpointsResponseBodyEndpoints
	SetMessage(v string) *DescribeServiceEndpointsResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeServiceEndpointsResponseBody
	GetRequestId() *string
}

type DescribeServiceEndpointsResponseBody struct {
	// The service token.
	//
	// example:
	//
	// Nzc5N2FhN****TQ0YzBmYTIyN2MxZTUxN2NkYjg4MTJmMWQxZmY1****
	AccessToken *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	// The service endpoints.
	Endpoints []*DescribeServiceEndpointsResponseBodyEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Repeated"`
	// The returned message.
	//
	// example:
	//
	// Execution successful.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 739998B5-FB39-12A3-8323-0FA340317298
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeServiceEndpointsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceEndpointsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeServiceEndpointsResponseBody) GetAccessToken() *string {
	return s.AccessToken
}

func (s *DescribeServiceEndpointsResponseBody) GetEndpoints() []*DescribeServiceEndpointsResponseBodyEndpoints {
	return s.Endpoints
}

func (s *DescribeServiceEndpointsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeServiceEndpointsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeServiceEndpointsResponseBody) SetAccessToken(v string) *DescribeServiceEndpointsResponseBody {
	s.AccessToken = &v
	return s
}

func (s *DescribeServiceEndpointsResponseBody) SetEndpoints(v []*DescribeServiceEndpointsResponseBodyEndpoints) *DescribeServiceEndpointsResponseBody {
	s.Endpoints = v
	return s
}

func (s *DescribeServiceEndpointsResponseBody) SetMessage(v string) *DescribeServiceEndpointsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeServiceEndpointsResponseBody) SetRequestId(v string) *DescribeServiceEndpointsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeServiceEndpointsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeServiceEndpointsResponseBodyEndpoints struct {
	BackendId         *string   `json:"BackendId,omitempty" xml:"BackendId,omitempty"`
	EndpointType      *string   `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	InternetEndpoints []*string `json:"InternetEndpoints,omitempty" xml:"InternetEndpoints,omitempty" type:"Repeated"`
	IntranetEndpoints []*string `json:"IntranetEndpoints,omitempty" xml:"IntranetEndpoints,omitempty" type:"Repeated"`
	PathType          *string   `json:"PathType,omitempty" xml:"PathType,omitempty"`
	Port              *int32    `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s DescribeServiceEndpointsResponseBodyEndpoints) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceEndpointsResponseBodyEndpoints) GoString() string {
	return s.String()
}

func (s *DescribeServiceEndpointsResponseBodyEndpoints) GetBackendId() *string {
	return s.BackendId
}

func (s *DescribeServiceEndpointsResponseBodyEndpoints) GetEndpointType() *string {
	return s.EndpointType
}

func (s *DescribeServiceEndpointsResponseBodyEndpoints) GetInternetEndpoints() []*string {
	return s.InternetEndpoints
}

func (s *DescribeServiceEndpointsResponseBodyEndpoints) GetIntranetEndpoints() []*string {
	return s.IntranetEndpoints
}

func (s *DescribeServiceEndpointsResponseBodyEndpoints) GetPathType() *string {
	return s.PathType
}

func (s *DescribeServiceEndpointsResponseBodyEndpoints) GetPort() *int32 {
	return s.Port
}

func (s *DescribeServiceEndpointsResponseBodyEndpoints) SetBackendId(v string) *DescribeServiceEndpointsResponseBodyEndpoints {
	s.BackendId = &v
	return s
}

func (s *DescribeServiceEndpointsResponseBodyEndpoints) SetEndpointType(v string) *DescribeServiceEndpointsResponseBodyEndpoints {
	s.EndpointType = &v
	return s
}

func (s *DescribeServiceEndpointsResponseBodyEndpoints) SetInternetEndpoints(v []*string) *DescribeServiceEndpointsResponseBodyEndpoints {
	s.InternetEndpoints = v
	return s
}

func (s *DescribeServiceEndpointsResponseBodyEndpoints) SetIntranetEndpoints(v []*string) *DescribeServiceEndpointsResponseBodyEndpoints {
	s.IntranetEndpoints = v
	return s
}

func (s *DescribeServiceEndpointsResponseBodyEndpoints) SetPathType(v string) *DescribeServiceEndpointsResponseBodyEndpoints {
	s.PathType = &v
	return s
}

func (s *DescribeServiceEndpointsResponseBodyEndpoints) SetPort(v int32) *DescribeServiceEndpointsResponseBodyEndpoints {
	s.Port = &v
	return s
}

func (s *DescribeServiceEndpointsResponseBodyEndpoints) Validate() error {
	return dara.Validate(s)
}
