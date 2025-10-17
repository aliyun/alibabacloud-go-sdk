// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSparkAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitSparkAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitSparkAppResponse
	GetStatusCode() *int32
	SetBody(v *SubmitSparkAppResponseBody) *SubmitSparkAppResponse
	GetBody() *SubmitSparkAppResponseBody
}

type SubmitSparkAppResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitSparkAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitSparkAppResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitSparkAppResponse) GoString() string {
	return s.String()
}

func (s *SubmitSparkAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitSparkAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitSparkAppResponse) GetBody() *SubmitSparkAppResponseBody {
	return s.Body
}

func (s *SubmitSparkAppResponse) SetHeaders(v map[string]*string) *SubmitSparkAppResponse {
	s.Headers = v
	return s
}

func (s *SubmitSparkAppResponse) SetStatusCode(v int32) *SubmitSparkAppResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitSparkAppResponse) SetBody(v *SubmitSparkAppResponseBody) *SubmitSparkAppResponse {
	s.Body = v
	return s
}

func (s *SubmitSparkAppResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
