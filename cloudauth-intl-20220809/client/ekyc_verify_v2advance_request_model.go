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
  // Specifies whether to enable authoritative identity verification. Currently, this feature is applicable only to second-generation ID cards of mainland China.
  // 
  // example:
  // 
  // T
  Authorize *string `json:"Authorize,omitempty" xml:"Authorize,omitempty"`
  // Specifies whether cropping is allowed. Not allowed by default. Valid values: T and F.
  // 
  // - T: Cropping is allowed.
  // 
  // - F: Cropping is not allowed. (Default: F)
  // 
  // example:
  // 
  // F
  Crop *string `json:"Crop,omitempty" xml:"Crop,omitempty"`
  // The real name of the user. When Authorize=\\"T\\" and the document type is a mainland China ID card, at least one of the following groups must be provided: document key information (DocName, DocNo) or document image (IdOcrPictureBase64/URL). Note: Supports combinations of Chinese characters with a length of at least 1 character. Special characters are not supported, except for the middle dot (·) used in ethnic minority names.
  // 
  // example:
  // 
  // 张**
  DocName *string `json:"DocName,omitempty" xml:"DocName,omitempty"`
  // The document number of the user. When Authorize=\\"T\\" and the document type is a mainland China ID card, at least one of the following groups must be provided: document key information (DocName, DocNo) or document image (IdOcrPictureBase64/URL). Note: Supports a combination of letters and digits with a length of 18 characters.
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
  // - If you use this method to pass the face image, check the photo size and do not pass an overly large photo.
  // 
  // - You can only specify one of FacePictureBase64, FacePictureUrl, and FacePictureFile.
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
  FaceQualityCheck *string `json:"FaceQualityCheck,omitempty" xml:"FaceQualityCheck,omitempty"`
  // The Base64-encoded document image. Note:
  // 
  // - If you use this method to pass the document image, check the photo size and do not pass an overly large photo.
  // 
  // - You can only specify one of IdOcrPictureBase64, IdOcrPictureUrl, and IdOcrPictureFile.
  // 
  // example:
  // 
  // base64
  IdOcrPictureBase64 *string `json:"IdOcrPictureBase64,omitempty" xml:"IdOcrPictureBase64,omitempty"`
  // The file stream of the front side of the document image.
  // 
  // example:
  // 
  // InputStream
  IdOcrPictureFileObject io.Reader `json:"IdOcrPictureFile,omitempty" xml:"IdOcrPictureFile,omitempty"`
  // The URL of the front side of the document image.
  // 
  // example:
  // 
  // https://digital-cardocr-prod8.oss-cn-hangzhou.aliyuncs.com/1669520556530-expo/default/face/20221127114236530_w3kx2e6t.jpg
  IdOcrPictureUrl *string `json:"IdOcrPictureUrl,omitempty" xml:"IdOcrPictureUrl,omitempty"`
  // The custom OCR quality detection threshold mode:
  // 
  // - 0: System default.
  // 
  // - 1: Strict mode.
  // 
  // - 2: Lenient mode.
  // 
  // - 3 (Default): Quality detection is disabled.
  // 
  // example:
  // 
  // 0
  IdThreshold *string `json:"IdThreshold,omitempty" xml:"IdThreshold,omitempty"`
  // A unique business identifier customized by the merchant, used for subsequent troubleshooting. Supports a combination of letters and digits with a length of 32 characters. Ensure that the value is unique.
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

