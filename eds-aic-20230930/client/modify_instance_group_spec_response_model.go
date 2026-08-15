// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceGroupSpecResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyInstanceGroupSpecResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyInstanceGroupSpecResponse
	GetStatusCode() *int32
	SetBody(v *ModifyInstanceGroupSpecResponseBody) *ModifyInstanceGroupSpecResponse
	GetBody() *ModifyInstanceGroupSpecResponseBody
}

type ModifyInstanceGroupSpecResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyInstanceGroupSpecResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyInstanceGroupSpecResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceGroupSpecResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceGroupSpecResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyInstanceGroupSpecResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyInstanceGroupSpecResponse) GetBody() *ModifyInstanceGroupSpecResponseBody {
	return s.Body
}

func (s *ModifyInstanceGroupSpecResponse) SetHeaders(v map[string]*string) *ModifyInstanceGroupSpecResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceGroupSpecResponse) SetStatusCode(v int32) *ModifyInstanceGroupSpecResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyInstanceGroupSpecResponse) SetBody(v *ModifyInstanceGroupSpecResponseBody) *ModifyInstanceGroupSpecResponse {
	s.Body = v
	return s
}

func (s *ModifyInstanceGroupSpecResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
