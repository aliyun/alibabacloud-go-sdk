// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHADiagnoseConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeHADiagnoseConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeHADiagnoseConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeHADiagnoseConfigResponseBody) *DescribeHADiagnoseConfigResponse
	GetBody() *DescribeHADiagnoseConfigResponseBody
}

type DescribeHADiagnoseConfigResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHADiagnoseConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHADiagnoseConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeHADiagnoseConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeHADiagnoseConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeHADiagnoseConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeHADiagnoseConfigResponse) GetBody() *DescribeHADiagnoseConfigResponseBody {
	return s.Body
}

func (s *DescribeHADiagnoseConfigResponse) SetHeaders(v map[string]*string) *DescribeHADiagnoseConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeHADiagnoseConfigResponse) SetStatusCode(v int32) *DescribeHADiagnoseConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHADiagnoseConfigResponse) SetBody(v *DescribeHADiagnoseConfigResponseBody) *DescribeHADiagnoseConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeHADiagnoseConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
