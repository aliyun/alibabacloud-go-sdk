// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFileTagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateFileTagResponseBody
	GetCode() *string
	SetData(v *UpdateFileTagResponseBodyData) *UpdateFileTagResponseBody
	GetData() *UpdateFileTagResponseBodyData
	SetMessage(v string) *UpdateFileTagResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateFileTagResponseBody
	GetRequestId() *string
	SetStatus(v string) *UpdateFileTagResponseBody
	GetStatus() *string
	SetSuccess(v bool) *UpdateFileTagResponseBody
	GetSuccess() *bool
}

type UpdateFileTagResponseBody struct {
	// example:
	//
	// Success
	Code *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *UpdateFileTagResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Requests throttling triggered.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// RequestId
	//
	// example:
	//
	// 35A267BF-xxxx-54DB-8394-AA3B0742D833
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 200
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateFileTagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateFileTagResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFileTagResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateFileTagResponseBody) GetData() *UpdateFileTagResponseBodyData {
	return s.Data
}

func (s *UpdateFileTagResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateFileTagResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateFileTagResponseBody) GetStatus() *string {
	return s.Status
}

func (s *UpdateFileTagResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateFileTagResponseBody) SetCode(v string) *UpdateFileTagResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateFileTagResponseBody) SetData(v *UpdateFileTagResponseBodyData) *UpdateFileTagResponseBody {
	s.Data = v
	return s
}

func (s *UpdateFileTagResponseBody) SetMessage(v string) *UpdateFileTagResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateFileTagResponseBody) SetRequestId(v string) *UpdateFileTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFileTagResponseBody) SetStatus(v string) *UpdateFileTagResponseBody {
	s.Status = &v
	return s
}

func (s *UpdateFileTagResponseBody) SetSuccess(v bool) *UpdateFileTagResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateFileTagResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateFileTagResponseBodyData struct {
	// example:
	//
	// file_9a65732555b54d5ea10796ca5742ba22_xxxxxxxx
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
}

func (s UpdateFileTagResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateFileTagResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateFileTagResponseBodyData) GetFileId() *string {
	return s.FileId
}

func (s *UpdateFileTagResponseBodyData) SetFileId(v string) *UpdateFileTagResponseBodyData {
	s.FileId = &v
	return s
}

func (s *UpdateFileTagResponseBodyData) Validate() error {
	return dara.Validate(s)
}
