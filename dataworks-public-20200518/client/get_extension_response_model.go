// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExtensionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetExtensionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetExtensionResponse
	GetStatusCode() *int32
	SetBody(v *GetExtensionResponseBody) *GetExtensionResponse
	GetBody() *GetExtensionResponseBody
}

type GetExtensionResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetExtensionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetExtensionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetExtensionResponse) GoString() string {
	return s.String()
}

func (s *GetExtensionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetExtensionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetExtensionResponse) GetBody() *GetExtensionResponseBody {
	return s.Body
}

func (s *GetExtensionResponse) SetHeaders(v map[string]*string) *GetExtensionResponse {
	s.Headers = v
	return s
}

func (s *GetExtensionResponse) SetStatusCode(v int32) *GetExtensionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetExtensionResponse) SetBody(v *GetExtensionResponseBody) *GetExtensionResponse {
	s.Body = v
	return s
}

func (s *GetExtensionResponse) Validate() error {
	return dara.Validate(s)
}
