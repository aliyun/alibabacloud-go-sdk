// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindHybridProxyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BindHybridProxyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BindHybridProxyResponse
	GetStatusCode() *int32
	SetBody(v *BindHybridProxyResponseBody) *BindHybridProxyResponse
	GetBody() *BindHybridProxyResponseBody
}

type BindHybridProxyResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindHybridProxyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindHybridProxyResponse) String() string {
	return dara.Prettify(s)
}

func (s BindHybridProxyResponse) GoString() string {
	return s.String()
}

func (s *BindHybridProxyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BindHybridProxyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BindHybridProxyResponse) GetBody() *BindHybridProxyResponseBody {
	return s.Body
}

func (s *BindHybridProxyResponse) SetHeaders(v map[string]*string) *BindHybridProxyResponse {
	s.Headers = v
	return s
}

func (s *BindHybridProxyResponse) SetStatusCode(v int32) *BindHybridProxyResponse {
	s.StatusCode = &v
	return s
}

func (s *BindHybridProxyResponse) SetBody(v *BindHybridProxyResponseBody) *BindHybridProxyResponse {
	s.Body = v
	return s
}

func (s *BindHybridProxyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
