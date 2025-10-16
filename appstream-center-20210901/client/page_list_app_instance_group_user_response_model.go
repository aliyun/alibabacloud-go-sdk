// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPageListAppInstanceGroupUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PageListAppInstanceGroupUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PageListAppInstanceGroupUserResponse
	GetStatusCode() *int32
	SetBody(v *PageListAppInstanceGroupUserResponseBody) *PageListAppInstanceGroupUserResponse
	GetBody() *PageListAppInstanceGroupUserResponseBody
}

type PageListAppInstanceGroupUserResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PageListAppInstanceGroupUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PageListAppInstanceGroupUserResponse) String() string {
	return dara.Prettify(s)
}

func (s PageListAppInstanceGroupUserResponse) GoString() string {
	return s.String()
}

func (s *PageListAppInstanceGroupUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PageListAppInstanceGroupUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PageListAppInstanceGroupUserResponse) GetBody() *PageListAppInstanceGroupUserResponseBody {
	return s.Body
}

func (s *PageListAppInstanceGroupUserResponse) SetHeaders(v map[string]*string) *PageListAppInstanceGroupUserResponse {
	s.Headers = v
	return s
}

func (s *PageListAppInstanceGroupUserResponse) SetStatusCode(v int32) *PageListAppInstanceGroupUserResponse {
	s.StatusCode = &v
	return s
}

func (s *PageListAppInstanceGroupUserResponse) SetBody(v *PageListAppInstanceGroupUserResponseBody) *PageListAppInstanceGroupUserResponse {
	s.Body = v
	return s
}

func (s *PageListAppInstanceGroupUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
