// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetTrafficControlApisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetTrafficControlApisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetTrafficControlApisResponse
	GetStatusCode() *int32
	SetBody(v *SetTrafficControlApisResponseBody) *SetTrafficControlApisResponse
	GetBody() *SetTrafficControlApisResponseBody
}

type SetTrafficControlApisResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetTrafficControlApisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetTrafficControlApisResponse) String() string {
	return dara.Prettify(s)
}

func (s SetTrafficControlApisResponse) GoString() string {
	return s.String()
}

func (s *SetTrafficControlApisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetTrafficControlApisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetTrafficControlApisResponse) GetBody() *SetTrafficControlApisResponseBody {
	return s.Body
}

func (s *SetTrafficControlApisResponse) SetHeaders(v map[string]*string) *SetTrafficControlApisResponse {
	s.Headers = v
	return s
}

func (s *SetTrafficControlApisResponse) SetStatusCode(v int32) *SetTrafficControlApisResponse {
	s.StatusCode = &v
	return s
}

func (s *SetTrafficControlApisResponse) SetBody(v *SetTrafficControlApisResponseBody) *SetTrafficControlApisResponse {
	s.Body = v
	return s
}

func (s *SetTrafficControlApisResponse) Validate() error {
	return dara.Validate(s)
}
