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
  // example:
  // 
  // T
  Authorize *string `json:"Authorize,omitempty" xml:"Authorize,omitempty"`
  // example:
  // 
  // F
  Crop *string `json:"Crop,omitempty" xml:"Crop,omitempty"`
  DocName *string `json:"DocName,omitempty" xml:"DocName,omitempty"`
  // example:
  // 
  // 410***************
  DocNo *string `json:"DocNo,omitempty" xml:"DocNo,omitempty"`
  // example:
  // 
  // 00000001
  DocType *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
  FacePictureBase64 *string `json:"FacePictureBase64,omitempty" xml:"FacePictureBase64,omitempty"`
  // example:
  // 
  // https://digital-face-prod8.oss-cn-hangzhou.aliyuncs.com/1669520556530-expo/default/face/20221127114236530_w3kx2e6t.jpg
  FacePictureUrl *string `json:"FacePictureUrl,omitempty" xml:"FacePictureUrl,omitempty"`
  IdOcrPictureBase64 *string `json:"IdOcrPictureBase64,omitempty" xml:"IdOcrPictureBase64,omitempty"`
  // example:
  // 
  // https://digital-cardocr-prod8.oss-cn-hangzhou.aliyuncs.com/1669520556530-expo/default/face/20221127114236530_w3kx2e6t.jpg
  IdOcrPictureUrl *string `json:"IdOcrPictureUrl,omitempty" xml:"IdOcrPictureUrl,omitempty"`
  IdThreshold *string `json:"IdThreshold,omitempty" xml:"IdThreshold,omitempty"`
  // example:
  // 
  // e0c34a77f5ac40a5aa5e6ed20c353888
  MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
  // example:
  // 
  // 123456
  MerchantUserId *string `json:"MerchantUserId,omitempty" xml:"MerchantUserId,omitempty"`
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

