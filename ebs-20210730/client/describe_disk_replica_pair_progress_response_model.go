// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiskReplicaPairProgressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDiskReplicaPairProgressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDiskReplicaPairProgressResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDiskReplicaPairProgressResponseBody) *DescribeDiskReplicaPairProgressResponse
	GetBody() *DescribeDiskReplicaPairProgressResponseBody
}

type DescribeDiskReplicaPairProgressResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDiskReplicaPairProgressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDiskReplicaPairProgressResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiskReplicaPairProgressResponse) GoString() string {
	return s.String()
}

func (s *DescribeDiskReplicaPairProgressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDiskReplicaPairProgressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDiskReplicaPairProgressResponse) GetBody() *DescribeDiskReplicaPairProgressResponseBody {
	return s.Body
}

func (s *DescribeDiskReplicaPairProgressResponse) SetHeaders(v map[string]*string) *DescribeDiskReplicaPairProgressResponse {
	s.Headers = v
	return s
}

func (s *DescribeDiskReplicaPairProgressResponse) SetStatusCode(v int32) *DescribeDiskReplicaPairProgressResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDiskReplicaPairProgressResponse) SetBody(v *DescribeDiskReplicaPairProgressResponseBody) *DescribeDiskReplicaPairProgressResponse {
	s.Body = v
	return s
}

func (s *DescribeDiskReplicaPairProgressResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
