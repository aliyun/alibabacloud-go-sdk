// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMPULayoutResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyMPULayoutResponseBody
	GetRequestId() *string
}

type ModifyMPULayoutResponseBody struct {
	// example:
	//
	// 760bad53276431c499e30dc36f6b26be
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyMPULayoutResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyMPULayoutResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyMPULayoutResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyMPULayoutResponseBody) SetRequestId(v string) *ModifyMPULayoutResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyMPULayoutResponseBody) Validate() error {
	return dara.Validate(s)
}
