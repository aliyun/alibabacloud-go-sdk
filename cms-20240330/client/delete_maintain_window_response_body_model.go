// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMaintainWindowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaintainWindowId(v string) *DeleteMaintainWindowResponseBody
	GetMaintainWindowId() *string
	SetRequestId(v string) *DeleteMaintainWindowResponseBody
	GetRequestId() *string
}

type DeleteMaintainWindowResponseBody struct {
	// example:
	//
	// 123-12-312-31-23123
	MaintainWindowId *string `json:"maintainWindowId,omitempty" xml:"maintainWindowId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 8FDE2569-626B-5176-9844-28877A*****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteMaintainWindowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMaintainWindowResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMaintainWindowResponseBody) GetMaintainWindowId() *string {
	return s.MaintainWindowId
}

func (s *DeleteMaintainWindowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMaintainWindowResponseBody) SetMaintainWindowId(v string) *DeleteMaintainWindowResponseBody {
	s.MaintainWindowId = &v
	return s
}

func (s *DeleteMaintainWindowResponseBody) SetRequestId(v string) *DeleteMaintainWindowResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMaintainWindowResponseBody) Validate() error {
	return dara.Validate(s)
}
