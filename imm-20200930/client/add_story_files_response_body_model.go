// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddStoryFilesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFiles(v []*AddStoryFilesResponseBodyFiles) *AddStoryFilesResponseBody
	GetFiles() []*AddStoryFilesResponseBodyFiles
	SetRequestId(v string) *AddStoryFilesResponseBody
	GetRequestId() *string
}

type AddStoryFilesResponseBody struct {
	// The objects that were added.
	Files []*AddStoryFilesResponseBodyFiles `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 6E93D6C9-5AC0-49F9-914D-E02678D3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddStoryFilesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddStoryFilesResponseBody) GoString() string {
	return s.String()
}

func (s *AddStoryFilesResponseBody) GetFiles() []*AddStoryFilesResponseBodyFiles {
	return s.Files
}

func (s *AddStoryFilesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddStoryFilesResponseBody) SetFiles(v []*AddStoryFilesResponseBodyFiles) *AddStoryFilesResponseBody {
	s.Files = v
	return s
}

func (s *AddStoryFilesResponseBody) SetRequestId(v string) *AddStoryFilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddStoryFilesResponseBody) Validate() error {
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

type AddStoryFilesResponseBodyFiles struct {
	// The error code.
	//
	// example:
	//
	// ResourceNotFound
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message that is returned.
	//
	// example:
	//
	// The specified resource %s is not found.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The URI of the object.
	//
	// The OSS URI follows the `oss://{bucketname}/{objectname}` format, where `bucketname` is the name of the bucket in the same region as the current project and `objectname` is the path of the object with the extension included.
	//
	// example:
	//
	// oss://test-bucket/test-object
	URI *string `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s AddStoryFilesResponseBodyFiles) String() string {
	return dara.Prettify(s)
}

func (s AddStoryFilesResponseBodyFiles) GoString() string {
	return s.String()
}

func (s *AddStoryFilesResponseBodyFiles) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *AddStoryFilesResponseBodyFiles) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *AddStoryFilesResponseBodyFiles) GetURI() *string {
	return s.URI
}

func (s *AddStoryFilesResponseBodyFiles) SetErrorCode(v string) *AddStoryFilesResponseBodyFiles {
	s.ErrorCode = &v
	return s
}

func (s *AddStoryFilesResponseBodyFiles) SetErrorMessage(v string) *AddStoryFilesResponseBodyFiles {
	s.ErrorMessage = &v
	return s
}

func (s *AddStoryFilesResponseBodyFiles) SetURI(v string) *AddStoryFilesResponseBodyFiles {
	s.URI = &v
	return s
}

func (s *AddStoryFilesResponseBodyFiles) Validate() error {
	return dara.Validate(s)
}
