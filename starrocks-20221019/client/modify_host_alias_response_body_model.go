// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHostAliasResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *ModifyHostAliasResponseBody
	GetData() *bool
	SetErrCode(v string) *ModifyHostAliasResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModifyHostAliasResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModifyHostAliasResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModifyHostAliasResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyHostAliasResponseBody
	GetSuccess() *bool
}

type ModifyHostAliasResponseBody struct {
	// example:
	//
	// 24151320976****
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// InvalidParams
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// Invalid params: [instance not exists].
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 32A44F0D-BFF6-5664-999A-218BBDE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyHostAliasResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyHostAliasResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyHostAliasResponseBody) GetData() *bool {
	return s.Data
}

func (s *ModifyHostAliasResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModifyHostAliasResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModifyHostAliasResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyHostAliasResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyHostAliasResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyHostAliasResponseBody) SetData(v bool) *ModifyHostAliasResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyHostAliasResponseBody) SetErrCode(v string) *ModifyHostAliasResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModifyHostAliasResponseBody) SetErrMessage(v string) *ModifyHostAliasResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModifyHostAliasResponseBody) SetHttpStatusCode(v int32) *ModifyHostAliasResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyHostAliasResponseBody) SetRequestId(v string) *ModifyHostAliasResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyHostAliasResponseBody) SetSuccess(v bool) *ModifyHostAliasResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyHostAliasResponseBody) Validate() error {
	return dara.Validate(s)
}
