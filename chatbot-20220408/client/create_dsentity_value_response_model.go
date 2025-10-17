// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDSEntityValueResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDSEntityValueResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDSEntityValueResponse
	GetStatusCode() *int32
	SetBody(v *CreateDSEntityValueResponseBody) *CreateDSEntityValueResponse
	GetBody() *CreateDSEntityValueResponseBody
}

type CreateDSEntityValueResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDSEntityValueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDSEntityValueResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDSEntityValueResponse) GoString() string {
	return s.String()
}

func (s *CreateDSEntityValueResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDSEntityValueResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDSEntityValueResponse) GetBody() *CreateDSEntityValueResponseBody {
	return s.Body
}

func (s *CreateDSEntityValueResponse) SetHeaders(v map[string]*string) *CreateDSEntityValueResponse {
	s.Headers = v
	return s
}

func (s *CreateDSEntityValueResponse) SetStatusCode(v int32) *CreateDSEntityValueResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDSEntityValueResponse) SetBody(v *CreateDSEntityValueResponseBody) *CreateDSEntityValueResponse {
	s.Body = v
	return s
}

func (s *CreateDSEntityValueResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
