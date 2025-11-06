// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMpaasAppInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMpaasAppInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMpaasAppInfoResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMpaasAppInfoResponseBody) *UpdateMpaasAppInfoResponse
	GetBody() *UpdateMpaasAppInfoResponseBody
}

type UpdateMpaasAppInfoResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMpaasAppInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMpaasAppInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMpaasAppInfoResponse) GoString() string {
	return s.String()
}

func (s *UpdateMpaasAppInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMpaasAppInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMpaasAppInfoResponse) GetBody() *UpdateMpaasAppInfoResponseBody {
	return s.Body
}

func (s *UpdateMpaasAppInfoResponse) SetHeaders(v map[string]*string) *UpdateMpaasAppInfoResponse {
	s.Headers = v
	return s
}

func (s *UpdateMpaasAppInfoResponse) SetStatusCode(v int32) *UpdateMpaasAppInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMpaasAppInfoResponse) SetBody(v *UpdateMpaasAppInfoResponseBody) *UpdateMpaasAppInfoResponse {
	s.Body = v
	return s
}

func (s *UpdateMpaasAppInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
