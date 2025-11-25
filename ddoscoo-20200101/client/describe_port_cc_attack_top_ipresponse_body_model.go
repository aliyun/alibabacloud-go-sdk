// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePortCcAttackTopIPResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribePortCcAttackTopIPResponseBody
	GetRequestId() *string
	SetTopIp(v []*DescribePortCcAttackTopIPResponseBodyTopIp) *DescribePortCcAttackTopIPResponseBody
	GetTopIp() []*DescribePortCcAttackTopIPResponseBodyTopIp
}

type DescribePortCcAttackTopIPResponseBody struct {
	// The request ID, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 24F36D81-5E2D-52E5-9DB6-A3ED23CF271A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The top IP addresses from which most attacks are initiated.
	TopIp []*DescribePortCcAttackTopIPResponseBodyTopIp `json:"TopIp,omitempty" xml:"TopIp,omitempty" type:"Repeated"`
}

func (s DescribePortCcAttackTopIPResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePortCcAttackTopIPResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePortCcAttackTopIPResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePortCcAttackTopIPResponseBody) GetTopIp() []*DescribePortCcAttackTopIPResponseBodyTopIp {
	return s.TopIp
}

func (s *DescribePortCcAttackTopIPResponseBody) SetRequestId(v string) *DescribePortCcAttackTopIPResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePortCcAttackTopIPResponseBody) SetTopIp(v []*DescribePortCcAttackTopIPResponseBodyTopIp) *DescribePortCcAttackTopIPResponseBody {
	s.TopIp = v
	return s
}

func (s *DescribePortCcAttackTopIPResponseBody) Validate() error {
	if s.TopIp != nil {
		for _, item := range s.TopIp {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePortCcAttackTopIPResponseBodyTopIp struct {
	// The code of the location from which the attack is initiated. For more information, see [Codes of administrative regions in China and codes of countries and areas](https://help.aliyun.com/document_detail/167926.html). For example, **110000*	- indicates Beijing, China, and **us*	- indicates the United States.
	//
	// example:
	//
	// cn-shanghai
	AreaId *string `json:"AreaId,omitempty" xml:"AreaId,omitempty"`
	// The number of attacks from the IP address.
	//
	// example:
	//
	// 33971
	Pv *int64 `json:"Pv,omitempty" xml:"Pv,omitempty"`
	// The source IP address of the attack.
	//
	// example:
	//
	// 172.18.XX.XX
	SrcIp *string `json:"SrcIp,omitempty" xml:"SrcIp,omitempty"`
}

func (s DescribePortCcAttackTopIPResponseBodyTopIp) String() string {
	return dara.Prettify(s)
}

func (s DescribePortCcAttackTopIPResponseBodyTopIp) GoString() string {
	return s.String()
}

func (s *DescribePortCcAttackTopIPResponseBodyTopIp) GetAreaId() *string {
	return s.AreaId
}

func (s *DescribePortCcAttackTopIPResponseBodyTopIp) GetPv() *int64 {
	return s.Pv
}

func (s *DescribePortCcAttackTopIPResponseBodyTopIp) GetSrcIp() *string {
	return s.SrcIp
}

func (s *DescribePortCcAttackTopIPResponseBodyTopIp) SetAreaId(v string) *DescribePortCcAttackTopIPResponseBodyTopIp {
	s.AreaId = &v
	return s
}

func (s *DescribePortCcAttackTopIPResponseBodyTopIp) SetPv(v int64) *DescribePortCcAttackTopIPResponseBodyTopIp {
	s.Pv = &v
	return s
}

func (s *DescribePortCcAttackTopIPResponseBodyTopIp) SetSrcIp(v string) *DescribePortCcAttackTopIPResponseBodyTopIp {
	s.SrcIp = &v
	return s
}

func (s *DescribePortCcAttackTopIPResponseBodyTopIp) Validate() error {
	return dara.Validate(s)
}
