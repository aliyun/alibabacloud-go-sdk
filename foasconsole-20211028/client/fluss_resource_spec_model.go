// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlussResourceSpec interface {
	dara.Model
	String() string
	GoString() string
	SetDiskSizeInGB(v int64) *FlussResourceSpec
	GetDiskSizeInGB() *int64
	SetSlaveModel(v string) *FlussResourceSpec
	GetSlaveModel() *string
	SetSlaveNum(v int64) *FlussResourceSpec
	GetSlaveNum() *int64
	SetTieringPostCu(v int64) *FlussResourceSpec
	GetTieringPostCu() *int64
	SetTieringPreCu(v int64) *FlussResourceSpec
	GetTieringPreCu() *int64
}

type FlussResourceSpec struct {
	// Disk size per node, in GB.
	DiskSizeInGB *int64 `json:"DiskSizeInGB,omitempty" xml:"DiskSizeInGB,omitempty"`
	// Instance type of the slave nodes.
	SlaveModel *string `json:"SlaveModel,omitempty" xml:"SlaveModel,omitempty"`
	// Number of slave nodes.
	SlaveNum *int64 `json:"SlaveNum,omitempty" xml:"SlaveNum,omitempty"`
	// Number of CUs for the post-tiering stage.
	TieringPostCu *int64 `json:"TieringPostCu,omitempty" xml:"TieringPostCu,omitempty"`
	// Number of compute units (CUs) for the pre-tiering stage.
	TieringPreCu *int64 `json:"TieringPreCu,omitempty" xml:"TieringPreCu,omitempty"`
}

func (s FlussResourceSpec) String() string {
	return dara.Prettify(s)
}

func (s FlussResourceSpec) GoString() string {
	return s.String()
}

func (s *FlussResourceSpec) GetDiskSizeInGB() *int64 {
	return s.DiskSizeInGB
}

func (s *FlussResourceSpec) GetSlaveModel() *string {
	return s.SlaveModel
}

func (s *FlussResourceSpec) GetSlaveNum() *int64 {
	return s.SlaveNum
}

func (s *FlussResourceSpec) GetTieringPostCu() *int64 {
	return s.TieringPostCu
}

func (s *FlussResourceSpec) GetTieringPreCu() *int64 {
	return s.TieringPreCu
}

func (s *FlussResourceSpec) SetDiskSizeInGB(v int64) *FlussResourceSpec {
	s.DiskSizeInGB = &v
	return s
}

func (s *FlussResourceSpec) SetSlaveModel(v string) *FlussResourceSpec {
	s.SlaveModel = &v
	return s
}

func (s *FlussResourceSpec) SetSlaveNum(v int64) *FlussResourceSpec {
	s.SlaveNum = &v
	return s
}

func (s *FlussResourceSpec) SetTieringPostCu(v int64) *FlussResourceSpec {
	s.TieringPostCu = &v
	return s
}

func (s *FlussResourceSpec) SetTieringPreCu(v int64) *FlussResourceSpec {
	s.TieringPreCu = &v
	return s
}

func (s *FlussResourceSpec) Validate() error {
	return dara.Validate(s)
}
