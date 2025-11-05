// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiskReplicaPairsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDiskReplicaPairsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDiskReplicaPairsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDiskReplicaPairsResponseBody) *DescribeDiskReplicaPairsResponse
	GetBody() *DescribeDiskReplicaPairsResponseBody
}

type DescribeDiskReplicaPairsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDiskReplicaPairsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDiskReplicaPairsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiskReplicaPairsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDiskReplicaPairsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDiskReplicaPairsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDiskReplicaPairsResponse) GetBody() *DescribeDiskReplicaPairsResponseBody {
	return s.Body
}

func (s *DescribeDiskReplicaPairsResponse) SetHeaders(v map[string]*string) *DescribeDiskReplicaPairsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDiskReplicaPairsResponse) SetStatusCode(v int32) *DescribeDiskReplicaPairsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponse) SetBody(v *DescribeDiskReplicaPairsResponseBody) *DescribeDiskReplicaPairsResponse {
	s.Body = v
	return s
}

func (s *DescribeDiskReplicaPairsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
