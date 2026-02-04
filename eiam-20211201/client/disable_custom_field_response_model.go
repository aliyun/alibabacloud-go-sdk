// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableCustomFieldResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableCustomFieldResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableCustomFieldResponse
	GetStatusCode() *int32
	SetBody(v *DisableCustomFieldResponseBody) *DisableCustomFieldResponse
	GetBody() *DisableCustomFieldResponseBody
}

type DisableCustomFieldResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableCustomFieldResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableCustomFieldResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableCustomFieldResponse) GoString() string {
	return s.String()
}

func (s *DisableCustomFieldResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableCustomFieldResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableCustomFieldResponse) GetBody() *DisableCustomFieldResponseBody {
	return s.Body
}

func (s *DisableCustomFieldResponse) SetHeaders(v map[string]*string) *DisableCustomFieldResponse {
	s.Headers = v
	return s
}

func (s *DisableCustomFieldResponse) SetStatusCode(v int32) *DisableCustomFieldResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableCustomFieldResponse) SetBody(v *DisableCustomFieldResponseBody) *DisableCustomFieldResponse {
	s.Body = v
	return s
}

func (s *DisableCustomFieldResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
