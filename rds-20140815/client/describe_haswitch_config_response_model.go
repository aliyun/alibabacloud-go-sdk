// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHASwitchConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeHASwitchConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeHASwitchConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeHASwitchConfigResponseBody) *DescribeHASwitchConfigResponse
	GetBody() *DescribeHASwitchConfigResponseBody
}

type DescribeHASwitchConfigResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHASwitchConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHASwitchConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeHASwitchConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeHASwitchConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeHASwitchConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeHASwitchConfigResponse) GetBody() *DescribeHASwitchConfigResponseBody {
	return s.Body
}

func (s *DescribeHASwitchConfigResponse) SetHeaders(v map[string]*string) *DescribeHASwitchConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeHASwitchConfigResponse) SetStatusCode(v int32) *DescribeHASwitchConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHASwitchConfigResponse) SetBody(v *DescribeHASwitchConfigResponseBody) *DescribeHASwitchConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeHASwitchConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
