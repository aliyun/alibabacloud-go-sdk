// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateJenkinsImageRegistryPersistenceDayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *UpdateJenkinsImageRegistryPersistenceDayResponseBody
	GetData() *bool
	SetHttpStatusCode(v int32) *UpdateJenkinsImageRegistryPersistenceDayResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *UpdateJenkinsImageRegistryPersistenceDayResponseBody
	GetRequestId() *string
	SetTimeCost(v int64) *UpdateJenkinsImageRegistryPersistenceDayResponseBody
	GetTimeCost() *int64
}

type UpdateJenkinsImageRegistryPersistenceDayResponseBody struct {
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 69F88BA1-004C-51E2-BF5C-A3220E5A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The time consumed. Unit: seconds.
	//
	// example:
	//
	// 1
	TimeCost *int64 `json:"TimeCost,omitempty" xml:"TimeCost,omitempty"`
}

func (s UpdateJenkinsImageRegistryPersistenceDayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateJenkinsImageRegistryPersistenceDayResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateJenkinsImageRegistryPersistenceDayResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateJenkinsImageRegistryPersistenceDayResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateJenkinsImageRegistryPersistenceDayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateJenkinsImageRegistryPersistenceDayResponseBody) GetTimeCost() *int64 {
	return s.TimeCost
}

func (s *UpdateJenkinsImageRegistryPersistenceDayResponseBody) SetData(v bool) *UpdateJenkinsImageRegistryPersistenceDayResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateJenkinsImageRegistryPersistenceDayResponseBody) SetHttpStatusCode(v int32) *UpdateJenkinsImageRegistryPersistenceDayResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateJenkinsImageRegistryPersistenceDayResponseBody) SetRequestId(v string) *UpdateJenkinsImageRegistryPersistenceDayResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateJenkinsImageRegistryPersistenceDayResponseBody) SetTimeCost(v int64) *UpdateJenkinsImageRegistryPersistenceDayResponseBody {
	s.TimeCost = &v
	return s
}

func (s *UpdateJenkinsImageRegistryPersistenceDayResponseBody) Validate() error {
	return dara.Validate(s)
}
