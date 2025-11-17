// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDataLevelPermissionWhiteListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddDataLevelPermissionWhiteListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddDataLevelPermissionWhiteListResponse
	GetStatusCode() *int32
	SetBody(v *AddDataLevelPermissionWhiteListResponseBody) *AddDataLevelPermissionWhiteListResponse
	GetBody() *AddDataLevelPermissionWhiteListResponseBody
}

type AddDataLevelPermissionWhiteListResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddDataLevelPermissionWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddDataLevelPermissionWhiteListResponse) String() string {
	return dara.Prettify(s)
}

func (s AddDataLevelPermissionWhiteListResponse) GoString() string {
	return s.String()
}

func (s *AddDataLevelPermissionWhiteListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddDataLevelPermissionWhiteListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddDataLevelPermissionWhiteListResponse) GetBody() *AddDataLevelPermissionWhiteListResponseBody {
	return s.Body
}

func (s *AddDataLevelPermissionWhiteListResponse) SetHeaders(v map[string]*string) *AddDataLevelPermissionWhiteListResponse {
	s.Headers = v
	return s
}

func (s *AddDataLevelPermissionWhiteListResponse) SetStatusCode(v int32) *AddDataLevelPermissionWhiteListResponse {
	s.StatusCode = &v
	return s
}

func (s *AddDataLevelPermissionWhiteListResponse) SetBody(v *AddDataLevelPermissionWhiteListResponseBody) *AddDataLevelPermissionWhiteListResponse {
	s.Body = v
	return s
}

func (s *AddDataLevelPermissionWhiteListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
