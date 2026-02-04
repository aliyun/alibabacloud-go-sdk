// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomFieldResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCustomFieldResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCustomFieldResponse
	GetStatusCode() *int32
	SetBody(v *CreateCustomFieldResponseBody) *CreateCustomFieldResponse
	GetBody() *CreateCustomFieldResponseBody
}

type CreateCustomFieldResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCustomFieldResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCustomFieldResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomFieldResponse) GoString() string {
	return s.String()
}

func (s *CreateCustomFieldResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCustomFieldResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCustomFieldResponse) GetBody() *CreateCustomFieldResponseBody {
	return s.Body
}

func (s *CreateCustomFieldResponse) SetHeaders(v map[string]*string) *CreateCustomFieldResponse {
	s.Headers = v
	return s
}

func (s *CreateCustomFieldResponse) SetStatusCode(v int32) *CreateCustomFieldResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCustomFieldResponse) SetBody(v *CreateCustomFieldResponseBody) *CreateCustomFieldResponse {
	s.Body = v
	return s
}

func (s *CreateCustomFieldResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
