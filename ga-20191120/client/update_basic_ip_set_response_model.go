// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBasicIpSetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateBasicIpSetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateBasicIpSetResponse
	GetStatusCode() *int32
	SetBody(v *UpdateBasicIpSetResponseBody) *UpdateBasicIpSetResponse
	GetBody() *UpdateBasicIpSetResponseBody
}

type UpdateBasicIpSetResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateBasicIpSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateBasicIpSetResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateBasicIpSetResponse) GoString() string {
	return s.String()
}

func (s *UpdateBasicIpSetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateBasicIpSetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateBasicIpSetResponse) GetBody() *UpdateBasicIpSetResponseBody {
	return s.Body
}

func (s *UpdateBasicIpSetResponse) SetHeaders(v map[string]*string) *UpdateBasicIpSetResponse {
	s.Headers = v
	return s
}

func (s *UpdateBasicIpSetResponse) SetStatusCode(v int32) *UpdateBasicIpSetResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateBasicIpSetResponse) SetBody(v *UpdateBasicIpSetResponseBody) *UpdateBasicIpSetResponse {
	s.Body = v
	return s
}

func (s *UpdateBasicIpSetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
