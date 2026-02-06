// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDialogueFlowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDialogueFlowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDialogueFlowResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDialogueFlowResponseBody) *DeleteDialogueFlowResponse
	GetBody() *DeleteDialogueFlowResponseBody
}

type DeleteDialogueFlowResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDialogueFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDialogueFlowResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDialogueFlowResponse) GoString() string {
	return s.String()
}

func (s *DeleteDialogueFlowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDialogueFlowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDialogueFlowResponse) GetBody() *DeleteDialogueFlowResponseBody {
	return s.Body
}

func (s *DeleteDialogueFlowResponse) SetHeaders(v map[string]*string) *DeleteDialogueFlowResponse {
	s.Headers = v
	return s
}

func (s *DeleteDialogueFlowResponse) SetStatusCode(v int32) *DeleteDialogueFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDialogueFlowResponse) SetBody(v *DeleteDialogueFlowResponseBody) *DeleteDialogueFlowResponse {
	s.Body = v
	return s
}

func (s *DeleteDialogueFlowResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
