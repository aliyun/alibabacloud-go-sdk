// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExampleQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetExampleQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetExampleQueryResponse
	GetStatusCode() *int32
	SetBody(v *GetExampleQueryResponseBody) *GetExampleQueryResponse
	GetBody() *GetExampleQueryResponseBody
}

type GetExampleQueryResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetExampleQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetExampleQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetExampleQueryResponse) GoString() string {
	return s.String()
}

func (s *GetExampleQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetExampleQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetExampleQueryResponse) GetBody() *GetExampleQueryResponseBody {
	return s.Body
}

func (s *GetExampleQueryResponse) SetHeaders(v map[string]*string) *GetExampleQueryResponse {
	s.Headers = v
	return s
}

func (s *GetExampleQueryResponse) SetStatusCode(v int32) *GetExampleQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetExampleQueryResponse) SetBody(v *GetExampleQueryResponseBody) *GetExampleQueryResponse {
	s.Body = v
	return s
}

func (s *GetExampleQueryResponse) Validate() error {
	return dara.Validate(s)
}
