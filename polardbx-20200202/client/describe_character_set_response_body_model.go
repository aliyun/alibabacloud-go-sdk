// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCharacterSetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeCharacterSetResponseBodyData) *DescribeCharacterSetResponseBody
	GetData() *DescribeCharacterSetResponseBodyData
	SetMessage(v string) *DescribeCharacterSetResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeCharacterSetResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeCharacterSetResponseBody
	GetSuccess() *bool
}

type DescribeCharacterSetResponseBody struct {
	Data *DescribeCharacterSetResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 709C1E40-092D-4A3D-9958-6D7438******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeCharacterSetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCharacterSetResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCharacterSetResponseBody) GetData() *DescribeCharacterSetResponseBodyData {
	return s.Data
}

func (s *DescribeCharacterSetResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeCharacterSetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCharacterSetResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeCharacterSetResponseBody) SetData(v *DescribeCharacterSetResponseBodyData) *DescribeCharacterSetResponseBody {
	s.Data = v
	return s
}

func (s *DescribeCharacterSetResponseBody) SetMessage(v string) *DescribeCharacterSetResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeCharacterSetResponseBody) SetRequestId(v string) *DescribeCharacterSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCharacterSetResponseBody) SetSuccess(v bool) *DescribeCharacterSetResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeCharacterSetResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCharacterSetResponseBodyData struct {
	CharacterSet []*string `json:"CharacterSet,omitempty" xml:"CharacterSet,omitempty" type:"Repeated"`
	// example:
	//
	// polarx
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
}

func (s DescribeCharacterSetResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeCharacterSetResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeCharacterSetResponseBodyData) GetCharacterSet() []*string {
	return s.CharacterSet
}

func (s *DescribeCharacterSetResponseBodyData) GetEngine() *string {
	return s.Engine
}

func (s *DescribeCharacterSetResponseBodyData) SetCharacterSet(v []*string) *DescribeCharacterSetResponseBodyData {
	s.CharacterSet = v
	return s
}

func (s *DescribeCharacterSetResponseBodyData) SetEngine(v string) *DescribeCharacterSetResponseBodyData {
	s.Engine = &v
	return s
}

func (s *DescribeCharacterSetResponseBodyData) Validate() error {
	return dara.Validate(s)
}
