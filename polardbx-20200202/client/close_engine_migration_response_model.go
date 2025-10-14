// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseEngineMigrationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloseEngineMigrationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloseEngineMigrationResponse
	GetStatusCode() *int32
	SetBody(v *CloseEngineMigrationResponseBody) *CloseEngineMigrationResponse
	GetBody() *CloseEngineMigrationResponseBody
}

type CloseEngineMigrationResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloseEngineMigrationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloseEngineMigrationResponse) String() string {
	return dara.Prettify(s)
}

func (s CloseEngineMigrationResponse) GoString() string {
	return s.String()
}

func (s *CloseEngineMigrationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloseEngineMigrationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloseEngineMigrationResponse) GetBody() *CloseEngineMigrationResponseBody {
	return s.Body
}

func (s *CloseEngineMigrationResponse) SetHeaders(v map[string]*string) *CloseEngineMigrationResponse {
	s.Headers = v
	return s
}

func (s *CloseEngineMigrationResponse) SetStatusCode(v int32) *CloseEngineMigrationResponse {
	s.StatusCode = &v
	return s
}

func (s *CloseEngineMigrationResponse) SetBody(v *CloseEngineMigrationResponseBody) *CloseEngineMigrationResponse {
	s.Body = v
	return s
}

func (s *CloseEngineMigrationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
