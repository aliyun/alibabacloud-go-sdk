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
	SetResourceDirectoryAccountId(v int64) *DescribeEventLevelCountRequest
	GetResourceDirectoryAccountId() *int64
	SetTargetType(v string) *DescribeEventLevelCountRequest
	GetTargetType() *string
}

type DescribeEventLevelCountRequest struct {
	// The ID of the container cluster that you want to query.
	//
	// > You can call the [DescribeGroupedContainerInstances](~~DescribeGroupedContainerInstances~~) operation to obtain this parameter.
	//
	// example:
	//
	// c7e3c5b420a7947c2933303144688****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The container search field. Valid values:
	//
	// - **instanceId**: instance ID
	//
	// - **appName**: application name
	//
	// - **clusterId**: cluster ID
	//
	// - **regionId**: region
	//
	// - **nodeName**: node name
	//
	// - **namespace**: namespace
	//
	// - **clusterName**: cluster name
	//
	// - **image**: image name
	//
	// - **imageRepoName**: image repository name
	//
	// - **imageRepoNamespace**: image repository namespace
	//
	// - **imageRepoTag**: image tag
	//
	// - **imageDigest**: image digest
	//
	// example:
	//
	// clusterId
	ContainerFieldName *string `json:"ContainerFieldName,omitempty" xml:"ContainerFieldName,omitempty"`
	// The value of the field that you want to query. Separate multiple values with commas (,).
	//
	// example:
	//
	// c951761046a9c4afe92be0a7b5bexxxxx
	ContainerFieldValue *string `json:"ContainerFieldValue,omitempty" xml:"ContainerFieldValue,omitempty"`
	// The container IDs.
	//
	// example:
	//
	// xxxxxx30389a10c28f6d38f2398f0dcexxxxxxx922b9e8290dc7c3019d4a8,48b87f2c0662e334820f436cb9133f1ae4e053d39b6fad42xxxxxxxxxx
	ContainerIds *string `json:"ContainerIds,omitempty" xml:"ContainerIds,omitempty"`
	// The source identifier of the request. Set this parameter to **sas**.
	//
	// example:
	//
	// sas
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// The multi-account query type. Default value: **0**. Valid values:
	//
	// - **0**: Query data of the current account.
	//
	// - **1**: Query data of all accounts.
	//
	// example:
	//
	// 0
	MultiAccountActionType *int32 `json:"MultiAccountActionType,omitempty" xml:"MultiAccountActionType,omitempty"`
	// The Alibaba Cloud account ID of the member accounts in the resource folder.
	//
	// >You can invoke the [DescribeMonitorAccounts](~~DescribeMonitorAccounts~~) operation to obtain this parameter.
	ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
	// The query type. Valid values:
	//
	// - **containerId**: container ID
	//
	// - **uuid**: asset ID
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

func (s *DescribeEventLevelCountRequest) GetResourceDirectoryAccountId() *int64 {
	return s.ResourceDirectoryAccountId
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

func (s *DescribeEventLevelCountRequest) SetResourceDirectoryAccountId(v int64) *DescribeEventLevelCountRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

func (s *DescribeEventLevelCountRequest) SetTargetType(v string) *DescribeEventLevelCountRequest {
	s.TargetType = &v
	return s
}

func (s *DescribeEventLevelCountRequest) Validate() error {
	return dara.Validate(s)
}
