// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIndexMonitorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetIndexMonitorResponseBody
	GetCode() *string
	SetData(v interface{}) *GetIndexMonitorResponseBody
	GetData() interface{}
	SetMessage(v string) *GetIndexMonitorResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetIndexMonitorResponseBody
	GetRequestId() *string
	SetStatus(v int32) *GetIndexMonitorResponseBody
	GetStatus() *int32
	SetSuccess(v bool) *GetIndexMonitorResponseBody
	GetSuccess() *bool
}

type GetIndexMonitorResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The core data object of the response.
	//
	// **pipelineCommercialType*	- (String): The edition of the knowledge base.
	//
	// - standard: Standard Edition
	//
	// - enterprise: Ultimate Edition
	//
	// **storageMonitorData*	- (Object): The storage monitoring data of the knowledge base.
	//
	// - indexStorageLimit (Number): The index storage limit of the knowledge base, in GB.
	//
	// - indexStorageUsage (Number): The current index storage usage of the knowledge base, in GB.
	//
	// **pipelineCommercialCu*	- (Integer): The number of RCU for the Ultimate Edition knowledge base. For example: 2.
	//
	// **qpsMonitorData*	- (Object): The aggregated retrieval monitoring data for the knowledge base over the entire query period.
	//
	// - peakQps (Integer): The peak QPS over the entire time period.
	//
	// - totalRequests (Integer): The total number of requests over the entire time period.
	//
	// - avgQpsOfActiveSeconds (Number): The average QPS during active seconds over the entire time period. Active seconds are seconds in which calls were made.
	//
	// - monitorData (Array): An array of detailed monitoring data broken down by time window. Each object in the array represents the performance statistics for a single time window.
	//
	//   <details>
	//
	//   <summary>
	//
	//   Sub-properties
	//
	//   </summary>
	//
	//   - successData (Object): The statistics for successful requests within this window.
	//
	//   - limitData (Object): The statistics for rate-limited requests within this window.
	//
	//   - failData (Object): The statistics for failed calls within this window.
	//
	//   - peakQpsInWindowRange (Integer): The total peak QPS within this window (successful + rate-limited + failed).
	//
	//   - totalRequests (Integer): The total number of requests within this window (successful + rate-limited + failed).
	//
	//   - windowRange (Integer): The start time of the time window (UNIX timestamp in seconds).
	//
	//   - windowRangeEnd (Integer): The end time of the time window (UNIX timestamp in seconds).
	//
	//   - avgQpsOfActiveSeconds (Number): The average QPS during active seconds within this window.
	//
	//   **The successData, limitData, and failData objects have the same internal structure, as described below:**
	//
	//   - peakQpsInWindowRange (Integer): The peak QPS for the corresponding status.
	//
	//   - totalRequests (Integer): The total number of requests for the corresponding status.
	//
	//   - avgQpsOfActiveSeconds (Number): The average QPS during active seconds for the corresponding status.
	//
	//   </details>
	//
	// example:
	//
	// {
	//
	//     "code": "Success",
	//
	//     "status_code": 200,
	//
	//     "data": {
	//
	// "pipelineCommercialType": "standard",       "storageMonitorData": Object{...},
	//
	//         "qpsMonitorData": Object{...}
	//
	//     },
	//
	//     "success": true,
	//
	//     "message": "success",
	//
	//     "request_id": "65d34b79-b97e-478e-a0a3-xxx",
	//
	//     "status": "SUCCESS"
	//
	// }
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// The status message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 778C0B3B-xxxx-5FC1-A947-36EDD13606AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status code returned by the operation.
	//
	// example:
	//
	// SUCCESS
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetIndexMonitorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetIndexMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *GetIndexMonitorResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetIndexMonitorResponseBody) GetData() interface{} {
	return s.Data
}

func (s *GetIndexMonitorResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetIndexMonitorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetIndexMonitorResponseBody) GetStatus() *int32 {
	return s.Status
}

func (s *GetIndexMonitorResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetIndexMonitorResponseBody) SetCode(v string) *GetIndexMonitorResponseBody {
	s.Code = &v
	return s
}

func (s *GetIndexMonitorResponseBody) SetData(v interface{}) *GetIndexMonitorResponseBody {
	s.Data = v
	return s
}

func (s *GetIndexMonitorResponseBody) SetMessage(v string) *GetIndexMonitorResponseBody {
	s.Message = &v
	return s
}

func (s *GetIndexMonitorResponseBody) SetRequestId(v string) *GetIndexMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIndexMonitorResponseBody) SetStatus(v int32) *GetIndexMonitorResponseBody {
	s.Status = &v
	return s
}

func (s *GetIndexMonitorResponseBody) SetSuccess(v bool) *GetIndexMonitorResponseBody {
	s.Success = &v
	return s
}

func (s *GetIndexMonitorResponseBody) Validate() error {
	return dara.Validate(s)
}
