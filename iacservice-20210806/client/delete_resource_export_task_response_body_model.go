// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteResourceExportTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteResourceExportTaskResponseBody
	GetRequestId() *string
}

type DeleteResourceExportTaskResponseBody struct {
	// example:
	//
	// 136B3926-DD90-5DB2-96EC-8BAD6407D1C9
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteResourceExportTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteResourceExportTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteResourceExportTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteResourceExportTaskResponseBody) SetRequestId(v string) *DeleteResourceExportTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteResourceExportTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
