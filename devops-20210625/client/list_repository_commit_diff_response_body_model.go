// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRepositoryCommitDiffResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListRepositoryCommitDiffResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListRepositoryCommitDiffResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListRepositoryCommitDiffResponseBody
	GetRequestId() *string
	SetResult(v []*ListRepositoryCommitDiffResponseBodyResult) *ListRepositoryCommitDiffResponseBody
	GetResult() []*ListRepositoryCommitDiffResponseBodyResult
	SetSuccess(v bool) *ListRepositoryCommitDiffResponseBody
	GetSuccess() *bool
}

type ListRepositoryCommitDiffResponseBody struct {
	// example:
	//
	// success
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 0EE38A4E-8991-532A-8E8B-5C22B5D2E058
	RequestId *string                                       `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*ListRepositoryCommitDiffResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListRepositoryCommitDiffResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRepositoryCommitDiffResponseBody) GoString() string {
	return s.String()
}

func (s *ListRepositoryCommitDiffResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListRepositoryCommitDiffResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListRepositoryCommitDiffResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRepositoryCommitDiffResponseBody) GetResult() []*ListRepositoryCommitDiffResponseBodyResult {
	return s.Result
}

func (s *ListRepositoryCommitDiffResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListRepositoryCommitDiffResponseBody) SetErrorCode(v string) *ListRepositoryCommitDiffResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListRepositoryCommitDiffResponseBody) SetErrorMessage(v string) *ListRepositoryCommitDiffResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListRepositoryCommitDiffResponseBody) SetRequestId(v string) *ListRepositoryCommitDiffResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRepositoryCommitDiffResponseBody) SetResult(v []*ListRepositoryCommitDiffResponseBodyResult) *ListRepositoryCommitDiffResponseBody {
	s.Result = v
	return s
}

func (s *ListRepositoryCommitDiffResponseBody) SetSuccess(v bool) *ListRepositoryCommitDiffResponseBody {
	s.Success = &v
	return s
}

func (s *ListRepositoryCommitDiffResponseBody) Validate() error {
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

type ListRepositoryCommitDiffResponseBodyResult struct {
	// example:
	//
	// 100644
	AMode *string `json:"aMode,omitempty" xml:"aMode,omitempty"`
	// example:
	//
	// 100755
	BMode       *string `json:"bMode,omitempty" xml:"bMode,omitempty"`
	DeletedFile *bool   `json:"deletedFile,omitempty" xml:"deletedFile,omitempty"`
	// example:
	//
	// --- /dev/null\\n+++ b/src/test/java/com/aliyun/codeupdemo/CodeupDemoApplicationTests.java\\n@@ -0,0 +1,13 @@\\n+package com.aliyun.codeupdemo;\\n+\\n+import org.junit.jupiter.api.Test;\\n+import org.springframework.boot.test.context.SpringBootTest;\\n+\\n+@SpringBootTest\\n+class CodeupDemoApplicationTest {\\n+\\n+ @Test\\n+ void contextLoads() {\\n+ }\\n+\\n+}\\n
	Diff     *string `json:"diff,omitempty" xml:"diff,omitempty"`
	IsBinary *bool   `json:"isBinary,omitempty" xml:"isBinary,omitempty"`
	IsNewLfs *bool   `json:"isNewLfs,omitempty" xml:"isNewLfs,omitempty"`
	IsOldLfs *bool   `json:"isOldLfs,omitempty" xml:"isOldLfs,omitempty"`
	NewFile  *bool   `json:"newFile,omitempty" xml:"newFile,omitempty"`
	// example:
	//
	// 6c268061a546378276559c713d0ad377d4xxxxxx
	NewId *string `json:"newId,omitempty" xml:"newId,omitempty"`
	// example:
	//
	// src/test/java/com/aliyun/codeupdemo/CodeupDemoApplicationTests.java
	NewPath *string `json:"newPath,omitempty" xml:"newPath,omitempty"`
	// example:
	//
	// 0000000000000000000000000000000000000000
	OldId *string `json:"oldId,omitempty" xml:"oldId,omitempty"`
	// example:
	//
	// src/test/java/com/aliyun/codeupdemo/CodeupDemoApplicationTests.java
	OldPath     *string `json:"oldPath,omitempty" xml:"oldPath,omitempty"`
	RenamedFile *bool   `json:"renamedFile,omitempty" xml:"renamedFile,omitempty"`
}

func (s ListRepositoryCommitDiffResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListRepositoryCommitDiffResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListRepositoryCommitDiffResponseBodyResult) GetAMode() *string {
	return s.AMode
}

func (s *ListRepositoryCommitDiffResponseBodyResult) GetBMode() *string {
	return s.BMode
}

func (s *ListRepositoryCommitDiffResponseBodyResult) GetDeletedFile() *bool {
	return s.DeletedFile
}

func (s *ListRepositoryCommitDiffResponseBodyResult) GetDiff() *string {
	return s.Diff
}

func (s *ListRepositoryCommitDiffResponseBodyResult) GetIsBinary() *bool {
	return s.IsBinary
}

func (s *ListRepositoryCommitDiffResponseBodyResult) GetIsNewLfs() *bool {
	return s.IsNewLfs
}

func (s *ListRepositoryCommitDiffResponseBodyResult) GetIsOldLfs() *bool {
	return s.IsOldLfs
}

func (s *ListRepositoryCommitDiffResponseBodyResult) GetNewFile() *bool {
	return s.NewFile
}

func (s *ListRepositoryCommitDiffResponseBodyResult) GetNewId() *string {
	return s.NewId
}

func (s *ListRepositoryCommitDiffResponseBodyResult) GetNewPath() *string {
	return s.NewPath
}

func (s *ListRepositoryCommitDiffResponseBodyResult) GetOldId() *string {
	return s.OldId
}

func (s *ListRepositoryCommitDiffResponseBodyResult) GetOldPath() *string {
	return s.OldPath
}

func (s *ListRepositoryCommitDiffResponseBodyResult) GetRenamedFile() *bool {
	return s.RenamedFile
}

func (s *ListRepositoryCommitDiffResponseBodyResult) SetAMode(v string) *ListRepositoryCommitDiffResponseBodyResult {
	s.AMode = &v
	return s
}

func (s *ListRepositoryCommitDiffResponseBodyResult) SetBMode(v string) *ListRepositoryCommitDiffResponseBodyResult {
	s.BMode = &v
	return s
}

func (s *ListRepositoryCommitDiffResponseBodyResult) SetDeletedFile(v bool) *ListRepositoryCommitDiffResponseBodyResult {
	s.DeletedFile = &v
	return s
}

func (s *ListRepositoryCommitDiffResponseBodyResult) SetDiff(v string) *ListRepositoryCommitDiffResponseBodyResult {
	s.Diff = &v
	return s
}

func (s *ListRepositoryCommitDiffResponseBodyResult) SetIsBinary(v bool) *ListRepositoryCommitDiffResponseBodyResult {
	s.IsBinary = &v
	return s
}

func (s *ListRepositoryCommitDiffResponseBodyResult) SetIsNewLfs(v bool) *ListRepositoryCommitDiffResponseBodyResult {
	s.IsNewLfs = &v
	return s
}

func (s *ListRepositoryCommitDiffResponseBodyResult) SetIsOldLfs(v bool) *ListRepositoryCommitDiffResponseBodyResult {
	s.IsOldLfs = &v
	return s
}

func (s *ListRepositoryCommitDiffResponseBodyResult) SetNewFile(v bool) *ListRepositoryCommitDiffResponseBodyResult {
	s.NewFile = &v
	return s
}

func (s *ListRepositoryCommitDiffResponseBodyResult) SetNewId(v string) *ListRepositoryCommitDiffResponseBodyResult {
	s.NewId = &v
	return s
}

func (s *ListRepositoryCommitDiffResponseBodyResult) SetNewPath(v string) *ListRepositoryCommitDiffResponseBodyResult {
	s.NewPath = &v
	return s
}

func (s *ListRepositoryCommitDiffResponseBodyResult) SetOldId(v string) *ListRepositoryCommitDiffResponseBodyResult {
	s.OldId = &v
	return s
}

func (s *ListRepositoryCommitDiffResponseBodyResult) SetOldPath(v string) *ListRepositoryCommitDiffResponseBodyResult {
	s.OldPath = &v
	return s
}

func (s *ListRepositoryCommitDiffResponseBodyResult) SetRenamedFile(v bool) *ListRepositoryCommitDiffResponseBodyResult {
	s.RenamedFile = &v
	return s
}

func (s *ListRepositoryCommitDiffResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
