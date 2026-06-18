// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBillDetailFileListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetBillDetailFileListResponseBody
	GetCode() *string
	SetData(v []*GetBillDetailFileListResponseBodyData) *GetBillDetailFileListResponseBody
	GetData() []*GetBillDetailFileListResponseBodyData
	SetMessage(v string) *GetBillDetailFileListResponseBody
	GetMessage() *string
	SetMsg(v string) *GetBillDetailFileListResponseBody
	GetMsg() *string
	SetRequestId(v string) *GetBillDetailFileListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetBillDetailFileListResponseBody
	GetSuccess() *bool
}

type GetBillDetailFileListResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data []*GetBillDetailFileListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The message.
	//
	// example:
	//
	// 成功
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Same as Message.
	//
	// example:
	//
	// 成功
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 210bc4b416874189683843905d9f9a
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetBillDetailFileListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBillDetailFileListResponseBody) GoString() string {
	return s.String()
}

func (s *GetBillDetailFileListResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetBillDetailFileListResponseBody) GetData() []*GetBillDetailFileListResponseBodyData {
	return s.Data
}

func (s *GetBillDetailFileListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetBillDetailFileListResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *GetBillDetailFileListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetBillDetailFileListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetBillDetailFileListResponseBody) SetCode(v string) *GetBillDetailFileListResponseBody {
	s.Code = &v
	return s
}

func (s *GetBillDetailFileListResponseBody) SetData(v []*GetBillDetailFileListResponseBodyData) *GetBillDetailFileListResponseBody {
	s.Data = v
	return s
}

func (s *GetBillDetailFileListResponseBody) SetMessage(v string) *GetBillDetailFileListResponseBody {
	s.Message = &v
	return s
}

func (s *GetBillDetailFileListResponseBody) SetMsg(v string) *GetBillDetailFileListResponseBody {
	s.Msg = &v
	return s
}

func (s *GetBillDetailFileListResponseBody) SetRequestId(v string) *GetBillDetailFileListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBillDetailFileListResponseBody) SetSuccess(v bool) *GetBillDetailFileListResponseBody {
	s.Success = &v
	return s
}

func (s *GetBillDetailFileListResponseBody) Validate() error {
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

type GetBillDetailFileListResponseBodyData struct {
	// The billing month.
	//
	// example:
	//
	// 202502
	BillMonth *string `json:"BillMonth,omitempty" xml:"BillMonth,omitempty"`
	// The file name.
	//
	// example:
	//
	// 账单202502021112
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The file URL.
	//
	// example:
	//
	// aps.ailyun.com/file/download?resourceId=1234&type=1
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// The OSS file push status. Valid values:
	//
	// - 1: pending
	//
	// - 2: processing
	//
	// - 3: completed.
	//
	// example:
	//
	// 3
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type. Valid values: customer acquisition or channel expansion.
	//
	// example:
	//
	// 拓渠
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetBillDetailFileListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetBillDetailFileListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetBillDetailFileListResponseBodyData) GetBillMonth() *string {
	return s.BillMonth
}

func (s *GetBillDetailFileListResponseBodyData) GetFileName() *string {
	return s.FileName
}

func (s *GetBillDetailFileListResponseBodyData) GetFileUrl() *string {
	return s.FileUrl
}

func (s *GetBillDetailFileListResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetBillDetailFileListResponseBodyData) GetType() *string {
	return s.Type
}

func (s *GetBillDetailFileListResponseBodyData) SetBillMonth(v string) *GetBillDetailFileListResponseBodyData {
	s.BillMonth = &v
	return s
}

func (s *GetBillDetailFileListResponseBodyData) SetFileName(v string) *GetBillDetailFileListResponseBodyData {
	s.FileName = &v
	return s
}

func (s *GetBillDetailFileListResponseBodyData) SetFileUrl(v string) *GetBillDetailFileListResponseBodyData {
	s.FileUrl = &v
	return s
}

func (s *GetBillDetailFileListResponseBodyData) SetStatus(v string) *GetBillDetailFileListResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetBillDetailFileListResponseBodyData) SetType(v string) *GetBillDetailFileListResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetBillDetailFileListResponseBodyData) Validate() error {
	return dara.Validate(s)
}
