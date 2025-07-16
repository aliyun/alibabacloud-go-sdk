// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserUsageDataExportTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteUserUsageDataExportTaskResponseBody
	GetRequestId() *string
}

type DeleteUserUsageDataExportTaskResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0AEDAF20-4DDF-4165-8750-47FF9C1929C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteUserUsageDataExportTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserUsageDataExportTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserUsageDataExportTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteUserUsageDataExportTaskResponseBody) SetRequestId(v string) *DeleteUserUsageDataExportTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteUserUsageDataExportTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
