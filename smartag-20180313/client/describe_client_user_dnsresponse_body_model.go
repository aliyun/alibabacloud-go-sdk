// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClientUserDNSResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppDNS(v []*string) *DescribeClientUserDNSResponseBody
	GetAppDNS() []*string
	SetRecoveredDNS(v []*string) *DescribeClientUserDNSResponseBody
	GetRecoveredDNS() []*string
	SetRequestId(v string) *DescribeClientUserDNSResponseBody
	GetRequestId() *string
}

type DescribeClientUserDNSResponseBody struct {
	// The active and standby DNS servers that the SAG app instance uses when it connects to private networks.
	AppDNS []*string `json:"AppDNS,omitempty" xml:"AppDNS,omitempty" type:"Repeated"`
	// The active and standby DNS servers that the SAG app instance uses when it disconnects from private networks.
	RecoveredDNS []*string `json:"RecoveredDNS,omitempty" xml:"RecoveredDNS,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 894AA8BD-0627-45B1-AA18-9CE1D50DA9D0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeClientUserDNSResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeClientUserDNSResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClientUserDNSResponseBody) GetAppDNS() []*string {
	return s.AppDNS
}

func (s *DescribeClientUserDNSResponseBody) GetRecoveredDNS() []*string {
	return s.RecoveredDNS
}

func (s *DescribeClientUserDNSResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeClientUserDNSResponseBody) SetAppDNS(v []*string) *DescribeClientUserDNSResponseBody {
	s.AppDNS = v
	return s
}

func (s *DescribeClientUserDNSResponseBody) SetRecoveredDNS(v []*string) *DescribeClientUserDNSResponseBody {
	s.RecoveredDNS = v
	return s
}

func (s *DescribeClientUserDNSResponseBody) SetRequestId(v string) *DescribeClientUserDNSResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeClientUserDNSResponseBody) Validate() error {
	return dara.Validate(s)
}
