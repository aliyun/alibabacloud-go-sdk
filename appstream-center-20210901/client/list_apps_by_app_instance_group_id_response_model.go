// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppsByAppInstanceGroupIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAppsByAppInstanceGroupIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAppsByAppInstanceGroupIdResponse
	GetStatusCode() *int32
	SetBody(v *ListAppsByAppInstanceGroupIdResponseBody) *ListAppsByAppInstanceGroupIdResponse
	GetBody() *ListAppsByAppInstanceGroupIdResponseBody
}

type ListAppsByAppInstanceGroupIdResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAppsByAppInstanceGroupIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAppsByAppInstanceGroupIdResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAppsByAppInstanceGroupIdResponse) GoString() string {
	return s.String()
}

func (s *ListAppsByAppInstanceGroupIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAppsByAppInstanceGroupIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAppsByAppInstanceGroupIdResponse) GetBody() *ListAppsByAppInstanceGroupIdResponseBody {
	return s.Body
}

func (s *ListAppsByAppInstanceGroupIdResponse) SetHeaders(v map[string]*string) *ListAppsByAppInstanceGroupIdResponse {
	s.Headers = v
	return s
}

func (s *ListAppsByAppInstanceGroupIdResponse) SetStatusCode(v int32) *ListAppsByAppInstanceGroupIdResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAppsByAppInstanceGroupIdResponse) SetBody(v *ListAppsByAppInstanceGroupIdResponseBody) *ListAppsByAppInstanceGroupIdResponse {
	s.Body = v
	return s
}

func (s *ListAppsByAppInstanceGroupIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
