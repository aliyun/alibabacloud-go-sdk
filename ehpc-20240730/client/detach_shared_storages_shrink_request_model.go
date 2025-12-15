// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachSharedStoragesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DetachSharedStoragesShrinkRequest
	GetClusterId() *string
	SetSharedStoragesShrink(v string) *DetachSharedStoragesShrinkRequest
	GetSharedStoragesShrink() *string
}

type DetachSharedStoragesShrinkRequest struct {
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
	// The information about mounted shared storage resources.
	//
	// This parameter is required.
	SharedStoragesShrink *string `json:"SharedStorages,omitempty" xml:"SharedStorages,omitempty"`
}

func (s DetachSharedStoragesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachSharedStoragesShrinkRequest) GoString() string {
	return s.String()
}

func (s *DetachSharedStoragesShrinkRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DetachSharedStoragesShrinkRequest) GetSharedStoragesShrink() *string {
	return s.SharedStoragesShrink
}

func (s *DetachSharedStoragesShrinkRequest) SetClusterId(v string) *DetachSharedStoragesShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *DetachSharedStoragesShrinkRequest) SetSharedStoragesShrink(v string) *DetachSharedStoragesShrinkRequest {
	s.SharedStoragesShrink = &v
	return s
}

func (s *DetachSharedStoragesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
