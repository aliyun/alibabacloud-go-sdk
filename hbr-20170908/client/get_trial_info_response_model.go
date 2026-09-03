// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTrialInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTrialInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTrialInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetTrialInfoResponseBody) *GetTrialInfoResponse
	GetBody() *GetTrialInfoResponseBody
}

type GetTrialInfoResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTrialInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTrialInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTrialInfoResponse) GoString() string {
	return s.String()
}

func (s *GetTrialInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTrialInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTrialInfoResponse) GetBody() *GetTrialInfoResponseBody {
	return s.Body
}

func (s *GetTrialInfoResponse) SetHeaders(v map[string]*string) *GetTrialInfoResponse {
	s.Headers = v
	return s
}

func (s *GetTrialInfoResponse) SetStatusCode(v int32) *GetTrialInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTrialInfoResponse) SetBody(v *GetTrialInfoResponseBody) *GetTrialInfoResponse {
	s.Body = v
	return s
}

func (s *GetTrialInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
