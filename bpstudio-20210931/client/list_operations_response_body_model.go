// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOperationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListOperationsResponseBody
	GetCode() *string
	SetData(v []*ListOperationsResponseBodyData) *ListOperationsResponseBody
	GetData() []*ListOperationsResponseBodyData
	SetMessage(v string) *ListOperationsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListOperationsResponseBody
	GetRequestId() *string
}

type ListOperationsResponseBody struct {
	// example:
	//
	// 200
	Code *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*ListOperationsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// NotFound
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// A07FFDF2-78FA-1B48-9E38-88E833A93187
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListOperationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListOperationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListOperationsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListOperationsResponseBody) GetData() []*ListOperationsResponseBodyData {
	return s.Data
}

func (s *ListOperationsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListOperationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListOperationsResponseBody) SetCode(v string) *ListOperationsResponseBody {
	s.Code = &v
	return s
}

func (s *ListOperationsResponseBody) SetData(v []*ListOperationsResponseBodyData) *ListOperationsResponseBody {
	s.Data = v
	return s
}

func (s *ListOperationsResponseBody) SetMessage(v string) *ListOperationsResponseBody {
	s.Message = &v
	return s
}

func (s *ListOperationsResponseBody) SetRequestId(v string) *ListOperationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOperationsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListOperationsResponseBodyData struct {
	// example:
	//
	// Add Tag
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// addTags
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// example:
	//
	// ecs
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
}

func (s ListOperationsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListOperationsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListOperationsResponseBodyData) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListOperationsResponseBodyData) GetOperation() *string {
	return s.Operation
}

func (s *ListOperationsResponseBodyData) GetService() *string {
	return s.Service
}

func (s *ListOperationsResponseBodyData) SetDisplayName(v string) *ListOperationsResponseBodyData {
	s.DisplayName = &v
	return s
}

func (s *ListOperationsResponseBodyData) SetOperation(v string) *ListOperationsResponseBodyData {
	s.Operation = &v
	return s
}

func (s *ListOperationsResponseBodyData) SetService(v string) *ListOperationsResponseBodyData {
	s.Service = &v
	return s
}

func (s *ListOperationsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
