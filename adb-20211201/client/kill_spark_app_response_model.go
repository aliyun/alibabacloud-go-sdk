// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKillSparkAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *KillSparkAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *KillSparkAppResponse
	GetStatusCode() *int32
	SetBody(v *KillSparkAppResponseBody) *KillSparkAppResponse
	GetBody() *KillSparkAppResponseBody
}

type KillSparkAppResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *KillSparkAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s KillSparkAppResponse) String() string {
	return dara.Prettify(s)
}

func (s KillSparkAppResponse) GoString() string {
	return s.String()
}

func (s *KillSparkAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *KillSparkAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *KillSparkAppResponse) GetBody() *KillSparkAppResponseBody {
	return s.Body
}

func (s *KillSparkAppResponse) SetHeaders(v map[string]*string) *KillSparkAppResponse {
	s.Headers = v
	return s
}

func (s *KillSparkAppResponse) SetStatusCode(v int32) *KillSparkAppResponse {
	s.StatusCode = &v
	return s
}

func (s *KillSparkAppResponse) SetBody(v *KillSparkAppResponseBody) *KillSparkAppResponse {
	s.Body = v
	return s
}

func (s *KillSparkAppResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
