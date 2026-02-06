// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDialogueFlowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDialogueFlowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDialogueFlowResponse
	GetStatusCode() *int32
	SetBody(v *CreateDialogueFlowResponseBody) *CreateDialogueFlowResponse
	GetBody() *CreateDialogueFlowResponseBody
}

type CreateDialogueFlowResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDialogueFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDialogueFlowResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDialogueFlowResponse) GoString() string {
	return s.String()
}

func (s *CreateDialogueFlowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDialogueFlowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDialogueFlowResponse) GetBody() *CreateDialogueFlowResponseBody {
	return s.Body
}

func (s *CreateDialogueFlowResponse) SetHeaders(v map[string]*string) *CreateDialogueFlowResponse {
	s.Headers = v
	return s
}

func (s *CreateDialogueFlowResponse) SetStatusCode(v int32) *CreateDialogueFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDialogueFlowResponse) SetBody(v *CreateDialogueFlowResponseBody) *CreateDialogueFlowResponse {
	s.Body = v
	return s
}

func (s *CreateDialogueFlowResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
