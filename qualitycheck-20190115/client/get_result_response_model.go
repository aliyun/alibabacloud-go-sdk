// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetResultResponse
	GetStatusCode() *int32
	SetBody(v *GetResultResponseBody) *GetResultResponse
	GetBody() *GetResultResponseBody
}

type GetResultResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetResultResponse) GoString() string {
	return s.String()
}

func (s *GetResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetResultResponse) GetBody() *GetResultResponseBody {
	return s.Body
}

func (s *GetResultResponse) SetHeaders(v map[string]*string) *GetResultResponse {
	s.Headers = v
	return s
}

func (s *GetResultResponse) SetStatusCode(v int32) *GetResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResultResponse) SetBody(v *GetResultResponseBody) *GetResultResponse {
	s.Body = v
	return s
}

func (s *GetResultResponse) Validate() error {
	return dara.Validate(s)
}
