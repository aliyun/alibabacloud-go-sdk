// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMmAppRagConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMmAppRagConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMmAppRagConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMmAppRagConfigResponseBody) *UpdateMmAppRagConfigResponse
	GetBody() *UpdateMmAppRagConfigResponseBody
}

type UpdateMmAppRagConfigResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMmAppRagConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMmAppRagConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMmAppRagConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateMmAppRagConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMmAppRagConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMmAppRagConfigResponse) GetBody() *UpdateMmAppRagConfigResponseBody {
	return s.Body
}

func (s *UpdateMmAppRagConfigResponse) SetHeaders(v map[string]*string) *UpdateMmAppRagConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateMmAppRagConfigResponse) SetStatusCode(v int32) *UpdateMmAppRagConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMmAppRagConfigResponse) SetBody(v *UpdateMmAppRagConfigResponseBody) *UpdateMmAppRagConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateMmAppRagConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
