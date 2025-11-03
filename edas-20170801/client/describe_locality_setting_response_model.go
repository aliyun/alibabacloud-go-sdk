// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLocalitySettingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLocalitySettingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLocalitySettingResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLocalitySettingResponseBody) *DescribeLocalitySettingResponse
	GetBody() *DescribeLocalitySettingResponseBody
}

type DescribeLocalitySettingResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLocalitySettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLocalitySettingResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLocalitySettingResponse) GoString() string {
	return s.String()
}

func (s *DescribeLocalitySettingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLocalitySettingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLocalitySettingResponse) GetBody() *DescribeLocalitySettingResponseBody {
	return s.Body
}

func (s *DescribeLocalitySettingResponse) SetHeaders(v map[string]*string) *DescribeLocalitySettingResponse {
	s.Headers = v
	return s
}

func (s *DescribeLocalitySettingResponse) SetStatusCode(v int32) *DescribeLocalitySettingResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLocalitySettingResponse) SetBody(v *DescribeLocalitySettingResponseBody) *DescribeLocalitySettingResponse {
	s.Body = v
	return s
}

func (s *DescribeLocalitySettingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
