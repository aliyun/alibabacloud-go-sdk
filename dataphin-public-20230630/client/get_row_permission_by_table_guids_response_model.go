// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRowPermissionByTableGuidsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRowPermissionByTableGuidsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRowPermissionByTableGuidsResponse
	GetStatusCode() *int32
	SetBody(v *GetRowPermissionByTableGuidsResponseBody) *GetRowPermissionByTableGuidsResponse
	GetBody() *GetRowPermissionByTableGuidsResponseBody
}

type GetRowPermissionByTableGuidsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRowPermissionByTableGuidsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRowPermissionByTableGuidsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRowPermissionByTableGuidsResponse) GoString() string {
	return s.String()
}

func (s *GetRowPermissionByTableGuidsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRowPermissionByTableGuidsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRowPermissionByTableGuidsResponse) GetBody() *GetRowPermissionByTableGuidsResponseBody {
	return s.Body
}

func (s *GetRowPermissionByTableGuidsResponse) SetHeaders(v map[string]*string) *GetRowPermissionByTableGuidsResponse {
	s.Headers = v
	return s
}

func (s *GetRowPermissionByTableGuidsResponse) SetStatusCode(v int32) *GetRowPermissionByTableGuidsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRowPermissionByTableGuidsResponse) SetBody(v *GetRowPermissionByTableGuidsResponseBody) *GetRowPermissionByTableGuidsResponse {
	s.Body = v
	return s
}

func (s *GetRowPermissionByTableGuidsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
