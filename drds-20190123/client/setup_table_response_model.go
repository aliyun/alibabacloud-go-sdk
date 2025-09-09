// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetupTableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetupTableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetupTableResponse
	GetStatusCode() *int32
	SetBody(v *SetupTableResponseBody) *SetupTableResponse
	GetBody() *SetupTableResponseBody
}

type SetupTableResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetupTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetupTableResponse) String() string {
	return dara.Prettify(s)
}

func (s SetupTableResponse) GoString() string {
	return s.String()
}

func (s *SetupTableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetupTableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetupTableResponse) GetBody() *SetupTableResponseBody {
	return s.Body
}

func (s *SetupTableResponse) SetHeaders(v map[string]*string) *SetupTableResponse {
	s.Headers = v
	return s
}

func (s *SetupTableResponse) SetStatusCode(v int32) *SetupTableResponse {
	s.StatusCode = &v
	return s
}

func (s *SetupTableResponse) SetBody(v *SetupTableResponseBody) *SetupTableResponse {
	s.Body = v
	return s
}

func (s *SetupTableResponse) Validate() error {
	return dara.Validate(s)
}
