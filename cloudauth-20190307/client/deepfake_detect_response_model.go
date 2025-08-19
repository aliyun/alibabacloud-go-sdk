// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeepfakeDetectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeepfakeDetectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeepfakeDetectResponse
	GetStatusCode() *int32
	SetBody(v *DeepfakeDetectResponseBody) *DeepfakeDetectResponse
	GetBody() *DeepfakeDetectResponseBody
}

type DeepfakeDetectResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeepfakeDetectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeepfakeDetectResponse) String() string {
	return dara.Prettify(s)
}

func (s DeepfakeDetectResponse) GoString() string {
	return s.String()
}

func (s *DeepfakeDetectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeepfakeDetectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeepfakeDetectResponse) GetBody() *DeepfakeDetectResponseBody {
	return s.Body
}

func (s *DeepfakeDetectResponse) SetHeaders(v map[string]*string) *DeepfakeDetectResponse {
	s.Headers = v
	return s
}

func (s *DeepfakeDetectResponse) SetStatusCode(v int32) *DeepfakeDetectResponse {
	s.StatusCode = &v
	return s
}

func (s *DeepfakeDetectResponse) SetBody(v *DeepfakeDetectResponseBody) *DeepfakeDetectResponse {
	s.Body = v
	return s
}

func (s *DeepfakeDetectResponse) Validate() error {
	return dara.Validate(s)
}
