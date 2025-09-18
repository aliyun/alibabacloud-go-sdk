// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMmAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMmAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMmAppResponse
	GetStatusCode() *int32
	SetBody(v *CreateMmAppResponseBody) *CreateMmAppResponse
	GetBody() *CreateMmAppResponseBody
}

type CreateMmAppResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMmAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMmAppResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMmAppResponse) GoString() string {
	return s.String()
}

func (s *CreateMmAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMmAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMmAppResponse) GetBody() *CreateMmAppResponseBody {
	return s.Body
}

func (s *CreateMmAppResponse) SetHeaders(v map[string]*string) *CreateMmAppResponse {
	s.Headers = v
	return s
}

func (s *CreateMmAppResponse) SetStatusCode(v int32) *CreateMmAppResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMmAppResponse) SetBody(v *CreateMmAppResponseBody) *CreateMmAppResponse {
	s.Body = v
	return s
}

func (s *CreateMmAppResponse) Validate() error {
	return dara.Validate(s)
}
