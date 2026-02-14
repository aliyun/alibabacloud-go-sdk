// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataCenterDatabaseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataCenterDatabaseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataCenterDatabaseResponse
	GetStatusCode() *int32
	SetBody(v *ListDataCenterDatabaseResponseBody) *ListDataCenterDatabaseResponse
	GetBody() *ListDataCenterDatabaseResponseBody
}

type ListDataCenterDatabaseResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataCenterDatabaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataCenterDatabaseResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataCenterDatabaseResponse) GoString() string {
	return s.String()
}

func (s *ListDataCenterDatabaseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataCenterDatabaseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataCenterDatabaseResponse) GetBody() *ListDataCenterDatabaseResponseBody {
	return s.Body
}

func (s *ListDataCenterDatabaseResponse) SetHeaders(v map[string]*string) *ListDataCenterDatabaseResponse {
	s.Headers = v
	return s
}

func (s *ListDataCenterDatabaseResponse) SetStatusCode(v int32) *ListDataCenterDatabaseResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataCenterDatabaseResponse) SetBody(v *ListDataCenterDatabaseResponseBody) *ListDataCenterDatabaseResponse {
	s.Body = v
	return s
}

func (s *ListDataCenterDatabaseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
