// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveMessageAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveMessageAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveMessageAppResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveMessageAppResponseBody) *DescribeLiveMessageAppResponse
	GetBody() *DescribeLiveMessageAppResponseBody
}

type DescribeLiveMessageAppResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveMessageAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveMessageAppResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveMessageAppResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveMessageAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveMessageAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveMessageAppResponse) GetBody() *DescribeLiveMessageAppResponseBody {
	return s.Body
}

func (s *DescribeLiveMessageAppResponse) SetHeaders(v map[string]*string) *DescribeLiveMessageAppResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveMessageAppResponse) SetStatusCode(v int32) *DescribeLiveMessageAppResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveMessageAppResponse) SetBody(v *DescribeLiveMessageAppResponseBody) *DescribeLiveMessageAppResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveMessageAppResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
