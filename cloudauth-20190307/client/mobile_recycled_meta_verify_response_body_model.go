// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMobileRecycledMetaVerifyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *MobileRecycledMetaVerifyResponseBody
	GetCode() *string
	SetMessage(v string) *MobileRecycledMetaVerifyResponseBody
	GetMessage() *string
	SetRequestId(v string) *MobileRecycledMetaVerifyResponseBody
	GetRequestId() *string
	SetResultObject(v *MobileRecycledMetaVerifyResponseBodyResultObject) *MobileRecycledMetaVerifyResponseBody
	GetResultObject() *MobileRecycledMetaVerifyResponseBodyResultObject
}

type MobileRecycledMetaVerifyResponseBody struct {
	// The response code. A value of 200 indicates success. Any other value indicates failure.
	//
	// > **Important**
	//
	// - This parameter indicates whether the API operation is called correctly. For more information about return codes, see error codes.
	//
	// - Check the business verification result in the fields of ResultObject.
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
	// 8FC3D6AC-9FED-4311-8DA7-C4BF47D9F260
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result.
	ResultObject *MobileRecycledMetaVerifyResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s MobileRecycledMetaVerifyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MobileRecycledMetaVerifyResponseBody) GoString() string {
	return s.String()
}

func (s *MobileRecycledMetaVerifyResponseBody) GetCode() *string {
	return s.Code
}

func (s *MobileRecycledMetaVerifyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *MobileRecycledMetaVerifyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MobileRecycledMetaVerifyResponseBody) GetResultObject() *MobileRecycledMetaVerifyResponseBodyResultObject {
	return s.ResultObject
}

func (s *MobileRecycledMetaVerifyResponseBody) SetCode(v string) *MobileRecycledMetaVerifyResponseBody {
	s.Code = &v
	return s
}

func (s *MobileRecycledMetaVerifyResponseBody) SetMessage(v string) *MobileRecycledMetaVerifyResponseBody {
	s.Message = &v
	return s
}

func (s *MobileRecycledMetaVerifyResponseBody) SetRequestId(v string) *MobileRecycledMetaVerifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *MobileRecycledMetaVerifyResponseBody) SetResultObject(v *MobileRecycledMetaVerifyResponseBodyResultObject) *MobileRecycledMetaVerifyResponseBody {
	s.ResultObject = v
	return s
}

func (s *MobileRecycledMetaVerifyResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MobileRecycledMetaVerifyResponseBodyResultObject struct {
	// The query result. Valid values:
	//
	// - 1: A query result is found.
	//
	// - 3: No query result is found.
	//
	// example:
	//
	// 1
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// The carrier name. China Mobile: CMCC. China Unicom: CUCC. China Telecom: CTCC.
	//
	// example:
	//
	// CMCC
	IspName *string `json:"IspName,omitempty" xml:"IspName,omitempty"`
	// The detailed verification result. Valid values:
	//
	// - 101: The registration date is equal to or later than the phone number activation date.
	//
	// - 102: The registration date is earlier than the phone number activation date.
	//
	// - 103: The new subscriber has not been synchronized yet.
	//
	// - 301: Data exception or the subscriber has been deactivated.
	//
	// example:
	//
	// 101
	SubCode *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
}

func (s MobileRecycledMetaVerifyResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s MobileRecycledMetaVerifyResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *MobileRecycledMetaVerifyResponseBodyResultObject) GetBizCode() *string {
	return s.BizCode
}

func (s *MobileRecycledMetaVerifyResponseBodyResultObject) GetIspName() *string {
	return s.IspName
}

func (s *MobileRecycledMetaVerifyResponseBodyResultObject) GetSubCode() *string {
	return s.SubCode
}

func (s *MobileRecycledMetaVerifyResponseBodyResultObject) SetBizCode(v string) *MobileRecycledMetaVerifyResponseBodyResultObject {
	s.BizCode = &v
	return s
}

func (s *MobileRecycledMetaVerifyResponseBodyResultObject) SetIspName(v string) *MobileRecycledMetaVerifyResponseBodyResultObject {
	s.IspName = &v
	return s
}

func (s *MobileRecycledMetaVerifyResponseBodyResultObject) SetSubCode(v string) *MobileRecycledMetaVerifyResponseBodyResultObject {
	s.SubCode = &v
	return s
}

func (s *MobileRecycledMetaVerifyResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
