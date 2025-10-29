// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveMessageGroupBandResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveMessageGroupBandResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveMessageGroupBandResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveMessageGroupBandResponseBody) *DescribeLiveMessageGroupBandResponse
	GetBody() *DescribeLiveMessageGroupBandResponseBody
}

type DescribeLiveMessageGroupBandResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveMessageGroupBandResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveMessageGroupBandResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveMessageGroupBandResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveMessageGroupBandResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveMessageGroupBandResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveMessageGroupBandResponse) GetBody() *DescribeLiveMessageGroupBandResponseBody {
	return s.Body
}

func (s *DescribeLiveMessageGroupBandResponse) SetHeaders(v map[string]*string) *DescribeLiveMessageGroupBandResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveMessageGroupBandResponse) SetStatusCode(v int32) *DescribeLiveMessageGroupBandResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveMessageGroupBandResponse) SetBody(v *DescribeLiveMessageGroupBandResponseBody) *DescribeLiveMessageGroupBandResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveMessageGroupBandResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
