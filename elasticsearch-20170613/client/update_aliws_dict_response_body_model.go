// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAliwsDictResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateAliwsDictResponseBody
	GetRequestId() *string
	SetResult(v []*UpdateAliwsDictResponseBodyResult) *UpdateAliwsDictResponseBody
	GetResult() []*UpdateAliwsDictResponseBodyResult
}

type UpdateAliwsDictResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result.
	Result []*UpdateAliwsDictResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s UpdateAliwsDictResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAliwsDictResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAliwsDictResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAliwsDictResponseBody) GetResult() []*UpdateAliwsDictResponseBodyResult {
	return s.Result
}

func (s *UpdateAliwsDictResponseBody) SetRequestId(v string) *UpdateAliwsDictResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAliwsDictResponseBody) SetResult(v []*UpdateAliwsDictResponseBodyResult) *UpdateAliwsDictResponseBody {
	s.Result = v
	return s
}

func (s *UpdateAliwsDictResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateAliwsDictResponseBodyResult struct {
	// The size of the dictionary file. Unit: bytes.
	//
	// example:
	//
	// 6226
	FileSize *int64 `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	// The name of the uploaded dictionary file.
	//
	// example:
	//
	// aliws_ext_dict.txt
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
	// The dictionary type. The value is fixed as ALI_WS.
	//
	// example:
	//
	// ALI_WS
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateAliwsDictResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s UpdateAliwsDictResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateAliwsDictResponseBodyResult) GetFileSize() *int64 {
	return s.FileSize
}

func (s *UpdateAliwsDictResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *UpdateAliwsDictResponseBodyResult) GetSourceType() *string {
	return s.SourceType
}

func (s *UpdateAliwsDictResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *UpdateAliwsDictResponseBodyResult) SetFileSize(v int64) *UpdateAliwsDictResponseBodyResult {
	s.FileSize = &v
	return s
}

func (s *UpdateAliwsDictResponseBodyResult) SetName(v string) *UpdateAliwsDictResponseBodyResult {
	s.Name = &v
	return s
}

func (s *UpdateAliwsDictResponseBodyResult) SetSourceType(v string) *UpdateAliwsDictResponseBodyResult {
	s.SourceType = &v
	return s
}

func (s *UpdateAliwsDictResponseBodyResult) SetType(v string) *UpdateAliwsDictResponseBodyResult {
	s.Type = &v
	return s
}

func (s *UpdateAliwsDictResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
