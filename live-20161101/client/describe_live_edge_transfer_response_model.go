// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveEdgeTransferResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveEdgeTransferResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveEdgeTransferResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveEdgeTransferResponseBody) *DescribeLiveEdgeTransferResponse
	GetBody() *DescribeLiveEdgeTransferResponseBody
}

type DescribeLiveEdgeTransferResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveEdgeTransferResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveEdgeTransferResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveEdgeTransferResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveEdgeTransferResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveEdgeTransferResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveEdgeTransferResponse) GetBody() *DescribeLiveEdgeTransferResponseBody {
	return s.Body
}

func (s *DescribeLiveEdgeTransferResponse) SetHeaders(v map[string]*string) *DescribeLiveEdgeTransferResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveEdgeTransferResponse) SetStatusCode(v int32) *DescribeLiveEdgeTransferResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveEdgeTransferResponse) SetBody(v *DescribeLiveEdgeTransferResponseBody) *DescribeLiveEdgeTransferResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveEdgeTransferResponse) Validate() error {
	return dara.Validate(s)
}
