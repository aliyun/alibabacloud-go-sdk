// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgenticDatabaseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAgenticDatabaseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAgenticDatabaseResponse
	GetStatusCode() *int32
	SetBody(v *GetAgenticDatabaseResponseBody) *GetAgenticDatabaseResponse
	GetBody() *GetAgenticDatabaseResponseBody
}

type GetAgenticDatabaseResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAgenticDatabaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAgenticDatabaseResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAgenticDatabaseResponse) GoString() string {
	return s.String()
}

func (s *GetAgenticDatabaseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAgenticDatabaseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAgenticDatabaseResponse) GetBody() *GetAgenticDatabaseResponseBody {
	return s.Body
}

func (s *GetAgenticDatabaseResponse) SetHeaders(v map[string]*string) *GetAgenticDatabaseResponse {
	s.Headers = v
	return s
}

func (s *GetAgenticDatabaseResponse) SetStatusCode(v int32) *GetAgenticDatabaseResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAgenticDatabaseResponse) SetBody(v *GetAgenticDatabaseResponseBody) *GetAgenticDatabaseResponse {
	s.Body = v
	return s
}

func (s *GetAgenticDatabaseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
