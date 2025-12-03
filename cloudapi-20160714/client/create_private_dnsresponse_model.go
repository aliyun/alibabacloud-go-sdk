// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePrivateDNSResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePrivateDNSResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePrivateDNSResponse
	GetStatusCode() *int32
	SetBody(v *CreatePrivateDNSResponseBody) *CreatePrivateDNSResponse
	GetBody() *CreatePrivateDNSResponseBody
}

type CreatePrivateDNSResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePrivateDNSResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePrivateDNSResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePrivateDNSResponse) GoString() string {
	return s.String()
}

func (s *CreatePrivateDNSResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePrivateDNSResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePrivateDNSResponse) GetBody() *CreatePrivateDNSResponseBody {
	return s.Body
}

func (s *CreatePrivateDNSResponse) SetHeaders(v map[string]*string) *CreatePrivateDNSResponse {
	s.Headers = v
	return s
}

func (s *CreatePrivateDNSResponse) SetStatusCode(v int32) *CreatePrivateDNSResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePrivateDNSResponse) SetBody(v *CreatePrivateDNSResponseBody) *CreatePrivateDNSResponse {
	s.Body = v
	return s
}

func (s *CreatePrivateDNSResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
