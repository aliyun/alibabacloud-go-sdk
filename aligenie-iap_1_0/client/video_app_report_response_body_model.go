// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVideoAppReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRetCode(v int32) *VideoAppReportResponseBody
	GetRetCode() *int32
	SetRetMsg(v string) *VideoAppReportResponseBody
	GetRetMsg() *string
	SetRetValue(v bool) *VideoAppReportResponseBody
	GetRetValue() *bool
}

type VideoAppReportResponseBody struct {
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

func (s VideoAppReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s VideoAppReportResponseBody) GoString() string {
	return s.String()
}

func (s *VideoAppReportResponseBody) GetRetCode() *int32 {
	return s.RetCode
}

func (s *VideoAppReportResponseBody) GetRetMsg() *string {
	return s.RetMsg
}

func (s *VideoAppReportResponseBody) GetRetValue() *bool {
	return s.RetValue
}

func (s *VideoAppReportResponseBody) SetRetCode(v int32) *VideoAppReportResponseBody {
	s.RetCode = &v
	return s
}

func (s *VideoAppReportResponseBody) SetRetMsg(v string) *VideoAppReportResponseBody {
	s.RetMsg = &v
	return s
}

func (s *VideoAppReportResponseBody) SetRetValue(v bool) *VideoAppReportResponseBody {
	s.RetValue = &v
	return s
}

func (s *VideoAppReportResponseBody) Validate() error {
	return dara.Validate(s)
}
