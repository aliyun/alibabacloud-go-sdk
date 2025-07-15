// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCloudDriveGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateCloudDriveGroupResponseBody
	GetCode() *string
	SetData(v string) *CreateCloudDriveGroupResponseBody
	GetData() *string
	SetMessage(v string) *CreateCloudDriveGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateCloudDriveGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateCloudDriveGroupResponseBody
	GetSuccess() *bool
}

type CreateCloudDriveGroupResponseBody struct {
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
	// The error message that is returned. This parameter is not returned if the value of Code is success.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// FD14D2A1-CC39-5ED3-8EE7-11FDF4B9D6D2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. true: The call is successful. false: The call fails.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateCloudDriveGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudDriveGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCloudDriveGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateCloudDriveGroupResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateCloudDriveGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateCloudDriveGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCloudDriveGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateCloudDriveGroupResponseBody) SetCode(v string) *CreateCloudDriveGroupResponseBody {
	s.Code = &v
	return s
}

func (s *CreateCloudDriveGroupResponseBody) SetData(v string) *CreateCloudDriveGroupResponseBody {
	s.Data = &v
	return s
}

func (s *CreateCloudDriveGroupResponseBody) SetMessage(v string) *CreateCloudDriveGroupResponseBody {
	s.Message = &v
	return s
}

func (s *CreateCloudDriveGroupResponseBody) SetRequestId(v string) *CreateCloudDriveGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCloudDriveGroupResponseBody) SetSuccess(v bool) *CreateCloudDriveGroupResponseBody {
	s.Success = &v
	return s
}

func (s *CreateCloudDriveGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
