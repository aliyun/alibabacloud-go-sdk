// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCasterLayoutResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddCasterLayoutResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddCasterLayoutResponse
	GetStatusCode() *int32
	SetBody(v *AddCasterLayoutResponseBody) *AddCasterLayoutResponse
	GetBody() *AddCasterLayoutResponseBody
}

type AddCasterLayoutResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddCasterLayoutResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddCasterLayoutResponse) String() string {
	return dara.Prettify(s)
}

func (s AddCasterLayoutResponse) GoString() string {
	return s.String()
}

func (s *AddCasterLayoutResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddCasterLayoutResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddCasterLayoutResponse) GetBody() *AddCasterLayoutResponseBody {
	return s.Body
}

func (s *AddCasterLayoutResponse) SetHeaders(v map[string]*string) *AddCasterLayoutResponse {
	s.Headers = v
	return s
}

func (s *AddCasterLayoutResponse) SetStatusCode(v int32) *AddCasterLayoutResponse {
	s.StatusCode = &v
	return s
}

func (s *AddCasterLayoutResponse) SetBody(v *AddCasterLayoutResponseBody) *AddCasterLayoutResponse {
	s.Body = v
	return s
}

func (s *AddCasterLayoutResponse) Validate() error {
	return dara.Validate(s)
}
