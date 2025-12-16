// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSortScriptFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetSortScriptFileResponseBody
	GetRequestId() *string
	SetResult(v *GetSortScriptFileResponseBodyResult) *GetSortScriptFileResponseBody
	GetResult() *GetSortScriptFileResponseBodyResult
}

type GetSortScriptFileResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ABCDEFGH
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The content of the sort script.
	Result *GetSortScriptFileResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetSortScriptFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSortScriptFileResponseBody) GoString() string {
	return s.String()
}

func (s *GetSortScriptFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSortScriptFileResponseBody) GetResult() *GetSortScriptFileResponseBodyResult {
	return s.Result
}

func (s *GetSortScriptFileResponseBody) SetRequestId(v string) *GetSortScriptFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSortScriptFileResponseBody) SetResult(v *GetSortScriptFileResponseBodyResult) *GetSortScriptFileResponseBody {
	s.Result = v
	return s
}

func (s *GetSortScriptFileResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSortScriptFileResponseBodyResult struct {
	// The script content that is encoded in the Base64 format.
	//
	// example:
	//
	// YWJjZGVmZw==
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The time when the script was created.
	//
	// example:
	//
	// 2020-04-02 20:21:14
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The time when the script was last modified.
	//
	// example:
	//
	// 2020-04-02 21:21:14
	ModifyTime *string `json:"modifyTime,omitempty" xml:"modifyTime,omitempty"`
	// The version of the script content.
	//
	// example:
	//
	// 123456
	Version *int64 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetSortScriptFileResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetSortScriptFileResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetSortScriptFileResponseBodyResult) GetContent() *string {
	return s.Content
}

func (s *GetSortScriptFileResponseBodyResult) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetSortScriptFileResponseBodyResult) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *GetSortScriptFileResponseBodyResult) GetVersion() *int64 {
	return s.Version
}

func (s *GetSortScriptFileResponseBodyResult) SetContent(v string) *GetSortScriptFileResponseBodyResult {
	s.Content = &v
	return s
}

func (s *GetSortScriptFileResponseBodyResult) SetCreateTime(v string) *GetSortScriptFileResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *GetSortScriptFileResponseBodyResult) SetModifyTime(v string) *GetSortScriptFileResponseBodyResult {
	s.ModifyTime = &v
	return s
}

func (s *GetSortScriptFileResponseBodyResult) SetVersion(v int64) *GetSortScriptFileResponseBodyResult {
	s.Version = &v
	return s
}

func (s *GetSortScriptFileResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
