// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySupabaseProjectResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySupabaseProjectResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySupabaseProjectResourceResponse
	GetStatusCode() *int32
	SetBody(v *ModifySupabaseProjectResourceResponseBody) *ModifySupabaseProjectResourceResponse
	GetBody() *ModifySupabaseProjectResourceResponseBody
}

type ModifySupabaseProjectResourceResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySupabaseProjectResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySupabaseProjectResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySupabaseProjectResourceResponse) GoString() string {
	return s.String()
}

func (s *ModifySupabaseProjectResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySupabaseProjectResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySupabaseProjectResourceResponse) GetBody() *ModifySupabaseProjectResourceResponseBody {
	return s.Body
}

func (s *ModifySupabaseProjectResourceResponse) SetHeaders(v map[string]*string) *ModifySupabaseProjectResourceResponse {
	s.Headers = v
	return s
}

func (s *ModifySupabaseProjectResourceResponse) SetStatusCode(v int32) *ModifySupabaseProjectResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySupabaseProjectResourceResponse) SetBody(v *ModifySupabaseProjectResourceResponseBody) *ModifySupabaseProjectResourceResponse {
	s.Body = v
	return s
}

func (s *ModifySupabaseProjectResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
