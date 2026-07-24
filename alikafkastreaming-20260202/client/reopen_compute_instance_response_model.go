// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReopenComputeInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReopenComputeInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReopenComputeInstanceResponse
	GetStatusCode() *int32
	SetBody(v *ReopenComputeInstanceResponseBody) *ReopenComputeInstanceResponse
	GetBody() *ReopenComputeInstanceResponseBody
}

type ReopenComputeInstanceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReopenComputeInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReopenComputeInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s ReopenComputeInstanceResponse) GoString() string {
	return s.String()
}

func (s *ReopenComputeInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReopenComputeInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReopenComputeInstanceResponse) GetBody() *ReopenComputeInstanceResponseBody {
	return s.Body
}

func (s *ReopenComputeInstanceResponse) SetHeaders(v map[string]*string) *ReopenComputeInstanceResponse {
	s.Headers = v
	return s
}

func (s *ReopenComputeInstanceResponse) SetStatusCode(v int32) *ReopenComputeInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ReopenComputeInstanceResponse) SetBody(v *ReopenComputeInstanceResponseBody) *ReopenComputeInstanceResponse {
	s.Body = v
	return s
}

func (s *ReopenComputeInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
