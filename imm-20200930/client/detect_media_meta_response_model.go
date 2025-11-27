// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectMediaMetaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetectMediaMetaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetectMediaMetaResponse
	GetStatusCode() *int32
	SetBody(v *DetectMediaMetaResponseBody) *DetectMediaMetaResponse
	GetBody() *DetectMediaMetaResponseBody
}

type DetectMediaMetaResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetectMediaMetaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetectMediaMetaResponse) String() string {
	return dara.Prettify(s)
}

func (s DetectMediaMetaResponse) GoString() string {
	return s.String()
}

func (s *DetectMediaMetaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetectMediaMetaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetectMediaMetaResponse) GetBody() *DetectMediaMetaResponseBody {
	return s.Body
}

func (s *DetectMediaMetaResponse) SetHeaders(v map[string]*string) *DetectMediaMetaResponse {
	s.Headers = v
	return s
}

func (s *DetectMediaMetaResponse) SetStatusCode(v int32) *DetectMediaMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectMediaMetaResponse) SetBody(v *DetectMediaMetaResponseBody) *DetectMediaMetaResponse {
	s.Body = v
	return s
}

func (s *DetectMediaMetaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
