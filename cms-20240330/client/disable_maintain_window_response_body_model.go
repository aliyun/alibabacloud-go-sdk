// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableMaintainWindowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaintainWindowId(v string) *DisableMaintainWindowResponseBody
	GetMaintainWindowId() *string
	SetRequestId(v string) *DisableMaintainWindowResponseBody
	GetRequestId() *string
}

type DisableMaintainWindowResponseBody struct {
	// example:
	//
	// 123-12-312-31-23123
	MaintainWindowId *string `json:"maintainWindowId,omitempty" xml:"maintainWindowId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 123123-3213-345-9941-345345345
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DisableMaintainWindowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableMaintainWindowResponseBody) GoString() string {
	return s.String()
}

func (s *DisableMaintainWindowResponseBody) GetMaintainWindowId() *string {
	return s.MaintainWindowId
}

func (s *DisableMaintainWindowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableMaintainWindowResponseBody) SetMaintainWindowId(v string) *DisableMaintainWindowResponseBody {
	s.MaintainWindowId = &v
	return s
}

func (s *DisableMaintainWindowResponseBody) SetRequestId(v string) *DisableMaintainWindowResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableMaintainWindowResponseBody) Validate() error {
	return dara.Validate(s)
}
