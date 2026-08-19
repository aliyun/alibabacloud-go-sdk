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
	// The unique identifier of the data development node.
	//
	// 	Notice: This field was of the Long type in SDK versions earlier than 8.0.0 and is of the String type in SDK 8.0.0 and later. **This change does not affect normal SDK usage, and the parameter is still returned in the type defined in the SDK**. Only when you upgrade across SDK version 8.0.0, the type change may cause project compilation failures, and you need to manually correct the data type.
	//
	// example:
	//
	// 860438872620113XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The request ID. You can use this ID to locate logs and troubleshoot issues.
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
