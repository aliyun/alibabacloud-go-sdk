// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchSubmitReviewInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *BatchSubmitReviewInfoRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *BatchSubmitReviewInfoRequest
	GetJsonStr() *string
}

type BatchSubmitReviewInfoRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// example:
	//
	// {"isSchemeData":1,"schemeTaskConfigId":334,"sourceDataType":2,"startTime":"2022-08-25 00:00:00","endTime":"2022-09-23 23:59:59","sessionList":[{"taskId":"20220831-F8D7F4DF-0A16-1A1C-BA63-28F203922692","fileId":"20220831-164343"},{"taskId":"20220831-F2A50A72-82C4-1E3F-A1FD-52A662283D25","fileId":"20220831-164343"}]}
	JsonStr *string `json:"jsonStr,omitempty" xml:"jsonStr,omitempty"`
}

func (s BatchSubmitReviewInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchSubmitReviewInfoRequest) GoString() string {
	return s.String()
}

func (s *BatchSubmitReviewInfoRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *BatchSubmitReviewInfoRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *BatchSubmitReviewInfoRequest) SetBaseMeAgentId(v int64) *BatchSubmitReviewInfoRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *BatchSubmitReviewInfoRequest) SetJsonStr(v string) *BatchSubmitReviewInfoRequest {
	s.JsonStr = &v
	return s
}

func (s *BatchSubmitReviewInfoRequest) Validate() error {
	return dara.Validate(s)
}
