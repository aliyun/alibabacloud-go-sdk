// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompleteModelingProjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CompleteModelingProjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CompleteModelingProjectResponse
	GetStatusCode() *int32
	SetBody(v *CompleteModelingProjectResponseBody) *CompleteModelingProjectResponse
	GetBody() *CompleteModelingProjectResponseBody
}

type CompleteModelingProjectResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CompleteModelingProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CompleteModelingProjectResponse) String() string {
	return dara.Prettify(s)
}

func (s CompleteModelingProjectResponse) GoString() string {
	return s.String()
}

func (s *CompleteModelingProjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CompleteModelingProjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CompleteModelingProjectResponse) GetBody() *CompleteModelingProjectResponseBody {
	return s.Body
}

func (s *CompleteModelingProjectResponse) SetHeaders(v map[string]*string) *CompleteModelingProjectResponse {
	s.Headers = v
	return s
}

func (s *CompleteModelingProjectResponse) SetStatusCode(v int32) *CompleteModelingProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *CompleteModelingProjectResponse) SetBody(v *CompleteModelingProjectResponseBody) *CompleteModelingProjectResponse {
	s.Body = v
	return s
}

func (s *CompleteModelingProjectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
