// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOutgoingAssetListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeOutgoingAssetListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeOutgoingAssetListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeOutgoingAssetListResponseBody) *DescribeOutgoingAssetListResponse
	GetBody() *DescribeOutgoingAssetListResponseBody
}

type DescribeOutgoingAssetListResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOutgoingAssetListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOutgoingAssetListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeOutgoingAssetListResponse) GoString() string {
	return s.String()
}

func (s *DescribeOutgoingAssetListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeOutgoingAssetListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeOutgoingAssetListResponse) GetBody() *DescribeOutgoingAssetListResponseBody {
	return s.Body
}

func (s *DescribeOutgoingAssetListResponse) SetHeaders(v map[string]*string) *DescribeOutgoingAssetListResponse {
	s.Headers = v
	return s
}

func (s *DescribeOutgoingAssetListResponse) SetStatusCode(v int32) *DescribeOutgoingAssetListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOutgoingAssetListResponse) SetBody(v *DescribeOutgoingAssetListResponseBody) *DescribeOutgoingAssetListResponse {
	s.Body = v
	return s
}

func (s *DescribeOutgoingAssetListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
