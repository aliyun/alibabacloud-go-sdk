// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePerspectiveResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePerspectiveResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePerspectiveResponse
	GetStatusCode() *int32
	SetBody(v *CreatePerspectiveResponseBody) *CreatePerspectiveResponse
	GetBody() *CreatePerspectiveResponseBody
}

type CreatePerspectiveResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePerspectiveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePerspectiveResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePerspectiveResponse) GoString() string {
	return s.String()
}

func (s *CreatePerspectiveResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePerspectiveResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePerspectiveResponse) GetBody() *CreatePerspectiveResponseBody {
	return s.Body
}

func (s *CreatePerspectiveResponse) SetHeaders(v map[string]*string) *CreatePerspectiveResponse {
	s.Headers = v
	return s
}

func (s *CreatePerspectiveResponse) SetStatusCode(v int32) *CreatePerspectiveResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePerspectiveResponse) SetBody(v *CreatePerspectiveResponseBody) *CreatePerspectiveResponse {
	s.Body = v
	return s
}

func (s *CreatePerspectiveResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
