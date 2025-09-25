// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetParseResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetParseResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetParseResultResponse
	GetStatusCode() *int32
	SetBody(v *GetParseResultResponseBody) *GetParseResultResponse
	GetBody() *GetParseResultResponseBody
}

type GetParseResultResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetParseResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetParseResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetParseResultResponse) GoString() string {
	return s.String()
}

func (s *GetParseResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetParseResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetParseResultResponse) GetBody() *GetParseResultResponseBody {
	return s.Body
}

func (s *GetParseResultResponse) SetHeaders(v map[string]*string) *GetParseResultResponse {
	s.Headers = v
	return s
}

func (s *GetParseResultResponse) SetStatusCode(v int32) *GetParseResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetParseResultResponse) SetBody(v *GetParseResultResponseBody) *GetParseResultResponse {
	s.Body = v
	return s
}

func (s *GetParseResultResponse) Validate() error {
	return dara.Validate(s)
}
