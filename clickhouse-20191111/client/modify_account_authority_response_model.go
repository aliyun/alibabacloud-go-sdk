// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAccountAuthorityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAccountAuthorityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAccountAuthorityResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAccountAuthorityResponseBody) *ModifyAccountAuthorityResponse
	GetBody() *ModifyAccountAuthorityResponseBody
}

type ModifyAccountAuthorityResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAccountAuthorityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAccountAuthorityResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAccountAuthorityResponse) GoString() string {
	return s.String()
}

func (s *ModifyAccountAuthorityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAccountAuthorityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAccountAuthorityResponse) GetBody() *ModifyAccountAuthorityResponseBody {
	return s.Body
}

func (s *ModifyAccountAuthorityResponse) SetHeaders(v map[string]*string) *ModifyAccountAuthorityResponse {
	s.Headers = v
	return s
}

func (s *ModifyAccountAuthorityResponse) SetStatusCode(v int32) *ModifyAccountAuthorityResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAccountAuthorityResponse) SetBody(v *ModifyAccountAuthorityResponseBody) *ModifyAccountAuthorityResponse {
	s.Body = v
	return s
}

func (s *ModifyAccountAuthorityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
