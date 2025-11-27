// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyLogisticsSmsMailNoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *VerifyLogisticsSmsMailNoResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *VerifyLogisticsSmsMailNoResponseBody
	GetCode() *string
	SetData(v *VerifyLogisticsSmsMailNoResponseBodyData) *VerifyLogisticsSmsMailNoResponseBody
	GetData() *VerifyLogisticsSmsMailNoResponseBodyData
	SetMessage(v string) *VerifyLogisticsSmsMailNoResponseBody
	GetMessage() *string
	SetRequestId(v string) *VerifyLogisticsSmsMailNoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *VerifyLogisticsSmsMailNoResponseBody
	GetSuccess() *bool
}

type VerifyLogisticsSmsMailNoResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 示例值
	Code *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *VerifyLogisticsSmsMailNoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 示例值示例值示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s VerifyLogisticsSmsMailNoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s VerifyLogisticsSmsMailNoResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyLogisticsSmsMailNoResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *VerifyLogisticsSmsMailNoResponseBody) GetCode() *string {
	return s.Code
}

func (s *VerifyLogisticsSmsMailNoResponseBody) GetData() *VerifyLogisticsSmsMailNoResponseBodyData {
	return s.Data
}

func (s *VerifyLogisticsSmsMailNoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *VerifyLogisticsSmsMailNoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *VerifyLogisticsSmsMailNoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *VerifyLogisticsSmsMailNoResponseBody) SetAccessDeniedDetail(v string) *VerifyLogisticsSmsMailNoResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *VerifyLogisticsSmsMailNoResponseBody) SetCode(v string) *VerifyLogisticsSmsMailNoResponseBody {
	s.Code = &v
	return s
}

func (s *VerifyLogisticsSmsMailNoResponseBody) SetData(v *VerifyLogisticsSmsMailNoResponseBodyData) *VerifyLogisticsSmsMailNoResponseBody {
	s.Data = v
	return s
}

func (s *VerifyLogisticsSmsMailNoResponseBody) SetMessage(v string) *VerifyLogisticsSmsMailNoResponseBody {
	s.Message = &v
	return s
}

func (s *VerifyLogisticsSmsMailNoResponseBody) SetRequestId(v string) *VerifyLogisticsSmsMailNoResponseBody {
	s.RequestId = &v
	return s
}

func (s *VerifyLogisticsSmsMailNoResponseBody) SetSuccess(v bool) *VerifyLogisticsSmsMailNoResponseBody {
	s.Success = &v
	return s
}

func (s *VerifyLogisticsSmsMailNoResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type VerifyLogisticsSmsMailNoResponseBodyData struct {
	// example:
	//
	// 示例值
	ExpressCompanyCode *string `json:"ExpressCompanyCode,omitempty" xml:"ExpressCompanyCode,omitempty"`
	// example:
	//
	// 示例值
	MobileSuffix *string `json:"MobileSuffix,omitempty" xml:"MobileSuffix,omitempty"`
	// example:
	//
	// false
	VerifyResult *bool `json:"VerifyResult,omitempty" xml:"VerifyResult,omitempty"`
}

func (s VerifyLogisticsSmsMailNoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s VerifyLogisticsSmsMailNoResponseBodyData) GoString() string {
	return s.String()
}

func (s *VerifyLogisticsSmsMailNoResponseBodyData) GetExpressCompanyCode() *string {
	return s.ExpressCompanyCode
}

func (s *VerifyLogisticsSmsMailNoResponseBodyData) GetMobileSuffix() *string {
	return s.MobileSuffix
}

func (s *VerifyLogisticsSmsMailNoResponseBodyData) GetVerifyResult() *bool {
	return s.VerifyResult
}

func (s *VerifyLogisticsSmsMailNoResponseBodyData) SetExpressCompanyCode(v string) *VerifyLogisticsSmsMailNoResponseBodyData {
	s.ExpressCompanyCode = &v
	return s
}

func (s *VerifyLogisticsSmsMailNoResponseBodyData) SetMobileSuffix(v string) *VerifyLogisticsSmsMailNoResponseBodyData {
	s.MobileSuffix = &v
	return s
}

func (s *VerifyLogisticsSmsMailNoResponseBodyData) SetVerifyResult(v bool) *VerifyLogisticsSmsMailNoResponseBodyData {
	s.VerifyResult = &v
	return s
}

func (s *VerifyLogisticsSmsMailNoResponseBodyData) Validate() error {
	return dara.Validate(s)
}
