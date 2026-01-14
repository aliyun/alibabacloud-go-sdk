// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSpareIpResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetSpareIpResponseBody
	GetRequestId() *string
	SetState(v string) *GetSpareIpResponseBody
	GetState() *string
}

type GetSpareIpResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 6FEA0CF3-D3B9-43E5-A304-D217037876A8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the secondary IP address. Valid values:
	//
	// 	- **active:*	- The secondary IP address is available.
	//
	// 	- **inuse:*	- The secondary IP address is in use.
	//
	// example:
	//
	// active
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s GetSpareIpResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSpareIpResponseBody) GoString() string {
	return s.String()
}

func (s *GetSpareIpResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSpareIpResponseBody) GetState() *string {
	return s.State
}

func (s *GetSpareIpResponseBody) SetRequestId(v string) *GetSpareIpResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSpareIpResponseBody) SetState(v string) *GetSpareIpResponseBody {
	s.State = &v
	return s
}

func (s *GetSpareIpResponseBody) Validate() error {
	return dara.Validate(s)
}
