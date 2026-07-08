// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTaskListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryTaskListResponseBody
	GetCode() *string
	SetData(v *QueryTaskListResponseBodyData) *QueryTaskListResponseBody
	GetData() *QueryTaskListResponseBodyData
	SetMessage(v string) *QueryTaskListResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryTaskListResponseBody
	GetRequestId() *string
}

type QueryTaskListResponseBody struct {
	// The request status code. Valid values:
	//
	// - **OK**: The request was successful.
	//
	// - **isv.INVALID_PARAMETERS**: The specified parameter is invalid.
	//
	// - **isp.SYSTEM_ERROR**: A system error occurred.
	//
	// > For more information, see the "Error codes" section of this topic.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	Data *QueryTaskListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The status code description.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 68A40250-50CD-034C-B728-0BD******177
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryTaskListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryTaskListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTaskListResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryTaskListResponseBody) GetData() *QueryTaskListResponseBodyData {
	return s.Data
}

func (s *QueryTaskListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryTaskListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryTaskListResponseBody) SetCode(v string) *QueryTaskListResponseBody {
	s.Code = &v
	return s
}

func (s *QueryTaskListResponseBody) SetData(v *QueryTaskListResponseBodyData) *QueryTaskListResponseBody {
	s.Data = v
	return s
}

func (s *QueryTaskListResponseBody) SetMessage(v string) *QueryTaskListResponseBody {
	s.Message = &v
	return s
}

func (s *QueryTaskListResponseBody) SetRequestId(v string) *QueryTaskListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTaskListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryTaskListResponseBodyData struct {
	// The details of the tasks.
	Data []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The current page number.
	//
	// example:
	//
	// 62
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The total count.
	//
	// example:
	//
	// 22
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryTaskListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryTaskListResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryTaskListResponseBodyData) GetData() []*string {
	return s.Data
}

func (s *QueryTaskListResponseBodyData) GetPageNo() *int64 {
	return s.PageNo
}

func (s *QueryTaskListResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *QueryTaskListResponseBodyData) SetData(v []*string) *QueryTaskListResponseBodyData {
	s.Data = v
	return s
}

func (s *QueryTaskListResponseBodyData) SetPageNo(v int64) *QueryTaskListResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *QueryTaskListResponseBodyData) SetTotalCount(v int64) *QueryTaskListResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *QueryTaskListResponseBodyData) Validate() error {
	return dara.Validate(s)
}
