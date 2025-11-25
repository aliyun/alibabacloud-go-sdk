// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAutoCcBlacklistResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAutoCcBlacklistResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAutoCcBlacklistResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAutoCcBlacklistResponseBody) *DescribeAutoCcBlacklistResponse
	GetBody() *DescribeAutoCcBlacklistResponseBody
}

type DescribeAutoCcBlacklistResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAutoCcBlacklistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAutoCcBlacklistResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoCcBlacklistResponse) GoString() string {
	return s.String()
}

func (s *DescribeAutoCcBlacklistResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAutoCcBlacklistResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAutoCcBlacklistResponse) GetBody() *DescribeAutoCcBlacklistResponseBody {
	return s.Body
}

func (s *DescribeAutoCcBlacklistResponse) SetHeaders(v map[string]*string) *DescribeAutoCcBlacklistResponse {
	s.Headers = v
	return s
}

func (s *DescribeAutoCcBlacklistResponse) SetStatusCode(v int32) *DescribeAutoCcBlacklistResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAutoCcBlacklistResponse) SetBody(v *DescribeAutoCcBlacklistResponseBody) *DescribeAutoCcBlacklistResponse {
	s.Body = v
	return s
}

func (s *DescribeAutoCcBlacklistResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
