// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendLogisticsSmsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *SendLogisticsSmsResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *SendLogisticsSmsResponseBody
	GetCode() *string
	SetData(v *SendLogisticsSmsResponseBodyData) *SendLogisticsSmsResponseBody
	GetData() *SendLogisticsSmsResponseBodyData
	SetMessage(v string) *SendLogisticsSmsResponseBody
	GetMessage() *string
	SetRequestId(v string) *SendLogisticsSmsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SendLogisticsSmsResponseBody
	GetSuccess() *bool
}

type SendLogisticsSmsResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	Code *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *SendLogisticsSmsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 示例值示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 示例值示例值
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SendLogisticsSmsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendLogisticsSmsResponseBody) GoString() string {
	return s.String()
}

func (s *SendLogisticsSmsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *SendLogisticsSmsResponseBody) GetCode() *string {
	return s.Code
}

func (s *SendLogisticsSmsResponseBody) GetData() *SendLogisticsSmsResponseBodyData {
	return s.Data
}

func (s *SendLogisticsSmsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SendLogisticsSmsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SendLogisticsSmsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SendLogisticsSmsResponseBody) SetAccessDeniedDetail(v string) *SendLogisticsSmsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *SendLogisticsSmsResponseBody) SetCode(v string) *SendLogisticsSmsResponseBody {
	s.Code = &v
	return s
}

func (s *SendLogisticsSmsResponseBody) SetData(v *SendLogisticsSmsResponseBodyData) *SendLogisticsSmsResponseBody {
	s.Data = v
	return s
}

func (s *SendLogisticsSmsResponseBody) SetMessage(v string) *SendLogisticsSmsResponseBody {
	s.Message = &v
	return s
}

func (s *SendLogisticsSmsResponseBody) SetRequestId(v string) *SendLogisticsSmsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendLogisticsSmsResponseBody) SetSuccess(v bool) *SendLogisticsSmsResponseBody {
	s.Success = &v
	return s
}

func (s *SendLogisticsSmsResponseBody) Validate() error {
	return dara.Validate(s)
}

type SendLogisticsSmsResponseBodyData struct {
	// example:
	//
	// 示例值示例值示例值
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
}

func (s SendLogisticsSmsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SendLogisticsSmsResponseBodyData) GoString() string {
	return s.String()
}

func (s *SendLogisticsSmsResponseBodyData) GetBizId() *string {
	return s.BizId
}

func (s *SendLogisticsSmsResponseBodyData) SetBizId(v string) *SendLogisticsSmsResponseBodyData {
	s.BizId = &v
	return s
}

func (s *SendLogisticsSmsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
