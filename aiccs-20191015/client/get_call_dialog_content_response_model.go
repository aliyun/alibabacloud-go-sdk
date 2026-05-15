// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCallDialogContentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCallDialogContentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCallDialogContentResponse
	GetStatusCode() *int32
	SetBody(v *GetCallDialogContentResponseBody) *GetCallDialogContentResponse
	GetBody() *GetCallDialogContentResponseBody
}

type GetCallDialogContentResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCallDialogContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCallDialogContentResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCallDialogContentResponse) GoString() string {
	return s.String()
}

func (s *GetCallDialogContentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCallDialogContentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCallDialogContentResponse) GetBody() *GetCallDialogContentResponseBody {
	return s.Body
}

func (s *GetCallDialogContentResponse) SetHeaders(v map[string]*string) *GetCallDialogContentResponse {
	s.Headers = v
	return s
}

func (s *GetCallDialogContentResponse) SetStatusCode(v int32) *GetCallDialogContentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCallDialogContentResponse) SetBody(v *GetCallDialogContentResponseBody) *GetCallDialogContentResponse {
	s.Body = v
	return s
}

func (s *GetCallDialogContentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
