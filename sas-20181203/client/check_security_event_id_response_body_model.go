// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckSecurityEventIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *CheckSecurityEventIdResponseBody
	GetData() *bool
	SetRequestId(v string) *CheckSecurityEventIdResponseBody
	GetRequestId() *string
}

type CheckSecurityEventIdResponseBody struct {
	// Indicates whether the alert events are generated on the server. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 1383B0DB-D5D6-4B0C-9E6B-75939C8E67FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckSecurityEventIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckSecurityEventIdResponseBody) GoString() string {
	return s.String()
}

func (s *CheckSecurityEventIdResponseBody) GetData() *bool {
	return s.Data
}

func (s *CheckSecurityEventIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckSecurityEventIdResponseBody) SetData(v bool) *CheckSecurityEventIdResponseBody {
	s.Data = &v
	return s
}

func (s *CheckSecurityEventIdResponseBody) SetRequestId(v string) *CheckSecurityEventIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckSecurityEventIdResponseBody) Validate() error {
	return dara.Validate(s)
}
