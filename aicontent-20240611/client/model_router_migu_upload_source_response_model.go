// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterMiguUploadSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterMiguUploadSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterMiguUploadSourceResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterMiguUploadSourceResponseBody) *ModelRouterMiguUploadSourceResponse
	GetBody() *ModelRouterMiguUploadSourceResponseBody
}

type ModelRouterMiguUploadSourceResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterMiguUploadSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterMiguUploadSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterMiguUploadSourceResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterMiguUploadSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterMiguUploadSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterMiguUploadSourceResponse) GetBody() *ModelRouterMiguUploadSourceResponseBody {
	return s.Body
}

func (s *ModelRouterMiguUploadSourceResponse) SetHeaders(v map[string]*string) *ModelRouterMiguUploadSourceResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterMiguUploadSourceResponse) SetStatusCode(v int32) *ModelRouterMiguUploadSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterMiguUploadSourceResponse) SetBody(v *ModelRouterMiguUploadSourceResponseBody) *ModelRouterMiguUploadSourceResponse {
	s.Body = v
	return s
}

func (s *ModelRouterMiguUploadSourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
