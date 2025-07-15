// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCasterVideoResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddCasterVideoResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddCasterVideoResourceResponse
	GetStatusCode() *int32
	SetBody(v *AddCasterVideoResourceResponseBody) *AddCasterVideoResourceResponse
	GetBody() *AddCasterVideoResourceResponseBody
}

type AddCasterVideoResourceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddCasterVideoResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddCasterVideoResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s AddCasterVideoResourceResponse) GoString() string {
	return s.String()
}

func (s *AddCasterVideoResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddCasterVideoResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddCasterVideoResourceResponse) GetBody() *AddCasterVideoResourceResponseBody {
	return s.Body
}

func (s *AddCasterVideoResourceResponse) SetHeaders(v map[string]*string) *AddCasterVideoResourceResponse {
	s.Headers = v
	return s
}

func (s *AddCasterVideoResourceResponse) SetStatusCode(v int32) *AddCasterVideoResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *AddCasterVideoResourceResponse) SetBody(v *AddCasterVideoResourceResponseBody) *AddCasterVideoResourceResponse {
	s.Body = v
	return s
}

func (s *AddCasterVideoResourceResponse) Validate() error {
	return dara.Validate(s)
}
