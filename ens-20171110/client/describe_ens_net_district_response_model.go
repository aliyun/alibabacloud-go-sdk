// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEnsNetDistrictResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEnsNetDistrictResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEnsNetDistrictResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEnsNetDistrictResponseBody) *DescribeEnsNetDistrictResponse
	GetBody() *DescribeEnsNetDistrictResponseBody
}

type DescribeEnsNetDistrictResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEnsNetDistrictResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEnsNetDistrictResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsNetDistrictResponse) GoString() string {
	return s.String()
}

func (s *DescribeEnsNetDistrictResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEnsNetDistrictResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEnsNetDistrictResponse) GetBody() *DescribeEnsNetDistrictResponseBody {
	return s.Body
}

func (s *DescribeEnsNetDistrictResponse) SetHeaders(v map[string]*string) *DescribeEnsNetDistrictResponse {
	s.Headers = v
	return s
}

func (s *DescribeEnsNetDistrictResponse) SetStatusCode(v int32) *DescribeEnsNetDistrictResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEnsNetDistrictResponse) SetBody(v *DescribeEnsNetDistrictResponseBody) *DescribeEnsNetDistrictResponse {
	s.Body = v
	return s
}

func (s *DescribeEnsNetDistrictResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
