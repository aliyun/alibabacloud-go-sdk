// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeListAntCloudAuthScenesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeListAntCloudAuthScenesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeListAntCloudAuthScenesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeListAntCloudAuthScenesResponseBody) *DescribeListAntCloudAuthScenesResponse
	GetBody() *DescribeListAntCloudAuthScenesResponseBody
}

type DescribeListAntCloudAuthScenesResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeListAntCloudAuthScenesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeListAntCloudAuthScenesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeListAntCloudAuthScenesResponse) GoString() string {
	return s.String()
}

func (s *DescribeListAntCloudAuthScenesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeListAntCloudAuthScenesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeListAntCloudAuthScenesResponse) GetBody() *DescribeListAntCloudAuthScenesResponseBody {
	return s.Body
}

func (s *DescribeListAntCloudAuthScenesResponse) SetHeaders(v map[string]*string) *DescribeListAntCloudAuthScenesResponse {
	s.Headers = v
	return s
}

func (s *DescribeListAntCloudAuthScenesResponse) SetStatusCode(v int32) *DescribeListAntCloudAuthScenesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeListAntCloudAuthScenesResponse) SetBody(v *DescribeListAntCloudAuthScenesResponseBody) *DescribeListAntCloudAuthScenesResponse {
	s.Body = v
	return s
}

func (s *DescribeListAntCloudAuthScenesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
