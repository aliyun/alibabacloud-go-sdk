// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyActiveOperationMaintenanceConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyActiveOperationMaintenanceConfigResponseBody
	GetRequestId() *string
}

type ModifyActiveOperationMaintenanceConfigResponseBody struct {
	// example:
	//
	// D8F1D721-6439-4257-A89C-F1E8E9C9****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyActiveOperationMaintenanceConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyActiveOperationMaintenanceConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyActiveOperationMaintenanceConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyActiveOperationMaintenanceConfigResponseBody) SetRequestId(v string) *ModifyActiveOperationMaintenanceConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyActiveOperationMaintenanceConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
