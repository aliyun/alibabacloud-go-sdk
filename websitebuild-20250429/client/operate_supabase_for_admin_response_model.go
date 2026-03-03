// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateSupabaseForAdminResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OperateSupabaseForAdminResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OperateSupabaseForAdminResponse
	GetStatusCode() *int32
	SetBody(v *OperateSupabaseForAdminResponseBody) *OperateSupabaseForAdminResponse
	GetBody() *OperateSupabaseForAdminResponseBody
}

type OperateSupabaseForAdminResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperateSupabaseForAdminResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OperateSupabaseForAdminResponse) String() string {
	return dara.Prettify(s)
}

func (s OperateSupabaseForAdminResponse) GoString() string {
	return s.String()
}

func (s *OperateSupabaseForAdminResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OperateSupabaseForAdminResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OperateSupabaseForAdminResponse) GetBody() *OperateSupabaseForAdminResponseBody {
	return s.Body
}

func (s *OperateSupabaseForAdminResponse) SetHeaders(v map[string]*string) *OperateSupabaseForAdminResponse {
	s.Headers = v
	return s
}

func (s *OperateSupabaseForAdminResponse) SetStatusCode(v int32) *OperateSupabaseForAdminResponse {
	s.StatusCode = &v
	return s
}

func (s *OperateSupabaseForAdminResponse) SetBody(v *OperateSupabaseForAdminResponseBody) *OperateSupabaseForAdminResponse {
	s.Body = v
	return s
}

func (s *OperateSupabaseForAdminResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
