// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRoutineCodeVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCodeDescription(v string) *GetRoutineCodeVersionResponseBody
	GetCodeDescription() *string
	SetCreateTime(v string) *GetRoutineCodeVersionResponseBody
	GetCreateTime() *string
	SetRequestId(v string) *GetRoutineCodeVersionResponseBody
	GetRequestId() *string
	SetRoutineCode(v string) *GetRoutineCodeVersionResponseBody
	GetRoutineCode() *string
}

type GetRoutineCodeVersionResponseBody struct {
	CodeDescription *string `json:"CodeDescription,omitempty" xml:"CodeDescription,omitempty"`
	CreateTime      *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Id of the request
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RoutineCode *string `json:"RoutineCode,omitempty" xml:"RoutineCode,omitempty"`
}

func (s GetRoutineCodeVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRoutineCodeVersionResponseBody) GoString() string {
	return s.String()
}

func (s *GetRoutineCodeVersionResponseBody) GetCodeDescription() *string {
	return s.CodeDescription
}

func (s *GetRoutineCodeVersionResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetRoutineCodeVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRoutineCodeVersionResponseBody) GetRoutineCode() *string {
	return s.RoutineCode
}

func (s *GetRoutineCodeVersionResponseBody) SetCodeDescription(v string) *GetRoutineCodeVersionResponseBody {
	s.CodeDescription = &v
	return s
}

func (s *GetRoutineCodeVersionResponseBody) SetCreateTime(v string) *GetRoutineCodeVersionResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetRoutineCodeVersionResponseBody) SetRequestId(v string) *GetRoutineCodeVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRoutineCodeVersionResponseBody) SetRoutineCode(v string) *GetRoutineCodeVersionResponseBody {
	s.RoutineCode = &v
	return s
}

func (s *GetRoutineCodeVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
