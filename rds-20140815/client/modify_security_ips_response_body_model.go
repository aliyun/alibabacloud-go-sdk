// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySecurityIpsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifySecurityIpsResponseBody
	GetRequestId() *string
	SetTaskId(v string) *ModifySecurityIpsResponseBody
	GetTaskId() *string
}

type ModifySecurityIpsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ModifySecurityIpsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySecurityIpsResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySecurityIpsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySecurityIpsResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *ModifySecurityIpsResponseBody) SetRequestId(v string) *ModifySecurityIpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySecurityIpsResponseBody) SetTaskId(v string) *ModifySecurityIpsResponseBody {
	s.TaskId = &v
	return s
}

func (s *ModifySecurityIpsResponseBody) Validate() error {
	return dara.Validate(s)
}
