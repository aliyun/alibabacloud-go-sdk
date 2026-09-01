// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInterceptionSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *GetInterceptionSummaryRequest
	GetClusterId() *string
	SetExcludeClusterTypes(v []*string) *GetInterceptionSummaryRequest
	GetExcludeClusterTypes() []*string
}

type GetInterceptionSummaryRequest struct {
	// The ID of the cluster to query. This parameter takes effect only on the InterceptionCountInDays response parameter.
	//
	// > You can call the [DescribeGroupedContainerInstances](~~DescribeGroupedContainerInstances~~) operation to obtain this parameter.
	//
	// example:
	//
	// c2999***bb61b
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The list of cluster types to exclude.
	ExcludeClusterTypes []*string `json:"ExcludeClusterTypes,omitempty" xml:"ExcludeClusterTypes,omitempty" type:"Repeated"`
}

func (s GetInterceptionSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInterceptionSummaryRequest) GoString() string {
	return s.String()
}

func (s *GetInterceptionSummaryRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetInterceptionSummaryRequest) GetExcludeClusterTypes() []*string {
	return s.ExcludeClusterTypes
}

func (s *GetInterceptionSummaryRequest) SetClusterId(v string) *GetInterceptionSummaryRequest {
	s.ClusterId = &v
	return s
}

func (s *GetInterceptionSummaryRequest) SetExcludeClusterTypes(v []*string) *GetInterceptionSummaryRequest {
	s.ExcludeClusterTypes = v
	return s
}

func (s *GetInterceptionSummaryRequest) Validate() error {
	return dara.Validate(s)
}
