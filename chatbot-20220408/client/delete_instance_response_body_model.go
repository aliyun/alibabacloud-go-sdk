// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBizTypeList(v []*string) *DeleteInstanceResponseBody
	GetBizTypeList() []*string
	SetCreateTime(v string) *DeleteInstanceResponseBody
	GetCreateTime() *string
	SetCreateUserId(v int64) *DeleteInstanceResponseBody
	GetCreateUserId() *int64
	SetCreateUserName(v string) *DeleteInstanceResponseBody
	GetCreateUserName() *string
	SetError(v string) *DeleteInstanceResponseBody
	GetError() *string
	SetId(v int64) *DeleteInstanceResponseBody
	GetId() *int64
	SetRequestId(v string) *DeleteInstanceResponseBody
	GetRequestId() *string
	SetResponse(v int64) *DeleteInstanceResponseBody
	GetResponse() *int64
	SetStatus(v string) *DeleteInstanceResponseBody
	GetStatus() *string
}

type DeleteInstanceResponseBody struct {
	BizTypeList []*string `json:"BizTypeList,omitempty" xml:"BizTypeList,omitempty" type:"Repeated"`
	// example:
	//
	// 2021-09-11T09:26:14Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 9052
	CreateUserId *int64 `json:"CreateUserId,omitempty" xml:"CreateUserId,omitempty"`
	// example:
	//
	// xuqiang_test
	CreateUserName *string `json:"CreateUserName,omitempty" xml:"CreateUserName,omitempty"`
	Error          *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// example:
	//
	// 8521
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 5CBF0581-EAE7-1DC4-95C6-A089656A1E2D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 8521
	Response *int64 `json:"Response,omitempty" xml:"Response,omitempty"`
	// example:
	//
	// FE_RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DeleteInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInstanceResponseBody) GetBizTypeList() []*string {
	return s.BizTypeList
}

func (s *DeleteInstanceResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DeleteInstanceResponseBody) GetCreateUserId() *int64 {
	return s.CreateUserId
}

func (s *DeleteInstanceResponseBody) GetCreateUserName() *string {
	return s.CreateUserName
}

func (s *DeleteInstanceResponseBody) GetError() *string {
	return s.Error
}

func (s *DeleteInstanceResponseBody) GetId() *int64 {
	return s.Id
}

func (s *DeleteInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteInstanceResponseBody) GetResponse() *int64 {
	return s.Response
}

func (s *DeleteInstanceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DeleteInstanceResponseBody) SetBizTypeList(v []*string) *DeleteInstanceResponseBody {
	s.BizTypeList = v
	return s
}

func (s *DeleteInstanceResponseBody) SetCreateTime(v string) *DeleteInstanceResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DeleteInstanceResponseBody) SetCreateUserId(v int64) *DeleteInstanceResponseBody {
	s.CreateUserId = &v
	return s
}

func (s *DeleteInstanceResponseBody) SetCreateUserName(v string) *DeleteInstanceResponseBody {
	s.CreateUserName = &v
	return s
}

func (s *DeleteInstanceResponseBody) SetError(v string) *DeleteInstanceResponseBody {
	s.Error = &v
	return s
}

func (s *DeleteInstanceResponseBody) SetId(v int64) *DeleteInstanceResponseBody {
	s.Id = &v
	return s
}

func (s *DeleteInstanceResponseBody) SetRequestId(v string) *DeleteInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteInstanceResponseBody) SetResponse(v int64) *DeleteInstanceResponseBody {
	s.Response = &v
	return s
}

func (s *DeleteInstanceResponseBody) SetStatus(v string) *DeleteInstanceResponseBody {
	s.Status = &v
	return s
}

func (s *DeleteInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
