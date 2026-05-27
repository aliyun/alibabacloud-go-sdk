// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWuyingServerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeWuyingServerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeWuyingServerResponse
	GetStatusCode() *int32
	SetBody(v *DescribeWuyingServerResponseBody) *DescribeWuyingServerResponse
	GetBody() *DescribeWuyingServerResponseBody
}

type DescribeWuyingServerResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeWuyingServerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeWuyingServerResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeWuyingServerResponse) GoString() string {
	return s.String()
}

func (s *DescribeWuyingServerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeWuyingServerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeWuyingServerResponse) GetBody() *DescribeWuyingServerResponseBody {
	return s.Body
}

func (s *DescribeWuyingServerResponse) SetHeaders(v map[string]*string) *DescribeWuyingServerResponse {
	s.Headers = v
	return s
}

func (s *DescribeWuyingServerResponse) SetStatusCode(v int32) *DescribeWuyingServerResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWuyingServerResponse) SetBody(v *DescribeWuyingServerResponseBody) *DescribeWuyingServerResponse {
	s.Body = v
	return s
}

func (s *DescribeWuyingServerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
