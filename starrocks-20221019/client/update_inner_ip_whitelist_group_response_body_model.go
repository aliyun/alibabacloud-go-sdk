// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInnerIpWhitelistGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrCode(v string) *UpdateInnerIpWhitelistGroupResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *UpdateInnerIpWhitelistGroupResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *UpdateInnerIpWhitelistGroupResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *UpdateInnerIpWhitelistGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateInnerIpWhitelistGroupResponseBody
	GetSuccess() *bool
}

type UpdateInnerIpWhitelistGroupResponseBody struct {
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
	// 32A44F0D-BFF6-5664-999A-218BBDE74XXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateInnerIpWhitelistGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateInnerIpWhitelistGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInnerIpWhitelistGroupResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *UpdateInnerIpWhitelistGroupResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *UpdateInnerIpWhitelistGroupResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateInnerIpWhitelistGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateInnerIpWhitelistGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateInnerIpWhitelistGroupResponseBody) SetErrCode(v string) *UpdateInnerIpWhitelistGroupResponseBody {
	s.ErrCode = &v
	return s
}

func (s *UpdateInnerIpWhitelistGroupResponseBody) SetErrMessage(v string) *UpdateInnerIpWhitelistGroupResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *UpdateInnerIpWhitelistGroupResponseBody) SetHttpStatusCode(v int32) *UpdateInnerIpWhitelistGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateInnerIpWhitelistGroupResponseBody) SetRequestId(v string) *UpdateInnerIpWhitelistGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateInnerIpWhitelistGroupResponseBody) SetSuccess(v bool) *UpdateInnerIpWhitelistGroupResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateInnerIpWhitelistGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
