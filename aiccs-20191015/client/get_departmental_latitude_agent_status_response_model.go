// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDepartmentalLatitudeAgentStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDepartmentalLatitudeAgentStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDepartmentalLatitudeAgentStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetDepartmentalLatitudeAgentStatusResponseBody) *GetDepartmentalLatitudeAgentStatusResponse
	GetBody() *GetDepartmentalLatitudeAgentStatusResponseBody
}

type GetDepartmentalLatitudeAgentStatusResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDepartmentalLatitudeAgentStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDepartmentalLatitudeAgentStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDepartmentalLatitudeAgentStatusResponse) GoString() string {
	return s.String()
}

func (s *GetDepartmentalLatitudeAgentStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDepartmentalLatitudeAgentStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDepartmentalLatitudeAgentStatusResponse) GetBody() *GetDepartmentalLatitudeAgentStatusResponseBody {
	return s.Body
}

func (s *GetDepartmentalLatitudeAgentStatusResponse) SetHeaders(v map[string]*string) *GetDepartmentalLatitudeAgentStatusResponse {
	s.Headers = v
	return s
}

func (s *GetDepartmentalLatitudeAgentStatusResponse) SetStatusCode(v int32) *GetDepartmentalLatitudeAgentStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDepartmentalLatitudeAgentStatusResponse) SetBody(v *GetDepartmentalLatitudeAgentStatusResponseBody) *GetDepartmentalLatitudeAgentStatusResponse {
	s.Body = v
	return s
}

func (s *GetDepartmentalLatitudeAgentStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
