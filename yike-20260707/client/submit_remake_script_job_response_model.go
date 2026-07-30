// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitRemakeScriptJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitRemakeScriptJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitRemakeScriptJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitRemakeScriptJobResponseBody) *SubmitRemakeScriptJobResponse
	GetBody() *SubmitRemakeScriptJobResponseBody
}

type SubmitRemakeScriptJobResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitRemakeScriptJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitRemakeScriptJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitRemakeScriptJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitRemakeScriptJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitRemakeScriptJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitRemakeScriptJobResponse) GetBody() *SubmitRemakeScriptJobResponseBody {
	return s.Body
}

func (s *SubmitRemakeScriptJobResponse) SetHeaders(v map[string]*string) *SubmitRemakeScriptJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitRemakeScriptJobResponse) SetStatusCode(v int32) *SubmitRemakeScriptJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitRemakeScriptJobResponse) SetBody(v *SubmitRemakeScriptJobResponseBody) *SubmitRemakeScriptJobResponse {
	s.Body = v
	return s
}

func (s *SubmitRemakeScriptJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
