// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAIQueryResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAIQueryResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAIQueryResultResponse
	GetStatusCode() *int32
	SetBody(v *GetAIQueryResultResponseBody) *GetAIQueryResultResponse
	GetBody() *GetAIQueryResultResponseBody
}

type GetAIQueryResultResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAIQueryResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAIQueryResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAIQueryResultResponse) GoString() string {
	return s.String()
}

func (s *GetAIQueryResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAIQueryResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAIQueryResultResponse) GetBody() *GetAIQueryResultResponseBody {
	return s.Body
}

func (s *GetAIQueryResultResponse) SetHeaders(v map[string]*string) *GetAIQueryResultResponse {
	s.Headers = v
	return s
}

func (s *GetAIQueryResultResponse) SetStatusCode(v int32) *GetAIQueryResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAIQueryResultResponse) SetBody(v *GetAIQueryResultResponseBody) *GetAIQueryResultResponse {
	s.Body = v
	return s
}

func (s *GetAIQueryResultResponse) Validate() error {
	return dara.Validate(s)
}
