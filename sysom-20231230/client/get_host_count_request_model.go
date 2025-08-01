// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHostCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCluster(v string) *GetHostCountRequest
	GetCluster() *string
	SetEnd(v float32) *GetHostCountRequest
	GetEnd() *float32
	SetInstance(v string) *GetHostCountRequest
	GetInstance() *string
	SetStart(v float32) *GetHostCountRequest
	GetStart() *float32
}

type GetHostCountRequest struct {
	// example:
	//
	// 1808078950770264
	Cluster *string `json:"cluster,omitempty" xml:"cluster,omitempty"`
	// example:
	//
	// 1725801327754
	End *float32 `json:"end,omitempty" xml:"end,omitempty"`
	// example:
	//
	// i-wz9d00ut2ska3mlyhn6j
	Instance *string `json:"instance,omitempty" xml:"instance,omitempty"`
	// example:
	//
	// 1725797727754
	Start *float32 `json:"start,omitempty" xml:"start,omitempty"`
}

func (s GetHostCountRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHostCountRequest) GoString() string {
	return s.String()
}

func (s *GetHostCountRequest) GetCluster() *string {
	return s.Cluster
}

func (s *GetHostCountRequest) GetEnd() *float32 {
	return s.End
}

func (s *GetHostCountRequest) GetInstance() *string {
	return s.Instance
}

func (s *GetHostCountRequest) GetStart() *float32 {
	return s.Start
}

func (s *GetHostCountRequest) SetCluster(v string) *GetHostCountRequest {
	s.Cluster = &v
	return s
}

func (s *GetHostCountRequest) SetEnd(v float32) *GetHostCountRequest {
	s.End = &v
	return s
}

func (s *GetHostCountRequest) SetInstance(v string) *GetHostCountRequest {
	s.Instance = &v
	return s
}

func (s *GetHostCountRequest) SetStart(v float32) *GetHostCountRequest {
	s.Start = &v
	return s
}

func (s *GetHostCountRequest) Validate() error {
	return dara.Validate(s)
}
