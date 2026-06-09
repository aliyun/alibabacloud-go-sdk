// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNamespaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetNamespaceResponseBody
	GetCode() *string
	SetData(v *Namespace) *GetNamespaceResponseBody
	GetData() *Namespace
	SetMessage(v string) *GetNamespaceResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetNamespaceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetNamespaceResponseBody
	GetSuccess() *bool
}

type GetNamespaceResponseBody struct {
	// example:
	//
	// 200
	Code *string    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *Namespace `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// Operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 34AD682D-5B91-5773-8132-AA38C130****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetNamespaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetNamespaceResponseBody) GoString() string {
	return s.String()
}

func (s *GetNamespaceResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetNamespaceResponseBody) GetData() *Namespace {
	return s.Data
}

func (s *GetNamespaceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetNamespaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetNamespaceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetNamespaceResponseBody) SetCode(v string) *GetNamespaceResponseBody {
	s.Code = &v
	return s
}

func (s *GetNamespaceResponseBody) SetData(v *Namespace) *GetNamespaceResponseBody {
	s.Data = v
	return s
}

func (s *GetNamespaceResponseBody) SetMessage(v string) *GetNamespaceResponseBody {
	s.Message = &v
	return s
}

func (s *GetNamespaceResponseBody) SetRequestId(v string) *GetNamespaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNamespaceResponseBody) SetSuccess(v bool) *GetNamespaceResponseBody {
	s.Success = &v
	return s
}

func (s *GetNamespaceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
