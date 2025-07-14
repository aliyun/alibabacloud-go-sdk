// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeConfigMapRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigMapId(v int64) *DescribeConfigMapRequest
	GetConfigMapId() *int64
}

type DescribeConfigMapRequest struct {
	// The ID of the ConfigMap whose details you want to query. You can call the [ListNamespacedConfigMaps](https://help.aliyun.com/document_detail/176917.html) operation to obtain the ID of a ConfigMap.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ConfigMapId *int64 `json:"ConfigMapId,omitempty" xml:"ConfigMapId,omitempty"`
}

func (s DescribeConfigMapRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeConfigMapRequest) GoString() string {
	return s.String()
}

func (s *DescribeConfigMapRequest) GetConfigMapId() *int64 {
	return s.ConfigMapId
}

func (s *DescribeConfigMapRequest) SetConfigMapId(v int64) *DescribeConfigMapRequest {
	s.ConfigMapId = &v
	return s
}

func (s *DescribeConfigMapRequest) Validate() error {
	return dara.Validate(s)
}
