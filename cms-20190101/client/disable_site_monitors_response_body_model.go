// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableSiteMonitorsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DisableSiteMonitorsResponseBody
	GetCode() *string
	SetData(v *DisableSiteMonitorsResponseBodyData) *DisableSiteMonitorsResponseBody
	GetData() *DisableSiteMonitorsResponseBodyData
	SetMessage(v string) *DisableSiteMonitorsResponseBody
	GetMessage() *string
	SetRequestId(v string) *DisableSiteMonitorsResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DisableSiteMonitorsResponseBody
	GetSuccess() *string
}

type DisableSiteMonitorsResponseBody struct {
	// The responses code.
	//
	// >  The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The number of detection points that are affected by the site monitoring tasks.
	Data *DisableSiteMonitorsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3fcd12e7-d387-42ee-b77e-661c775bb17f
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DisableSiteMonitorsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableSiteMonitorsResponseBody) GoString() string {
	return s.String()
}

func (s *DisableSiteMonitorsResponseBody) GetCode() *string {
	return s.Code
}

func (s *DisableSiteMonitorsResponseBody) GetData() *DisableSiteMonitorsResponseBodyData {
	return s.Data
}

func (s *DisableSiteMonitorsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DisableSiteMonitorsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableSiteMonitorsResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DisableSiteMonitorsResponseBody) SetCode(v string) *DisableSiteMonitorsResponseBody {
	s.Code = &v
	return s
}

func (s *DisableSiteMonitorsResponseBody) SetData(v *DisableSiteMonitorsResponseBodyData) *DisableSiteMonitorsResponseBody {
	s.Data = v
	return s
}

func (s *DisableSiteMonitorsResponseBody) SetMessage(v string) *DisableSiteMonitorsResponseBody {
	s.Message = &v
	return s
}

func (s *DisableSiteMonitorsResponseBody) SetRequestId(v string) *DisableSiteMonitorsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableSiteMonitorsResponseBody) SetSuccess(v string) *DisableSiteMonitorsResponseBody {
	s.Success = &v
	return s
}

func (s *DisableSiteMonitorsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DisableSiteMonitorsResponseBodyData struct {
	// The number of detection points.
	//
	// example:
	//
	// 0
	Count *int32 `json:"count,omitempty" xml:"count,omitempty"`
}

func (s DisableSiteMonitorsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DisableSiteMonitorsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DisableSiteMonitorsResponseBodyData) GetCount() *int32 {
	return s.Count
}

func (s *DisableSiteMonitorsResponseBodyData) SetCount(v int32) *DisableSiteMonitorsResponseBodyData {
	s.Count = &v
	return s
}

func (s *DisableSiteMonitorsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
