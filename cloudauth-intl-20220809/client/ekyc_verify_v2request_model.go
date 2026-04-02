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
  SetIdOcrPictureBase64(v string) *EkycVerifyV2Request
  GetIdOcrPictureBase64() *string 
  SetIdOcrPictureFile(v string) *EkycVerifyV2Request
  GetIdOcrPictureFile() *string 
  SetIdOcrPictureUrl(v string) *EkycVerifyV2Request
  GetIdOcrPictureUrl() *string 
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
  FacePictureFile *string `json:"FacePictureFile,omitempty" xml:"FacePictureFile,omitempty"`
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
  IdOcrPictureFile *string `json:"IdOcrPictureFile,omitempty" xml:"IdOcrPictureFile,omitempty"`
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

func (s *EkycVerifyV2Request) GetIdOcrPictureBase64() *string  {
  return s.IdOcrPictureBase64
}

func (s *EkycVerifyV2Request) GetIdOcrPictureFile() *string  {
  return s.IdOcrPictureFile
}

func (s *EkycVerifyV2Request) GetIdOcrPictureUrl() *string  {
  return s.IdOcrPictureUrl
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

