// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectImageScoreResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetectImageScoreResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetectImageScoreResponse
	GetStatusCode() *int32
	SetBody(v *DetectImageScoreResponseBody) *DetectImageScoreResponse
	GetBody() *DetectImageScoreResponseBody
}

type DetectImageScoreResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetectImageScoreResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetectImageScoreResponse) String() string {
	return dara.Prettify(s)
}

func (s DetectImageScoreResponse) GoString() string {
	return s.String()
}

func (s *DetectImageScoreResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetectImageScoreResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetectImageScoreResponse) GetBody() *DetectImageScoreResponseBody {
	return s.Body
}

func (s *DetectImageScoreResponse) SetHeaders(v map[string]*string) *DetectImageScoreResponse {
	s.Headers = v
	return s
}

func (s *DetectImageScoreResponse) SetStatusCode(v int32) *DetectImageScoreResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectImageScoreResponse) SetBody(v *DetectImageScoreResponseBody) *DetectImageScoreResponse {
	s.Body = v
	return s
}

func (s *DetectImageScoreResponse) Validate() error {
	return dara.Validate(s)
}
