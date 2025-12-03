// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadDeviceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UploadDeviceResponseBody
	GetCode() *string
	SetData(v *UploadDeviceResponseBodyData) *UploadDeviceResponseBody
	GetData() *UploadDeviceResponseBodyData
	SetHttpStatusCode(v int32) *UploadDeviceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UploadDeviceResponseBody
	GetMessage() *string
	SetRequestId(v string) *UploadDeviceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UploadDeviceResponseBody
	GetSuccess() *bool
}

type UploadDeviceResponseBody struct {
	// example:
	//
	// 0
	Code *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *UploadDeviceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// null
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 86C4236B-D6C2-1E31-8370-2FAEC5CFE012
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UploadDeviceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UploadDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *UploadDeviceResponseBody) GetCode() *string {
	return s.Code
}

func (s *UploadDeviceResponseBody) GetData() *UploadDeviceResponseBodyData {
	return s.Data
}

func (s *UploadDeviceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UploadDeviceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UploadDeviceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UploadDeviceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UploadDeviceResponseBody) SetCode(v string) *UploadDeviceResponseBody {
	s.Code = &v
	return s
}

func (s *UploadDeviceResponseBody) SetData(v *UploadDeviceResponseBodyData) *UploadDeviceResponseBody {
	s.Data = v
	return s
}

func (s *UploadDeviceResponseBody) SetHttpStatusCode(v int32) *UploadDeviceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UploadDeviceResponseBody) SetMessage(v string) *UploadDeviceResponseBody {
	s.Message = &v
	return s
}

func (s *UploadDeviceResponseBody) SetRequestId(v string) *UploadDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadDeviceResponseBody) SetSuccess(v bool) *UploadDeviceResponseBody {
	s.Success = &v
	return s
}

func (s *UploadDeviceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UploadDeviceResponseBodyData struct {
	// example:
	//
	// PF835431668603208261
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
}

func (s UploadDeviceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UploadDeviceResponseBodyData) GoString() string {
	return s.String()
}

func (s *UploadDeviceResponseBodyData) GetFileId() *string {
	return s.FileId
}

func (s *UploadDeviceResponseBodyData) SetFileId(v string) *UploadDeviceResponseBodyData {
	s.FileId = &v
	return s
}

func (s *UploadDeviceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
