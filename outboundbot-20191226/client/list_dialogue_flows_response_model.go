// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDialogueFlowsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDialogueFlowsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDialogueFlowsResponse
	GetStatusCode() *int32
	SetBody(v *ListDialogueFlowsResponseBody) *ListDialogueFlowsResponse
	GetBody() *ListDialogueFlowsResponseBody
}

type ListDialogueFlowsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDialogueFlowsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDialogueFlowsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDialogueFlowsResponse) GoString() string {
	return s.String()
}

func (s *ListDialogueFlowsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDialogueFlowsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDialogueFlowsResponse) GetBody() *ListDialogueFlowsResponseBody {
	return s.Body
}

func (s *ListDialogueFlowsResponse) SetHeaders(v map[string]*string) *ListDialogueFlowsResponse {
	s.Headers = v
	return s
}

func (s *ListDialogueFlowsResponse) SetStatusCode(v int32) *ListDialogueFlowsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDialogueFlowsResponse) SetBody(v *ListDialogueFlowsResponseBody) *ListDialogueFlowsResponse {
	s.Body = v
	return s
}

func (s *ListDialogueFlowsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
