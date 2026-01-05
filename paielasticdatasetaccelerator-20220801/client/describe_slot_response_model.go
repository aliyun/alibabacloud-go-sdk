// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlotResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSlotResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSlotResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSlotResponseBody) *DescribeSlotResponse
	GetBody() *DescribeSlotResponseBody
}

type DescribeSlotResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSlotResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSlotResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlotResponse) GoString() string {
	return s.String()
}

func (s *DescribeSlotResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSlotResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSlotResponse) GetBody() *DescribeSlotResponseBody {
	return s.Body
}

func (s *DescribeSlotResponse) SetHeaders(v map[string]*string) *DescribeSlotResponse {
	s.Headers = v
	return s
}

func (s *DescribeSlotResponse) SetStatusCode(v int32) *DescribeSlotResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSlotResponse) SetBody(v *DescribeSlotResponseBody) *DescribeSlotResponse {
	s.Body = v
	return s
}

func (s *DescribeSlotResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
