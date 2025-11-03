// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRepoTagScanSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRepoTagScanSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRepoTagScanSummaryResponse
	GetStatusCode() *int32
	SetBody(v *GetRepoTagScanSummaryResponseBody) *GetRepoTagScanSummaryResponse
	GetBody() *GetRepoTagScanSummaryResponseBody
}

type GetRepoTagScanSummaryResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRepoTagScanSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRepoTagScanSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRepoTagScanSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetRepoTagScanSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRepoTagScanSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRepoTagScanSummaryResponse) GetBody() *GetRepoTagScanSummaryResponseBody {
	return s.Body
}

func (s *GetRepoTagScanSummaryResponse) SetHeaders(v map[string]*string) *GetRepoTagScanSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetRepoTagScanSummaryResponse) SetStatusCode(v int32) *GetRepoTagScanSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRepoTagScanSummaryResponse) SetBody(v *GetRepoTagScanSummaryResponseBody) *GetRepoTagScanSummaryResponse {
	s.Body = v
	return s
}

func (s *GetRepoTagScanSummaryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
