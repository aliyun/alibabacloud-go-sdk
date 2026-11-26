// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeployHttpApiRoutesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BatchDeployHttpApiRoutesResponseBody
	GetCode() *string
	SetData(v *BatchDeployHttpApiRoutesResponseBodyData) *BatchDeployHttpApiRoutesResponseBody
	GetData() *BatchDeployHttpApiRoutesResponseBodyData
	SetMessage(v string) *BatchDeployHttpApiRoutesResponseBody
	GetMessage() *string
	SetRequestId(v string) *BatchDeployHttpApiRoutesResponseBody
	GetRequestId() *string
}

type BatchDeployHttpApiRoutesResponseBody struct {
	// The response code of the API call. For the acceptance or publish result of each route, see data.results.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The route-level results of the batch publish. A synchronous publish also returns the number of successful and failed routes. An asynchronous publish returns only the route-level acceptance results.
	Data *BatchDeployHttpApiRoutesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The response message of the API call. If a route fails or is rejected, see the errorMessage of the corresponding route for the cause.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID of the API call. You can invoke this ID to locate the active request.
	//
	// example:
	//
	// F7B8A1C2-1234-5678-90AB-CDEF12345678
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s BatchDeployHttpApiRoutesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchDeployHttpApiRoutesResponseBody) GoString() string {
	return s.String()
}

func (s *BatchDeployHttpApiRoutesResponseBody) GetCode() *string {
	return s.Code
}

func (s *BatchDeployHttpApiRoutesResponseBody) GetData() *BatchDeployHttpApiRoutesResponseBodyData {
	return s.Data
}

func (s *BatchDeployHttpApiRoutesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BatchDeployHttpApiRoutesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchDeployHttpApiRoutesResponseBody) SetCode(v string) *BatchDeployHttpApiRoutesResponseBody {
	s.Code = &v
	return s
}

func (s *BatchDeployHttpApiRoutesResponseBody) SetData(v *BatchDeployHttpApiRoutesResponseBodyData) *BatchDeployHttpApiRoutesResponseBody {
	s.Data = v
	return s
}

func (s *BatchDeployHttpApiRoutesResponseBody) SetMessage(v string) *BatchDeployHttpApiRoutesResponseBody {
	s.Message = &v
	return s
}

func (s *BatchDeployHttpApiRoutesResponseBody) SetRequestId(v string) *BatchDeployHttpApiRoutesResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchDeployHttpApiRoutesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BatchDeployHttpApiRoutesResponseBodyData struct {
	// The number of routes that fail to be published in the batch. This parameter is returned only for a synchronous publish and can be 0.
	//
	// example:
	//
	// 0
	FailedCount *int32 `json:"failedCount,omitempty" xml:"failedCount,omitempty"`
	// A map of results with route IDs as keys. If other routes in the same batch fail, the successfully published routes remain valid.
	//
	// example:
	//
	// {"hr-example001":{"routeId":"hr-example001","status":"Succeeded"},"hr-example002":{"routeId":"hr-example002","status":"Succeeded"}}
	Results map[string]*DataResultsValue `json:"results,omitempty" xml:"results,omitempty"`
	// The number of routes that are successfully published in the batch. This parameter is returned only for a synchronous publish and can be 0.
	//
	// example:
	//
	// 2
	SuccessCount *int32 `json:"successCount,omitempty" xml:"successCount,omitempty"`
}

func (s BatchDeployHttpApiRoutesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s BatchDeployHttpApiRoutesResponseBodyData) GoString() string {
	return s.String()
}

func (s *BatchDeployHttpApiRoutesResponseBodyData) GetFailedCount() *int32 {
	return s.FailedCount
}

func (s *BatchDeployHttpApiRoutesResponseBodyData) GetResults() map[string]*DataResultsValue {
	return s.Results
}

func (s *BatchDeployHttpApiRoutesResponseBodyData) GetSuccessCount() *int32 {
	return s.SuccessCount
}

func (s *BatchDeployHttpApiRoutesResponseBodyData) SetFailedCount(v int32) *BatchDeployHttpApiRoutesResponseBodyData {
	s.FailedCount = &v
	return s
}

func (s *BatchDeployHttpApiRoutesResponseBodyData) SetResults(v map[string]*DataResultsValue) *BatchDeployHttpApiRoutesResponseBodyData {
	s.Results = v
	return s
}

func (s *BatchDeployHttpApiRoutesResponseBodyData) SetSuccessCount(v int32) *BatchDeployHttpApiRoutesResponseBodyData {
	s.SuccessCount = &v
	return s
}

func (s *BatchDeployHttpApiRoutesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
