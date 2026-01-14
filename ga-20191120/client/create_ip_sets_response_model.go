// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIpSetsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateIpSetsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateIpSetsResponse
	GetStatusCode() *int32
	SetBody(v *CreateIpSetsResponseBody) *CreateIpSetsResponse
	GetBody() *CreateIpSetsResponseBody
}

type CreateIpSetsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateIpSetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateIpSetsResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateIpSetsResponse) GoString() string {
	return s.String()
}

func (s *CreateIpSetsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateIpSetsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateIpSetsResponse) GetBody() *CreateIpSetsResponseBody {
	return s.Body
}

func (s *CreateIpSetsResponse) SetHeaders(v map[string]*string) *CreateIpSetsResponse {
	s.Headers = v
	return s
}

func (s *CreateIpSetsResponse) SetStatusCode(v int32) *CreateIpSetsResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateIpSetsResponse) SetBody(v *CreateIpSetsResponseBody) *CreateIpSetsResponse {
	s.Body = v
	return s
}

func (s *CreateIpSetsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
