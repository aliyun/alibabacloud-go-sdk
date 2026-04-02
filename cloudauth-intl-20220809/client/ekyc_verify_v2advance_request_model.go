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
  // example:
  // 
  // Base64
  FacePictureBase64 *string `json:"FacePictureBase64,omitempty" xml:"FacePictureBase64,omitempty"`
  // example:
  // 
  // InputStream
  FacePictureFileObject io.Reader `json:"FacePictureFile,omitempty" xml:"FacePictureFile,omitempty"`
  // example:
  // 
  // https://digital-face-prod8.oss-cn-hangzhou.aliyuncs.com/1669520556530-expo/default/face/20221127114236530_w3kx2e6t.jpg
  FacePictureUrl *string `json:"FacePictureUrl,omitempty" xml:"FacePictureUrl,omitempty"`
  // example:
  // 
  // base64
  IdOcrPictureBase64 *string `json:"IdOcrPictureBase64,omitempty" xml:"IdOcrPictureBase64,omitempty"`
  // example:
  // 
  // InputStream
  IdOcrPictureFileObject io.Reader `json:"IdOcrPictureFile,omitempty" xml:"IdOcrPictureFile,omitempty"`
  // example:
  // 
  // https://digital-cardocr-prod8.oss-cn-hangzhou.aliyuncs.com/1669520556530-expo/default/face/20221127114236530_w3kx2e6t.jpg
  IdOcrPictureUrl *string `json:"IdOcrPictureUrl,omitempty" xml:"IdOcrPictureUrl,omitempty"`
  // example:
  // 
  // 0
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

