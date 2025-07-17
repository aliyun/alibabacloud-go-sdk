// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLogicDatabaseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLogicDatabaseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLogicDatabaseResponse
	GetStatusCode() *int32
	SetBody(v *GetLogicDatabaseResponseBody) *GetLogicDatabaseResponse
	GetBody() *GetLogicDatabaseResponseBody
}

type GetLogicDatabaseResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLogicDatabaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLogicDatabaseResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLogicDatabaseResponse) GoString() string {
	return s.String()
}

func (s *GetLogicDatabaseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLogicDatabaseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLogicDatabaseResponse) GetBody() *GetLogicDatabaseResponseBody {
	return s.Body
}

func (s *GetLogicDatabaseResponse) SetHeaders(v map[string]*string) *GetLogicDatabaseResponse {
	s.Headers = v
	return s
}

func (s *GetLogicDatabaseResponse) SetStatusCode(v int32) *GetLogicDatabaseResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLogicDatabaseResponse) SetBody(v *GetLogicDatabaseResponseBody) *GetLogicDatabaseResponse {
	s.Body = v
	return s
}

func (s *GetLogicDatabaseResponse) Validate() error {
	return dara.Validate(s)
}
