// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCloudDriveGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyCloudDriveGroupsResponseBody
	GetCode() *string
	SetData(v string) *ModifyCloudDriveGroupsResponseBody
	GetData() *string
	SetMessage(v string) *ModifyCloudDriveGroupsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyCloudDriveGroupsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyCloudDriveGroupsResponseBody
	GetSuccess() *bool
}

type ModifyCloudDriveGroupsResponseBody struct {
	// The returned results. A value of success indicates that the operation is successful. If the operation failed, an error message is returned.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data information.
	//
	// example:
	//
	// []
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The message returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values: true and false.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyCloudDriveGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyCloudDriveGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCloudDriveGroupsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyCloudDriveGroupsResponseBody) GetData() *string {
	return s.Data
}

func (s *ModifyCloudDriveGroupsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyCloudDriveGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyCloudDriveGroupsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyCloudDriveGroupsResponseBody) SetCode(v string) *ModifyCloudDriveGroupsResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyCloudDriveGroupsResponseBody) SetData(v string) *ModifyCloudDriveGroupsResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyCloudDriveGroupsResponseBody) SetMessage(v string) *ModifyCloudDriveGroupsResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyCloudDriveGroupsResponseBody) SetRequestId(v string) *ModifyCloudDriveGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyCloudDriveGroupsResponseBody) SetSuccess(v bool) *ModifyCloudDriveGroupsResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyCloudDriveGroupsResponseBody) Validate() error {
	return dara.Validate(s)
}
