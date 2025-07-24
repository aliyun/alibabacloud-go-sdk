// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDecreaseNodesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DecreaseNodesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DecreaseNodesResponse
	GetStatusCode() *int32
	SetBody(v *DecreaseNodesResponseBody) *DecreaseNodesResponse
	GetBody() *DecreaseNodesResponseBody
}

type DecreaseNodesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DecreaseNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DecreaseNodesResponse) String() string {
	return dara.Prettify(s)
}

func (s DecreaseNodesResponse) GoString() string {
	return s.String()
}

func (s *DecreaseNodesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DecreaseNodesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DecreaseNodesResponse) GetBody() *DecreaseNodesResponseBody {
	return s.Body
}

func (s *DecreaseNodesResponse) SetHeaders(v map[string]*string) *DecreaseNodesResponse {
	s.Headers = v
	return s
}

func (s *DecreaseNodesResponse) SetStatusCode(v int32) *DecreaseNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *DecreaseNodesResponse) SetBody(v *DecreaseNodesResponseBody) *DecreaseNodesResponse {
	s.Body = v
	return s
}

func (s *DecreaseNodesResponse) Validate() error {
	return dara.Validate(s)
}
