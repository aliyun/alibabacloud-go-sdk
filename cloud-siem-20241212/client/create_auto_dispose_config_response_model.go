// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAutoDisposeConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAutoDisposeConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAutoDisposeConfigResponse
	GetStatusCode() *int32
	SetBody(v *CreateAutoDisposeConfigResponseBody) *CreateAutoDisposeConfigResponse
	GetBody() *CreateAutoDisposeConfigResponseBody
}

type CreateAutoDisposeConfigResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAutoDisposeConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAutoDisposeConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoDisposeConfigResponse) GoString() string {
	return s.String()
}

func (s *CreateAutoDisposeConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAutoDisposeConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAutoDisposeConfigResponse) GetBody() *CreateAutoDisposeConfigResponseBody {
	return s.Body
}

func (s *CreateAutoDisposeConfigResponse) SetHeaders(v map[string]*string) *CreateAutoDisposeConfigResponse {
	s.Headers = v
	return s
}

func (s *CreateAutoDisposeConfigResponse) SetStatusCode(v int32) *CreateAutoDisposeConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAutoDisposeConfigResponse) SetBody(v *CreateAutoDisposeConfigResponseBody) *CreateAutoDisposeConfigResponse {
	s.Body = v
	return s
}

func (s *CreateAutoDisposeConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
