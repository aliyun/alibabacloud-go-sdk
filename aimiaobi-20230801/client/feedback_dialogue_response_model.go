// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFeedbackDialogueResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FeedbackDialogueResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FeedbackDialogueResponse
	GetStatusCode() *int32
	SetBody(v *FeedbackDialogueResponseBody) *FeedbackDialogueResponse
	GetBody() *FeedbackDialogueResponseBody
}

type FeedbackDialogueResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FeedbackDialogueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FeedbackDialogueResponse) String() string {
	return dara.Prettify(s)
}

func (s FeedbackDialogueResponse) GoString() string {
	return s.String()
}

func (s *FeedbackDialogueResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FeedbackDialogueResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FeedbackDialogueResponse) GetBody() *FeedbackDialogueResponseBody {
	return s.Body
}

func (s *FeedbackDialogueResponse) SetHeaders(v map[string]*string) *FeedbackDialogueResponse {
	s.Headers = v
	return s
}

func (s *FeedbackDialogueResponse) SetStatusCode(v int32) *FeedbackDialogueResponse {
	s.StatusCode = &v
	return s
}

func (s *FeedbackDialogueResponse) SetBody(v *FeedbackDialogueResponseBody) *FeedbackDialogueResponse {
	s.Body = v
	return s
}

func (s *FeedbackDialogueResponse) Validate() error {
	return dara.Validate(s)
}
