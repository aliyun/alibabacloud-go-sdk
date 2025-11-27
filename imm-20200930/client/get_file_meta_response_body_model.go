// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileMetaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFiles(v []*File) *GetFileMetaResponseBody
	GetFiles() []*File
	SetRequestId(v string) *GetFileMetaResponseBody
	GetRequestId() *string
}

type GetFileMetaResponseBody struct {
	// The metadata returned.
	Files []*File `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 7F84C6D9-5AC0-49F9-914D-F02678E3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetFileMetaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFileMetaResponseBody) GoString() string {
	return s.String()
}

func (s *GetFileMetaResponseBody) GetFiles() []*File {
	return s.Files
}

func (s *GetFileMetaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFileMetaResponseBody) SetFiles(v []*File) *GetFileMetaResponseBody {
	s.Files = v
	return s
}

func (s *GetFileMetaResponseBody) SetRequestId(v string) *GetFileMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFileMetaResponseBody) Validate() error {
	if s.Files != nil {
		for _, item := range s.Files {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
