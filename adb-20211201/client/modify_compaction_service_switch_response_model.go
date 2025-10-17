// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCompactionServiceSwitchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyCompactionServiceSwitchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyCompactionServiceSwitchResponse
	GetStatusCode() *int32
	SetBody(v *ModifyCompactionServiceSwitchResponseBody) *ModifyCompactionServiceSwitchResponse
	GetBody() *ModifyCompactionServiceSwitchResponseBody
}

type ModifyCompactionServiceSwitchResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyCompactionServiceSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyCompactionServiceSwitchResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyCompactionServiceSwitchResponse) GoString() string {
	return s.String()
}

func (s *ModifyCompactionServiceSwitchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyCompactionServiceSwitchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyCompactionServiceSwitchResponse) GetBody() *ModifyCompactionServiceSwitchResponseBody {
	return s.Body
}

func (s *ModifyCompactionServiceSwitchResponse) SetHeaders(v map[string]*string) *ModifyCompactionServiceSwitchResponse {
	s.Headers = v
	return s
}

func (s *ModifyCompactionServiceSwitchResponse) SetStatusCode(v int32) *ModifyCompactionServiceSwitchResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyCompactionServiceSwitchResponse) SetBody(v *ModifyCompactionServiceSwitchResponseBody) *ModifyCompactionServiceSwitchResponse {
	s.Body = v
	return s
}

func (s *ModifyCompactionServiceSwitchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
