// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySupabaseAutoScalePolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySupabaseAutoScalePolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySupabaseAutoScalePolicyResponse
	GetStatusCode() *int32
	SetBody(v *ModifySupabaseAutoScalePolicyResponseBody) *ModifySupabaseAutoScalePolicyResponse
	GetBody() *ModifySupabaseAutoScalePolicyResponseBody
}

type ModifySupabaseAutoScalePolicyResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySupabaseAutoScalePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySupabaseAutoScalePolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySupabaseAutoScalePolicyResponse) GoString() string {
	return s.String()
}

func (s *ModifySupabaseAutoScalePolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySupabaseAutoScalePolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySupabaseAutoScalePolicyResponse) GetBody() *ModifySupabaseAutoScalePolicyResponseBody {
	return s.Body
}

func (s *ModifySupabaseAutoScalePolicyResponse) SetHeaders(v map[string]*string) *ModifySupabaseAutoScalePolicyResponse {
	s.Headers = v
	return s
}

func (s *ModifySupabaseAutoScalePolicyResponse) SetStatusCode(v int32) *ModifySupabaseAutoScalePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySupabaseAutoScalePolicyResponse) SetBody(v *ModifySupabaseAutoScalePolicyResponseBody) *ModifySupabaseAutoScalePolicyResponse {
	s.Body = v
	return s
}

func (s *ModifySupabaseAutoScalePolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
