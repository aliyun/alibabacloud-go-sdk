// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRotateTokenPlanKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RotateTokenPlanKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RotateTokenPlanKeyResponse
	GetStatusCode() *int32
	SetBody(v *RotateTokenPlanKeyResponseBody) *RotateTokenPlanKeyResponse
	GetBody() *RotateTokenPlanKeyResponseBody
}

type RotateTokenPlanKeyResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RotateTokenPlanKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RotateTokenPlanKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s RotateTokenPlanKeyResponse) GoString() string {
	return s.String()
}

func (s *RotateTokenPlanKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RotateTokenPlanKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RotateTokenPlanKeyResponse) GetBody() *RotateTokenPlanKeyResponseBody {
	return s.Body
}

func (s *RotateTokenPlanKeyResponse) SetHeaders(v map[string]*string) *RotateTokenPlanKeyResponse {
	s.Headers = v
	return s
}

func (s *RotateTokenPlanKeyResponse) SetStatusCode(v int32) *RotateTokenPlanKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *RotateTokenPlanKeyResponse) SetBody(v *RotateTokenPlanKeyResponseBody) *RotateTokenPlanKeyResponse {
	s.Body = v
	return s
}

func (s *RotateTokenPlanKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
