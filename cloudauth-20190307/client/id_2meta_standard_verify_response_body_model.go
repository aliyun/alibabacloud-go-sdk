// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iId2MetaStandardVerifyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *Id2MetaStandardVerifyResponseBody
	GetCode() *string
	SetMessage(v string) *Id2MetaStandardVerifyResponseBody
	GetMessage() *string
	SetRequestId(v string) *Id2MetaStandardVerifyResponseBody
	GetRequestId() *string
	SetResultObject(v *Id2MetaStandardVerifyResponseBodyResultObject) *Id2MetaStandardVerifyResponseBody
	GetResultObject() *Id2MetaStandardVerifyResponseBodyResultObject
}

type Id2MetaStandardVerifyResponseBody struct {
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
	// D6163397-15C5-419C-9ACC-B7C83E0B4C10
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returned result information.
	ResultObject *Id2MetaStandardVerifyResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s Id2MetaStandardVerifyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s Id2MetaStandardVerifyResponseBody) GoString() string {
	return s.String()
}

func (s *Id2MetaStandardVerifyResponseBody) GetCode() *string {
	return s.Code
}

func (s *Id2MetaStandardVerifyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *Id2MetaStandardVerifyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *Id2MetaStandardVerifyResponseBody) GetResultObject() *Id2MetaStandardVerifyResponseBodyResultObject {
	return s.ResultObject
}

func (s *Id2MetaStandardVerifyResponseBody) SetCode(v string) *Id2MetaStandardVerifyResponseBody {
	s.Code = &v
	return s
}

func (s *Id2MetaStandardVerifyResponseBody) SetMessage(v string) *Id2MetaStandardVerifyResponseBody {
	s.Message = &v
	return s
}

func (s *Id2MetaStandardVerifyResponseBody) SetRequestId(v string) *Id2MetaStandardVerifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *Id2MetaStandardVerifyResponseBody) SetResultObject(v *Id2MetaStandardVerifyResponseBodyResultObject) *Id2MetaStandardVerifyResponseBody {
	s.ResultObject = v
	return s
}

func (s *Id2MetaStandardVerifyResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type Id2MetaStandardVerifyResponseBodyResultObject struct {
	// Verification result code:
	//
	// - **1**: verification matches.
	//
	// - **2**: verification does not match.
	//
	// - **3**: no record found.
	//
	// example:
	//
	// 1
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
}

func (s Id2MetaStandardVerifyResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s Id2MetaStandardVerifyResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *Id2MetaStandardVerifyResponseBodyResultObject) GetBizCode() *string {
	return s.BizCode
}

func (s *Id2MetaStandardVerifyResponseBodyResultObject) SetBizCode(v string) *Id2MetaStandardVerifyResponseBodyResultObject {
	s.BizCode = &v
	return s
}

func (s *Id2MetaStandardVerifyResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
