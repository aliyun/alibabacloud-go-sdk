// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSecondRankResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateSecondRankResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *CreateSecondRankResponseBody
	GetResult() map[string]interface{}
}

type CreateSecondRankResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ABCDEFGH
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the fine sort expression.
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s CreateSecondRankResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSecondRankResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSecondRankResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSecondRankResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *CreateSecondRankResponseBody) SetRequestId(v string) *CreateSecondRankResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSecondRankResponseBody) SetResult(v map[string]interface{}) *CreateSecondRankResponseBody {
	s.Result = v
	return s
}

func (s *CreateSecondRankResponseBody) Validate() error {
	return dara.Validate(s)
}
