// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySecurityCheckScheduleConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifySecurityCheckScheduleConfigResponseBody
	GetRequestId() *string
}

type ModifySecurityCheckScheduleConfigResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 02287C0D-8DA9-5766-B51A-A63192BD3E80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySecurityCheckScheduleConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySecurityCheckScheduleConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySecurityCheckScheduleConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySecurityCheckScheduleConfigResponseBody) SetRequestId(v string) *ModifySecurityCheckScheduleConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySecurityCheckScheduleConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
