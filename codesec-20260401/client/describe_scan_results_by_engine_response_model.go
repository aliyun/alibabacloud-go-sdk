// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScanResultsByEngineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeScanResultsByEngineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeScanResultsByEngineResponse
	GetStatusCode() *int32
	SetBody(v *DescribeScanResultsByEngineResponseBody) *DescribeScanResultsByEngineResponse
	GetBody() *DescribeScanResultsByEngineResponseBody
}

type DescribeScanResultsByEngineResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeScanResultsByEngineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeScanResultsByEngineResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeScanResultsByEngineResponse) GoString() string {
	return s.String()
}

func (s *DescribeScanResultsByEngineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeScanResultsByEngineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeScanResultsByEngineResponse) GetBody() *DescribeScanResultsByEngineResponseBody {
	return s.Body
}

func (s *DescribeScanResultsByEngineResponse) SetHeaders(v map[string]*string) *DescribeScanResultsByEngineResponse {
	s.Headers = v
	return s
}

func (s *DescribeScanResultsByEngineResponse) SetStatusCode(v int32) *DescribeScanResultsByEngineResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScanResultsByEngineResponse) SetBody(v *DescribeScanResultsByEngineResponseBody) *DescribeScanResultsByEngineResponse {
	s.Body = v
	return s
}

func (s *DescribeScanResultsByEngineResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
