// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMaintainWindowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaintainWindowId(v string) *UpdateMaintainWindowResponseBody
	GetMaintainWindowId() *string
	SetRequestId(v string) *UpdateMaintainWindowResponseBody
	GetRequestId() *string
}

type UpdateMaintainWindowResponseBody struct {
	// example:
	//
	// 123-12-312-31-23123
	MaintainWindowId *string `json:"maintainWindowId,omitempty" xml:"maintainWindowId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 0CEC5375-C554-562B-A65F-*****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateMaintainWindowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMaintainWindowResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMaintainWindowResponseBody) GetMaintainWindowId() *string {
	return s.MaintainWindowId
}

func (s *UpdateMaintainWindowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMaintainWindowResponseBody) SetMaintainWindowId(v string) *UpdateMaintainWindowResponseBody {
	s.MaintainWindowId = &v
	return s
}

func (s *UpdateMaintainWindowResponseBody) SetRequestId(v string) *UpdateMaintainWindowResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMaintainWindowResponseBody) Validate() error {
	return dara.Validate(s)
}
