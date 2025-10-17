// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDatabaseZonalResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDatabaseZonalResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDatabaseZonalResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDatabaseZonalResponseBody) *DeleteDatabaseZonalResponse
	GetBody() *DeleteDatabaseZonalResponseBody
}

type DeleteDatabaseZonalResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDatabaseZonalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDatabaseZonalResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDatabaseZonalResponse) GoString() string {
	return s.String()
}

func (s *DeleteDatabaseZonalResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDatabaseZonalResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDatabaseZonalResponse) GetBody() *DeleteDatabaseZonalResponseBody {
	return s.Body
}

func (s *DeleteDatabaseZonalResponse) SetHeaders(v map[string]*string) *DeleteDatabaseZonalResponse {
	s.Headers = v
	return s
}

func (s *DeleteDatabaseZonalResponse) SetStatusCode(v int32) *DeleteDatabaseZonalResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDatabaseZonalResponse) SetBody(v *DeleteDatabaseZonalResponseBody) *DeleteDatabaseZonalResponse {
	s.Body = v
	return s
}

func (s *DeleteDatabaseZonalResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
