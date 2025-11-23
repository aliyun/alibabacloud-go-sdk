// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLogicDatabaseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLogicDatabaseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLogicDatabaseResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLogicDatabaseResponseBody) *DeleteLogicDatabaseResponse
	GetBody() *DeleteLogicDatabaseResponseBody
}

type DeleteLogicDatabaseResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLogicDatabaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLogicDatabaseResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLogicDatabaseResponse) GoString() string {
	return s.String()
}

func (s *DeleteLogicDatabaseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLogicDatabaseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLogicDatabaseResponse) GetBody() *DeleteLogicDatabaseResponseBody {
	return s.Body
}

func (s *DeleteLogicDatabaseResponse) SetHeaders(v map[string]*string) *DeleteLogicDatabaseResponse {
	s.Headers = v
	return s
}

func (s *DeleteLogicDatabaseResponse) SetStatusCode(v int32) *DeleteLogicDatabaseResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLogicDatabaseResponse) SetBody(v *DeleteLogicDatabaseResponseBody) *DeleteLogicDatabaseResponse {
	s.Body = v
	return s
}

func (s *DeleteLogicDatabaseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
