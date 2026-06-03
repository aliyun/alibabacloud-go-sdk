// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSyncInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSyncInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSyncInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSyncInfoResponseBody) *DescribeSyncInfoResponse
	GetBody() *DescribeSyncInfoResponseBody
}

type DescribeSyncInfoResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSyncInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSyncInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSyncInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeSyncInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSyncInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSyncInfoResponse) GetBody() *DescribeSyncInfoResponseBody {
	return s.Body
}

func (s *DescribeSyncInfoResponse) SetHeaders(v map[string]*string) *DescribeSyncInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeSyncInfoResponse) SetStatusCode(v int32) *DescribeSyncInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSyncInfoResponse) SetBody(v *DescribeSyncInfoResponseBody) *DescribeSyncInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeSyncInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
