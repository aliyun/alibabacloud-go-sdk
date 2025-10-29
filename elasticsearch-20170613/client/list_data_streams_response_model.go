// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataStreamsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataStreamsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataStreamsResponse
	GetStatusCode() *int32
	SetBody(v *ListDataStreamsResponseBody) *ListDataStreamsResponse
	GetBody() *ListDataStreamsResponseBody
}

type ListDataStreamsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataStreamsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataStreamsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataStreamsResponse) GoString() string {
	return s.String()
}

func (s *ListDataStreamsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataStreamsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataStreamsResponse) GetBody() *ListDataStreamsResponseBody {
	return s.Body
}

func (s *ListDataStreamsResponse) SetHeaders(v map[string]*string) *ListDataStreamsResponse {
	s.Headers = v
	return s
}

func (s *ListDataStreamsResponse) SetStatusCode(v int32) *ListDataStreamsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataStreamsResponse) SetBody(v *ListDataStreamsResponseBody) *ListDataStreamsResponse {
	s.Body = v
	return s
}

func (s *ListDataStreamsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
