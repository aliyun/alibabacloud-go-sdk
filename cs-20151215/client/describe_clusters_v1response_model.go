// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClustersV1Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeClustersV1Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeClustersV1Response
	GetStatusCode() *int32
	SetBody(v *DescribeClustersV1ResponseBody) *DescribeClustersV1Response
	GetBody() *DescribeClustersV1ResponseBody
}

type DescribeClustersV1Response struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeClustersV1ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeClustersV1Response) String() string {
	return dara.Prettify(s)
}

func (s DescribeClustersV1Response) GoString() string {
	return s.String()
}

func (s *DescribeClustersV1Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeClustersV1Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeClustersV1Response) GetBody() *DescribeClustersV1ResponseBody {
	return s.Body
}

func (s *DescribeClustersV1Response) SetHeaders(v map[string]*string) *DescribeClustersV1Response {
	s.Headers = v
	return s
}

func (s *DescribeClustersV1Response) SetStatusCode(v int32) *DescribeClustersV1Response {
	s.StatusCode = &v
	return s
}

func (s *DescribeClustersV1Response) SetBody(v *DescribeClustersV1ResponseBody) *DescribeClustersV1Response {
	s.Body = v
	return s
}

func (s *DescribeClustersV1Response) Validate() error {
	return dara.Validate(s)
}
