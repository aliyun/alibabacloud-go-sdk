// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFlowVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateFlowVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateFlowVersionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateFlowVersionResponseBody) *UpdateFlowVersionResponse
	GetBody() *UpdateFlowVersionResponseBody
}

type UpdateFlowVersionResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateFlowVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFlowVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateFlowVersionResponse) GoString() string {
	return s.String()
}

func (s *UpdateFlowVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateFlowVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateFlowVersionResponse) GetBody() *UpdateFlowVersionResponseBody {
	return s.Body
}

func (s *UpdateFlowVersionResponse) SetHeaders(v map[string]*string) *UpdateFlowVersionResponse {
	s.Headers = v
	return s
}

func (s *UpdateFlowVersionResponse) SetStatusCode(v int32) *UpdateFlowVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFlowVersionResponse) SetBody(v *UpdateFlowVersionResponseBody) *UpdateFlowVersionResponse {
	s.Body = v
	return s
}

func (s *UpdateFlowVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
