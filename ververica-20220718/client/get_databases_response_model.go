// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDatabasesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDatabasesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDatabasesResponse
	GetStatusCode() *int32
	SetBody(v *GetDatabasesResponseBody) *GetDatabasesResponse
	GetBody() *GetDatabasesResponseBody
}

type GetDatabasesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDatabasesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDatabasesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDatabasesResponse) GoString() string {
	return s.String()
}

func (s *GetDatabasesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDatabasesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDatabasesResponse) GetBody() *GetDatabasesResponseBody {
	return s.Body
}

func (s *GetDatabasesResponse) SetHeaders(v map[string]*string) *GetDatabasesResponse {
	s.Headers = v
	return s
}

func (s *GetDatabasesResponse) SetStatusCode(v int32) *GetDatabasesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDatabasesResponse) SetBody(v *GetDatabasesResponseBody) *GetDatabasesResponse {
	s.Body = v
	return s
}

func (s *GetDatabasesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
