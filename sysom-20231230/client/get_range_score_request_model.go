// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRangeScoreRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCluster(v string) *GetRangeScoreRequest
	GetCluster() *string
	SetEnd(v float32) *GetRangeScoreRequest
	GetEnd() *float32
	SetInstance(v string) *GetRangeScoreRequest
	GetInstance() *string
	SetStart(v float32) *GetRangeScoreRequest
	GetStart() *float32
}

type GetRangeScoreRequest struct {
	// example:
	//
	// 1808078950770264
	Cluster *string `json:"cluster,omitempty" xml:"cluster,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1725801327754
	End *float32 `json:"end,omitempty" xml:"end,omitempty"`
	// example:
	//
	// i-wz9d00ut2ska3mlyhn6j
	Instance *string `json:"instance,omitempty" xml:"instance,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1725797727754
	Start *float32 `json:"start,omitempty" xml:"start,omitempty"`
}

func (s GetRangeScoreRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRangeScoreRequest) GoString() string {
	return s.String()
}

func (s *GetRangeScoreRequest) GetCluster() *string {
	return s.Cluster
}

func (s *GetRangeScoreRequest) GetEnd() *float32 {
	return s.End
}

func (s *GetRangeScoreRequest) GetInstance() *string {
	return s.Instance
}

func (s *GetRangeScoreRequest) GetStart() *float32 {
	return s.Start
}

func (s *GetRangeScoreRequest) SetCluster(v string) *GetRangeScoreRequest {
	s.Cluster = &v
	return s
}

func (s *GetRangeScoreRequest) SetEnd(v float32) *GetRangeScoreRequest {
	s.End = &v
	return s
}

func (s *GetRangeScoreRequest) SetInstance(v string) *GetRangeScoreRequest {
	s.Instance = &v
	return s
}

func (s *GetRangeScoreRequest) SetStart(v float32) *GetRangeScoreRequest {
	s.Start = &v
	return s
}

func (s *GetRangeScoreRequest) Validate() error {
	return dara.Validate(s)
}
