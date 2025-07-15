// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDetectNotifyConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveDetectNotifyConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveDetectNotifyConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveDetectNotifyConfigResponseBody) *DescribeLiveDetectNotifyConfigResponse
	GetBody() *DescribeLiveDetectNotifyConfigResponseBody
}

type DescribeLiveDetectNotifyConfigResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveDetectNotifyConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveDetectNotifyConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDetectNotifyConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveDetectNotifyConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveDetectNotifyConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveDetectNotifyConfigResponse) GetBody() *DescribeLiveDetectNotifyConfigResponseBody {
	return s.Body
}

func (s *DescribeLiveDetectNotifyConfigResponse) SetHeaders(v map[string]*string) *DescribeLiveDetectNotifyConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveDetectNotifyConfigResponse) SetStatusCode(v int32) *DescribeLiveDetectNotifyConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveDetectNotifyConfigResponse) SetBody(v *DescribeLiveDetectNotifyConfigResponseBody) *DescribeLiveDetectNotifyConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveDetectNotifyConfigResponse) Validate() error {
	return dara.Validate(s)
}
