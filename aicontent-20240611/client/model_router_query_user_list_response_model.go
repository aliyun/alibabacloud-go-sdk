// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryUserListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterQueryUserListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterQueryUserListResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterQueryUserListResponseBody) *ModelRouterQueryUserListResponse
	GetBody() *ModelRouterQueryUserListResponseBody
}

type ModelRouterQueryUserListResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterQueryUserListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterQueryUserListResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryUserListResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryUserListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterQueryUserListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterQueryUserListResponse) GetBody() *ModelRouterQueryUserListResponseBody {
	return s.Body
}

func (s *ModelRouterQueryUserListResponse) SetHeaders(v map[string]*string) *ModelRouterQueryUserListResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterQueryUserListResponse) SetStatusCode(v int32) *ModelRouterQueryUserListResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterQueryUserListResponse) SetBody(v *ModelRouterQueryUserListResponseBody) *ModelRouterQueryUserListResponse {
	s.Body = v
	return s
}

func (s *ModelRouterQueryUserListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
