// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterMonitorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDBClusterMonitorResponseBody
	GetRequestId() *string
}

type ModifyDBClusterMonitorResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 75B92353-73B4-447B-8477-C85F3C******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBClusterMonitorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterMonitorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBClusterMonitorResponseBody) SetRequestId(v string) *ModifyDBClusterMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBClusterMonitorResponseBody) Validate() error {
	return dara.Validate(s)
}
