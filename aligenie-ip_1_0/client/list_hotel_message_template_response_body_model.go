// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotelMessageTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListHotelMessageTemplateResponseBody
	GetCode() *int32
	SetMessage(v string) *ListHotelMessageTemplateResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListHotelMessageTemplateResponseBody
	GetRequestId() *string
	SetResult(v []*ListHotelMessageTemplateResponseBodyResult) *ListHotelMessageTemplateResponseBody
	GetResult() []*ListHotelMessageTemplateResponseBodyResult
}

type ListHotelMessageTemplateResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message   *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListHotelMessageTemplateResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListHotelMessageTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListHotelMessageTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ListHotelMessageTemplateResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListHotelMessageTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListHotelMessageTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListHotelMessageTemplateResponseBody) GetResult() []*ListHotelMessageTemplateResponseBodyResult {
	return s.Result
}

func (s *ListHotelMessageTemplateResponseBody) SetCode(v int32) *ListHotelMessageTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *ListHotelMessageTemplateResponseBody) SetMessage(v string) *ListHotelMessageTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *ListHotelMessageTemplateResponseBody) SetRequestId(v string) *ListHotelMessageTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHotelMessageTemplateResponseBody) SetResult(v []*ListHotelMessageTemplateResponseBodyResult) *ListHotelMessageTemplateResponseBody {
	s.Result = v
	return s
}

func (s *ListHotelMessageTemplateResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListHotelMessageTemplateResponseBodyResult struct {
	// example:
	//
	// 不通过
	AuditMark *string `json:"AuditMark,omitempty" xml:"AuditMark,omitempty"`
	// example:
	//
	// COMMIT
	AuditStatus *string `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	// example:
	//
	// 这是${hotel}的一个测试模板
	TemplateDetail *string `json:"TemplateDetail,omitempty" xml:"TemplateDetail,omitempty"`
	// example:
	//
	// 1
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// example:
	//
	// 测试模板
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s ListHotelMessageTemplateResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListHotelMessageTemplateResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListHotelMessageTemplateResponseBodyResult) GetAuditMark() *string {
	return s.AuditMark
}

func (s *ListHotelMessageTemplateResponseBodyResult) GetAuditStatus() *string {
	return s.AuditStatus
}

func (s *ListHotelMessageTemplateResponseBodyResult) GetTemplateDetail() *string {
	return s.TemplateDetail
}

func (s *ListHotelMessageTemplateResponseBodyResult) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *ListHotelMessageTemplateResponseBodyResult) GetTemplateName() *string {
	return s.TemplateName
}

func (s *ListHotelMessageTemplateResponseBodyResult) SetAuditMark(v string) *ListHotelMessageTemplateResponseBodyResult {
	s.AuditMark = &v
	return s
}

func (s *ListHotelMessageTemplateResponseBodyResult) SetAuditStatus(v string) *ListHotelMessageTemplateResponseBodyResult {
	s.AuditStatus = &v
	return s
}

func (s *ListHotelMessageTemplateResponseBodyResult) SetTemplateDetail(v string) *ListHotelMessageTemplateResponseBodyResult {
	s.TemplateDetail = &v
	return s
}

func (s *ListHotelMessageTemplateResponseBodyResult) SetTemplateId(v int64) *ListHotelMessageTemplateResponseBodyResult {
	s.TemplateId = &v
	return s
}

func (s *ListHotelMessageTemplateResponseBodyResult) SetTemplateName(v string) *ListHotelMessageTemplateResponseBodyResult {
	s.TemplateName = &v
	return s
}

func (s *ListHotelMessageTemplateResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
