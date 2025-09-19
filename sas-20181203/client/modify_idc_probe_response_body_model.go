// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyIdcProbeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyIdcProbeResponseBody
	GetRequestId() *string
}

type ModifyIdcProbeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 20456DD5-5CBF-5015-9173-12CA4246B***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyIdcProbeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyIdcProbeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyIdcProbeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyIdcProbeResponseBody) SetRequestId(v string) *ModifyIdcProbeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyIdcProbeResponseBody) Validate() error {
	return dara.Validate(s)
}
