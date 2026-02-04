// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceTrialStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInstanceTrialStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInstanceTrialStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetInstanceTrialStatusResponseBody) *GetInstanceTrialStatusResponse
	GetBody() *GetInstanceTrialStatusResponseBody
}

type GetInstanceTrialStatusResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceTrialStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceTrialStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceTrialStatusResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceTrialStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInstanceTrialStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInstanceTrialStatusResponse) GetBody() *GetInstanceTrialStatusResponseBody {
	return s.Body
}

func (s *GetInstanceTrialStatusResponse) SetHeaders(v map[string]*string) *GetInstanceTrialStatusResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceTrialStatusResponse) SetStatusCode(v int32) *GetInstanceTrialStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceTrialStatusResponse) SetBody(v *GetInstanceTrialStatusResponseBody) *GetInstanceTrialStatusResponse {
	s.Body = v
	return s
}

func (s *GetInstanceTrialStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
