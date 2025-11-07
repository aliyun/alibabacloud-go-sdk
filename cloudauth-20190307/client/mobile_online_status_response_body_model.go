// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMobileOnlineStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *MobileOnlineStatusResponseBody
	GetCode() *string
	SetMessage(v string) *MobileOnlineStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *MobileOnlineStatusResponseBody
	GetRequestId() *string
	SetResultObject(v *MobileOnlineStatusResponseBodyResultObject) *MobileOnlineStatusResponseBody
	GetResultObject() *MobileOnlineStatusResponseBodyResultObject
}

type MobileOnlineStatusResponseBody struct {
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
	ResultObject *MobileOnlineStatusResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s MobileOnlineStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MobileOnlineStatusResponseBody) GoString() string {
	return s.String()
}

func (s *MobileOnlineStatusResponseBody) GetCode() *string {
	return s.Code
}

func (s *MobileOnlineStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *MobileOnlineStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MobileOnlineStatusResponseBody) GetResultObject() *MobileOnlineStatusResponseBodyResultObject {
	return s.ResultObject
}

func (s *MobileOnlineStatusResponseBody) SetCode(v string) *MobileOnlineStatusResponseBody {
	s.Code = &v
	return s
}

func (s *MobileOnlineStatusResponseBody) SetMessage(v string) *MobileOnlineStatusResponseBody {
	s.Message = &v
	return s
}

func (s *MobileOnlineStatusResponseBody) SetRequestId(v string) *MobileOnlineStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *MobileOnlineStatusResponseBody) SetResultObject(v *MobileOnlineStatusResponseBodyResultObject) *MobileOnlineStatusResponseBody {
	s.ResultObject = v
	return s
}

func (s *MobileOnlineStatusResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MobileOnlineStatusResponseBodyResultObject struct {
	// Verification result
	//
	// - 1: Available online
	//
	// - 2: Not available online (see subCode for details)
	//
	// - 3: No query result
	//
	// example:
	//
	// 1
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// ISP name
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
	// Verification details
	//
	// - 101: Available online
	//
	// - 201: Suspended
	//
	// - 202: Disconnected
	//
	// - 203: Online but not available
	//
	// - 204: Not online
	//
	// - 301: No record found
	//
	// example:
	//
	// 101
	SubCode *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
}

func (s MobileOnlineStatusResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s MobileOnlineStatusResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *MobileOnlineStatusResponseBodyResultObject) GetBizCode() *string {
	return s.BizCode
}

func (s *MobileOnlineStatusResponseBodyResultObject) GetIspName() *string {
	return s.IspName
}

func (s *MobileOnlineStatusResponseBodyResultObject) GetSubCode() *string {
	return s.SubCode
}

func (s *MobileOnlineStatusResponseBodyResultObject) SetBizCode(v string) *MobileOnlineStatusResponseBodyResultObject {
	s.BizCode = &v
	return s
}

func (s *MobileOnlineStatusResponseBodyResultObject) SetIspName(v string) *MobileOnlineStatusResponseBodyResultObject {
	s.IspName = &v
	return s
}

func (s *MobileOnlineStatusResponseBodyResultObject) SetSubCode(v string) *MobileOnlineStatusResponseBodyResultObject {
	s.SubCode = &v
	return s
}

func (s *MobileOnlineStatusResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
