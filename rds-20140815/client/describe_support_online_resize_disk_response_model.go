// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSupportOnlineResizeDiskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSupportOnlineResizeDiskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSupportOnlineResizeDiskResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSupportOnlineResizeDiskResponseBody) *DescribeSupportOnlineResizeDiskResponse
	GetBody() *DescribeSupportOnlineResizeDiskResponseBody
}

type DescribeSupportOnlineResizeDiskResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSupportOnlineResizeDiskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSupportOnlineResizeDiskResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSupportOnlineResizeDiskResponse) GoString() string {
	return s.String()
}

func (s *DescribeSupportOnlineResizeDiskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSupportOnlineResizeDiskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSupportOnlineResizeDiskResponse) GetBody() *DescribeSupportOnlineResizeDiskResponseBody {
	return s.Body
}

func (s *DescribeSupportOnlineResizeDiskResponse) SetHeaders(v map[string]*string) *DescribeSupportOnlineResizeDiskResponse {
	s.Headers = v
	return s
}

func (s *DescribeSupportOnlineResizeDiskResponse) SetStatusCode(v int32) *DescribeSupportOnlineResizeDiskResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSupportOnlineResizeDiskResponse) SetBody(v *DescribeSupportOnlineResizeDiskResponseBody) *DescribeSupportOnlineResizeDiskResponse {
	s.Body = v
	return s
}

func (s *DescribeSupportOnlineResizeDiskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
