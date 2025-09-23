// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterVpcResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDBClusterVpcResponseBody
	GetRequestId() *string
}

type ModifyDBClusterVpcResponseBody struct {
	// example:
	//
	// EBEAA83D-1734-42E3-85E3-E25F6E******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBClusterVpcResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterVpcResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterVpcResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBClusterVpcResponseBody) SetRequestId(v string) *ModifyDBClusterVpcResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBClusterVpcResponseBody) Validate() error {
	return dara.Validate(s)
}
