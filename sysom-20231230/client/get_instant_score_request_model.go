// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstantScoreRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCluster(v string) *GetInstantScoreRequest
	GetCluster() *string
	SetInstance(v string) *GetInstantScoreRequest
	GetInstance() *string
}

type GetInstantScoreRequest struct {
	// example:
	//
	// 1808078950770264
	Cluster *string `json:"cluster,omitempty" xml:"cluster,omitempty"`
	// example:
	//
	// i-wz9d00ut2ska3mlyhn6j
	Instance *string `json:"instance,omitempty" xml:"instance,omitempty"`
}

func (s GetInstantScoreRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInstantScoreRequest) GoString() string {
	return s.String()
}

func (s *GetInstantScoreRequest) GetCluster() *string {
	return s.Cluster
}

func (s *GetInstantScoreRequest) GetInstance() *string {
	return s.Instance
}

func (s *GetInstantScoreRequest) SetCluster(v string) *GetInstantScoreRequest {
	s.Cluster = &v
	return s
}

func (s *GetInstantScoreRequest) SetInstance(v string) *GetInstantScoreRequest {
	s.Instance = &v
	return s
}

func (s *GetInstantScoreRequest) Validate() error {
	return dara.Validate(s)
}
