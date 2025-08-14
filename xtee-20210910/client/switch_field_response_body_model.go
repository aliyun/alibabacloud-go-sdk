// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchFieldResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SwitchFieldResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *SwitchFieldResponseBody
	GetResultObject() *bool
}

type SwitchFieldResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return Object
	//
	// example:
	//
	// true
	ResultObject *bool `json:"resultObject,omitempty" xml:"resultObject,omitempty"`
}

func (s SwitchFieldResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SwitchFieldResponseBody) GoString() string {
	return s.String()
}

func (s *SwitchFieldResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SwitchFieldResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *SwitchFieldResponseBody) SetRequestId(v string) *SwitchFieldResponseBody {
	s.RequestId = &v
	return s
}

func (s *SwitchFieldResponseBody) SetResultObject(v bool) *SwitchFieldResponseBody {
	s.ResultObject = &v
	return s
}

func (s *SwitchFieldResponseBody) Validate() error {
	return dara.Validate(s)
}
