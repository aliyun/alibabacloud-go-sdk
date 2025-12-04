// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSingleCallByVideoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SingleCallByVideoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SingleCallByVideoResponse
	GetStatusCode() *int32
	SetBody(v *SingleCallByVideoResponseBody) *SingleCallByVideoResponse
	GetBody() *SingleCallByVideoResponseBody
}

type SingleCallByVideoResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SingleCallByVideoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SingleCallByVideoResponse) String() string {
	return dara.Prettify(s)
}

func (s SingleCallByVideoResponse) GoString() string {
	return s.String()
}

func (s *SingleCallByVideoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SingleCallByVideoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SingleCallByVideoResponse) GetBody() *SingleCallByVideoResponseBody {
	return s.Body
}

func (s *SingleCallByVideoResponse) SetHeaders(v map[string]*string) *SingleCallByVideoResponse {
	s.Headers = v
	return s
}

func (s *SingleCallByVideoResponse) SetStatusCode(v int32) *SingleCallByVideoResponse {
	s.StatusCode = &v
	return s
}

func (s *SingleCallByVideoResponse) SetBody(v *SingleCallByVideoResponseBody) *SingleCallByVideoResponse {
	s.Body = v
	return s
}

func (s *SingleCallByVideoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
