// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceComplianceByPackResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetResourceComplianceByPackResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetResourceComplianceByPackResponse
	GetStatusCode() *int32
	SetBody(v *GetResourceComplianceByPackResponseBody) *GetResourceComplianceByPackResponse
	GetBody() *GetResourceComplianceByPackResponseBody
}

type GetResourceComplianceByPackResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResourceComplianceByPackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResourceComplianceByPackResponse) String() string {
	return dara.Prettify(s)
}

func (s GetResourceComplianceByPackResponse) GoString() string {
	return s.String()
}

func (s *GetResourceComplianceByPackResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetResourceComplianceByPackResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetResourceComplianceByPackResponse) GetBody() *GetResourceComplianceByPackResponseBody {
	return s.Body
}

func (s *GetResourceComplianceByPackResponse) SetHeaders(v map[string]*string) *GetResourceComplianceByPackResponse {
	s.Headers = v
	return s
}

func (s *GetResourceComplianceByPackResponse) SetStatusCode(v int32) *GetResourceComplianceByPackResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceComplianceByPackResponse) SetBody(v *GetResourceComplianceByPackResponseBody) *GetResourceComplianceByPackResponse {
	s.Body = v
	return s
}

func (s *GetResourceComplianceByPackResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
