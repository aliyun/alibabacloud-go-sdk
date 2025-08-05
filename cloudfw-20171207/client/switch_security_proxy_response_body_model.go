// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchSecurityProxyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SwitchSecurityProxyResponseBody
	GetRequestId() *string
}

type SwitchSecurityProxyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F1E55690-3ABA-58FA-90E3-593EF05B73ED
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SwitchSecurityProxyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SwitchSecurityProxyResponseBody) GoString() string {
	return s.String()
}

func (s *SwitchSecurityProxyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SwitchSecurityProxyResponseBody) SetRequestId(v string) *SwitchSecurityProxyResponseBody {
	s.RequestId = &v
	return s
}

func (s *SwitchSecurityProxyResponseBody) Validate() error {
	return dara.Validate(s)
}
