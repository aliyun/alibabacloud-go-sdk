// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCrossRegionBackupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCrossRegionBackupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCrossRegionBackupsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCrossRegionBackupsResponseBody) *DescribeCrossRegionBackupsResponse
	GetBody() *DescribeCrossRegionBackupsResponseBody
}

type DescribeCrossRegionBackupsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCrossRegionBackupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCrossRegionBackupsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCrossRegionBackupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeCrossRegionBackupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCrossRegionBackupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCrossRegionBackupsResponse) GetBody() *DescribeCrossRegionBackupsResponseBody {
	return s.Body
}

func (s *DescribeCrossRegionBackupsResponse) SetHeaders(v map[string]*string) *DescribeCrossRegionBackupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeCrossRegionBackupsResponse) SetStatusCode(v int32) *DescribeCrossRegionBackupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCrossRegionBackupsResponse) SetBody(v *DescribeCrossRegionBackupsResponseBody) *DescribeCrossRegionBackupsResponse {
	s.Body = v
	return s
}

func (s *DescribeCrossRegionBackupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
