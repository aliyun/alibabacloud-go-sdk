// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCenChildInstanceRouteEntriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCenChildInstanceRouteEntriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCenChildInstanceRouteEntriesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCenChildInstanceRouteEntriesResponseBody) *DescribeCenChildInstanceRouteEntriesResponse
	GetBody() *DescribeCenChildInstanceRouteEntriesResponseBody
}

type DescribeCenChildInstanceRouteEntriesResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCenChildInstanceRouteEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCenChildInstanceRouteEntriesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenChildInstanceRouteEntriesResponse) GoString() string {
	return s.String()
}

func (s *DescribeCenChildInstanceRouteEntriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCenChildInstanceRouteEntriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCenChildInstanceRouteEntriesResponse) GetBody() *DescribeCenChildInstanceRouteEntriesResponseBody {
	return s.Body
}

func (s *DescribeCenChildInstanceRouteEntriesResponse) SetHeaders(v map[string]*string) *DescribeCenChildInstanceRouteEntriesResponse {
	s.Headers = v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponse) SetStatusCode(v int32) *DescribeCenChildInstanceRouteEntriesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponse) SetBody(v *DescribeCenChildInstanceRouteEntriesResponseBody) *DescribeCenChildInstanceRouteEntriesResponse {
	s.Body = v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
