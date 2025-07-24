// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClickhouseConf interface {
	dara.Model
	String() string
	GoString() string
	SetInitialReplica(v int32) *ClickhouseConf
	GetInitialReplica() *int32
	SetInitialShard(v int32) *ClickhouseConf
	GetInitialShard() *int32
	SetNewNodeCount(v int32) *ClickhouseConf
	GetNewNodeCount() *int32
	SetResizeType(v string) *ClickhouseConf
	GetResizeType() *string
}

type ClickhouseConf struct {
	InitialReplica *int32  `json:"InitialReplica,omitempty" xml:"InitialReplica,omitempty"`
	InitialShard   *int32  `json:"InitialShard,omitempty" xml:"InitialShard,omitempty"`
	NewNodeCount   *int32  `json:"NewNodeCount,omitempty" xml:"NewNodeCount,omitempty"`
	ResizeType     *string `json:"ResizeType,omitempty" xml:"ResizeType,omitempty"`
}

func (s ClickhouseConf) String() string {
	return dara.Prettify(s)
}

func (s ClickhouseConf) GoString() string {
	return s.String()
}

func (s *ClickhouseConf) GetInitialReplica() *int32 {
	return s.InitialReplica
}

func (s *ClickhouseConf) GetInitialShard() *int32 {
	return s.InitialShard
}

func (s *ClickhouseConf) GetNewNodeCount() *int32 {
	return s.NewNodeCount
}

func (s *ClickhouseConf) GetResizeType() *string {
	return s.ResizeType
}

func (s *ClickhouseConf) SetInitialReplica(v int32) *ClickhouseConf {
	s.InitialReplica = &v
	return s
}

func (s *ClickhouseConf) SetInitialShard(v int32) *ClickhouseConf {
	s.InitialShard = &v
	return s
}

func (s *ClickhouseConf) SetNewNodeCount(v int32) *ClickhouseConf {
	s.NewNodeCount = &v
	return s
}

func (s *ClickhouseConf) SetResizeType(v string) *ClickhouseConf {
	s.ResizeType = &v
	return s
}

func (s *ClickhouseConf) Validate() error {
	return dara.Validate(s)
}
