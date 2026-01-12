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
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 8FC3D6AC-9FED-4311-8DA7-C4BF47D9F260
	RequestId    *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// 1
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
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

func (s *MobileRecycledMetaVerifyResponseBodyResultObject) GetSubCode() *string {
	return s.SubCode
}

func (s *MobileRecycledMetaVerifyResponseBodyResultObject) SetBizCode(v string) *MobileRecycledMetaVerifyResponseBodyResultObject {
	s.BizCode = &v
	return s
}

func (s *MobileRecycledMetaVerifyResponseBodyResultObject) SetSubCode(v string) *MobileRecycledMetaVerifyResponseBodyResultObject {
	s.SubCode = &v
	return s
}

func (s *MobileRecycledMetaVerifyResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
