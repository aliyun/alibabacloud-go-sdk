// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeviceBelongIdentifiersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIdentifiers(v []*ListDeviceBelongIdentifiersResponseBodyIdentifiers) *ListDeviceBelongIdentifiersResponseBody
	GetIdentifiers() []*ListDeviceBelongIdentifiersResponseBodyIdentifiers
	SetMaxResults(v int32) *ListDeviceBelongIdentifiersResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListDeviceBelongIdentifiersResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListDeviceBelongIdentifiersResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListDeviceBelongIdentifiersResponseBody
	GetTotalCount() *int64
}

type ListDeviceBelongIdentifiersResponseBody struct {
	Identifiers []*ListDeviceBelongIdentifiersResponseBodyIdentifiers `json:"Identifiers,omitempty" xml:"Identifiers,omitempty" type:"Repeated"`
	MaxResults  *int32                                                `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken   *string                                               `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId   *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount  *int64                                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDeviceBelongIdentifiersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDeviceBelongIdentifiersResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeviceBelongIdentifiersResponseBody) GetIdentifiers() []*ListDeviceBelongIdentifiersResponseBodyIdentifiers {
	return s.Identifiers
}

func (s *ListDeviceBelongIdentifiersResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDeviceBelongIdentifiersResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDeviceBelongIdentifiersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDeviceBelongIdentifiersResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListDeviceBelongIdentifiersResponseBody) SetIdentifiers(v []*ListDeviceBelongIdentifiersResponseBodyIdentifiers) *ListDeviceBelongIdentifiersResponseBody {
	s.Identifiers = v
	return s
}

func (s *ListDeviceBelongIdentifiersResponseBody) SetMaxResults(v int32) *ListDeviceBelongIdentifiersResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListDeviceBelongIdentifiersResponseBody) SetNextToken(v string) *ListDeviceBelongIdentifiersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDeviceBelongIdentifiersResponseBody) SetRequestId(v string) *ListDeviceBelongIdentifiersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDeviceBelongIdentifiersResponseBody) SetTotalCount(v int64) *ListDeviceBelongIdentifiersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDeviceBelongIdentifiersResponseBody) Validate() error {
	if s.Identifiers != nil {
		for _, item := range s.Identifiers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDeviceBelongIdentifiersResponseBodyIdentifiers struct {
	CreateTime      *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	IdentifierId    *string `json:"IdentifierId,omitempty" xml:"IdentifierId,omitempty"`
	IdentifierType  *string `json:"IdentifierType,omitempty" xml:"IdentifierType,omitempty"`
	IdentifierValue *string `json:"IdentifierValue,omitempty" xml:"IdentifierValue,omitempty"`
}

func (s ListDeviceBelongIdentifiersResponseBodyIdentifiers) String() string {
	return dara.Prettify(s)
}

func (s ListDeviceBelongIdentifiersResponseBodyIdentifiers) GoString() string {
	return s.String()
}

func (s *ListDeviceBelongIdentifiersResponseBodyIdentifiers) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListDeviceBelongIdentifiersResponseBodyIdentifiers) GetIdentifierId() *string {
	return s.IdentifierId
}

func (s *ListDeviceBelongIdentifiersResponseBodyIdentifiers) GetIdentifierType() *string {
	return s.IdentifierType
}

func (s *ListDeviceBelongIdentifiersResponseBodyIdentifiers) GetIdentifierValue() *string {
	return s.IdentifierValue
}

func (s *ListDeviceBelongIdentifiersResponseBodyIdentifiers) SetCreateTime(v string) *ListDeviceBelongIdentifiersResponseBodyIdentifiers {
	s.CreateTime = &v
	return s
}

func (s *ListDeviceBelongIdentifiersResponseBodyIdentifiers) SetIdentifierId(v string) *ListDeviceBelongIdentifiersResponseBodyIdentifiers {
	s.IdentifierId = &v
	return s
}

func (s *ListDeviceBelongIdentifiersResponseBodyIdentifiers) SetIdentifierType(v string) *ListDeviceBelongIdentifiersResponseBodyIdentifiers {
	s.IdentifierType = &v
	return s
}

func (s *ListDeviceBelongIdentifiersResponseBodyIdentifiers) SetIdentifierValue(v string) *ListDeviceBelongIdentifiersResponseBodyIdentifiers {
	s.IdentifierValue = &v
	return s
}

func (s *ListDeviceBelongIdentifiersResponseBodyIdentifiers) Validate() error {
	return dara.Validate(s)
}
