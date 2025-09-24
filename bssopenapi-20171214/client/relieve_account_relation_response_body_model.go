// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRelieveAccountRelationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RelieveAccountRelationResponseBody
	GetCode() *string
	SetData(v *RelieveAccountRelationResponseBodyData) *RelieveAccountRelationResponseBody
	GetData() *RelieveAccountRelationResponseBodyData
	SetMessage(v string) *RelieveAccountRelationResponseBody
	GetMessage() *string
	SetRequestId(v string) *RelieveAccountRelationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RelieveAccountRelationResponseBody
	GetSuccess() *bool
}

type RelieveAccountRelationResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *RelieveAccountRelationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message returned.
	//
	// example:
	//
	// Message returned
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique ID of the request.
	//
	// example:
	//
	// request_id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RelieveAccountRelationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RelieveAccountRelationResponseBody) GoString() string {
	return s.String()
}

func (s *RelieveAccountRelationResponseBody) GetCode() *string {
	return s.Code
}

func (s *RelieveAccountRelationResponseBody) GetData() *RelieveAccountRelationResponseBodyData {
	return s.Data
}

func (s *RelieveAccountRelationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RelieveAccountRelationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RelieveAccountRelationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RelieveAccountRelationResponseBody) SetCode(v string) *RelieveAccountRelationResponseBody {
	s.Code = &v
	return s
}

func (s *RelieveAccountRelationResponseBody) SetData(v *RelieveAccountRelationResponseBodyData) *RelieveAccountRelationResponseBody {
	s.Data = v
	return s
}

func (s *RelieveAccountRelationResponseBody) SetMessage(v string) *RelieveAccountRelationResponseBody {
	s.Message = &v
	return s
}

func (s *RelieveAccountRelationResponseBody) SetRequestId(v string) *RelieveAccountRelationResponseBody {
	s.RequestId = &v
	return s
}

func (s *RelieveAccountRelationResponseBody) SetSuccess(v bool) *RelieveAccountRelationResponseBody {
	s.Success = &v
	return s
}

func (s *RelieveAccountRelationResponseBody) Validate() error {
	return dara.Validate(s)
}

type RelieveAccountRelationResponseBodyData struct {
	// hostid
	//
	// example:
	//
	// HostId
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
}

func (s RelieveAccountRelationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RelieveAccountRelationResponseBodyData) GoString() string {
	return s.String()
}

func (s *RelieveAccountRelationResponseBodyData) GetHostId() *string {
	return s.HostId
}

func (s *RelieveAccountRelationResponseBodyData) SetHostId(v string) *RelieveAccountRelationResponseBodyData {
	s.HostId = &v
	return s
}

func (s *RelieveAccountRelationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
