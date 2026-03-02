// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeveloperEnterpriseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDeveloperEnterpriseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDeveloperEnterpriseResponse
	GetStatusCode() *int32
	SetBody(v *GetDeveloperEnterpriseResponseBody) *GetDeveloperEnterpriseResponse
	GetBody() *GetDeveloperEnterpriseResponseBody
}

type GetDeveloperEnterpriseResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDeveloperEnterpriseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDeveloperEnterpriseResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDeveloperEnterpriseResponse) GoString() string {
	return s.String()
}

func (s *GetDeveloperEnterpriseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDeveloperEnterpriseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDeveloperEnterpriseResponse) GetBody() *GetDeveloperEnterpriseResponseBody {
	return s.Body
}

func (s *GetDeveloperEnterpriseResponse) SetHeaders(v map[string]*string) *GetDeveloperEnterpriseResponse {
	s.Headers = v
	return s
}

func (s *GetDeveloperEnterpriseResponse) SetStatusCode(v int32) *GetDeveloperEnterpriseResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeveloperEnterpriseResponse) SetBody(v *GetDeveloperEnterpriseResponseBody) *GetDeveloperEnterpriseResponse {
	s.Body = v
	return s
}

func (s *GetDeveloperEnterpriseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
