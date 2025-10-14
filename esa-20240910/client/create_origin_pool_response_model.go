// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOriginPoolResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateOriginPoolResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateOriginPoolResponse
	GetStatusCode() *int32
	SetBody(v *CreateOriginPoolResponseBody) *CreateOriginPoolResponse
	GetBody() *CreateOriginPoolResponseBody
}

type CreateOriginPoolResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateOriginPoolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateOriginPoolResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateOriginPoolResponse) GoString() string {
	return s.String()
}

func (s *CreateOriginPoolResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateOriginPoolResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateOriginPoolResponse) GetBody() *CreateOriginPoolResponseBody {
	return s.Body
}

func (s *CreateOriginPoolResponse) SetHeaders(v map[string]*string) *CreateOriginPoolResponse {
	s.Headers = v
	return s
}

func (s *CreateOriginPoolResponse) SetStatusCode(v int32) *CreateOriginPoolResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOriginPoolResponse) SetBody(v *CreateOriginPoolResponseBody) *CreateOriginPoolResponse {
	s.Body = v
	return s
}

func (s *CreateOriginPoolResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
