// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFaqResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateFaqResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateFaqResponse
	GetStatusCode() *int32
	SetBody(v *CreateFaqResponseBody) *CreateFaqResponse
	GetBody() *CreateFaqResponseBody
}

type CreateFaqResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFaqResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFaqResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateFaqResponse) GoString() string {
	return s.String()
}

func (s *CreateFaqResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateFaqResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateFaqResponse) GetBody() *CreateFaqResponseBody {
	return s.Body
}

func (s *CreateFaqResponse) SetHeaders(v map[string]*string) *CreateFaqResponse {
	s.Headers = v
	return s
}

func (s *CreateFaqResponse) SetStatusCode(v int32) *CreateFaqResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFaqResponse) SetBody(v *CreateFaqResponseBody) *CreateFaqResponse {
	s.Body = v
	return s
}

func (s *CreateFaqResponse) Validate() error {
	return dara.Validate(s)
}
