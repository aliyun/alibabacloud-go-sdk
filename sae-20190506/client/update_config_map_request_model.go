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
	// The ID of the ConfigMap instance that you want to update. To view the ID, call the [ListNamespacedConfigMaps](https://help.aliyun.com/document_detail/176917.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ConfigMapId *int64 `json:"ConfigMapId,omitempty" xml:"ConfigMapId,omitempty"`
	// The key-value pairs for the ConfigMap. The value must be a JSON-formatted string, as shown in the following example:
	//
	// {"Data":"{"k1":"v1", "k2":"v2"}"}
	//
	// In the JSON string, k represents a key and v represents a value. For more information about configuration items, see [Managing and using configuration items](https://help.aliyun.com/document_detail/171326.html).
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
