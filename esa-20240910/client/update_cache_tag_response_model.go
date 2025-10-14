// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCacheTagResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCacheTagResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCacheTagResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCacheTagResponseBody) *UpdateCacheTagResponse
	GetBody() *UpdateCacheTagResponseBody
}

type UpdateCacheTagResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCacheTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCacheTagResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCacheTagResponse) GoString() string {
	return s.String()
}

func (s *UpdateCacheTagResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCacheTagResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCacheTagResponse) GetBody() *UpdateCacheTagResponseBody {
	return s.Body
}

func (s *UpdateCacheTagResponse) SetHeaders(v map[string]*string) *UpdateCacheTagResponse {
	s.Headers = v
	return s
}

func (s *UpdateCacheTagResponse) SetStatusCode(v int32) *UpdateCacheTagResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCacheTagResponse) SetBody(v *UpdateCacheTagResponseBody) *UpdateCacheTagResponse {
	s.Body = v
	return s
}

func (s *UpdateCacheTagResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
