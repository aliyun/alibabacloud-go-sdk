// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConfigMapRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigMapId(v int64) *UpdateConfigMapRequest
	GetConfigMapId() *int64
	SetData(v string) *UpdateConfigMapRequest
	GetData() *string
	SetDescription(v string) *UpdateConfigMapRequest
	GetDescription() *string
}

type UpdateConfigMapRequest struct {
	// Instance ID of the ConfigMap instance to be updated. To obtain instance ID, call the [ListNamespacedConfigMaps](https://help.aliyun.com/document_detail/176917.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ConfigMapId *int64 `json:"ConfigMapId,omitempty" xml:"ConfigMapId,omitempty"`
	// The key-value pair data of the ConfigMap in JSON format. The format is as follows:
	//
	// {"Data":"{\\"k1\\":\\"v1\\", \\"k2\\":\\"v2\\"}"}
	//
	// k represents the key and v represents the value. For more information about ConfigMaps, see [Manage and use ConfigMaps](https://help.aliyun.com/document_detail/171326.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// {"env.shell": "/bin/sh"}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The description.
	//
	// example:
	//
	// test-desc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s UpdateConfigMapRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateConfigMapRequest) GoString() string {
	return s.String()
}

func (s *UpdateConfigMapRequest) GetConfigMapId() *int64 {
	return s.ConfigMapId
}

func (s *UpdateConfigMapRequest) GetData() *string {
	return s.Data
}

func (s *UpdateConfigMapRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateConfigMapRequest) SetConfigMapId(v int64) *UpdateConfigMapRequest {
	s.ConfigMapId = &v
	return s
}

func (s *UpdateConfigMapRequest) SetData(v string) *UpdateConfigMapRequest {
	s.Data = &v
	return s
}

func (s *UpdateConfigMapRequest) SetDescription(v string) *UpdateConfigMapRequest {
	s.Description = &v
	return s
}

func (s *UpdateConfigMapRequest) Validate() error {
	return dara.Validate(s)
}
