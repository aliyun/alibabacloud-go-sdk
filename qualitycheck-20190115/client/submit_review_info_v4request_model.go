// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitReviewInfoV4Request interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *SubmitReviewInfoV4Request
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *SubmitReviewInfoV4Request
	GetJsonStr() *string
}

type SubmitReviewInfoV4Request struct {
	// example:
	//
	// 12345
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// example:
	//
	// {"comments":"tidComment","jsonReviewResult":"{\\"reviewInfoList\\":[{\\"changed\\":true,\\"comment\\":\\"ridComment\\",\\"matched\\":true,\\"reviewHitResult\\":0,\\"reviewResult\\":1,\\"rid\\":31459,\\"sentenceReviewResults\\":[{\\"changed\\":true,\\"cid\\":95302,\\"comment\\":\\"pidComment\\",\\"hitStatus\\":0,\\"pid\\":\\"0\\",\\"reviewDimensionType\\":\\"2\\",\\"rid\\":31459,\\"sid\\":54104}]}]}","taskId":"20251224-62931498-881B-1436-A93D-1FFBC5D7D4A0","vid":"8cbe2bccf3be4b42bada45136f77d4e9"}
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s SubmitReviewInfoV4Request) String() string {
	return dara.Prettify(s)
}

func (s SubmitReviewInfoV4Request) GoString() string {
	return s.String()
}

func (s *SubmitReviewInfoV4Request) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *SubmitReviewInfoV4Request) GetJsonStr() *string {
	return s.JsonStr
}

func (s *SubmitReviewInfoV4Request) SetBaseMeAgentId(v int64) *SubmitReviewInfoV4Request {
	s.BaseMeAgentId = &v
	return s
}

func (s *SubmitReviewInfoV4Request) SetJsonStr(v string) *SubmitReviewInfoV4Request {
	s.JsonStr = &v
	return s
}

func (s *SubmitReviewInfoV4Request) Validate() error {
	return dara.Validate(s)
}
