// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveRuntimeChannelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveRuntimeChannelResponseBody
	GetRequestId() *string
}

type RemoveRuntimeChannelResponseBody struct {
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveRuntimeChannelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveRuntimeChannelResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveRuntimeChannelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveRuntimeChannelResponseBody) SetRequestId(v string) *RemoveRuntimeChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveRuntimeChannelResponseBody) Validate() error {
	return dara.Validate(s)
}
