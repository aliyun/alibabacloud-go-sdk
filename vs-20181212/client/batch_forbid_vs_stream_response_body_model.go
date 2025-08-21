// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchForbidVsStreamResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetForbidResult(v *BatchForbidVsStreamResponseBodyForbidResult) *BatchForbidVsStreamResponseBody
	GetForbidResult() *BatchForbidVsStreamResponseBodyForbidResult
	SetRequestId(v string) *BatchForbidVsStreamResponseBody
	GetRequestId() *string
}

type BatchForbidVsStreamResponseBody struct {
	ForbidResult *BatchForbidVsStreamResponseBodyForbidResult `json:"ForbidResult,omitempty" xml:"ForbidResult,omitempty" type:"Struct"`
	// example:
	//
	// B058D71B-76EA-5DF6-ACAF-A617C1E7937F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchForbidVsStreamResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchForbidVsStreamResponseBody) GoString() string {
	return s.String()
}

func (s *BatchForbidVsStreamResponseBody) GetForbidResult() *BatchForbidVsStreamResponseBodyForbidResult {
	return s.ForbidResult
}

func (s *BatchForbidVsStreamResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchForbidVsStreamResponseBody) SetForbidResult(v *BatchForbidVsStreamResponseBodyForbidResult) *BatchForbidVsStreamResponseBody {
	s.ForbidResult = v
	return s
}

func (s *BatchForbidVsStreamResponseBody) SetRequestId(v string) *BatchForbidVsStreamResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchForbidVsStreamResponseBody) Validate() error {
	return dara.Validate(s)
}

type BatchForbidVsStreamResponseBodyForbidResult struct {
	ForbidResultInfo []*BatchForbidVsStreamResponseBodyForbidResultForbidResultInfo `json:"ForbidResultInfo,omitempty" xml:"ForbidResultInfo,omitempty" type:"Repeated"`
}

func (s BatchForbidVsStreamResponseBodyForbidResult) String() string {
	return dara.Prettify(s)
}

func (s BatchForbidVsStreamResponseBodyForbidResult) GoString() string {
	return s.String()
}

func (s *BatchForbidVsStreamResponseBodyForbidResult) GetForbidResultInfo() []*BatchForbidVsStreamResponseBodyForbidResultForbidResultInfo {
	return s.ForbidResultInfo
}

func (s *BatchForbidVsStreamResponseBodyForbidResult) SetForbidResultInfo(v []*BatchForbidVsStreamResponseBodyForbidResultForbidResultInfo) *BatchForbidVsStreamResponseBodyForbidResult {
	s.ForbidResultInfo = v
	return s
}

func (s *BatchForbidVsStreamResponseBodyForbidResult) Validate() error {
	return dara.Validate(s)
}

type BatchForbidVsStreamResponseBodyForbidResultForbidResultInfo struct {
	Channels *BatchForbidVsStreamResponseBodyForbidResultForbidResultInfoChannels `json:"Channels,omitempty" xml:"Channels,omitempty" type:"Struct"`
	// example:
	//
	// 2
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// ok
	Detail *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// example:
	//
	// success
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s BatchForbidVsStreamResponseBodyForbidResultForbidResultInfo) String() string {
	return dara.Prettify(s)
}

func (s BatchForbidVsStreamResponseBodyForbidResultForbidResultInfo) GoString() string {
	return s.String()
}

func (s *BatchForbidVsStreamResponseBodyForbidResultForbidResultInfo) GetChannels() *BatchForbidVsStreamResponseBodyForbidResultForbidResultInfoChannels {
	return s.Channels
}

func (s *BatchForbidVsStreamResponseBodyForbidResultForbidResultInfo) GetCount() *int32 {
	return s.Count
}

func (s *BatchForbidVsStreamResponseBodyForbidResultForbidResultInfo) GetDetail() *string {
	return s.Detail
}

func (s *BatchForbidVsStreamResponseBodyForbidResultForbidResultInfo) GetResult() *string {
	return s.Result
}

func (s *BatchForbidVsStreamResponseBodyForbidResultForbidResultInfo) SetChannels(v *BatchForbidVsStreamResponseBodyForbidResultForbidResultInfoChannels) *BatchForbidVsStreamResponseBodyForbidResultForbidResultInfo {
	s.Channels = v
	return s
}

func (s *BatchForbidVsStreamResponseBodyForbidResultForbidResultInfo) SetCount(v int32) *BatchForbidVsStreamResponseBodyForbidResultForbidResultInfo {
	s.Count = &v
	return s
}

func (s *BatchForbidVsStreamResponseBodyForbidResultForbidResultInfo) SetDetail(v string) *BatchForbidVsStreamResponseBodyForbidResultForbidResultInfo {
	s.Detail = &v
	return s
}

func (s *BatchForbidVsStreamResponseBodyForbidResultForbidResultInfo) SetResult(v string) *BatchForbidVsStreamResponseBodyForbidResultForbidResultInfo {
	s.Result = &v
	return s
}

func (s *BatchForbidVsStreamResponseBodyForbidResultForbidResultInfo) Validate() error {
	return dara.Validate(s)
}

type BatchForbidVsStreamResponseBodyForbidResultForbidResultInfoChannels struct {
	Channel []*string `json:"Channel,omitempty" xml:"Channel,omitempty" type:"Repeated"`
}

func (s BatchForbidVsStreamResponseBodyForbidResultForbidResultInfoChannels) String() string {
	return dara.Prettify(s)
}

func (s BatchForbidVsStreamResponseBodyForbidResultForbidResultInfoChannels) GoString() string {
	return s.String()
}

func (s *BatchForbidVsStreamResponseBodyForbidResultForbidResultInfoChannels) GetChannel() []*string {
	return s.Channel
}

func (s *BatchForbidVsStreamResponseBodyForbidResultForbidResultInfoChannels) SetChannel(v []*string) *BatchForbidVsStreamResponseBodyForbidResultForbidResultInfoChannels {
	s.Channel = v
	return s
}

func (s *BatchForbidVsStreamResponseBodyForbidResultForbidResultInfoChannels) Validate() error {
	return dara.Validate(s)
}
