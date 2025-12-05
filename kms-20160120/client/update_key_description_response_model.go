// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKeyDescriptionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateKeyDescriptionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateKeyDescriptionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateKeyDescriptionResponseBody) *UpdateKeyDescriptionResponse
	GetBody() *UpdateKeyDescriptionResponseBody
}

type UpdateKeyDescriptionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateKeyDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateKeyDescriptionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateKeyDescriptionResponse) GoString() string {
	return s.String()
}

func (s *UpdateKeyDescriptionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateKeyDescriptionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateKeyDescriptionResponse) GetBody() *UpdateKeyDescriptionResponseBody {
	return s.Body
}

func (s *UpdateKeyDescriptionResponse) SetHeaders(v map[string]*string) *UpdateKeyDescriptionResponse {
	s.Headers = v
	return s
}

func (s *UpdateKeyDescriptionResponse) SetStatusCode(v int32) *UpdateKeyDescriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateKeyDescriptionResponse) SetBody(v *UpdateKeyDescriptionResponseBody) *UpdateKeyDescriptionResponse {
	s.Body = v
	return s
}

func (s *UpdateKeyDescriptionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
