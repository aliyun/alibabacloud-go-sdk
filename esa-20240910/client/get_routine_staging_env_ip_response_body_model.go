// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRoutineStagingEnvIpResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIPV4(v []*string) *GetRoutineStagingEnvIpResponseBody
	GetIPV4() []*string
	SetRequestId(v string) *GetRoutineStagingEnvIpResponseBody
	GetRequestId() *string
}

type GetRoutineStagingEnvIpResponseBody struct {
	IPV4      []*string `json:"IPV4,omitempty" xml:"IPV4,omitempty" type:"Repeated"`
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetRoutineStagingEnvIpResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRoutineStagingEnvIpResponseBody) GoString() string {
	return s.String()
}

func (s *GetRoutineStagingEnvIpResponseBody) GetIPV4() []*string {
	return s.IPV4
}

func (s *GetRoutineStagingEnvIpResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRoutineStagingEnvIpResponseBody) SetIPV4(v []*string) *GetRoutineStagingEnvIpResponseBody {
	s.IPV4 = v
	return s
}

func (s *GetRoutineStagingEnvIpResponseBody) SetRequestId(v string) *GetRoutineStagingEnvIpResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRoutineStagingEnvIpResponseBody) Validate() error {
	return dara.Validate(s)
}
