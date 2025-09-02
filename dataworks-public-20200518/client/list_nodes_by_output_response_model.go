// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodesByOutputResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListNodesByOutputResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListNodesByOutputResponse
	GetStatusCode() *int32
	SetBody(v *ListNodesByOutputResponseBody) *ListNodesByOutputResponse
	GetBody() *ListNodesByOutputResponseBody
}

type ListNodesByOutputResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNodesByOutputResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNodesByOutputResponse) String() string {
	return dara.Prettify(s)
}

func (s ListNodesByOutputResponse) GoString() string {
	return s.String()
}

func (s *ListNodesByOutputResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListNodesByOutputResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListNodesByOutputResponse) GetBody() *ListNodesByOutputResponseBody {
	return s.Body
}

func (s *ListNodesByOutputResponse) SetHeaders(v map[string]*string) *ListNodesByOutputResponse {
	s.Headers = v
	return s
}

func (s *ListNodesByOutputResponse) SetStatusCode(v int32) *ListNodesByOutputResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNodesByOutputResponse) SetBody(v *ListNodesByOutputResponseBody) *ListNodesByOutputResponse {
	s.Body = v
	return s
}

func (s *ListNodesByOutputResponse) Validate() error {
	return dara.Validate(s)
}
