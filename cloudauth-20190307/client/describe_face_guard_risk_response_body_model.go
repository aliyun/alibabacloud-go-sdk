// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFaceGuardRiskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeFaceGuardRiskResponseBody
	GetCode() *string
	SetMessage(v string) *DescribeFaceGuardRiskResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeFaceGuardRiskResponseBody
	GetRequestId() *string
	SetResultObject(v *DescribeFaceGuardRiskResponseBodyResultObject) *DescribeFaceGuardRiskResponseBody
	GetResultObject() *DescribeFaceGuardRiskResponseBodyResultObject
}

type DescribeFaceGuardRiskResponseBody struct {
	// The response code. **200*	- indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// D6163397-15C5-419C-9ACC-B7C83E0B4C10
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result information.
	ResultObject *DescribeFaceGuardRiskResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s DescribeFaceGuardRiskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeFaceGuardRiskResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFaceGuardRiskResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeFaceGuardRiskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeFaceGuardRiskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeFaceGuardRiskResponseBody) GetResultObject() *DescribeFaceGuardRiskResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeFaceGuardRiskResponseBody) SetCode(v string) *DescribeFaceGuardRiskResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeFaceGuardRiskResponseBody) SetMessage(v string) *DescribeFaceGuardRiskResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeFaceGuardRiskResponseBody) SetRequestId(v string) *DescribeFaceGuardRiskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFaceGuardRiskResponseBody) SetResultObject(v *DescribeFaceGuardRiskResponseBodyResultObject) *DescribeFaceGuardRiskResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeFaceGuardRiskResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeFaceGuardRiskResponseBodyResultObject struct {
	// The unique identifier for ID Verification.
	//
	// example:
	//
	// sha43d9cabd52d370d9f4cca9468f71e
	CertifyId *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	// The device risk probability predicted by the device assistant algorithm. A higher score indicates a higher device risk.
	//
	// Value range: 0 to 100.
	//
	// example:
	//
	// 80
	GuardRiskScore *float64 `json:"GuardRiskScore,omitempty" xml:"GuardRiskScore,omitempty"`
	// The extended information in JSON format. The returned content is customized based on tenant requirements.
	//
	// example:
	//
	// {
	//
	//   "code": 200
	//
	//   "badNet":false,
	//
	//   "umid":"74e37355171ab62230063569350d368e",
	//
	//   "fileTags":"basic_root,basic_hook",
	//
	//   "queryCount":1,
	//
	//   "querySessionCount":1,
	//
	//   "queryUmidCount":1
	//
	//   "platform":"Android"
	//
	// }
	RiskExtends *string `json:"RiskExtends,omitempty" xml:"RiskExtends,omitempty"`
	// The device risk labels.
	//
	// - Multiple device risk labels are separated by commas (,), such as "ROOT,VPN,HOOK".
	//
	// - For more information about device risk labels and their descriptions, see the Face Guard label description in the official documentation.
	//
	// example:
	//
	// ROOT,VPN,HOOK
	RiskTags *string `json:"RiskTags,omitempty" xml:"RiskTags,omitempty"`
}

func (s DescribeFaceGuardRiskResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeFaceGuardRiskResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeFaceGuardRiskResponseBodyResultObject) GetCertifyId() *string {
	return s.CertifyId
}

func (s *DescribeFaceGuardRiskResponseBodyResultObject) GetGuardRiskScore() *float64 {
	return s.GuardRiskScore
}

func (s *DescribeFaceGuardRiskResponseBodyResultObject) GetRiskExtends() *string {
	return s.RiskExtends
}

func (s *DescribeFaceGuardRiskResponseBodyResultObject) GetRiskTags() *string {
	return s.RiskTags
}

func (s *DescribeFaceGuardRiskResponseBodyResultObject) SetCertifyId(v string) *DescribeFaceGuardRiskResponseBodyResultObject {
	s.CertifyId = &v
	return s
}

func (s *DescribeFaceGuardRiskResponseBodyResultObject) SetGuardRiskScore(v float64) *DescribeFaceGuardRiskResponseBodyResultObject {
	s.GuardRiskScore = &v
	return s
}

func (s *DescribeFaceGuardRiskResponseBodyResultObject) SetRiskExtends(v string) *DescribeFaceGuardRiskResponseBodyResultObject {
	s.RiskExtends = &v
	return s
}

func (s *DescribeFaceGuardRiskResponseBodyResultObject) SetRiskTags(v string) *DescribeFaceGuardRiskResponseBodyResultObject {
	s.RiskTags = &v
	return s
}

func (s *DescribeFaceGuardRiskResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
