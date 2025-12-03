// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyInstanceTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyInstanceTypeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyInstanceTypeResponseBody) *ModifyInstanceTypeResponse
	GetBody() *ModifyInstanceTypeResponseBody
}

type ModifyInstanceTypeResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyInstanceTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyInstanceTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceTypeResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyInstanceTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyInstanceTypeResponse) GetBody() *ModifyInstanceTypeResponseBody {
	return s.Body
}

func (s *ModifyInstanceTypeResponse) SetHeaders(v map[string]*string) *ModifyInstanceTypeResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceTypeResponse) SetStatusCode(v int32) *ModifyInstanceTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyInstanceTypeResponse) SetBody(v *ModifyInstanceTypeResponseBody) *ModifyInstanceTypeResponse {
	s.Body = v
	return s
}

func (s *ModifyInstanceTypeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
