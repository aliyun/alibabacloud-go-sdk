// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHotIkDictsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateHotIkDictsResponseBody
	GetRequestId() *string
	SetResult(v []*UpdateHotIkDictsResponseBodyResult) *UpdateHotIkDictsResponseBody
	GetResult() []*UpdateHotIkDictsResponseBodyResult
}

type UpdateHotIkDictsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result.
	Result []*UpdateHotIkDictsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s UpdateHotIkDictsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateHotIkDictsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateHotIkDictsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateHotIkDictsResponseBody) GetResult() []*UpdateHotIkDictsResponseBodyResult {
	return s.Result
}

func (s *UpdateHotIkDictsResponseBody) SetRequestId(v string) *UpdateHotIkDictsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateHotIkDictsResponseBody) SetResult(v []*UpdateHotIkDictsResponseBodyResult) *UpdateHotIkDictsResponseBody {
	s.Result = v
	return s
}

func (s *UpdateHotIkDictsResponseBody) Validate() error {
	return dara.Validate(s)
}

type UpdateHotIkDictsResponseBodyResult struct {
	// The size of the dictionary file. Unit: bytes.
	//
	// example:
	//
	// 6
	FileSize *int64 `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	// The name of the dictionary file.
	//
	// example:
	//
	// deploy_0.dic
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The source type of the dictionary file. Valid values:
	//
	// 	- OSS
	//
	// 	- ORIGIN
	//
	// example:
	//
	// OSS
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// The type of the dictionaries. Valid values:
	//
	// 	- MAIN: IK main dictionary
	//
	// 	- STOP: IK stopword list
	//
	// example:
	//
	// MAIN
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateHotIkDictsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s UpdateHotIkDictsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateHotIkDictsResponseBodyResult) GetFileSize() *int64 {
	return s.FileSize
}

func (s *UpdateHotIkDictsResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *UpdateHotIkDictsResponseBodyResult) GetSourceType() *string {
	return s.SourceType
}

func (s *UpdateHotIkDictsResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *UpdateHotIkDictsResponseBodyResult) SetFileSize(v int64) *UpdateHotIkDictsResponseBodyResult {
	s.FileSize = &v
	return s
}

func (s *UpdateHotIkDictsResponseBodyResult) SetName(v string) *UpdateHotIkDictsResponseBodyResult {
	s.Name = &v
	return s
}

func (s *UpdateHotIkDictsResponseBodyResult) SetSourceType(v string) *UpdateHotIkDictsResponseBodyResult {
	s.SourceType = &v
	return s
}

func (s *UpdateHotIkDictsResponseBodyResult) SetType(v string) *UpdateHotIkDictsResponseBodyResult {
	s.Type = &v
	return s
}

func (s *UpdateHotIkDictsResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
