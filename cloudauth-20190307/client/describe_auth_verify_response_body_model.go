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
	// The return code.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The return message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result.
	Result *DescribeAuthVerifyResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
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
	// - Card information read by OCR (ocrIdCardInfo)
	//
	// - Card information photo edited by the client (ocrIdEditInfo)
	//
	// - OSS storage location and link of the OCR photo (ocrPictureFront).
	//
	// example:
	//
	// "ocrIdCardInfo": {
	//
	//     "certName": "张先生",
	//
	//     "sex": "男",
	//
	//     "nationality": "汉",
	//
	//     "birth": "20010213",
	//
	//     "address": "XXX省XX市XXX区XX街道X号",
	//
	//     "certNo": "4****************X",
	//
	//     "authority": "XXX公安局XXX分局",
	//
	//     "startDate": "20250523",
	//
	//     "endDate": "20450523"
	//
	//   },
	//
	// "ocrIdEditInfo": {
	//
	//     "certName": "张先生",
	//
	//     "sex": "男",
	//
	//     "nationality": "汉",
	//
	//     "birth": "20010213",
	//
	//     "address": "XXX省XX市XXX区XX街道X号",
	//
	//     "certNo": "4****************X",
	//
	//     "authority": "XXX公安局XXX分局",
	//
	//     "startDate": "20250523",
	//
	//     "endDate": "20450523"
	//
	//   },
	//
	//   "ocrPictureFront": {
	//
	//     "ossBucketName": "cn-shanghai-aliyun-cloudauth-12********",
	//
	//     "ossIdFaceObjectName": "verify/1234567890/f7ed1ef80ad1234fdsd95c********cd_ocridface_b749.jpeg",
	//
	//     "ossIdFaceUrl": "http://cn-shanghai-aliyun-cloudauth-1234567890.oss-cn-shanghai.aliyuncs.com/verify/1234567890/f7ed1ef80ad1234fdsd95c66d83340cd_ocridface_b749.jpeg?security-token=CAISjdfgeJ1q6F...",
	//
	//     "ossIdNationalEmblemObjectName": "verify/1234567890/f7ed1ef80ad1234fdsd95c********cd_ocridnationalemblem_a3hf.jpeg",
	//
	//     "ossIdNationalEmblemUrl": "http://cn-shanghai-aliyun-cloudauth-1234567890.oss-cn-shanghai.aliyuncs.com/verify/1234567890/f7ed1ef80ad1234fdsd95c66d83340cd_ocridnationalemblem_a3hf.jpeg?security-token=CAISjgJ1q6..."
	//
	//   }
	MaterialInfo *string `json:"MaterialInfo,omitempty" xml:"MaterialInfo,omitempty"`
	// The anti-spoofing detection result for the back side of the document, including the risk determination result and risk type:
	//
	// > -
	//
	// Card front anti-spoofing detection is enabled only when IdSpoof = Y is set in the Initialize operation. Otherwise, spoofRiskResult returns N by default, and spoofType is empty.
	//
	// spoofRiskResult:
	//
	// - Y: Risk detected.
	//
	//  - N: No risk detected.
	//
	// spoofType:
	//
	// - SCREEN_REMARK: Recaptured photo.
	//
	// - PHOTO_COPY: Photocopy.
	//
	// - TAMPER: Digitally tampered.
	//
	// > - This is an algorithm prediction result. This field may not be returned. Avoid setting a mandatory dependency on this field in your business logic.
	//
	// example:
	//
	// spoofRiskResult：Y
	//
	// spoofType：SCREEN_REMARK
	SpoofBackInfo *string `json:"SpoofBackInfo,omitempty" xml:"SpoofBackInfo,omitempty"`
	// The anti-spoofing detection result for the front side of the document, including the risk determination result and risk type:
	//
	// > -
	//
	// Card front anti-spoofing detection is enabled only when IdSpoof = Y is set in the Initialize operation. Otherwise, spoofRiskResult returns N by default, and spoofType is empty.
	//
	// spoofRiskResult:
	//
	// - Y: Risk detected.
	//
	//  - N: No risk detected.
	//
	// spoofType:
	//
	// - SCREEN_REMARK: Recaptured photo.
	//
	// - PHOTO_COPY: Photocopy.
	//
	// - TAMPER: Digitally tampered.
	//
	// > - This is an algorithm prediction result. This field may not be returned. Avoid setting a mandatory dependency on this field in your business logic.
	//
	// example:
	//
	// spoofRiskResult：Y
	//
	// spoofType：SCREEN_REMARK
	SpoofInfo *string `json:"SpoofInfo,omitempty" xml:"SpoofInfo,omitempty"`
	// The result description.
	//
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
