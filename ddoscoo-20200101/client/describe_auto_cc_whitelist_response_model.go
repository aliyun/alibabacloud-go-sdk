// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAutoCcWhitelistResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAutoCcWhitelistResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAutoCcWhitelistResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAutoCcWhitelistResponseBody) *DescribeAutoCcWhitelistResponse
	GetBody() *DescribeAutoCcWhitelistResponseBody
}

type DescribeAutoCcWhitelistResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAutoCcWhitelistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAutoCcWhitelistResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoCcWhitelistResponse) GoString() string {
	return s.String()
}

func (s *DescribeAutoCcWhitelistResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAutoCcWhitelistResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAutoCcWhitelistResponse) GetBody() *DescribeAutoCcWhitelistResponseBody {
	return s.Body
}

func (s *DescribeAutoCcWhitelistResponse) SetHeaders(v map[string]*string) *DescribeAutoCcWhitelistResponse {
	s.Headers = v
	return s
}

func (s *DescribeAutoCcWhitelistResponse) SetStatusCode(v int32) *DescribeAutoCcWhitelistResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAutoCcWhitelistResponse) SetBody(v *DescribeAutoCcWhitelistResponseBody) *DescribeAutoCcWhitelistResponse {
	s.Body = v
	return s
}

func (s *DescribeAutoCcWhitelistResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
