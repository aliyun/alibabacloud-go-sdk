// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAclGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateAclGroupResponseBody
	GetAccessDeniedDetail() *string
	SetData(v bool) *CreateAclGroupResponseBody
	GetData() *bool
	SetErrCode(v string) *CreateAclGroupResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *CreateAclGroupResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *CreateAclGroupResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *CreateAclGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateAclGroupResponseBody
	GetSuccess() *bool
}

type CreateAclGroupResponseBody struct {
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"accessDeniedDetail,omitempty" xml:"accessDeniedDetail,omitempty"`
	// example:
	//
	// {\\"instanceId\\": \\"c-fa521a2393de4623\\", \\"orderId\\": \\"271129670950939\\"}
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// 0
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// [%60wget Nds0v6lx.popscan.xaliyun.com%60]
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// BE7E95C4-10DE-5EA7-9D10-2D3E0FCCE68C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateAclGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAclGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAclGroupResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateAclGroupResponseBody) GetData() *bool {
	return s.Data
}

func (s *CreateAclGroupResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *CreateAclGroupResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *CreateAclGroupResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateAclGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAclGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateAclGroupResponseBody) SetAccessDeniedDetail(v string) *CreateAclGroupResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateAclGroupResponseBody) SetData(v bool) *CreateAclGroupResponseBody {
	s.Data = &v
	return s
}

func (s *CreateAclGroupResponseBody) SetErrCode(v string) *CreateAclGroupResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateAclGroupResponseBody) SetErrMessage(v string) *CreateAclGroupResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *CreateAclGroupResponseBody) SetHttpStatusCode(v int32) *CreateAclGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateAclGroupResponseBody) SetRequestId(v string) *CreateAclGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAclGroupResponseBody) SetSuccess(v bool) *CreateAclGroupResponseBody {
	s.Success = &v
	return s
}

func (s *CreateAclGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
