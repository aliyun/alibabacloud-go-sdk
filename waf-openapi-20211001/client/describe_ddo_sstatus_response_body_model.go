// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDDoSStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDDoSStatus(v []*DescribeDDoSStatusResponseBodyDDoSStatus) *DescribeDDoSStatusResponseBody
	GetDDoSStatus() []*DescribeDDoSStatusResponseBodyDDoSStatus
	SetRequestId(v string) *DescribeDDoSStatusResponseBody
	GetRequestId() *string
}

type DescribeDDoSStatusResponseBody struct {
	// Indicates whether DDoS attacks occur on specific domain names.
	DDoSStatus []*DescribeDDoSStatusResponseBodyDDoSStatus `json:"DDoSStatus,omitempty" xml:"DDoSStatus,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// D7861F61-5B61-46CE-A47C-***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDDoSStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDDoSStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDDoSStatusResponseBody) GetDDoSStatus() []*DescribeDDoSStatusResponseBodyDDoSStatus {
	return s.DDoSStatus
}

func (s *DescribeDDoSStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDDoSStatusResponseBody) SetDDoSStatus(v []*DescribeDDoSStatusResponseBodyDDoSStatus) *DescribeDDoSStatusResponseBody {
	s.DDoSStatus = v
	return s
}

func (s *DescribeDDoSStatusResponseBody) SetRequestId(v string) *DescribeDDoSStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDDoSStatusResponseBody) Validate() error {
	if s.DDoSStatus != nil {
		for _, item := range s.DDoSStatus {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDDoSStatusResponseBodyDDoSStatus struct {
	// The type of events that are triggered by DDoS attacks. Valid values:
	//
	// 	- defense: traffic scrubbing events.
	//
	// 	- blackhole: blackhole filtering events.
	//
	// example:
	//
	// blackhole
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// Indicates whether DDoS attacks occur on specific domain names. Valid value:
	//
	// 	- **doing**: DDoS attacks occur on specific domain names.
	//
	// example:
	//
	// doing
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDDoSStatusResponseBodyDDoSStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeDDoSStatusResponseBodyDDoSStatus) GoString() string {
	return s.String()
}

func (s *DescribeDDoSStatusResponseBodyDDoSStatus) GetEventType() *string {
	return s.EventType
}

func (s *DescribeDDoSStatusResponseBodyDDoSStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeDDoSStatusResponseBodyDDoSStatus) SetEventType(v string) *DescribeDDoSStatusResponseBodyDDoSStatus {
	s.EventType = &v
	return s
}

func (s *DescribeDDoSStatusResponseBodyDDoSStatus) SetStatus(v string) *DescribeDDoSStatusResponseBodyDDoSStatus {
	s.Status = &v
	return s
}

func (s *DescribeDDoSStatusResponseBodyDDoSStatus) Validate() error {
	return dara.Validate(s)
}
