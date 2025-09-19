// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateWebLockFileEventsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *OperateWebLockFileEventsResponseBody
	GetRequestId() *string
}

type OperateWebLockFileEventsResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 2884C186-E8C0-5611-8207-3FF15EE7D9B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OperateWebLockFileEventsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OperateWebLockFileEventsResponseBody) GoString() string {
	return s.String()
}

func (s *OperateWebLockFileEventsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OperateWebLockFileEventsResponseBody) SetRequestId(v string) *OperateWebLockFileEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *OperateWebLockFileEventsResponseBody) Validate() error {
	return dara.Validate(s)
}
