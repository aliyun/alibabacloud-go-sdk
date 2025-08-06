// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBizUserTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribeBizUserTypeResponseBodyData) *DescribeBizUserTypeResponseBody
	GetData() []*DescribeBizUserTypeResponseBodyData
	SetRequestId(v string) *DescribeBizUserTypeResponseBody
	GetRequestId() *string
}

type DescribeBizUserTypeResponseBody struct {
	Data      []*DescribeBizUserTypeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeBizUserTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBizUserTypeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBizUserTypeResponseBody) GetData() []*DescribeBizUserTypeResponseBodyData {
	return s.Data
}

func (s *DescribeBizUserTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBizUserTypeResponseBody) SetData(v []*DescribeBizUserTypeResponseBodyData) *DescribeBizUserTypeResponseBody {
	s.Data = v
	return s
}

func (s *DescribeBizUserTypeResponseBody) SetRequestId(v string) *DescribeBizUserTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBizUserTypeResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeBizUserTypeResponseBodyData struct {
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeBizUserTypeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeBizUserTypeResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeBizUserTypeResponseBodyData) GetType() *string {
	return s.Type
}

func (s *DescribeBizUserTypeResponseBodyData) SetType(v string) *DescribeBizUserTypeResponseBodyData {
	s.Type = &v
	return s
}

func (s *DescribeBizUserTypeResponseBodyData) Validate() error {
	return dara.Validate(s)
}
