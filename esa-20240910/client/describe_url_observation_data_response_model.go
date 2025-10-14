// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUrlObservationDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUrlObservationDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUrlObservationDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUrlObservationDataResponseBody) *DescribeUrlObservationDataResponse
	GetBody() *DescribeUrlObservationDataResponseBody
}

type DescribeUrlObservationDataResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUrlObservationDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUrlObservationDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUrlObservationDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeUrlObservationDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUrlObservationDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUrlObservationDataResponse) GetBody() *DescribeUrlObservationDataResponseBody {
	return s.Body
}

func (s *DescribeUrlObservationDataResponse) SetHeaders(v map[string]*string) *DescribeUrlObservationDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeUrlObservationDataResponse) SetStatusCode(v int32) *DescribeUrlObservationDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUrlObservationDataResponse) SetBody(v *DescribeUrlObservationDataResponseBody) *DescribeUrlObservationDataResponse {
	s.Body = v
	return s
}

func (s *DescribeUrlObservationDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
