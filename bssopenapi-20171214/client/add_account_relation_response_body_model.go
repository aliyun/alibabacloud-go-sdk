// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAccountRelationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddAccountRelationResponseBody
	GetCode() *string
	SetData(v *AddAccountRelationResponseBodyData) *AddAccountRelationResponseBody
	GetData() *AddAccountRelationResponseBodyData
	SetMessage(v string) *AddAccountRelationResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddAccountRelationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddAccountRelationResponseBody
	GetSuccess() *bool
}

type AddAccountRelationResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *AddAccountRelationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message returned.
	//
	// example:
	//
	// Message returned
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// Request ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddAccountRelationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddAccountRelationResponseBody) GoString() string {
	return s.String()
}

func (s *AddAccountRelationResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddAccountRelationResponseBody) GetData() *AddAccountRelationResponseBodyData {
	return s.Data
}

func (s *AddAccountRelationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddAccountRelationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddAccountRelationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddAccountRelationResponseBody) SetCode(v string) *AddAccountRelationResponseBody {
	s.Code = &v
	return s
}

func (s *AddAccountRelationResponseBody) SetData(v *AddAccountRelationResponseBodyData) *AddAccountRelationResponseBody {
	s.Data = v
	return s
}

func (s *AddAccountRelationResponseBody) SetMessage(v string) *AddAccountRelationResponseBody {
	s.Message = &v
	return s
}

func (s *AddAccountRelationResponseBody) SetRequestId(v string) *AddAccountRelationResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddAccountRelationResponseBody) SetSuccess(v bool) *AddAccountRelationResponseBody {
	s.Success = &v
	return s
}

func (s *AddAccountRelationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddAccountRelationResponseBodyData struct {
	// The IP address of the request
	//
	// example:
	//
	// HostId
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// The ID of the financial relationship.
	//
	// example:
	//
	// RelationId
	RelationId *int64 `json:"RelationId,omitempty" xml:"RelationId,omitempty"`
}

func (s AddAccountRelationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AddAccountRelationResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddAccountRelationResponseBodyData) GetHostId() *string {
	return s.HostId
}

func (s *AddAccountRelationResponseBodyData) GetRelationId() *int64 {
	return s.RelationId
}

func (s *AddAccountRelationResponseBodyData) SetHostId(v string) *AddAccountRelationResponseBodyData {
	s.HostId = &v
	return s
}

func (s *AddAccountRelationResponseBodyData) SetRelationId(v int64) *AddAccountRelationResponseBodyData {
	s.RelationId = &v
	return s
}

func (s *AddAccountRelationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
