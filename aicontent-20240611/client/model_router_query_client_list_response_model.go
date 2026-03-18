// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryClientListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterQueryClientListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterQueryClientListResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterQueryClientListResponseBody) *ModelRouterQueryClientListResponse
	GetBody() *ModelRouterQueryClientListResponseBody
}

type ModelRouterQueryClientListResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterQueryClientListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterQueryClientListResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryClientListResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryClientListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterQueryClientListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterQueryClientListResponse) GetBody() *ModelRouterQueryClientListResponseBody {
	return s.Body
}

func (s *ModelRouterQueryClientListResponse) SetHeaders(v map[string]*string) *ModelRouterQueryClientListResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterQueryClientListResponse) SetStatusCode(v int32) *ModelRouterQueryClientListResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterQueryClientListResponse) SetBody(v *ModelRouterQueryClientListResponseBody) *ModelRouterQueryClientListResponse {
	s.Body = v
	return s
}

func (s *ModelRouterQueryClientListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
