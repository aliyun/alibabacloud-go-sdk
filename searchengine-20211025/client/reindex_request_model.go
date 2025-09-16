// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReindexRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataTimeSec(v int32) *ReindexRequest
	GetDataTimeSec() *int32
	SetOssDataPath(v string) *ReindexRequest
	GetOssDataPath() *string
	SetPartition(v string) *ReindexRequest
	GetPartition() *string
}

type ReindexRequest struct {
	// The timestamp in seconds. The value must be of the INTEGER type. This parameter is required if you specify an API data source.
	//
	// example:
	//
	// 1640867288
	DataTimeSec *int32 `json:"dataTimeSec,omitempty" xml:"dataTimeSec,omitempty"`
	// oss data path
	//
	// example:
	//
	// oss://opensearch
	OssDataPath *string `json:"ossDataPath,omitempty" xml:"ossDataPath,omitempty"`
	// The partition in the MaxCompute table. This parameter is required if type is set to odps.
	//
	// example:
	//
	// ds=20220713
	Partition *string `json:"partition,omitempty" xml:"partition,omitempty"`
}

func (s ReindexRequest) String() string {
	return dara.Prettify(s)
}

func (s ReindexRequest) GoString() string {
	return s.String()
}

func (s *ReindexRequest) GetDataTimeSec() *int32 {
	return s.DataTimeSec
}

func (s *ReindexRequest) GetOssDataPath() *string {
	return s.OssDataPath
}

func (s *ReindexRequest) GetPartition() *string {
	return s.Partition
}

func (s *ReindexRequest) SetDataTimeSec(v int32) *ReindexRequest {
	s.DataTimeSec = &v
	return s
}

func (s *ReindexRequest) SetOssDataPath(v string) *ReindexRequest {
	s.OssDataPath = &v
	return s
}

func (s *ReindexRequest) SetPartition(v string) *ReindexRequest {
	s.Partition = &v
	return s
}

func (s *ReindexRequest) Validate() error {
	return dara.Validate(s)
}
