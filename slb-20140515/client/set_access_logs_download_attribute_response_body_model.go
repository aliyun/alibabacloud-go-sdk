// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetAccessLogsDownloadAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetAccessLogsDownloadAttributeResponseBody
	GetRequestId() *string
}

type SetAccessLogsDownloadAttributeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 9DEC9C28-AB05-4DDF-9A78-6B08EC9C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetAccessLogsDownloadAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetAccessLogsDownloadAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *SetAccessLogsDownloadAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetAccessLogsDownloadAttributeResponseBody) SetRequestId(v string) *SetAccessLogsDownloadAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetAccessLogsDownloadAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
