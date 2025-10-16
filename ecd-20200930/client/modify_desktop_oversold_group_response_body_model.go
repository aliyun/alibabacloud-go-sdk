// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDesktopOversoldGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ModifyDesktopOversoldGroupResponseBodyData) *ModifyDesktopOversoldGroupResponseBody
	GetData() *ModifyDesktopOversoldGroupResponseBodyData
	SetRequestId(v string) *ModifyDesktopOversoldGroupResponseBody
	GetRequestId() *string
}

type ModifyDesktopOversoldGroupResponseBody struct {
	Data      *ModifyDesktopOversoldGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDesktopOversoldGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDesktopOversoldGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDesktopOversoldGroupResponseBody) GetData() *ModifyDesktopOversoldGroupResponseBodyData {
	return s.Data
}

func (s *ModifyDesktopOversoldGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDesktopOversoldGroupResponseBody) SetData(v *ModifyDesktopOversoldGroupResponseBodyData) *ModifyDesktopOversoldGroupResponseBody {
	s.Data = v
	return s
}

func (s *ModifyDesktopOversoldGroupResponseBody) SetRequestId(v string) *ModifyDesktopOversoldGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDesktopOversoldGroupResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyDesktopOversoldGroupResponseBodyData struct {
	OversoldGroupId *string `json:"OversoldGroupId,omitempty" xml:"OversoldGroupId,omitempty"`
}

func (s ModifyDesktopOversoldGroupResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModifyDesktopOversoldGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifyDesktopOversoldGroupResponseBodyData) GetOversoldGroupId() *string {
	return s.OversoldGroupId
}

func (s *ModifyDesktopOversoldGroupResponseBodyData) SetOversoldGroupId(v string) *ModifyDesktopOversoldGroupResponseBodyData {
	s.OversoldGroupId = &v
	return s
}

func (s *ModifyDesktopOversoldGroupResponseBodyData) Validate() error {
	return dara.Validate(s)
}
