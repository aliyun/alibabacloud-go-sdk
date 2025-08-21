// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindTagResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BindTagResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BindTagResponse
	GetStatusCode() *int32
	SetBody(v *BindTagResponseBody) *BindTagResponse
	GetBody() *BindTagResponseBody
}

type BindTagResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindTagResponse) String() string {
	return dara.Prettify(s)
}

func (s BindTagResponse) GoString() string {
	return s.String()
}

func (s *BindTagResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BindTagResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BindTagResponse) GetBody() *BindTagResponseBody {
	return s.Body
}

func (s *BindTagResponse) SetHeaders(v map[string]*string) *BindTagResponse {
	s.Headers = v
	return s
}

func (s *BindTagResponse) SetStatusCode(v int32) *BindTagResponse {
	s.StatusCode = &v
	return s
}

func (s *BindTagResponse) SetBody(v *BindTagResponseBody) *BindTagResponse {
	s.Body = v
	return s
}

func (s *BindTagResponse) Validate() error {
	return dara.Validate(s)
}
