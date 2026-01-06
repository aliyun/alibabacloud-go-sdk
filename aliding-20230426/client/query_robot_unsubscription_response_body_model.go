// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRobotUnsubscriptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNo(v int32) *QueryRobotUnsubscriptionResponseBody
	GetPageNo() *int32
	SetPageSize(v int32) *QueryRobotUnsubscriptionResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *QueryRobotUnsubscriptionResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *QueryRobotUnsubscriptionResponseBody
	GetTotalCount() *int32
	SetUnsubscribedStaffIds(v []*string) *QueryRobotUnsubscriptionResponseBody
	GetUnsubscribedStaffIds() []*string
}

type QueryRobotUnsubscriptionResponseBody struct {
	// example:
	//
	// 1
	PageNo *int32 `json:"pageNo,omitempty" xml:"pageNo,omitempty"`
	// example:
	//
	// 50
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// requestId
	//
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 500
	TotalCount           *int32    `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	UnsubscribedStaffIds []*string `json:"unsubscribedStaffIds,omitempty" xml:"unsubscribedStaffIds,omitempty" type:"Repeated"`
}

func (s QueryRobotUnsubscriptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryRobotUnsubscriptionResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRobotUnsubscriptionResponseBody) GetPageNo() *int32 {
	return s.PageNo
}

func (s *QueryRobotUnsubscriptionResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryRobotUnsubscriptionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryRobotUnsubscriptionResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *QueryRobotUnsubscriptionResponseBody) GetUnsubscribedStaffIds() []*string {
	return s.UnsubscribedStaffIds
}

func (s *QueryRobotUnsubscriptionResponseBody) SetPageNo(v int32) *QueryRobotUnsubscriptionResponseBody {
	s.PageNo = &v
	return s
}

func (s *QueryRobotUnsubscriptionResponseBody) SetPageSize(v int32) *QueryRobotUnsubscriptionResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryRobotUnsubscriptionResponseBody) SetRequestId(v string) *QueryRobotUnsubscriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRobotUnsubscriptionResponseBody) SetTotalCount(v int32) *QueryRobotUnsubscriptionResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryRobotUnsubscriptionResponseBody) SetUnsubscribedStaffIds(v []*string) *QueryRobotUnsubscriptionResponseBody {
	s.UnsubscribedStaffIds = v
	return s
}

func (s *QueryRobotUnsubscriptionResponseBody) Validate() error {
	return dara.Validate(s)
}
