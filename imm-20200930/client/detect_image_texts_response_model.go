// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectImageTextsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetectImageTextsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetectImageTextsResponse
	GetStatusCode() *int32
	SetBody(v *DetectImageTextsResponseBody) *DetectImageTextsResponse
	GetBody() *DetectImageTextsResponseBody
}

type DetectImageTextsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetectImageTextsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetectImageTextsResponse) String() string {
	return dara.Prettify(s)
}

func (s DetectImageTextsResponse) GoString() string {
	return s.String()
}

func (s *DetectImageTextsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetectImageTextsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetectImageTextsResponse) GetBody() *DetectImageTextsResponseBody {
	return s.Body
}

func (s *DetectImageTextsResponse) SetHeaders(v map[string]*string) *DetectImageTextsResponse {
	s.Headers = v
	return s
}

func (s *DetectImageTextsResponse) SetStatusCode(v int32) *DetectImageTextsResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectImageTextsResponse) SetBody(v *DetectImageTextsResponseBody) *DetectImageTextsResponse {
	s.Body = v
	return s
}

func (s *DetectImageTextsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
