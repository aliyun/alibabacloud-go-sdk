// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEkycVerifyV2Request interface {
  dara.Model
  String() string
  GoString() string
  SetAuthorize(v string) *EkycVerifyV2Request
  GetAuthorize() *string 
  SetCrop(v string) *EkycVerifyV2Request
  GetCrop() *string 
  SetDocName(v string) *EkycVerifyV2Request
  GetDocName() *string 
  SetDocNo(v string) *EkycVerifyV2Request
  GetDocNo() *string 
  SetDocType(v string) *EkycVerifyV2Request
  GetDocType() *string 
  SetFacePictureBase64(v string) *EkycVerifyV2Request
  GetFacePictureBase64() *string 
  SetFacePictureFile(v string) *EkycVerifyV2Request
  GetFacePictureFile() *string 
  SetFacePictureUrl(v string) *EkycVerifyV2Request
  GetFacePictureUrl() *string 
  SetFaceQualityCheck(v string) *EkycVerifyV2Request
  GetFaceQualityCheck() *string 
  SetIdOcrPictureBase64(v string) *EkycVerifyV2Request
  GetIdOcrPictureBase64() *string 
  SetIdOcrPictureFile(v string) *EkycVerifyV2Request
  GetIdOcrPictureFile() *string 
  SetIdOcrPictureUrl(v string) *EkycVerifyV2Request
  GetIdOcrPictureUrl() *string 
  SetIdSpoof(v string) *EkycVerifyV2Request
  GetIdSpoof() *string 
  SetIdThreshold(v string) *EkycVerifyV2Request
  GetIdThreshold() *string 
  SetMerchantBizId(v string) *EkycVerifyV2Request
  GetMerchantBizId() *string 
  SetMerchantUserId(v string) *EkycVerifyV2Request
  GetMerchantUserId() *string 
  SetProductCode(v string) *EkycVerifyV2Request
  GetProductCode() *string 
}

type EkycVerifyV2Request struct {
  // Specifies whether to enable authoritative identity verification. Currently, this parameter applies only to second-generation mainland China ID cards.
  // 
  // example:
  // 
  // T
  Authorize *string `json:"Authorize,omitempty" xml:"Authorize,omitempty"`
  // Specifies whether cropping is allowed. By default, cropping is not allowed. Valid values:
  // 
  // - T: Detection is required.
  // 
  // - F: Detection is required (default value: F).
  // 
  // example:
  // 
  // F
  Crop *string `json:"Crop,omitempty" xml:"Crop,omitempty"`
  // The real name of the user. When Authorize=\\"T\\" and the document type is a mainland China ID card, you must provide at least one of the following: key document information (DocName, DocNo) or document images (IdOcrPictureBase64/URL). Note: Supports a combination of Chinese characters with a minimum length of 1 character. No special characters are allowed, except for the middle dot (·) used in ethnic minority names.
  // 
  // example:
  // 
  // Zhang**
  DocName *string `json:"DocName,omitempty" xml:"DocName,omitempty"`
  // The document number of the user. When Authorize=\\"T\\" and the document type is a mainland China ID card, you must provide at least one of the following: key document information (DocName, DocNo) or document images (IdOcrPictureBase64/URL). Note: Supports a combination of letters and numbers with a length of 18 characters.
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
  // - Specify one of the following parameters: FacePictureBase64, FacePictureUrl, or FacePictureFile.
  // 
  // example:
  // 
  // Base64
  FacePictureBase64 *string `json:"FacePictureBase64,omitempty" xml:"FacePictureBase64,omitempty"`
  // The file stream of the face photo.
  // 
  // example:
  // 
  // InputStream
  FacePictureFile *string `json:"FacePictureFile,omitempty" xml:"FacePictureFile,omitempty"`
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
  // The Base64-encoded identity document image. Note:
  // 
  // - If you choose this method to pass in the document image, check the photo size and do not pass in an excessively large photo.
  // 
  // - Specify one of the following parameters: IdOcrPictureBase64, IdOcrPictureUrl, or IdOcrPictureFile.
  // 
  // example:
  // 
  // base64
  IdOcrPictureBase64 *string `json:"IdOcrPictureBase64,omitempty" xml:"IdOcrPictureBase64,omitempty"`
  // The file stream of the front side of the identity document image.
  // 
  // example:
  // 
  // InputStream
  IdOcrPictureFile *string `json:"IdOcrPictureFile,omitempty" xml:"IdOcrPictureFile,omitempty"`
  // The URL of the front side of the identity document image.
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
  // - 0: system default.
  // 
  // - 1: strict mode.
  // 
  // - 2: loose mode.
  // 
  // - 3 (default): quality detection disabled.
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
  // A custom user ID or other identifier that can identify a specific user, such as a phone number or email address. We strongly recommend that you desensitize the value of this field in advance, for example, by hashing the value.
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

func (s EkycVerifyV2Request) String() string {
  return dara.Prettify(s)
}

func (s EkycVerifyV2Request) GoString() string {
  return s.String()
}

func (s *EkycVerifyV2Request) GetAuthorize() *string  {
  return s.Authorize
}

func (s *EkycVerifyV2Request) GetCrop() *string  {
  return s.Crop
}

func (s *EkycVerifyV2Request) GetDocName() *string  {
  return s.DocName
}

func (s *EkycVerifyV2Request) GetDocNo() *string  {
  return s.DocNo
}

func (s *EkycVerifyV2Request) GetDocType() *string  {
  return s.DocType
}

func (s *EkycVerifyV2Request) GetFacePictureBase64() *string  {
  return s.FacePictureBase64
}

func (s *EkycVerifyV2Request) GetFacePictureFile() *string  {
  return s.FacePictureFile
}

func (s *EkycVerifyV2Request) GetFacePictureUrl() *string  {
  return s.FacePictureUrl
}

func (s *EkycVerifyV2Request) GetFaceQualityCheck() *string  {
  return s.FaceQualityCheck
}

func (s *EkycVerifyV2Request) GetIdOcrPictureBase64() *string  {
  return s.IdOcrPictureBase64
}

func (s *EkycVerifyV2Request) GetIdOcrPictureFile() *string  {
  return s.IdOcrPictureFile
}

func (s *EkycVerifyV2Request) GetIdOcrPictureUrl() *string  {
  return s.IdOcrPictureUrl
}

func (s *EkycVerifyV2Request) GetIdSpoof() *string  {
  return s.IdSpoof
}

func (s *EkycVerifyV2Request) GetIdThreshold() *string  {
  return s.IdThreshold
}

func (s *EkycVerifyV2Request) GetMerchantBizId() *string  {
  return s.MerchantBizId
}

func (s *EkycVerifyV2Request) GetMerchantUserId() *string  {
  return s.MerchantUserId
}

func (s *EkycVerifyV2Request) GetProductCode() *string  {
  return s.ProductCode
}

func (s *EkycVerifyV2Request) SetAuthorize(v string) *EkycVerifyV2Request {
  s.Authorize = &v
  return s
}

func (s *EkycVerifyV2Request) SetCrop(v string) *EkycVerifyV2Request {
  s.Crop = &v
  return s
}

func (s *EkycVerifyV2Request) SetDocName(v string) *EkycVerifyV2Request {
  s.DocName = &v
  return s
}

func (s *EkycVerifyV2Request) SetDocNo(v string) *EkycVerifyV2Request {
  s.DocNo = &v
  return s
}

func (s *EkycVerifyV2Request) SetDocType(v string) *EkycVerifyV2Request {
  s.DocType = &v
  return s
}

func (s *EkycVerifyV2Request) SetFacePictureBase64(v string) *EkycVerifyV2Request {
  s.FacePictureBase64 = &v
  return s
}

func (s *EkycVerifyV2Request) SetFacePictureFile(v string) *EkycVerifyV2Request {
  s.FacePictureFile = &v
  return s
}

func (s *EkycVerifyV2Request) SetFacePictureUrl(v string) *EkycVerifyV2Request {
  s.FacePictureUrl = &v
  return s
}

func (s *EkycVerifyV2Request) SetFaceQualityCheck(v string) *EkycVerifyV2Request {
  s.FaceQualityCheck = &v
  return s
}

func (s *EkycVerifyV2Request) SetIdOcrPictureBase64(v string) *EkycVerifyV2Request {
  s.IdOcrPictureBase64 = &v
  return s
}

func (s *EkycVerifyV2Request) SetIdOcrPictureFile(v string) *EkycVerifyV2Request {
  s.IdOcrPictureFile = &v
  return s
}

func (s *EkycVerifyV2Request) SetIdOcrPictureUrl(v string) *EkycVerifyV2Request {
  s.IdOcrPictureUrl = &v
  return s
}

func (s *EkycVerifyV2Request) SetIdSpoof(v string) *EkycVerifyV2Request {
  s.IdSpoof = &v
  return s
}

func (s *EkycVerifyV2Request) SetIdThreshold(v string) *EkycVerifyV2Request {
  s.IdThreshold = &v
  return s
}

func (s *EkycVerifyV2Request) SetMerchantBizId(v string) *EkycVerifyV2Request {
  s.MerchantBizId = &v
  return s
}

func (s *EkycVerifyV2Request) SetMerchantUserId(v string) *EkycVerifyV2Request {
  s.MerchantUserId = &v
  return s
}

func (s *EkycVerifyV2Request) SetProductCode(v string) *EkycVerifyV2Request {
  s.ProductCode = &v
  return s
}

func (s *EkycVerifyV2Request) Validate() error {
  return dara.Validate(s)
}

