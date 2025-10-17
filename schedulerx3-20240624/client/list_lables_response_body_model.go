// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLablesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListLablesResponseBody
	GetCode() *int32
	SetData(v []*ListLablesResponseBodyData) *ListLablesResponseBody
	GetData() []*ListLablesResponseBodyData
	SetMessage(v string) *ListLablesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListLablesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListLablesResponseBody
	GetSuccess() *bool
}

type ListLablesResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// -
	Data []*ListLablesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// Parameter check error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9079A828-9138-50F1-801E-F2BC3D222A06
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListLablesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLablesResponseBody) GoString() string {
	return s.String()
}

func (s *ListLablesResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListLablesResponseBody) GetData() []*ListLablesResponseBodyData {
	return s.Data
}

func (s *ListLablesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListLablesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLablesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListLablesResponseBody) SetCode(v int32) *ListLablesResponseBody {
	s.Code = &v
	return s
}

func (s *ListLablesResponseBody) SetData(v []*ListLablesResponseBodyData) *ListLablesResponseBody {
	s.Data = v
	return s
}

func (s *ListLablesResponseBody) SetMessage(v string) *ListLablesResponseBody {
	s.Message = &v
	return s
}

func (s *ListLablesResponseBody) SetRequestId(v string) *ListLablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLablesResponseBody) SetSuccess(v bool) *ListLablesResponseBody {
	s.Success = &v
	return s
}

func (s *ListLablesResponseBody) Validate() error {
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

type ListLablesResponseBodyData struct {
	// example:
	//
	// true
	IsDesignated *bool `json:"IsDesignated,omitempty" xml:"IsDesignated,omitempty"`
	// example:
	//
	// gray
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// example:
	//
	// true
	Online *bool `json:"Online,omitempty" xml:"Online,omitempty"`
	// example:
	//
	// 2
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s ListLablesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListLablesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListLablesResponseBodyData) GetIsDesignated() *bool {
	return s.IsDesignated
}

func (s *ListLablesResponseBodyData) GetLabel() *string {
	return s.Label
}

func (s *ListLablesResponseBodyData) GetOnline() *bool {
	return s.Online
}

func (s *ListLablesResponseBodyData) GetSize() *int32 {
	return s.Size
}

func (s *ListLablesResponseBodyData) SetIsDesignated(v bool) *ListLablesResponseBodyData {
	s.IsDesignated = &v
	return s
}

func (s *ListLablesResponseBodyData) SetLabel(v string) *ListLablesResponseBodyData {
	s.Label = &v
	return s
}

func (s *ListLablesResponseBodyData) SetOnline(v bool) *ListLablesResponseBodyData {
	s.Online = &v
	return s
}

func (s *ListLablesResponseBodyData) SetSize(v int32) *ListLablesResponseBodyData {
	s.Size = &v
	return s
}

func (s *ListLablesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
