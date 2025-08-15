// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTrailStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIsLogging(v bool) *GetTrailStatusResponseBody
	GetIsLogging() *bool
	SetLatestDeliveryError(v string) *GetTrailStatusResponseBody
	GetLatestDeliveryError() *string
	SetLatestDeliveryLogServiceError(v string) *GetTrailStatusResponseBody
	GetLatestDeliveryLogServiceError() *string
	SetLatestDeliveryLogServiceTime(v string) *GetTrailStatusResponseBody
	GetLatestDeliveryLogServiceTime() *string
	SetLatestDeliveryTime(v string) *GetTrailStatusResponseBody
	GetLatestDeliveryTime() *string
	SetOssBucketStatus(v bool) *GetTrailStatusResponseBody
	GetOssBucketStatus() *bool
	SetRequestId(v string) *GetTrailStatusResponseBody
	GetRequestId() *string
	SetSlsLogStoreStatus(v bool) *GetTrailStatusResponseBody
	GetSlsLogStoreStatus() *bool
	SetStartLoggingTime(v string) *GetTrailStatusResponseBody
	GetStartLoggingTime() *string
	SetStopLoggingTime(v string) *GetTrailStatusResponseBody
	GetStopLoggingTime() *string
}

type GetTrailStatusResponseBody struct {
	// Indicates whether logging is enabled for the trail. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	IsLogging *bool `json:"IsLogging,omitempty" xml:"IsLogging,omitempty"`
	// The log of the last failed delivery.
	//
	// example:
	//
	// write sls failed, exception: the parent of sub user must be project owner, itemscount: 1
	LatestDeliveryError *string `json:"LatestDeliveryError,omitempty" xml:"LatestDeliveryError,omitempty"`
	// The log of the last failed delivery to Log Service.
	//
	// example:
	//
	// write sls failed, exception: the parent of sub user must be project owner, itemscount: 1
	LatestDeliveryLogServiceError *string `json:"LatestDeliveryLogServiceError,omitempty" xml:"LatestDeliveryLogServiceError,omitempty"`
	// The most recent time when an event was delivered to Log Service.
	//
	// example:
	//
	// 2021-02-26T09:19:44Z
	LatestDeliveryLogServiceTime *string `json:"LatestDeliveryLogServiceTime,omitempty" xml:"LatestDeliveryLogServiceTime,omitempty"`
	// The most recent time when an event was delivered by the trail.
	//
	// example:
	//
	// 2021-02-26T09:19:44Z
	LatestDeliveryTime *string `json:"LatestDeliveryTime,omitempty" xml:"LatestDeliveryTime,omitempty"`
	// Indicates whether the destination Object Storage Service (OSS) bucket is available. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	OssBucketStatus *bool `json:"OssBucketStatus,omitempty" xml:"OssBucketStatus,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 8067369B-B923-4D26-85BC-61BF33922505
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the destination Log Service Logstore is available. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	SlsLogStoreStatus *bool `json:"SlsLogStoreStatus,omitempty" xml:"SlsLogStoreStatus,omitempty"`
	// The time when logging was last enabled for the trail.
	//
	// example:
	//
	// 2021-02-24T09:19:44Z
	StartLoggingTime *string `json:"StartLoggingTime,omitempty" xml:"StartLoggingTime,omitempty"`
	// The time when logging was last disabled for the trail.
	//
	// example:
	//
	// 2021-02-25T09:19:44Z
	StopLoggingTime *string `json:"StopLoggingTime,omitempty" xml:"StopLoggingTime,omitempty"`
}

func (s GetTrailStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTrailStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetTrailStatusResponseBody) GetIsLogging() *bool {
	return s.IsLogging
}

func (s *GetTrailStatusResponseBody) GetLatestDeliveryError() *string {
	return s.LatestDeliveryError
}

func (s *GetTrailStatusResponseBody) GetLatestDeliveryLogServiceError() *string {
	return s.LatestDeliveryLogServiceError
}

func (s *GetTrailStatusResponseBody) GetLatestDeliveryLogServiceTime() *string {
	return s.LatestDeliveryLogServiceTime
}

func (s *GetTrailStatusResponseBody) GetLatestDeliveryTime() *string {
	return s.LatestDeliveryTime
}

func (s *GetTrailStatusResponseBody) GetOssBucketStatus() *bool {
	return s.OssBucketStatus
}

func (s *GetTrailStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTrailStatusResponseBody) GetSlsLogStoreStatus() *bool {
	return s.SlsLogStoreStatus
}

func (s *GetTrailStatusResponseBody) GetStartLoggingTime() *string {
	return s.StartLoggingTime
}

func (s *GetTrailStatusResponseBody) GetStopLoggingTime() *string {
	return s.StopLoggingTime
}

func (s *GetTrailStatusResponseBody) SetIsLogging(v bool) *GetTrailStatusResponseBody {
	s.IsLogging = &v
	return s
}

func (s *GetTrailStatusResponseBody) SetLatestDeliveryError(v string) *GetTrailStatusResponseBody {
	s.LatestDeliveryError = &v
	return s
}

func (s *GetTrailStatusResponseBody) SetLatestDeliveryLogServiceError(v string) *GetTrailStatusResponseBody {
	s.LatestDeliveryLogServiceError = &v
	return s
}

func (s *GetTrailStatusResponseBody) SetLatestDeliveryLogServiceTime(v string) *GetTrailStatusResponseBody {
	s.LatestDeliveryLogServiceTime = &v
	return s
}

func (s *GetTrailStatusResponseBody) SetLatestDeliveryTime(v string) *GetTrailStatusResponseBody {
	s.LatestDeliveryTime = &v
	return s
}

func (s *GetTrailStatusResponseBody) SetOssBucketStatus(v bool) *GetTrailStatusResponseBody {
	s.OssBucketStatus = &v
	return s
}

func (s *GetTrailStatusResponseBody) SetRequestId(v string) *GetTrailStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTrailStatusResponseBody) SetSlsLogStoreStatus(v bool) *GetTrailStatusResponseBody {
	s.SlsLogStoreStatus = &v
	return s
}

func (s *GetTrailStatusResponseBody) SetStartLoggingTime(v string) *GetTrailStatusResponseBody {
	s.StartLoggingTime = &v
	return s
}

func (s *GetTrailStatusResponseBody) SetStopLoggingTime(v string) *GetTrailStatusResponseBody {
	s.StopLoggingTime = &v
	return s
}

func (s *GetTrailStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
