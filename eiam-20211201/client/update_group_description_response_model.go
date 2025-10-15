// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGroupDescriptionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateGroupDescriptionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateGroupDescriptionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateGroupDescriptionResponseBody) *UpdateGroupDescriptionResponse
	GetBody() *UpdateGroupDescriptionResponseBody
}

type UpdateGroupDescriptionResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGroupDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGroupDescriptionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateGroupDescriptionResponse) GoString() string {
	return s.String()
}

func (s *UpdateGroupDescriptionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateGroupDescriptionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateGroupDescriptionResponse) GetBody() *UpdateGroupDescriptionResponseBody {
	return s.Body
}

func (s *UpdateGroupDescriptionResponse) SetHeaders(v map[string]*string) *UpdateGroupDescriptionResponse {
	s.Headers = v
	return s
}

func (s *UpdateGroupDescriptionResponse) SetStatusCode(v int32) *UpdateGroupDescriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGroupDescriptionResponse) SetBody(v *UpdateGroupDescriptionResponseBody) *UpdateGroupDescriptionResponse {
	s.Body = v
	return s
}

func (s *UpdateGroupDescriptionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
