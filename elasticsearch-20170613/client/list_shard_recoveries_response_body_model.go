// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListShardRecoveriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListShardRecoveriesResponseBody
	GetRequestId() *string
	SetResult(v []*ListShardRecoveriesResponseBodyResult) *ListShardRecoveriesResponseBody
	GetResult() []*ListShardRecoveriesResponseBodyResult
}

type ListShardRecoveriesResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// F99407AB-2FA9-489E-A259-40CF6DCC47D9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result.
	Result []*ListShardRecoveriesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListShardRecoveriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListShardRecoveriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListShardRecoveriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListShardRecoveriesResponseBody) GetResult() []*ListShardRecoveriesResponseBodyResult {
	return s.Result
}

func (s *ListShardRecoveriesResponseBody) SetRequestId(v string) *ListShardRecoveriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListShardRecoveriesResponseBody) SetResult(v []*ListShardRecoveriesResponseBodyResult) *ListShardRecoveriesResponseBody {
	s.Result = v
	return s
}

func (s *ListShardRecoveriesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListShardRecoveriesResponseBodyResult struct {
	// The data restoration progress.
	//
	// example:
	//
	// 80%
	BytesPercent *string `json:"bytesPercent,omitempty" xml:"bytesPercent,omitempty"`
	// The total amount of data that is restored.
	//
	// example:
	//
	// 12086
	BytesTotal *int64 `json:"bytesTotal,omitempty" xml:"bytesTotal,omitempty"`
	// The file execution progress.
	//
	// example:
	//
	// 80.0%
	FilesPercent *string `json:"filesPercent,omitempty" xml:"filesPercent,omitempty"`
	// The total number of files.
	//
	// example:
	//
	// 79
	FilesTotal *int64 `json:"filesTotal,omitempty" xml:"filesTotal,omitempty"`
	// The name of the index.
	//
	// example:
	//
	// my-index-000001
	Index *string `json:"index,omitempty" xml:"index,omitempty"`
	// The IP address of the source node.
	//
	// example:
	//
	// 192.168.XX.XX
	SourceHost *string `json:"sourceHost,omitempty" xml:"sourceHost,omitempty"`
	// The name of the source node.
	//
	// example:
	//
	// 2Kni3dJ
	SourceNode *string `json:"sourceNode,omitempty" xml:"sourceNode,omitempty"`
	// The data restoration status. Valid values:
	//
	// 	- done: Data restoration is complete.
	//
	// 	- finalize: Data is being cleared.
	//
	// 	- index: Index metadata is being read, and bytes are being copied from source to destination.
	//
	// 	- init: Data restoration is not started.
	//
	// 	- start: Data restoration is started.
	//
	// 	- translog: Translogs are being redone.
	//
	// example:
	//
	// done
	Stage *string `json:"stage,omitempty" xml:"stage,omitempty"`
	// The IP address of the destination node.
	//
	// example:
	//
	// 192.168.XX.XX
	TargetHost *string `json:"targetHost,omitempty" xml:"targetHost,omitempty"`
	// The name of the destination node.
	//
	// example:
	//
	// YVVKLmW
	TargetNode *string `json:"targetNode,omitempty" xml:"targetNode,omitempty"`
	// The number of translog operations to be restored.
	//
	// example:
	//
	// 12086
	TranslogOps *int64 `json:"translogOps,omitempty" xml:"translogOps,omitempty"`
	// The restoration progress of translog operations.
	//
	// example:
	//
	// 80%
	TranslogOpsPercent *string `json:"translogOpsPercent,omitempty" xml:"translogOpsPercent,omitempty"`
}

func (s ListShardRecoveriesResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListShardRecoveriesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListShardRecoveriesResponseBodyResult) GetBytesPercent() *string {
	return s.BytesPercent
}

func (s *ListShardRecoveriesResponseBodyResult) GetBytesTotal() *int64 {
	return s.BytesTotal
}

func (s *ListShardRecoveriesResponseBodyResult) GetFilesPercent() *string {
	return s.FilesPercent
}

func (s *ListShardRecoveriesResponseBodyResult) GetFilesTotal() *int64 {
	return s.FilesTotal
}

func (s *ListShardRecoveriesResponseBodyResult) GetIndex() *string {
	return s.Index
}

func (s *ListShardRecoveriesResponseBodyResult) GetSourceHost() *string {
	return s.SourceHost
}

func (s *ListShardRecoveriesResponseBodyResult) GetSourceNode() *string {
	return s.SourceNode
}

func (s *ListShardRecoveriesResponseBodyResult) GetStage() *string {
	return s.Stage
}

func (s *ListShardRecoveriesResponseBodyResult) GetTargetHost() *string {
	return s.TargetHost
}

func (s *ListShardRecoveriesResponseBodyResult) GetTargetNode() *string {
	return s.TargetNode
}

func (s *ListShardRecoveriesResponseBodyResult) GetTranslogOps() *int64 {
	return s.TranslogOps
}

func (s *ListShardRecoveriesResponseBodyResult) GetTranslogOpsPercent() *string {
	return s.TranslogOpsPercent
}

func (s *ListShardRecoveriesResponseBodyResult) SetBytesPercent(v string) *ListShardRecoveriesResponseBodyResult {
	s.BytesPercent = &v
	return s
}

func (s *ListShardRecoveriesResponseBodyResult) SetBytesTotal(v int64) *ListShardRecoveriesResponseBodyResult {
	s.BytesTotal = &v
	return s
}

func (s *ListShardRecoveriesResponseBodyResult) SetFilesPercent(v string) *ListShardRecoveriesResponseBodyResult {
	s.FilesPercent = &v
	return s
}

func (s *ListShardRecoveriesResponseBodyResult) SetFilesTotal(v int64) *ListShardRecoveriesResponseBodyResult {
	s.FilesTotal = &v
	return s
}

func (s *ListShardRecoveriesResponseBodyResult) SetIndex(v string) *ListShardRecoveriesResponseBodyResult {
	s.Index = &v
	return s
}

func (s *ListShardRecoveriesResponseBodyResult) SetSourceHost(v string) *ListShardRecoveriesResponseBodyResult {
	s.SourceHost = &v
	return s
}

func (s *ListShardRecoveriesResponseBodyResult) SetSourceNode(v string) *ListShardRecoveriesResponseBodyResult {
	s.SourceNode = &v
	return s
}

func (s *ListShardRecoveriesResponseBodyResult) SetStage(v string) *ListShardRecoveriesResponseBodyResult {
	s.Stage = &v
	return s
}

func (s *ListShardRecoveriesResponseBodyResult) SetTargetHost(v string) *ListShardRecoveriesResponseBodyResult {
	s.TargetHost = &v
	return s
}

func (s *ListShardRecoveriesResponseBodyResult) SetTargetNode(v string) *ListShardRecoveriesResponseBodyResult {
	s.TargetNode = &v
	return s
}

func (s *ListShardRecoveriesResponseBodyResult) SetTranslogOps(v int64) *ListShardRecoveriesResponseBodyResult {
	s.TranslogOps = &v
	return s
}

func (s *ListShardRecoveriesResponseBodyResult) SetTranslogOpsPercent(v string) *ListShardRecoveriesResponseBodyResult {
	s.TranslogOpsPercent = &v
	return s
}

func (s *ListShardRecoveriesResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
