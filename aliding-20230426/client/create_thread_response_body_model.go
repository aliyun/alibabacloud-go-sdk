// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateThreadResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateThreadResponseBody
	GetRequestId() *string
	SetThread(v *CreateThreadResponseBodyThread) *CreateThreadResponseBody
	GetThread() *CreateThreadResponseBodyThread
}

type CreateThreadResponseBody struct {
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string                         `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Thread    *CreateThreadResponseBodyThread `json:"thread,omitempty" xml:"thread,omitempty" type:"Struct"`
}

func (s CreateThreadResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateThreadResponseBody) GoString() string {
	return s.String()
}

func (s *CreateThreadResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateThreadResponseBody) GetThread() *CreateThreadResponseBodyThread {
	return s.Thread
}

func (s *CreateThreadResponseBody) SetRequestId(v string) *CreateThreadResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateThreadResponseBody) SetThread(v *CreateThreadResponseBodyThread) *CreateThreadResponseBody {
	s.Thread = v
	return s
}

func (s *CreateThreadResponseBody) Validate() error {
	if s.Thread != nil {
		if err := s.Thread.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateThreadResponseBodyThread struct {
	CreateAt *int64  `json:"createAt,omitempty" xml:"createAt,omitempty"`
	Id       *string `json:"id,omitempty" xml:"id,omitempty"`
}

func (s CreateThreadResponseBodyThread) String() string {
	return dara.Prettify(s)
}

func (s CreateThreadResponseBodyThread) GoString() string {
	return s.String()
}

func (s *CreateThreadResponseBodyThread) GetCreateAt() *int64 {
	return s.CreateAt
}

func (s *CreateThreadResponseBodyThread) GetId() *string {
	return s.Id
}

func (s *CreateThreadResponseBodyThread) SetCreateAt(v int64) *CreateThreadResponseBodyThread {
	s.CreateAt = &v
	return s
}

func (s *CreateThreadResponseBodyThread) SetId(v string) *CreateThreadResponseBodyThread {
	s.Id = &v
	return s
}

func (s *CreateThreadResponseBodyThread) Validate() error {
	return dara.Validate(s)
}
