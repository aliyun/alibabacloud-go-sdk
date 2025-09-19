// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryJenkinsImageRegistryPersistenceDayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v int32) *QueryJenkinsImageRegistryPersistenceDayResponseBody
	GetData() *int32
	SetHttpStatusCode(v int32) *QueryJenkinsImageRegistryPersistenceDayResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *QueryJenkinsImageRegistryPersistenceDayResponseBody
	GetRequestId() *string
	SetTimeCost(v int64) *QueryJenkinsImageRegistryPersistenceDayResponseBody
	GetTimeCost() *int64
}

type QueryJenkinsImageRegistryPersistenceDayResponseBody struct {
	// The retention period. Unit: days.
	//
	// example:
	//
	// 30
	Data *int32 `json:"Data,omitempty" xml:"Data,omitempty"`
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
	// EA4AC8B7-0C18-5BC1-9DA4-798B3BE4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The time consumed. Unit: seconds.
	//
	// example:
	//
	// 1
	TimeCost *int64 `json:"TimeCost,omitempty" xml:"TimeCost,omitempty"`
}

func (s QueryJenkinsImageRegistryPersistenceDayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryJenkinsImageRegistryPersistenceDayResponseBody) GoString() string {
	return s.String()
}

func (s *QueryJenkinsImageRegistryPersistenceDayResponseBody) GetData() *int32 {
	return s.Data
}

func (s *QueryJenkinsImageRegistryPersistenceDayResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *QueryJenkinsImageRegistryPersistenceDayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryJenkinsImageRegistryPersistenceDayResponseBody) GetTimeCost() *int64 {
	return s.TimeCost
}

func (s *QueryJenkinsImageRegistryPersistenceDayResponseBody) SetData(v int32) *QueryJenkinsImageRegistryPersistenceDayResponseBody {
	s.Data = &v
	return s
}

func (s *QueryJenkinsImageRegistryPersistenceDayResponseBody) SetHttpStatusCode(v int32) *QueryJenkinsImageRegistryPersistenceDayResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryJenkinsImageRegistryPersistenceDayResponseBody) SetRequestId(v string) *QueryJenkinsImageRegistryPersistenceDayResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryJenkinsImageRegistryPersistenceDayResponseBody) SetTimeCost(v int64) *QueryJenkinsImageRegistryPersistenceDayResponseBody {
	s.TimeCost = &v
	return s
}

func (s *QueryJenkinsImageRegistryPersistenceDayResponseBody) Validate() error {
	return dara.Validate(s)
}
