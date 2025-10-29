// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveGrtnDurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveGrtnDurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveGrtnDurationResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveGrtnDurationResponseBody) *DescribeLiveGrtnDurationResponse
	GetBody() *DescribeLiveGrtnDurationResponseBody
}

type DescribeLiveGrtnDurationResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveGrtnDurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveGrtnDurationResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveGrtnDurationResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveGrtnDurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveGrtnDurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveGrtnDurationResponse) GetBody() *DescribeLiveGrtnDurationResponseBody {
	return s.Body
}

func (s *DescribeLiveGrtnDurationResponse) SetHeaders(v map[string]*string) *DescribeLiveGrtnDurationResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveGrtnDurationResponse) SetStatusCode(v int32) *DescribeLiveGrtnDurationResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveGrtnDurationResponse) SetBody(v *DescribeLiveGrtnDurationResponseBody) *DescribeLiveGrtnDurationResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveGrtnDurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
