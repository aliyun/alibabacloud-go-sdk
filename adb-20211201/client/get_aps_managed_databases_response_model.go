// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApsManagedDatabasesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetApsManagedDatabasesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetApsManagedDatabasesResponse
	GetStatusCode() *int32
	SetBody(v *GetApsManagedDatabasesResponseBody) *GetApsManagedDatabasesResponse
	GetBody() *GetApsManagedDatabasesResponseBody
}

type GetApsManagedDatabasesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetApsManagedDatabasesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetApsManagedDatabasesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetApsManagedDatabasesResponse) GoString() string {
	return s.String()
}

func (s *GetApsManagedDatabasesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetApsManagedDatabasesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetApsManagedDatabasesResponse) GetBody() *GetApsManagedDatabasesResponseBody {
	return s.Body
}

func (s *GetApsManagedDatabasesResponse) SetHeaders(v map[string]*string) *GetApsManagedDatabasesResponse {
	s.Headers = v
	return s
}

func (s *GetApsManagedDatabasesResponse) SetStatusCode(v int32) *GetApsManagedDatabasesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetApsManagedDatabasesResponse) SetBody(v *GetApsManagedDatabasesResponseBody) *GetApsManagedDatabasesResponse {
	s.Body = v
	return s
}

func (s *GetApsManagedDatabasesResponse) Validate() error {
	return dara.Validate(s)
}
