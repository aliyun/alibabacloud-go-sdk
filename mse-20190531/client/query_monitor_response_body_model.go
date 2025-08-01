// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMonitorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*QueryMonitorResponseBodyData) *QueryMonitorResponseBody
	GetData() []*QueryMonitorResponseBodyData
	SetErrorCode(v string) *QueryMonitorResponseBody
	GetErrorCode() *string
	SetMessage(v string) *QueryMonitorResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryMonitorResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryMonitorResponseBody
	GetSuccess() *bool
}

type QueryMonitorResponseBody struct {
	// The details of the data.
	//
	// example:
	//
	// 6
	Data []*QueryMonitorResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// mse-100-000
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The message returned.
	//
	// example:
	//
	// The request was successfully processed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// ADDD8AB7-8D1C-4697-A83E-410D2607****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`: The request was successful.
	//
	// 	- `false`: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryMonitorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMonitorResponseBody) GetData() []*QueryMonitorResponseBodyData {
	return s.Data
}

func (s *QueryMonitorResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *QueryMonitorResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryMonitorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryMonitorResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryMonitorResponseBody) SetData(v []*QueryMonitorResponseBodyData) *QueryMonitorResponseBody {
	s.Data = v
	return s
}

func (s *QueryMonitorResponseBody) SetErrorCode(v string) *QueryMonitorResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *QueryMonitorResponseBody) SetMessage(v string) *QueryMonitorResponseBody {
	s.Message = &v
	return s
}

func (s *QueryMonitorResponseBody) SetRequestId(v string) *QueryMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMonitorResponseBody) SetSuccess(v bool) *QueryMonitorResponseBody {
	s.Success = &v
	return s
}

func (s *QueryMonitorResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryMonitorResponseBodyData struct {
	// The prefix of the name.
	//
	// example:
	//
	// mse-xxxx-xxxxxx
	ClusterNamePrefix *string `json:"clusterNamePrefix,omitempty" xml:"clusterNamePrefix,omitempty"`
	// The name of the pod.
	//
	// example:
	//
	// mse-xxxxxx-xxxxxx-reg-center-0-0
	PodName *string `json:"podName,omitempty" xml:"podName,omitempty"`
	// The details of the data.
	Values []map[string]interface{} `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s QueryMonitorResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryMonitorResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryMonitorResponseBodyData) GetClusterNamePrefix() *string {
	return s.ClusterNamePrefix
}

func (s *QueryMonitorResponseBodyData) GetPodName() *string {
	return s.PodName
}

func (s *QueryMonitorResponseBodyData) GetValues() []map[string]interface{} {
	return s.Values
}

func (s *QueryMonitorResponseBodyData) SetClusterNamePrefix(v string) *QueryMonitorResponseBodyData {
	s.ClusterNamePrefix = &v
	return s
}

func (s *QueryMonitorResponseBodyData) SetPodName(v string) *QueryMonitorResponseBodyData {
	s.PodName = &v
	return s
}

func (s *QueryMonitorResponseBodyData) SetValues(v []map[string]interface{}) *QueryMonitorResponseBodyData {
	s.Values = v
	return s
}

func (s *QueryMonitorResponseBodyData) Validate() error {
	return dara.Validate(s)
}
