// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMediaConnectFlowStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMediaConnectFlowStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMediaConnectFlowStatusResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMediaConnectFlowStatusResponseBody) *UpdateMediaConnectFlowStatusResponse
	GetBody() *UpdateMediaConnectFlowStatusResponseBody
}

type UpdateMediaConnectFlowStatusResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMediaConnectFlowStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMediaConnectFlowStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaConnectFlowStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateMediaConnectFlowStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMediaConnectFlowStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMediaConnectFlowStatusResponse) GetBody() *UpdateMediaConnectFlowStatusResponseBody {
	return s.Body
}

func (s *UpdateMediaConnectFlowStatusResponse) SetHeaders(v map[string]*string) *UpdateMediaConnectFlowStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateMediaConnectFlowStatusResponse) SetStatusCode(v int32) *UpdateMediaConnectFlowStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMediaConnectFlowStatusResponse) SetBody(v *UpdateMediaConnectFlowStatusResponseBody) *UpdateMediaConnectFlowStatusResponse {
	s.Body = v
	return s
}

func (s *UpdateMediaConnectFlowStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
