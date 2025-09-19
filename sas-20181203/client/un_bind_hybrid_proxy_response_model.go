// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnBindHybridProxyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnBindHybridProxyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnBindHybridProxyResponse
	GetStatusCode() *int32
	SetBody(v *UnBindHybridProxyResponseBody) *UnBindHybridProxyResponse
	GetBody() *UnBindHybridProxyResponseBody
}

type UnBindHybridProxyResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnBindHybridProxyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnBindHybridProxyResponse) String() string {
	return dara.Prettify(s)
}

func (s UnBindHybridProxyResponse) GoString() string {
	return s.String()
}

func (s *UnBindHybridProxyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnBindHybridProxyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnBindHybridProxyResponse) GetBody() *UnBindHybridProxyResponseBody {
	return s.Body
}

func (s *UnBindHybridProxyResponse) SetHeaders(v map[string]*string) *UnBindHybridProxyResponse {
	s.Headers = v
	return s
}

func (s *UnBindHybridProxyResponse) SetStatusCode(v int32) *UnBindHybridProxyResponse {
	s.StatusCode = &v
	return s
}

func (s *UnBindHybridProxyResponse) SetBody(v *UnBindHybridProxyResponseBody) *UnBindHybridProxyResponse {
	s.Body = v
	return s
}

func (s *UnBindHybridProxyResponse) Validate() error {
	return dara.Validate(s)
}
