// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContactsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeletedCount(v int32) *DeleteContactsResponseBody
	GetDeletedCount() *int32
	SetRequestId(v string) *DeleteContactsResponseBody
	GetRequestId() *string
}

type DeleteContactsResponseBody struct {
	// The actual number of deleted records.
	//
	// example:
	//
	// 0
	DeletedCount *int32 `json:"deletedCount,omitempty" xml:"deletedCount,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8FDE2569-626B-5176-9844-28877A*****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteContactsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteContactsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteContactsResponseBody) GetDeletedCount() *int32 {
	return s.DeletedCount
}

func (s *DeleteContactsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteContactsResponseBody) SetDeletedCount(v int32) *DeleteContactsResponseBody {
	s.DeletedCount = &v
	return s
}

func (s *DeleteContactsResponseBody) SetRequestId(v string) *DeleteContactsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteContactsResponseBody) Validate() error {
	return dara.Validate(s)
}
