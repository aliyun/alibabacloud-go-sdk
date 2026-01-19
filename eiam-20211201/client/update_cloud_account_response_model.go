// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloudAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCloudAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCloudAccountResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCloudAccountResponseBody) *UpdateCloudAccountResponse
	GetBody() *UpdateCloudAccountResponseBody
}

type UpdateCloudAccountResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCloudAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCloudAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudAccountResponse) GoString() string {
	return s.String()
}

func (s *UpdateCloudAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCloudAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCloudAccountResponse) GetBody() *UpdateCloudAccountResponseBody {
	return s.Body
}

func (s *UpdateCloudAccountResponse) SetHeaders(v map[string]*string) *UpdateCloudAccountResponse {
	s.Headers = v
	return s
}

func (s *UpdateCloudAccountResponse) SetStatusCode(v int32) *UpdateCloudAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCloudAccountResponse) SetBody(v *UpdateCloudAccountResponseBody) *UpdateCloudAccountResponse {
	s.Body = v
	return s
}

func (s *UpdateCloudAccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
