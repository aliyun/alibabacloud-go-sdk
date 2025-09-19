// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateJenkinsImageRegistryNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *UpdateJenkinsImageRegistryNameResponseBody
	GetData() *bool
	SetHttpStatusCode(v int32) *UpdateJenkinsImageRegistryNameResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *UpdateJenkinsImageRegistryNameResponseBody
	GetRequestId() *string
	SetTimeCost(v int64) *UpdateJenkinsImageRegistryNameResponseBody
	GetTimeCost() *int64
}

type UpdateJenkinsImageRegistryNameResponseBody struct {
	// The result of the operation. Valid values:
	//
	// 	- **true**: successful
	//
	// 	- **false**: failed
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
	// 4347E985-6E64-467B-96EC-30D4EA9E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The time consumed. Unit: seconds.
	//
	// example:
	//
	// 1
	TimeCost *int64 `json:"TimeCost,omitempty" xml:"TimeCost,omitempty"`
}

func (s UpdateJenkinsImageRegistryNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateJenkinsImageRegistryNameResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateJenkinsImageRegistryNameResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateJenkinsImageRegistryNameResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateJenkinsImageRegistryNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateJenkinsImageRegistryNameResponseBody) GetTimeCost() *int64 {
	return s.TimeCost
}

func (s *UpdateJenkinsImageRegistryNameResponseBody) SetData(v bool) *UpdateJenkinsImageRegistryNameResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateJenkinsImageRegistryNameResponseBody) SetHttpStatusCode(v int32) *UpdateJenkinsImageRegistryNameResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateJenkinsImageRegistryNameResponseBody) SetRequestId(v string) *UpdateJenkinsImageRegistryNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateJenkinsImageRegistryNameResponseBody) SetTimeCost(v int64) *UpdateJenkinsImageRegistryNameResponseBody {
	s.TimeCost = &v
	return s
}

func (s *UpdateJenkinsImageRegistryNameResponseBody) Validate() error {
	return dara.Validate(s)
}
