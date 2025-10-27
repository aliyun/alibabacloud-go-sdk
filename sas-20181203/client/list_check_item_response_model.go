// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCheckItemResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCheckItemResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCheckItemResponse
	GetStatusCode() *int32
	SetBody(v *ListCheckItemResponseBody) *ListCheckItemResponse
	GetBody() *ListCheckItemResponseBody
}

type ListCheckItemResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCheckItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCheckItemResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCheckItemResponse) GoString() string {
	return s.String()
}

func (s *ListCheckItemResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCheckItemResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCheckItemResponse) GetBody() *ListCheckItemResponseBody {
	return s.Body
}

func (s *ListCheckItemResponse) SetHeaders(v map[string]*string) *ListCheckItemResponse {
	s.Headers = v
	return s
}

func (s *ListCheckItemResponse) SetStatusCode(v int32) *ListCheckItemResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCheckItemResponse) SetBody(v *ListCheckItemResponseBody) *ListCheckItemResponse {
	s.Body = v
	return s
}

func (s *ListCheckItemResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
