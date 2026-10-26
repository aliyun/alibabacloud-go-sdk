// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAiproEmbedConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyAiproEmbedConfigResponseBody
	GetRequestId() *string
	SetResult(v bool) *ModifyAiproEmbedConfigResponseBody
	GetResult() *bool
	SetSuccess(v bool) *ModifyAiproEmbedConfigResponseBody
	GetSuccess() *bool
}

type ModifyAiproEmbedConfigResponseBody struct {
	// example:
	//
	// 46e53746***70
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyAiproEmbedConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAiproEmbedConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAiproEmbedConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAiproEmbedConfigResponseBody) GetResult() *bool {
	return s.Result
}

func (s *ModifyAiproEmbedConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyAiproEmbedConfigResponseBody) SetRequestId(v string) *ModifyAiproEmbedConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAiproEmbedConfigResponseBody) SetResult(v bool) *ModifyAiproEmbedConfigResponseBody {
	s.Result = &v
	return s
}

func (s *ModifyAiproEmbedConfigResponseBody) SetSuccess(v bool) *ModifyAiproEmbedConfigResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyAiproEmbedConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
