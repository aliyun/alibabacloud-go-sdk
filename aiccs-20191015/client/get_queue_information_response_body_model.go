// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQueueInformationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetQueueInformationResponseBody
	GetCode() *string
	SetData(v *GetQueueInformationResponseBodyData) *GetQueueInformationResponseBody
	GetData() *GetQueueInformationResponseBodyData
	SetMessage(v string) *GetQueueInformationResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetQueueInformationResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetQueueInformationResponseBody
	GetSuccess() *string
}

type GetQueueInformationResponseBody struct {
	// example:
	//
	// 200
	Code *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetQueueInformationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// EE338D98-9BD3-4413-B165
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetQueueInformationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetQueueInformationResponseBody) GoString() string {
	return s.String()
}

func (s *GetQueueInformationResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetQueueInformationResponseBody) GetData() *GetQueueInformationResponseBodyData {
	return s.Data
}

func (s *GetQueueInformationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetQueueInformationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetQueueInformationResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetQueueInformationResponseBody) SetCode(v string) *GetQueueInformationResponseBody {
	s.Code = &v
	return s
}

func (s *GetQueueInformationResponseBody) SetData(v *GetQueueInformationResponseBodyData) *GetQueueInformationResponseBody {
	s.Data = v
	return s
}

func (s *GetQueueInformationResponseBody) SetMessage(v string) *GetQueueInformationResponseBody {
	s.Message = &v
	return s
}

func (s *GetQueueInformationResponseBody) SetRequestId(v string) *GetQueueInformationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQueueInformationResponseBody) SetSuccess(v string) *GetQueueInformationResponseBody {
	s.Success = &v
	return s
}

func (s *GetQueueInformationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetQueueInformationResponseBodyData struct {
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 2000
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Rows     *string `json:"Rows,omitempty" xml:"Rows,omitempty"`
	// example:
	//
	// 4
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s GetQueueInformationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetQueueInformationResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetQueueInformationResponseBodyData) GetPageNum() *int32 {
	return s.PageNum
}

func (s *GetQueueInformationResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetQueueInformationResponseBodyData) GetRows() *string {
	return s.Rows
}

func (s *GetQueueInformationResponseBodyData) GetTotalNum() *int32 {
	return s.TotalNum
}

func (s *GetQueueInformationResponseBodyData) SetPageNum(v int32) *GetQueueInformationResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *GetQueueInformationResponseBodyData) SetPageSize(v int32) *GetQueueInformationResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetQueueInformationResponseBodyData) SetRows(v string) *GetQueueInformationResponseBodyData {
	s.Rows = &v
	return s
}

func (s *GetQueueInformationResponseBodyData) SetTotalNum(v int32) *GetQueueInformationResponseBodyData {
	s.TotalNum = &v
	return s
}

func (s *GetQueueInformationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
