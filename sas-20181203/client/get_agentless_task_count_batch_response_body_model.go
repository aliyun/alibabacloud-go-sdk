// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentlessTaskCountBatchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v map[string]*DataValue) *GetAgentlessTaskCountBatchResponseBody
	GetData() map[string]*DataValue
	SetRequestId(v string) *GetAgentlessTaskCountBatchResponseBody
	GetRequestId() *string
}

type GetAgentlessTaskCountBatchResponseBody struct {
	// The statistics grouped by resource UUID. The key of the map is the resource UUID.
	//
	// example:
	//
	// {"3bb30859-b3b5-4f28-868f-b0892c98****":{"RiskMachine":1,"ScanMachine":1}}
	Data map[string]*DataValue `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID. Alibaba Cloud generates a unique identifier for each request. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// D03DD0FD-6041-5107-AC00-383E28F1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAgentlessTaskCountBatchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAgentlessTaskCountBatchResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgentlessTaskCountBatchResponseBody) GetData() map[string]*DataValue {
	return s.Data
}

func (s *GetAgentlessTaskCountBatchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAgentlessTaskCountBatchResponseBody) SetData(v map[string]*DataValue) *GetAgentlessTaskCountBatchResponseBody {
	s.Data = v
	return s
}

func (s *GetAgentlessTaskCountBatchResponseBody) SetRequestId(v string) *GetAgentlessTaskCountBatchResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAgentlessTaskCountBatchResponseBody) Validate() error {
	return dara.Validate(s)
}
