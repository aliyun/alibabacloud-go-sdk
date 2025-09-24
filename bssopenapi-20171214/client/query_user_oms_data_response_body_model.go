// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUserOmsDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryUserOmsDataResponseBody
	GetCode() *string
	SetData(v *QueryUserOmsDataResponseBodyData) *QueryUserOmsDataResponseBody
	GetData() *QueryUserOmsDataResponseBodyData
	SetMessage(v string) *QueryUserOmsDataResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryUserOmsDataResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryUserOmsDataResponseBody
	GetSuccess() *bool
}

type QueryUserOmsDataResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *QueryUserOmsDataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message returned.
	//
	// example:
	//
	// Successful！
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 6000EE23-274B-4E07-A697-FF2E999520A4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryUserOmsDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryUserOmsDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUserOmsDataResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryUserOmsDataResponseBody) GetData() *QueryUserOmsDataResponseBodyData {
	return s.Data
}

func (s *QueryUserOmsDataResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryUserOmsDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryUserOmsDataResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryUserOmsDataResponseBody) SetCode(v string) *QueryUserOmsDataResponseBody {
	s.Code = &v
	return s
}

func (s *QueryUserOmsDataResponseBody) SetData(v *QueryUserOmsDataResponseBodyData) *QueryUserOmsDataResponseBody {
	s.Data = v
	return s
}

func (s *QueryUserOmsDataResponseBody) SetMessage(v string) *QueryUserOmsDataResponseBody {
	s.Message = &v
	return s
}

func (s *QueryUserOmsDataResponseBody) SetRequestId(v string) *QueryUserOmsDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryUserOmsDataResponseBody) SetSuccess(v bool) *QueryUserOmsDataResponseBody {
	s.Success = &v
	return s
}

func (s *QueryUserOmsDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryUserOmsDataResponseBodyData struct {
	// The ID of the host.
	//
	// example:
	//
	// cn
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// Indicates that the returned usage data starts from the next page. If no value is returned for this parameter or this parameter is not returned, no data can be queried.
	//
	// example:
	//
	// NextToken
	Marker  *string                  `json:"Marker,omitempty" xml:"Marker,omitempty"`
	OmsData []map[string]interface{} `json:"OmsData,omitempty" xml:"OmsData,omitempty" type:"Repeated"`
}

func (s QueryUserOmsDataResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryUserOmsDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryUserOmsDataResponseBodyData) GetHostId() *string {
	return s.HostId
}

func (s *QueryUserOmsDataResponseBodyData) GetMarker() *string {
	return s.Marker
}

func (s *QueryUserOmsDataResponseBodyData) GetOmsData() []map[string]interface{} {
	return s.OmsData
}

func (s *QueryUserOmsDataResponseBodyData) SetHostId(v string) *QueryUserOmsDataResponseBodyData {
	s.HostId = &v
	return s
}

func (s *QueryUserOmsDataResponseBodyData) SetMarker(v string) *QueryUserOmsDataResponseBodyData {
	s.Marker = &v
	return s
}

func (s *QueryUserOmsDataResponseBodyData) SetOmsData(v []map[string]interface{}) *QueryUserOmsDataResponseBodyData {
	s.OmsData = v
	return s
}

func (s *QueryUserOmsDataResponseBodyData) Validate() error {
	return dara.Validate(s)
}
