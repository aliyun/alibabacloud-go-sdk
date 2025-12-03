// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyClusterNetTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyClusterNetTypeResponseBody
	GetRequestId() *string
}

type ModifyClusterNetTypeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyClusterNetTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterNetTypeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyClusterNetTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyClusterNetTypeResponseBody) SetRequestId(v string) *ModifyClusterNetTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyClusterNetTypeResponseBody) Validate() error {
	return dara.Validate(s)
}
