// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchTraceAppByNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *SearchTraceAppByNameRequest
	GetRegionId() *string
	SetTraceAppName(v string) *SearchTraceAppByNameRequest
	GetTraceAppName() *string
}

type SearchTraceAppByNameRequest struct {
	// This parameter is required.
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TraceAppName *string `json:"TraceAppName,omitempty" xml:"TraceAppName,omitempty"`
}

func (s SearchTraceAppByNameRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchTraceAppByNameRequest) GoString() string {
	return s.String()
}

func (s *SearchTraceAppByNameRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SearchTraceAppByNameRequest) GetTraceAppName() *string {
	return s.TraceAppName
}

func (s *SearchTraceAppByNameRequest) SetRegionId(v string) *SearchTraceAppByNameRequest {
	s.RegionId = &v
	return s
}

func (s *SearchTraceAppByNameRequest) SetTraceAppName(v string) *SearchTraceAppByNameRequest {
	s.TraceAppName = &v
	return s
}

func (s *SearchTraceAppByNameRequest) Validate() error {
	return dara.Validate(s)
}
