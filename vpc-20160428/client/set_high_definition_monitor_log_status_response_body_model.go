// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetHighDefinitionMonitorLogStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetHighDefinitionMonitorLogStatusResponseBody
	GetRequestId() *string
	SetSuccess(v string) *SetHighDefinitionMonitorLogStatusResponseBody
	GetSuccess() *string
}

type SetHighDefinitionMonitorLogStatusResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 54B48E3D-DF70-471B-AA93-08E683A1B457
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation is performed. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetHighDefinitionMonitorLogStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetHighDefinitionMonitorLogStatusResponseBody) GoString() string {
	return s.String()
}

func (s *SetHighDefinitionMonitorLogStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetHighDefinitionMonitorLogStatusResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *SetHighDefinitionMonitorLogStatusResponseBody) SetRequestId(v string) *SetHighDefinitionMonitorLogStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetHighDefinitionMonitorLogStatusResponseBody) SetSuccess(v string) *SetHighDefinitionMonitorLogStatusResponseBody {
	s.Success = &v
	return s
}

func (s *SetHighDefinitionMonitorLogStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
