// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPfsSqlSummariesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPfsSqlSummariesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPfsSqlSummariesResponse
	GetStatusCode() *int32
	SetBody(v *GetPfsSqlSummariesResponseBody) *GetPfsSqlSummariesResponse
	GetBody() *GetPfsSqlSummariesResponseBody
}

type GetPfsSqlSummariesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPfsSqlSummariesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPfsSqlSummariesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPfsSqlSummariesResponse) GoString() string {
	return s.String()
}

func (s *GetPfsSqlSummariesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPfsSqlSummariesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPfsSqlSummariesResponse) GetBody() *GetPfsSqlSummariesResponseBody {
	return s.Body
}

func (s *GetPfsSqlSummariesResponse) SetHeaders(v map[string]*string) *GetPfsSqlSummariesResponse {
	s.Headers = v
	return s
}

func (s *GetPfsSqlSummariesResponse) SetStatusCode(v int32) *GetPfsSqlSummariesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPfsSqlSummariesResponse) SetBody(v *GetPfsSqlSummariesResponseBody) *GetPfsSqlSummariesResponse {
	s.Body = v
	return s
}

func (s *GetPfsSqlSummariesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
