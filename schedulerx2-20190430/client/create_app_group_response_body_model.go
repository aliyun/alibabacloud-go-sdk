// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CreateAppGroupResponseBody
	GetCode() *int32
	SetData(v *CreateAppGroupResponseBodyData) *CreateAppGroupResponseBody
	GetData() *CreateAppGroupResponseBodyData
	SetMessage(v string) *CreateAppGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateAppGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateAppGroupResponseBody
	GetSuccess() *bool
}

type CreateAppGroupResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the job group.
	Data *CreateAppGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message that is returned only if the corresponding error occurs.
	//
	// example:
	//
	// Your request is denied as lack of ssl protect.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 883AFE93-FB03-4FA9-A958-E750C6DE120C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the application was created. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateAppGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAppGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppGroupResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreateAppGroupResponseBody) GetData() *CreateAppGroupResponseBodyData {
	return s.Data
}

func (s *CreateAppGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateAppGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAppGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateAppGroupResponseBody) SetCode(v int32) *CreateAppGroupResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAppGroupResponseBody) SetData(v *CreateAppGroupResponseBodyData) *CreateAppGroupResponseBody {
	s.Data = v
	return s
}

func (s *CreateAppGroupResponseBody) SetMessage(v string) *CreateAppGroupResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAppGroupResponseBody) SetRequestId(v string) *CreateAppGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAppGroupResponseBody) SetSuccess(v bool) *CreateAppGroupResponseBody {
	s.Success = &v
	return s
}

func (s *CreateAppGroupResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateAppGroupResponseBodyData struct {
	// The job group ID.
	//
	// example:
	//
	// 6607
	AppGroupId *int64 `json:"AppGroupId,omitempty" xml:"AppGroupId,omitempty"`
	// The AppKey for the application.
	//
	// example:
	//
	// adcExHZviL******
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
}

func (s CreateAppGroupResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateAppGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateAppGroupResponseBodyData) GetAppGroupId() *int64 {
	return s.AppGroupId
}

func (s *CreateAppGroupResponseBodyData) GetAppKey() *string {
	return s.AppKey
}

func (s *CreateAppGroupResponseBodyData) SetAppGroupId(v int64) *CreateAppGroupResponseBodyData {
	s.AppGroupId = &v
	return s
}

func (s *CreateAppGroupResponseBodyData) SetAppKey(v string) *CreateAppGroupResponseBodyData {
	s.AppKey = &v
	return s
}

func (s *CreateAppGroupResponseBodyData) Validate() error {
	return dara.Validate(s)
}
