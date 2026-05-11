// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceRolloutResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateServiceRolloutResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateServiceRolloutResponse
	GetStatusCode() *int32
	SetBody(v *UpdateServiceRolloutResponseBody) *UpdateServiceRolloutResponse
	GetBody() *UpdateServiceRolloutResponseBody
}

type UpdateServiceRolloutResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateServiceRolloutResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateServiceRolloutResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceRolloutResponse) GoString() string {
	return s.String()
}

func (s *UpdateServiceRolloutResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateServiceRolloutResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateServiceRolloutResponse) GetBody() *UpdateServiceRolloutResponseBody {
	return s.Body
}

func (s *UpdateServiceRolloutResponse) SetHeaders(v map[string]*string) *UpdateServiceRolloutResponse {
	s.Headers = v
	return s
}

func (s *UpdateServiceRolloutResponse) SetStatusCode(v int32) *UpdateServiceRolloutResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateServiceRolloutResponse) SetBody(v *UpdateServiceRolloutResponseBody) *UpdateServiceRolloutResponse {
	s.Body = v
	return s
}

func (s *UpdateServiceRolloutResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
