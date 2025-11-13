// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRspDomainStatusOteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateRspDomainStatusOteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateRspDomainStatusOteResponse
	GetStatusCode() *int32
	SetBody(v *UpdateRspDomainStatusOteResponseBody) *UpdateRspDomainStatusOteResponse
	GetBody() *UpdateRspDomainStatusOteResponseBody
}

type UpdateRspDomainStatusOteResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRspDomainStatusOteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRspDomainStatusOteResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateRspDomainStatusOteResponse) GoString() string {
	return s.String()
}

func (s *UpdateRspDomainStatusOteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateRspDomainStatusOteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateRspDomainStatusOteResponse) GetBody() *UpdateRspDomainStatusOteResponseBody {
	return s.Body
}

func (s *UpdateRspDomainStatusOteResponse) SetHeaders(v map[string]*string) *UpdateRspDomainStatusOteResponse {
	s.Headers = v
	return s
}

func (s *UpdateRspDomainStatusOteResponse) SetStatusCode(v int32) *UpdateRspDomainStatusOteResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRspDomainStatusOteResponse) SetBody(v *UpdateRspDomainStatusOteResponseBody) *UpdateRspDomainStatusOteResponse {
	s.Body = v
	return s
}

func (s *UpdateRspDomainStatusOteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
