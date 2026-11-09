// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySessionByClientIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOnlineStatus(v bool) *QuerySessionByClientIdResponseBody
	GetOnlineStatus() *bool
	SetRequestId(v string) *QuerySessionByClientIdResponseBody
	GetRequestId() *string
}

type QuerySessionByClientIdResponseBody struct {
	// Indicates whether the ApsaraMQ for MQTT client is connected to the ApsaraMQ for MQTT broker. Valid values:
	//
	// - **true**
	//
	// - **false**
	//
	// example:
	//
	// true
	OnlineStatus *bool `json:"OnlineStatus,omitempty" xml:"OnlineStatus,omitempty"`
	// The request ID. This parameter is a common parameter.
	//
	// example:
	//
	// E4581CCF-62AF-44D9-B5B4-D1DBBC0E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QuerySessionByClientIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QuerySessionByClientIdResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySessionByClientIdResponseBody) GetOnlineStatus() *bool {
	return s.OnlineStatus
}

func (s *QuerySessionByClientIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QuerySessionByClientIdResponseBody) SetOnlineStatus(v bool) *QuerySessionByClientIdResponseBody {
	s.OnlineStatus = &v
	return s
}

func (s *QuerySessionByClientIdResponseBody) SetRequestId(v string) *QuerySessionByClientIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySessionByClientIdResponseBody) Validate() error {
	return dara.Validate(s)
}
