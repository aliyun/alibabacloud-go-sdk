// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSyntheticProbeListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSyntheticProbeListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSyntheticProbeListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSyntheticProbeListResponseBody) *DescribeSyntheticProbeListResponse
	GetBody() *DescribeSyntheticProbeListResponseBody
}

type DescribeSyntheticProbeListResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSyntheticProbeListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSyntheticProbeListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSyntheticProbeListResponse) GoString() string {
	return s.String()
}

func (s *DescribeSyntheticProbeListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSyntheticProbeListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSyntheticProbeListResponse) GetBody() *DescribeSyntheticProbeListResponseBody {
	return s.Body
}

func (s *DescribeSyntheticProbeListResponse) SetHeaders(v map[string]*string) *DescribeSyntheticProbeListResponse {
	s.Headers = v
	return s
}

func (s *DescribeSyntheticProbeListResponse) SetStatusCode(v int32) *DescribeSyntheticProbeListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSyntheticProbeListResponse) SetBody(v *DescribeSyntheticProbeListResponseBody) *DescribeSyntheticProbeListResponse {
	s.Body = v
	return s
}

func (s *DescribeSyntheticProbeListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
