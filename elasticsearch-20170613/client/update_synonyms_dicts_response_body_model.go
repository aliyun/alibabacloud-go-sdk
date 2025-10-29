// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSynonymsDictsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateSynonymsDictsResponseBody
	GetRequestId() *string
	SetResult(v []*UpdateSynonymsDictsResponseBodyResult) *UpdateSynonymsDictsResponseBody
	GetResult() []*UpdateSynonymsDictsResponseBodyResult
}

type UpdateSynonymsDictsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 7C5622CC-B312-426F-85AA-B0271*******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result.
	Result []*UpdateSynonymsDictsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s UpdateSynonymsDictsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSynonymsDictsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSynonymsDictsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSynonymsDictsResponseBody) GetResult() []*UpdateSynonymsDictsResponseBodyResult {
	return s.Result
}

func (s *UpdateSynonymsDictsResponseBody) SetRequestId(v string) *UpdateSynonymsDictsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSynonymsDictsResponseBody) SetResult(v []*UpdateSynonymsDictsResponseBodyResult) *UpdateSynonymsDictsResponseBody {
	s.Result = v
	return s
}

func (s *UpdateSynonymsDictsResponseBody) Validate() error {
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

type UpdateSynonymsDictsResponseBodyResult struct {
	// The size of the dictionary file. Unit: bytes.
	//
	// example:
	//
	// 220
	FileSize *int64 `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	// The name of the dictionary file.
	//
	// example:
	//
	// deploy_0.txt
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
	// The dictionary type. The value is fixed as SYNONYMS.
	//
	// example:
	//
	// SYNONYMS
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateSynonymsDictsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s UpdateSynonymsDictsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateSynonymsDictsResponseBodyResult) GetFileSize() *int64 {
	return s.FileSize
}

func (s *UpdateSynonymsDictsResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *UpdateSynonymsDictsResponseBodyResult) GetSourceType() *string {
	return s.SourceType
}

func (s *UpdateSynonymsDictsResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *UpdateSynonymsDictsResponseBodyResult) SetFileSize(v int64) *UpdateSynonymsDictsResponseBodyResult {
	s.FileSize = &v
	return s
}

func (s *UpdateSynonymsDictsResponseBodyResult) SetName(v string) *UpdateSynonymsDictsResponseBodyResult {
	s.Name = &v
	return s
}

func (s *UpdateSynonymsDictsResponseBodyResult) SetSourceType(v string) *UpdateSynonymsDictsResponseBodyResult {
	s.SourceType = &v
	return s
}

func (s *UpdateSynonymsDictsResponseBodyResult) SetType(v string) *UpdateSynonymsDictsResponseBodyResult {
	s.Type = &v
	return s
}

func (s *UpdateSynonymsDictsResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
