// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLogStoreConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateLogStoreConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateLogStoreConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdateLogStoreConfigResponseBody) *UpdateLogStoreConfigResponse
	GetBody() *UpdateLogStoreConfigResponseBody
}

type UpdateLogStoreConfigResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLogStoreConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLogStoreConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateLogStoreConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateLogStoreConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateLogStoreConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateLogStoreConfigResponse) GetBody() *UpdateLogStoreConfigResponseBody {
	return s.Body
}

func (s *UpdateLogStoreConfigResponse) SetHeaders(v map[string]*string) *UpdateLogStoreConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateLogStoreConfigResponse) SetStatusCode(v int32) *UpdateLogStoreConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLogStoreConfigResponse) SetBody(v *UpdateLogStoreConfigResponseBody) *UpdateLogStoreConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateLogStoreConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
