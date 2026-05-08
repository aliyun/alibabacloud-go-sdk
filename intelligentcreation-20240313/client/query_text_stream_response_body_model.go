// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTextStreamResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEnd(v bool) *QueryTextStreamResponseBody
	GetEnd() *bool
	SetIndex(v int32) *QueryTextStreamResponseBody
	GetIndex() *int32
	SetMessage(v string) *QueryTextStreamResponseBody
	GetMessage() *string
	SetType(v int32) *QueryTextStreamResponseBody
	GetType() *int32
}

type QueryTextStreamResponseBody struct {
	// example:
	//
	// false
	End *bool `json:"end,omitempty" xml:"end,omitempty"`
	// example:
	//
	// 1
	Index *int32 `json:"index,omitempty" xml:"index,omitempty"`
	// Id of the request
	//
	// example:
	//
	// None
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 1
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s QueryTextStreamResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryTextStreamResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTextStreamResponseBody) GetEnd() *bool {
	return s.End
}

func (s *QueryTextStreamResponseBody) GetIndex() *int32 {
	return s.Index
}

func (s *QueryTextStreamResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryTextStreamResponseBody) GetType() *int32 {
	return s.Type
}

func (s *QueryTextStreamResponseBody) SetEnd(v bool) *QueryTextStreamResponseBody {
	s.End = &v
	return s
}

func (s *QueryTextStreamResponseBody) SetIndex(v int32) *QueryTextStreamResponseBody {
	s.Index = &v
	return s
}

func (s *QueryTextStreamResponseBody) SetMessage(v string) *QueryTextStreamResponseBody {
	s.Message = &v
	return s
}

func (s *QueryTextStreamResponseBody) SetType(v int32) *QueryTextStreamResponseBody {
	s.Type = &v
	return s
}

func (s *QueryTextStreamResponseBody) Validate() error {
	return dara.Validate(s)
}
