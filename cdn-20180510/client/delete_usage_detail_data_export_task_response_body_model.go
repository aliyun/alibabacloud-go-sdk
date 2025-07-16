// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUsageDetailDataExportTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteUsageDetailDataExportTaskResponseBody
	GetRequestId() *string
}

type DeleteUsageDetailDataExportTaskResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0AEDAF20-4DDF-4165-8750-47FF9C1929C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteUsageDetailDataExportTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteUsageDetailDataExportTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUsageDetailDataExportTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteUsageDetailDataExportTaskResponseBody) SetRequestId(v string) *DeleteUsageDetailDataExportTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteUsageDetailDataExportTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
