// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMapPkToHidResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MapPkToHidResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MapPkToHidResponse
	GetStatusCode() *int32
	SetBody(v *MapPkToHidResponseBody) *MapPkToHidResponse
	GetBody() *MapPkToHidResponseBody
}

type MapPkToHidResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MapPkToHidResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MapPkToHidResponse) String() string {
	return dara.Prettify(s)
}

func (s MapPkToHidResponse) GoString() string {
	return s.String()
}

func (s *MapPkToHidResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MapPkToHidResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MapPkToHidResponse) GetBody() *MapPkToHidResponseBody {
	return s.Body
}

func (s *MapPkToHidResponse) SetHeaders(v map[string]*string) *MapPkToHidResponse {
	s.Headers = v
	return s
}

func (s *MapPkToHidResponse) SetStatusCode(v int32) *MapPkToHidResponse {
	s.StatusCode = &v
	return s
}

func (s *MapPkToHidResponse) SetBody(v *MapPkToHidResponseBody) *MapPkToHidResponse {
	s.Body = v
	return s
}

func (s *MapPkToHidResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
