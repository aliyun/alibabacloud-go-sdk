// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyFlowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyFlowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyFlowResponse
	GetStatusCode() *int32
	SetBody(v *ModifyFlowResponseBody) *ModifyFlowResponse
	GetBody() *ModifyFlowResponseBody
}

type ModifyFlowResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyFlowResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyFlowResponse) GoString() string {
	return s.String()
}

func (s *ModifyFlowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyFlowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyFlowResponse) GetBody() *ModifyFlowResponseBody {
	return s.Body
}

func (s *ModifyFlowResponse) SetHeaders(v map[string]*string) *ModifyFlowResponse {
	s.Headers = v
	return s
}

func (s *ModifyFlowResponse) SetStatusCode(v int32) *ModifyFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyFlowResponse) SetBody(v *ModifyFlowResponseBody) *ModifyFlowResponse {
	s.Body = v
	return s
}

func (s *ModifyFlowResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
