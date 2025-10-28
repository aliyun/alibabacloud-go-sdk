// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMultimodalAsyncModerationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *MultimodalAsyncModerationResponseBody
	GetCode() *int32
	SetData(v *MultimodalAsyncModerationResponseBodyData) *MultimodalAsyncModerationResponseBody
	GetData() *MultimodalAsyncModerationResponseBodyData
	SetMsg(v string) *MultimodalAsyncModerationResponseBody
	GetMsg() *string
	SetRequestId(v string) *MultimodalAsyncModerationResponseBody
	GetRequestId() *string
}

type MultimodalAsyncModerationResponseBody struct {
	Code      *int32                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *MultimodalAsyncModerationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Msg       *string                                    `json:"Msg,omitempty" xml:"Msg,omitempty"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MultimodalAsyncModerationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MultimodalAsyncModerationResponseBody) GoString() string {
	return s.String()
}

func (s *MultimodalAsyncModerationResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *MultimodalAsyncModerationResponseBody) GetData() *MultimodalAsyncModerationResponseBodyData {
	return s.Data
}

func (s *MultimodalAsyncModerationResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *MultimodalAsyncModerationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MultimodalAsyncModerationResponseBody) SetCode(v int32) *MultimodalAsyncModerationResponseBody {
	s.Code = &v
	return s
}

func (s *MultimodalAsyncModerationResponseBody) SetData(v *MultimodalAsyncModerationResponseBodyData) *MultimodalAsyncModerationResponseBody {
	s.Data = v
	return s
}

func (s *MultimodalAsyncModerationResponseBody) SetMsg(v string) *MultimodalAsyncModerationResponseBody {
	s.Msg = &v
	return s
}

func (s *MultimodalAsyncModerationResponseBody) SetRequestId(v string) *MultimodalAsyncModerationResponseBody {
	s.RequestId = &v
	return s
}

func (s *MultimodalAsyncModerationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MultimodalAsyncModerationResponseBodyData struct {
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	ReqId  *string `json:"ReqId,omitempty" xml:"ReqId,omitempty"`
}

func (s MultimodalAsyncModerationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s MultimodalAsyncModerationResponseBodyData) GoString() string {
	return s.String()
}

func (s *MultimodalAsyncModerationResponseBodyData) GetDataId() *string {
	return s.DataId
}

func (s *MultimodalAsyncModerationResponseBodyData) GetReqId() *string {
	return s.ReqId
}

func (s *MultimodalAsyncModerationResponseBodyData) SetDataId(v string) *MultimodalAsyncModerationResponseBodyData {
	s.DataId = &v
	return s
}

func (s *MultimodalAsyncModerationResponseBodyData) SetReqId(v string) *MultimodalAsyncModerationResponseBodyData {
	s.ReqId = &v
	return s
}

func (s *MultimodalAsyncModerationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
