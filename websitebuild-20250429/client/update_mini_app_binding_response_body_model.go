// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMiniAppBindingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *UpdateMiniAppBindingResponseBodyData) *UpdateMiniAppBindingResponseBody
	GetData() *UpdateMiniAppBindingResponseBodyData
	SetRequestId(v string) *UpdateMiniAppBindingResponseBody
	GetRequestId() *string
}

type UpdateMiniAppBindingResponseBody struct {
	Data *UpdateMiniAppBindingResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 6C6B99AC-39EC-5350-874C-204128C905E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateMiniAppBindingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMiniAppBindingResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMiniAppBindingResponseBody) GetData() *UpdateMiniAppBindingResponseBodyData {
	return s.Data
}

func (s *UpdateMiniAppBindingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMiniAppBindingResponseBody) SetData(v *UpdateMiniAppBindingResponseBodyData) *UpdateMiniAppBindingResponseBody {
	s.Data = v
	return s
}

func (s *UpdateMiniAppBindingResponseBody) SetRequestId(v string) *UpdateMiniAppBindingResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMiniAppBindingResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateMiniAppBindingResponseBodyData struct {
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateMiniAppBindingResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateMiniAppBindingResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateMiniAppBindingResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateMiniAppBindingResponseBodyData) SetSuccess(v bool) *UpdateMiniAppBindingResponseBodyData {
	s.Success = &v
	return s
}

func (s *UpdateMiniAppBindingResponseBodyData) Validate() error {
	return dara.Validate(s)
}
