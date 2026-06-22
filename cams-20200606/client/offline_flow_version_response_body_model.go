// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOfflineFlowVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *OfflineFlowVersionResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *OfflineFlowVersionResponseBody
	GetCode() *string
	SetMessage(v string) *OfflineFlowVersionResponseBody
	GetMessage() *string
	SetRequestId(v string) *OfflineFlowVersionResponseBody
	GetRequestId() *string
	SetResponse(v map[string]interface{}) *OfflineFlowVersionResponseBody
	GetResponse() map[string]interface{}
	SetSuccess(v bool) *OfflineFlowVersionResponseBody
	GetSuccess() *bool
}

type OfflineFlowVersionResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 示例值示例值
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 示例值示例值
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Response  map[string]interface{} `json:"Response,omitempty" xml:"Response,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s OfflineFlowVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OfflineFlowVersionResponseBody) GoString() string {
	return s.String()
}

func (s *OfflineFlowVersionResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *OfflineFlowVersionResponseBody) GetCode() *string {
	return s.Code
}

func (s *OfflineFlowVersionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *OfflineFlowVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OfflineFlowVersionResponseBody) GetResponse() map[string]interface{} {
	return s.Response
}

func (s *OfflineFlowVersionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *OfflineFlowVersionResponseBody) SetAccessDeniedDetail(v string) *OfflineFlowVersionResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *OfflineFlowVersionResponseBody) SetCode(v string) *OfflineFlowVersionResponseBody {
	s.Code = &v
	return s
}

func (s *OfflineFlowVersionResponseBody) SetMessage(v string) *OfflineFlowVersionResponseBody {
	s.Message = &v
	return s
}

func (s *OfflineFlowVersionResponseBody) SetRequestId(v string) *OfflineFlowVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *OfflineFlowVersionResponseBody) SetResponse(v map[string]interface{}) *OfflineFlowVersionResponseBody {
	s.Response = v
	return s
}

func (s *OfflineFlowVersionResponseBody) SetSuccess(v bool) *OfflineFlowVersionResponseBody {
	s.Success = &v
	return s
}

func (s *OfflineFlowVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
