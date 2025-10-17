// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeepfakeDetectIntlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeepfakeDetectIntlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeepfakeDetectIntlResponse
	GetStatusCode() *int32
	SetBody(v *DeepfakeDetectIntlResponseBody) *DeepfakeDetectIntlResponse
	GetBody() *DeepfakeDetectIntlResponseBody
}

type DeepfakeDetectIntlResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeepfakeDetectIntlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeepfakeDetectIntlResponse) String() string {
	return dara.Prettify(s)
}

func (s DeepfakeDetectIntlResponse) GoString() string {
	return s.String()
}

func (s *DeepfakeDetectIntlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeepfakeDetectIntlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeepfakeDetectIntlResponse) GetBody() *DeepfakeDetectIntlResponseBody {
	return s.Body
}

func (s *DeepfakeDetectIntlResponse) SetHeaders(v map[string]*string) *DeepfakeDetectIntlResponse {
	s.Headers = v
	return s
}

func (s *DeepfakeDetectIntlResponse) SetStatusCode(v int32) *DeepfakeDetectIntlResponse {
	s.StatusCode = &v
	return s
}

func (s *DeepfakeDetectIntlResponse) SetBody(v *DeepfakeDetectIntlResponseBody) *DeepfakeDetectIntlResponse {
	s.Body = v
	return s
}

func (s *DeepfakeDetectIntlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
