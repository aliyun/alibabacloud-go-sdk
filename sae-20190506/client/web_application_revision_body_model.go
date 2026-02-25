// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWebApplicationRevisionBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *WebApplicationRevisionBody
	GetCode() *int32
	SetData(v *Revision) *WebApplicationRevisionBody
	GetData() *Revision
	SetMessage(v string) *WebApplicationRevisionBody
	GetMessage() *string
	SetRequestId(v string) *WebApplicationRevisionBody
	GetRequestId() *string
	SetSuccess(v bool) *WebApplicationRevisionBody
	GetSuccess() *bool
}

type WebApplicationRevisionBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned for the operation.
	Data *Revision `json:"Data,omitempty" xml:"Data,omitempty"`
	// The message returned for the operation.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 91F93257-7A4A-4BD3-9A7E-2F6EAE6D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s WebApplicationRevisionBody) String() string {
	return dara.Prettify(s)
}

func (s WebApplicationRevisionBody) GoString() string {
	return s.String()
}

func (s *WebApplicationRevisionBody) GetCode() *int32 {
	return s.Code
}

func (s *WebApplicationRevisionBody) GetData() *Revision {
	return s.Data
}

func (s *WebApplicationRevisionBody) GetMessage() *string {
	return s.Message
}

func (s *WebApplicationRevisionBody) GetRequestId() *string {
	return s.RequestId
}

func (s *WebApplicationRevisionBody) GetSuccess() *bool {
	return s.Success
}

func (s *WebApplicationRevisionBody) SetCode(v int32) *WebApplicationRevisionBody {
	s.Code = &v
	return s
}

func (s *WebApplicationRevisionBody) SetData(v *Revision) *WebApplicationRevisionBody {
	s.Data = v
	return s
}

func (s *WebApplicationRevisionBody) SetMessage(v string) *WebApplicationRevisionBody {
	s.Message = &v
	return s
}

func (s *WebApplicationRevisionBody) SetRequestId(v string) *WebApplicationRevisionBody {
	s.RequestId = &v
	return s
}

func (s *WebApplicationRevisionBody) SetSuccess(v bool) *WebApplicationRevisionBody {
	s.Success = &v
	return s
}

func (s *WebApplicationRevisionBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
