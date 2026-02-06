// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDialogueFlowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDialogueFlowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDialogueFlowResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDialogueFlowResponseBody) *ModifyDialogueFlowResponse
	GetBody() *ModifyDialogueFlowResponseBody
}

type ModifyDialogueFlowResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDialogueFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDialogueFlowResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDialogueFlowResponse) GoString() string {
	return s.String()
}

func (s *ModifyDialogueFlowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDialogueFlowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDialogueFlowResponse) GetBody() *ModifyDialogueFlowResponseBody {
	return s.Body
}

func (s *ModifyDialogueFlowResponse) SetHeaders(v map[string]*string) *ModifyDialogueFlowResponse {
	s.Headers = v
	return s
}

func (s *ModifyDialogueFlowResponse) SetStatusCode(v int32) *ModifyDialogueFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDialogueFlowResponse) SetBody(v *ModifyDialogueFlowResponseBody) *ModifyDialogueFlowResponse {
	s.Body = v
	return s
}

func (s *ModifyDialogueFlowResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
