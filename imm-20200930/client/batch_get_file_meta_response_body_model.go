// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGetFileMetaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFiles(v []*File) *BatchGetFileMetaResponseBody
	GetFiles() []*File
	SetRequestId(v string) *BatchGetFileMetaResponseBody
	GetRequestId() *string
}

type BatchGetFileMetaResponseBody struct {
	// The metadata returned.
	Files []*File `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 7F84C6D9-5AC0-49F9-914D-F02678E3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchGetFileMetaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchGetFileMetaResponseBody) GoString() string {
	return s.String()
}

func (s *BatchGetFileMetaResponseBody) GetFiles() []*File {
	return s.Files
}

func (s *BatchGetFileMetaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchGetFileMetaResponseBody) SetFiles(v []*File) *BatchGetFileMetaResponseBody {
	s.Files = v
	return s
}

func (s *BatchGetFileMetaResponseBody) SetRequestId(v string) *BatchGetFileMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchGetFileMetaResponseBody) Validate() error {
	return dara.Validate(s)
}
