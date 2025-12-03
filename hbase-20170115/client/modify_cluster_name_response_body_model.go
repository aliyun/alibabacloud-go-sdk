// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyClusterNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyClusterNameResponseBody
	GetRequestId() *string
}

type ModifyClusterNameResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyClusterNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterNameResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyClusterNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyClusterNameResponseBody) SetRequestId(v string) *ModifyClusterNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyClusterNameResponseBody) Validate() error {
	return dara.Validate(s)
}
