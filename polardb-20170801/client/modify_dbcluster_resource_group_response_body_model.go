// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterResourceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDBClusterResourceGroupResponseBody
	GetRequestId() *string
}

type ModifyDBClusterResourceGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 70656639-1416-479F-AF13-D08197******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBClusterResourceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterResourceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBClusterResourceGroupResponseBody) SetRequestId(v string) *ModifyDBClusterResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBClusterResourceGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
