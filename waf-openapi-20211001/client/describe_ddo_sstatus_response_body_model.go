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
	// The current DDoS status.
	DDoSStatus []*DescribeDDoSStatusResponseBodyDDoSStatus `json:"DDoSStatus,omitempty" xml:"DDoSStatus,omitempty" type:"Repeated"`
	// The request ID.
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
	// The event type. Valid values:
	//
	// - defense: DDoS scrubbing events.
	//
	// - blackhole: DDoS blackhole filtering events.
	//
	// example:
	//
	// blackhole
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The DDoS status. Valid values:
	//
	// - **doing**: under DDoS attack.
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
