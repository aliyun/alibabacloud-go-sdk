// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseDdosOriginInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ReleaseDdosOriginInstanceRequest
	GetInstanceId() *string
}

type ReleaseDdosOriginInstanceRequest struct {
	// The ID of the Anti-DDoS Origin instance that you want to release.
	//
	// >  You can call the [DescribeInstanceList](https://help.aliyun.com/document_detail/118698.html) operation to query the IDs of all Anti-DDoS Origin instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddosorigin_cn-pe335v7gs01
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ReleaseDdosOriginInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ReleaseDdosOriginInstanceRequest) GoString() string {
	return s.String()
}

func (s *ReleaseDdosOriginInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ReleaseDdosOriginInstanceRequest) SetInstanceId(v string) *ReleaseDdosOriginInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *ReleaseDdosOriginInstanceRequest) Validate() error {
	return dara.Validate(s)
}
