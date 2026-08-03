// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAndPayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAndPayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAndPayResponse
	GetStatusCode() *int32
	SetBody(v *CreateAndPayResponseBody) *CreateAndPayResponse
	GetBody() *CreateAndPayResponseBody
}

type CreateAndPayResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAndPayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAndPayResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAndPayResponse) GoString() string {
	return s.String()
}

func (s *CreateAndPayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAndPayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAndPayResponse) GetBody() *CreateAndPayResponseBody {
	return s.Body
}

func (s *CreateAndPayResponse) SetHeaders(v map[string]*string) *CreateAndPayResponse {
	s.Headers = v
	return s
}

func (s *CreateAndPayResponse) SetStatusCode(v int32) *CreateAndPayResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAndPayResponse) SetBody(v *CreateAndPayResponseBody) *CreateAndPayResponse {
	s.Body = v
	return s
}

func (s *CreateAndPayResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
