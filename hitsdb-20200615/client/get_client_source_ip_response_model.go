// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClientSourceIpResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetClientSourceIpResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetClientSourceIpResponse
	GetStatusCode() *int32
	SetBody(v *GetClientSourceIpResponseBody) *GetClientSourceIpResponse
	GetBody() *GetClientSourceIpResponseBody
}

type GetClientSourceIpResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetClientSourceIpResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetClientSourceIpResponse) String() string {
	return dara.Prettify(s)
}

func (s GetClientSourceIpResponse) GoString() string {
	return s.String()
}

func (s *GetClientSourceIpResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetClientSourceIpResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetClientSourceIpResponse) GetBody() *GetClientSourceIpResponseBody {
	return s.Body
}

func (s *GetClientSourceIpResponse) SetHeaders(v map[string]*string) *GetClientSourceIpResponse {
	s.Headers = v
	return s
}

func (s *GetClientSourceIpResponse) SetStatusCode(v int32) *GetClientSourceIpResponse {
	s.StatusCode = &v
	return s
}

func (s *GetClientSourceIpResponse) SetBody(v *GetClientSourceIpResponseBody) *GetClientSourceIpResponse {
	s.Body = v
	return s
}

func (s *GetClientSourceIpResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
