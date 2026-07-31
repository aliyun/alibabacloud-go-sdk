// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryModelGroupListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterQueryModelGroupListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterQueryModelGroupListResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterQueryModelGroupListResponseBody) *ModelRouterQueryModelGroupListResponse
	GetBody() *ModelRouterQueryModelGroupListResponseBody
}

type ModelRouterQueryModelGroupListResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterQueryModelGroupListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterQueryModelGroupListResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryModelGroupListResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryModelGroupListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterQueryModelGroupListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterQueryModelGroupListResponse) GetBody() *ModelRouterQueryModelGroupListResponseBody {
	return s.Body
}

func (s *ModelRouterQueryModelGroupListResponse) SetHeaders(v map[string]*string) *ModelRouterQueryModelGroupListResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterQueryModelGroupListResponse) SetStatusCode(v int32) *ModelRouterQueryModelGroupListResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterQueryModelGroupListResponse) SetBody(v *ModelRouterQueryModelGroupListResponseBody) *ModelRouterQueryModelGroupListResponse {
	s.Body = v
	return s
}

func (s *ModelRouterQueryModelGroupListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
