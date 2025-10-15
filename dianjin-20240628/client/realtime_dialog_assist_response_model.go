// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRealtimeDialogAssistResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RealtimeDialogAssistResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RealtimeDialogAssistResponse
	GetStatusCode() *int32
	SetBody(v *RealtimeDialogAssistResponseBody) *RealtimeDialogAssistResponse
	GetBody() *RealtimeDialogAssistResponseBody
}

type RealtimeDialogAssistResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RealtimeDialogAssistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RealtimeDialogAssistResponse) String() string {
	return dara.Prettify(s)
}

func (s RealtimeDialogAssistResponse) GoString() string {
	return s.String()
}

func (s *RealtimeDialogAssistResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RealtimeDialogAssistResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RealtimeDialogAssistResponse) GetBody() *RealtimeDialogAssistResponseBody {
	return s.Body
}

func (s *RealtimeDialogAssistResponse) SetHeaders(v map[string]*string) *RealtimeDialogAssistResponse {
	s.Headers = v
	return s
}

func (s *RealtimeDialogAssistResponse) SetStatusCode(v int32) *RealtimeDialogAssistResponse {
	s.StatusCode = &v
	return s
}

func (s *RealtimeDialogAssistResponse) SetBody(v *RealtimeDialogAssistResponseBody) *RealtimeDialogAssistResponse {
	s.Body = v
	return s
}

func (s *RealtimeDialogAssistResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
