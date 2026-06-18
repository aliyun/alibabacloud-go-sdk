// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTokenPlanKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTokenPlanKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTokenPlanKeyResponse
	GetStatusCode() *int32
	SetBody(v *CreateTokenPlanKeyResponseBody) *CreateTokenPlanKeyResponse
	GetBody() *CreateTokenPlanKeyResponseBody
}

type CreateTokenPlanKeyResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTokenPlanKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTokenPlanKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTokenPlanKeyResponse) GoString() string {
	return s.String()
}

func (s *CreateTokenPlanKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTokenPlanKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTokenPlanKeyResponse) GetBody() *CreateTokenPlanKeyResponseBody {
	return s.Body
}

func (s *CreateTokenPlanKeyResponse) SetHeaders(v map[string]*string) *CreateTokenPlanKeyResponse {
	s.Headers = v
	return s
}

func (s *CreateTokenPlanKeyResponse) SetStatusCode(v int32) *CreateTokenPlanKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTokenPlanKeyResponse) SetBody(v *CreateTokenPlanKeyResponseBody) *CreateTokenPlanKeyResponse {
	s.Body = v
	return s
}

func (s *CreateTokenPlanKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
