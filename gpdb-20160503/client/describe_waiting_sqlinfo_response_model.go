// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWaitingSQLInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeWaitingSQLInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeWaitingSQLInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeWaitingSQLInfoResponseBody) *DescribeWaitingSQLInfoResponse
	GetBody() *DescribeWaitingSQLInfoResponseBody
}

type DescribeWaitingSQLInfoResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeWaitingSQLInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeWaitingSQLInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeWaitingSQLInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeWaitingSQLInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeWaitingSQLInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeWaitingSQLInfoResponse) GetBody() *DescribeWaitingSQLInfoResponseBody {
	return s.Body
}

func (s *DescribeWaitingSQLInfoResponse) SetHeaders(v map[string]*string) *DescribeWaitingSQLInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeWaitingSQLInfoResponse) SetStatusCode(v int32) *DescribeWaitingSQLInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWaitingSQLInfoResponse) SetBody(v *DescribeWaitingSQLInfoResponseBody) *DescribeWaitingSQLInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeWaitingSQLInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
