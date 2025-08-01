// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddServiceSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddServiceSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddServiceSourceResponse
	GetStatusCode() *int32
	SetBody(v *AddServiceSourceResponseBody) *AddServiceSourceResponse
	GetBody() *AddServiceSourceResponseBody
}

type AddServiceSourceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddServiceSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddServiceSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s AddServiceSourceResponse) GoString() string {
	return s.String()
}

func (s *AddServiceSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddServiceSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddServiceSourceResponse) GetBody() *AddServiceSourceResponseBody {
	return s.Body
}

func (s *AddServiceSourceResponse) SetHeaders(v map[string]*string) *AddServiceSourceResponse {
	s.Headers = v
	return s
}

func (s *AddServiceSourceResponse) SetStatusCode(v int32) *AddServiceSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *AddServiceSourceResponse) SetBody(v *AddServiceSourceResponseBody) *AddServiceSourceResponse {
	s.Body = v
	return s
}

func (s *AddServiceSourceResponse) Validate() error {
	return dara.Validate(s)
}
