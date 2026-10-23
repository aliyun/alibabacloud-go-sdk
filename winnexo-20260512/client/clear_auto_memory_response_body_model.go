// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClearAutoMemoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClearedCount(v int64) *ClearAutoMemoryResponseBody
	GetClearedCount() *int64
	SetCode(v string) *ClearAutoMemoryResponseBody
	GetCode() *string
	SetMessage(v string) *ClearAutoMemoryResponseBody
	GetMessage() *string
	SetRequestId(v string) *ClearAutoMemoryResponseBody
	GetRequestId() *string
	SetType(v string) *ClearAutoMemoryResponseBody
	GetType() *string
}

type ClearAutoMemoryResponseBody struct {
	// The number of automatic memory files successfully deleted in this request. The value is 0 if no memory is available to be cleared.
	//
	// example:
	//
	// 3
	ClearedCount *int64 `json:"clearedCount,omitempty" xml:"clearedCount,omitempty"`
	// The business status code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The error description.
	//
	// example:
	//
	// The current user is not an active member of the group cg_12345
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request tracking ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The type of automatic memory cleared in this request. Valid values: user, feedback, reference, project, and all.
	//
	// example:
	//
	// project
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ClearAutoMemoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ClearAutoMemoryResponseBody) GoString() string {
	return s.String()
}

func (s *ClearAutoMemoryResponseBody) GetClearedCount() *int64 {
	return s.ClearedCount
}

func (s *ClearAutoMemoryResponseBody) GetCode() *string {
	return s.Code
}

func (s *ClearAutoMemoryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ClearAutoMemoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ClearAutoMemoryResponseBody) GetType() *string {
	return s.Type
}

func (s *ClearAutoMemoryResponseBody) SetClearedCount(v int64) *ClearAutoMemoryResponseBody {
	s.ClearedCount = &v
	return s
}

func (s *ClearAutoMemoryResponseBody) SetCode(v string) *ClearAutoMemoryResponseBody {
	s.Code = &v
	return s
}

func (s *ClearAutoMemoryResponseBody) SetMessage(v string) *ClearAutoMemoryResponseBody {
	s.Message = &v
	return s
}

func (s *ClearAutoMemoryResponseBody) SetRequestId(v string) *ClearAutoMemoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ClearAutoMemoryResponseBody) SetType(v string) *ClearAutoMemoryResponseBody {
	s.Type = &v
	return s
}

func (s *ClearAutoMemoryResponseBody) Validate() error {
	return dara.Validate(s)
}
