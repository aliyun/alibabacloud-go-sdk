// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourceLogStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeResourceLogStatusResponseBody
	GetRequestId() *string
	SetResult(v []*DescribeResourceLogStatusResponseBodyResult) *DescribeResourceLogStatusResponseBody
	GetResult() []*DescribeResourceLogStatusResponseBodyResult
}

type DescribeResourceLogStatusResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0DABF8AB-2321-5F8D-A8D7-922D757FBFFE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result.
	Result []*DescribeResourceLogStatusResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s DescribeResourceLogStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceLogStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeResourceLogStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeResourceLogStatusResponseBody) GetResult() []*DescribeResourceLogStatusResponseBodyResult {
	return s.Result
}

func (s *DescribeResourceLogStatusResponseBody) SetRequestId(v string) *DescribeResourceLogStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeResourceLogStatusResponseBody) SetResult(v []*DescribeResourceLogStatusResponseBodyResult) *DescribeResourceLogStatusResponseBody {
	s.Result = v
	return s
}

func (s *DescribeResourceLogStatusResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeResourceLogStatusResponseBodyResult struct {
	// The protected object.
	//
	// example:
	//
	// alb-wewbb23dfsetetcic****
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// Indicates whether the log collection feature is enabled for the protected object. Valid values:
	//
	// 	- **true:*	- The log collection feature is enabled.
	//
	// 	- **false:*	- The log collection feature is disabled.
	//
	// example:
	//
	// true
	Status      *bool                                                   `json:"Status,omitempty" xml:"Status,omitempty"`
	TraceConfig *DescribeResourceLogStatusResponseBodyResultTraceConfig `json:"TraceConfig,omitempty" xml:"TraceConfig,omitempty" type:"Struct"`
	TraceStatus *bool                                                   `json:"TraceStatus,omitempty" xml:"TraceStatus,omitempty"`
}

func (s DescribeResourceLogStatusResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceLogStatusResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeResourceLogStatusResponseBodyResult) GetResource() *string {
	return s.Resource
}

func (s *DescribeResourceLogStatusResponseBodyResult) GetStatus() *bool {
	return s.Status
}

func (s *DescribeResourceLogStatusResponseBodyResult) GetTraceConfig() *DescribeResourceLogStatusResponseBodyResultTraceConfig {
	return s.TraceConfig
}

func (s *DescribeResourceLogStatusResponseBodyResult) GetTraceStatus() *bool {
	return s.TraceStatus
}

func (s *DescribeResourceLogStatusResponseBodyResult) SetResource(v string) *DescribeResourceLogStatusResponseBodyResult {
	s.Resource = &v
	return s
}

func (s *DescribeResourceLogStatusResponseBodyResult) SetStatus(v bool) *DescribeResourceLogStatusResponseBodyResult {
	s.Status = &v
	return s
}

func (s *DescribeResourceLogStatusResponseBodyResult) SetTraceConfig(v *DescribeResourceLogStatusResponseBodyResultTraceConfig) *DescribeResourceLogStatusResponseBodyResult {
	s.TraceConfig = v
	return s
}

func (s *DescribeResourceLogStatusResponseBodyResult) SetTraceStatus(v bool) *DescribeResourceLogStatusResponseBodyResult {
	s.TraceStatus = &v
	return s
}

func (s *DescribeResourceLogStatusResponseBodyResult) Validate() error {
	if s.TraceConfig != nil {
		if err := s.TraceConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeResourceLogStatusResponseBodyResultTraceConfig struct {
	RatePerMille *int32  `json:"RatePerMille,omitempty" xml:"RatePerMille,omitempty"`
	Workspace    *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s DescribeResourceLogStatusResponseBodyResultTraceConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceLogStatusResponseBodyResultTraceConfig) GoString() string {
	return s.String()
}

func (s *DescribeResourceLogStatusResponseBodyResultTraceConfig) GetRatePerMille() *int32 {
	return s.RatePerMille
}

func (s *DescribeResourceLogStatusResponseBodyResultTraceConfig) GetWorkspace() *string {
	return s.Workspace
}

func (s *DescribeResourceLogStatusResponseBodyResultTraceConfig) SetRatePerMille(v int32) *DescribeResourceLogStatusResponseBodyResultTraceConfig {
	s.RatePerMille = &v
	return s
}

func (s *DescribeResourceLogStatusResponseBodyResultTraceConfig) SetWorkspace(v string) *DescribeResourceLogStatusResponseBodyResultTraceConfig {
	s.Workspace = &v
	return s
}

func (s *DescribeResourceLogStatusResponseBodyResultTraceConfig) Validate() error {
	return dara.Validate(s)
}
