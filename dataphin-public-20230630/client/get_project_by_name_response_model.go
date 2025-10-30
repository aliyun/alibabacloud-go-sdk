// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProjectByNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetProjectByNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetProjectByNameResponse
	GetStatusCode() *int32
	SetBody(v *GetProjectByNameResponseBody) *GetProjectByNameResponse
	GetBody() *GetProjectByNameResponseBody
}

type GetProjectByNameResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetProjectByNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetProjectByNameResponse) String() string {
	return dara.Prettify(s)
}

func (s GetProjectByNameResponse) GoString() string {
	return s.String()
}

func (s *GetProjectByNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetProjectByNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetProjectByNameResponse) GetBody() *GetProjectByNameResponseBody {
	return s.Body
}

func (s *GetProjectByNameResponse) SetHeaders(v map[string]*string) *GetProjectByNameResponse {
	s.Headers = v
	return s
}

func (s *GetProjectByNameResponse) SetStatusCode(v int32) *GetProjectByNameResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProjectByNameResponse) SetBody(v *GetProjectByNameResponseBody) *GetProjectByNameResponse {
	s.Body = v
	return s
}

func (s *GetProjectByNameResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
