// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAlertSceneByEventResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAlertSceneByEventResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAlertSceneByEventResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAlertSceneByEventResponseBody) *DescribeAlertSceneByEventResponse
	GetBody() *DescribeAlertSceneByEventResponseBody
}

type DescribeAlertSceneByEventResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAlertSceneByEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAlertSceneByEventResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertSceneByEventResponse) GoString() string {
	return s.String()
}

func (s *DescribeAlertSceneByEventResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAlertSceneByEventResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAlertSceneByEventResponse) GetBody() *DescribeAlertSceneByEventResponseBody {
	return s.Body
}

func (s *DescribeAlertSceneByEventResponse) SetHeaders(v map[string]*string) *DescribeAlertSceneByEventResponse {
	s.Headers = v
	return s
}

func (s *DescribeAlertSceneByEventResponse) SetStatusCode(v int32) *DescribeAlertSceneByEventResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAlertSceneByEventResponse) SetBody(v *DescribeAlertSceneByEventResponseBody) *DescribeAlertSceneByEventResponse {
	s.Body = v
	return s
}

func (s *DescribeAlertSceneByEventResponse) Validate() error {
	return dara.Validate(s)
}
