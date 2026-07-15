// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageRemovalProResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ImageRemovalProResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ImageRemovalProResponse
	GetStatusCode() *int32
	SetBody(v *ImageRemovalProResponseBody) *ImageRemovalProResponse
	GetBody() *ImageRemovalProResponseBody
}

type ImageRemovalProResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImageRemovalProResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImageRemovalProResponse) String() string {
	return dara.Prettify(s)
}

func (s ImageRemovalProResponse) GoString() string {
	return s.String()
}

func (s *ImageRemovalProResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ImageRemovalProResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImageRemovalProResponse) GetBody() *ImageRemovalProResponseBody {
	return s.Body
}

func (s *ImageRemovalProResponse) SetHeaders(v map[string]*string) *ImageRemovalProResponse {
	s.Headers = v
	return s
}

func (s *ImageRemovalProResponse) SetStatusCode(v int32) *ImageRemovalProResponse {
	s.StatusCode = &v
	return s
}

func (s *ImageRemovalProResponse) SetBody(v *ImageRemovalProResponseBody) *ImageRemovalProResponse {
	s.Body = v
	return s
}

func (s *ImageRemovalProResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
