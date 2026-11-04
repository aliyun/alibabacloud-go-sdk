// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePictureSearchAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePictureSearchAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePictureSearchAppResponse
	GetStatusCode() *int32
	SetBody(v *CreatePictureSearchAppResponseBody) *CreatePictureSearchAppResponse
	GetBody() *CreatePictureSearchAppResponseBody
}

type CreatePictureSearchAppResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePictureSearchAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePictureSearchAppResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePictureSearchAppResponse) GoString() string {
	return s.String()
}

func (s *CreatePictureSearchAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePictureSearchAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePictureSearchAppResponse) GetBody() *CreatePictureSearchAppResponseBody {
	return s.Body
}

func (s *CreatePictureSearchAppResponse) SetHeaders(v map[string]*string) *CreatePictureSearchAppResponse {
	s.Headers = v
	return s
}

func (s *CreatePictureSearchAppResponse) SetStatusCode(v int32) *CreatePictureSearchAppResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePictureSearchAppResponse) SetBody(v *CreatePictureSearchAppResponseBody) *CreatePictureSearchAppResponse {
	s.Body = v
	return s
}

func (s *CreatePictureSearchAppResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
