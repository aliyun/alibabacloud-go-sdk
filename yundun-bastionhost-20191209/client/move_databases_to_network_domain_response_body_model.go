// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveDatabasesToNetworkDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *MoveDatabasesToNetworkDomainResponseBody
	GetRequestId() *string
	SetResults(v []*MoveDatabasesToNetworkDomainResponseBodyResults) *MoveDatabasesToNetworkDomainResponseBody
	GetResults() []*MoveDatabasesToNetworkDomainResponseBodyResults
}

type MoveDatabasesToNetworkDomainResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// E5B1BC32-72B2-5BFD-BF75-5D38261264D1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The results of the call.
	Results []*MoveDatabasesToNetworkDomainResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s MoveDatabasesToNetworkDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MoveDatabasesToNetworkDomainResponseBody) GoString() string {
	return s.String()
}

func (s *MoveDatabasesToNetworkDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MoveDatabasesToNetworkDomainResponseBody) GetResults() []*MoveDatabasesToNetworkDomainResponseBodyResults {
	return s.Results
}

func (s *MoveDatabasesToNetworkDomainResponseBody) SetRequestId(v string) *MoveDatabasesToNetworkDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *MoveDatabasesToNetworkDomainResponseBody) SetResults(v []*MoveDatabasesToNetworkDomainResponseBodyResults) *MoveDatabasesToNetworkDomainResponseBody {
	s.Results = v
	return s
}

func (s *MoveDatabasesToNetworkDomainResponseBody) Validate() error {
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

type MoveDatabasesToNetworkDomainResponseBodyResults struct {
	// Indicates whether the database is added to the network domain.
	//
	// > The code LICENSE_OUT_OF_LIMIT indicates that the network domain feature is not supported by the current Bastionhost edition.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The database ID.
	//
	// example:
	//
	// 45
	DatabaseId *string `json:"DatabaseId,omitempty" xml:"DatabaseId,omitempty"`
	// The error message that is returned.
	//
	// example:
	//
	// N/A
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s MoveDatabasesToNetworkDomainResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s MoveDatabasesToNetworkDomainResponseBodyResults) GoString() string {
	return s.String()
}

func (s *MoveDatabasesToNetworkDomainResponseBodyResults) GetCode() *string {
	return s.Code
}

func (s *MoveDatabasesToNetworkDomainResponseBodyResults) GetDatabaseId() *string {
	return s.DatabaseId
}

func (s *MoveDatabasesToNetworkDomainResponseBodyResults) GetMessage() *string {
	return s.Message
}

func (s *MoveDatabasesToNetworkDomainResponseBodyResults) SetCode(v string) *MoveDatabasesToNetworkDomainResponseBodyResults {
	s.Code = &v
	return s
}

func (s *MoveDatabasesToNetworkDomainResponseBodyResults) SetDatabaseId(v string) *MoveDatabasesToNetworkDomainResponseBodyResults {
	s.DatabaseId = &v
	return s
}

func (s *MoveDatabasesToNetworkDomainResponseBodyResults) SetMessage(v string) *MoveDatabasesToNetworkDomainResponseBodyResults {
	s.Message = &v
	return s
}

func (s *MoveDatabasesToNetworkDomainResponseBodyResults) Validate() error {
	return dara.Validate(s)
}
