// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDropProjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DropProjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DropProjectResponse
	GetStatusCode() *int32
	SetBody(v *DropProjectResponseBody) *DropProjectResponse
	GetBody() *DropProjectResponseBody
}

type DropProjectResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DropProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DropProjectResponse) String() string {
	return dara.Prettify(s)
}

func (s DropProjectResponse) GoString() string {
	return s.String()
}

func (s *DropProjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DropProjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DropProjectResponse) GetBody() *DropProjectResponseBody {
	return s.Body
}

func (s *DropProjectResponse) SetHeaders(v map[string]*string) *DropProjectResponse {
	s.Headers = v
	return s
}

func (s *DropProjectResponse) SetStatusCode(v int32) *DropProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *DropProjectResponse) SetBody(v *DropProjectResponseBody) *DropProjectResponse {
	s.Body = v
	return s
}

func (s *DropProjectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
