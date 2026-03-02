// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPermissionResourceForFrontResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPermissionResourceForFrontResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPermissionResourceForFrontResponse
	GetStatusCode() *int32
	SetBody(v *ListPermissionResourceForFrontResponseBody) *ListPermissionResourceForFrontResponse
	GetBody() *ListPermissionResourceForFrontResponseBody
}

type ListPermissionResourceForFrontResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPermissionResourceForFrontResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPermissionResourceForFrontResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPermissionResourceForFrontResponse) GoString() string {
	return s.String()
}

func (s *ListPermissionResourceForFrontResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPermissionResourceForFrontResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPermissionResourceForFrontResponse) GetBody() *ListPermissionResourceForFrontResponseBody {
	return s.Body
}

func (s *ListPermissionResourceForFrontResponse) SetHeaders(v map[string]*string) *ListPermissionResourceForFrontResponse {
	s.Headers = v
	return s
}

func (s *ListPermissionResourceForFrontResponse) SetStatusCode(v int32) *ListPermissionResourceForFrontResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPermissionResourceForFrontResponse) SetBody(v *ListPermissionResourceForFrontResponseBody) *ListPermissionResourceForFrontResponse {
	s.Body = v
	return s
}

func (s *ListPermissionResourceForFrontResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
