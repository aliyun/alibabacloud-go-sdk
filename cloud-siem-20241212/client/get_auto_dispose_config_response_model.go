// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAutoDisposeConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAutoDisposeConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAutoDisposeConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetAutoDisposeConfigResponseBody) *GetAutoDisposeConfigResponse
	GetBody() *GetAutoDisposeConfigResponseBody
}

type GetAutoDisposeConfigResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAutoDisposeConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAutoDisposeConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAutoDisposeConfigResponse) GoString() string {
	return s.String()
}

func (s *GetAutoDisposeConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAutoDisposeConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAutoDisposeConfigResponse) GetBody() *GetAutoDisposeConfigResponseBody {
	return s.Body
}

func (s *GetAutoDisposeConfigResponse) SetHeaders(v map[string]*string) *GetAutoDisposeConfigResponse {
	s.Headers = v
	return s
}

func (s *GetAutoDisposeConfigResponse) SetStatusCode(v int32) *GetAutoDisposeConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAutoDisposeConfigResponse) SetBody(v *GetAutoDisposeConfigResponseBody) *GetAutoDisposeConfigResponse {
	s.Body = v
	return s
}

func (s *GetAutoDisposeConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
