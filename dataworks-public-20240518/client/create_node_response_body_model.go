// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *CreateNodeResponseBody
	GetId() *string
	SetRequestId(v string) *CreateNodeResponseBody
	GetRequestId() *string
}

type CreateNodeResponseBody struct {
	// The unique identifier of the Data Studio node.
	//
	// >  Prior to SDK version 8.0.0, this field is of type Long. In SDK version 8.0.0 and later, it is of type String. This change does not affect the normal use of the SDK. The parameter is returned based on the type defined in the SDK. Compilation failures caused by the type change may occur only when you upgrade the SDK across version 8.0.0. In this case, you must manually update the data type.
	//
	// example:
	//
	// 860438872620113XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
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

func (s *CreateNodeResponseBody) GetId() *string {
	return s.Id
}

func (s *CreateNodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateNodeResponseBody) SetId(v string) *CreateNodeResponseBody {
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
