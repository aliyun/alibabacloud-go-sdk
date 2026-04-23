// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryCostModelListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterQueryCostModelListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterQueryCostModelListResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterQueryCostModelListResponseBody) *ModelRouterQueryCostModelListResponse
	GetBody() *ModelRouterQueryCostModelListResponseBody
}

type ModelRouterQueryCostModelListResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterQueryCostModelListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterQueryCostModelListResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryCostModelListResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryCostModelListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterQueryCostModelListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterQueryCostModelListResponse) GetBody() *ModelRouterQueryCostModelListResponseBody {
	return s.Body
}

func (s *ModelRouterQueryCostModelListResponse) SetHeaders(v map[string]*string) *ModelRouterQueryCostModelListResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterQueryCostModelListResponse) SetStatusCode(v int32) *ModelRouterQueryCostModelListResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterQueryCostModelListResponse) SetBody(v *ModelRouterQueryCostModelListResponseBody) *ModelRouterQueryCostModelListResponse {
	s.Body = v
	return s
}

func (s *ModelRouterQueryCostModelListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
