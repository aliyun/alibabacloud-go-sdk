// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyStageModelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyStageModelResponseBody
	GetRequestId() *string
}

type ModifyStageModelResponseBody struct {
	// example:
	//
	// 6C87A26A-6A18-4B8E-8099-705278381xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyStageModelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyStageModelResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyStageModelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyStageModelResponseBody) SetRequestId(v string) *ModifyStageModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyStageModelResponseBody) Validate() error {
	return dara.Validate(s)
}
