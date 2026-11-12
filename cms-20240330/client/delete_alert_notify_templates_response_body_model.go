// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAlertNotifyTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeletedCount(v int32) *DeleteAlertNotifyTemplatesResponseBody
	GetDeletedCount() *int32
	SetRequestId(v string) *DeleteAlertNotifyTemplatesResponseBody
	GetRequestId() *string
}

type DeleteAlertNotifyTemplatesResponseBody struct {
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
	// 0B9377D9-C56B-5C2E-A8A4-************
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteAlertNotifyTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAlertNotifyTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAlertNotifyTemplatesResponseBody) GetDeletedCount() *int32 {
	return s.DeletedCount
}

func (s *DeleteAlertNotifyTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAlertNotifyTemplatesResponseBody) SetDeletedCount(v int32) *DeleteAlertNotifyTemplatesResponseBody {
	s.DeletedCount = &v
	return s
}

func (s *DeleteAlertNotifyTemplatesResponseBody) SetRequestId(v string) *DeleteAlertNotifyTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAlertNotifyTemplatesResponseBody) Validate() error {
	return dara.Validate(s)
}
