// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageScanNumInPeriodResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetImageScanNumInPeriodResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetImageScanNumInPeriodResponse
	GetStatusCode() *int32
	SetBody(v *GetImageScanNumInPeriodResponseBody) *GetImageScanNumInPeriodResponse
	GetBody() *GetImageScanNumInPeriodResponseBody
}

type GetImageScanNumInPeriodResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetImageScanNumInPeriodResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetImageScanNumInPeriodResponse) String() string {
	return dara.Prettify(s)
}

func (s GetImageScanNumInPeriodResponse) GoString() string {
	return s.String()
}

func (s *GetImageScanNumInPeriodResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetImageScanNumInPeriodResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetImageScanNumInPeriodResponse) GetBody() *GetImageScanNumInPeriodResponseBody {
	return s.Body
}

func (s *GetImageScanNumInPeriodResponse) SetHeaders(v map[string]*string) *GetImageScanNumInPeriodResponse {
	s.Headers = v
	return s
}

func (s *GetImageScanNumInPeriodResponse) SetStatusCode(v int32) *GetImageScanNumInPeriodResponse {
	s.StatusCode = &v
	return s
}

func (s *GetImageScanNumInPeriodResponse) SetBody(v *GetImageScanNumInPeriodResponseBody) *GetImageScanNumInPeriodResponse {
	s.Body = v
	return s
}

func (s *GetImageScanNumInPeriodResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
