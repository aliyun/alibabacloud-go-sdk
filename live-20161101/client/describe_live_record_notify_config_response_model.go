// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveRecordNotifyConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveRecordNotifyConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveRecordNotifyConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveRecordNotifyConfigResponseBody) *DescribeLiveRecordNotifyConfigResponse
	GetBody() *DescribeLiveRecordNotifyConfigResponseBody
}

type DescribeLiveRecordNotifyConfigResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveRecordNotifyConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveRecordNotifyConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveRecordNotifyConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveRecordNotifyConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveRecordNotifyConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveRecordNotifyConfigResponse) GetBody() *DescribeLiveRecordNotifyConfigResponseBody {
	return s.Body
}

func (s *DescribeLiveRecordNotifyConfigResponse) SetHeaders(v map[string]*string) *DescribeLiveRecordNotifyConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveRecordNotifyConfigResponse) SetStatusCode(v int32) *DescribeLiveRecordNotifyConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveRecordNotifyConfigResponse) SetBody(v *DescribeLiveRecordNotifyConfigResponseBody) *DescribeLiveRecordNotifyConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveRecordNotifyConfigResponse) Validate() error {
	return dara.Validate(s)
}
