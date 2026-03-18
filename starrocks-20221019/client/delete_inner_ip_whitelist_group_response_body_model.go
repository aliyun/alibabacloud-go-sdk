// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInnerIpWhitelistGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrCode(v string) *DeleteInnerIpWhitelistGroupResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DeleteInnerIpWhitelistGroupResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *DeleteInnerIpWhitelistGroupResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DeleteInnerIpWhitelistGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteInnerIpWhitelistGroupResponseBody
	GetSuccess() *bool
}

type DeleteInnerIpWhitelistGroupResponseBody struct {
	// example:
	//
	// InvalidParams
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// Invalid params: [Region id should be select from set [cn-beijing, cn-hangzhou]]
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

func (s DeleteInnerIpWhitelistGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteInnerIpWhitelistGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInnerIpWhitelistGroupResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DeleteInnerIpWhitelistGroupResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DeleteInnerIpWhitelistGroupResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteInnerIpWhitelistGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteInnerIpWhitelistGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteInnerIpWhitelistGroupResponseBody) SetErrCode(v string) *DeleteInnerIpWhitelistGroupResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DeleteInnerIpWhitelistGroupResponseBody) SetErrMessage(v string) *DeleteInnerIpWhitelistGroupResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DeleteInnerIpWhitelistGroupResponseBody) SetHttpStatusCode(v int32) *DeleteInnerIpWhitelistGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteInnerIpWhitelistGroupResponseBody) SetRequestId(v string) *DeleteInnerIpWhitelistGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteInnerIpWhitelistGroupResponseBody) SetSuccess(v bool) *DeleteInnerIpWhitelistGroupResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteInnerIpWhitelistGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
