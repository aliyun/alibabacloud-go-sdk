// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEkycVerifyResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EkycVerifyResponseBody
  GetCode() *string 
  SetMessage(v string) *EkycVerifyResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EkycVerifyResponseBody
  GetRequestId() *string 
  SetResult(v *EkycVerifyResponseBodyResult) *EkycVerifyResponseBody
  GetResult() *EkycVerifyResponseBodyResult 
}

type EkycVerifyResponseBody struct {
  // The [response code](https://www.alibabacloud.com/help/en/ekyc/latest/im1u641gyesiqmbg?spm=a2c63.p38356.0.i18#GiGmf).
  // 
  // example:
  // 
  // Success
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  // A detailed description of the response code.
  // 
  // example:
  // 
  // success
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // Id of the request
  // 
  // example:
  // 
  // 4EB356FE-BB6A-5DCC-B4C5-E8051787EBA1
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // Result object
  Result *EkycVerifyResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s EkycVerifyResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EkycVerifyResponseBody) GoString() string {
  return s.String()
}

func (s *EkycVerifyResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EkycVerifyResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EkycVerifyResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EkycVerifyResponseBody) GetResult() *EkycVerifyResponseBodyResult  {
  return s.Result
}

func (s *EkycVerifyResponseBody) SetCode(v string) *EkycVerifyResponseBody {
  s.Code = &v
  return s
}

func (s *EkycVerifyResponseBody) SetMessage(v string) *EkycVerifyResponseBody {
  s.Message = &v
  return s
}

func (s *EkycVerifyResponseBody) SetRequestId(v string) *EkycVerifyResponseBody {
  s.RequestId = &v
  return s
}

func (s *EkycVerifyResponseBody) SetResult(v *EkycVerifyResponseBodyResult) *EkycVerifyResponseBody {
  s.Result = v
  return s
}

func (s *EkycVerifyResponseBody) Validate() error {
  if s.Result != nil {
    if err := s.Result.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type EkycVerifyResponseBodyResult struct {
  // Information about the face liveness verification result. For the JSON format, see the example on the right. For more information, see [ExtFaceInfo](https://www.alibabacloud.com/help/en/ekyc/latest/im1u641gyesiqmbg?spm=a2c63.p38356.0.i18#JJ40j).
  // 
  // example:
  // 
  // {
  // 
  // "faceAttack": "N",
  // 
  // "faceComparisonScore": 52.57,
  // 
  // "facePassed": "N",
  // 
  // "authorityComparisonScore": 80.39
  // 
  // }
  ExtFaceInfo *string `json:"ExtFaceInfo,omitempty" xml:"ExtFaceInfo,omitempty"`
  // Information about the certificate detection result.
  // 
  // For the JSON format, see the example on the right. For more information, see [ExtIdInfo](https://www.alibabacloud.com/help/en/ekyc/latest/im1u641gyesiqmbg?spm=a2c63.p38356.0.i18#iWOBY).
  // 
  // example:
  // 
  // {
  // 
  //   "ocrIdInfo": {
  // 
  //     "expiryDate": "",
  // 
  //     "originOfIssue": "Exit and Entry Administration of the Ministry of Public Security",
  // 
  //     "englishName": "LI SI",
  // 
  //     "sex": "Male",
  // 
  //     "name": "Li Si",
  // 
  //     "idNumber": "H11111112",
  // 
  //     "issueDate": "2013-01-02",
  // 
  //     "birthDate": "1990-02-21"
  // 
  //   },
  // 
  //   "ocrIdPassed": "N",
  // 
  //   "spoofInfo": {
  // 
  //     "spoofResult": "Y",
  // 
  //     "spoofType": [
  // 
  //       "SCREEN_REMARK"
  // 
  //     ]
  // 
  //   }
  // 
  // }
  ExtIdInfo *string `json:"ExtIdInfo,omitempty" xml:"ExtIdInfo,omitempty"`
  // The final authentication result. Valid values:
  // 
  // - **Y**: The authentication is passed.
  // 
  // - **N**: The authentication fails.
  // 
  // example:
  // 
  // Y
  Passed *string `json:"Passed,omitempty" xml:"Passed,omitempty"`
  // A description of the authentication result. For more information, see [Error codes for ResultObject.SubCode](https://www.alibabacloud.com/help/en/ekyc/latest/im1u641gyesiqmbg?spm=a2c63.p38356.0.i18#HCGLb).
  // 
  // example:
  // 
  // 205
  SubCode *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
  // The transaction ID.
  // 
  // example:
  // 
  // 4ab0b***cbde97
  TransactionId *string `json:"TransactionId,omitempty" xml:"TransactionId,omitempty"`
}

func (s EkycVerifyResponseBodyResult) String() string {
  return dara.Prettify(s)
}

func (s EkycVerifyResponseBodyResult) GoString() string {
  return s.String()
}

func (s *EkycVerifyResponseBodyResult) GetExtFaceInfo() *string  {
  return s.ExtFaceInfo
}

func (s *EkycVerifyResponseBodyResult) GetExtIdInfo() *string  {
  return s.ExtIdInfo
}

func (s *EkycVerifyResponseBodyResult) GetPassed() *string  {
  return s.Passed
}

func (s *EkycVerifyResponseBodyResult) GetSubCode() *string  {
  return s.SubCode
}

func (s *EkycVerifyResponseBodyResult) GetTransactionId() *string  {
  return s.TransactionId
}

func (s *EkycVerifyResponseBodyResult) SetExtFaceInfo(v string) *EkycVerifyResponseBodyResult {
  s.ExtFaceInfo = &v
  return s
}

func (s *EkycVerifyResponseBodyResult) SetExtIdInfo(v string) *EkycVerifyResponseBodyResult {
  s.ExtIdInfo = &v
  return s
}

func (s *EkycVerifyResponseBodyResult) SetPassed(v string) *EkycVerifyResponseBodyResult {
  s.Passed = &v
  return s
}

func (s *EkycVerifyResponseBodyResult) SetSubCode(v string) *EkycVerifyResponseBodyResult {
  s.SubCode = &v
  return s
}

func (s *EkycVerifyResponseBodyResult) SetTransactionId(v string) *EkycVerifyResponseBodyResult {
  s.TransactionId = &v
  return s
}

func (s *EkycVerifyResponseBodyResult) Validate() error {
  return dara.Validate(s)
}

