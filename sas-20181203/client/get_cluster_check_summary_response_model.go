// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClusterCheckSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetClusterCheckSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetClusterCheckSummaryResponse
	GetStatusCode() *int32
	SetBody(v *GetClusterCheckSummaryResponseBody) *GetClusterCheckSummaryResponse
	GetBody() *GetClusterCheckSummaryResponseBody
}

type GetClusterCheckSummaryResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetClusterCheckSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetClusterCheckSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetClusterCheckSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetClusterCheckSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetClusterCheckSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetClusterCheckSummaryResponse) GetBody() *GetClusterCheckSummaryResponseBody {
	return s.Body
}

func (s *GetClusterCheckSummaryResponse) SetHeaders(v map[string]*string) *GetClusterCheckSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetClusterCheckSummaryResponse) SetStatusCode(v int32) *GetClusterCheckSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetClusterCheckSummaryResponse) SetBody(v *GetClusterCheckSummaryResponseBody) *GetClusterCheckSummaryResponse {
	s.Body = v
	return s
}

func (s *GetClusterCheckSummaryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
