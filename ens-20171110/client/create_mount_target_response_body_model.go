// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMountTargetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateMountTargetResponseBody
	GetRequestId() *string
	SetStatus(v string) *CreateMountTargetResponseBody
	GetStatus() *string
}

type CreateMountTargetResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 70EACC9C-D07A-4A34-ADA4-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The state of the mount target. Valid values:
	//
	// 	- active
	//
	// 	- inactive
	//
	// 	- pending
	//
	// 	- deleting
	//
	// example:
	//
	// pending
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateMountTargetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMountTargetResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMountTargetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMountTargetResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateMountTargetResponseBody) SetRequestId(v string) *CreateMountTargetResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMountTargetResponseBody) SetStatus(v string) *CreateMountTargetResponseBody {
	s.Status = &v
	return s
}

func (s *CreateMountTargetResponseBody) Validate() error {
	return dara.Validate(s)
}
