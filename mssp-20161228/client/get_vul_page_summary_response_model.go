// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVulPageSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVulPageSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVulPageSummaryResponse
	GetStatusCode() *int32
	SetBody(v *GetVulPageSummaryResponseBody) *GetVulPageSummaryResponse
	GetBody() *GetVulPageSummaryResponseBody
}

type GetVulPageSummaryResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVulPageSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVulPageSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVulPageSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetVulPageSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVulPageSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVulPageSummaryResponse) GetBody() *GetVulPageSummaryResponseBody {
	return s.Body
}

func (s *GetVulPageSummaryResponse) SetHeaders(v map[string]*string) *GetVulPageSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetVulPageSummaryResponse) SetStatusCode(v int32) *GetVulPageSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVulPageSummaryResponse) SetBody(v *GetVulPageSummaryResponseBody) *GetVulPageSummaryResponse {
	s.Body = v
	return s
}

func (s *GetVulPageSummaryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
