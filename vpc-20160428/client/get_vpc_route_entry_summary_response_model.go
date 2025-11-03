// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVpcRouteEntrySummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVpcRouteEntrySummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVpcRouteEntrySummaryResponse
	GetStatusCode() *int32
	SetBody(v *GetVpcRouteEntrySummaryResponseBody) *GetVpcRouteEntrySummaryResponse
	GetBody() *GetVpcRouteEntrySummaryResponseBody
}

type GetVpcRouteEntrySummaryResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVpcRouteEntrySummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVpcRouteEntrySummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVpcRouteEntrySummaryResponse) GoString() string {
	return s.String()
}

func (s *GetVpcRouteEntrySummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVpcRouteEntrySummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVpcRouteEntrySummaryResponse) GetBody() *GetVpcRouteEntrySummaryResponseBody {
	return s.Body
}

func (s *GetVpcRouteEntrySummaryResponse) SetHeaders(v map[string]*string) *GetVpcRouteEntrySummaryResponse {
	s.Headers = v
	return s
}

func (s *GetVpcRouteEntrySummaryResponse) SetStatusCode(v int32) *GetVpcRouteEntrySummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVpcRouteEntrySummaryResponse) SetBody(v *GetVpcRouteEntrySummaryResponseBody) *GetVpcRouteEntrySummaryResponse {
	s.Body = v
	return s
}

func (s *GetVpcRouteEntrySummaryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
