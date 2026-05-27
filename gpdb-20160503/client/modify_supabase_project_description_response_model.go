// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySupabaseProjectDescriptionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySupabaseProjectDescriptionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySupabaseProjectDescriptionResponse
	GetStatusCode() *int32
	SetBody(v *ModifySupabaseProjectDescriptionResponseBody) *ModifySupabaseProjectDescriptionResponse
	GetBody() *ModifySupabaseProjectDescriptionResponseBody
}

type ModifySupabaseProjectDescriptionResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySupabaseProjectDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySupabaseProjectDescriptionResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySupabaseProjectDescriptionResponse) GoString() string {
	return s.String()
}

func (s *ModifySupabaseProjectDescriptionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySupabaseProjectDescriptionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySupabaseProjectDescriptionResponse) GetBody() *ModifySupabaseProjectDescriptionResponseBody {
	return s.Body
}

func (s *ModifySupabaseProjectDescriptionResponse) SetHeaders(v map[string]*string) *ModifySupabaseProjectDescriptionResponse {
	s.Headers = v
	return s
}

func (s *ModifySupabaseProjectDescriptionResponse) SetStatusCode(v int32) *ModifySupabaseProjectDescriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySupabaseProjectDescriptionResponse) SetBody(v *ModifySupabaseProjectDescriptionResponseBody) *ModifySupabaseProjectDescriptionResponse {
	s.Body = v
	return s
}

func (s *ModifySupabaseProjectDescriptionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
