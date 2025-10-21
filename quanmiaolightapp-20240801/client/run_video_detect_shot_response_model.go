// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunVideoDetectShotResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunVideoDetectShotResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunVideoDetectShotResponse
	GetStatusCode() *int32
	SetBody(v *RunVideoDetectShotResponseBody) *RunVideoDetectShotResponse
	GetBody() *RunVideoDetectShotResponseBody
}

type RunVideoDetectShotResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunVideoDetectShotResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunVideoDetectShotResponse) String() string {
	return dara.Prettify(s)
}

func (s RunVideoDetectShotResponse) GoString() string {
	return s.String()
}

func (s *RunVideoDetectShotResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunVideoDetectShotResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunVideoDetectShotResponse) GetBody() *RunVideoDetectShotResponseBody {
	return s.Body
}

func (s *RunVideoDetectShotResponse) SetHeaders(v map[string]*string) *RunVideoDetectShotResponse {
	s.Headers = v
	return s
}

func (s *RunVideoDetectShotResponse) SetStatusCode(v int32) *RunVideoDetectShotResponse {
	s.StatusCode = &v
	return s
}

func (s *RunVideoDetectShotResponse) SetBody(v *RunVideoDetectShotResponseBody) *RunVideoDetectShotResponse {
	s.Body = v
	return s
}

func (s *RunVideoDetectShotResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
