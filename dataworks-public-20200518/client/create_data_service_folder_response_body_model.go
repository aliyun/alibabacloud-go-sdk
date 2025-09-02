// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataServiceFolderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFolderId(v int64) *CreateDataServiceFolderResponseBody
	GetFolderId() *int64
	SetRequestId(v string) *CreateDataServiceFolderResponseBody
	GetRequestId() *string
}

type CreateDataServiceFolderResponseBody struct {
	// The ID of the created folder.
	//
	// example:
	//
	// 123
	FolderId *int64 `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0000-ABCD-E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDataServiceFolderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDataServiceFolderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataServiceFolderResponseBody) GetFolderId() *int64 {
	return s.FolderId
}

func (s *CreateDataServiceFolderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDataServiceFolderResponseBody) SetFolderId(v int64) *CreateDataServiceFolderResponseBody {
	s.FolderId = &v
	return s
}

func (s *CreateDataServiceFolderResponseBody) SetRequestId(v string) *CreateDataServiceFolderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDataServiceFolderResponseBody) Validate() error {
	return dara.Validate(s)
}
