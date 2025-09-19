// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDynamicDictResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDynamicDictResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDynamicDictResponse
	GetStatusCode() *int32
	SetBody(v *CreateDynamicDictResponseBody) *CreateDynamicDictResponse
	GetBody() *CreateDynamicDictResponseBody
}

type CreateDynamicDictResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDynamicDictResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDynamicDictResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDynamicDictResponse) GoString() string {
	return s.String()
}

func (s *CreateDynamicDictResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDynamicDictResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDynamicDictResponse) GetBody() *CreateDynamicDictResponseBody {
	return s.Body
}

func (s *CreateDynamicDictResponse) SetHeaders(v map[string]*string) *CreateDynamicDictResponse {
	s.Headers = v
	return s
}

func (s *CreateDynamicDictResponse) SetStatusCode(v int32) *CreateDynamicDictResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDynamicDictResponse) SetBody(v *CreateDynamicDictResponseBody) *CreateDynamicDictResponse {
	s.Body = v
	return s
}

func (s *CreateDynamicDictResponse) Validate() error {
	return dara.Validate(s)
}
