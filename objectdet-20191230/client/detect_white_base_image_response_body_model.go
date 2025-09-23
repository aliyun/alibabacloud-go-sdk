// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectWhiteBaseImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DetectWhiteBaseImageResponseBodyData) *DetectWhiteBaseImageResponseBody
	GetData() *DetectWhiteBaseImageResponseBodyData
	SetRequestId(v string) *DetectWhiteBaseImageResponseBody
	GetRequestId() *string
}

type DetectWhiteBaseImageResponseBody struct {
	Data *DetectWhiteBaseImageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 7A7F9EEB-44C4-4592-9089-A6185B222B23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetectWhiteBaseImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetectWhiteBaseImageResponseBody) GoString() string {
	return s.String()
}

func (s *DetectWhiteBaseImageResponseBody) GetData() *DetectWhiteBaseImageResponseBodyData {
	return s.Data
}

func (s *DetectWhiteBaseImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetectWhiteBaseImageResponseBody) SetData(v *DetectWhiteBaseImageResponseBodyData) *DetectWhiteBaseImageResponseBody {
	s.Data = v
	return s
}

func (s *DetectWhiteBaseImageResponseBody) SetRequestId(v string) *DetectWhiteBaseImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetectWhiteBaseImageResponseBody) Validate() error {
	return dara.Validate(s)
}

type DetectWhiteBaseImageResponseBodyData struct {
	Elements []*DetectWhiteBaseImageResponseBodyDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
}

func (s DetectWhiteBaseImageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DetectWhiteBaseImageResponseBodyData) GoString() string {
	return s.String()
}

func (s *DetectWhiteBaseImageResponseBodyData) GetElements() []*DetectWhiteBaseImageResponseBodyDataElements {
	return s.Elements
}

func (s *DetectWhiteBaseImageResponseBodyData) SetElements(v []*DetectWhiteBaseImageResponseBodyDataElements) *DetectWhiteBaseImageResponseBodyData {
	s.Elements = v
	return s
}

func (s *DetectWhiteBaseImageResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DetectWhiteBaseImageResponseBodyDataElements struct {
	// example:
	//
	// 0
	WhiteBase *int32 `json:"WhiteBase,omitempty" xml:"WhiteBase,omitempty"`
}

func (s DetectWhiteBaseImageResponseBodyDataElements) String() string {
	return dara.Prettify(s)
}

func (s DetectWhiteBaseImageResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *DetectWhiteBaseImageResponseBodyDataElements) GetWhiteBase() *int32 {
	return s.WhiteBase
}

func (s *DetectWhiteBaseImageResponseBodyDataElements) SetWhiteBase(v int32) *DetectWhiteBaseImageResponseBodyDataElements {
	s.WhiteBase = &v
	return s
}

func (s *DetectWhiteBaseImageResponseBodyDataElements) Validate() error {
	return dara.Validate(s)
}
