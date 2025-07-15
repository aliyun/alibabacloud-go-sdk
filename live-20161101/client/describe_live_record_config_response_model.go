// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveRecordConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveRecordConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveRecordConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveRecordConfigResponseBody) *DescribeLiveRecordConfigResponse
	GetBody() *DescribeLiveRecordConfigResponseBody
}

type DescribeLiveRecordConfigResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveRecordConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveRecordConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveRecordConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveRecordConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveRecordConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveRecordConfigResponse) GetBody() *DescribeLiveRecordConfigResponseBody {
	return s.Body
}

func (s *DescribeLiveRecordConfigResponse) SetHeaders(v map[string]*string) *DescribeLiveRecordConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveRecordConfigResponse) SetStatusCode(v int32) *DescribeLiveRecordConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveRecordConfigResponse) SetBody(v *DescribeLiveRecordConfigResponseBody) *DescribeLiveRecordConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveRecordConfigResponse) Validate() error {
	return dara.Validate(s)
}
