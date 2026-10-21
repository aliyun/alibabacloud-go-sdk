// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIncidentInvestigationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetIncidentInvestigationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetIncidentInvestigationResponse
	GetStatusCode() *int32
	SetBody(v *GetIncidentInvestigationResponseBody) *GetIncidentInvestigationResponse
	GetBody() *GetIncidentInvestigationResponseBody
}

type GetIncidentInvestigationResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetIncidentInvestigationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetIncidentInvestigationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetIncidentInvestigationResponse) GoString() string {
	return s.String()
}

func (s *GetIncidentInvestigationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetIncidentInvestigationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetIncidentInvestigationResponse) GetBody() *GetIncidentInvestigationResponseBody {
	return s.Body
}

func (s *GetIncidentInvestigationResponse) SetHeaders(v map[string]*string) *GetIncidentInvestigationResponse {
	s.Headers = v
	return s
}

func (s *GetIncidentInvestigationResponse) SetStatusCode(v int32) *GetIncidentInvestigationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIncidentInvestigationResponse) SetBody(v *GetIncidentInvestigationResponseBody) *GetIncidentInvestigationResponse {
	s.Body = v
	return s
}

func (s *GetIncidentInvestigationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
