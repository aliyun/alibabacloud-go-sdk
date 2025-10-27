// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFunctionTrialStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFunctionTrialStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFunctionTrialStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetFunctionTrialStatusResponseBody) *GetFunctionTrialStatusResponse
	GetBody() *GetFunctionTrialStatusResponseBody
}

type GetFunctionTrialStatusResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFunctionTrialStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFunctionTrialStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFunctionTrialStatusResponse) GoString() string {
	return s.String()
}

func (s *GetFunctionTrialStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFunctionTrialStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFunctionTrialStatusResponse) GetBody() *GetFunctionTrialStatusResponseBody {
	return s.Body
}

func (s *GetFunctionTrialStatusResponse) SetHeaders(v map[string]*string) *GetFunctionTrialStatusResponse {
	s.Headers = v
	return s
}

func (s *GetFunctionTrialStatusResponse) SetStatusCode(v int32) *GetFunctionTrialStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFunctionTrialStatusResponse) SetBody(v *GetFunctionTrialStatusResponseBody) *GetFunctionTrialStatusResponse {
	s.Body = v
	return s
}

func (s *GetFunctionTrialStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
