// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveHostsToNetworkDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *MoveHostsToNetworkDomainResponseBody
	GetRequestId() *string
	SetResults(v []*MoveHostsToNetworkDomainResponseBodyResults) *MoveHostsToNetworkDomainResponseBody
	GetResults() []*MoveHostsToNetworkDomainResponseBodyResults
}

type MoveHostsToNetworkDomainResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F9B9E190-9C8E-5FEE-B963-7E9F1FD7FB4E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the call.
	Results []*MoveHostsToNetworkDomainResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s MoveHostsToNetworkDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MoveHostsToNetworkDomainResponseBody) GoString() string {
	return s.String()
}

func (s *MoveHostsToNetworkDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MoveHostsToNetworkDomainResponseBody) GetResults() []*MoveHostsToNetworkDomainResponseBodyResults {
	return s.Results
}

func (s *MoveHostsToNetworkDomainResponseBody) SetRequestId(v string) *MoveHostsToNetworkDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *MoveHostsToNetworkDomainResponseBody) SetResults(v []*MoveHostsToNetworkDomainResponseBodyResults) *MoveHostsToNetworkDomainResponseBody {
	s.Results = v
	return s
}

func (s *MoveHostsToNetworkDomainResponseBody) Validate() error {
	if s.Results != nil {
		for _, item := range s.Results {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type MoveHostsToNetworkDomainResponseBodyResults struct {
	// The return code that indicates whether the host is added to the network domain.
	//
	// > The code LICENSE_OUT_OF_LIMIT indicates that the network domain feature is not supported by the current Bastionhost edition.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The host ID.
	//
	// example:
	//
	// 1
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// The error message that is returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s MoveHostsToNetworkDomainResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s MoveHostsToNetworkDomainResponseBodyResults) GoString() string {
	return s.String()
}

func (s *MoveHostsToNetworkDomainResponseBodyResults) GetCode() *string {
	return s.Code
}

func (s *MoveHostsToNetworkDomainResponseBodyResults) GetHostId() *string {
	return s.HostId
}

func (s *MoveHostsToNetworkDomainResponseBodyResults) GetMessage() *string {
	return s.Message
}

func (s *MoveHostsToNetworkDomainResponseBodyResults) SetCode(v string) *MoveHostsToNetworkDomainResponseBodyResults {
	s.Code = &v
	return s
}

func (s *MoveHostsToNetworkDomainResponseBodyResults) SetHostId(v string) *MoveHostsToNetworkDomainResponseBodyResults {
	s.HostId = &v
	return s
}

func (s *MoveHostsToNetworkDomainResponseBodyResults) SetMessage(v string) *MoveHostsToNetworkDomainResponseBodyResults {
	s.Message = &v
	return s
}

func (s *MoveHostsToNetworkDomainResponseBodyResults) Validate() error {
	return dara.Validate(s)
}
