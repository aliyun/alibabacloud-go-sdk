// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetGroupAuthAppCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetGroupAuthAppCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetGroupAuthAppCodeResponse
	GetStatusCode() *int32
	SetBody(v *SetGroupAuthAppCodeResponseBody) *SetGroupAuthAppCodeResponse
	GetBody() *SetGroupAuthAppCodeResponseBody
}

type SetGroupAuthAppCodeResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetGroupAuthAppCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetGroupAuthAppCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s SetGroupAuthAppCodeResponse) GoString() string {
	return s.String()
}

func (s *SetGroupAuthAppCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetGroupAuthAppCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetGroupAuthAppCodeResponse) GetBody() *SetGroupAuthAppCodeResponseBody {
	return s.Body
}

func (s *SetGroupAuthAppCodeResponse) SetHeaders(v map[string]*string) *SetGroupAuthAppCodeResponse {
	s.Headers = v
	return s
}

func (s *SetGroupAuthAppCodeResponse) SetStatusCode(v int32) *SetGroupAuthAppCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *SetGroupAuthAppCodeResponse) SetBody(v *SetGroupAuthAppCodeResponseBody) *SetGroupAuthAppCodeResponse {
	s.Body = v
	return s
}

func (s *SetGroupAuthAppCodeResponse) Validate() error {
	return dara.Validate(s)
}
