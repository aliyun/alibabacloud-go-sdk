// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadDataV4ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UploadDataV4ResponseBody
	GetCode() *string
	SetData(v string) *UploadDataV4ResponseBody
	GetData() *string
	SetMessage(v string) *UploadDataV4ResponseBody
	GetMessage() *string
	SetRequestId(v string) *UploadDataV4ResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UploadDataV4ResponseBody
	GetSuccess() *bool
}

type UploadDataV4ResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 6F5934C7-C223-4F0F-BBF3-5B3594***
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 6F5934C7-C223-4F0F-BBF3-5B3594***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UploadDataV4ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UploadDataV4ResponseBody) GoString() string {
	return s.String()
}

func (s *UploadDataV4ResponseBody) GetCode() *string {
	return s.Code
}

func (s *UploadDataV4ResponseBody) GetData() *string {
	return s.Data
}

func (s *UploadDataV4ResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UploadDataV4ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UploadDataV4ResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UploadDataV4ResponseBody) SetCode(v string) *UploadDataV4ResponseBody {
	s.Code = &v
	return s
}

func (s *UploadDataV4ResponseBody) SetData(v string) *UploadDataV4ResponseBody {
	s.Data = &v
	return s
}

func (s *UploadDataV4ResponseBody) SetMessage(v string) *UploadDataV4ResponseBody {
	s.Message = &v
	return s
}

func (s *UploadDataV4ResponseBody) SetRequestId(v string) *UploadDataV4ResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadDataV4ResponseBody) SetSuccess(v bool) *UploadDataV4ResponseBody {
	s.Success = &v
	return s
}

func (s *UploadDataV4ResponseBody) Validate() error {
	return dara.Validate(s)
}
