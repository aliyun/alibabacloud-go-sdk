// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainSnapshotDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveDomainSnapshotDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveDomainSnapshotDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveDomainSnapshotDataResponseBody) *DescribeLiveDomainSnapshotDataResponse
	GetBody() *DescribeLiveDomainSnapshotDataResponseBody
}

type DescribeLiveDomainSnapshotDataResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveDomainSnapshotDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveDomainSnapshotDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainSnapshotDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainSnapshotDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveDomainSnapshotDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveDomainSnapshotDataResponse) GetBody() *DescribeLiveDomainSnapshotDataResponseBody {
	return s.Body
}

func (s *DescribeLiveDomainSnapshotDataResponse) SetHeaders(v map[string]*string) *DescribeLiveDomainSnapshotDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveDomainSnapshotDataResponse) SetStatusCode(v int32) *DescribeLiveDomainSnapshotDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveDomainSnapshotDataResponse) SetBody(v *DescribeLiveDomainSnapshotDataResponseBody) *DescribeLiveDomainSnapshotDataResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveDomainSnapshotDataResponse) Validate() error {
	return dara.Validate(s)
}
