// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWarmPoolResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateWarmPoolResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateWarmPoolResponse
	GetStatusCode() *int32
	SetBody(v *UpdateWarmPoolResponseBody) *UpdateWarmPoolResponse
	GetBody() *UpdateWarmPoolResponseBody
}

type UpdateWarmPoolResponse struct {
	Headers    map[string]*string          `json:"headers" xml:"headers"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateWarmPoolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWarmPoolResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateWarmPoolResponse) GoString() string {
	return s.String()
}

func (s *UpdateWarmPoolResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateWarmPoolResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateWarmPoolResponse) GetBody() *UpdateWarmPoolResponseBody {
	return s.Body
}

func (s *UpdateWarmPoolResponse) SetHeaders(v map[string]*string) *UpdateWarmPoolResponse {
	s.Headers = v
	return s
}

func (s *UpdateWarmPoolResponse) SetStatusCode(v int32) *UpdateWarmPoolResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWarmPoolResponse) SetBody(v *UpdateWarmPoolResponseBody) *UpdateWarmPoolResponse {
	s.Body = v
	return s
}

func (s *UpdateWarmPoolResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
