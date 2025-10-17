// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCharacterSetNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCharacterSetNameItems(v *DescribeCharacterSetNameResponseBodyCharacterSetNameItems) *DescribeCharacterSetNameResponseBody
	GetCharacterSetNameItems() *DescribeCharacterSetNameResponseBodyCharacterSetNameItems
	SetEngine(v string) *DescribeCharacterSetNameResponseBody
	GetEngine() *string
	SetRequestId(v string) *DescribeCharacterSetNameResponseBody
	GetRequestId() *string
}

type DescribeCharacterSetNameResponseBody struct {
	// The character sets that are supported.
	CharacterSetNameItems *DescribeCharacterSetNameResponseBodyCharacterSetNameItems `json:"CharacterSetNameItems,omitempty" xml:"CharacterSetNameItems,omitempty" type:"Struct"`
	// The type of the database engine.
	//
	// example:
	//
	// POLARDB
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 34458CD3-33E0-4624-BFEF-840C15******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCharacterSetNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCharacterSetNameResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCharacterSetNameResponseBody) GetCharacterSetNameItems() *DescribeCharacterSetNameResponseBodyCharacterSetNameItems {
	return s.CharacterSetNameItems
}

func (s *DescribeCharacterSetNameResponseBody) GetEngine() *string {
	return s.Engine
}

func (s *DescribeCharacterSetNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCharacterSetNameResponseBody) SetCharacterSetNameItems(v *DescribeCharacterSetNameResponseBodyCharacterSetNameItems) *DescribeCharacterSetNameResponseBody {
	s.CharacterSetNameItems = v
	return s
}

func (s *DescribeCharacterSetNameResponseBody) SetEngine(v string) *DescribeCharacterSetNameResponseBody {
	s.Engine = &v
	return s
}

func (s *DescribeCharacterSetNameResponseBody) SetRequestId(v string) *DescribeCharacterSetNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCharacterSetNameResponseBody) Validate() error {
	if s.CharacterSetNameItems != nil {
		if err := s.CharacterSetNameItems.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCharacterSetNameResponseBodyCharacterSetNameItems struct {
	CharacterSetName []*string `json:"CharacterSetName,omitempty" xml:"CharacterSetName,omitempty" type:"Repeated"`
}

func (s DescribeCharacterSetNameResponseBodyCharacterSetNameItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeCharacterSetNameResponseBodyCharacterSetNameItems) GoString() string {
	return s.String()
}

func (s *DescribeCharacterSetNameResponseBodyCharacterSetNameItems) GetCharacterSetName() []*string {
	return s.CharacterSetName
}

func (s *DescribeCharacterSetNameResponseBodyCharacterSetNameItems) SetCharacterSetName(v []*string) *DescribeCharacterSetNameResponseBodyCharacterSetNameItems {
	s.CharacterSetName = v
	return s
}

func (s *DescribeCharacterSetNameResponseBodyCharacterSetNameItems) Validate() error {
	return dara.Validate(s)
}
