// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMessageLanguageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMessageLanguageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMessageLanguageResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMessageLanguageResponseBody) *UpdateMessageLanguageResponse
	GetBody() *UpdateMessageLanguageResponseBody
}

type UpdateMessageLanguageResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMessageLanguageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMessageLanguageResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMessageLanguageResponse) GoString() string {
	return s.String()
}

func (s *UpdateMessageLanguageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMessageLanguageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMessageLanguageResponse) GetBody() *UpdateMessageLanguageResponseBody {
	return s.Body
}

func (s *UpdateMessageLanguageResponse) SetHeaders(v map[string]*string) *UpdateMessageLanguageResponse {
	s.Headers = v
	return s
}

func (s *UpdateMessageLanguageResponse) SetStatusCode(v int32) *UpdateMessageLanguageResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMessageLanguageResponse) SetBody(v *UpdateMessageLanguageResponseBody) *UpdateMessageLanguageResponse {
	s.Body = v
	return s
}

func (s *UpdateMessageLanguageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
