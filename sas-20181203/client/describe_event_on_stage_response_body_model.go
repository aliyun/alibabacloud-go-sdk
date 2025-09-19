// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventOnStageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeEventOnStageResponseBody
	GetRequestId() *string
	SetSecurityEventStageResponse(v *DescribeEventOnStageResponseBodySecurityEventStageResponse) *DescribeEventOnStageResponseBody
	GetSecurityEventStageResponse() *DescribeEventOnStageResponseBodySecurityEventStageResponse
}

type DescribeEventOnStageResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 24A20733-10A0-4AF6-BE6B-E332241XXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The platforms that are supported by the feature of container threat detection.
	SecurityEventStageResponse *DescribeEventOnStageResponseBodySecurityEventStageResponse `json:"SecurityEventStageResponse,omitempty" xml:"SecurityEventStageResponse,omitempty" type:"Struct"`
}

func (s DescribeEventOnStageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventOnStageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEventOnStageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEventOnStageResponseBody) GetSecurityEventStageResponse() *DescribeEventOnStageResponseBodySecurityEventStageResponse {
	return s.SecurityEventStageResponse
}

func (s *DescribeEventOnStageResponseBody) SetRequestId(v string) *DescribeEventOnStageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEventOnStageResponseBody) SetSecurityEventStageResponse(v *DescribeEventOnStageResponseBodySecurityEventStageResponse) *DescribeEventOnStageResponseBody {
	s.SecurityEventStageResponse = v
	return s
}

func (s *DescribeEventOnStageResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeEventOnStageResponseBodySecurityEventStageResponse struct {
	// The platform that is supported by the feature of container threat detection. Valid values:
	//
	// 	- **container**
	//
	// 	- **linux**
	//
	// 	- **windows**
	SecurityEventOnStag map[string]interface{} `json:"SecurityEventOnStag,omitempty" xml:"SecurityEventOnStag,omitempty"`
}

func (s DescribeEventOnStageResponseBodySecurityEventStageResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventOnStageResponseBodySecurityEventStageResponse) GoString() string {
	return s.String()
}

func (s *DescribeEventOnStageResponseBodySecurityEventStageResponse) GetSecurityEventOnStag() map[string]interface{} {
	return s.SecurityEventOnStag
}

func (s *DescribeEventOnStageResponseBodySecurityEventStageResponse) SetSecurityEventOnStag(v map[string]interface{}) *DescribeEventOnStageResponseBodySecurityEventStageResponse {
	s.SecurityEventOnStag = v
	return s
}

func (s *DescribeEventOnStageResponseBodySecurityEventStageResponse) Validate() error {
	return dara.Validate(s)
}
