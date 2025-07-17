// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDashboardsByNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListDashboardsByNameResponseBody
	GetCode() *int32
	SetData(v string) *ListDashboardsByNameResponseBody
	GetData() *string
	SetMessage(v string) *ListDashboardsByNameResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListDashboardsByNameResponseBody
	GetRequestId() *string
}

type ListDashboardsByNameResponseBody struct {
	// The status code. The HTTP 200 status code indicates a successful request, while others indicate error conditions.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned struct.
	//
	// example:
	//
	// [{"name":"Edas Ingress Url Analysis","type":"edas-ingress-url-analysis","url":"https://g.console.aliyun.com/d/1036052989950239-11040375-66-3/edas-ingress-url-analysis?var-clusterId=29ksa&var-regionId=cn-hangzhou"}]
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request, You can use the ID to locate logs and troubleshoot issues.
	//
	// example:
	//
	// 2983BEF7-4A0D-47A2-94A2-8E9C5E63****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDashboardsByNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDashboardsByNameResponseBody) GoString() string {
	return s.String()
}

func (s *ListDashboardsByNameResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListDashboardsByNameResponseBody) GetData() *string {
	return s.Data
}

func (s *ListDashboardsByNameResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListDashboardsByNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDashboardsByNameResponseBody) SetCode(v int32) *ListDashboardsByNameResponseBody {
	s.Code = &v
	return s
}

func (s *ListDashboardsByNameResponseBody) SetData(v string) *ListDashboardsByNameResponseBody {
	s.Data = &v
	return s
}

func (s *ListDashboardsByNameResponseBody) SetMessage(v string) *ListDashboardsByNameResponseBody {
	s.Message = &v
	return s
}

func (s *ListDashboardsByNameResponseBody) SetRequestId(v string) *ListDashboardsByNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDashboardsByNameResponseBody) Validate() error {
	return dara.Validate(s)
}
