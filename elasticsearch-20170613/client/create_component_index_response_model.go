// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateComponentIndexResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateComponentIndexResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateComponentIndexResponse
	GetStatusCode() *int32
	SetBody(v *CreateComponentIndexResponseBody) *CreateComponentIndexResponse
	GetBody() *CreateComponentIndexResponseBody
}

type CreateComponentIndexResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateComponentIndexResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateComponentIndexResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateComponentIndexResponse) GoString() string {
	return s.String()
}

func (s *CreateComponentIndexResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateComponentIndexResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateComponentIndexResponse) GetBody() *CreateComponentIndexResponseBody {
	return s.Body
}

func (s *CreateComponentIndexResponse) SetHeaders(v map[string]*string) *CreateComponentIndexResponse {
	s.Headers = v
	return s
}

func (s *CreateComponentIndexResponse) SetStatusCode(v int32) *CreateComponentIndexResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateComponentIndexResponse) SetBody(v *CreateComponentIndexResponseBody) *CreateComponentIndexResponse {
	s.Body = v
	return s
}

func (s *CreateComponentIndexResponse) Validate() error {
	return dara.Validate(s)
}
