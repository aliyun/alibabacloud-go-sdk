// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelPublishTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBizTypeList(v []*string) *CancelPublishTaskResponseBody
	GetBizTypeList() []*string
	SetCreateTime(v string) *CancelPublishTaskResponseBody
	GetCreateTime() *string
	SetError(v string) *CancelPublishTaskResponseBody
	GetError() *string
	SetId(v int64) *CancelPublishTaskResponseBody
	GetId() *int64
	SetModifyTime(v string) *CancelPublishTaskResponseBody
	GetModifyTime() *string
	SetRequestId(v string) *CancelPublishTaskResponseBody
	GetRequestId() *string
	SetResponse(v string) *CancelPublishTaskResponseBody
	GetResponse() *string
	SetStatus(v string) *CancelPublishTaskResponseBody
	GetStatus() *string
}

type CancelPublishTaskResponseBody struct {
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
	// FE_ABORTED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CancelPublishTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelPublishTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CancelPublishTaskResponseBody) GetBizTypeList() []*string {
	return s.BizTypeList
}

func (s *CancelPublishTaskResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CancelPublishTaskResponseBody) GetError() *string {
	return s.Error
}

func (s *CancelPublishTaskResponseBody) GetId() *int64 {
	return s.Id
}

func (s *CancelPublishTaskResponseBody) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *CancelPublishTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelPublishTaskResponseBody) GetResponse() *string {
	return s.Response
}

func (s *CancelPublishTaskResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CancelPublishTaskResponseBody) SetBizTypeList(v []*string) *CancelPublishTaskResponseBody {
	s.BizTypeList = v
	return s
}

func (s *CancelPublishTaskResponseBody) SetCreateTime(v string) *CancelPublishTaskResponseBody {
	s.CreateTime = &v
	return s
}

func (s *CancelPublishTaskResponseBody) SetError(v string) *CancelPublishTaskResponseBody {
	s.Error = &v
	return s
}

func (s *CancelPublishTaskResponseBody) SetId(v int64) *CancelPublishTaskResponseBody {
	s.Id = &v
	return s
}

func (s *CancelPublishTaskResponseBody) SetModifyTime(v string) *CancelPublishTaskResponseBody {
	s.ModifyTime = &v
	return s
}

func (s *CancelPublishTaskResponseBody) SetRequestId(v string) *CancelPublishTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelPublishTaskResponseBody) SetResponse(v string) *CancelPublishTaskResponseBody {
	s.Response = &v
	return s
}

func (s *CancelPublishTaskResponseBody) SetStatus(v string) *CancelPublishTaskResponseBody {
	s.Status = &v
	return s
}

func (s *CancelPublishTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
