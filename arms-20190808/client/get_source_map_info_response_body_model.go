// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSourceMapInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetSourceMapInfoResponseBody
	GetRequestId() *string
	SetSourceMapList(v []*GetSourceMapInfoResponseBodySourceMapList) *GetSourceMapInfoResponseBody
	GetSourceMapList() []*GetSourceMapInfoResponseBodySourceMapList
}

type GetSourceMapInfoResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// C21AB7CF-B7AF-410F-BD61-82D1567F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the SourceMap file.
	SourceMapList []*GetSourceMapInfoResponseBodySourceMapList `json:"SourceMapList,omitempty" xml:"SourceMapList,omitempty" type:"Repeated"`
}

func (s GetSourceMapInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSourceMapInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetSourceMapInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSourceMapInfoResponseBody) GetSourceMapList() []*GetSourceMapInfoResponseBodySourceMapList {
	return s.SourceMapList
}

func (s *GetSourceMapInfoResponseBody) SetRequestId(v string) *GetSourceMapInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSourceMapInfoResponseBody) SetSourceMapList(v []*GetSourceMapInfoResponseBodySourceMapList) *GetSourceMapInfoResponseBody {
	s.SourceMapList = v
	return s
}

func (s *GetSourceMapInfoResponseBody) Validate() error {
	if s.SourceMapList != nil {
		for _, item := range s.SourceMapList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetSourceMapInfoResponseBodySourceMapList struct {
	// The ID of the SourceMap file.
	Fid *string `json:"Fid,omitempty" xml:"Fid,omitempty"`
	// The name of the SourceMap file.
	//
	// example:
	//
	// test.sourcemap.js
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The size of the file. Unit: KB.
	//
	// example:
	//
	// 201223
	Size *string `json:"Size,omitempty" xml:"Size,omitempty"`
	// The timestamp that indicates when the file was uploaded.
	//
	// example:
	//
	// 1590657842000
	UploadTime *string `json:"UploadTime,omitempty" xml:"UploadTime,omitempty"`
	// The version of the file.
	//
	// example:
	//
	// 0.0.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetSourceMapInfoResponseBodySourceMapList) String() string {
	return dara.Prettify(s)
}

func (s GetSourceMapInfoResponseBodySourceMapList) GoString() string {
	return s.String()
}

func (s *GetSourceMapInfoResponseBodySourceMapList) GetFid() *string {
	return s.Fid
}

func (s *GetSourceMapInfoResponseBodySourceMapList) GetFileName() *string {
	return s.FileName
}

func (s *GetSourceMapInfoResponseBodySourceMapList) GetSize() *string {
	return s.Size
}

func (s *GetSourceMapInfoResponseBodySourceMapList) GetUploadTime() *string {
	return s.UploadTime
}

func (s *GetSourceMapInfoResponseBodySourceMapList) GetVersion() *string {
	return s.Version
}

func (s *GetSourceMapInfoResponseBodySourceMapList) SetFid(v string) *GetSourceMapInfoResponseBodySourceMapList {
	s.Fid = &v
	return s
}

func (s *GetSourceMapInfoResponseBodySourceMapList) SetFileName(v string) *GetSourceMapInfoResponseBodySourceMapList {
	s.FileName = &v
	return s
}

func (s *GetSourceMapInfoResponseBodySourceMapList) SetSize(v string) *GetSourceMapInfoResponseBodySourceMapList {
	s.Size = &v
	return s
}

func (s *GetSourceMapInfoResponseBodySourceMapList) SetUploadTime(v string) *GetSourceMapInfoResponseBodySourceMapList {
	s.UploadTime = &v
	return s
}

func (s *GetSourceMapInfoResponseBodySourceMapList) SetVersion(v string) *GetSourceMapInfoResponseBodySourceMapList {
	s.Version = &v
	return s
}

func (s *GetSourceMapInfoResponseBodySourceMapList) Validate() error {
	return dara.Validate(s)
}
