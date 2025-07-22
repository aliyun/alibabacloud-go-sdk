// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExpressConnectRouterSupportedRegionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListExpressConnectRouterSupportedRegionResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListExpressConnectRouterSupportedRegionResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListExpressConnectRouterSupportedRegionResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListExpressConnectRouterSupportedRegionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListExpressConnectRouterSupportedRegionResponseBody
	GetSuccess() *bool
	SetSupportedRegionIdList(v []*string) *ListExpressConnectRouterSupportedRegionResponseBody
	GetSupportedRegionIdList() []*string
}

type ListExpressConnectRouterSupportedRegionResponseBody struct {
	// The response code. The status code 200 indicates that the request was successful. Other status codes indicate that the request failed. For more information, see Error codes.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6FABF516-FED3-5697-BDA2-B18C5D9A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- True
	//
	// 	- False
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The regions in which the ECR feature is activated.
	SupportedRegionIdList []*string `json:"SupportedRegionIdList,omitempty" xml:"SupportedRegionIdList,omitempty" type:"Repeated"`
}

func (s ListExpressConnectRouterSupportedRegionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListExpressConnectRouterSupportedRegionResponseBody) GoString() string {
	return s.String()
}

func (s *ListExpressConnectRouterSupportedRegionResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListExpressConnectRouterSupportedRegionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListExpressConnectRouterSupportedRegionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListExpressConnectRouterSupportedRegionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListExpressConnectRouterSupportedRegionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListExpressConnectRouterSupportedRegionResponseBody) GetSupportedRegionIdList() []*string {
	return s.SupportedRegionIdList
}

func (s *ListExpressConnectRouterSupportedRegionResponseBody) SetCode(v string) *ListExpressConnectRouterSupportedRegionResponseBody {
	s.Code = &v
	return s
}

func (s *ListExpressConnectRouterSupportedRegionResponseBody) SetHttpStatusCode(v int32) *ListExpressConnectRouterSupportedRegionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListExpressConnectRouterSupportedRegionResponseBody) SetMessage(v string) *ListExpressConnectRouterSupportedRegionResponseBody {
	s.Message = &v
	return s
}

func (s *ListExpressConnectRouterSupportedRegionResponseBody) SetRequestId(v string) *ListExpressConnectRouterSupportedRegionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListExpressConnectRouterSupportedRegionResponseBody) SetSuccess(v bool) *ListExpressConnectRouterSupportedRegionResponseBody {
	s.Success = &v
	return s
}

func (s *ListExpressConnectRouterSupportedRegionResponseBody) SetSupportedRegionIdList(v []*string) *ListExpressConnectRouterSupportedRegionResponseBody {
	s.SupportedRegionIdList = v
	return s
}

func (s *ListExpressConnectRouterSupportedRegionResponseBody) Validate() error {
	return dara.Validate(s)
}
