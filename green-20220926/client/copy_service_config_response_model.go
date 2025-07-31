// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyServiceConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CopyServiceConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CopyServiceConfigResponse
	GetStatusCode() *int32
	SetBody(v *CopyServiceConfigResponseBody) *CopyServiceConfigResponse
	GetBody() *CopyServiceConfigResponseBody
}

type CopyServiceConfigResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CopyServiceConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CopyServiceConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s CopyServiceConfigResponse) GoString() string {
	return s.String()
}

func (s *CopyServiceConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CopyServiceConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CopyServiceConfigResponse) GetBody() *CopyServiceConfigResponseBody {
	return s.Body
}

func (s *CopyServiceConfigResponse) SetHeaders(v map[string]*string) *CopyServiceConfigResponse {
	s.Headers = v
	return s
}

func (s *CopyServiceConfigResponse) SetStatusCode(v int32) *CopyServiceConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *CopyServiceConfigResponse) SetBody(v *CopyServiceConfigResponseBody) *CopyServiceConfigResponse {
	s.Body = v
	return s
}

func (s *CopyServiceConfigResponse) Validate() error {
	return dara.Validate(s)
}
