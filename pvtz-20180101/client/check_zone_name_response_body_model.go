// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckZoneNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCheck(v bool) *CheckZoneNameResponseBody
	GetCheck() *bool
	SetRequestId(v string) *CheckZoneNameResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CheckZoneNameResponseBody
	GetSuccess() *bool
}

type CheckZoneNameResponseBody struct {
	// Indicates whether the zone name can be added. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Check *bool `json:"Check,omitempty" xml:"Check,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CA29B88F-A571-4123-80D5-768AC2F7F806
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckZoneNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckZoneNameResponseBody) GoString() string {
	return s.String()
}

func (s *CheckZoneNameResponseBody) GetCheck() *bool {
	return s.Check
}

func (s *CheckZoneNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckZoneNameResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CheckZoneNameResponseBody) SetCheck(v bool) *CheckZoneNameResponseBody {
	s.Check = &v
	return s
}

func (s *CheckZoneNameResponseBody) SetRequestId(v string) *CheckZoneNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckZoneNameResponseBody) SetSuccess(v bool) *CheckZoneNameResponseBody {
	s.Success = &v
	return s
}

func (s *CheckZoneNameResponseBody) Validate() error {
	return dara.Validate(s)
}
