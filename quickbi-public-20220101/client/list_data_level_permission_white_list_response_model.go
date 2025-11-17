// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataLevelPermissionWhiteListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataLevelPermissionWhiteListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataLevelPermissionWhiteListResponse
	GetStatusCode() *int32
	SetBody(v *ListDataLevelPermissionWhiteListResponseBody) *ListDataLevelPermissionWhiteListResponse
	GetBody() *ListDataLevelPermissionWhiteListResponseBody
}

type ListDataLevelPermissionWhiteListResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataLevelPermissionWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataLevelPermissionWhiteListResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataLevelPermissionWhiteListResponse) GoString() string {
	return s.String()
}

func (s *ListDataLevelPermissionWhiteListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataLevelPermissionWhiteListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataLevelPermissionWhiteListResponse) GetBody() *ListDataLevelPermissionWhiteListResponseBody {
	return s.Body
}

func (s *ListDataLevelPermissionWhiteListResponse) SetHeaders(v map[string]*string) *ListDataLevelPermissionWhiteListResponse {
	s.Headers = v
	return s
}

func (s *ListDataLevelPermissionWhiteListResponse) SetStatusCode(v int32) *ListDataLevelPermissionWhiteListResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataLevelPermissionWhiteListResponse) SetBody(v *ListDataLevelPermissionWhiteListResponseBody) *ListDataLevelPermissionWhiteListResponse {
	s.Body = v
	return s
}

func (s *ListDataLevelPermissionWhiteListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
