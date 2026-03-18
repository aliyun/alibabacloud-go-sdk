// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryModifyDiskPerformanceLevelPriceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryModifyDiskPerformanceLevelPriceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryModifyDiskPerformanceLevelPriceResponse
	GetStatusCode() *int32
	SetBody(v *QueryModifyDiskPerformanceLevelPriceResponseBody) *QueryModifyDiskPerformanceLevelPriceResponse
	GetBody() *QueryModifyDiskPerformanceLevelPriceResponseBody
}

type QueryModifyDiskPerformanceLevelPriceResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryModifyDiskPerformanceLevelPriceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryModifyDiskPerformanceLevelPriceResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryModifyDiskPerformanceLevelPriceResponse) GoString() string {
	return s.String()
}

func (s *QueryModifyDiskPerformanceLevelPriceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryModifyDiskPerformanceLevelPriceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryModifyDiskPerformanceLevelPriceResponse) GetBody() *QueryModifyDiskPerformanceLevelPriceResponseBody {
	return s.Body
}

func (s *QueryModifyDiskPerformanceLevelPriceResponse) SetHeaders(v map[string]*string) *QueryModifyDiskPerformanceLevelPriceResponse {
	s.Headers = v
	return s
}

func (s *QueryModifyDiskPerformanceLevelPriceResponse) SetStatusCode(v int32) *QueryModifyDiskPerformanceLevelPriceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryModifyDiskPerformanceLevelPriceResponse) SetBody(v *QueryModifyDiskPerformanceLevelPriceResponseBody) *QueryModifyDiskPerformanceLevelPriceResponse {
	s.Body = v
	return s
}

func (s *QueryModifyDiskPerformanceLevelPriceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
