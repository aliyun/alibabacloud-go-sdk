// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckCreatedByEnabledResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOpenStatus(v bool) *CheckCreatedByEnabledResponseBody
	GetOpenStatus() *bool
	SetRequestId(v string) *CheckCreatedByEnabledResponseBody
	GetRequestId() *string
}

type CheckCreatedByEnabledResponseBody struct {
	// example:
	//
	// false
	OpenStatus *bool `json:"OpenStatus,omitempty" xml:"OpenStatus,omitempty"`
	// example:
	//
	// 682DD9E1-F530-5D14-A839-A6787FA82B74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckCreatedByEnabledResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckCreatedByEnabledResponseBody) GoString() string {
	return s.String()
}

func (s *CheckCreatedByEnabledResponseBody) GetOpenStatus() *bool {
	return s.OpenStatus
}

func (s *CheckCreatedByEnabledResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckCreatedByEnabledResponseBody) SetOpenStatus(v bool) *CheckCreatedByEnabledResponseBody {
	s.OpenStatus = &v
	return s
}

func (s *CheckCreatedByEnabledResponseBody) SetRequestId(v string) *CheckCreatedByEnabledResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckCreatedByEnabledResponseBody) Validate() error {
	return dara.Validate(s)
}
