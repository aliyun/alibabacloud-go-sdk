// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAnycastEipAddressSpecResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAnycastEipAddressSpecResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAnycastEipAddressSpecResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAnycastEipAddressSpecResponseBody) *ModifyAnycastEipAddressSpecResponse
	GetBody() *ModifyAnycastEipAddressSpecResponseBody
}

type ModifyAnycastEipAddressSpecResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAnycastEipAddressSpecResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAnycastEipAddressSpecResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAnycastEipAddressSpecResponse) GoString() string {
	return s.String()
}

func (s *ModifyAnycastEipAddressSpecResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAnycastEipAddressSpecResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAnycastEipAddressSpecResponse) GetBody() *ModifyAnycastEipAddressSpecResponseBody {
	return s.Body
}

func (s *ModifyAnycastEipAddressSpecResponse) SetHeaders(v map[string]*string) *ModifyAnycastEipAddressSpecResponse {
	s.Headers = v
	return s
}

func (s *ModifyAnycastEipAddressSpecResponse) SetStatusCode(v int32) *ModifyAnycastEipAddressSpecResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAnycastEipAddressSpecResponse) SetBody(v *ModifyAnycastEipAddressSpecResponseBody) *ModifyAnycastEipAddressSpecResponse {
	s.Body = v
	return s
}

func (s *ModifyAnycastEipAddressSpecResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
