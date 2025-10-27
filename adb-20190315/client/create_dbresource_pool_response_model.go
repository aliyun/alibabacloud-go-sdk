// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBResourcePoolResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDBResourcePoolResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDBResourcePoolResponse
	GetStatusCode() *int32
	SetBody(v *CreateDBResourcePoolResponseBody) *CreateDBResourcePoolResponse
	GetBody() *CreateDBResourcePoolResponseBody
}

type CreateDBResourcePoolResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDBResourcePoolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDBResourcePoolResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDBResourcePoolResponse) GoString() string {
	return s.String()
}

func (s *CreateDBResourcePoolResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDBResourcePoolResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDBResourcePoolResponse) GetBody() *CreateDBResourcePoolResponseBody {
	return s.Body
}

func (s *CreateDBResourcePoolResponse) SetHeaders(v map[string]*string) *CreateDBResourcePoolResponse {
	s.Headers = v
	return s
}

func (s *CreateDBResourcePoolResponse) SetStatusCode(v int32) *CreateDBResourcePoolResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDBResourcePoolResponse) SetBody(v *CreateDBResourcePoolResponseBody) *CreateDBResourcePoolResponse {
	s.Body = v
	return s
}

func (s *CreateDBResourcePoolResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
