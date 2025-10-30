// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDevObjectDependencyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDevObjectDependencyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDevObjectDependencyResponse
	GetStatusCode() *int32
	SetBody(v *GetDevObjectDependencyResponseBody) *GetDevObjectDependencyResponse
	GetBody() *GetDevObjectDependencyResponseBody
}

type GetDevObjectDependencyResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDevObjectDependencyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDevObjectDependencyResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDevObjectDependencyResponse) GoString() string {
	return s.String()
}

func (s *GetDevObjectDependencyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDevObjectDependencyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDevObjectDependencyResponse) GetBody() *GetDevObjectDependencyResponseBody {
	return s.Body
}

func (s *GetDevObjectDependencyResponse) SetHeaders(v map[string]*string) *GetDevObjectDependencyResponse {
	s.Headers = v
	return s
}

func (s *GetDevObjectDependencyResponse) SetStatusCode(v int32) *GetDevObjectDependencyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDevObjectDependencyResponse) SetBody(v *GetDevObjectDependencyResponseBody) *GetDevObjectDependencyResponse {
	s.Body = v
	return s
}

func (s *GetDevObjectDependencyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
