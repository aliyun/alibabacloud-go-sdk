// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLayer interface {
	dara.Model
	String() string
	GoString() string
	SetAcl(v string) *Layer
	GetAcl() *string
	SetCode(v *OutputCodeLocation) *Layer
	GetCode() *OutputCodeLocation
	SetCodeChecksum(v string) *Layer
	GetCodeChecksum() *string
	SetCodeSize(v int64) *Layer
	GetCodeSize() *int64
	SetCompatibleRuntime(v []*string) *Layer
	GetCompatibleRuntime() []*string
	SetCreateTime(v string) *Layer
	GetCreateTime() *string
	SetDescription(v string) *Layer
	GetDescription() *string
	SetLayerName(v string) *Layer
	GetLayerName() *string
	SetLayerVersionArn(v string) *Layer
	GetLayerVersionArn() *string
	SetLicense(v string) *Layer
	GetLicense() *string
	SetVersion(v int32) *Layer
	GetVersion() *int32
}

type Layer struct {
	// example:
	//
	// 0
	Acl  *string             `json:"acl,omitempty" xml:"acl,omitempty"`
	Code *OutputCodeLocation `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// 2825179536350****
	CodeChecksum *string `json:"codeChecksum,omitempty" xml:"codeChecksum,omitempty"`
	// example:
	//
	// 421
	CodeSize          *int64    `json:"codeSize,omitempty" xml:"codeSize,omitempty"`
	CompatibleRuntime []*string `json:"compatibleRuntime" xml:"compatibleRuntime" type:"Repeated"`
	// example:
	//
	// 2023-03-30T11:08:00Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// My first layer
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// MyLayer
	LayerName *string `json:"layerName,omitempty" xml:"layerName,omitempty"`
	// example:
	//
	// acs:fc:cn-beijing:186824xxxxxx:layers/fc_layer/versions/1
	LayerVersionArn *string `json:"layerVersionArn,omitempty" xml:"layerVersionArn,omitempty"`
	// example:
	//
	// Apache
	License *string `json:"license,omitempty" xml:"license,omitempty"`
	// example:
	//
	// 1
	Version *int32 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s Layer) String() string {
	return dara.Prettify(s)
}

func (s Layer) GoString() string {
	return s.String()
}

func (s *Layer) GetAcl() *string {
	return s.Acl
}

func (s *Layer) GetCode() *OutputCodeLocation {
	return s.Code
}

func (s *Layer) GetCodeChecksum() *string {
	return s.CodeChecksum
}

func (s *Layer) GetCodeSize() *int64 {
	return s.CodeSize
}

func (s *Layer) GetCompatibleRuntime() []*string {
	return s.CompatibleRuntime
}

func (s *Layer) GetCreateTime() *string {
	return s.CreateTime
}

func (s *Layer) GetDescription() *string {
	return s.Description
}

func (s *Layer) GetLayerName() *string {
	return s.LayerName
}

func (s *Layer) GetLayerVersionArn() *string {
	return s.LayerVersionArn
}

func (s *Layer) GetLicense() *string {
	return s.License
}

func (s *Layer) GetVersion() *int32 {
	return s.Version
}

func (s *Layer) SetAcl(v string) *Layer {
	s.Acl = &v
	return s
}

func (s *Layer) SetCode(v *OutputCodeLocation) *Layer {
	s.Code = v
	return s
}

func (s *Layer) SetCodeChecksum(v string) *Layer {
	s.CodeChecksum = &v
	return s
}

func (s *Layer) SetCodeSize(v int64) *Layer {
	s.CodeSize = &v
	return s
}

func (s *Layer) SetCompatibleRuntime(v []*string) *Layer {
	s.CompatibleRuntime = v
	return s
}

func (s *Layer) SetCreateTime(v string) *Layer {
	s.CreateTime = &v
	return s
}

func (s *Layer) SetDescription(v string) *Layer {
	s.Description = &v
	return s
}

func (s *Layer) SetLayerName(v string) *Layer {
	s.LayerName = &v
	return s
}

func (s *Layer) SetLayerVersionArn(v string) *Layer {
	s.LayerVersionArn = &v
	return s
}

func (s *Layer) SetLicense(v string) *Layer {
	s.License = &v
	return s
}

func (s *Layer) SetVersion(v int32) *Layer {
	s.Version = &v
	return s
}

func (s *Layer) Validate() error {
	if s.Code != nil {
		if err := s.Code.Validate(); err != nil {
			return err
		}
	}
	return nil
}
