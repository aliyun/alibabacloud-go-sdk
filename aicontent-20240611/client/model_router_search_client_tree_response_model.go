// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterSearchClientTreeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterSearchClientTreeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterSearchClientTreeResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterSearchClientTreeResponseBody) *ModelRouterSearchClientTreeResponse
	GetBody() *ModelRouterSearchClientTreeResponseBody
}

type ModelRouterSearchClientTreeResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterSearchClientTreeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterSearchClientTreeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterSearchClientTreeResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterSearchClientTreeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterSearchClientTreeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterSearchClientTreeResponse) GetBody() *ModelRouterSearchClientTreeResponseBody {
	return s.Body
}

func (s *ModelRouterSearchClientTreeResponse) SetHeaders(v map[string]*string) *ModelRouterSearchClientTreeResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterSearchClientTreeResponse) SetStatusCode(v int32) *ModelRouterSearchClientTreeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterSearchClientTreeResponse) SetBody(v *ModelRouterSearchClientTreeResponseBody) *ModelRouterSearchClientTreeResponse {
	s.Body = v
	return s
}

func (s *ModelRouterSearchClientTreeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
