// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteConfigMapRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigMapId(v int64) *DeleteConfigMapRequest
	GetConfigMapId() *int64
}

type DeleteConfigMapRequest struct {
	// The ID of the ConfigMap that you want to delete. You can call the [ListNamespacedConfigMaps](https://help.aliyun.com/document_detail/176917.html) operation to obtain the ID of a ConfigMap.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ConfigMapId *int64 `json:"ConfigMapId,omitempty" xml:"ConfigMapId,omitempty"`
}

func (s DeleteConfigMapRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteConfigMapRequest) GoString() string {
	return s.String()
}

func (s *DeleteConfigMapRequest) GetConfigMapId() *int64 {
	return s.ConfigMapId
}

func (s *DeleteConfigMapRequest) SetConfigMapId(v int64) *DeleteConfigMapRequest {
	s.ConfigMapId = &v
	return s
}

func (s *DeleteConfigMapRequest) Validate() error {
	return dara.Validate(s)
}
