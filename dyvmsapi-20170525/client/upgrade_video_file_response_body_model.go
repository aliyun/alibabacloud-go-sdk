// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeVideoFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UpgradeVideoFileResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *UpgradeVideoFileResponseBody
	GetCode() *string
	SetData(v map[string]interface{}) *UpgradeVideoFileResponseBody
	GetData() map[string]interface{}
	SetMessage(v string) *UpgradeVideoFileResponseBody
	GetMessage() *string
	SetSuccess(v bool) *UpgradeVideoFileResponseBody
	GetSuccess() *bool
}

type UpgradeVideoFileResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 示例值示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpgradeVideoFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpgradeVideoFileResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeVideoFileResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UpgradeVideoFileResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpgradeVideoFileResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *UpgradeVideoFileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpgradeVideoFileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpgradeVideoFileResponseBody) SetAccessDeniedDetail(v string) *UpgradeVideoFileResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpgradeVideoFileResponseBody) SetCode(v string) *UpgradeVideoFileResponseBody {
	s.Code = &v
	return s
}

func (s *UpgradeVideoFileResponseBody) SetData(v map[string]interface{}) *UpgradeVideoFileResponseBody {
	s.Data = v
	return s
}

func (s *UpgradeVideoFileResponseBody) SetMessage(v string) *UpgradeVideoFileResponseBody {
	s.Message = &v
	return s
}

func (s *UpgradeVideoFileResponseBody) SetSuccess(v bool) *UpgradeVideoFileResponseBody {
	s.Success = &v
	return s
}

func (s *UpgradeVideoFileResponseBody) Validate() error {
	return dara.Validate(s)
}
