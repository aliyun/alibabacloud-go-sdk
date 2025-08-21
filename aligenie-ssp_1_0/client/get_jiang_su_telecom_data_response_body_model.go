// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJiangSuTelecomDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetJiangSuTelecomDataResponseBody
	GetCode() *int32
	SetMessage(v string) *GetJiangSuTelecomDataResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetJiangSuTelecomDataResponseBody
	GetRequestId() *string
	SetResult(v *GetJiangSuTelecomDataResponseBodyResult) *GetJiangSuTelecomDataResponseBody
	GetResult() *GetJiangSuTelecomDataResponseBodyResult
}

type GetJiangSuTelecomDataResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Id of the request
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 26C9C3D0-160D-5CDE-BF7A-B3C8D14AA949
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *GetJiangSuTelecomDataResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetJiangSuTelecomDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetJiangSuTelecomDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetJiangSuTelecomDataResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetJiangSuTelecomDataResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetJiangSuTelecomDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetJiangSuTelecomDataResponseBody) GetResult() *GetJiangSuTelecomDataResponseBodyResult {
	return s.Result
}

func (s *GetJiangSuTelecomDataResponseBody) SetCode(v int32) *GetJiangSuTelecomDataResponseBody {
	s.Code = &v
	return s
}

func (s *GetJiangSuTelecomDataResponseBody) SetMessage(v string) *GetJiangSuTelecomDataResponseBody {
	s.Message = &v
	return s
}

func (s *GetJiangSuTelecomDataResponseBody) SetRequestId(v string) *GetJiangSuTelecomDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetJiangSuTelecomDataResponseBody) SetResult(v *GetJiangSuTelecomDataResponseBodyResult) *GetJiangSuTelecomDataResponseBody {
	s.Result = v
	return s
}

func (s *GetJiangSuTelecomDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetJiangSuTelecomDataResponseBodyResult struct {
	// example:
	//
	// http://jiangsu-telecom.oss-cn-zhangjiakou.aliyuncs.com/jiangsuTelecom/jiangsu_telecom_2024-11-06_data.xls?Expires=1731056700&OSSAccessKeyId=STS.NUqZx6e1HjWYHhYwyDqbRNySp&Signature=wqMK%2Bspo08cg7xDIrzJdgFpZT3U%3D&security-token=CAIS6wJ1q6Ft5B2yfSjIr5bEEcKCiO5p3ZWySk7ok3kRfe1%2Bobz4kjz2IHhMeXJsBuketv42nmxV7%2FoblrN0UIQAT1HPbsZsq84Pq1%2F4O9GY%2FpXrseBZ08VJ18si00SpsvXJasDVEfn%2FGJ70GX2m%2BwZ3xbzlD0bAO3WuLZyOj7N%2Bc90TRXPWRDFaBdBQVGAAwY1gQhm3D%2Fu2NQPwiWf9FVdhvhEG6Vly8qOi2MaRmHG85R%2FYsrZL%2B9uuc8b5P5A0Y8wlAo6PsbYoJvab4kl58ANX8ap6tqtA9Arcs8uVa1sruE3ebrGIrYQ3dFUgPPRnQvIdtrP1nvt5%2FOXS0p%2Fs01NHNOpWXiLTAoe7247OBeiqO8p%2FKeyjZGQuOTooxiRr2elNRQX4VGsiE7JJQhf7CU293KO0YeZxdAR%2FoOMPnzRBdA2yRaA0rjOmtJPBTFOEIL7ymXpDY8bnsxtwDRYu1%2BXrCqUee2Ik3j4vnOf9Je0agAEDuPfR8GLB8uVv4ZCGRRAM5mV3gKCxl07flVk1UPsxMZWBeVwwALT34lxk4x6ivWA7ZXjXdMmIapup%2FEb6UU%2BWhJH1G4sevkw5%2BGb8h8aRjJHBoxB4YLXNTEb5Rk6sVv%2BoCfkljy3%2FWImWGIMigurkAfOIBgUSql8JgFTxQrOHbiAA
	OssUrl *string `json:"OssUrl,omitempty" xml:"OssUrl,omitempty"`
}

func (s GetJiangSuTelecomDataResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetJiangSuTelecomDataResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetJiangSuTelecomDataResponseBodyResult) GetOssUrl() *string {
	return s.OssUrl
}

func (s *GetJiangSuTelecomDataResponseBodyResult) SetOssUrl(v string) *GetJiangSuTelecomDataResponseBodyResult {
	s.OssUrl = &v
	return s
}

func (s *GetJiangSuTelecomDataResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
