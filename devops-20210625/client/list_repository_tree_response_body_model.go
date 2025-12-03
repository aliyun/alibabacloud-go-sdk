// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRepositoryTreeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListRepositoryTreeResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListRepositoryTreeResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListRepositoryTreeResponseBody
	GetRequestId() *string
	SetResult(v []*ListRepositoryTreeResponseBodyResult) *ListRepositoryTreeResponseBody
	GetResult() []*ListRepositoryTreeResponseBodyResult
	SetSuccess(v bool) *ListRepositoryTreeResponseBody
	GetSuccess() *bool
}

type ListRepositoryTreeResponseBody struct {
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
	// 6557983C-FB08-51A9-AC5A-A7A0D0950A07
	RequestId *string                                 `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*ListRepositoryTreeResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListRepositoryTreeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRepositoryTreeResponseBody) GoString() string {
	return s.String()
}

func (s *ListRepositoryTreeResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListRepositoryTreeResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListRepositoryTreeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRepositoryTreeResponseBody) GetResult() []*ListRepositoryTreeResponseBodyResult {
	return s.Result
}

func (s *ListRepositoryTreeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListRepositoryTreeResponseBody) SetErrorCode(v string) *ListRepositoryTreeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListRepositoryTreeResponseBody) SetErrorMessage(v string) *ListRepositoryTreeResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListRepositoryTreeResponseBody) SetRequestId(v string) *ListRepositoryTreeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRepositoryTreeResponseBody) SetResult(v []*ListRepositoryTreeResponseBodyResult) *ListRepositoryTreeResponseBody {
	s.Result = v
	return s
}

func (s *ListRepositoryTreeResponseBody) SetSuccess(v bool) *ListRepositoryTreeResponseBody {
	s.Success = &v
	return s
}

func (s *ListRepositoryTreeResponseBody) Validate() error {
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

type ListRepositoryTreeResponseBodyResult struct {
	// example:
	//
	// 76c3f251f414ac31f2e01faf6f2008a9d756a437
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// false
	IsLFS *bool `json:"isLFS,omitempty" xml:"isLFS,omitempty"`
	// example:
	//
	// 100644
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
	// example:
	//
	// test-codeup
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// test-codeup
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// example:
	//
	// blob
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListRepositoryTreeResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListRepositoryTreeResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListRepositoryTreeResponseBodyResult) GetId() *string {
	return s.Id
}

func (s *ListRepositoryTreeResponseBodyResult) GetIsLFS() *bool {
	return s.IsLFS
}

func (s *ListRepositoryTreeResponseBodyResult) GetMode() *string {
	return s.Mode
}

func (s *ListRepositoryTreeResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *ListRepositoryTreeResponseBodyResult) GetPath() *string {
	return s.Path
}

func (s *ListRepositoryTreeResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *ListRepositoryTreeResponseBodyResult) SetId(v string) *ListRepositoryTreeResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListRepositoryTreeResponseBodyResult) SetIsLFS(v bool) *ListRepositoryTreeResponseBodyResult {
	s.IsLFS = &v
	return s
}

func (s *ListRepositoryTreeResponseBodyResult) SetMode(v string) *ListRepositoryTreeResponseBodyResult {
	s.Mode = &v
	return s
}

func (s *ListRepositoryTreeResponseBodyResult) SetName(v string) *ListRepositoryTreeResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListRepositoryTreeResponseBodyResult) SetPath(v string) *ListRepositoryTreeResponseBodyResult {
	s.Path = &v
	return s
}

func (s *ListRepositoryTreeResponseBodyResult) SetType(v string) *ListRepositoryTreeResponseBodyResult {
	s.Type = &v
	return s
}

func (s *ListRepositoryTreeResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
