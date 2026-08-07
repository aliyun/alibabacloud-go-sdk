// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloneVoicesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCloneVoicesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCloneVoicesResponse
	GetStatusCode() *int32
	SetBody(v *ListCloneVoicesResponseBody) *ListCloneVoicesResponse
	GetBody() *ListCloneVoicesResponseBody
}

type ListCloneVoicesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCloneVoicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCloneVoicesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCloneVoicesResponse) GoString() string {
	return s.String()
}

func (s *ListCloneVoicesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCloneVoicesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCloneVoicesResponse) GetBody() *ListCloneVoicesResponseBody {
	return s.Body
}

func (s *ListCloneVoicesResponse) SetHeaders(v map[string]*string) *ListCloneVoicesResponse {
	s.Headers = v
	return s
}

func (s *ListCloneVoicesResponse) SetStatusCode(v int32) *ListCloneVoicesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCloneVoicesResponse) SetBody(v *ListCloneVoicesResponseBody) *ListCloneVoicesResponse {
	s.Body = v
	return s
}

func (s *ListCloneVoicesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
