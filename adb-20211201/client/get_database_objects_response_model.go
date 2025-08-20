// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDatabaseObjectsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDatabaseObjectsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDatabaseObjectsResponse
	GetStatusCode() *int32
	SetBody(v *GetDatabaseObjectsResponseBody) *GetDatabaseObjectsResponse
	GetBody() *GetDatabaseObjectsResponseBody
}

type GetDatabaseObjectsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDatabaseObjectsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDatabaseObjectsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDatabaseObjectsResponse) GoString() string {
	return s.String()
}

func (s *GetDatabaseObjectsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDatabaseObjectsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDatabaseObjectsResponse) GetBody() *GetDatabaseObjectsResponseBody {
	return s.Body
}

func (s *GetDatabaseObjectsResponse) SetHeaders(v map[string]*string) *GetDatabaseObjectsResponse {
	s.Headers = v
	return s
}

func (s *GetDatabaseObjectsResponse) SetStatusCode(v int32) *GetDatabaseObjectsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDatabaseObjectsResponse) SetBody(v *GetDatabaseObjectsResponseBody) *GetDatabaseObjectsResponse {
	s.Body = v
	return s
}

func (s *GetDatabaseObjectsResponse) Validate() error {
	return dara.Validate(s)
}
