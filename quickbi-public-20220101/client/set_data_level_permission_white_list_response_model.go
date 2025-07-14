// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDataLevelPermissionWhiteListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetDataLevelPermissionWhiteListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetDataLevelPermissionWhiteListResponse
	GetStatusCode() *int32
	SetBody(v *SetDataLevelPermissionWhiteListResponseBody) *SetDataLevelPermissionWhiteListResponse
	GetBody() *SetDataLevelPermissionWhiteListResponseBody
}

type SetDataLevelPermissionWhiteListResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetDataLevelPermissionWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetDataLevelPermissionWhiteListResponse) String() string {
	return dara.Prettify(s)
}

func (s SetDataLevelPermissionWhiteListResponse) GoString() string {
	return s.String()
}

func (s *SetDataLevelPermissionWhiteListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetDataLevelPermissionWhiteListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetDataLevelPermissionWhiteListResponse) GetBody() *SetDataLevelPermissionWhiteListResponseBody {
	return s.Body
}

func (s *SetDataLevelPermissionWhiteListResponse) SetHeaders(v map[string]*string) *SetDataLevelPermissionWhiteListResponse {
	s.Headers = v
	return s
}

func (s *SetDataLevelPermissionWhiteListResponse) SetStatusCode(v int32) *SetDataLevelPermissionWhiteListResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDataLevelPermissionWhiteListResponse) SetBody(v *SetDataLevelPermissionWhiteListResponseBody) *SetDataLevelPermissionWhiteListResponse {
	s.Body = v
	return s
}

func (s *SetDataLevelPermissionWhiteListResponse) Validate() error {
	return dara.Validate(s)
}
