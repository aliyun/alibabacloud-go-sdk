// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetVpcAccessResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetVpcAccessResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetVpcAccessResponse
	GetStatusCode() *int32
	SetBody(v *SetVpcAccessResponseBody) *SetVpcAccessResponse
	GetBody() *SetVpcAccessResponseBody
}

type SetVpcAccessResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetVpcAccessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetVpcAccessResponse) String() string {
	return dara.Prettify(s)
}

func (s SetVpcAccessResponse) GoString() string {
	return s.String()
}

func (s *SetVpcAccessResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetVpcAccessResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetVpcAccessResponse) GetBody() *SetVpcAccessResponseBody {
	return s.Body
}

func (s *SetVpcAccessResponse) SetHeaders(v map[string]*string) *SetVpcAccessResponse {
	s.Headers = v
	return s
}

func (s *SetVpcAccessResponse) SetStatusCode(v int32) *SetVpcAccessResponse {
	s.StatusCode = &v
	return s
}

func (s *SetVpcAccessResponse) SetBody(v *SetVpcAccessResponseBody) *SetVpcAccessResponse {
	s.Body = v
	return s
}

func (s *SetVpcAccessResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
