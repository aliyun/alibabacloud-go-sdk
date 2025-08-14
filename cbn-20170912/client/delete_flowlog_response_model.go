// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFlowlogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteFlowlogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteFlowlogResponse
	GetStatusCode() *int32
	SetBody(v *DeleteFlowlogResponseBody) *DeleteFlowlogResponse
	GetBody() *DeleteFlowlogResponseBody
}

type DeleteFlowlogResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFlowlogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFlowlogResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteFlowlogResponse) GoString() string {
	return s.String()
}

func (s *DeleteFlowlogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteFlowlogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteFlowlogResponse) GetBody() *DeleteFlowlogResponseBody {
	return s.Body
}

func (s *DeleteFlowlogResponse) SetHeaders(v map[string]*string) *DeleteFlowlogResponse {
	s.Headers = v
	return s
}

func (s *DeleteFlowlogResponse) SetStatusCode(v int32) *DeleteFlowlogResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFlowlogResponse) SetBody(v *DeleteFlowlogResponseBody) *DeleteFlowlogResponse {
	s.Body = v
	return s
}

func (s *DeleteFlowlogResponse) Validate() error {
	return dara.Validate(s)
}
