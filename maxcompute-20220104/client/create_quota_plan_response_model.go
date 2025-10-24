// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateQuotaPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateQuotaPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateQuotaPlanResponse
	GetStatusCode() *int32
	SetBody(v *CreateQuotaPlanResponseBody) *CreateQuotaPlanResponse
	GetBody() *CreateQuotaPlanResponseBody
}

type CreateQuotaPlanResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateQuotaPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateQuotaPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateQuotaPlanResponse) GoString() string {
	return s.String()
}

func (s *CreateQuotaPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateQuotaPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateQuotaPlanResponse) GetBody() *CreateQuotaPlanResponseBody {
	return s.Body
}

func (s *CreateQuotaPlanResponse) SetHeaders(v map[string]*string) *CreateQuotaPlanResponse {
	s.Headers = v
	return s
}

func (s *CreateQuotaPlanResponse) SetStatusCode(v int32) *CreateQuotaPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateQuotaPlanResponse) SetBody(v *CreateQuotaPlanResponseBody) *CreateQuotaPlanResponse {
	s.Body = v
	return s
}

func (s *CreateQuotaPlanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
