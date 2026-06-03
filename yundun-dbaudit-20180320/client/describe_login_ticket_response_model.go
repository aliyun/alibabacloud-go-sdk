// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLoginTicketResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLoginTicketResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLoginTicketResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLoginTicketResponseBody) *DescribeLoginTicketResponse
	GetBody() *DescribeLoginTicketResponseBody
}

type DescribeLoginTicketResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLoginTicketResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLoginTicketResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoginTicketResponse) GoString() string {
	return s.String()
}

func (s *DescribeLoginTicketResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLoginTicketResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLoginTicketResponse) GetBody() *DescribeLoginTicketResponseBody {
	return s.Body
}

func (s *DescribeLoginTicketResponse) SetHeaders(v map[string]*string) *DescribeLoginTicketResponse {
	s.Headers = v
	return s
}

func (s *DescribeLoginTicketResponse) SetStatusCode(v int32) *DescribeLoginTicketResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLoginTicketResponse) SetBody(v *DescribeLoginTicketResponseBody) *DescribeLoginTicketResponse {
	s.Body = v
	return s
}

func (s *DescribeLoginTicketResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
