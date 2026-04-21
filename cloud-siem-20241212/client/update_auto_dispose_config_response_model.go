// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAutoDisposeConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAutoDisposeConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAutoDisposeConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAutoDisposeConfigResponseBody) *UpdateAutoDisposeConfigResponse
	GetBody() *UpdateAutoDisposeConfigResponseBody
}

type UpdateAutoDisposeConfigResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAutoDisposeConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAutoDisposeConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAutoDisposeConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateAutoDisposeConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAutoDisposeConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAutoDisposeConfigResponse) GetBody() *UpdateAutoDisposeConfigResponseBody {
	return s.Body
}

func (s *UpdateAutoDisposeConfigResponse) SetHeaders(v map[string]*string) *UpdateAutoDisposeConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateAutoDisposeConfigResponse) SetStatusCode(v int32) *UpdateAutoDisposeConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAutoDisposeConfigResponse) SetBody(v *UpdateAutoDisposeConfigResponseBody) *UpdateAutoDisposeConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateAutoDisposeConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
