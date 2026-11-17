// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchConfigRuntimeChannelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*BatchConfigRuntimeChannelResponseBodyData) *BatchConfigRuntimeChannelResponseBody
	GetData() []*BatchConfigRuntimeChannelResponseBodyData
	SetRequestId(v string) *BatchConfigRuntimeChannelResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *BatchConfigRuntimeChannelResponseBody
	GetTotalCount() *int32
}

type BatchConfigRuntimeChannelResponseBody struct {
	// The list of batch configuration results. Each channel configuration item corresponds to one result.
	Data []*BatchConfigRuntimeChannelResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of configuration results, which is the same as the number of configuration items in ChannelConfigs.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s BatchConfigRuntimeChannelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchConfigRuntimeChannelResponseBody) GoString() string {
	return s.String()
}

func (s *BatchConfigRuntimeChannelResponseBody) GetData() []*BatchConfigRuntimeChannelResponseBodyData {
	return s.Data
}

func (s *BatchConfigRuntimeChannelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchConfigRuntimeChannelResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *BatchConfigRuntimeChannelResponseBody) SetData(v []*BatchConfigRuntimeChannelResponseBodyData) *BatchConfigRuntimeChannelResponseBody {
	s.Data = v
	return s
}

func (s *BatchConfigRuntimeChannelResponseBody) SetRequestId(v string) *BatchConfigRuntimeChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchConfigRuntimeChannelResponseBody) SetTotalCount(v int32) *BatchConfigRuntimeChannelResponseBody {
	s.TotalCount = &v
	return s
}

func (s *BatchConfigRuntimeChannelResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BatchConfigRuntimeChannelResponseBodyData struct {
	// The third-party channel code.
	//
	// example:
	//
	// dingtalk-connector
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error code returned when the configuration fails. This parameter is returned only when Success is false.
	//
	// example:
	//
	// ConfigFailed
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error description returned when the configuration fails. This parameter is returned only when Success is false.
	//
	// example:
	//
	// The runtime is not found.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The agent runtime ID.
	//
	// example:
	//
	// ecd-xgads
	RuntimeId *string `json:"RuntimeId,omitempty" xml:"RuntimeId,omitempty"`
	// Indicates whether the channel configuration item is configured successfully. Valid values:
	//
	// - true: The configuration is successful.
	//
	// - false: The configuration fails. For the cause of the failure, see ErrorCode and Message.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BatchConfigRuntimeChannelResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s BatchConfigRuntimeChannelResponseBodyData) GoString() string {
	return s.String()
}

func (s *BatchConfigRuntimeChannelResponseBodyData) GetCode() *string {
	return s.Code
}

func (s *BatchConfigRuntimeChannelResponseBodyData) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *BatchConfigRuntimeChannelResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *BatchConfigRuntimeChannelResponseBodyData) GetRuntimeId() *string {
	return s.RuntimeId
}

func (s *BatchConfigRuntimeChannelResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *BatchConfigRuntimeChannelResponseBodyData) SetCode(v string) *BatchConfigRuntimeChannelResponseBodyData {
	s.Code = &v
	return s
}

func (s *BatchConfigRuntimeChannelResponseBodyData) SetErrorCode(v string) *BatchConfigRuntimeChannelResponseBodyData {
	s.ErrorCode = &v
	return s
}

func (s *BatchConfigRuntimeChannelResponseBodyData) SetMessage(v string) *BatchConfigRuntimeChannelResponseBodyData {
	s.Message = &v
	return s
}

func (s *BatchConfigRuntimeChannelResponseBodyData) SetRuntimeId(v string) *BatchConfigRuntimeChannelResponseBodyData {
	s.RuntimeId = &v
	return s
}

func (s *BatchConfigRuntimeChannelResponseBodyData) SetSuccess(v bool) *BatchConfigRuntimeChannelResponseBodyData {
	s.Success = &v
	return s
}

func (s *BatchConfigRuntimeChannelResponseBodyData) Validate() error {
	return dara.Validate(s)
}
