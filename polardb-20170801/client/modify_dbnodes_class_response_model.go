// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBNodesClassResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBNodesClassResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBNodesClassResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBNodesClassResponseBody) *ModifyDBNodesClassResponse
	GetBody() *ModifyDBNodesClassResponseBody
}

type ModifyDBNodesClassResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBNodesClassResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBNodesClassResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBNodesClassResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBNodesClassResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBNodesClassResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBNodesClassResponse) GetBody() *ModifyDBNodesClassResponseBody {
	return s.Body
}

func (s *ModifyDBNodesClassResponse) SetHeaders(v map[string]*string) *ModifyDBNodesClassResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBNodesClassResponse) SetStatusCode(v int32) *ModifyDBNodesClassResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBNodesClassResponse) SetBody(v *ModifyDBNodesClassResponseBody) *ModifyDBNodesClassResponse {
	s.Body = v
	return s
}

func (s *ModifyDBNodesClassResponse) Validate() error {
	return dara.Validate(s)
}
