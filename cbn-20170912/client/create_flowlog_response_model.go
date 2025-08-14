// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFlowlogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateFlowlogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateFlowlogResponse
	GetStatusCode() *int32
	SetBody(v *CreateFlowlogResponseBody) *CreateFlowlogResponse
	GetBody() *CreateFlowlogResponseBody
}

type CreateFlowlogResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFlowlogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFlowlogResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateFlowlogResponse) GoString() string {
	return s.String()
}

func (s *CreateFlowlogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateFlowlogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateFlowlogResponse) GetBody() *CreateFlowlogResponseBody {
	return s.Body
}

func (s *CreateFlowlogResponse) SetHeaders(v map[string]*string) *CreateFlowlogResponse {
	s.Headers = v
	return s
}

func (s *CreateFlowlogResponse) SetStatusCode(v int32) *CreateFlowlogResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFlowlogResponse) SetBody(v *CreateFlowlogResponseBody) *CreateFlowlogResponse {
	s.Body = v
	return s
}

func (s *CreateFlowlogResponse) Validate() error {
	return dara.Validate(s)
}
