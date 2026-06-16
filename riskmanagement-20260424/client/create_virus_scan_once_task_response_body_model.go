// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVirusScanOnceTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateVirusScanOnceTaskResponseBody
	GetCode() *string
	SetData(v *CreateVirusScanOnceTaskResponseBodyData) *CreateVirusScanOnceTaskResponseBody
	GetData() *CreateVirusScanOnceTaskResponseBodyData
	SetMessage(v string) *CreateVirusScanOnceTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateVirusScanOnceTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateVirusScanOnceTaskResponseBody
	GetSuccess() *bool
}

type CreateVirusScanOnceTaskResponseBody struct {
	Code      *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *CreateVirusScanOnceTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateVirusScanOnceTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVirusScanOnceTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVirusScanOnceTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateVirusScanOnceTaskResponseBody) GetData() *CreateVirusScanOnceTaskResponseBodyData {
	return s.Data
}

func (s *CreateVirusScanOnceTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateVirusScanOnceTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVirusScanOnceTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateVirusScanOnceTaskResponseBody) SetCode(v string) *CreateVirusScanOnceTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateVirusScanOnceTaskResponseBody) SetData(v *CreateVirusScanOnceTaskResponseBodyData) *CreateVirusScanOnceTaskResponseBody {
	s.Data = v
	return s
}

func (s *CreateVirusScanOnceTaskResponseBody) SetMessage(v string) *CreateVirusScanOnceTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateVirusScanOnceTaskResponseBody) SetRequestId(v string) *CreateVirusScanOnceTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVirusScanOnceTaskResponseBody) SetSuccess(v bool) *CreateVirusScanOnceTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CreateVirusScanOnceTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateVirusScanOnceTaskResponseBodyData struct {
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	Platform     *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SelectionKey *int32  `json:"SelectionKey,omitempty" xml:"SelectionKey,omitempty"`
	TargetType   *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	Uuid         *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s CreateVirusScanOnceTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateVirusScanOnceTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateVirusScanOnceTaskResponseBodyData) GetBusinessType() *string {
	return s.BusinessType
}

func (s *CreateVirusScanOnceTaskResponseBodyData) GetPlatform() *string {
	return s.Platform
}

func (s *CreateVirusScanOnceTaskResponseBodyData) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVirusScanOnceTaskResponseBodyData) GetSelectionKey() *int32 {
	return s.SelectionKey
}

func (s *CreateVirusScanOnceTaskResponseBodyData) GetTargetType() *string {
	return s.TargetType
}

func (s *CreateVirusScanOnceTaskResponseBodyData) GetUuid() *string {
	return s.Uuid
}

func (s *CreateVirusScanOnceTaskResponseBodyData) SetBusinessType(v string) *CreateVirusScanOnceTaskResponseBodyData {
	s.BusinessType = &v
	return s
}

func (s *CreateVirusScanOnceTaskResponseBodyData) SetPlatform(v string) *CreateVirusScanOnceTaskResponseBodyData {
	s.Platform = &v
	return s
}

func (s *CreateVirusScanOnceTaskResponseBodyData) SetRequestId(v string) *CreateVirusScanOnceTaskResponseBodyData {
	s.RequestId = &v
	return s
}

func (s *CreateVirusScanOnceTaskResponseBodyData) SetSelectionKey(v int32) *CreateVirusScanOnceTaskResponseBodyData {
	s.SelectionKey = &v
	return s
}

func (s *CreateVirusScanOnceTaskResponseBodyData) SetTargetType(v string) *CreateVirusScanOnceTaskResponseBodyData {
	s.TargetType = &v
	return s
}

func (s *CreateVirusScanOnceTaskResponseBodyData) SetUuid(v string) *CreateVirusScanOnceTaskResponseBodyData {
	s.Uuid = &v
	return s
}

func (s *CreateVirusScanOnceTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
