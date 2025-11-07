// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iId2MetaPeriodVerifyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *Id2MetaPeriodVerifyResponseBody
	GetCode() *string
	SetMessage(v string) *Id2MetaPeriodVerifyResponseBody
	GetMessage() *string
	SetRequestId(v string) *Id2MetaPeriodVerifyResponseBody
	GetRequestId() *string
	SetResultObject(v *Id2MetaPeriodVerifyResponseBodyResultObject) *Id2MetaPeriodVerifyResponseBody
	GetResultObject() *Id2MetaPeriodVerifyResponseBodyResultObject
}

type Id2MetaPeriodVerifyResponseBody struct {
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
	// ID of the request
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returned result information.
	ResultObject *Id2MetaPeriodVerifyResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s Id2MetaPeriodVerifyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s Id2MetaPeriodVerifyResponseBody) GoString() string {
	return s.String()
}

func (s *Id2MetaPeriodVerifyResponseBody) GetCode() *string {
	return s.Code
}

func (s *Id2MetaPeriodVerifyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *Id2MetaPeriodVerifyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *Id2MetaPeriodVerifyResponseBody) GetResultObject() *Id2MetaPeriodVerifyResponseBodyResultObject {
	return s.ResultObject
}

func (s *Id2MetaPeriodVerifyResponseBody) SetCode(v string) *Id2MetaPeriodVerifyResponseBody {
	s.Code = &v
	return s
}

func (s *Id2MetaPeriodVerifyResponseBody) SetMessage(v string) *Id2MetaPeriodVerifyResponseBody {
	s.Message = &v
	return s
}

func (s *Id2MetaPeriodVerifyResponseBody) SetRequestId(v string) *Id2MetaPeriodVerifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *Id2MetaPeriodVerifyResponseBody) SetResultObject(v *Id2MetaPeriodVerifyResponseBodyResultObject) *Id2MetaPeriodVerifyResponseBody {
	s.ResultObject = v
	return s
}

func (s *Id2MetaPeriodVerifyResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type Id2MetaPeriodVerifyResponseBodyResultObject struct {
	// Verification result code:
	//
	// - **1**: Verification consistent.
	//
	// - **2**: Verification inconsistent.
	//
	// - **3**: No record found.
	//
	// example:
	//
	// 1
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
}

func (s Id2MetaPeriodVerifyResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s Id2MetaPeriodVerifyResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *Id2MetaPeriodVerifyResponseBodyResultObject) GetBizCode() *string {
	return s.BizCode
}

func (s *Id2MetaPeriodVerifyResponseBodyResultObject) SetBizCode(v string) *Id2MetaPeriodVerifyResponseBodyResultObject {
	s.BizCode = &v
	return s
}

func (s *Id2MetaPeriodVerifyResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
