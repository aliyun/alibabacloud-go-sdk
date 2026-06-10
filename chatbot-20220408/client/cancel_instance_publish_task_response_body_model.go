// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelInstancePublishTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBizTypeList(v []*string) *CancelInstancePublishTaskResponseBody
	GetBizTypeList() []*string
	SetCreateTime(v string) *CancelInstancePublishTaskResponseBody
	GetCreateTime() *string
	SetError(v string) *CancelInstancePublishTaskResponseBody
	GetError() *string
	SetId(v int64) *CancelInstancePublishTaskResponseBody
	GetId() *int64
	SetModifyTime(v string) *CancelInstancePublishTaskResponseBody
	GetModifyTime() *string
	SetRequestId(v string) *CancelInstancePublishTaskResponseBody
	GetRequestId() *string
	SetResponse(v string) *CancelInstancePublishTaskResponseBody
	GetResponse() *string
	SetStatus(v string) *CancelInstancePublishTaskResponseBody
	GetStatus() *string
}

type CancelInstancePublishTaskResponseBody struct {
	// A list of business types.
	BizTypeList []*string `json:"BizTypeList,omitempty" xml:"BizTypeList,omitempty" type:"Repeated"`
	// The creation time of the task in UTC.
	//
	// example:
	//
	// 2022-04-12T06:30:17Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The error message if the task fails.
	//
	// example:
	//
	// 检查待发布模块是否空闲发生错误，faq
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 8522
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The last modification time of the task in UTC.
	//
	// example:
	//
	// 2022-04-12T06:30:33Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 5CBF0581-EAE7-1DC4-95C6-A089656A1E2D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task ID, returned as a string.
	//
	// example:
	//
	// 8522
	Response *string `json:"Response,omitempty" xml:"Response,omitempty"`
	// The task status.
	//
	// example:
	//
	// FE_ABORTED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CancelInstancePublishTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelInstancePublishTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CancelInstancePublishTaskResponseBody) GetBizTypeList() []*string {
	return s.BizTypeList
}

func (s *CancelInstancePublishTaskResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CancelInstancePublishTaskResponseBody) GetError() *string {
	return s.Error
}

func (s *CancelInstancePublishTaskResponseBody) GetId() *int64 {
	return s.Id
}

func (s *CancelInstancePublishTaskResponseBody) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *CancelInstancePublishTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelInstancePublishTaskResponseBody) GetResponse() *string {
	return s.Response
}

func (s *CancelInstancePublishTaskResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CancelInstancePublishTaskResponseBody) SetBizTypeList(v []*string) *CancelInstancePublishTaskResponseBody {
	s.BizTypeList = v
	return s
}

func (s *CancelInstancePublishTaskResponseBody) SetCreateTime(v string) *CancelInstancePublishTaskResponseBody {
	s.CreateTime = &v
	return s
}

func (s *CancelInstancePublishTaskResponseBody) SetError(v string) *CancelInstancePublishTaskResponseBody {
	s.Error = &v
	return s
}

func (s *CancelInstancePublishTaskResponseBody) SetId(v int64) *CancelInstancePublishTaskResponseBody {
	s.Id = &v
	return s
}

func (s *CancelInstancePublishTaskResponseBody) SetModifyTime(v string) *CancelInstancePublishTaskResponseBody {
	s.ModifyTime = &v
	return s
}

func (s *CancelInstancePublishTaskResponseBody) SetRequestId(v string) *CancelInstancePublishTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelInstancePublishTaskResponseBody) SetResponse(v string) *CancelInstancePublishTaskResponseBody {
	s.Response = &v
	return s
}

func (s *CancelInstancePublishTaskResponseBody) SetStatus(v string) *CancelInstancePublishTaskResponseBody {
	s.Status = &v
	return s
}

func (s *CancelInstancePublishTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
