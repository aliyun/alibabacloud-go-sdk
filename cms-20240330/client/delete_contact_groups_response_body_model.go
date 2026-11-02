// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContactGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeletedCount(v int32) *DeleteContactGroupsResponseBody
	GetDeletedCount() *int32
	SetRequestId(v string) *DeleteContactGroupsResponseBody
	GetRequestId() *string
}

type DeleteContactGroupsResponseBody struct {
	// The actual number of deleted records.
	//
	// example:
	//
	// 0
	DeletedCount *int32 `json:"deletedCount,omitempty" xml:"deletedCount,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 8FDE2569-626B-5176-9844-28877A*****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteContactGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteContactGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteContactGroupsResponseBody) GetDeletedCount() *int32 {
	return s.DeletedCount
}

func (s *DeleteContactGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteContactGroupsResponseBody) SetDeletedCount(v int32) *DeleteContactGroupsResponseBody {
	s.DeletedCount = &v
	return s
}

func (s *DeleteContactGroupsResponseBody) SetRequestId(v string) *DeleteContactGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteContactGroupsResponseBody) Validate() error {
	return dara.Validate(s)
}
