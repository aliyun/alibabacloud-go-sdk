// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveRTCStatsV2ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SaveRTCStatsV2ResponseBody
	GetCode() *string
	SetHttpStatusCode(v int64) *SaveRTCStatsV2ResponseBody
	GetHttpStatusCode() *int64
	SetMessage(v string) *SaveRTCStatsV2ResponseBody
	GetMessage() *string
	SetRequestId(v string) *SaveRTCStatsV2ResponseBody
	GetRequestId() *string
	SetRowCount(v int64) *SaveRTCStatsV2ResponseBody
	GetRowCount() *int64
	SetSuccess(v bool) *SaveRTCStatsV2ResponseBody
	GetSuccess() *bool
	SetTimeStamp(v int64) *SaveRTCStatsV2ResponseBody
	GetTimeStamp() *int64
}

type SaveRTCStatsV2ResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int64  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// BA03159C-E808-4FF1-B27E-A61B6E888D7F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	RowCount *int64 `json:"RowCount,omitempty" xml:"RowCount,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 1647309061000
	TimeStamp *int64 `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s SaveRTCStatsV2ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveRTCStatsV2ResponseBody) GoString() string {
	return s.String()
}

func (s *SaveRTCStatsV2ResponseBody) GetCode() *string {
	return s.Code
}

func (s *SaveRTCStatsV2ResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *SaveRTCStatsV2ResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SaveRTCStatsV2ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveRTCStatsV2ResponseBody) GetRowCount() *int64 {
	return s.RowCount
}

func (s *SaveRTCStatsV2ResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SaveRTCStatsV2ResponseBody) GetTimeStamp() *int64 {
	return s.TimeStamp
}

func (s *SaveRTCStatsV2ResponseBody) SetCode(v string) *SaveRTCStatsV2ResponseBody {
	s.Code = &v
	return s
}

func (s *SaveRTCStatsV2ResponseBody) SetHttpStatusCode(v int64) *SaveRTCStatsV2ResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SaveRTCStatsV2ResponseBody) SetMessage(v string) *SaveRTCStatsV2ResponseBody {
	s.Message = &v
	return s
}

func (s *SaveRTCStatsV2ResponseBody) SetRequestId(v string) *SaveRTCStatsV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveRTCStatsV2ResponseBody) SetRowCount(v int64) *SaveRTCStatsV2ResponseBody {
	s.RowCount = &v
	return s
}

func (s *SaveRTCStatsV2ResponseBody) SetSuccess(v bool) *SaveRTCStatsV2ResponseBody {
	s.Success = &v
	return s
}

func (s *SaveRTCStatsV2ResponseBody) SetTimeStamp(v int64) *SaveRTCStatsV2ResponseBody {
	s.TimeStamp = &v
	return s
}

func (s *SaveRTCStatsV2ResponseBody) Validate() error {
	return dara.Validate(s)
}
