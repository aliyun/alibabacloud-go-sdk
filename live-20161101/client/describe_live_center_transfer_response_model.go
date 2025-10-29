// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveCenterTransferResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveCenterTransferResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveCenterTransferResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveCenterTransferResponseBody) *DescribeLiveCenterTransferResponse
	GetBody() *DescribeLiveCenterTransferResponseBody
}

type DescribeLiveCenterTransferResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveCenterTransferResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveCenterTransferResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveCenterTransferResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveCenterTransferResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveCenterTransferResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveCenterTransferResponse) GetBody() *DescribeLiveCenterTransferResponseBody {
	return s.Body
}

func (s *DescribeLiveCenterTransferResponse) SetHeaders(v map[string]*string) *DescribeLiveCenterTransferResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveCenterTransferResponse) SetStatusCode(v int32) *DescribeLiveCenterTransferResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveCenterTransferResponse) SetBody(v *DescribeLiveCenterTransferResponseBody) *DescribeLiveCenterTransferResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveCenterTransferResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
