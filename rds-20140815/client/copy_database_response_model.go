// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyDatabaseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CopyDatabaseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CopyDatabaseResponse
	GetStatusCode() *int32
	SetBody(v *CopyDatabaseResponseBody) *CopyDatabaseResponse
	GetBody() *CopyDatabaseResponseBody
}

type CopyDatabaseResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CopyDatabaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CopyDatabaseResponse) String() string {
	return dara.Prettify(s)
}

func (s CopyDatabaseResponse) GoString() string {
	return s.String()
}

func (s *CopyDatabaseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CopyDatabaseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CopyDatabaseResponse) GetBody() *CopyDatabaseResponseBody {
	return s.Body
}

func (s *CopyDatabaseResponse) SetHeaders(v map[string]*string) *CopyDatabaseResponse {
	s.Headers = v
	return s
}

func (s *CopyDatabaseResponse) SetStatusCode(v int32) *CopyDatabaseResponse {
	s.StatusCode = &v
	return s
}

func (s *CopyDatabaseResponse) SetBody(v *CopyDatabaseResponseBody) *CopyDatabaseResponse {
	s.Body = v
	return s
}

func (s *CopyDatabaseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
