// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCustVariableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyCustVariableResponseBody
	GetRequestId() *string
	SetResultObject(v []*ModifyCustVariableResponseBodyResultObject) *ModifyCustVariableResponseBody
	GetResultObject() []*ModifyCustVariableResponseBodyResultObject
}

type ModifyCustVariableResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return object
	ResultObject []*ModifyCustVariableResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Repeated"`
}

func (s ModifyCustVariableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyCustVariableResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCustVariableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyCustVariableResponseBody) GetResultObject() []*ModifyCustVariableResponseBodyResultObject {
	return s.ResultObject
}

func (s *ModifyCustVariableResponseBody) SetRequestId(v string) *ModifyCustVariableResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyCustVariableResponseBody) SetResultObject(v []*ModifyCustVariableResponseBodyResultObject) *ModifyCustVariableResponseBody {
	s.ResultObject = v
	return s
}

func (s *ModifyCustVariableResponseBody) Validate() error {
	if s.ResultObject != nil {
		for _, item := range s.ResultObject {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyCustVariableResponseBodyResultObject struct {
	// Failure type
	//
	// example:
	//
	// rule
	FailType *string `json:"failType,omitempty" xml:"failType,omitempty"`
	// Detailed information.
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Whether the operation was successful
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ModifyCustVariableResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s ModifyCustVariableResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *ModifyCustVariableResponseBodyResultObject) GetFailType() *string {
	return s.FailType
}

func (s *ModifyCustVariableResponseBodyResultObject) GetMessage() *string {
	return s.Message
}

func (s *ModifyCustVariableResponseBodyResultObject) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyCustVariableResponseBodyResultObject) SetFailType(v string) *ModifyCustVariableResponseBodyResultObject {
	s.FailType = &v
	return s
}

func (s *ModifyCustVariableResponseBodyResultObject) SetMessage(v string) *ModifyCustVariableResponseBodyResultObject {
	s.Message = &v
	return s
}

func (s *ModifyCustVariableResponseBodyResultObject) SetSuccess(v bool) *ModifyCustVariableResponseBodyResultObject {
	s.Success = &v
	return s
}

func (s *ModifyCustVariableResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
