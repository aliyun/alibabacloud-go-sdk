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
	// Return code: 200 indicates success, all others indicate failure.
	//
	// Important
	//
	// - This parameter indicates whether the interface was called correctly. For detailed return code explanations, please refer to the error codes.
	//
	// - Please check the business verification results through the fields in ResultObject.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Interface call return message.
	//
	// Important
	//
	// - This parameter only indicates whether there was an exception with the interface.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 130A2C10-B9EE-4D84-88E3-5384FF03****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Result object.
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
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCardVerifyResponseBodyResultObject struct {
	// Identity verification result:
	//
	// - 1: Consistent
	//
	// - 2: Inconsistent
	//
	// - 3: No Record Found
	//
	// example:
	//
	// 1
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// Submitted ID card information for verification.
	//
	// example:
	//
	// {"address":"浙江省杭州市余杭区文一西路969号","birthDate":"19901226","certName":"张三","certNo":"1234561990122*****","nationality":"汉","authority":"xxx公安局","startDate":"20201130","endDate":"20301130"}
	CardInfo *string `json:"CardInfo,omitempty" xml:"CardInfo,omitempty"`
	// Image comparison score.
	//
	// example:
	//
	// {
	//
	//  "verifyScore": 50.28594166529785
	//
	// }
	FaceDetail *string `json:"FaceDetail,omitempty" xml:"FaceDetail,omitempty"`
	// ID card information read by OCR.
	//
	// example:
	//
	// {"address":"浙江省杭州市余杭区文一西路969号","birthDate":"19901226","certName":"张三","certNo":"1234561990122*****","nationality":"汉","authority":"xxx公安局","startDate":"20201130","endDate":"20301130"}
	OcrCardInfo *string `json:"OcrCardInfo,omitempty" xml:"OcrCardInfo,omitempty"`
	// Returned photo URLs.
	//
	// - certUrl  Front side
	//
	// - certNationalUrl  National emblem side
	//
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
