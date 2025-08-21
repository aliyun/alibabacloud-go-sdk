// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstancePublishTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBizTypeList(v []*string) *CreateInstancePublishTaskResponseBody
	GetBizTypeList() []*string
	SetCreateTime(v string) *CreateInstancePublishTaskResponseBody
	GetCreateTime() *string
	SetError(v string) *CreateInstancePublishTaskResponseBody
	GetError() *string
	SetId(v int64) *CreateInstancePublishTaskResponseBody
	GetId() *int64
	SetModifyTime(v string) *CreateInstancePublishTaskResponseBody
	GetModifyTime() *string
	SetRequestId(v string) *CreateInstancePublishTaskResponseBody
	GetRequestId() *string
	SetResponse(v string) *CreateInstancePublishTaskResponseBody
	GetResponse() *string
	SetStatus(v string) *CreateInstancePublishTaskResponseBody
	GetStatus() *string
}

type CreateInstancePublishTaskResponseBody struct {
	BizTypeList []*string `json:"BizTypeList,omitempty" xml:"BizTypeList,omitempty" type:"Repeated"`
	// example:
	//
	// 2022-04-12T06:30:17Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Error      *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// example:
	//
	// 8522
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 2022-04-12T06:30:33Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// 5CBF0581-EAE7-1DC4-95C6-A089656A1E2D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 8522
	Response *string `json:"Response,omitempty" xml:"Response,omitempty"`
	// example:
	//
	// FE_RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateInstancePublishTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateInstancePublishTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstancePublishTaskResponseBody) GetBizTypeList() []*string {
	return s.BizTypeList
}

func (s *CreateInstancePublishTaskResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CreateInstancePublishTaskResponseBody) GetError() *string {
	return s.Error
}

func (s *CreateInstancePublishTaskResponseBody) GetId() *int64 {
	return s.Id
}

func (s *CreateInstancePublishTaskResponseBody) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *CreateInstancePublishTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateInstancePublishTaskResponseBody) GetResponse() *string {
	return s.Response
}

func (s *CreateInstancePublishTaskResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateInstancePublishTaskResponseBody) SetBizTypeList(v []*string) *CreateInstancePublishTaskResponseBody {
	s.BizTypeList = v
	return s
}

func (s *CreateInstancePublishTaskResponseBody) SetCreateTime(v string) *CreateInstancePublishTaskResponseBody {
	s.CreateTime = &v
	return s
}

func (s *CreateInstancePublishTaskResponseBody) SetError(v string) *CreateInstancePublishTaskResponseBody {
	s.Error = &v
	return s
}

func (s *CreateInstancePublishTaskResponseBody) SetId(v int64) *CreateInstancePublishTaskResponseBody {
	s.Id = &v
	return s
}

func (s *CreateInstancePublishTaskResponseBody) SetModifyTime(v string) *CreateInstancePublishTaskResponseBody {
	s.ModifyTime = &v
	return s
}

func (s *CreateInstancePublishTaskResponseBody) SetRequestId(v string) *CreateInstancePublishTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInstancePublishTaskResponseBody) SetResponse(v string) *CreateInstancePublishTaskResponseBody {
	s.Response = &v
	return s
}

func (s *CreateInstancePublishTaskResponseBody) SetStatus(v string) *CreateInstancePublishTaskResponseBody {
	s.Status = &v
	return s
}

func (s *CreateInstancePublishTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
