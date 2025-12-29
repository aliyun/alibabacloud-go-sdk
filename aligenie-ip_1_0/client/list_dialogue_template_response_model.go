// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDialogueTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDialogueTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDialogueTemplateResponse
	GetStatusCode() *int32
	SetBody(v *ListDialogueTemplateResponseBody) *ListDialogueTemplateResponse
	GetBody() *ListDialogueTemplateResponseBody
}

type ListDialogueTemplateResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDialogueTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDialogueTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDialogueTemplateResponse) GoString() string {
	return s.String()
}

func (s *ListDialogueTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDialogueTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDialogueTemplateResponse) GetBody() *ListDialogueTemplateResponseBody {
	return s.Body
}

func (s *ListDialogueTemplateResponse) SetHeaders(v map[string]*string) *ListDialogueTemplateResponse {
	s.Headers = v
	return s
}

func (s *ListDialogueTemplateResponse) SetStatusCode(v int32) *ListDialogueTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDialogueTemplateResponse) SetBody(v *ListDialogueTemplateResponseBody) *ListDialogueTemplateResponse {
	s.Body = v
	return s
}

func (s *ListDialogueTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
