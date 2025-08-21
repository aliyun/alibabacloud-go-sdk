// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloudAppInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCloudAppInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCloudAppInfoResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCloudAppInfoResponseBody) *UpdateCloudAppInfoResponse
	GetBody() *UpdateCloudAppInfoResponseBody
}

type UpdateCloudAppInfoResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCloudAppInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCloudAppInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudAppInfoResponse) GoString() string {
	return s.String()
}

func (s *UpdateCloudAppInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCloudAppInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCloudAppInfoResponse) GetBody() *UpdateCloudAppInfoResponseBody {
	return s.Body
}

func (s *UpdateCloudAppInfoResponse) SetHeaders(v map[string]*string) *UpdateCloudAppInfoResponse {
	s.Headers = v
	return s
}

func (s *UpdateCloudAppInfoResponse) SetStatusCode(v int32) *UpdateCloudAppInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCloudAppInfoResponse) SetBody(v *UpdateCloudAppInfoResponseBody) *UpdateCloudAppInfoResponse {
	s.Body = v
	return s
}

func (s *UpdateCloudAppInfoResponse) Validate() error {
	return dara.Validate(s)
}
