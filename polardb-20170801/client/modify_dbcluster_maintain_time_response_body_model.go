// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterMaintainTimeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDBClusterMaintainTimeResponseBody
	GetRequestId() *string
}

type ModifyDBClusterMaintainTimeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 70656639-1416-479F-AF13-D08197******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBClusterMaintainTimeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterMaintainTimeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterMaintainTimeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBClusterMaintainTimeResponseBody) SetRequestId(v string) *ModifyDBClusterMaintainTimeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBClusterMaintainTimeResponseBody) Validate() error {
	return dara.Validate(s)
}
