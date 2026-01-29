// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfirmRelationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ConfirmRelationResponseBody
	GetCode() *string
	SetData(v *ConfirmRelationResponseBodyData) *ConfirmRelationResponseBody
	GetData() *ConfirmRelationResponseBodyData
	SetMessage(v string) *ConfirmRelationResponseBody
	GetMessage() *string
	SetRequestId(v string) *ConfirmRelationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ConfirmRelationResponseBody
	GetSuccess() *bool
}

type ConfirmRelationResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *ConfirmRelationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s ConfirmRelationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfirmRelationResponseBody) GoString() string {
	return s.String()
}

func (s *ConfirmRelationResponseBody) GetCode() *string {
	return s.Code
}

func (s *ConfirmRelationResponseBody) GetData() *ConfirmRelationResponseBodyData {
	return s.Data
}

func (s *ConfirmRelationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ConfirmRelationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfirmRelationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ConfirmRelationResponseBody) SetCode(v string) *ConfirmRelationResponseBody {
	s.Code = &v
	return s
}

func (s *ConfirmRelationResponseBody) SetData(v *ConfirmRelationResponseBodyData) *ConfirmRelationResponseBody {
	s.Data = v
	return s
}

func (s *ConfirmRelationResponseBody) SetMessage(v string) *ConfirmRelationResponseBody {
	s.Message = &v
	return s
}

func (s *ConfirmRelationResponseBody) SetRequestId(v string) *ConfirmRelationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfirmRelationResponseBody) SetSuccess(v bool) *ConfirmRelationResponseBody {
	s.Success = &v
	return s
}

func (s *ConfirmRelationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ConfirmRelationResponseBodyData struct {
	// HostId
	//
	// example:
	//
	// HostId
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
}

func (s ConfirmRelationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ConfirmRelationResponseBodyData) GoString() string {
	return s.String()
}

func (s *ConfirmRelationResponseBodyData) GetHostId() *string {
	return s.HostId
}

func (s *ConfirmRelationResponseBodyData) SetHostId(v string) *ConfirmRelationResponseBodyData {
	s.HostId = &v
	return s
}

func (s *ConfirmRelationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
