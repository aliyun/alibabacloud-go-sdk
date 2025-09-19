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
}

type GetInterceptionSummaryRequest struct {
	// The ID of the cluster.
	//
	// > You can call the [DescribeGroupedContainerInstances](https://help.aliyun.com/document_detail/421736.html) operation to query the IDs of clusters.
	//
	// example:
	//
	// c2999***bb61b
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
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

func (s *GetInterceptionSummaryRequest) SetClusterId(v string) *GetInterceptionSummaryRequest {
	s.ClusterId = &v
	return s
}

func (s *GetInterceptionSummaryRequest) Validate() error {
	return dara.Validate(s)
}
