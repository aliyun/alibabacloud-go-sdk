// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMetaEntityDefResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMetaEntityDefResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMetaEntityDefResponse
	GetStatusCode() *int32
	SetBody(v *CreateMetaEntityDefResponseBody) *CreateMetaEntityDefResponse
	GetBody() *CreateMetaEntityDefResponseBody
}

type CreateMetaEntityDefResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMetaEntityDefResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMetaEntityDefResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMetaEntityDefResponse) GoString() string {
	return s.String()
}

func (s *CreateMetaEntityDefResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMetaEntityDefResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMetaEntityDefResponse) GetBody() *CreateMetaEntityDefResponseBody {
	return s.Body
}

func (s *CreateMetaEntityDefResponse) SetHeaders(v map[string]*string) *CreateMetaEntityDefResponse {
	s.Headers = v
	return s
}

func (s *CreateMetaEntityDefResponse) SetStatusCode(v int32) *CreateMetaEntityDefResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMetaEntityDefResponse) SetBody(v *CreateMetaEntityDefResponseBody) *CreateMetaEntityDefResponse {
	s.Body = v
	return s
}

func (s *CreateMetaEntityDefResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
