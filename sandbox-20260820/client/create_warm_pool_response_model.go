// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWarmPoolResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateWarmPoolResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateWarmPoolResponse
	GetStatusCode() *int32
	SetBody(v *CreateWarmPoolResponseBody) *CreateWarmPoolResponse
	GetBody() *CreateWarmPoolResponseBody
}

type CreateWarmPoolResponse struct {
	Headers    map[string]*string          `json:"headers" xml:"headers"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateWarmPoolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWarmPoolResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateWarmPoolResponse) GoString() string {
	return s.String()
}

func (s *CreateWarmPoolResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateWarmPoolResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateWarmPoolResponse) GetBody() *CreateWarmPoolResponseBody {
	return s.Body
}

func (s *CreateWarmPoolResponse) SetHeaders(v map[string]*string) *CreateWarmPoolResponse {
	s.Headers = v
	return s
}

func (s *CreateWarmPoolResponse) SetStatusCode(v int32) *CreateWarmPoolResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWarmPoolResponse) SetBody(v *CreateWarmPoolResponseBody) *CreateWarmPoolResponse {
	s.Body = v
	return s
}

func (s *CreateWarmPoolResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
