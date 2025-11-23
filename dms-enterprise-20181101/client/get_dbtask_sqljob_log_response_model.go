// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDBTaskSQLJobLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDBTaskSQLJobLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDBTaskSQLJobLogResponse
	GetStatusCode() *int32
	SetBody(v *GetDBTaskSQLJobLogResponseBody) *GetDBTaskSQLJobLogResponse
	GetBody() *GetDBTaskSQLJobLogResponseBody
}

type GetDBTaskSQLJobLogResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDBTaskSQLJobLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDBTaskSQLJobLogResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDBTaskSQLJobLogResponse) GoString() string {
	return s.String()
}

func (s *GetDBTaskSQLJobLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDBTaskSQLJobLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDBTaskSQLJobLogResponse) GetBody() *GetDBTaskSQLJobLogResponseBody {
	return s.Body
}

func (s *GetDBTaskSQLJobLogResponse) SetHeaders(v map[string]*string) *GetDBTaskSQLJobLogResponse {
	s.Headers = v
	return s
}

func (s *GetDBTaskSQLJobLogResponse) SetStatusCode(v int32) *GetDBTaskSQLJobLogResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDBTaskSQLJobLogResponse) SetBody(v *GetDBTaskSQLJobLogResponseBody) *GetDBTaskSQLJobLogResponse {
	s.Body = v
	return s
}

func (s *GetDBTaskSQLJobLogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
