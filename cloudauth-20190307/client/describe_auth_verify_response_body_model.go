// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAuthVerifyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeAuthVerifyResponseBody
	GetCode() *string
	SetMessage(v string) *DescribeAuthVerifyResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeAuthVerifyResponseBody
	GetRequestId() *string
	SetResult(v *DescribeAuthVerifyResponseBodyResult) *DescribeAuthVerifyResponseBody
	GetResult() *DescribeAuthVerifyResponseBodyResult
}

type DescribeAuthVerifyResponseBody struct {
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeAuthVerifyResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribeAuthVerifyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAuthVerifyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAuthVerifyResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeAuthVerifyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeAuthVerifyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAuthVerifyResponseBody) GetResult() *DescribeAuthVerifyResponseBodyResult {
	return s.Result
}

func (s *DescribeAuthVerifyResponseBody) SetCode(v string) *DescribeAuthVerifyResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAuthVerifyResponseBody) SetMessage(v string) *DescribeAuthVerifyResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAuthVerifyResponseBody) SetRequestId(v string) *DescribeAuthVerifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAuthVerifyResponseBody) SetResult(v *DescribeAuthVerifyResponseBodyResult) *DescribeAuthVerifyResponseBody {
	s.Result = v
	return s
}

func (s *DescribeAuthVerifyResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAuthVerifyResponseBodyResult struct {
	MaterialInfo *string `json:"MaterialInfo,omitempty" xml:"MaterialInfo,omitempty"`
	// example:
	//
	// spoofRiskResult：Y
	//
	// spoofType：SCREEN_REMARK
	SpoofBackInfo *string `json:"SpoofBackInfo,omitempty" xml:"SpoofBackInfo,omitempty"`
	// example:
	//
	// spoofRiskResult：Y
	//
	// spoofType：SCREEN_REMARK
	SpoofInfo *string `json:"SpoofInfo,omitempty" xml:"SpoofInfo,omitempty"`
	// example:
	//
	// 200
	SubCode *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
}

func (s DescribeAuthVerifyResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DescribeAuthVerifyResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeAuthVerifyResponseBodyResult) GetMaterialInfo() *string {
	return s.MaterialInfo
}

func (s *DescribeAuthVerifyResponseBodyResult) GetSpoofBackInfo() *string {
	return s.SpoofBackInfo
}

func (s *DescribeAuthVerifyResponseBodyResult) GetSpoofInfo() *string {
	return s.SpoofInfo
}

func (s *DescribeAuthVerifyResponseBodyResult) GetSubCode() *string {
	return s.SubCode
}

func (s *DescribeAuthVerifyResponseBodyResult) SetMaterialInfo(v string) *DescribeAuthVerifyResponseBodyResult {
	s.MaterialInfo = &v
	return s
}

func (s *DescribeAuthVerifyResponseBodyResult) SetSpoofBackInfo(v string) *DescribeAuthVerifyResponseBodyResult {
	s.SpoofBackInfo = &v
	return s
}

func (s *DescribeAuthVerifyResponseBodyResult) SetSpoofInfo(v string) *DescribeAuthVerifyResponseBodyResult {
	s.SpoofInfo = &v
	return s
}

func (s *DescribeAuthVerifyResponseBodyResult) SetSubCode(v string) *DescribeAuthVerifyResponseBodyResult {
	s.SubCode = &v
	return s
}

func (s *DescribeAuthVerifyResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
