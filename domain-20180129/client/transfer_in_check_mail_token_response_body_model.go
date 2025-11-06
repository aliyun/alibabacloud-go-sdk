// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferInCheckMailTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFailList(v *TransferInCheckMailTokenResponseBodyFailList) *TransferInCheckMailTokenResponseBody
	GetFailList() *TransferInCheckMailTokenResponseBodyFailList
	SetRequestId(v string) *TransferInCheckMailTokenResponseBody
	GetRequestId() *string
	SetSuccessList(v *TransferInCheckMailTokenResponseBodySuccessList) *TransferInCheckMailTokenResponseBody
	GetSuccessList() *TransferInCheckMailTokenResponseBodySuccessList
}

type TransferInCheckMailTokenResponseBody struct {
	FailList *TransferInCheckMailTokenResponseBodyFailList `json:"FailList,omitempty" xml:"FailList,omitempty" type:"Struct"`
	// example:
	//
	// AF7D4DCE-0776-47F2-A9B2-6FB85A87AA60
	RequestId   *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SuccessList *TransferInCheckMailTokenResponseBodySuccessList `json:"SuccessList,omitempty" xml:"SuccessList,omitempty" type:"Struct"`
}

func (s TransferInCheckMailTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TransferInCheckMailTokenResponseBody) GoString() string {
	return s.String()
}

func (s *TransferInCheckMailTokenResponseBody) GetFailList() *TransferInCheckMailTokenResponseBodyFailList {
	return s.FailList
}

func (s *TransferInCheckMailTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TransferInCheckMailTokenResponseBody) GetSuccessList() *TransferInCheckMailTokenResponseBodySuccessList {
	return s.SuccessList
}

func (s *TransferInCheckMailTokenResponseBody) SetFailList(v *TransferInCheckMailTokenResponseBodyFailList) *TransferInCheckMailTokenResponseBody {
	s.FailList = v
	return s
}

func (s *TransferInCheckMailTokenResponseBody) SetRequestId(v string) *TransferInCheckMailTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *TransferInCheckMailTokenResponseBody) SetSuccessList(v *TransferInCheckMailTokenResponseBodySuccessList) *TransferInCheckMailTokenResponseBody {
	s.SuccessList = v
	return s
}

func (s *TransferInCheckMailTokenResponseBody) Validate() error {
	if s.FailList != nil {
		if err := s.FailList.Validate(); err != nil {
			return err
		}
	}
	if s.SuccessList != nil {
		if err := s.SuccessList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TransferInCheckMailTokenResponseBodyFailList struct {
	FailDomain []*string `json:"FailDomain,omitempty" xml:"FailDomain,omitempty" type:"Repeated"`
}

func (s TransferInCheckMailTokenResponseBodyFailList) String() string {
	return dara.Prettify(s)
}

func (s TransferInCheckMailTokenResponseBodyFailList) GoString() string {
	return s.String()
}

func (s *TransferInCheckMailTokenResponseBodyFailList) GetFailDomain() []*string {
	return s.FailDomain
}

func (s *TransferInCheckMailTokenResponseBodyFailList) SetFailDomain(v []*string) *TransferInCheckMailTokenResponseBodyFailList {
	s.FailDomain = v
	return s
}

func (s *TransferInCheckMailTokenResponseBodyFailList) Validate() error {
	return dara.Validate(s)
}

type TransferInCheckMailTokenResponseBodySuccessList struct {
	SuccessDomain []*string `json:"SuccessDomain,omitempty" xml:"SuccessDomain,omitempty" type:"Repeated"`
}

func (s TransferInCheckMailTokenResponseBodySuccessList) String() string {
	return dara.Prettify(s)
}

func (s TransferInCheckMailTokenResponseBodySuccessList) GoString() string {
	return s.String()
}

func (s *TransferInCheckMailTokenResponseBodySuccessList) GetSuccessDomain() []*string {
	return s.SuccessDomain
}

func (s *TransferInCheckMailTokenResponseBodySuccessList) SetSuccessDomain(v []*string) *TransferInCheckMailTokenResponseBodySuccessList {
	s.SuccessDomain = v
	return s
}

func (s *TransferInCheckMailTokenResponseBodySuccessList) Validate() error {
	return dara.Validate(s)
}
