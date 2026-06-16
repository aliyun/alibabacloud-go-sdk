// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageRemoveResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ImageRemoveResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ImageRemoveResponse
	GetStatusCode() *int32
	SetBody(v *ImageRemoveResponseBody) *ImageRemoveResponse
	GetBody() *ImageRemoveResponseBody
}

type ImageRemoveResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImageRemoveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImageRemoveResponse) String() string {
	return dara.Prettify(s)
}

func (s ImageRemoveResponse) GoString() string {
	return s.String()
}

func (s *ImageRemoveResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ImageRemoveResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImageRemoveResponse) GetBody() *ImageRemoveResponseBody {
	return s.Body
}

func (s *ImageRemoveResponse) SetHeaders(v map[string]*string) *ImageRemoveResponse {
	s.Headers = v
	return s
}

func (s *ImageRemoveResponse) SetStatusCode(v int32) *ImageRemoveResponse {
	s.StatusCode = &v
	return s
}

func (s *ImageRemoveResponse) SetBody(v *ImageRemoveResponseBody) *ImageRemoveResponse {
	s.Body = v
	return s
}

func (s *ImageRemoveResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
