// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCardVerifyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeCardVerifyResponseBody
	GetCode() *string
	SetMessage(v string) *DescribeCardVerifyResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeCardVerifyResponseBody
	GetRequestId() *string
	SetResultObject(v *DescribeCardVerifyResponseBodyResultObject) *DescribeCardVerifyResponseBody
	GetResultObject() *DescribeCardVerifyResponseBodyResultObject
}

type DescribeCardVerifyResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 130A2C10-B9EE-4D84-88E3-5384FF03****
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultObject *DescribeCardVerifyResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s DescribeCardVerifyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCardVerifyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCardVerifyResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeCardVerifyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeCardVerifyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCardVerifyResponseBody) GetResultObject() *DescribeCardVerifyResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeCardVerifyResponseBody) SetCode(v string) *DescribeCardVerifyResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeCardVerifyResponseBody) SetMessage(v string) *DescribeCardVerifyResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeCardVerifyResponseBody) SetRequestId(v string) *DescribeCardVerifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCardVerifyResponseBody) SetResultObject(v *DescribeCardVerifyResponseBodyResultObject) *DescribeCardVerifyResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeCardVerifyResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeCardVerifyResponseBodyResultObject struct {
	// example:
	//
	// 1
	BizCode     *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	CardInfo    *string `json:"CardInfo,omitempty" xml:"CardInfo,omitempty"`
	FaceDetail  *string `json:"FaceDetail,omitempty" xml:"FaceDetail,omitempty"`
	OcrCardInfo *string `json:"OcrCardInfo,omitempty" xml:"OcrCardInfo,omitempty"`
	// example:
	//
	// {
	//
	//     "certUrl": "https://cn-shanghai-aliyun-nmeta.cn-shanghai.oss.aliyuncs.com/verify/xxxxxxx/xxxxxxxxxx_ocridface_dbf2_normal.jpeg",
	//
	//     "certNationalUrl": "https://cn-shanghai-aliyun-nmeta.cn-shanghai.oss.aliyuncs.com/verify/xxxxxxxxxx/xxxxxxxxxxx_ocridback_e8a3_normal.jpeg"
	//
	// }
	PictureInfo *string `json:"PictureInfo,omitempty" xml:"PictureInfo,omitempty"`
}

func (s DescribeCardVerifyResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeCardVerifyResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeCardVerifyResponseBodyResultObject) GetBizCode() *string {
	return s.BizCode
}

func (s *DescribeCardVerifyResponseBodyResultObject) GetCardInfo() *string {
	return s.CardInfo
}

func (s *DescribeCardVerifyResponseBodyResultObject) GetFaceDetail() *string {
	return s.FaceDetail
}

func (s *DescribeCardVerifyResponseBodyResultObject) GetOcrCardInfo() *string {
	return s.OcrCardInfo
}

func (s *DescribeCardVerifyResponseBodyResultObject) GetPictureInfo() *string {
	return s.PictureInfo
}

func (s *DescribeCardVerifyResponseBodyResultObject) SetBizCode(v string) *DescribeCardVerifyResponseBodyResultObject {
	s.BizCode = &v
	return s
}

func (s *DescribeCardVerifyResponseBodyResultObject) SetCardInfo(v string) *DescribeCardVerifyResponseBodyResultObject {
	s.CardInfo = &v
	return s
}

func (s *DescribeCardVerifyResponseBodyResultObject) SetFaceDetail(v string) *DescribeCardVerifyResponseBodyResultObject {
	s.FaceDetail = &v
	return s
}

func (s *DescribeCardVerifyResponseBodyResultObject) SetOcrCardInfo(v string) *DescribeCardVerifyResponseBodyResultObject {
	s.OcrCardInfo = &v
	return s
}

func (s *DescribeCardVerifyResponseBodyResultObject) SetPictureInfo(v string) *DescribeCardVerifyResponseBodyResultObject {
	s.PictureInfo = &v
	return s
}

func (s *DescribeCardVerifyResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
