// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCrowdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCrowdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCrowdResponse
	GetStatusCode() *int32
	SetBody(v *CreateCrowdResponseBody) *CreateCrowdResponse
	GetBody() *CreateCrowdResponseBody
}

type CreateCrowdResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCrowdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCrowdResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCrowdResponse) GoString() string {
	return s.String()
}

func (s *CreateCrowdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCrowdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCrowdResponse) GetBody() *CreateCrowdResponseBody {
	return s.Body
}

func (s *CreateCrowdResponse) SetHeaders(v map[string]*string) *CreateCrowdResponse {
	s.Headers = v
	return s
}

func (s *CreateCrowdResponse) SetStatusCode(v int32) *CreateCrowdResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCrowdResponse) SetBody(v *CreateCrowdResponseBody) *CreateCrowdResponse {
	s.Body = v
	return s
}

func (s *CreateCrowdResponse) Validate() error {
	return dara.Validate(s)
}
