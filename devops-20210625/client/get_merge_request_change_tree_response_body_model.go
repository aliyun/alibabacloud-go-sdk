// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMergeRequestChangeTreeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetMergeRequestChangeTreeResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetMergeRequestChangeTreeResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetMergeRequestChangeTreeResponseBody
	GetRequestId() *string
	SetResult(v *GetMergeRequestChangeTreeResponseBodyResult) *GetMergeRequestChangeTreeResponseBody
	GetResult() *GetMergeRequestChangeTreeResponseBodyResult
	SetSuccess(v bool) *GetMergeRequestChangeTreeResponseBody
	GetSuccess() *bool
}

type GetMergeRequestChangeTreeResponseBody struct {
	// example:
	//
	// SYSTEM_UNKNOWN_ERROR
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// HC93CE1A-8D7A-13A9-8306-7465DE2E5C0F
	RequestId *string                                      `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *GetMergeRequestChangeTreeResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetMergeRequestChangeTreeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMergeRequestChangeTreeResponseBody) GoString() string {
	return s.String()
}

func (s *GetMergeRequestChangeTreeResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetMergeRequestChangeTreeResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetMergeRequestChangeTreeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMergeRequestChangeTreeResponseBody) GetResult() *GetMergeRequestChangeTreeResponseBodyResult {
	return s.Result
}

func (s *GetMergeRequestChangeTreeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetMergeRequestChangeTreeResponseBody) SetErrorCode(v string) *GetMergeRequestChangeTreeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetMergeRequestChangeTreeResponseBody) SetErrorMessage(v string) *GetMergeRequestChangeTreeResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetMergeRequestChangeTreeResponseBody) SetRequestId(v string) *GetMergeRequestChangeTreeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMergeRequestChangeTreeResponseBody) SetResult(v *GetMergeRequestChangeTreeResponseBodyResult) *GetMergeRequestChangeTreeResponseBody {
	s.Result = v
	return s
}

func (s *GetMergeRequestChangeTreeResponseBody) SetSuccess(v bool) *GetMergeRequestChangeTreeResponseBody {
	s.Success = &v
	return s
}

func (s *GetMergeRequestChangeTreeResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMergeRequestChangeTreeResponseBodyResult struct {
	// example:
	//
	// 20
	ChangedFilesCount *int64                                                          `json:"changedFilesCount,omitempty" xml:"changedFilesCount,omitempty"`
	ChangedFilesInfos []*GetMergeRequestChangeTreeResponseBodyResultChangedFilesInfos `json:"changedFilesInfos,omitempty" xml:"changedFilesInfos,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	TotalAddLines *int64 `json:"totalAddLines,omitempty" xml:"totalAddLines,omitempty"`
	// example:
	//
	// 50
	TotalDelLines *int64 `json:"totalDelLines,omitempty" xml:"totalDelLines,omitempty"`
}

func (s GetMergeRequestChangeTreeResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetMergeRequestChangeTreeResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetMergeRequestChangeTreeResponseBodyResult) GetChangedFilesCount() *int64 {
	return s.ChangedFilesCount
}

func (s *GetMergeRequestChangeTreeResponseBodyResult) GetChangedFilesInfos() []*GetMergeRequestChangeTreeResponseBodyResultChangedFilesInfos {
	return s.ChangedFilesInfos
}

func (s *GetMergeRequestChangeTreeResponseBodyResult) GetTotalAddLines() *int64 {
	return s.TotalAddLines
}

func (s *GetMergeRequestChangeTreeResponseBodyResult) GetTotalDelLines() *int64 {
	return s.TotalDelLines
}

func (s *GetMergeRequestChangeTreeResponseBodyResult) SetChangedFilesCount(v int64) *GetMergeRequestChangeTreeResponseBodyResult {
	s.ChangedFilesCount = &v
	return s
}

func (s *GetMergeRequestChangeTreeResponseBodyResult) SetChangedFilesInfos(v []*GetMergeRequestChangeTreeResponseBodyResultChangedFilesInfos) *GetMergeRequestChangeTreeResponseBodyResult {
	s.ChangedFilesInfos = v
	return s
}

func (s *GetMergeRequestChangeTreeResponseBodyResult) SetTotalAddLines(v int64) *GetMergeRequestChangeTreeResponseBodyResult {
	s.TotalAddLines = &v
	return s
}

func (s *GetMergeRequestChangeTreeResponseBodyResult) SetTotalDelLines(v int64) *GetMergeRequestChangeTreeResponseBodyResult {
	s.TotalDelLines = &v
	return s
}

func (s *GetMergeRequestChangeTreeResponseBodyResult) Validate() error {
	if s.ChangedFilesInfos != nil {
		for _, item := range s.ChangedFilesInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetMergeRequestChangeTreeResponseBodyResultChangedFilesInfos struct {
	// example:
	//
	// 10
	AddLines *int64 `json:"addLines,omitempty" xml:"addLines,omitempty"`
	// example:
	//
	// false
	BinaryFile *bool `json:"binaryFile,omitempty" xml:"binaryFile,omitempty"`
	// example:
	//
	// 0
	DelLines *int64 `json:"delLines,omitempty" xml:"delLines,omitempty"`
	// example:
	//
	// false
	DeletedFile *bool `json:"deletedFile,omitempty" xml:"deletedFile,omitempty"`
	// example:
	//
	// true
	NewFile *bool `json:"newFile,omitempty" xml:"newFile,omitempty"`
	// example:
	//
	// test.txt
	NewPath *string `json:"newPath,omitempty" xml:"newPath,omitempty"`
	// example:
	//
	// test.txt
	OldPath *string `json:"oldPath,omitempty" xml:"oldPath,omitempty"`
	// example:
	//
	// false
	RenamedFile *bool `json:"renamedFile,omitempty" xml:"renamedFile,omitempty"`
}

func (s GetMergeRequestChangeTreeResponseBodyResultChangedFilesInfos) String() string {
	return dara.Prettify(s)
}

func (s GetMergeRequestChangeTreeResponseBodyResultChangedFilesInfos) GoString() string {
	return s.String()
}

func (s *GetMergeRequestChangeTreeResponseBodyResultChangedFilesInfos) GetAddLines() *int64 {
	return s.AddLines
}

func (s *GetMergeRequestChangeTreeResponseBodyResultChangedFilesInfos) GetBinaryFile() *bool {
	return s.BinaryFile
}

func (s *GetMergeRequestChangeTreeResponseBodyResultChangedFilesInfos) GetDelLines() *int64 {
	return s.DelLines
}

func (s *GetMergeRequestChangeTreeResponseBodyResultChangedFilesInfos) GetDeletedFile() *bool {
	return s.DeletedFile
}

func (s *GetMergeRequestChangeTreeResponseBodyResultChangedFilesInfos) GetNewFile() *bool {
	return s.NewFile
}

func (s *GetMergeRequestChangeTreeResponseBodyResultChangedFilesInfos) GetNewPath() *string {
	return s.NewPath
}

func (s *GetMergeRequestChangeTreeResponseBodyResultChangedFilesInfos) GetOldPath() *string {
	return s.OldPath
}

func (s *GetMergeRequestChangeTreeResponseBodyResultChangedFilesInfos) GetRenamedFile() *bool {
	return s.RenamedFile
}

func (s *GetMergeRequestChangeTreeResponseBodyResultChangedFilesInfos) SetAddLines(v int64) *GetMergeRequestChangeTreeResponseBodyResultChangedFilesInfos {
	s.AddLines = &v
	return s
}

func (s *GetMergeRequestChangeTreeResponseBodyResultChangedFilesInfos) SetBinaryFile(v bool) *GetMergeRequestChangeTreeResponseBodyResultChangedFilesInfos {
	s.BinaryFile = &v
	return s
}

func (s *GetMergeRequestChangeTreeResponseBodyResultChangedFilesInfos) SetDelLines(v int64) *GetMergeRequestChangeTreeResponseBodyResultChangedFilesInfos {
	s.DelLines = &v
	return s
}

func (s *GetMergeRequestChangeTreeResponseBodyResultChangedFilesInfos) SetDeletedFile(v bool) *GetMergeRequestChangeTreeResponseBodyResultChangedFilesInfos {
	s.DeletedFile = &v
	return s
}

func (s *GetMergeRequestChangeTreeResponseBodyResultChangedFilesInfos) SetNewFile(v bool) *GetMergeRequestChangeTreeResponseBodyResultChangedFilesInfos {
	s.NewFile = &v
	return s
}

func (s *GetMergeRequestChangeTreeResponseBodyResultChangedFilesInfos) SetNewPath(v string) *GetMergeRequestChangeTreeResponseBodyResultChangedFilesInfos {
	s.NewPath = &v
	return s
}

func (s *GetMergeRequestChangeTreeResponseBodyResultChangedFilesInfos) SetOldPath(v string) *GetMergeRequestChangeTreeResponseBodyResultChangedFilesInfos {
	s.OldPath = &v
	return s
}

func (s *GetMergeRequestChangeTreeResponseBodyResultChangedFilesInfos) SetRenamedFile(v bool) *GetMergeRequestChangeTreeResponseBodyResultChangedFilesInfos {
	s.RenamedFile = &v
	return s
}

func (s *GetMergeRequestChangeTreeResponseBodyResultChangedFilesInfos) Validate() error {
	return dara.Validate(s)
}
