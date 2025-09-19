// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRuleTargetAllRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ListRuleTargetAllRequest
	GetClusterId() *string
}

type ListRuleTargetAllRequest struct {
	// The ID of the container cluster.
	//
	// > You can call the [DescribeGroupedContainerInstances](https://help.aliyun.com/document_detail/182997.html) operation to query the IDs of container clusters.
	//
	// This parameter is required.
	//
	// example:
	//
	// cfa7e2fb8c221483ba59e098c34c6****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s ListRuleTargetAllRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRuleTargetAllRequest) GoString() string {
	return s.String()
}

func (s *ListRuleTargetAllRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListRuleTargetAllRequest) SetClusterId(v string) *ListRuleTargetAllRequest {
	s.ClusterId = &v
	return s
}

func (s *ListRuleTargetAllRequest) Validate() error {
	return dara.Validate(s)
}
