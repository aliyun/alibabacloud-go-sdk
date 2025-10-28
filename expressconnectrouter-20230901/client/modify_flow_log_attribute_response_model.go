// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyFlowLogAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyFlowLogAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyFlowLogAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyFlowLogAttributeResponseBody) *ModifyFlowLogAttributeResponse
	GetBody() *ModifyFlowLogAttributeResponseBody
}

type ModifyFlowLogAttributeResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyFlowLogAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyFlowLogAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyFlowLogAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyFlowLogAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyFlowLogAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyFlowLogAttributeResponse) GetBody() *ModifyFlowLogAttributeResponseBody {
	return s.Body
}

func (s *ModifyFlowLogAttributeResponse) SetHeaders(v map[string]*string) *ModifyFlowLogAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyFlowLogAttributeResponse) SetStatusCode(v int32) *ModifyFlowLogAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyFlowLogAttributeResponse) SetBody(v *ModifyFlowLogAttributeResponseBody) *ModifyFlowLogAttributeResponse {
	s.Body = v
	return s
}

func (s *ModifyFlowLogAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
