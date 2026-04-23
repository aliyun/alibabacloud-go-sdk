// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryClientDiscountLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterQueryClientDiscountLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterQueryClientDiscountLogsResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterQueryClientDiscountLogsResponseBody) *ModelRouterQueryClientDiscountLogsResponse
	GetBody() *ModelRouterQueryClientDiscountLogsResponseBody
}

type ModelRouterQueryClientDiscountLogsResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterQueryClientDiscountLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterQueryClientDiscountLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryClientDiscountLogsResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryClientDiscountLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterQueryClientDiscountLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterQueryClientDiscountLogsResponse) GetBody() *ModelRouterQueryClientDiscountLogsResponseBody {
	return s.Body
}

func (s *ModelRouterQueryClientDiscountLogsResponse) SetHeaders(v map[string]*string) *ModelRouterQueryClientDiscountLogsResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterQueryClientDiscountLogsResponse) SetStatusCode(v int32) *ModelRouterQueryClientDiscountLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterQueryClientDiscountLogsResponse) SetBody(v *ModelRouterQueryClientDiscountLogsResponseBody) *ModelRouterQueryClientDiscountLogsResponse {
	s.Body = v
	return s
}

func (s *ModelRouterQueryClientDiscountLogsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
