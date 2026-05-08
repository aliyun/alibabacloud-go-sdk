// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContextsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeletedCount(v int32) *DeleteContextsResponseBody
	GetDeletedCount() *int32
	SetRequestId(v string) *DeleteContextsResponseBody
	GetRequestId() *string
}

type DeleteContextsResponseBody struct {
	// example:
	//
	// 55
	DeletedCount *int32 `json:"deletedCount,omitempty" xml:"deletedCount,omitempty"`
	// example:
	//
	// 8FDE2569-626B-5176-9844-28877A*****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteContextsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteContextsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteContextsResponseBody) GetDeletedCount() *int32 {
	return s.DeletedCount
}

func (s *DeleteContextsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteContextsResponseBody) SetDeletedCount(v int32) *DeleteContextsResponseBody {
	s.DeletedCount = &v
	return s
}

func (s *DeleteContextsResponseBody) SetRequestId(v string) *DeleteContextsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteContextsResponseBody) Validate() error {
	return dara.Validate(s)
}
