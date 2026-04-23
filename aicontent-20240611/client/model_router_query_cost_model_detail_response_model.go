// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryCostModelDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterQueryCostModelDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterQueryCostModelDetailResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterQueryCostModelDetailResponseBody) *ModelRouterQueryCostModelDetailResponse
	GetBody() *ModelRouterQueryCostModelDetailResponseBody
}

type ModelRouterQueryCostModelDetailResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterQueryCostModelDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterQueryCostModelDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryCostModelDetailResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryCostModelDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterQueryCostModelDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterQueryCostModelDetailResponse) GetBody() *ModelRouterQueryCostModelDetailResponseBody {
	return s.Body
}

func (s *ModelRouterQueryCostModelDetailResponse) SetHeaders(v map[string]*string) *ModelRouterQueryCostModelDetailResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterQueryCostModelDetailResponse) SetStatusCode(v int32) *ModelRouterQueryCostModelDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterQueryCostModelDetailResponse) SetBody(v *ModelRouterQueryCostModelDetailResponseBody) *ModelRouterQueryCostModelDetailResponse {
	s.Body = v
	return s
}

func (s *ModelRouterQueryCostModelDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
