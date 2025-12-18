// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartRemediationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *StartRemediationResponseBody
	GetData() *bool
	SetRequestId(v string) *StartRemediationResponseBody
	GetRequestId() *string
}

type StartRemediationResponseBody struct {
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C7817373-78CB-4F9A-8AFA-E7A88E9D64A2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartRemediationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartRemediationResponseBody) GoString() string {
	return s.String()
}

func (s *StartRemediationResponseBody) GetData() *bool {
	return s.Data
}

func (s *StartRemediationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartRemediationResponseBody) SetData(v bool) *StartRemediationResponseBody {
	s.Data = &v
	return s
}

func (s *StartRemediationResponseBody) SetRequestId(v string) *StartRemediationResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartRemediationResponseBody) Validate() error {
	return dara.Validate(s)
}
