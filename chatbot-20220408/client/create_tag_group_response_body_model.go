// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTagGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *CreateTagGroupResponseBody
	GetId() *int64
	SetRequestId(v string) *CreateTagGroupResponseBody
	GetRequestId() *string
}

type CreateTagGroupResponseBody struct {
	// example:
	//
	// 7393472
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// xxxXxxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateTagGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTagGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTagGroupResponseBody) GetId() *int64 {
	return s.Id
}

func (s *CreateTagGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTagGroupResponseBody) SetId(v int64) *CreateTagGroupResponseBody {
	s.Id = &v
	return s
}

func (s *CreateTagGroupResponseBody) SetRequestId(v string) *CreateTagGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTagGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
