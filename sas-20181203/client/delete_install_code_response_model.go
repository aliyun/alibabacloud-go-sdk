// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstallCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteInstallCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteInstallCodeResponse
	GetStatusCode() *int32
	SetBody(v *DeleteInstallCodeResponseBody) *DeleteInstallCodeResponse
	GetBody() *DeleteInstallCodeResponseBody
}

type DeleteInstallCodeResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteInstallCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteInstallCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstallCodeResponse) GoString() string {
	return s.String()
}

func (s *DeleteInstallCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteInstallCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteInstallCodeResponse) GetBody() *DeleteInstallCodeResponseBody {
	return s.Body
}

func (s *DeleteInstallCodeResponse) SetHeaders(v map[string]*string) *DeleteInstallCodeResponse {
	s.Headers = v
	return s
}

func (s *DeleteInstallCodeResponse) SetStatusCode(v int32) *DeleteInstallCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteInstallCodeResponse) SetBody(v *DeleteInstallCodeResponseBody) *DeleteInstallCodeResponse {
	s.Body = v
	return s
}

func (s *DeleteInstallCodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
