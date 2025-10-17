// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppAgentFunctionStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAppAgentFunctionStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAppAgentFunctionStatusResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAppAgentFunctionStatusResponseBody) *ModifyAppAgentFunctionStatusResponse
	GetBody() *ModifyAppAgentFunctionStatusResponseBody
}

type ModifyAppAgentFunctionStatusResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAppAgentFunctionStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAppAgentFunctionStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppAgentFunctionStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyAppAgentFunctionStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAppAgentFunctionStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAppAgentFunctionStatusResponse) GetBody() *ModifyAppAgentFunctionStatusResponseBody {
	return s.Body
}

func (s *ModifyAppAgentFunctionStatusResponse) SetHeaders(v map[string]*string) *ModifyAppAgentFunctionStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyAppAgentFunctionStatusResponse) SetStatusCode(v int32) *ModifyAppAgentFunctionStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAppAgentFunctionStatusResponse) SetBody(v *ModifyAppAgentFunctionStatusResponseBody) *ModifyAppAgentFunctionStatusResponse {
	s.Body = v
	return s
}

func (s *ModifyAppAgentFunctionStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
