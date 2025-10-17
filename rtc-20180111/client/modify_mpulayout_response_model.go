// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMPULayoutResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyMPULayoutResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyMPULayoutResponse
	GetStatusCode() *int32
	SetBody(v *ModifyMPULayoutResponseBody) *ModifyMPULayoutResponse
	GetBody() *ModifyMPULayoutResponseBody
}

type ModifyMPULayoutResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyMPULayoutResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyMPULayoutResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyMPULayoutResponse) GoString() string {
	return s.String()
}

func (s *ModifyMPULayoutResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyMPULayoutResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyMPULayoutResponse) GetBody() *ModifyMPULayoutResponseBody {
	return s.Body
}

func (s *ModifyMPULayoutResponse) SetHeaders(v map[string]*string) *ModifyMPULayoutResponse {
	s.Headers = v
	return s
}

func (s *ModifyMPULayoutResponse) SetStatusCode(v int32) *ModifyMPULayoutResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyMPULayoutResponse) SetBody(v *ModifyMPULayoutResponseBody) *ModifyMPULayoutResponse {
	s.Body = v
	return s
}

func (s *ModifyMPULayoutResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
