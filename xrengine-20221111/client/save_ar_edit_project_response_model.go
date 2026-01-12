// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveArEditProjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveArEditProjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveArEditProjectResponse
	GetStatusCode() *int32
	SetBody(v *SaveArEditProjectResponseBody) *SaveArEditProjectResponse
	GetBody() *SaveArEditProjectResponseBody
}

type SaveArEditProjectResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveArEditProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveArEditProjectResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveArEditProjectResponse) GoString() string {
	return s.String()
}

func (s *SaveArEditProjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveArEditProjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveArEditProjectResponse) GetBody() *SaveArEditProjectResponseBody {
	return s.Body
}

func (s *SaveArEditProjectResponse) SetHeaders(v map[string]*string) *SaveArEditProjectResponse {
	s.Headers = v
	return s
}

func (s *SaveArEditProjectResponse) SetStatusCode(v int32) *SaveArEditProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveArEditProjectResponse) SetBody(v *SaveArEditProjectResponseBody) *SaveArEditProjectResponse {
	s.Body = v
	return s
}

func (s *SaveArEditProjectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
