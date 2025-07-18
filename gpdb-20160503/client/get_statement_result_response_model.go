// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStatementResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetStatementResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetStatementResultResponse
	GetStatusCode() *int32
	SetBody(v *GetStatementResultResponseBody) *GetStatementResultResponse
	GetBody() *GetStatementResultResponseBody
}

type GetStatementResultResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetStatementResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetStatementResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetStatementResultResponse) GoString() string {
	return s.String()
}

func (s *GetStatementResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetStatementResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetStatementResultResponse) GetBody() *GetStatementResultResponseBody {
	return s.Body
}

func (s *GetStatementResultResponse) SetHeaders(v map[string]*string) *GetStatementResultResponse {
	s.Headers = v
	return s
}

func (s *GetStatementResultResponse) SetStatusCode(v int32) *GetStatementResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStatementResultResponse) SetBody(v *GetStatementResultResponseBody) *GetStatementResultResponse {
	s.Body = v
	return s
}

func (s *GetStatementResultResponse) Validate() error {
	return dara.Validate(s)
}
