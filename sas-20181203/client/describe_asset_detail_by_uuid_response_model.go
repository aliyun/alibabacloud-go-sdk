// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAssetDetailByUuidResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAssetDetailByUuidResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAssetDetailByUuidResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAssetDetailByUuidResponseBody) *DescribeAssetDetailByUuidResponse
	GetBody() *DescribeAssetDetailByUuidResponseBody
}

type DescribeAssetDetailByUuidResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAssetDetailByUuidResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAssetDetailByUuidResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAssetDetailByUuidResponse) GoString() string {
	return s.String()
}

func (s *DescribeAssetDetailByUuidResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAssetDetailByUuidResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAssetDetailByUuidResponse) GetBody() *DescribeAssetDetailByUuidResponseBody {
	return s.Body
}

func (s *DescribeAssetDetailByUuidResponse) SetHeaders(v map[string]*string) *DescribeAssetDetailByUuidResponse {
	s.Headers = v
	return s
}

func (s *DescribeAssetDetailByUuidResponse) SetStatusCode(v int32) *DescribeAssetDetailByUuidResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAssetDetailByUuidResponse) SetBody(v *DescribeAssetDetailByUuidResponseBody) *DescribeAssetDetailByUuidResponse {
	s.Body = v
	return s
}

func (s *DescribeAssetDetailByUuidResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
