// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataLevelRuleConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDataLevelRuleConfigResponseBody
	GetRequestId() *string
	SetResult(v bool) *DeleteDataLevelRuleConfigResponseBody
	GetResult() *bool
	SetSuccess(v bool) *DeleteDataLevelRuleConfigResponseBody
	GetSuccess() *bool
}

type DeleteDataLevelRuleConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D8749D65-E80A-433C-AF1B-CE9C180FF3B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The execution result of the operation. Valid values:
	//
	// - true: The request was successful.
	//
	// - false: The request failed.
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - true: The request was successful.
	//
	// - false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteDataLevelRuleConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataLevelRuleConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDataLevelRuleConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDataLevelRuleConfigResponseBody) GetResult() *bool {
	return s.Result
}

func (s *DeleteDataLevelRuleConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteDataLevelRuleConfigResponseBody) SetRequestId(v string) *DeleteDataLevelRuleConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDataLevelRuleConfigResponseBody) SetResult(v bool) *DeleteDataLevelRuleConfigResponseBody {
	s.Result = &v
	return s
}

func (s *DeleteDataLevelRuleConfigResponseBody) SetSuccess(v bool) *DeleteDataLevelRuleConfigResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteDataLevelRuleConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
