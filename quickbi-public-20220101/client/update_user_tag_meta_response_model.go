// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserTagMetaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateUserTagMetaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateUserTagMetaResponse
	GetStatusCode() *int32
	SetBody(v *UpdateUserTagMetaResponseBody) *UpdateUserTagMetaResponse
	GetBody() *UpdateUserTagMetaResponseBody
}

type UpdateUserTagMetaResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateUserTagMetaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateUserTagMetaResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserTagMetaResponse) GoString() string {
	return s.String()
}

func (s *UpdateUserTagMetaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateUserTagMetaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateUserTagMetaResponse) GetBody() *UpdateUserTagMetaResponseBody {
	return s.Body
}

func (s *UpdateUserTagMetaResponse) SetHeaders(v map[string]*string) *UpdateUserTagMetaResponse {
	s.Headers = v
	return s
}

func (s *UpdateUserTagMetaResponse) SetStatusCode(v int32) *UpdateUserTagMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUserTagMetaResponse) SetBody(v *UpdateUserTagMetaResponseBody) *UpdateUserTagMetaResponse {
	s.Body = v
	return s
}

func (s *UpdateUserTagMetaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
