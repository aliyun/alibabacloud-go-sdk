// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectImageLabelsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetectImageLabelsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetectImageLabelsResponse
	GetStatusCode() *int32
	SetBody(v *DetectImageLabelsResponseBody) *DetectImageLabelsResponse
	GetBody() *DetectImageLabelsResponseBody
}

type DetectImageLabelsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetectImageLabelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetectImageLabelsResponse) String() string {
	return dara.Prettify(s)
}

func (s DetectImageLabelsResponse) GoString() string {
	return s.String()
}

func (s *DetectImageLabelsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetectImageLabelsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetectImageLabelsResponse) GetBody() *DetectImageLabelsResponseBody {
	return s.Body
}

func (s *DetectImageLabelsResponse) SetHeaders(v map[string]*string) *DetectImageLabelsResponse {
	s.Headers = v
	return s
}

func (s *DetectImageLabelsResponse) SetStatusCode(v int32) *DetectImageLabelsResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectImageLabelsResponse) SetBody(v *DetectImageLabelsResponseBody) *DetectImageLabelsResponse {
	s.Body = v
	return s
}

func (s *DetectImageLabelsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
