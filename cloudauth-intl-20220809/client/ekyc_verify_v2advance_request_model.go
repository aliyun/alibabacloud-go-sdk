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
  // Indicates whether to enable authoritative identity verification. This parameter currently applies only to second-generation ID cards issued in the Chinese mainland.
  // 
  // example:
  // 
  // T
  Authorize *string `json:"Authorize,omitempty" xml:"Authorize,omitempty"`
  // Indicates whether clipping is allowed. Clipping is disabled by default (T/F).
  // 
  // - T: Detection is required.
  // 
  // - F: Detection is required (default is F).
  // 
  // example:
  // 
  // F
  Crop *string `json:"Crop,omitempty" xml:"Crop,omitempty"`
  // The user\\"s real name. When Authorize=\\"T\\" and the certificate type is a Chinese mainland ID card, you must provide at least one of the following: key certificate information (DocName, DocNo) or certificate image (IdOcrPictureBase64/URL). Note: It supports combinations of one or more Chinese characters, excluding special characters except for the interpunct 【·】 used in ethnic minority names.
  // 
  // example:
  // 
  // 张**
  DocName *string `json:"DocName,omitempty" xml:"DocName,omitempty"`
  // The user\\"s certificate number. When Authorize=\\"T\\" and the certificate type is a Chinese mainland ID card, you must provide at least one of the following: key certificate information (DocName, DocNo) or certificate image (IdOcrPictureBase64/URL). Note: It supports a combination of letters and digits with a length of 18 characters.
  // 
  // example:
  // 
  // 410***************
  DocNo *string `json:"DocNo,omitempty" xml:"DocNo,omitempty"`
  // Certificate type
  // 
  // example:
  // 
  // 00000001
  DocType *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
  // Base64 encoding of the facial image.
  // 
  // Notes:
  // 
  // - If you choose this method to submit the certificate image, check the image size and avoid uploading excessively large images.
  // 
  // - You must specify exactly one of FacePictureBase64, FacePictureUrl, or FacePictureFile.
  // 
  // example:
  // 
  // Base64
  FacePictureBase64 *string `json:"FacePictureBase64,omitempty" xml:"FacePictureBase64,omitempty"`
  // File stream of the facial photo
  // 
  // example:
  // 
  // InputStream
  FacePictureFileObject io.Reader `json:"FacePictureFile,omitempty" xml:"FacePictureFile,omitempty"`
  // URL of the facial photo
  // 
  // example:
  // 
  // https://digital-face-prod8.oss-cn-hangzhou.aliyuncs.com/1669520556530-expo/default/face/20221127114236530_w3kx2e6t.jpg
  FacePictureUrl *string `json:"FacePictureUrl,omitempty" xml:"FacePictureUrl,omitempty"`
  // Base64-encoded certificate Image. Notes:
  // 
  // - If you use this method to submit the certificate image, check the image size and avoid uploading excessively large images.
  // 
  // - You must specify exactly one of IdOcrPictureBase64, IdOcrPictureUrl, or IdOcrPictureFile.
  // 
  // example:
  // 
  // base64
  IdOcrPictureBase64 *string `json:"IdOcrPictureBase64,omitempty" xml:"IdOcrPictureBase64,omitempty"`
  // File stream of the front side of the certificate
  // 
  // example:
  // 
  // InputStream
  IdOcrPictureFileObject io.Reader `json:"IdOcrPictureFile,omitempty" xml:"IdOcrPictureFile,omitempty"`
  // URL of the front side of the certificate
  // 
  // example:
  // 
  // https://digital-cardocr-prod8.oss-cn-hangzhou.aliyuncs.com/1669520556530-expo/default/face/20221127114236530_w3kx2e6t.jpg
  IdOcrPictureUrl *string `json:"IdOcrPictureUrl,omitempty" xml:"IdOcrPictureUrl,omitempty"`
  // Custom OCR quality detection threshold mode:
  // 
  // - 0: System default
  // 
  // - 1: Strict mode
  // 
  // - 2: Loose mode
  // 
  // - 3 (default): Shutdown quality detection
  // 
  // example:
  // 
  // 0
  IdThreshold *string `json:"IdThreshold,omitempty" xml:"IdThreshold,omitempty"`
  // A custom business UUID defined by the merchant, used for subsequent issue tracking and troubleshooting. It supports a combination of letters and digits with a length of 32 characters. Ensure its uniqueness.
  // 
  // example:
  // 
  // e0c34a77f5ac40a5aa5e6ed20c353888
  MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
  // Your custom user ID or another identifier that can uniquely identify a specific user, such as a mobile phone number or mailbox address. We strongly recommend pre-masking the value of this field—for example, by applying a hash function.
  // 
  // example:
  // 
  // 123456
  MerchantUserId *string `json:"MerchantUserId,omitempty" xml:"MerchantUserId,omitempty"`
  // Product code
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

