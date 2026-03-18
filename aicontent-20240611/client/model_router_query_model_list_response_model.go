// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryModelListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterQueryModelListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterQueryModelListResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterQueryModelListResponseBody) *ModelRouterQueryModelListResponse
	GetBody() *ModelRouterQueryModelListResponseBody
}

type ModelRouterQueryModelListResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterQueryModelListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterQueryModelListResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryModelListResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryModelListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterQueryModelListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterQueryModelListResponse) GetBody() *ModelRouterQueryModelListResponseBody {
	return s.Body
}

func (s *ModelRouterQueryModelListResponse) SetHeaders(v map[string]*string) *ModelRouterQueryModelListResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterQueryModelListResponse) SetStatusCode(v int32) *ModelRouterQueryModelListResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterQueryModelListResponse) SetBody(v *ModelRouterQueryModelListResponseBody) *ModelRouterQueryModelListResponse {
	s.Body = v
	return s
}

func (s *ModelRouterQueryModelListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
