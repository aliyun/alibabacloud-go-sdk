// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryClientTreeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterQueryClientTreeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterQueryClientTreeResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterQueryClientTreeResponseBody) *ModelRouterQueryClientTreeResponse
	GetBody() *ModelRouterQueryClientTreeResponseBody
}

type ModelRouterQueryClientTreeResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterQueryClientTreeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterQueryClientTreeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryClientTreeResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryClientTreeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterQueryClientTreeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterQueryClientTreeResponse) GetBody() *ModelRouterQueryClientTreeResponseBody {
	return s.Body
}

func (s *ModelRouterQueryClientTreeResponse) SetHeaders(v map[string]*string) *ModelRouterQueryClientTreeResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterQueryClientTreeResponse) SetStatusCode(v int32) *ModelRouterQueryClientTreeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterQueryClientTreeResponse) SetBody(v *ModelRouterQueryClientTreeResponseBody) *ModelRouterQueryClientTreeResponse {
	s.Body = v
	return s
}

func (s *ModelRouterQueryClientTreeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
