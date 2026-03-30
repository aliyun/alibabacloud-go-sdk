// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCloneVoiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteCloneVoiceResponseBody
	GetCode() *string
	SetData(v string) *DeleteCloneVoiceResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *DeleteCloneVoiceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteCloneVoiceResponseBody
	GetMessage() *string
	SetParams(v []*string) *DeleteCloneVoiceResponseBody
	GetParams() []*string
	SetRequestId(v string) *DeleteCloneVoiceResponseBody
	GetRequestId() *string
}

type DeleteCloneVoiceResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// e48e45dd-e47a-4744-a063-f08cbebb1c5b
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Instance af81a389-91f0-4157-8d82-720edd02b66a
	//
	//  does not exist.
	Message *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params  []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// 14C39896-AE6D-4643-9C9A-E0566B2C2DDD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCloneVoiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCloneVoiceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCloneVoiceResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteCloneVoiceResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteCloneVoiceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteCloneVoiceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteCloneVoiceResponseBody) GetParams() []*string {
	return s.Params
}

func (s *DeleteCloneVoiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCloneVoiceResponseBody) SetCode(v string) *DeleteCloneVoiceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteCloneVoiceResponseBody) SetData(v string) *DeleteCloneVoiceResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteCloneVoiceResponseBody) SetHttpStatusCode(v int32) *DeleteCloneVoiceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteCloneVoiceResponseBody) SetMessage(v string) *DeleteCloneVoiceResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteCloneVoiceResponseBody) SetParams(v []*string) *DeleteCloneVoiceResponseBody {
	s.Params = v
	return s
}

func (s *DeleteCloneVoiceResponseBody) SetRequestId(v string) *DeleteCloneVoiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCloneVoiceResponseBody) Validate() error {
	return dara.Validate(s)
}
