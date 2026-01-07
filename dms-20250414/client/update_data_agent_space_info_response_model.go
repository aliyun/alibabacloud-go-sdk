// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataAgentSpaceInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDataAgentSpaceInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDataAgentSpaceInfoResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDataAgentSpaceInfoResponseBody) *UpdateDataAgentSpaceInfoResponse
	GetBody() *UpdateDataAgentSpaceInfoResponseBody
}

type UpdateDataAgentSpaceInfoResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDataAgentSpaceInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDataAgentSpaceInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataAgentSpaceInfoResponse) GoString() string {
	return s.String()
}

func (s *UpdateDataAgentSpaceInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDataAgentSpaceInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDataAgentSpaceInfoResponse) GetBody() *UpdateDataAgentSpaceInfoResponseBody {
	return s.Body
}

func (s *UpdateDataAgentSpaceInfoResponse) SetHeaders(v map[string]*string) *UpdateDataAgentSpaceInfoResponse {
	s.Headers = v
	return s
}

func (s *UpdateDataAgentSpaceInfoResponse) SetStatusCode(v int32) *UpdateDataAgentSpaceInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDataAgentSpaceInfoResponse) SetBody(v *UpdateDataAgentSpaceInfoResponseBody) *UpdateDataAgentSpaceInfoResponse {
	s.Body = v
	return s
}

func (s *UpdateDataAgentSpaceInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
