// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePortAutoCcStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPortAutoCcStatus(v []*DescribePortAutoCcStatusResponseBodyPortAutoCcStatus) *DescribePortAutoCcStatusResponseBody
	GetPortAutoCcStatus() []*DescribePortAutoCcStatusResponseBodyPortAutoCcStatus
	SetRequestId(v string) *DescribePortAutoCcStatusResponseBody
	GetRequestId() *string
}

type DescribePortAutoCcStatusResponseBody struct {
	// An array that consists of the configurations of the Intelligent Protection policy.
	PortAutoCcStatus []*DescribePortAutoCcStatusResponseBodyPortAutoCcStatus `json:"PortAutoCcStatus,omitempty" xml:"PortAutoCcStatus,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// BC3C6403-F248-4125-B2C9-8733ED94EA85
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePortAutoCcStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePortAutoCcStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePortAutoCcStatusResponseBody) GetPortAutoCcStatus() []*DescribePortAutoCcStatusResponseBodyPortAutoCcStatus {
	return s.PortAutoCcStatus
}

func (s *DescribePortAutoCcStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePortAutoCcStatusResponseBody) SetPortAutoCcStatus(v []*DescribePortAutoCcStatusResponseBodyPortAutoCcStatus) *DescribePortAutoCcStatusResponseBody {
	s.PortAutoCcStatus = v
	return s
}

func (s *DescribePortAutoCcStatusResponseBody) SetRequestId(v string) *DescribePortAutoCcStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePortAutoCcStatusResponseBody) Validate() error {
	if s.PortAutoCcStatus != nil {
		for _, item := range s.PortAutoCcStatus {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePortAutoCcStatusResponseBodyPortAutoCcStatus struct {
	// The mode of the Intelligent Protection policy. Valid values:
	//
	// 	- **normal**
	//
	// 	- **loose**
	//
	// 	- **strict**
	//
	// example:
	//
	// normal
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The status of the Intelligent Protection policy. Valid values:
	//
	// 	- **on**: enabled
	//
	// 	- **off**: disabled
	//
	// example:
	//
	// on
	Switch *string `json:"Switch,omitempty" xml:"Switch,omitempty"`
	// The protection mode for ports 80 and 443. Valid values:
	//
	// 	- **normal**
	//
	// 	- **loose**
	//
	// 	- **strict**
	//
	// example:
	//
	// normal
	WebMode *string `json:"WebMode,omitempty" xml:"WebMode,omitempty"`
	// The status of the Intelligent Protection policy for ports 80 and 443. Valid values:
	//
	// 	- **on**: enabled
	//
	// 	- **off**: disabled
	//
	// example:
	//
	// off
	WebSwitch *string `json:"WebSwitch,omitempty" xml:"WebSwitch,omitempty"`
}

func (s DescribePortAutoCcStatusResponseBodyPortAutoCcStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribePortAutoCcStatusResponseBodyPortAutoCcStatus) GoString() string {
	return s.String()
}

func (s *DescribePortAutoCcStatusResponseBodyPortAutoCcStatus) GetMode() *string {
	return s.Mode
}

func (s *DescribePortAutoCcStatusResponseBodyPortAutoCcStatus) GetSwitch() *string {
	return s.Switch
}

func (s *DescribePortAutoCcStatusResponseBodyPortAutoCcStatus) GetWebMode() *string {
	return s.WebMode
}

func (s *DescribePortAutoCcStatusResponseBodyPortAutoCcStatus) GetWebSwitch() *string {
	return s.WebSwitch
}

func (s *DescribePortAutoCcStatusResponseBodyPortAutoCcStatus) SetMode(v string) *DescribePortAutoCcStatusResponseBodyPortAutoCcStatus {
	s.Mode = &v
	return s
}

func (s *DescribePortAutoCcStatusResponseBodyPortAutoCcStatus) SetSwitch(v string) *DescribePortAutoCcStatusResponseBodyPortAutoCcStatus {
	s.Switch = &v
	return s
}

func (s *DescribePortAutoCcStatusResponseBodyPortAutoCcStatus) SetWebMode(v string) *DescribePortAutoCcStatusResponseBodyPortAutoCcStatus {
	s.WebMode = &v
	return s
}

func (s *DescribePortAutoCcStatusResponseBodyPortAutoCcStatus) SetWebSwitch(v string) *DescribePortAutoCcStatusResponseBodyPortAutoCcStatus {
	s.WebSwitch = &v
	return s
}

func (s *DescribePortAutoCcStatusResponseBodyPortAutoCcStatus) Validate() error {
	return dara.Validate(s)
}
