// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTestConnectivityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *TestConnectivityResponseBodyData) *TestConnectivityResponseBody
	GetData() *TestConnectivityResponseBodyData
	SetRequestId(v string) *TestConnectivityResponseBody
	GetRequestId() *string
}

type TestConnectivityResponseBody struct {
	// The result details of the connectivity test.
	Data *TestConnectivityResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The unique identifier of the request, used for troubleshooting and log tracing.
	//
	// example:
	//
	// 1EBD0C05-6C1F-4C95-9C63-B7AB7B5A9C8E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TestConnectivityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TestConnectivityResponseBody) GoString() string {
	return s.String()
}

func (s *TestConnectivityResponseBody) GetData() *TestConnectivityResponseBodyData {
	return s.Data
}

func (s *TestConnectivityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TestConnectivityResponseBody) SetData(v *TestConnectivityResponseBodyData) *TestConnectivityResponseBody {
	s.Data = v
	return s
}

func (s *TestConnectivityResponseBody) SetRequestId(v string) *TestConnectivityResponseBody {
	s.RequestId = &v
	return s
}

func (s *TestConnectivityResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TestConnectivityResponseBodyData struct {
	// The tracking identifier of the connectivity test. The system generates this value for the first call. For polling calls, this value is the same as the CheckId in the request.
	//
	// example:
	//
	// conn-a1b2c3d4e5f67890
	CheckId *string `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	// The end-to-end latency of the target service response, in milliseconds. This value may be 0 or not returned when VerifyStatus is checking.
	//
	// example:
	//
	// 1200
	LatencyMs *int64 `json:"LatencyMs,omitempty" xml:"LatencyMs,omitempty"`
	// The detailed information of the verification result. When VerifyStatus is verified, this value is a snippet of the response returned by the target service. When VerifyStatus is failed, this value describes the error cause, such as authentication failure, timeout, or empty response.
	//
	// example:
	//
	// Connectivity verification succeeded
	VerifyMessage *string `json:"VerifyMessage,omitempty" xml:"VerifyMessage,omitempty"`
	// The current status of the connectivity verification.
	//
	// example:
	//
	// verified
	VerifyStatus *string `json:"VerifyStatus,omitempty" xml:"VerifyStatus,omitempty"`
}

func (s TestConnectivityResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s TestConnectivityResponseBodyData) GoString() string {
	return s.String()
}

func (s *TestConnectivityResponseBodyData) GetCheckId() *string {
	return s.CheckId
}

func (s *TestConnectivityResponseBodyData) GetLatencyMs() *int64 {
	return s.LatencyMs
}

func (s *TestConnectivityResponseBodyData) GetVerifyMessage() *string {
	return s.VerifyMessage
}

func (s *TestConnectivityResponseBodyData) GetVerifyStatus() *string {
	return s.VerifyStatus
}

func (s *TestConnectivityResponseBodyData) SetCheckId(v string) *TestConnectivityResponseBodyData {
	s.CheckId = &v
	return s
}

func (s *TestConnectivityResponseBodyData) SetLatencyMs(v int64) *TestConnectivityResponseBodyData {
	s.LatencyMs = &v
	return s
}

func (s *TestConnectivityResponseBodyData) SetVerifyMessage(v string) *TestConnectivityResponseBodyData {
	s.VerifyMessage = &v
	return s
}

func (s *TestConnectivityResponseBodyData) SetVerifyStatus(v string) *TestConnectivityResponseBodyData {
	s.VerifyStatus = &v
	return s
}

func (s *TestConnectivityResponseBodyData) Validate() error {
	return dara.Validate(s)
}
