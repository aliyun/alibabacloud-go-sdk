// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLumaNamespaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetLumaNamespaceResponseBody
	GetCode() *string
	SetData(v *Namespace) *GetLumaNamespaceResponseBody
	GetData() *Namespace
	SetMessage(v string) *GetLumaNamespaceResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetLumaNamespaceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetLumaNamespaceResponseBody
	GetSuccess() *bool
}

type GetLumaNamespaceResponseBody struct {
	// The response code. A value of Success indicates that the call succeeded. If the call fails, a specific error code is returned.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the namespace bound to the Agent.
	Data *Namespace `json:"Data,omitempty" xml:"Data,omitempty"`
	// The message returned by the operation. The value is Operation success if the call succeeds, or a specific error description if the call fails.
	//
	// example:
	//
	// Operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique identifier of the request, used for troubleshooting and ticket submission.
	//
	// example:
	//
	// 34AD682D-5B91-5773-8132-AA38C130****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. A value of true indicates success.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetLumaNamespaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLumaNamespaceResponseBody) GoString() string {
	return s.String()
}

func (s *GetLumaNamespaceResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetLumaNamespaceResponseBody) GetData() *Namespace {
	return s.Data
}

func (s *GetLumaNamespaceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetLumaNamespaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLumaNamespaceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetLumaNamespaceResponseBody) SetCode(v string) *GetLumaNamespaceResponseBody {
	s.Code = &v
	return s
}

func (s *GetLumaNamespaceResponseBody) SetData(v *Namespace) *GetLumaNamespaceResponseBody {
	s.Data = v
	return s
}

func (s *GetLumaNamespaceResponseBody) SetMessage(v string) *GetLumaNamespaceResponseBody {
	s.Message = &v
	return s
}

func (s *GetLumaNamespaceResponseBody) SetRequestId(v string) *GetLumaNamespaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLumaNamespaceResponseBody) SetSuccess(v bool) *GetLumaNamespaceResponseBody {
	s.Success = &v
	return s
}

func (s *GetLumaNamespaceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
