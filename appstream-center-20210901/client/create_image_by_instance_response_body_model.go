// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateImageByInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateImageByInstanceResponseBody
	GetCode() *string
	SetData(v *CreateImageByInstanceResponseBodyData) *CreateImageByInstanceResponseBody
	GetData() *CreateImageByInstanceResponseBodyData
	SetMessage(v string) *CreateImageByInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateImageByInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateImageByInstanceResponseBody
	GetSuccess() *bool
}

type CreateImageByInstanceResponseBody struct {
	// The error code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *CreateImageByInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// NULL
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

func (s CreateImageByInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateImageByInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateImageByInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateImageByInstanceResponseBody) GetData() *CreateImageByInstanceResponseBodyData {
	return s.Data
}

func (s *CreateImageByInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateImageByInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateImageByInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateImageByInstanceResponseBody) SetCode(v string) *CreateImageByInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateImageByInstanceResponseBody) SetData(v *CreateImageByInstanceResponseBodyData) *CreateImageByInstanceResponseBody {
	s.Data = v
	return s
}

func (s *CreateImageByInstanceResponseBody) SetMessage(v string) *CreateImageByInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *CreateImageByInstanceResponseBody) SetRequestId(v string) *CreateImageByInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateImageByInstanceResponseBody) SetSuccess(v bool) *CreateImageByInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *CreateImageByInstanceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateImageByInstanceResponseBodyData struct {
	// The ID of the RDS image.
	//
	// example:
	//
	// imgc-07hnjj5fp****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The ID of the image creation task.
	//
	// example:
	//
	// tid-0abxi0lbih******
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The version of the image.
	//
	// example:
	//
	// iv-0abxi0lbi*****
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s CreateImageByInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateImageByInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateImageByInstanceResponseBodyData) GetImageId() *string {
	return s.ImageId
}

func (s *CreateImageByInstanceResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateImageByInstanceResponseBodyData) GetVersionId() *string {
	return s.VersionId
}

func (s *CreateImageByInstanceResponseBodyData) SetImageId(v string) *CreateImageByInstanceResponseBodyData {
	s.ImageId = &v
	return s
}

func (s *CreateImageByInstanceResponseBodyData) SetTaskId(v string) *CreateImageByInstanceResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *CreateImageByInstanceResponseBodyData) SetVersionId(v string) *CreateImageByInstanceResponseBodyData {
	s.VersionId = &v
	return s
}

func (s *CreateImageByInstanceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
