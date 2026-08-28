// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "io"
  "github.com/alibabacloud-go/tea/dara"
)

type iEkycVerifyV2AdvanceRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAuthorize(v string) *EkycVerifyV2AdvanceRequest
  GetAuthorize() *string 
  SetCrop(v string) *EkycVerifyV2AdvanceRequest
  GetCrop() *string 
  SetDocName(v string) *EkycVerifyV2AdvanceRequest
  GetDocName() *string 
  SetDocNo(v string) *EkycVerifyV2AdvanceRequest
  GetDocNo() *string 
  SetDocType(v string) *EkycVerifyV2AdvanceRequest
  GetDocType() *string 
  SetFacePictureBase64(v string) *EkycVerifyV2AdvanceRequest
  GetFacePictureBase64() *string 
  SetFacePictureFileObject(v io.Reader) *EkycVerifyV2AdvanceRequest
  GetFacePictureFileObject() io.Reader 
  SetFacePictureUrl(v string) *EkycVerifyV2AdvanceRequest
  GetFacePictureUrl() *string 
  SetFaceQualityCheck(v string) *EkycVerifyV2AdvanceRequest
  GetFaceQualityCheck() *string 
  SetIdOcrPictureBase64(v string) *EkycVerifyV2AdvanceRequest
  GetIdOcrPictureBase64() *string 
  SetIdOcrPictureFileObject(v io.Reader) *EkycVerifyV2AdvanceRequest
  GetIdOcrPictureFileObject() io.Reader 
  SetIdOcrPictureUrl(v string) *EkycVerifyV2AdvanceRequest
  GetIdOcrPictureUrl() *string 
  SetIdSpoof(v string) *EkycVerifyV2AdvanceRequest
  GetIdSpoof() *string 
  SetIdThreshold(v string) *EkycVerifyV2AdvanceRequest
  GetIdThreshold() *string 
  SetMerchantBizId(v string) *EkycVerifyV2AdvanceRequest
  GetMerchantBizId() *string 
  SetMerchantUserId(v string) *EkycVerifyV2AdvanceRequest
  GetMerchantUserId() *string 
  SetProductCode(v string) *EkycVerifyV2AdvanceRequest
  GetProductCode() *string 
}

type EkycVerifyV2AdvanceRequest struct {
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
  FacePictureFileObject io.Reader `json:"FacePictureFile,omitempty" xml:"FacePictureFile,omitempty"`
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
  IdOcrPictureFileObject io.Reader `json:"IdOcrPictureFile,omitempty" xml:"IdOcrPictureFile,omitempty"`
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

func (s EkycVerifyV2AdvanceRequest) String() string {
  return dara.Prettify(s)
}

func (s EkycVerifyV2AdvanceRequest) GoString() string {
  return s.String()
}

func (s *EkycVerifyV2AdvanceRequest) GetAuthorize() *string  {
  return s.Authorize
}

func (s *EkycVerifyV2AdvanceRequest) GetCrop() *string  {
  return s.Crop
}

func (s *EkycVerifyV2AdvanceRequest) GetDocName() *string  {
  return s.DocName
}

func (s *EkycVerifyV2AdvanceRequest) GetDocNo() *string  {
  return s.DocNo
}

func (s *EkycVerifyV2AdvanceRequest) GetDocType() *string  {
  return s.DocType
}

func (s *EkycVerifyV2AdvanceRequest) GetFacePictureBase64() *string  {
  return s.FacePictureBase64
}

func (s *EkycVerifyV2AdvanceRequest) GetFacePictureFileObject() io.Reader  {
  return s.FacePictureFileObject
}

func (s *EkycVerifyV2AdvanceRequest) GetFacePictureUrl() *string  {
  return s.FacePictureUrl
}

func (s *EkycVerifyV2AdvanceRequest) GetFaceQualityCheck() *string  {
  return s.FaceQualityCheck
}

func (s *EkycVerifyV2AdvanceRequest) GetIdOcrPictureBase64() *string  {
  return s.IdOcrPictureBase64
}

func (s *EkycVerifyV2AdvanceRequest) GetIdOcrPictureFileObject() io.Reader  {
  return s.IdOcrPictureFileObject
}

func (s *EkycVerifyV2AdvanceRequest) GetIdOcrPictureUrl() *string  {
  return s.IdOcrPictureUrl
}

func (s *EkycVerifyV2AdvanceRequest) GetIdSpoof() *string  {
  return s.IdSpoof
}

func (s *EkycVerifyV2AdvanceRequest) GetIdThreshold() *string  {
  return s.IdThreshold
}

func (s *EkycVerifyV2AdvanceRequest) GetMerchantBizId() *string  {
  return s.MerchantBizId
}

func (s *EkycVerifyV2AdvanceRequest) GetMerchantUserId() *string  {
  return s.MerchantUserId
}

func (s *EkycVerifyV2AdvanceRequest) GetProductCode() *string  {
  return s.ProductCode
}

func (s *EkycVerifyV2AdvanceRequest) SetAuthorize(v string) *EkycVerifyV2AdvanceRequest {
  s.Authorize = &v
  return s
}

func (s *EkycVerifyV2AdvanceRequest) SetCrop(v string) *EkycVerifyV2AdvanceRequest {
  s.Crop = &v
  return s
}

func (s *EkycVerifyV2AdvanceRequest) SetDocName(v string) *EkycVerifyV2AdvanceRequest {
  s.DocName = &v
  return s
}

func (s *EkycVerifyV2AdvanceRequest) SetDocNo(v string) *EkycVerifyV2AdvanceRequest {
  s.DocNo = &v
  return s
}

func (s *EkycVerifyV2AdvanceRequest) SetDocType(v string) *EkycVerifyV2AdvanceRequest {
  s.DocType = &v
  return s
}

func (s *EkycVerifyV2AdvanceRequest) SetFacePictureBase64(v string) *EkycVerifyV2AdvanceRequest {
  s.FacePictureBase64 = &v
  return s
}

func (s *EkycVerifyV2AdvanceRequest) SetFacePictureFileObject(v io.Reader) *EkycVerifyV2AdvanceRequest {
  s.FacePictureFileObject = v
  return s
}

func (s *EkycVerifyV2AdvanceRequest) SetFacePictureUrl(v string) *EkycVerifyV2AdvanceRequest {
  s.FacePictureUrl = &v
  return s
}

func (s *EkycVerifyV2AdvanceRequest) SetFaceQualityCheck(v string) *EkycVerifyV2AdvanceRequest {
  s.FaceQualityCheck = &v
  return s
}

func (s *EkycVerifyV2AdvanceRequest) SetIdOcrPictureBase64(v string) *EkycVerifyV2AdvanceRequest {
  s.IdOcrPictureBase64 = &v
  return s
}

func (s *EkycVerifyV2AdvanceRequest) SetIdOcrPictureFileObject(v io.Reader) *EkycVerifyV2AdvanceRequest {
  s.IdOcrPictureFileObject = v
  return s
}

func (s *EkycVerifyV2AdvanceRequest) SetIdOcrPictureUrl(v string) *EkycVerifyV2AdvanceRequest {
  s.IdOcrPictureUrl = &v
  return s
}

func (s *EkycVerifyV2AdvanceRequest) SetIdSpoof(v string) *EkycVerifyV2AdvanceRequest {
  s.IdSpoof = &v
  return s
}

func (s *EkycVerifyV2AdvanceRequest) SetIdThreshold(v string) *EkycVerifyV2AdvanceRequest {
  s.IdThreshold = &v
  return s
}

func (s *EkycVerifyV2AdvanceRequest) SetMerchantBizId(v string) *EkycVerifyV2AdvanceRequest {
  s.MerchantBizId = &v
  return s
}

func (s *EkycVerifyV2AdvanceRequest) SetMerchantUserId(v string) *EkycVerifyV2AdvanceRequest {
  s.MerchantUserId = &v
  return s
}

func (s *EkycVerifyV2AdvanceRequest) SetProductCode(v string) *EkycVerifyV2AdvanceRequest {
  s.ProductCode = &v
  return s
}

func (s *EkycVerifyV2AdvanceRequest) Validate() error {
  return dara.Validate(s)
}

