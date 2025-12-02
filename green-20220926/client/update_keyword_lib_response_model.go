// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKeywordLibResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateKeywordLibResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateKeywordLibResponse
	GetStatusCode() *int32
	SetBody(v *UpdateKeywordLibResponseBody) *UpdateKeywordLibResponse
	GetBody() *UpdateKeywordLibResponseBody
}

type UpdateKeywordLibResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateKeywordLibResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateKeywordLibResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateKeywordLibResponse) GoString() string {
	return s.String()
}

func (s *UpdateKeywordLibResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateKeywordLibResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateKeywordLibResponse) GetBody() *UpdateKeywordLibResponseBody {
	return s.Body
}

func (s *UpdateKeywordLibResponse) SetHeaders(v map[string]*string) *UpdateKeywordLibResponse {
	s.Headers = v
	return s
}

func (s *UpdateKeywordLibResponse) SetStatusCode(v int32) *UpdateKeywordLibResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateKeywordLibResponse) SetBody(v *UpdateKeywordLibResponseBody) *UpdateKeywordLibResponse {
	s.Body = v
	return s
}

func (s *UpdateKeywordLibResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
