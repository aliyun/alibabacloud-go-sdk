// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSmsSignResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListSmsSignResponseBody
	GetCode() *string
	SetData(v []*ListSmsSignResponseBodyData) *ListSmsSignResponseBody
	GetData() []*ListSmsSignResponseBodyData
	SetMessage(v string) *ListSmsSignResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListSmsSignResponseBody
	GetRequestId() *string
}

type ListSmsSignResponseBody struct {
	Code      *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*ListSmsSignResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListSmsSignResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSmsSignResponseBody) GoString() string {
	return s.String()
}

func (s *ListSmsSignResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListSmsSignResponseBody) GetData() []*ListSmsSignResponseBodyData {
	return s.Data
}

func (s *ListSmsSignResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListSmsSignResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSmsSignResponseBody) SetCode(v string) *ListSmsSignResponseBody {
	s.Code = &v
	return s
}

func (s *ListSmsSignResponseBody) SetData(v []*ListSmsSignResponseBodyData) *ListSmsSignResponseBody {
	s.Data = v
	return s
}

func (s *ListSmsSignResponseBody) SetMessage(v string) *ListSmsSignResponseBody {
	s.Message = &v
	return s
}

func (s *ListSmsSignResponseBody) SetRequestId(v string) *ListSmsSignResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSmsSignResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSmsSignResponseBodyData struct {
	AuditResult *string `json:"AuditResult,omitempty" xml:"AuditResult,omitempty"`
	CreateDate  *int64  `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	DefaultFlag *bool   `json:"DefaultFlag,omitempty" xml:"DefaultFlag,omitempty"`
	SmsSignName *string `json:"SmsSignName,omitempty" xml:"SmsSignName,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TestFlag    *bool   `json:"TestFlag,omitempty" xml:"TestFlag,omitempty"`
}

func (s ListSmsSignResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListSmsSignResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSmsSignResponseBodyData) GetAuditResult() *string {
	return s.AuditResult
}

func (s *ListSmsSignResponseBodyData) GetCreateDate() *int64 {
	return s.CreateDate
}

func (s *ListSmsSignResponseBodyData) GetDefaultFlag() *bool {
	return s.DefaultFlag
}

func (s *ListSmsSignResponseBodyData) GetSmsSignName() *string {
	return s.SmsSignName
}

func (s *ListSmsSignResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ListSmsSignResponseBodyData) GetTestFlag() *bool {
	return s.TestFlag
}

func (s *ListSmsSignResponseBodyData) SetAuditResult(v string) *ListSmsSignResponseBodyData {
	s.AuditResult = &v
	return s
}

func (s *ListSmsSignResponseBodyData) SetCreateDate(v int64) *ListSmsSignResponseBodyData {
	s.CreateDate = &v
	return s
}

func (s *ListSmsSignResponseBodyData) SetDefaultFlag(v bool) *ListSmsSignResponseBodyData {
	s.DefaultFlag = &v
	return s
}

func (s *ListSmsSignResponseBodyData) SetSmsSignName(v string) *ListSmsSignResponseBodyData {
	s.SmsSignName = &v
	return s
}

func (s *ListSmsSignResponseBodyData) SetStatus(v string) *ListSmsSignResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListSmsSignResponseBodyData) SetTestFlag(v bool) *ListSmsSignResponseBodyData {
	s.TestFlag = &v
	return s
}

func (s *ListSmsSignResponseBodyData) Validate() error {
	return dara.Validate(s)
}
