// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSyncJobListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSyncJobListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSyncJobListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSyncJobListResponseBody) *DescribeSyncJobListResponse
	GetBody() *DescribeSyncJobListResponseBody
}

type DescribeSyncJobListResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSyncJobListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSyncJobListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSyncJobListResponse) GoString() string {
	return s.String()
}

func (s *DescribeSyncJobListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSyncJobListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSyncJobListResponse) GetBody() *DescribeSyncJobListResponseBody {
	return s.Body
}

func (s *DescribeSyncJobListResponse) SetHeaders(v map[string]*string) *DescribeSyncJobListResponse {
	s.Headers = v
	return s
}

func (s *DescribeSyncJobListResponse) SetStatusCode(v int32) *DescribeSyncJobListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSyncJobListResponse) SetBody(v *DescribeSyncJobListResponseBody) *DescribeSyncJobListResponse {
	s.Body = v
	return s
}

func (s *DescribeSyncJobListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
