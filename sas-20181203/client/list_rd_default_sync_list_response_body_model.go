// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRdDefaultSyncListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListRdDefaultSyncListResponseBodyData) *ListRdDefaultSyncListResponseBody
	GetData() *ListRdDefaultSyncListResponseBodyData
	SetRequestId(v string) *ListRdDefaultSyncListResponseBody
	GetRequestId() *string
}

type ListRdDefaultSyncListResponseBody struct {
	// The data returned if the call is successful.
	Data *ListRdDefaultSyncListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 7532B7EE-7CE7-5F4D-BF04-B12447DDCAE1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListRdDefaultSyncListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRdDefaultSyncListResponseBody) GoString() string {
	return s.String()
}

func (s *ListRdDefaultSyncListResponseBody) GetData() *ListRdDefaultSyncListResponseBodyData {
	return s.Data
}

func (s *ListRdDefaultSyncListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRdDefaultSyncListResponseBody) SetData(v *ListRdDefaultSyncListResponseBodyData) *ListRdDefaultSyncListResponseBody {
	s.Data = v
	return s
}

func (s *ListRdDefaultSyncListResponseBody) SetRequestId(v string) *ListRdDefaultSyncListResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRdDefaultSyncListResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListRdDefaultSyncListResponseBodyData struct {
	// The IDs of the folders in the resource directory.
	//
	// example:
	//
	// fd-BwoXuf****,fd-CFamY7****
	FolderIds *string `json:"FolderIds,omitempty" xml:"FolderIds,omitempty"`
}

func (s ListRdDefaultSyncListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListRdDefaultSyncListResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListRdDefaultSyncListResponseBodyData) GetFolderIds() *string {
	return s.FolderIds
}

func (s *ListRdDefaultSyncListResponseBodyData) SetFolderIds(v string) *ListRdDefaultSyncListResponseBodyData {
	s.FolderIds = &v
	return s
}

func (s *ListRdDefaultSyncListResponseBodyData) Validate() error {
	return dara.Validate(s)
}
