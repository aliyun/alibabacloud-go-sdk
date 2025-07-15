// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchApiResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SwitchApiResponseBody
	GetRequestId() *string
}

type SwitchApiResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BZ001
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SwitchApiResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SwitchApiResponseBody) GoString() string {
	return s.String()
}

func (s *SwitchApiResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SwitchApiResponseBody) SetRequestId(v string) *SwitchApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *SwitchApiResponseBody) Validate() error {
	return dara.Validate(s)
}
