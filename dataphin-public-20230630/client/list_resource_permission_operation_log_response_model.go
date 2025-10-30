// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourcePermissionOperationLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListResourcePermissionOperationLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListResourcePermissionOperationLogResponse
	GetStatusCode() *int32
	SetBody(v *ListResourcePermissionOperationLogResponseBody) *ListResourcePermissionOperationLogResponse
	GetBody() *ListResourcePermissionOperationLogResponseBody
}

type ListResourcePermissionOperationLogResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListResourcePermissionOperationLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListResourcePermissionOperationLogResponse) String() string {
	return dara.Prettify(s)
}

func (s ListResourcePermissionOperationLogResponse) GoString() string {
	return s.String()
}

func (s *ListResourcePermissionOperationLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListResourcePermissionOperationLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListResourcePermissionOperationLogResponse) GetBody() *ListResourcePermissionOperationLogResponseBody {
	return s.Body
}

func (s *ListResourcePermissionOperationLogResponse) SetHeaders(v map[string]*string) *ListResourcePermissionOperationLogResponse {
	s.Headers = v
	return s
}

func (s *ListResourcePermissionOperationLogResponse) SetStatusCode(v int32) *ListResourcePermissionOperationLogResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourcePermissionOperationLogResponse) SetBody(v *ListResourcePermissionOperationLogResponseBody) *ListResourcePermissionOperationLogResponse {
	s.Body = v
	return s
}

func (s *ListResourcePermissionOperationLogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
