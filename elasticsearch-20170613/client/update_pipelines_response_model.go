// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePipelinesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePipelinesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePipelinesResponse
	GetStatusCode() *int32
	SetBody(v *UpdatePipelinesResponseBody) *UpdatePipelinesResponse
	GetBody() *UpdatePipelinesResponseBody
}

type UpdatePipelinesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePipelinesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePipelinesResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePipelinesResponse) GoString() string {
	return s.String()
}

func (s *UpdatePipelinesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePipelinesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePipelinesResponse) GetBody() *UpdatePipelinesResponseBody {
	return s.Body
}

func (s *UpdatePipelinesResponse) SetHeaders(v map[string]*string) *UpdatePipelinesResponse {
	s.Headers = v
	return s
}

func (s *UpdatePipelinesResponse) SetStatusCode(v int32) *UpdatePipelinesResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePipelinesResponse) SetBody(v *UpdatePipelinesResponseBody) *UpdatePipelinesResponse {
	s.Body = v
	return s
}

func (s *UpdatePipelinesResponse) Validate() error {
	return dara.Validate(s)
}
