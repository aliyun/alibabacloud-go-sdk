// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchUpdateFileMetaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFiles(v []*BatchUpdateFileMetaResponseBodyFiles) *BatchUpdateFileMetaResponseBody
	GetFiles() []*BatchUpdateFileMetaResponseBodyFiles
	SetRequestId(v string) *BatchUpdateFileMetaResponseBody
	GetRequestId() *string
}

type BatchUpdateFileMetaResponseBody struct {
	// The files whose metadata was updated.
	Files []*BatchUpdateFileMetaResponseBodyFiles `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// F5BF215E-3237-0852-B9C6-F233D44A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchUpdateFileMetaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchUpdateFileMetaResponseBody) GoString() string {
	return s.String()
}

func (s *BatchUpdateFileMetaResponseBody) GetFiles() []*BatchUpdateFileMetaResponseBodyFiles {
	return s.Files
}

func (s *BatchUpdateFileMetaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchUpdateFileMetaResponseBody) SetFiles(v []*BatchUpdateFileMetaResponseBodyFiles) *BatchUpdateFileMetaResponseBody {
	s.Files = v
	return s
}

func (s *BatchUpdateFileMetaResponseBody) SetRequestId(v string) *BatchUpdateFileMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchUpdateFileMetaResponseBody) Validate() error {
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

type BatchUpdateFileMetaResponseBodyFiles struct {
	// The error message returned when the value of the Success parameter is false.
	//
	// example:
	//
	// *error.OpError : InvalidArgument | Index KV count exceeded, should be no more than 100.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// Enumerated values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The URI of the file.
	//
	// example:
	//
	// oss://examplebucket/example.jpg
	URI *string `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s BatchUpdateFileMetaResponseBodyFiles) String() string {
	return dara.Prettify(s)
}

func (s BatchUpdateFileMetaResponseBodyFiles) GoString() string {
	return s.String()
}

func (s *BatchUpdateFileMetaResponseBodyFiles) GetMessage() *string {
	return s.Message
}

func (s *BatchUpdateFileMetaResponseBodyFiles) GetSuccess() *bool {
	return s.Success
}

func (s *BatchUpdateFileMetaResponseBodyFiles) GetURI() *string {
	return s.URI
}

func (s *BatchUpdateFileMetaResponseBodyFiles) SetMessage(v string) *BatchUpdateFileMetaResponseBodyFiles {
	s.Message = &v
	return s
}

func (s *BatchUpdateFileMetaResponseBodyFiles) SetSuccess(v bool) *BatchUpdateFileMetaResponseBodyFiles {
	s.Success = &v
	return s
}

func (s *BatchUpdateFileMetaResponseBodyFiles) SetURI(v string) *BatchUpdateFileMetaResponseBodyFiles {
	s.URI = &v
	return s
}

func (s *BatchUpdateFileMetaResponseBodyFiles) Validate() error {
	return dara.Validate(s)
}
