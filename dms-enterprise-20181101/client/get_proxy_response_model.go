// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProxyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetProxyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetProxyResponse
	GetStatusCode() *int32
	SetBody(v *GetProxyResponseBody) *GetProxyResponse
	GetBody() *GetProxyResponseBody
}

type GetProxyResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetProxyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetProxyResponse) String() string {
	return dara.Prettify(s)
}

func (s GetProxyResponse) GoString() string {
	return s.String()
}

func (s *GetProxyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetProxyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetProxyResponse) GetBody() *GetProxyResponseBody {
	return s.Body
}

func (s *GetProxyResponse) SetHeaders(v map[string]*string) *GetProxyResponse {
	s.Headers = v
	return s
}

func (s *GetProxyResponse) SetStatusCode(v int32) *GetProxyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProxyResponse) SetBody(v *GetProxyResponseBody) *GetProxyResponse {
	s.Body = v
	return s
}

func (s *GetProxyResponse) Validate() error {
	return dara.Validate(s)
}
