// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDownloadTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDownloadTaskResponseBody
	GetRequestId() *string
}

type DeleteDownloadTaskResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 064022A8-F415-572C-B3C1-657152833F11
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDownloadTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDownloadTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDownloadTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDownloadTaskResponseBody) SetRequestId(v string) *DeleteDownloadTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDownloadTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
