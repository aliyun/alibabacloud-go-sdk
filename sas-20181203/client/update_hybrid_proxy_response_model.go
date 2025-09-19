// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHybridProxyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateHybridProxyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateHybridProxyResponse
	GetStatusCode() *int32
	SetBody(v *UpdateHybridProxyResponseBody) *UpdateHybridProxyResponse
	GetBody() *UpdateHybridProxyResponseBody
}

type UpdateHybridProxyResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateHybridProxyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateHybridProxyResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateHybridProxyResponse) GoString() string {
	return s.String()
}

func (s *UpdateHybridProxyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateHybridProxyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateHybridProxyResponse) GetBody() *UpdateHybridProxyResponseBody {
	return s.Body
}

func (s *UpdateHybridProxyResponse) SetHeaders(v map[string]*string) *UpdateHybridProxyResponse {
	s.Headers = v
	return s
}

func (s *UpdateHybridProxyResponse) SetStatusCode(v int32) *UpdateHybridProxyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateHybridProxyResponse) SetBody(v *UpdateHybridProxyResponseBody) *UpdateHybridProxyResponse {
	s.Body = v
	return s
}

func (s *UpdateHybridProxyResponse) Validate() error {
	return dara.Validate(s)
}
