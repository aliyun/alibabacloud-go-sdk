// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRemakeScriptJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRemakeScriptJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRemakeScriptJobResponse
	GetStatusCode() *int32
	SetBody(v *GetRemakeScriptJobResponseBody) *GetRemakeScriptJobResponse
	GetBody() *GetRemakeScriptJobResponseBody
}

type GetRemakeScriptJobResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRemakeScriptJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRemakeScriptJobResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRemakeScriptJobResponse) GoString() string {
	return s.String()
}

func (s *GetRemakeScriptJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRemakeScriptJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRemakeScriptJobResponse) GetBody() *GetRemakeScriptJobResponseBody {
	return s.Body
}

func (s *GetRemakeScriptJobResponse) SetHeaders(v map[string]*string) *GetRemakeScriptJobResponse {
	s.Headers = v
	return s
}

func (s *GetRemakeScriptJobResponse) SetStatusCode(v int32) *GetRemakeScriptJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRemakeScriptJobResponse) SetBody(v *GetRemakeScriptJobResponseBody) *GetRemakeScriptJobResponse {
	s.Body = v
	return s
}

func (s *GetRemakeScriptJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
