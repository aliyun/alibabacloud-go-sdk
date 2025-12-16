// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveSecondRankResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveSecondRankResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *RemoveSecondRankResponseBody
	GetResult() map[string]interface{}
}

type RemoveSecondRankResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0A6EB64B-B4C8-CF02-810F-E660812972FF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The returned results.
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s RemoveSecondRankResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveSecondRankResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveSecondRankResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveSecondRankResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *RemoveSecondRankResponseBody) SetRequestId(v string) *RemoveSecondRankResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveSecondRankResponseBody) SetResult(v map[string]interface{}) *RemoveSecondRankResponseBody {
	s.Result = v
	return s
}

func (s *RemoveSecondRankResponseBody) Validate() error {
	return dara.Validate(s)
}
