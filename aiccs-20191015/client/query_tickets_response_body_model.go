// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTicketsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryTicketsResponseBody
	GetCode() *string
	SetData(v string) *QueryTicketsResponseBody
	GetData() *string
	SetMessage(v string) *QueryTicketsResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryTicketsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryTicketsResponseBody
	GetSuccess() *bool
}

type QueryTicketsResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// { "totalResults":1 "previousPage":1 "data":[ 0:{ "serviceType":1 "lastUrgeTime":0 "queueId":0 "sopCateId":252011 "totalUrgeMemo":"" "taskGmtModified":0 "departmentId":10 "groupId":0 "channelType":1 "questionInfo":"" "templateId":0 "deadLine":0 "srType":30701 "caseId":2000000001165962}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// EE338D98-9BD3-4413-B165
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryTicketsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryTicketsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTicketsResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryTicketsResponseBody) GetData() *string {
	return s.Data
}

func (s *QueryTicketsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryTicketsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryTicketsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryTicketsResponseBody) SetCode(v string) *QueryTicketsResponseBody {
	s.Code = &v
	return s
}

func (s *QueryTicketsResponseBody) SetData(v string) *QueryTicketsResponseBody {
	s.Data = &v
	return s
}

func (s *QueryTicketsResponseBody) SetMessage(v string) *QueryTicketsResponseBody {
	s.Message = &v
	return s
}

func (s *QueryTicketsResponseBody) SetRequestId(v string) *QueryTicketsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTicketsResponseBody) SetSuccess(v bool) *QueryTicketsResponseBody {
	s.Success = &v
	return s
}

func (s *QueryTicketsResponseBody) Validate() error {
	return dara.Validate(s)
}
