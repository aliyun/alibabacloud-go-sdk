// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWuyingServerSpecResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyWuyingServerSpecResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyWuyingServerSpecResponse
	GetStatusCode() *int32
	SetBody(v *ModifyWuyingServerSpecResponseBody) *ModifyWuyingServerSpecResponse
	GetBody() *ModifyWuyingServerSpecResponseBody
}

type ModifyWuyingServerSpecResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyWuyingServerSpecResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyWuyingServerSpecResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyWuyingServerSpecResponse) GoString() string {
	return s.String()
}

func (s *ModifyWuyingServerSpecResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyWuyingServerSpecResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyWuyingServerSpecResponse) GetBody() *ModifyWuyingServerSpecResponseBody {
	return s.Body
}

func (s *ModifyWuyingServerSpecResponse) SetHeaders(v map[string]*string) *ModifyWuyingServerSpecResponse {
	s.Headers = v
	return s
}

func (s *ModifyWuyingServerSpecResponse) SetStatusCode(v int32) *ModifyWuyingServerSpecResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyWuyingServerSpecResponse) SetBody(v *ModifyWuyingServerSpecResponseBody) *ModifyWuyingServerSpecResponse {
	s.Body = v
	return s
}

func (s *ModifyWuyingServerSpecResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
