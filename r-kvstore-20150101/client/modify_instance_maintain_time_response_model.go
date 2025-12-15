// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceMaintainTimeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyInstanceMaintainTimeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyInstanceMaintainTimeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyInstanceMaintainTimeResponseBody) *ModifyInstanceMaintainTimeResponse
	GetBody() *ModifyInstanceMaintainTimeResponseBody
}

type ModifyInstanceMaintainTimeResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyInstanceMaintainTimeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyInstanceMaintainTimeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceMaintainTimeResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceMaintainTimeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyInstanceMaintainTimeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyInstanceMaintainTimeResponse) GetBody() *ModifyInstanceMaintainTimeResponseBody {
	return s.Body
}

func (s *ModifyInstanceMaintainTimeResponse) SetHeaders(v map[string]*string) *ModifyInstanceMaintainTimeResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceMaintainTimeResponse) SetStatusCode(v int32) *ModifyInstanceMaintainTimeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyInstanceMaintainTimeResponse) SetBody(v *ModifyInstanceMaintainTimeResponseBody) *ModifyInstanceMaintainTimeResponse {
	s.Body = v
	return s
}

func (s *ModifyInstanceMaintainTimeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
