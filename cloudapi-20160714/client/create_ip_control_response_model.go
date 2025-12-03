// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIpControlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateIpControlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateIpControlResponse
	GetStatusCode() *int32
	SetBody(v *CreateIpControlResponseBody) *CreateIpControlResponse
	GetBody() *CreateIpControlResponseBody
}

type CreateIpControlResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateIpControlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateIpControlResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateIpControlResponse) GoString() string {
	return s.String()
}

func (s *CreateIpControlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateIpControlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateIpControlResponse) GetBody() *CreateIpControlResponseBody {
	return s.Body
}

func (s *CreateIpControlResponse) SetHeaders(v map[string]*string) *CreateIpControlResponse {
	s.Headers = v
	return s
}

func (s *CreateIpControlResponse) SetStatusCode(v int32) *CreateIpControlResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateIpControlResponse) SetBody(v *CreateIpControlResponseBody) *CreateIpControlResponse {
	s.Body = v
	return s
}

func (s *CreateIpControlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
