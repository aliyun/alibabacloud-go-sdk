// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePdnsThreatLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePdnsThreatLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePdnsThreatLogsResponse
	GetStatusCode() *int32
	SetBody(v *DescribePdnsThreatLogsResponseBody) *DescribePdnsThreatLogsResponse
	GetBody() *DescribePdnsThreatLogsResponseBody
}

type DescribePdnsThreatLogsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePdnsThreatLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePdnsThreatLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePdnsThreatLogsResponse) GoString() string {
	return s.String()
}

func (s *DescribePdnsThreatLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePdnsThreatLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePdnsThreatLogsResponse) GetBody() *DescribePdnsThreatLogsResponseBody {
	return s.Body
}

func (s *DescribePdnsThreatLogsResponse) SetHeaders(v map[string]*string) *DescribePdnsThreatLogsResponse {
	s.Headers = v
	return s
}

func (s *DescribePdnsThreatLogsResponse) SetStatusCode(v int32) *DescribePdnsThreatLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePdnsThreatLogsResponse) SetBody(v *DescribePdnsThreatLogsResponseBody) *DescribePdnsThreatLogsResponse {
	s.Body = v
	return s
}

func (s *DescribePdnsThreatLogsResponse) Validate() error {
	return dara.Validate(s)
}
