// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfirmTransferInEmailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFailList(v *ConfirmTransferInEmailResponseBodyFailList) *ConfirmTransferInEmailResponseBody
	GetFailList() *ConfirmTransferInEmailResponseBodyFailList
	SetRequestId(v string) *ConfirmTransferInEmailResponseBody
	GetRequestId() *string
	SetSuccessList(v *ConfirmTransferInEmailResponseBodySuccessList) *ConfirmTransferInEmailResponseBody
	GetSuccessList() *ConfirmTransferInEmailResponseBodySuccessList
}

type ConfirmTransferInEmailResponseBody struct {
	FailList *ConfirmTransferInEmailResponseBodyFailList `json:"FailList,omitempty" xml:"FailList,omitempty" type:"Struct"`
	// example:
	//
	// 40F46D3D-F4F3-4CCB-AC30-2DD20E32E528
	RequestId   *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SuccessList *ConfirmTransferInEmailResponseBodySuccessList `json:"SuccessList,omitempty" xml:"SuccessList,omitempty" type:"Struct"`
}

func (s ConfirmTransferInEmailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfirmTransferInEmailResponseBody) GoString() string {
	return s.String()
}

func (s *ConfirmTransferInEmailResponseBody) GetFailList() *ConfirmTransferInEmailResponseBodyFailList {
	return s.FailList
}

func (s *ConfirmTransferInEmailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfirmTransferInEmailResponseBody) GetSuccessList() *ConfirmTransferInEmailResponseBodySuccessList {
	return s.SuccessList
}

func (s *ConfirmTransferInEmailResponseBody) SetFailList(v *ConfirmTransferInEmailResponseBodyFailList) *ConfirmTransferInEmailResponseBody {
	s.FailList = v
	return s
}

func (s *ConfirmTransferInEmailResponseBody) SetRequestId(v string) *ConfirmTransferInEmailResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfirmTransferInEmailResponseBody) SetSuccessList(v *ConfirmTransferInEmailResponseBodySuccessList) *ConfirmTransferInEmailResponseBody {
	s.SuccessList = v
	return s
}

func (s *ConfirmTransferInEmailResponseBody) Validate() error {
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

type ConfirmTransferInEmailResponseBodyFailList struct {
	FailDomain []*string `json:"FailDomain,omitempty" xml:"FailDomain,omitempty" type:"Repeated"`
}

func (s ConfirmTransferInEmailResponseBodyFailList) String() string {
	return dara.Prettify(s)
}

func (s ConfirmTransferInEmailResponseBodyFailList) GoString() string {
	return s.String()
}

func (s *ConfirmTransferInEmailResponseBodyFailList) GetFailDomain() []*string {
	return s.FailDomain
}

func (s *ConfirmTransferInEmailResponseBodyFailList) SetFailDomain(v []*string) *ConfirmTransferInEmailResponseBodyFailList {
	s.FailDomain = v
	return s
}

func (s *ConfirmTransferInEmailResponseBodyFailList) Validate() error {
	return dara.Validate(s)
}

type ConfirmTransferInEmailResponseBodySuccessList struct {
	SuccessDomain []*string `json:"SuccessDomain,omitempty" xml:"SuccessDomain,omitempty" type:"Repeated"`
}

func (s ConfirmTransferInEmailResponseBodySuccessList) String() string {
	return dara.Prettify(s)
}

func (s ConfirmTransferInEmailResponseBodySuccessList) GoString() string {
	return s.String()
}

func (s *ConfirmTransferInEmailResponseBodySuccessList) GetSuccessDomain() []*string {
	return s.SuccessDomain
}

func (s *ConfirmTransferInEmailResponseBodySuccessList) SetSuccessDomain(v []*string) *ConfirmTransferInEmailResponseBodySuccessList {
	s.SuccessDomain = v
	return s
}

func (s *ConfirmTransferInEmailResponseBodySuccessList) Validate() error {
	return dara.Validate(s)
}
