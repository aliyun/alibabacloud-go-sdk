// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDomainSuffixResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryDomainSuffixResponseBody
	GetRequestId() *string
	SetSuffixList(v *QueryDomainSuffixResponseBodySuffixList) *QueryDomainSuffixResponseBody
	GetSuffixList() *QueryDomainSuffixResponseBodySuffixList
}

type QueryDomainSuffixResponseBody struct {
	// example:
	//
	// D1C9DE44-1D7F-4F66-9653-00000
	RequestId  *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SuffixList *QueryDomainSuffixResponseBodySuffixList `json:"SuffixList,omitempty" xml:"SuffixList,omitempty" type:"Struct"`
}

func (s QueryDomainSuffixResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryDomainSuffixResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDomainSuffixResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryDomainSuffixResponseBody) GetSuffixList() *QueryDomainSuffixResponseBodySuffixList {
	return s.SuffixList
}

func (s *QueryDomainSuffixResponseBody) SetRequestId(v string) *QueryDomainSuffixResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDomainSuffixResponseBody) SetSuffixList(v *QueryDomainSuffixResponseBodySuffixList) *QueryDomainSuffixResponseBody {
	s.SuffixList = v
	return s
}

func (s *QueryDomainSuffixResponseBody) Validate() error {
	if s.SuffixList != nil {
		if err := s.SuffixList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryDomainSuffixResponseBodySuffixList struct {
	Suffix []*string `json:"Suffix,omitempty" xml:"Suffix,omitempty" type:"Repeated"`
}

func (s QueryDomainSuffixResponseBodySuffixList) String() string {
	return dara.Prettify(s)
}

func (s QueryDomainSuffixResponseBodySuffixList) GoString() string {
	return s.String()
}

func (s *QueryDomainSuffixResponseBodySuffixList) GetSuffix() []*string {
	return s.Suffix
}

func (s *QueryDomainSuffixResponseBodySuffixList) SetSuffix(v []*string) *QueryDomainSuffixResponseBodySuffixList {
	s.Suffix = v
	return s
}

func (s *QueryDomainSuffixResponseBodySuffixList) Validate() error {
	return dara.Validate(s)
}
