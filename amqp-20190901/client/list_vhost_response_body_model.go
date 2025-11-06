// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVhostResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListVhostResponseBody
	GetCode() *int32
	SetData(v *ListVhostResponseBodyData) *ListVhostResponseBody
	GetData() *ListVhostResponseBodyData
	SetMessage(v string) *ListVhostResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListVhostResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListVhostResponseBody
	GetSuccess() *bool
}

type ListVhostResponseBody struct {
	Code      *int32                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListVhostResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListVhostResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVhostResponseBody) GoString() string {
	return s.String()
}

func (s *ListVhostResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListVhostResponseBody) GetData() *ListVhostResponseBodyData {
	return s.Data
}

func (s *ListVhostResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListVhostResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVhostResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListVhostResponseBody) SetCode(v int32) *ListVhostResponseBody {
	s.Code = &v
	return s
}

func (s *ListVhostResponseBody) SetData(v *ListVhostResponseBodyData) *ListVhostResponseBody {
	s.Data = v
	return s
}

func (s *ListVhostResponseBody) SetMessage(v string) *ListVhostResponseBody {
	s.Message = &v
	return s
}

func (s *ListVhostResponseBody) SetRequestId(v string) *ListVhostResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVhostResponseBody) SetSuccess(v bool) *ListVhostResponseBody {
	s.Success = &v
	return s
}

func (s *ListVhostResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListVhostResponseBodyData struct {
	Vhosts []*ListVhostResponseBodyDataVhosts `json:"Vhosts,omitempty" xml:"Vhosts,omitempty" type:"Repeated"`
}

func (s ListVhostResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListVhostResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListVhostResponseBodyData) GetVhosts() []*ListVhostResponseBodyDataVhosts {
	return s.Vhosts
}

func (s *ListVhostResponseBodyData) SetVhosts(v []*ListVhostResponseBodyDataVhosts) *ListVhostResponseBodyData {
	s.Vhosts = v
	return s
}

func (s *ListVhostResponseBodyData) Validate() error {
	if s.Vhosts != nil {
		for _, item := range s.Vhosts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListVhostResponseBodyDataVhosts struct {
	ChannelNum    *int32  `json:"ChannelNum,omitempty" xml:"ChannelNum,omitempty"`
	ConnectionNum *int32  `json:"ConnectionNum,omitempty" xml:"ConnectionNum,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListVhostResponseBodyDataVhosts) String() string {
	return dara.Prettify(s)
}

func (s ListVhostResponseBodyDataVhosts) GoString() string {
	return s.String()
}

func (s *ListVhostResponseBodyDataVhosts) GetChannelNum() *int32 {
	return s.ChannelNum
}

func (s *ListVhostResponseBodyDataVhosts) GetConnectionNum() *int32 {
	return s.ConnectionNum
}

func (s *ListVhostResponseBodyDataVhosts) GetName() *string {
	return s.Name
}

func (s *ListVhostResponseBodyDataVhosts) SetChannelNum(v int32) *ListVhostResponseBodyDataVhosts {
	s.ChannelNum = &v
	return s
}

func (s *ListVhostResponseBodyDataVhosts) SetConnectionNum(v int32) *ListVhostResponseBodyDataVhosts {
	s.ConnectionNum = &v
	return s
}

func (s *ListVhostResponseBodyDataVhosts) SetName(v string) *ListVhostResponseBodyDataVhosts {
	s.Name = &v
	return s
}

func (s *ListVhostResponseBodyDataVhosts) Validate() error {
	return dara.Validate(s)
}
