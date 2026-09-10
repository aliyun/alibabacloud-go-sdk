// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSqlTableLineageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSqlTableLineageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSqlTableLineageResponse
	GetStatusCode() *int32
	SetBody(v *GetSqlTableLineageResponseBody) *GetSqlTableLineageResponse
	GetBody() *GetSqlTableLineageResponseBody
}

type GetSqlTableLineageResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSqlTableLineageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSqlTableLineageResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSqlTableLineageResponse) GoString() string {
	return s.String()
}

func (s *GetSqlTableLineageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSqlTableLineageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSqlTableLineageResponse) GetBody() *GetSqlTableLineageResponseBody {
	return s.Body
}

func (s *GetSqlTableLineageResponse) SetHeaders(v map[string]*string) *GetSqlTableLineageResponse {
	s.Headers = v
	return s
}

func (s *GetSqlTableLineageResponse) SetStatusCode(v int32) *GetSqlTableLineageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSqlTableLineageResponse) SetBody(v *GetSqlTableLineageResponseBody) *GetSqlTableLineageResponse {
	s.Body = v
	return s
}

func (s *GetSqlTableLineageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
