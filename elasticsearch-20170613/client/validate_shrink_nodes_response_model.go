// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateShrinkNodesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ValidateShrinkNodesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ValidateShrinkNodesResponse
	GetStatusCode() *int32
	SetBody(v *ValidateShrinkNodesResponseBody) *ValidateShrinkNodesResponse
	GetBody() *ValidateShrinkNodesResponseBody
}

type ValidateShrinkNodesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ValidateShrinkNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ValidateShrinkNodesResponse) String() string {
	return dara.Prettify(s)
}

func (s ValidateShrinkNodesResponse) GoString() string {
	return s.String()
}

func (s *ValidateShrinkNodesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ValidateShrinkNodesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ValidateShrinkNodesResponse) GetBody() *ValidateShrinkNodesResponseBody {
	return s.Body
}

func (s *ValidateShrinkNodesResponse) SetHeaders(v map[string]*string) *ValidateShrinkNodesResponse {
	s.Headers = v
	return s
}

func (s *ValidateShrinkNodesResponse) SetStatusCode(v int32) *ValidateShrinkNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *ValidateShrinkNodesResponse) SetBody(v *ValidateShrinkNodesResponseBody) *ValidateShrinkNodesResponse {
	s.Body = v
	return s
}

func (s *ValidateShrinkNodesResponse) Validate() error {
	return dara.Validate(s)
}
