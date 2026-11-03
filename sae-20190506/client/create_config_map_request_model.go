// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConfigMapRequest interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *CreateConfigMapRequest
	GetData() *string
	SetDescription(v string) *CreateConfigMapRequest
	GetDescription() *string
	SetName(v string) *CreateConfigMapRequest
	GetName() *string
	SetNamespaceId(v string) *CreateConfigMapRequest
	GetNamespaceId() *string
}

type CreateConfigMapRequest struct {
	// The key-value pair data of the ConfigMap in JSON format. The format is as follows:
	//
	// {"Data":"{\\"k1\\":\\"v1\\", \\"k2\\":\\"v2\\"}"}
	//
	// k represents the key, and v represents the value. For more information about ConfigMaps, see [Manage and use ConfigMaps](https://help.aliyun.com/document_detail/171326.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// {"env.shell": "/bin/sh"}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The description. The description cannot contain spaces and can be up to 255 characters in length.
	//
	// example:
	//
	// test-desc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the ConfigMap instance. The name can contain digits, letters, and underscores (_), and must start with a letter.
	//
	// This parameter is required.
	//
	// example:
	//
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the namespace where the ConfigMap instance resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
}

func (s CreateConfigMapRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateConfigMapRequest) GoString() string {
	return s.String()
}

func (s *CreateConfigMapRequest) GetData() *string {
	return s.Data
}

func (s *CreateConfigMapRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateConfigMapRequest) GetName() *string {
	return s.Name
}

func (s *CreateConfigMapRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *CreateConfigMapRequest) SetData(v string) *CreateConfigMapRequest {
	s.Data = &v
	return s
}

func (s *CreateConfigMapRequest) SetDescription(v string) *CreateConfigMapRequest {
	s.Description = &v
	return s
}

func (s *CreateConfigMapRequest) SetName(v string) *CreateConfigMapRequest {
	s.Name = &v
	return s
}

func (s *CreateConfigMapRequest) SetNamespaceId(v string) *CreateConfigMapRequest {
	s.NamespaceId = &v
	return s
}

func (s *CreateConfigMapRequest) Validate() error {
	return dara.Validate(s)
}
