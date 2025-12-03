// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAccessLogsDownloadAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAccessLogsDownloadAttributeResponseBody
	GetRequestId() *string
}

type DeleteAccessLogsDownloadAttributeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 9DEC9C28-AB05-4DDF-9A78-6B08EC9C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAccessLogsDownloadAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAccessLogsDownloadAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAccessLogsDownloadAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAccessLogsDownloadAttributeResponseBody) SetRequestId(v string) *DeleteAccessLogsDownloadAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAccessLogsDownloadAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
