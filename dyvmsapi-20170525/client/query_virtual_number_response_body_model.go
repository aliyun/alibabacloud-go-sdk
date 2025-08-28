// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryVirtualNumberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryVirtualNumberResponseBody
	GetCode() *string
	SetData(v string) *QueryVirtualNumberResponseBody
	GetData() *string
	SetRequestId(v string) *QueryVirtualNumberResponseBody
	GetRequestId() *string
}

type QueryVirtualNumberResponseBody struct {
	// The response code. The value 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the numbers associated with the virtual numbers. The following fields are returned:
	//
	// 	- createTime: the time when the number was activated.
	//
	// 	- qualificationCount: the number of qualifications.
	//
	// 	- cityCount: the number of cities.
	//
	// 	- phoneNumCount: the number of virtual numbers.
	//
	// 	- remark: the additional information.
	//
	// 	- phoneNum: the virtual number.
	//
	// 	- routeType: the route type.
	//
	// 	- canCancel: indicates whether the number can be deactivated.
	//
	// 	- specCount: the number of Internet service providers (ISPs).
	//
	// 	- status: the number state. Valid values: **1**, **0**, and **-1**. The value 1 indicates that the number is valid. The value 0 indicates that the number is invalid. The value -1 indicates that the number was deactivated.
	//
	// 	- pageNo: the page number.
	//
	// 	- pageSize: the number of entries per page.
	//
	// 	- total: the total number of virtual numbers.
	//
	// example:
	//
	// {"data":[{"createTime":"2020-07-15 04:00:00","qualificationCount":0,"cityCount":0,"phoneNumCount":1,"remark":"20200715Unicom CTD shut down","phoneNum":"05516214****","routeType":1,"canCancel":true,"specCount":0,"status":"1"}],"pageSize":1,"total":17,"pageNo":1}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9FF70B74-1B3C-44C1-ACDF-8DF962988F0E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryVirtualNumberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryVirtualNumberResponseBody) GoString() string {
	return s.String()
}

func (s *QueryVirtualNumberResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryVirtualNumberResponseBody) GetData() *string {
	return s.Data
}

func (s *QueryVirtualNumberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryVirtualNumberResponseBody) SetCode(v string) *QueryVirtualNumberResponseBody {
	s.Code = &v
	return s
}

func (s *QueryVirtualNumberResponseBody) SetData(v string) *QueryVirtualNumberResponseBody {
	s.Data = &v
	return s
}

func (s *QueryVirtualNumberResponseBody) SetRequestId(v string) *QueryVirtualNumberResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryVirtualNumberResponseBody) Validate() error {
	return dara.Validate(s)
}
