// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRestartClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyRestartClusterResponseBody
	GetRequestId() *string
}

type ModifyRestartClusterResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyRestartClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyRestartClusterResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyRestartClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyRestartClusterResponseBody) SetRequestId(v string) *ModifyRestartClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyRestartClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
