// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventLevelCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeEventLevelCountRequest
	GetClusterId() *string
	SetContainerFieldName(v string) *DescribeEventLevelCountRequest
	GetContainerFieldName() *string
	SetContainerFieldValue(v string) *DescribeEventLevelCountRequest
	GetContainerFieldValue() *string
	SetContainerIds(v string) *DescribeEventLevelCountRequest
	GetContainerIds() *string
	SetFrom(v string) *DescribeEventLevelCountRequest
	GetFrom() *string
	SetMultiAccountActionType(v int32) *DescribeEventLevelCountRequest
	GetMultiAccountActionType() *int32
	SetTargetType(v string) *DescribeEventLevelCountRequest
	GetTargetType() *string
}

type DescribeEventLevelCountRequest struct {
	// The ID of the container cluster.
	//
	// > You can call the [DescribeGroupedContainerInstances](~~DescribeGroupedContainerInstances~~) operation to query the IDs of container clusters.
	//
	// example:
	//
	// c7e3c5b420a7947c2933303144688****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The key of the condition that is used to query alert events on containers. Valid values:
	//
	// 	- **instanceId**: the ID of the asset
	//
	// 	- **appName**: the name of the application
	//
	// 	- **clusterId**: the ID of the cluster
	//
	// 	- **regionId**: the ID of the region
	//
	// 	- **nodeName**: the name of the node
	//
	// 	- **namespace**: the namespace
	//
	// 	- **clusterName**: the name of the cluster
	//
	// 	- **image**: the name of the image
	//
	// 	- **imageRepoName**: the name of the image repository
	//
	// 	- **imageRepoNamespace**: the namespace to which the image repository belongs
	//
	// 	- **imageRepoTag**: the tag that is added to the image
	//
	// 	- **imageDigest**: the digest of the image
	//
	// example:
	//
	// clusterId
	ContainerFieldName *string `json:"ContainerFieldName,omitempty" xml:"ContainerFieldName,omitempty"`
	// The value of the condition that is used to query alert events on containers. If you specify multiple values, separate them with commas (,).
	//
	// example:
	//
	// c951761046a9c4afe92be0a7b5bexxxxx
	ContainerFieldValue *string `json:"ContainerFieldValue,omitempty" xml:"ContainerFieldValue,omitempty"`
	// The ID of the container.
	//
	// example:
	//
	// xxxxxx30389a10c28f6d38f2398f0dcexxxxxxx922b9e8290dc7c3019d4a8,48b87f2c0662e334820f436cb9133f1ae4e053d39b6fad42xxxxxxxxxx
	ContainerIds *string `json:"ContainerIds,omitempty" xml:"ContainerIds,omitempty"`
	// The ID of the request source. Set the value to **sas**.
	//
	// example:
	//
	// sas
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// The type of the accounts that you want to query. Default value: **0**. Valid values:
	//
	// 	- **0**: the current account.
	//
	// 	- **1**: all accounts.
	//
	// example:
	//
	// 0
	MultiAccountActionType *int32 `json:"MultiAccountActionType,omitempty" xml:"MultiAccountActionType,omitempty"`
	// The type of the query condition. Valid values:
	//
	// 	- **containerId**: the ID of the container
	//
	// 	- **uuid**: the UUID of the asset
	//
	// example:
	//
	// uuid
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s DescribeEventLevelCountRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventLevelCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeEventLevelCountRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeEventLevelCountRequest) GetContainerFieldName() *string {
	return s.ContainerFieldName
}

func (s *DescribeEventLevelCountRequest) GetContainerFieldValue() *string {
	return s.ContainerFieldValue
}

func (s *DescribeEventLevelCountRequest) GetContainerIds() *string {
	return s.ContainerIds
}

func (s *DescribeEventLevelCountRequest) GetFrom() *string {
	return s.From
}

func (s *DescribeEventLevelCountRequest) GetMultiAccountActionType() *int32 {
	return s.MultiAccountActionType
}

func (s *DescribeEventLevelCountRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *DescribeEventLevelCountRequest) SetClusterId(v string) *DescribeEventLevelCountRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeEventLevelCountRequest) SetContainerFieldName(v string) *DescribeEventLevelCountRequest {
	s.ContainerFieldName = &v
	return s
}

func (s *DescribeEventLevelCountRequest) SetContainerFieldValue(v string) *DescribeEventLevelCountRequest {
	s.ContainerFieldValue = &v
	return s
}

func (s *DescribeEventLevelCountRequest) SetContainerIds(v string) *DescribeEventLevelCountRequest {
	s.ContainerIds = &v
	return s
}

func (s *DescribeEventLevelCountRequest) SetFrom(v string) *DescribeEventLevelCountRequest {
	s.From = &v
	return s
}

func (s *DescribeEventLevelCountRequest) SetMultiAccountActionType(v int32) *DescribeEventLevelCountRequest {
	s.MultiAccountActionType = &v
	return s
}

func (s *DescribeEventLevelCountRequest) SetTargetType(v string) *DescribeEventLevelCountRequest {
	s.TargetType = &v
	return s
}

func (s *DescribeEventLevelCountRequest) Validate() error {
	return dara.Validate(s)
}
