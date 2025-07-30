// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReverseTwoWayDirectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReverseTwoWayDirectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReverseTwoWayDirectionResponse
	GetStatusCode() *int32
	SetBody(v *ReverseTwoWayDirectionResponseBody) *ReverseTwoWayDirectionResponse
	GetBody() *ReverseTwoWayDirectionResponseBody
}

type ReverseTwoWayDirectionResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReverseTwoWayDirectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReverseTwoWayDirectionResponse) String() string {
	return dara.Prettify(s)
}

func (s ReverseTwoWayDirectionResponse) GoString() string {
	return s.String()
}

func (s *ReverseTwoWayDirectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReverseTwoWayDirectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReverseTwoWayDirectionResponse) GetBody() *ReverseTwoWayDirectionResponseBody {
	return s.Body
}

func (s *ReverseTwoWayDirectionResponse) SetHeaders(v map[string]*string) *ReverseTwoWayDirectionResponse {
	s.Headers = v
	return s
}

func (s *ReverseTwoWayDirectionResponse) SetStatusCode(v int32) *ReverseTwoWayDirectionResponse {
	s.StatusCode = &v
	return s
}

func (s *ReverseTwoWayDirectionResponse) SetBody(v *ReverseTwoWayDirectionResponseBody) *ReverseTwoWayDirectionResponse {
	s.Body = v
	return s
}

func (s *ReverseTwoWayDirectionResponse) Validate() error {
	return dara.Validate(s)
}
