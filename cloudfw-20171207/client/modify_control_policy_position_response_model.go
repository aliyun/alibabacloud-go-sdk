// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyControlPolicyPositionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyControlPolicyPositionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyControlPolicyPositionResponse
	GetStatusCode() *int32
	SetBody(v *ModifyControlPolicyPositionResponseBody) *ModifyControlPolicyPositionResponse
	GetBody() *ModifyControlPolicyPositionResponseBody
}

type ModifyControlPolicyPositionResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyControlPolicyPositionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyControlPolicyPositionResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyControlPolicyPositionResponse) GoString() string {
	return s.String()
}

func (s *ModifyControlPolicyPositionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyControlPolicyPositionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyControlPolicyPositionResponse) GetBody() *ModifyControlPolicyPositionResponseBody {
	return s.Body
}

func (s *ModifyControlPolicyPositionResponse) SetHeaders(v map[string]*string) *ModifyControlPolicyPositionResponse {
	s.Headers = v
	return s
}

func (s *ModifyControlPolicyPositionResponse) SetStatusCode(v int32) *ModifyControlPolicyPositionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyControlPolicyPositionResponse) SetBody(v *ModifyControlPolicyPositionResponseBody) *ModifyControlPolicyPositionResponse {
	s.Body = v
	return s
}

func (s *ModifyControlPolicyPositionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
