// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProduceForPartnerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateProduceForPartnerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateProduceForPartnerResponse
	GetStatusCode() *int32
	SetBody(v *CreateProduceForPartnerResponseBody) *CreateProduceForPartnerResponse
	GetBody() *CreateProduceForPartnerResponseBody
}

type CreateProduceForPartnerResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateProduceForPartnerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateProduceForPartnerResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateProduceForPartnerResponse) GoString() string {
	return s.String()
}

func (s *CreateProduceForPartnerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateProduceForPartnerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateProduceForPartnerResponse) GetBody() *CreateProduceForPartnerResponseBody {
	return s.Body
}

func (s *CreateProduceForPartnerResponse) SetHeaders(v map[string]*string) *CreateProduceForPartnerResponse {
	s.Headers = v
	return s
}

func (s *CreateProduceForPartnerResponse) SetStatusCode(v int32) *CreateProduceForPartnerResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateProduceForPartnerResponse) SetBody(v *CreateProduceForPartnerResponseBody) *CreateProduceForPartnerResponse {
	s.Body = v
	return s
}

func (s *CreateProduceForPartnerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
