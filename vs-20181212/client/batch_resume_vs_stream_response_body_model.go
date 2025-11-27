// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchResumeVsStreamResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BatchResumeVsStreamResponseBody
	GetRequestId() *string
	SetResumeResult(v *BatchResumeVsStreamResponseBodyResumeResult) *BatchResumeVsStreamResponseBody
	GetResumeResult() *BatchResumeVsStreamResponseBodyResumeResult
}

type BatchResumeVsStreamResponseBody struct {
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId    *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResumeResult *BatchResumeVsStreamResponseBodyResumeResult `json:"ResumeResult,omitempty" xml:"ResumeResult,omitempty" type:"Struct"`
}

func (s BatchResumeVsStreamResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchResumeVsStreamResponseBody) GoString() string {
	return s.String()
}

func (s *BatchResumeVsStreamResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchResumeVsStreamResponseBody) GetResumeResult() *BatchResumeVsStreamResponseBodyResumeResult {
	return s.ResumeResult
}

func (s *BatchResumeVsStreamResponseBody) SetRequestId(v string) *BatchResumeVsStreamResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchResumeVsStreamResponseBody) SetResumeResult(v *BatchResumeVsStreamResponseBodyResumeResult) *BatchResumeVsStreamResponseBody {
	s.ResumeResult = v
	return s
}

func (s *BatchResumeVsStreamResponseBody) Validate() error {
	if s.ResumeResult != nil {
		if err := s.ResumeResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BatchResumeVsStreamResponseBodyResumeResult struct {
	ResumeResultInfo []*BatchResumeVsStreamResponseBodyResumeResultResumeResultInfo `json:"ResumeResultInfo,omitempty" xml:"ResumeResultInfo,omitempty" type:"Repeated"`
}

func (s BatchResumeVsStreamResponseBodyResumeResult) String() string {
	return dara.Prettify(s)
}

func (s BatchResumeVsStreamResponseBodyResumeResult) GoString() string {
	return s.String()
}

func (s *BatchResumeVsStreamResponseBodyResumeResult) GetResumeResultInfo() []*BatchResumeVsStreamResponseBodyResumeResultResumeResultInfo {
	return s.ResumeResultInfo
}

func (s *BatchResumeVsStreamResponseBodyResumeResult) SetResumeResultInfo(v []*BatchResumeVsStreamResponseBodyResumeResultResumeResultInfo) *BatchResumeVsStreamResponseBodyResumeResult {
	s.ResumeResultInfo = v
	return s
}

func (s *BatchResumeVsStreamResponseBodyResumeResult) Validate() error {
	if s.ResumeResultInfo != nil {
		for _, item := range s.ResumeResultInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BatchResumeVsStreamResponseBodyResumeResultResumeResultInfo struct {
	Channels *BatchResumeVsStreamResponseBodyResumeResultResumeResultInfoChannels `json:"Channels,omitempty" xml:"Channels,omitempty" type:"Struct"`
	// example:
	//
	// 1
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

func (s BatchResumeVsStreamResponseBodyResumeResultResumeResultInfo) String() string {
	return dara.Prettify(s)
}

func (s BatchResumeVsStreamResponseBodyResumeResultResumeResultInfo) GoString() string {
	return s.String()
}

func (s *BatchResumeVsStreamResponseBodyResumeResultResumeResultInfo) GetChannels() *BatchResumeVsStreamResponseBodyResumeResultResumeResultInfoChannels {
	return s.Channels
}

func (s *BatchResumeVsStreamResponseBodyResumeResultResumeResultInfo) GetCount() *int32 {
	return s.Count
}

func (s *BatchResumeVsStreamResponseBodyResumeResultResumeResultInfo) GetDetail() *string {
	return s.Detail
}

func (s *BatchResumeVsStreamResponseBodyResumeResultResumeResultInfo) GetResult() *string {
	return s.Result
}

func (s *BatchResumeVsStreamResponseBodyResumeResultResumeResultInfo) SetChannels(v *BatchResumeVsStreamResponseBodyResumeResultResumeResultInfoChannels) *BatchResumeVsStreamResponseBodyResumeResultResumeResultInfo {
	s.Channels = v
	return s
}

func (s *BatchResumeVsStreamResponseBodyResumeResultResumeResultInfo) SetCount(v int32) *BatchResumeVsStreamResponseBodyResumeResultResumeResultInfo {
	s.Count = &v
	return s
}

func (s *BatchResumeVsStreamResponseBodyResumeResultResumeResultInfo) SetDetail(v string) *BatchResumeVsStreamResponseBodyResumeResultResumeResultInfo {
	s.Detail = &v
	return s
}

func (s *BatchResumeVsStreamResponseBodyResumeResultResumeResultInfo) SetResult(v string) *BatchResumeVsStreamResponseBodyResumeResultResumeResultInfo {
	s.Result = &v
	return s
}

func (s *BatchResumeVsStreamResponseBodyResumeResultResumeResultInfo) Validate() error {
	if s.Channels != nil {
		if err := s.Channels.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BatchResumeVsStreamResponseBodyResumeResultResumeResultInfoChannels struct {
	Channel []*string `json:"Channel,omitempty" xml:"Channel,omitempty" type:"Repeated"`
}

func (s BatchResumeVsStreamResponseBodyResumeResultResumeResultInfoChannels) String() string {
	return dara.Prettify(s)
}

func (s BatchResumeVsStreamResponseBodyResumeResultResumeResultInfoChannels) GoString() string {
	return s.String()
}

func (s *BatchResumeVsStreamResponseBodyResumeResultResumeResultInfoChannels) GetChannel() []*string {
	return s.Channel
}

func (s *BatchResumeVsStreamResponseBodyResumeResultResumeResultInfoChannels) SetChannel(v []*string) *BatchResumeVsStreamResponseBodyResumeResultResumeResultInfoChannels {
	s.Channel = v
	return s
}

func (s *BatchResumeVsStreamResponseBodyResumeResultResumeResultInfoChannels) Validate() error {
	return dara.Validate(s)
}
