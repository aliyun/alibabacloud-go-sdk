// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOncallSchedulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeletedCount(v int32) *DeleteOncallSchedulesResponseBody
	GetDeletedCount() *int32
	SetRequestId(v string) *DeleteOncallSchedulesResponseBody
	GetRequestId() *string
}

type DeleteOncallSchedulesResponseBody struct {
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

func (s DeleteOncallSchedulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteOncallSchedulesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteOncallSchedulesResponseBody) GetDeletedCount() *int32 {
	return s.DeletedCount
}

func (s *DeleteOncallSchedulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteOncallSchedulesResponseBody) SetDeletedCount(v int32) *DeleteOncallSchedulesResponseBody {
	s.DeletedCount = &v
	return s
}

func (s *DeleteOncallSchedulesResponseBody) SetRequestId(v string) *DeleteOncallSchedulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteOncallSchedulesResponseBody) Validate() error {
	return dara.Validate(s)
}
