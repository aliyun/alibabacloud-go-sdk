// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveWebRtcInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SaveWebRtcInfoResponseBody
	GetCode() *string
	SetHttpStatusCode(v int64) *SaveWebRtcInfoResponseBody
	GetHttpStatusCode() *int64
	SetMessage(v string) *SaveWebRtcInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *SaveWebRtcInfoResponseBody
	GetRequestId() *string
	SetRowCount(v int64) *SaveWebRtcInfoResponseBody
	GetRowCount() *int64
	SetSuccess(v bool) *SaveWebRtcInfoResponseBody
	GetSuccess() *bool
	SetTimeStamp(v int64) *SaveWebRtcInfoResponseBody
	GetTimeStamp() *int64
}

type SaveWebRtcInfoResponseBody struct {
	// The response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int64 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The response message.
	//
	// example:
	//
	// 无
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CF1C21B9-2D49-4B54-880F-FBE248C16903
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of rows returned upon data storage, indicating how many records were successfully written. If this value is not 1, an abnormal condition may have occurred and requires attention.
	//
	// example:
	//
	// 1
	RowCount *int64 `json:"RowCount,omitempty" xml:"RowCount,omitempty"`
	// Indicates whether the operation succeeded.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The time when the information was stored, formatted as a Unix timestamp in milliseconds.
	//
	// example:
	//
	// 1647262108395
	TimeStamp *int64 `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s SaveWebRtcInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveWebRtcInfoResponseBody) GoString() string {
	return s.String()
}

func (s *SaveWebRtcInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *SaveWebRtcInfoResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *SaveWebRtcInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SaveWebRtcInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveWebRtcInfoResponseBody) GetRowCount() *int64 {
	return s.RowCount
}

func (s *SaveWebRtcInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SaveWebRtcInfoResponseBody) GetTimeStamp() *int64 {
	return s.TimeStamp
}

func (s *SaveWebRtcInfoResponseBody) SetCode(v string) *SaveWebRtcInfoResponseBody {
	s.Code = &v
	return s
}

func (s *SaveWebRtcInfoResponseBody) SetHttpStatusCode(v int64) *SaveWebRtcInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SaveWebRtcInfoResponseBody) SetMessage(v string) *SaveWebRtcInfoResponseBody {
	s.Message = &v
	return s
}

func (s *SaveWebRtcInfoResponseBody) SetRequestId(v string) *SaveWebRtcInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveWebRtcInfoResponseBody) SetRowCount(v int64) *SaveWebRtcInfoResponseBody {
	s.RowCount = &v
	return s
}

func (s *SaveWebRtcInfoResponseBody) SetSuccess(v bool) *SaveWebRtcInfoResponseBody {
	s.Success = &v
	return s
}

func (s *SaveWebRtcInfoResponseBody) SetTimeStamp(v int64) *SaveWebRtcInfoResponseBody {
	s.TimeStamp = &v
	return s
}

func (s *SaveWebRtcInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
