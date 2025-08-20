// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAppUseTimeReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRetCode(v int32) *AppUseTimeReportResponseBody
	GetRetCode() *int32
	SetRetMsg(v string) *AppUseTimeReportResponseBody
	GetRetMsg() *string
	SetRetValue(v bool) *AppUseTimeReportResponseBody
	GetRetValue() *bool
}

type AppUseTimeReportResponseBody struct {
	// example:
	//
	// 0
	RetCode *int32  `json:"RetCode,omitempty" xml:"RetCode,omitempty"`
	RetMsg  *string `json:"RetMsg,omitempty" xml:"RetMsg,omitempty"`
	// example:
	//
	// true
	RetValue *bool `json:"RetValue,omitempty" xml:"RetValue,omitempty"`
}

func (s AppUseTimeReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AppUseTimeReportResponseBody) GoString() string {
	return s.String()
}

func (s *AppUseTimeReportResponseBody) GetRetCode() *int32 {
	return s.RetCode
}

func (s *AppUseTimeReportResponseBody) GetRetMsg() *string {
	return s.RetMsg
}

func (s *AppUseTimeReportResponseBody) GetRetValue() *bool {
	return s.RetValue
}

func (s *AppUseTimeReportResponseBody) SetRetCode(v int32) *AppUseTimeReportResponseBody {
	s.RetCode = &v
	return s
}

func (s *AppUseTimeReportResponseBody) SetRetMsg(v string) *AppUseTimeReportResponseBody {
	s.RetMsg = &v
	return s
}

func (s *AppUseTimeReportResponseBody) SetRetValue(v bool) *AppUseTimeReportResponseBody {
	s.RetValue = &v
	return s
}

func (s *AppUseTimeReportResponseBody) Validate() error {
	return dara.Validate(s)
}
