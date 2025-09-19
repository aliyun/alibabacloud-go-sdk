// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClusterRuleSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *GetClusterRuleSummaryRequest
	GetClusterId() *string
}

type GetClusterRuleSummaryRequest struct {
	// The ID of the container cluster.
	//
	// >  You can call the [DescribeGroupedContainerInstances](~~DescribeGroupedContainerInstances~~) operation to query the IDs of container clusters.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc49d88d1exxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s GetClusterRuleSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetClusterRuleSummaryRequest) GoString() string {
	return s.String()
}

func (s *GetClusterRuleSummaryRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetClusterRuleSummaryRequest) SetClusterId(v string) *GetClusterRuleSummaryRequest {
	s.ClusterId = &v
	return s
}

func (s *GetClusterRuleSummaryRequest) Validate() error {
	return dara.Validate(s)
}
