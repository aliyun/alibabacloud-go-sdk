// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *CreateNodeResponseBody
	GetId() *int64
	SetRequestId(v string) *CreateNodeResponseBody
	GetRequestId() *string
}

type CreateNodeResponseBody struct {
	// The ID of the node.
	//
	// example:
	//
	// 860438872620113XXXX
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The request ID.
	//
	// example:
	//
	// AFBB799F-8578-51C5-A766-E922EDB8XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateNodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateNodeResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNodeResponseBody) GetId() *int64 {
	return s.Id
}

func (s *CreateNodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateNodeResponseBody) SetId(v int64) *CreateNodeResponseBody {
	s.Id = &v
	return s
}

func (s *CreateNodeResponseBody) SetRequestId(v string) *CreateNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateNodeResponseBody) Validate() error {
	return dara.Validate(s)
}
