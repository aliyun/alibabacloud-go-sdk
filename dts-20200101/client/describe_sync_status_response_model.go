// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSyncStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSyncStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSyncStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSyncStatusResponseBody) *DescribeSyncStatusResponse
	GetBody() *DescribeSyncStatusResponseBody
}

type DescribeSyncStatusResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSyncStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSyncStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSyncStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeSyncStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSyncStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSyncStatusResponse) GetBody() *DescribeSyncStatusResponseBody {
	return s.Body
}

func (s *DescribeSyncStatusResponse) SetHeaders(v map[string]*string) *DescribeSyncStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeSyncStatusResponse) SetStatusCode(v int32) *DescribeSyncStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSyncStatusResponse) SetBody(v *DescribeSyncStatusResponseBody) *DescribeSyncStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeSyncStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
