// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableRightsSeparationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *DisableRightsSeparationResponseBody
	GetMessage() *string
	SetRequestId(v string) *DisableRightsSeparationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DisableRightsSeparationResponseBody
	GetSuccess() *bool
}

type DisableRightsSeparationResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// FE5D94E3-3C93-3594-95D9-AAED2A980915
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DisableRightsSeparationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableRightsSeparationResponseBody) GoString() string {
	return s.String()
}

func (s *DisableRightsSeparationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DisableRightsSeparationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableRightsSeparationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DisableRightsSeparationResponseBody) SetMessage(v string) *DisableRightsSeparationResponseBody {
	s.Message = &v
	return s
}

func (s *DisableRightsSeparationResponseBody) SetRequestId(v string) *DisableRightsSeparationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableRightsSeparationResponseBody) SetSuccess(v bool) *DisableRightsSeparationResponseBody {
	s.Success = &v
	return s
}

func (s *DisableRightsSeparationResponseBody) Validate() error {
	return dara.Validate(s)
}
