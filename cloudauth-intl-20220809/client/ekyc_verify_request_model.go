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
  SetIdOcrPictureBase64(v string) *EkycVerifyRequest
  GetIdOcrPictureBase64() *string 
  SetIdOcrPictureUrl(v string) *EkycVerifyRequest
  GetIdOcrPictureUrl() *string 
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
  // Specifies whether to enable authoritative identity verification. This parameter is currently applicable only to second-generation ID cards in the Chinese mainland.
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
  // The real name of the user. If Authorize is set to T and the document type is a Chinese mainland ID card, you must provide at least one of the following: document key information (DocName and DocNo) or document images (IdOcrPictureBase64 or IdOcrPictureUrl).
  // 
  // Note: The value must contain at least one Chinese character and no special characters, except for the middle dot (·) used in ethnic minority names.
  // 
  // example:
  // 
  // 张**
  DocName *string `json:"DocName,omitempty" xml:"DocName,omitempty"`
  // The document number of the user. If Authorize is set to T and the document type is a Chinese mainland ID card, you must provide at least one of the following: document key information (DocName and DocNo) or document images (IdOcrPictureBase64 or IdOcrPictureUrl).
  // 
  // Note: The value is a combination of letters and digits up to 18 characters in length.
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
  // - If you use this method to submit the face image, check the image size and do not submit an excessively large image.
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
  // The Base64-encoded document image.
  // 
  // Note:
  // 
  // - If you use this method to submit the document image, check the image size and do not submit an excessively large image.
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
  // The custom OCR quality detection threshold mode. Valid values:
  // 
  // - 0: system default
  // 
  // - 1: strict mode
  // 
  // - 2: loose mode
  // 
  // - 3 (default): quality detection disabled.
  // 
  // example:
  // 
  // 0
  IdThreshold *string `json:"IdThreshold,omitempty" xml:"IdThreshold,omitempty"`
  // The merchant-defined unique business identifier, used for subsequent troubleshooting. The value is a combination of letters and digits up to 32 characters in length. Ensure that the value is unique.
  // 
  // example:
  // 
  // e0c34a77f5ac40a5aa5e6ed20c353888
  MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
  // Your custom user ID, or another identifier that can identify a specific user, such as a phone number or email address. We strongly recommend that you mask this field value in advance, for example, by hashing the value.
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

func (s *EkycVerifyRequest) GetIdOcrPictureBase64() *string  {
  return s.IdOcrPictureBase64
}

func (s *EkycVerifyRequest) GetIdOcrPictureUrl() *string  {
  return s.IdOcrPictureUrl
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

func (s *EkycVerifyRequest) SetIdOcrPictureBase64(v string) *EkycVerifyRequest {
  s.IdOcrPictureBase64 = &v
  return s
}

func (s *EkycVerifyRequest) SetIdOcrPictureUrl(v string) *EkycVerifyRequest {
  s.IdOcrPictureUrl = &v
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

