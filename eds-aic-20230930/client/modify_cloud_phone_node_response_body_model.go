// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCloudPhoneNodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyCloudPhoneNodeResponseBody
	GetRequestId() *string
}

type ModifyCloudPhoneNodeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 7B9EFA4F-4305-5968-BAEE-BD8B8DE5****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyCloudPhoneNodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyCloudPhoneNodeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCloudPhoneNodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyCloudPhoneNodeResponseBody) SetRequestId(v string) *ModifyCloudPhoneNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyCloudPhoneNodeResponseBody) Validate() error {
	return dara.Validate(s)
}
