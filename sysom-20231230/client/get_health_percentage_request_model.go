// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHealthPercentageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCluster(v string) *GetHealthPercentageRequest
	GetCluster() *string
	SetEnd(v float32) *GetHealthPercentageRequest
	GetEnd() *float32
	SetInstance(v string) *GetHealthPercentageRequest
	GetInstance() *string
	SetStart(v float32) *GetHealthPercentageRequest
	GetStart() *float32
}

type GetHealthPercentageRequest struct {
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

func (s GetHealthPercentageRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHealthPercentageRequest) GoString() string {
	return s.String()
}

func (s *GetHealthPercentageRequest) GetCluster() *string {
	return s.Cluster
}

func (s *GetHealthPercentageRequest) GetEnd() *float32 {
	return s.End
}

func (s *GetHealthPercentageRequest) GetInstance() *string {
	return s.Instance
}

func (s *GetHealthPercentageRequest) GetStart() *float32 {
	return s.Start
}

func (s *GetHealthPercentageRequest) SetCluster(v string) *GetHealthPercentageRequest {
	s.Cluster = &v
	return s
}

func (s *GetHealthPercentageRequest) SetEnd(v float32) *GetHealthPercentageRequest {
	s.End = &v
	return s
}

func (s *GetHealthPercentageRequest) SetInstance(v string) *GetHealthPercentageRequest {
	s.Instance = &v
	return s
}

func (s *GetHealthPercentageRequest) SetStart(v float32) *GetHealthPercentageRequest {
	s.Start = &v
	return s
}

func (s *GetHealthPercentageRequest) Validate() error {
	return dara.Validate(s)
}
