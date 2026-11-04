// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePictureSearchJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePictureSearchJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePictureSearchJobResponse
	GetStatusCode() *int32
	SetBody(v *CreatePictureSearchJobResponseBody) *CreatePictureSearchJobResponse
	GetBody() *CreatePictureSearchJobResponseBody
}

type CreatePictureSearchJobResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePictureSearchJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePictureSearchJobResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePictureSearchJobResponse) GoString() string {
	return s.String()
}

func (s *CreatePictureSearchJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePictureSearchJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePictureSearchJobResponse) GetBody() *CreatePictureSearchJobResponseBody {
	return s.Body
}

func (s *CreatePictureSearchJobResponse) SetHeaders(v map[string]*string) *CreatePictureSearchJobResponse {
	s.Headers = v
	return s
}

func (s *CreatePictureSearchJobResponse) SetStatusCode(v int32) *CreatePictureSearchJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePictureSearchJobResponse) SetBody(v *CreatePictureSearchJobResponseBody) *CreatePictureSearchJobResponse {
	s.Body = v
	return s
}

func (s *CreatePictureSearchJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
