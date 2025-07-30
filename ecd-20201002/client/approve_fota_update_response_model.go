// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApproveFotaUpdateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ApproveFotaUpdateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ApproveFotaUpdateResponse
	GetStatusCode() *int32
	SetBody(v *ApproveFotaUpdateResponseBody) *ApproveFotaUpdateResponse
	GetBody() *ApproveFotaUpdateResponseBody
}

type ApproveFotaUpdateResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApproveFotaUpdateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApproveFotaUpdateResponse) String() string {
	return dara.Prettify(s)
}

func (s ApproveFotaUpdateResponse) GoString() string {
	return s.String()
}

func (s *ApproveFotaUpdateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ApproveFotaUpdateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ApproveFotaUpdateResponse) GetBody() *ApproveFotaUpdateResponseBody {
	return s.Body
}

func (s *ApproveFotaUpdateResponse) SetHeaders(v map[string]*string) *ApproveFotaUpdateResponse {
	s.Headers = v
	return s
}

func (s *ApproveFotaUpdateResponse) SetStatusCode(v int32) *ApproveFotaUpdateResponse {
	s.StatusCode = &v
	return s
}

func (s *ApproveFotaUpdateResponse) SetBody(v *ApproveFotaUpdateResponseBody) *ApproveFotaUpdateResponse {
	s.Body = v
	return s
}

func (s *ApproveFotaUpdateResponse) Validate() error {
	return dara.Validate(s)
}
