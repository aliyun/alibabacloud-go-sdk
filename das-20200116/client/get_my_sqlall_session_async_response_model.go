// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMySQLAllSessionAsyncResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMySQLAllSessionAsyncResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMySQLAllSessionAsyncResponse
	GetStatusCode() *int32
	SetBody(v *GetMySQLAllSessionAsyncResponseBody) *GetMySQLAllSessionAsyncResponse
	GetBody() *GetMySQLAllSessionAsyncResponseBody
}

type GetMySQLAllSessionAsyncResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMySQLAllSessionAsyncResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMySQLAllSessionAsyncResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMySQLAllSessionAsyncResponse) GoString() string {
	return s.String()
}

func (s *GetMySQLAllSessionAsyncResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMySQLAllSessionAsyncResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMySQLAllSessionAsyncResponse) GetBody() *GetMySQLAllSessionAsyncResponseBody {
	return s.Body
}

func (s *GetMySQLAllSessionAsyncResponse) SetHeaders(v map[string]*string) *GetMySQLAllSessionAsyncResponse {
	s.Headers = v
	return s
}

func (s *GetMySQLAllSessionAsyncResponse) SetStatusCode(v int32) *GetMySQLAllSessionAsyncResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMySQLAllSessionAsyncResponse) SetBody(v *GetMySQLAllSessionAsyncResponseBody) *GetMySQLAllSessionAsyncResponse {
	s.Body = v
	return s
}

func (s *GetMySQLAllSessionAsyncResponse) Validate() error {
	return dara.Validate(s)
}
