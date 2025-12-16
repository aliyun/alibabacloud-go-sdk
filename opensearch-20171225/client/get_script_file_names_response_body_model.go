// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScriptFileNamesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetScriptFileNamesResponseBody
	GetRequestId() *string
	SetResult(v []*GetScriptFileNamesResponseBodyResult) *GetScriptFileNamesResponseBody
	GetResult() []*GetScriptFileNamesResponseBodyResult
}

type GetScriptFileNamesResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// ABCDEFGH
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The files of the script.
	Result []*GetScriptFileNamesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s GetScriptFileNamesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetScriptFileNamesResponseBody) GoString() string {
	return s.String()
}

func (s *GetScriptFileNamesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetScriptFileNamesResponseBody) GetResult() []*GetScriptFileNamesResponseBodyResult {
	return s.Result
}

func (s *GetScriptFileNamesResponseBody) SetRequestId(v string) *GetScriptFileNamesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetScriptFileNamesResponseBody) SetResult(v []*GetScriptFileNamesResponseBodyResult) *GetScriptFileNamesResponseBody {
	s.Result = v
	return s
}

func (s *GetScriptFileNamesResponseBody) Validate() error {
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

type GetScriptFileNamesResponseBodyResult struct {
	// The time when the script file was created.
	//
	// example:
	//
	// 2020-04-02 20:21:14
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The name of the script file.
	//
	// example:
	//
	// my_cava_script.cava
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// The time when the script file was last modified.
	//
	// example:
	//
	// 2020-04-02 21:21:14
	ModifyTime *string `json:"modifyTime,omitempty" xml:"modifyTime,omitempty"`
	// The path name of the script file.
	//
	// example:
	//
	// src
	PathName *string `json:"pathName,omitempty" xml:"pathName,omitempty"`
}

func (s GetScriptFileNamesResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetScriptFileNamesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetScriptFileNamesResponseBodyResult) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetScriptFileNamesResponseBodyResult) GetFileName() *string {
	return s.FileName
}

func (s *GetScriptFileNamesResponseBodyResult) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *GetScriptFileNamesResponseBodyResult) GetPathName() *string {
	return s.PathName
}

func (s *GetScriptFileNamesResponseBodyResult) SetCreateTime(v string) *GetScriptFileNamesResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *GetScriptFileNamesResponseBodyResult) SetFileName(v string) *GetScriptFileNamesResponseBodyResult {
	s.FileName = &v
	return s
}

func (s *GetScriptFileNamesResponseBodyResult) SetModifyTime(v string) *GetScriptFileNamesResponseBodyResult {
	s.ModifyTime = &v
	return s
}

func (s *GetScriptFileNamesResponseBodyResult) SetPathName(v string) *GetScriptFileNamesResponseBodyResult {
	s.PathName = &v
	return s
}

func (s *GetScriptFileNamesResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
