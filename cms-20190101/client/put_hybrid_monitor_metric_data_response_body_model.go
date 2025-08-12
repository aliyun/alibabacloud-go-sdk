// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutHybridMonitorMetricDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *PutHybridMonitorMetricDataResponseBody
	GetCode() *string
	SetErrorDetail(v []*PutHybridMonitorMetricDataResponseBodyErrorDetail) *PutHybridMonitorMetricDataResponseBody
	GetErrorDetail() []*PutHybridMonitorMetricDataResponseBodyErrorDetail
	SetRequestId(v string) *PutHybridMonitorMetricDataResponseBody
	GetRequestId() *string
}

type PutHybridMonitorMetricDataResponseBody struct {
	// The HTTP status code.
	//
	// >  The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of invalid parameters.
	//
	// If a request parameter is invalid, the details of the invalid parameter are returned.
	ErrorDetail []*PutHybridMonitorMetricDataResponseBodyErrorDetail `json:"ErrorDetail,omitempty" xml:"ErrorDetail,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 5DB1CBCA-D14A-55FA-814F-B4DBD9735F68
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PutHybridMonitorMetricDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PutHybridMonitorMetricDataResponseBody) GoString() string {
	return s.String()
}

func (s *PutHybridMonitorMetricDataResponseBody) GetCode() *string {
	return s.Code
}

func (s *PutHybridMonitorMetricDataResponseBody) GetErrorDetail() []*PutHybridMonitorMetricDataResponseBodyErrorDetail {
	return s.ErrorDetail
}

func (s *PutHybridMonitorMetricDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PutHybridMonitorMetricDataResponseBody) SetCode(v string) *PutHybridMonitorMetricDataResponseBody {
	s.Code = &v
	return s
}

func (s *PutHybridMonitorMetricDataResponseBody) SetErrorDetail(v []*PutHybridMonitorMetricDataResponseBodyErrorDetail) *PutHybridMonitorMetricDataResponseBody {
	s.ErrorDetail = v
	return s
}

func (s *PutHybridMonitorMetricDataResponseBody) SetRequestId(v string) *PutHybridMonitorMetricDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutHybridMonitorMetricDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type PutHybridMonitorMetricDataResponseBodyErrorDetail struct {
	// The error message of the invalid parameter.
	//
	// example:
	//
	// label name :123 not match [a-zA-Z_][a-zA-Z0-9_]*
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The position of the error message in the array.
	//
	// example:
	//
	// 0
	Index *int64 `json:"Index,omitempty" xml:"Index,omitempty"`
}

func (s PutHybridMonitorMetricDataResponseBodyErrorDetail) String() string {
	return dara.Prettify(s)
}

func (s PutHybridMonitorMetricDataResponseBodyErrorDetail) GoString() string {
	return s.String()
}

func (s *PutHybridMonitorMetricDataResponseBodyErrorDetail) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *PutHybridMonitorMetricDataResponseBodyErrorDetail) GetIndex() *int64 {
	return s.Index
}

func (s *PutHybridMonitorMetricDataResponseBodyErrorDetail) SetErrorMessage(v string) *PutHybridMonitorMetricDataResponseBodyErrorDetail {
	s.ErrorMessage = &v
	return s
}

func (s *PutHybridMonitorMetricDataResponseBodyErrorDetail) SetIndex(v int64) *PutHybridMonitorMetricDataResponseBodyErrorDetail {
	s.Index = &v
	return s
}

func (s *PutHybridMonitorMetricDataResponseBodyErrorDetail) Validate() error {
	return dara.Validate(s)
}
