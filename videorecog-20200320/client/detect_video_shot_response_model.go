// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectVideoShotResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetectVideoShotResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetectVideoShotResponse
	GetStatusCode() *int32
	SetBody(v *DetectVideoShotResponseBody) *DetectVideoShotResponse
	GetBody() *DetectVideoShotResponseBody
}

type DetectVideoShotResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetectVideoShotResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetectVideoShotResponse) String() string {
	return dara.Prettify(s)
}

func (s DetectVideoShotResponse) GoString() string {
	return s.String()
}

func (s *DetectVideoShotResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetectVideoShotResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetectVideoShotResponse) GetBody() *DetectVideoShotResponseBody {
	return s.Body
}

func (s *DetectVideoShotResponse) SetHeaders(v map[string]*string) *DetectVideoShotResponse {
	s.Headers = v
	return s
}

func (s *DetectVideoShotResponse) SetStatusCode(v int32) *DetectVideoShotResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectVideoShotResponse) SetBody(v *DetectVideoShotResponseBody) *DetectVideoShotResponse {
	s.Body = v
	return s
}

func (s *DetectVideoShotResponse) Validate() error {
	return dara.Validate(s)
}
