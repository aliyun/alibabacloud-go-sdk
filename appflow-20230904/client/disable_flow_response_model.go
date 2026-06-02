// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableFlowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableFlowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableFlowResponse
	GetStatusCode() *int32
	SetBody(v *DisableFlowResponseBody) *DisableFlowResponse
	GetBody() *DisableFlowResponseBody
}

type DisableFlowResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableFlowResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableFlowResponse) GoString() string {
	return s.String()
}

func (s *DisableFlowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableFlowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableFlowResponse) GetBody() *DisableFlowResponseBody {
	return s.Body
}

func (s *DisableFlowResponse) SetHeaders(v map[string]*string) *DisableFlowResponse {
	s.Headers = v
	return s
}

func (s *DisableFlowResponse) SetStatusCode(v int32) *DisableFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableFlowResponse) SetBody(v *DisableFlowResponseBody) *DisableFlowResponse {
	s.Body = v
	return s
}

func (s *DisableFlowResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
