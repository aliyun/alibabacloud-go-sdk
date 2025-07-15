// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCloudDriveGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteCloudDriveGroupsResponseBody
	GetCode() *string
	SetData(v string) *DeleteCloudDriveGroupsResponseBody
	GetData() *string
	SetMessage(v string) *DeleteCloudDriveGroupsResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteCloudDriveGroupsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteCloudDriveGroupsResponseBody
	GetSuccess() *bool
}

type DeleteCloudDriveGroupsResponseBody struct {
	// The result of the operation. A value of success indicates that the operation is successful. If the operation failed, an error message is returned.
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
	// The error message returned. This parameter is not returned if the value of Code is `success`.
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
	// Indicates whether the request was successful. Valid values: true: The request is successful. false: The request fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteCloudDriveGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCloudDriveGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCloudDriveGroupsResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteCloudDriveGroupsResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteCloudDriveGroupsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteCloudDriveGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCloudDriveGroupsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteCloudDriveGroupsResponseBody) SetCode(v string) *DeleteCloudDriveGroupsResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteCloudDriveGroupsResponseBody) SetData(v string) *DeleteCloudDriveGroupsResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteCloudDriveGroupsResponseBody) SetMessage(v string) *DeleteCloudDriveGroupsResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteCloudDriveGroupsResponseBody) SetRequestId(v string) *DeleteCloudDriveGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCloudDriveGroupsResponseBody) SetSuccess(v bool) *DeleteCloudDriveGroupsResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteCloudDriveGroupsResponseBody) Validate() error {
	return dara.Validate(s)
}
