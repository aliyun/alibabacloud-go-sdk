// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeCheckCustomConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChangeCheckCustomConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChangeCheckCustomConfigResponse
	GetStatusCode() *int32
	SetBody(v *ChangeCheckCustomConfigResponseBody) *ChangeCheckCustomConfigResponse
	GetBody() *ChangeCheckCustomConfigResponseBody
}

type ChangeCheckCustomConfigResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeCheckCustomConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeCheckCustomConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ChangeCheckCustomConfigResponse) GoString() string {
	return s.String()
}

func (s *ChangeCheckCustomConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChangeCheckCustomConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChangeCheckCustomConfigResponse) GetBody() *ChangeCheckCustomConfigResponseBody {
	return s.Body
}

func (s *ChangeCheckCustomConfigResponse) SetHeaders(v map[string]*string) *ChangeCheckCustomConfigResponse {
	s.Headers = v
	return s
}

func (s *ChangeCheckCustomConfigResponse) SetStatusCode(v int32) *ChangeCheckCustomConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeCheckCustomConfigResponse) SetBody(v *ChangeCheckCustomConfigResponseBody) *ChangeCheckCustomConfigResponse {
	s.Body = v
	return s
}

func (s *ChangeCheckCustomConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
