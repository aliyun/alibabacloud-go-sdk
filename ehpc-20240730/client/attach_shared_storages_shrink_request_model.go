// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachSharedStoragesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *AttachSharedStoragesShrinkRequest
	GetClusterId() *string
	SetSharedStoragesShrink(v string) *AttachSharedStoragesShrinkRequest
	GetSharedStoragesShrink() *string
}

type AttachSharedStoragesShrinkRequest struct {
	// The cluster ID.
	//
	// You can call the [ListClusters](https://help.aliyun.com/document_detail/87116.html) operation to query the cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The information about the shared storage resources that you want to attach to the cluster.
	//
	// This parameter is required.
	SharedStoragesShrink *string `json:"SharedStorages,omitempty" xml:"SharedStorages,omitempty"`
}

func (s AttachSharedStoragesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachSharedStoragesShrinkRequest) GoString() string {
	return s.String()
}

func (s *AttachSharedStoragesShrinkRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *AttachSharedStoragesShrinkRequest) GetSharedStoragesShrink() *string {
	return s.SharedStoragesShrink
}

func (s *AttachSharedStoragesShrinkRequest) SetClusterId(v string) *AttachSharedStoragesShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *AttachSharedStoragesShrinkRequest) SetSharedStoragesShrink(v string) *AttachSharedStoragesShrinkRequest {
	s.SharedStoragesShrink = &v
	return s
}

func (s *AttachSharedStoragesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
