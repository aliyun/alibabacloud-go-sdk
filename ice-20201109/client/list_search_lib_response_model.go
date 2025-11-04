// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSearchLibResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSearchLibResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSearchLibResponse
	GetStatusCode() *int32
	SetBody(v *ListSearchLibResponseBody) *ListSearchLibResponse
	GetBody() *ListSearchLibResponseBody
}

type ListSearchLibResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSearchLibResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSearchLibResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSearchLibResponse) GoString() string {
	return s.String()
}

func (s *ListSearchLibResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSearchLibResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSearchLibResponse) GetBody() *ListSearchLibResponseBody {
	return s.Body
}

func (s *ListSearchLibResponse) SetHeaders(v map[string]*string) *ListSearchLibResponse {
	s.Headers = v
	return s
}

func (s *ListSearchLibResponse) SetStatusCode(v int32) *ListSearchLibResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSearchLibResponse) SetBody(v *ListSearchLibResponseBody) *ListSearchLibResponse {
	s.Body = v
	return s
}

func (s *ListSearchLibResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
