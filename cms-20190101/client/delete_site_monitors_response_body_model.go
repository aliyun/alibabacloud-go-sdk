// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSiteMonitorsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteSiteMonitorsResponseBody
	GetCode() *string
	SetData(v *DeleteSiteMonitorsResponseBodyData) *DeleteSiteMonitorsResponseBody
	GetData() *DeleteSiteMonitorsResponseBodyData
	SetMessage(v string) *DeleteSiteMonitorsResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteSiteMonitorsResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DeleteSiteMonitorsResponseBody
	GetSuccess() *string
}

type DeleteSiteMonitorsResponseBody struct {
	// The HTTP status code.
	//
	// >  The status code 200 indicates that the call was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the site monitoring tasks that were deleted.
	Data *DeleteSiteMonitorsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message. If the call was successful, the value success is returned. If the call failed, an error message such as `TaskId not found` is returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 123BCC5D-8B63-48EA-B747-9A8995BE7AA6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. The value true indicates success. The value false indicates failure.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteSiteMonitorsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSiteMonitorsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSiteMonitorsResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteSiteMonitorsResponseBody) GetData() *DeleteSiteMonitorsResponseBodyData {
	return s.Data
}

func (s *DeleteSiteMonitorsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteSiteMonitorsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSiteMonitorsResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DeleteSiteMonitorsResponseBody) SetCode(v string) *DeleteSiteMonitorsResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteSiteMonitorsResponseBody) SetData(v *DeleteSiteMonitorsResponseBodyData) *DeleteSiteMonitorsResponseBody {
	s.Data = v
	return s
}

func (s *DeleteSiteMonitorsResponseBody) SetMessage(v string) *DeleteSiteMonitorsResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteSiteMonitorsResponseBody) SetRequestId(v string) *DeleteSiteMonitorsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSiteMonitorsResponseBody) SetSuccess(v string) *DeleteSiteMonitorsResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteSiteMonitorsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DeleteSiteMonitorsResponseBodyData struct {
	// The number of the site monitoring tasks that were deleted.
	//
	// example:
	//
	// 0
	Count *int32 `json:"count,omitempty" xml:"count,omitempty"`
}

func (s DeleteSiteMonitorsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteSiteMonitorsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteSiteMonitorsResponseBodyData) GetCount() *int32 {
	return s.Count
}

func (s *DeleteSiteMonitorsResponseBodyData) SetCount(v int32) *DeleteSiteMonitorsResponseBodyData {
	s.Count = &v
	return s
}

func (s *DeleteSiteMonitorsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
