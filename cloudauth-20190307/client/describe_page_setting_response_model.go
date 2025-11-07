// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePageSettingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePageSettingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePageSettingResponse
	GetStatusCode() *int32
	SetBody(v *DescribePageSettingResponseBody) *DescribePageSettingResponse
	GetBody() *DescribePageSettingResponseBody
}

type DescribePageSettingResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePageSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePageSettingResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePageSettingResponse) GoString() string {
	return s.String()
}

func (s *DescribePageSettingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePageSettingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePageSettingResponse) GetBody() *DescribePageSettingResponseBody {
	return s.Body
}

func (s *DescribePageSettingResponse) SetHeaders(v map[string]*string) *DescribePageSettingResponse {
	s.Headers = v
	return s
}

func (s *DescribePageSettingResponse) SetStatusCode(v int32) *DescribePageSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePageSettingResponse) SetBody(v *DescribePageSettingResponseBody) *DescribePageSettingResponse {
	s.Body = v
	return s
}

func (s *DescribePageSettingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
