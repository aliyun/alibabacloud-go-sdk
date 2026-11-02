// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAlertActionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeletedCount(v int32) *DeleteAlertActionsResponseBody
	GetDeletedCount() *int32
	SetRequestId(v string) *DeleteAlertActionsResponseBody
	GetRequestId() *string
}

type DeleteAlertActionsResponseBody struct {
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

func (s DeleteAlertActionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAlertActionsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAlertActionsResponseBody) GetDeletedCount() *int32 {
	return s.DeletedCount
}

func (s *DeleteAlertActionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAlertActionsResponseBody) SetDeletedCount(v int32) *DeleteAlertActionsResponseBody {
	s.DeletedCount = &v
	return s
}

func (s *DeleteAlertActionsResponseBody) SetRequestId(v string) *DeleteAlertActionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAlertActionsResponseBody) Validate() error {
	return dara.Validate(s)
}
