// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetIMVMonitorDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ResetIMVMonitorDataResponseBody
	GetRequestId() *string
	SetStatus(v bool) *ResetIMVMonitorDataResponseBody
	GetStatus() *bool
}

type ResetIMVMonitorDataResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **false**
	//
	// 	- **true**
	//
	// example:
	//
	// true
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ResetIMVMonitorDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResetIMVMonitorDataResponseBody) GoString() string {
	return s.String()
}

func (s *ResetIMVMonitorDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResetIMVMonitorDataResponseBody) GetStatus() *bool {
	return s.Status
}

func (s *ResetIMVMonitorDataResponseBody) SetRequestId(v string) *ResetIMVMonitorDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetIMVMonitorDataResponseBody) SetStatus(v bool) *ResetIMVMonitorDataResponseBody {
	s.Status = &v
	return s
}

func (s *ResetIMVMonitorDataResponseBody) Validate() error {
	return dara.Validate(s)
}
