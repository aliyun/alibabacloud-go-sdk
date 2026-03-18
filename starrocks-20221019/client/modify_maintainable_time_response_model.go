// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMaintainableTimeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyMaintainableTimeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyMaintainableTimeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyMaintainableTimeResponseBody) *ModifyMaintainableTimeResponse
	GetBody() *ModifyMaintainableTimeResponseBody
}

type ModifyMaintainableTimeResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyMaintainableTimeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyMaintainableTimeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyMaintainableTimeResponse) GoString() string {
	return s.String()
}

func (s *ModifyMaintainableTimeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyMaintainableTimeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyMaintainableTimeResponse) GetBody() *ModifyMaintainableTimeResponseBody {
	return s.Body
}

func (s *ModifyMaintainableTimeResponse) SetHeaders(v map[string]*string) *ModifyMaintainableTimeResponse {
	s.Headers = v
	return s
}

func (s *ModifyMaintainableTimeResponse) SetStatusCode(v int32) *ModifyMaintainableTimeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyMaintainableTimeResponse) SetBody(v *ModifyMaintainableTimeResponseBody) *ModifyMaintainableTimeResponse {
	s.Body = v
	return s
}

func (s *ModifyMaintainableTimeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
