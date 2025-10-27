// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSyncAvailableDBClusterListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSyncAvailableDBClusterListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSyncAvailableDBClusterListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSyncAvailableDBClusterListResponseBody) *DescribeSyncAvailableDBClusterListResponse
	GetBody() *DescribeSyncAvailableDBClusterListResponseBody
}

type DescribeSyncAvailableDBClusterListResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSyncAvailableDBClusterListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSyncAvailableDBClusterListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSyncAvailableDBClusterListResponse) GoString() string {
	return s.String()
}

func (s *DescribeSyncAvailableDBClusterListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSyncAvailableDBClusterListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSyncAvailableDBClusterListResponse) GetBody() *DescribeSyncAvailableDBClusterListResponseBody {
	return s.Body
}

func (s *DescribeSyncAvailableDBClusterListResponse) SetHeaders(v map[string]*string) *DescribeSyncAvailableDBClusterListResponse {
	s.Headers = v
	return s
}

func (s *DescribeSyncAvailableDBClusterListResponse) SetStatusCode(v int32) *DescribeSyncAvailableDBClusterListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSyncAvailableDBClusterListResponse) SetBody(v *DescribeSyncAvailableDBClusterListResponseBody) *DescribeSyncAvailableDBClusterListResponse {
	s.Body = v
	return s
}

func (s *DescribeSyncAvailableDBClusterListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
