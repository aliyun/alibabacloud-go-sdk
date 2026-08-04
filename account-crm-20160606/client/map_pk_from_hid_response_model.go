// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMapPkFromHidResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MapPkFromHidResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MapPkFromHidResponse
	GetStatusCode() *int32
	SetBody(v *MapPkFromHidResponseBody) *MapPkFromHidResponse
	GetBody() *MapPkFromHidResponseBody
}

type MapPkFromHidResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MapPkFromHidResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MapPkFromHidResponse) String() string {
	return dara.Prettify(s)
}

func (s MapPkFromHidResponse) GoString() string {
	return s.String()
}

func (s *MapPkFromHidResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MapPkFromHidResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MapPkFromHidResponse) GetBody() *MapPkFromHidResponseBody {
	return s.Body
}

func (s *MapPkFromHidResponse) SetHeaders(v map[string]*string) *MapPkFromHidResponse {
	s.Headers = v
	return s
}

func (s *MapPkFromHidResponse) SetStatusCode(v int32) *MapPkFromHidResponse {
	s.StatusCode = &v
	return s
}

func (s *MapPkFromHidResponse) SetBody(v *MapPkFromHidResponseBody) *MapPkFromHidResponse {
	s.Body = v
	return s
}

func (s *MapPkFromHidResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
