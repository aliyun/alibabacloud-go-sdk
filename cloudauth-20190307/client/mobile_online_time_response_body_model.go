// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMobileOnlineTimeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *MobileOnlineTimeResponseBody
	GetCode() *string
	SetMessage(v string) *MobileOnlineTimeResponseBody
	GetMessage() *string
	SetRequestId(v string) *MobileOnlineTimeResponseBody
	GetRequestId() *string
	SetResultObject(v *MobileOnlineTimeResponseBodyResultObject) *MobileOnlineTimeResponseBody
	GetResultObject() *MobileOnlineTimeResponseBodyResultObject
}

type MobileOnlineTimeResponseBody struct {
	// Return code: 200 for success, others for failure.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Return message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// B506328A-D84B-4750-82C7-6A207C585CF1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returned result information
	ResultObject *MobileOnlineTimeResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s MobileOnlineTimeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MobileOnlineTimeResponseBody) GoString() string {
	return s.String()
}

func (s *MobileOnlineTimeResponseBody) GetCode() *string {
	return s.Code
}

func (s *MobileOnlineTimeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *MobileOnlineTimeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MobileOnlineTimeResponseBody) GetResultObject() *MobileOnlineTimeResponseBodyResultObject {
	return s.ResultObject
}

func (s *MobileOnlineTimeResponseBody) SetCode(v string) *MobileOnlineTimeResponseBody {
	s.Code = &v
	return s
}

func (s *MobileOnlineTimeResponseBody) SetMessage(v string) *MobileOnlineTimeResponseBody {
	s.Message = &v
	return s
}

func (s *MobileOnlineTimeResponseBody) SetRequestId(v string) *MobileOnlineTimeResponseBody {
	s.RequestId = &v
	return s
}

func (s *MobileOnlineTimeResponseBody) SetResultObject(v *MobileOnlineTimeResponseBodyResultObject) *MobileOnlineTimeResponseBody {
	s.ResultObject = v
	return s
}

func (s *MobileOnlineTimeResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MobileOnlineTimeResponseBodyResultObject struct {
	// Verification result code.
	//
	// - 1: Verification consistent
	//
	// - 2: Verification inconsistent
	//
	// - 3: No record found
	//
	// example:
	//
	// 1
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// Operator name
	//
	// - CMCC: China Mobile
	//
	// - CUCC: China Unicom
	//
	// - CTCC: China Telecom
	//
	// example:
	//
	// CMCC
	IspName *string `json:"IspName,omitempty" xml:"IspName,omitempty"`
	// - 1: [0,3) indicates the online duration is 0~3 months
	//
	// - 2: [3,6) indicates the online duration is 3~6 months
	//
	// - 3: [6,12) indicates the online duration is 6~12 months
	//
	// - 4: [12,24) indicates the online duration is 12~24 months
	//
	// - 5: [24,+) indicates the online duration is more than 24 months
	//
	// example:
	//
	// 5
	TimeCode *string `json:"TimeCode,omitempty" xml:"TimeCode,omitempty"`
}

func (s MobileOnlineTimeResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s MobileOnlineTimeResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *MobileOnlineTimeResponseBodyResultObject) GetBizCode() *string {
	return s.BizCode
}

func (s *MobileOnlineTimeResponseBodyResultObject) GetIspName() *string {
	return s.IspName
}

func (s *MobileOnlineTimeResponseBodyResultObject) GetTimeCode() *string {
	return s.TimeCode
}

func (s *MobileOnlineTimeResponseBodyResultObject) SetBizCode(v string) *MobileOnlineTimeResponseBodyResultObject {
	s.BizCode = &v
	return s
}

func (s *MobileOnlineTimeResponseBodyResultObject) SetIspName(v string) *MobileOnlineTimeResponseBodyResultObject {
	s.IspName = &v
	return s
}

func (s *MobileOnlineTimeResponseBodyResultObject) SetTimeCode(v string) *MobileOnlineTimeResponseBodyResultObject {
	s.TimeCode = &v
	return s
}

func (s *MobileOnlineTimeResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
