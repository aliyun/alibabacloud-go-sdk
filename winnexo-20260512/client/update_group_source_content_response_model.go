// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGroupSourceContentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateGroupSourceContentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateGroupSourceContentResponse
	GetStatusCode() *int32
	SetBody(v *UpdateGroupSourceContentResponseBody) *UpdateGroupSourceContentResponse
	GetBody() *UpdateGroupSourceContentResponseBody
}

type UpdateGroupSourceContentResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGroupSourceContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGroupSourceContentResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateGroupSourceContentResponse) GoString() string {
	return s.String()
}

func (s *UpdateGroupSourceContentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateGroupSourceContentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateGroupSourceContentResponse) GetBody() *UpdateGroupSourceContentResponseBody {
	return s.Body
}

func (s *UpdateGroupSourceContentResponse) SetHeaders(v map[string]*string) *UpdateGroupSourceContentResponse {
	s.Headers = v
	return s
}

func (s *UpdateGroupSourceContentResponse) SetStatusCode(v int32) *UpdateGroupSourceContentResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGroupSourceContentResponse) SetBody(v *UpdateGroupSourceContentResponseBody) *UpdateGroupSourceContentResponse {
	s.Body = v
	return s
}

func (s *UpdateGroupSourceContentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
