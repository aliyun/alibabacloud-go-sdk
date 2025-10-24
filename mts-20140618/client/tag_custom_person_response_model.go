// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTagCustomPersonResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TagCustomPersonResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TagCustomPersonResponse
	GetStatusCode() *int32
	SetBody(v *TagCustomPersonResponseBody) *TagCustomPersonResponse
	GetBody() *TagCustomPersonResponseBody
}

type TagCustomPersonResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TagCustomPersonResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TagCustomPersonResponse) String() string {
	return dara.Prettify(s)
}

func (s TagCustomPersonResponse) GoString() string {
	return s.String()
}

func (s *TagCustomPersonResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TagCustomPersonResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TagCustomPersonResponse) GetBody() *TagCustomPersonResponseBody {
	return s.Body
}

func (s *TagCustomPersonResponse) SetHeaders(v map[string]*string) *TagCustomPersonResponse {
	s.Headers = v
	return s
}

func (s *TagCustomPersonResponse) SetStatusCode(v int32) *TagCustomPersonResponse {
	s.StatusCode = &v
	return s
}

func (s *TagCustomPersonResponse) SetBody(v *TagCustomPersonResponseBody) *TagCustomPersonResponse {
	s.Body = v
	return s
}

func (s *TagCustomPersonResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
