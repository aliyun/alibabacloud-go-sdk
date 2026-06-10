// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePublishTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBizTypeList(v []*string) *CreatePublishTaskResponseBody
	GetBizTypeList() []*string
	SetCreateTime(v string) *CreatePublishTaskResponseBody
	GetCreateTime() *string
	SetError(v string) *CreatePublishTaskResponseBody
	GetError() *string
	SetId(v int64) *CreatePublishTaskResponseBody
	GetId() *int64
	SetModifyTime(v string) *CreatePublishTaskResponseBody
	GetModifyTime() *string
	SetRequestId(v string) *CreatePublishTaskResponseBody
	GetRequestId() *string
	SetResponse(v string) *CreatePublishTaskResponseBody
	GetResponse() *string
	SetStatus(v string) *CreatePublishTaskResponseBody
	GetStatus() *string
}

type CreatePublishTaskResponseBody struct {
	// The list of business types.
	BizTypeList []*string `json:"BizTypeList,omitempty" xml:"BizTypeList,omitempty" type:"Repeated"`
	// The time when the task was created, in UTC.
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
	// The publish task ID.
	//
	// example:
	//
	// 8522
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The time when the task was last modified, in UTC.
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
	// The publish task ID. This parameter is redundant. Use the `Id` parameter instead.
	//
	// example:
	//
	// 8522
	Response *string `json:"Response,omitempty" xml:"Response,omitempty"`
	// The task status. Valid values:
	//
	// `FE_RUNNING`: Publishing<br><br>`FE_SUCCESS`: Published<br><br>`FE_FAILED`: Failed<br><br>`FE_ABORTED`: Canceled<br><br><br><br><br><br>
	//
	// example:
	//
	// FE_RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreatePublishTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePublishTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePublishTaskResponseBody) GetBizTypeList() []*string {
	return s.BizTypeList
}

func (s *CreatePublishTaskResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CreatePublishTaskResponseBody) GetError() *string {
	return s.Error
}

func (s *CreatePublishTaskResponseBody) GetId() *int64 {
	return s.Id
}

func (s *CreatePublishTaskResponseBody) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *CreatePublishTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePublishTaskResponseBody) GetResponse() *string {
	return s.Response
}

func (s *CreatePublishTaskResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreatePublishTaskResponseBody) SetBizTypeList(v []*string) *CreatePublishTaskResponseBody {
	s.BizTypeList = v
	return s
}

func (s *CreatePublishTaskResponseBody) SetCreateTime(v string) *CreatePublishTaskResponseBody {
	s.CreateTime = &v
	return s
}

func (s *CreatePublishTaskResponseBody) SetError(v string) *CreatePublishTaskResponseBody {
	s.Error = &v
	return s
}

func (s *CreatePublishTaskResponseBody) SetId(v int64) *CreatePublishTaskResponseBody {
	s.Id = &v
	return s
}

func (s *CreatePublishTaskResponseBody) SetModifyTime(v string) *CreatePublishTaskResponseBody {
	s.ModifyTime = &v
	return s
}

func (s *CreatePublishTaskResponseBody) SetRequestId(v string) *CreatePublishTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePublishTaskResponseBody) SetResponse(v string) *CreatePublishTaskResponseBody {
	s.Response = &v
	return s
}

func (s *CreatePublishTaskResponseBody) SetStatus(v string) *CreatePublishTaskResponseBody {
	s.Status = &v
	return s
}

func (s *CreatePublishTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
