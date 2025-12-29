// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncDeviceStatusWithAkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *SyncDeviceStatusWithAkResponseBody
	GetMessage() *string
	SetResult(v bool) *SyncDeviceStatusWithAkResponseBody
	GetResult() *bool
	SetStatusCode(v int32) *SyncDeviceStatusWithAkResponseBody
	GetStatusCode() *int32
	SetRequestId(v string) *SyncDeviceStatusWithAkResponseBody
	GetRequestId() *string
}

type SyncDeviceStatusWithAkResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	// example:
	//
	// F12B6147-5925-19E5-A3AD-E1EE1360F34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s SyncDeviceStatusWithAkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SyncDeviceStatusWithAkResponseBody) GoString() string {
	return s.String()
}

func (s *SyncDeviceStatusWithAkResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SyncDeviceStatusWithAkResponseBody) GetResult() *bool {
	return s.Result
}

func (s *SyncDeviceStatusWithAkResponseBody) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SyncDeviceStatusWithAkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SyncDeviceStatusWithAkResponseBody) SetMessage(v string) *SyncDeviceStatusWithAkResponseBody {
	s.Message = &v
	return s
}

func (s *SyncDeviceStatusWithAkResponseBody) SetResult(v bool) *SyncDeviceStatusWithAkResponseBody {
	s.Result = &v
	return s
}

func (s *SyncDeviceStatusWithAkResponseBody) SetStatusCode(v int32) *SyncDeviceStatusWithAkResponseBody {
	s.StatusCode = &v
	return s
}

func (s *SyncDeviceStatusWithAkResponseBody) SetRequestId(v string) *SyncDeviceStatusWithAkResponseBody {
	s.RequestId = &v
	return s
}

func (s *SyncDeviceStatusWithAkResponseBody) Validate() error {
	return dara.Validate(s)
}
