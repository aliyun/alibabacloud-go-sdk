// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryOneKeyDeleteBlockListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryOneKeyDeleteBlockListResponseBody
	GetCode() *string
	SetData(v []*QueryOneKeyDeleteBlockListResponseBodyData) *QueryOneKeyDeleteBlockListResponseBody
	GetData() []*QueryOneKeyDeleteBlockListResponseBodyData
	SetMessage(v string) *QueryOneKeyDeleteBlockListResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryOneKeyDeleteBlockListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryOneKeyDeleteBlockListResponseBody
	GetSuccess() *bool
}

type QueryOneKeyDeleteBlockListResponseBody struct {
	Code      *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*QueryOneKeyDeleteBlockListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryOneKeyDeleteBlockListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryOneKeyDeleteBlockListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryOneKeyDeleteBlockListResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryOneKeyDeleteBlockListResponseBody) GetData() []*QueryOneKeyDeleteBlockListResponseBodyData {
	return s.Data
}

func (s *QueryOneKeyDeleteBlockListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryOneKeyDeleteBlockListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryOneKeyDeleteBlockListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryOneKeyDeleteBlockListResponseBody) SetCode(v string) *QueryOneKeyDeleteBlockListResponseBody {
	s.Code = &v
	return s
}

func (s *QueryOneKeyDeleteBlockListResponseBody) SetData(v []*QueryOneKeyDeleteBlockListResponseBodyData) *QueryOneKeyDeleteBlockListResponseBody {
	s.Data = v
	return s
}

func (s *QueryOneKeyDeleteBlockListResponseBody) SetMessage(v string) *QueryOneKeyDeleteBlockListResponseBody {
	s.Message = &v
	return s
}

func (s *QueryOneKeyDeleteBlockListResponseBody) SetRequestId(v string) *QueryOneKeyDeleteBlockListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryOneKeyDeleteBlockListResponseBody) SetSuccess(v bool) *QueryOneKeyDeleteBlockListResponseBody {
	s.Success = &v
	return s
}

func (s *QueryOneKeyDeleteBlockListResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryOneKeyDeleteBlockListResponseBodyData struct {
	BlockCode *string `json:"BlockCode,omitempty" xml:"BlockCode,omitempty"`
	BlockMsg  *string `json:"BlockMsg,omitempty" xml:"BlockMsg,omitempty"`
}

func (s QueryOneKeyDeleteBlockListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryOneKeyDeleteBlockListResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryOneKeyDeleteBlockListResponseBodyData) GetBlockCode() *string {
	return s.BlockCode
}

func (s *QueryOneKeyDeleteBlockListResponseBodyData) GetBlockMsg() *string {
	return s.BlockMsg
}

func (s *QueryOneKeyDeleteBlockListResponseBodyData) SetBlockCode(v string) *QueryOneKeyDeleteBlockListResponseBodyData {
	s.BlockCode = &v
	return s
}

func (s *QueryOneKeyDeleteBlockListResponseBodyData) SetBlockMsg(v string) *QueryOneKeyDeleteBlockListResponseBodyData {
	s.BlockMsg = &v
	return s
}

func (s *QueryOneKeyDeleteBlockListResponseBodyData) Validate() error {
	return dara.Validate(s)
}
