// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteQueueResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteQueueResponseBody
	GetRequestId() *string
}

type DeleteQueueResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 92385FD2-624A-48C9-8FB5-753F2AFA***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteQueueResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteQueueResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteQueueResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteQueueResponseBody) SetRequestId(v string) *DeleteQueueResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteQueueResponseBody) Validate() error {
	return dara.Validate(s)
}
