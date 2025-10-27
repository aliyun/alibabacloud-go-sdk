// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetModuleTrialAuthInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetModuleTrialAuthInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetModuleTrialAuthInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetModuleTrialAuthInfoResponseBody) *GetModuleTrialAuthInfoResponse
	GetBody() *GetModuleTrialAuthInfoResponseBody
}

type GetModuleTrialAuthInfoResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetModuleTrialAuthInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetModuleTrialAuthInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetModuleTrialAuthInfoResponse) GoString() string {
	return s.String()
}

func (s *GetModuleTrialAuthInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetModuleTrialAuthInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetModuleTrialAuthInfoResponse) GetBody() *GetModuleTrialAuthInfoResponseBody {
	return s.Body
}

func (s *GetModuleTrialAuthInfoResponse) SetHeaders(v map[string]*string) *GetModuleTrialAuthInfoResponse {
	s.Headers = v
	return s
}

func (s *GetModuleTrialAuthInfoResponse) SetStatusCode(v int32) *GetModuleTrialAuthInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetModuleTrialAuthInfoResponse) SetBody(v *GetModuleTrialAuthInfoResponseBody) *GetModuleTrialAuthInfoResponse {
	s.Body = v
	return s
}

func (s *GetModuleTrialAuthInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
