// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveCdsFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *MoveCdsFileResponseBody
	GetCode() *string
	SetMessage(v string) *MoveCdsFileResponseBody
	GetMessage() *string
	SetMoveCdsFileModel(v *MoveCdsFileResponseBodyMoveCdsFileModel) *MoveCdsFileResponseBody
	GetMoveCdsFileModel() *MoveCdsFileResponseBodyMoveCdsFileModel
	SetRequestId(v string) *MoveCdsFileResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *MoveCdsFileResponseBody
	GetSuccess() *bool
}

type MoveCdsFileResponseBody struct {
	// The result of the modification. A value of success indicates that the modification is successful. If the modification failed, an error message is returned.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message that is returned. This parameter is not returned if the value of Code is success.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The response object when you move a file.
	MoveCdsFileModel *MoveCdsFileResponseBodyMoveCdsFileModel `json:"MoveCdsFileModel,omitempty" xml:"MoveCdsFileModel,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// Valid values:
	//
	// 	- <!-- -->
	//
	//     true
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- <!-- -->
	//
	//     false
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s MoveCdsFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MoveCdsFileResponseBody) GoString() string {
	return s.String()
}

func (s *MoveCdsFileResponseBody) GetCode() *string {
	return s.Code
}

func (s *MoveCdsFileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *MoveCdsFileResponseBody) GetMoveCdsFileModel() *MoveCdsFileResponseBodyMoveCdsFileModel {
	return s.MoveCdsFileModel
}

func (s *MoveCdsFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MoveCdsFileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *MoveCdsFileResponseBody) SetCode(v string) *MoveCdsFileResponseBody {
	s.Code = &v
	return s
}

func (s *MoveCdsFileResponseBody) SetMessage(v string) *MoveCdsFileResponseBody {
	s.Message = &v
	return s
}

func (s *MoveCdsFileResponseBody) SetMoveCdsFileModel(v *MoveCdsFileResponseBodyMoveCdsFileModel) *MoveCdsFileResponseBody {
	s.MoveCdsFileModel = v
	return s
}

func (s *MoveCdsFileResponseBody) SetRequestId(v string) *MoveCdsFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *MoveCdsFileResponseBody) SetSuccess(v bool) *MoveCdsFileResponseBody {
	s.Success = &v
	return s
}

func (s *MoveCdsFileResponseBody) Validate() error {
	if s.MoveCdsFileModel != nil {
		if err := s.MoveCdsFileModel.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MoveCdsFileResponseBodyMoveCdsFileModel struct {
	// The ID of the asynchronous task. This parameter is not returned if you copy files. This parameter is returned if you copy folders in the backend in an asynchronous manner. You can call the GetAsyncTask operation to obtain the ID and details of an asynchronous task.
	//
	// example:
	//
	// fe307518-825a-4c8b-a69c-958f0e8a****
	AsyncTaskId *string `json:"AsyncTaskId,omitempty" xml:"AsyncTaskId,omitempty"`
	// Indicates whether the file exists.
	//
	// Valid values:
	//
	// 	- <!-- -->
	//
	//     true
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- <!-- -->
	//
	//     false
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// false
	Exist *bool `json:"Exist,omitempty" xml:"Exist,omitempty"`
	// The ID of the file.
	//
	// example:
	//
	// 63636837e47e5a24a8a940218bef395c210e****
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
}

func (s MoveCdsFileResponseBodyMoveCdsFileModel) String() string {
	return dara.Prettify(s)
}

func (s MoveCdsFileResponseBodyMoveCdsFileModel) GoString() string {
	return s.String()
}

func (s *MoveCdsFileResponseBodyMoveCdsFileModel) GetAsyncTaskId() *string {
	return s.AsyncTaskId
}

func (s *MoveCdsFileResponseBodyMoveCdsFileModel) GetExist() *bool {
	return s.Exist
}

func (s *MoveCdsFileResponseBodyMoveCdsFileModel) GetFileId() *string {
	return s.FileId
}

func (s *MoveCdsFileResponseBodyMoveCdsFileModel) SetAsyncTaskId(v string) *MoveCdsFileResponseBodyMoveCdsFileModel {
	s.AsyncTaskId = &v
	return s
}

func (s *MoveCdsFileResponseBodyMoveCdsFileModel) SetExist(v bool) *MoveCdsFileResponseBodyMoveCdsFileModel {
	s.Exist = &v
	return s
}

func (s *MoveCdsFileResponseBodyMoveCdsFileModel) SetFileId(v string) *MoveCdsFileResponseBodyMoveCdsFileModel {
	s.FileId = &v
	return s
}

func (s *MoveCdsFileResponseBodyMoveCdsFileModel) Validate() error {
	return dara.Validate(s)
}
