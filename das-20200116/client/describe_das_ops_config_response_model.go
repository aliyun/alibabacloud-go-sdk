// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDasOpsConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDasOpsConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDasOpsConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDasOpsConfigResponseBody) *DescribeDasOpsConfigResponse
	GetBody() *DescribeDasOpsConfigResponseBody
}

type DescribeDasOpsConfigResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDasOpsConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDasOpsConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDasOpsConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeDasOpsConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDasOpsConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDasOpsConfigResponse) GetBody() *DescribeDasOpsConfigResponseBody {
	return s.Body
}

func (s *DescribeDasOpsConfigResponse) SetHeaders(v map[string]*string) *DescribeDasOpsConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeDasOpsConfigResponse) SetStatusCode(v int32) *DescribeDasOpsConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDasOpsConfigResponse) SetBody(v *DescribeDasOpsConfigResponseBody) *DescribeDasOpsConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeDasOpsConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
