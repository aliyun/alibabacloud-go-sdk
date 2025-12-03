// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceVpcAttributeForConsoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyInstanceVpcAttributeForConsoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyInstanceVpcAttributeForConsoleResponse
	GetStatusCode() *int32
	SetBody(v *ModifyInstanceVpcAttributeForConsoleResponseBody) *ModifyInstanceVpcAttributeForConsoleResponse
	GetBody() *ModifyInstanceVpcAttributeForConsoleResponseBody
}

type ModifyInstanceVpcAttributeForConsoleResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyInstanceVpcAttributeForConsoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyInstanceVpcAttributeForConsoleResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceVpcAttributeForConsoleResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceVpcAttributeForConsoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyInstanceVpcAttributeForConsoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyInstanceVpcAttributeForConsoleResponse) GetBody() *ModifyInstanceVpcAttributeForConsoleResponseBody {
	return s.Body
}

func (s *ModifyInstanceVpcAttributeForConsoleResponse) SetHeaders(v map[string]*string) *ModifyInstanceVpcAttributeForConsoleResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceVpcAttributeForConsoleResponse) SetStatusCode(v int32) *ModifyInstanceVpcAttributeForConsoleResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyInstanceVpcAttributeForConsoleResponse) SetBody(v *ModifyInstanceVpcAttributeForConsoleResponseBody) *ModifyInstanceVpcAttributeForConsoleResponse {
	s.Body = v
	return s
}

func (s *ModifyInstanceVpcAttributeForConsoleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
