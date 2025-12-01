// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVulSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVulSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVulSummaryResponse
	GetStatusCode() *int32
	SetBody(v *GetVulSummaryResponseBody) *GetVulSummaryResponse
	GetBody() *GetVulSummaryResponseBody
}

type GetVulSummaryResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVulSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVulSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVulSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetVulSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVulSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVulSummaryResponse) GetBody() *GetVulSummaryResponseBody {
	return s.Body
}

func (s *GetVulSummaryResponse) SetHeaders(v map[string]*string) *GetVulSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetVulSummaryResponse) SetStatusCode(v int32) *GetVulSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVulSummaryResponse) SetBody(v *GetVulSummaryResponseBody) *GetVulSummaryResponse {
	s.Body = v
	return s
}

func (s *GetVulSummaryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
