// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRumSymbolFilesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeletedObjectCount(v int32) *DeleteRumSymbolFilesResponseBody
	GetDeletedObjectCount() *int32
	SetRequestId(v string) *DeleteRumSymbolFilesResponseBody
	GetRequestId() *string
	SetRequestedFileCount(v int32) *DeleteRumSymbolFilesResponseBody
	GetRequestedFileCount() *int32
	SetRequestedObjectCount(v int32) *DeleteRumSymbolFilesResponseBody
	GetRequestedObjectCount() *int32
	SetSuccess(v bool) *DeleteRumSymbolFilesResponseBody
	GetSuccess() *bool
}

type DeleteRumSymbolFilesResponseBody struct {
	// The actual number of objects that were successfully deleted.
	//
	// example:
	//
	// 1
	DeletedObjectCount *int32 `json:"deletedObjectCount,omitempty" xml:"deletedObjectCount,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// request-1
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The number of file paths requested for deletion.
	//
	// example:
	//
	// 1
	RequestedFileCount *int32 `json:"requestedFileCount,omitempty" xml:"requestedFileCount,omitempty"`
	// The total number of deduplicated objects to be deleted.
	//
	// example:
	//
	// 1
	RequestedObjectCount *int32 `json:"requestedObjectCount,omitempty" xml:"requestedObjectCount,omitempty"`
	// Indicates whether all delete operations succeeded. The value is true only if all deduplicated objects to be deleted are successfully deleted. If a partial deletion failure or exception occurs, the API directly returns an error.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteRumSymbolFilesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRumSymbolFilesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRumSymbolFilesResponseBody) GetDeletedObjectCount() *int32 {
	return s.DeletedObjectCount
}

func (s *DeleteRumSymbolFilesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRumSymbolFilesResponseBody) GetRequestedFileCount() *int32 {
	return s.RequestedFileCount
}

func (s *DeleteRumSymbolFilesResponseBody) GetRequestedObjectCount() *int32 {
	return s.RequestedObjectCount
}

func (s *DeleteRumSymbolFilesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteRumSymbolFilesResponseBody) SetDeletedObjectCount(v int32) *DeleteRumSymbolFilesResponseBody {
	s.DeletedObjectCount = &v
	return s
}

func (s *DeleteRumSymbolFilesResponseBody) SetRequestId(v string) *DeleteRumSymbolFilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRumSymbolFilesResponseBody) SetRequestedFileCount(v int32) *DeleteRumSymbolFilesResponseBody {
	s.RequestedFileCount = &v
	return s
}

func (s *DeleteRumSymbolFilesResponseBody) SetRequestedObjectCount(v int32) *DeleteRumSymbolFilesResponseBody {
	s.RequestedObjectCount = &v
	return s
}

func (s *DeleteRumSymbolFilesResponseBody) SetSuccess(v bool) *DeleteRumSymbolFilesResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteRumSymbolFilesResponseBody) Validate() error {
	return dara.Validate(s)
}
