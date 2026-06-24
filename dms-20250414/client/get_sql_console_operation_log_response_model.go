// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSqlConsoleOperationLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSqlConsoleOperationLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSqlConsoleOperationLogResponse
	GetStatusCode() *int32
	SetBody(v *GetSqlConsoleOperationLogResponseBody) *GetSqlConsoleOperationLogResponse
	GetBody() *GetSqlConsoleOperationLogResponseBody
}

type GetSqlConsoleOperationLogResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSqlConsoleOperationLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSqlConsoleOperationLogResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSqlConsoleOperationLogResponse) GoString() string {
	return s.String()
}

func (s *GetSqlConsoleOperationLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSqlConsoleOperationLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSqlConsoleOperationLogResponse) GetBody() *GetSqlConsoleOperationLogResponseBody {
	return s.Body
}

func (s *GetSqlConsoleOperationLogResponse) SetHeaders(v map[string]*string) *GetSqlConsoleOperationLogResponse {
	s.Headers = v
	return s
}

func (s *GetSqlConsoleOperationLogResponse) SetStatusCode(v int32) *GetSqlConsoleOperationLogResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSqlConsoleOperationLogResponse) SetBody(v *GetSqlConsoleOperationLogResponseBody) *GetSqlConsoleOperationLogResponse {
	s.Body = v
	return s
}

func (s *GetSqlConsoleOperationLogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
