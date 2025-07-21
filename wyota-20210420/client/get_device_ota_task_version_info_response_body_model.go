// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeviceOtaTaskVersionInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetDeviceOtaTaskVersionInfoResponseBody
	GetCode() *string
	SetData(v *GetDeviceOtaTaskVersionInfoResponseBodyData) *GetDeviceOtaTaskVersionInfoResponseBody
	GetData() *GetDeviceOtaTaskVersionInfoResponseBodyData
	SetMessage(v string) *GetDeviceOtaTaskVersionInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDeviceOtaTaskVersionInfoResponseBody
	GetRequestId() *string
}

type GetDeviceOtaTaskVersionInfoResponseBody struct {
	Code      *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetDeviceOtaTaskVersionInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDeviceOtaTaskVersionInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceOtaTaskVersionInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeviceOtaTaskVersionInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDeviceOtaTaskVersionInfoResponseBody) GetData() *GetDeviceOtaTaskVersionInfoResponseBodyData {
	return s.Data
}

func (s *GetDeviceOtaTaskVersionInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDeviceOtaTaskVersionInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDeviceOtaTaskVersionInfoResponseBody) SetCode(v string) *GetDeviceOtaTaskVersionInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetDeviceOtaTaskVersionInfoResponseBody) SetData(v *GetDeviceOtaTaskVersionInfoResponseBodyData) *GetDeviceOtaTaskVersionInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetDeviceOtaTaskVersionInfoResponseBody) SetMessage(v string) *GetDeviceOtaTaskVersionInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetDeviceOtaTaskVersionInfoResponseBody) SetRequestId(v string) *GetDeviceOtaTaskVersionInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDeviceOtaTaskVersionInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDeviceOtaTaskVersionInfoResponseBodyData struct {
	ReleaseNote *string `json:"ReleaseNote,omitempty" xml:"ReleaseNote,omitempty"`
	Size        *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
	Version     *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetDeviceOtaTaskVersionInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceOtaTaskVersionInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDeviceOtaTaskVersionInfoResponseBodyData) GetReleaseNote() *string {
	return s.ReleaseNote
}

func (s *GetDeviceOtaTaskVersionInfoResponseBodyData) GetSize() *int64 {
	return s.Size
}

func (s *GetDeviceOtaTaskVersionInfoResponseBodyData) GetVersion() *string {
	return s.Version
}

func (s *GetDeviceOtaTaskVersionInfoResponseBodyData) SetReleaseNote(v string) *GetDeviceOtaTaskVersionInfoResponseBodyData {
	s.ReleaseNote = &v
	return s
}

func (s *GetDeviceOtaTaskVersionInfoResponseBodyData) SetSize(v int64) *GetDeviceOtaTaskVersionInfoResponseBodyData {
	s.Size = &v
	return s
}

func (s *GetDeviceOtaTaskVersionInfoResponseBodyData) SetVersion(v string) *GetDeviceOtaTaskVersionInfoResponseBodyData {
	s.Version = &v
	return s
}

func (s *GetDeviceOtaTaskVersionInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}
