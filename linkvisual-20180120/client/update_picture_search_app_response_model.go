// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePictureSearchAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePictureSearchAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePictureSearchAppResponse
	GetStatusCode() *int32
	SetBody(v *UpdatePictureSearchAppResponseBody) *UpdatePictureSearchAppResponse
	GetBody() *UpdatePictureSearchAppResponseBody
}

type UpdatePictureSearchAppResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePictureSearchAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePictureSearchAppResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePictureSearchAppResponse) GoString() string {
	return s.String()
}

func (s *UpdatePictureSearchAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePictureSearchAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePictureSearchAppResponse) GetBody() *UpdatePictureSearchAppResponseBody {
	return s.Body
}

func (s *UpdatePictureSearchAppResponse) SetHeaders(v map[string]*string) *UpdatePictureSearchAppResponse {
	s.Headers = v
	return s
}

func (s *UpdatePictureSearchAppResponse) SetStatusCode(v int32) *UpdatePictureSearchAppResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePictureSearchAppResponse) SetBody(v *UpdatePictureSearchAppResponseBody) *UpdatePictureSearchAppResponse {
	s.Body = v
	return s
}

func (s *UpdatePictureSearchAppResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
