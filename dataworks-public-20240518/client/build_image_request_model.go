// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBuildImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCu(v float64) *BuildImageRequest
	GetCu() *float64
	SetId(v string) *BuildImageRequest
	GetId() *string
	SetProcessId(v string) *BuildImageRequest
	GetProcessId() *string
	SetResourceGroupId(v string) *BuildImageRequest
	GetResourceGroupId() *string
}

type BuildImageRequest struct {
	// The number of compute units (CUs) used for the build.
	//
	// example:
	//
	// 0.5
	Cu *float64 `json:"Cu,omitempty" xml:"Cu,omitempty"`
	// The image ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// Custom_image_xxxx_xxxx
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The image build execution ID, which is used as an idempotence identifier.
	//
	// example:
	//
	// 582d4896-d224-413b-b883-239eeebe0bc5
	ProcessId *string `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
	// The unique identifier of the general-purpose resource group.
	//
	// This parameter is required.
	//
	// example:
	//
	// Serverless_res_group_****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s BuildImageRequest) String() string {
	return dara.Prettify(s)
}

func (s BuildImageRequest) GoString() string {
	return s.String()
}

func (s *BuildImageRequest) GetCu() *float64 {
	return s.Cu
}

func (s *BuildImageRequest) GetId() *string {
	return s.Id
}

func (s *BuildImageRequest) GetProcessId() *string {
	return s.ProcessId
}

func (s *BuildImageRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *BuildImageRequest) SetCu(v float64) *BuildImageRequest {
	s.Cu = &v
	return s
}

func (s *BuildImageRequest) SetId(v string) *BuildImageRequest {
	s.Id = &v
	return s
}

func (s *BuildImageRequest) SetProcessId(v string) *BuildImageRequest {
	s.ProcessId = &v
	return s
}

func (s *BuildImageRequest) SetResourceGroupId(v string) *BuildImageRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *BuildImageRequest) Validate() error {
	return dara.Validate(s)
}
