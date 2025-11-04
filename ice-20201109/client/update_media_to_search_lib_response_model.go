// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMediaToSearchLibResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMediaToSearchLibResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMediaToSearchLibResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMediaToSearchLibResponseBody) *UpdateMediaToSearchLibResponse
	GetBody() *UpdateMediaToSearchLibResponseBody
}

type UpdateMediaToSearchLibResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMediaToSearchLibResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMediaToSearchLibResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaToSearchLibResponse) GoString() string {
	return s.String()
}

func (s *UpdateMediaToSearchLibResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMediaToSearchLibResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMediaToSearchLibResponse) GetBody() *UpdateMediaToSearchLibResponseBody {
	return s.Body
}

func (s *UpdateMediaToSearchLibResponse) SetHeaders(v map[string]*string) *UpdateMediaToSearchLibResponse {
	s.Headers = v
	return s
}

func (s *UpdateMediaToSearchLibResponse) SetStatusCode(v int32) *UpdateMediaToSearchLibResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMediaToSearchLibResponse) SetBody(v *UpdateMediaToSearchLibResponseBody) *UpdateMediaToSearchLibResponse {
	s.Body = v
	return s
}

func (s *UpdateMediaToSearchLibResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
