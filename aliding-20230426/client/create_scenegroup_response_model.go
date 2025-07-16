// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScenegroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateScenegroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateScenegroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateScenegroupResponseBody) *CreateScenegroupResponse
	GetBody() *CreateScenegroupResponseBody
}

type CreateScenegroupResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateScenegroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateScenegroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateScenegroupResponse) GoString() string {
	return s.String()
}

func (s *CreateScenegroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateScenegroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateScenegroupResponse) GetBody() *CreateScenegroupResponseBody {
	return s.Body
}

func (s *CreateScenegroupResponse) SetHeaders(v map[string]*string) *CreateScenegroupResponse {
	s.Headers = v
	return s
}

func (s *CreateScenegroupResponse) SetStatusCode(v int32) *CreateScenegroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateScenegroupResponse) SetBody(v *CreateScenegroupResponseBody) *CreateScenegroupResponse {
	s.Body = v
	return s
}

func (s *CreateScenegroupResponse) Validate() error {
	return dara.Validate(s)
}
