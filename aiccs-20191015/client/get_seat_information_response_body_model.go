// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSeatInformationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetSeatInformationResponseBody
	GetCode() *string
	SetData(v *GetSeatInformationResponseBodyData) *GetSeatInformationResponseBody
	GetData() *GetSeatInformationResponseBodyData
	SetMessage(v string) *GetSeatInformationResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetSeatInformationResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetSeatInformationResponseBody
	GetSuccess() *string
}

type GetSeatInformationResponseBody struct {
	// example:
	//
	// 200
	Code *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetSeatInformationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s GetSeatInformationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSeatInformationResponseBody) GoString() string {
	return s.String()
}

func (s *GetSeatInformationResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetSeatInformationResponseBody) GetData() *GetSeatInformationResponseBodyData {
	return s.Data
}

func (s *GetSeatInformationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSeatInformationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSeatInformationResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetSeatInformationResponseBody) SetCode(v string) *GetSeatInformationResponseBody {
	s.Code = &v
	return s
}

func (s *GetSeatInformationResponseBody) SetData(v *GetSeatInformationResponseBodyData) *GetSeatInformationResponseBody {
	s.Data = v
	return s
}

func (s *GetSeatInformationResponseBody) SetMessage(v string) *GetSeatInformationResponseBody {
	s.Message = &v
	return s
}

func (s *GetSeatInformationResponseBody) SetRequestId(v string) *GetSeatInformationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSeatInformationResponseBody) SetSuccess(v string) *GetSeatInformationResponseBody {
	s.Success = &v
	return s
}

func (s *GetSeatInformationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSeatInformationResponseBodyData struct {
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 2000
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Rowr     *string `json:"Rowr,omitempty" xml:"Rowr,omitempty"`
	// example:
	//
	// 4
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s GetSeatInformationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSeatInformationResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSeatInformationResponseBodyData) GetPageNum() *int32 {
	return s.PageNum
}

func (s *GetSeatInformationResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetSeatInformationResponseBodyData) GetRowr() *string {
	return s.Rowr
}

func (s *GetSeatInformationResponseBodyData) GetTotalNum() *int32 {
	return s.TotalNum
}

func (s *GetSeatInformationResponseBodyData) SetPageNum(v int32) *GetSeatInformationResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *GetSeatInformationResponseBodyData) SetPageSize(v int32) *GetSeatInformationResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetSeatInformationResponseBodyData) SetRowr(v string) *GetSeatInformationResponseBodyData {
	s.Rowr = &v
	return s
}

func (s *GetSeatInformationResponseBodyData) SetTotalNum(v int32) *GetSeatInformationResponseBodyData {
	s.TotalNum = &v
	return s
}

func (s *GetSeatInformationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
