// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveWebRTCStatsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SaveWebRTCStatsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int64) *SaveWebRTCStatsResponseBody
	GetHttpStatusCode() *int64
	SetMessage(v string) *SaveWebRTCStatsResponseBody
	GetMessage() *string
	SetRequestId(v string) *SaveWebRTCStatsResponseBody
	GetRequestId() *string
	SetRowCount(v int64) *SaveWebRTCStatsResponseBody
	GetRowCount() *int64
	SetSuccess(v bool) *SaveWebRTCStatsResponseBody
	GetSuccess() *bool
	SetTimeStamp(v int64) *SaveWebRTCStatsResponseBody
	GetTimeStamp() *int64
}

type SaveWebRTCStatsResponseBody struct {
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
	// CF1C21B9-2D49-4B54-880F-FBE248C16903
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
	// 1555492246000
	TimeStamp *int64 `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s SaveWebRTCStatsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveWebRTCStatsResponseBody) GoString() string {
	return s.String()
}

func (s *SaveWebRTCStatsResponseBody) GetCode() *string {
	return s.Code
}

func (s *SaveWebRTCStatsResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *SaveWebRTCStatsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SaveWebRTCStatsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveWebRTCStatsResponseBody) GetRowCount() *int64 {
	return s.RowCount
}

func (s *SaveWebRTCStatsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SaveWebRTCStatsResponseBody) GetTimeStamp() *int64 {
	return s.TimeStamp
}

func (s *SaveWebRTCStatsResponseBody) SetCode(v string) *SaveWebRTCStatsResponseBody {
	s.Code = &v
	return s
}

func (s *SaveWebRTCStatsResponseBody) SetHttpStatusCode(v int64) *SaveWebRTCStatsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SaveWebRTCStatsResponseBody) SetMessage(v string) *SaveWebRTCStatsResponseBody {
	s.Message = &v
	return s
}

func (s *SaveWebRTCStatsResponseBody) SetRequestId(v string) *SaveWebRTCStatsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveWebRTCStatsResponseBody) SetRowCount(v int64) *SaveWebRTCStatsResponseBody {
	s.RowCount = &v
	return s
}

func (s *SaveWebRTCStatsResponseBody) SetSuccess(v bool) *SaveWebRTCStatsResponseBody {
	s.Success = &v
	return s
}

func (s *SaveWebRTCStatsResponseBody) SetTimeStamp(v int64) *SaveWebRTCStatsResponseBody {
	s.TimeStamp = &v
	return s
}

func (s *SaveWebRTCStatsResponseBody) Validate() error {
	return dara.Validate(s)
}
