// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCenRegionDomainRouteEntriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCenRegionDomainRouteEntriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCenRegionDomainRouteEntriesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCenRegionDomainRouteEntriesResponseBody) *DescribeCenRegionDomainRouteEntriesResponse
	GetBody() *DescribeCenRegionDomainRouteEntriesResponseBody
}

type DescribeCenRegionDomainRouteEntriesResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCenRegionDomainRouteEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCenRegionDomainRouteEntriesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenRegionDomainRouteEntriesResponse) GoString() string {
	return s.String()
}

func (s *DescribeCenRegionDomainRouteEntriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCenRegionDomainRouteEntriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCenRegionDomainRouteEntriesResponse) GetBody() *DescribeCenRegionDomainRouteEntriesResponseBody {
	return s.Body
}

func (s *DescribeCenRegionDomainRouteEntriesResponse) SetHeaders(v map[string]*string) *DescribeCenRegionDomainRouteEntriesResponse {
	s.Headers = v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesResponse) SetStatusCode(v int32) *DescribeCenRegionDomainRouteEntriesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesResponse) SetBody(v *DescribeCenRegionDomainRouteEntriesResponseBody) *DescribeCenRegionDomainRouteEntriesResponse {
	s.Body = v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
