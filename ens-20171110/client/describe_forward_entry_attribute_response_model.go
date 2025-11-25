// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeForwardEntryAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeForwardEntryAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeForwardEntryAttributeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeForwardEntryAttributeResponseBody) *DescribeForwardEntryAttributeResponse
	GetBody() *DescribeForwardEntryAttributeResponseBody
}

type DescribeForwardEntryAttributeResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeForwardEntryAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeForwardEntryAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeForwardEntryAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeForwardEntryAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeForwardEntryAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeForwardEntryAttributeResponse) GetBody() *DescribeForwardEntryAttributeResponseBody {
	return s.Body
}

func (s *DescribeForwardEntryAttributeResponse) SetHeaders(v map[string]*string) *DescribeForwardEntryAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeForwardEntryAttributeResponse) SetStatusCode(v int32) *DescribeForwardEntryAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeForwardEntryAttributeResponse) SetBody(v *DescribeForwardEntryAttributeResponseBody) *DescribeForwardEntryAttributeResponse {
	s.Body = v
	return s
}

func (s *DescribeForwardEntryAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
