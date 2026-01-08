// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMMLActiveResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QueryMMLActiveResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *QueryMMLActiveResponseBody
	GetCode() *string
	SetData(v string) *QueryMMLActiveResponseBody
	GetData() *string
	SetMessage(v string) *QueryMMLActiveResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryMMLActiveResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryMMLActiveResponseBody
	GetSuccess() *bool
}

type QueryMMLActiveResponseBody struct {
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 示例值
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值示例值
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryMMLActiveResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryMMLActiveResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMMLActiveResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QueryMMLActiveResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryMMLActiveResponseBody) GetData() *string {
	return s.Data
}

func (s *QueryMMLActiveResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryMMLActiveResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryMMLActiveResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryMMLActiveResponseBody) SetAccessDeniedDetail(v string) *QueryMMLActiveResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QueryMMLActiveResponseBody) SetCode(v string) *QueryMMLActiveResponseBody {
	s.Code = &v
	return s
}

func (s *QueryMMLActiveResponseBody) SetData(v string) *QueryMMLActiveResponseBody {
	s.Data = &v
	return s
}

func (s *QueryMMLActiveResponseBody) SetMessage(v string) *QueryMMLActiveResponseBody {
	s.Message = &v
	return s
}

func (s *QueryMMLActiveResponseBody) SetRequestId(v string) *QueryMMLActiveResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMMLActiveResponseBody) SetSuccess(v bool) *QueryMMLActiveResponseBody {
	s.Success = &v
	return s
}

func (s *QueryMMLActiveResponseBody) Validate() error {
	return dara.Validate(s)
}
