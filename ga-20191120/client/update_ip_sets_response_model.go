// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIpSetsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateIpSetsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateIpSetsResponse
	GetStatusCode() *int32
	SetBody(v *UpdateIpSetsResponseBody) *UpdateIpSetsResponse
	GetBody() *UpdateIpSetsResponseBody
}

type UpdateIpSetsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateIpSetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateIpSetsResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateIpSetsResponse) GoString() string {
	return s.String()
}

func (s *UpdateIpSetsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateIpSetsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateIpSetsResponse) GetBody() *UpdateIpSetsResponseBody {
	return s.Body
}

func (s *UpdateIpSetsResponse) SetHeaders(v map[string]*string) *UpdateIpSetsResponse {
	s.Headers = v
	return s
}

func (s *UpdateIpSetsResponse) SetStatusCode(v int32) *UpdateIpSetsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateIpSetsResponse) SetBody(v *UpdateIpSetsResponseBody) *UpdateIpSetsResponse {
	s.Body = v
	return s
}

func (s *UpdateIpSetsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
