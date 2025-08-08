// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMdsMiniConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddMdsMiniConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddMdsMiniConfigResponse
	GetStatusCode() *int32
	SetBody(v *AddMdsMiniConfigResponseBody) *AddMdsMiniConfigResponse
	GetBody() *AddMdsMiniConfigResponseBody
}

type AddMdsMiniConfigResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddMdsMiniConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddMdsMiniConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s AddMdsMiniConfigResponse) GoString() string {
	return s.String()
}

func (s *AddMdsMiniConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddMdsMiniConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddMdsMiniConfigResponse) GetBody() *AddMdsMiniConfigResponseBody {
	return s.Body
}

func (s *AddMdsMiniConfigResponse) SetHeaders(v map[string]*string) *AddMdsMiniConfigResponse {
	s.Headers = v
	return s
}

func (s *AddMdsMiniConfigResponse) SetStatusCode(v int32) *AddMdsMiniConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *AddMdsMiniConfigResponse) SetBody(v *AddMdsMiniConfigResponseBody) *AddMdsMiniConfigResponse {
	s.Body = v
	return s
}

func (s *AddMdsMiniConfigResponse) Validate() error {
	return dara.Validate(s)
}
