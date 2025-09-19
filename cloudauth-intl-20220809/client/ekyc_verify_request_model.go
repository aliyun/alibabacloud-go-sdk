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
  // Specifies whether to enable identity verification against the official database:
  // 
  // - **T**: Enable.
  // 
  // - **F**: Disable. (Default)
  // 
  // > This feature is currently available only for second-generation resident ID cards of the Chinese mainland.
  // 
  // example:
  // 
  // T
  Authorize *string `json:"Authorize,omitempty" xml:"Authorize,omitempty"`
  // Specifies whether to crop the face image:
  // 
  // - **T**: Allows cropping.
  // 
  // - **F**: Disallows cropping. (Default)
  // 
  // example:
  // 
  // F
  Crop *string `json:"Crop,omitempty" xml:"Crop,omitempty"`
  // The user\\"s real name.
  // 
  // > If Authorize is set to T and the certificate type is Chinese mainland resident ID card, you must enter at least one of the following groups of information:
  // 
  // > - DocName and DocNo.
  // 
  // > - IdOcrPictureBase64 or IdOcrPictureUrl.
  // 
  // example:
  // 
  // Zhang San
  DocName *string `json:"DocName,omitempty" xml:"DocName,omitempty"`
  // The user\\"s certificate number.
  // 
  // 
  // > If Authorize is set to **T*	- and the certificate type is Chinese mainland resident ID card, you must enter at least one of the following groups of information:
  // 
  // > - DocName and DocNo.
  // 
  // > - IdOcrPictureBase64 or IdOcrPictureUrl.
  // 
  // example:
  // 
  // 410***************
  DocNo *string `json:"DocNo,omitempty" xml:"DocNo,omitempty"`
  // The certificate type, which is uniquely identified by an 8-digit number. For more information, see [Certificate types](https://www.alibabacloud.com/help/en/ekyc/latest/im1u641gyesiqmbg?spm=a2c63.p38356.0.i18#Hu5TG).
  // 
  // example:
  // 
  // 00000001
  DocType *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
  FacePictureBase64 *string `json:"FacePictureBase64,omitempty" xml:"FacePictureBase64,omitempty"`
  // The URL of the portrait image. The URL must be an HTTP or HTTPS link accessible over the Internet.
  // 
  // example:
  // 
  // https://digital-face-prod8.oss-cn-hangzhou.aliyuncs.com/1669520556530-expo/default/face/20221127114236530_w3kx2e6t.jpg
  FacePictureUrl *string `json:"FacePictureUrl,omitempty" xml:"FacePictureUrl,omitempty"`
  IdOcrPictureBase64 *string `json:"IdOcrPictureBase64,omitempty" xml:"IdOcrPictureBase64,omitempty"`
  // The URL of the certificate image. The URL must be an HTTP or HTTPS link accessible over the Internet.
  // 
  // example:
  // 
  // https://digital-cardocr-prod8.oss-cn-hangzhou.aliyuncs.com/1669520556530-expo/default/face/20221127114236530_w3kx2e6t.jpg
  IdOcrPictureUrl *string `json:"IdOcrPictureUrl,omitempty" xml:"IdOcrPictureUrl,omitempty"`
  // The custom OCR quality detection threshold mode:
  // 
  // - **0**: Standard mode
  // 
  // - **1**: Strict mode
  // 
  // - **2**: Loose mode
  // 
  // - **3*	- (default): Disables quality detection
  // 
  // example:
  // 
  // 0
  IdThreshold *string `json:"IdThreshold,omitempty" xml:"IdThreshold,omitempty"`
  // A unique business identifier that you customize. It is used to locate and troubleshoot issues. The identifier can be up to 32 characters in length and can contain letters and digits. Make sure that the identifier is unique.
  // 
  // example:
  // 
  // e0c34a77f5ac40a5aa5e6ed20c353888
  MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
  // A custom user ID or another identifier that can identify a specific user, such as a mobile number or an email address. Desensitize the value of this field in advance, for example, by hashing the value.
  // 
  // example:
  // 
  // 123456
  MerchantUserId *string `json:"MerchantUserId,omitempty" xml:"MerchantUserId,omitempty"`
  // The product solution to integrate. Set the value to **eKYC_MIN**.
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

