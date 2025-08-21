// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDictResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateDictResponseBody
	GetRequestId() *string
	SetResult(v []*UpdateDictResponseBodyResult) *UpdateDictResponseBody
	GetResult() []*UpdateDictResponseBodyResult
}

type UpdateDictResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result.
	Result []*UpdateDictResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s UpdateDictResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDictResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDictResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDictResponseBody) GetResult() []*UpdateDictResponseBodyResult {
	return s.Result
}

func (s *UpdateDictResponseBody) SetRequestId(v string) *UpdateDictResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDictResponseBody) SetResult(v []*UpdateDictResponseBodyResult) *UpdateDictResponseBody {
	s.Result = v
	return s
}

func (s *UpdateDictResponseBody) Validate() error {
	return dara.Validate(s)
}

type UpdateDictResponseBodyResult struct {
	// The size of the dictionary file. Unit: bytes.
	//
	// example:
	//
	// 2782602
	FileSize *int64 `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	// The name of the dictionary file.
	//
	// example:
	//
	// SYSTEM_MAIN.dic
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The source type of the dictionary file. Valid values:
	//
	// 	- OSS
	//
	// 	- ORIGIN
	//
	// example:
	//
	// ORIGIN
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// The dictionary type. Valid values:
	//
	// 	- MAIN: IK main dicrionary
	//
	// 	- STOP: IK stopword list
	//
	// example:
	//
	// MAIN
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateDictResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s UpdateDictResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateDictResponseBodyResult) GetFileSize() *int64 {
	return s.FileSize
}

func (s *UpdateDictResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *UpdateDictResponseBodyResult) GetSourceType() *string {
	return s.SourceType
}

func (s *UpdateDictResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *UpdateDictResponseBodyResult) SetFileSize(v int64) *UpdateDictResponseBodyResult {
	s.FileSize = &v
	return s
}

func (s *UpdateDictResponseBodyResult) SetName(v string) *UpdateDictResponseBodyResult {
	s.Name = &v
	return s
}

func (s *UpdateDictResponseBodyResult) SetSourceType(v string) *UpdateDictResponseBodyResult {
	s.SourceType = &v
	return s
}

func (s *UpdateDictResponseBodyResult) SetType(v string) *UpdateDictResponseBodyResult {
	s.Type = &v
	return s
}

func (s *UpdateDictResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
