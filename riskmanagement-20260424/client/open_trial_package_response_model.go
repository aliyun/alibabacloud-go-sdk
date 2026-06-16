// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenTrialPackageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OpenTrialPackageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OpenTrialPackageResponse
	GetStatusCode() *int32
	SetBody(v *OpenTrialPackageResponseBody) *OpenTrialPackageResponse
	GetBody() *OpenTrialPackageResponseBody
}

type OpenTrialPackageResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenTrialPackageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenTrialPackageResponse) String() string {
	return dara.Prettify(s)
}

func (s OpenTrialPackageResponse) GoString() string {
	return s.String()
}

func (s *OpenTrialPackageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OpenTrialPackageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OpenTrialPackageResponse) GetBody() *OpenTrialPackageResponseBody {
	return s.Body
}

func (s *OpenTrialPackageResponse) SetHeaders(v map[string]*string) *OpenTrialPackageResponse {
	s.Headers = v
	return s
}

func (s *OpenTrialPackageResponse) SetStatusCode(v int32) *OpenTrialPackageResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenTrialPackageResponse) SetBody(v *OpenTrialPackageResponseBody) *OpenTrialPackageResponse {
	s.Body = v
	return s
}

func (s *OpenTrialPackageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
