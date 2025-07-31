// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateNodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateNodeResponse
	GetStatusCode() *int32
	SetBody(v *CreateNodeResponseBody) *CreateNodeResponse
	GetBody() *CreateNodeResponseBody
}

type CreateNodeResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateNodeResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateNodeResponse) GoString() string {
	return s.String()
}

func (s *CreateNodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateNodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateNodeResponse) GetBody() *CreateNodeResponseBody {
	return s.Body
}

func (s *CreateNodeResponse) SetHeaders(v map[string]*string) *CreateNodeResponse {
	s.Headers = v
	return s
}

func (s *CreateNodeResponse) SetStatusCode(v int32) *CreateNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateNodeResponse) SetBody(v *CreateNodeResponseBody) *CreateNodeResponse {
	s.Body = v
	return s
}

func (s *CreateNodeResponse) Validate() error {
	return dara.Validate(s)
}
