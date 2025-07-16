// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPartitionSummariesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPartitionSummariesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPartitionSummariesResponse
	GetStatusCode() *int32
	SetBody(v *PartitionSummaries) *ListPartitionSummariesResponse
	GetBody() *PartitionSummaries
}

type ListPartitionSummariesResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PartitionSummaries `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPartitionSummariesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPartitionSummariesResponse) GoString() string {
	return s.String()
}

func (s *ListPartitionSummariesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPartitionSummariesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPartitionSummariesResponse) GetBody() *PartitionSummaries {
	return s.Body
}

func (s *ListPartitionSummariesResponse) SetHeaders(v map[string]*string) *ListPartitionSummariesResponse {
	s.Headers = v
	return s
}

func (s *ListPartitionSummariesResponse) SetStatusCode(v int32) *ListPartitionSummariesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPartitionSummariesResponse) SetBody(v *PartitionSummaries) *ListPartitionSummariesResponse {
	s.Body = v
	return s
}

func (s *ListPartitionSummariesResponse) Validate() error {
	return dara.Validate(s)
}
