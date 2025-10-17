// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSparkReplStatementResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSparkReplStatementResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSparkReplStatementResponse
	GetStatusCode() *int32
	SetBody(v *GetSparkReplStatementResponseBody) *GetSparkReplStatementResponse
	GetBody() *GetSparkReplStatementResponseBody
}

type GetSparkReplStatementResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSparkReplStatementResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSparkReplStatementResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSparkReplStatementResponse) GoString() string {
	return s.String()
}

func (s *GetSparkReplStatementResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSparkReplStatementResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSparkReplStatementResponse) GetBody() *GetSparkReplStatementResponseBody {
	return s.Body
}

func (s *GetSparkReplStatementResponse) SetHeaders(v map[string]*string) *GetSparkReplStatementResponse {
	s.Headers = v
	return s
}

func (s *GetSparkReplStatementResponse) SetStatusCode(v int32) *GetSparkReplStatementResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSparkReplStatementResponse) SetBody(v *GetSparkReplStatementResponseBody) *GetSparkReplStatementResponse {
	s.Body = v
	return s
}

func (s *GetSparkReplStatementResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
