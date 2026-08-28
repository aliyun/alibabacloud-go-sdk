// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEkycVerifyRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAuthorize(v string) *EkycVerifyRequest
  GetAuthorize() *string 
  SetCrop(v string) *EkycVerifyRequest
  GetCrop() *string 
  SetDocName(v string) *EkycVerifyRequest
  GetDocName() *string 
  SetDocNo(v string) *EkycVerifyRequest
  GetDocNo() *string 
  SetDocType(v string) *EkycVerifyRequest
  GetDocType() *string 
  SetFacePictureBase64(v string) *EkycVerifyRequest
  GetFacePictureBase64() *string 
  SetFacePictureUrl(v string) *EkycVerifyRequest
  GetFacePictureUrl() *string 
  SetFaceQualityCheck(v string) *EkycVerifyRequest
  GetFaceQualityCheck() *string 
  SetIdOcrPictureBase64(v string) *EkycVerifyRequest
  GetIdOcrPictureBase64() *string 
  SetIdOcrPictureUrl(v string) *EkycVerifyRequest
  GetIdOcrPictureUrl() *string 
  SetIdSpoof(v string) *EkycVerifyRequest
  GetIdSpoof() *string 
  SetIdThreshold(v string) *EkycVerifyRequest
  GetIdThreshold() *string 
  SetMerchantBizId(v string) *EkycVerifyRequest
  GetMerchantBizId() *string 
  SetMerchantUserId(v string) *EkycVerifyRequest
  GetMerchantUserId() *string 
  SetProductCode(v string) *EkycVerifyRequest
  GetProductCode() *string 
}

type EkycVerifyRequest struct {
  // Specifies whether to enable authoritative identity verification. Currently, this applies only to second-generation ID cards in mainland China.
  // 
  // example:
  // 
  // T
  Authorize *string `json:"Authorize,omitempty" xml:"Authorize,omitempty"`
  // Specifies whether cropping is allowed. By default, cropping is not allowed. Valid values:
  // 
  // - T: Detection is required.
  // 
  // - F: Detection is required. (Default value: F)
  // 
  // example:
  // 
  // F
  Crop *string `json:"Crop,omitempty" xml:"Crop,omitempty"`
  // The real name of the user. When Authorize=\\"T\\" and the document type is a mainland China ID card, either the key document information (DocName, DocNo) or the document image (IdOcrPictureBase64/URL) must be provided.
  // 
  // Note: Supports a combination of Chinese characters with a minimum length of one character. No special characters are allowed, except for the middle dot (·) used in ethnic minority names.
  // 
  // example:
  // 
  // Zhang**
  DocName *string `json:"DocName,omitempty" xml:"DocName,omitempty"`
  // The document number of the user. When Authorize=\\"T\\" and the document type is a mainland China ID card, either the key document information (DocName, DocNo) or the document image (IdOcrPictureBase64/URL) must be provided.
  // 
  // Note: Supports a combination of letters and numbers with a length of 18 characters.
  // 
  // example:
  // 
  // 410***************
  DocNo *string `json:"DocNo,omitempty" xml:"DocNo,omitempty"`
  // The document type.
  // 
  // example:
  // 
  // 00000001
  DocType *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
  // The Base64-encoded face image.
  // 
  // Note:
  // 
  // - If you choose this method to pass in the face image, check the photo size and do not pass in an excessively large photo.
  // 
  // - Specify either FacePictureBase64 or FacePictureUrl.
  // 
  // example:
  // 
  // Base64
  FacePictureBase64 *string `json:"FacePictureBase64,omitempty" xml:"FacePictureBase64,omitempty"`
  // The URL of the face photo.
  // 
  // example:
  // 
  // https://digital-face-prod8.oss-cn-hangzhou.aliyuncs.com/1669520556530-expo/default/face/20221127114236530_w3kx2e6t.jpg
  FacePictureUrl *string `json:"FacePictureUrl,omitempty" xml:"FacePictureUrl,omitempty"`
  // Specifies whether to enable face quality detection.
  // 
  // example:
  // 
  // Y
  FaceQualityCheck *string `json:"FaceQualityCheck,omitempty" xml:"FaceQualityCheck,omitempty"`
  // The Base64-encoded document image.
  // 
  // Note:
  // 
  // - If you choose this method to pass in the document image, check the photo size and do not pass in an excessively large photo.
  // 
  // - Specify either IdOcrPictureBase64 or IdOcrPictureUrl.
  // 
  // example:
  // 
  // base64
  IdOcrPictureBase64 *string `json:"IdOcrPictureBase64,omitempty" xml:"IdOcrPictureBase64,omitempty"`
  // The URL of the front side of the document image.
  // 
  // example:
  // 
  // https://digital-cardocr-prod8.oss-cn-hangzhou.aliyuncs.com/1669520556530-expo/default/face/20221127114236530_w3kx2e6t.jpg
  IdOcrPictureUrl *string `json:"IdOcrPictureUrl,omitempty" xml:"IdOcrPictureUrl,omitempty"`
  // Specifies whether to enable document anti-spoofing.
  // 
  // example:
  // 
  // Y
  IdSpoof *string `json:"IdSpoof,omitempty" xml:"IdSpoof,omitempty"`
  // The custom OCR quality detection threshold mode. Valid values:
  // 
  // - 0: system default
  // 
  // - 1: strict mode
  // 
  // - 2: loose mode
  // 
  // - 3 (default): disable quality detection
  // 
  // example:
  // 
  // 0
  IdThreshold *string `json:"IdThreshold,omitempty" xml:"IdThreshold,omitempty"`
  // A custom business unique identifier defined by the merchant, used for subsequent issue tracking and troubleshooting. Supports a combination of letters and numbers up to 32 characters in length. Ensure that this value is unique.
  // 
  // example:
  // 
  // e0c34a77f5ac40a5aa5e6ed20c353888
  MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
  // A custom user ID or other identifier that can identify a specific user, such as a phone number or email address. We strongly recommend that you desensitize this field value in advance, for example, by hashing the value.
  // 
  // example:
  // 
  // 123456
  MerchantUserId *string `json:"MerchantUserId,omitempty" xml:"MerchantUserId,omitempty"`
  // The product code.
  // 
  // example:
  // 
  // eKYC_MIN
  ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s EkycVerifyRequest) String() string {
  return dara.Prettify(s)
}

func (s EkycVerifyRequest) GoString() string {
  return s.String()
}

func (s *EkycVerifyRequest) GetAuthorize() *string  {
  return s.Authorize
}

func (s *EkycVerifyRequest) GetCrop() *string  {
  return s.Crop
}

func (s *EkycVerifyRequest) GetDocName() *string  {
  return s.DocName
}

func (s *EkycVerifyRequest) GetDocNo() *string  {
  return s.DocNo
}

func (s *EkycVerifyRequest) GetDocType() *string  {
  return s.DocType
}

func (s *EkycVerifyRequest) GetFacePictureBase64() *string  {
  return s.FacePictureBase64
}

func (s *EkycVerifyRequest) GetFacePictureUrl() *string  {
  return s.FacePictureUrl
}

func (s *EkycVerifyRequest) GetFaceQualityCheck() *string  {
  return s.FaceQualityCheck
}

func (s *EkycVerifyRequest) GetIdOcrPictureBase64() *string  {
  return s.IdOcrPictureBase64
}

func (s *EkycVerifyRequest) GetIdOcrPictureUrl() *string  {
  return s.IdOcrPictureUrl
}

func (s *EkycVerifyRequest) GetIdSpoof() *string  {
  return s.IdSpoof
}

func (s *EkycVerifyRequest) GetIdThreshold() *string  {
  return s.IdThreshold
}

func (s *EkycVerifyRequest) GetMerchantBizId() *string  {
  return s.MerchantBizId
}

func (s *EkycVerifyRequest) GetMerchantUserId() *string  {
  return s.MerchantUserId
}

func (s *EkycVerifyRequest) GetProductCode() *string  {
  return s.ProductCode
}

func (s *EkycVerifyRequest) SetAuthorize(v string) *EkycVerifyRequest {
  s.Authorize = &v
  return s
}

func (s *EkycVerifyRequest) SetCrop(v string) *EkycVerifyRequest {
  s.Crop = &v
  return s
}

func (s *EkycVerifyRequest) SetDocName(v string) *EkycVerifyRequest {
  s.DocName = &v
  return s
}

func (s *EkycVerifyRequest) SetDocNo(v string) *EkycVerifyRequest {
  s.DocNo = &v
  return s
}

func (s *EkycVerifyRequest) SetDocType(v string) *EkycVerifyRequest {
  s.DocType = &v
  return s
}

func (s *EkycVerifyRequest) SetFacePictureBase64(v string) *EkycVerifyRequest {
  s.FacePictureBase64 = &v
  return s
}

func (s *EkycVerifyRequest) SetFacePictureUrl(v string) *EkycVerifyRequest {
  s.FacePictureUrl = &v
  return s
}

func (s *EkycVerifyRequest) SetFaceQualityCheck(v string) *EkycVerifyRequest {
  s.FaceQualityCheck = &v
  return s
}

func (s *EkycVerifyRequest) SetIdOcrPictureBase64(v string) *EkycVerifyRequest {
  s.IdOcrPictureBase64 = &v
  return s
}

func (s *EkycVerifyRequest) SetIdOcrPictureUrl(v string) *EkycVerifyRequest {
  s.IdOcrPictureUrl = &v
  return s
}

func (s *EkycVerifyRequest) SetIdSpoof(v string) *EkycVerifyRequest {
  s.IdSpoof = &v
  return s
}

func (s *EkycVerifyRequest) SetIdThreshold(v string) *EkycVerifyRequest {
  s.IdThreshold = &v
  return s
}

func (s *EkycVerifyRequest) SetMerchantBizId(v string) *EkycVerifyRequest {
  s.MerchantBizId = &v
  return s
}

func (s *EkycVerifyRequest) SetMerchantUserId(v string) *EkycVerifyRequest {
  s.MerchantUserId = &v
  return s
}

func (s *EkycVerifyRequest) SetProductCode(v string) *EkycVerifyRequest {
  s.ProductCode = &v
  return s
}

func (s *EkycVerifyRequest) Validate() error {
  return dara.Validate(s)
}

