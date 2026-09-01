// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentlessTaskUsedSizeEstimateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEstimateUsedSize(v int64) *GetAgentlessTaskUsedSizeEstimateResponseBody
	GetEstimateUsedSize() *int64
	SetRequestId(v string) *GetAgentlessTaskUsedSizeEstimateResponseBody
	GetRequestId() *string
}

type GetAgentlessTaskUsedSizeEstimateResponseBody struct {
	// The estimated scan volume of the detection task. Unit: GB.
	//
	// example:
	//
	// 1
	EstimateUsedSize *int64 `json:"EstimateUsedSize,omitempty" xml:"EstimateUsedSize,omitempty"`
	// The request ID. Alibaba Cloud generates a unique identifier for each API request. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// CD380235-A0B8-540D-A0D5-D62884469E3C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAgentlessTaskUsedSizeEstimateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAgentlessTaskUsedSizeEstimateResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgentlessTaskUsedSizeEstimateResponseBody) GetEstimateUsedSize() *int64 {
	return s.EstimateUsedSize
}

func (s *GetAgentlessTaskUsedSizeEstimateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAgentlessTaskUsedSizeEstimateResponseBody) SetEstimateUsedSize(v int64) *GetAgentlessTaskUsedSizeEstimateResponseBody {
	s.EstimateUsedSize = &v
	return s
}

func (s *GetAgentlessTaskUsedSizeEstimateResponseBody) SetRequestId(v string) *GetAgentlessTaskUsedSizeEstimateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAgentlessTaskUsedSizeEstimateResponseBody) Validate() error {
	return dara.Validate(s)
}
